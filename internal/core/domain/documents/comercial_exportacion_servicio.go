package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
)

// FacturaComercialExportacionServicio representa la estructura completa de una factura comercial
// de exportación de servicios para el SIAT.
type FacturaComercialExportacionServicio struct {
	XMLName           xml.Name                              `json:"-"`
	XmlnsXsi          string                                `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                                `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraComercialExportacionServicio  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleComercialExportacionServicio `xml:"detalle" json:"detalle"`
}

// CabeceraComercialExportacionServicio contiene la información general y del comprador
// de la factura comercial de exportación de servicios.
type CabeceraComercialExportacionServicio struct {
	NitEmisor                    int64                     `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string                    `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string                    `xml:"municipio" json:"municipio"`
	Telefono                     datatype.Nilable[string]  `xml:"telefono" json:"telefono"`
	NumeroFactura                int64                     `xml:"numeroFactura" json:"numeroFactura"`
	Cuf                          string                    `xml:"cuf" json:"cuf"`
	Cufd                         string                    `xml:"cufd" json:"cufd"`
	CodigoSucursal               int                       `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                    string                    `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta             datatype.Nilable[int]     `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	FechaEmision                 datatype.TimeSiat         `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            datatype.Nilable[string]  `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int                       `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string                    `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                  string                    `xml:"complemento" json:"complemento"`
	DireccionComprador           string                    `xml:"direccionComprador" json:"direccionComprador"`
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	LugarDestino                 string                    `xml:"lugarDestino" json:"lugarDestino"`
	CodigoPais                   int                       `xml:"codigoPais" json:"codigoPais"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	SubTotalNeto                 float64                   `xml:"subTotalNeto" json:"subTotalNeto"`
	TotalPrecioNeto              float64                   `xml:"totalPrecioNeto" json:"totalPrecioNeto"`
	Iva                          float64                   `xml:"iva" json:"iva"`
	MontoAPagar                  float64                   `xml:"montoaPagar" json:"montoaPagar"`
	CodigoMoneda                 int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	InformacionAdicional         datatype.Nilable[string]  `xml:"informacionAdicional" json:"informacionAdicional"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	CodigoInterno                string                    `xml:"codigoInterno" json:"codigoInterno"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleComercialExportacionServicio representa un ítem o servicio dentro de la factura.
type DetalleComercialExportacionServicio struct {
	ActividadEconomica string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion        string                    `xml:"descripcion" json:"descripcion"`
	Cantidad           int                       `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioNetoUnitario float64                   `xml:"precioNetoUnitario" json:"precioNetoUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64                   `xml:"subTotal" json:"subTotal"`
}
