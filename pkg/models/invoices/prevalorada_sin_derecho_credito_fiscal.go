package invoices

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
)

type FacturaPrevaloradaSinDerechoCreditoFiscal struct {
	models.RequestWrapper[documents.FacturaPrevaloradaSinDerechoCreditoFiscal]
}

type FacturaPrevaloradaSinDerechoCreditoFiscalCabecera struct {
	models.RequestWrapper[documents.CabeceraPrevaloradaSinDerechoCreditoFiscal]
}

type FacturaPrevaloradaSinDerechoCreditoFiscalDetalle struct {
	models.RequestWrapper[documents.DetallePrevaloradaSinDerechoCreditoFiscal]
}

func NewPrevaloradaSinDerechoCreditoFiscalBuilder() *prevaloradaSinDerechoCreditoFiscalBuilder {
	return &prevaloradaSinDerechoCreditoFiscalBuilder{
		factura: &documents.FacturaPrevaloradaSinDerechoCreditoFiscal{
			XMLName:           xml.Name{Local: "facturaElectronicaPrevaloradaSD"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaPrevaloradaSD.xsd",
		},
	}
}

type prevaloradaSinDerechoCreditoFiscalBuilder struct {
	factura *documents.FacturaPrevaloradaSinDerechoCreditoFiscal
}

func (b *prevaloradaSinDerechoCreditoFiscalBuilder) WithCabecera(req FacturaPrevaloradaSinDerechoCreditoFiscalCabecera) *prevaloradaSinDerechoCreditoFiscalBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraPrevaloradaSinDerechoCreditoFiscal](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalBuilder) WithModalidad(tipo int) *prevaloradaSinDerechoCreditoFiscalBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaPrevaloradaSD"}
		b.factura.XsiSchemaLocation = "facturaElectronicaPrevaloradaSD.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaPrevaloradaSD"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaPrevaloradaSD.xsd"
	}
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalBuilder) WithDetalle(req FacturaPrevaloradaSinDerechoCreditoFiscalDetalle) *prevaloradaSinDerechoCreditoFiscalBuilder {
	if internal := models.UnwrapInternalRequest[documents.DetallePrevaloradaSinDerechoCreditoFiscal](req); internal != nil {
		b.factura.Detalle = *internal
	}
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalBuilder) Build() FacturaPrevaloradaSinDerechoCreditoFiscal {
	return FacturaPrevaloradaSinDerechoCreditoFiscal{models.NewRequestWrapper(b.factura)}
}

type prevaloradaSinDerechoCreditoFiscalCabeceraBuilder struct {
	cabecera *documents.CabeceraPrevaloradaSinDerechoCreditoFiscal
}

func NewPrevaloradaSinDerechoCreditoFiscalCabeceraBuilder() *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	return &prevaloradaSinDerechoCreditoFiscalCabeceraBuilder{
		cabecera: &documents.CabeceraPrevaloradaSinDerechoCreditoFiscal{
			CodigoDocumentoSector:        36,
			NombreRazonSocial:            "S/N",
			CodigoTipoDocumentoIdentidad: 5,
			NumeroDocumento:              0,
			CodigoCliente:                "N/A",
			MontoTotalSujetoIva:          0,
		},
	}
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithNitEmisor(v int64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithRazonSocialEmisor(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithMunicipio(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithTelefono(v *string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	if v == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.Telefono = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithNumeroFactura(v int64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCuf(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCufd(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCodigoSucursal(v int) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithDireccion(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCodigoPuntoVenta(v *int) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithFechaEmision(v time.Time) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithNombreRazonSocial(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.NombreRazonSocial = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithNumeroDocumento(v int64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCodigoCliente(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCodigoMetodoPago(v int) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithNumeroTarjeta(v *int64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithMontoTotal(v float64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithCodigoMoneda(v int) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithTipoCambio(v float64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithMontoTotalMoneda(v float64) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithLeyenda(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) WithUsuario(v string) *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalCabeceraBuilder) Build() FacturaPrevaloradaSinDerechoCreditoFiscalCabecera {
	return FacturaPrevaloradaSinDerechoCreditoFiscalCabecera{models.NewRequestWrapper(b.cabecera)}
}

type prevaloradaSinDerechoCreditoFiscalDetalleBuilder struct {
	detalle *documents.DetallePrevaloradaSinDerechoCreditoFiscal
}

func NewPrevaloradaSinDerechoCreditoFiscalDetalleBuilder() *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	return &prevaloradaSinDerechoCreditoFiscalDetalleBuilder{
		detalle: &documents.DetallePrevaloradaSinDerechoCreditoFiscal{},
	}
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithActividadEconomica(v string) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithCodigoProductoSin(v int64) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithCodigoProducto(v string) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithDescripcion(v string) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithCantidad(v float64) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithUnidadMedida(v int) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithPrecioUnitario(v float64) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithMontoDescuento(v *float64) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
	} else {
		val := *v
		val, _ = strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64), 64)
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &val}
	}
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) WithSubTotal(v float64) *prevaloradaSinDerechoCreditoFiscalDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalDetalleBuilder) Build() FacturaPrevaloradaSinDerechoCreditoFiscalDetalle {
	return FacturaPrevaloradaSinDerechoCreditoFiscalDetalle{models.NewRequestWrapper(b.detalle)}
}
