package operaciones

import "encoding/xml"

// Mensaje contiene un código numérico y una descripción devuelta por el SIN.
type Mensaje struct {
	Codigo      int    `xml:"codigo"`
	Descripcion string `xml:"descripcion"`
}

// Evento contiene los datos de un evento significativo registrado o consultado.
type Evento struct {
	CodigoEvento                       int    `xml:"codigoEvento"`
	CodigoRecepcionEventoSignificativo int64  `xml:"codigoRecepcionEventoSignificativo"`
	Descripcion                        string `xml:"descripcion"`
	FechaFin                           string `xml:"fechaFin"`
	FechaInicio                        string `xml:"fechaInicio"`
}

// RespuestaEventos conserva el resultado y listado de eventos significativos.
type RespuestaEventos struct {
	XMLName                            xml.Name  `xml:"RespuestaListaEventos"`
	CodigoRecepcionEventoSignificativo int64     `xml:"codigoRecepcionEventoSignificativo"`
	Eventos                            []Evento  `xml:"listaCodigos"`
	Mensajes                           []Mensaje `xml:"mensajesList"`
	Transaccion                        *bool     `xml:"transaccion"`
}

// PuntoVenta contiene la información de un punto de venta registrado ante el SIN.
type PuntoVenta struct {
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta"`
	NombrePuntoVenta string `xml:"nombrePuntoVenta"`
	TipoPuntoVenta   string `xml:"tipoPuntoVenta"`
}

// RespuestaRegistroPuntoVenta conserva la confirmación de registro de un punto de venta.
type RespuestaRegistroPuntoVenta struct {
	XMLName          xml.Name  `xml:"RespuestaRegistroPuntoVenta"`
	CodigoPuntoVenta *int      `xml:"codigoPuntoVenta"`
	Mensajes         []Mensaje `xml:"mensajesList"`
	Transaccion      *bool     `xml:"transaccion"`
}

// RespuestaConsultaPuntoVenta contiene los puntos de venta vigentes devueltos por el SIN.
type RespuestaConsultaPuntoVenta struct {
	XMLName      xml.Name     `xml:"RespuestaConsultaPuntoVenta"`
	PuntosVentas []PuntoVenta `xml:"listaPuntosVentas"`
	Mensajes     []Mensaje    `xml:"mensajesList"`
	Transaccion  *bool        `xml:"transaccion"`
}

// RespuestaCierrePuntoVenta conserva la respuesta de cierre de un punto de venta.
type RespuestaCierrePuntoVenta struct {
	XMLName          xml.Name  `xml:"RespuestaCierrePuntoVenta"`
	CodigoPuntoVenta *int      `xml:"codigoPuntoVenta"`
	Mensajes         []Mensaje `xml:"mensajesList"`
	Transaccion      *bool     `xml:"transaccion"`
}

// RespuestaRegistroComisionista conserva la respuesta de registro de comisionista.
type RespuestaRegistroComisionista struct {
	XMLName          xml.Name  `xml:"RespuestaPuntoVentaComisionista"`
	CodigoPuntoVenta *int      `xml:"codigoPuntoVenta"`
	Mensajes         []Mensaje `xml:"mensajesList"`
	Transaccion      *bool     `xml:"transaccion"`
}

// RespuestaCierreSistema contiene el resultado del cierre de operaciones del sistema.
type RespuestaCierreSistema struct {
	XMLName       xml.Name  `xml:"RespuestaCierreSistemas"`
	CodigoSistema string    `xml:"codigoSistema"`
	Mensajes      []Mensaje `xml:"mensajesList"`
	Transaccion   *bool     `xml:"transaccion"`
}

// RespuestaComunicacion conserva el resultado de verificar conectividad con el servicio.
type RespuestaComunicacion struct {
	Mensajes    []Mensaje `xml:"mensajesList"`
	Transaccion *bool     `xml:"transaccion"`
}
