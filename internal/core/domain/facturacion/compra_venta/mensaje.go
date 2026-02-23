package compra_venta

// MensajeServicio representa una estructura de notificación o error devuelta por los servicios del SIAT.
// Contiene un código identificador y una descripción detallada que permite entender el resultado de una operación.
type MensajeServicio struct {
	Codigo      int    `xml:"codigo" json:"codigo"`
	Descripcion string `xml:"descripcion" json:"descripcion"`
}

type RespuestaServicioFacturacion struct {
	CodigoDescripcion string            `xml:"codigoDescripcion,omitempty" json:"codigoDescripcion"`
	CodigoEstado      int               `xml:"codigoEstado,omitempty" json:"codigoEstado"`
	CodigoRecepcion   string            `xml:"codigoRecepcion,omitempty" json:"codigoRecepcion"`
	MensajesList      []MensajeServicio `xml:"mensajesList,omitempty" json:"mensajesList"`
	Transaccion       bool              `xml:"transaccion,omitempty" json:"transaccion"`
}
