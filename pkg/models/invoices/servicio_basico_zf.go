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

// ServicioBasicoZF representa la estructura completa de una factura de servicio básico ZF lista para ser procesada.
type ServicioBasicoZF struct {
	models.RequestWrapper[documents.FacturaServicioBasicoZF]
}

// ServicioBasicoZFCabecera representa la sección de cabecera de una factura de servicio básico ZF.
type ServicioBasicoZFCabecera struct {
	models.RequestWrapper[documents.CabeceraServicioBasicoZF]
}

// ServicioBasicoZFDetalle representa un ítem individual dentro del detalle de una factura de servicio básico ZF.
type ServicioBasicoZFDetalle struct {
	models.RequestWrapper[documents.DetalleServicioBasico]
}

type servicioBasicoZFBuilder struct {
	factura *documents.FacturaServicioBasicoZF
}

func NewServicioBasicoZFBuilder() *servicioBasicoZFBuilder {
	return &servicioBasicoZFBuilder{
		factura: &documents.FacturaServicioBasicoZF{
			XMLName:           xml.Name{Local: "facturaElectronicaServicioBasicoZf"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaServicioBasicoZf.xsd",
		},
	}
}

func (b *servicioBasicoZFBuilder) WithModalidad(tipo int) *servicioBasicoZFBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaServicioBasicoZf"}
		b.factura.XsiSchemaLocation = "facturaElectronicaServicioBasicoZf.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaServicioBasicoZf"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaServicioBasicoZf.xsd"
	}
	return b
}

func (b *servicioBasicoZFBuilder) WithCabecera(req ServicioBasicoZFCabecera) *servicioBasicoZFBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraServicioBasicoZF](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *servicioBasicoZFBuilder) AddDetalle(req ServicioBasicoZFDetalle) *servicioBasicoZFBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetalleServicioBasico](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *servicioBasicoZFBuilder) Build() ServicioBasicoZF {
	return ServicioBasicoZF{models.NewRequestWrapper(b.factura)}
}

// Cabecera Builder
type servicioBasicoZFCabeceraBuilder struct {
	cabecera *documents.CabeceraServicioBasicoZF
}

func NewServicioBasicoZFCabeceraBuilder() *servicioBasicoZFCabeceraBuilder {
	return &servicioBasicoZFCabeceraBuilder{
		cabecera: &documents.CabeceraServicioBasicoZF{
			CodigoDocumentoSector: 40,
			MontoTotalSujetoIva:   0,
		},
	}
}

func (b *servicioBasicoZFCabeceraBuilder) WithNitEmisor(v int64) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithRazonSocialEmisor(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithMunicipio(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithTelefono(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithNumeroFactura(v int64) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCuf(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCufd(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCodigoSucursal(v int) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithDireccion(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCodigoPuntoVenta(v *int) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithMes(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Mes = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Mes = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithGestion(v *int) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Gestion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.Gestion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCiudad(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Ciudad = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithZona(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Zona = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Zona = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithNumeroMedidor(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.NumeroMedidor = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithFechaEmision(v time.Time) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithNombreRazonSocial(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithDomicilioCliente(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DomicilioCliente = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DomicilioCliente = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithNumeroDocumento(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithComplemento(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Complemento = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCodigoCliente(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCodigoMetodoPago(v int) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithNumeroTarjeta(v *int64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithMontoTotal(v float64) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.MontoTotal = utils.Round(v, 2)
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithConsumoPeriodo(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.ConsumoPeriodo = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.ConsumoPeriodo = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithBeneficiarioLey1886(v *int64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.BeneficiarioLey1886 = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.BeneficiarioLey1886 = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithMontoDescuentoLey1886(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoLey1886 = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.MontoDescuentoLey1886 = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithMontoDescuentoTarifaDignidad(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoDescuentoTarifaDignidad = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.MontoDescuentoTarifaDignidad = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithTasaAseo(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.TasaAseo = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.TasaAseo = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithTasaAlumbrado(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.TasaAlumbrado = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.TasaAlumbrado = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithAjusteNoSujetoIva(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.AjusteNoSujetoIva = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.AjusteNoSujetoIva = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithDetalleAjusteNoSujetoIva(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DetalleAjusteNoSujetoIva = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DetalleAjusteNoSujetoIva = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithAjusteSujetoIva(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.AjusteSujetoIva = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.AjusteSujetoIva = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithDetalleAjusteSujetoIva(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DetalleAjusteSujetoIva = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DetalleAjusteSujetoIva = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithOtrosPagosNoSujetoIva(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.OtrosPagosNoSujetoIva = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.OtrosPagosNoSujetoIva = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithDetalleOtrosPagosNoSujetoIva(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DetalleOtrosPagosNoSujetoIva = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.DetalleOtrosPagosNoSujetoIva = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithOtrasTasas(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.OtrasTasas = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.OtrasTasas = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCodigoMoneda(v int) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithTipoCambio(v float64) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.TipoCambio = utils.Round(v, 2)
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithMontoTotalMoneda(v float64) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = utils.Round(v, 2)
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithDescuentoAdicional(v *float64) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCodigoExcepcion(v *int) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithCafc(v *string) *servicioBasicoZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Cafc = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithLeyenda(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) WithUsuario(v string) *servicioBasicoZFCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *servicioBasicoZFCabeceraBuilder) Build() ServicioBasicoZFCabecera {
	return ServicioBasicoZFCabecera{models.NewRequestWrapper(b.cabecera)}
}

// Detalle Builder
type servicioBasicoZFDetalleBuilder struct {
	detalle *documents.DetalleServicioBasico
}

func NewServicioBasicoZFDetalleBuilder() *servicioBasicoZFDetalleBuilder {
	return &servicioBasicoZFDetalleBuilder{
		detalle: &documents.DetalleServicioBasico{},
	}
}

func (b *servicioBasicoZFDetalleBuilder) WithActividadEconomica(v string) *servicioBasicoZFDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithCodigoProductoSin(v int64) *servicioBasicoZFDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithCodigoProducto(v string) *servicioBasicoZFDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithDescripcion(v string) *servicioBasicoZFDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithCantidad(v float64) *servicioBasicoZFDetalleBuilder {
	b.detalle.Cantidad = utils.Round(v, 2)
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithUnidadMedida(v int) *servicioBasicoZFDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithPrecioUnitario(v float64) *servicioBasicoZFDetalleBuilder {
	b.detalle.PrecioUnitario = utils.Round(v, 2)
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithMontoDescuento(v *float64) *servicioBasicoZFDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		r := utils.Round(*v, 2)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &r}
	}
	return b
}

func (b *servicioBasicoZFDetalleBuilder) WithSubTotal(v float64) *servicioBasicoZFDetalleBuilder {
	b.detalle.SubTotal = utils.Round(v, 2)
	return b
}

func (b *servicioBasicoZFDetalleBuilder) Build() ServicioBasicoZFDetalle {
	return ServicioBasicoZFDetalle{models.NewRequestWrapper(b.detalle)}
}
