package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// FacturaCompraVentaTasas representa una factura de compra-venta con tasas lista para procesar.
type FacturaCompraVentaTasas interface{}

// FacturaCompraVentaTasasCabecera representa la cabecera de una factura de tasas.
type FacturaCompraVentaTasasCabecera interface{}

// FacturaCompraVentaTasasDetalle representa un ítem de detalle de una factura de tasas.
type FacturaCompraVentaTasasDetalle interface{}

// NewFacturaCompraVentaTasasBuilder inicia la construcción de una factura de compra-venta con tasas.
func NewFacturaCompraVentaTasasBuilder() *FacturaCompraVentaTasasBuilder {
	return &FacturaCompraVentaTasasBuilder{
		factura: &documentos.FacturaCompraVentaTasas{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVentaTasas"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVentaTasas.xsd",
		},
	}
}

// NewFacturaCompraVentaTasasCabeceraBuilder crea el constructor de la cabecera de tasas.
func NewFacturaCompraVentaTasasCabeceraBuilder() *FacturaCompraVentaTasasCabeceraBuilder {
	return &FacturaCompraVentaTasasCabeceraBuilder{
		cabecera: &documentos.CabeceraCompraVentaTasas{},
	}
}

// NewFacturaCompraVentaTasasDetalleBuilder crea el constructor del detalle de tasas.
func NewFacturaCompraVentaTasasDetalleBuilder() *DetalleTasasBuilder {
	return &DetalleTasasBuilder{
		detalle: &documentos.DetalleCompraVentaTasas{},
	}
}

// --- Builder Factura Tasas ---

type FacturaCompraVentaTasasBuilder struct {
	factura *documentos.FacturaCompraVentaTasas
}

// WithModalidad configura los metadatos XML según la modalidad.
func (b *FacturaCompraVentaTasasBuilder) WithModalidad(tipo int) *FacturaCompraVentaTasasBuilder {
	switch tipo {
	case ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVentaTasas"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVentaTasas.xsd"
	case ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVentaTasas"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVentaTasas.xsd"
	}
	return b
}

// WithCabecera asocia la cabecera a la factura.
func (b *FacturaCompraVentaTasasBuilder) WithCabecera(req FacturaCompraVentaTasasCabecera) *FacturaCompraVentaTasasBuilder {
	if c := getInternalRequest[documentos.CabeceraCompraVentaTasas](req); c != nil {
		b.factura.Cabecera = *c
	}
	return b
}

// AddDetalle añade un ítem de detalle a la factura.
func (b *FacturaCompraVentaTasasBuilder) AddDetalle(req FacturaCompraVentaTasasDetalle) *FacturaCompraVentaTasasBuilder {
	if d := getInternalRequest[documentos.DetalleCompraVentaTasas](req); d != nil {
		b.factura.Detalle = append(b.factura.Detalle, *d)
	}
	return b
}

// Build finaliza la construcción y retorna la interfaz opaca.
func (b *FacturaCompraVentaTasasBuilder) Build() FacturaCompraVentaTasas {
	return requestWrapper[documentos.FacturaCompraVentaTasas]{request: b.factura}
}

// --- Builder Cabecera Tasas ---

type FacturaCompraVentaTasasCabeceraBuilder struct {
	cabecera *documentos.CabeceraCompraVentaTasas
}

func (b *FacturaCompraVentaTasasCabeceraBuilder) WithNitEmisor(nitEmisor int64) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.NitEmisor = nitEmisor
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithRazonSocialEmisor(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithMunicipio(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithTelefono(v *string) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithNumeroFactura(v int64) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCuf(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCufd(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoSucursal(v int) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithDireccion(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

// WithCodigoPuntoVenta es nillable en tasas; pasar nil para omitir.
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoPuntoVenta(v *int) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithFechaEmision(v time.Time) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithNombreRazonSocial(v *string) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithNumeroDocumento(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithComplemento(v *string) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoCliente(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoMetodoPago(v int) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithNumeroTarjeta(v *int64) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithMontoTotal(v float64) *FacturaCompraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *FacturaCompraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoMoneda(v int) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithTipoCambio(v float64) *FacturaCompraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithMontoTotalMoneda(v float64) *FacturaCompraVentaTasasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithMontoGiftCard(v *float64) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &cp}
	return b
}

// WithMontoTasa configura el monto de tasa (campo exclusivo de esta modalidad).
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithMontoTasa(v *float64) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoTasa = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.MontoTasa = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithDescuentoAdicional(v *float64) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoExcepcion(v *int64) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{}
		return b
	}
	cp := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCafc(v *string) *FacturaCompraVentaTasasCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithLeyenda(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithUsuario(v string) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) WithCodigoDocumentoSector(v int) *FacturaCompraVentaTasasCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}
func (b *FacturaCompraVentaTasasCabeceraBuilder) Build() FacturaCompraVentaTasasCabecera {
	return requestWrapper[documentos.CabeceraCompraVentaTasas]{request: b.cabecera}
}

// --- Builder Detalle Tasas ---

type DetalleTasasBuilder struct {
	detalle *documentos.DetalleCompraVentaTasas
}

func (b *DetalleTasasBuilder) WithActividadEconomica(v string) *DetalleTasasBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

// WithCodigoProductoSin en tasas es integer.
func (b *DetalleTasasBuilder) WithCodigoProductoSin(v int64) *DetalleTasasBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}
func (b *DetalleTasasBuilder) WithCodigoProducto(v string) *DetalleTasasBuilder {
	b.detalle.CodigoProducto = v
	return b
}
func (b *DetalleTasasBuilder) WithDescripcion(v string) *DetalleTasasBuilder {
	b.detalle.Descripcion = v
	return b
}
func (b *DetalleTasasBuilder) WithCantidad(v float64) *DetalleTasasBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}
func (b *DetalleTasasBuilder) WithUnidadMedida(v int) *DetalleTasasBuilder {
	b.detalle.UnidadMedida = v
	return b
}
func (b *DetalleTasasBuilder) WithPrecioUnitario(v float64) *DetalleTasasBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}
func (b *DetalleTasasBuilder) WithMontoDescuento(v *float64) *DetalleTasasBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	cp, _ := strconv.ParseFloat(strconv.FormatFloat(*v, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &cp}
	return b
}
func (b *DetalleTasasBuilder) WithSubTotal(v float64) *DetalleTasasBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}
func (b *DetalleTasasBuilder) WithNumeroSerie(v *string) *DetalleTasasBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *DetalleTasasBuilder) WithNumeroImei(v *string) *DetalleTasasBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{}
		return b
	}
	cp := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &cp}
	return b
}
func (b *DetalleTasasBuilder) Build() FacturaCompraVentaTasasDetalle {
	return requestWrapper[documentos.DetalleCompraVentaTasas]{request: b.detalle}
}


