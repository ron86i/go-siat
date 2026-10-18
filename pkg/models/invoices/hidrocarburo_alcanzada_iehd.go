package invoices

import (
	"encoding/xml"
	"time"

	"github.com/shopspring/decimal"

	siat "github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
)

// HidrocarburoAlcanzadaIehd representa una factura de Hidrocarburos Alcanzada por el IEHD
// (Sector 19) lista para ser procesada.
type HidrocarburoAlcanzadaIehd struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaHidrocarburoAlcanzadaIehd]
}

// HidrocarburoAlcanzadaIehdCabecera representa la sección de cabecera de la factura.
type HidrocarburoAlcanzadaIehdCabecera struct {
	models.RequestWrapper[documents.CabeceraHidrocarburoAlcanzadaIehd]
}

// HidrocarburoAlcanzadaIehdDetalle representa un ítem individual dentro del detalle.
type HidrocarburoAlcanzadaIehdDetalle struct {
	models.RequestWrapper[documents.DetalleHidrocarburoAlcanzadaIehd]
}

// NewHidrocarburoAlcanzadaIehdBuilder inicia la construcción de una Factura de
// Hidrocarburos Alcanzada por el IEHD (Sector 19).
func NewHidrocarburoAlcanzadaIehdBuilder() *hidrocarburoAlcanzadaIehdBuilder {
	return &hidrocarburoAlcanzadaIehdBuilder{
		metadatos: models.NuevosMetadatosFactura(19),
		factura: &documents.FacturaHidrocarburoAlcanzadaIehd{
			XMLName:           xml.Name{Local: "facturaElectronicaHidrocarburoAlcanzadaIehd"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaHidrocarburoAlcanzadaIehd.xsd",
		},
	}
}

// NewHidrocarburoAlcanzadaIehdCabeceraBuilder crea el constructor para la cabecera.
func NewHidrocarburoAlcanzadaIehdCabeceraBuilder() *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	return &hidrocarburoAlcanzadaIehdCabeceraBuilder{
		cabecera: &documents.CabeceraHidrocarburoAlcanzadaIehd{
			CodigoDocumentoSector: 19, // Sector 19 para Hidrocarburos Alcanzada IEHD
		},
	}
}

// NewHidrocarburoAlcanzadaIehdDetalleBuilder crea el constructor para los ítems de detalle.
func NewHidrocarburoAlcanzadaIehdDetalleBuilder() *hidrocarburoAlcanzadaIehdDetalleBuilder {
	return &hidrocarburoAlcanzadaIehdDetalleBuilder{
		detalle: &documents.DetalleHidrocarburoAlcanzadaIehd{},
	}
}

type hidrocarburoAlcanzadaIehdBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaHidrocarburoAlcanzadaIehd
}

func (b *hidrocarburoAlcanzadaIehdBuilder) WithCabecera(req HidrocarburoAlcanzadaIehdCabecera) *hidrocarburoAlcanzadaIehdBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraHidrocarburoAlcanzadaIehd](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *hidrocarburoAlcanzadaIehdBuilder) AddDetalle(req HidrocarburoAlcanzadaIehdDetalle) *hidrocarburoAlcanzadaIehdBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleHidrocarburoAlcanzadaIehd](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *hidrocarburoAlcanzadaIehdBuilder) WithModalidad(tipo int) *hidrocarburoAlcanzadaIehdBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaHidrocarburoAlcanzadaIehd"}
		b.factura.XsiSchemaLocation = "facturaElectronicaHidrocarburoAlcanzadaIehd.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaHidrocarburoAlcanzadaIehd"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaHidrocarburoAlcanzadaIehd.xsd"
	}
	return b
}

func (b *hidrocarburoAlcanzadaIehdBuilder) Build() HidrocarburoAlcanzadaIehd {
	return HidrocarburoAlcanzadaIehd{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type hidrocarburoAlcanzadaIehdCabeceraBuilder struct {
	cabecera *documents.CabeceraHidrocarburoAlcanzadaIehd
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) Build() HidrocarburoAlcanzadaIehdCabecera {
	return HidrocarburoAlcanzadaIehdCabecera{models.NewRequestWrapper(b.cabecera)}
}

type hidrocarburoAlcanzadaIehdDetalleBuilder struct {
	detalle *documents.DetalleHidrocarburoAlcanzadaIehd
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) Build() HidrocarburoAlcanzadaIehdDetalle {
	return HidrocarburoAlcanzadaIehdDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithNitEmisor(v int64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithRazonSocialEmisor(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithMunicipio(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithTelefono(v *string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithNumeroFactura(v int64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCuf(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCufd(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoSucursal(v int) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithDireccion(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoPuntoVenta(v *int) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithFechaEmision(v time.Time) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithNombreRazonSocial(v *string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithNumeroDocumento(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithComplemento(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoCliente(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCiudad(v *string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Ciudad = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Ciudad = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithNombrePropietario(v *string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombrePropietario = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithNombreRepresentanteLegal(v *string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCondicionPago(v *string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CondicionPago = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.CondicionPago = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithPeriodoEntrega(v *string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.PeriodoEntrega = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoMetodoPago(v int) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithNumeroTarjeta(v *int64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithMontoIehd(v *float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoIehd = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoIehd = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithSubTotalNeto(v float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithTotalPrecioNeto(v float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithIva(v float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithIvad(v float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Ivad = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithMontoAPagar(v float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoMoneda(v int) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithTipoCambio(v float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithMontoTotalMoneda(v float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithDescuentoAdicional(v *float64) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoExcepcion(v *int) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithLeyenda(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithUsuario(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoInterno(v string) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdCabeceraBuilder) WithCodigoDocumentoSector(v int) *hidrocarburoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithActividadEconomica(v string) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithCodigoProductoSin(v int64) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithCodigoProducto(v string) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithDescripcion(v string) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithCantidad(v float64) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(10).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithUnidadMedida(v int) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithPrecioNetoUnitario(v float64) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(10).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithMontoDescuento(v *float64) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(10).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithSubTotal(v float64) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(10).InexactFloat64()
	return b
}

func (b *hidrocarburoAlcanzadaIehdDetalleBuilder) WithPorcentajeIehd(v *float64) *hidrocarburoAlcanzadaIehdDetalleBuilder {
	if v == nil {
		b.detalle.PorcentajeIehd = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(10).InexactFloat64()
	b.detalle.PorcentajeIehd = datatype.Nilable[float64]{Value: &value}
	return b
}
