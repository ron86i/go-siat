package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
)

// FacturaServicioBasico representa la estructura completa de una factura de Servicio Básico para el SIAT.
type FacturaServicioBasico struct {
	XMLName           xml.Name                `json:"-"`
	XmlnsXsi          string                  `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                  `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraServicioBasico  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleServicioBasico `xml:"detalle" json:"detalle"`
}

// CabeceraServicioBasico contiene la información general del servicio básico.
type CabeceraServicioBasico struct {
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
	Mes                          datatype.Nilable[string]  `xml:"mes" json:"mes"`
	Gestion                      datatype.Nilable[int]     `xml:"gestion" json:"gestion"`
	Ciudad                       datatype.Nilable[string]  `xml:"ciudad" json:"ciudad"`
	Zona                         datatype.Nilable[string]  `xml:"zona" json:"zona"`
	NumeroMedidor                string                    `xml:"numeroMedidor" json:"numeroMedidor"`
	FechaEmision                 datatype.TimeSiat         `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            datatype.Nilable[string]  `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	DomicilioCliente             datatype.Nilable[string]  `xml:"domicilioCliente" json:"domicilioCliente"`
	CodigoTipoDocumentoIdentidad int                       `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string                    `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                  string                    `xml:"complemento" json:"complemento"`
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	SubTotalNeto                 float64                   `xml:"subTotalNeto" json:"subTotalNeto"`
	TotalPrecioNeto              float64                   `xml:"totalPrecioNeto" json:"totalPrecioNeto"`
	Iva                          float64                   `xml:"iva" json:"iva"`
	MontoAPagar                  float64                   `xml:"montoaPagar" json:"montoaPagar"`
	ConsumoPeriodo               datatype.Nilable[float64] `xml:"consumoPeriodo" json:"consumoPeriodo"`
	BeneficiarioLey1886          datatype.Nilable[int64]   `xml:"beneficiarioLey1886" json:"beneficiarioLey1886"`
	ComplementoBeneficiario      string                    `xml:"complementoBeneficiario" json:"complementoBeneficiario"`
	MontoDescuentoLey1886        datatype.Nilable[float64] `xml:"montoDescuentoLey1886" json:"montoDescuentoLey1886"`
	MontoDescuentoSocial         datatype.Nilable[float64] `xml:"montoDescuentoSocial" json:"montoDescuentoSocial"`
	TasaAseo                     datatype.Nilable[float64] `xml:"tasaAseo" json:"tasaAseo"`
	TasaAlumbrado                datatype.Nilable[float64] `xml:"tasaAlumbrado" json:"tasaAlumbrado"`
	AjusteNoSujetoIva            datatype.Nilable[float64] `xml:"ajusteNoSujetoIva" json:"ajusteNoSujetoIva"`
	DetalleAjusteNoSujetoIva     datatype.Nilable[string]  `xml:"detalleAjusteNoSujetoIva" json:"detalleAjusteNoSujetoIva"`
	OtrosPagosNoSujetoIva        datatype.Nilable[float64] `xml:"otrosPagosNoSujetoIva" json:"otrosPagosNoSujetoIva"`
	DetalleOtrosPagosNoSujetoIva datatype.Nilable[string]  `xml:"detalleOtrosPagosNoSujetoIva" json:"detalleOtrosPagosNoSujetoIva"`
	OtrasTasas                   datatype.Nilable[float64] `xml:"otrasTasas" json:"otrasTasas"`
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

// DetalleServicioBasico representa un ítem o producto dentro de la factura de Servicio Básico.
type DetalleServicioBasico struct {
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
