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

func TestNotaCreditoDebitoIceBuilder(t *testing.T) {
	fecha := time.Now()

	cabecera := invoices.NewNotaCreditoDebitoIceCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("Empresa Test ICE").
		WithMunicipio("Santa Cruz").
		WithNumeroNotaCreditoDebito(2).
		WithCuf("ABC123DEF").
		WithCufd("XYZ789").
		WithCodigoSucursal(0).
		WithDireccion("Av. Banzer 123").
		WithFechaEmision(fecha).
		WithNombreRazonSocial(ptr("Cliente Test ICE")).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5544332").
		WithCodigoCliente("CLI-002").
		WithNumeroFactura(101).
		WithNumeroAutorizacionCuf("FACT-ABC-ICE").
		WithFechaEmisionFactura(fecha.Add(-48 * time.Hour)).
		WithMontoTotalOriginal(2000.00).
		WithMontoTotalDevuelto(1000.00).
		WithMontoEfectivoCreditoDebito(130.00).
		WithLeyenda("Ley Nro 453 ICE").
		WithUsuario("admin-ice").
		Build()

	precioNeto := 90.0
	alicuotaIva := 13.0
	detalle1 := invoices.NewNotaDetalleCreditoDebitoIceBuilder().
		WithNroItem(1).
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(1234).
		WithCodigoProducto("P-ICE-001").
		WithDescripcion("Producto ICE A").
		WithCantidad(20).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(2000).
		WithMarcaIce(1).
		WithAlicuotaIva(&alicuotaIva).
		WithPrecioNetoVentaIce(&precioNeto).
		WithCodigoDetalleTransaccion(1).
		Build()

	detalle2 := invoices.NewNotaDetalleCreditoDebitoIceBuilder().
		WithNroItem(2).
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(1234).
		WithCodigoProducto("P-ICE-002").
		WithDescripcion("Producto ICE B devuelto").
		WithCantidad(10).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(1000).
		WithMarcaIce(1).
		WithAlicuotaIva(&alicuotaIva).
		WithPrecioNetoVentaIce(&precioNeto).
		WithCodigoDetalleTransaccion(2).
		Build()

	nota := invoices.NewNotaCreditoDebitoIceBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle1).
		AddDetalle(detalle2).
		Build()

	output, err := xml.MarshalIndent(nota, "", "  ")
	if err != nil {
		t.Fatalf("Error al serializar: %v", err)
	}

	t.Logf("Nota XML ICE:\n%s", string(output))
}

func TestNotaCreditoDebitoIceIntegration_Computarizada(t *testing.T) {
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

	// 3. Generar CUF
	cuf, err := utils.GenerarCUF(nit, fecha, 0, siat.ModalidadComputarizada, siat.EmisionOnline, 3, 48, 2, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoIceBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(invoices.NewNotaCreditoDebitoIceCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Empresa Test ICE").
			WithMunicipio("Santa Cruz").
			WithNumeroNotaCreditoDebito(2).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Banzer 123").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test ICE")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-002").
			WithNumeroFactura(101).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF-ICE").
			WithFechaEmisionFactura(fecha.Add(-48 * time.Hour)).
			WithMontoTotalOriginal(2000.00).
			WithMontoTotalDevuelto(1000.00).
			WithMontoEfectivoCreditoDebito(130.00).
			WithLeyenda("Ley Nro 453 ICE").
			WithUsuario("admin-ice").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoIceBuilder().
			WithNroItem(1).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-ICE-001").
			WithDescripcion("Producto ICE A").
			WithCantidad(20).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(2000).
			WithMarcaIce(1).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoIceBuilder().
			WithNroItem(2).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-ICE-002").
			WithDescripcion("Producto ICE B devuelto").
			WithCantidad(10).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1000).
			WithMarcaIce(1).
			WithCodigoDetalleTransaccion(2).
			Build()).
		Build()

	xmlBytes, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	// 5. Preparar envío
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	// 6. Solicitud de recepción
	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithDocumentoSector(48).
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

func TestNotaCreditoDebitoIceIntegration_Electronica(t *testing.T) {
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

	// 3. Generar CUF
	cuf, err := utils.GenerarCUF(nit, fecha, 0, siat.ModalidadElectronica, siat.EmisionOnline, 3, 48, 2, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoIceBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(invoices.NewNotaCreditoDebitoIceCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Empresa Test ICE").
			WithMunicipio("Santa Cruz").
			WithNumeroNotaCreditoDebito(2).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Banzer 123").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test ICE")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-002").
			WithNumeroFactura(101).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF-ICE").
			WithFechaEmisionFactura(fecha.Add(-48 * time.Hour)).
			WithMontoTotalOriginal(2000.00).
			WithMontoTotalDevuelto(1000.00).
			WithMontoEfectivoCreditoDebito(130.00).
			WithLeyenda("Ley Nro 453 ICE").
			WithUsuario("admin-ice").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoIceBuilder().
			WithNroItem(1).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-ICE-001").
			WithDescripcion("Producto ICE A").
			WithCantidad(20).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(2000).
			WithMarcaIce(1).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoIceBuilder().
			WithNroItem(2).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-ICE-002").
			WithDescripcion("Producto ICE B devuelto").
			WithCantidad(10).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1000).
			WithMarcaIce(1).
			WithCodigoDetalleTransaccion(2).
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

	// 5. Preparar envío
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	// 6. Solicitud de recepción
	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithDocumentoSector(48).
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
