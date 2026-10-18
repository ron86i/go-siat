package invoices_test

import (
	"encoding/xml"
	"testing"

	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
)

func TestVentaMineralBCBConservaFormatoElectronico(t *testing.T) {
	factura := invoices.NewVentaMineralBCBBuilder().WithModalidad(2).Build()
	data, err := xml.Marshal(factura)
	if err != nil {
		t.Fatal(err)
	}
	var node xsdNode
	if err := xml.Unmarshal(data, &node); err != nil {
		t.Fatal(err)
	}
	if node.XMLName.Local != "facturaElectronicaVentaMineralBCB" {
		t.Fatalf("no existe un XSD computarizado para Venta Mineral BCB: %s", node.XMLName.Local)
	}
}
