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

// VentaCombustibleSinSubvencion representa la estructura completa de una factura de Venta de Combustible Sin Subvención lista para ser procesada.
type VentaCombustibleSinSubvencion struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaVentaCombustibleSinSubvencion]
}

// VentaCombustibleSinSubvencionCabecera representa la sección de cabecera de una factura de Venta de Combustible Sin Subvención.
type VentaCombustibleSinSubvencionCabecera struct {
	models.RequestWrapper[documents.CabeceraVentaCombustibleSinSubvencion]
}

// VentaCombustibleSinSubvencionDetalle representa un ítem individual dentro del detalle de una factura de Venta de Combustible Sin Subvención.
type VentaCombustibleSinSubvencionDetalle struct {
	models.RequestWrapper[documents.DetalleVentaCombustibleSinSubvencion]
}

// NewVentaCombustibleSinSubvencionBuilder inicia el proceso de construcción de una Factura de Venta de Combustible Sin Subvención.
func NewVentaCombustibleSinSubvencionBuilder() *ventaCombustibleSinSubvencionBuilder {
	return &ventaCombustibleSinSubvencionBuilder{
		metadatos: models.NuevosMetadatosFactura(55),
		factura: &documents.FacturaVentaCombustibleSinSubvencion{
			XMLName:           xml.Name{Local: "facturaElectronicaVentaCombustibleSinSubvencion"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaVentaCombustibleSinSubvencion.xsd",
		},
	}
}

// NewVentaCombustibleSinSubvencionCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewVentaCombustibleSinSubvencionCabeceraBuilder() *ventaCombustibleSinSubvencionCabeceraBuilder {
	return &ventaCombustibleSinSubvencionCabeceraBuilder{
		cabecera: &documents.CabeceraVentaCombustibleSinSubvencion{
			CodigoDocumentoSector: 55, // Sector 55 para Venta de Combustible Sin Subvención
		},
	}
}

// NewVentaCombustibleSinSubvencionDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewVentaCombustibleSinSubvencionDetalleBuilder() *ventaCombustibleSinSubvencionDetalleBuilder {
	return &ventaCombustibleSinSubvencionDetalleBuilder{
		detalle: &documents.DetalleVentaCombustibleSinSubvencion{},
	}
}

type ventaCombustibleSinSubvencionBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaVentaCombustibleSinSubvencion
}

func (b *ventaCombustibleSinSubvencionBuilder) WithCabecera(req VentaCombustibleSinSubvencionCabecera) *ventaCombustibleSinSubvencionBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraVentaCombustibleSinSubvencion](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *ventaCombustibleSinSubvencionBuilder) AddDetalle(req VentaCombustibleSinSubvencionDetalle) *ventaCombustibleSinSubvencionBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleVentaCombustibleSinSubvencion](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *ventaCombustibleSinSubvencionBuilder) WithModalidad(tipo int) *ventaCombustibleSinSubvencionBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaVentaCombustibleSinSubvencion"}
		b.factura.XsiSchemaLocation = "facturaElectronicaVentaCombustibleSinSubvencion.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaVentaCombustibleSinSubvencion"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaVentaCombustibleSinSubvencion.xsd"
	}
	return b
}

func (b *ventaCombustibleSinSubvencionBuilder) Build() VentaCombustibleSinSubvencion {
	return VentaCombustibleSinSubvencion{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type ventaCombustibleSinSubvencionCabeceraBuilder struct {
	cabecera *documents.CabeceraVentaCombustibleSinSubvencion
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) Build() VentaCombustibleSinSubvencionCabecera {
	return VentaCombustibleSinSubvencionCabecera{models.NewRequestWrapper(b.cabecera)}
}

type ventaCombustibleSinSubvencionDetalleBuilder struct {
	detalle *documents.DetalleVentaCombustibleSinSubvencion
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) Build() VentaCombustibleSinSubvencionDetalle {
	return VentaCombustibleSinSubvencionDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNitEmisor(v int64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithRazonSocialEmisor(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithMunicipio(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithTelefono(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNumeroFactura(v int64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCuf(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCufd(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoSucursal(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithDireccion(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoPuntoVenta(v *int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNitEmisorEeSs(v int64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.NitEmisorEeSs = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithFechaEmision(v time.Time) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNombreRazonSocial(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNumeroDocumento(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithComplemento(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoCliente(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoPais(v *int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPais = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPais = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithPlacaVehiculo(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.PlacaVehiculo = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithTipoEnvase(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.TipoEnvase = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.TipoEnvase = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoMetodoPago(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNumeroTarjeta(v *int64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithSubTotalNeto(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithTotalPrecioNeto(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithIva(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithMontoAPagar(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoAutorizacionSC(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoAutorizacionSC = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.CodigoAutorizacionSC = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithObservacion(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.Observacion = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Observacion = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoMoneda(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithTipoCambio(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithMontoTotalMoneda(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithDescuentoAdicional(v *float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoExcepcion(v *int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithLeyenda(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithUsuario(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoInterno(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoDocumentoSector(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithActividadEconomica(v string) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithCodigoProductoSin(v int64) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithCodigoProducto(v string) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithDescripcion(v string) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithCantidad(v float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithUnidadMedida(v int) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithPrecioNetoUnitario(v float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithMontoDescuento(v *float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithSubTotal(v float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}
