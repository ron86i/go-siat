package documentos

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaLibreConsignacion representa la estructura completa de una factura Libre Consignación para el SIAT.
type FacturaLibreConsignacion struct {
	XMLName           xml.Name                   `json:"-"`
	XmlnsXsi          string                     `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                     `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraLibreConsignacion  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleLibreConsignacion `xml:"detalle" json:"detalle"`
}

// CabeceraLibreConsignacion contiene la información general y del cliente de la factura Libre Consignación.
type CabeceraLibreConsignacion struct {
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
	NumeroDocumento              string                   `xml:"numeroDocumento" json:"numeroDocumento"`
	CodigoCliente                string                   `xml:"codigoCliente" json:"codigoCliente"`
	CodigoPais                   int                      `xml:"codigoPais" json:"codigoPais"`
	PuertoDestino                string                   `xml:"puertoDestino" json:"puertoDestino"`
	CodigoMetodoPago             int                      `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]  `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                  `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64                  `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                 int                      `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                  `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                  `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	CodigoExcepcion              datatype.Nilable[int]    `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Cafc                         datatype.Nilable[string] `xml:"cafc" json:"cafc"`
	Leyenda                      string                   `xml:"leyenda" json:"leyenda"`
	Usuario                      string                   `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                      `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleLibreConsignacion representa un ítem o servicio dentro de la factura Libre Consignación.
type DetalleLibreConsignacion struct {
	ActividadEconomica string  `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  int64   `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string  `xml:"codigoProducto" json:"codigoProducto"`
	CodigoNandina      string  `xml:"codigoNandina" json:"codigoNandina"`
	Descripcion        string  `xml:"descripcion" json:"descripcion"`
	UnidadMedida       int     `xml:"unidadMedida" json:"unidadMedida"`
	Cantidad           float64 `xml:"cantidad" json:"cantidad"`
	PrecioUnitario     float64 `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento     float64 `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64 `xml:"subTotal" json:"subTotal"`
}
