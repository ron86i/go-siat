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

// LibreConsignacion representa la estructura completa de una factura Libre Consignación lista para ser procesada.
type LibreConsignacion struct {
	models.RequestWrapper[documentos.FacturaLibreConsignacion]
}

// LibreConsignacionCabecera representa la sección de cabecera de una factura Libre Consignación.
type LibreConsignacionCabecera struct {
	models.RequestWrapper[documentos.CabeceraLibreConsignacion]
}

// LibreConsignacionDetalle representa un ítem individual dentro del detalle de una factura Libre Consignación.
type LibreConsignacionDetalle struct {
	models.RequestWrapper[documentos.DetalleLibreConsignacion]
}

// NewLibreConsignacionBuilder inicia el proceso de construcción de una Factura Libre Consignación.
func NewLibreConsignacionBuilder() *libreConsignacionBuilder {
	return &libreConsignacionBuilder{
		factura: &documentos.FacturaLibreConsignacion{
			XMLName:           xml.Name{Local: "facturaElectronicaLibreConsignacion"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaLibreConsignacion.xsd",
		},
	}
}

// NewLibreConsignacionCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewLibreConsignacionCabeceraBuilder() *libreConsignacionCabeceraBuilder {
	return &libreConsignacionCabeceraBuilder{
		cabecera: &documentos.CabeceraLibreConsignacion{
			CodigoDocumentoSector:        4,                  // Sector 4 para Libre Consignación
			NombreRazonSocial:            "SIN RAZON SOCIAL", // Fijo según XSD
			CodigoTipoDocumentoIdentidad: 5,                  // Fijo según XSD (NIT/Extranjero)
			NumeroDocumento:              "0",                // Fijo según XSD
			MontoTotalSujetoIva:          0,                  // Fijo según XSD
		},
	}
}

// NewLibreConsignacionDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewLibreConsignacionDetalleBuilder() *libreConsignacionDetalleBuilder {
	return &libreConsignacionDetalleBuilder{
		detalle: &documentos.DetalleLibreConsignacion{
			MontoDescuento: 0, // Fijo según XSD
		},
	}
}

type libreConsignacionBuilder struct {
	factura *documentos.FacturaLibreConsignacion
}

func (b *libreConsignacionBuilder) WithCabecera(req LibreConsignacionCabecera) *libreConsignacionBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraLibreConsignacion](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *libreConsignacionBuilder) AddDetalle(req LibreConsignacionDetalle) *libreConsignacionBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleLibreConsignacion](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *libreConsignacionBuilder) WithModalidad(tipo int) *libreConsignacionBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaLibreConsignacion"}
		b.factura.XsiSchemaLocation = "facturaElectronicaLibreConsignacion.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaLibreConsignacion"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaLibreConsignacion.xsd"
	}
	return b
}

func (b *libreConsignacionBuilder) Build() LibreConsignacion {
	return LibreConsignacion{models.NewRequestWrapper(b.factura)}
}

type libreConsignacionCabeceraBuilder struct {
	cabecera *documentos.CabeceraLibreConsignacion
}

func (b *libreConsignacionCabeceraBuilder) WithNitEmisor(v int64) *libreConsignacionCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithRazonSocialEmisor(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithMunicipio(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithTelefono(telefono *string) *libreConsignacionCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithNumeroFactura(v int64) *libreConsignacionCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCuf(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCufd(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoSucursal(v int) *libreConsignacionCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithDireccion(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoPuntoVenta(v *int) *libreConsignacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *libreConsignacionCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoCliente(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoPais(v int) *libreConsignacionCabeceraBuilder {
	b.cabecera.CodigoPais = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithPuertoDestino(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.PuertoDestino = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoMetodoPago(v int) *libreConsignacionCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithNumeroTarjeta(v *int64) *libreConsignacionCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithMontoTotal(v float64) *libreConsignacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoMoneda(v int) *libreConsignacionCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithTipoCambio(v float64) *libreConsignacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithMontoTotalMoneda(v float64) *libreConsignacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoExcepcion(v *int) *libreConsignacionCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCafc(v *string) *libreConsignacionCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithLeyenda(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithUsuario(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *libreConsignacionCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithNumeroDocumento(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithNombreRazonSocial(v string) *libreConsignacionCabeceraBuilder {
	b.cabecera.NombreRazonSocial = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *libreConsignacionCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *libreConsignacionCabeceraBuilder) WithCodigoDocumentoSector(v int) *libreConsignacionCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *libreConsignacionCabeceraBuilder) Build() LibreConsignacionCabecera {
	return LibreConsignacionCabecera{models.NewRequestWrapper(b.cabecera)}
}

type libreConsignacionDetalleBuilder struct {
	detalle *documentos.DetalleLibreConsignacion
}

func (b *libreConsignacionDetalleBuilder) WithActividadEconomica(v string) *libreConsignacionDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithCodigoProductoSin(v int64) *libreConsignacionDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithCodigoProducto(v string) *libreConsignacionDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithCodigoNandina(v string) *libreConsignacionDetalleBuilder {
	b.detalle.CodigoNandina = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithDescripcion(v string) *libreConsignacionDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithUnidadMedida(v int) *libreConsignacionDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithCantidad(v float64) *libreConsignacionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithPrecioUnitario(v float64) *libreConsignacionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithSubTotal(v float64) *libreConsignacionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *libreConsignacionDetalleBuilder) WithMontoDescuento(v float64) *libreConsignacionDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = v
	return b
}

func (b *libreConsignacionDetalleBuilder) Build() LibreConsignacionDetalle {
	return LibreConsignacionDetalle{models.NewRequestWrapper(b.detalle)}
}
