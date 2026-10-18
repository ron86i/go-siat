package invoices_test

import (
	"context"
	"encoding/xml"
	"strings"
	"testing"
	"time"

	siat "github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// TestHidrocarburoNoAlcanzadaIehd_OrdenXSD verifica el Sector 38, idéntico al 19
// salvo por la ausencia de montoIehd y porcentajeIehd. XSD ver.11/03/2022.
func TestHidrocarburoNoAlcanzadaIehd_OrdenXSD(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 10, 30, 0, 0, time.UTC)
	nombre := "CLIENTE DE PRUEBA"

	cabecera := invoices.NewHidrocarburoNoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("EMPRESA DE PRUEBA S.R.L.").
		WithMunicipio("La Paz").
		WithTelefono(nil).
		WithNumeroFactura(1).
		WithCuf("CUF-DE-PRUEBA").
		WithCufd("CUFD-DE-PRUEBA").
		WithCodigoSucursal(0).
		WithDireccion("Av. Siempre Viva 742").
		WithCodigoPuntoVenta(utils.IntPtr(0)).
		WithFechaEmision(fecha).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithComplemento("").
		WithCodigoCliente("CLI-001").
		WithCiudad(nil).
		WithNombrePropietario(nil).
		WithNombreRepresentanteLegal(nil).
		WithCondicionPago(nil).
		WithPeriodoEntrega(nil).
		WithCodigoMetodoPago(1).
		WithNumeroTarjeta(nil).
		WithMontoAPagar(100).
		WithTotalPrecioNeto(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithDescuentoAdicional(nil).
		WithCodigoExcepcion(nil).
		WithLeyenda("Ley N 453").
		WithUsuario("tester").
		Build()

	detalle := invoices.NewHidrocarburoNoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("PROD-001").
		WithDescripcion("Gas Natural").
		WithCantidad(2).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(50).
		WithMontoDescuento(nil).
		WithSubTotal(100).
		Build()

	factura := invoices.NewHidrocarburoNoAlcanzadaIehdBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	raw, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error al serializar: %v", err)
	}
	got := string(raw)

	if !strings.HasPrefix(got, "<facturaElectronicaHidrocarburoNoAlcanzadaIehd") {
		t.Errorf("elemento raíz incorrecto:\n%.120s", got)
	}

	assertOrder(t, elementOrder(t, got, "cabecera"), []string{"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura", "cuf", "cufd", "codigoSucursal", "direccion", "codigoPuntoVenta", "fechaEmision", "nombreRazonSocial", "codigoTipoDocumentoIdentidad", "numeroDocumento", "complemento", "codigoCliente", "ciudad", "nombrePropietario", "nombreRepresentanteLegal", "condicionPago", "periodoEntrega", "codigoMetodoPago", "numeroTarjeta", "subTotalNeto", "totalPrecioNeto", "iva", "montoaPagar", "codigoMoneda", "tipoCambio", "montoTotalMoneda", "descuentoAdicional", "codigoExcepcion", "leyenda", "usuario", "codigoInterno", "codigoDocumentoSector"}, "cabecera sector 38")

	assertOrder(t, elementOrder(t, got, "detalle"), []string{"actividadEconomica", "codigoProductoSin", "codigoProducto", "descripcion", "cantidad", "unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal"}, "detalle sector 38")

	if !strings.Contains(got, "<codigoDocumentoSector>38</codigoDocumentoSector>") {
		t.Error("el código de documento sector no quedó fijado en 38")
	}
	if strings.Contains(got, "montoIehd") || strings.Contains(got, "porcentajeIehd") {
		t.Error("el sector 38 no debe emitir campos IEHD")
	}
}

// TestHidrocarburoNoAlcanzadaIehd_Computarizada verifica el cambio de elemento raíz en modalidad computarizada.
func TestHidrocarburoNoAlcanzadaIehd_Computarizada(t *testing.T) {
	raw, err := xml.Marshal(invoices.NewHidrocarburoNoAlcanzadaIehdBuilder().
		WithModalidad(siat.ModalidadComputarizada).Build())
	if err != nil {
		t.Fatalf("error al serializar: %v", err)
	}
	esperado := "<facturaComputarizadaHidrocarburoNoAlcanzadaIehd"
	if !strings.HasPrefix(string(raw), esperado) {
		t.Errorf("elemento raíz es %.60s, se esperaba %s", raw, esperado)
	}
}

