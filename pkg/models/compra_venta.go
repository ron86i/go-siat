package models

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/compra_venta"
)

// --- Interfaces opacas para restringir el acceso a los atributos ---

// AnulacionFacturaRequest representa una solicitud para anular una factura emitida.
type AnulacionFacturaRequest interface{ commonRequest() }

// RecepcionFacturaRequest representa una solicitud para el envío de una factura al SIAT.
type RecepcionFacturaRequest interface{ commonRequest() }

// FacturaRequest representa la estructura de una factura lista para ser procesada.
type FacturaRequest interface{ commonRequest() }

// CabeceraRequest representa la cabecera de una factura.
type CabeceraRequest interface{ commonRequest() }

// DetalleRequest representa un ítem de detalle dentro de una factura.
type DetalleRequest interface{ commonRequest() }

// requestWrapper satisface todas estas interfaces mediante el método commonRequest() en common.go

type compraVentaNamespace struct{}

// CompraVenta expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
var CompraVenta = compraVentaNamespace{}

// --- Builders para la creación de solicitudes ---

// NewAnulacionFacturaRequest inicia la construcción de una solicitud de anulación.
func (compraVentaNamespace) NewAnulacionFacturaRequest() *AnulacionFacturaBuilder {
	return &AnulacionFacturaBuilder{
		request: &compra_venta.AnulacionFactura{},
	}
}

type AnulacionFacturaBuilder struct {
	request *compra_venta.AnulacionFactura
}

