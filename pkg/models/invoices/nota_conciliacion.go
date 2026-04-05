package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
)

// NotaConciliacion representa la estructura completa de una nota de conciliación lista para ser procesada.
type NotaConciliacion struct {
	models.RequestWrapper[documents.NotaConciliacion]
}

// NotaConciliacionCabecera representa la sección de cabecera de una nota de conciliación.
type NotaConciliacionCabecera struct {
	models.RequestWrapper[documents.CabeceraConciliacion]
}

// NotaDetalleOriginal representa un ítem individual dentro del detalle original.
type NotaDetalleOriginal struct {
	models.RequestWrapper[documents.DetalleOriginal]
}

// NotaDetalleConciliacion representa un ítem individual dentro del detalle de conciliación.
type NotaDetalleConciliacion struct {
	models.RequestWrapper[documents.DetalleConciliacion]
}

// NewNotaConciliacionBuilder inicia el proceso de construcción de una Nota de Conciliación.
func NewNotaConciliacionBuilder() *notaConciliacionBuilder {
	return &notaConciliacionBuilder{
		nota: &documents.NotaConciliacion{
			XMLName:           xml.Name{Local: "notaElectronicaConciliacion"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "notaElectronicaConciliacion.xsd",
		},
	}
}

// NewNotaConciliacionCabeceraBuilder crea una nueva instancia del constructor para la cabecera.
func NewNotaConciliacionCabeceraBuilder() *notaConciliacionCabeceraBuilder {
	return &notaConciliacionCabeceraBuilder{
		cabecera: &documents.CabeceraConciliacion{
			CodigoDocumentoSector: 29, // Sector 29 para Nota de Conciliación
		},
	}
}

// NewNotaDetalleOriginalBuilder crea una nueva instancia del constructor para los ítems de detalle original.
func NewNotaDetalleOriginalBuilder() *notaDetalleOriginalBuilder {
	return &notaDetalleOriginalBuilder{
		detalle: &documents.DetalleOriginal{},
	}
}

// NewNotaDetalleConciliacionBuilder crea una nueva instancia del constructor para los ítems de detalle de conciliación.
func NewNotaDetalleConciliacionBuilder() *notaDetalleConciliacionBuilder {
	return &notaDetalleConciliacionBuilder{
		detalle: &documents.DetalleConciliacion{},
	}
}

type notaConciliacionBuilder struct {
	nota *documents.NotaConciliacion
}

func (b *notaConciliacionBuilder) WithCabecera(req NotaConciliacionCabecera) *notaConciliacionBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraConciliacion](req); internal != nil {
		b.nota.Cabecera = *internal
	}
	return b
}

func (b *notaConciliacionBuilder) AddDetalleOriginal(req NotaDetalleOriginal) *notaConciliacionBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleOriginal](req); internal != nil {
		b.nota.DetalleOriginal = append(b.nota.DetalleOriginal, *internal)
	}
	return b
}

func (b *notaConciliacionBuilder) AddDetalleConciliacion(req NotaDetalleConciliacion) *notaConciliacionBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleConciliacion](req); internal != nil {
		b.nota.DetalleConciliacion = append(b.nota.DetalleConciliacion, *internal)
	}
	return b
}

func (b *notaConciliacionBuilder) WithModalidad(tipo int) *notaConciliacionBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.nota.XMLName = xml.Name{Local: "notaElectronicaConciliacion"}
		b.nota.XsiSchemaLocation = "notaElectronicaConciliacion.xsd"
	case siat.ModalidadComputarizada:
		// Se asume el mismo patrón que otros documentos sectoriales
		b.nota.XMLName = xml.Name{Local: "notaComputarizadaConciliacion"}
		b.nota.XsiSchemaLocation = "notaComputarizadaConciliacion.xsd"
	}
	return b
}

func (b *notaConciliacionBuilder) Build() NotaConciliacion {
	return NotaConciliacion{models.NewRequestWrapper(b.nota)}
}

type notaConciliacionCabeceraBuilder struct {
	cabecera *documents.CabeceraConciliacion
}

