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

// JuegoAzar representa la estructura completa de una factura de Juego de Azar lista para ser procesada.
type JuegoAzar struct {
	models.RequestWrapper[documentos.FacturaJuegoAzar]
}

// JuegoAzarCabecera representa la sección de cabecera de una factura de Juego de Azar.
type JuegoAzarCabecera struct {
	models.RequestWrapper[documentos.CabeceraJuegoAzar]
}

// JuegoAzarDetalle representa un ítem individual dentro del detalle de una factura de Juego de Azar.
type JuegoAzarDetalle struct {
	models.RequestWrapper[documentos.DetalleJuegoAzar]
}

// NewJuegoAzarBuilder inicia el proceso de construcción de una Factura de Juego de Azar.
func NewJuegoAzarBuilder() *juegoAzarBuilder {
	return &juegoAzarBuilder{
		factura: &documentos.FacturaJuegoAzar{
			XMLName:           xml.Name{Local: "facturaElectronicaJuegoAzar"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaJuegoAzar.xsd",
		},
	}
}

// NewJuegoAzarCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewJuegoAzarCabeceraBuilder() *juegoAzarCabeceraBuilder {
	return &juegoAzarCabeceraBuilder{
		cabecera: &documentos.CabeceraJuegoAzar{
			CodigoDocumentoSector: 18, // Sector 18 para Juego de Azar
		},
	}
}

// NewJuegoAzarDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewJuegoAzarDetalleBuilder() *juegoAzarDetalleBuilder {
	return &juegoAzarDetalleBuilder{
		detalle: &documentos.DetalleJuegoAzar{},
	}
}

type juegoAzarBuilder struct {
	factura *documentos.FacturaJuegoAzar
}

func (b *juegoAzarBuilder) WithCabecera(req JuegoAzarCabecera) *juegoAzarBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraJuegoAzar](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *juegoAzarBuilder) AddDetalle(req JuegoAzarDetalle) *juegoAzarBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleJuegoAzar](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *juegoAzarBuilder) WithModalidad(tipo int) *juegoAzarBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaJuegoAzar"}
		b.factura.XsiSchemaLocation = "facturaElectronicaJuegoAzar.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaJuegoAzar"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaJuegoAzar.xsd"
	}
	return b
}

func (b *juegoAzarBuilder) Build() JuegoAzar {
	return JuegoAzar{models.NewRequestWrapper(b.factura)}
}

type juegoAzarCabeceraBuilder struct {
	cabecera *documentos.CabeceraJuegoAzar
}

func (b *juegoAzarCabeceraBuilder) WithNitEmisor(v int64) *juegoAzarCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithRazonSocialEmisor(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithMunicipio(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithTelefono(telefono *string) *juegoAzarCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithNumeroFactura(v int64) *juegoAzarCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCuf(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCufd(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoSucursal(v int) *juegoAzarCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithDireccion(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoPuntoVenta(v *int) *juegoAzarCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithFechaEmision(v time.Time) *juegoAzarCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *juegoAzarCabeceraBuilder) WithNombreRazonSocial(v *string) *juegoAzarCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *juegoAzarCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithNumeroDocumento(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithComplemento(v *string) *juegoAzarCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoCliente(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoMetodoPago(v int) *juegoAzarCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithNumeroTarjeta(v *int64) *juegoAzarCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithMontoTotal(v float64) *juegoAzarCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithMontoTotalIj(v float64) *juegoAzarCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalIj = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithMontoTotalSujetoIpj(v float64) *juegoAzarCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIpj = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *juegoAzarCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoMoneda(v int) *juegoAzarCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithTipoCambio(v float64) *juegoAzarCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithMontoTotalMoneda(v float64) *juegoAzarCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithDescuentoAdicional(v *float64) *juegoAzarCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoExcepcion(v *int) *juegoAzarCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCafc(v *string) *juegoAzarCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *juegoAzarCabeceraBuilder) WithLeyenda(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithUsuario(v string) *juegoAzarCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *juegoAzarCabeceraBuilder) WithCodigoDocumentoSector(v int) *juegoAzarCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *juegoAzarCabeceraBuilder) Build() JuegoAzarCabecera {
	return JuegoAzarCabecera{models.NewRequestWrapper(b.cabecera)}
}

type juegoAzarDetalleBuilder struct {
	detalle *documentos.DetalleJuegoAzar
}

func (b *juegoAzarDetalleBuilder) WithActividadEconomica(v string) *juegoAzarDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *juegoAzarDetalleBuilder) WithCodigoProductoSin(v int64) *juegoAzarDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *juegoAzarDetalleBuilder) WithCodigoProducto(v string) *juegoAzarDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *juegoAzarDetalleBuilder) WithDescripcion(v string) *juegoAzarDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *juegoAzarDetalleBuilder) WithCantidad(v float64) *juegoAzarDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *juegoAzarDetalleBuilder) WithUnidadMedida(v int) *juegoAzarDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *juegoAzarDetalleBuilder) WithPrecioUnitario(v float64) *juegoAzarDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *juegoAzarDetalleBuilder) WithMontoDescuento(v *float64) *juegoAzarDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *juegoAzarDetalleBuilder) WithSubTotal(v float64) *juegoAzarDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *juegoAzarDetalleBuilder) Build() JuegoAzarDetalle {
	return JuegoAzarDetalle{models.NewRequestWrapper(b.detalle)}
}
