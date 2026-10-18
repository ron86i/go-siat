package datatype

import "time"

// TimeSiat es un alias de time.Time. Usa la serialización estándar de Go,
// con zona horaria y sin conversiones implícitas a UTC.
// Conserva el nombre utilizado por los modelos, pero no define un tipo nuevo.
type TimeSiat = time.Time

// NewTimeSiat devuelve t sin modificar el instante ni la zona horaria.
func NewTimeSiat(t time.Time) time.Time {
	return t
}
