package facturas

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// Hotel representa la estructura completa de una factura de hotel
// lista para ser procesada.
type Hotel struct {
	requestWrapper[documentos.FacturaHotel]
}

// HotelCabecera representa la sección de cabecera de la factura.
type HotelCabecera struct {
	requestWrapper[documentos.CabeceraHotel]
}

// HotelDetalle representa un ítem individual dentro del detalle.
type HotelDetalle struct {
	requestWrapper[documentos.DetalleHotel]
}

// NewHotel inicia el proceso de construcción de la factura.
func NewHotelBuilder() *hotelBuilder {
	return &hotelBuilder{
		factura: &documentos.FacturaHotel{
			XMLName:           xml.Name{Local: "facturaElectronicaHotel"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaHotel.xsd",
		},
	}
}

// NewHotelCabecera crea el constructor para la cabecera.
func NewHotelCabeceraBuilder() *hotelCabeceraBuilder {
	return &hotelCabeceraBuilder{
		cabecera: &documentos.CabeceraHotel{
			CodigoDocumentoSector: 16, // Fixed in XSD
		},
	}
}

// NewHotelDetalle crea el constructor para los detalles.
func NewHotelDetalleBuilder() *hotelDetalleBuilder {
	return &hotelDetalleBuilder{
		detalle: &documentos.DetalleHotel{},
	}
}

type hotelBuilder struct {
	factura *documentos.FacturaHotel
}

func (b *hotelBuilder) WithCabecera(req HotelCabecera) *hotelBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *hotelBuilder) AddDetalle(req HotelDetalle) *hotelBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

func (b *hotelBuilder) WithModalidad(tipo int) *hotelBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaHotel"}
		b.factura.XsiSchemaLocation = "facturaElectronicaHotel.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaHotel"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaHotel.xsd"
	}
	return b
}

func (b *hotelBuilder) Build() Hotel {
	return Hotel{requestWrapper[documentos.FacturaHotel]{request: b.factura}}
}

type hotelCabeceraBuilder struct {
	cabecera *documentos.CabeceraHotel
}

func (b *hotelCabeceraBuilder) WithNitEmisor(v int64) *hotelCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *hotelCabeceraBuilder) WithRazonSocialEmisor(v string) *hotelCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *hotelCabeceraBuilder) WithMunicipio(v string) *hotelCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *hotelCabeceraBuilder) WithTelefono(v *string) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithNumeroFactura(v int64) *hotelCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *hotelCabeceraBuilder) WithCuf(v string) *hotelCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *hotelCabeceraBuilder) WithCufd(v string) *hotelCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoSucursal(v int) *hotelCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *hotelCabeceraBuilder) WithDireccion(v string) *hotelCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoPuntoVenta(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithFechaEmision(v time.Time) *hotelCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.TimeSiat(v)
	return b
}

func (b *hotelCabeceraBuilder) WithNombreRazonSocial(v *string) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *hotelCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *hotelCabeceraBuilder) WithNumeroDocumento(v string) *hotelCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *hotelCabeceraBuilder) WithComplemento(v *string) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoCliente(v string) *hotelCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadHuespedes(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadHuespedes = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CantidadHuespedes = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadHabitaciones(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadHabitaciones = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CantidadHabitaciones = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadMayores(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadMayores = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CantidadMayores = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadMenores(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadMenores = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CantidadMenores = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithFechaIngresoHospedaje(v time.Time) *hotelCabeceraBuilder {
	b.cabecera.FechaIngresoHospedaje = datatype.TimeSiat(v)
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoMetodoPago(v int) *hotelCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *hotelCabeceraBuilder) WithNumeroTarjeta(v *int64) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithMontoTotal(v float64) *hotelCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *hotelCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *hotelCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoMoneda(v int) *hotelCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *hotelCabeceraBuilder) WithTipoCambio(v float64) *hotelCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *hotelCabeceraBuilder) WithMontoTotalMoneda(v float64) *hotelCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *hotelCabeceraBuilder) WithMontoGiftCard(v *float64) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithDescuentoAdicional(v *float64) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoExcepcion(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithCafc(v *string) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	val := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *hotelCabeceraBuilder) WithLeyenda(v string) *hotelCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *hotelCabeceraBuilder) WithUsuario(v string) *hotelCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *hotelCabeceraBuilder) Build() HotelCabecera {
	return HotelCabecera{requestWrapper[documentos.CabeceraHotel]{request: b.cabecera}}
}

type hotelDetalleBuilder struct {
	detalle *documentos.DetalleHotel
}

func (b *hotelDetalleBuilder) WithActividadEconomica(v string) *hotelDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *hotelDetalleBuilder) WithCodigoProductoSin(v int64) *hotelDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *hotelDetalleBuilder) WithCodigoProducto(v string) *hotelDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *hotelDetalleBuilder) WithCodigoTipoHabitacion(v *int) *hotelDetalleBuilder {
	if v == nil {
		b.detalle.CodigoTipoHabitacion = datatype.Nilable[int]{Value: nil}
		return b
	}
	val := *v
	b.detalle.CodigoTipoHabitacion = datatype.Nilable[int]{Value: &val}
	return b
}

func (b *hotelDetalleBuilder) WithDescripcion(v string) *hotelDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *hotelDetalleBuilder) WithCantidad(v float64) *hotelDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *hotelDetalleBuilder) WithUnidadMedida(v int) *hotelDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *hotelDetalleBuilder) WithPrecioUnitario(v float64) *hotelDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *hotelDetalleBuilder) WithMontoDescuento(v *float64) *hotelDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	val := *v
	val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	return b
}

func (b *hotelDetalleBuilder) WithSubTotal(v float64) *hotelDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *hotelDetalleBuilder) WithDetalleHuespedes(v any) *hotelDetalleBuilder {
	if v == nil {
		b.detalle.DetalleHuespedes = datatype.Nilable[string]{Value: nil}
		return b
	}
	// If it's already a string, use it. Otherwise marshal to JSON.
	if str, ok := v.(string); ok {
		b.detalle.DetalleHuespedes = datatype.Nilable[string]{Value: &str}
	} else {
		jsonData, _ := json.Marshal(v)
		val := string(jsonData)
		b.detalle.DetalleHuespedes = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *hotelDetalleBuilder) Build() HotelDetalle {
	return HotelDetalle{requestWrapper[documentos.DetalleHotel]{request: b.detalle}}
}
