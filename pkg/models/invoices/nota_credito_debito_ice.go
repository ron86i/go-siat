package invoices

import (
	"encoding/xml"
	"github.com/shopspring/decimal"
	"time"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/ron86i/go-siat/v3/pkg/models"
)

// NotaCreditoDebitoIce representa la estructura completa de una nota de crédito/débito/descuento ICE lista para procesarse.
type NotaCreditoDebitoIce struct {
	models.RequestWrapper[documents.NotaCreditoDebitoIce]
	models.MetadatosFactura
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
		metadatos: models.NuevosMetadatosFactura(48),
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
	metadatos models.MetadatosFactura
	nota      *documents.NotaCreditoDebitoIce
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
	return NotaCreditoDebitoIce{RequestWrapper: models.NewRequestWrapper(b.nota), MetadatosFactura: b.metadatos}
}

type notaCreditoDebitoIceCabeceraBuilder struct {
	cabecera *documents.CabeceraNotaCreditoDebitoIce
}

func (b *notaCreditoDebitoIceCabeceraBuilder) Build() NotaCreditoDebitoIceCabecera {
	return NotaCreditoDebitoIceCabecera{models.NewRequestWrapper(b.cabecera)}
}

type notaDetalleCreditoDebitoIceBuilder struct {
	detalle *documents.DetalleNotaCreditoDebitoIce
}

func (b *notaDetalleCreditoDebitoIceBuilder) Build() NotaDetalleCreditoDebitoIce {
	return NotaDetalleCreditoDebitoIce{models.NewRequestWrapper(b.detalle)}
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
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
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
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithFechaEmision(v time.Time) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithNombreRazonSocial(v *string) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
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

func (b *notaCreditoDebitoIceCabeceraBuilder) WithComplemento(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.Complemento = v
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

func (b *notaCreditoDebitoIceCabeceraBuilder) WithSubTotalNetoOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.SubTotalNetoOriginal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithTotalPrecioNetoOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.TotalPrecioNetoOriginal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithTotalIcePorcentualOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.TotalIcePorcentualOriginal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithTotalIceEspecificoOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.TotalIceEspecificoOriginal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithIvaOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.IvaOriginal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMontoAPagarOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.MontoAPagarOriginal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithSubTotalNetoDevuelto(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.SubTotalNetoDevuelto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithTotalPrecioNetoDevuelto(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.TotalPrecioNetoDevuelto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithIvaDevuelto(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.IvaDevuelto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMontoAPagarDevuelto(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.MontoAPagarDevuelto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithDescuentoAdicionalOriginal(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.DescuentoAdicionalOriginal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithMontoDescuentoDevuelto(v float64) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.MontoDescuentoDevuelto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoExcepcion(v *int) *notaCreditoDebitoIceCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
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

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoInterno(v string) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *notaCreditoDebitoIceCabeceraBuilder) WithCodigoDocumentoSector(v int) *notaCreditoDebitoIceCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
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
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithUnidadMedida(v int) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithPrecioNetoUnitario(v float64) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMontoDescuento(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithSubTotal(v float64) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMarcaIce(v int) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.MarcaIce = v
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithPrecioNetoVentaIce(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithAlicuotaEspecifica(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithAlicuotaPorcentual(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMontoIceEspecifico(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithMontoIcePorcentual(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithCantidadIce(v *float64) *notaDetalleCreditoDebitoIceBuilder {
	if v == nil {
		b.detalle.CantidadIce = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.CantidadIce = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *notaDetalleCreditoDebitoIceBuilder) WithCodigoDetalleTransaccion(v int) *notaDetalleCreditoDebitoIceBuilder {
	b.detalle.CodigoDetalleTransaccion = v
	return b
}
