package facturacion

// MensajeRecepcion contiene un mensaje de servicio, observación o advertencia devuelto por el SIN.
type MensajeRecepcion struct {
	Codigo        int    `xml:"codigo"`
	Descripcion   string `xml:"descripcion"`
	Advertencia   bool   `xml:"advertencia"`
	NumeroArchivo int    `xml:"numeroArchivo"`
	NumeroDetalle int    `xml:"numeroDetalle"`
}

// RespuestaRecepcion conserva el resultado fiscal, estado y mensajes emitidos por el SIN.
type RespuestaRecepcion struct {
	CodigoDescripcion string             `xml:"codigoDescripcion"`
	CodigoEstado      int                `xml:"codigoEstado"`
	CodigoRecepcion   string             `xml:"codigoRecepcion"`
	Mensajes          []MensajeRecepcion `xml:"mensajesList"`
	Transaccion       *bool              `xml:"transaccion"`
}

// RespuestaComunicacion conserva el resultado de verificar comunicación con el endpoint.
type RespuestaComunicacion struct {
	Mensajes    []MensajeRecepcion `xml:"mensajesList"`
	Transaccion *bool              `xml:"transaccion"`
}
