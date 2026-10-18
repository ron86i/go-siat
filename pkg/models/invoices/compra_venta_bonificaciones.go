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

// CompraVentaBonificaciones representa una factura de compra-venta con bonificaciones lista para procesar.
type CompraVentaBonificaciones struct {
	models.RequestWrapper[documents.FacturaCompraVentaBonificaciones]
	models.MetadatosFactura
}

// CompraVentaBonificacionesCabecera representa la cabecera de una factura de bonificaciones.
type CompraVentaBonificacionesCabecera struct {
	models.RequestWrapper[documents.CabeceraCompraVentaBonificaciones]
}

// CompraVentaBonificacionesDetalle representa un ítem de detalle de una factura de bonificaciones.
type CompraVentaBonificacionesDetalle struct {
	models.RequestWrapper[documents.DetalleCompraVentaBonificaciones]
}

// NewCompraVentaBonificacionesBuilder inicia la construcción de una factura de compra-venta con bonificaciones.
func NewCompraVentaBonificacionesBuilder() *compraVentaBonificacionesBuilder {
	return &compraVentaBonificacionesBuilder{
		metadatos: models.NuevosMetadatosFactura(35),
		factura: &documents.FacturaCompraVentaBonificaciones{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVentaBon"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVentaBon.xsd",
		},
	}
}

// NewCompraVentaBonificacionesCabeceraBuilder crea el constructor de la cabecera de bonificaciones.
func NewCompraVentaBonificacionesCabeceraBuilder() *compraVentaBonificacionesCabeceraBuilder {
	return &compraVentaBonificacionesCabeceraBuilder{
		cabecera: &documents.CabeceraCompraVentaBonificaciones{
			CodigoDocumentoSector: 35, // Sector 35 para Compra y Venta Bonificaciones
		},
	}
}

// NewCompraVentaBonificacionesDetalleBuilder crea el constructor del detalle de bonificaciones.
func NewCompraVentaBonificacionesDetalleBuilder() *detalleBonificacionesBuilder {
	return &detalleBonificacionesBuilder{
		detalle: &documents.DetalleCompraVentaBonificaciones{},
	}
}

// --- Builder Factura Bonificaciones ---

type compraVentaBonificacionesBuilder struct {
	factura   *documents.FacturaCompraVentaBonificaciones
	metadatos models.MetadatosFactura
}

// WithTipoFacturaDocumento reemplaza el tipo fiscal predeterminado de bonificaciones.
func (b *compraVentaBonificacionesBuilder) WithTipoFacturaDocumento(tipo int) *compraVentaBonificacionesBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

// WithModalidad configura los metadatos XML según la modalidad.
func (b *compraVentaBonificacionesBuilder) WithModalidad(tipo int) *compraVentaBonificacionesBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVentaBon"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVentaBon.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVentaBon"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVentaBon.xsd"
	}
	return b
}

func (b *compraVentaBonificacionesBuilder) WithCabecera(req CompraVentaBonificacionesCabecera) *compraVentaBonificacionesBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraCompraVentaBonificaciones](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *compraVentaBonificacionesBuilder) AddDetalle(req CompraVentaBonificacionesDetalle) *compraVentaBonificacionesBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleCompraVentaBonificaciones](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *compraVentaBonificacionesBuilder) Build() CompraVentaBonificaciones {
	return CompraVentaBonificaciones{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

// --- Builder Cabecera Bonificaciones ---

type compraVentaBonificacionesCabeceraBuilder struct {
	cabecera *documents.CabeceraCompraVentaBonificaciones
}

func (b *compraVentaBonificacionesCabeceraBuilder) Build() CompraVentaBonificacionesCabecera {
	return CompraVentaBonificacionesCabecera{models.NewRequestWrapper(b.cabecera)}
}

// --- Builder Detalle Bonificaciones ---

type detalleBonificacionesBuilder struct {
	detalle *documents.DetalleCompraVentaBonificaciones
}

func (b *detalleBonificacionesBuilder) Build() CompraVentaBonificacionesDetalle {
	return CompraVentaBonificacionesDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithNitEmisor(v int64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithRazonSocialEmisor(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithMunicipio(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithTelefono(v *string) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithNumeroFactura(v int64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCuf(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCufd(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoSucursal(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithDireccion(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoPuntoVenta(v *int) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithFechaEmision(v time.Time) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithNombreRazonSocial(v *string) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithNumeroDocumento(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithComplemento(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoCliente(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoMetodoPago(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithNumeroTarjeta(v *int64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithSubTotalNeto(v float64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithTotalPrecioNeto(v float64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithIva(v float64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithMontoAPagar(v float64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoMoneda(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithTipoCambio(v float64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithMontoTotalMoneda(v float64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithMontoGiftCard(v *float64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithDescuentoAdicional(v *float64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoExcepcion(v *int64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithLeyenda(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithUsuario(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoInterno(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoDocumentoSector(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *detalleBonificacionesBuilder) WithActividadEconomica(v string) *detalleBonificacionesBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *detalleBonificacionesBuilder) WithCodigoProductoSin(v int64) *detalleBonificacionesBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *detalleBonificacionesBuilder) WithCodigoProducto(v string) *detalleBonificacionesBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *detalleBonificacionesBuilder) WithDescripcion(v string) *detalleBonificacionesBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *detalleBonificacionesBuilder) WithCantidad(v float64) *detalleBonificacionesBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *detalleBonificacionesBuilder) WithUnidadMedida(v int) *detalleBonificacionesBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *detalleBonificacionesBuilder) WithPrecioNetoUnitario(v float64) *detalleBonificacionesBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *detalleBonificacionesBuilder) WithMontoDescuento(v *float64) *detalleBonificacionesBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *detalleBonificacionesBuilder) WithSubTotal(v float64) *detalleBonificacionesBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *detalleBonificacionesBuilder) WithNumeroSerie(v *string) *detalleBonificacionesBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *detalleBonificacionesBuilder) WithNumeroImei(v *string) *detalleBonificacionesBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &value}
	return b
}
