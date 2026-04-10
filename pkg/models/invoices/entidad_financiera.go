package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
)

// EntidadFinanciera representa la estructura completa de una factura de entidad financiera lista para ser procesada.
type EntidadFinanciera struct {
	models.RequestWrapper[documents.FacturaEntidadFinanciera]
}

// EntidadFinancieraCabecera representa la sección de cabecera de una factura de entidad financiera.
type EntidadFinancieraCabecera struct {
	models.RequestWrapper[documents.CabeceraEntidadFinanciera]
}

// EntidadFinancieraDetalle representa un ítem individual dentro del detalle de una factura de entidad financiera.
type EntidadFinancieraDetalle struct {
	models.RequestWrapper[documents.DetalleEntidadFinanciera]
}

type entidadFinancieraBuilder struct {
	factura *documents.FacturaEntidadFinanciera
}

func NewEntidadFinancieraBuilder() *entidadFinancieraBuilder {
	return &entidadFinancieraBuilder{
		factura: &documents.FacturaEntidadFinanciera{
			XMLName:           xml.Name{Local: "facturaElectronicaEntidadFinanciera"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaEntidadFinanciera.xsd",
		},
	}
}

func (b *entidadFinancieraBuilder) WithModalidad(tipo int) *entidadFinancieraBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaEntidadFinanciera"}
		b.factura.XsiSchemaLocation = "facturaElectronicaEntidadFinanciera.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaEntidadFinanciera"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaEntidadFinanciera.xsd"
	}
	return b
}

func (b *entidadFinancieraBuilder) WithCabecera(req EntidadFinancieraCabecera) *entidadFinancieraBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraEntidadFinanciera](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *entidadFinancieraBuilder) AddDetalle(req EntidadFinancieraDetalle) *entidadFinancieraBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleEntidadFinanciera](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *entidadFinancieraBuilder) Build() EntidadFinanciera {
	return EntidadFinanciera{models.NewRequestWrapper(b.factura)}
}

// Cabecera Builder
type entidadFinancieraCabeceraBuilder struct {
	cabecera *documents.CabeceraEntidadFinanciera
}

func NewEntidadFinancieraCabeceraBuilder() *entidadFinancieraCabeceraBuilder {
	return &entidadFinancieraCabeceraBuilder{
		cabecera: &documents.CabeceraEntidadFinanciera{
			CodigoDocumentoSector: 15,
		},
	}
}

func (b *entidadFinancieraCabeceraBuilder) WithNitEmisor(v int64) *entidadFinancieraCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithRazonSocialEmisor(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithMunicipio(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithTelefono(v *string) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithNumeroFactura(v int64) *entidadFinancieraCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCuf(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCufd(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCodigoSucursal(v int) *entidadFinancieraCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithDireccion(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCodigoPuntoVenta(v *int) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithFechaEmision(v time.Time) *entidadFinancieraCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithNombreRazonSocial(v *string) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *entidadFinancieraCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithNumeroDocumento(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithComplemento(v *string) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCodigoCliente(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCodigoMetodoPago(v int) *entidadFinancieraCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithNumeroTarjeta(v *int64) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithMontoTotalArrendamientoFinanciero(v *float64) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoTotalArrendamientoFinanciero = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.MontoTotalArrendamientoFinanciero = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithMontoTotal(v float64) *entidadFinancieraCabeceraBuilder {
	b.cabecera.MontoTotal = utils.Round(v, 2)
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *entidadFinancieraCabeceraBuilder {
	b.cabecera.MontoTotalSujetoIva = utils.Round(v, 2)
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCodigoMoneda(v int) *entidadFinancieraCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithTipoCambio(v float64) *entidadFinancieraCabeceraBuilder {
	b.cabecera.TipoCambio = utils.Round(v, 2)
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithMontoTotalMoneda(v float64) *entidadFinancieraCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = utils.Round(v, 2)
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithDescuentoAdicional(v *float64) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCodigoExcepcion(v *int) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithCafc(v *string) *entidadFinancieraCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithLeyenda(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) WithUsuario(v string) *entidadFinancieraCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *entidadFinancieraCabeceraBuilder) Build() EntidadFinancieraCabecera {
	return EntidadFinancieraCabecera{models.NewRequestWrapper(b.cabecera)}
}

// Detalle Builder
type entidadFinancieraDetalleBuilder struct {
	detalle *documents.DetalleEntidadFinanciera
}

func NewEntidadFinancieraDetalleBuilder() *entidadFinancieraDetalleBuilder {
	return &entidadFinancieraDetalleBuilder{
		detalle: &documents.DetalleEntidadFinanciera{},
	}
}

func (b *entidadFinancieraDetalleBuilder) WithActividadEconomica(v string) *entidadFinancieraDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithCodigoProductoSin(v int64) *entidadFinancieraDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithCodigoProducto(v string) *entidadFinancieraDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithDescripcion(v string) *entidadFinancieraDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithCantidad(v float64) *entidadFinancieraDetalleBuilder {
	b.detalle.Cantidad = utils.Round(v, 2)
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithUnidadMedida(v int) *entidadFinancieraDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithPrecioUnitario(v float64) *entidadFinancieraDetalleBuilder {
	b.detalle.PrecioUnitario = utils.Round(v, 2)
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithMontoDescuento(v *float64) *entidadFinancieraDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *entidadFinancieraDetalleBuilder) WithSubTotal(v float64) *entidadFinancieraDetalleBuilder {
	b.detalle.SubTotal = utils.Round(v, 2)
	return b
}

func (b *entidadFinancieraDetalleBuilder) Build() EntidadFinancieraDetalle {
	return EntidadFinancieraDetalle{models.NewRequestWrapper(b.detalle)}
}
