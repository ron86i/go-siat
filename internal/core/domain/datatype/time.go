package datatype

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

const layoutSiat = "2006-01-02T15:04:05.000"

// TimeSiat es un tipo personalizado que envuelve time.Time para manejar
// la serialización y deserialización XML/JSON con el formato específico del SIAT ("2006-01-02T15:04:05.000").
// Implementa las interfaces de marshaling para asegurar el formato requerido por el SIAT.
type TimeSiat time.Time

// MarshalXML Codificador de XML
func (t TimeSiat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	v := time.Time(t)
	if v.IsZero() {
		return e.EncodeElement("", start)
	}
	return e.EncodeElement(v.Format(layoutSiat), start)
}

// UnmarshalXML Decodificador de XML
func (t *TimeSiat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	elem, err := time.Parse(layoutSiat, s)
	if err != nil {
		return err
	}

	*t = TimeSiat(elem)
	return nil
}

// MarshalJSON Codificador de JSON
func (t TimeSiat) MarshalJSON() ([]byte, error) {
	v := time.Time(t)
	if v.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", v.Format(layoutSiat))), nil
}

// UnmarshalJSON Decodificador de JSON
func (t *TimeSiat) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == "\"\"" {
		return nil
	}

	s = strings.Trim(s, "\"")
	elem, err := time.Parse(layoutSiat, s)
	if err != nil {
		return err
	}

	*t = TimeSiat(elem)
	return nil
}

// ToTime Helper para convertir de vuelta a time.Time fácilmente
func (t TimeSiat) ToTime() time.Time {
	return time.Time(t)
}

// String devuelve la representación en cadena de texto de la fecha en formato SIAT.
// Esto permite que al imprimir el objeto (ej. en logs) se vea la fecha formateada
// en lugar de los campos internos de time.Time.
func (t TimeSiat) String() string {
	v := time.Time(t)
	if v.IsZero() {
		return ""
	}
	return v.Format(layoutSiat)
}

func NewTimeSiat(t time.Time) TimeSiat {
	return TimeSiat(t)
}
