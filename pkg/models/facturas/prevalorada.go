package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// Prevalorada representa la estructura completa de una factura Prevalorada lista para ser procesada.
type Prevalorada struct {
	models.RequestWrapper[documentos.FacturaPrevalorada]
}

// PrevaloradaCabecera representa la sección de cabecera de una factura Prevalorada.
type PrevaloradaCabecera struct {
	models.RequestWrapper[documentos.CabeceraPrevalorada]
}

// PrevaloradaDetalle representa el ítem único dentro del detalle de una factura Prevalorada.
type PrevaloradaDetalle struct {
	models.RequestWrapper[documentos.DetallePrevalorada]
}

// NewPrevaloradaBuilder inicia el proceso de construcción de una Factura Prevalorada.
func NewPrevaloradaBuilder() *prevaloradaBuilder {
	return &prevaloradaBuilder{
		factura: &documentos.FacturaPrevalorada{
			XMLName:           xml.Name{Local: "facturaElectronicaPrevalorada"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaPrevalorada.xsd",
		},
	}
}

// NewPrevaloradaCabeceraBuilder crea una instancia del constructor para la cabecera con campos fijos inicializados.
func NewPrevaloradaCabeceraBuilder() *prevaloradaCabeceraBuilder {
	return &prevaloradaCabeceraBuilder{
		cabecera: &documentos.CabeceraPrevalorada{
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
		detalle: &documentos.DetallePrevalorada{},
	}
}

type prevaloradaBuilder struct {
	factura *documentos.FacturaPrevalorada
}

func (b *prevaloradaBuilder) WithCabecera(req PrevaloradaCabecera) *prevaloradaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraPrevalorada](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *prevaloradaBuilder) WithDetalle(req PrevaloradaDetalle) *prevaloradaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetallePrevalorada](req); internal != nil {
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
	return Prevalorada{models.NewRequestWrapper(b.factura)}
}

type prevaloradaCabeceraBuilder struct {
	cabecera *documentos.CabeceraPrevalorada
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

func (b *prevaloradaCabeceraBuilder) WithTelefono(telefono *string) *prevaloradaCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
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
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *prevaloradaCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *prevaloradaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoMetodoPago(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithNumeroTarjeta(v *int64) *prevaloradaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *prevaloradaCabeceraBuilder) WithMontoTotal(v float64) *prevaloradaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *prevaloradaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithCodigoMoneda(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithTipoCambio(v float64) *prevaloradaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithMontoTotalMoneda(v float64) *prevaloradaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
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

func (b *prevaloradaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *prevaloradaCabeceraBuilder) WithNombreRazonSocial(v string) *prevaloradaCabeceraBuilder {
	b.cabecera.NombreRazonSocial = v
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

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *prevaloradaCabeceraBuilder) WithCodigoDocumentoSector(v int) *prevaloradaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *prevaloradaCabeceraBuilder) Build() PrevaloradaCabecera {
	return PrevaloradaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type prevaloradaDetalleBuilder struct {
	detalle *documentos.DetallePrevalorada
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
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithUnidadMedida(v int) *prevaloradaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithPrecioUnitario(v float64) *prevaloradaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *prevaloradaDetalleBuilder) WithMontoDescuento(v *float64) *prevaloradaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *prevaloradaDetalleBuilder) WithSubTotal(v float64) *prevaloradaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *prevaloradaDetalleBuilder) Build() PrevaloradaDetalle {
	return PrevaloradaDetalle{models.NewRequestWrapper(b.detalle)}
}
