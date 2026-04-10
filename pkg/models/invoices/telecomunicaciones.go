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

// Telecomunicaciones representa la estructura completa de una factura de telecomunicaciones lista para ser procesada.
type Telecomunicaciones struct {
	models.RequestWrapper[documents.FacturaTelecomunicaciones]
}

// TelecomunicacionesCabecera representa la sección de cabecera de una factura de telecomunicaciones.
type TelecomunicacionesCabecera struct {
	models.RequestWrapper[documents.CabeceraTelecomunicaciones]
}

// TelecomunicacionesDetalle representa un ítem individual dentro del detalle de una factura de telecomunicaciones.
type TelecomunicacionesDetalle struct {
	models.RequestWrapper[documents.DetalleTelecomunicaciones]
}

// NewTelecomunicacionesBuilder inicia el proceso de construcción de una factura de Telecomunicaciones,
// configurando los nombres de nodo XML necesarios para el SIAT.
// Por defecto, la factura se configura para ser emitida electrónica.
func NewTelecomunicacionesBuilder() *telecomunicacionesBuilder {
	return &telecomunicacionesBuilder{
		factura: &documents.FacturaTelecomunicaciones{
			XMLName:           xml.Name{Local: "facturaElectronicaTelecomunicacion"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaTelecomunicacion.xsd",
		},
	}
}

// NewTelecomunicacionesCabeceraBuilder crea una nueva instancia del constructor para la cabecera
// de facturas de telecomunicaciones.
func NewTelecomunicacionesCabeceraBuilder() *telecomunicacionesCabeceraBuilder {
	return &telecomunicacionesCabeceraBuilder{
		cabecera: &documents.CabeceraTelecomunicaciones{
			CodigoDocumentoSector: 22, // Sector 22 para Telecomunicaciones
		},
	}
}

// NewTelecomunicacionesDetalleBuilder crea una nueva instancia del constructor para los ítems
// de detalle de la factura.
func NewTelecomunicacionesDetalleBuilder() *telecomunicacionesDetalleBuilder {
	return &telecomunicacionesDetalleBuilder{
		detalle: &documents.DetalleTelecomunicaciones{},
	}
}

type telecomunicacionesBuilder struct {
	factura *documents.FacturaTelecomunicaciones
}

// WithCabecera asocia la cabecera construida previamente a la factura.
func (b *telecomunicacionesBuilder) WithCabecera(req TelecomunicacionesCabecera) *telecomunicacionesBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraTelecomunicaciones](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

// AddDetalle añade un ítem de detalle a la lista de detalles de la factura.
func (b *telecomunicacionesBuilder) AddDetalle(req TelecomunicacionesDetalle) *telecomunicacionesBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleTelecomunicaciones](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

// WithModalidad configura los metadatos XML de la factura según la modalidad (Electrónica o Computarizada).
func (b *telecomunicacionesBuilder) WithModalidad(tipo int) *telecomunicacionesBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaTelecomunicacion"}
		b.factura.XsiSchemaLocation = "facturaElectronicaTelecomunicacion.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaTelecomunicacion"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaTelecomunicacion.xsd"
	}
	return b
}

// Build finaliza la construcción y retorna la estructura opaca lista para ser firmada y enviada.
func (b *telecomunicacionesBuilder) Build() Telecomunicaciones {
	return Telecomunicaciones{models.NewRequestWrapper(b.factura)}
}

// telecomunicacionesCabeceraBuilder ayuda a configurar la cabecera de la factura.
type telecomunicacionesCabeceraBuilder struct {
	cabecera *documents.CabeceraTelecomunicaciones
}

func (b *telecomunicacionesCabeceraBuilder) WithNitEmisor(v int64) *telecomunicacionesCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithRazonSocialEmisor(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithMunicipio(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithTelefono(v *string) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithNitConjunto(v *int64) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NitConjunto = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NitConjunto = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithNumeroFactura(v int64) *telecomunicacionesCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCuf(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCufd(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCodigoSucursal(v int) *telecomunicacionesCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithDireccion(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCodigoPuntoVenta(v *int) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithFechaEmision(v time.Time) *telecomunicacionesCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithNombreRazonSocial(v *string) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *telecomunicacionesCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithNumeroDocumento(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithComplemento(v *string) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCodigoCliente(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCodigoMetodoPago(v int) *telecomunicacionesCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithNumeroTarjeta(v *int64) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithMontoTotal(v float64) *telecomunicacionesCabeceraBuilder {
	b.cabecera.MontoTotal = utils.Round(v, 2)
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *telecomunicacionesCabeceraBuilder {
	b.cabecera.MontoTotalSujetoIva = utils.Round(v, 2)
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCodigoMoneda(v int) *telecomunicacionesCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithTipoCambio(v float64) *telecomunicacionesCabeceraBuilder {
	b.cabecera.TipoCambio = utils.Round(v, 2)
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithMontoTotalMoneda(v float64) *telecomunicacionesCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = utils.Round(v, 2)
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithMontoGiftCard(v *float64) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
	} else {
		val := utils.Round(*v, 2)
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithDescuentoAdicional(v *float64) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		val := utils.Round(*v, 2)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCodigoExcepcion(v *int) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithCafc(v *string) *telecomunicacionesCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithLeyenda(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *telecomunicacionesCabeceraBuilder) WithUsuario(v string) *telecomunicacionesCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// Build finaliza la construcción de la cabecera retornando la estructura opaca.
func (b *telecomunicacionesCabeceraBuilder) Build() TelecomunicacionesCabecera {
	return TelecomunicacionesCabecera{models.NewRequestWrapper(b.cabecera)}
}

// telecomunicacionesDetalleBuilder ayuda a configurar un ítem individual de detalle.
type telecomunicacionesDetalleBuilder struct {
	detalle *documents.DetalleTelecomunicaciones
}

func (b *telecomunicacionesDetalleBuilder) WithActividadEconomica(v string) *telecomunicacionesDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithCodigoProductoSin(v int64) *telecomunicacionesDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithCodigoProducto(v string) *telecomunicacionesDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithDescripcion(v string) *telecomunicacionesDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithCantidad(v float64) *telecomunicacionesDetalleBuilder {
	b.detalle.Cantidad = utils.Round(v, 5)
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithUnidadMedida(v int) *telecomunicacionesDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithPrecioUnitario(v float64) *telecomunicacionesDetalleBuilder {
	b.detalle.PrecioUnitario = utils.Round(v, 5)
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithMontoDescuento(v *float64) *telecomunicacionesDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := utils.Round(*v, 5)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithSubTotal(v float64) *telecomunicacionesDetalleBuilder {
	b.detalle.SubTotal = utils.Round(v, 5)
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithNumeroSerie(v *string) *telecomunicacionesDetalleBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesDetalleBuilder) WithNumeroImei(v *string) *telecomunicacionesDetalleBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.detalle.NumeroImei = datatype.Nilable[string]{Value: &val}
	}
	return b
}

// Build finaliza la construcción del detalle retornando la estructura opaca.
func (b *telecomunicacionesDetalleBuilder) Build() TelecomunicacionesDetalle {
	return TelecomunicacionesDetalle{models.NewRequestWrapper(b.detalle)}
}
