package documentos

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaTurismoHospedaje representa la estructura completa de una factura de Servicio Turístico y Hospedaje para el SIAT.
type FacturaTurismoHospedaje struct {
	XMLName           xml.Name                  `json:"-"`
	XmlnsXsi          string                    `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                    `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraTurismoHospedaje  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleTurismoHospedaje `xml:"detalle" json:"detalle"`
}

// CabeceraTurismoHospedaje contiene la información general y del cliente de la factura de Turismo y Hospedaje.
type CabeceraTurismoHospedaje struct {
	NitEmisor                    int64                               `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string                              `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string                              `xml:"municipio" json:"municipio"`
	Telefono                     datatype.Nilable[string]            `xml:"telefono" json:"telefono"`
	NumeroFactura                int64                               `xml:"numeroFactura" json:"numeroFactura"`
	Cuf                          string                              `xml:"cuf" json:"cuf"`
	Cufd                         string                              `xml:"cufd" json:"cufd"`
	CodigoSucursal               int                                 `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                    string                              `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta             datatype.Nilable[int]               `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	FechaEmision                 datatype.TimeSiat                   `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            datatype.Nilable[string]            `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int                                 `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string                              `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                  datatype.Nilable[string]            `xml:"complemento" json:"complemento"`
	CodigoCliente                string                              `xml:"codigoCliente" json:"codigoCliente"`
	RazonSocialOperadorTurismo   datatype.Nilable[string]            `xml:"razonSocialOperadorTurismo" json:"razonSocialOperadorTurismo"`
	CantidadHuespedes            datatype.Nilable[int]               `xml:"cantidadHuespedes" json:"cantidadHuespedes"`
	CantidadHabitaciones         datatype.Nilable[int]               `xml:"cantidadHabitaciones" json:"cantidadHabitaciones"`
	CantidadMayores              datatype.Nilable[int]               `xml:"cantidadMayores" json:"cantidadMayores"`
	CantidadMenores              datatype.Nilable[int]               `xml:"cantidadMenores" json:"cantidadMenores"`
	FechaIngresoHospedaje        datatype.Nilable[datatype.TimeSiat] `xml:"fechaIngresoHospedaje" json:"fechaIngresoHospedaje"`
	CodigoMetodoPago             int                                 `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]             `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                             `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64                             `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                 int                                 `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                             `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                             `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	MontoGiftCard                datatype.Nilable[float64]           `xml:"montoGiftCard" json:"montoGiftCard"`
	DescuentoAdicional           datatype.Nilable[float64]           `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]               `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Cafc                         datatype.Nilable[string]            `xml:"cafc" json:"cafc"`
	Leyenda                      string                              `xml:"leyenda" json:"leyenda"`
	Usuario                      string                              `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                                 `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleTurismoHospedaje representa un ítem o producto dentro de la factura de Turismo y Hospedaje.
type DetalleTurismoHospedaje struct {
	ActividadEconomica   string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin    int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto       string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion          string                    `xml:"descripcion" json:"descripcion"`
	CodigoTipoHabitacion datatype.Nilable[int]     `xml:"codigoTipoHabitacion" json:"codigoTipoHabitacion"`
	Cantidad             float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida         int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario       float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento       datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal             float64                   `xml:"subTotal" json:"subTotal"`
	DetalleHuespedes     datatype.Nilable[string]  `xml:"detalleHuespedes" json:"detalleHuespedes"`
}
