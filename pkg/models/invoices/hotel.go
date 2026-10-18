package invoices

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/shopspring/decimal"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
)

// Hotel representa la estructura completa de una factura de hotel
// lista para ser procesada.
type Hotel struct {
	models.RequestWrapper[documents.FacturaHotel]
	models.MetadatosFactura
}

// HotelCabecera representa la sección de cabecera de la factura.
type HotelCabecera struct {
	models.RequestWrapper[documents.CabeceraHotel]
}

// HotelDetalle representa un ítem individual dentro del detalle.
type HotelDetalle struct {
	models.RequestWrapper[documents.DetalleHotel]
}

// NewHotel inicia el proceso de construcción de la factura.
func NewHotelBuilder() *hotelBuilder {
	return &hotelBuilder{
		metadatos: models.NuevosMetadatosFactura(16),
		factura: &documents.FacturaHotel{
			XMLName:           xml.Name{Local: "facturaElectronicaHotel"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaHotel.xsd",
		},
	}
}

// NewHotelCabecera crea el constructor para la cabecera.
func NewHotelCabeceraBuilder() *hotelCabeceraBuilder {
	return &hotelCabeceraBuilder{
		cabecera: &documents.CabeceraHotel{
			CodigoDocumentoSector: 16, // Fixed in XSD
		},
	}
}

// NewHotelDetalle crea el constructor para los detalles.
func NewHotelDetalleBuilder() *hotelDetalleBuilder {
	return &hotelDetalleBuilder{
		detalle: &documents.DetalleHotel{},
	}
}

type hotelBuilder struct {
	factura   *documents.FacturaHotel
	metadatos models.MetadatosFactura
}

// WithTipoFacturaDocumento reemplaza el tipo fiscal predeterminado de hotel.
func (b *hotelBuilder) WithTipoFacturaDocumento(tipo int) *hotelBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *hotelBuilder) WithCabecera(req HotelCabecera) *hotelBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraHotel](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *hotelBuilder) AddDetalle(req HotelDetalle) *hotelBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleHotel](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
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
	return Hotel{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type hotelCabeceraBuilder struct {
	cabecera *documents.CabeceraHotel
}

func (b *hotelCabeceraBuilder) Build() HotelCabecera {
	return HotelCabecera{models.NewRequestWrapper(b.cabecera)}
}

type hotelDetalleBuilder struct {
	detalle *documents.DetalleHotel
	err     error
}

func (b *hotelDetalleBuilder) Build() HotelDetalle {
	return HotelDetalle{models.NewRequestWrapper(b.detalle)}
}

// Err devuelve el último error de serialización de un campo JSON libre.
func (b *hotelDetalleBuilder) Err() error {
	return b.err
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
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
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
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithFechaEmision(v time.Time) *hotelCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *hotelCabeceraBuilder) WithNombreRazonSocial(v *string) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
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

func (b *hotelCabeceraBuilder) WithComplemento(v string) *hotelCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoCliente(v string) *hotelCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadHuespedes(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadHuespedes = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CantidadHuespedes = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadHabitaciones(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadHabitaciones = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CantidadHabitaciones = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadMayores(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadMayores = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CantidadMayores = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithCantidadMenores(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadMenores = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CantidadMenores = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithFechaIngresoHospedaje(v time.Time) *hotelCabeceraBuilder {
	b.cabecera.FechaIngresoHospedaje = datatype.NewTimeSiat(v)
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoMetodoPago(v int) *hotelCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *hotelCabeceraBuilder) WithNumeroTarjeta(v *int64) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithSubTotalNeto(v float64) *hotelCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelCabeceraBuilder) WithTotalPrecioNeto(v float64) *hotelCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelCabeceraBuilder) WithIva(v float64) *hotelCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelCabeceraBuilder) WithMontoAPagar(v float64) *hotelCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoMoneda(v int) *hotelCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *hotelCabeceraBuilder) WithTipoCambio(v float64) *hotelCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelCabeceraBuilder) WithMontoTotalMoneda(v float64) *hotelCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelCabeceraBuilder) WithMontoGiftCard(v *float64) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithDescuentoAdicional(v *float64) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoExcepcion(v *int) *hotelCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
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

func (b *hotelCabeceraBuilder) WithCodigoInterno(v string) *hotelCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *hotelCabeceraBuilder) WithCodigoDocumentoSector(v int) *hotelCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
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
		b.detalle.CodigoTipoHabitacion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.detalle.CodigoTipoHabitacion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hotelDetalleBuilder) WithDescripcion(v string) *hotelDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *hotelDetalleBuilder) WithCantidad(v float64) *hotelDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelDetalleBuilder) WithUnidadMedida(v int) *hotelDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *hotelDetalleBuilder) WithPrecioNetoUnitario(v float64) *hotelDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *hotelDetalleBuilder) WithMontoDescuento(v *float64) *hotelDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hotelDetalleBuilder) WithSubTotal(v float64) *hotelDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
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
		jsonData, err := json.Marshal(v)
		if err != nil {
			b.err = err
			return b
		}
		val := string(jsonData)
		b.detalle.DetalleHuespedes = datatype.Nilable[string]{Value: &val}
	}
	return b
}
