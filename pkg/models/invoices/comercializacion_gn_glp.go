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

// ComercializacionGnGlp representa la estructura completa de una factura de Comercialización de GN y GLP lista para ser procesada.
type ComercializacionGnGlp struct {
	models.RequestWrapper[documents.FacturaComercializacionGnGlp]
}

// ComercializacionGnGlpCabecera representa la sección de cabecera de una factura de Comercialización de GN y GLP.
type ComercializacionGnGlpCabecera struct {
	models.RequestWrapper[documents.CabeceraComercializacionGnGlp]
}

// ComercializacionGnGlpDetalle representa un ítem individual dentro del detalle de una factura de Comercialización de GN y GLP.
type ComercializacionGnGlpDetalle struct {
	models.RequestWrapper[documents.DetalleComercializacionGnGlp]
}

// NewComercializacionGnGlpBuilder inicia el proceso de construcción de una Factura de Comercialización de GN y GLP.
func NewComercializacionGnGlpBuilder() *comercializacionGnGlpBuilder {
	return &comercializacionGnGlpBuilder{
		factura: &documents.FacturaComercializacionGnGlp{
			XMLName:           xml.Name{Local: "facturaElectronicaComercializacionGnGlp"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercializacionGnGlp.xsd",
		},
	}
}

// NewComercializacionGnGlpCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewComercializacionGnGlpCabeceraBuilder() *comercializacionGnGlpCabeceraBuilder {
	return &comercializacionGnGlpCabeceraBuilder{
		cabecera: &documents.CabeceraComercializacionGnGlp{
			CodigoDocumentoSector: 39, // Sector 39 para Comercialización de GN y GLP
		},
	}
}

// NewComercializacionGnGlpDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewComercializacionGnGlpDetalleBuilder() *comercializacionGnGlpDetalleBuilder {
	return &comercializacionGnGlpDetalleBuilder{
		detalle: &documents.DetalleComercializacionGnGlp{},
	}
}

type comercializacionGnGlpBuilder struct {
	factura *documents.FacturaComercializacionGnGlp
}

func (b *comercializacionGnGlpBuilder) WithCabecera(req ComercializacionGnGlpCabecera) *comercializacionGnGlpBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraComercializacionGnGlp](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercializacionGnGlpBuilder) AddDetalle(req ComercializacionGnGlpDetalle) *comercializacionGnGlpBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleComercializacionGnGlp](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *comercializacionGnGlpBuilder) WithModalidad(tipo int) *comercializacionGnGlpBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercializacionGnGlp"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercializacionGnGlp.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercializacionGnGlp"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercializacionGnGlp.xsd"
	}
	return b
}

func (b *comercializacionGnGlpBuilder) Build() ComercializacionGnGlp {
	return ComercializacionGnGlp{models.NewRequestWrapper(b.factura)}
}

type comercializacionGnGlpCabeceraBuilder struct {
	cabecera *documents.CabeceraComercializacionGnGlp
}

func (b *comercializacionGnGlpCabeceraBuilder) WithNitEmisor(v int64) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithRazonSocialEmisor(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithMunicipio(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithTelefono(telefono *string) *comercializacionGnGlpCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithNumeroFactura(v int64) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCuf(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCufd(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoSucursal(v int) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithDireccion(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercializacionGnGlpCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithFechaEmision(v time.Time) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithNombreRazonSocial(v *string) *comercializacionGnGlpCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithNumeroDocumento(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithComplemento(v *string) *comercializacionGnGlpCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoCliente(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoMetodoPago(v int) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercializacionGnGlpCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithMontoTotal(v float64) *comercializacionGnGlpCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *comercializacionGnGlpCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoMoneda(v int) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithTipoCambio(v float64) *comercializacionGnGlpCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercializacionGnGlpCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercializacionGnGlpCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoExcepcion(v *int) *comercializacionGnGlpCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCafc(v *string) *comercializacionGnGlpCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithLeyenda(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithUsuario(v string) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercializacionGnGlpCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *comercializacionGnGlpCabeceraBuilder) Build() ComercializacionGnGlpCabecera {
	return ComercializacionGnGlpCabecera{models.NewRequestWrapper(b.cabecera)}
}

type comercializacionGnGlpDetalleBuilder struct {
	detalle *documents.DetalleComercializacionGnGlp
}

func (b *comercializacionGnGlpDetalleBuilder) WithActividadEconomica(v string) *comercializacionGnGlpDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithCodigoProductoSin(v int64) *comercializacionGnGlpDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithCodigoProducto(v string) *comercializacionGnGlpDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithDescripcion(v string) *comercializacionGnGlpDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithCantidad(v float64) *comercializacionGnGlpDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithUnidadMedida(v int) *comercializacionGnGlpDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithPrecioUnitario(v float64) *comercializacionGnGlpDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithMontoDescuento(v *float64) *comercializacionGnGlpDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) WithSubTotal(v float64) *comercializacionGnGlpDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercializacionGnGlpDetalleBuilder) Build() ComercializacionGnGlpDetalle {
	return ComercializacionGnGlpDetalle{models.NewRequestWrapper(b.detalle)}
}
