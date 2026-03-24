package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// Seguros representa la estructura completa de una factura de Seguros lista para ser procesada.
type Seguros struct {
	models.RequestWrapper[documentos.FacturaSeguros]
}

// SegurosCabecera representa la sección de cabecera de una factura de Seguros.
type SegurosCabecera struct {
	models.RequestWrapper[documentos.CabeceraSeguros]
}

// SegurosDetalle representa un ítem individual dentro del detalle de una factura de Seguros.
type SegurosDetalle struct {
	models.RequestWrapper[documentos.DetalleSeguros]
}

// NewSegurosBuilder inicia el proceso de construcción de una Factura de Seguros.
func NewSegurosBuilder() *segurosBuilder {
	return &segurosBuilder{
		factura: &documentos.FacturaSeguros{
			XMLName:           xml.Name{Local: "facturaElectronicaSeguros"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaSeguros.xsd",
		},
	}
}

// NewSegurosCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewSegurosCabeceraBuilder() *segurosCabeceraBuilder {
	return &segurosCabeceraBuilder{
		cabecera: &documentos.CabeceraSeguros{
			CodigoDocumentoSector: 34, // Sector 34 para Seguros
		},
	}
}

// NewSegurosDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewSegurosDetalleBuilder() *segurosDetalleBuilder {
	return &segurosDetalleBuilder{
		detalle: &documentos.DetalleSeguros{},
	}
}

type segurosBuilder struct {
	factura *documentos.FacturaSeguros
}

func (b *segurosBuilder) WithCabecera(req SegurosCabecera) *segurosBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraSeguros](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *segurosBuilder) AddDetalle(req SegurosDetalle) *segurosBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleSeguros](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *segurosBuilder) WithModalidad(tipo int) *segurosBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaSeguros"}
		b.factura.XsiSchemaLocation = "facturaElectronicaSeguros.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaSeguros"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaSeguros.xsd"
	}
	return b
}

func (b *segurosBuilder) Build() Seguros {
	return Seguros{models.NewRequestWrapper(b.factura)}
}

type segurosCabeceraBuilder struct {
	cabecera *documentos.CabeceraSeguros
}

func (b *segurosCabeceraBuilder) WithNitEmisor(v int64) *segurosCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *segurosCabeceraBuilder) WithRazonSocialEmisor(v string) *segurosCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *segurosCabeceraBuilder) WithMunicipio(v string) *segurosCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *segurosCabeceraBuilder) WithTelefono(telefono *string) *segurosCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithNumeroFactura(v int64) *segurosCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *segurosCabeceraBuilder) WithCuf(v string) *segurosCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *segurosCabeceraBuilder) WithCufd(v string) *segurosCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *segurosCabeceraBuilder) WithCodigoSucursal(v int) *segurosCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *segurosCabeceraBuilder) WithDireccion(v string) *segurosCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *segurosCabeceraBuilder) WithCodigoPuntoVenta(v *int) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *segurosCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *segurosCabeceraBuilder) WithNombreRazonSocial(v *string) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *segurosCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *segurosCabeceraBuilder) WithNumeroDocumento(v string) *segurosCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *segurosCabeceraBuilder) WithComplemento(v *string) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithCodigoCliente(v string) *segurosCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *segurosCabeceraBuilder) WithCodigoMetodoPago(v int) *segurosCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *segurosCabeceraBuilder) WithNumeroTarjeta(v *int64) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithMontoTotal(v float64) *segurosCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *segurosCabeceraBuilder) WithAjusteAfectacionIva(v float64) *segurosCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.AjusteAfectacionIva = v
	return b
}

func (b *segurosCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *segurosCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *segurosCabeceraBuilder) WithCodigoMoneda(v int) *segurosCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *segurosCabeceraBuilder) WithTipoCambio(v float64) *segurosCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *segurosCabeceraBuilder) WithMontoTotalMoneda(v float64) *segurosCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *segurosCabeceraBuilder) WithMontoGiftCard(v *float64) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithDescuentoAdicional(v *float64) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithCodigoExcepcion(v *int) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithCafc(v *string) *segurosCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *segurosCabeceraBuilder) WithLeyenda(v string) *segurosCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *segurosCabeceraBuilder) WithUsuario(v string) *segurosCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *segurosCabeceraBuilder) WithCodigoDocumentoSector(v int) *segurosCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *segurosCabeceraBuilder) Build() SegurosCabecera {
	return SegurosCabecera{models.NewRequestWrapper(b.cabecera)}
}

type segurosDetalleBuilder struct {
	detalle *documentos.DetalleSeguros
}

func (b *segurosDetalleBuilder) WithActividadEconomica(v string) *segurosDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *segurosDetalleBuilder) WithCodigoProductoSin(v int64) *segurosDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *segurosDetalleBuilder) WithCodigoProducto(v string) *segurosDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *segurosDetalleBuilder) WithDescripcion(v string) *segurosDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *segurosDetalleBuilder) WithCantidad(v float64) *segurosDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *segurosDetalleBuilder) WithUnidadMedida(v int) *segurosDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *segurosDetalleBuilder) WithPrecioUnitario(v float64) *segurosDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *segurosDetalleBuilder) WithMontoDescuento(v *float64) *segurosDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *segurosDetalleBuilder) WithSubTotal(v float64) *segurosDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *segurosDetalleBuilder) Build() SegurosDetalle {
	return SegurosDetalle{models.NewRequestWrapper(b.detalle)}
}
