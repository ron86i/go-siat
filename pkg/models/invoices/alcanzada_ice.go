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

// AlcanzadaIce representa la estructura completa de una factura Sector 14
// lista para ser procesada.
type AlcanzadaIce struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaAlcanzadaIce]
}

// AlcanzadaIceCabecera representa la sección de cabecera de la factura.
type AlcanzadaIceCabecera struct {
	models.RequestWrapper[documents.CabeceraAlcanzadaIce]
}

// AlcanzadaIceDetalle representa un ítem individual dentro del detalle.
type AlcanzadaIceDetalle struct {
	models.RequestWrapper[documents.DetalleAlcanzadaIce]
}

// NewAlcanzadaIceBuilder inicia el proceso de construcción de la factura.
func NewAlcanzadaIceBuilder() *alcanzadaIceBuilder {
	return &alcanzadaIceBuilder{
		metadatos: models.NuevosMetadatosFactura(14),
		factura: &documents.FacturaAlcanzadaIce{
			XMLName:           xml.Name{Local: "facturaElectronicaAlcanzadaIce"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaAlcanzadaIce.xsd",
		},
	}
}

// NewAlcanzadaIceCabeceraBuilder crea el constructor para la cabecera.
func NewAlcanzadaIceCabeceraBuilder() *alcanzadaIceCabeceraBuilder {
	return &alcanzadaIceCabeceraBuilder{
		cabecera: &documents.CabeceraAlcanzadaIce{
			CodigoDocumentoSector: 14, // Fixed in XSD
		},
	}
}

// NewAlcanzadaIceDetalleBuilder crea el constructor para los detalles.
func NewAlcanzadaIceDetalleBuilder() *alcanzadaIceDetalleBuilder {
	return &alcanzadaIceDetalleBuilder{
		detalle: &documents.DetalleAlcanzadaIce{},
	}
}

type alcanzadaIceBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaAlcanzadaIce
}

func (b *alcanzadaIceBuilder) WithCabecera(req AlcanzadaIceCabecera) *alcanzadaIceBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraAlcanzadaIce](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *alcanzadaIceBuilder) AddDetalle(req AlcanzadaIceDetalle) *alcanzadaIceBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleAlcanzadaIce](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *alcanzadaIceBuilder) WithModalidad(tipo int) *alcanzadaIceBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaAlcanzadaIce"}
		b.factura.XsiSchemaLocation = "facturaElectronicaAlcanzadaIce.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaAlcanzadaIce"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaAlcanzadaIce.xsd"
	}
	return b
}

func (b *alcanzadaIceBuilder) Build() AlcanzadaIce {
	return AlcanzadaIce{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type alcanzadaIceCabeceraBuilder struct {
	cabecera *documents.CabeceraAlcanzadaIce
}

func (b *alcanzadaIceCabeceraBuilder) Build() AlcanzadaIceCabecera {
	return AlcanzadaIceCabecera{models.NewRequestWrapper(b.cabecera)}
}

type alcanzadaIceDetalleBuilder struct {
	detalle *documents.DetalleAlcanzadaIce
}

func (b *alcanzadaIceDetalleBuilder) Build() AlcanzadaIceDetalle {
	return AlcanzadaIceDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *alcanzadaIceCabeceraBuilder) WithNitEmisor(v int64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithRazonSocialEmisor(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMunicipio(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithTelefono(v *string) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithNumeroFactura(v int64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCuf(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCufd(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoSucursal(v int) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithDireccion(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoPuntoVenta(v *int) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithFechaEmision(v time.Time) *alcanzadaIceCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithNombreRazonSocial(v *string) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithNumeroDocumento(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithComplemento(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoCliente(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoMetodoPago(v int) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithNumeroTarjeta(v *int64) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithSubTotalNeto(v float64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithTotalPrecioNeto(v float64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithIva(v float64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMontoAPagar(v float64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithTotalIceEspecifico(v *float64) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.TotalIceEspecifico = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.TotalIceEspecifico = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithTotalIcePorcentual(v *float64) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.TotalIcePorcentual = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.TotalIcePorcentual = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoMoneda(v int) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithTipoCambio(v float64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMontoTotalMoneda(v float64) *alcanzadaIceCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithDescuentoAdicional(v *float64) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoExcepcion(v *int) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithLeyenda(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithUsuario(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoInterno(v string) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoDocumentoSector(v int) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithActividadEconomica(v string) *alcanzadaIceDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithCodigoProductoSin(v int64) *alcanzadaIceDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithCodigoProducto(v string) *alcanzadaIceDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithDescripcion(v string) *alcanzadaIceDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithCantidad(v float64) *alcanzadaIceDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithUnidadMedida(v int) *alcanzadaIceDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithPrecioNetoUnitario(v float64) *alcanzadaIceDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMontoDescuento(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithSubTotal(v float64) *alcanzadaIceDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMarcaIce(v int) *alcanzadaIceDetalleBuilder {
	b.detalle.MarcaIce = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithPrecioNetoVentaIce(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithAlicuotaEspecifica(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithAlicuotaPorcentual(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMontoIceEspecifico(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMontoIcePorcentual(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithCantidadIce(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.CantidadIce = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.CantidadIce = datatype.Nilable[float64]{Value: &value}
	return b
}
