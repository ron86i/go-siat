package codigos

import "encoding/xml"

// SolicitudCUIS contiene los parámetros para solicitar un Código Único de Inicio de Sistemas (CUIS).
type SolicitudCUIS struct {
	CodigoModalidad  int
	CodigoPuntoVenta *int
	CodigoSucursal   int
}

// SolicitudCuis es un alias para SolicitudCUIS.
type SolicitudCuis = SolicitudCUIS

// SolicitudCUFD contiene los parámetros para solicitar un Código Único de Facturación Diaria (CUFD).
type SolicitudCUFD struct {
	CodigoModalidad  int
	CodigoPuntoVenta *int
	CodigoSucursal   int
	Cuis             string
}

// SolicitudCufd es un alias para SolicitudCUFD.
type SolicitudCufd = SolicitudCUFD

type solicitudSoap struct {
	XMLName          xml.Name
	CodigoAmbiente   int     `xml:"codigoAmbiente"`
	CodigoModalidad  int     `xml:"codigoModalidad"`
	CodigoPuntoVenta *int    `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema    string  `xml:"codigoSistema"`
	CodigoSucursal   int     `xml:"codigoSucursal"`
	Cuis             *string `xml:"cuis,omitempty"`
	Nit              int64   `xml:"nit"`
}