func (b *AnulacionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoEmision(codigoEmision int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoSistema(codigoSistema string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

func (b *AnulacionFacturaBuilder) WithCufd(cufd string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

func (b *AnulacionFacturaBuilder) WithCuf(cuf string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

func (b *AnulacionFacturaBuilder) WithCuis(cuis string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

func (b *AnulacionFacturaBuilder) WithNit(nit int64) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

func (b *AnulacionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoMotivo(codigoMotivo int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = codigoMotivo
	return b
}

func (b *AnulacionFacturaBuilder) Build() AnulacionFacturaRequest {
	return requestWrapper[compra_venta.AnulacionFactura]{request: b.request}
}

// NewRecepcionFacturaRequest inicia la construcción de una solicitud de recepción de factura.
func (compraVentaNamespace) NewRecepcionFacturaRequest() *RecepcionFacturaBuilder {
	return &RecepcionFacturaBuilder{
		request: &compra_venta.RecepcionFactura{},
	}
}

type RecepcionFacturaBuilder struct {
	request *compra_venta.RecepcionFactura
}

func (b *RecepcionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoEmision(codigoEmision int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoSistema(codigoSistema string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *RecepcionFacturaBuilder) WithCufd(cufd string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *RecepcionFacturaBuilder) WithCuis(cuis string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *RecepcionFacturaBuilder) WithNit(nit int64) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}

func (b *RecepcionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *RecepcionFacturaBuilder) WithArchivo(archivo string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

func (b *RecepcionFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *RecepcionFacturaBuilder) WithHashArchivo(hashArchivo string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hashArchivo
	return b
}

func (b *RecepcionFacturaBuilder) Build() RecepcionFacturaRequest {
	return requestWrapper[compra_venta.RecepcionFactura]{request: b.request}
}

// NewFacturaCompraVenta inicia la construcción de una FacturaCompraVenta.
func (compraVentaNamespace) NewFacturaCompraVenta() *FacturaCompraVentaBuilder {
	return &FacturaCompraVentaBuilder{
		factura: &compra_venta.FacturaCompraVenta{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVenta"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVenta.xsd",
		},
	}
}

// NewCabecera inicia la construcción de la cabecera de una factura.
func (compraVentaNamespace) NewCabecera() *CabeceraBuilder {
	return &CabeceraBuilder{
		cabecera: &compra_venta.Cabecera{},
	}
}

// NewDetalle inicia la construcción de un ítem de detalle de una factura.
func (compraVentaNamespace) NewDetalle() *DetalleBuilder {
	return &DetalleBuilder{
		detalle: &compra_venta.Detalle{},
	}
}

type FacturaCompraVentaBuilder struct {
	factura *compra_venta.FacturaCompraVenta
}

func (b *FacturaCompraVentaBuilder) WithCabecera(req CabeceraRequest) *FacturaCompraVentaBuilder {
	if c := GetInternalRequest[compra_venta.Cabecera](req); c != nil {
		b.factura.Cabecera = *c
	}
	return b
}

func (b *FacturaCompraVentaBuilder) AddDetalle(req DetalleRequest) *FacturaCompraVentaBuilder {
	if d := GetInternalRequest[compra_venta.Detalle](req); d != nil {
		b.factura.Detalle = append(b.factura.Detalle, *d)
	}
	return b
}

func (b *FacturaCompraVentaBuilder) WithModalidad(tipo int) *FacturaCompraVentaBuilder {

	switch tipo {
	case 1:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVenta.xsd"
	case 2:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVenta.xsd"
	}
	return b
}

func (b *FacturaCompraVentaBuilder) Build() FacturaRequest {
	return requestWrapper[compra_venta.FacturaCompraVenta]{request: b.factura}
}

// CabeceraBuilder ayuda a configurar la cabecera de la factura.
type CabeceraBuilder struct {
	cabecera *compra_venta.Cabecera
}

func (b *CabeceraBuilder) WithNitEmisor(nitEmisor int64) *CabeceraBuilder {
	b.cabecera.NitEmisor = nitEmisor
	return b
}

func (b *CabeceraBuilder) WithRazonSocialEmisor(razonSocialEmisor string) *CabeceraBuilder {
	b.cabecera.RazonSocialEmisor = razonSocialEmisor
	return b
}

func (b *CabeceraBuilder) WithMunicipio(municipio string) *CabeceraBuilder {
	b.cabecera.Municipio = municipio
	return b
}

func (b *CabeceraBuilder) WithTelefono(telefono string) *CabeceraBuilder {
	b.cabecera.Telefono = compra_venta.Nilable[string]{Value: &telefono}
	return b
}

func (b *CabeceraBuilder) WithNumeroFactura(numeroFactura int64) *CabeceraBuilder {
	b.cabecera.NumeroFactura = numeroFactura
	return b
}

func (b *CabeceraBuilder) WithCuf(cuf string) *CabeceraBuilder {
	b.cabecera.Cuf = cuf
	return b
}

func (b *CabeceraBuilder) WithCufd(cufd string) *CabeceraBuilder {
	b.cabecera.Cufd = cufd
	return b
}

func (b *CabeceraBuilder) WithCodigoSucursal(codigoSucursal int) *CabeceraBuilder {
	b.cabecera.CodigoSucursal = codigoSucursal
	return b
}

func (b *CabeceraBuilder) WithDireccion(direccion string) *CabeceraBuilder {
	b.cabecera.Direccion = direccion
	return b
}

func (b *CabeceraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *CabeceraBuilder {
	b.cabecera.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *CabeceraBuilder) WithFechaEmision(fechaEmision string) *CabeceraBuilder {
	b.cabecera.FechaEmision = fechaEmision
	return b
}

func (b *CabeceraBuilder) WithNombreRazonSocial(nombreRazonSocial string) *CabeceraBuilder {
	b.cabecera.NombreRazonSocial = compra_venta.Nilable[string]{Value: &nombreRazonSocial}
	return b
}

func (b *CabeceraBuilder) WithCodigoTipoDocumentoIdentidad(codigoTipoDocumentoIdentidad int) *CabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = codigoTipoDocumentoIdentidad
	return b
}

func (b *CabeceraBuilder) WithNumeroDocumento(numeroDocumento string) *CabeceraBuilder {
	b.cabecera.NumeroDocumento = numeroDocumento
	return b
}

func (b *CabeceraBuilder) WithComplemento(complemento string) *CabeceraBuilder {
	b.cabecera.Complemento = compra_venta.Nilable[string]{Value: &complemento}
	return b
}

func (b *CabeceraBuilder) WithCodigoCliente(codigoCliente string) *CabeceraBuilder {
	b.cabecera.CodigoCliente = codigoCliente
	return b
}

func (b *CabeceraBuilder) WithCodigoMetodoPago(codigoMetodoPago int) *CabeceraBuilder {
	b.cabecera.CodigoMetodoPago = codigoMetodoPago
	return b
}

func (b *CabeceraBuilder) WithNumeroTarjeta(numeroTarjeta int64) *CabeceraBuilder {
	b.cabecera.NumeroTarjeta = compra_venta.Nilable[int64]{Value: &numeroTarjeta}
	return b
}

func (b *CabeceraBuilder) WithMontoTotal(montoTotal float64) *CabeceraBuilder {
	b.cabecera.MontoTotal = montoTotal
	return b
}

func (b *CabeceraBuilder) WithMontoTotalSujetoIva(montoTotalSujetoIva float64) *CabeceraBuilder {
	b.cabecera.MontoTotalSujetoIva = montoTotalSujetoIva
	return b
}

func (b *CabeceraBuilder) WithCodigoMoneda(codigoMoneda int) *CabeceraBuilder {
	b.cabecera.CodigoMoneda = codigoMoneda
	return b
}

func (b *CabeceraBuilder) WithTipoCambio(tipoCambio float64) *CabeceraBuilder {
	b.cabecera.TipoCambio = tipoCambio
	return b
}

func (b *CabeceraBuilder) WithMontoTotalMoneda(montoTotalMoneda float64) *CabeceraBuilder {
	b.cabecera.MontoTotalMoneda = montoTotalMoneda
	return b
}

func (b *CabeceraBuilder) WithMontoGiftCard(montoGiftCard float64) *CabeceraBuilder {
	b.cabecera.MontoGiftCard = compra_venta.Nilable[float64]{Value: &montoGiftCard}
	return b
}

func (b *CabeceraBuilder) WithDescuentoAdicional(descuentoAdicional float64) *CabeceraBuilder {
	b.cabecera.DescuentoAdicional = compra_venta.Nilable[float64]{Value: &descuentoAdicional}
	return b
}

func (b *CabeceraBuilder) WithCodigoExcepcion(codigoExcepcion int64) *CabeceraBuilder {
	b.cabecera.CodigoExcepcion = compra_venta.Nilable[int64]{Value: &codigoExcepcion}
	return b
}

func (b *CabeceraBuilder) WithCafc(cafc string) *CabeceraBuilder {
	b.cabecera.Cafc = compra_venta.Nilable[string]{Value: &cafc}
	return b
}

func (b *CabeceraBuilder) WithLeyenda(leyenda string) *CabeceraBuilder {
	b.cabecera.Leyenda = leyenda
	return b
}

func (b *CabeceraBuilder) WithUsuario(usuario string) *CabeceraBuilder {
	b.cabecera.Usuario = usuario
	return b
}

func (b *CabeceraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *CabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *CabeceraBuilder) Build() CabeceraRequest {
	return requestWrapper[compra_venta.Cabecera]{request: b.cabecera}
}

// DetalleBuilder ayuda a configurar un ítem de detalle de la factura.
type DetalleBuilder struct {
	detalle *compra_venta.Detalle
}

func (b *DetalleBuilder) WithActividadEconomica(actividadEconomica string) *DetalleBuilder {
	b.detalle.ActividadEconomica = actividadEconomica
	return b
}

func (b *DetalleBuilder) WithCodigoProductoSin(codigoProductoSin string) *DetalleBuilder {
	b.detalle.CodigoProductoSin = codigoProductoSin
	return b
}

func (b *DetalleBuilder) WithCodigoProducto(codigoProducto string) *DetalleBuilder {
	b.detalle.CodigoProducto = codigoProducto
	return b
}

func (b *DetalleBuilder) WithDescripcion(descripcion string) *DetalleBuilder {
	b.detalle.Descripcion = descripcion
	return b
}

func (b *DetalleBuilder) WithCantidad(cantidad float64) *DetalleBuilder {
	b.detalle.Cantidad = cantidad
	return b
}

func (b *DetalleBuilder) WithUnidadMedida(unidadMedida int) *DetalleBuilder {
	b.detalle.UnidadMedida = unidadMedida
	return b
}

func (b *DetalleBuilder) WithPrecioUnitario(precioUnitario float64) *DetalleBuilder {
	b.detalle.PrecioUnitario = precioUnitario
	return b
}

func (b *DetalleBuilder) WithMontoDescuento(montoDescuento float64) *DetalleBuilder {
	b.detalle.MontoDescuento = compra_venta.Nilable[float64]{Value: &montoDescuento}
	return b
}

func (b *DetalleBuilder) WithSubTotal(subTotal float64) *DetalleBuilder {
	b.detalle.SubTotal = subTotal
	return b
}

func (b *DetalleBuilder) WithNumeroSerie(numeroSerie string) *DetalleBuilder {
	b.detalle.NumeroSerie = compra_venta.Nilable[string]{Value: &numeroSerie}
	return b
}

func (b *DetalleBuilder) WithNumeroImei(numeroImei string) *DetalleBuilder {
	b.detalle.NumeroImei = compra_venta.Nilable[string]{Value: &numeroImei}
	return b
}

func (b *DetalleBuilder) Build() DetalleRequest {
	return requestWrapper[compra_venta.Detalle]{request: b.detalle}
}
