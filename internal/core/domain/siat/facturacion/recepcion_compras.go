package facturacion

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// SolicitudCompras representa la base de las solicitudes para el servicio de recepción de compras.
type SolicitudCompras struct {
	CodigoAmbiente   int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSistema    string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal   int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cufd             string `xml:"cufd" json:"cufd"`
	Cuis             string `xml:"cuis" json:"cuis"`
	Nit              int64  `xml:"nit" json:"nit"`
}

// RecepcionPaqueteCompras representa el envoltorio para la recepción de paquetes de compras.
type RecepcionPaqueteCompras struct {
	XMLName                   xml.Name                  `xml:"ns:recepcionPaqueteCompras" json:"-"`
	SolicitudRecepcionCompras SolicitudRecepcionCompras `xml:"SolicitudRecepcionCompras" json:"SolicitudRecepcionCompras"`
}

// RecepcionPaqueteComprasResponse representa la respuesta a la recepción de paquetes de compras.
type RecepcionPaqueteComprasResponse struct {
	RespuestaRecepcion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"RespuestaServicioFacturacion"`
}

// ValidacionRecepcionPaqueteCompras representa el envoltorio para la validación de paquetes de compras.
type ValidacionRecepcionPaqueteCompras struct {
	XMLName                            xml.Name                           `xml:"ns:validacionRecepcionPaqueteCompras" json:"-"`
	SolicitudValidacionRecepcionCompras SolicitudValidacionRecepcionCompras `xml:"SolicitudValidacionRecepcionCompras" json:"SolicitudValidacionRecepcionCompras"`
}

// ValidacionRecepcionPaqueteComprasResponse representa la respuesta a la validación de paquetes de compras.
type ValidacionRecepcionPaqueteComprasResponse struct {
	RespuestaRecepcion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"RespuestaServicioFacturacion"`
}

// AnulacionCompra representa el envoltorio para la anulación de una compra.
type AnulacionCompra struct {
	XMLName                  xml.Name                 `xml:"ns:anulacionCompra" json:"-"`
	SolicitudAnulacionCompra SolicitudAnulacionCompra `xml:"SolicitudAnulacionCompra" json:"SolicitudAnulacionCompra"`
}

// AnulacionCompraResponse representa la respuesta a la anulación de una compra.
type AnulacionCompraResponse struct {
	RespuestaRecepcion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"RespuestaServicioFacturacion"`
}

// ConfirmacionCompras representa el envoltorio para la confirmación de compras.
type ConfirmacionCompras struct {
	XMLName                     xml.Name                  `xml:"ns:confirmacionCompras" json:"-"`
	SolicitudConfirmacionCompras SolicitudRecepcionCompras `xml:"SolicitudConfirmacionCompras" json:"SolicitudConfirmacionCompras"`
}

// ConfirmacionComprasResponse representa la respuesta a la confirmación de compras.
type ConfirmacionComprasResponse struct {
	RespuestaRecepcion RespuestaRecepcion `xml:"RespuestaServicioFacturacion" json:"RespuestaServicioFacturacion"`
}

// ConsultaCompras representa el envoltorio para la consulta de compras.
type ConsultaCompras struct {
	XMLName                  xml.Name                 `xml:"ns:consultaCompras" json:"-"`
	SolicitudConsultaCompras SolicitudConsultaCompras `xml:"SolicitudConsultaCompras" json:"SolicitudConsultaCompras"`
}

// ConsultaComprasResponse representa la respuesta a la consulta de compras.
type ConsultaComprasResponse struct {
	RespuestaConsultaCompras RespuestaConsultaCompras `xml:"RespuestaServicioFacturacion" json:"RespuestaServicioFacturacion"`
}

// SolicitudRecepcionCompras representa los datos necesarios para la recepción de compras.
type SolicitudRecepcionCompras struct {
	SolicitudCompras
	Archivo          string            `xml:"archivo" json:"archivo"`
	CantidadFacturas int               `xml:"cantidadFacturas" json:"cantidadFacturas"`
	FechaEnvio       datatype.TimeSiat `xml:"fechaEnvio" json:"fechaEnvio"`
	Gestion          int               `xml:"gestion" json:"gestion"`
	HashArchivo      string            `xml:"hashArchivo" json:"hashArchivo"`
	Periodo          int               `xml:"periodo" json:"periodo"`
}

// SolicitudValidacionRecepcionCompras representa los datos para validar la recepción de un paquete de compras.
type SolicitudValidacionRecepcionCompras struct {
	SolicitudCompras
	CodigoRecepcion string `xml:"codigoRecepcion" json:"codigoRecepcion"`
}

// SolicitudAnulacionCompra representa los datos para anular una compra.
type SolicitudAnulacionCompra struct {
	SolicitudCompras
	CodAutorizacion string `xml:"codAutorizacion" json:"codAutorizacion"`
	NitProveedor    int64  `xml:"nitProveedor" json:"nitProveedor"`
	NroFactura      int64  `xml:"nroFactura" json:"nroFactura"`
	NroDuiDim       string `xml:"nroDuiDim" json:"nroDuiDim"`
}

// SolicitudConsultaCompras representa los datos para consultar compras.
type SolicitudConsultaCompras struct {
	SolicitudCompras
	Fecha datatype.TimeSiat `xml:"fecha" json:"fecha"`
}

// RespuestaConsultaCompras representa el detalle de la consulta de compras.
type RespuestaConsultaCompras struct {
	Archivo           string            `xml:"archivo" json:"archivo"`
	CodigoDescripcion string            `xml:"codigoDescripcion" json:"codigoDescripcion"`
	CodigoEstado      int               `xml:"codigoEstado" json:"codigoEstado"`
	MensajesList      []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion       bool              `xml:"transaccion" json:"transaccion"`
}

// VerificarComunicacionRecepcionCompras envoltorio para verificar comunicacion.
type VerificarComunicacionRecepcionCompras struct {
	XMLName xml.Name `xml:"ns:verificarComunicacion" json:"-"`
}
