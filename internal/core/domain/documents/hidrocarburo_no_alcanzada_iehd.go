package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
)

// FacturaHidrocarburoNoAlcanzadaIehd representa la estructura completa de una factura de
// Hidrocarburos No Alcanzada por el IEHD (Sector 38) para el SIAT.
type FacturaHidrocarburoNoAlcanzadaIehd struct {
	XMLName           xml.Name                             `json:"-"`
	XmlnsXsi          string                               `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                               `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraHidrocarburoNoAlcanzadaIehd  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleHidrocarburoNoAlcanzadaIehd `xml:"detalle" json:"detalle"`
}

// CabeceraHidrocarburoNoAlcanzadaIehd contiene la información general y del cliente de la factura.
// El orden de los campos es EXTREMADAMENTE CRÍTICO y sigue el XSD ver.11/03/2022.
//
// Es idéntica a CabeceraHidrocarburoAlcanzadaIehd salvo por la ausencia de MontoIehd,
// que no aplica a las actividades exentas del pago del IEHD.
type CabeceraHidrocarburoNoAlcanzadaIehd struct {
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
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	Ciudad                       datatype.Nilable[string]  `xml:"ciudad" json:"ciudad"`
	NombrePropietario            datatype.Nilable[string]  `xml:"nombrePropietario" json:"nombrePropietario"`
	NombreRepresentanteLegal     datatype.Nilable[string]  `xml:"nombreRepresentanteLegal" json:"nombreRepresentanteLegal"`
	CondicionPago                datatype.Nilable[string]  `xml:"condicionPago" json:"condicionPago"`
	PeriodoEntrega               datatype.Nilable[string]  `xml:"periodoEntrega" json:"periodoEntrega"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	SubTotalNeto                 float64                   `xml:"subTotalNeto" json:"subTotalNeto"`
	TotalPrecioNeto              float64                   `xml:"totalPrecioNeto" json:"totalPrecioNeto"`
	Iva                          float64                   `xml:"iva" json:"iva"`
	MontoAPagar                  float64                   `xml:"montoaPagar" json:"montoaPagar"`
	CodigoMoneda                 int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	CodigoInterno                string                    `xml:"codigoInterno" json:"codigoInterno"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleHidrocarburoNoAlcanzadaIehd representa un ítem dentro de la factura.
// Es idéntico a DetalleHidrocarburoAlcanzadaIehd salvo por la ausencia de PorcentajeIehd.
type DetalleHidrocarburoNoAlcanzadaIehd struct {
	ActividadEconomica string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion        string                    `xml:"descripcion" json:"descripcion"`
	Cantidad           float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioNetoUnitario float64                   `xml:"precioNetoUnitario" json:"precioNetoUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64                   `xml:"subTotal" json:"subTotal"`
}
