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

// Prevalorada representa la estructura completa de una factura Prevalorada lista para ser procesada.
type Prevalorada struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaPrevalorada]
}

// PrevaloradaCabecera representa la sección de cabecera de una factura Prevalorada.
type PrevaloradaCabecera struct {
	models.RequestWrapper[documents.CabeceraPrevalorada]
}

// PrevaloradaDetalle representa el ítem único dentro del detalle de una factura Prevalorada.
type PrevaloradaDetalle struct {
	models.RequestWrapper[documents.DetallePrevalorada]
}

// NewPrevaloradaBuilder inicia el proceso de construcción de una Factura Prevalorada.
func NewPrevaloradaBuilder() *prevaloradaBuilder {
	return &prevaloradaBuilder{
		metadatos: models.NuevosMetadatosFactura(23),
		factura: &documents.FacturaPrevalorada{
			XMLName:           xml.Name{Local: "facturaElectronicaPrevalorada"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaPrevalorada.xsd",
		},
	}
}

// NewPrevaloradaCabeceraBuilder crea una instancia del constructor para la cabecera con campos fijos inicializados.
func NewPrevaloradaCabeceraBuilder() *prevaloradaCabeceraBuilder {
	return &prevaloradaCabeceraBuilder{
		cabecera: &documents.CabeceraPrevalorada{
			NombreRazonSocial:            "S/N",
			CodigoTipoDocumentoIdentidad: 5,
			NumeroDocumento:              0,
			CodigoCliente:                "N/A",
			CodigoDocumentoSector:        23,
		},
	}
}

// NewPrevaloradaDetalleBuilder crea una instancia del constructor para el ítem de detalle.
func NewPrevaloradaDetalleBuilder() *prevaloradaDetalleBuilder {
	return &prevaloradaDetalleBuilder{
		detalle: &documents.DetallePrevalorada{},
	}
}

type prevaloradaBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaPrevalorada
}

func (b *prevaloradaBuilder) WithCabecera(req PrevaloradaCabecera) *prevaloradaBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraPrevalorada](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *prevaloradaBuilder) WithDetalle(req PrevaloradaDetalle) *prevaloradaBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetallePrevalorada](req); internal != nil {
		b.factura.Detalle = *internal
	}
	return b
}

func (b *prevaloradaBuilder) WithModalidad(tipo int) *prevaloradaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaPrevalorada"}
		b.factura.XsiSchemaLocation = "facturaElectronicaPrevalorada.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaPrevalorada"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaPrevalorada.xsd"
	}
	return b
}

func (b *prevaloradaBuilder) Build() Prevalorada {
	return Prevalorada{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type prevaloradaCabeceraBuilder struct {
	cabecera *documents.CabeceraPrevalorada
}

func (b *prevaloradaCabeceraBuilder) Build() PrevaloradaCabecera {
	return PrevaloradaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type prevaloradaDetalleBuilder struct {
	detalle *documents.DetallePrevalorada
}

func (b *prevaloradaDetalleBuilder) Build() PrevaloradaDetalle {
	return PrevaloradaDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *prevaloradaCabeceraBuilder) WithNitEmisor(v int64) *prevaloradaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithRazonSocialEmisor(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithMunicipio(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithTelefono(v *string) *prevaloradaCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *prevaloradaCabeceraBuilder) WithNumeroFactura(v int64) *prevaloradaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCuf(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCufd(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoSucursal(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithDireccion(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *prevaloradaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *prevaloradaCabeceraBuilder) WithFechaEmision(v time.Time) *prevaloradaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *prevaloradaCabeceraBuilder) WithNombreRazonSocial(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.NombreRazonSocial = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithNumeroDocumento(v int64) *prevaloradaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoCliente(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoMetodoPago(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithNumeroTarjeta(v *int64) *prevaloradaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *prevaloradaCabeceraBuilder) WithSubTotalNeto(v float64) *prevaloradaCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaCabeceraBuilder) WithTotalPrecioNeto(v float64) *prevaloradaCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaCabeceraBuilder) WithIva(v float64) *prevaloradaCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaCabeceraBuilder) WithMontoAPagar(v float64) *prevaloradaCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoMoneda(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithTipoCambio(v float64) *prevaloradaCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaCabeceraBuilder) WithMontoTotalMoneda(v float64) *prevaloradaCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaCabeceraBuilder) WithLeyenda(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithUsuario(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoInterno(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoDocumentoSector(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithActividadEconomica(v string) *prevaloradaDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithCodigoProductoSin(v int64) *prevaloradaDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithCodigoProducto(v string) *prevaloradaDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithDescripcion(v string) *prevaloradaDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithCantidad(v float64) *prevaloradaDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaDetalleBuilder) WithUnidadMedida(v int) *prevaloradaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithPrecioNetoUnitario(v float64) *prevaloradaDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *prevaloradaDetalleBuilder) WithMontoDescuento(v *float64) *prevaloradaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *prevaloradaDetalleBuilder) WithSubTotal(v float64) *prevaloradaDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}
