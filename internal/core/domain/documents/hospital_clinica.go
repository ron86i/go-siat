package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaHospitalClinica representa la estructura completa de una factura de Hospital o Clínica para el SIAT.
type FacturaHospitalClinica struct {
	XMLName           xml.Name                 `json:"-"`
	XmlnsXsi          string                   `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string                   `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraHospitalClinica  `xml:"cabecera" json:"cabecera"`
	Detalle           []DetalleHospitalClinica `xml:"detalle" json:"detalle"`
}

// CabeceraHospitalClinica contiene la información general y del cliente de la factura.
type CabeceraHospitalClinica struct {
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
	Complemento                  datatype.Nilable[string]  `xml:"complemento" json:"complemento"`
	CodigoCliente                string                    `xml:"codigoCliente" json:"codigoCliente"`
	ModalidadServicio            datatype.Nilable[string]  `xml:"modalidadServicio" json:"modalidadServicio"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64                   `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64                   `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                 int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	MontoGiftCard                datatype.Nilable[float64] `xml:"montoGiftCard" json:"montoGiftCard"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Cafc                         datatype.Nilable[string]  `xml:"cafc" json:"cafc"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// DetalleHospitalClinica representa un ítem de servicio médico.
type DetalleHospitalClinica struct {
	ActividadEconomica          string                    `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin           int64                     `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto              string                    `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion                 string                    `xml:"descripcion" json:"descripcion"`
	Especialidad                datatype.Nilable[string]  `xml:"especialidad" json:"especialidad"`
	EspecialidadDetalle         datatype.Nilable[string]  `xml:"especialidadDetalle" json:"especialidadDetalle"`
	NroQuirofanoSalaOperaciones int                       `xml:"nroQuirofanoSalaOperaciones" json:"nroQuirofanoSalaOperaciones"`
	EspecialidadMedico          datatype.Nilable[string]  `xml:"especialidadMedico" json:"especialidadMedico"`
	NombreApellidoMedico        string                    `xml:"nombreApellidoMedico" json:"nombreApellidoMedico"`
	NitDocumentoMedico          int64                     `xml:"nitDocumentoMedico" json:"nitDocumentoMedico"`
	NroMatriculaMedico          datatype.Nilable[string]  `xml:"nroMatriculaMedico" json:"nroMatriculaMedico"`
	NroFacturaMedico            datatype.Nilable[int]     `xml:"nroFacturaMedico" json:"nroFacturaMedico"`
	Cantidad                    float64                   `xml:"cantidad" json:"cantidad"`
	UnidadMedida                int                       `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario              float64                   `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento              datatype.Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal                    float64                   `xml:"subTotal" json:"subTotal"`
}
