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

// AlquilerBienInmueble representa la estructura completa de una factura de Alquiler de Bienes Inmuebles lista para ser procesada.
type AlquilerBienInmueble struct {
	models.RequestWrapper[documents.FacturaAlquilerBienInmueble]
	models.MetadatosFactura
}

// AlquilerBienInmuebleCabecera representa la sección de cabecera de una factura de Alquiler de Bienes Inmuebles.
type AlquilerBienInmuebleCabecera struct {
	models.RequestWrapper[documents.CabeceraAlquilerBienInmueble]
}

// AlquilerBienInmuebleDetalle representa un ítem individual dentro del detalle de una factura de Alquiler de Bienes Inmuebles.
type AlquilerBienInmuebleDetalle struct {
	models.RequestWrapper[documents.DetalleAlquilerBienInmueble]
}

// NewAlquilerBienInmuebleBuilder inicia el proceso de construcción de una Factura de Alquiler de Bienes Inmuebles.
func NewAlquilerBienInmuebleBuilder() *alquilerBienInmuebleBuilder {
	return &alquilerBienInmuebleBuilder{
		metadatos: models.NuevosMetadatosFactura(2),
		factura: &documents.FacturaAlquilerBienInmueble{
			XMLName:           xml.Name{Local: "facturaElectronicaAlquilerBienInmueble"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaAlquilerBienInmueble.xsd",
		},
	}
}

// NewAlquilerBienInmuebleCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewAlquilerBienInmuebleCabeceraBuilder() *alquilerBienInmuebleCabeceraBuilder {
	return &alquilerBienInmuebleCabeceraBuilder{
		cabecera: &documents.CabeceraAlquilerBienInmueble{
			CodigoDocumentoSector: 2, // Sector 2 para Alquiler de Bienes Inmuebles
		},
	}
}

// NewAlquilerBienInmuebleDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewAlquilerBienInmuebleDetalleBuilder() *alquilerBienInmuebleDetalleBuilder {
	return &alquilerBienInmuebleDetalleBuilder{
		detalle: &documents.DetalleAlquilerBienInmueble{},
	}
}

type alquilerBienInmuebleBuilder struct {
	factura   *documents.FacturaAlquilerBienInmueble
	metadatos models.MetadatosFactura
}

// WithTipoFacturaDocumento reemplaza el tipo fiscal predeterminado de alquiler.
func (b *alquilerBienInmuebleBuilder) WithTipoFacturaDocumento(tipo int) *alquilerBienInmuebleBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *alquilerBienInmuebleBuilder) WithCabecera(req AlquilerBienInmuebleCabecera) *alquilerBienInmuebleBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraAlquilerBienInmueble](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *alquilerBienInmuebleBuilder) AddDetalle(req AlquilerBienInmuebleDetalle) *alquilerBienInmuebleBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleAlquilerBienInmueble](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *alquilerBienInmuebleBuilder) WithModalidad(tipo int) *alquilerBienInmuebleBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaAlquilerBienInmueble"}
		b.factura.XsiSchemaLocation = "facturaElectronicaAlquilerBienInmueble.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaAlquilerBienInmueble"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaAlquilerBienInmueble.xsd"
	}
	return b
}

func (b *alquilerBienInmuebleBuilder) Build() AlquilerBienInmueble {
	return AlquilerBienInmueble{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type alquilerBienInmuebleCabeceraBuilder struct {
	cabecera *documents.CabeceraAlquilerBienInmueble
}

func (b *alquilerBienInmuebleCabeceraBuilder) Build() AlquilerBienInmuebleCabecera {
	return AlquilerBienInmuebleCabecera{models.NewRequestWrapper(b.cabecera)}
}

type alquilerBienInmuebleDetalleBuilder struct {
	detalle *documents.DetalleAlquilerBienInmueble
}

func (b *alquilerBienInmuebleDetalleBuilder) Build() AlquilerBienInmuebleDetalle {
	return AlquilerBienInmuebleDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNitEmisor(v int64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithRazonSocialEmisor(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithMunicipio(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithTelefono(v *string) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNumeroFactura(v int64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCuf(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCufd(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoSucursal(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithDireccion(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoPuntoVenta(v *int) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithFechaEmision(v time.Time) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNombreRazonSocial(v *string) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNumeroDocumento(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithComplemento(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoCliente(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithPeriodoFacturado(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.PeriodoFacturado = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoMetodoPago(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithNumeroTarjeta(v *int64) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithSubTotalNeto(v float64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithTotalPrecioNeto(v float64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithIva(v float64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithMontoAPagar(v float64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoMoneda(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithTipoCambio(v float64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithMontoTotalMoneda(v float64) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithDescuentoAdicional(v *float64) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoExcepcion(v *int) *alquilerBienInmuebleCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithLeyenda(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithUsuario(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoInterno(v string) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *alquilerBienInmuebleCabeceraBuilder) WithCodigoDocumentoSector(v int) *alquilerBienInmuebleCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithActividadEconomica(v string) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithCodigoProductoSin(v int64) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithCodigoProducto(v string) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithDescripcion(v string) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithCantidad(v float64) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithUnidadMedida(v int) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithPrecioNetoUnitario(v float64) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithMontoDescuento(v *float64) *alquilerBienInmuebleDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alquilerBienInmuebleDetalleBuilder) WithSubTotal(v float64) *alquilerBienInmuebleDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}
