package exportacionminerales_test

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/exportacionminerales"
)

func TestFacturaV3Marshal(t *testing.T) {
	doc := exportacionminerales.Factura{
		Modalidad: 1,
		Cabecera: exportacionminerales.Cabecera{
			Incoterm: 1, TipoCambio: 1.234567, KilosNetosHumedos: 1, KilosNetosSecos: 1,
			SubTotalNeto: 1, TotalPrecioNeto: 1, MontoAPagar: 1, MontoTotalMoneda: 1, GastosRealizacion: 1,
		},
		Detalles: []exportacionminerales.Detalle{{
			CantidadExtraccion: 1.23456,
			Cantidad:           1.23456,
			PrecioNetoUnitario: 1.23456,
			SubTotal:           1.23456,
		}},
	}
	data, err := xml.Marshal(doc)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}
	if len(data) == 0 {
		t.Error("se esperaba xml serializado no vacío")
	}
	for _, expected := range []string{"<complemento>0</complemento>", "<tipoCambio>1.23457</tipoCambio>", "<cantidadExtraccion>1.23456</cantidadExtraccion>", "<cantidad>1.23456</cantidad>", "<precioNetoUnitario>1.23456</precioNetoUnitario>", "<subTotal>1.23456</subTotal>"} {
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
		factura := exportacionminerales.Factura{Modalidad: 1, Cabecera: exportacionminerales.Cabecera{Incoterm: tc.value}}
		if err := factura.Validate(); (err != nil) != tc.invalid {
			t.Errorf("Incoterm %d: Validate() error = %v, invalid=%t", tc.value, err, tc.invalid)
		}
	}
}

func TestFacturaV3ValidateFixedFields(t *testing.T) {
	for _, tc := range []struct {
		name       string
		complement string
		iva        float64
		wantErr    bool
	}{
		{name: "valores fijos permitidos", complement: "0"},
		{name: "complemento incorrecto", complement: "1", wantErr: true},
		{name: "iva distinto de cero", complement: "0", iva: 1, wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			factura := exportacionminerales.Factura{Modalidad: 1, Cabecera: exportacionminerales.Cabecera{Incoterm: 1, Complemento: tc.complement, Iva: tc.iva}}
			if err := factura.Validate(); (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr=%t", err, tc.wantErr)
			}
		})
	}
}
