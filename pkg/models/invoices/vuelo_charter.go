package invoices

import (
	"encoding/xml"
	"time"

	"github.com/shopspring/decimal"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/ron86i/go-siat/v3/pkg/models"
)

// VueloCharter representa la estructura completa de una factura lista para ser procesada.
type VueloCharter struct {
	models.RequestWrapper[documents.FacturaVueloCharter]
	models.MetadatosFactura
}

// VueloCharterCabecera representa la sección de cabecera de una factura de vuelo chárter.
type VueloCharterCabecera struct {
	models.RequestWrapper[documents.CabeceraVueloCharter]
}

// VueloCharterDetalle representa un ítem individual dentro del detalle de una factura.
type VueloCharterDetalle struct {
	models.RequestWrapper[documents.DetalleVueloCharter]
}

// NewVueloCharterBuilder inicia el proceso de construcción de una VueloCharter,
// configurando los nombres de nodo XML necesarios para el SIAT.
//
// Por defecto, la factura se configura para ser emitida electrónica.
func NewVueloCharterBuilder() *vueloCharterBuilder {
	return &vueloCharterBuilder{
		metadatos: models.NuevosMetadatosFactura(41),
		factura: &documents.FacturaVueloCharter{
			XMLName:           xml.Name{Local: "facturaElectronicaVueloCharter"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaVueloCharter.xsd",
		},
	}
}

// NewVueloCharterCabeceraBuilder crea una nueva instancia del constructor para la cabecera
// de facturas de vuelo chárter.
func NewVueloCharterCabeceraBuilder() *vueloCharterCabeceraBuilder {
	return &vueloCharterCabeceraBuilder{
		cabecera: &documents.CabeceraVueloCharter{
			CodigoDocumentoSector: 41, // Sector 41 para Vuelo Chárter
		},
	}
}

// NewVueloCharterDetalleBuilder crea una nueva instancia del constructor para los ítems
// de detalle de la factura.
func NewVueloCharterDetalleBuilder() *vueloCharterDetalleBuilder {
	return &vueloCharterDetalleBuilder{
		detalle: &documents.DetalleVueloCharter{},
	}
}

type vueloCharterBuilder struct {
	factura   *documents.FacturaVueloCharter
	metadatos models.MetadatosFactura
}

// WithTipoFacturaDocumento reemplaza el tipo fiscal predeterminado de vuelo chárter.
func (b *vueloCharterBuilder) WithTipoFacturaDocumento(tipo int) *vueloCharterBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

// WithCabecera asocia la cabecera construida previamente a la factura.
func (b *vueloCharterBuilder) WithCabecera(req VueloCharterCabecera) *vueloCharterBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraVueloCharter](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

// AddDetalle añade un ítem de detalle a la lista de detalles de la factura.
func (b *vueloCharterBuilder) AddDetalle(req VueloCharterDetalle) *vueloCharterBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleVueloCharter](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

// WithModalidad configura los metadatos XML de la factura según la modalidad (Electrónica o Computarizada).
func (b *vueloCharterBuilder) WithModalidad(tipo int) *vueloCharterBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaVueloCharter"}
		b.factura.XsiSchemaLocation = "facturaElectronicaVueloCharter.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaVueloCharter"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaVueloCharter.xsd"
	}
	return b
}

// Build finaliza la construcción y retorna la estructura opaca lista para ser firmada y enviada.
func (b *vueloCharterBuilder) Build() VueloCharter {
	return VueloCharter{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

// vueloCharterCabeceraBuilder ayuda a configurar la cabecera de la factura de vuelo chárter.
type vueloCharterCabeceraBuilder struct {
	cabecera *documents.CabeceraVueloCharter
}

// Build finaliza la construcción de la cabecera retornando la estructura opaca.
func (b *vueloCharterCabeceraBuilder) Build() VueloCharterCabecera {
	return VueloCharterCabecera{models.NewRequestWrapper(b.cabecera)}
}

// vueloCharterDetalleBuilder ayuda a configurar un ítem individual de detalle de la factura.
type vueloCharterDetalleBuilder struct {
	detalle *documents.DetalleVueloCharter
}

// Build finaliza la construcción del detalle retornando la estructura opaca.
func (b *vueloCharterDetalleBuilder) Build() VueloCharterDetalle {
	return VueloCharterDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *vueloCharterCabeceraBuilder) WithNitEmisor(v int64) *vueloCharterCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithRazonSocialEmisor(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithMunicipio(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithTelefono(v *string) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithNumeroFactura(v int64) *vueloCharterCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCuf(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCufd(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoSucursal(v int) *vueloCharterCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithDireccion(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoPuntoVenta(v *int) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithFechaEmision(v time.Time) *vueloCharterCabeceraBuilder {
	b.cabecera.FechaEmision = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithNombreRazonSocial(v *string) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *vueloCharterCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithNumeroDocumento(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithComplemento(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoCliente(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoMetodoPago(v int) *vueloCharterCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithNumeroTarjeta(v *int64) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithSubTotalNeto(v float64) *vueloCharterCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *vueloCharterCabeceraBuilder) WithTotalPrecioNeto(v float64) *vueloCharterCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *vueloCharterCabeceraBuilder) WithIva(v float64) *vueloCharterCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *vueloCharterCabeceraBuilder) WithMontoAPagar(v float64) *vueloCharterCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoMoneda(v int) *vueloCharterCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithTipoCambio(v float64) *vueloCharterCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *vueloCharterCabeceraBuilder) WithMontoTotalMoneda(v float64) *vueloCharterCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *vueloCharterCabeceraBuilder) WithMontoGiftCard(v *float64) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithMontoTasa(v *float64) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoTasa = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoTasa = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithDescuentoAdicional(v *float64) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoExcepcion(v *int64) *vueloCharterCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *vueloCharterCabeceraBuilder) WithLeyenda(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithUsuario(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoInterno(v string) *vueloCharterCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *vueloCharterCabeceraBuilder) WithCodigoDocumentoSector(v int) *vueloCharterCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *vueloCharterDetalleBuilder) WithActividadEconomica(v string) *vueloCharterDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *vueloCharterDetalleBuilder) WithCodigoProductoSin(v int64) *vueloCharterDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *vueloCharterDetalleBuilder) WithCodigoProducto(v string) *vueloCharterDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *vueloCharterDetalleBuilder) WithDescripcion(v string) *vueloCharterDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *vueloCharterDetalleBuilder) WithCantidad(v float64) *vueloCharterDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *vueloCharterDetalleBuilder) WithUnidadMedida(v int) *vueloCharterDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *vueloCharterDetalleBuilder) WithPrecioNetoUnitario(v float64) *vueloCharterDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *vueloCharterDetalleBuilder) WithMontoDescuento(v *float64) *vueloCharterDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *vueloCharterDetalleBuilder) WithSubTotal(v float64) *vueloCharterDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *vueloCharterDetalleBuilder) WithNumeroSerie(v *string) *vueloCharterDetalleBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *vueloCharterDetalleBuilder) WithNumeroImei(v *string) *vueloCharterDetalleBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &value}
	return b
}
