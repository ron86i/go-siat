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

// VentaMineralBCB representa la estructura completa de una factura de Venta de Minerales para el BCB lista para ser procesada.
type VentaMineralBCB struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaVentaMineralBCB]
}

// VentaMineralBCBCabecera representa la sección de cabecera de la factura.
type VentaMineralBCBCabecera struct {
	models.RequestWrapper[documents.CabeceraVentaMineralBCB]
}

// VentaMineralBCBDetalle representa un ítem individual dentro del detalle.
type VentaMineralBCBDetalle struct {
	models.RequestWrapper[documents.DetalleVentaMineralBCB]
}

// NewVentaMineralBCBBuilder inicia el proceso de construcción de la factura.
func NewVentaMineralBCBBuilder() *ventaMineralBCBBuilder {
	return &ventaMineralBCBBuilder{
		metadatos: models.NuevosMetadatosFactura(52),
		factura: &documents.FacturaVentaMineralBCB{
			XMLName:           xml.Name{Local: "facturaElectronicaVentaMineralBCB"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaVentaMineralBCB.xsd",
		},
	}
}

// NewVentaMineralBCBCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewVentaMineralBCBCabeceraBuilder() *ventaMineralBCBCabeceraBuilder {
	return &ventaMineralBCBCabeceraBuilder{
		cabecera: &documents.CabeceraVentaMineralBCB{
			CodigoDocumentoSector: 52, // Sector 52
			TipoCambio:            1,  // Fijo 1 según XSD
		},
	}
}

// NewVentaMineralBCBDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewVentaMineralBCBDetalleBuilder() *ventaMineralBCBDetalleBuilder {
	return &ventaMineralBCBDetalleBuilder{
		detalle: &documents.DetalleVentaMineralBCB{
			MontoDescuento: 0, // Fijo 0 según XSD
		},
	}
}

type ventaMineralBCBBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaVentaMineralBCB
}

func (b *ventaMineralBCBBuilder) WithCabecera(req VentaMineralBCBCabecera) *ventaMineralBCBBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraVentaMineralBCB](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *ventaMineralBCBBuilder) AddDetalle(req VentaMineralBCBDetalle) *ventaMineralBCBBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleVentaMineralBCB](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

// WithModalidad conserva el formato electrónico de Venta Mineral BCB.
// El paquete XSD v3 sólo define esta modalidad; otros valores conservan el formato actual.
func (b *ventaMineralBCBBuilder) WithModalidad(tipo int) *ventaMineralBCBBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaVentaMineralBCB"}
		b.factura.XsiSchemaLocation = "facturaElectronicaVentaMineralBCB.xsd"
	}
	return b
}

func (b *ventaMineralBCBBuilder) Build() VentaMineralBCB {
	return VentaMineralBCB{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type ventaMineralBCBCabeceraBuilder struct {
	cabecera *documents.CabeceraVentaMineralBCB
}

func (b *ventaMineralBCBCabeceraBuilder) Build() VentaMineralBCBCabecera {
	return VentaMineralBCBCabecera{models.NewRequestWrapper(b.cabecera)}
}

type ventaMineralBCBDetalleBuilder struct {
	detalle *documents.DetalleVentaMineralBCB
}

func (b *ventaMineralBCBDetalleBuilder) Build() VentaMineralBCBDetalle {
	return VentaMineralBCBDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *ventaMineralBCBCabeceraBuilder) WithNitEmisor(v int64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithRazonSocialEmisor(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithMunicipio(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithTelefono(v *string) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithNumeroFactura(v int64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCuf(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCufd(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoSucursal(v int) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithDireccion(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoPuntoVenta(v *int) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithFechaEmision(v time.Time) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithNombreRazonSocial(v *string) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithDireccionComprador(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithNumeroDocumento(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithComplemento(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithConcentradoGranel(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.ConcentradoGranel = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithOrigen(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Origen = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoCliente(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoMoneda(v int) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithTipoCambio(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithNumeroLote(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.NumeroLote = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithKilosNetosHumedos(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.KilosNetosHumedos = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithHumedadPorcentaje(v *float64) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.HumedadPorcentaje = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.HumedadPorcentaje = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithHumedadValor(v *float64) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.HumedadValor = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.HumedadValor = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithMermaPorcentaje(v *float64) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.MermaPorcentaje = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MermaPorcentaje = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithMermaValor(v *float64) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.MermaValor = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MermaValor = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithKilosNetosSecos(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.KilosNetosSecos = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoMetodoPago(v int) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithNumeroTarjeta(v *int64) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithSubTotalNeto(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithTotalPrecioNeto(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithIva(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithMontoAPagar(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithMontoTotalMoneda(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithGastosRealizacion(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.GastosRealizacion = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithLiquidacionPreliminar(v *float64) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.LiquidacionPreliminar = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.LiquidacionPreliminar = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithOtrosDatos(v *string) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.OtrosDatos = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.OtrosDatos = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithDescuentoAdicional(v float64) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.DescuentoAdicional = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoExcepcion(v *int) *ventaMineralBCBCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithLeyenda(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithUsuario(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoInterno(v string) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *ventaMineralBCBCabeceraBuilder) WithCodigoDocumentoSector(v int) *ventaMineralBCBCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithActividadEconomica(v string) *ventaMineralBCBDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithCodigoProductoSin(v int64) *ventaMineralBCBDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithCodigoProducto(v string) *ventaMineralBCBDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithCodigoNandina(v string) *ventaMineralBCBDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithDescripcion(v string) *ventaMineralBCBDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithDescripcionLeyes(v string) *ventaMineralBCBDetalleBuilder {
	b.detalle.DescripcionLeyes = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithCantidadExtraccion(v float64) *ventaMineralBCBDetalleBuilder {
	b.detalle.CantidadExtraccion = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithCantidad(v float64) *ventaMineralBCBDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithUnidadMedidaExtraccion(v int) *ventaMineralBCBDetalleBuilder {
	b.detalle.UnidadMedidaExtraccion = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithUnidadMedida(v int) *ventaMineralBCBDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithPrecioNetoUnitario(v float64) *ventaMineralBCBDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithMontoDescuento(v int) *ventaMineralBCBDetalleBuilder {
	b.detalle.MontoDescuento = v
	return b
}

func (b *ventaMineralBCBDetalleBuilder) WithSubTotal(v float64) *ventaMineralBCBDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}
