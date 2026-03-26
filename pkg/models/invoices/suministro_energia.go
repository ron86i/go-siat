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

// SuministroEnergia representa la estructura completa de una factura de Suministro de Energía lista para ser procesada.
type SuministroEnergia struct {
	models.RequestWrapper[documents.FacturaSuministroEnergia]
}

// SuministroEnergiaCabecera representa la sección de cabecera de una factura de Energía.
type SuministroEnergiaCabecera struct {
	models.RequestWrapper[documents.CabeceraSuministroEnergia]
}

// SuministroEnergiaDetalle representa un ítem individual dentro del detalle de una factura de Energía.
type SuministroEnergiaDetalle struct {
	models.RequestWrapper[documents.DetalleSuministroEnergia]
}

// NewSuministroEnergiaBuilder inicia el proceso de construcción de una Factura de Suministro de Energía.
func NewSuministroEnergiaBuilder() *suministroEnergiaBuilder {
	return &suministroEnergiaBuilder{
		factura: &documents.FacturaSuministroEnergia{
			XMLName:           xml.Name{Local: "facturaElectronicaSuministroEnergia"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaSuministroEnergia.xsd",
		},
	}
}

// NewSuministroEnergiaCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewSuministroEnergiaCabeceraBuilder() *suministroEnergiaCabeceraBuilder {
	return &suministroEnergiaCabeceraBuilder{
		cabecera: &documents.CabeceraSuministroEnergia{
			CodigoDocumentoSector: 31, // Sector 31 para Energía
		},
	}
}

// NewSuministroEnergiaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewSuministroEnergiaDetalleBuilder() *suministroEnergiaDetalleBuilder {
	return &suministroEnergiaDetalleBuilder{
		detalle: &documents.DetalleSuministroEnergia{},
	}
}

type suministroEnergiaBuilder struct {
	factura *documents.FacturaSuministroEnergia
}

func (b *suministroEnergiaBuilder) WithCabecera(req SuministroEnergiaCabecera) *suministroEnergiaBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraSuministroEnergia](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *suministroEnergiaBuilder) AddDetalle(req SuministroEnergiaDetalle) *suministroEnergiaBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleSuministroEnergia](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *suministroEnergiaBuilder) WithModalidad(tipo int) *suministroEnergiaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaSuministroEnergia"}
		b.factura.XsiSchemaLocation = "facturaElectronicaSuministroEnergia.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaSuministroEnergia"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaSuministroEnergia.xsd"
	}
	return b
}

func (b *suministroEnergiaBuilder) Build() SuministroEnergia {
	return SuministroEnergia{models.NewRequestWrapper(b.factura)}
}

type suministroEnergiaCabeceraBuilder struct {
	cabecera *documents.CabeceraSuministroEnergia
}

func (b *suministroEnergiaCabeceraBuilder) WithNitEmisor(v int64) *suministroEnergiaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithRazonSocialEmisor(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithMunicipio(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithTelefono(telefono *string) *suministroEnergiaCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithNumeroFactura(v int64) *suministroEnergiaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCuf(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCufd(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoSucursal(v int) *suministroEnergiaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithDireccion(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *suministroEnergiaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithNombreRazonSocial(v *string) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *suministroEnergiaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithNumeroDocumento(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithComplemento(v *string) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoCliente(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoPais(v *int) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPais = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPais = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithPlacaVehiculo(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.PlacaVehiculo = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoMetodoPago(v int) *suministroEnergiaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithNumeroTarjeta(v *int64) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithMontoTotal(v float64) *suministroEnergiaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *suministroEnergiaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithMontoGiftCard(v *float64) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithDescuentoAdicional(v *float64) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoExcepcion(v *int) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCafc(v *string) *suministroEnergiaCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithCodigoMoneda(v int) *suministroEnergiaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithTipoCambio(v float64) *suministroEnergiaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithMontoTotalMoneda(v float64) *suministroEnergiaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithLeyenda(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) WithUsuario(v string) *suministroEnergiaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *suministroEnergiaCabeceraBuilder) WithCodigoDocumentoSector(v int) *suministroEnergiaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *suministroEnergiaCabeceraBuilder) Build() SuministroEnergiaCabecera {
	return SuministroEnergiaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type suministroEnergiaDetalleBuilder struct {
	detalle *documents.DetalleSuministroEnergia
}

func (b *suministroEnergiaDetalleBuilder) WithActividadEconomica(v string) *suministroEnergiaDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithCodigoProductoSin(v int64) *suministroEnergiaDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithCodigoProducto(v string) *suministroEnergiaDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithDescripcion(v string) *suministroEnergiaDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithCantidad(v float64) *suministroEnergiaDetalleBuilder {
	// Sector 31 usa 4 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 4, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithUnidadMedida(v int) *suministroEnergiaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithPrecioUnitario(v float64) *suministroEnergiaDetalleBuilder {
	// Sector 31 usa 4 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 4, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithMontoDescuento(v *float64) *suministroEnergiaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	// Sector 31 usa 4 decimales
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 4, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *suministroEnergiaDetalleBuilder) WithSubTotal(v float64) *suministroEnergiaDetalleBuilder {
	// Sector 31 usa 4 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 4, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *suministroEnergiaDetalleBuilder) Build() SuministroEnergiaDetalle {
	return SuministroEnergiaDetalle{models.NewRequestWrapper(b.detalle)}
}
