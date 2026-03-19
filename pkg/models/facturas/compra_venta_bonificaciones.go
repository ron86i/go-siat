package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// CompraVentaBonificaciones representa una factura de compra-venta con bonificaciones lista para procesar.
type CompraVentaBonificaciones struct {
	requestWrapper[documentos.FacturaCompraVentaBonificaciones]
}

// CompraVentaBonificacionesCabecera representa la cabecera de una factura de bonificaciones.
type CompraVentaBonificacionesCabecera struct {
	requestWrapper[documentos.CabeceraCompraVentaBonificaciones]
}

// CompraVentaBonificacionesDetalle representa un ítem de detalle de una factura de bonificaciones.
type CompraVentaBonificacionesDetalle struct {
	requestWrapper[documentos.DetalleCompraVentaBonificaciones]
}

// NewCompraVentaBonificacionesBuilder inicia la construcción de una factura de compra-venta con bonificaciones.
func NewCompraVentaBonificacionesBuilder() *compraVentaBonificacionesBuilder {
	return &compraVentaBonificacionesBuilder{
		factura: &documentos.FacturaCompraVentaBonificaciones{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVentaBon"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVentaBon.xsd",
		},
	}
}

// NewCompraVentaBonificacionesCabeceraBuilder crea el constructor de la cabecera de bonificaciones.
func NewCompraVentaBonificacionesCabeceraBuilder() *compraVentaBonificacionesCabeceraBuilder {
	return &compraVentaBonificacionesCabeceraBuilder{
		cabecera: &documentos.CabeceraCompraVentaBonificaciones{},
	}
}

// NewCompraVentaBonificacionesDetalleBuilder crea el constructor del detalle de bonificaciones.
func NewCompraVentaBonificacionesDetalleBuilder() *detalleBonificacionesBuilder {
	return &detalleBonificacionesBuilder{
		detalle: &documentos.DetalleCompraVentaBonificaciones{},
	}
}

// --- Builder Factura Bonificaciones ---

type compraVentaBonificacionesBuilder struct {
	factura *documentos.FacturaCompraVentaBonificaciones
}

// WithModalidad configura los metadatos XML según la modalidad.
func (b *compraVentaBonificacionesBuilder) WithModalidad(tipo int) *compraVentaBonificacionesBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVentaBon"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVentaBon.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVentaBon"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVentaBon.xsd"
	}
	return b
}

func (b *compraVentaBonificacionesBuilder) WithCabecera(req CompraVentaBonificacionesCabecera) *compraVentaBonificacionesBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *compraVentaBonificacionesBuilder) AddDetalle(req CompraVentaBonificacionesDetalle) *compraVentaBonificacionesBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

func (b *compraVentaBonificacionesBuilder) Build() CompraVentaBonificaciones {
	return CompraVentaBonificaciones{requestWrapper[documentos.FacturaCompraVentaBonificaciones]{request: b.factura}}
}

// --- Builder Cabecera Bonificaciones ---

type compraVentaBonificacionesCabeceraBuilder struct {
	cabecera *documentos.CabeceraCompraVentaBonificaciones
}

func (b *compraVentaBonificacionesCabeceraBuilder) WithNitEmisor(v int64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithRazonSocialEmisor(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithMunicipio(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithTelefono(v *string) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithNumeroFactura(v int64) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCuf(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCufd(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoSucursal(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithDireccion(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

// WithCodigoPuntoVenta es nillable en bonificaciones; pasar nil para omitir.
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoPuntoVenta(v *int) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithFechaEmision(v time.Time) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithNombreRazonSocial(v *string) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithNumeroDocumento(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithComplemento(v *string) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoCliente(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoMetodoPago(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithNumeroTarjeta(v *int64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithMontoTotal(v float64) *compraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *compraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoMoneda(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithTipoCambio(v float64) *compraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithMontoTotalMoneda(v float64) *compraVentaBonificacionesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithMontoGiftCard(v *float64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithDescuentoAdicional(v *float64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoExcepcion(v *int64) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCafc(v *string) *compraVentaBonificacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithLeyenda(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithUsuario(v string) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) WithCodigoDocumentoSector(v int) *compraVentaBonificacionesCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}
func (b *compraVentaBonificacionesCabeceraBuilder) Build() CompraVentaBonificacionesCabecera {
	return CompraVentaBonificacionesCabecera{requestWrapper[documentos.CabeceraCompraVentaBonificaciones]{request: b.cabecera}}
}

// --- Builder Detalle Bonificaciones ---

type detalleBonificacionesBuilder struct {
	detalle *documentos.DetalleCompraVentaBonificaciones
}

func (b *detalleBonificacionesBuilder) WithActividadEconomica(v string) *detalleBonificacionesBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

// WithCodigoProductoSin en bonificaciones es integer.
func (b *detalleBonificacionesBuilder) WithCodigoProductoSin(v int64) *detalleBonificacionesBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}
func (b *detalleBonificacionesBuilder) WithCodigoProducto(v string) *detalleBonificacionesBuilder {
	b.detalle.CodigoProducto = v
	return b
}
func (b *detalleBonificacionesBuilder) WithDescripcion(v string) *detalleBonificacionesBuilder {
	b.detalle.Descripcion = v
	return b
}
func (b *detalleBonificacionesBuilder) WithCantidad(v float64) *detalleBonificacionesBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}
func (b *detalleBonificacionesBuilder) WithUnidadMedida(v int) *detalleBonificacionesBuilder {
	b.detalle.UnidadMedida = v
	return b
}
func (b *detalleBonificacionesBuilder) WithPrecioUnitario(v float64) *detalleBonificacionesBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}
func (b *detalleBonificacionesBuilder) WithMontoDescuento(v *float64) *detalleBonificacionesBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *detalleBonificacionesBuilder) WithSubTotal(v float64) *detalleBonificacionesBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}
func (b *detalleBonificacionesBuilder) WithNumeroSerie(v *string) *detalleBonificacionesBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *detalleBonificacionesBuilder) WithNumeroImei(v *string) *detalleBonificacionesBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *detalleBonificacionesBuilder) Build() CompraVentaBonificacionesDetalle {
	return CompraVentaBonificacionesDetalle{requestWrapper[documentos.DetalleCompraVentaBonificaciones]{request: b.detalle}}
}
