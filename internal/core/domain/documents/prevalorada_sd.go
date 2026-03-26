package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaPrevaloradaSd representa la solicitud para el Sector 35 (Prevalorada SD).
type FacturaPrevaloradaSd struct {
	XMLName           xml.Name              `json:"-"`
	XmlnsXsi          string                `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraPrevaloradaSd `xml:"cabecera" json:"cabecera"`
	Detalle           DetallePrevaloradaSd  `xml:"detalle" json:"detalle"`
}

// CabeceraPrevaloradaSd contiene los datos generales de la factura prevalorada.
type CabeceraPrevaloradaSd struct {
	NitEmisor                    int64                    `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string                   `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string                   `xml:"municipio" json:"municipio"`
	Telefono                     datatype.Nilable[string] `xml:"telefono,omitempty" json:"telefono,omitempty"`
	NumeroFactura                int64                    `xml:"numeroFactura" json:"numeroFactura"`
	Cuf                          string                   `xml:"cuf" json:"cuf"`
	Cufd                         string                   `xml:"cufd" json:"cufd"`
	CodigoSucursal               int                      `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                    string                   `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta             datatype.Nilable[int]    `xml:"codigoPuntoVenta,omitempty" json:"codigoPuntoVenta,omitempty"`
	FechaEmision                 datatype.TimeSiat        `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            string                   `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int                      `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              int64                    `xml:"numeroDocumento" json:"numeroDocumento"`
	CodigoCliente                string                   `xml:"codigoCliente" json:"codigoCliente"`
	CodigoMetodoPago             int                      `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]  `xml:"numeroTarjeta,omitempty" json:"numeroTarjeta,omitempty"`
	MontoTotal                   float64                  `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64                  `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                 int                      `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                  `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                  `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	Leyenda                      string                   `xml:"leyenda" json:"leyenda"`
	Usuario                      string                   `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                      `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetallePrevaloradaSd representa un ítem de la factura prevalorada.
type DetallePrevaloradaSd struct {
	ActividadEconomica string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion        string                    `xml:"descripcion" json:"descripcion"`
	Cantidad           float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario     float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento,omitempty" json:"montoDescuento,omitempty"`
	SubTotal           float64                   `xml:"subTotal" json:"subTotal"`
}
