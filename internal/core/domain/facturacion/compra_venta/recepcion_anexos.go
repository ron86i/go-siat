package compra_venta

import "encoding/xml"

type RecepcionAnexos struct {
	XMLName xml.Name `xml:"ns:recepcionAnexos" json:"-"`
}

type SolicitudRecepcionAnexos struct {
	SolicitudRecepcion
	AnexosList []VentaAnexo `xml:"anexosList" json:"anexosList"`
	Cuf        string       `xml:"cuf" json:"cuf"`
}

type VentaAnexo struct {
	Codigo            string `xml:"codigo" json:"codigo"`
	CodigoProducto    string `xml:"codigoProducto" json:"codigoProducto"`
	CodigoProductoSin int64  `xml:"codigoProductoSin" json:"codigoProductoSin"`
	TipoCodigo        string `xml:"tipoCodigo" json:"tipoCodigo"`
}
