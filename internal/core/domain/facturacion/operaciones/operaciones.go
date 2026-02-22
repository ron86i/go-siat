package operaciones

import "encoding/xml"

// CierreOperacionesSistema es el wrapper para cerrar operaciones del sistema
type CierreOperacionesSistema struct {
	XMLName              xml.Name             `xml:"ns:cierreOperacionesSistema"`
	SolicitudOperaciones SolicitudOperaciones `xml:"SolicitudOperaciones"`
}

// SolicitudOperaciones es la base para las solicitudes de operaciones
type SolicitudOperaciones struct {
	CodigoAmbiente   int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad  int    `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta,omitempty" json:"codigoPuntoVenta,omitempty"`
	CodigoSistema    string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal   int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis             string `xml:"cuis" json:"cuis"`
	Nit              int64  `xml:"nit" json:"nit"`
}

// CierreOperacionesSistemaResponse es el wrapper para la respuesta de cierre de sistema
type CierreOperacionesSistemaResponse struct {
	Respuesta RespuestaCierreSistemas `xml:"RespuestaCierreSistemas"`
}

// RespuestaCierreSistemas representa el resultado del cierre de operaciones
type RespuestaCierreSistemas struct {
	CodigoSistema string            `xml:"codigoSistema" json:"codigoSistema"`
	MensajesList  []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion   bool              `xml:"transaccion" json:"transaccion"`
}
