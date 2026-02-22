package operaciones

import (
	"encoding/xml"
	"time"
)

// RegistroPuntoVenta es el wrapper para registrar un punto de venta
type RegistroPuntoVenta struct {
	XMLName                     xml.Name                    `xml:"ns:registroPuntoVenta"`
	SolicitudRegistroPuntoVenta SolicitudRegistroPuntoVenta `xml:"SolicitudRegistroPuntoVenta"`
}

// SolicitudRegistroPuntoVenta representa los datos para registrar un punto de venta
type SolicitudRegistroPuntoVenta struct {
	CodigoAmbiente       int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad      int    `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoSistema        string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal       int    `xml:"codigoSucursal" json:"codigoSucursal"`
	CodigoTipoPuntoVenta int    `xml:"codigoTipoPuntoVenta" json:"codigoTipoPuntoVenta"`
	Cuis                 string `xml:"cuis" json:"cuis"`
	Descripcion          string `xml:"descripcion" json:"descripcion"`
	Nit                  int64  `xml:"nit" json:"nit"`
	NombrePuntoVenta     string `xml:"nombrePuntoVenta" json:"nombrePuntoVenta"`
}

// RegistroPuntoVentaResponse es el wrapper para la respuesta de registro
type RegistroPuntoVentaResponse struct {
	Respuesta RespuestaRegistroPuntoVenta `xml:"RespuestaRegistroPuntoVenta"`
}

// RespuestaRegistroPuntoVenta representa el resultado del registro
type RespuestaRegistroPuntoVenta struct {
	CodigoPuntoVenta int               `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	MensajesList     []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion      bool              `xml:"transaccion" json:"transaccion"`
}

// ConsultaPuntoVenta es el wrapper para consultar puntos de venta
type ConsultaPuntoVenta struct {
	XMLName                     xml.Name                    `xml:"ns:consultaPuntoVenta"`
	SolicitudConsultaPuntoVenta SolicitudConsultaPuntoVenta `xml:"SolicitudConsultaPuntoVenta"`
}

// SolicitudConsultaPuntoVenta representa los datos para consultar puntos de venta
type SolicitudConsultaPuntoVenta struct {
	CodigoAmbiente int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoSistema  string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis           string `xml:"cuis" json:"cuis"`
	Nit            int64  `xml:"nit" json:"nit"`
}

// ConsultaPuntoVentaResponse es el wrapper para la respuesta de consulta
type ConsultaPuntoVentaResponse struct {
	Respuesta RespuestaConsultaPuntoVenta `xml:"RespuestaConsultaPuntoVenta"`
}

// RespuestaConsultaPuntoVenta representa el resultado de la consulta
type RespuestaConsultaPuntoVenta struct {
	ListaPuntosVentas []PuntosVentasDto `xml:"listaPuntosVentas" json:"listaPuntosVentas"`
	MensajesList      []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion       bool              `xml:"transaccion" json:"transaccion"`
}

// PuntosVentasDto representa la información de un punto de venta
type PuntosVentasDto struct {
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	NombrePuntoVenta string `xml:"nombrePuntoVenta" json:"nombrePuntoVenta"`
	TipoPuntoVenta   string `xml:"tipoPuntoVenta" json:"tipoPuntoVenta"`
}

// CierrePuntoVenta es el wrapper para cerrar un punto de venta
type CierrePuntoVenta struct {
	XMLName                   xml.Name                  `xml:"ns:cierrePuntoVenta"`
	SolicitudCierrePuntoVenta SolicitudCierrePuntoVenta `xml:"SolicitudCierrePuntoVenta"`
}

// SolicitudCierrePuntoVenta representa los datos para cerrar un punto de venta
type SolicitudCierrePuntoVenta struct {
	CodigoAmbiente   int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSistema    string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal   int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis             string `xml:"cuis" json:"cuis"`
	Nit              int64  `xml:"nit" json:"nit"`
}

// CierrePuntoVentaResponse es el wrapper para la respuesta de cierre
type CierrePuntoVentaResponse struct {
	Respuesta RespuestaCierrePuntoVenta `xml:"RespuestaCierrePuntoVenta"`
}

// RespuestaCierrePuntoVenta representa el resultado del cierre
type RespuestaCierrePuntoVenta struct {
	CodigoPuntoVenta int               `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	MensajesList     []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion      bool              `xml:"transaccion" json:"transaccion"`
}

// RegistroPuntoVentaComisionista es el wrapper para registrar un comisionista
type RegistroPuntoVentaComisionista struct {
	XMLName                         xml.Name                        `xml:"ns:registroPuntoVentaComisionista"`
	SolicitudPuntoVentaComisionista SolicitudPuntoVentaComisionista `xml:"SolicitudPuntoVentaComisionista"`
}

// SolicitudPuntoVentaComisionista representa los datos para registrar un comisionista
type SolicitudPuntoVentaComisionista struct {
	CodigoAmbiente   int       `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad  int       `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoSistema    string    `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal   int       `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis             string    `xml:"cuis" json:"cuis"`
	Descripcion      string    `xml:"descripcion" json:"descripcion"`
	FechaFin         time.Time `xml:"fechaFin" json:"fechaFin"`
	FechaInicio      time.Time `xml:"fechaInicio" json:"fechaInicio"`
	Nit              int64     `xml:"nit" json:"nit"`
	NitComisionista  int64     `xml:"nitComisionista" json:"nitComisionista"`
	NombrePuntoVenta string    `xml:"nombrePuntoVenta" json:"nombrePuntoVenta"`
	NumeroContrato   string    `xml:"numeroContrato" json:"numeroContrato"`
}

// RegistroPuntoVentaComisionistaResponse es el wrapper para la respuesta de comisionista
type RegistroPuntoVentaComisionistaResponse struct {
	Respuesta RespuestaPuntoVentaComisionista `xml:"RespuestaPuntoVentaComisionista"`
}

// RespuestaPuntoVentaComisionista representa el resultado del registro de comisionista
type RespuestaPuntoVentaComisionista struct {
	CodigoPuntoVenta int               `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	MensajesList     []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion      bool              `xml:"transaccion" json:"transaccion"`
}
