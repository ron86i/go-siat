package invoices_test

import (
	"context"
	"encoding/xml"
	"regexp"
	"strings"
	"testing"
	"time"

	siat "github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// elementOrder extrae los nombres de los elementos hijos directos de <cabecera> o <detalle>
// en el orden en que aparecen en el XML serializado.
func elementOrder(t *testing.T, xmlStr, parent string) []string {
	t.Helper()

	start := strings.Index(xmlStr, "<"+parent+">")
	if start < 0 {
		t.Fatalf("no se encontró <%s> en el XML", parent)
	}
	end := strings.Index(xmlStr[start:], "</"+parent+">")
	if end < 0 {
		t.Fatalf("no se encontró </%s> en el XML", parent)
	}
	block := xmlStr[start+len(parent)+2 : start+end]

	re := regexp.MustCompile(`<([a-zA-Z][a-zA-Z0-9]*)[ />]`)
	var names []string
	for _, m := range re.FindAllStringSubmatch(block, -1) {
		names = append(names, m[1])
	}
	return names
}

func assertOrder(t *testing.T, got, want []string, label string) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("%s: %d elementos, se esperaban %d\n  obtenido: %v\n  esperado: %v",
			label, len(got), len(want), got, want)
		return
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("%s: posición %d es %q, se esperaba %q", label, i, got[i], want[i])
		}
	}
}

// TestHidrocarburoAlcanzadaIehd_OrdenXSD verifica que la factura del Sector 19 se serialice
// con los elementos exactamente en el orden del XSD ver.23/08/2021. El SIAT rechaza con
// código 939 cualquier desviación del orden.
func TestHidrocarburoAlcanzadaIehd_OrdenXSD(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 10, 30, 0, 0, time.UTC)
	nombre := "CLIENTE DE PRUEBA"
	montoIehd := 12.50
	porcentaje := 5.0

	cabecera := invoices.NewHidrocarburoAlcanzadaIehdCabeceraBuilder().
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
		WithMontoIehd(&montoIehd).
		WithTotalPrecioNeto(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithDescuentoAdicional(nil).
		WithCodigoExcepcion(nil).
		WithLeyenda("Ley N 453").
		WithUsuario("tester").
		Build()

	detalle := invoices.NewHidrocarburoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("PROD-001").
		WithDescripcion("Diesel Oil").
		WithCantidad(2).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(50).
		WithMontoDescuento(nil).
		WithSubTotal(100).
		WithPorcentajeIehd(&porcentaje).
		Build()

	factura := invoices.NewHidrocarburoAlcanzadaIehdBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	raw, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error al serializar: %v", err)
	}
	got := string(raw)

	if !strings.HasPrefix(got, "<facturaElectronicaHidrocarburoAlcanzadaIehd") {
		t.Errorf("elemento raíz incorrecto:\n%.120s", got)
	}

	assertOrder(t, elementOrder(t, got, "cabecera"), []string{"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura", "cuf", "cufd", "codigoSucursal", "direccion", "codigoPuntoVenta", "fechaEmision", "nombreRazonSocial", "codigoTipoDocumentoIdentidad", "numeroDocumento", "complemento", "codigoCliente", "ciudad", "nombrePropietario", "nombreRepresentanteLegal", "condicionPago", "periodoEntrega", "codigoMetodoPago", "numeroTarjeta", "montoIehd", "subTotalNeto", "totalPrecioNeto", "iva", "ivad", "montoaPagar", "codigoMoneda", "tipoCambio", "montoTotalMoneda", "descuentoAdicional", "codigoExcepcion", "leyenda", "usuario", "codigoInterno", "codigoDocumentoSector"}, "cabecera sector 19")

	assertOrder(t, elementOrder(t, got, "detalle"), []string{"actividadEconomica", "codigoProductoSin", "codigoProducto", "descripcion", "cantidad", "unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal", "porcentajeIehd"}, "detalle sector 19")

	if !strings.Contains(got, "<codigoDocumentoSector>19</codigoDocumentoSector>") {
		t.Error("el código de documento sector no quedó fijado en 19")
	}
	if !strings.Contains(got, `<telefono xsi:nil="true">`) {
		t.Error("un campo nillable sin valor debe emitirse con xsi:nil=\"true\"")
	}
}

// TestHidrocarburoAlcanzadaIehd_Computarizada verifica el cambio de elemento raíz en modalidad computarizada.
func TestHidrocarburoAlcanzadaIehd_Computarizada(t *testing.T) {
	raw, err := xml.Marshal(invoices.NewHidrocarburoAlcanzadaIehdBuilder().
		WithModalidad(siat.ModalidadComputarizada).Build())
	if err != nil {
		t.Fatalf("error al serializar: %v", err)
	}
	esperado := "<facturaComputarizadaHidrocarburoAlcanzadaIehd"
	if !strings.HasPrefix(string(raw), esperado) {
		t.Errorf("elemento raíz es %.60s, se esperaba %s", raw, esperado)
	}
}

