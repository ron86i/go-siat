package facturas

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// ComercialExportacionPVenta representa la estructura completa de una factura PVenta lista para ser procesada.
type ComercialExportacionPVenta struct {
	models.RequestWrapper[documentos.FacturaComercialExportacionPVenta]
}

// ComercialExportacionPVentaCabecera representa la sección de cabecera de una factura PVenta.
type ComercialExportacionPVentaCabecera struct {
	models.RequestWrapper[documentos.CabeceraComercialExportacionPVenta]
}

// ComercialExportacionPVentaDetalle representa un ítem individual dentro del detalle de una factura PVenta.
type ComercialExportacionPVentaDetalle struct {
	models.RequestWrapper[documentos.DetalleComercialExportacionPVenta]
}

// NewComercialExportacionPVentaBuilder inicia el proceso de construcción de una Factura PVenta.
func NewComercialExportacionPVentaBuilder() *comercialExportacionPVentaBuilder {
	return &comercialExportacionPVentaBuilder{
		factura: &documentos.FacturaComercialExportacionPVenta{
			XMLName:           xml.Name{Local: "facturaElectronicaComercialExportacionPVenta"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercialExportacionPVenta.xsd",
		},
	}
}

// NewComercialExportacionPVentaCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewComercialExportacionPVentaCabeceraBuilder() *comercialExportacionPVentaCabeceraBuilder {
	return &comercialExportacionPVentaCabeceraBuilder{
		cabecera: &documentos.CabeceraComercialExportacionPVenta{
			CodigoDocumentoSector: 45, // Sector 45 para Comercial Exportación PVenta
			MontoTotalSujetoIva:   0,  // Siempre 0 para exportación
		},
	}
}

// NewComercialExportacionPVentaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewComercialExportacionPVentaDetalleBuilder() *comercialExportacionPVentaDetalleBuilder {
	return &comercialExportacionPVentaDetalleBuilder{
		detalle: &documentos.DetalleComercialExportacionPVenta{},
	}
}

type comercialExportacionPVentaBuilder struct {
	factura *documentos.FacturaComercialExportacionPVenta
}

func (b *comercialExportacionPVentaBuilder) WithCabecera(req ComercialExportacionPVentaCabecera) *comercialExportacionPVentaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraComercialExportacionPVenta](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercialExportacionPVentaBuilder) AddDetalle(req ComercialExportacionPVentaDetalle) *comercialExportacionPVentaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleComercialExportacionPVenta](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *comercialExportacionPVentaBuilder) WithModalidad(tipo int) *comercialExportacionPVentaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercialExportacionPVenta"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercialExportacionPVenta.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercialExportacionPVenta"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercialExportacionPVenta.xsd"
	}
	return b
}

func (b *comercialExportacionPVentaBuilder) Build() ComercialExportacionPVenta {
	return ComercialExportacionPVenta{models.NewRequestWrapper(b.factura)}
}

type comercialExportacionPVentaCabeceraBuilder struct {
	cabecera *documentos.CabeceraComercialExportacionPVenta
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithNitEmisor(v int64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithRazonSocialEmisor(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithMunicipio(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTelefono(telefono *string) *comercialExportacionPVentaCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithNumeroFactura(v int64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCuf(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCufd(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoSucursal(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithDireccion(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithNombreRazonSocial(v *string) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithNumeroDocumento(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithComplemento(v *string) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithDireccionComprador(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoCliente(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithIncoterm(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Incoterm = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithIncotermDetalle(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.IncotermDetalle = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithPuertoDestino(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.PuertoDestino = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithLugarDestino(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.LugarDestino = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoPais(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoPais = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoMetodoPago(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithMontoTotal(v float64) *comercialExportacionPVentaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCostosGastosNacionales(v any) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	switch val := v.(type) {
	case string:
		b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: &val}
	default:
		jsonData, _ := json.Marshal(v)
		valStr := string(jsonData)
		b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: &valStr}
	}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTotalGastosNacionalesFob(v float64) *comercialExportacionPVentaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TotalGastosNacionalesFob = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCostosGastosInternacionales(v any) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	switch val := v.(type) {
	case string:
		b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: &val}
	default:
		jsonData, _ := json.Marshal(v)
		valStr := string(jsonData)
		b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: &valStr}
	}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTotalGastosInternacionales(v *float64) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithPrecioValorBruto(v *float64) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.PrecioValorBruto = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.PrecioValorBruto = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithMontoDetalle(v float64) *comercialExportacionPVentaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoDetalle = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoMoneda(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTipoCambio(v float64) *comercialExportacionPVentaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercialExportacionPVentaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithNumeroDescripcionPaquetesBultos(v any) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: nil}
		return b
	}
	switch val := v.(type) {
	case string:
		b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: &val}
	default:
		jsonData, _ := json.Marshal(v)
		valStr := string(jsonData)
		b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: &valStr}
	}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithInformacionAdicional(v *string) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoExcepcion(v *int) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCafc(v *string) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithLeyenda(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithUsuario(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *comercialExportacionPVentaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) Build() ComercialExportacionPVentaCabecera {
	return ComercialExportacionPVentaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type comercialExportacionPVentaDetalleBuilder struct {
	detalle *documentos.DetalleComercialExportacionPVenta
}

func (b *comercialExportacionPVentaDetalleBuilder) WithActividadEconomica(v string) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithCodigoProductoSin(v int64) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithCodigoProducto(v string) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithCodigoNandina(v string) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithDescripcion(v string) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithCantidad(v float64) *comercialExportacionPVentaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithUnidadMedida(v int) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithPrecioUnitario(v float64) *comercialExportacionPVentaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithMontoDescuento(v *float64) *comercialExportacionPVentaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithSubTotal(v float64) *comercialExportacionPVentaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) Build() ComercialExportacionPVentaDetalle {
	return ComercialExportacionPVentaDetalle{models.NewRequestWrapper(b.detalle)}
}
