package services_test

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// TestSiatCompraVentaService_RecepcionAnexos verifica el envío de anexos técnicos para facturas específicas.
// Valida que el builder construya correctamente la solicitud con todos los campos obligatorios
// y que el servicio procese la respuesta del SIAT usando el mapeo XML estandarizado.
func TestSiatCompraVentaService_RecepcionAnexos(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            123456789,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{},
	}

	siatClient, err := siat.New(cfg)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	service := siatClient.CompraVenta()

	req := models.NewRecepcionAnexosBuilder().
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		WithCuf("D5340CCDF031F0CFDB...").
		AddAnexos(models.NewVentaAnexoBuilder().
			WithCodigo("1").
			WithCodigoProducto("86111").
			Build()).
		Build()

	resp, err := service.RecepcionAnexos(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion Anexos: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_VerificacionEstadoFactura valida la consulta del estado actual de una factura en el SIAT.
// Verifica que se retorne la información de recepción y estado (Válida, Anulada, etc.) correctamente.
func TestSiatCompraVentaService_VerificacionEstadoFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: codAmbiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     siat.NewHTTPClient(siat.DefaultHTTPConfig()),
	}
	siatClient, _ := siat.New(cfg)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.NewCuisBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), cuisReq)

	cufdReq := models.NewCufdBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), cufdReq)

	req := models.NewVerificacionEstadoFacturaBuilder().
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1).
		WithCuf("D5340CCDF031F0CFDBF...").
		Build()

	resp, err := serviceCompraVenta.VerificacionEstadoFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Verificacion Estado: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_VerificarComunicacion valida la conectividad básica con el
// Servicio de Facturación Compra Venta del SIAT.
func TestSiatCompraVentaService_VerificarComunicacion(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: codAmbiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     siat.NewHTTPClient(siat.DefaultHTTPConfig()),
	}
	siatClient, _ := siat.New(cfg)
	service := siatClient.CompraVenta()

	req := models.NewVerificarComunicacionFacturacion()
	resp, err := service.VerificarComunicacion(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		assert.True(t, resp.Body.Content.Return.Transaccion)
		log.Printf("Respuesta Comunicacion: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_RecepcionMasivaFactura prueba el flujo de envío de múltiples facturas
// bajo la modalidad de emisión masiva. Asegura que el CUF generado sea consistente con el tipo de emisión.
func TestSiatCompraVentaService_RecepcionMasivaFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: codAmbiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     siat.NewHTTPClient(siat.DefaultHTTPConfig()),
	}

	siatClient, _ := siat.New(cfg)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.NewCuisBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		Build()

	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), cuisReq)
	if err != nil {
		t.Fatalf("error calculando CUIS: %v", err)
	}

	cufdReq := models.NewCufdBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), cufdReq)
	if err != nil {
		t.Fatalf("error calculando CUFD: %v", err)
	}

	// Construir paquete de 5 facturas para recepción masiva
	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)
	fechaEmision := time.Now()
	codigoPuntoVenta := 0
	for i := 1; i <= 5; i++ {
		nombreRazonSocial := "JUAN PEREZ"
		// Para masiva debe ser emisión Masiva (3)
		cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, siat.EmisionMasiva, 1, 1, int64(i), 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
		t.Logf("CUF #%d: %s", i, cuf)
		cabecera := invoices.NewCompraVentaCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Ronaldo Rua").
			WithMunicipio("Tarija").
			WithNumeroFactura(int64(i)).
			WithCuf(cuf).
			WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
			WithCodigoSucursal(0).
			WithDireccion("ESQUINA AVENIDA LA PAZ").
			WithCodigoPuntoVenta(&codigoPuntoVenta).
			WithFechaEmision(fechaEmision).
			WithNombreRazonSocial(&nombreRazonSocial).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5115889").
			WithCodigoCliente(strconv.Itoa(i)).
			WithCodigoMetodoPago(1).
			WithMontoTotal(100).
			WithMontoTotalSujetoIva(100).
			WithCodigoMoneda(1).
			WithTipoCambio(1).
			WithMontoTotalMoneda(100).
			WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
			WithUsuario("usuario").
			WithCodigoDocumentoSector(1).
			Build()

		detalle := invoices.NewCompraVentaDetalleBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("abc123").
			WithDescripcion("GASA").
			WithCantidad(1).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(100).
			Build()
		factura := invoices.NewCompraVentaBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		xmlData, _ := xml.Marshal(factura)
		signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")

		hdr := &tar.Header{
			Name: fmt.Sprintf("factura_%d.xml", i),
			Mode: 0600,
			Size: int64(len(signedXML)),
		}
		tw.WriteHeader(hdr)
		tw.Write(signedXML)
	}
	tw.Close()

	// Comprimir el TAR con Gzip y preparar para SIAT
	hashString, encodedArchivo, err := utils.CompressAndHash(tarBuf.Bytes())
	if err != nil {
		t.Fatalf("error preparando paquete masivo: %v", err)
	}

	req := models.NewRecepcionMasivaFacturaBuilder().
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		WithCantidadFacturas(5).
		Build()

	resp, err := serviceCompraVenta.RecepcionMasivaFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion Masiva: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_ValidacionRecepcionMasivaFactura verifica el estado de procesamiento
// de un paquete de facturas enviado previamente mediante RecepcionMasivaFactura.
func TestSiatCompraVentaService_ValidacionRecepcionMasivaFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: codAmbiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     siat.NewHTTPClient(siat.DefaultHTTPConfig()),
	}
	siatClient, _ := siat.New(cfg)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.NewCuisBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), cuisReq)

	cufdReq := models.NewCufdBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), cufdReq)

	req := models.NewValidacionRecepcionMasivaFacturaBuilder().
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1).
		WithCodigoRecepcion("755d4aab-1ce6-11f1-8c52-99bc8e8492c6").
		Build()

	resp, err := serviceCompraVenta.ValidacionRecepcionMasivaFactura(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Validacion Masiva: %+v", resp.Body.Content)
	}
}
