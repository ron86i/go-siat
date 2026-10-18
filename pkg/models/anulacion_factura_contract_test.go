package models_test

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/pkg/models"
)

func TestAnulacionFacturaXMLOrder(t *testing.T) {
	request := models.NewAnulacionFacturaBuilder().WithCodigoMotivo(1).WithCuf("CUF").Build()
	data, err := xml.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "<codigoMotivo>1</codigoMotivo><cuf>CUF</cuf>") {
		t.Fatalf("WSDL requires codigoMotivo before cuf: %s", data)
	}
}

func TestConsultaEventoSignificativoIncludesCufd(t *testing.T) {
	request := models.NewConsultaEventoSignificativoBuilder().WithCufd("CUFD").WithCuis("CUIS").Build()
	data, err := xml.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "<cufd>CUFD</cufd><cuis>CUIS</cuis>") {
		t.Fatalf("event query is missing the ordered CUFD field: %s", data)
	}
}
