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
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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

	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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

// TestSiatComputarizadaService_VerificacionEstadoFactura valida la consulta del estado actual
// de una factura en los servidores del SIAT.
func TestSiatComputarizadaService_VerificacionEstadoFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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

func TestSiatComputarizadaService_RecepcionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadComputarizada
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

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
	// 1. Generar CUF
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, siat.ModalidadComputarizada, 1, 2, 10, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0
	cantidad := 1.0
	precioUnitario := 100.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

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
		WithModalidad(codModalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, _ := xml.Marshal(factura)

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
		WithCodigoDocumentoSector(10).
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
	log.Printf("Respuesta SIAT (Computarizada): %+v", resp.Body.Content)
}