func TestHidrocarburoNoAlcanzadaIehd_CompleteBuilder(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	telefono := "44123456"
	puntoVenta := 2
	nombre := "CLIENTE SECTOR 38 S.A."
	complemento := "2B"
	ciudad := "Tarija"
	propietario := "Ana López"
	repLegal := "Mario Silva"
	condicionPago := "Crédito"
	periodoEntrega := "Semanal"
	tarjeta := int64(4500987654321098)
	descuentoAdic := 10.00
	codigoExcepcion := 0
	montoDesc := 5.00

	cabecera := invoices.NewHidrocarburoNoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(8877665544).
		WithRazonSocialEmisor("EMPRESA SECTOR 38 S.R.L.").
		WithMunicipio("Tarija").
		WithTelefono(&telefono).
		WithNumeroFactura(250).
		WithCuf("CUF-SEC38").
		WithCufd("CUFD-SEC38").
		WithCodigoSucursal(1).
		WithDireccion("Av. Las Panteras 123").
		WithCodigoPuntoVenta(&puntoVenta).
		WithFechaEmision(fecha).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(5).
		WithNumeroDocumento("8877665").
		WithComplemento(complemento).
		WithCodigoCliente("CLI-SEC38").
		WithCiudad(&ciudad).
		WithNombrePropietario(&propietario).
		WithNombreRepresentanteLegal(&repLegal).
		WithCondicionPago(&condicionPago).
		WithPeriodoEntrega(&periodoEntrega).
		WithCodigoMetodoPago(1).
		WithNumeroTarjeta(&tarjeta).
		WithMontoAPagar(500.00).
		WithTotalPrecioNeto(495.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(500.00).
		WithDescuentoAdicional(&descuentoAdic).
		WithCodigoExcepcion(&codigoExcepcion).
		WithLeyenda("Ley N 453").
		WithUsuario("operador").
		WithCodigoDocumentoSector(38).
		Build()

	det1 := invoices.NewHidrocarburoNoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("352000").
		WithCodigoProductoSin(333444).
		WithCodigoProducto("GN-001").
		WithDescripcion("Gas Natural Vehicular").
		WithCantidad(100.0).
		WithUnidadMedida(2).
		WithPrecioNetoUnitario(3.0).
		WithMontoDescuento(&montoDesc).
		WithSubTotal(295.00).
		Build()

	det2 := invoices.NewHidrocarburoNoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("352000").
		WithCodigoProductoSin(333445).
		WithCodigoProducto("GN-002").
		WithDescripcion("Gas Natural Industrial").
		WithCantidad(50.0).
		WithUnidadMedida(2).
		WithPrecioNetoUnitario(4.1).
		WithMontoDescuento(nil).
		WithSubTotal(205.00).
		Build()

	factura := invoices.NewHidrocarburoNoAlcanzadaIehdBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(det1).
		AddDetalle(det2).
		Build()

	raw, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error al serializar XML: %v", err)
	}
	xmlStr := string(raw)

	if !strings.Contains(xmlStr, "<telefono>44123456</telefono>") {
		t.Errorf("telefono no se serializó correctamente: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<nombreRepresentanteLegal>Mario Silva</nombreRepresentanteLegal>") {
		t.Errorf("nombreRepresentanteLegal no se serializó correctamente: %s", xmlStr)
	}
	if strings.Contains(xmlStr, "montoIehd") {
		t.Errorf("sector 38 no debe incluir campo montoIehd: %s", xmlStr)
	}
}

func TestHidrocarburoNoAlcanzadaIehd_NilAndUnwrap(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)

	cabecera := invoices.NewHidrocarburoNoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(123456).
		WithRazonSocialEmisor("EMPRESA TEST").
		WithMunicipio("Sucre").
		WithTelefono(nil).
		WithNumeroFactura(20).
		WithCuf("CUF-TEST38").
		WithCufd("CUFD-TEST38").
		WithCodigoSucursal(0).
		WithDireccion("Calle 2").
		WithCodigoPuntoVenta(nil).
		WithFechaEmision(fecha).
		WithNombreRazonSocial(nil).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("445566").
		WithComplemento("").
		WithCodigoCliente("CLI-NIL38").
		WithCiudad(nil).
		WithNombrePropietario(nil).
		WithNombreRepresentanteLegal(nil).
		WithCondicionPago(nil).
		WithPeriodoEntrega(nil).
		WithCodigoMetodoPago(1).
		WithNumeroTarjeta(nil).
		WithMontoAPagar(150.0).
		WithTotalPrecioNeto(150.0).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(150.0).
		WithDescuentoAdicional(nil).
		WithCodigoExcepcion(nil).
		WithLeyenda("Ley N 453").
		WithUsuario("tester").
		Build()

	detalle := invoices.NewHidrocarburoNoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("352000").
		WithCodigoProductoSin(200).
		WithCodigoProducto("PROD-NIL38").
		WithDescripcion("Producto No IEHD").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(150.0).
		WithMontoDescuento(nil).
		WithSubTotal(150.0).
		Build()

	factura := invoices.NewHidrocarburoNoAlcanzadaIehdBuilder().
		WithCabecera(cabecera).
		WithCabecera(invoices.HidrocarburoNoAlcanzadaIehdCabecera{}). // test unwrap nil
		WithCabecera(cabecera).
		AddDetalle(invoices.HidrocarburoNoAlcanzadaIehdDetalle{}). // test unwrap nil
		AddDetalle(detalle).
		WithModalidad(88). // test default fallback modalidad
		Build()

	raw, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error al serializar: %v", err)
	}
	xmlStr := string(raw)

	if !strings.Contains(xmlStr, `<telefono xsi:nil="true">`) {
		t.Errorf("telefono nil debe serializarse con xsi:nil=\"true\": %s", xmlStr)
	}
}

