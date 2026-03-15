package compra_venta

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

type RecepcionAnexos struct {
	XMLName                  xml.Name                 `xml:"ns:recepcionAnexos" json:"-"`
	SolicitudRecepcionAnexos SolicitudRecepcionAnexos `xml:"SolicitudRecepcionAnexos" json:"solicitudRecepcionAnexos"`
}

type SolicitudRecepcionAnexos struct {
	facturacion.SolicitudRecepcion
	AnexosList []VentaAnexo `xml:"anexosList" json:"anexosList"`
	Cuf        string       `xml:"cuf" json:"cuf"`
}

type RecepcionAnexosResponse struct {
	RespuestaRecepcionAnexos facturacion.RespuestaRecepcion `xml:"RespuestaRecepcionAnexos" json:"respuestaRecepcionAnexos"`
}

type VentaAnexo struct {
	Codigo            string `xml:"codigo" json:"codigo"`
	CodigoProducto    string `xml:"codigoProducto" json:"codigoProducto"`
	CodigoProductoSin int64  `xml:"codigoProductoSin" json:"codigoProductoSin"`
	TipoCodigo        string `xml:"tipoCodigo" json:"tipoCodigo"`
}
