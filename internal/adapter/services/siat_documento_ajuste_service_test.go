package services_test

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/ron86i/go-siat/v2/pkg/models/invoices"
	"github.com/ron86i/go-siat/v2/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func getDocumentoAjusteSetup(t *testing.T) (*siat.SiatServices, string, string, string, int64, int) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	codModalidad := siat.ModalidadElectronica // Asumiendo modalidad por defecto

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: codAmbiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     siat.NewHTTPClient(siat.DefaultHTTPConfig()),
	}

	siatClient, err := siat.New(cfg)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	serviceCodigos := siatClient.Codigos()

	cuisReq := models.NewCuisBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		Build()
	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), cuisReq)
	if err != nil {
		t.Fatalf("error solicitando CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	cufdReq := models.NewCufdBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCuis(cuis).
		Build()
	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), cufdReq)
	if err != nil {
		t.Fatalf("error solicitando CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
	cufdControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	return siatClient, cuis, cufd, cufdControl, nit, codModalidad
}

func TestSiatDocumentoAjuste_VerificarComunicacion(t *testing.T) {
	siatClient, _, _, _, _, _ := getDocumentoAjusteSetup(t)
	service := siatClient.DocumentoAjuste()

	req := models.NewVerificarComunicacionDocumentoAjusteBuilder().Build()
	resp, err := service.VerificarComunicacion(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		assert.True(t, resp.Body.Content.Return.Transaccion)
		log.Printf("Respuesta Comunicacion: %+v", resp.Body.Content)
	}
}

func TestSiatDocumentoAjuste_RecepcionDocumentoAjuste(t *testing.T) {
	siatClient, cuis, cufd, cufdControl, nit, codModalidad := getDocumentoAjusteSetup(t)
	service := siatClient.DocumentoAjuste()

	fechaEmision := time.Now()
	// Generar CUF para Nota Conciliacion (Sector 29)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, siat.EmisionOnline, 3, 29, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	cabecera := invoices.NewNotaConciliacionCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroNotaConciliacion(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("ESQUINA AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithNumeroFactura(10).
		WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
		WithFechaEmisionFactura(fechaEmision.Add(-24 * time.Hour)).
		WithMontoTotalOriginal(200.0).
		WithMontoTotalConciliado(150.0).
		WithCreditoFiscalIva(0).
		WithDebitoFiscalIva(0).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalleA := invoices.NewNotaDetalleOriginalBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("PROD-01").
		WithDescripcion("Producto DUMMY").
		WithCantidad(2.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(100.0).
		WithSubTotal(200.0).
		Build()

	detalleB := invoices.NewNotaDetalleConciliacionBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("PROD-01").
		WithDescripcion("Producto DUMMY").
		WithMontoOriginal(200.0).
		WithMontoFinal(50.0).
		WithMontoConciliado(150.0).
		Build()

	nota := invoices.NewNotaConciliacionBuilder().
		WithModalidad(codModalidad).
		WithCabecera(cabecera).
		AddDetalleOriginal(detalleA).
		AddDetalleConciliacion(detalleB).
		Build()

	xmlData, _ := xml.Marshal(nota)

	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.NewRecepcionDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(3).
		WithArchivo(encodedArchivo).
		WithHashArchivo(hashString).
		WithFechaEnvio(fechaEmision).
		Build()

	resp, err := service.RecepcionDocumentoAjuste(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion: %+v", resp.Body.Content)
	}
}

func TestSiatDocumentoAjuste_AnulacionDocumentoAjuste(t *testing.T) {
	siatClient, cuis, cufd, _, _, _ := getDocumentoAjusteSetup(t)
	service := siatClient.DocumentoAjuste()

	cuf := "CUF_FROM_INVOICE"

	req := models.NewAnulacionDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(3).
		WithCuf(cuf).
		WithCodigoMotivo(1).
		Build()

	resp, err := service.AnulacionDocumentoAjuste(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Anulacion: %+v", resp.Body.Content)
	}
}

func TestSiatDocumentoAjuste_ReversionAnulacionDocumentoAjuste(t *testing.T) {
	siatClient, cuis, cufd, _, _, _ := getDocumentoAjusteSetup(t)
	service := siatClient.DocumentoAjuste()

	cuf := "CUF_FROM_INVOICE"

	req := models.NewReversionAnulacionDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(3).
		WithCuf(cuf).
		Build()

	resp, err := service.ReversionAnulacionDocumentoAjuste(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Reversion Anulacion: %+v", resp.Body.Content)
	}
}

func TestSiatDocumentoAjuste_VerificacionEstadoDocumentoAjuste(t *testing.T) {
	siatClient, cuis, cufd, _, _, _ := getDocumentoAjusteSetup(t)
	service := siatClient.DocumentoAjuste()

	cuf := "CUF_FROM_INVOICE"

	req := models.NewVerificacionEstadoDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(3).
		WithCuf(cuf).
		Build()

	resp, err := service.VerificacionEstadoDocumentoAjuste(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Verificacion Estado: %+v", resp.Body.Content)
	}
}