func TestHidrocarburoAlcanzadaIehd_CompleteBuilder(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)
	telefono := "22123456"
	puntoVenta := 1
	nombre := "EMPRESA COMPRADORA S.A."
	complemento := "1A"
	ciudad := "La Paz"
	propietario := "Juan Pérez"
	repLegal := "Carlos Gómez"
	condicionPago := "Contado"
	periodoEntrega := "Mensual"
	tarjeta := int64(4500123456789012)
	montoIehd := 15.50
	descuentoAdic := 5.00
	codigoExcepcion := 1
	montoDesc := 2.50
	porcentajeIehd := 7.50

	cabecera := invoices.NewHidrocarburoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(9900112233).
		WithRazonSocialEmisor("YPFB CORPORACION").
		WithMunicipio("Santa Cruz").
		WithTelefono(&telefono).
		WithNumeroFactura(500).
		WithCuf("CUF1234567890").
		WithCufd("CUFD1234567890").
		WithCodigoSucursal(0).
		WithDireccion("Av. Banzer Km 3").
		WithCodigoPuntoVenta(&puntoVenta).
		WithFechaEmision(fecha).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(5).
		WithNumeroDocumento("1029384").
		WithComplemento(complemento).
		WithCodigoCliente("CLI-100").
		WithCiudad(&ciudad).
		WithNombrePropietario(&propietario).
		WithNombreRepresentanteLegal(&repLegal).
		WithCondicionPago(&condicionPago).
		WithPeriodoEntrega(&periodoEntrega).
		WithCodigoMetodoPago(2).
		WithNumeroTarjeta(&tarjeta).
		WithMontoAPagar(200.00).
		WithMontoIehd(&montoIehd).
		WithTotalPrecioNeto(184.50).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(200.00).
		WithDescuentoAdicional(&descuentoAdic).
		WithCodigoExcepcion(&codigoExcepcion).
		WithLeyenda("Ley N 453").
		WithUsuario("admin").
		WithCodigoDocumentoSector(19).
		Build()

	det1 := invoices.NewHidrocarburoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("061000").
		WithCodigoProductoSin(111222).
		WithCodigoProducto("GAS-001").
		WithDescripcion("Gasolina Especial").
		WithCantidad(10.0).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(15.0).
		WithMontoDescuento(&montoDesc).
		WithSubTotal(147.50).
		WithPorcentajeIehd(&porcentajeIehd).
		Build()

	det2 := invoices.NewHidrocarburoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("061000").
		WithCodigoProductoSin(111223).
		WithCodigoProducto("DIE-001").
		WithDescripcion("Diesel Premium").
		WithCantidad(5.0).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(10.5).
		WithMontoDescuento(nil).
		WithSubTotal(52.50).
		WithPorcentajeIehd(nil).
		Build()

	factura := invoices.NewHidrocarburoAlcanzadaIehdBuilder().
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

	if !strings.Contains(xmlStr, "<telefono>22123456</telefono>") {
		t.Errorf("telefono no se serializó correctamente: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<nombrePropietario>Juan Pérez</nombrePropietario>") {
		t.Errorf("nombrePropietario no se serializó correctamente: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<montoIehd>15.5</montoIehd>") && !strings.Contains(xmlStr, "<montoIehd>15.50</montoIehd>") {
		t.Errorf("montoIehd no se serializó correctamente: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<porcentajeIehd>7.5</porcentajeIehd>") && !strings.Contains(xmlStr, "<porcentajeIehd>7.50</porcentajeIehd>") {
		t.Errorf("porcentajeIehd no se serializó correctamente: %s", xmlStr)
	}
}

func TestHidrocarburoAlcanzadaIehd_NilAndUnwrap(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)

	cabecera := invoices.NewHidrocarburoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(123456).
		WithRazonSocialEmisor("EMPRESA TEST").
		WithMunicipio("Cochabamba").
		WithTelefono(nil).
		WithNumeroFactura(10).
		WithCuf("CUF-TEST").
		WithCufd("CUFD-TEST").
		WithCodigoSucursal(0).
		WithDireccion("Calle 1").
		WithCodigoPuntoVenta(nil).
		WithFechaEmision(fecha).
		WithNombreRazonSocial(nil).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("998877").
		WithComplemento("").
		WithCodigoCliente("CLI-NIL").
		WithCiudad(nil).
		WithNombrePropietario(nil).
		WithNombreRepresentanteLegal(nil).
		WithCondicionPago(nil).
		WithPeriodoEntrega(nil).
		WithCodigoMetodoPago(1).
		WithNumeroTarjeta(nil).
		WithMontoAPagar(100.0).
		WithMontoIehd(nil).
		WithTotalPrecioNeto(100.0).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(100.0).
		WithDescuentoAdicional(nil).
		WithCodigoExcepcion(nil).
		WithLeyenda("Ley N 453").
		WithUsuario("tester").
		Build()

	detalle := invoices.NewHidrocarburoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("061000").
		WithCodigoProductoSin(100).
		WithCodigoProducto("PROD-NIL").
		WithDescripcion("Producto").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(100.0).
		WithMontoDescuento(nil).
		WithSubTotal(100.0).
		WithPorcentajeIehd(nil).
		Build()

	factura := invoices.NewHidrocarburoAlcanzadaIehdBuilder().
		WithCabecera(cabecera).
		WithCabecera(invoices.HidrocarburoAlcanzadaIehdCabecera{}). // test unwrap nil
		WithCabecera(cabecera).
		AddDetalle(invoices.HidrocarburoAlcanzadaIehdDetalle{}). // test unwrap nil
		AddDetalle(detalle).
		WithModalidad(99). // test default fallback modalidad
		Build()

	raw, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error al serializar: %v", err)
	}
	xmlStr := string(raw)

	if !strings.Contains(xmlStr, `<montoIehd xsi:nil="true">`) {
		t.Errorf("montoIehd nil debe serializarse con xsi:nil=\"true\": %s", xmlStr)
	}
	if !strings.Contains(xmlStr, `<porcentajeIehd xsi:nil="true">`) {
		t.Errorf("porcentajeIehd nil debe serializarse con xsi:nil=\"true\": %s", xmlStr)
	}
}

