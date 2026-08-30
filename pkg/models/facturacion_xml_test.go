package models

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestRecepcionFacturaBuilderConservaXMLIndentado(t *testing.T) {
	type facturaPrueba struct {
		XMLName xml.Name `xml:"facturaPrueba"`
		Nombre  string   `xml:"nombre"`
	}
	builder := NewRecepcionFacturaBuilder().
		WithCodigoModalidad(ModalidadComputarizada).
		WithXMLIndentado()
	if err := builder.WithFactura(facturaPrueba{Nombre: "Falmus"}, nil); err != nil {
		t.Fatal(err)
	}
	xmlPreparado, _, disponible := builder.XMLPreparado()
	if !disponible {
		t.Fatal("XML preparado no disponible")
	}
	contenido := string(xmlPreparado)
	if !strings.HasPrefix(contenido, xml.Header) || !strings.Contains(contenido, "\n    <nombre>Falmus</nombre>\n") {
		t.Fatalf("XML no quedó indentado: %q", contenido)
	}
}
