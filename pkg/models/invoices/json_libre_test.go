package invoices_test

import (
	"encoding/json"
	"testing"

	"github.com/ron86i/go-siat/v2/pkg/models/invoices"
	"github.com/stretchr/testify/assert"
)

func TestBuildersConJSONLibreExponenErrorDeSerializacion(t *testing.T) {
	valorNoSerializable := map[string]any{"canal": make(chan int)}
	casos := []struct {
		nombre string
		err    error
	}{
		{
			nombre: "exportacion comercial",
			err: invoices.NewComercialExportacionCabeceraBuilder().
				WithCostosGastosNacionales(valorNoSerializable).
				Err(),
		},
		{
			nombre: "exportacion punto de venta",
			err: invoices.NewComercialExportacionPVentaCabeceraBuilder().
				WithCostosGastosNacionales(valorNoSerializable).
				Err(),
		},
		{
			nombre: "exportacion hidrocarburos",
			err: invoices.NewComercialExportacionHidroCabeceraBuilder().
				WithCostosGastosNacionales(valorNoSerializable).
				Err(),
		},
		{
			nombre: "exportacion minera",
			err: invoices.NewComercialExportacionMineraCabeceraBuilder().
				WithOtrosDatos(valorNoSerializable).
				Err(),
		},
		{
			nombre: "hotel",
			err: invoices.NewHotelDetalleBuilder().
				WithDetalleHuespedes(valorNoSerializable).
				Err(),
		},
	}

	for _, tc := range casos {
		t.Run(tc.nombre, func(t *testing.T) {
			var unsupportedTypeError *json.UnsupportedTypeError
			assert.ErrorAs(t, tc.err, &unsupportedTypeError)
		})
	}
}
