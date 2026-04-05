package invoices_test

import (
	"context"
	"encoding/xml"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
)

func TestNotaConciliacionBuilder(t *testing.T) {
	fecha := time.Now()

	// Construir la cabecera
	cabecera := invoices.NewNotaConciliacionCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("Empresa Test").
		WithMunicipio("La Paz").
		WithNumeroNotaConciliacion(1).
		WithCuf("ABC123DEF").
		WithCufd("XYZ789").
		WithCodigoSucursal(0).
		WithDireccion("Av. Principal 123").
		WithFechaEmision(fecha).
		WithNombreRazonSocial(ptr("Cliente Test")).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5544332").
		WithCodigoCliente("CLI-001").
		WithNumeroFactura(100).
		WithNumeroAutorizacionCuf("FACT-ABC").
		WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
		WithMontoTotalOriginal(1000.00).
		WithMontoTotalConciliado(900.00).
		WithCreditoFiscalIva(0.00).
		WithDebitoFiscalIva(13.00). // 13% de la diferencia si aplica
		WithLeyenda("Ley Nro 453: El proveedor deberá entregar el bien...").
		WithUsuario("admin").
		Build()

	// Construir detalles
	detalleOri := invoices.NewNotaDetalleOriginalBuilder().
		WithActividadEconomica("Venta de productos").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Producto A").
		WithCantidad(10).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(1000).
		Build()

	detalleCon := invoices.NewNotaDetalleConciliacionBuilder().
		WithActividadEconomica("Venta de productos").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Producto A").
		WithMontoOriginal(1000).
		WithMontoFinal(900).
		WithMontoConciliado(100).
		Build()

	// Construir la nota completa
	nota := invoices.NewNotaConciliacionBuilder().
		WithCabecera(cabecera).
		AddDetalleOriginal(detalleOri).
		AddDetalleConciliacion(detalleCon).
		Build()

	output, err := xml.MarshalIndent(nota, "", "  ")
	if err != nil {
		t.Fatalf("Error al serializar: %v", err)
	}

	t.Logf("Nota XML:\n%s", string(output))
}

func TestNotaConciliacionIntegration_Computarizada(t *testing.T) {
	godotenv.Load("../../.env")
	godotenv.Load()

	token := os.Getenv("SIAT_TOKEN")
	url := os.Getenv("SIAT_URL")
	if token == "" || url == "" {
		t.Skip("Saltando test de integración: Token o URL no configurados")
	}

	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: token}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(url, client)

	serviceCodigos := siatClient.Codigos()
	service := siatClient.DocumentoAjuste()
	// codModalidad := siat.ModalidadComputarizada

	// 1. Obtener CUIS
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(siat.ModalidadComputarizada).
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
		WithCodigoModalidad(siat.ModalidadComputarizada).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()
	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
	cufdControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	fecha := time.Now()

	// 3. Generar CUF real para Nota Conciliacion (Sector 29) Tipo Documento (3)
	cuf, err := utils.GenerarCUF(nit, fecha, 0, siat.ModalidadComputarizada, siat.EmisionOnline, 3, 29, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	codigoPuntoVenta := 0

	// 4. Construir el documento XML de la Nota de Conciliación
	nota := invoices.NewNotaConciliacionBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(invoices.NewNotaConciliacionCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroNotaConciliacion(1).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Principal 123").
			WithCodigoPuntoVenta(&codigoPuntoVenta).
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-001").
			WithNumeroFactura(100).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
			WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
			WithMontoTotalOriginal(1000.00).
			WithMontoTotalConciliado(10.00).
			WithCreditoFiscalIva(0.00).
			WithDebitoFiscalIva(13.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalleOriginal(invoices.NewNotaDetalleOriginalBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Test Original").
			WithCantidad(1).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(100).
			Build()).
		AddDetalleConciliacion(invoices.NewNotaDetalleConciliacionBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Test Conciliado").
			WithMontoOriginal(100).
			WithMontoFinal(90).
			WithMontoConciliado(10).
			Build()).
		Build()

	xmlBytes, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	// 5. Preparar el envío (Gzip + Base64 + Hash)
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	// 6. Crear solicitud de recepción
	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(siat.ModalidadComputarizada).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(nit).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionDocumentoAjuste(context.Background(), config, req)
	if err != nil {
		t.Fatalf("Error en la comunicación con el SIAT: %v", err)
	}

	t.Logf("Respuesta SIAT: %+v", resp.Body.Content.RespuestaRecepcionFactura)
}

func TestNotaConciliacionIntegration_Electronica(t *testing.T) {
	godotenv.Load("../../.env")
	godotenv.Load()

	token := os.Getenv("SIAT_TOKEN")
	url := os.Getenv("SIAT_URL")
	if token == "" || url == "" {
		t.Skip("Saltando test de integración: Token o URL no configurados")
	}

	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: token}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(url, client)

	serviceCodigos := siatClient.Codigos()
	service := siatClient.DocumentoAjuste()
	// codModalidad := siat.ModalidadElectronica

	// 1. Obtener CUIS
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(siat.ModalidadElectronica).
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
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		Build()
	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
	cufdControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	fecha := time.Now()

	// 3. Generar CUF real para Nota Conciliacion (Sector 29) Tipo Documento (3)
	cuf, err := utils.GenerarCUF(nit, fecha, 0, siat.ModalidadElectronica, siat.EmisionOnline, 3, 29, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	codigoPuntoVenta := 0

	// 4. Construir el documento XML de la Nota de Conciliación
	nota := invoices.NewNotaConciliacionBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(invoices.NewNotaConciliacionCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroNotaConciliacion(1).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Principal 123").
			WithCodigoPuntoVenta(&codigoPuntoVenta).
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-001").
			WithNumeroFactura(100).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
			WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
			WithMontoTotalOriginal(1000.00).
			WithMontoTotalConciliado(10.00).
			WithCreditoFiscalIva(0.00).
			WithDebitoFiscalIva(13.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalleOriginal(invoices.NewNotaDetalleOriginalBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Test Original").
			WithCantidad(1).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(100).
			Build()).
		AddDetalleConciliacion(invoices.NewNotaDetalleConciliacionBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Test Conciliado").
			WithMontoOriginal(100).
			WithMontoFinal(90).
			WithMontoConciliado(10).
			Build()).
		Build()

	xmlBytes, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	xmlBytes, err = utils.SignXML(xmlBytes, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error al firmar: %v", err)
	}
	// 5. Preparar el envío (Gzip + Base64 + Hash)
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	// 6. Crear solicitud de recepción
	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithDocumentoSector(29).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(nit).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionDocumentoAjuste(context.Background(), config, req)
	if err != nil {
		t.Fatalf("Error en la comunicación con el SIAT: %v", err)
	}

	t.Logf("Respuesta SIAT: %+v", resp.Body.Content.RespuestaRecepcionFactura)
}

func ptr[T any](v T) *T {
	return &v
}
