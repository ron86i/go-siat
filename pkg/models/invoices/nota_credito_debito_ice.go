package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
)

// NotaCreditoDebitoIce representa la estructura completa de una nota de crédito/débito/descuento ICE lista para procesarse.
type NotaCreditoDebitoIce struct {
	models.RequestWrapper[documents.NotaCreditoDebitoIce]
}

// NotaCreditoDebitoIceCabecera representa la sección de cabecera de la nota ICE.
type NotaCreditoDebitoIceCabecera struct {
	models.RequestWrapper[documents.CabeceraNotaCreditoDebitoIce]
}

// NotaDetalleCreditoDebitoIce representa un ítem individual dentro del detalle.
type NotaDetalleCreditoDebitoIce struct {
	models.RequestWrapper[documents.DetalleNotaCreditoDebitoIce]
}

// NewNotaCreditoDebitoIceBuilder inicia el proceso de construcción de una Nota ICE.
func NewNotaCreditoDebitoIceBuilder() *notaCreditoDebitoIceBuilder {
	return &notaCreditoDebitoIceBuilder{
		nota: &documents.NotaCreditoDebitoIce{
			XMLName:           xml.Name{Local: "notaElectronicaCreditoDebitoIce"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "notaElectronicaCreditoDebitoIce.xsd",
		},
	}
}

// NewNotaCreditoDebitoIceCabeceraBuilder crea una nueva instancia del constructor para la cabecera.
func NewNotaCreditoDebitoIceCabeceraBuilder() *notaCreditoDebitoIceCabeceraBuilder {
	return &notaCreditoDebitoIceCabeceraBuilder{
		cabecera: &documents.CabeceraNotaCreditoDebitoIce{
			CodigoDocumentoSector: 48, // Sector 48 para Nota de Crédito/Débito ICE
		},
	}
}

// NewNotaDetalleCreditoDebitoIceBuilder crea una nueva instancia del constructor para el detalle.
func NewNotaDetalleCreditoDebitoIceBuilder() *notaDetalleCreditoDebitoIceBuilder {
	return &notaDetalleCreditoDebitoIceBuilder{
		detalle: &documents.DetalleNotaCreditoDebitoIce{},
	}
}

type notaCreditoDebitoIceBuilder struct {
	nota *documents.NotaCreditoDebitoIce
}

func (b *notaCreditoDebitoIceBuilder) WithCabecera(req NotaCreditoDebitoIceCabecera) *notaCreditoDebitoIceBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraNotaCreditoDebitoIce](req); internal != nil {
		b.nota.Cabecera = *internal
	}
	return b
}

func (b *notaCreditoDebitoIceBuilder) AddDetalle(req NotaDetalleCreditoDebitoIce) *notaCreditoDebitoIceBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleNotaCreditoDebitoIce](req); internal != nil {
		b.nota.Detalle = append(b.nota.Detalle, *internal)
	}
	return b
}

func (b *notaCreditoDebitoIceBuilder) WithModalidad(tipo int) *notaCreditoDebitoIceBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.nota.XMLName = xml.Name{Local: "notaElectronicaCreditoDebitoIce"}
		b.nota.XsiSchemaLocation = "notaElectronicaCreditoDebitoIce.xsd"
	case siat.ModalidadComputarizada:
		b.nota.XMLName = xml.Name{Local: "notaComputarizadaCreditoDebitoIce"}
		b.nota.XsiSchemaLocation = "notaComputarizadaCreditoDebitoIce.xsd"
	}
	return b
}

func (b *notaCreditoDebitoIceBuilder) Build() NotaCreditoDebitoIce {
	return NotaCreditoDebitoIce{models.NewRequestWrapper(b.nota)}
}

type notaCreditoDebitoIceCabeceraBuilder struct {
	cabecera *documents.CabeceraNotaCreditoDebitoIce
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithNitEmisor(v int64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithRazonSocialEmisor(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMunicipio(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithTelefono(v *string) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithNumeroNotaCreditoDebito(v int64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.NumeroNotaCreditoDebito = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCuf(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCufd(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoSucursal(v int) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithDireccion(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoPuntoVenta(v *int) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithFechaEmision(v time.Time) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithNombreRazonSocial(v *string) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithNumeroDocumento(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithComplemento(v *string) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoCliente(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithNumeroFactura(v int64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithNumeroAutorizacionCuf(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.NumeroAutorizacionCuf = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithFechaEmisionFactura(v time.Time) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.FechaEmisionFactura = datatype.NewTimeSiat(v)
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMontoTotalOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.MontoTotalOriginal = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithDescuentoAdicional(v *float64) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 2)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMontoTotalDevuelto(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.MontoTotalDevuelto = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMontoDescuentoCreditoDebito(v *float64) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 2)
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMontoEfectivoCreditoDebito(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.MontoEfectivoCreditoDebito = datatype.Float64Round(v, 2)
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoExcepcion(v *int) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithLeyenda(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithUsuario(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) Build() NotaCreditoDebitoIceCabecera {
	return NotaCreditoDebitoIceCabecera{models.NewRequestWrapper(b.cabecera)}
}

type notaDetalleCreditoDebitoIceBuilder struct {
	detalle *documents.DetalleNotaCreditoDebitoIce
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithNroItem(v int) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.NroItem = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithActividadEconomica(v string) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithCodigoProductoSin(v int64) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithCodigoProducto(v string) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithDescripcion(v string) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithCantidad(v float64) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.Cantidad = datatype.Float64Round(v, 5)
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithUnidadMedida(v int) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithPrecioUnitario(v float64) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.PrecioUnitario = datatype.Float64Round(v, 5)
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMontoDescuento(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithSubTotal(v float64) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.SubTotal = datatype.Float64Round(v, 5)
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMarcaIce(v int) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.MarcaIce = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithAlicuotaIva(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.AlicuotaIva = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.AlicuotaIva = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithPrecioNetoVentaIce(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithAlicuotaEspecifica(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithAlicuotaPorcentual(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMontoIceEspecifico(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMontoIcePorcentual(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithCantidadIce(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.CantidadIce = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 5)
		b.detalle.CantidadIce = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithCodigoDetalleTransaccion(v int) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.CodigoDetalleTransaccion = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) Build() NotaDetalleCreditoDebitoIce {
	return NotaDetalleCreditoDebitoIce{models.NewRequestWrapper(b.detalle)}
}
