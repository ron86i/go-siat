package invoices

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
)

// ComercialExportacionMinera representa la estructura completa de una factura de Comercial Exportación Minera lista para ser procesada.
type ComercialExportacionMinera struct {
	models.RequestWrapper[documents.FacturaComercialExportacionMinera]
}

// ComercialExportacionMineraCabecera representa la sección de cabecera de una factura de Comercial Exportación Minera.
type ComercialExportacionMineraCabecera struct {
	models.RequestWrapper[documents.CabeceraComercialExportacionMinera]
}

// ComercialExportacionMineraDetalle representa un ítem individual dentro del detalle de una factura de Comercial Exportación Minera.
type ComercialExportacionMineraDetalle struct {
	models.RequestWrapper[documents.DetalleComercialExportacionMinera]
}

// NewComercialExportacionMineraBuilder inicia el proceso de construcción de una Factura de Comercial Exportación Minera.
func NewComercialExportacionMineraBuilder() *comercialExportacionMineraBuilder {
	return &comercialExportacionMineraBuilder{
		factura: &documents.FacturaComercialExportacionMinera{
			XMLName:           xml.Name{Local: "facturaElectronicaComercialExportacionMinera"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercialExportacionMinera.xsd",
		},
	}
}

// NewComercialExportacionMineraCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewComercialExportacionMineraCabeceraBuilder() *comercialExportacionMineraCabeceraBuilder {
	return &comercialExportacionMineraCabeceraBuilder{
		cabecera: &documents.CabeceraComercialExportacionMinera{
			CodigoDocumentoSector: 20, // Sector 20 para Comercial Exportación Minera
		},
	}
}

// NewComercialExportacionMineraDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewComercialExportacionMineraDetalleBuilder() *comercialExportacionMineraDetalleBuilder {
	return &comercialExportacionMineraDetalleBuilder{
		detalle: &documents.DetalleComercialExportacionMinera{},
	}
}

type comercialExportacionMineraBuilder struct {
	factura *documents.FacturaComercialExportacionMinera
}

func (b *comercialExportacionMineraBuilder) WithCabecera(req ComercialExportacionMineraCabecera) *comercialExportacionMineraBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraComercialExportacionMinera](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercialExportacionMineraBuilder) AddDetalle(req ComercialExportacionMineraDetalle) *comercialExportacionMineraBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleComercialExportacionMinera](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *comercialExportacionMineraBuilder) WithModalidad(tipo int) *comercialExportacionMineraBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercialExportacionMinera"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercialExportacionMinera.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercialExportacionMinera"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercialExportacionMinera.xsd"
	}
	return b
}

func (b *comercialExportacionMineraBuilder) Build() ComercialExportacionMinera {
	return ComercialExportacionMinera{models.NewRequestWrapper(b.factura)}
}

type comercialExportacionMineraCabeceraBuilder struct {
	cabecera *documents.CabeceraComercialExportacionMinera
}

func (b *comercialExportacionMineraCabeceraBuilder) WithNitEmisor(v int64) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithRazonSocialEmisor(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithMunicipio(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithTelefono(telefono *string) *comercialExportacionMineraCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithNumeroFactura(v int64) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCuf(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCufd(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoSucursal(v int) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithDireccion(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithNombreRazonSocial(v *string) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithDireccionComprador(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithNumeroDocumento(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithComplemento(v *string) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithRuex(v *string) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.Ruex = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Ruex = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithNim(v *string) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.Nim = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Nim = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithConcentradoGranel(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.ConcentradoGranel = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithOrigen(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Origen = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithPuertoTransito(v *string) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.PuertoTransito = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.PuertoTransito = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithPuertoDestino(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.PuertoDestino = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithPaisDestino(v int) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.PaisDestino = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithIncoterm(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Incoterm = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoCliente(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoMoneda(v int) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithTipoCambio(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithTipoCambioANB(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambioANB = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithNumeroLote(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.NumeroLote = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithKilosNetosHumedos(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.KilosNetosHumedos = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithHumedadPorcentaje(v *float64) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.HumedadPorcentaje = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.HumedadPorcentaje = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithHumedadValor(v *float64) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.HumedadValor = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.HumedadValor = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithMermaPorcentaje(v *float64) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.MermaPorcentaje = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MermaPorcentaje = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithMermaValor(v *float64) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.MermaValor = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MermaValor = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithKilosNetosSecos(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.KilosNetosSecos = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoMetodoPago(v int) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithMontoTotal(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithGastosRealizacion(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.GastosRealizacion = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithOtrosDatos(v any) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.OtrosDatos = datatype.Nilable[string]{Value: nil}
		return b
	}
	var value string
	switch val := v.(type) {
	case string:
		value = val
	case *string:
		if val != nil {
			value = *val
		}
	default:
		jsonData, _ := json.Marshal(v)
		value = string(jsonData)
	}
	b.cabecera.OtrosDatos = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoExcepcion(v *int) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithCafc(v *string) *comercialExportacionMineraCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithLeyenda(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithUsuario(v string) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *comercialExportacionMineraCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *comercialExportacionMineraCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercialExportacionMineraCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *comercialExportacionMineraCabeceraBuilder) Build() ComercialExportacionMineraCabecera {
	return ComercialExportacionMineraCabecera{models.NewRequestWrapper(b.cabecera)}
}

type comercialExportacionMineraDetalleBuilder struct {
	detalle *documents.DetalleComercialExportacionMinera
}

func (b *comercialExportacionMineraDetalleBuilder) WithActividadEconomica(v string) *comercialExportacionMineraDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithCodigoProductoSin(v int64) *comercialExportacionMineraDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithCodigoProducto(v string) *comercialExportacionMineraDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithCodigoNandina(v string) *comercialExportacionMineraDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithDescripcion(v string) *comercialExportacionMineraDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithDescripcionLeyes(v string) *comercialExportacionMineraDetalleBuilder {
	b.detalle.DescripcionLeyes = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithCantidadExtraccion(v float64) *comercialExportacionMineraDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.CantidadExtraccion = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithCantidad(v float64) *comercialExportacionMineraDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithUnidadMedidaExtraccion(v int) *comercialExportacionMineraDetalleBuilder {
	b.detalle.UnidadMedidaExtraccion = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithUnidadMedida(v int) *comercialExportacionMineraDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithPrecioUnitario(v float64) *comercialExportacionMineraDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithMontoDescuento(v *float64) *comercialExportacionMineraDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) WithSubTotal(v float64) *comercialExportacionMineraDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercialExportacionMineraDetalleBuilder) Build() ComercialExportacionMineraDetalle {
	return ComercialExportacionMineraDetalle{models.NewRequestWrapper(b.detalle)}
}
