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

// ComercialExportacion representa la estructura completa de una factura comercial
// de exportación lista para ser procesada.
type ComercialExportacion struct {
	models.RequestWrapper[documents.FacturaComercialExportacion]
	models.MetadatosFactura
}

// ComercialExportacionCabecera representa la sección de cabecera de la factura.
type ComercialExportacionCabecera struct {
	models.RequestWrapper[documents.CabeceraComercialExportacion]
}

// ComercialExportacionDetalle representa un ítem individual dentro del detalle.
type ComercialExportacionDetalle struct {
	models.RequestWrapper[documents.DetalleComercialExportacion]
}

// NewComercialExportacionBuilder inicia el proceso de construcción de la factura.
func NewComercialExportacionBuilder() *comercialExportacionBuilder {
	return &comercialExportacionBuilder{
		metadatos: models.NuevosMetadatosFactura(3),
		factura: &documents.FacturaComercialExportacion{
			XMLName:           xml.Name{Local: "facturaElectronicaComercialExportacion"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercialExportacion.xsd",
		},
	}
}

// NewComercialExportacionCabeceraBuilder crea el constructor para la cabecera.
func NewComercialExportacionCabeceraBuilder() *comercialExportacionCabeceraBuilder {
	return &comercialExportacionCabeceraBuilder{
		cabecera: &documents.CabeceraComercialExportacion{
			CodigoDocumentoSector: 3, // Fixed in XSD
		},
	}
}

// NewComercialExportacionDetalleBuilder crea el constructor para los detalles.
func NewComercialExportacionDetalleBuilder() *comercialExportacionDetalleBuilder {
	return &comercialExportacionDetalleBuilder{
		detalle: &documents.DetalleComercialExportacion{},
	}
}

type comercialExportacionBuilder struct {
	factura   *documents.FacturaComercialExportacion
	metadatos models.MetadatosFactura
}

// WithTipoFacturaDocumento reemplaza el tipo fiscal predeterminado de exportación.
func (b *comercialExportacionBuilder) WithTipoFacturaDocumento(tipo int) *comercialExportacionBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercialExportacionBuilder) WithCabecera(req ComercialExportacionCabecera) *comercialExportacionBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraComercialExportacion](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *comercialExportacionBuilder) AddDetalle(req ComercialExportacionDetalle) *comercialExportacionBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleComercialExportacion](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *comercialExportacionBuilder) WithModalidad(tipo int) *comercialExportacionBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercialExportacion"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercialExportacion.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercialExportacion"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercialExportacion.xsd"
	}
	return b
}

func (b *comercialExportacionBuilder) Build() ComercialExportacion {
	return ComercialExportacion{RequestWrapper: models.NewRequestWrapper(b.factura), MetadatosFactura: b.metadatos}
}

type comercialExportacionCabeceraBuilder struct {
	cabecera *documents.CabeceraComercialExportacion
	err      error
}

func (b *comercialExportacionCabeceraBuilder) Build() ComercialExportacionCabecera {
	return ComercialExportacionCabecera{models.NewRequestWrapper(b.cabecera)}
}

// Err devuelve el último error de serialización detectado durante la construcción.
func (b *comercialExportacionCabeceraBuilder) Err() error {
	return b.err
}

type comercialExportacionDetalleBuilder struct {
	detalle *documents.DetalleComercialExportacion
}

func (b *comercialExportacionDetalleBuilder) Build() ComercialExportacionDetalle {
	return ComercialExportacionDetalle{models.NewRequestWrapper(b.detalle)}
}

func (b *comercialExportacionCabeceraBuilder) WithNitEmisor(v int64) *comercialExportacionCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithRazonSocialEmisor(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithMunicipio(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTelefono(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroFactura(v int64) *comercialExportacionCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCuf(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCufd(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoSucursal(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithDireccion(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithFechaEmision(v time.Time) *comercialExportacionCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNombreRazonSocial(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroDocumento(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithComplemento(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Complemento = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithDireccionComprador(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoCliente(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithIncoterm(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Incoterm = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithIncotermAlcance(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.IncotermAlcance = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithIncotermDetalle(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.IncotermDetalle = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithPuertoDestino(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.PuertoDestino = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithLugarDestino(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.LugarDestino = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoPais(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoPais = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoMetodoPago(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTotalPrecioNeto(v float64) *comercialExportacionCabeceraBuilder {
	b.cabecera.TotalPrecioNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithMontoAPagar(v float64) *comercialExportacionCabeceraBuilder {
	b.cabecera.MontoAPagar = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCostosGastosNacionales(v map[string]any) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		b.err = err
		return b
	}
	val := string(jsonData)
	b.cabecera.CostosGastosNacionales = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTotalGastosNacionalesFob(v float64) *comercialExportacionCabeceraBuilder {
	b.cabecera.TotalGastosNacionalesFob = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCostosGastosInternacionales(v map[string]any) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: nil}
		return b
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		b.err = err
		return b
	}
	val := string(jsonData)
	b.cabecera.CostosGastosInternacionales = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTotalGastosInternacionales(v *float64) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.TotalGastosInternacionales = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithSubTotalNeto(v float64) *comercialExportacionCabeceraBuilder {
	b.cabecera.SubTotalNeto = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithIva(v float64) *comercialExportacionCabeceraBuilder {
	b.cabecera.Iva = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoMoneda(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithTipoCambio(v float64) *comercialExportacionCabeceraBuilder {
	b.cabecera.TipoCambio = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercialExportacionCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = decimal.NewFromFloat(v).Round(2).InexactFloat64()
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithNumeroDescripcionPaquetesBultos(v map[string]any) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: nil}
		return b
	}
	jsonData, err := json.Marshal(v)
	if err != nil {
		b.err = err
		return b
	}
	val := string(jsonData)
	b.cabecera.NumeroDescripcionPaquetesBultos = datatype.Nilable[string]{Value: &val}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithInformacionAdicional(v *string) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.InformacionAdicional = datatype.Nilable[string]{}
		return b
	}
	value := *v
	b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(2).InexactFloat64()
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoExcepcion(v *int) *comercialExportacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithLeyenda(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithUsuario(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoInterno(v string) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoInterno = v
	return b
}

func (b *comercialExportacionCabeceraBuilder) WithCodigoDocumentoSector(v int) *comercialExportacionCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithActividadEconomica(v string) *comercialExportacionDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCodigoProductoSin(v int64) *comercialExportacionDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCodigoProducto(v string) *comercialExportacionDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCodigoNandina(v string) *comercialExportacionDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithDescripcion(v string) *comercialExportacionDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithCantidad(v float64) *comercialExportacionDetalleBuilder {
	b.detalle.Cantidad = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *comercialExportacionDetalleBuilder) WithUnidadMedida(v int) *comercialExportacionDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercialExportacionDetalleBuilder) WithPrecioNetoUnitario(v float64) *comercialExportacionDetalleBuilder {
	b.detalle.PrecioNetoUnitario = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}

func (b *comercialExportacionDetalleBuilder) WithMontoDescuento(v *float64) *comercialExportacionDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{}
		return b
	}
	value := *v
	value = decimal.NewFromFloat(value).Round(5).InexactFloat64()
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionDetalleBuilder) WithSubTotal(v float64) *comercialExportacionDetalleBuilder {
	b.detalle.SubTotal = decimal.NewFromFloat(v).Round(5).InexactFloat64()
	return b
}
