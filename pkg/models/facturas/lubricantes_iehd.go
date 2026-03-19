package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// LubricantesIehd representa la estructura completa de una factura Sector 53
// lista para ser procesada.
type LubricantesIehd struct {
	requestWrapper[documentos.FacturaLubricantesIehd]
}

// LubricantesIehdCabecera representa la sección de cabecera de la factura.
type LubricantesIehdCabecera struct {
	requestWrapper[documentos.CabeceraLubricantesIehd]
}

// LubricantesIehdDetalle representa un ítem individual dentro del detalle.
type LubricantesIehdDetalle struct {
	requestWrapper[documentos.DetalleLubricantesIehd]
}

// NewLubricantesIehd inicia el proceso de construcción de la factura.
func NewLubricantesIehdBuilder() *lubricantesIehdBuilder {
	return &lubricantesIehdBuilder{
		factura: &documentos.FacturaLubricantesIehd{
			XMLName:           xml.Name{Local: "facturaElectronicaImportacionComercializacionLubricantesIEHD"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaImportacionComercializacionLubricantesIEHD.xsd",
		},
	}
}

// NewLubricantesIehdCabecera crea el constructor para la cabecera.
func NewLubricantesIehdCabeceraBuilder() *lubricantesIehdCabeceraBuilder {
	return &lubricantesIehdCabeceraBuilder{
		cabecera: &documentos.CabeceraLubricantesIehd{
			CodigoDocumentoSector: 53, // Fixed in XSD
		},
	}
}

// NewLubricantesIehdDetalle crea el constructor para los detalles.
func NewLubricantesIehdDetalleBuilder() *lubricantesIehdDetalleBuilder {
	return &lubricantesIehdDetalleBuilder{
		detalle: &documentos.DetalleLubricantesIehd{},
	}
}

type lubricantesIehdBuilder struct {
	factura *documentos.FacturaLubricantesIehd
}

func (b *lubricantesIehdBuilder) WithCabecera(req LubricantesIehdCabecera) *lubricantesIehdBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *lubricantesIehdBuilder) AddDetalle(req LubricantesIehdDetalle) *lubricantesIehdBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

func (b *lubricantesIehdBuilder) WithModalidad(tipo int) *lubricantesIehdBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaImportacionComercializacionLubricantesIEHD"}
		b.factura.XsiSchemaLocation = "facturaElectronicaImportacionComercializacionLubricantesIEHD.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaImportacionComercializacionLubricantesIEHD"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaImportacionComercializacionLubricantesIEHD.xsd"
	}
	return b
}

func (b *lubricantesIehdBuilder) Build() LubricantesIehd {
	return LubricantesIehd{requestWrapper[documentos.FacturaLubricantesIehd]{request: b.factura}}
}

type lubricantesIehdCabeceraBuilder struct {
	cabecera *documentos.CabeceraLubricantesIehd
}

func (b *lubricantesIehdCabeceraBuilder) WithNitEmisor(v int64) *lubricantesIehdCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithRazonSocialEmisor(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithMunicipio(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithTelefono(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithNumeroFactura(v int64) *lubricantesIehdCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCuf(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCufd(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCodigoSucursal(v int) *lubricantesIehdCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithDireccion(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCodigoPuntoVenta(v *int) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithFechaEmision(v time.Time) *lubricantesIehdCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.TimeSiat(v)
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithNombreRazonSocial(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *lubricantesIehdCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithNumeroDocumento(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithComplemento(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCodigoCliente(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCiudad(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Ciudad = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithNombrePropietario(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithNombreRepresentanteLegal(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCondicionPago(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CondicionPago = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CondicionPago = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithPeriodoEntrega(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCodigoMetodoPago(v int) *lubricantesIehdCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithNumeroTarjeta(v *int64) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithMontoTotal(v float64) *lubricantesIehdCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithMontoDeduccionIehdDS25530(v *float64) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.MontoDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *lubricantesIehdCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCodigoMoneda(v int) *lubricantesIehdCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithTipoCambio(v float64) *lubricantesIehdCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithMontoTotalMoneda(v float64) *lubricantesIehdCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithDescuentoAdicional(v *float64) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCodigoExcepcion(v *int) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithCafc(v *string) *lubricantesIehdCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithLeyenda(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) WithUsuario(v string) *lubricantesIehdCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *lubricantesIehdCabeceraBuilder) Build() LubricantesIehdCabecera {
	return LubricantesIehdCabecera{requestWrapper[documentos.CabeceraLubricantesIehd]{request: b.cabecera}}
}

type lubricantesIehdDetalleBuilder struct {
	detalle *documentos.DetalleLubricantesIehd
}

func (b *lubricantesIehdDetalleBuilder) WithActividadEconomica(v string) *lubricantesIehdDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithCodigoProductoSin(v int64) *lubricantesIehdDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithCodigoProducto(v string) *lubricantesIehdDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithDescripcion(v string) *lubricantesIehdDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithCantidad(v float64) *lubricantesIehdDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithUnidadMedida(v int) *lubricantesIehdDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithPrecioUnitario(v float64) *lubricantesIehdDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithMontoDescuento(v *float64) *lubricantesIehdDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithSubTotal(v float64) *lubricantesIehdDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithCantidadLitros(v float64) *lubricantesIehdDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.CantidadLitros = v
	return b
}

func (b *lubricantesIehdDetalleBuilder) WithPorcentajeDeduccionIehdDS25530(v *float64) *lubricantesIehdDetalleBuilder {
	if v == nil {
		b.detalle.PorcentajeDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
	b.detalle.PorcentajeDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *lubricantesIehdDetalleBuilder) Build() LubricantesIehdDetalle {
	return LubricantesIehdDetalle{requestWrapper[documentos.DetalleLubricantesIehd]{request: b.detalle}}
}
