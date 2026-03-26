package facturacion

import "github.com/ron86i/go-siat/internal/core/domain/siat/common"

// Re-exportamos MensajeServicio para no romper el código existente, pero usamos el tipo común.
type MensajeServicio = common.MensajeServicio

type RespuestaRecepcion struct {
	CodigoDescripcion string            `xml:"codigoDescripcion" json:"codigoDescripcion"`
	CodigoEstado      int               `xml:"codigoEstado" json:"codigoEstado"`
	CodigoRecepcion   string            `xml:"codigoRecepcion" json:"codigoRecepcion"`
	MensajesList      []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion       bool              `xml:"transaccion" json:"transaccion"`
}

// IsSuccess implementa la interfaz common.Result.
func (r RespuestaRecepcion) IsSuccess() bool {
	return r.Transaccion
}

// GetMessages implementa la interfaz common.Result.
func (r RespuestaRecepcion) GetMessages() []common.MensajeServicio {
	return r.MensajesList
}
