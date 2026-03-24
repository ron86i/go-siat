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

// SectorEducativoZF representa la estructura completa de una factura del Sector Educativo ZF lista para ser procesada.
type SectorEducativoZF struct {
	models.RequestWrapper[documentos.FacturaSectorEducativoZF]
}

// SectorEducativoZFCabecera representa la sección de cabecera de una factura del Sector Educativo Zona Franca.
type SectorEducativoZFCabecera struct {
	models.RequestWrapper[documentos.CabeceraSectorEducativoZF]
}

// SectorEducativoZFDetalle representa un ítem individual dentro del detalle de una factura del Sector Educativo Zona Franca.
type SectorEducativoZFDetalle struct {
	models.RequestWrapper[documentos.DetalleSectorEducativoZF]
}

// NewSectorEducativoZFBuilder inicia el proceso de construcción de una Factura del Sector Educativo ZonaFranca.
func NewSectorEducativoZFBuilder() *sectorEducativoZFBuilder {
	return &sectorEducativoZFBuilder{
		factura: &documentos.FacturaSectorEducativoZF{
			XMLName:           xml.Name{Local: "facturaElectronicaSectorEducativoZF"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaSectorEducativoZF.xsd",
		},
	}
}

// NewSectorEducativoZFCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewSectorEducativoZFCabeceraBuilder() *sectorEducativoZFCabeceraBuilder {
	return &sectorEducativoZFCabeceraBuilder{
		cabecera: &documentos.CabeceraSectorEducativoZF{
			CodigoDocumentoSector: 46, // Sector 46 para Sector Educativo ZF
		},
	}
}

// NewSectorEducativoZFDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewSectorEducativoZFDetalleBuilder() *sectorEducativoZFDetalleBuilder {
	return &sectorEducativoZFDetalleBuilder{
		detalle: &documentos.DetalleSectorEducativoZF{},
	}
}

type sectorEducativoZFBuilder struct {
	factura *documentos.FacturaSectorEducativoZF
}

func (b *sectorEducativoZFBuilder) WithCabecera(req SectorEducativoZFCabecera) *sectorEducativoZFBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraSectorEducativoZF](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *sectorEducativoZFBuilder) AddDetalle(req SectorEducativoZFDetalle) *sectorEducativoZFBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleSectorEducativoZF](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *sectorEducativoZFBuilder) WithModalidad(tipo int) *sectorEducativoZFBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaSectorEducativoZF"}
		b.factura.XsiSchemaLocation = "facturaElectronicaSectorEducativoZF.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaSectorEducativoZF"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaSectorEducativoZF.xsd"
	}
	return b
}

func (b *sectorEducativoZFBuilder) Build() SectorEducativoZF {
	return SectorEducativoZF{models.NewRequestWrapper(b.factura)}
}

type sectorEducativoZFCabeceraBuilder struct {
	cabecera *documentos.CabeceraSectorEducativoZF
}

func (b *sectorEducativoZFCabeceraBuilder) WithNitEmisor(v int64) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithRazonSocialEmisor(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithMunicipio(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithTelefono(telefono *string) *sectorEducativoZFCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithNumeroFactura(v int64) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCuf(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCufd(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCodigoSucursal(v int) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithDireccion(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCodigoPuntoVenta(v *int) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithNombreRazonSocial(v *string) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithNumeroDocumento(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithComplemento(v *string) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCodigoCliente(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithNombreEstudiante(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.NombreEstudiante = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithPeriodoFacturado(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.PeriodoFacturado = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCodigoMetodoPago(v int) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithNumeroTarjeta(v *int64) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithMontoTotal(v float64) *sectorEducativoZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCodigoMoneda(v int) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithTipoCambio(v float64) *sectorEducativoZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithMontoTotalMoneda(v float64) *sectorEducativoZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithMontoGiftCard(v *float64) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithDescuentoAdicional(v *float64) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCodigoExcepcion(v *int) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithCafc(v *string) *sectorEducativoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithLeyenda(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithUsuario(v string) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *sectorEducativoZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *sectorEducativoZFCabeceraBuilder) WithCodigoDocumentoSector(v int) *sectorEducativoZFCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *sectorEducativoZFCabeceraBuilder) Build() SectorEducativoZFCabecera {
	return SectorEducativoZFCabecera{models.NewRequestWrapper(b.cabecera)}
}

type sectorEducativoZFDetalleBuilder struct {
	detalle *documentos.DetalleSectorEducativoZF
}

func (b *sectorEducativoZFDetalleBuilder) WithActividadEconomica(v string) *sectorEducativoZFDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithCodigoProductoSin(v int64) *sectorEducativoZFDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithCodigoProducto(v string) *sectorEducativoZFDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithDescripcion(v string) *sectorEducativoZFDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithCantidad(v float64) *sectorEducativoZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithUnidadMedida(v int) *sectorEducativoZFDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithPrecioUnitario(v float64) *sectorEducativoZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithMontoDescuento(v *float64) *sectorEducativoZFDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *sectorEducativoZFDetalleBuilder) WithSubTotal(v float64) *sectorEducativoZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *sectorEducativoZFDetalleBuilder) Build() SectorEducativoZFDetalle {
	return SectorEducativoZFDetalle{models.NewRequestWrapper(b.detalle)}
}
