package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// Engarrafadoras representa la estructura completa de una factura de engarrafadoras lista para ser procesada.
type Engarrafadoras struct {
	requestWrapper[documentos.FacturaEngarrafadoras]
}

// EngarrafadorasCabecera representa la sección de cabecera de la factura.
type EngarrafadorasCabecera struct {
	requestWrapper[documentos.CabeceraEngarrafadoras]
}

// EngarrafadorasDetalle representa un ítem individual dentro del detalle.
type EngarrafadorasDetalle struct {
	requestWrapper[documentos.DetalleEngarrafadoras]
}

// NewEngarrafadoras inicia el proceso de construcción de la factura.
func NewEngarrafadorasBuilder() *engarrafadorasBuilder {
	return &engarrafadorasBuilder{
		factura: &documentos.FacturaEngarrafadoras{
			XMLName:           xml.Name{Local: "facturaElectronicaEngarrafadoras"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaEngarrafadoras.xsd",
		},
	}
}

// NewEngarrafadorasCabecera crea el constructor para la cabecera.
func NewEngarrafadorasCabeceraBuilder() *engarrafadorasCabeceraBuilder {
	return &engarrafadorasCabeceraBuilder{
		cabecera: &documentos.CabeceraEngarrafadoras{
			CodigoDocumentoSector: 51, // Fixed in XSD
		},
	}
}

// NewEngarrafadorasDetalle crea el constructor para los detalles.
func NewEngarrafadorasDetalleBuilder() *engarrafadorasDetalleBuilder {
	return &engarrafadorasDetalleBuilder{
		detalle: &documentos.DetalleEngarrafadoras{},
	}
}

type engarrafadorasBuilder struct {
	factura *documentos.FacturaEngarrafadoras
}

func (b *engarrafadorasBuilder) WithCabecera(req EngarrafadorasCabecera) *engarrafadorasBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *engarrafadorasBuilder) AddDetalle(req EngarrafadorasDetalle) *engarrafadorasBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

func (b *engarrafadorasBuilder) WithModalidad(tipo int) *engarrafadorasBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaEngarrafadoras"}
		b.factura.XsiSchemaLocation = "facturaElectronicaEngarrafadoras.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaEngarrafadoras"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaEngarrafadoras.xsd"
	}
	return b
}

func (b *engarrafadorasBuilder) Build() Engarrafadoras {
	return Engarrafadoras{requestWrapper[documentos.FacturaEngarrafadoras]{request: b.factura}}
}

type engarrafadorasCabeceraBuilder struct {
	cabecera *documentos.CabeceraEngarrafadoras
}

func (b *engarrafadorasCabeceraBuilder) WithNitEmisor(v int64) *engarrafadorasCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithRazonSocialEmisor(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithMunicipio(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithTelefono(telefono *string) *engarrafadorasCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithNumeroFactura(v int64) *engarrafadorasCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCuf(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCufd(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCodigoSucursal(v int) *engarrafadorasCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithDireccion(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCodigoPuntoVenta(v *int) *engarrafadorasCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithFechaEmision(v time.Time) *engarrafadorasCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithNombreRazonSocial(v *string) *engarrafadorasCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *engarrafadorasCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithNumeroDocumento(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithComplemento(v *string) *engarrafadorasCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCodigoCliente(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCodigoMetodoPago(v int) *engarrafadorasCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithNumeroTarjeta(v *int64) *engarrafadorasCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithMontoTotal(v float64) *engarrafadorasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *engarrafadorasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCodigoMoneda(v int) *engarrafadorasCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithTipoCambio(v float64) *engarrafadorasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithMontoTotalMoneda(v float64) *engarrafadorasCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithDescuentoAdicional(v *float64) *engarrafadorasCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCodigoExcepcion(v *int) *engarrafadorasCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithCafc(v *string) *engarrafadorasCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithLeyenda(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) WithUsuario(v string) *engarrafadorasCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *engarrafadorasCabeceraBuilder) Build() EngarrafadorasCabecera {
	return EngarrafadorasCabecera{requestWrapper[documentos.CabeceraEngarrafadoras]{request: b.cabecera}}
}

type engarrafadorasDetalleBuilder struct {
	detalle *documentos.DetalleEngarrafadoras
}

func (b *engarrafadorasDetalleBuilder) WithActividadEconomica(v string) *engarrafadorasDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *engarrafadorasDetalleBuilder) WithCodigoProductoSin(v int64) *engarrafadorasDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *engarrafadorasDetalleBuilder) WithCodigoProducto(v string) *engarrafadorasDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *engarrafadorasDetalleBuilder) WithDescripcion(v string) *engarrafadorasDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *engarrafadorasDetalleBuilder) WithCantidad(v float64) *engarrafadorasDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *engarrafadorasDetalleBuilder) WithUnidadMedida(v int) *engarrafadorasDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *engarrafadorasDetalleBuilder) WithPrecioUnitario(v float64) *engarrafadorasDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *engarrafadorasDetalleBuilder) WithMontoDescuento(v *float64) *engarrafadorasDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *engarrafadorasDetalleBuilder) WithSubTotal(v float64) *engarrafadorasDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *engarrafadorasDetalleBuilder) Build() EngarrafadorasDetalle {
	return EngarrafadorasDetalle{requestWrapper[documentos.DetalleEngarrafadoras]{request: b.detalle}}
}
