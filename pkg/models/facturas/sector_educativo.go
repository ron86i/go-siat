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

// SectorEducativo representa la estructura completa de una factura del Sector Educativo lista para ser procesada.
type SectorEducativo struct {
	models.RequestWrapper[documentos.FacturaSectorEducativo]
}

// SectorEducativoCabecera representa la sección de cabecera de una factura del Sector Educativo.
type SectorEducativoCabecera struct {
	models.RequestWrapper[documentos.CabeceraSectorEducativo]
}

// SectorEducativoDetalle representa un ítem individual dentro del detalle de una factura del Sector Educativo.
type SectorEducativoDetalle struct {
	models.RequestWrapper[documentos.DetalleSectorEducativo]
}

// NewSectorEducativoBuilder inicia el proceso de construcción de una Factura del Sector Educativo.
func NewSectorEducativoBuilder() *sectorEducativoBuilder {
	return &sectorEducativoBuilder{
		factura: &documentos.FacturaSectorEducativo{
			XMLName:           xml.Name{Local: "facturaElectronicaSectorEducativo"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaSectorEducativo.xsd",
		},
	}
}

// NewSectorEducativoCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewSectorEducativoCabeceraBuilder() *sectorEducativoCabeceraBuilder {
	return &sectorEducativoCabeceraBuilder{
		cabecera: &documentos.CabeceraSectorEducativo{
			CodigoDocumentoSector: 11, // Sector 11 para Sector Educativo
		},
	}
}

// NewSectorEducativoDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewSectorEducativoDetalleBuilder() *sectorEducativoDetalleBuilder {
	return &sectorEducativoDetalleBuilder{
		detalle: &documentos.DetalleSectorEducativo{},
	}
}

type sectorEducativoBuilder struct {
	factura *documentos.FacturaSectorEducativo
}

func (b *sectorEducativoBuilder) WithCabecera(req SectorEducativoCabecera) *sectorEducativoBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraSectorEducativo](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *sectorEducativoBuilder) AddDetalle(req SectorEducativoDetalle) *sectorEducativoBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleSectorEducativo](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *sectorEducativoBuilder) WithModalidad(tipo int) *sectorEducativoBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaSectorEducativo"}
		b.factura.XsiSchemaLocation = "facturaElectronicaSectorEducativo.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaSectorEducativo"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaSectorEducativo.xsd"
	}
	return b
}

func (b *sectorEducativoBuilder) Build() SectorEducativo {
	return SectorEducativo{models.NewRequestWrapper(b.factura)}
}

type sectorEducativoCabeceraBuilder struct {
	cabecera *documentos.CabeceraSectorEducativo
}

func (b *sectorEducativoCabeceraBuilder) WithNitEmisor(v int64) *sectorEducativoCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithRazonSocialEmisor(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithMunicipio(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithTelefono(telefono *string) *sectorEducativoCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithNumeroFactura(v int64) *sectorEducativoCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCuf(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCufd(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCodigoSucursal(v int) *sectorEducativoCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithDireccion(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCodigoPuntoVenta(v *int) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *sectorEducativoCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithNombreRazonSocial(v *string) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *sectorEducativoCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithNumeroDocumento(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithComplemento(v *string) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCodigoCliente(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithNombreEstudiante(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.NombreEstudiante = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithPeriodoFacturado(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.PeriodoFacturado = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCodigoMetodoPago(v int) *sectorEducativoCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithNumeroTarjeta(v *int64) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithMontoTotal(v float64) *sectorEducativoCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *sectorEducativoCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCodigoMoneda(v int) *sectorEducativoCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithTipoCambio(v float64) *sectorEducativoCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithMontoTotalMoneda(v float64) *sectorEducativoCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithMontoGiftCard(v *float64) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithDescuentoAdicional(v *float64) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCodigoExcepcion(v *int) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithCafc(v *string) *sectorEducativoCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithLeyenda(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) WithUsuario(v string) *sectorEducativoCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *sectorEducativoCabeceraBuilder) WithCodigoDocumentoSector(v int) *sectorEducativoCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *sectorEducativoCabeceraBuilder) Build() SectorEducativoCabecera {
	return SectorEducativoCabecera{models.NewRequestWrapper(b.cabecera)}
}

type sectorEducativoDetalleBuilder struct {
	detalle *documentos.DetalleSectorEducativo
}

func (b *sectorEducativoDetalleBuilder) WithActividadEconomica(v string) *sectorEducativoDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *sectorEducativoDetalleBuilder) WithCodigoProductoSin(v int64) *sectorEducativoDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *sectorEducativoDetalleBuilder) WithCodigoProducto(v string) *sectorEducativoDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *sectorEducativoDetalleBuilder) WithDescripcion(v string) *sectorEducativoDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *sectorEducativoDetalleBuilder) WithCantidad(v float64) *sectorEducativoDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *sectorEducativoDetalleBuilder) WithUnidadMedida(v int) *sectorEducativoDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *sectorEducativoDetalleBuilder) WithPrecioUnitario(v float64) *sectorEducativoDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *sectorEducativoDetalleBuilder) WithMontoDescuento(v *float64) *sectorEducativoDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *sectorEducativoDetalleBuilder) WithSubTotal(v float64) *sectorEducativoDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *sectorEducativoDetalleBuilder) Build() SectorEducativoDetalle {
	return SectorEducativoDetalle{models.NewRequestWrapper(b.detalle)}
}
