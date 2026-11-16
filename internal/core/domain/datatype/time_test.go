package datatype

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"testing"
	"time"
)

func TestTimeSiatConservaZonaEnJSON(t *testing.T) {
	for _, zone := range []string{"Z", "-04:00"} {
		t.Run(zone, func(t *testing.T) {
			input := "2026-10-18T11:30:00" + zone
			var value TimeSiat
			if err := xml.Unmarshal([]byte("<fecha>"+input+"</fecha>"), &value); err != nil {
				t.Fatal(err)
			}
			standard := NewTimeSiat(value.Time)
			encoded, err := json.Marshal(standard)
			if err != nil {
				t.Fatal(err)
			}
			if string(encoded) != `"`+input+`"` {
				t.Fatalf("fecha modificada: %s", encoded)
			}
			data, err := xml.Marshal(struct {
				XMLName xml.Name `xml:"documento"`
				Fecha   TimeSiat `xml:"fecha"`
			}{Fecha: standard})
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != "<documento><fecha>2026-10-18T11:30:00.000</fecha></documento>" {
				t.Fatalf("XML inesperado: %s", data)
			}
		})
	}
}

func TestTimeSiatXMLSinZonaConMilisegundos(t *testing.T) {
	for _, zona := range []*time.Location{time.UTC, time.FixedZone("Bolivia", -4*60*60)} {
		fecha := NewTimeSiat(time.Date(2026, 10, 19, 23, 50, 43, 498123456, zona))
		data, err := xml.Marshal(struct {
			XMLName xml.Name `xml:"documento"`
			Fecha   TimeSiat `xml:"fecha"`
		}{Fecha: fecha})
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(data), "<fecha>2026-10-19T23:50:43.498</fecha>") {
			t.Fatalf("XML debe conservar la hora y omitir la zona: %s", data)
		}
	}
}

func TestTimeSiatXMLAceptaFechaSinZona(t *testing.T) {
	var fecha TimeSiat
	if err := xml.Unmarshal([]byte("<fecha>2026-10-19T23:50:43.498</fecha>"), &fecha); err != nil {
		t.Fatal(err)
	}
}
