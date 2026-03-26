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

// MonedaExtranjera representa la estructura completa de una factura de Sector 9.
type MonedaExtranjera struct {
	models.RequestWrapper[documents.FacturaMonedaExtranjera]
}

// MonedaExtranjeraCabecera representa la sección de cabecera de Sector 9.
type MonedaExtranjeraCabecera struct {
	models.RequestWrapper[documents.CabeceraMonedaExtranjera]
}

// MonedaExtranjeraDetalle representa un ítem individual de Sector 9.
type MonedaExtranjeraDetalle struct {
	models.RequestWrapper[documents.DetalleMonedaExtranjera]
}

// NewMonedaExtranjeraBuilder inicia el proceso de construcción.
func NewMonedaExtranjeraBuilder() *monedaExtranjeraBuilder {
	return &monedaExtranjeraBuilder{
		factura: &documents.FacturaMonedaExtranjera{
			XMLName:  xml.Name{Local: "facturaElectronicaMonedaExtranjera"},
			XmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		},
	}
}

// NewMonedaExtranjeraCabeceraBuilder crea el constructor para la cabecera.
func NewMonedaExtranjeraCabeceraBuilder() *monedaExtranjeraCabeceraBuilder {
	return &monedaExtranjeraCabeceraBuilder{
		cabecera: &documents.CabeceraMonedaExtranjera{
			CodigoDocumentoSector: 9,
		},
	}
}

// NewMonedaExtranjeraDetalleBuilder crea el constructor para los ítems.
func NewMonedaExtranjeraDetalleBuilder() *monedaExtranjeraDetalleBuilder {
	return &monedaExtranjeraDetalleBuilder{
		detalle: &documents.DetalleMonedaExtranjera{},
	}
}

type monedaExtranjeraBuilder struct {
	factura *documents.FacturaMonedaExtranjera
}

func (b *monedaExtranjeraBuilder) WithCabecera(req MonedaExtranjeraCabecera) *monedaExtranjeraBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraMonedaExtranjera](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *monedaExtranjeraBuilder) AddDetalle(req MonedaExtranjeraDetalle) *monedaExtranjeraBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleMonedaExtranjera](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *monedaExtranjeraBuilder) WithModalidad(tipo int) *monedaExtranjeraBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaMonedaExtranjera"}
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaMonedaExtranjera"}
	}
	return b
}

func (b *monedaExtranjeraBuilder) Build() MonedaExtranjera {
	return MonedaExtranjera{models.NewRequestWrapper(b.factura)}
}

type monedaExtranjeraCabeceraBuilder struct {
	cabecera *documents.CabeceraMonedaExtranjera
}

func (b *monedaExtranjeraCabeceraBuilder) WithNitEmisor(v int64) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithRazonSocialEmisor(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithMunicipio(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithTelefono(telefono *string) *monedaExtranjeraCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *telefono
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithNumeroFactura(v int64) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCuf(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCufd(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoSucursal(v int) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithDireccion(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoPuntoVenta(v *int) *monedaExtranjeraCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithFechaEmision(v time.Time) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithNombreRazonSocial(v *string) *monedaExtranjeraCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithNumeroDocumento(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithComplemento(v *string) *monedaExtranjeraCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoCliente(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoTipoOperacion(v int) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.CodigoTipoOperacion = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoMetodoPago(v int) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithNumeroTarjeta(v *int64) *monedaExtranjeraCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithMontoTotal(v float64) *monedaExtranjeraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *monedaExtranjeraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithIngresoDiferenciaCambio(v float64) *monedaExtranjeraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.IngresoDiferenciaCambio = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoMoneda(v int) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithTipoCambio(v float64) *monedaExtranjeraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithMontoTotalMoneda(v float64) *monedaExtranjeraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithDescuentoAdicional(v *float64) *monedaExtranjeraCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCodigoExcepcion(v *int) *monedaExtranjeraCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithCafc(v *string) *monedaExtranjeraCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithLeyenda(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithUsuario(v string) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) WithTipoCambioOficial(v float64) *monedaExtranjeraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambioOficial = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *monedaExtranjeraCabeceraBuilder) WithCodigoDocumentoSector(v int) *monedaExtranjeraCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *monedaExtranjeraCabeceraBuilder) Build() MonedaExtranjeraCabecera {
	return MonedaExtranjeraCabecera{models.NewRequestWrapper(b.cabecera)}
}

type monedaExtranjeraDetalleBuilder struct {
	detalle *documents.DetalleMonedaExtranjera
}

func (b *monedaExtranjeraDetalleBuilder) WithActividadEconomica(v string) *monedaExtranjeraDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithCodigoProductoSin(v int64) *monedaExtranjeraDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithCodigoProducto(v string) *monedaExtranjeraDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithDescripcion(v string) *monedaExtranjeraDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithCantidad(v float64) *monedaExtranjeraDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithUnidadMedida(v int) *monedaExtranjeraDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithPrecioUnitario(v float64) *monedaExtranjeraDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithMontoDescuento(v *float64) *monedaExtranjeraDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 5, 64), 64)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *monedaExtranjeraDetalleBuilder) WithSubTotal(v float64) *monedaExtranjeraDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *monedaExtranjeraDetalleBuilder) Build() MonedaExtranjeraDetalle {
	return MonedaExtranjeraDetalle{models.NewRequestWrapper(b.detalle)}
}
