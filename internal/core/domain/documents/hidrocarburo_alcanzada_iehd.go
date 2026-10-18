package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
)

// FacturaHidrocarburoAlcanzadaIehd representa la estructura completa de una factura de
// Hidrocarburos Alcanzada por el IEHD (Sector 19) para el SIAT.
type FacturaHidrocarburoAlcanzadaIehd struct {
	XMLName           xml.Name                           `json:"-"`
	XmlnsXsi          string                             `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                             `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraHidrocarburoAlcanzadaIehd  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleHidrocarburoAlcanzadaIehd `xml:"detalle" json:"detalle"`
}

// CabeceraHidrocarburoAlcanzadaIehd contiene la información general y del cliente de la factura.
// El orden de los campos es EXTREMADAMENTE CRÍTICO y sigue el XSD ver.23/08/2021.
type CabeceraHidrocarburoAlcanzadaIehd struct {
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
	Ciudad                       datatype.Nilable[string]  `xml:"ciudad" json:"ciudad"`
	NombrePropietario            datatype.Nilable[string]  `xml:"nombrePropietario" json:"nombrePropietario"`
	NombreRepresentanteLegal     datatype.Nilable[string]  `xml:"nombreRepresentanteLegal" json:"nombreRepresentanteLegal"`
	CondicionPago                datatype.Nilable[string]  `xml:"condicionPago" json:"condicionPago"`
	PeriodoEntrega               datatype.Nilable[string]  `xml:"periodoEntrega" json:"periodoEntrega"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                   `xml:"montoTotal" json:"montoTotal"`
	MontoIehd                    datatype.Nilable[float64] `xml:"montoIehd" json:"montoIehd"`
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

// DetalleHidrocarburoAlcanzadaIehd representa un ítem dentro de la factura.
type DetalleHidrocarburoAlcanzadaIehd struct {
	ActividadEconomica string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion        string                    `xml:"descripcion" json:"descripcion"`
	Cantidad           float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario     float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64                   `xml:"subTotal" json:"subTotal"`
	PorcentajeIehd     datatype.Nilable[float64] `xml:"porcentajeIehd" json:"porcentajeIehd"`
}