func (b *notaConciliacionCabeceraBuilder) WithNitEmisor(v int64) *notaConciliacionCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithRazonSocialEmisor(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithMunicipio(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithTelefono(v *string) *notaConciliacionCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithNumeroNotaConciliacion(v int64) *notaConciliacionCabeceraBuilder {
	b.cabecera.NumeroNotaConciliacion = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCuf(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCufd(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCodigoSucursal(v int) *notaConciliacionCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithDireccion(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCodigoPuntoVenta(v *int) *notaConciliacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithFechaEmision(v time.Time) *notaConciliacionCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithNombreRazonSocial(v *string) *notaConciliacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *notaConciliacionCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithNumeroDocumento(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithComplemento(v *string) *notaConciliacionCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCodigoCliente(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithNumeroFactura(v int64) *notaConciliacionCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithNumeroAutorizacionCuf(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.NumeroAutorizacionCuf = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCodigoControl(v *string) *notaConciliacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoControl = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoControl = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithFechaEmisionFactura(v time.Time) *notaConciliacionCabeceraBuilder {
	b.cabecera.FechaEmisionFactura = datatype.NewTimeSiat(v)
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithMontoTotalOriginal(v float64) *notaConciliacionCabeceraBuilder {
	b.cabecera.MontoTotalOriginal = datatype.Float64Round(v, 2)
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithMontoTotalConciliado(v float64) *notaConciliacionCabeceraBuilder {
	b.cabecera.MontoTotalConciliado = datatype.Float64Round(v, 2)
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCreditoFiscalIva(v float64) *notaConciliacionCabeceraBuilder {
	b.cabecera.CreditoFiscalIva = datatype.Float64Round(v, 2)
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithDebitoFiscalIva(v float64) *notaConciliacionCabeceraBuilder {
	b.cabecera.DebitoFiscalIva = datatype.Float64Round(v, 2)
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithCodigoExcepcion(v *int) *notaConciliacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithLeyenda(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) WithUsuario(v string) *notaConciliacionCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *notaConciliacionCabeceraBuilder) Build() NotaConciliacionCabecera {
	return NotaConciliacionCabecera{models.NewRequestWrapper(b.cabecera)}
}

type notaDetalleOriginalBuilder struct {
	detalle *documents.DetalleOriginal
}

func (b *notaDetalleOriginalBuilder) WithActividadEconomica(v string) *notaDetalleOriginalBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *notaDetalleOriginalBuilder) WithCodigoProductoSin(v int64) *notaDetalleOriginalBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *notaDetalleOriginalBuilder) WithCodigoProducto(v string) *notaDetalleOriginalBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *notaDetalleOriginalBuilder) WithDescripcion(v string) *notaDetalleOriginalBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *notaDetalleOriginalBuilder) WithCantidad(v float64) *notaDetalleOriginalBuilder {
	b.detalle.Cantidad = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleOriginalBuilder) WithUnidadMedida(v int) *notaDetalleOriginalBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *notaDetalleOriginalBuilder) WithPrecioUnitario(v float64) *notaDetalleOriginalBuilder {
	b.detalle.PrecioUnitario = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleOriginalBuilder) WithMontoDescuento(v *float64) *notaDetalleOriginalBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 10)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleOriginalBuilder) WithSubTotal(v float64) *notaDetalleOriginalBuilder {
	b.detalle.SubTotal = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleOriginalBuilder) Build() NotaDetalleOriginal {
	return NotaDetalleOriginal{models.NewRequestWrapper(b.detalle)}
}

type notaDetalleConciliacionBuilder struct {
	detalle *documents.DetalleConciliacion
}

func (b *notaDetalleConciliacionBuilder) WithActividadEconomica(v string) *notaDetalleConciliacionBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *notaDetalleConciliacionBuilder) WithCodigoProductoSin(v int64) *notaDetalleConciliacionBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *notaDetalleConciliacionBuilder) WithCodigoProducto(v string) *notaDetalleConciliacionBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *notaDetalleConciliacionBuilder) WithDescripcion(v string) *notaDetalleConciliacionBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *notaDetalleConciliacionBuilder) WithMontoOriginal(v float64) *notaDetalleConciliacionBuilder {
	b.detalle.MontoOriginal = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleConciliacionBuilder) WithMontoFinal(v float64) *notaDetalleConciliacionBuilder {
	b.detalle.MontoFinal = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleConciliacionBuilder) WithMontoConciliado(v float64) *notaDetalleConciliacionBuilder {
	b.detalle.MontoConciliado = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleConciliacionBuilder) Build() NotaDetalleConciliacion {
	return NotaDetalleConciliacion{models.NewRequestWrapper(b.detalle)}
}
