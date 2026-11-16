package models

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"
)

func TestRecepcionesSerializanFechaEnvioSinZona(t *testing.T) {
	fecha := time.Date(2026, 10, 19, 23, 50, 43, 498123456, time.FixedZone("Bolivia", -4*60*60))
	for nombre, solicitud := range map[string]any{
		"individual": NewRecepcionFacturaBuilder().WithFechaEnvio(fecha).Build(),
		"paquete":    NewRecepcionPaqueteFacturaBuilder().WithFechaEnvio(fecha).Build(),
		"masiva":     NewRecepcionMasivaFacturaBuilder().WithFechaEnvio(fecha).Build(),
	} {
		t.Run(nombre, func(t *testing.T) {
			data, err := xml.Marshal(solicitud)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(data), "<fechaEnvio>2026-10-19T23:50:43.498</fechaEnvio>") {
				t.Fatalf("fechaEnvio incompatible con SIAT: %s", data)
			}
		})
	}
}

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
