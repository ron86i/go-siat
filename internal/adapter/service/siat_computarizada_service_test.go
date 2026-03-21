package service_test

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/facturas"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// TestSiatComputarizadaService_ReversionAnulacionFactura valida el flujo de reversión de una
// anulación de factura previamente realizada.
func TestSiatComputarizadaService_ReversionAnulacionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()
	// 1. Obtener CUIS
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo
	log.Printf("CUIS: %s", cuis)

	// 2. Obtener CUFD
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()

	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
	log.Printf("CUFD: %s", cufd)

	cuf := "D5340CCDF031F2596FC03311F6F76AB5334D0A86A626F497FCE6AAF74"

	req := models.Computarizada().NewReversionAnulacionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCodigoDocumentoSector(10).
		WithTipoFacturaDocumento(2).
		WithCodigoEmision(1).
		WithCuf(cuf).
		Build()
	resp, err := serviceComputarizada.ReversionAnulacionFactura(context.Background(), config, req)
	assert.NoError(t, err)         // Verifica que no haya errores de red o de parseo
	assert.Nil(t, resp.Body.Fault) // Verifica que no haya errores en la respuesta
	// Verifica que el código de estado sea el esperado
	assert.Equal(t, siat.CodeReversionAnulacionRechazada, resp.Body.Content.RespuestaServicioFacturacion.CodigoEstado, resp.Body.Content.RespuestaServicioFacturacion.MensajesList)
	log.Printf("Respuesta ReversionAnulacionFactura: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_VerificarComunicacion valida que el servicio responda correctamente
// a una solicitud de verificación de comunicación con el SIAT.
func TestSiatComputarizadaService_VerificarComunicacion(t *testing.T) {
	godotenv.Load(".env")

	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceComputarizada := siatClient.Computarizada()

	// 3. Verificar comunicación
	req := models.Computarizada().NewVerificarComunicacionBuilder().
		Build()

	resp, err := serviceComputarizada.VerificarComunicacion(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en verificación de comunicación: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta Verificación Comunicación SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_AnulacionDuttyFree valida el flujo completo de anulación
// para una factura del sector Dutty Free (Sector 10).
func TestSiatComputarizadaService_AnulacionDuttyFree(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceComputarizada := siatClient.Computarizada()
	serviceCodigos := siatClient.Codigos()

	// 1. Obtener CUIS
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo
	log.Printf("CUIS: %s", cuis)

	// 2. Obtener CUFD
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()

	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
	log.Printf("CUFD (Codigo): %s", cufd)

	// 3. Preparar solicitud de anulación para Dutty Free (Sector 10 - Tipo Factura 2)
	// CUF de ejemplo (debe ser uno válido emitido previamente)
	cuf := "D5340CCDF031F2596FC03311F6F76AB5334D0A86A626F497FCE6AAF74"

	req := models.Computarizada().NewAnulacionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithTipoFacturaDocumento(2).   // Tipo Factura 2: Sin derecho a crédito fiscal
		WithCodigoEmision(1).          // 1: Online
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis(cuis).
		WithNit(nit).
		WithCufd(cufd).
		WithCuf(cuf).
		WithCodigoMotivo(1). // 1: Factura mal emitida
		Build()

	xmlData, _ := xml.MarshalIndent(req, "", "  ")
	log.Printf("XML Solicitud Anulación: \n%s", string(xmlData))

	// 4. Ejecutar anulación
	resp, err := serviceComputarizada.AnulacionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud de anulación: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta Anulación SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionDuttyFree valida el flujo de emisión y recepción para
// facturas de sector Dutty Free (Sector 10).
// Sigue la regla: subTotal = (cantidad * precioUnitario) - montoDescuento.
func TestSiatComputarizadaService_RecepcionDuttyFree(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}

	serviceComputarizada := siatClient.Computarizada()

	fechaEmision := time.Now()
	// 1. Generar CUF (Usando sector 10 para Dutty Free, Tipo Factura 2 - Sin Crédito Fiscal)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 2, 10, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}
	log.Printf("CUF: %+v", cuf)
	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	cantidad := 1.0
	precioUnitario := 100.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	// Crear objeto de factura usando el nuevo constructor de Dutty Free (Sector 10)
	cabecera := facturas.NewDuttyFreeCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("ESQUINA AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalle := facturas.NewDuttyFreeDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("abc123").
		WithDescripcion("GASA").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := facturas.NewDuttyFreeBuilder().
		WithModalidad(siat.ModalidadComputarizada). // Usar modalidad del ambiente
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar (No se firma en modalidad computarizada)
	xmlData, _ := xml.MarshalIndent(factura, "", "  ")
	log.Printf("XML Factura: \n%s", string(xmlData))

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(10). // Sector 10 para Dutty Free
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(2).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	if resp.Body.Fault != nil {
		t.Fatalf("error en solicitud: %v", resp.Body.Fault.Detail)
	}
	if resp.Body.Content.RespuestaServicioFacturacion.CodigoEstado != siat.CodeRecepcionValidada {
		t.Fatalf("error en solicitud: %v", resp.Body.Content.RespuestaServicioFacturacion.MensajesList)
	}
	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionComercialExportacionServicio valida el flujo para
// facturas de Exportación de Servicios (Sector 28).
// Nota: Se utiliza UnidadMedida 58 (Servicio) y cantidad 1 según normativa SIAT.
// Sigue la regla: subTotal = (cantidad * precioUnitario) - montoDescuento.
func TestSiatComputarizadaService_RecepcionComercialExportacionServicio(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	serviceComputarizada := siatClient.Computarizada()

	fechaEmision := time.Now()
	// 1. Generar CUF (Usando sector 28 para Comercial Exportación de Servicios)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 2, 28, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "CLIENTE EXTRANJERO"
	codigoPuntoVenta := 0

	cantidad := 1.0
	precioUnitario := 100.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem
	tipoCambio := 1.0
	montoTotalMoneda := montoTotal

	// Crear objeto de factura usando el constructor de Sector 28
	cabecera := facturas.NewComercialExportacionServicioCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("AV. HEROES DEL CHACO").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithDireccionComprador("CALLE EXTERIOR 456").
		WithLugarDestino("MIAMI").
		WithCodigoPais(1).
		WithCodigoCliente("EXP-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(tipoCambio).
		WithMontoTotalMoneda(montoTotalMoneda).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalle := facturas.NewComercialExportacionServicioDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("EXP-PROD").
		WithDescripcion("SERVICIO DE EXPORTACION").
		WithUnidadMedida(58).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := facturas.NewComercialExportacionServicioBuilder().
		WithModalidad(codModalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar
	xmlData, _ := xml.Marshal(factura)

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(28).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(2).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionEngarrafadoras valida el flujo de recepción para
// comercializadores de gas licuado de petróleo (Sector 51).
// El subTotal es calculado manualmente siguiendo la regla: (cantidad * precioUnitario) - descuento.
func TestSiatComputarizadaService_RecepcionEngarrafadoras(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	serviceComputarizada := siatClient.Computarizada()

	fechaEmision := time.Now()
	// 1. Generar CUF (Usando sector 51 para Engarrafadoras)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 51, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "CLIENTE GAS"
	codigoPuntoVenta := 0

	cantidad := 1.0
	precioUnitario := 50.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem // En este caso solo un item

	// Crear objeto de factura usando el constructor de Sector 51
	cabecera := facturas.NewEngarrafadorasCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("AV. HEROES DEL CHACO").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("GAS-CLI-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalle := facturas.NewEngarrafadorasDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("GAS-01").
		WithDescripcion("GARRAFA 10KG").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := facturas.NewEngarrafadorasBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar
	xmlData, _ := xml.Marshal(factura)

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(51). // Sector 51
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionComercialExportacion valida el flujo para
// facturas comerciales de exportación (Sector 3).
// El montoTotal de la cabecera debe reflejar la suma de montoDetalle y totalGastosNacionalesFob.
// Sigue la regla: subTotal = (cantidad * precioUnitario) - montoDescuento.
func TestSiatComputarizadaService_RecepcionComercialExportacion(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	serviceComputarizada := siatClient.Computarizada()

	fechaEmision := time.Now()
	// 1. Generar CUF (Usando sector 3 para Comercial Exportación)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 2, 3, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "CLIENTE EXPORTACION"
	codigoPuntoVenta := 0
	costosNacionalesMap := siat.Map{
		"transporte": 100,
	}
	// Aplicar redondeo manual a los costos nacionales si se requiere
	costosNacionales := costosNacionalesMap.Sum()
	cantidad := 1.0
	precioUnitario := 100.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoDetalle := subTotalItem
	totalGastosNacionalesFob := montoDetalle + costosNacionales // montoDetalle + costosNacionales
	montoTotal := totalGastosNacionalesFob

	// Crear objeto de factura usando el constructor de Sector 3
	cabecera := facturas.NewComercialExportacionCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("AV. HEROES DEL CHACO").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithDireccionComprador("CALLE EXTERIOR 101").
		WithCodigoCliente("EXP-CLI-001").
		WithIncoterm("FOB").
		WithIncotermDetalle("FREE ON BOARD").
		WithPuertoDestino("ARICA").
		WithLugarDestino("CHILE").
		WithCodigoPais(14).
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithCostosGastosNacionales(costosNacionalesMap).
		WithTotalGastosNacionalesFob(totalGastosNacionalesFob).
		WithMontoDetalle(montoDetalle).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalle := facturas.NewComercialExportacionDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("EXP-01").
		WithCodigoNandina("100110").
		WithDescripcion("PRODUCTO EXPORTACION").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := facturas.NewComercialExportacionBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar
	xmlData, _ := xml.Marshal(factura)

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(3). // Sector 3
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(2).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionHotel valida el flujo de emisión y recepción para
// facturas de sector Hotel (Sector 16).
// Sigue la regla: subTotal = (cantidad * precioUnitario) - montoDescuento.
func TestSiatComputarizadaService_RecepcionHotel(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	siatClient, err := siat.New(os.Getenv("SIAT_URL"), nil)
	if err != nil {
		t.Fatalf("error initializing SIAT client: %v", err)
	}

	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	serviceComputarizada := siatClient.Computarizada()

	fechaEmision := time.Now()
	// 1. Generar CUF (Usando sector 16 para Hotel)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 16, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "HUESPED PRUEBA"
	codigoPuntoVenta := 0

	cantidad := 1.0
	precioUnitario := 150.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	cantHuespedes := 1
	cantHabitaciones := 1
	cantMayores := 1
	cantMenores := 0

	// Crear objeto de factura usando el constructor de Sector 16
	cabecera := facturas.NewHotelCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("HOTEL PRUEBA").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("AV. PRUEBA").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("CLI-HOT-01").
		WithCantidadHuespedes(&cantHuespedes).
		WithCantidadHabitaciones(&cantHabitaciones).
		WithCantidadMayores(&cantMayores).
		WithCantidadMenores(&cantMenores).
		WithFechaIngresoHospedaje(fechaEmision.Add(-24 * time.Hour)).
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	codTipoHabitacion := 1 // Habitacion simple
	detalleHuespedes := []siat.Map{
		{
			"nombreHuesped":           "HUESPED PRUEBA",
			"documentoIdentificacion": "5115889",
			"codigoPais":              "1",
		},
	}

	detalle := facturas.NewHotelDetalleBuilder().
		WithActividadEconomica("551010").
		WithCodigoProductoSin(99100).
		WithCodigoProducto("HAB-01").
		WithCodigoTipoHabitacion(&codTipoHabitacion).
		WithDescripcion("HOSPEDAJE POR 1 NOCHE").
		WithCantidad(cantidad).
		WithUnidadMedida(58). // Servicio
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		WithDetalleHuespedes(detalleHuespedes).
		Build()

	factura := facturas.NewHotelBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar
	xmlData, _ := xml.Marshal(factura)

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(16). // Sector 16
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionLubricantesIehd valida el flujo de emisión y recepción para
// facturas de sector Lubricantes IEHD (Sector 53).
// Sigue la regla: subTotal = (cantidad * precioUnitario) - montoDescuento.
func TestSiatComputarizadaService_RecepcionLubricantesIehd(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	siatClient, err := siat.New(os.Getenv("SIAT_URL"), nil)
	if err != nil {
		t.Fatalf("error initializing SIAT client: %v", err)
	}

	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	serviceComputarizada := siatClient.Computarizada()

	fechaEmision := time.Now()
	// 1. Generar CUF (Usando sector 53)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 53, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "CLIENTE PRUEBA"
	codigoPuntoVenta := 0

	cantidad := 5.0
	precioUnitario := 40.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	ciudad := "SANTA CRUZ"
	propietario := "JUAN PEREZ"
	representante := "PEDRO REPRESENTANTE"
	condicionPago := "EFECTIVO"
	periodoEntrega := "INMEDIATA"

	porcentajeDeduccion := 2.5
	montoDeduccion := cantidad * porcentajeDeduccion

	// Crear objeto de factura usando el constructor de Sector 53
	cabecera := facturas.NewLubricantesIehdCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("EMPRESA LUBRICANTES PRUEBA").
		WithMunicipio("SANTA CRUZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("AV. PRUEBA").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1). // CI
		WithNumeroDocumento("5115889").
		WithCodigoCliente("CLI-LUB-01").
		WithCiudad(&ciudad).
		WithNombrePropietario(&propietario).
		WithNombreRepresentanteLegal(&representante).
		WithCondicionPago(&condicionPago).
		WithPeriodoEntrega(&periodoEntrega).
		WithCodigoMetodoPago(1). // Efectivo
		WithMontoTotal(montoTotal).
		WithMontoDeduccionIehdDS25530(&montoDeduccion).
		WithMontoTotalSujetoIva(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithDescuentoAdicional(nil).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalle := facturas.NewLubricantesIehdDetalleBuilder().
		WithActividadEconomica("466110"). // Venta de lubricantes (ejemplo)
		WithCodigoProductoSin(99100).     // Código SIN
		WithCodigoProducto("LUB-001").
		WithDescripcion("LUBRICANTE 20W50").
		WithCantidad(cantidad).
		WithUnidadMedida(26). // Litros u otra unidad según catálogo
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		WithCantidadLitros(cantidad).
		WithPorcentajeDeduccionIehdDS25530(&porcentajeDeduccion).
		Build()

	factura := facturas.NewLubricantesIehdBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar
	xmlData, _ := xml.Marshal(factura)

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(53). // Sector 53
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1). // Factura con derecho a crédito fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionAlcanzadaIce valida el flujo de emisión y recepción para
// facturas Alcanzadas por ICE (Sector 14).
func TestSiatComputarizadaService_RecepcionAlcanzadaIce(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	siatClient, err := siat.New(os.Getenv("SIAT_URL"), nil)
	if err != nil {
		t.Fatalf("error initializing SIAT client: %v", err)
	}

	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	serviceComputarizada := siatClient.Computarizada()

	fechaEmision := time.Now()
	// 1. Generar CUF (Usando sector 14)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 14, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "CLIENTE PRUEBA ICE"
	codigoPuntoVenta := 0

	cantidad := 10.0
	precioUnitario := 50.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	montoIceEspecifico := 15.50
	montoIcePorcentual := 5.00

	// Crear objeto de factura usando el constructor de Sector 14
	cabecera := facturas.NewAlcanzadaIceCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("EMPRESA BEBIDAS PRUEBA").
		WithMunicipio("COCHABAMBA").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("AV. PRUEBA ICE").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1). // CI
		WithNumeroDocumento("5115889").
		WithCodigoCliente("CLI-ICE-01").
		WithCodigoMetodoPago(1). // Efectivo
		WithMontoTotal(montoTotal).
		WithMontoIceEspecifico(&montoIceEspecifico).
		WithMontoIcePorcentual(&montoIcePorcentual).
		WithMontoTotalSujetoIva(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithDescuentoAdicional(nil).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario_ice").
		Build()

	alicuotaIva := 14.94
	precioNetoVentaIce := precioUnitario
	alicuotaEspecifica := 1.55
	alicuotaPorcentual := 0.50
	cantidadIce := cantidad

	detalle := facturas.NewAlcanzadaIceDetalleBuilder().
		WithActividadEconomica("110100").
		WithCodigoProductoSin(99100).
		WithCodigoProducto("BEB-001").
		WithDescripcion("BEBIDA ALCOHOLICA PRUEBA").
		WithCantidad(cantidad).
		WithUnidadMedida(26).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		WithMarcaIce(1).
		WithAlicuotaIva(&alicuotaIva).
		WithPrecioNetoVentaIce(&precioNetoVentaIce).
		WithAlicuotaEspecifica(&alicuotaEspecifica).
		WithAlicuotaPorcentual(&alicuotaPorcentual).
		WithMontoIceEspecifico(&montoIceEspecifico).
		WithMontoIcePorcentual(&montoIcePorcentual).
		WithCantidadIce(&cantidadIce).
		Build()

	factura := facturas.NewAlcanzadaIceBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar
	xmlData, _ := xml.Marshal(factura)

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(14). // Sector 14
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1). // Factura con derecho a crédito fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionAnexosSuministroEnergia valida el servicio de envío de anexos
// para el suministro de energía eléctrica (Electrolineras).
// Según la documentación de SIAT, para el Sector 41 (Electrolineras), se debe utilizar la
// Modalidad 1 (Electrónica en Línea). El test verifica la construcción de la solicitud
// con CUIS, CUFD y el nuevo campo TipoFacturaDocumento.
func TestSiatComputarizadaService_RecepcionAnexosSuministroEnergia(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()

	// 1. Obtener CUIS
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	// 2. Obtener CUFD
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()

	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo

	// 3. Generar CUF para el anexo
	fechaRecarga := time.Now()
	numeroFactura := 1
	tipoFactura := 2 // Sin Crédito Fiscal / Sectorial
	cufAnexo, _ := utils.GenerarCUF(
		nit,
		fechaRecarga,
		0,                  // Sucursal
		codModalidad,       // Modalidad
		siat.EmisionOnline, // Tipo Emisión
		tipoFactura,        // Tipo Factura
		34,                 // Tipo Documento Sector (Suministro de Energía)
		numeroFactura,      // Número de Factura
		0,                  // Punto de Venta
		cufd,               // Código de Control (CUFD)
	)

	// 4. Preparar anexo
	anexo := models.Computarizada().NewSuministroEnergiaAnexoBuilder().
		WithCufFactSuministro(cufAnexo).
		WithFechaRecarga(fechaRecarga).
		WithMontoRecarga(100.50).
		Build()

	// 5. Preparar solicitud principal
	req := models.Computarizada().NewRecepcionAnexosSuministroEnergiaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(34). // Sector 34: Suministro de Energía Eléctrica
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(tipoFactura).
		AddAnexos(anexo).
		WithGiftCard(0).
		Build()

	// 5. Ejecutar solicitud
	resp, err := serviceComputarizada.RecepcionAnexosSuministroEnergia(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en RecepcionAnexosSuministroEnergia: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta RecepcionAnexosSuministroEnergia: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_VerificacionEstadoFactura valida la consulta del estado actual
// de una factura en los servidores del SIAT.
func TestSiatComputarizadaService_VerificacionEstadoFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()

	// 1. Obtener CUIS
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	// 2. Obtener CUFD
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()

	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo

	// 3. Preparar solicitud de verificación
	req := models.Computarizada().NewVerificacionEstadoFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty free
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2).
		WithCuf("D5340CCDF031F2596F...").
		Build()

	// 4. Ejecutar solicitud
	resp, err := serviceComputarizada.VerificacionEstadoFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en VerificacionEstadoFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta VerificacionEstadoFactura: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionPaqueteFactura valida el envío de un paquete de
// facturas (contingencia) al SIAT.
func TestSiatComputarizadaService_RecepcionPaqueteFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()

	// 1. Obtener CUIS y CUFD
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()
	cuisResp, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()
	cufdResp, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo

	// 2. Construir Factura Dutty Free (Ejemplo para Paquete)
	fechaEmision := time.Now()
	numeroFactura := int64(200)

	// Obtener el código de control original para el CUF desde la respuesta del CUFD
	cufdOriginal := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	cuf, err := utils.GenerarCUF(
		nit,                 // NIT
		fechaEmision,        // Fecha
		0,                   // Sucursal
		codModalidad,        // Modalidad
		siat.EmisionOffline, // Tipo Emisión (Offline para Paquete)
		2,                   // Tipo Factura (Sin Crédito Fiscal)
		10,                  // Sector (Dutty Free)
		int(numeroFactura),  // Número Factura
		0,                   // Punto Venta
		cufdOriginal,        // CUFD (Código Control)
	)
	if err != nil {
		t.Fatalf("error generando CUF: %v", err)
	}
	razonSocial := "RAZON SOCIAL TEST"
	factura := facturas.NewDuttyFreeBuilder().
		WithModalidad(codModalidad).
		WithCabecera(facturas.NewDuttyFreeCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor(razonSocial).
			WithMunicipio("LA PAZ").
			WithNumeroFactura(numeroFactura).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("CALLE TEST 123").
			WithFechaEmision(fechaEmision).
			WithNombreRazonSocial(&razonSocial).
			WithCodigoTipoDocumentoIdentidad(1). // CI
			WithNumeroDocumento("1234567").
			WithCodigoCliente("CLI-001").
			WithCodigoMetodoPago(1). // Efectivo
			WithMontoTotal(250.0).
			WithCodigoMoneda(1). // Boliviano
			WithTipoCambio(1.0).
			WithMontoTotalMoneda(250.0).
			WithLeyenda("Leyenda Test Paquete").
			WithUsuario("usuario_test").
			Build()).
		AddDetalle(facturas.NewDuttyFreeDetalleBuilder().
			WithActividadEconomica("86111").
			WithCodigoProductoSin(86111).
			WithCodigoProducto("PROD-DF-01").
			WithDescripcion("Item Duty Free Paquete").
			WithCantidad(2.0).
			WithUnidadMedida(1).
			WithPrecioUnitario(125.0).
			WithSubTotal(250.0).
			Build()).
		Build()

	xmlData, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error serializando factura: %v", err)
	}

	// SIAT Paquete requiere un archivo .tar.gz que contenga los XMLs
	tarGz, err := utils.CreateTarGz(map[string][]byte{
		"factura.xml": xmlData,
	})
	if err != nil {
		t.Fatalf("error creando tar.gz: %v", err)
	}

	hashArchivo := utils.SHA256Hex(tarGz)
	encodedArchivo := base64.StdEncoding.EncodeToString(tarGz)

	// 3. Preparar solicitud
	req := models.Computarizada().NewRecepcionPaqueteFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithCodigoEmision(siat.EmisionOffline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2). // 2: Sin Crédito Fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashArchivo).
		WithCantidadFacturas(1).
		WithCodigoEvento(0).
		Build()

	// 4. Ejecutar
	resp, err := serviceComputarizada.RecepcionPaqueteFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en RecepcionPaqueteFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta RecepcionPaqueteFactura: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_ValidacionRecepcionPaqueteFactura valida el estado de un paquete
// de facturas enviado previamente bajo contingencia (Emisión Offline).
// Se utiliza el Sector 10 (Dutty Free) como ejemplo de sector sin crédito fiscal.
func TestSiatComputarizadaService_ValidacionRecepcionPaqueteFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()

	// 1. Obtener CUIS y CUFD
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()
	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()
	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo

	// 2. Preparar solicitud
	req := models.Computarizada().NewValidacionRecepcionPaqueteFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionOffline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoRecepcion("123ABCD").
		WithTipoFacturaDocumento(2).
		Build()

	// 3. Ejecutar
	resp, err := serviceComputarizada.ValidacionRecepcionPaqueteFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en ValidacionRecepcionPaqueteFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta ValidacionRecepcionPaqueteFactura: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionMasivaFactura valida el envío masivo de facturas (Emisión Masiva).
// Requiere que las facturas XML sean empaquetadas en un archivo .tar.gz codificado en Base64.
// En este test se utiliza el Sector 10 (Dutty Free) y la Modalidad Computarizada (2),
// empaquetando una factura estructurada real mediante la utilidad CreateTarGz.
func TestSiatComputarizadaService_RecepcionMasivaFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()

	// 1. Obtener CUIS y CUFD
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()
	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()
	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo

	// 2. Construir Factura Dutty Free (Ejemplo para Masiva)
	fechaEmision := time.Now()
	numeroFactura := int64(100) // Diferente para evitar colisión

	// Obtener el código de control original para el CUF
	cufdOriginal := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	cuf, err := utils.GenerarCUF(
		nit,                // NIT
		fechaEmision,       // Fecha
		0,                  // Sucursal
		codModalidad,       // Modalidad
		siat.EmisionMasiva, // Tipo Emisión (Masiva)
		2,                  // Tipo Factura (Sin Crédito Fiscal)
		10,                 // Sector (Dutty Free)
		int(numeroFactura), // Número Factura
		0,                  // Punto Venta
		cufdOriginal,       // CUFD (Código Control)
	)
	if err != nil {
		t.Fatalf("error generando CUF: %v", err)
	}

	factura := facturas.NewDuttyFreeBuilder().
		WithModalidad(codModalidad).
		WithCabecera(facturas.NewDuttyFreeCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("RAZON SOCIAL TEST").
			WithMunicipio("LA PAZ").
			WithNumeroFactura(numeroFactura).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("CALLE TEST 123").
			WithFechaEmision(fechaEmision).
			WithCodigoTipoDocumentoIdentidad(1). // CI
			WithNumeroDocumento("1234567").
			WithCodigoCliente("CLI-001").
			WithCodigoMetodoPago(1). // Efectivo
			WithMontoTotal(250.0).
			WithCodigoMoneda(1). // Boliviano
			WithTipoCambio(1.0).
			WithMontoTotalMoneda(250.0).
			WithLeyenda("Leyenda Test Masiva").
			WithUsuario("usuario_test").
			Build()).
		AddDetalle(facturas.NewDuttyFreeDetalleBuilder().
			WithActividadEconomica("86111").
			WithCodigoProductoSin(86111).
			WithCodigoProducto("PROD-DF-01").
			WithDescripcion("Item Duty Free Masiva").
			WithCantidad(2.0).
			WithUnidadMedida(1).
			WithPrecioUnitario(125.0).
			WithSubTotal(250.0).
			Build()).
		Build()

	xmlData, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error serializando factura: %v", err)
	}

	// SIAT Masiva requiere un archivo .tar.gz que contenga los XMLs
	tarGz, err := utils.CreateTarGz(map[string][]byte{
		"factura.xml": xmlData,
	})
	if err != nil {
		t.Fatalf("error creando tar.gz: %v", err)
	}

	hashArchivo := utils.SHA256Hex(tarGz)
	encodedArchivo := base64.StdEncoding.EncodeToString(tarGz)

	// 3. Preparar solicitud
	req := models.Computarizada().NewRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty free
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2). // 2: Sin Crédito Fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(time.Now()).
		WithHashArchivo(hashArchivo).
		WithCantidadFacturas(1).
		Build()

	// 4. Ejecutar
	resp, err := serviceComputarizada.RecepcionMasivaFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en RecepcionMasivaFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta RecepcionMasivaFactura: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_ValidacionRecepcionMasivaFactura valida el estado de un envío masivo
// utilizando el código de recepción obtenido en una solicitud previa.
func TestSiatComputarizadaService_ValidacionRecepcionMasivaFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()

	// 1. Obtener CUIS y CUFD
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()
	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()
	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo

	// 2. Preparar solicitud
	req := models.Computarizada().NewValidacionRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty free
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2). // 2: Sin Crédito Fiscal
		WithCodigoRecepcion("0e7c84b7-23da-11f1-b19f-51e43057f8a6").
		Build()

	// 3. Ejecutar
	resp, err := serviceComputarizada.ValidacionRecepcionMasivaFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en ValidacionRecepcionMasivaFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta ValidacionRecepcionMasivaFactura: %+v", resp.Body.Content)
}

// TestSiatComputarizadaService_RecepcionFacturaDuttyFree valida la recepción individual de una
// factura Dutty Free (Sector 10). A diferencia de la modalidad electrónica, la computarizada
// no requiere firma digital pero sí la generación de un CUF válido basado en el CUFD vigente.
func TestSiatComputarizadaService_RecepcionFacturaDuttyFree(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceComputarizada := siatClient.Computarizada()

	// 1. Obtener CUIS y CUFD
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()
	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()
	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
	cufdOriginal := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	// 2. Construir Factura Dutty Free
	fechaEmision := time.Now()
	numeroFactura := int64(1)

	// Generar CUF
	cuf, err := utils.GenerarCUF(
		nit,                // NIT
		fechaEmision,       // Fecha
		0,                  // Sucursal
		codModalidad,       // Modalidad
		siat.EmisionOnline, // Tipo Emisión
		2,                  // Tipo Factura (Sin Crédito Fiscal)
		10,                 // Sector (Dutty Free)
		int(numeroFactura), // Número Factura
		0,                  // Punto Venta
		cufdOriginal,       // CUFD (Código Control)
	)
	if err != nil {
		t.Fatalf("error generando CUF: %v", err)
	}

	factura := facturas.NewDuttyFreeBuilder().
		WithModalidad(codModalidad).
		WithCabecera(facturas.NewDuttyFreeCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("RAZON SOCIAL TEST").
			WithMunicipio("LA PAZ").
			WithNumeroFactura(numeroFactura).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("CALLE TEST 123").
			WithFechaEmision(fechaEmision).
			WithCodigoTipoDocumentoIdentidad(1). // CI
			WithNumeroDocumento("1234567").
			WithCodigoCliente("CLI-001").
			WithCodigoMetodoPago(1). // Efectivo
			WithMontoTotal(100.0).
			WithCodigoMoneda(1). // Boliviano
			WithTipoCambio(1.0).
			WithMontoTotalMoneda(100.0).
			WithLeyenda("Leyenda Test").
			WithUsuario("usuario_test").
			Build()).
		AddDetalle(facturas.NewDuttyFreeDetalleBuilder().
			WithActividadEconomica("86111"). // Ejemplo
			WithCodigoProductoSin(86111).
			WithCodigoProducto("PROD-001").
			WithDescripcion("Producto de prueba").
			WithCantidad(1.0).
			WithUnidadMedida(1).
			WithPrecioUnitario(100.0).
			WithSubTotal(100.0).
			Build()).
		Build()

	// 3. Serializar y Comprimir
	xmlData, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error serializando factura: %v", err)
	}

	hashArchivo, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error comprimiendo factura: %v", err)
	}

	// 4. Preparar solicitud de recepción
	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(nit).
		WithTipoFacturaDocumento(2).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(time.Now()).
		WithHashArchivo(hashArchivo).
		Build()

	// 5. Enviar al SIAT
	resp, err := serviceComputarizada.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en RecepcionFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta RecepcionFactura Dutty Free: %+v", resp.Body.Content)
}
