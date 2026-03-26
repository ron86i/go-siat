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

// TasaCero representa la estructura completa de una factura Tasa Cero lista para ser procesada.
type TasaCero struct {
	models.RequestWrapper[documents.FacturaTasaCero]
}

// TasaCeroCabecera representa la sección de cabecera de una factura Tasa Cero.
type TasaCeroCabecera struct {
	models.RequestWrapper[documents.CabeceraTasaCero]
}

// TasaCeroDetalle representa un ítem individual dentro del detalle de una factura Tasa Cero.
type TasaCeroDetalle struct {
	models.RequestWrapper[documents.DetalleTasaCero]
}

// NewTasaCeroBuilder inicia el proceso de construcción de una Factura Tasa Cero.
func NewTasaCeroBuilder() *tasaCeroBuilder {
	return &tasaCeroBuilder{
		factura: &documents.FacturaTasaCero{
			XMLName:           xml.Name{Local: "facturaElectronicaTasaCero"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaTasaCero.xsd",
		},
	}
}

// NewTasaCeroCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewTasaCeroCabeceraBuilder() *tasaCeroCabeceraBuilder {
	return &tasaCeroCabeceraBuilder{
		cabecera: &documents.CabeceraTasaCero{
			CodigoDocumentoSector: 8, // Sector 8 para Tasa Cero
			MontoTotalSujetoIva:   0, // Siempre 0 para Tasa Cero
		},
	}
}

// NewTasaCeroDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewTasaCeroDetalleBuilder() *tasaCeroDetalleBuilder {
	return &tasaCeroDetalleBuilder{
		detalle: &documents.DetalleTasaCero{},
	}
}

type tasaCeroBuilder struct {
	factura *documents.FacturaTasaCero
}

func (b *tasaCeroBuilder) WithCabecera(req TasaCeroCabecera) *tasaCeroBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraTasaCero](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *tasaCeroBuilder) AddDetalle(req TasaCeroDetalle) *tasaCeroBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleTasaCero](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *tasaCeroBuilder) WithModalidad(tipo int) *tasaCeroBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaTasaCero"}
		b.factura.XsiSchemaLocation = "facturaElectronicaTasaCero.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaTasaCero"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaTasaCero.xsd"
	}
	return b
}

func (b *tasaCeroBuilder) Build() TasaCero {
	return TasaCero{models.NewRequestWrapper(b.factura)}
}

type tasaCeroCabeceraBuilder struct {
	cabecera *documents.CabeceraTasaCero
}

func (b *tasaCeroCabeceraBuilder) WithNitEmisor(v int64) *tasaCeroCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithRazonSocialEmisor(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithMunicipio(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithTelefono(telefono *string) *tasaCeroCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithNumeroFactura(v int64) *tasaCeroCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCuf(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCufd(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCodigoSucursal(v int) *tasaCeroCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithDireccion(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCodigoPuntoVenta(v *int) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *tasaCeroCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *tasaCeroCabeceraBuilder) WithNombreRazonSocial(v *string) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *tasaCeroCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithNumeroDocumento(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithComplemento(v *string) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCodigoCliente(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCodigoMetodoPago(v int) *tasaCeroCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithNumeroTarjeta(v *int64) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithMontoTotal(v float64) *tasaCeroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCodigoMoneda(v int) *tasaCeroCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithTipoCambio(v float64) *tasaCeroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithMontoTotalMoneda(v float64) *tasaCeroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithMontoGiftCard(v *float64) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithDescuentoAdicional(v *float64) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCodigoExcepcion(v *int) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithCafc(v *string) *tasaCeroCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *tasaCeroCabeceraBuilder) WithLeyenda(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithUsuario(v string) *tasaCeroCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *tasaCeroCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *tasaCeroCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *tasaCeroCabeceraBuilder) WithCodigoDocumentoSector(v int) *tasaCeroCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *tasaCeroCabeceraBuilder) Build() TasaCeroCabecera {
	return TasaCeroCabecera{models.NewRequestWrapper(b.cabecera)}
}

type tasaCeroDetalleBuilder struct {
	detalle *documents.DetalleTasaCero
}

func (b *tasaCeroDetalleBuilder) WithActividadEconomica(v string) *tasaCeroDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *tasaCeroDetalleBuilder) WithCodigoProductoSin(v int64) *tasaCeroDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *tasaCeroDetalleBuilder) WithCodigoProducto(v string) *tasaCeroDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *tasaCeroDetalleBuilder) WithDescripcion(v string) *tasaCeroDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *tasaCeroDetalleBuilder) WithCantidad(v float64) *tasaCeroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *tasaCeroDetalleBuilder) WithUnidadMedida(v int) *tasaCeroDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *tasaCeroDetalleBuilder) WithPrecioUnitario(v float64) *tasaCeroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *tasaCeroDetalleBuilder) WithMontoDescuento(v *float64) *tasaCeroDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *tasaCeroDetalleBuilder) WithSubTotal(v float64) *tasaCeroDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *tasaCeroDetalleBuilder) Build() TasaCeroDetalle {
	return TasaCeroDetalle{models.NewRequestWrapper(b.detalle)}
}
