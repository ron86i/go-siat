package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// TurismoHospedaje representa la estructura completa de una factura de Turismo y Hospedaje lista para ser procesada.
type TurismoHospedaje struct {
	models.RequestWrapper[documentos.FacturaTurismoHospedaje]
}

// TurismoHospedajeCabecera representa la sección de cabecera de una factura de Turismo y Hospedaje.
type TurismoHospedajeCabecera struct {
	models.RequestWrapper[documentos.CabeceraTurismoHospedaje]
}

// TurismoHospedajeDetalle representa un ítem individual dentro del detalle de una factura de Turismo y Hospedaje.
type TurismoHospedajeDetalle struct {
	models.RequestWrapper[documentos.DetalleTurismoHospedaje]
}

// NewTurismoHospedajeBuilder inicia el proceso de construcción de una Factura de Turismo y Hospedaje.
func NewTurismoHospedajeBuilder() *turismoHospedajeBuilder {
	return &turismoHospedajeBuilder{
		factura: &documentos.FacturaTurismoHospedaje{
			XMLName:           xml.Name{Local: "facturaElectronicaServicioTuristicoHospedaje"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaServicioTuristicoHospedaje.xsd",
		},
	}
}

// NewTurismoHospedajeCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewTurismoHospedajeCabeceraBuilder() *turismoHospedajeCabeceraBuilder {
	return &turismoHospedajeCabeceraBuilder{
		cabecera: &documentos.CabeceraTurismoHospedaje{
			CodigoDocumentoSector: 6, // Sector 6 para Turismo y Hospedaje
			MontoTotalSujetoIva:   0, // Siempre 0 para este sector
		},
	}
}

// NewTurismoHospedajeDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewTurismoHospedajeDetalleBuilder() *turismoHospedajeDetalleBuilder {
	return &turismoHospedajeDetalleBuilder{
		detalle: &documentos.DetalleTurismoHospedaje{},
	}
}

type turismoHospedajeBuilder struct {
	factura *documentos.FacturaTurismoHospedaje
}

func (b *turismoHospedajeBuilder) WithCabecera(req TurismoHospedajeCabecera) *turismoHospedajeBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraTurismoHospedaje](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *turismoHospedajeBuilder) AddDetalle(req TurismoHospedajeDetalle) *turismoHospedajeBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleTurismoHospedaje](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *turismoHospedajeBuilder) WithModalidad(tipo int) *turismoHospedajeBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaServicioTuristicoHospedaje"}
		b.factura.XsiSchemaLocation = "facturaElectronicaServicioTuristicoHospedaje.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaServicioTuristicoHospedaje"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaServicioTuristicoHospedaje.xsd"
	}
	return b
}

func (b *turismoHospedajeBuilder) Build() TurismoHospedaje {
	return TurismoHospedaje{models.NewRequestWrapper(b.factura)}
}

type turismoHospedajeCabeceraBuilder struct {
	cabecera *documentos.CabeceraTurismoHospedaje
}

func (b *turismoHospedajeCabeceraBuilder) WithNitEmisor(v int64) *turismoHospedajeCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithRazonSocialEmisor(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithMunicipio(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithTelefono(telefono *string) *turismoHospedajeCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithNumeroFactura(v int64) *turismoHospedajeCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCuf(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCufd(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCodigoSucursal(v int) *turismoHospedajeCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithDireccion(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCodigoPuntoVenta(v *int) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *turismoHospedajeCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithNombreRazonSocial(v *string) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *turismoHospedajeCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithNumeroDocumento(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithComplemento(v *string) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCodigoCliente(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithRazonSocialOperadorTurismo(v *string) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.RazonSocialOperadorTurismo = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.RazonSocialOperadorTurismo = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCantidadHuespedes(v *int) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadHuespedes = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CantidadHuespedes = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCantidadHabitaciones(v *int) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadHabitaciones = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CantidadHabitaciones = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCantidadMayores(v *int) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadMayores = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CantidadMayores = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCantidadMenores(v *int) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.CantidadMenores = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CantidadMenores = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithFechaIngresoHospedaje(v *time.Time) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.FechaIngresoHospedaje = datatype.Nilable[datatype.TimeSiat]{Value: nil}
		return b
	}
	value := datatype.NewTimeSiat(*v)
	b.cabecera.FechaIngresoHospedaje = datatype.Nilable[datatype.TimeSiat]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCodigoMetodoPago(v int) *turismoHospedajeCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithNumeroTarjeta(v *int64) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithMontoTotal(v float64) *turismoHospedajeCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCodigoMoneda(v int) *turismoHospedajeCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithTipoCambio(v float64) *turismoHospedajeCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithMontoTotalMoneda(v float64) *turismoHospedajeCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithMontoGiftCard(v *float64) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithDescuentoAdicional(v *float64) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCodigoExcepcion(v *int) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithCafc(v *string) *turismoHospedajeCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithLeyenda(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithUsuario(v string) *turismoHospedajeCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *turismoHospedajeCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *turismoHospedajeCabeceraBuilder) WithCodigoDocumentoSector(v int) *turismoHospedajeCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *turismoHospedajeCabeceraBuilder) Build() TurismoHospedajeCabecera {
	return TurismoHospedajeCabecera{models.NewRequestWrapper(b.cabecera)}
}

type turismoHospedajeDetalleBuilder struct {
	detalle *documentos.DetalleTurismoHospedaje
}

func (b *turismoHospedajeDetalleBuilder) WithActividadEconomica(v string) *turismoHospedajeDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithCodigoProductoSin(v int64) *turismoHospedajeDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithCodigoProducto(v string) *turismoHospedajeDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithDescripcion(v string) *turismoHospedajeDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithCodigoTipoHabitacion(v *int) *turismoHospedajeDetalleBuilder {
	if v == nil {
		b.detalle.CodigoTipoHabitacion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.detalle.CodigoTipoHabitacion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithCantidad(v float64) *turismoHospedajeDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithUnidadMedida(v int) *turismoHospedajeDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithPrecioUnitario(v float64) *turismoHospedajeDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithMontoDescuento(v *float64) *turismoHospedajeDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithSubTotal(v float64) *turismoHospedajeDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *turismoHospedajeDetalleBuilder) WithDetalleHuespedes(v *string) *turismoHospedajeDetalleBuilder {
	if v == nil {
		b.detalle.DetalleHuespedes = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.DetalleHuespedes = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *turismoHospedajeDetalleBuilder) Build() TurismoHospedajeDetalle {
	return TurismoHospedajeDetalle{models.NewRequestWrapper(b.detalle)}
}
