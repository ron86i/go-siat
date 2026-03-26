package invoices

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
)

// AlquilerBienInmueble representa la estructura completa de una factura de Alquiler de Bienes Inmuebles lista para ser procesada.
type AlquilerBienInmueble struct {
	models.RequestWrapper[documents.FacturaAlquilerBienInmueble]
}

// AlquilerBienInmuebleCabecera representa la sección de cabecera de una factura de Alquiler de Bienes Inmuebles.
type AlquilerBienInmuebleCabecera struct {
	models.RequestWrapper[documents.CabeceraAlquilerBienInmueble]
}

// AlquilerBienInmuebleDetalle representa un ítem individual dentro del detalle de una factura de Alquiler de Bienes Inmuebles.
type AlquilerBienInmuebleDetalle struct {
	models.RequestWrapper[documents.DetalleAlquilerBienInmueble]
}

// NewAlquilerBienInmuebleBuilder inicia el proceso de construcción de una Factura de Alquiler de Bienes Inmuebles.
func NewAlquilerBienInmuebleBuilder() *alquilerBienInmuebleBuilder {
	return &alquilerBienInmuebleBuilder{
		factura: &documents.FacturaAlquilerBienInmueble{
			XMLName:           xml.Name{Local: "facturaElectronicaAlquilerBienInmueble"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaAlquilerBienInmueble.xsd",
		},
	}
}

// NewAlquilerBienInmuebleCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewAlquilerBienInmuebleCabeceraBuilder() *alquilerBienInmuebleCabeceraBuilder {
	return &alquilerBienInmuebleCabeceraBuilder{
		cabecera: &documents.CabeceraAlquilerBienInmueble{
			CodigoDocumentoSector: 2, // Sector 2 para Alquiler de Bienes Inmuebles
		},
	}
}

// NewAlquilerBienInmuebleDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewAlquilerBienInmuebleDetalleBuilder() *alquilerBienInmuebleDetalleBuilder {
	return &alquilerBienInmuebleDetalleBuilder{
		detalle: &documents.DetalleAlquilerBienInmueble{},
	}
}

type alquilerBienInmuebleBuilder struct {
	factura *documents.FacturaAlquilerBienInmueble
}

func (b *alquilerBienInmuebleBuilder) WithCabecera(req AlquilerBienInmuebleCabecera) *alquilerBienInmuebleBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraAlquilerBienInmueble](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *alquilerBienInmuebleBuilder) AddDetalle(req AlquilerBienInmuebleDetalle) *alquilerBienInmuebleBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleAlquilerBienInmueble](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *alquilerBienInmuebleBuilder) WithModalidad(tipo int) *alquilerBienInmuebleBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaAlquilerBienInmueble"}
		b.factura.XsiSchemaLocation = "facturaElectronicaAlquilerBienInmueble.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaAlquilerBienInmueble"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaAlquilerBienInmueble.xsd"
	}
	return b
}

func (b *alquilerBienInmuebleBuilder) Build() AlquilerBienInmueble {
	return AlquilerBienInmueble{models.NewRequestWrapper(b.factura)}
}

type alquilerBienInmuebleCabeceraBuilder struct {
	cabecera *documents.CabeceraAlquilerBienInmueble
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNitEmisor(v int64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithRazonSocialEmisor(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithMunicipio(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithTelefono(telefono *string) *alquilerBienInmuebleCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNumeroFactura(v int64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCuf(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCufd(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoSucursal(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithDireccion(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoPuntoVenta(v *int) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNombreRazonSocial(v *string) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNumeroDocumento(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithComplemento(v *string) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoCliente(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithPeriodoFacturado(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.PeriodoFacturado = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoMetodoPago(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNumeroTarjeta(v *int64) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithMontoTotal(v float64) *alquilerBienInmuebleCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *alquilerBienInmuebleCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoMoneda(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithTipoCambio(v float64) *alquilerBienInmuebleCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithMontoTotalMoneda(v float64) *alquilerBienInmuebleCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithDescuentoAdicional(v *float64) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoExcepcion(v *int) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCafc(v *string) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithLeyenda(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithUsuario(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoDocumentoSector(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) Build() AlquilerBienInmuebleCabecera {
	return AlquilerBienInmuebleCabecera{models.NewRequestWrapper(b.cabecera)}
}

type alquilerBienInmuebleDetalleBuilder struct {
	detalle *documents.DetalleAlquilerBienInmueble
}

func (b *alquilerBienInmuebleDetalleBuilder) WithActividadEconomica(v string) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithCodigoProductoSin(v int64) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithCodigoProducto(v string) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithDescripcion(v string) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithCantidad(v float64) *alquilerBienInmuebleDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithUnidadMedida(v int) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithPrecioUnitario(v float64) *alquilerBienInmuebleDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithMontoDescuento(v *float64) *alquilerBienInmuebleDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithSubTotal(v float64) *alquilerBienInmuebleDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) Build() AlquilerBienInmuebleDetalle {
	return AlquilerBienInmuebleDetalle{models.NewRequestWrapper(b.detalle)}
}
