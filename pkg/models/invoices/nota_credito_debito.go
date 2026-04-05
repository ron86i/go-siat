package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
)

// NotaCreditoDebito representa la estructura completa de una nota de crédito/débito/descuento lista para ser procesada.
type NotaCreditoDebito struct {
	models.RequestWrapper[documents.NotaCreditoDebito]
}

// NotaCreditoDebitoCabecera representa la sección de cabecera de la nota.
type NotaCreditoDebitoCabecera struct {
	models.RequestWrapper[documents.CabeceraNotaCreditoDebito]
}

// NotaDetalleCreditoDebito representa un ítem individual dentro del detalle.
type NotaDetalleCreditoDebito struct {
	models.RequestWrapper[documents.DetalleNotaCreditoDebito]
}

// NewNotaCreditoDebitoBuilder inicia el proceso de construcción de una Nota.
func NewNotaCreditoDebitoBuilder() *notaCreditoDebitoBuilder {
	return &notaCreditoDebitoBuilder{
		nota: &documents.NotaCreditoDebito{
			XMLName:           xml.Name{Local: "notaElectronicaCreditoDebitoDescuento"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "notaElectronicaCreditoDebitoDescuento.xsd",
		},
	}
}

// NewNotaCreditoDebitoCabeceraBuilder crea una nueva instancia del constructor para la cabecera.
func NewNotaCreditoDebitoCabeceraBuilder() *notaCreditoDebitoCabeceraBuilder {
	return &notaCreditoDebitoCabeceraBuilder{
		cabecera: &documents.CabeceraNotaCreditoDebito{
			CodigoDocumentoSector: 47, // Sector 47
		},
	}
}

// NewNotaDetalleCreditoDebitoBuilder crea una nueva instancia del constructor para el detalle.
func NewNotaDetalleCreditoDebitoBuilder() *notaDetalleCreditoDebitoBuilder {
	return &notaDetalleCreditoDebitoBuilder{
		detalle: &documents.DetalleNotaCreditoDebito{},
	}
}

type notaCreditoDebitoBuilder struct {
	nota *documents.NotaCreditoDebito
}

func (b *notaCreditoDebitoBuilder) WithCabecera(req NotaCreditoDebitoCabecera) *notaCreditoDebitoBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraNotaCreditoDebito](req); internal != nil {
		b.nota.Cabecera = *internal
	}
	return b
}

func (b *notaCreditoDebitoBuilder) AddDetalle(req NotaDetalleCreditoDebito) *notaCreditoDebitoBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleNotaCreditoDebito](req); internal != nil {
		b.nota.Detalle = append(b.nota.Detalle, *internal)
	}
	return b
}

func (b *notaCreditoDebitoBuilder) WithModalidad(tipo int) *notaCreditoDebitoBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.nota.XMLName = xml.Name{Local: "notaElectronicaCreditoDebitoDescuento"}
		b.nota.XsiSchemaLocation = "notaElectronicaCreditoDebitoDescuento.xsd"
	case siat.ModalidadComputarizada:
		b.nota.XMLName = xml.Name{Local: "notaComputarizadaCreditoDebitoDescuento"}
		b.nota.XsiSchemaLocation = "notaComputarizadaCreditoDebitoDescuento.xsd"
	}
	return b
}

func (b *notaCreditoDebitoBuilder) Build() NotaCreditoDebito {
	return NotaCreditoDebito{models.NewRequestWrapper(b.nota)}
}

type notaCreditoDebitoCabeceraBuilder struct {
	cabecera *documents.CabeceraNotaCreditoDebito
}

func (b *notaCreditoDebitoCabeceraBuilder) WithNitEmisor(v int64) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithRazonSocialEmisor(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithMunicipio(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithTelefono(v *string) *notaCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithNumeroNotaCreditoDebito(v int64) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroNotaCreditoDebito = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithCuf(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithCufd(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithCodigoSucursal(v int) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithDireccion(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithCodigoPuntoVenta(v *int) *notaCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithFechaEmision(v time.Time) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithNombreRazonSocial(v *string) *notaCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithNumeroDocumento(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithComplemento(v *string) *notaCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithCodigoCliente(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithNumeroFactura(v int64) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithNumeroAutorizacionCuf(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroAutorizacionCuf = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithFechaEmisionFactura(v time.Time) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.FechaEmisionFactura = datatype.NewTimeSiat(v)
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithMontoTotalOriginal(v float64) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.MontoTotalOriginal = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithDescuentoAdicional(v *float64) *notaCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 2)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithMontoTotalDevuelto(v float64) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.MontoTotalDevuelto = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithMontoDescuentoCreditoDebito(v *float64) *notaCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 2)
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithMontoEfectivoCreditoDebito(v float64) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.MontoEfectivoCreditoDebito = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithCodigoExcepcion(v *int) *notaCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithLeyenda(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) WithUsuario(v string) *notaCreditoDebitoCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *notaCreditoDebitoCabeceraBuilder) Build() NotaCreditoDebitoCabecera {
	return NotaCreditoDebitoCabecera{models.NewRequestWrapper(b.cabecera)}
}

type notaDetalleCreditoDebitoBuilder struct {
	detalle *documents.DetalleNotaCreditoDebito
}

func (b *notaDetalleCreditoDebitoBuilder) WithNroItem(v int) *notaDetalleCreditoDebitoBuilder {
	b.detalle.NroItem = v
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithActividadEconomica(v string) *notaDetalleCreditoDebitoBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithCodigoProductoSin(v int64) *notaDetalleCreditoDebitoBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithCodigoProducto(v string) *notaDetalleCreditoDebitoBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithDescripcion(v string) *notaDetalleCreditoDebitoBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithCantidad(v float64) *notaDetalleCreditoDebitoBuilder {
	b.detalle.Cantidad = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithUnidadMedida(v int) *notaDetalleCreditoDebitoBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithPrecioUnitario(v float64) *notaDetalleCreditoDebitoBuilder {
	b.detalle.PrecioUnitario = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithMontoDescuento(v *float64) *notaDetalleCreditoDebitoBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 10)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithSubTotal(v float64) *notaDetalleCreditoDebitoBuilder {
	b.detalle.SubTotal = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) WithCodigoDetalleTransaccion(v int) *notaDetalleCreditoDebitoBuilder {
	b.detalle.CodigoDetalleTransaccion = v
	return b
}

func (b *notaDetalleCreditoDebitoBuilder) Build() NotaDetalleCreditoDebito {
	return NotaDetalleCreditoDebito{models.NewRequestWrapper(b.detalle)}
}
