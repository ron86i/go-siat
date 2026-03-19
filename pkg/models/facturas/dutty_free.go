package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// DuttyFree representa la estructura completa de una factura Dutty Free lista para ser procesada.
type DuttyFree struct {
	requestWrapper[documentos.FacturaDuttyFree]
}

// DuttyFreeCabecera representa la sección de cabecera de una factura Dutty Free.
type DuttyFreeCabecera struct {
	requestWrapper[documentos.CabeceraDuttyFree]
}

// DuttyFreeDetalle representa un ítem individual dentro del detalle de una factura Dutty Free.
type DuttyFreeDetalle struct {
	requestWrapper[documentos.DetalleDuttyFree]
}

// NewDuttyFree inicia el proceso de construcción de una DuttyFree.
// NewDuttyFree inicia el proceso de construcción de una DuttyFree.
func NewDuttyFreeBuilder() *duttyFreeBuilder {
	return &duttyFreeBuilder{
		factura: &documentos.FacturaDuttyFree{
			XMLName:           xml.Name{Local: "facturaElectronicaDuttyFree"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaDuttyFree.xsd",
		},
	}
}

// NewDuttyFreeCabecera crea una instancia del constructor para la cabecera.
// NewDuttyFreeCabecera crea una instancia del constructor para la cabecera.
func NewDuttyFreeCabeceraBuilder() *duttyFreeCabeceraBuilder {
	return &duttyFreeCabeceraBuilder{
		cabecera: &documentos.CabeceraDuttyFree{
			MontoTotalSujetoIva:   0,  // Fixed in XSD
			CodigoDocumentoSector: 10, // Fixed in XSD
		},
	}
}

// NewDuttyFreeDetalle crea una instancia del constructor para los ítems de detalle.
// NewDuttyFreeDetalle crea una instancia del constructor para los ítems de detalle.
func NewDuttyFreeDetalleBuilder() *duttyFreeDetalleBuilder {
	return &duttyFreeDetalleBuilder{
		detalle: &documentos.DetalleDuttyFree{},
	}
}

type duttyFreeBuilder struct {
	factura *documentos.FacturaDuttyFree
}

func (b *duttyFreeBuilder) WithCabecera(req DuttyFreeCabecera) *duttyFreeBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *duttyFreeBuilder) AddDetalle(req DuttyFreeDetalle) *duttyFreeBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

func (b *duttyFreeBuilder) WithModalidad(tipo int) *duttyFreeBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaDuttyFree"}
		b.factura.XsiSchemaLocation = "facturaElectronicaDuttyFree.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaDuttyFree"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaDuttyFree.xsd"
	}
	return b
}

func (b *duttyFreeBuilder) Build() DuttyFree {
	return DuttyFree{requestWrapper[documentos.FacturaDuttyFree]{request: b.factura}}
}

type duttyFreeCabeceraBuilder struct {
	cabecera *documentos.CabeceraDuttyFree
}

func (b *duttyFreeCabeceraBuilder) WithNitEmisor(v int64) *duttyFreeCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithRazonSocialEmisor(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithMunicipio(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithTelefono(telefono *string) *duttyFreeCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithNumeroFactura(v int64) *duttyFreeCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCuf(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCufd(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCodigoSucursal(v int) *duttyFreeCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithDireccion(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCodigoPuntoVenta(v *int) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *duttyFreeCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *duttyFreeCabeceraBuilder) WithNombreRazonSocial(v *string) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *duttyFreeCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithNumeroDocumento(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithComplemento(v *string) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCodigoCliente(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCodigoMetodoPago(v int) *duttyFreeCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithNumeroTarjeta(v *int64) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithMontoTotal(v float64) *duttyFreeCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCodigoMoneda(v int) *duttyFreeCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithTipoCambio(v float64) *duttyFreeCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithMontoTotalMoneda(v float64) *duttyFreeCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithMontoGiftCard(v *float64) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithDescuentoAdicional(v *float64) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCodigoExcepcion(v *int) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithCafc(v *string) *duttyFreeCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *duttyFreeCabeceraBuilder) WithLeyenda(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *duttyFreeCabeceraBuilder) WithUsuario(v string) *duttyFreeCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *duttyFreeCabeceraBuilder) Build() DuttyFreeCabecera {
	return DuttyFreeCabecera{requestWrapper[documentos.CabeceraDuttyFree]{request: b.cabecera}}
}

type duttyFreeDetalleBuilder struct {
	detalle *documentos.DetalleDuttyFree
}

func (b *duttyFreeDetalleBuilder) WithActividadEconomica(v string) *duttyFreeDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *duttyFreeDetalleBuilder) WithCodigoProductoSin(v int64) *duttyFreeDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *duttyFreeDetalleBuilder) WithCodigoProducto(v string) *duttyFreeDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *duttyFreeDetalleBuilder) WithDescripcion(v string) *duttyFreeDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *duttyFreeDetalleBuilder) WithCantidad(v float64) *duttyFreeDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *duttyFreeDetalleBuilder) WithUnidadMedida(v int) *duttyFreeDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *duttyFreeDetalleBuilder) WithPrecioUnitario(v float64) *duttyFreeDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *duttyFreeDetalleBuilder) WithMontoDescuento(v *float64) *duttyFreeDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *duttyFreeDetalleBuilder) WithSubTotal(v float64) *duttyFreeDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *duttyFreeDetalleBuilder) Build() DuttyFreeDetalle {
	return DuttyFreeDetalle{requestWrapper[documentos.DetalleDuttyFree]{request: b.detalle}}
}
