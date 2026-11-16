package services_test

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func getDocumentoAjusteSetup(t *testing.T) (*siat.SiatServices, string, string, string, int64, int) {
	if os.Getenv("SIAT_RUN") != "1" {
		t.Skip("prueba de integración real: requiere SIAT_RUN=1")
	}

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

	nota := buildNotaConciliacionFixture(nit, cuf, cufd, fechaEmision, codModalidad)

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
