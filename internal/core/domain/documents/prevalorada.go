package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaPrevalorada representa la estructura completa de una factura Prevalorada para el SIAT.
type FacturaPrevalorada struct {
	XMLName           xml.Name            `json:"-"`
	XmlnsXsi          string              `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string              `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraPrevalorada `xml:"cabecera" json:"cabecera"`
	Detalle           DetallePrevalorada  `xml:"detalle" json:"detalle"`
}

// CabeceraPrevalorada contiene la información general y del cliente de la factura Prevalorada.
type CabeceraPrevalorada struct {
	NitEmisor                    int64                    `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string                   `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string                   `xml:"municipio" json:"municipio"`
	Telefono                     datatype.Nilable[string] `xml:"telefono" json:"telefono"`
	NumeroFactura                int64                    `xml:"numeroFactura" json:"numeroFactura"`
	Cuf                          string                   `xml:"cuf" json:"cuf"`
	Cufd                         string                   `xml:"cufd" json:"cufd"`
	CodigoSucursal               int                      `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                    string                   `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta             datatype.Nilable[int]    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	FechaEmision                 datatype.TimeSiat        `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            string                   `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int                      `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              int64                    `xml:"numeroDocumento" json:"numeroDocumento"`
	CodigoCliente                string                   `xml:"codigoCliente" json:"codigoCliente"`
	CodigoMetodoPago             int                      `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]  `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                  `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64                  `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                 int                      `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                  `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                  `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	Leyenda                      string                   `xml:"leyenda" json:"leyenda"`
	Usuario                      string                   `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                      `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetallePrevalorada representa el ítem único dentro de la factura Prevalorada.
type DetallePrevalorada struct {
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
