package comercialexportacion_test

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/comercialexportacion"
)

func TestFacturaV3ComercialExportacionMarshal(t *testing.T) {
	factura := comercialexportacion.Factura{
		Modalidad: 1, // Electrónica
		Cabecera: comercialexportacion.Cabecera{
			NitEmisor:                    123456789,
			RazonSocialEmisor:            "Exportadora Boliviana V3",
			Municipio:                    "Santa Cruz",
			NumeroFactura:                51,
			Cuf:                          "CUF-EXP-123V3",
			Cufd:                         "CUFD-EXP-123",
			Direccion:                    "Parque Industrial",
			FechaEmision:                 time.Date(2026, 10, 31, 12, 0, 0, 0, time.UTC),
			CodigoTipoDocumentoIdentidad: 1,
			NumeroDocumento:              "88776655",
			DireccionComprador:           "Avenida Paulista 1000",
			CodigoCliente:                "CLI-BR",
			Incoterm:                     4, // FOB
			IncotermAlcance:              "FOB Puerto",
			IncotermDetalle:              "Puerto Suarez",
			PuertoDestino:                "Santos",
			LugarDestino:                 "Brasil",
			CodigoPais:                   105,
			CodigoMetodoPago:             1,
			TotalPrecioNeto:              1000.0,
			MontoAPagar:                  1000.0,
			TotalGastosNacionalesFob:     100.0,
			SubTotalNeto:                 900.0,
			Iva:                          0,
			CodigoMoneda:                 2, // USD
			TipoCambio:                   6.96123,
			MontoTotalMoneda:             6960.0,
			Leyenda:                      "Ley N 453",
			Usuario:                      "agente",
			CodigoInterno:                "EXP-2026",
		},
		Detalles: []comercialexportacion.Detalle{
			{
				ActividadEconomica: "011100",
				CodigoProductoSin:  55443,
				CodigoProducto:     "SOYA-01",
				CodigoNandina:      "12010000",
				Descripcion:        "Granos de soya",
				Cantidad:           10.12345,
				UnidadMedida:       25,
				PrecioNetoUnitario: 90.12345,
				SubTotal:           900.0,
			},
		},
	}

	data, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}

	xmlStr := string(data)
	if !strings.Contains(xmlStr, "<facturaElectronicaComercialExportacion") {
		t.Errorf("elemento raíz incorrecto: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<incoterm>4</incoterm>") {
		t.Errorf("incoterm FOB numérico 4 no encontrado: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<precioNetoUnitario>90.12345</precioNetoUnitario>") {
		t.Errorf("precioNetoUnitario no encontrado: %s", xmlStr)
	}
	for _, expected := range []string{"<tipoCambio>6.96123</tipoCambio>", "<cantidad>10.12345</cantidad>"} {
		if !strings.Contains(xmlStr, expected) {
			t.Errorf("XML no contiene %s: %s", expected, xmlStr)
		}
	}
}

func TestFacturaV3IncotermRange(t *testing.T) {
	for _, tc := range []struct {
		value   int
		invalid bool
	}{{value: 1}, {value: 11}, {value: 0, invalid: true}, {value: 12, invalid: true}} {
		factura := comercialexportacion.Factura{Modalidad: 1, Cabecera: comercialexportacion.Cabecera{Incoterm: tc.value}}
		if err := factura.Validate(); (err != nil) != tc.invalid {
			t.Errorf("Incoterm %d: Validate() error = %v, invalid=%t", tc.value, err, tc.invalid)
		}
	}
}
