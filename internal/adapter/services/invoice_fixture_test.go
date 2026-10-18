package services_test

import (
	"encoding/xml"
	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
	"strconv"
	"testing"
	"time"
)

func buildBoletoAereoFixture(nit int64, cuf, cufd string, fechaEmision time.Time, modalidad int) invoices.BoletoAereo {
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
		WithMontoAPagar(1000).
		WithTotalPrecioNeto(870).
		WithIva(130).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(1000).
		WithCodigoTipoTransaccion("I").
		WithLeyenda("Ley 453...").
		WithUsuario("operador").
		WithCodigoInterno("BOLETO-1").
		Build()

	factura := invoices.NewBoletoAereoBuilder().
		WithModalidad(modalidad).
		WithCabecera(cabecera).
		Build()

	return factura
}

func buildNotaConciliacionFixture(nit int64, cuf, cufd string, fechaEmision time.Time, codModalidad int) invoices.NotaConciliacion {
	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	cabecera := invoices.NewNotaConciliacionCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroNotaConciliacion(1).
		WithCuf(cuf).
		WithCufd(cufd).
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
		WithTotalPrecioNetoOriginal(200).
		WithMontoActualizacion(50).
		WithUfvFactura(1).
		WithUfvConciliacion(1.25).
		WithMontoTotalConciliado(250.0).
		WithCreditoFiscalIva(0).
		WithDebitoFiscalIva(0).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		WithCodigoInterno("NOTA-1").
		Build()

	detalleA := invoices.NewNotaDetalleOriginalBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("PROD-01").
		WithDescripcion("Producto DUMMY").
		WithCantidad(2.0).
		WithUnidadMedida(1).
		WithPrecioNetoUnitario(100.0).
		WithSubTotal(200.0).
		Build()

	detalleB := invoices.NewNotaDetalleConciliacionBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("PROD-01").
		WithDescripcion("Producto DUMMY").
		WithMontoConciliado(250.0).
		Build()

	nota := invoices.NewNotaConciliacionBuilder().
		WithModalidad(codModalidad).
		WithCabecera(cabecera).
		AddDetalleOriginal(detalleA).
		AddDetalleConciliacion(detalleB).
		Build()

	return nota
}

func TestFacturasFixturesSinRed(t *testing.T) {
	fecha := time.Date(2026, 10, 18, 10, 0, 0, 0, time.UTC)
	_, dutty := buildDuttyFree(t, &FacturacionTestContext{Nit: 123456789, Modalidad: 1, Cufd: "CUFD", CufdControl: "ABCD"}, 1, fecha, 1)
	cases := []struct {
		name     string
		document any
		positive []string
	}{
		{"DuttyFree", dutty, []string{"subTotalNeto", "totalPrecioNeto", "montoaPagar"}},
		{"BoletoAereo", buildBoletoAereoFixture(123456789, "CUF", "CUFD", fecha, 1), []string{"iva", "totalPrecioNeto", "montoaPagar"}},
		{"NotaConciliacion", buildNotaConciliacionFixture(123456789, "CUF", "CUFD", fecha, 1), []string{"totalPrecioNetoOriginal", "montoActualizacion", "ufvFactura", "ufvConciliacion"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, err := xml.Marshal(c.document)
			if err != nil {
				t.Fatal(err)
			}
			// Decode names explicitly; this test checks the data used by the live fixtures.
			var doc struct {
				Cabecera struct {
					Fields []struct {
						XMLName xml.Name
						Value   string `xml:",chardata"`
					} `xml:",any"`
				} `xml:"cabecera"`
			}
			if err := xml.Unmarshal(data, &doc); err != nil {
				t.Fatal(err)
			}
			fields := map[string]string{}
			for _, f := range doc.Cabecera.Fields {
				fields[f.XMLName.Local] = f.Value
			}
			for _, name := range c.positive {
				v, err := strconv.ParseFloat(fields[name], 64)
				if err != nil || v <= 0 {
					t.Errorf("%s debe ser positivo según XSD; valor=%q", name, fields[name])
				}
			}
			if len(fields["codigoInterno"]) < 1 || len(fields["codigoInterno"]) > 10 {
				t.Errorf("codigoInterno inválido: %q", fields["codigoInterno"])
			}
		})
	}
}
