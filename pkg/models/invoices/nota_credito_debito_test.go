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

func TestNotaCreditoDebitoBuilder(t *testing.T) {
	fecha := time.Now()

	cabecera := invoices.NewNotaCreditoDebitoCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("Empresa Test").
		WithMunicipio("La Paz").
		WithNumeroNotaCreditoDebito(1).
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
		WithMontoTotalDevuelto(900.00).
		WithMontoEfectivoCreditoDebito(100.00).
		WithLeyenda("Ley Nro 453").
		WithUsuario("admin").
		Build()

	detalle1 := invoices.NewNotaDetalleCreditoDebitoBuilder().
		WithNroItem(1).
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Producto A").
		WithCantidad(10).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(1000).
		WithCodigoDetalleTransaccion(1).
		Build()

	detalle2 := invoices.NewNotaDetalleCreditoDebitoBuilder().
		WithNroItem(2).
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P002").
		WithDescripcion("Producto B devuelto").
		WithCantidad(5).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(500).
		WithCodigoDetalleTransaccion(2).
		Build()

	nota := invoices.NewNotaCreditoDebitoBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle1).
		AddDetalle(detalle2).
		Build()

	output, err := xml.MarshalIndent(nota, "", "  ")
	if err != nil {
		t.Fatalf("Error al serializar: %v", err)
	}

	t.Logf("Nota XML:\n%s", string(output))
}

func TestNotaCreditoDebitoIntegration_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)

	service := tc.Client.DocumentoAjuste()
	// codModalidad := siat.ModalidadComputarizada

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fecha := time.Now()

	// 3. Generar CUF real para Nota (Sector 47) Tipo Documento (3)
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, 0, tc.Modalidad, siat.EmisionOnline, 3, 47, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroNotaCreditoDebito(1).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Principal 123").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-001").
			WithNumeroFactura(100).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
			WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
			WithMontoTotalOriginal(1000.00).
			WithMontoTotalDevuelto(500.00).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithNroItem(1).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Producto A").
			WithCantidad(10).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1000).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithNroItem(2).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P002").
			WithDescripcion("Producto B devuelto").
			WithCantidad(5).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(500).
			WithCodigoDetalleTransaccion(2).
			Build()).
		Build()

	xmlBytes, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	// 5. Preparar el envío
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	// 6. Crear solicitud de recepción
	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(47).
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

func TestNotaCreditoDebitoIntegration_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)

	service := tc.Client.DocumentoAjuste()
	// codModalidad := siat.ModalidadElectronica

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fecha := time.Now()

	// 3. Generar CUF real para Nota (Sector 47) Tipo Documento (3)
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, 0, tc.Modalidad, siat.EmisionOnline, 3, 47, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroNotaCreditoDebito(1).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Principal 123").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-001").
			WithNumeroFactura(100).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
			WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
			WithMontoTotalOriginal(1000.00).
			WithMontoTotalDevuelto(500.00).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithNroItem(1).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Producto A").
			WithCantidad(10).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1000).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithNroItem(2).
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P002").
			WithDescripcion("Producto B devuelto").
			WithCantidad(5).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(500).
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

	// 5. Preparar el envío
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	// 6. Crear solicitud de recepción
	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(47).
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
