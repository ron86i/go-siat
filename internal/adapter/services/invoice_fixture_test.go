package services_test

import (
	"encoding/xml"
	"strconv"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/boletoaereo"
	"github.com/ron86i/go-siat/v3/documents/notaconciliacion"
)

func buildBoletoAereoFixture(nit int64, cuf, cufd string, fechaEmision time.Time, modalidad int) boletoaereo.Factura {
	nombre := "JUAN PEREZ"
	return boletoaereo.Factura{
		Modalidad: modalidad,
		Cabecera: boletoaereo.Cabecera{
			NitEmisor:                    nit,
			RazonSocialEmisor:            "AEROLINEAS TEST S.A.",
			NumeroFactura:                1,
			Cuf:                          cuf,
			Cufd:                         cufd,
			CodigoSucursal:               0,
			Direccion:                    "Aeropuerto Internacional El Alto",
			FechaEmision:                 fechaEmision,
			NombreRazonSocial:            &nombre,
			CodigoTipoDocumentoIdentidad: 1,
			NumeroDocumento:              "1234567",
			NombrePasajero:               "JUAN PEREZ",
			CodigoIataLineaAerea:         900,
			CodigoOrigenServicio:         "LPB",
			CodigoMetodoPago:             1,
			MontoTarifa:                  1000,
			MontoAPagar:                  1000,
			TotalPrecioNeto:              870,
			Iva:                          130,
			CodigoMoneda:                 1,
			TipoCambio:                   1,
			MontoTotalMoneda:             1000,
			CodigoTipoTransaccion:        "I",
			Leyenda:                      "Ley 453...",
			Usuario:                      "operador",
			CodigoInterno:                "BOLETO-1",
			CodigoDocumentoSector:        30,
		},
	}
}

func buildNotaConciliacionFixture(nit int64, cuf, cufd string, fechaEmision time.Time, codModalidad int) notaconciliacion.NotaCreditoDebito {
	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	return notaconciliacion.NotaCreditoDebito{
		Modalidad: codModalidad,
		Cabecera: notaconciliacion.Cabecera{
			NitEmisor:                    nit,
			RazonSocialEmisor:            "Ronaldo Rua",
			Municipio:                    "Tarija",
			NumeroNotaConciliacion:       1,
			Cuf:                          cuf,
			Cufd:                         cufd,
			CodigoSucursal:               0,
			Direccion:                    "ESQUINA AVENIDA LA PAZ",
			CodigoPuntoVenta:             &codigoPuntoVenta,
			FechaEmision:                 fechaEmision,
			NombreRazonSocial:            &nombreRazonSocial,
			CodigoTipoDocumentoIdentidad: 1,
			NumeroDocumento:              "5115889",
			CodigoCliente:                "1",
			NumeroFactura:                10,
			NumeroAutorizacionCuf:        "DUMMY_AUT_CUF",
			FechaEmisionFactura:          fechaEmision.Add(-24 * time.Hour),
			TotalPrecioNetoOriginal:      200,
			MontoActualizacion:           50,
			UfvFactura:                   1,
			UfvConciliacion:              1.25,
			MontoTotalConciliado:         250.0,
			CreditoFiscalIva:             0,
			DebitoFiscalIva:              0,
			Leyenda:                      "Ley N° 453: Tienes derecho a recibir información...",
			Usuario:                      "usuario",
			CodigoInterno:                "NOTA-1",
			CodigoDocumentoSector:        29,
		},
		Detalles: []notaconciliacion.Detalle{
			{
				ActividadEconomica:   "477300",
				CodigoProductoSin:    622539,
				CodigoProducto:       "PROD-01",
				Descripcion:          "Producto DUMMY",
				TotalPrecioNeto:      200.0,
				TotalPrecioNetoFinal: 250.0,
				MontoConciliado:      250.0,
			},
		},
	}
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
