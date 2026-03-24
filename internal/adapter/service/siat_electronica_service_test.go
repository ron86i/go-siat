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

func TestSiatElectronicaService_ReversionAnulacionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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

	cuf := "D5340CCDF031F2596FC03311F6F76AB5334D0A86A626F497FCE6AAF74"

	req := models.Electronica().NewReversionAnulacionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithTipoFacturaDocumento(2).   // Sin Crédito Fiscal
		WithCodigoEmision(1).
		WithCuf(cuf).
		Build()

	resp, err := serviceElectronica.ReversionAnulacionFactura(context.Background(), config, req)
	assert.NoError(t, err)
	assert.Nil(t, resp.Body.Fault)
	log.Printf("Respuesta ReversionAnulacionFactura: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_VerificarComunicacion(t *testing.T) {
	godotenv.Load(".env")

	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceElectronica := siatClient.Electronica()

	req := models.Electronica().NewVerificarComunicacionBuilder().
		Build()

	resp, err := serviceElectronica.VerificarComunicacion(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en verificación de comunicación: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta Verificación Comunicación SIAT: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_VerificacionEstadoFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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
	req := models.Electronica().NewVerificacionEstadoFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2). // Sin Crédito Fiscal
		WithCuf("D5340CCDF031F2596F...").
		Build()

	resp, err := serviceElectronica.VerificacionEstadoFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en VerificacionEstadoFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta VerificacionEstadoFactura: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_RecepcionPaqueteFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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
	cufdControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	// 2. Construir Factura (Ejemplo para Paquete)
	fechaEmision := time.Now()
	numeroFactura := 1

	cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, siat.EmisionOffline, 2, 10, numeroFactura, 0, cufdControl)

	factura := facturas.NewDuttyFreeBuilder().
		WithModalidad(codModalidad).
		WithCabecera(facturas.NewDuttyFreeCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroFactura(int64(numeroFactura)).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Calle 1").
			WithFechaEmision(fechaEmision).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("1234567").
			WithCodigoCliente("CLI01").
			WithCodigoMetodoPago(1).
			WithMontoTotal(100.0).
			WithCodigoMoneda(1).
			WithTipoCambio(1.0).
			WithMontoTotalMoneda(100.0).
			WithLeyenda("Leyenda").
			WithUsuario("user").
			Build()).
		AddDetalle(facturas.NewDuttyFreeDetalleBuilder().
			WithActividadEconomica("123456").
			WithCodigoProductoSin(123456).
			WithCodigoProducto("P01").
			WithDescripcion("Producto").
			WithCantidad(1.0).
			WithUnidadMedida(1).
			WithPrecioUnitario(100.0).
			WithSubTotal(100.0).
			Build()).
		Build()

	xmlData, _ := xml.Marshal(factura)
	signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")

	tarGz, _ := utils.CreateTarGz(map[string][]byte{
		"factura.xml": signedXML,
	})

	hashArchivo := utils.SHA256Hex(tarGz)
	encodedArchivo := base64.StdEncoding.EncodeToString(tarGz)

	// 3. Preparar solicitud
	req := models.Electronica().NewRecepcionPaqueteFacturaBuilder().
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
		WithTipoFacturaDocumento(2). // Sin Crédito Fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashArchivo).
		WithCantidadFacturas(1).
		Build()

	// 4. Ejecutar
	resp, err := serviceElectronica.RecepcionPaqueteFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en RecepcionPaqueteFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta RecepcionPaqueteFactura: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_ValidacionRecepcionPaqueteFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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

	// 2. Preparar solicitud
	req := models.Electronica().NewValidacionRecepcionPaqueteFacturaBuilder().
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
		WithCodigoRecepcion("123ABCD").
		WithTipoFacturaDocumento(2). // Sin Crédito Fiscal
		Build()

	// 3. Ejecutar
	resp, err := serviceElectronica.ValidacionRecepcionPaqueteFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en ValidacionRecepcionPaqueteFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta ValidacionRecepcionPaqueteFactura: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_RecepcionMasivaFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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
	cufdControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	// 2. Construir Factura (Ejemplo para Masiva)
	fechaEmision := time.Now()
	numeroFactura := 1

	cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, siat.EmisionMasiva, 2, 10, numeroFactura, 0, cufdControl)

	factura := facturas.NewDuttyFreeBuilder().
		WithModalidad(codModalidad).
		WithCabecera(facturas.NewDuttyFreeCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroFactura(int64(numeroFactura)).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Calle 1").
			WithFechaEmision(fechaEmision).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("1234567").
			WithCodigoCliente("CLI01").
			WithCodigoMetodoPago(1).
			WithMontoTotal(100.0).
			WithCodigoMoneda(1).
			WithTipoCambio(1.0).
			WithMontoTotalMoneda(100.0).
			WithLeyenda("Leyenda").
			WithUsuario("user").
			Build()).
		AddDetalle(facturas.NewDuttyFreeDetalleBuilder().
			WithActividadEconomica("123456").
			WithCodigoProductoSin(123456).
			WithCodigoProducto("P01").
			WithDescripcion("Producto").
			WithCantidad(1.0).
			WithUnidadMedida(1).
			WithPrecioUnitario(100.0).
			WithSubTotal(100.0).
			Build()).
		Build()

	xmlData, _ := xml.Marshal(factura)
	signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")

	tarGz, _ := utils.CreateTarGz(map[string][]byte{
		"factura.xml": signedXML,
	})

	hashArchivo := utils.SHA256Hex(tarGz)
	encodedArchivo := base64.StdEncoding.EncodeToString(tarGz)

	// 3. Preparar solicitud
	req := models.Electronica().NewRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2). // Sin Crédito Fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashArchivo).
		WithCantidadFacturas(1).
		Build()

	// 4. Ejecutar
	resp, err := serviceElectronica.RecepcionMasivaFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en RecepcionMasivaFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta RecepcionMasivaFactura: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_ValidacionRecepcionMasivaFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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

	// 2. Preparar solicitud
	req := models.Electronica().NewValidacionRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2). // Sin Crédito Fiscal
		WithCodigoRecepcion("123ABCD").
		Build()

	// 3. Ejecutar
	resp, err := serviceElectronica.ValidacionRecepcionMasivaFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en ValidacionRecepcionMasivaFactura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta ValidacionRecepcionMasivaFactura: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_RecepcionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
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

	serviceElectronica := siatClient.Electronica()

	fechaEmision := time.Now()
	// 1. Generar CUF
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 2, 10, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
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
		WithRazonSocialEmisor("Empresa Test").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("Calle Test").
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
		WithLeyenda("Leyenda").
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
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error firmando XML: %v", err)
	}

	hashString, encodedArchivo, err := utils.CompressAndHash(signedXML)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Electronica().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(2). // Sin Crédito Fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceElectronica.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT (Electronica): %+v", resp.Body.Content)
}

func TestSiatElectronicaService_AnulacionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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

	req := models.Electronica().NewAnulacionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCuis(cuis).
		WithCufd(cufd).
		WithCodigoMotivo(1).
		WithCuf("D5340CCDF031F2596FC03311F6F76AB5334D0A86A626F497FCE6AAF74").
		Build()

	resp, err := serviceElectronica.AnulacionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en anulación de factura: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta Anulación SIAT: %+v", resp.Body.Content)
}

func TestSiatElectronicaService_RecepcionAnexosSuministroEnergia(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceElectronica := siatClient.Electronica()

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

	req := models.Electronica().NewRecepcionAnexosSuministroEnergiaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCuis(cuis).
		WithCufd(cufd).
		AddAnexos(
			models.Electronica().NewSuministroEnergiaAnexoBuilder().
				WithCufFactSuministro("CUF123").
				WithFechaRecarga(time.Now()).
				WithMontoRecarga(10.0).
				Build(),
		).
		Build()

	resp, err := serviceElectronica.RecepcionAnexosSuministroEnergia(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en recepción de anexos: %v", err)
	}

	assert.NotNil(t, resp)
}
