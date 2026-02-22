package compra_venta

import (
	"encoding/xml"
	"go-siat/internal/core/domain/datatype"
)

type RecepcionFactura struct {
	XMLName                           xml.Name                          `xml:"ns:recepcionFactura" json:"-"`
	SolicitudServicioRecepcionFactura SolicitudServicioRecepcionFactura `xml:"SolicitudServicioRecepcionFactura" json:"-"`
}
type SolicitudRecepcion struct {
	CodigoAmbiente        int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
	CodigoEmision         int    `xml:"codigoEmision" json:"codigoEmision"`
	CodigoModalidad       int    `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoPuntoVenta      int    `xml:"codigoPuntoVenta,omitempty" json:"codigoPuntoVenta" `
	CodigoSistema         string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal        int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cufd                  string `xml:"cufd" json:"cufd"`
	Cuis                  string `xml:"cuis" json:"cuis"`
	Nit                   int64  `xml:"nit" json:"nit"`
	TipoFacturaDocumento  int    `xml:"tipoFacturaDocumento" json:"tipoFacturaDocumento"`
}

type SolicitudServicioRecepcionFactura struct {
	SolicitudRecepcion
	//FacturaCompraVenta domain.FacturaCompraVenta `xml:"-" json:"facturaCompraVenta"`
	Archivo     []byte            `xml:"archivo" json:"archivo"`       // se serializa en base64 automáticamente
	FechaEnvio  datatype.TimeSiat `xml:"fechaEnvio" json:"fechaEnvio"` // enviar como RFC3339 string
	HashArchivo string            `xml:"hashArchivo" json:"hashArchivo"`
}
