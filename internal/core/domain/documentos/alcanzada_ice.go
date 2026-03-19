package documentos

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaAlcanzadaIce representa la estructura de una factura del Sector 14
// (Factura Alcanzada por el ICE).
type FacturaAlcanzadaIce struct {
	XMLName           xml.Name              `json:"-"`
	XmlnsXsi          string                `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraAlcanzadaIce  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleAlcanzadaIce `xml:"detalle" json:"detalle"`
}

// CabeceraAlcanzadaIce contiene la información general y los montos ICE totales.
type CabeceraAlcanzadaIce struct {
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
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                   `xml:"montoTotal" json:"montoTotal"`
	MontoIceEspecifico           datatype.Nilable[float64] `xml:"montoIceEspecifico" json:"montoIceEspecifico"`
	MontoIcePorcentual           datatype.Nilable[float64] `xml:"montoIcePorcentual" json:"montoIcePorcentual"`
	MontoTotalSujetoIva          float64                   `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
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

// DetalleAlcanzadaIce representa un ítem individual, incluyendo la lógica de ICE específico y porcentual.
type DetalleAlcanzadaIce struct {
	ActividadEconomica string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion        string                    `xml:"descripcion" json:"descripcion"`
	Cantidad           float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario     float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64                   `xml:"subTotal" json:"subTotal"`
	MarcaIce           int                       `xml:"marcaIce" json:"marcaIce"`
	AlicuotaIva        datatype.Nilable[float64] `xml:"alicuotaIva" json:"alicuotaIva"`
	PrecioNetoVentaIce datatype.Nilable[float64] `xml:"precioNetoVentaIce" json:"precioNetoVentaIce"`
	AlicuotaEspecifica datatype.Nilable[float64] `xml:"alicuotaEspecifica" json:"alicuotaEspecifica"`
	AlicuotaPorcentual datatype.Nilable[float64] `xml:"alicuotaPorcentual" json:"alicuotaPorcentual"`
	MontoIceEspecifico datatype.Nilable[float64] `xml:"montoIceEspecifico" json:"montoIceEspecifico"`
	MontoIcePorcentual datatype.Nilable[float64] `xml:"montoIcePorcentual" json:"montoIcePorcentual"`
	CantidadIce        datatype.Nilable[float64] `xml:"cantidadIce" json:"cantidadIce"`
}
