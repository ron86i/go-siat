package invoices

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/shopspring/decimal"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
)

// ComercialExportacionPVenta representa la estructura completa de una factura PVenta lista para ser procesada.
type ComercialExportacionPVenta struct {
	models.MetadatosFactura
	models.RequestWrapper[documents.FacturaComercialExportacionPVenta]
}

// ComercialExportacionPVentaCabecera representa la sección de cabecera de una factura PVenta.
type ComercialExportacionPVentaCabecera struct {
	models.RequestWrapper[documents.CabeceraComercialExportacionPVenta]
}

// ComercialExportacionPVentaDetalle representa un ítem individual dentro del detalle de una factura PVenta.
type ComercialExportacionPVentaDetalle struct {
	models.RequestWrapper[documents.DetalleComercialExportacionPVenta]
}

// NewComercialExportacionPVentaBuilder inicia el proceso de construcción de una Factura PVenta.
func NewComercialExportacionPVentaBuilder() *comercialExportacionPVentaBuilder {
	return &comercialExportacionPVentaBuilder{
		metadatos: models.NuevosMetadatosFactura(45),
		factura: &documents.FacturaComercialExportacionPVenta{
			XMLName:           xml.Name{Local: "facturaElectronicaComercialExportacionPVenta"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercialExportacionPVenta.xsd",
		},
	}
}

// NewComercialExportacionPVentaCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewComercialExportacionPVentaCabeceraBuilder() *comercialExportacionPVentaCabeceraBuilder {
	return &comercialExportacionPVentaCabeceraBuilder{
		cabecera: &documents.CabeceraComercialExportacionPVenta{
			CodigoDocumentoSector: 45, // Sector 45 para Comercial Exportación PVenta
		},
	}
}

// NewComercialExportacionPVentaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewComercialExportacionPVentaDetalleBuilder() *comercialExportacionPVentaDetalleBuilder {
	return &comercialExportacionPVentaDetalleBuilder{
		detalle: &documents.DetalleComercialExportacionPVenta{},
	}
}

type comercialExportacionPVentaBuilder struct {
	metadatos models.MetadatosFactura
	factura   *documents.FacturaComercialExportacionPVenta
}

func (b *comercialExportacionPVentaBuilder) WithCabecera(req ComercialExportacionPVentaCabecera) *comercialExportacionPVentaBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraComercialExportacionPVenta](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercialExportacionPVentaBuilder) AddDetalle(req ComercialExportacionPVentaDetalle) *comercialExportacionPVentaBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleComercialExportacionPVenta](req); internal != nil {
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
	return ComercialExportacionPVenta{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type comercialExportacionPVentaCabeceraBuilder struct {
	cabecera *documents.CabeceraComercialExportacionPVenta
	err      error
}

func (b *comercialExportacionPVentaCabeceraBuilder) Build() ComercialExportacionPVentaCabecera {
	return ComercialExportacionPVentaCabecera{models.NewRequestWrapper(b.cabecera)}
}

// Err devuelve el último error de serialización de un campo JSON libre.
func (b *comercialExportacionPVentaCabeceraBuilder) Err() error {
	return b.err
}

type comercialExportacionPVentaDetalleBuilder struct {
	detalle *documents.DetalleComercialExportacionPVenta
}

func (b *comercialExportacionPVentaDetalleBuilder) Build() ComercialExportacionPVentaDetalle {
	return ComercialExportacionPVentaDetalle{models.NewRequestWrapper(b.detalle)}
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

func (b *comercialExportacionPVentaCabeceraBuilder) WithTelefono(v *string) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
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
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithFechaEmision(v time.Time) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithNombreRazonSocial(v *string) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
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

func (b *comercialExportacionPVentaCabeceraBuilder) WithComplemento(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Complemento = v
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

func (b *comercialExportacionPVentaCabeceraBuilder) WithAlcanceIncoterm(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.AlcanceIncoterm = v
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
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithMontoAPagar(v float64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
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
		jsonData, err := json.Marshal(v)
		if err != nil {
			b.err = err
			return b
		}
		valStr := string(jsonData)
		b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: &valStr}
	}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTotalGastosNacionalesFob(v float64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.TotalGastosNacionalesFob = decimal.NewFromFloat(v).Round(2).InexactFloat64()
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
		jsonData, err := json.Marshal(v)
		if err != nil {
			b.err = err
			return b
		}
		valStr := string(jsonData)
		b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: &valStr}
	}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTotalGastosInternacionales(v *float64) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithPrecioValorBruto(v *float64) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.PrecioValorBruto = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.PrecioValorBruto = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithSubTotalNeto(v float64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTotalPrecioNeto(v float64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithIva(v float64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoMoneda(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithTipoCambio(v float64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
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
		jsonData, err := json.Marshal(v)
		if err != nil {
			b.err = err
			return b
		}
		valStr := string(jsonData)
		b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: &valStr}
	}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithInformacionAdicional(v *string) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.InformacionAdicional = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoExcepcion(v *int) *comercialExportacionPVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
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

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoInterno(v string) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *comercialExportacionPVentaCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercialExportacionPVentaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
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
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithUnidadMedida(v int) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithPrecioNetoUnitario(v float64) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithMontoDescuento(v *float64) *comercialExportacionPVentaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionPVentaDetalleBuilder) WithSubTotal(v float64) *comercialExportacionPVentaDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}
