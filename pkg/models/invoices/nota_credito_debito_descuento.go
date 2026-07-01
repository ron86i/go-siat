package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v2/internal/core/domain/documents"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

type NotaCreditoDebitoDescuento struct {
	models.RequestWrapper[documents.NotaCreditoDebitoDescuento]
}

type NotaCreditoDebitoDescuentoCabecera struct {
	models.RequestWrapper[documents.CabeceraNotaCreditoDebitoDescuento]
}

type NotaDetalleCreditoDebitoDescuento struct {
	models.RequestWrapper[documents.DetalleNotaCreditoDebitoDescuento]
}

func NewNotaCreditoDebitoDescuentoBuilder() *notaCreditoDebitoDescuentoBuilder {
	return &notaCreditoDebitoDescuentoBuilder{
		nota: &documents.NotaCreditoDebitoDescuento{
			XMLName:           xml.Name{Local: "notaElectronicaCreditoDebitoDescuento"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "notaElectronicaCreditoDebitoDescuento.xsd",
		},
	}
}

func NewNotaCreditoDebitoDescuentoCabeceraBuilder() *notaCreditoDebitoDescuentoCabeceraBuilder {
	return &notaCreditoDebitoDescuentoCabeceraBuilder{
		cabecera: &documents.CabeceraNotaCreditoDebitoDescuento{
			CodigoDocumentoSector: 47,
		},
	}
}

func NewNotaDetalleCreditoDebitoDescuentoBuilder() *notaDetalleCreditoDebitoDescuentoBuilder {
	return &notaDetalleCreditoDebitoDescuentoBuilder{
		detalle: &documents.DetalleNotaCreditoDebitoDescuento{},
	}
}

type notaCreditoDebitoDescuentoBuilder struct {
	nota *documents.NotaCreditoDebitoDescuento
}

func (b *notaCreditoDebitoDescuentoBuilder) WithCabecera(req NotaCreditoDebitoDescuentoCabecera) *notaCreditoDebitoDescuentoBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraNotaCreditoDebitoDescuento](req); internal != nil {
		b.nota.Cabecera = *internal
	}
	return b
}

func (b *notaCreditoDebitoDescuentoBuilder) AddDetalle(req NotaDetalleCreditoDebitoDescuento) *notaCreditoDebitoDescuentoBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleNotaCreditoDebitoDescuento](req); internal != nil {
		b.nota.Detalle = append(b.nota.Detalle, *internal)
	}
	return b
}

func (b *notaCreditoDebitoDescuentoBuilder) WithModalidad(tipo int) *notaCreditoDebitoDescuentoBuilder {
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

func (b *notaCreditoDebitoDescuentoBuilder) Build() NotaCreditoDebitoDescuento {
	return NotaCreditoDebitoDescuento{models.NewRequestWrapper(b.nota)}
}

type notaCreditoDebitoDescuentoCabeceraBuilder struct {
	cabecera *documents.CabeceraNotaCreditoDebitoDescuento
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithNitEmisor(v int64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithRazonSocialEmisor(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithMunicipio(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithTelefono(v *string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithNumeroNotaCreditoDebito(v int64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.NumeroNotaCreditoDebito = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithCuf(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithCufd(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithCodigoSucursal(v int) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithDireccion(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithCodigoPuntoVenta(v *int) *notaCreditoDebitoDescuentoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithFechaEmision(v time.Time) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithNombreRazonSocial(v *string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithNumeroDocumento(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithComplemento(v *string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithCodigoCliente(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithNumeroFactura(v int64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithNumeroAutorizacionCuf(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.NumeroAutorizacionCuf = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithFechaEmisionFactura(v time.Time) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.FechaEmisionFactura = datatype.NewTimeSiat(v)
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithMontoTotalOriginal(v float64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.MontoTotalOriginal = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithDescuentoAdicional(v *float64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 2)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithMontoTotalDevuelto(v float64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.MontoTotalDevuelto = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithMontoDescuentoCreditoDebito(v *float64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 2)
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithMontoEfectivoCreditoDebito(v float64) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.MontoEfectivoCreditoDebito = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithCodigoExcepcion(v *int) *notaCreditoDebitoDescuentoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithLeyenda(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) WithUsuario(v string) *notaCreditoDebitoDescuentoCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *notaCreditoDebitoDescuentoCabeceraBuilder) Build() NotaCreditoDebitoDescuentoCabecera {
	return NotaCreditoDebitoDescuentoCabecera{models.NewRequestWrapper(b.cabecera)}
}

type notaDetalleCreditoDebitoDescuentoBuilder struct {
	detalle *documents.DetalleNotaCreditoDebitoDescuento
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithNroItem(v int) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.NroItem = v
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithActividadEconomica(v string) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithCodigoProductoSin(v int64) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithCodigoProducto(v string) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithDescripcion(v string) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithCantidad(v float64) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.Cantidad = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithUnidadMedida(v int) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithPrecioUnitario(v float64) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.PrecioUnitario = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithMontoDescuento(v *float64) *notaDetalleCreditoDebitoDescuentoBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 10)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithSubTotal(v float64) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.SubTotal = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) WithCodigoDetalleTransaccion(v int) *notaDetalleCreditoDebitoDescuentoBuilder {
	b.detalle.CodigoDetalleTransaccion = v
	return b
}

func (b *notaDetalleCreditoDebitoDescuentoBuilder) Build() NotaDetalleCreditoDebitoDescuento {
	return NotaDetalleCreditoDebitoDescuento{models.NewRequestWrapper(b.detalle)}
}
