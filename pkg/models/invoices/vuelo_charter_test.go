package invoices_test

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
)

func TestVueloCharterBuilders(t *testing.T) {
	for modalidad, root := range map[int]string{1: "facturaElectronicaVueloCharter", 2: "facturaComputarizadaVueloCharter"} {
		t.Run(root, func(t *testing.T) {
			cabecera := invoices.NewVueloCharterCabeceraBuilder().
				WithComplemento("").WithCodigoInterno("VC-1").WithMontoTasa(new(5.25)).Build()
			detalle := invoices.NewVueloCharterDetalleBuilder().
				WithPrecioNetoUnitario(12.34567).WithNumeroSerie(new("SERIE-1")).WithNumeroImei(nil).Build()
			factura := invoices.NewVueloCharterBuilder().WithModalidad(modalidad).
				WithCabecera(cabecera).AddDetalle(detalle).Build()
			data, err := xml.Marshal(factura)
			if err != nil {
				t.Fatal(err)
			}
			for _, fragment := range []string{"<" + root, "<codigoDocumentoSector>41</codigoDocumentoSector>", "<montoTasa>5.25</montoTasa>", "<precioNetoUnitario>12.34567</precioNetoUnitario>", "<numeroSerie>SERIE-1</numeroSerie>"} {
				if !strings.Contains(string(data), fragment) {
					t.Fatalf("falta %s en %s", fragment, data)
				}
			}
		})
	}
}
