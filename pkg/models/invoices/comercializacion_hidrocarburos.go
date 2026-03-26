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

// ComercializacionHidro representa la estructura completa de una factura de Comercialización de Hidrocarburos lista para ser procesada.
type ComercializacionHidro struct {
	models.RequestWrapper[documents.FacturaComercializacionHidro]
}

// ComercializacionHidroCabecera representa la sección de cabecera de una factura de Comercialización de Hidrocarburos.
type ComercializacionHidroCabecera struct {
	models.RequestWrapper[documents.CabeceraComercializacionHidro]
}

// ComercializacionHidroDetalle representa un ítem individual dentro del detalle de una factura de Comercialización de Hidrocarburos.
type ComercializacionHidroDetalle struct {
	models.RequestWrapper[documents.DetalleComercializacionHidro]
}

// NewComercializacionHidroBuilder inicia el proceso de construcción de una Factura de Comercialización de Hidrocarburos.
func NewComercializacionHidroBuilder() *comercializacionHidroBuilder {
	return &comercializacionHidroBuilder{
		factura: &documents.FacturaComercializacionHidro{
			XMLName:           xml.Name{Local: "facturaElectronicaComercializacionHidrocarburo"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercializacionHidrocarburo.xsd",
		},
	}
}

// NewComercializacionHidroCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewComercializacionHidroCabeceraBuilder() *comercializacionHidroCabeceraBuilder {
	return &comercializacionHidroCabeceraBuilder{
		cabecera: &documents.CabeceraComercializacionHidro{
			CodigoDocumentoSector: 12, // Sector 12 para Comercialización de Hidrocarburos
		},
	}
}

// NewComercializacionHidroDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewComercializacionHidroDetalleBuilder() *comercializacionHidroDetalleBuilder {
	return &comercializacionHidroDetalleBuilder{
		detalle: &documents.DetalleComercializacionHidro{},
	}
}

type comercializacionHidroBuilder struct {
	factura *documents.FacturaComercializacionHidro
}

func (b *comercializacionHidroBuilder) WithCabecera(req ComercializacionHidroCabecera) *comercializacionHidroBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraComercializacionHidro](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercializacionHidroBuilder) AddDetalle(req ComercializacionHidroDetalle) *comercializacionHidroBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleComercializacionHidro](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *comercializacionHidroBuilder) WithModalidad(tipo int) *comercializacionHidroBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercializacionHidrocarburo"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercializacionHidrocarburo.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercializacionHidrocarburo"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercializacionHidrocarburo.xsd"
	}
	return b
}

func (b *comercializacionHidroBuilder) Build() ComercializacionHidro {
	return ComercializacionHidro{models.NewRequestWrapper(b.factura)}
}

type comercializacionHidroCabeceraBuilder struct {
	cabecera *documents.CabeceraComercializacionHidro
}

func (b *comercializacionHidroCabeceraBuilder) WithNitEmisor(v int64) *comercializacionHidroCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithRazonSocialEmisor(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithMunicipio(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithTelefono(telefono *string) *comercializacionHidroCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithNumeroFactura(v int64) *comercializacionHidroCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCuf(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCufd(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoSucursal(v int) *comercializacionHidroCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithDireccion(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *comercializacionHidroCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithNombreRazonSocial(v *string) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercializacionHidroCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithNumeroDocumento(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithComplemento(v *string) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoCliente(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoPais(v *int) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPais = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPais = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithPlacaVehiculo(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.PlacaVehiculo = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithTipoEnvase(v *string) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.TipoEnvase = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.TipoEnvase = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoMetodoPago(v int) *comercializacionHidroCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithMontoTotal(v float64) *comercializacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *comercializacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoAutorizacionSC(v *string) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoAutorizacionSC = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoAutorizacionSC = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithObservacion(v *string) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.Observacion = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Observacion = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoMoneda(v int) *comercializacionHidroCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithTipoCambio(v float64) *comercializacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercializacionHidroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCodigoExcepcion(v *int) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithCafc(v *string) *comercializacionHidroCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithLeyenda(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) WithUsuario(v string) *comercializacionHidroCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *comercializacionHidroCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercializacionHidroCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *comercializacionHidroCabeceraBuilder) Build() ComercializacionHidroCabecera {
	return ComercializacionHidroCabecera{models.NewRequestWrapper(b.cabecera)}
}

type comercializacionHidroDetalleBuilder struct {
	detalle *documents.DetalleComercializacionHidro
}

func (b *comercializacionHidroDetalleBuilder) WithActividadEconomica(v string) *comercializacionHidroDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithCodigoProductoSin(v int64) *comercializacionHidroDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithCodigoProducto(v string) *comercializacionHidroDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithDescripcion(v string) *comercializacionHidroDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithCantidad(v float64) *comercializacionHidroDetalleBuilder {
	// Notamos que el XSD pide 5 decimales para cantidad en este sector (totalDigits 20, fraction 5)
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithUnidadMedida(v int) *comercializacionHidroDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithPrecioUnitario(v float64) *comercializacionHidroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithMontoDescuento(v *float64) *comercializacionHidroDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercializacionHidroDetalleBuilder) WithSubTotal(v float64) *comercializacionHidroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercializacionHidroDetalleBuilder) Build() ComercializacionHidroDetalle {
	return ComercializacionHidroDetalle{models.NewRequestWrapper(b.detalle)}
}
