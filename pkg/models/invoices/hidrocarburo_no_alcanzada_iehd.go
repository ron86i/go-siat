package invoices

import (
	"encoding/xml"
	"time"

	"github.com/shopspring/decimal"

	siat "github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v2/internal/core/domain/documents"
)

// HidrocarburoNoAlcanzadaIehd representa una factura de Hidrocarburos No Alcanzada por el IEHD
// (Sector 38) lista para ser procesada.
type HidrocarburoNoAlcanzadaIehd struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaHidrocarburoNoAlcanzadaIehd]
}

// HidrocarburoNoAlcanzadaIehdCabecera representa la sección de cabecera de la factura.
type HidrocarburoNoAlcanzadaIehdCabecera struct {
	models.RequestWrapper[documents.CabeceraHidrocarburoNoAlcanzadaIehd]
}

// HidrocarburoNoAlcanzadaIehdDetalle representa un ítem individual dentro del detalle.
type HidrocarburoNoAlcanzadaIehdDetalle struct {
	models.RequestWrapper[documents.DetalleHidrocarburoNoAlcanzadaIehd]
}

// NewHidrocarburoNoAlcanzadaIehdBuilder inicia la construcción de una Factura de
// Hidrocarburos No Alcanzada por el IEHD (Sector 38).
func NewHidrocarburoNoAlcanzadaIehdBuilder() *hidrocarburoNoAlcanzadaIehdBuilder {
	return &hidrocarburoNoAlcanzadaIehdBuilder{
		metadatos: models.NuevosMetadatosFactura(38),
		factura: &documents.FacturaHidrocarburoNoAlcanzadaIehd{
			XMLName:           xml.Name{Local: "facturaElectronicaHidrocarburoNoAlcanzadaIehd"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaHidrocarburoNoAlcanzadaIehd.xsd",
		},
	}
}

// NewHidrocarburoNoAlcanzadaIehdCabeceraBuilder crea el constructor para la cabecera.
func NewHidrocarburoNoAlcanzadaIehdCabeceraBuilder() *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	return &hidrocarburoNoAlcanzadaIehdCabeceraBuilder{
		cabecera: &documents.CabeceraHidrocarburoNoAlcanzadaIehd{
			CodigoDocumentoSector: 38, // Sector 38 para Hidrocarburos No Alcanzada IEHD
		},
	}
}

// NewHidrocarburoNoAlcanzadaIehdDetalleBuilder crea el constructor para los ítems de detalle.
func NewHidrocarburoNoAlcanzadaIehdDetalleBuilder() *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	return &hidrocarburoNoAlcanzadaIehdDetalleBuilder{
		detalle: &documents.DetalleHidrocarburoNoAlcanzadaIehd{},
	}
}

type hidrocarburoNoAlcanzadaIehdBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaHidrocarburoNoAlcanzadaIehd
}

func (b *hidrocarburoNoAlcanzadaIehdBuilder) WithCabecera(req HidrocarburoNoAlcanzadaIehdCabecera) *hidrocarburoNoAlcanzadaIehdBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraHidrocarburoNoAlcanzadaIehd](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdBuilder) AddDetalle(req HidrocarburoNoAlcanzadaIehdDetalle) *hidrocarburoNoAlcanzadaIehdBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleHidrocarburoNoAlcanzadaIehd](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdBuilder) WithModalidad(tipo int) *hidrocarburoNoAlcanzadaIehdBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaHidrocarburoNoAlcanzadaIehd"}
		b.factura.XsiSchemaLocation = "facturaElectronicaHidrocarburoNoAlcanzadaIehd.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaHidrocarburoNoAlcanzadaIehd"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaHidrocarburoNoAlcanzadaIehd.xsd"
	}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdBuilder) Build() HidrocarburoNoAlcanzadaIehd {
	return HidrocarburoNoAlcanzadaIehd{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type hidrocarburoNoAlcanzadaIehdCabeceraBuilder struct {
	cabecera *documents.CabeceraHidrocarburoNoAlcanzadaIehd
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithNitEmisor(v int64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithRazonSocialEmisor(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithMunicipio(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithTelefono(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithNumeroFactura(v int64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCuf(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCufd(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoSucursal(v int) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithDireccion(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoPuntoVenta(v *int) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithNombreRazonSocial(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithNumeroDocumento(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithComplemento(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoCliente(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCiudad(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Ciudad = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithNombrePropietario(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithNombreRepresentanteLegal(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCondicionPago(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CondicionPago = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CondicionPago = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithPeriodoEntrega(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoMetodoPago(v int) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithNumeroTarjeta(v *int64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithMontoTotal(v float64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	v = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	b.cabecera.MontoTotal = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	v = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoMoneda(v int) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithTipoCambio(v float64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	v = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	b.cabecera.TipoCambio = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithMontoTotalMoneda(v float64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	v = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithDescuentoAdicional(v *float64) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoExcepcion(v *int) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCafc(v *string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithLeyenda(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithUsuario(v string) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) WithCodigoDocumentoSector(v int) *hidrocarburoNoAlcanzadaIehdCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdCabeceraBuilder) Build() HidrocarburoNoAlcanzadaIehdCabecera {
	return HidrocarburoNoAlcanzadaIehdCabecera{models.NewRequestWrapper(b.cabecera)}
}

type hidrocarburoNoAlcanzadaIehdDetalleBuilder struct {
	detalle *documents.DetalleHidrocarburoNoAlcanzadaIehd
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithActividadEconomica(v string) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithCodigoProductoSin(v int64) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithCodigoProducto(v string) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithDescripcion(v string) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithCantidad(v float64) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	v = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	b.detalle.Cantidad = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithUnidadMedida(v int) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithPrecioUnitario(v float64) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	v = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	b.detalle.PrecioUnitario = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithMontoDescuento(v *float64) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) WithSubTotal(v float64) *hidrocarburoNoAlcanzadaIehdDetalleBuilder {
	v = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	b.detalle.SubTotal = v
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdDetalleBuilder) Build() HidrocarburoNoAlcanzadaIehdDetalle {
	return HidrocarburoNoAlcanzadaIehdDetalle{models.NewRequestWrapper(b.detalle)}
}
