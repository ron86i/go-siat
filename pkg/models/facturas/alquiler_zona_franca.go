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

// AlquilerZF representa la estructura completa de una factura de Alquiler ZF lista para ser procesada.
type AlquilerZF struct {
	models.RequestWrapper[documentos.FacturaAlquilerZF]
}

// AlquilerZFCabecera representa la sección de cabecera de una factura de Alquiler ZF.
type AlquilerZFCabecera struct {
	models.RequestWrapper[documentos.CabeceraAlquilerZF]
}

// AlquilerZFDetalle representa un ítem individual dentro del detalle de una factura de Alquiler ZF.
type AlquilerZFDetalle struct {
	models.RequestWrapper[documentos.DetalleAlquilerZF]
}

// NewAlquilerZFBuilder inicia el proceso de construcción de una Factura de Alquiler ZF.
func NewAlquilerZFBuilder() *alquilerZFBuilder {
	return &alquilerZFBuilder{
		factura: &documentos.FacturaAlquilerZF{
			XMLName:           xml.Name{Local: "facturaElectronicaAlquilerZF"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaAlquilerZF.xsd",
		},
	}
}

// NewAlquilerZFCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewAlquilerZFCabeceraBuilder() *alquilerZFCabeceraBuilder {
	return &alquilerZFCabeceraBuilder{
		cabecera: &documentos.CabeceraAlquilerZF{
			CodigoDocumentoSector: 42, // Sector 42 para Alquiler ZF
			MontoTotalSujetoIva:   0,  // Fijo 0 según XSD
		},
	}
}

// NewAlquilerZFDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewAlquilerZFDetalleBuilder() *alquilerZFDetalleBuilder {
	return &alquilerZFDetalleBuilder{
		detalle: &documentos.DetalleAlquilerZF{},
	}
}

type alquilerZFBuilder struct {
	factura *documentos.FacturaAlquilerZF
}

func (b *alquilerZFBuilder) WithCabecera(req AlquilerZFCabecera) *alquilerZFBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraAlquilerZF](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *alquilerZFBuilder) AddDetalle(req AlquilerZFDetalle) *alquilerZFBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleAlquilerZF](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *alquilerZFBuilder) WithModalidad(tipo int) *alquilerZFBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaAlquilerZF"}
		b.factura.XsiSchemaLocation = "facturaElectronicaAlquilerZF.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaAlquilerZF"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaAlquilerZF.xsd"
	}
	return b
}

func (b *alquilerZFBuilder) Build() AlquilerZF {
	return AlquilerZF{models.NewRequestWrapper(b.factura)}
}

type alquilerZFCabeceraBuilder struct {
	cabecera *documentos.CabeceraAlquilerZF
}

func (b *alquilerZFCabeceraBuilder) WithNitEmisor(v int64) *alquilerZFCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithRazonSocialEmisor(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithMunicipio(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithTelefono(telefono *string) *alquilerZFCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithNumeroFactura(v int64) *alquilerZFCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCuf(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCufd(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCodigoSucursal(v int) *alquilerZFCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithDireccion(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCodigoPuntoVenta(v *int) *alquilerZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *alquilerZFCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *alquilerZFCabeceraBuilder) WithNombreRazonSocial(v *string) *alquilerZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *alquilerZFCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithNumeroDocumento(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithComplemento(v *string) *alquilerZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCodigoCliente(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithPeriodoFacturado(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.PeriodoFacturado = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCodigoMetodoPago(v int) *alquilerZFCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithNumeroTarjeta(v *int64) *alquilerZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithMontoTotal(v float64) *alquilerZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCodigoMoneda(v int) *alquilerZFCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithTipoCambio(v float64) *alquilerZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithMontoTotalMoneda(v float64) *alquilerZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithDescuentoAdicional(v *float64) *alquilerZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCodigoExcepcion(v *int) *alquilerZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithCafc(v *string) *alquilerZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *alquilerZFCabeceraBuilder) WithLeyenda(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithUsuario(v string) *alquilerZFCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *alquilerZFCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *alquilerZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *alquilerZFCabeceraBuilder) WithCodigoDocumentoSector(v int) *alquilerZFCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *alquilerZFCabeceraBuilder) Build() AlquilerZFCabecera {
	return AlquilerZFCabecera{models.NewRequestWrapper(b.cabecera)}
}

type alquilerZFDetalleBuilder struct {
	detalle *documentos.DetalleAlquilerZF
}

func (b *alquilerZFDetalleBuilder) WithActividadEconomica(v string) *alquilerZFDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *alquilerZFDetalleBuilder) WithCodigoProductoSin(v int64) *alquilerZFDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *alquilerZFDetalleBuilder) WithCodigoProducto(v string) *alquilerZFDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *alquilerZFDetalleBuilder) WithDescripcion(v string) *alquilerZFDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *alquilerZFDetalleBuilder) WithCantidad(v float64) *alquilerZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *alquilerZFDetalleBuilder) WithUnidadMedida(v int) *alquilerZFDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *alquilerZFDetalleBuilder) WithPrecioUnitario(v float64) *alquilerZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *alquilerZFDetalleBuilder) WithMontoDescuento(v *float64) *alquilerZFDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *alquilerZFDetalleBuilder) WithSubTotal(v float64) *alquilerZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *alquilerZFDetalleBuilder) Build() AlquilerZFDetalle {
	return AlquilerZFDetalle{models.NewRequestWrapper(b.detalle)}
}
