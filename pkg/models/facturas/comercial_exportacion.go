package facturas

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// ComercialExportacion representa la estructura completa de una factura comercial
// de exportación lista para ser procesada.
type ComercialExportacion struct {
	requestWrapper[documentos.FacturaComercialExportacion]
}

// ComercialExportacionCabecera representa la sección de cabecera de la factura.
type ComercialExportacionCabecera struct {
	requestWrapper[documentos.CabeceraComercialExportacion]
}

// ComercialExportacionDetalle representa un ítem individual dentro del detalle.
type ComercialExportacionDetalle struct {
	requestWrapper[documentos.DetalleComercialExportacion]
}

// NewComercialExportacionBuilder inicia el proceso de construcción de la factura.
func NewComercialExportacionBuilder() *comercialExportacionBuilder {
	return &comercialExportacionBuilder{
		factura: &documentos.FacturaComercialExportacion{
			XMLName:           xml.Name{Local: "facturaElectronicaComercialExportacion"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercialExportacion.xsd",
		},
	}
}

// NewComercialExportacionCabeceraBuilder crea el constructor para la cabecera.
func NewComercialExportacionCabeceraBuilder() *comercialExportacionCabeceraBuilder {
	return &comercialExportacionCabeceraBuilder{
		cabecera: &documentos.CabeceraComercialExportacion{
			MontoTotalSujetoIva:   0, // Fixed in XSD
			CodigoDocumentoSector: 3, // Fixed in XSD
		},
	}
}

// NewComercialExportacionDetalleBuilder crea el constructor para los detalles.
func NewComercialExportacionDetalleBuilder() *comercialExportacionDetalleBuilder {
	return &comercialExportacionDetalleBuilder{
		detalle: &documentos.DetalleComercialExportacion{},
	}
}

type comercialExportacionBuilder struct {
	factura *documentos.FacturaComercialExportacion
}

func (b *comercialExportacionBuilder) WithCabecera(req ComercialExportacionCabecera) *comercialExportacionBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *comercialExportacionBuilder) AddDetalle(req ComercialExportacionDetalle) *comercialExportacionBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

func (b *comercialExportacionBuilder) WithModalidad(tipo int) *comercialExportacionBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercialExportacion"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercialExportacion.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercialExportacion"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercialExportacion.xsd"
	}
	return b
}

func (b *comercialExportacionBuilder) Build() ComercialExportacion {
	return ComercialExportacion{requestWrapper[documentos.FacturaComercialExportacion]{request: b.factura}}
}

type comercialExportacionCabeceraBuilder struct {
	cabecera *documentos.CabeceraComercialExportacion
}

func (b *comercialExportacionCabeceraBuilder) WithNitEmisor(v int64) *comercialExportacionCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithRazonSocialEmisor(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithMunicipio(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTelefono(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroFactura(v int64) *comercialExportacionCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCuf(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCufd(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoSucursal(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithDireccion(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithFechaEmision(v time.Time) *comercialExportacionCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNombreRazonSocial(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroDocumento(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithComplemento(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithDireccionComprador(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoCliente(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithIncoterm(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Incoterm = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithIncotermDetalle(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.IncotermDetalle = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithPuertoDestino(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.PuertoDestino = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithLugarDestino(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.LugarDestino = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoPais(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoPais = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoMetodoPago(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithMontoTotal(v float64) *comercialExportacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCostosGastosNacionales(v map[string]interface{}) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	jsonData, _ := json.Marshal(v)
	val := string(jsonData)
	b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTotalGastosNacionalesFob(v float64) *comercialExportacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TotalGastosNacionalesFob = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCostosGastosInternacionales(v map[string]interface{}) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	jsonData, _ := json.Marshal(v)
	val := string(jsonData)
	b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTotalGastosInternacionales(v *float64) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithMontoDetalle(v float64) *comercialExportacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoDetalle = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoMoneda(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTipoCambio(v float64) *comercialExportacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercialExportacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroDescripcionPaquetesBultos(v map[string]interface{}) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: nil}
		return b
	}
	jsonData, _ := json.Marshal(v)
	val := string(jsonData)
	b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithInformacionAdicional(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoExcepcion(v *int) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCafc(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithLeyenda(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithUsuario(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) Build() ComercialExportacionCabecera {
	return ComercialExportacionCabecera{requestWrapper[documentos.CabeceraComercialExportacion]{request: b.cabecera}}
}

type comercialExportacionDetalleBuilder struct {
	detalle *documentos.DetalleComercialExportacion
}

func (b *comercialExportacionDetalleBuilder) WithActividadEconomica(v string) *comercialExportacionDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCodigoProductoSin(v int64) *comercialExportacionDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCodigoProducto(v string) *comercialExportacionDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCodigoNandina(v string) *comercialExportacionDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithDescripcion(v string) *comercialExportacionDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCantidad(v float64) *comercialExportacionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithUnidadMedida(v int) *comercialExportacionDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithPrecioUnitario(v float64) *comercialExportacionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithMontoDescuento(v *float64) *comercialExportacionDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *comercialExportacionDetalleBuilder) WithSubTotal(v float64) *comercialExportacionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercialExportacionDetalleBuilder) Build() ComercialExportacionDetalle {
	return ComercialExportacionDetalle{requestWrapper[documentos.DetalleComercialExportacion]{request: b.detalle}}
}
