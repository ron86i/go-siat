package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaComercialExportacionPVenta representa la estructura completa de una factura Comercial de Exportación Precio Venta para el SIAT.
type FacturaComercialExportacionPVenta struct {
	XMLName           xml.Name                            `json:"-"`
	XmlnsXsi          string                              `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                              `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraComercialExportacionPVenta  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleComercialExportacionPVenta `xml:"detalle" json:"detalle"`
}

// CabeceraComercialExportacionPVenta contiene la información general y de exportación de la factura Precio Venta.
type CabeceraComercialExportacionPVenta struct {
	NitEmisor                       int64                     `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor               string                    `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                       string                    `xml:"municipio" json:"municipio"`
	Telefono                        datatype.Nilable[string]  `xml:"telefono" json:"telefono"`
	NumeroFactura                   int64                     `xml:"numeroFactura" json:"numeroFactura"`
	Cuf                             string                    `xml:"cuf" json:"cuf"`
	Cufd                            string                    `xml:"cufd" json:"cufd"`
	CodigoSucursal                  int                       `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                       string                    `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta                datatype.Nilable[int]     `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	FechaEmision                    datatype.TimeSiat         `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial               datatype.Nilable[string]  `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad    int                       `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento                 string                    `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                     datatype.Nilable[string]  `xml:"complemento" json:"complemento"`
	DireccionComprador              string                    `xml:"direccionComprador" json:"direccionComprador"`
	CodigoCliente                   string                    `xml:"codigoCliente" json:"codigoCliente"`
	Incoterm                        string                    `xml:"incoterm" json:"incoterm"`
	IncotermDetalle                 string                    `xml:"incotermDetalle" json:"incotermDetalle"`
	PuertoDestino                   string                    `xml:"puertoDestino" json:"puertoDestino"`
	LugarDestino                    string                    `xml:"lugarDestino" json:"lugarDestino"`
	CodigoPais                      int                       `xml:"codigoPais" json:"codigoPais"`
	CodigoMetodoPago                int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                   datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                      float64                   `xml:"montoTotal" json:"montoTotal"`
	CostosGastosNacionales          datatype.Nilable[string]  `xml:"costosGastosNacionales" json:"costosGastosNacionales"`
	TotalGastosNacionalesFob        float64                   `xml:"totalGastosNacionalesFob" json:"totalGastosNacionalesFob"`
	CostosGastosInternacionales     datatype.Nilable[string]  `xml:"costosGastosInternacionales" json:"costosGastosInternacionales"`
	TotalGastosInternacionales      datatype.Nilable[float64] `xml:"totalGastosInternacionales" json:"totalGastosInternacionales"`
	PrecioValorBruto                datatype.Nilable[float64] `xml:"precioValorBruto" json:"precioValorBruto"`
	MontoDetalle                    float64                   `xml:"montoDetalle" json:"montoDetalle"`
	MontoTotalSujetoIva             float64                   `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                    int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                      float64                   `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda                float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	NumeroDescripcionPaquetesBultos datatype.Nilable[string]  `xml:"numeroDescripcionPaquetesBultos" json:"numeroDescripcionPaquetesBultos"`
	InformacionAdicional            datatype.Nilable[string]  `xml:"informacionAdicional" json:"informacionAdicional"`
	DescuentoAdicional              datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion                 datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Cafc                            datatype.Nilable[string]  `xml:"cafc" json:"cafc"`
	Leyenda                         string                    `xml:"leyenda" json:"leyenda"`
	Usuario                         string                    `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector           int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleComercialExportacionPVenta representa un ítem dentro de la factura Precio Venta.
type DetalleComercialExportacionPVenta struct {
	ActividadEconomica string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto" json:"codigoProducto"`
	CodigoNandina      string                    `xml:"codigoNandina" json:"codigoNandina"`
	Descripcion        string                    `xml:"descripcion" json:"descripcion"`
	Cantidad           float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario     float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64                   `xml:"subTotal" json:"subTotal"`
}