func TestHidrocarburoAlcanzadaIehd_RecepcionFacturaBuilder(t *testing.T) {
	fecha := time.Date(2026, 7, 26, 12, 0, 0, 0, time.UTC)

	cab19 := invoices.NewHidrocarburoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("EMPRESA S.A.").
		WithMunicipio("La Paz").
		WithNumeroFactura(1).
		WithCuf("CUF19").
		WithCufd("CUFD19").
		WithCodigoSucursal(0).
		WithDireccion("Av. Principal").
		WithFechaEmision(fecha).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("12345").
		WithCodigoCliente("C19").
		WithCodigoMetodoPago(1).
		WithMontoAPagar(100).
		WithTotalPrecioNeto(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley 453").
		WithUsuario("usr").
		Build()

	det19 := invoices.NewHidrocarburoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("061000").
		WithCodigoProductoSin(111).
		WithCodigoProducto("P19").
		WithDescripcion("Combustible").
		WithCantidad(2).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(50).
		WithSubTotal(100).
		Build()

	factura19 := invoices.NewHidrocarburoAlcanzadaIehdBuilder().
		WithModalidad(siat.ModalidadComputarizada).
		WithCabecera(cab19).
		AddDetalle(det19).
		Build()

	builder19 := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(siat.ModalidadComputarizada).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(19).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd("CUFD19").
		WithCuis("CUIS19").
		WithTipoFacturaDocumento(1).
		WithFechaEnvio(fecha)

	if err := builder19.WithFactura(factura19, nil); err != nil {
		t.Fatalf("error en WithFactura sector 19: %v", err)
	}
	req19 := builder19.Build()
	raw19, err := xml.Marshal(req19)
	if err != nil {
		t.Fatalf("error al serializar solicitud sector 19: %v", err)
	}
	if !strings.Contains(string(raw19), "<archivo>") {
		t.Error("el campo archivo en la solicitud sector 19 no debe estar vacío")
	}
}

func TestHidrocarburoAlcanzadaIehd_ElectronicaIntegration(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 19, 1, 0, cufdControl)

	nombreRazonSocial := "COMPRADOR ALCANZADA IEHD"
	montoIehd := 10.0
	porcentaje := 5.0

	cabecera := invoices.NewHidrocarburoAlcanzadaIehdCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("PETROBOL S.A.").
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
		WithCodigoCliente("CLI-IEHD-01").
		WithCodigoMetodoPago(1).
		WithMontoAPagar(100.00).
		WithMontoIehd(&montoIehd).
		WithTotalPrecioNeto(90.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(100.00).
		WithLeyenda("Leyenda IEHD").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewHidrocarburoAlcanzadaIehdDetalleBuilder().
		WithActividadEconomica("061000").
		WithCodigoProductoSin(54321).
		WithCodigoProducto("IEHD-001").
		WithDescripcion("GASOLINA PREMIUM").
		WithCantidad(10.0).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(10.0).
		WithSubTotal(100.0).
		WithPorcentajeIehd(&porcentaje).
		Build()

	factura := invoices.NewHidrocarburoAlcanzadaIehdBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(19).
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
