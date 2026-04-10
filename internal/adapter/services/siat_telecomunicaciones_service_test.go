package services_test

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSiatTelecomunicacionesService_VerificarComunicacion(t *testing.T) {
	godotenv.Load(".env")
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Telecomunicaciones()

	req := models.Telecomunicaciones().NewVerificarComunicacionBuilder().Build()
	resp, err := service.VerificarComunicacion(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		assert.True(t, resp.Body.Content.Return.Transaccion)
		log.Printf("Respuesta Comunicacion: %+v", resp.Body.Content)
	}
}

func TestSiatTelecomunicacionesService_RecepcionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceTelecom := siatClient.Telecomunicaciones()

	// 1. Obtener CUIS
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

	// 2. Obtener CUFD
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

	fechaEmision := time.Now()
	// 3. Generar CUF (Sector 22)
	// Args: nit, fecha, sucursal, modalidad, emision, tipoFactura, sector, nFactura, puntoVenta, control
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 22, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir Factura Telecomunicaciones (Sector 22)
	nombre := "JUAN PEREZ"
	tel := "2222222"
	cabecera := invoices.NewTelecomunicacionesCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("EMPRESA TELECOM S.A.").
		WithMunicipio("La Paz").
		WithTelefono(&tel).
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("Av. Principal 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley 453...").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewTelecomunicacionesDetalleBuilder().
		WithActividadEconomica("611000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Servicio de Internet").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := invoices.NewTelecomunicacionesBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 5. Serializar, Firmar, Comprimir y Hashear
	xmlData, _ := xml.Marshal(factura)
	signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")
	hashString, encodedArchivo, _ := utils.CompressAndHash(signedXML)

	// 6. Preparar solicitud de recepción
	req := models.Telecomunicaciones().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(22).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceTelecom.RecepcionFactura(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion Telecom: %+v", resp.Body.Content)
	}
}

func TestSiatTelecomunicacionesService_VerificacionEstadoFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	serviceCodigos := siatClient.Codigos()
	serviceTelecom := siatClient.Telecomunicaciones()

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

	req := models.Telecomunicaciones().NewVerificacionEstadoFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(22).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithCuf("ABC123FAKE").
		Build()

	resp, err := serviceTelecom.VerificacionEstadoFactura(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Verificacion Estado: %+v", resp.Body.Content)
	}
}

func TestSiatTelecomunicacionesService_AnulacionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	serviceCodigos := siatClient.Codigos()
	serviceTelecom := siatClient.Telecomunicaciones()

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

	req := models.Telecomunicaciones().NewAnulacionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(22).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithCuf("ABC123FAKE").
		WithCodigoMotivo(1).
		Build()

	resp, err := serviceTelecom.AnulacionFactura(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Anulación: %+v", resp.Body.Content)
	}
}
