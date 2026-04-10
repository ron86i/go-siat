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
	"github.com/ron86i/go-siat/internal/core/ports"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func runDocumentoAjusteTest[ReqType any, RespType any](
	t *testing.T,
	name string,
	req ReqType,
	fn func(context.Context, ports.Config, ReqType) (*RespType, error),
) {
	t.Run(name, func(t *testing.T) {
		if _, err := os.Stat(".env"); os.IsNotExist(err) {
			t.Skip("Saltando prueba de integración: .env no encontrado")
		}
		godotenv.Load()

		config := siat.Config{
			Token: os.Getenv("SIAT_TOKEN"),
		}

		_, err := siat.New(os.Getenv("SIAT_URL"), nil)
		if err != nil {
			t.Fatalf("No se pudo inicializar el cliente SIAT: %v", err)
		}

		resp, err := fn(context.Background(), config, req)
		if err != nil {
			t.Fatalf("Error en %s: %v", name, err)
		}

		assert.NotNil(t, resp)
		log.Printf("Resultado de %s: %+v", name, resp)
	})
}

func TestSiatDocumentoAjuste_VerificarComunicacion(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load()
	ctx := context.Background()
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.DocumentoAjuste()

	req := models.DocumentoAjuste().NewVerificarComunicacionBuilder().Build()

	resp, err := service.VerificarComunicacion(ctx, config, req)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta comunicación: %+v", resp)
}

func TestSiatDocumentoAjuste_RecepcionDocumentoAjuste(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load()

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

	fechaEmision := time.Now()
	// 1. Generar CUF para Nota Conciliacion (Sector 29)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, siat.EmisionOnline, 3, 29, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
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
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
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

	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithNit(nit).
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(codModalidad).
		WithTipoFacturaDocumento(3).
		WithArchivo(encodedArchivo).
		WithHashArchivo(hashString).
		WithFechaEnvio(fechaEmision).
		Build()

	serviceDocumentoAjuste := siatClient.DocumentoAjuste()

	runDocumentoAjusteTest(t, "RecepcionDocumentoAjuste", req, serviceDocumentoAjuste.RecepcionDocumentoAjuste)
}

func TestSiatDocumentoAjuste_AnulacionDocumentoAjuste(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load()

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

	cuf := "CUF_FROM_INVOICE"

	req := models.DocumentoAjuste().NewAnulacionBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis(cuis).
		WithCufd(cufd).
		WithNit(nit).
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(codModalidad).
		WithTipoFacturaDocumento(3).
		WithCuf(cuf).
		WithCodigoMotivo(1).
		Build()

	serviceDocumentoAjuste := siatClient.DocumentoAjuste()
	runDocumentoAjusteTest(t, "AnulacionDocumentoAjuste", req, serviceDocumentoAjuste.AnulacionDocumentoAjuste)
}

func TestSiatDocumentoAjuste_ReversionAnulacionDocumentoAjuste(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load()

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
	// Generar CUF para Nota Conciliacion (Sector 29)
	cuf := "CUF_FROM_INVOICE"

	req := models.DocumentoAjuste().NewReversionAnulacionBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis(cuis).
		WithCufd(cufd).
		WithNit(nit).
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(codModalidad).
		WithTipoFacturaDocumento(3).
		WithCuf(cuf).
		Build()

	serviceDocumentoAjuste := siatClient.DocumentoAjuste()

	runDocumentoAjusteTest(t, "ReversionAnulacionDocumentoAjuste", req, serviceDocumentoAjuste.ReversionAnulacionDocumentoAjuste)
}

func TestSiatDocumentoAjuste_VerificacionEstadoDocumentoAjuste(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load()

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

	// Generar CUF para Nota Conciliacion (Sector 29)
	cuf := "CUF_FROM_INVOICE"

	req := models.DocumentoAjuste().NewVerificacionEstadoBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis(cuis).
		WithCufd(cufd).
		WithNit(nit).
		WithCodigoDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(codModalidad).
		WithTipoFacturaDocumento(3).
		WithCuf(cuf).
		Build()

	serviceDocumentoAjuste := siatClient.DocumentoAjuste()

	runDocumentoAjusteTest(t, "VerificacionEstadoDocumentoAjuste", req, serviceDocumentoAjuste.VerificacionEstadoDocumentoAjuste)
}
