package invoices

import (
	"encoding/xml"
	"time"

	"github.com/shopspring/decimal"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
)

// ZonaFranca representa la estructura completa de una factura Zona Franca lista para ser procesada.
type ZonaFranca struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaZonaFranca]
}

// ZonaFrancaCabecera representa la sección de cabecera de una factura Zona Franca.
type ZonaFrancaCabecera struct {
	models.RequestWrapper[documents.CabeceraZonaFranca]
}

// ZonaFrancaDetalle representa un ítem individual dentro del detalle de una factura Zona Franca.
type ZonaFrancaDetalle struct {
	models.RequestWrapper[documents.DetalleZonaFranca]
}

// NewZonaFrancaBuilder inicia el proceso de construcción de una Factura Zona Franca.
func NewZonaFrancaBuilder() *zonaFrancaBuilder {
	return &zonaFrancaBuilder{
		metadatos: models.NuevosMetadatosFactura(5),
		factura: &documents.FacturaZonaFranca{
			XMLName:           xml.Name{Local: "facturaElectronicaZonaFranca"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaZonaFranca.xsd",
		},
	}
}

// NewZonaFrancaCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewZonaFrancaCabeceraBuilder() *zonaFrancaCabeceraBuilder {
	return &zonaFrancaCabeceraBuilder{
		cabecera: &documents.CabeceraZonaFranca{
			CodigoDocumentoSector: 5, // Sector 5 para Zona Franca
		},
	}
}

// NewZonaFrancaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewZonaFrancaDetalleBuilder() *zonaFrancaDetalleBuilder {
	return &zonaFrancaDetalleBuilder{
		detalle: &documents.DetalleZonaFranca{},
	}
}

type zonaFrancaBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaZonaFranca
}

func (b *zonaFrancaBuilder) WithCabecera(req ZonaFrancaCabecera) *zonaFrancaBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraZonaFranca](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *zonaFrancaBuilder) AddDetalle(req ZonaFrancaDetalle) *zonaFrancaBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleZonaFranca](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *zonaFrancaBuilder) WithModalidad(tipo int) *zonaFrancaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaZonaFranca"}
		b.factura.XsiSchemaLocation = "facturaElectronicaZonaFranca.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaZonaFranca"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaZonaFranca.xsd"
	}
	return b
}

func (b *zonaFrancaBuilder) Build() ZonaFranca {
	return ZonaFranca{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type zonaFrancaCabeceraBuilder struct {
	cabecera *documents.CabeceraZonaFranca
}

func (b *zonaFrancaCabeceraBuilder) Build() ZonaFrancaCabecera {
	return ZonaFrancaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type zonaFrancaDetalleBuilder struct {
	detalle *documents.DetalleZonaFranca
}

func (b *zonaFrancaDetalleBuilder) Build() ZonaFrancaDetalle {
	return ZonaFrancaDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *zonaFrancaCabeceraBuilder) WithNitEmisor(v int64) *zonaFrancaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithRazonSocialEmisor(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMunicipio(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithTelefono(v *string) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroFactura(v int64) *zonaFrancaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCuf(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCufd(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoSucursal(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithDireccion(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithFechaEmision(v time.Time) *zonaFrancaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNombreRazonSocial(v *string) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroDocumento(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithComplemento(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoCliente(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoMetodoPago(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroTarjeta(v *int64) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithSubTotalNeto(v float64) *zonaFrancaCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithTotalPrecioNeto(v float64) *zonaFrancaCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithIva(v float64) *zonaFrancaCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMontoAPagar(v float64) *zonaFrancaCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoMoneda(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithTipoCambio(v float64) *zonaFrancaCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMontoTotalMoneda(v float64) *zonaFrancaCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroParteRecepcion(v *string) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroParteRecepcion = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NumeroParteRecepcion = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMontoGiftCard(v *float64) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithDescuentoAdicional(v *float64) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoExcepcion(v *int) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithLeyenda(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithUsuario(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoInterno(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoDocumentoSector(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithActividadEconomica(v string) *zonaFrancaDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithCodigoProductoSin(v int64) *zonaFrancaDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithCodigoProducto(v string) *zonaFrancaDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithDescripcion(v string) *zonaFrancaDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithCantidad(v float64) *zonaFrancaDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaDetalleBuilder) WithUnidadMedida(v int) *zonaFrancaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithPrecioNetoUnitario(v float64) *zonaFrancaDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaDetalleBuilder) WithMontoDescuento(v *float64) *zonaFrancaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *zonaFrancaDetalleBuilder) WithSubTotal(v float64) *zonaFrancaDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *zonaFrancaDetalleBuilder) WithNumeroSerie(v *string) *zonaFrancaDetalleBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaDetalleBuilder) WithNumeroImei(v *string) *zonaFrancaDetalleBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &value}
	return b
}
