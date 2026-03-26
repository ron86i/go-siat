package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaVentaCombustibleSinSubvencion representa la estructura completa de una factura de Venta de Combustible Sin Subvención para el SIAT.
type FacturaVentaCombustibleSinSubvencion struct {
	XMLName           xml.Name                               `json:"-"`
	XmlnsXsi          string                                 `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                                 `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraVentaCombustibleSinSubvencion  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleVentaCombustibleSinSubvencion `xml:"detalle" json:"detalle"`
}

// CabeceraVentaCombustibleSinSubvencion contiene la información general y del cliente de la factura de Combustible Sin Subvención.
type CabeceraVentaCombustibleSinSubvencion struct {
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
	Complemento                  datatype.Nilable[string]  `xml:"complemento" json:"complemento"`
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	CodigoPais                   datatype.Nilable[int]     `xml:"codigoPais" json:"codigoPais"`
	PlacaVehiculo                string                    `xml:"placaVehiculo" json:"placaVehiculo"`
	TipoEnvase                   datatype.Nilable[string]  `xml:"tipoEnvase" json:"tipoEnvase"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                   `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64                   `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoAutorizacionSC         datatype.Nilable[string]  `xml:"codigoAutorizacionSC" json:"codigoAutorizacionSC"`
	Observacion                  datatype.Nilable[string]  `xml:"observacion" json:"observacion"`
	CodigoMoneda                 int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Cafc                         datatype.Nilable[string]  `xml:"cafc" json:"cafc"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleVentaCombustibleSinSubvencion representa un ítem o servicio dentro de la factura de Combustible Sin Subvención.
type DetalleVentaCombustibleSinSubvencion struct {
	ActividadEconomica string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion        string                    `xml:"descripcion" json:"descripcion"`
	Cantidad           float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario     float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64                   `xml:"subTotal" json:"subTotal"`
}
