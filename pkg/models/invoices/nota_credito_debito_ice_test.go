package invoices_test

import (
	"context"
	"encoding/xml"
	"testing"
	"time"

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
	tc := setupTestContext(t, siat.ModalidadComputarizada)

	service := tc.Client.DocumentoAjuste()
	// codModalidad := siat.ModalidadComputarizada

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fecha := time.Now()

	// 3. Generar CUF
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, 0, tc.Modalidad, siat.EmisionOnline, 3, 48, 2, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoIceBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoIceCabeceraBuilder().
			WithNitEmisor(tc.Nit).
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
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(48).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(tc.Nit).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionDocumentoAjuste(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("Error en la comunicación con el SIAT: %v", err)
	}

	t.Logf("Respuesta SIAT: %+v", resp.Body.Content.RespuestaRecepcionFactura)
}

func TestNotaCreditoDebitoIceIntegration_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)

	service := tc.Client.DocumentoAjuste()
	// codModalidad := siat.ModalidadElectronica

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fecha := time.Now()

	// 3. Generar CUF
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, 0, tc.Modalidad, siat.EmisionOnline, 3, 48, 2, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoIceBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoIceCabeceraBuilder().
			WithNitEmisor(tc.Nit).
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
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(48).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(tc.Nit).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionDocumentoAjuste(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("Error en la comunicación con el SIAT: %v", err)
	}

	t.Logf("Respuesta SIAT: %+v", resp.Body.Content.RespuestaRecepcionFactura)
}
