package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
)

// ServicioBasico representa la estructura completa de una factura de servicio básico lista para ser procesada.
type ServicioBasico struct {
	models.RequestWrapper[documents.FacturaServicioBasico]
}

// ServicioBasicoCabecera representa la sección de cabecera de una factura de servicio básico.
type ServicioBasicoCabecera struct {
	models.RequestWrapper[documents.CabeceraServicioBasico]
}

// ServicioBasicoDetalle representa un ítem individual dentro del detalle de una factura de servicio básico.
type ServicioBasicoDetalle struct {
	models.RequestWrapper[documents.DetalleServicioBasico]
}

type servicioBasicoBuilder struct {
	factura *documents.FacturaServicioBasico
}

func NewServicioBasicoBuilder() *servicioBasicoBuilder {
	return &servicioBasicoBuilder{
		factura: &documents.FacturaServicioBasico{
			XMLName:           xml.Name{Local: "facturaElectronicaServicioBasico"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaServicioBasico.xsd",
		},
	}
}

func (b *servicioBasicoBuilder) WithModalidad(tipo int) *servicioBasicoBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaServicioBasico"}
		b.factura.XsiSchemaLocation = "facturaElectronicaServicioBasico.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaServicioBasico"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaServicioBasico.xsd"
	}
	return b
}

func (b *servicioBasicoBuilder) WithCabecera(req ServicioBasicoCabecera) *servicioBasicoBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraServicioBasico](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *servicioBasicoBuilder) AddDetalle(req ServicioBasicoDetalle) *servicioBasicoBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleServicioBasico](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *servicioBasicoBuilder) Build() ServicioBasico {
	return ServicioBasico{models.NewRequestWrapper(b.factura)}
}

// Cabecera Builder
type servicioBasicoCabeceraBuilder struct {
	cabecera *documents.CabeceraServicioBasico
}

func NewServicioBasicoCabeceraBuilder() *servicioBasicoCabeceraBuilder {
	return &servicioBasicoCabeceraBuilder{
		cabecera: &documents.CabeceraServicioBasico{
			CodigoDocumentoSector: 13,
		},
	}
}

func (b *servicioBasicoCabeceraBuilder) WithNitEmisor(v int64) *servicioBasicoCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithRazonSocialEmisor(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithMunicipio(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithTelefono(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithNumeroFactura(v int64) *servicioBasicoCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCuf(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCufd(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCodigoSucursal(v int) *servicioBasicoCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithDireccion(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCodigoPuntoVenta(v *int) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithMes(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.Mes = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Mes = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithGestion(v *int) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.Gestion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.Gestion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCiudad(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithZona(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.Zona = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Zona = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithNumeroMedidor(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.NumeroMedidor = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithFechaEmision(v time.Time) *servicioBasicoCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithNombreRazonSocial(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithDomicilioCliente(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.DomicilioCliente = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DomicilioCliente = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *servicioBasicoCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithNumeroDocumento(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithComplemento(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCodigoCliente(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCodigoMetodoPago(v int) *servicioBasicoCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithNumeroTarjeta(v *int64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithMontoTotal(v float64) *servicioBasicoCabeceraBuilder {
	b.cabecera.MontoTotal = datatype.Float64Round(v, 2)
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *servicioBasicoCabeceraBuilder {
	b.cabecera.MontoTotalSujetoIva = datatype.Float64Round(v, 2)
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithConsumoPeriodo(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.ConsumoPeriodo = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.ConsumoPeriodo = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithBeneficiarioLey1886(v *int64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.BeneficiarioLey1886 = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.BeneficiarioLey1886 = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithMontoDescuentoLey1886(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoLey1886 = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.MontoDescuentoLey1886 = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithMontoDescuentoTarifaDignidad(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoTarifaDignidad = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.MontoDescuentoTarifaDignidad = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithTasaAseo(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.TasaAseo = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.TasaAseo = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithTasaAlumbrado(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.TasaAlumbrado = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.TasaAlumbrado = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithAjusteNoSujetoIva(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.AjusteNoSujetoIva = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.AjusteNoSujetoIva = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithDetalleAjusteNoSujetoIva(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.DetalleAjusteNoSujetoIva = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DetalleAjusteNoSujetoIva = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithAjusteSujetoIva(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.AjusteSujetoIva = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.AjusteSujetoIva = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithDetalleAjusteSujetoIva(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.DetalleAjusteSujetoIva = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DetalleAjusteSujetoIva = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithOtrosPagosNoSujetoIva(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.OtrosPagosNoSujetoIva = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.OtrosPagosNoSujetoIva = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithDetalleOtrosPagosNoSujetoIva(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.DetalleOtrosPagosNoSujetoIva = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DetalleOtrosPagosNoSujetoIva = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithOtrasTasas(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.OtrasTasas = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.OtrasTasas = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCodigoMoneda(v int) *servicioBasicoCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithTipoCambio(v float64) *servicioBasicoCabeceraBuilder {
	b.cabecera.TipoCambio = datatype.Float64Round(v, 2)
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithMontoTotalMoneda(v float64) *servicioBasicoCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = datatype.Float64Round(v, 2)
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithDescuentoAdicional(v *float64) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCodigoExcepcion(v *int) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithCafc(v *string) *servicioBasicoCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithLeyenda(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) WithUsuario(v string) *servicioBasicoCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *servicioBasicoCabeceraBuilder) Build() ServicioBasicoCabecera {
	return ServicioBasicoCabecera{models.NewRequestWrapper(b.cabecera)}
}

// Detalle Builder
type servicioBasicoDetalleBuilder struct {
	detalle *documents.DetalleServicioBasico
}

func NewServicioBasicoDetalleBuilder() *servicioBasicoDetalleBuilder {
	return &servicioBasicoDetalleBuilder{
		detalle: &documents.DetalleServicioBasico{},
	}
}

func (b *servicioBasicoDetalleBuilder) WithActividadEconomica(v string) *servicioBasicoDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *servicioBasicoDetalleBuilder) WithCodigoProductoSin(v int64) *servicioBasicoDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *servicioBasicoDetalleBuilder) WithCodigoProducto(v string) *servicioBasicoDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *servicioBasicoDetalleBuilder) WithDescripcion(v string) *servicioBasicoDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *servicioBasicoDetalleBuilder) WithCantidad(v float64) *servicioBasicoDetalleBuilder {
	b.detalle.Cantidad = datatype.Float64Round(v, 2)
	return b
}

func (b *servicioBasicoDetalleBuilder) WithUnidadMedida(v int) *servicioBasicoDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *servicioBasicoDetalleBuilder) WithPrecioUnitario(v float64) *servicioBasicoDetalleBuilder {
	b.detalle.PrecioUnitario = datatype.Float64Round(v, 2)
	return b
}

func (b *servicioBasicoDetalleBuilder) WithMontoDescuento(v *float64) *servicioBasicoDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		r := datatype.Float64Round(*v, 2)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoDetalleBuilder) WithSubTotal(v float64) *servicioBasicoDetalleBuilder {
	b.detalle.SubTotal = datatype.Float64Round(v, 2)
	return b
}

func (b *servicioBasicoDetalleBuilder) Build() ServicioBasicoDetalle {
	return ServicioBasicoDetalle{models.NewRequestWrapper(b.detalle)}
}