func TestHidrocarburoNoAlcanzadaIehd_RecepcionFacturaBuilder(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)

	cab38 := invoices.NewHidrocarburoNoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("EMPRESA S.A.").
		WithMunicipio("La Paz").
		WithNumeroFactura(2).
		WithCuf("CUF38").
		WithCufd("CUFD38").
		WithCodigoSucursal(0).
		WithDireccion("Av. Principal").
		WithFechaEmision(fecha).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("12345").
		WithCodigoCliente("C38").
		WithCodigoMetodoPago(1).
		WithMontoAPagar(200).
		WithTotalPrecioNeto(200).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(200).
		WithLeyenda("Ley 453").
		WithUsuario("usr").
		Build()

	det38 := invoices.NewHidrocarburoNoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("352000").
		WithCodigoProductoSin(222).
		WithCodigoProducto("P38").
		WithDescripcion("Gas").
		WithCantidad(4).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(50).
		WithSubTotal(200).
		Build()

	factura38 := invoices.NewHidrocarburoNoAlcanzadaIehdBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(cab38).
		AddDetalle(det38).
		Build()

	builder38 := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(siat.ModalidadComputarizada).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(38).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd("CUFD38").
		WithCuis("CUIS38").
		WithTipoFacturaDocumento(1).
		WithFechaEnvio(fecha)

	if err := builder38.WithFactura(factura38, nil); err != nil {
		t.Fatalf("error en WithFactura sector 38: %v", err)
	}
	req38 := builder38.Build()
	raw38, err := xml.Marshal(req38)
	if err != nil {
		t.Fatalf("error al serializar solicitud sector 38: %v", err)
	}
	if !strings.Contains(string(raw38), "<archivo>") {
		t.Error("el campo archivo en la solicitud sector 38 no debe estar vacío")
	}
}

func TestHidrocarburoNoAlcanzadaIehd_ElectronicaIntegration(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 38, 1, 0, cufdControl)

	nombreRazonSocial := "COMPRADOR NO ALCANZADA IEHD"

	cabecera := invoices.NewHidrocarburoNoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("GASBOL S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. PRADO").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-NOIEHD-01").
		WithCodigoMetodoPago(1).
		WithMontoAPagar(100.00).
		WithTotalPrecioNeto(100.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(100.00).
		WithLeyenda("Leyenda No IEHD").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewHidrocarburoNoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("352000").
		WithCodigoProductoSin(54321).
		WithCodigoProducto("NOIEHD-001").
		WithDescripcion("GAS NATURAL").
		WithCantidad(10.0).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(10.0).
		WithSubTotal(100.0).
		Build()

	factura := invoices.NewHidrocarburoNoAlcanzadaIehdBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(38).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithFechaEnvio(fechaEmision)

	err := builderReq.WithFactura(factura, tc.Client.Config())
	if err != nil {
		t.Fatalf("error al preparar factura: %v", err)
	}

	req := builderReq.Build()

	resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
}
