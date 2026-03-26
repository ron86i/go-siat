package common

import "fmt"

// MensajeServicio representa una estructura de notificación o error devuelta por los servicios del SIAT.
type MensajeServicio struct {
	Codigo      int    `xml:"codigo" json:"codigo"`
	Descripcion string `xml:"descripcion" json:"descripcion"`
}

// Result define una interfaz común para todas las respuestas del SIAT.
type Result interface {
	IsSuccess() bool
	GetMessages() []MensajeServicio
}

// Summary retorna un string con todos los mensajes de error/notificación formateados.
func Summary(r Result) string {
	msgs := r.GetMessages()
	if len(msgs) == 0 {
		return ""
	}
	var res string
	for _, m := range msgs {
		res += fmt.Sprintf("[%d] %s; ", m.Codigo, m.Descripcion)
	}
	return res
}
