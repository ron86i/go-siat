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

// ZonaFranca representa la estructura completa de una factura Zona Franca lista para ser procesada.
type ZonaFranca struct {
	models.RequestWrapper[documents.FacturaZonaFranca]
}

// ZonaFrancaCabecera representa la sección de cabecera de una factura Zona Franca.
type ZonaFrancaCabecera struct {
	models.RequestWrapper[documents.CabeceraZonaFranca]
}

// ZonaFrancaDetalle representa un ítem individual dentro del detalle de una factura Zona Franca.
type ZonaFrancaDetalle struct {
	models.RequestWrapper[documents.DetalleZonaFranca]
}

// NewZonaFrancaBuilder inicia el proceso de construcción de una Factura Zona Franca.
func NewZonaFrancaBuilder() *zonaFrancaBuilder {
	return &zonaFrancaBuilder{
		factura: &documents.FacturaZonaFranca{
			XMLName:           xml.Name{Local: "facturaElectronicaZonaFranca"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaZonaFranca.xsd",
		},
	}
}

// NewZonaFrancaCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewZonaFrancaCabeceraBuilder() *zonaFrancaCabeceraBuilder {
	return &zonaFrancaCabeceraBuilder{
		cabecera: &documents.CabeceraZonaFranca{
			CodigoDocumentoSector: 5, // Sector 5 para Zona Franca
			MontoTotalSujetoIva:   0, // Siempre 0 para Zona Franca
		},
	}
}

// NewZonaFrancaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewZonaFrancaDetalleBuilder() *zonaFrancaDetalleBuilder {
	return &zonaFrancaDetalleBuilder{
		detalle: &documents.DetalleZonaFranca{},
	}
}

type zonaFrancaBuilder struct {
	factura *documents.FacturaZonaFranca
}

func (b *zonaFrancaBuilder) WithCabecera(req ZonaFrancaCabecera) *zonaFrancaBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraZonaFranca](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *zonaFrancaBuilder) AddDetalle(req ZonaFrancaDetalle) *zonaFrancaBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleZonaFranca](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *zonaFrancaBuilder) WithModalidad(tipo int) *zonaFrancaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaZonaFranca"}
		b.factura.XsiSchemaLocation = "facturaElectronicaZonaFranca.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaZonaFranca"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaZonaFranca.xsd"
	}
	return b
}

func (b *zonaFrancaBuilder) Build() ZonaFranca {
	return ZonaFranca{models.NewRequestWrapper(b.factura)}
}

type zonaFrancaCabeceraBuilder struct {
	cabecera *documents.CabeceraZonaFranca
}

func (b *zonaFrancaCabeceraBuilder) WithNitEmisor(v int64) *zonaFrancaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithRazonSocialEmisor(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMunicipio(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithTelefono(telefono *string) *zonaFrancaCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroFactura(v int64) *zonaFrancaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCuf(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCufd(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoSucursal(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithDireccion(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *zonaFrancaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNombreRazonSocial(v *string) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroDocumento(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithComplemento(v *string) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoCliente(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoMetodoPago(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroTarjeta(v *int64) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMontoTotal(v float64) *zonaFrancaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoMoneda(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithTipoCambio(v float64) *zonaFrancaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMontoTotalMoneda(v float64) *zonaFrancaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithNumeroParteRecepcion(v *string) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroParteRecepcion = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroParteRecepcion = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMontoGiftCard(v *float64) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithDescuentoAdicional(v *float64) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCodigoExcepcion(v *int) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithCafc(v *string) *zonaFrancaCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithLeyenda(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithUsuario(v string) *zonaFrancaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *zonaFrancaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *zonaFrancaCabeceraBuilder) WithCodigoDocumentoSector(v int) *zonaFrancaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *zonaFrancaCabeceraBuilder) Build() ZonaFrancaCabecera {
	return ZonaFrancaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type zonaFrancaDetalleBuilder struct {
	detalle *documents.DetalleZonaFranca
}

func (b *zonaFrancaDetalleBuilder) WithActividadEconomica(v string) *zonaFrancaDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithCodigoProductoSin(v int64) *zonaFrancaDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithCodigoProducto(v string) *zonaFrancaDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithDescripcion(v string) *zonaFrancaDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithCantidad(v float64) *zonaFrancaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithUnidadMedida(v int) *zonaFrancaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithPrecioUnitario(v float64) *zonaFrancaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *zonaFrancaDetalleBuilder) WithMontoDescuento(v *float64) *zonaFrancaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *zonaFrancaDetalleBuilder) WithSubTotal(v float64) *zonaFrancaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *zonaFrancaDetalleBuilder) Build() ZonaFrancaDetalle {
	return ZonaFrancaDetalle{models.NewRequestWrapper(b.detalle)}
}
