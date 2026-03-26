package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaComercialExportacionMinera representa la estructura completa de una factura de Comercial Exportación Minera para el SIAT.
type FacturaComercialExportacionMinera struct {
	XMLName           xml.Name                            `json:"-"`
	XmlnsXsi          string                              `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                              `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraComercialExportacionMinera  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleComercialExportacionMinera `xml:"detalle" json:"detalle"`
}

// CabeceraComercialExportacionMinera contiene la información general y del cliente de la factura de Comercial Exportación Minera.
type CabeceraComercialExportacionMinera struct {
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
	DireccionComprador           string                    `xml:"direccionComprador" json:"direccionComprador"`
	CodigoTipoDocumentoIdentidad int                       `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string                    `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                  datatype.Nilable[string]  `xml:"complemento" json:"complemento"`
	Ruex                         datatype.Nilable[string]  `xml:"ruex" json:"ruex"`
	Nim                          datatype.Nilable[string]  `xml:"nim" json:"nim"`
	ConcentradoGranel            string                    `xml:"concentradoGranel" json:"concentradoGranel"`
	Origen                       string                    `xml:"origen" json:"origen"`
	PuertoTransito               datatype.Nilable[string]  `xml:"puertoTransito" json:"puertoTransito"`
	PuertoDestino                string                    `xml:"puertoDestino" json:"puertoDestino"`
	PaisDestino                  int                       `xml:"paisDestino" json:"paisDestino"`
	Incoterm                     string                    `xml:"incoterm" json:"incoterm"`
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	MontoTotalSujetoIva          float64                   `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                 int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio" json:"tipoCambio"`
	TipoCambioANB                float64                   `xml:"tipoCambioANB" json:"tipoCambioANB"`
	NumeroLote                   string                    `xml:"numeroLote" json:"numeroLote"`
	KilosNetosHumedos            float64                   `xml:"kilosNetosHumedos" json:"kilosNetosHumedos"`
	HumedadPorcentaje            datatype.Nilable[float64] `xml:"humedadPorcentaje" json:"humedadPorcentaje"`
	HumedadValor                 datatype.Nilable[float64] `xml:"humedadValor" json:"humedadValor"`
	MermaPorcentaje              datatype.Nilable[float64] `xml:"mermaPorcentaje" json:"mermaPorcentaje"`
	MermaValor                   datatype.Nilable[float64] `xml:"mermaValor" json:"mermaValor"`
	KilosNetosSecos              float64                   `xml:"kilosNetosSecos" json:"kilosNetosSecos"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                   `xml:"montoTotal" json:"montoTotal"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	GastosRealizacion            float64                   `xml:"gastosRealizacion" json:"gastosRealizacion"`
	OtrosDatos                   datatype.Nilable[string]  `xml:"otrosDatos" json:"otrosDatos"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Cafc                         datatype.Nilable[string]  `xml:"cafc" json:"cafc"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleComercialExportacionMinera representa un ítem o servicio dentro de la factura de Comercial Exportación Minera.
type DetalleComercialExportacionMinera struct {
	ActividadEconomica     string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin      int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto         string                    `xml:"codigoProducto" json:"codigoProducto"`
	CodigoNandina          string                    `xml:"codigoNandina" json:"codigoNandina"`
	Descripcion            string                    `xml:"descripcion" json:"descripcion"`
	DescripcionLeyes       string                    `xml:"descripcionLeyes" json:"descripcionLeyes"`
	CantidadExtraccion     float64                   `xml:"cantidadExtraccion" json:"cantidadExtraccion"`
	Cantidad               float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedidaExtraccion int                       `xml:"unidadMedidaExtraccion" json:"unidadMedidaExtraccion"`
	UnidadMedida           int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario         float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento         datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal               float64                   `xml:"subTotal" json:"subTotal"`
}
