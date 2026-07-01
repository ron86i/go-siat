package services_test

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/adapter/services"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func getBoletoAereoSetup(t *testing.T) (*siat.SiatServices, string, string, string, int64, int) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	codModalidad := siat.ModalidadElectronica

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

func TestSiatBoletoAereoService_NewSiatBoletoAereoService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		service, err := services.NewSiatBoletoAereoService("https://example.com", nil, siat.Config{})
		assert.NoError(t, err)
		assert.NotNil(t, service)
	})

	t.Run("Empty URL", func(t *testing.T) {
		service, err := services.NewSiatBoletoAereoService("", nil, siat.Config{})
		assert.Error(t, err)
		assert.Nil(t, service)
	})
}

func TestSiatBoletoAereoService_VerificarComunicacion(t *testing.T) {
	siatClient, _, _, _, _, _ := getBoletoAereoSetup(t)
	service := siatClient.BoletoAereo()

	req := models.NewVerificarComunicacionFacturacion()
	resp, err := service.VerificarComunicacion(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		assert.True(t, resp.Body.Content.Return.Transaccion)
		log.Printf("Respuesta Comunicacion Boleto Aereo: %+v", resp.Body.Content)
	}
}

func TestSiatBoletoAereoService_RecepcionMasivaFactura(t *testing.T) {
	siatClient, cuis, cufd, cufdControl, nit, codModalidad := getBoletoAereoSetup(t)
	serviceBoleto := siatClient.BoletoAereo()

	fechaEmision := time.Now()
	// Generar CUF (Sector 30)
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 3, 4, 30, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombre := "JUAN PEREZ"
	cabecera := invoices.NewBoletoAereoCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("AEROLINEAS TEST S.A.").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("Aropuerto Internacional El Alto").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithNombrePasajero("JUAN PEREZ").
		WithCodigoIataLineaAerea(900).
		WithCodigoOrigenServicio("LPB").
		WithCodigoMetodoPago(1).
		WithMontoTarifa(1000).
		WithMontoTotal(1000).
		WithMontoTotalSujetoIva(1000).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(1000).
		WithCodigoTipoTransaccion("I").
		WithLeyenda("Ley 453...").
		WithUsuario("operador").
		Build()

	factura := invoices.NewBoletoAereoBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		Build()

	xmlData, _ := xml.Marshal(factura)
	signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")

	// Para Boleto Aéreo se requiere empaquetado TAR.GZ
	tarGz, err := utils.CreateTarGz(map[string][]byte{"factura.xml": signedXML})
	if err != nil {
		t.Fatalf("error creando tar.gz: %v", err)
	}

	hashSum := sha256.Sum256(tarGz)
	hashString := fmt.Sprintf("%x", hashSum)
	encodedArchivo := base64.StdEncoding.EncodeToString(tarGz)

	req := models.NewRecepcionMasivaFacturaBuilder().
		WithCodigoDocumentoSector(30).
		WithCodigoEmision(3).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(4).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		WithCantidadFacturas(1).
		Build()

	resp, err := serviceBoleto.RecepcionMasivaFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion Masiva Boleto Aereo: %+v", resp.Body.Content)
	}
}

func TestSiatBoletoAereoService_ValidacionRecepcionMasivaFactura(t *testing.T) {
	siatClient, cuis, cufd, _, _, codModalidad := getBoletoAereoSetup(t)
	serviceBoleto := siatClient.BoletoAereo()

	req := models.NewValidacionRecepcionMasivaFacturaBuilder().
		WithCodigoDocumentoSector(30).
		WithCodigoEmision(3).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(4).
		WithCodigoRecepcion("aaa617de-350b-11f1-bc7c-b31977654538").
		Build()

	resp, err := serviceBoleto.ValidacionRecepcionMasivaFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Validación Masiva Boleto Aereo: %+v", resp.Body.Content)
	}
}

func TestSiatBoletoAereoService_VerificacionEstadoFactura(t *testing.T) {
	siatClient, cuis, cufd, _, _, codModalidad := getBoletoAereoSetup(t)
	serviceBoleto := siatClient.BoletoAereo()

	req := models.NewVerificacionEstadoFacturaBuilder().
		WithCodigoDocumentoSector(30).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(4).
		WithCuf("ABC123FAKE").
		Build()

	resp, err := serviceBoleto.VerificacionEstadoFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Verificacion Estado Boleto Aereo: %+v", resp.Body.Content)
	}
}

func TestSiatBoletoAereoService_AnulacionFactura(t *testing.T) {
	siatClient, cuis, cufd, _, _, codModalidad := getBoletoAereoSetup(t)
	serviceBoleto := siatClient.BoletoAereo()

	req := models.NewAnulacionFacturaBuilder().
		WithCodigoDocumentoSector(30).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(4).
		WithCuf("ABC123FAKE").
		WithCodigoMotivo(1).
		Build()

	resp, err := serviceBoleto.AnulacionFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Anulación Boleto Aereo: %+v", resp.Body.Content)
	}
}

func TestSiatBoletoAereoService_ReversionAnulacionFactura(t *testing.T) {
	siatClient, cuis, cufd, _, _, codModalidad := getBoletoAereoSetup(t)
	serviceBoleto := siatClient.BoletoAereo()

	req := models.NewReversionAnulacionFacturaBuilder().
		WithCodigoDocumentoSector(30).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(4).
		WithCuf("ABC123FAKE").
		Build()

	resp, err := serviceBoleto.ReversionAnulacionFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Reversión Anulación Boleto Aereo: %+v", resp.Body.Content)
	}
}
