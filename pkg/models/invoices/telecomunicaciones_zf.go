package invoices

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
)

// TelecomunicacionesZF representa la estructura completa de una factura de telecomunicaciones en Zona Franca.
type TelecomunicacionesZF struct {
	models.RequestWrapper[documents.FacturaTelecomunicacionesZF]
}

// TelecomunicacionesZFCabecera representa la sección de cabecera de la factura ZF.
type TelecomunicacionesZFCabecera struct {
	models.RequestWrapper[documents.CabeceraTelecomunicacionesZF]
}

// TelecomunicacionesZFDetalle representa un ítem individual dentro del detalle de la factura ZF.
type TelecomunicacionesZFDetalle struct {
	models.RequestWrapper[documents.DetalleTelecomunicacionesZF]
}

// NewTelecomunicacionesZFBuilder inicia el proceso de construcción de una factura de Telecomunicaciones ZF.
func NewTelecomunicacionesZFBuilder() *telecomunicacionesZFBuilder {
	return &telecomunicacionesZFBuilder{
		factura: &documents.FacturaTelecomunicacionesZF{
			XMLName:           xml.Name{Local: "facturaElectronicaTelecomunicacionZF"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaTelecomunicacionZF.xsd",
		},
	}
}

// NewTelecomunicacionesZFCabeceraBuilder crea una nueva instancia del constructor para la cabecera.
func NewTelecomunicacionesZFCabeceraBuilder() *telecomunicacionesZFCabeceraBuilder {
	return &telecomunicacionesZFCabeceraBuilder{
		cabecera: &documents.CabeceraTelecomunicacionesZF{
			CodigoDocumentoSector: 49, // Sector 49 para Telecomunicaciones ZF
			MontoTotalSujetoIva:   0,  // Siempre 0 en Zona Franca
		},
	}
}

// NewTelecomunicacionesZFDetalleBuilder crea una nueva instancia del constructor para los ítems.
func NewTelecomunicacionesZFDetalleBuilder() *telecomunicacionesZFDetalleBuilder {
	return &telecomunicacionesZFDetalleBuilder{
		detalle: &documents.DetalleTelecomunicacionesZF{},
	}
}

type telecomunicacionesZFBuilder struct {
	factura *documents.FacturaTelecomunicacionesZF
}

// WithCabecera asocia la cabecera construida previamente a la factura.
func (b *telecomunicacionesZFBuilder) WithCabecera(req TelecomunicacionesZFCabecera) *telecomunicacionesZFBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraTelecomunicacionesZF](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

// AddDetalle añade un ítem de detalle a la lista de detalles de la factura.
func (b *telecomunicacionesZFBuilder) AddDetalle(req TelecomunicacionesZFDetalle) *telecomunicacionesZFBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleTelecomunicacionesZF](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

// WithModalidad configura los metadatos XML de la factura según la modalidad.
// Nota: ZF suele ser electrónica, pero se mantiene la flexibilidad.
func (b *telecomunicacionesZFBuilder) WithModalidad(tipo int) *telecomunicacionesZFBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaTelecomunicacionZF"}
		b.factura.XsiSchemaLocation = "facturaElectronicaTelecomunicacionZF.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaTelecomunicacionZF"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaTelecomunicacionZF.xsd"
	}
	return b
}

// Build finaliza la construcción y retorna la estructura opaca.
func (b *telecomunicacionesZFBuilder) Build() TelecomunicacionesZF {
	return TelecomunicacionesZF{models.NewRequestWrapper(b.factura)}
}

// telecomunicacionesZFCabeceraBuilder ayuda a configurar la cabecera de la factura ZF.
type telecomunicacionesZFCabeceraBuilder struct {
	cabecera *documents.CabeceraTelecomunicacionesZF
}

func (b *telecomunicacionesZFCabeceraBuilder) WithNitEmisor(v int64) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithRazonSocialEmisor(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithMunicipio(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithTelefono(v *string) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithNitConjunto(v *int64) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NitConjunto = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NitConjunto = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithNumeroFactura(v int64) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCuf(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCufd(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCodigoSucursal(v int) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithDireccion(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCodigoPuntoVenta(v *int) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithFechaEmision(v time.Time) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithNombreRazonSocial(v *string) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithNumeroDocumento(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithComplemento(v *string) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCodigoCliente(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCodigoMetodoPago(v int) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithNumeroTarjeta(v *int64) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithMontoTotal(v float64) *telecomunicacionesZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCodigoMoneda(v int) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithTipoCambio(v float64) *telecomunicacionesZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithMontoTotalMoneda(v float64) *telecomunicacionesZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithMontoGiftCard(v *float64) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithDescuentoAdicional(v *float64) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCodigoExcepcion(v *int) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithCafc(v *string) *telecomunicacionesZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithLeyenda(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *telecomunicacionesZFCabeceraBuilder) WithUsuario(v string) *telecomunicacionesZFCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// Build finaliza la construcción de la cabecera retornando la estructura opaca.
func (b *telecomunicacionesZFCabeceraBuilder) Build() TelecomunicacionesZFCabecera {
	return TelecomunicacionesZFCabecera{models.NewRequestWrapper(b.cabecera)}
}

// telecomunicacionesZFDetalleBuilder ayuda a configurar un ítem individual de detalle.
type telecomunicacionesZFDetalleBuilder struct {
	detalle *documents.DetalleTelecomunicacionesZF
}

func (b *telecomunicacionesZFDetalleBuilder) WithActividadEconomica(v string) *telecomunicacionesZFDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithCodigoProductoSin(v int64) *telecomunicacionesZFDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithCodigoProducto(v string) *telecomunicacionesZFDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithDescripcion(v string) *telecomunicacionesZFDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithCantidad(v float64) *telecomunicacionesZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithUnidadMedida(v int) *telecomunicacionesZFDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithPrecioUnitario(v float64) *telecomunicacionesZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithMontoDescuento(v *float64) *telecomunicacionesZFDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithSubTotal(v float64) *telecomunicacionesZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithNumeroSerie(v *string) *telecomunicacionesZFDetalleBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *telecomunicacionesZFDetalleBuilder) WithNumeroImei(v *string) *telecomunicacionesZFDetalleBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.detalle.NumeroImei = datatype.Nilable[string]{Value: &val}
	}
	return b
}

// Build finaliza la construcción del detalle retornando la estructura opaca.
func (b *telecomunicacionesZFDetalleBuilder) Build() TelecomunicacionesZFDetalle {
	return TelecomunicacionesZFDetalle{models.NewRequestWrapper(b.detalle)}
}
