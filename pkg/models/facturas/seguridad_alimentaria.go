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

// SeguridadAlimentaria representa la estructura completa de una factura de Seguridad Alimentaria lista para ser procesada.
type SeguridadAlimentaria struct {
	models.RequestWrapper[documentos.FacturaSeguridadAlimentaria]
}

// SeguridadAlimentariaCabecera representa la sección de cabecera de la factura.
type SeguridadAlimentariaCabecera struct {
	models.RequestWrapper[documentos.CabeceraSeguridadAlimentaria]
}

// SeguridadAlimentariaDetalle representa un ítem individual dentro del detalle de la factura.
type SeguridadAlimentariaDetalle struct {
	models.RequestWrapper[documentos.DetalleSeguridadAlimentaria]
}

// NewSeguridadAlimentariaBuilder inicia el proceso de construcción de la factura.
func NewSeguridadAlimentariaBuilder() *seguridadAlimentariaBuilder {
	return &seguridadAlimentariaBuilder{
		factura: &documentos.FacturaSeguridadAlimentaria{
			XMLName:           xml.Name{Local: "facturaElectronicaSeguridadAlimentaria"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaSeguridadAlimentaria.xsd",
		},
	}
}

// NewSeguridadAlimentariaCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewSeguridadAlimentariaCabeceraBuilder() *seguridadAlimentariaCabeceraBuilder {
	return &seguridadAlimentariaCabeceraBuilder{
		cabecera: &documentos.CabeceraSeguridadAlimentaria{
			CodigoDocumentoSector: 7, // Sector 7
			MontoTotalSujetoIva:   0, // Fijo 0 según XSD
		},
	}
}

// NewSeguridadAlimentariaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewSeguridadAlimentariaDetalleBuilder() *seguridadAlimentariaDetalleBuilder {
	return &seguridadAlimentariaDetalleBuilder{
		detalle: &documentos.DetalleSeguridadAlimentaria{},
	}
}

type seguridadAlimentariaBuilder struct {
	factura *documentos.FacturaSeguridadAlimentaria
}

func (b *seguridadAlimentariaBuilder) WithCabecera(req SeguridadAlimentariaCabecera) *seguridadAlimentariaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraSeguridadAlimentaria](req); internal != nil {
		b.factura.Cabecera = *internal
		b.factura.Cabecera.MontoTotalSujetoIva = 0 // Garantizamos que sea 0
	}
	return b
}

func (b *seguridadAlimentariaBuilder) AddDetalle(req SeguridadAlimentariaDetalle) *seguridadAlimentariaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleSeguridadAlimentaria](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *seguridadAlimentariaBuilder) WithModalidad(tipo int) *seguridadAlimentariaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaSeguridadAlimentaria"}
		b.factura.XsiSchemaLocation = "facturaElectronicaSeguridadAlimentaria.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaSeguridadAlimentaria"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaSeguridadAlimentaria.xsd"
	}
	return b
}

func (b *seguridadAlimentariaBuilder) Build() SeguridadAlimentaria {
	return SeguridadAlimentaria{models.NewRequestWrapper(b.factura)}
}

type seguridadAlimentariaCabeceraBuilder struct {
	cabecera *documentos.CabeceraSeguridadAlimentaria
}

func (b *seguridadAlimentariaCabeceraBuilder) WithNitEmisor(v int64) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithRazonSocialEmisor(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithMunicipio(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithTelefono(telefono *string) *seguridadAlimentariaCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithNumeroFactura(v int64) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCuf(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCufd(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoSucursal(v int) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithDireccion(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithFechaEmision(v time.Time) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithNombreRazonSocial(v *string) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithNumeroDocumento(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithComplemento(v *string) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoCliente(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoMetodoPago(v int) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithNumeroTarjeta(v *int64) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithMontoTotal(v float64) *seguridadAlimentariaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoMoneda(v int) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithTipoCambio(v float64) *seguridadAlimentariaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithMontoTotalMoneda(v float64) *seguridadAlimentariaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithMontoGiftCard(v *float64) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithDescuentoAdicional(v *float64) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoExcepcion(v *int) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCafc(v *string) *seguridadAlimentariaCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithLeyenda(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithUsuario(v string) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) WithCodigoDocumentoSector(v int) *seguridadAlimentariaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *seguridadAlimentariaCabeceraBuilder) Build() SeguridadAlimentariaCabecera {
	return SeguridadAlimentariaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type seguridadAlimentariaDetalleBuilder struct {
	detalle *documentos.DetalleSeguridadAlimentaria
}

func (b *seguridadAlimentariaDetalleBuilder) WithActividadEconomica(v string) *seguridadAlimentariaDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithCodigoProductoSin(v int64) *seguridadAlimentariaDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithCodigoProducto(v string) *seguridadAlimentariaDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithDescripcion(v string) *seguridadAlimentariaDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithCantidad(v float64) *seguridadAlimentariaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithUnidadMedida(v int) *seguridadAlimentariaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithPrecioUnitario(v float64) *seguridadAlimentariaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithMontoDescuento(v *float64) *seguridadAlimentariaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) WithSubTotal(v float64) *seguridadAlimentariaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *seguridadAlimentariaDetalleBuilder) Build() SeguridadAlimentariaDetalle {
	return SeguridadAlimentariaDetalle{models.NewRequestWrapper(b.detalle)}
}
