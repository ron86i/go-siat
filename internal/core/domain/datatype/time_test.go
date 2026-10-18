package datatype

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"
)

func TestTimeSiatStandardSerialization(t *testing.T) {
	for _, zone := range []string{"Z", "-04:00"} {
		t.Run(zone, func(t *testing.T) {
			input := "2026-10-18T11:30:00" + zone
			var value TimeSiat
			if err := xml.Unmarshal([]byte("<fecha>"+input+"</fecha>"), &value); err != nil {
				t.Fatal(err)
			}
			var standard time.Time = NewTimeSiat(value)
			encoded, err := json.Marshal(standard)
			if err != nil {
				t.Fatal(err)
			}
			if string(encoded) != `"`+input+`"` {
				t.Fatalf("fecha modificada: %s", encoded)
			}
			data, err := xml.Marshal(struct {
				XMLName xml.Name  `xml:"documento"`
				Fecha   time.Time `xml:"fecha"`
			}{Fecha: standard})
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != "<documento><fecha>"+input+"</fecha></documento>" {
				t.Fatalf("XML inesperado: %s", data)
			}
		})
	}
}
