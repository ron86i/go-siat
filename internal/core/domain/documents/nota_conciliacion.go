package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
)

// NotaConciliacion representa la estructura completa de una nota electrónica de conciliación (Sector 29).
type NotaConciliacion struct {
	XMLName             xml.Name              `json:"-"`
	XmlnsXsi            string                `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation   string                `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera            CabeceraConciliacion  `xml:"cabecera" json:"cabecera"`
	DetalleOriginal     []DetalleOriginal     `xml:"detalleOriginal" json:"detalleOriginal"`
	DetalleConciliacion []DetalleConciliacion `xml:"detalleConciliacion" json:"detalleConciliacion"`
}

// CabeceraConciliacion contiene la información general de la nota de conciliación.
type CabeceraConciliacion struct {
	NitEmisor                    int64                    `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string                   `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string                   `xml:"municipio" json:"municipio"`
	Telefono                     datatype.Nilable[string] `xml:"telefono" json:"telefono"`
	NumeroNotaConciliacion       int64                    `xml:"numeroNotaConciliacion" json:"numeroNotaConciliacion"`
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
	CodigoControl                datatype.Nilable[string] `xml:"codigoControl" json:"codigoControl"`
	FechaEmisionFactura          datatype.TimeSiat        `xml:"fechaEmisionFactura" json:"fechaEmisionFactura"`
	TotalPrecioNetoOriginal      float64                  `xml:"totalPrecioNetoOriginal" json:"totalPrecioNetoOriginal"`
	MontoTotalConciliado         float64                  `xml:"montoTotalConciliado" json:"montoTotalConciliado"`
	IvaConciliado                float64                  `xml:"ivaConciliado" json:"ivaConciliado"`
	MontoActualizacion           float64                  `xml:"montoActualizacion" json:"montoActualizacion"`
	CreditoFiscalIva             float64                  `xml:"creditoFiscalIva" json:"creditoFiscalIva"`
	DebitoFiscalIva              float64                  `xml:"debitoFiscalIva" json:"debitoFiscalIva"`
	UfvFactura                   float64                  `xml:"ufvFactura" json:"ufvFactura"`
	UfvConciliacion              float64                  `xml:"ufvConciliacion" json:"ufvConciliacion"`
	CodigoExcepcion              datatype.Nilable[int]    `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Leyenda                      string                   `xml:"leyenda" json:"leyenda"`
	Usuario                      string                   `xml:"usuario" json:"usuario"`
	CodigoInterno                string                   `xml:"codigoInterno" json:"codigoInterno"`
	CodigoDocumentoSector        int                      `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleOriginal representa un ítem de la factura original.
type DetalleOriginal struct {
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

// DetalleConciliacion representa un ítem conciliado.
type DetalleConciliacion struct {
	ActividadEconomica   string  `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin    int64   `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto       string  `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion          string  `xml:"descripcion" json:"descripcion"`
	TotalPrecioNeto      float64 `xml:"totalPrecioNeto" json:"totalPrecioNeto"`
	TotalPrecioNetoFinal float64 `xml:"totalPrecioNetoFinal" json:"totalPrecioNetoFinal"`
	MontoConciliado      float64 `xml:"montoConciliado" json:"montoConciliado"`
}
