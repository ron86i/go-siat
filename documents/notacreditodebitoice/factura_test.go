package notacreditodebitoice_test

import (
	"encoding/xml"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/notacreditodebitoice"
)

func TestNotaCreditoDebitoV3Marshal(t *testing.T) {
	doc := notacreditodebitoice.NotaCreditoDebito{
		Modalidad: 1,
		Cabecera:  notacreditodebitoice.Cabecera{},
	}
	data, err := xml.Marshal(doc)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}
	if len(data) == 0 {
		t.Error("se esperaba xml serializado no vacío")
	}
}
