package documents

import "encoding/xml"

// ConfirmacionCompra representa el XML individual que se incluye en un archivo
// de confirmaciones de compras enviado al SIAT.
type ConfirmacionCompra struct {
	XMLName            xml.Name `xml:"confirmacionCompra" json:"-"`
	Nro                int      `xml:"nro" json:"nro"`
	NitEmisor          int64    `xml:"nitEmisor" json:"nitEmisor"`
	CodigoAutorizacion string   `xml:"codigoAutorizacion" json:"codigoAutorizacion"`
	NumeroFactura      string   `xml:"numeroFactura" json:"numeroFactura"`
	TipoCompra         int      `xml:"tipoCompra" json:"tipoCompra"`
}
