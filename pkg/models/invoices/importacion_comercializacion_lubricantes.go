package invoices

import (
	"encoding/xml"
	"time"

	"github.com/shopspring/decimal"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
)

// ImportacionComercializacionLubricantes representa la estructura completa de una factura de Sector 44.
type ImportacionComercializacionLubricantes struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaImportacionComercializacionLubricantes]
}

// ImportacionComercializacionLubricantesCabecera representa la sección de cabecera de Sector 44.
type ImportacionComercializacionLubricantesCabecera struct {
	models.RequestWrapper[documents.CabeceraImportacionComercializacionLubricantes]
}

// ImportacionComercializacionLubricantesDetalle representa un ítem individual de Sector 44.
type ImportacionComercializacionLubricantesDetalle struct {
	models.RequestWrapper[documents.DetalleImportacionComercializacionLubricantes]
}

// NewImportacionComercializacionLubricantesBuilder inicia el proceso de construcción.
func NewImportacionComercializacionLubricantesBuilder() *importacionComercializacionLubricantesBuilder {
	return &importacionComercializacionLubricantesBuilder{
		metadatos: models.NuevosMetadatosFactura(44),
		factura: &documents.FacturaImportacionComercializacionLubricantes{
			XMLName:           xml.Name{Local: "facturaElectronicaImportacionComercializacionLubricantes"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaImportacionComercializacionLubricantes.xsd",
		},
	}
}

// NewImportacionComercializacionLubricantesCabeceraBuilder crea el constructor para la cabecera.
func NewImportacionComercializacionLubricantesCabeceraBuilder() *importacionComercializacionLubricantesCabeceraBuilder {
	return &importacionComercializacionLubricantesCabeceraBuilder{
		cabecera: &documents.CabeceraImportacionComercializacionLubricantes{
			CodigoDocumentoSector: 44,
		},
	}
}

// NewImportacionComercializacionLubricantesDetalleBuilder crea el constructor para los ítems.
func NewImportacionComercializacionLubricantesDetalleBuilder() *importacionComercializacionLubricantesDetalleBuilder {
	return &importacionComercializacionLubricantesDetalleBuilder{
		detalle: &documents.DetalleImportacionComercializacionLubricantes{},
	}
}

type importacionComercializacionLubricantesBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaImportacionComercializacionLubricantes
}

func (b *importacionComercializacionLubricantesBuilder) WithCabecera(req ImportacionComercializacionLubricantesCabecera) *importacionComercializacionLubricantesBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraImportacionComercializacionLubricantes](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *importacionComercializacionLubricantesBuilder) AddDetalle(req ImportacionComercializacionLubricantesDetalle) *importacionComercializacionLubricantesBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleImportacionComercializacionLubricantes](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *importacionComercializacionLubricantesBuilder) WithModalidad(tipo int) *importacionComercializacionLubricantesBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaImportacionComercializacionLubricantes"}
		b.factura.XsiSchemaLocation = "facturaElectronicaImportacionComercializacionLubricantes.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaImportacionComercializacionLubricantes"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaImportacionComercializacionLubricantes.xsd"
	}
	return b
}

func (b *importacionComercializacionLubricantesBuilder) Build() ImportacionComercializacionLubricantes {
	return ImportacionComercializacionLubricantes{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type importacionComercializacionLubricantesCabeceraBuilder struct {
	cabecera *documents.CabeceraImportacionComercializacionLubricantes
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) Build() ImportacionComercializacionLubricantesCabecera {
	return ImportacionComercializacionLubricantesCabecera{models.NewRequestWrapper(b.cabecera)}
}

type importacionComercializacionLubricantesDetalleBuilder struct {
	detalle *documents.DetalleImportacionComercializacionLubricantes
}

func (b *importacionComercializacionLubricantesDetalleBuilder) Build() ImportacionComercializacionLubricantesDetalle {
	return ImportacionComercializacionLubricantesDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNitEmisor(v int64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithRazonSocialEmisor(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMunicipio(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithTelefono(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNumeroFactura(v int64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCuf(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCufd(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoSucursal(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithDireccion(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoPuntoVenta(v *int) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithFechaEmision(v time.Time) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNombreRazonSocial(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNumeroDocumento(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithComplemento(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoCliente(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCiudad(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.Ciudad = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Ciudad = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNombrePropietario(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombrePropietario = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNombreRepresentanteLegal(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCondicionPago(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.CondicionPago = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.CondicionPago = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithPeriodoEntrega(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.PeriodoEntrega = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoMetodoPago(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNumeroTarjeta(v *int64) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMontoDeduccionIehdDS25530(v *float64) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDeduccionIehdDS25530 = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.MontoDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithSubTotalNeto(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithTotalPrecioNeto(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithIva(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMontoAPagar(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoMoneda(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithTipoCambio(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMontoTotalMoneda(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithDescuentoAdicional(v *float64) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoExcepcion(v *int) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithLeyenda(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithUsuario(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoInterno(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoDocumentoSector(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithActividadEconomica(v string) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithCodigoProductoSin(v int64) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithCodigoProducto(v string) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithDescripcion(v string) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithCantidad(v float64) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithUnidadMedida(v int) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithPrecioNetoUnitario(v float64) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithMontoDescuento(v *float64) *importacionComercializacionLubricantesDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithSubTotal(v float64) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithCantidadLitros(v float64) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.CantidadLitros = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithPorcentajeDeduccionIehdDS25530(v *float64) *importacionComercializacionLubricantesDetalleBuilder {
	if v == nil {
		b.detalle.PorcentajeDeduccionIehdDS25530 = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.PorcentajeDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: &value}
	return b
}
