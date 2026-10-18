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

// CompraVenta representa la estructura completa de una factura lista para ser procesada.
type CompraVenta struct {
	models.RequestWrapper[documents.FacturaCompraVenta]
	models.MetadatosFactura
}

// CompraVentaCabecera representa la sección de cabecera de una factura de compra y venta.
type CompraVentaCabecera struct {
	models.RequestWrapper[documents.CabeceraCompraVenta]
}

// CompraVentaDetalle representa un ítem individual dentro del detalle de una factura.
type CompraVentaDetalle struct {
	models.RequestWrapper[documents.DetalleCompraVenta]
}

// NewCompraVentaBuilder inicia el proceso de construcción de una CompraVenta,
// configurando los nombres de nodo XML necesarios para el SIAT.
//
// Por defecto, la factura se configura para ser emitida electrónica.
func NewCompraVentaBuilder() *compraVentaBuilder {
	return &compraVentaBuilder{
		metadatos: models.NuevosMetadatosFactura(1),
		factura: &documents.FacturaCompraVenta{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVenta"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVenta.xsd",
		},
	}
}

// NewCompraVentaCabeceraBuilder crea una nueva instancia del constructor para la cabecera
// de facturas de compra y venta.
func NewCompraVentaCabeceraBuilder() *compraVentaCabeceraBuilder {
	return &compraVentaCabeceraBuilder{
		cabecera: &documents.CabeceraCompraVenta{
			CodigoDocumentoSector: 1, // Sector 1 para Compra y Venta
		},
	}
}

// NewCompraVentaDetalleBuilder crea una nueva instancia del constructor para los ítems
// de detalle de la factura.
func NewCompraVentaDetalleBuilder() *detalleBuilder {
	return &detalleBuilder{
		detalle: &documents.DetalleCompraVenta{},
	}
}

type compraVentaBuilder struct {
	factura   *documents.FacturaCompraVenta
	metadatos models.MetadatosFactura
}

// WithTipoFacturaDocumento reemplaza el tipo fiscal predeterminado de compra y venta.
func (b *compraVentaBuilder) WithTipoFacturaDocumento(tipo int) *compraVentaBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

// WithCabecera asocia la cabecera construida previamente a la factura.
func (b *compraVentaBuilder) WithCabecera(req CompraVentaCabecera) *compraVentaBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraCompraVenta](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

// AddDetalle añade un ítem de detalle a la lista de detalles de la factura.
func (b *compraVentaBuilder) AddDetalle(req CompraVentaDetalle) *compraVentaBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleCompraVenta](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

// WithModalidad configura los metadatos XML de la factura según la modalidad (Electrónica o Computarizada).
func (b *compraVentaBuilder) WithModalidad(tipo int) *compraVentaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVenta.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVenta.xsd"
	}
	return b
}

// Build finaliza la construcción y retorna la estructura opaca lista para ser firmada y enviada.
func (b *compraVentaBuilder) Build() CompraVenta {
	return CompraVenta{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

// compraVentaCabeceraBuilder ayuda a configurar la cabecera de la factura de compra y venta.
type compraVentaCabeceraBuilder struct {
	cabecera *documents.CabeceraCompraVenta
}

// Build finaliza la construcción de la cabecera retornando la estructura opaca.
func (b *compraVentaCabeceraBuilder) Build() CompraVentaCabecera {
	return CompraVentaCabecera{models.NewRequestWrapper(b.cabecera)}
}

// detalleBuilder ayuda a configurar un ítem individual de detalle de la factura.
type detalleBuilder struct {
	detalle *documents.DetalleCompraVenta
}

// Build finaliza la construcción del detalle retornando la estructura opaca.
func (b *detalleBuilder) Build() CompraVentaDetalle {
	return CompraVentaDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *compraVentaCabeceraBuilder) WithNitEmisor(v int64) *compraVentaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithRazonSocialEmisor(v string) *compraVentaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithMunicipio(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithTelefono(v *string) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *compraVentaCabeceraBuilder) WithNumeroFactura(v int64) *compraVentaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCuf(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCufd(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoSucursal(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithDireccion(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *compraVentaCabeceraBuilder) WithFechaEmision(v time.Time) *compraVentaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *compraVentaCabeceraBuilder) WithNombreRazonSocial(v *string) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithNumeroDocumento(v string) *compraVentaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithComplemento(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoCliente(v string) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoMetodoPago(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithNumeroTarjeta(v *int64) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *compraVentaCabeceraBuilder) WithSubTotalNeto(v float64) *compraVentaCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaCabeceraBuilder) WithTotalPrecioNeto(v float64) *compraVentaCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaCabeceraBuilder) WithIva(v float64) *compraVentaCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaCabeceraBuilder) WithMontoAPagar(v float64) *compraVentaCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoMoneda(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithTipoCambio(v float64) *compraVentaCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaCabeceraBuilder) WithMontoTotalMoneda(v float64) *compraVentaCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaCabeceraBuilder) WithMontoGiftCard(v *float64) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *compraVentaCabeceraBuilder) WithDescuentoAdicional(v *float64) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoExcepcion(v *int64) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *compraVentaCabeceraBuilder) WithLeyenda(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithUsuario(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoInterno(v string) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *compraVentaCabeceraBuilder) WithCodigoDocumentoSector(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *detalleBuilder) WithActividadEconomica(v string) *detalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *detalleBuilder) WithCodigoProductoSin(v int64) *detalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *detalleBuilder) WithCodigoProducto(v string) *detalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *detalleBuilder) WithDescripcion(v string) *detalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *detalleBuilder) WithCantidad(v float64) *detalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *detalleBuilder) WithUnidadMedida(v int) *detalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *detalleBuilder) WithPrecioNetoUnitario(v float64) *detalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *detalleBuilder) WithMontoDescuento(v *float64) *detalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *detalleBuilder) WithSubTotal(v float64) *detalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *detalleBuilder) WithNumeroSerie(v *string) *detalleBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *detalleBuilder) WithNumeroImei(v *string) *detalleBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &value}
	return b
}
