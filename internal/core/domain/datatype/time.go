package datatype

import (
	"encoding/xml"
	"time"
)

const layoutSiat = "2006-01-02T15:04:05.000"

// TimeSiat es un tipo personalizado que envuelve time.Time para manejar
// la serialización y deserialización XML con el formato específico del SIAT ("2006-01-02T15:04:05.000").
// Implementa las interfaces xml.Marshaler y xml.Unmarshaler para codificar/decodificar
// fechas en el formato requerido por los servicios del SIAT.
type TimeSiat time.Time

// MarshalXML Codificador de XML
func (t TimeSiat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Convertimos a `time.Time` para usar sus métodos
	v := time.Time(t)
	// Opcional: Si la fecha es zero, podrías querer omitirla o enviar string vacío
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

	// Si el tag está vacío, evitamos el error de parseo
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

// ToTime Helper para convertir de vuelta a time.Time fácilmente
func (t TimeSiat) ToTime() time.Time {
	return time.Time(t)
}
