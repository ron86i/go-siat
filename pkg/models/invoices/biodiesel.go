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

// Biodiesel representa la estructura completa de una factura de Biodiesel lista para ser procesada.
type Biodiesel struct {
	models.RequestWrapper[documents.FacturaBiodiesel]
}

// BiodieselCabecera representa la sección de cabecera de una factura de Biodiesel.
type BiodieselCabecera struct {
	models.RequestWrapper[documents.CabeceraBiodiesel]
}

// BiodieselDetalle representa un ítem individual dentro del detalle de una factura de Biodiesel.
type BiodieselDetalle struct {
	models.RequestWrapper[documents.DetalleBiodiesel]
}

// NewBiodieselBuilder inicia el proceso de construcción de una Factura de Biodiesel.
func NewBiodieselBuilder() *biodieselBuilder {
	return &biodieselBuilder{
		factura: &documents.FacturaBiodiesel{
			XMLName:           xml.Name{Local: "facturaElectronicaBiodiesel"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaBiodiesel.xsd",
		},
	}
}

// NewBiodieselCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewBiodieselCabeceraBuilder() *biodieselCabeceraBuilder {
	return &biodieselCabeceraBuilder{
		cabecera: &documents.CabeceraBiodiesel{
			CodigoDocumentoSector: 54, // Sector 54 para Biodiesel
			MontoTotalSujetoIva:   0,  // Fijo 0 según XSD
		},
	}
}

// NewBiodieselDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewBiodieselDetalleBuilder() *biodieselDetalleBuilder {
	return &biodieselDetalleBuilder{
		detalle: &documents.DetalleBiodiesel{},
	}
}

type biodieselBuilder struct {
	factura *documents.FacturaBiodiesel
}

func (b *biodieselBuilder) WithCabecera(req BiodieselCabecera) *biodieselBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraBiodiesel](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *biodieselBuilder) AddDetalle(req BiodieselDetalle) *biodieselBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleBiodiesel](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *biodieselBuilder) WithModalidad(tipo int) *biodieselBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaBiodiesel"}
		b.factura.XsiSchemaLocation = "facturaElectronicaBiodiesel.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaBiodiesel"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaBiodiesel.xsd"
	}
	return b
}

func (b *biodieselBuilder) Build() Biodiesel {
	return Biodiesel{models.NewRequestWrapper(b.factura)}
}

type biodieselCabeceraBuilder struct {
	cabecera *documents.CabeceraBiodiesel
}

func (b *biodieselCabeceraBuilder) WithNitEmisor(v int64) *biodieselCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *biodieselCabeceraBuilder) WithRazonSocialEmisor(v string) *biodieselCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *biodieselCabeceraBuilder) WithMunicipio(v string) *biodieselCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *biodieselCabeceraBuilder) WithTelefono(telefono *string) *biodieselCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithNumeroFactura(v int64) *biodieselCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *biodieselCabeceraBuilder) WithCuf(v string) *biodieselCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *biodieselCabeceraBuilder) WithCufd(v string) *biodieselCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoSucursal(v int) *biodieselCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *biodieselCabeceraBuilder) WithDireccion(v string) *biodieselCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoPuntoVenta(v *int) *biodieselCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithFechaEmision(v time.Time) *biodieselCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *biodieselCabeceraBuilder) WithNombreRazonSocial(v *string) *biodieselCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *biodieselCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *biodieselCabeceraBuilder) WithNumeroDocumento(v string) *biodieselCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *biodieselCabeceraBuilder) WithComplemento(v *string) *biodieselCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoCliente(v string) *biodieselCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoMetodoPago(v int) *biodieselCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *biodieselCabeceraBuilder) WithNumeroTarjeta(v *int64) *biodieselCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithMontoTotal(v float64) *biodieselCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoMoneda(v int) *biodieselCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *biodieselCabeceraBuilder) WithTipoCambio(v float64) *biodieselCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *biodieselCabeceraBuilder) WithMontoTotalMoneda(v float64) *biodieselCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *biodieselCabeceraBuilder) WithDescuentoAdicional(v *float64) *biodieselCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoExcepcion(v *int) *biodieselCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithCafc(v *string) *biodieselCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *biodieselCabeceraBuilder) WithLeyenda(v string) *biodieselCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *biodieselCabeceraBuilder) WithUsuario(v string) *biodieselCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *biodieselCabeceraBuilder) WithCodigoDocumentoSector(v int) *biodieselCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *biodieselCabeceraBuilder) Build() BiodieselCabecera {
	return BiodieselCabecera{models.NewRequestWrapper(b.cabecera)}
}

type biodieselDetalleBuilder struct {
	detalle *documents.DetalleBiodiesel
}

func (b *biodieselDetalleBuilder) WithActividadEconomica(v string) *biodieselDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *biodieselDetalleBuilder) WithCodigoProductoSin(v int64) *biodieselDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *biodieselDetalleBuilder) WithCodigoProducto(v string) *biodieselDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *biodieselDetalleBuilder) WithDescripcion(v string) *biodieselDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *biodieselDetalleBuilder) WithCantidad(v float64) *biodieselDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *biodieselDetalleBuilder) WithUnidadMedida(v int) *biodieselDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *biodieselDetalleBuilder) WithPrecioUnitario(v float64) *biodieselDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *biodieselDetalleBuilder) WithMontoDescuento(v *float64) *biodieselDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *biodieselDetalleBuilder) WithSubTotal(v float64) *biodieselDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *biodieselDetalleBuilder) Build() BiodieselDetalle {
	return BiodieselDetalle{models.NewRequestWrapper(b.detalle)}
}
