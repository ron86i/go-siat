package sincronizacion

import "encoding/xml"

// SolicitudSincronizacion identifica el establecimiento y su CUIS para consultar catálogos.
type SolicitudSincronizacion struct {
	CodigoPuntoVenta *int
	CodigoSucursal   int
	Cuis             string
}

// Solicitud es un alias para SolicitudSincronizacion.
type Solicitud = SolicitudSincronizacion

type solicitudXML struct {
	XMLName          xml.Name `xml:"SolicitudSincronizacion"`
	CodigoAmbiente   int      `xml:"codigoAmbiente"`
	CodigoPuntoVenta *int     `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema    string   `xml:"codigoSistema"`
	CodigoSucursal   int      `xml:"codigoSucursal"`
	Cuis             string   `xml:"cuis"`
	Nit              int64    `xml:"nit"`
}
