package compra_venta

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

type RecepcionFactura struct {
	XMLName                           xml.Name                  `xml:"ns:recepcionFactura" json:"-"`
	SolicitudServicioRecepcionFactura SolicitudRecepcionFactura `xml:"SolicitudServicioRecepcionFactura" json:"solicitudServicioRecepcionFactura"`
}

// SolicitudRecepcion
type SolicitudRecepcion struct {
	CodigoAmbiente        int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
	CodigoEmision         int    `xml:"codigoEmision" json:"codigoEmision"`
	CodigoModalidad       int    `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoPuntoVenta      int    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSistema         string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal        int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cufd                  string `xml:"cufd" json:"cufd"`
	Cuis                  string `xml:"cuis" json:"cuis"`
	Nit                   int64  `xml:"nit" json:"nit"`
	TipoFacturaDocumento  int    `xml:"tipoFacturaDocumento" json:"tipoFacturaDocumento"`
}

type SolicitudRecepcionFactura struct {
	SolicitudRecepcion
	Archivo     string            `xml:"archivo" json:"archivo"`
	FechaEnvio  datatype.TimeSiat `xml:"fechaEnvio" json:"fechaEnvio"`
	HashArchivo string            `xml:"hashArchivo" json:"hashArchivo"`
}

type RecepcionFacturaResponse struct {
	XMLName                      xml.Name           `xml:"recepcionFacturaResponse" json:"-"`
	RespuestaServicioFacturacion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"respuestaServicioFacturacion"`
}
