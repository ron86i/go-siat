package invoices

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
)

// ComercialExportacionHidro representa la estructura completa de una factura de Comercial Exportación Hidrocarburos lista para ser procesada.
type ComercialExportacionHidro struct {
	models.RequestWrapper[documents.FacturaComercialExportacionHidro]
}

// ComercialExportacionHidroCabecera representa la sección de cabecera de una factura de Comercial Exportación Hidrocarburos.
type ComercialExportacionHidroCabecera struct {
	models.RequestWrapper[documents.CabeceraComercialExportacionHidro]
}

// ComercialExportacionHidroDetalle representa un ítem individual dentro del detalle de una factura de Comercial Exportación Hidrocarburos.
type ComercialExportacionHidroDetalle struct {
	models.RequestWrapper[documents.DetalleComercialExportacionHidro]
}

// NewComercialExportacionHidroBuilder inicia el proceso de construcción de una Factura de Comercial Exportación Hidrocarburos.
func NewComercialExportacionHidroBuilder() *comercialExportacionHidroBuilder {
	return &comercialExportacionHidroBuilder{
		factura: &documents.FacturaComercialExportacionHidro{
			XMLName:           xml.Name{Local: "facturaElectronicaComercialExportacionHidro"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercialExportacionHidro.xsd",
		},
	}
}

// NewComercialExportacionHidroCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewComercialExportacionHidroCabeceraBuilder() *comercialExportacionHidroCabeceraBuilder {
	return &comercialExportacionHidroCabeceraBuilder{
		cabecera: &documents.CabeceraComercialExportacionHidro{
			CodigoDocumentoSector: 43, // Sector 43 para Comercial Exportación Hidrocarburos
		},
	}
}

// NewComercialExportacionHidroDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewComercialExportacionHidroDetalleBuilder() *comercialExportacionHidroDetalleBuilder {
	return &comercialExportacionHidroDetalleBuilder{
		detalle: &documents.DetalleComercialExportacionHidro{},
	}
}

type comercialExportacionHidroBuilder struct {
	factura *documents.FacturaComercialExportacionHidro
}

func (b *comercialExportacionHidroBuilder) WithCabecera(req ComercialExportacionHidroCabecera) *comercialExportacionHidroBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraComercialExportacionHidro](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercialExportacionHidroBuilder) AddDetalle(req ComercialExportacionHidroDetalle) *comercialExportacionHidroBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleComercialExportacionHidro](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *comercialExportacionHidroBuilder) WithModalidad(tipo int) *comercialExportacionHidroBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercialExportacionHidro"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercialExportacionHidro.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercialExportacionHidro"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercialExportacionHidro.xsd"
	}
	return b
}

func (b *comercialExportacionHidroBuilder) Build() ComercialExportacionHidro {
	return ComercialExportacionHidro{models.NewRequestWrapper(b.factura)}
}

type comercialExportacionHidroCabeceraBuilder struct {
	cabecera *documents.CabeceraComercialExportacionHidro
}

func (b *comercialExportacionHidroCabeceraBuilder) WithNitEmisor(v int64) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithRazonSocialEmisor(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithMunicipio(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithTelefono(telefono *string) *comercialExportacionHidroCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithNumeroFactura(v int64) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCuf(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCufd(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoSucursal(v int) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithDireccion(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithNombreRazonSocial(v *string) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithNumeroDocumento(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithComplemento(v *string) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithDireccionComprador(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoCliente(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithIncoterm(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.Incoterm = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithIncotermDetalle(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.IncotermDetalle = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithPuertoDestino(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.PuertoDestino = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithLugarDestino(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.LugarDestino = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoPais(v int) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.CodigoPais = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoMetodoPago(v int) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithMontoTotal(v float64) *comercialExportacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCostosGastosNacionales(v any) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	var value string
	switch val := v.(type) {
	case string:
		value = val
	case *string:
		if val != nil {
			value = *val
		}
	default:
		jsonData, _ := json.Marshal(v)
		value = string(jsonData)
	}
	b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithTotalGastosNacionalesFob(v float64) *comercialExportacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TotalGastosNacionalesFob = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCostosGastosInternacionales(v any) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	var value string
	switch val := v.(type) {
	case string:
		value = val
	case *string:
		if val != nil {
			value = *val
		}
	default:
		jsonData, _ := json.Marshal(v)
		value = string(jsonData)
	}
	b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithTotalGastosInternacionales(v *float64) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithMontoDetalle(v float64) *comercialExportacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoDetalle = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoMoneda(v int) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithTipoCambio(v float64) *comercialExportacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercialExportacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithNumeroDescripcionPaquetesBultos(v any) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: nil}
		return b
	}
	var value string
	switch val := v.(type) {
	case string:
		value = val
	case *string:
		if val != nil {
			value = *val
		}
	default:
		jsonData, _ := json.Marshal(v)
		value = string(jsonData)
	}
	b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithInformacionAdicional(v *string) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoExcepcion(v *int) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithCafc(v *string) *comercialExportacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithLeyenda(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithUsuario(v string) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *comercialExportacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *comercialExportacionHidroCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercialExportacionHidroCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *comercialExportacionHidroCabeceraBuilder) Build() ComercialExportacionHidroCabecera {
	return ComercialExportacionHidroCabecera{models.NewRequestWrapper(b.cabecera)}
}

type comercialExportacionHidroDetalleBuilder struct {
	detalle *documents.DetalleComercialExportacionHidro
}

func (b *comercialExportacionHidroDetalleBuilder) WithActividadEconomica(v string) *comercialExportacionHidroDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithCodigoProductoSin(v int64) *comercialExportacionHidroDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithCodigoProducto(v string) *comercialExportacionHidroDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithCodigoNandina(v string) *comercialExportacionHidroDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithDescripcion(v string) *comercialExportacionHidroDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithCantidad(v float64) *comercialExportacionHidroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 6, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithUnidadMedida(v int) *comercialExportacionHidroDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithPrecioUnitario(v float64) *comercialExportacionHidroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 6, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithMontoDescuento(v *float64) *comercialExportacionHidroDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 6, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) WithSubTotal(v float64) *comercialExportacionHidroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 6, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercialExportacionHidroDetalleBuilder) Build() ComercialExportacionHidroDetalle {
	return ComercialExportacionHidroDetalle{models.NewRequestWrapper(b.detalle)}
}
