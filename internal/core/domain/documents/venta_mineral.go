package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
)

// FacturaVentaMineral representa la estructura completa de una factura de Venta de Minerales para el SIAT.
type FacturaVentaMineral struct {
	XMLName           xml.Name              `json:"-"`
	XmlnsXsi          string                `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraVentaMineral  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleVentaMineral `xml:"detalle" json:"detalle"`
}

// CabeceraVentaMineral contiene la información general, del comprador y los datos técnicos de la exportación de minerales.
type CabeceraVentaMineral struct {
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
	Nim                          datatype.Nilable[string]  `xml:"nim" json:"nim"`
	DescripcionProducto          string                    `xml:"descripcionProducto" json:"descripcionProducto"`
	Localidad                    string                    `xml:"localidad" json:"localidad"`
	DepartamentoOrigen           datatype.Nilable[string]  `xml:"departamentoOrigen" json:"departamentoOrigen"`
	MunicipioOrigen              datatype.Nilable[string]  `xml:"municipioOrigen" json:"municipioOrigen"`
	CodigoMoneda                 int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio" json:"tipoCambio"`
	NumeroLote                   string                    `xml:"numeroLote" json:"numeroLote"`
	NumeroContrato               string                    `xml:"numeroContrato" json:"numeroContrato"`
	KilosNetosHumedos            float64                   `xml:"kilosNetosHumedos" json:"kilosNetosHumedos"`
	HumedadPorcentaje            datatype.Nilable[float64] `xml:"humedadPorcentaje" json:"humedadPorcentaje"`
	HumedadValor                 datatype.Nilable[float64] `xml:"humedadValor" json:"humedadValor"`
	MermaPorcentaje              datatype.Nilable[float64] `xml:"mermaPorcentaje" json:"mermaPorcentaje"`
	MermaValor                   datatype.Nilable[float64] `xml:"mermaValor" json:"mermaValor"`
	KilosNetosSecos              float64                   `xml:"kilosNetosSecos" json:"kilosNetosSecos"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	SubTotalNeto                 float64                   `xml:"subTotalNeto" json:"subTotalNeto"`
	TotalPrecioNeto              float64                   `xml:"totalPrecioNeto" json:"totalPrecioNeto"`
	Iva                          float64                   `xml:"iva" json:"iva"`
	MontoAPagar                  float64                   `xml:"montoaPagar" json:"montoaPagar"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	GastosRealizacion            float64                   `xml:"gastosRealizacion" json:"gastosRealizacion"`
	LiquidacionPreliminar        datatype.Nilable[float64] `xml:"liquidacionPreliminar" json:"liquidacionPreliminar"`
	OtrosDatos                   datatype.Nilable[string]  `xml:"otrosDatos" json:"otrosDatos"`
	MontoGiftCard                datatype.Nilable[float64] `xml:"montoGiftCard" json:"montoGiftCard"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	CodigoInterno                string                    `xml:"codigoInterno" json:"codigoInterno"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleVentaMineral representa un ítem de mineral detallando su ley y cantidades de extracción.
type DetalleVentaMineral struct {
	ActividadEconomica     string  `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin      int64   `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto         string  `xml:"codigoProducto" json:"codigoProducto"`
	CodigoNandina          string  `xml:"codigoNandina" json:"codigoNandina"`
	Descripcion            string  `xml:"descripcion" json:"descripcion"`
	DescripcionLeyes       string  `xml:"descripcionLeyes" json:"descripcionLeyes"`
	CantidadExtraccion     float64 `xml:"cantidadExtraccion" json:"cantidadExtraccion"`
	Cantidad               float64 `xml:"cantidad" json:"cantidad"`
	UnidadMedidaExtraccion int     `xml:"unidadMedidaExtraccion" json:"unidadMedidaExtraccion"`
	UnidadMedida           int     `xml:"unidadMedida" json:"unidadMedida"`
	PrecioNetoUnitario     float64 `xml:"precioNetoUnitario" json:"precioNetoUnitario"`
	MontoDescuento         float64 `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal               float64 `xml:"subTotal" json:"subTotal"`
}
