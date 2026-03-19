package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// ComercialExportacionServicio representa la estructura completa de una factura comercial
// de exportación de servicios lista para ser procesada.
type ComercialExportacionServicio struct {
	requestWrapper[documentos.FacturaComercialExportacionServicio]
}

// ComercialExportacionServicioCabecera representa la sección de cabecera de la factura.
type ComercialExportacionServicioCabecera struct {
	requestWrapper[documentos.CabeceraComercialExportacionServicio]
}

// ComercialExportacionServicioDetalle representa un ítem individual dentro del detalle.
type ComercialExportacionServicioDetalle struct {
	requestWrapper[documentos.DetalleComercialExportacionServicio]
}

// NewComercialExportacionServicioBuilder inicia el proceso de construcción de la factura.
func NewComercialExportacionServicioBuilder() *comercialExportacionServicioBuilder {
	return &comercialExportacionServicioBuilder{
		factura: &documentos.FacturaComercialExportacionServicio{
			XMLName:           xml.Name{Local: "facturaElectronicaComercialExportacionServicio"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaComercialExportacionServicio.xsd",
		},
	}
}

// NewComercialExportacionServicioCabeceraBuilder crea el constructor para la cabecera.
func NewComercialExportacionServicioCabeceraBuilder() *comercialExportacionServicioCabeceraBuilder {
	return &comercialExportacionServicioCabeceraBuilder{
		cabecera: &documentos.CabeceraComercialExportacionServicio{
			MontoTotalSujetoIva:   0,  // Fixed in XSD
			CodigoDocumentoSector: 28, // Fixed in XSD
		},
	}
}

// NewComercialExportacionServicioDetalleBuilder crea el constructor para los detalles.
func NewComercialExportacionServicioDetalleBuilder() *comercialExportacionServicioDetalleBuilder {
	return &comercialExportacionServicioDetalleBuilder{
		detalle: &documentos.DetalleComercialExportacionServicio{
			Cantidad: 1, // Fixed in XSD
		},
	}
}

type comercialExportacionServicioBuilder struct {
	factura *documentos.FacturaComercialExportacionServicio
}

func (b *comercialExportacionServicioBuilder) WithCabecera(req ComercialExportacionServicioCabecera) *comercialExportacionServicioBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

func (b *comercialExportacionServicioBuilder) AddDetalle(req ComercialExportacionServicioDetalle) *comercialExportacionServicioBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

func (b *comercialExportacionServicioBuilder) WithModalidad(tipo int) *comercialExportacionServicioBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaComercialExportacionServicio"}
		b.factura.XsiSchemaLocation = "facturaElectronicaComercialExportacionServicio.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaComercialExportacionServicio"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaComercialExportacionServicio.xsd"
	}
	return b
}

func (b *comercialExportacionServicioBuilder) Build() ComercialExportacionServicio {
	return ComercialExportacionServicio{requestWrapper[documentos.FacturaComercialExportacionServicio]{request: b.factura}}
}

type comercialExportacionServicioCabeceraBuilder struct {
	cabecera *documentos.CabeceraComercialExportacionServicio
}

func (b *comercialExportacionServicioCabeceraBuilder) WithNitEmisor(v int64) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithRazonSocialEmisor(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithMunicipio(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithTelefono(telefono *string) *comercialExportacionServicioCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithNumeroFactura(v int64) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCuf(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCufd(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoSucursal(v int) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithDireccion(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoPuntoVenta(v *int) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithFechaEmision(v time.Time) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithNombreRazonSocial(v *string) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithNumeroDocumento(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithComplemento(v *string) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithDireccionComprador(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.DireccionComprador = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoCliente(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithLugarDestino(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.LugarDestino = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoPais(v int) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.CodigoPais = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoMetodoPago(v int) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithNumeroTarjeta(v *int64) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithMontoTotal(v float64) *comercialExportacionServicioCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoMoneda(v int) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithTipoCambio(v float64) *comercialExportacionServicioCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithMontoTotalMoneda(v float64) *comercialExportacionServicioCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithInformacionAdicional(v *string) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.InformacionAdicional = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithDescuentoAdicional(v *float64) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCodigoExcepcion(v *int) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithCafc(v *string) *comercialExportacionServicioCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithLeyenda(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) WithUsuario(v string) *comercialExportacionServicioCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *comercialExportacionServicioCabeceraBuilder) Build() ComercialExportacionServicioCabecera {
	return ComercialExportacionServicioCabecera{requestWrapper[documentos.CabeceraComercialExportacionServicio]{request: b.cabecera}}
}

type comercialExportacionServicioDetalleBuilder struct {
	detalle *documentos.DetalleComercialExportacionServicio
}

func (b *comercialExportacionServicioDetalleBuilder) WithActividadEconomica(v string) *comercialExportacionServicioDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) WithCodigoProductoSin(v int64) *comercialExportacionServicioDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) WithCodigoProducto(v string) *comercialExportacionServicioDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) WithDescripcion(v string) *comercialExportacionServicioDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) WithUnidadMedida(v int) *comercialExportacionServicioDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) WithPrecioUnitario(v float64) *comercialExportacionServicioDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) WithMontoDescuento(v *float64) *comercialExportacionServicioDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) WithSubTotal(v float64) *comercialExportacionServicioDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *comercialExportacionServicioDetalleBuilder) Build() ComercialExportacionServicioDetalle {
	return ComercialExportacionServicioDetalle{requestWrapper[documentos.DetalleComercialExportacionServicio]{request: b.detalle}}
}
