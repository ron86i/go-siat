package facturacion

import "github.com/ron86i/go-siat/internal/core/domain/siat/common"

// MensajeServicio representa un mensaje devuelto por el servidor del SIAT
type MensajeServicio = common.MensajeServicio

type RespuestaRecepcion struct {
	CodigoDescripcion string            `xml:"codigoDescripcion" json:"codigoDescripcion"`
	CodigoEstado      int               `xml:"codigoEstado" json:"codigoEstado"`
	CodigoRecepcion   string            `xml:"codigoRecepcion" json:"codigoRecepcion"`
	MensajesList      []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion       bool              `xml:"transaccion" json:"transaccion"`
}
