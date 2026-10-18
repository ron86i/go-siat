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

// VentaMineral representa la estructura completa de una factura de Venta de Minerales lista para ser procesada.
type VentaMineral struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaVentaMineral]
}

// VentaMineralCabecera representa la sección de cabecera de la factura de minerales.
type VentaMineralCabecera struct {
	models.RequestWrapper[documents.CabeceraVentaMineral]
}

// VentaMineralDetalle representa un ítem individual dentro del detalle de la factura de minerales.
type VentaMineralDetalle struct {
	models.RequestWrapper[documents.DetalleVentaMineral]
}

// NewVentaMineralBuilder inicia el proceso de construcción de la factura.
func NewVentaMineralBuilder() *ventaMineralBuilder {
	return &ventaMineralBuilder{
		metadatos: models.NuevosMetadatosFactura(21),
		factura: &documents.FacturaVentaMineral{
			XMLName:           xml.Name{Local: "facturaElectronicaVentaMineral"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaVentaMineral.xsd",
		},
	}
}

// NewVentaMineralCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewVentaMineralCabeceraBuilder() *ventaMineralCabeceraBuilder {
	return &ventaMineralCabeceraBuilder{
		cabecera: &documents.CabeceraVentaMineral{
			CodigoDocumentoSector: 21, // Sector 21
		},
	}
}

// NewVentaMineralDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewVentaMineralDetalleBuilder() *ventaMineralDetalleBuilder {
	return &ventaMineralDetalleBuilder{
		detalle: &documents.DetalleVentaMineral{
			MontoDescuento: 0, // Fijo 0 según XSD
		},
	}
}

type ventaMineralBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaVentaMineral
}

func (b *ventaMineralBuilder) WithCabecera(req VentaMineralCabecera) *ventaMineralBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraVentaMineral](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *ventaMineralBuilder) AddDetalle(req VentaMineralDetalle) *ventaMineralBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleVentaMineral](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *ventaMineralBuilder) WithModalidad(tipo int) *ventaMineralBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaVentaMineral"}
		b.factura.XsiSchemaLocation = "facturaElectronicaVentaMineral.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaVentaMineral"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaVentaMineral.xsd"
	}
	return b
}

func (b *ventaMineralBuilder) Build() VentaMineral {
	return VentaMineral{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type ventaMineralCabeceraBuilder struct {
	cabecera *documents.CabeceraVentaMineral
}

func (b *ventaMineralCabeceraBuilder) Build() VentaMineralCabecera {
	return VentaMineralCabecera{models.NewRequestWrapper(b.cabecera)}
}

type ventaMineralDetalleBuilder struct {
	detalle *documents.DetalleVentaMineral
}

func (b *ventaMineralDetalleBuilder) Build() VentaMineralDetalle {
	return VentaMineralDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *ventaMineralCabeceraBuilder) WithNitEmisor(v int64) *ventaMineralCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithRazonSocialEmisor(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithMunicipio(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithTelefono(v *string) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithNumeroFactura(v int64) *ventaMineralCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCuf(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCufd(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoSucursal(v int) *ventaMineralCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithDireccion(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoPuntoVenta(v *int) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithFechaEmision(v time.Time) *ventaMineralCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *ventaMineralCabeceraBuilder) WithNombreRazonSocial(v *string) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithDireccionComprador(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *ventaMineralCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithNumeroDocumento(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithComplemento(v *string) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithNim(v *string) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.Nim = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Nim = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithDescripcionProducto(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.DescripcionProducto = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithLocalidad(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.Localidad = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithDepartamentoOrigen(v *string) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.DepartamentoOrigen = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.DepartamentoOrigen = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithMunicipioOrigen(v *string) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.MunicipioOrigen = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.MunicipioOrigen = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoMoneda(v int) *ventaMineralCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithTipoCambio(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithNumeroLote(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.NumeroLote = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithNumeroContrato(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.NumeroContrato = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithKilosNetosHumedos(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.KilosNetosHumedos = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithHumedadPorcentaje(v *float64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.HumedadPorcentaje = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.HumedadPorcentaje = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithHumedadValor(v *float64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.HumedadValor = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.HumedadValor = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithMermaPorcentaje(v *float64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.MermaPorcentaje = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MermaPorcentaje = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithMermaValor(v *float64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.MermaValor = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MermaValor = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithKilosNetosSecos(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.KilosNetosSecos = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoMetodoPago(v int) *ventaMineralCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithNumeroTarjeta(v *int64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithSubTotalNeto(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithTotalPrecioNeto(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithIva(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithMontoAPagar(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithMontoTotalMoneda(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithGastosRealizacion(v float64) *ventaMineralCabeceraBuilder {
	b.cabecera.GastosRealizacion = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralCabeceraBuilder) WithLiquidacionPreliminar(v *float64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.LiquidacionPreliminar = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.LiquidacionPreliminar = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithOtrosDatos(v *string) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.OtrosDatos = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.OtrosDatos = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithMontoGiftCard(v *float64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithDescuentoAdicional(v *float64) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoExcepcion(v *int) *ventaMineralCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaMineralCabeceraBuilder) WithLeyenda(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithUsuario(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoInterno(v string) *ventaMineralCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *ventaMineralCabeceraBuilder) WithCodigoDocumentoSector(v int) *ventaMineralCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithActividadEconomica(v string) *ventaMineralDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithCodigoProductoSin(v int64) *ventaMineralDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithCodigoProducto(v string) *ventaMineralDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithCodigoNandina(v string) *ventaMineralDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithDescripcion(v string) *ventaMineralDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithDescripcionLeyes(v string) *ventaMineralDetalleBuilder {
	b.detalle.DescripcionLeyes = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithCantidadExtraccion(v float64) *ventaMineralDetalleBuilder {
	b.detalle.CantidadExtraccion = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaMineralDetalleBuilder) WithCantidad(v float64) *ventaMineralDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaMineralDetalleBuilder) WithUnidadMedidaExtraccion(v int) *ventaMineralDetalleBuilder {
	b.detalle.UnidadMedidaExtraccion = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithUnidadMedida(v int) *ventaMineralDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *ventaMineralDetalleBuilder) WithPrecioNetoUnitario(v float64) *ventaMineralDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaMineralDetalleBuilder) WithMontoDescuento(v float64) *ventaMineralDetalleBuilder {
	b.detalle.MontoDescuento = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaMineralDetalleBuilder) WithSubTotal(v float64) *ventaMineralDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}
