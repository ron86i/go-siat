package datatype

import (
	"encoding/xml"
	"time"
)

// TimeSiat serializa XML con milisegundos y sin zona, conservando la hora local.
// JSON conserva la serialización de time.Time, incluida su zona horaria.
type TimeSiat struct {
	time.Time
}

// NewTimeSiat conserva el instante y la zona horaria recibidos.
func NewTimeSiat(t time.Time) TimeSiat {
	return TimeSiat{Time: t}
}

func (t TimeSiat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.Format("2006-01-02T15:04:05.000"), start)
}

// UnmarshalXML interpreta fechas sin zona como hora boliviana; si incluyen
// una zona explícita, la conserva.
func (t *TimeSiat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	if err := d.DecodeElement(&value, &start); err != nil {
		return err
	}
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		parsed, err = time.ParseInLocation("2006-01-02T15:04:05.999999999", value, time.FixedZone("America/La_Paz", -4*60*60))
	}
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}
