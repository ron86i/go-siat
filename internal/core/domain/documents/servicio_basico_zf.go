package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaServicioBasicoZF representa la estructura raíz para documentos del sector 40 (Servicio Básico Zona Franca).
type FacturaServicioBasicoZF struct {
	XMLName           xml.Name                `json:"-"`
	XmlnsXsi          string                  `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                  `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraServicioBasicoZF  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleServicioBasico `xml:"detalle" json:"detalle"`
}

// CabeceraServicioBasicoZF define los campos de cabecera específicos para el sector 40.
type CabeceraServicioBasicoZF struct {
	NitEmisor                      int64                     `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor              string                    `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                      string                    `xml:"municipio" json:"municipio"`
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
	Complemento                  datatype.Nilable[string]  `xml:"complemento" json:"complemento"`
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                   `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64                   `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	ConsumoPeriodo               datatype.Nilable[float64] `xml:"consumoPeriodo" json:"consumoPeriodo"`
	BeneficiarioLey1886          datatype.Nilable[int64]   `xml:"beneficiarioLey1886" json:"beneficiarioLey1886"`
	MontoDescuentoLey1886        datatype.Nilable[float64] `xml:"montoDescuentoLey1886" json:"montoDescuentoLey1886"`
	MontoDescuentoTarifaDignidad datatype.Nilable[float64] `xml:"montoDescuentoTarifaDignidad" json:"montoDescuentoTarifaDignidad"`
	TasaAseo                     datatype.Nilable[float64] `xml:"tasaAseo" json:"tasaAseo"`
	TasaAlumbrado                datatype.Nilable[float64] `xml:"tasaAlumbrado" json:"tasaAlumbrado"`
	AjusteNoSujetoIva            datatype.Nilable[float64] `xml:"ajusteNoSujetoIva" json:"ajusteNoSujetoIva"`
	DetalleAjusteNoSujetoIva     datatype.Nilable[string]  `xml:"detalleAjusteNoSujetoIva" json:"detalleAjusteNoSujetoIva"`
	AjusteSujetoIva              datatype.Nilable[float64] `xml:"ajusteSujetoIva" json:"ajusteSujetoIva"`
	DetalleAjusteSujetoIva       datatype.Nilable[string]  `xml:"detalleAjusteSujetoIva" json:"detalleAjusteSujetoIva"`
	OtrosPagosNoSujetoIva        datatype.Nilable[float64] `xml:"otrosPagosNoSujetoIva" json:"otrosPagosNoSujetoIva"`
	DetalleOtrosPagosNoSujetoIva datatype.Nilable[string]  `xml:"detalleOtrosPagosNoSujetoIva" json:"detalleOtrosPagosNoSujetoIva"`
	OtrasTasas                   datatype.Nilable[float64] `xml:"otrasTasas" json:"otrasTasas"`
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
