package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// FacturaCompraVentaBonificaciones representa una factura de compra-venta con bonificaciones lista para procesar.
type FacturaCompraVentaBonificaciones struct {
	requestWrapper[documentos.FacturaCompraVentaBonificaciones]
}

// FacturaCompraVentaBonificacionesCabecera representa la cabecera de una factura de bonificaciones.
type FacturaCompraVentaBonificacionesCabecera struct {
	requestWrapper[documentos.CabeceraCompraVentaBonificaciones]
}

// FacturaCompraVentaBonificacionesDetalle representa un ítem de detalle de una factura de bonificaciones.
type FacturaCompraVentaBonificacionesDetalle struct {
	requestWrapper[documentos.DetalleCompraVentaBonificaciones]
}

// NewFacturaCompraVentaBonificacionesBuilder inicia la construcción de una factura de compra-venta con bonificaciones.
func NewFacturaCompraVentaBonificacionesBuilder() *FacturaCompraVentaBonificacionesBuilder {
	return &FacturaCompraVentaBonificacionesBuilder{
		factura: &documentos.FacturaCompraVentaBonificaciones{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVentaBon"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVentaBon.xsd",
		},
	}
}

// NewFacturaCompraVentaBonificacionesCabeceraBuilder crea el constructor de la cabecera de bonificaciones.
func NewFacturaCompraVentaBonificacionesCabeceraBuilder() *FacturaCompraVentaBonificacionesCabeceraBuilder {
	return &FacturaCompraVentaBonificacionesCabeceraBuilder{
		cabecera: &documentos.CabeceraCompraVentaBonificaciones{},
	}
}

// NewFacturaCompraVentaBonificacionesDetalleBuilder crea el constructor del detalle de bonificaciones.
func NewFacturaCompraVentaBonificacionesDetalleBuilder() *DetalleBonificacionesBuilder {
	return &DetalleBonificacionesBuilder{
		detalle: &documentos.DetalleCompraVentaBonificaciones{},
	}
}

// --- Builder Factura Bonificaciones ---

type FacturaCompraVentaBonificacionesBuilder struct {
	factura *documentos.FacturaCompraVentaBonificaciones
}

// WithModalidad configura los metadatos XML según la modalidad.
func (b *FacturaCompraVentaBonificacionesBuilder) WithModalidad(tipo int) *FacturaCompraVentaBonificacionesBuilder {
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

func (b *FacturaCompraVentaBonificacionesBuilder) WithCabecera(req FacturaCompraVentaBonificacionesCabecera) *FacturaCompraVentaBonificacionesBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *FacturaCompraVentaBonificacionesBuilder) AddDetalle(req FacturaCompraVentaBonificacionesDetalle) *FacturaCompraVentaBonificacionesBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}



func (b *FacturaCompraVentaBonificacionesBuilder) Build() FacturaCompraVentaBonificaciones {
	return FacturaCompraVentaBonificaciones{requestWrapper[documentos.FacturaCompraVentaBonificaciones]{request: b.factura}}
}

// --- Builder Cabecera Bonificaciones ---

type FacturaCompraVentaBonificacionesCabeceraBuilder struct {
	cabecera *documentos.CabeceraCompraVentaBonificaciones
}

func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithNitEmisor(v int64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithRazonSocialEmisor(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithMunicipio(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithTelefono(v *string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithNumeroFactura(v int64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCuf(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCufd(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoSucursal(v int) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithDireccion(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

// WithCodigoPuntoVenta es nillable en bonificaciones; pasar nil para omitir.
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoPuntoVenta(v *int) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithFechaEmision(v time.Time) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithNombreRazonSocial(v *string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithNumeroDocumento(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithComplemento(v *string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoCliente(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoMetodoPago(v int) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithNumeroTarjeta(v *int64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithMontoTotal(v float64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoMoneda(v int) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithTipoCambio(v float64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithMontoTotalMoneda(v float64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithMontoGiftCard(v *float64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithDescuentoAdicional(v *float64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoExcepcion(v *int64) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCafc(v *string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithLeyenda(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithUsuario(v string) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) WithCodigoDocumentoSector(v int) *FacturaCompraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}
func (b *FacturaCompraVentaBonificacionesCabeceraBuilder) Build() FacturaCompraVentaBonificacionesCabecera {
	return FacturaCompraVentaBonificacionesCabecera{requestWrapper[documentos.CabeceraCompraVentaBonificaciones]{request: b.cabecera}}
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
func (b *DetalleBonificacionesBuilder) Build() FacturaCompraVentaBonificacionesDetalle {
	return FacturaCompraVentaBonificacionesDetalle{requestWrapper[documentos.DetalleCompraVentaBonificaciones]{request: b.detalle}}
}
