package compras

// Mensaje contiene un código numérico y una descripción devuelta por el SIN.
type Mensaje struct {
	Codigo        int    `xml:"codigo"`
	Descripcion   string `xml:"descripcion"`
	Advertencia   bool   `xml:"advertencia"`
	NumeroArchivo int    `xml:"numeroArchivo"`
	NumeroDetalle int    `xml:"numeroDetalle"`
}

// RespuestaRecepcion conserva el estado y los mensajes del SIN para operaciones de compras.
type RespuestaRecepcion struct {
	CodigoDescripcion string    `xml:"codigoDescripcion"`
	CodigoEstado      int       `xml:"codigoEstado"`
	CodigoRecepcion   string    `xml:"codigoRecepcion"`
	Mensajes          []Mensaje `xml:"mensajesList"`
	Transaccion       *bool     `xml:"transaccion"`
}

// RespuestaConsulta contiene el archivo decodificado y los mensajes devueltos por el SIN.
type RespuestaConsulta struct {
	Archivo     []byte
	Mensajes    []Mensaje
	Transaccion *bool
}

// RespuestaComunicacion conserva el resultado de verificar comunicación con el servicio.
type RespuestaComunicacion struct {
	Mensajes    []Mensaje `xml:"mensajesList"`
	Transaccion *bool     `xml:"transaccion"`
}
