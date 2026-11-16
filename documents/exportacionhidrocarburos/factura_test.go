package exportacionhidrocarburos_test

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/exportacionhidrocarburos"
)

func TestFacturaV3Marshal(t *testing.T) {
	descuento := 0.123456
	doc := exportacionhidrocarburos.Factura{
		Modalidad: 1,
		Cabecera:  exportacionhidrocarburos.Cabecera{Incoterm: 1, TipoCambio: 1.234567},
		Detalles:  []exportacionhidrocarburos.Detalle{{Cantidad: 1.234567, PrecioNetoUnitario: 1.234567, MontoDescuento: &descuento, SubTotal: 1.234567}},
	}
	data, err := xml.Marshal(doc)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}
	if len(data) == 0 {
		t.Error("se esperaba xml serializado no vacío")
	}
	for _, expected := range []string{"<tipoCambio>1.23457</tipoCambio>", "<cantidad>1.234567</cantidad>", "<precioNetoUnitario>1.234567</precioNetoUnitario>", "<montoDescuento>0.123456</montoDescuento>", "<subTotal>1.234567</subTotal>"} {
		if !strings.Contains(string(data), expected) {
			t.Errorf("XML no contiene %s: %s", expected, data)
		}
	}
}

func TestFacturaV3IncotermRange(t *testing.T) {
	for _, tc := range []struct {
		value   int
		invalid bool
	}{{value: 1}, {value: 11}, {value: 0, invalid: true}, {value: 12, invalid: true}} {
		factura := exportacionhidrocarburos.Factura{Modalidad: 1, Cabecera: exportacionhidrocarburos.Cabecera{Incoterm: tc.value}}
		if err := factura.Validate(); (err != nil) != tc.invalid {
			t.Errorf("Incoterm %d: Validate() error = %v, invalid=%t", tc.value, err, tc.invalid)
		}
	}
}
