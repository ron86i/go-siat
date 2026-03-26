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

// ComercializacionGnv representa la estructura completa de una factura de Comercialización de GNV lista para ser procesada.
type ComercializacionGnv struct {
	models.RequestWrapper[documents.FacturaComercializacionGnv]
}

// ComercializacionGnvCabecera representa la sección de cabecera de una factura de GNV.
type ComercializacionGnvCabecera struct {
	models.RequestWrapper[documents.CabeceraComercializacionGnv]
}

// ComercializacionGnvDetalle representa un ítem individual dentro del detalle de una factura de GNV.
type ComercializacionGnvDetalle struct {
	models.RequestWrapper[documents.DetalleComercializacionGnv]
}

// NewComercializacionGnvBuilder inicia el proceso de construcción de una Factura de Comercialización de GNV.
func NewComercializacionGnvBuilder() *comercializacionGnvBuilder {
	return &comercializacionGnvBuilder{
		factura: &documents.FacturaComercializacionGnv{
			XMLName:           xml.Name{Local: "facturaElectronicaComercializacionGnv"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercializacionGnv.xsd",
		},
	}
}

// NewComercializacionGnvCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewComercializacionGnvCabeceraBuilder() *comercializacionGnvCabeceraBuilder {
	return &comercializacionGnvCabeceraBuilder{
		cabecera: &documents.CabeceraComercializacionGnv{
			CodigoDocumentoSector: 37, // Sector 37 para GNV
		},
	}
}

// NewComercializacionGnvDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewComercializacionGnvDetalleBuilder() *comercializacionGnvDetalleBuilder {
	return &comercializacionGnvDetalleBuilder{
		detalle: &documents.DetalleComercializacionGnv{},
	}
}

type comercializacionGnvBuilder struct {
	factura *documents.FacturaComercializacionGnv
}

func (b *comercializacionGnvBuilder) WithCabecera(req ComercializacionGnvCabecera) *comercializacionGnvBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraComercializacionGnv](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercializacionGnvBuilder) AddDetalle(req ComercializacionGnvDetalle) *comercializacionGnvBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleComercializacionGnv](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *comercializacionGnvBuilder) WithModalidad(tipo int) *comercializacionGnvBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercializacionGnv"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercializacionGnv.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercializacionGnv"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercializacionGnv.xsd"
	}
	return b
}

func (b *comercializacionGnvBuilder) Build() ComercializacionGnv {
	return ComercializacionGnv{models.NewRequestWrapper(b.factura)}
}

type comercializacionGnvCabeceraBuilder struct {
	cabecera *documents.CabeceraComercializacionGnv
}

func (b *comercializacionGnvCabeceraBuilder) WithNitEmisor(v int64) *comercializacionGnvCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithRazonSocialEmisor(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithMunicipio(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithTelefono(telefono *string) *comercializacionGnvCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithNumeroFactura(v int64) *comercializacionGnvCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCuf(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCufd(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoSucursal(v int) *comercializacionGnvCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithDireccion(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *comercializacionGnvCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithNombreRazonSocial(v *string) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercializacionGnvCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithNumeroDocumento(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithComplemento(v *string) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoCliente(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoPais(v *int) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPais = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPais = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithPlacaVehiculo(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.PlacaVehiculo = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithTipoEnvase(v *string) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.TipoEnvase = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.TipoEnvase = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoMetodoPago(v int) *comercializacionGnvCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithMontoTotal(v float64) *comercializacionGnvCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *comercializacionGnvCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoMoneda(v int) *comercializacionGnvCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithTipoCambio(v float64) *comercializacionGnvCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercializacionGnvCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithMontoVale(v *float64) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoVale = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoVale = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCodigoExcepcion(v *int) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithCafc(v *string) *comercializacionGnvCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithLeyenda(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) WithUsuario(v string) *comercializacionGnvCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *comercializacionGnvCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercializacionGnvCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *comercializacionGnvCabeceraBuilder) Build() ComercializacionGnvCabecera {
	return ComercializacionGnvCabecera{models.NewRequestWrapper(b.cabecera)}
}

type comercializacionGnvDetalleBuilder struct {
	detalle *documents.DetalleComercializacionGnv
}

func (b *comercializacionGnvDetalleBuilder) WithActividadEconomica(v string) *comercializacionGnvDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithCodigoProductoSin(v int64) *comercializacionGnvDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithCodigoProducto(v string) *comercializacionGnvDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithDescripcion(v string) *comercializacionGnvDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithCantidad(v float64) *comercializacionGnvDetalleBuilder {
	// GNV usa 5 decimales para cantidad
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithUnidadMedida(v int) *comercializacionGnvDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithPrecioUnitario(v float64) *comercializacionGnvDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithMontoDescuento(v *float64) *comercializacionGnvDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercializacionGnvDetalleBuilder) WithSubTotal(v float64) *comercializacionGnvDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercializacionGnvDetalleBuilder) Build() ComercializacionGnvDetalle {
	return ComercializacionGnvDetalle{models.NewRequestWrapper(b.detalle)}
}
