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

func TestNotaFiscalCreditoDebitoBuilder(t *testing.T) {
	fecha := time.Now()

	cabecera := invoices.NewNotaFiscalCreditoDebitoCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("Empresa Test Fiscal").
		WithMunicipio("Cochabamba").
		WithNumeroNotaCreditoDebito(3).
		WithCuf("ABC123DEF").
		WithCufd("XYZ789").
		WithCodigoSucursal(0).
		WithDireccion("Av. Blanco Galindo").
		WithFechaEmision(fecha).
		WithNombreRazonSocial(ptr("Cliente Test Fiscal")).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5544332").
		WithCodigoCliente("CLI-003").
		WithNumeroFactura(105).
		WithNumeroAutorizacionCuf("FACT-ABC-FISCAL").
		WithFechaEmisionFactura(fecha.Add(-72 * time.Hour)).
		WithMontoTotalOriginal(2000.00).
		WithMontoTotalDevuelto(1500.00).
		WithMontoEfectivoCreditoDebito(195.00).
		WithLeyenda("Ley Nro 453 Normal").
		WithUsuario("admin-fiscal").
		Build()

	detalle1 := invoices.NewNotaDetalleFiscalCreditoDebitoBuilder().
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(5566).
		WithCodigoProducto("P-F-001").
		WithDescripcion("Producto F A").
		WithCantidad(20).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(2000).
		WithCodigoDetalleTransaccion(1).
		Build()

	detalle2 := invoices.NewNotaDetalleFiscalCreditoDebitoBuilder().
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(5566).
		WithCodigoProducto("P-F-002").
		WithDescripcion("Producto F B").
		WithCantidad(15).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(1500).
		WithCodigoDetalleTransaccion(2).
		Build()

	nota := invoices.NewNotaFiscalCreditoDebitoBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle1).
		AddDetalle(detalle2).
		Build()

	output, err := xml.MarshalIndent(nota, "", "  ")
	if err != nil {
		t.Fatalf("Error al serializar: %v", err)
	}

	t.Logf("Nota XML Fiscal:\n%s", string(output))
}

func TestNotaFiscalCreditoDebitoIntegration_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)

	service := tc.Client.DocumentoAjuste()
	// codModalidad := siat.ModalidadComputarizada

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fecha := time.Now()

	// 3. Generar CUF
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, 0, tc.Modalidad, siat.EmisionOnline, 3, 24, 3, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaFiscalCreditoDebitoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaFiscalCreditoDebitoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Empresa Test Fiscal").
			WithMunicipio("Cochabamba").
			WithNumeroNotaCreditoDebito(3).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Blanco Galindo").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test Fiscal")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-003").
			WithNumeroFactura(105).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF-FISCAL").
			WithFechaEmisionFactura(fecha.Add(-72 * time.Hour)).
			WithMontoTotalOriginal(2000.00).
			WithMontoTotalDevuelto(1500.00).
			WithMontoEfectivoCreditoDebito(195.00).
			WithLeyenda("Ley Nro 453 Normal").
			WithUsuario("admin-fiscal").
			Build()).
		AddDetalle(invoices.NewNotaDetalleFiscalCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-F-001").
			WithDescripcion("Producto F A").
			WithCantidad(20).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(2000).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleFiscalCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-F-002").
			WithDescripcion("Producto F B").
			WithCantidad(15).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1500).
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
		WithCodigoDocumentoSector(24).
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

func TestNotaFiscalCreditoDebitoIntegration_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)

	service := tc.Client.DocumentoAjuste()
	// codModalidad := siat.ModalidadElectronica

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fecha := time.Now()

	// 3. Generar CUF
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, 0, tc.Modalidad, siat.EmisionOnline, 3, 24, 3, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaFiscalCreditoDebitoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaFiscalCreditoDebitoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Empresa Test Fiscal").
			WithMunicipio("Cochabamba").
			WithNumeroNotaCreditoDebito(3).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Blanco Galindo").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test Fiscal")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-003").
			WithNumeroFactura(105).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF-FISCAL").
			WithFechaEmisionFactura(fecha.Add(-72 * time.Hour)).
			WithMontoTotalOriginal(2000.00).
			WithMontoTotalDevuelto(1500.00).
			WithMontoEfectivoCreditoDebito(195.00).
			WithLeyenda("Ley Nro 453 Normal").
			WithUsuario("admin-fiscal").
			Build()).
		AddDetalle(invoices.NewNotaDetalleFiscalCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-F-001").
			WithDescripcion("Producto F A").
			WithCantidad(20).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(2000).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleFiscalCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P-F-002").
			WithDescripcion("Producto F B").
			WithCantidad(15).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1500).
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
		WithCodigoDocumentoSector(24).
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
