package sincronizacion

import "github.com/ron86i/go-siat/v2/internal/core/domain/siat/common"

// MensajeServicio representa un mensaje devuelto por el servidor del SIAT
type MensajeServicio = common.MensajeServicio

// RespuestaServicio representa la estructura genérica común para las respuestas SOAP del SIAT.
type RespuestaServicio struct {
	Transaccion  bool              `xml:"transaccion" json:"transaccion"`
	MensajesList []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
}
