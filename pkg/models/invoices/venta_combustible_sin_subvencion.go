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

// VentaCombustibleSinSubvencion representa la estructura completa de una factura de Venta de Combustible Sin Subvención lista para ser procesada.
type VentaCombustibleSinSubvencion struct {
	models.RequestWrapper[documents.FacturaVentaCombustibleSinSubvencion]
}

// VentaCombustibleSinSubvencionCabecera representa la sección de cabecera de una factura de Venta de Combustible Sin Subvención.
type VentaCombustibleSinSubvencionCabecera struct {
	models.RequestWrapper[documents.CabeceraVentaCombustibleSinSubvencion]
}

// VentaCombustibleSinSubvencionDetalle representa un ítem individual dentro del detalle de una factura de Venta de Combustible Sin Subvención.
type VentaCombustibleSinSubvencionDetalle struct {
	models.RequestWrapper[documents.DetalleVentaCombustibleSinSubvencion]
}

// NewVentaCombustibleSinSubvencionBuilder inicia el proceso de construcción de una Factura de Venta de Combustible Sin Subvención.
func NewVentaCombustibleSinSubvencionBuilder() *ventaCombustibleSinSubvencionBuilder {
	return &ventaCombustibleSinSubvencionBuilder{
		factura: &documents.FacturaVentaCombustibleSinSubvencion{
			XMLName:           xml.Name{Local: "facturaElectronicaVentaCombustibleSinSubvencion"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaVentaCombustibleSinSubvencion.xsd",
		},
	}
}

// NewVentaCombustibleSinSubvencionCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewVentaCombustibleSinSubvencionCabeceraBuilder() *ventaCombustibleSinSubvencionCabeceraBuilder {
	return &ventaCombustibleSinSubvencionCabeceraBuilder{
		cabecera: &documents.CabeceraVentaCombustibleSinSubvencion{
			CodigoDocumentoSector: 55, // Sector 55 para Venta de Combustible Sin Subvención
		},
	}
}

// NewVentaCombustibleSinSubvencionDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewVentaCombustibleSinSubvencionDetalleBuilder() *ventaCombustibleSinSubvencionDetalleBuilder {
	return &ventaCombustibleSinSubvencionDetalleBuilder{
		detalle: &documents.DetalleVentaCombustibleSinSubvencion{},
	}
}

type ventaCombustibleSinSubvencionBuilder struct {
	factura *documents.FacturaVentaCombustibleSinSubvencion
}

func (b *ventaCombustibleSinSubvencionBuilder) WithCabecera(req VentaCombustibleSinSubvencionCabecera) *ventaCombustibleSinSubvencionBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraVentaCombustibleSinSubvencion](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *ventaCombustibleSinSubvencionBuilder) AddDetalle(req VentaCombustibleSinSubvencionDetalle) *ventaCombustibleSinSubvencionBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleVentaCombustibleSinSubvencion](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *ventaCombustibleSinSubvencionBuilder) WithModalidad(tipo int) *ventaCombustibleSinSubvencionBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaVentaCombustibleSinSubvencion"}
		b.factura.XsiSchemaLocation = "facturaElectronicaVentaCombustibleSinSubvencion.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaVentaCombustibleSinSubvencion"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaVentaCombustibleSinSubvencion.xsd"
	}
	return b
}

func (b *ventaCombustibleSinSubvencionBuilder) Build() VentaCombustibleSinSubvencion {
	return VentaCombustibleSinSubvencion{models.NewRequestWrapper(b.factura)}
}

type ventaCombustibleSinSubvencionCabeceraBuilder struct {
	cabecera *documents.CabeceraVentaCombustibleSinSubvencion
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNitEmisor(v int64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithRazonSocialEmisor(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithMunicipio(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithTelefono(telefono *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNumeroFactura(v int64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCuf(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCufd(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoSucursal(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithDireccion(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoPuntoVenta(v *int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNombreRazonSocial(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNumeroDocumento(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithComplemento(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoCliente(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoPais(v *int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPais = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPais = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithPlacaVehiculo(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.PlacaVehiculo = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithTipoEnvase(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.TipoEnvase = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.TipoEnvase = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoMetodoPago(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithNumeroTarjeta(v *int64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithMontoTotal(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoAutorizacionSC(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoAutorizacionSC = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoAutorizacionSC = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithObservacion(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.Observacion = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Observacion = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoMoneda(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithTipoCambio(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithMontoTotalMoneda(v float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithDescuentoAdicional(v *float64) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoExcepcion(v *int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCafc(v *string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithLeyenda(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithUsuario(v string) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *ventaCombustibleSinSubvencionCabeceraBuilder) WithCodigoDocumentoSector(v int) *ventaCombustibleSinSubvencionCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *ventaCombustibleSinSubvencionCabeceraBuilder) Build() VentaCombustibleSinSubvencionCabecera {
	return VentaCombustibleSinSubvencionCabecera{models.NewRequestWrapper(b.cabecera)}
}

type ventaCombustibleSinSubvencionDetalleBuilder struct {
	detalle *documents.DetalleVentaCombustibleSinSubvencion
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithActividadEconomica(v string) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithCodigoProductoSin(v int64) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithCodigoProducto(v string) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithDescripcion(v string) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithCantidad(v float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithUnidadMedida(v int) *ventaCombustibleSinSubvencionDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithPrecioUnitario(v float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithMontoDescuento(v *float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) WithSubTotal(v float64) *ventaCombustibleSinSubvencionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *ventaCombustibleSinSubvencionDetalleBuilder) Build() VentaCombustibleSinSubvencionDetalle {
	return VentaCombustibleSinSubvencionDetalle{models.NewRequestWrapper(b.detalle)}
}
