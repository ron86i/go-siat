package codigos

// Mensaje contiene el código y descripción técnica devueltos por el SIN.
type Mensaje struct {
	Codigo      int    `xml:"codigo"`
	Descripcion string `xml:"descripcion"`
}

// RespuestaCUIS conserva el CUIS generado y su fecha límite de vigencia literal.
type RespuestaCUIS struct {
	Codigo        string    `xml:"codigo"`
	FechaVigencia string    `xml:"fechaVigencia"`
	Mensajes      []Mensaje `xml:"mensajesList"`
	Transaccion   *bool     `xml:"transaccion"`
}

// RespuestaCuis es un alias para RespuestaCUIS.
type RespuestaCuis = RespuestaCUIS

// RespuestaCUFD conserva el CUFD, datos de control y código interno devuelto en SIAT v3.
type RespuestaCUFD struct {
	Codigo        string    `xml:"codigo"`
	CodigoControl string    `xml:"codigoControl"`
	CodigoInterno string    `xml:"codigoInterno"`
	Direccion     string    `xml:"direccion"`
	FechaVigencia string    `xml:"fechaVigencia"`
	Mensajes      []Mensaje `xml:"mensajesList"`
	Transaccion   *bool     `xml:"transaccion"`
}

// RespuestaCufd es un alias para RespuestaCUFD.
type RespuestaCufd = RespuestaCUFD
