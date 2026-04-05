package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// NotaCreditoDebito representa la estructura completa de una nota de crédito, débito y descuento (Sector 47).
type NotaCreditoDebito struct {
	XMLName           xml.Name                   `json:"-"`
	XmlnsXsi          string                     `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                     `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraNotaCreditoDebito  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleNotaCreditoDebito `xml:"detalle" json:"detalle"`
}

// CabeceraNotaCreditoDebito contiene la información general de la nota de crédito/débito.
type CabeceraNotaCreditoDebito struct {
	NitEmisor                    int64                     `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string                    `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string                    `xml:"municipio" json:"municipio"`
	Telefono                     datatype.Nilable[string]  `xml:"telefono" json:"telefono"`
	NumeroNotaCreditoDebito      int64                     `xml:"numeroNotaCreditoDebito" json:"numeroNotaCreditoDebito"`
	Cuf                          string                    `xml:"cuf" json:"cuf"`
	Cufd                         string                    `xml:"cufd" json:"cufd"`
	CodigoSucursal               int                       `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                    string                    `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta             datatype.Nilable[int]     `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	FechaEmision                 datatype.TimeSiat         `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            datatype.Nilable[string]  `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int                       `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string                    `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                  datatype.Nilable[string]  `xml:"complemento" json:"complemento"`
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	NumeroFactura                int64                     `xml:"numeroFactura" json:"numeroFactura"`
	NumeroAutorizacionCuf        string                    `xml:"numeroAutorizacionCuf" json:"numeroAutorizacionCuf"`
	FechaEmisionFactura          datatype.TimeSiat         `xml:"fechaEmisionFactura" json:"fechaEmisionFactura"`
	MontoTotalOriginal           float64                   `xml:"montoTotalOriginal" json:"montoTotalOriginal"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	MontoTotalDevuelto           float64                   `xml:"montoTotalDevuelto" json:"montoTotalDevuelto"`
	MontoDescuentoCreditoDebito  datatype.Nilable[float64] `xml:"montoDescuentoCreditoDebito" json:"montoDescuentoCreditoDebito"`
	MontoEfectivoCreditoDebito   float64                   `xml:"montoEfectivoCreditoDebito" json:"montoEfectivoCreditoDebito"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleNotaCreditoDebito representa un ítem individual de la nota.
type DetalleNotaCreditoDebito struct {
	NroItem                  int                       `xml:"nroItem" json:"nroItem"`
	ActividadEconomica       string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin        int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto           string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion              string                    `xml:"descripcion" json:"descripcion"`
	Cantidad                 float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida             int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario           float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento           datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal                 float64                   `xml:"subTotal" json:"subTotal"`
	CodigoDetalleTransaccion int                       `xml:"codigoDetalleTransaccion" json:"codigoDetalleTransaccion"`
}
