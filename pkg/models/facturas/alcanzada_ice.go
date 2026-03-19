package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// AlcanzadaIce representa la estructura completa de una factura Sector 14
// lista para ser procesada.
type AlcanzadaIce struct {
	requestWrapper[documentos.FacturaAlcanzadaIce]
}

// AlcanzadaIceCabecera representa la sección de cabecera de la factura.
type AlcanzadaIceCabecera struct {
	requestWrapper[documentos.CabeceraAlcanzadaIce]
}

// AlcanzadaIceDetalle representa un ítem individual dentro del detalle.
type AlcanzadaIceDetalle struct {
	requestWrapper[documentos.DetalleAlcanzadaIce]
}

// NewAlcanzadaIceBuilder inicia el proceso de construcción de la factura.
func NewAlcanzadaIceBuilder() *alcanzadaIceBuilder {
	return &alcanzadaIceBuilder{
		factura: &documentos.FacturaAlcanzadaIce{
			XMLName:           xml.Name{Local: "facturaElectronicaAlcanzadaIce"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaAlcanzadaIce.xsd",
		},
	}
}

// NewAlcanzadaIceCabeceraBuilder crea el constructor para la cabecera.
func NewAlcanzadaIceCabeceraBuilder() *alcanzadaIceCabeceraBuilder {
	return &alcanzadaIceCabeceraBuilder{
		cabecera: &documentos.CabeceraAlcanzadaIce{
			CodigoDocumentoSector: 14, // Fixed in XSD
		},
	}
}

// NewAlcanzadaIceDetalleBuilder crea el constructor para los detalles.
func NewAlcanzadaIceDetalleBuilder() *alcanzadaIceDetalleBuilder {
	return &alcanzadaIceDetalleBuilder{
		detalle: &documentos.DetalleAlcanzadaIce{},
	}
}

type alcanzadaIceBuilder struct {
	factura *documentos.FacturaAlcanzadaIce
}

func (b *alcanzadaIceBuilder) WithCabecera(req AlcanzadaIceCabecera) *alcanzadaIceBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *alcanzadaIceBuilder) AddDetalle(req AlcanzadaIceDetalle) *alcanzadaIceBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
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
	return AlcanzadaIce{requestWrapper[documentos.FacturaAlcanzadaIce]{request: b.factura}}
}

type alcanzadaIceCabeceraBuilder struct {
	cabecera *documentos.CabeceraAlcanzadaIce
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
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
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
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithFechaEmision(v time.Time) *alcanzadaIceCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.TimeSiat(v)
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithNombreRazonSocial(v *string) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
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

func (b *alcanzadaIceCabeceraBuilder) WithComplemento(v *string) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
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
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMontoTotal(v float64) *alcanzadaIceCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMontoIceEspecifico(v *float64) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoIceEspecifico = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.MontoIceEspecifico = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMontoIcePorcentual(v *float64) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoIcePorcentual = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.MontoIcePorcentual = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *alcanzadaIceCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoMoneda(v int) *alcanzadaIceCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithTipoCambio(v float64) *alcanzadaIceCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithMontoTotalMoneda(v float64) *alcanzadaIceCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithDescuentoAdicional(v *float64) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCodigoExcepcion(v *int) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *alcanzadaIceCabeceraBuilder) WithCafc(v *string) *alcanzadaIceCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
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

func (b *alcanzadaIceCabeceraBuilder) Build() AlcanzadaIceCabecera {
	return AlcanzadaIceCabecera{requestWrapper[documentos.CabeceraAlcanzadaIce]{request: b.cabecera}}
}

type alcanzadaIceDetalleBuilder struct {
	detalle *documentos.DetalleAlcanzadaIce
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
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithUnidadMedida(v int) *alcanzadaIceDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithPrecioUnitario(v float64) *alcanzadaIceDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMontoDescuento(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithSubTotal(v float64) *alcanzadaIceDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMarcaIce(v int) *alcanzadaIceDetalleBuilder {
	b.detalle.MarcaIce = v
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithAlicuotaIva(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.AlicuotaIva = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.AlicuotaIva = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithPrecioNetoVentaIce(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.PrecioNetoVentaIce = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithAlicuotaEspecifica(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.AlicuotaEspecifica = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithAlicuotaPorcentual(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.AlicuotaPorcentual = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMontoIceEspecifico(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.MontoIceEspecifico = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithMontoIcePorcentual(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.MontoIcePorcentual = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) WithCantidadIce(v *float64) *alcanzadaIceDetalleBuilder {
	if v == nil {
		b.detalle.CantidadIce = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.CantidadIce = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *alcanzadaIceDetalleBuilder) Build() AlcanzadaIceDetalle {
	return AlcanzadaIceDetalle{requestWrapper[documentos.DetalleAlcanzadaIce]{request: b.detalle}}
}
