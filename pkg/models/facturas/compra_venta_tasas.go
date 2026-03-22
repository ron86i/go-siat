package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// CompraVentaTasas representa una factura de compra-venta con tasas lista para procesar.
type CompraVentaTasas struct {
	requestWrapper[documentos.FacturaCompraVentaTasas]
}

// CompraVentaTasasCabecera representa la cabecera de una factura de tasas.
type CompraVentaTasasCabecera struct {
	requestWrapper[documentos.CabeceraCompraVentaTasas]
}

// CompraVentaTasasDetalle representa un ítem de detalle de una factura de tasas.
type CompraVentaTasasDetalle struct {
	requestWrapper[documentos.DetalleCompraVentaTasas]
}

// NewCompraVentaTasasBuilder inicia la construcción de una factura de compra-venta con tasas.
func NewCompraVentaTasasBuilder() *compraVentaTasasBuilder {
	return &compraVentaTasasBuilder{
		factura: &documentos.FacturaCompraVentaTasas{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVentaTasas"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVentaTasas.xsd",
		},
	}
}

// NewCompraVentaTasasCabeceraBuilder crea el constructor de la cabecera de tasas.
func NewCompraVentaTasasCabeceraBuilder() *compraVentaTasasCabeceraBuilder {
	return &compraVentaTasasCabeceraBuilder{
		cabecera: &documentos.CabeceraCompraVentaTasas{
			CodigoDocumentoSector: 41, // Sector 41 para Compra y Venta de Tasa
		},
	}
}

// NewCompraVentaTasasDetalleBuilder crea el constructor del detalle de tasas.
func NewCompraVentaTasasDetalleBuilder() *detalleTasasBuilder {
	return &detalleTasasBuilder{
		detalle: &documentos.DetalleCompraVentaTasas{},
	}
}

// --- Builder Factura Tasas ---

type compraVentaTasasBuilder struct {
	factura *documentos.FacturaCompraVentaTasas
}

// WithModalidad configura los metadatos XML según la modalidad.
func (b *compraVentaTasasBuilder) WithModalidad(tipo int) *compraVentaTasasBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVentaTasas"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVentaTasas.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVentaTasas"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVentaTasas.xsd"
	}
	return b
}

// WithCabecera asocia la cabecera a la factura.
func (b *compraVentaTasasBuilder) WithCabecera(req CompraVentaTasasCabecera) *compraVentaTasasBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

// AddDetalle añade un ítem de detalle a la factura.
func (b *compraVentaTasasBuilder) AddDetalle(req CompraVentaTasasDetalle) *compraVentaTasasBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

// Build finaliza la construcción y retorna la estructura opaca.
func (b *compraVentaTasasBuilder) Build() CompraVentaTasas {
	return CompraVentaTasas{requestWrapper[documentos.FacturaCompraVentaTasas]{request: b.factura}}
}

// --- Builder Cabecera Tasas ---

type compraVentaTasasCabeceraBuilder struct {
	cabecera *documentos.CabeceraCompraVentaTasas
}

func (b *compraVentaTasasCabeceraBuilder) WithNitEmisor(nitEmisor int64) *compraVentaTasasCabeceraBuilder {
	b.cabecera.NitEmisor = nitEmisor
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithRazonSocialEmisor(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithMunicipio(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithTelefono(v *string) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithNumeroFactura(v int64) *compraVentaTasasCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCuf(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCufd(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCodigoSucursal(v int) *compraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithDireccion(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

// WithCodigoPuntoVenta es nillable en tasas; pasar nil para omitir.
func (b *compraVentaTasasCabeceraBuilder) WithCodigoPuntoVenta(v *int) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithFechaEmision(v time.Time) *compraVentaTasasCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithNombreRazonSocial(v *string) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *compraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithNumeroDocumento(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithComplemento(v *string) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCodigoCliente(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCodigoMetodoPago(v int) *compraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithNumeroTarjeta(v *int64) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithMontoTotal(v float64) *compraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *compraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCodigoMoneda(v int) *compraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithTipoCambio(v float64) *compraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithMontoTotalMoneda(v float64) *compraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithMontoGiftCard(v *float64) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &cp}
	return b
}

// WithMontoTasa configura el monto de tasa (campo exclusivo de esta modalidad).
func (b *compraVentaTasasCabeceraBuilder) WithMontoTasa(v *float64) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoTasa = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.MontoTasa = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithDescuentoAdicional(v *float64) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCodigoExcepcion(v *int64) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCafc(v *string) *compraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithLeyenda(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithUsuario(v string) *compraVentaTasasCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) WithCodigoDocumentoSector(v int) *compraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}
func (b *compraVentaTasasCabeceraBuilder) Build() CompraVentaTasasCabecera {
	return CompraVentaTasasCabecera{requestWrapper[documentos.CabeceraCompraVentaTasas]{request: b.cabecera}}
}

// --- Builder Detalle Tasas ---

type detalleTasasBuilder struct {
	detalle *documentos.DetalleCompraVentaTasas
}

func (b *detalleTasasBuilder) WithActividadEconomica(v string) *detalleTasasBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

// WithCodigoProductoSin en tasas es integer.
func (b *detalleTasasBuilder) WithCodigoProductoSin(v int64) *detalleTasasBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}
func (b *detalleTasasBuilder) WithCodigoProducto(v string) *detalleTasasBuilder {
	b.detalle.CodigoProducto = v
	return b
}
func (b *detalleTasasBuilder) WithDescripcion(v string) *detalleTasasBuilder {
	b.detalle.Descripcion = v
	return b
}
func (b *detalleTasasBuilder) WithCantidad(v float64) *detalleTasasBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}
func (b *detalleTasasBuilder) WithUnidadMedida(v int) *detalleTasasBuilder {
	b.detalle.UnidadMedida = v
	return b
}
func (b *detalleTasasBuilder) WithPrecioUnitario(v float64) *detalleTasasBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}
func (b *detalleTasasBuilder) WithMontoDescuento(v *float64) *detalleTasasBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *detalleTasasBuilder) WithSubTotal(v float64) *detalleTasasBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}
func (b *detalleTasasBuilder) WithNumeroSerie(v *string) *detalleTasasBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *detalleTasasBuilder) WithNumeroImei(v *string) *detalleTasasBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *detalleTasasBuilder) Build() CompraVentaTasasDetalle {
	return CompraVentaTasasDetalle{requestWrapper[documentos.DetalleCompraVentaTasas]{request: b.detalle}}
}
