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

// ImportacionComercializacionLubricantes representa la estructura completa de una factura de Sector 44.
type ImportacionComercializacionLubricantes struct {
	models.RequestWrapper[documentos.FacturaImportacionComercializacionLubricantes]
}

// ImportacionComercializacionLubricantesCabecera representa la sección de cabecera de Sector 44.
type ImportacionComercializacionLubricantesCabecera struct {
	models.RequestWrapper[documentos.CabeceraImportacionComercializacionLubricantes]
}

// ImportacionComercializacionLubricantesDetalle representa un ítem individual de Sector 44.
type ImportacionComercializacionLubricantesDetalle struct {
	models.RequestWrapper[documentos.DetalleImportacionComercializacionLubricantes]
}

// NewImportacionComercializacionLubricantesBuilder inicia el proceso de construcción.
func NewImportacionComercializacionLubricantesBuilder() *importacionComercializacionLubricantesBuilder {
	return &importacionComercializacionLubricantesBuilder{
		factura: &documentos.FacturaImportacionComercializacionLubricantes{
			XMLName:  xml.Name{Local: "facturaElectronicaImportacionComercializacionLubricantes"},
			XmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		},
	}
}

// NewImportacionComercializacionLubricantesCabeceraBuilder crea el constructor para la cabecera.
func NewImportacionComercializacionLubricantesCabeceraBuilder() *importacionComercializacionLubricantesCabeceraBuilder {
	return &importacionComercializacionLubricantesCabeceraBuilder{
		cabecera: &documentos.CabeceraImportacionComercializacionLubricantes{
			CodigoDocumentoSector: 44,
		},
	}
}

// NewImportacionComercializacionLubricantesDetalleBuilder crea el constructor para los ítems.
func NewImportacionComercializacionLubricantesDetalleBuilder() *importacionComercializacionLubricantesDetalleBuilder {
	return &importacionComercializacionLubricantesDetalleBuilder{
		detalle: &documentos.DetalleImportacionComercializacionLubricantes{},
	}
}

type importacionComercializacionLubricantesBuilder struct {
	factura *documentos.FacturaImportacionComercializacionLubricantes
}

func (b *importacionComercializacionLubricantesBuilder) WithCabecera(req ImportacionComercializacionLubricantesCabecera) *importacionComercializacionLubricantesBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraImportacionComercializacionLubricantes](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *importacionComercializacionLubricantesBuilder) AddDetalle(req ImportacionComercializacionLubricantesDetalle) *importacionComercializacionLubricantesBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleImportacionComercializacionLubricantes](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *importacionComercializacionLubricantesBuilder) WithModalidad(tipo int) *importacionComercializacionLubricantesBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaImportacionComercializacionLubricantes"}
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaImportacionComercializacionLubricantes"}
	}
	return b
}

func (b *importacionComercializacionLubricantesBuilder) Build() ImportacionComercializacionLubricantes {
	return ImportacionComercializacionLubricantes{models.NewRequestWrapper(b.factura)}
}

type importacionComercializacionLubricantesCabeceraBuilder struct {
	cabecera *documentos.CabeceraImportacionComercializacionLubricantes
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

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithTelefono(telefono *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *telefono
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
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
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithFechaEmision(v time.Time) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNombreRazonSocial(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
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

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithComplemento(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoCliente(v string) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCiudad(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNombrePropietario(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombrePropietario = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNombreRepresentanteLegal(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRepresentanteLegal = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCondicionPago(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.CondicionPago = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.CondicionPago = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithPeriodoEntrega(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.PeriodoEntrega = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoMetodoPago(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithNumeroTarjeta(v *int64) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMontoTotal(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMontoDeduccionIehdDS25530(v *float64) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
		b.cabecera.MontoDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoMoneda(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithTipoCambio(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithMontoTotalMoneda(v float64) *importacionComercializacionLubricantesCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithDescuentoAdicional(v *float64) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoExcepcion(v *int) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCafc(v *string) *importacionComercializacionLubricantesCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	}
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

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *importacionComercializacionLubricantesCabeceraBuilder) WithCodigoDocumentoSector(v int) *importacionComercializacionLubricantesCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *importacionComercializacionLubricantesCabeceraBuilder) Build() ImportacionComercializacionLubricantesCabecera {
	return ImportacionComercializacionLubricantesCabecera{models.NewRequestWrapper(b.cabecera)}
}

type importacionComercializacionLubricantesDetalleBuilder struct {
	detalle *documentos.DetalleImportacionComercializacionLubricantes
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
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithUnidadMedida(v int) *importacionComercializacionLubricantesDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithPrecioUnitario(v float64) *importacionComercializacionLubricantesDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithMontoDescuento(v *float64) *importacionComercializacionLubricantesDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithSubTotal(v float64) *importacionComercializacionLubricantesDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithCantidadLitros(v float64) *importacionComercializacionLubricantesDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.CantidadLitros = v
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) WithPorcentajeDeduccionIehdDS25530(v *float64) *importacionComercializacionLubricantesDetalleBuilder {
	if v == nil {
		b.detalle.PorcentajeDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
		b.detalle.PorcentajeDeduccionIehdDS25530 = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *importacionComercializacionLubricantesDetalleBuilder) Build() ImportacionComercializacionLubricantesDetalle {
	return ImportacionComercializacionLubricantesDetalle{models.NewRequestWrapper(b.detalle)}
}
