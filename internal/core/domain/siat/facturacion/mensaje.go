package facturacion

import "github.com/ron86i/go-siat/v3/internal/core/domain/siat/common"

// MensajeServicio representa un mensaje devuelto por el servidor del SIAT
type MensajeServicio = common.MensajeServicio

// MensajeRecepcion añade la ubicación y advertencia del mensaje de facturación.
// El WSDL distingue este tipo del mensaje común usado por otros servicios.
type MensajeRecepcion struct {
	common.MensajeServicio
	Advertencia   bool `xml:"advertencia,omitempty" json:"advertencia,omitzero"`
	NumeroArchivo int  `xml:"numeroArchivo,omitempty" json:"numeroArchivo,omitzero"`
	NumeroDetalle int  `xml:"numeroDetalle,omitempty" json:"numeroDetalle,omitzero"`
}

type RespuestaRecepcion struct {
	CodigoDescripcion string             `xml:"codigoDescripcion" json:"codigoDescripcion"`
	CodigoEstado      int                `xml:"codigoEstado" json:"codigoEstado"`
	CodigoRecepcion   string             `xml:"codigoRecepcion" json:"codigoRecepcion"`
	MensajesList      []MensajeRecepcion `xml:"mensajesList" json:"mensajesList"`
	Transaccion       bool               `xml:"transaccion" json:"transaccion"`
}
