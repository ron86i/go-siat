package hidrocarburosiehd_test

import (
	"encoding/xml"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/hidrocarburosiehd"
)

func TestFacturaV3Marshal(t *testing.T) {
	doc := hidrocarburosiehd.Factura{
		Modalidad: 1,
		Cabecera:  hidrocarburosiehd.Cabecera{},
	}
	data, err := xml.Marshal(doc)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}
	if len(data) == 0 {
		t.Error("se esperaba xml serializado no vacío")
	}
}
