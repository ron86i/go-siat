package vuelocharter_test

import (
	"encoding/xml"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/vuelocharter"
)

func TestFacturaV3Marshal(t *testing.T) {
	doc := vuelocharter.Factura{
		Modalidad: 1,
		Cabecera:  vuelocharter.Cabecera{},
	}
	data, err := xml.Marshal(doc)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}
	if len(data) == 0 {
		t.Error("se esperaba xml serializado no vacío")
	}
}
