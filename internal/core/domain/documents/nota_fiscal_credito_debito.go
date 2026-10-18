package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
)

// NotaFiscalCreditoDebito representa la estructura de una Nota de Crédito/Débito standard (Sector 24).
type NotaFiscalCreditoDebito struct {
	XMLName           xml.Name                         `json:"-"`
	XmlnsXsi          string                           `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                           `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraNotaFiscalCreditoDebito  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleNotaFiscalCreditoDebito `xml:"detalle" json:"detalle"`
}

// CabeceraNotaFiscalCreditoDebito contiene la información general de la nota (Sector 24).
type CabeceraNotaFiscalCreditoDebito struct {
	NitEmisor                    int64                    `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string                   `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string                   `xml:"municipio" json:"municipio"`
	Telefono                     datatype.Nilable[string] `xml:"telefono" json:"telefono"`
	NumeroNotaCreditoDebito      int64                    `xml:"numeroNotaCreditoDebito" json:"numeroNotaCreditoDebito"`
	Cuf                          string                   `xml:"cuf" json:"cuf"`
	Cufd                         string                   `xml:"cufd" json:"cufd"`
	CodigoSucursal               int                      `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                    string                   `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta             datatype.Nilable[int]    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	FechaEmision                 datatype.TimeSiat        `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            datatype.Nilable[string] `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int                      `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string                   `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                  string                   `xml:"complemento" json:"complemento"`
	CodigoCliente                string                   `xml:"codigoCliente" json:"codigoCliente"`
	NumeroFactura                int64                    `xml:"numeroFactura" json:"numeroFactura"`
	NumeroAutorizacionCuf        string                   `xml:"numeroAutorizacionCuf" json:"numeroAutorizacionCuf"`
	FechaEmisionFactura          datatype.TimeSiat        `xml:"fechaEmisionFactura" json:"fechaEmisionFactura"`
	SubTotalNetoOriginal         float64                  `xml:"subTotalNetoOriginal" json:"subTotalNetoOriginal"`
	TotalPrecioNetoOriginal      float64                  `xml:"totalPrecioNetoOriginal" json:"totalPrecioNetoOriginal"`
	IvaOriginal                  float64                  `xml:"ivaOriginal" json:"ivaOriginal"`
	MontoAPagarOriginal          float64                  `xml:"montoaPagarOriginal" json:"montoaPagarOriginal"`
	SubTotalNetoDevuelto         float64                  `xml:"subTotalNetoDevuelto" json:"subTotalNetoDevuelto"`
	TotalPrecioNetoDevuelto      float64                  `xml:"totalPrecioNetoDevuelto" json:"totalPrecioNetoDevuelto"`
	IvaDevuelto                  float64                  `xml:"ivaDevuelto" json:"ivaDevuelto"`
	MontoAPagarDevuelto          float64                  `xml:"montoaPagarDevuelto" json:"montoaPagarDevuelto"`
	DescuentoAdicionalOriginal   float64                  `xml:"descuentoAdicionalOriginal" json:"descuentoAdicionalOriginal"`
	MontoDescuentoDevuelto       float64                  `xml:"montoDescuentoDevuelto" json:"montoDescuentoDevuelto"`
	CodigoExcepcion              datatype.Nilable[int]    `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Leyenda                      string                   `xml:"leyenda" json:"leyenda"`
	Usuario                      string                   `xml:"usuario" json:"usuario"`
	CodigoInterno                string                   `xml:"codigoInterno" json:"codigoInterno"`
	CodigoDocumentoSector        int                      `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleNotaFiscalCreditoDebito representa un ítem individual de la nota (Sector 24).
// Nota: A diferencia de otros sectores, este detalle NO contiene nroItem.
type DetalleNotaFiscalCreditoDebito struct {
	NroItem                  int                       `xml:"nroItem" json:"nroItem"`
	ActividadEconomica       string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin        int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto           string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion              string                    `xml:"descripcion" json:"descripcion"`
	Cantidad                 float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida             int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioNetoUnitario       float64                   `xml:"precioNetoUnitario" json:"precioNetoUnitario"`
	MontoDescuento           datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal                 float64                   `xml:"subTotal" json:"subTotal"`
	CodigoDetalleTransaccion int                       `xml:"codigoDetalleTransaccion" json:"codigoDetalleTransaccion"`
}
