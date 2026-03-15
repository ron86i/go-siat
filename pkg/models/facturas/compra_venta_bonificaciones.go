package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// FacturaCompraVentaBonificaciones representa una factura de compra-venta con bonificaciones lista para procesar.
type FacturaCompraVentaBonificaciones interface{}

// FacturaCompraVentaBonCabecera representa la cabecera de una factura de bonificaciones.
type FacturaCompraVentaBonCabecera interface{}

// FacturaCompraVentaBonDetalle representa un ítem de detalle de una factura de bonificaciones.
type FacturaCompraVentaBonDetalle interface{}

// NewFacturaCompraVentaBonBuilder inicia la construcción de una factura de compra-venta con bonificaciones.
func NewFacturaCompraVentaBonBuilder() *FacturaCompraVentaBonBuilder {
	return &FacturaCompraVentaBonBuilder{
		factura: &documentos.FacturaCompraVentaBonificaciones{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVentaBon"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVentaBon.xsd",
		},
	}
}

// NewFacturaCompraVentaBonCabeceraBuilder crea el constructor de la cabecera de bonificaciones.
func NewFacturaCompraVentaBonCabeceraBuilder() *FacturaCompraVentaBonCabeceraBuilder {
	return &FacturaCompraVentaBonCabeceraBuilder{
		cabecera: &documentos.CabeceraCompraVentaBonificaciones{},
	}
}

// NewFacturaCompraVentaBonDetalleBuilder crea el constructor del detalle de bonificaciones.
func NewFacturaCompraVentaBonDetalleBuilder() *DetalleBonificacionesBuilder {
	return &DetalleBonificacionesBuilder{
		detalle: &documentos.DetalleCompraVentaBonificaciones{},
	}
}

// --- Builder Factura Bonificaciones ---

type FacturaCompraVentaBonBuilder struct {
	factura *documentos.FacturaCompraVentaBonificaciones
}

// WithModalidad configura los metadatos XML según la modalidad.
func (b *FacturaCompraVentaBonBuilder) WithModalidad(tipo int) *FacturaCompraVentaBonBuilder {
	switch tipo {
	case ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVentaBon"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVentaBon.xsd"
	case ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVentaBon"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVentaBon.xsd"
	}
	return b
}

func (b *FacturaCompraVentaBonBuilder) WithCabecera(req FacturaCompraVentaBonCabecera) *FacturaCompraVentaBonBuilder {
	if c := getInternalRequest[documentos.CabeceraCompraVentaBonificaciones](req); c != nil {
		b.factura.Cabecera = *c
	}
	return b
}

func (b *FacturaCompraVentaBonBuilder) AddDetalle(req FacturaCompraVentaBonDetalle) *FacturaCompraVentaBonBuilder {
	if d := getInternalRequest[documentos.DetalleCompraVentaBonificaciones](req); d != nil {
		b.factura.Detalle = append(b.factura.Detalle, *d)
	}
	return b
}

func (b *FacturaCompraVentaBonBuilder) Build() FacturaCompraVentaBonificaciones {
	return requestWrapper[documentos.FacturaCompraVentaBonificaciones]{request: b.factura}
}

// --- Builder Cabecera Bonificaciones ---

type FacturaCompraVentaBonCabeceraBuilder struct {
	cabecera *documentos.CabeceraCompraVentaBonificaciones
}

func (b *FacturaCompraVentaBonCabeceraBuilder) WithNitEmisor(v int64) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithRazonSocialEmisor(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithMunicipio(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithTelefono(v *string) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithNumeroFactura(v int64) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCuf(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCufd(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoSucursal(v int) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithDireccion(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

// WithCodigoPuntoVenta es nillable en bonificaciones; pasar nil para omitir.
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoPuntoVenta(v *int) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithFechaEmision(v time.Time) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithNombreRazonSocial(v *string) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithNumeroDocumento(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithComplemento(v *string) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoCliente(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoMetodoPago(v int) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithNumeroTarjeta(v *int64) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithMontoTotal(v float64) *FacturaCompraVentaBonCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *FacturaCompraVentaBonCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoMoneda(v int) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithTipoCambio(v float64) *FacturaCompraVentaBonCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithMontoTotalMoneda(v float64) *FacturaCompraVentaBonCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithMontoGiftCard(v *float64) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithDescuentoAdicional(v *float64) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoExcepcion(v *int64) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCafc(v *string) *FacturaCompraVentaBonCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithLeyenda(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithUsuario(v string) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) WithCodigoDocumentoSector(v int) *FacturaCompraVentaBonCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}
func (b *FacturaCompraVentaBonCabeceraBuilder) Build() FacturaCompraVentaBonCabecera {
	return requestWrapper[documentos.CabeceraCompraVentaBonificaciones]{request: b.cabecera}
}

// --- Builder Detalle Bonificaciones ---

type DetalleBonificacionesBuilder struct {
	detalle *documentos.DetalleCompraVentaBonificaciones
}

func (b *DetalleBonificacionesBuilder) WithActividadEconomica(v string) *DetalleBonificacionesBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

// WithCodigoProductoSin en bonificaciones es integer.
func (b *DetalleBonificacionesBuilder) WithCodigoProductoSin(v int64) *DetalleBonificacionesBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}
func (b *DetalleBonificacionesBuilder) WithCodigoProducto(v string) *DetalleBonificacionesBuilder {
	b.detalle.CodigoProducto = v
	return b
}
func (b *DetalleBonificacionesBuilder) WithDescripcion(v string) *DetalleBonificacionesBuilder {
	b.detalle.Descripcion = v
	return b
}
func (b *DetalleBonificacionesBuilder) WithCantidad(v float64) *DetalleBonificacionesBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}
func (b *DetalleBonificacionesBuilder) WithUnidadMedida(v int) *DetalleBonificacionesBuilder {
	b.detalle.UnidadMedida = v
	return b
}
func (b *DetalleBonificacionesBuilder) WithPrecioUnitario(v float64) *DetalleBonificacionesBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}
func (b *DetalleBonificacionesBuilder) WithMontoDescuento(v *float64) *DetalleBonificacionesBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *DetalleBonificacionesBuilder) WithSubTotal(v float64) *DetalleBonificacionesBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}
func (b *DetalleBonificacionesBuilder) WithNumeroSerie(v *string) *DetalleBonificacionesBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *DetalleBonificacionesBuilder) WithNumeroImei(v *string) *DetalleBonificacionesBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *DetalleBonificacionesBuilder) Build() FacturaCompraVentaBonDetalle {
	return requestWrapper[documentos.DetalleCompraVentaBonificaciones]{request: b.detalle}
}


