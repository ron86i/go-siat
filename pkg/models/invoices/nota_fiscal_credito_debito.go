package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
)

// NotaFiscalCreditoDebito representa la estructura completa de una nota de crédito/débito standard (Sector 24) lista para ser procesada.
type NotaFiscalCreditoDebito struct {
	models.RequestWrapper[documents.NotaFiscalCreditoDebito]
}

// NotaFiscalCreditoDebitoCabecera representa la sección de cabecera de la nota (Sector 24).
type NotaFiscalCreditoDebitoCabecera struct {
	models.RequestWrapper[documents.CabeceraNotaFiscalCreditoDebito]
}

// NotaDetalleFiscalCreditoDebito representa un ítem individual dentro del detalle (Sector 24).
type NotaDetalleFiscalCreditoDebito struct {
	models.RequestWrapper[documents.DetalleNotaFiscalCreditoDebito]
}

// NewNotaFiscalCreditoDebitoBuilder inicia el proceso de construcción de una Nota Fiscal.
func NewNotaFiscalCreditoDebitoBuilder() *notaFiscalCreditoDebitoBuilder {
	return &notaFiscalCreditoDebitoBuilder{
		nota: &documents.NotaFiscalCreditoDebito{
			XMLName:           xml.Name{Local: "notaFiscalElectronicaCreditoDebito"}, // Default a electrónica
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "notaFiscalElectronicaCreditoDebito.xsd",
		},
	}
}

// NewNotaFiscalCreditoDebitoCabeceraBuilder crea una nueva instancia del constructor para la cabecera (Sector 24).
func NewNotaFiscalCreditoDebitoCabeceraBuilder() *notaFiscalCreditoDebitoCabeceraBuilder {
	return &notaFiscalCreditoDebitoCabeceraBuilder{
		cabecera: &documents.CabeceraNotaFiscalCreditoDebito{
			CodigoDocumentoSector: 24, // Sector 24 para Nota Fiscal Crédito/Débito
		},
	}
}

// NewNotaDetalleFiscalCreditoDebitoBuilder crea una nueva instancia del constructor para el detalle (Sector 24).
func NewNotaDetalleFiscalCreditoDebitoBuilder() *notaDetalleFiscalCreditoDebitoBuilder {
	return &notaDetalleFiscalCreditoDebitoBuilder{
		detalle: &documents.DetalleNotaFiscalCreditoDebito{},
	}
}

type notaFiscalCreditoDebitoBuilder struct {
	nota *documents.NotaFiscalCreditoDebito
}

func (b *notaFiscalCreditoDebitoBuilder) WithCabecera(req NotaFiscalCreditoDebitoCabecera) *notaFiscalCreditoDebitoBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraNotaFiscalCreditoDebito](req); internal != nil {
		b.nota.Cabecera = *internal
	}
	return b
}

func (b *notaFiscalCreditoDebitoBuilder) AddDetalle(req NotaDetalleFiscalCreditoDebito) *notaFiscalCreditoDebitoBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleNotaFiscalCreditoDebito](req); internal != nil {
		b.nota.Detalle = append(b.nota.Detalle, *internal)
	}
	return b
}

func (b *notaFiscalCreditoDebitoBuilder) WithModalidad(tipo int) *notaFiscalCreditoDebitoBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.nota.XMLName = xml.Name{Local: "notaFiscalElectronicaCreditoDebito"}
		b.nota.XsiSchemaLocation = "notaFiscalElectronicaCreditoDebito.xsd"
	case siat.ModalidadComputarizada:
		b.nota.XMLName = xml.Name{Local: "notaFiscalComputarizadaCreditoDebito"}
		b.nota.XsiSchemaLocation = "notaFiscalComputarizadaCreditoDebito.xsd"
	}
	return b
}

func (b *notaFiscalCreditoDebitoBuilder) Build() NotaFiscalCreditoDebito {
	return NotaFiscalCreditoDebito{models.NewRequestWrapper(b.nota)}
}

type notaFiscalCreditoDebitoCabeceraBuilder struct {
	cabecera *documents.CabeceraNotaFiscalCreditoDebito
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithNitEmisor(v int64) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithRazonSocialEmisor(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithMunicipio(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithTelefono(v *string) *notaFiscalCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithNumeroNotaCreditoDebito(v int64) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroNotaCreditoDebito = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithCuf(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithCufd(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithCodigoSucursal(v int) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithDireccion(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithCodigoPuntoVenta(v *int) *notaFiscalCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithFechaEmision(v time.Time) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithNombreRazonSocial(v *string) *notaFiscalCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithNumeroDocumento(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithComplemento(v *string) *notaFiscalCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithCodigoCliente(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithNumeroFactura(v int64) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithNumeroAutorizacionCuf(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.NumeroAutorizacionCuf = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithFechaEmisionFactura(v time.Time) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.FechaEmisionFactura = datatype.NewTimeSiat(v)
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithMontoTotalOriginal(v float64) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.MontoTotalOriginal = datatype.Float64Round(v, 2)
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithMontoTotalDevuelto(v float64) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.MontoTotalDevuelto = datatype.Float64Round(v, 2)
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithMontoDescuentoCreditoDebito(v *float64) *notaFiscalCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 2)
		b.cabecera.MontoDescuentoCreditoDebito = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithMontoEfectivoCreditoDebito(v float64) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.MontoEfectivoCreditoDebito = datatype.Float64Round(v, 2)
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithCodigoExcepcion(v *int) *notaFiscalCreditoDebitoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithLeyenda(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) WithUsuario(v string) *notaFiscalCreditoDebitoCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *notaFiscalCreditoDebitoCabeceraBuilder) Build() NotaFiscalCreditoDebitoCabecera {
	return NotaFiscalCreditoDebitoCabecera{models.NewRequestWrapper(b.cabecera)}
}

type notaDetalleFiscalCreditoDebitoBuilder struct {
	detalle *documents.DetalleNotaFiscalCreditoDebito
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithActividadEconomica(v string) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithCodigoProductoSin(v int64) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithCodigoProducto(v string) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithDescripcion(v string) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithCantidad(v float64) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.Cantidad = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithUnidadMedida(v int) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithPrecioUnitario(v float64) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.PrecioUnitario = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithMontoDescuento(v *float64) *notaDetalleFiscalCreditoDebitoBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := datatype.Float64Round(*v, 10)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithSubTotal(v float64) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.SubTotal = datatype.Float64Round(v, 10)
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) WithCodigoDetalleTransaccion(v int) *notaDetalleFiscalCreditoDebitoBuilder {
	b.detalle.CodigoDetalleTransaccion = v
	return b
}

func (b *notaDetalleFiscalCreditoDebitoBuilder) Build() NotaDetalleFiscalCreditoDebito {
	return NotaDetalleFiscalCreditoDebito{models.NewRequestWrapper(b.detalle)}
}
