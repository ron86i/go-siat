package models

import (
	"time"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
)

// RecepcionPaqueteCompras representa la solicitud opaca para el envío de paquetes de facturas de compras al SIAT.
type RecepcionPaqueteCompras struct {
	RequestWrapper[facturacion.RecepcionPaqueteCompras]
}

// ValidacionRecepcionPaqueteCompras representa la solicitud opaca para validar el estado de un paquete de compras previamente enviado.
type ValidacionRecepcionPaqueteCompras struct {
	RequestWrapper[facturacion.ValidacionRecepcionPaqueteCompras]
}

// AnulacionCompra representa la solicitud opaca para anular un registro de compra en el SIAT.
type AnulacionCompra struct {
	RequestWrapper[facturacion.AnulacionCompra]
}

// ConfirmacionCompras representa la solicitud opaca para confirmar la recepción de facturas de compras.
type ConfirmacionCompras struct {
	RequestWrapper[facturacion.ConfirmacionCompras]
}

// ConsultaCompras representa la solicitud opaca para consultar el historial de facturas de compras registradas.
type ConsultaCompras struct {
	RequestWrapper[facturacion.ConsultaCompras]
}

// VerificarComunicacionRecepcionCompras representa la solicitud opaca para realizar una prueba de conectividad con el servicio de compras.
type VerificarComunicacionRecepcionCompras struct {
	RequestWrapper[facturacion.VerificarComunicacionRecepcionCompras]
}

// --- Builders a nivel de paquete ---

func NewRecepcionPaqueteComprasBuilder() *recepcionPaqueteComprasBuilder {
	return &recepcionPaqueteComprasBuilder{
		request: &facturacion.RecepcionPaqueteCompras{},
	}
}

func NewValidacionRecepcionPaqueteComprasBuilder() *validacionRecepcionPaqueteComprasBuilder {
	return &validacionRecepcionPaqueteComprasBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteCompras{},
	}
}

func NewAnulacionCompraBuilder() *anulacionCompraBuilder {
	return &anulacionCompraBuilder{
		request: &facturacion.AnulacionCompra{},
	}
}

func NewConfirmacionComprasBuilder() *confirmacionComprasBuilder {
	return &confirmacionComprasBuilder{
		request: &facturacion.ConfirmacionCompras{},
	}
}

func NewConsultaComprasBuilder() *consultaComprasBuilder {
	return &consultaComprasBuilder{
		request: &facturacion.ConsultaCompras{},
	}
}

func NewVerificarComunicacionRecepcionComprasBuilder() *verificarComunicacionRecepcionComprasBuilder {
	return &verificarComunicacionRecepcionComprasBuilder{
		request: &facturacion.VerificarComunicacionRecepcionCompras{},
	}
}

// --- Implementaciones de Builders ---

type recepcionPaqueteComprasBuilder struct {
	request *facturacion.RecepcionPaqueteCompras
}

func (b *recepcionPaqueteComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CodigoSucursal = codigoSucursal
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithCufd(cufd string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Cufd = cufd
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithCuis(cuis string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Cuis = cuis
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithArchivo(archivo string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Archivo = archivo
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithCantidadFacturas(cantidad int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CantidadFacturas = cantidad
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithFechaEnvio(fecha time.Time) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithGestion(gestion int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Gestion = gestion
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithHashArchivo(hash string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.HashArchivo = hash
	return b
}

func (b *recepcionPaqueteComprasBuilder) WithPeriodo(periodo int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Periodo = periodo
	return b
}

func (b *recepcionPaqueteComprasBuilder) Build() RecepcionPaqueteCompras {
	return RecepcionPaqueteCompras{RequestWrapper[facturacion.RecepcionPaqueteCompras]{request: b.request}}
}

type validacionRecepcionPaqueteComprasBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteCompras
}

func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoSucursal = codigoSucursal
	return b
}

func (b *validacionRecepcionPaqueteComprasBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.Cufd = cufd
	return b
}

func (b *validacionRecepcionPaqueteComprasBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.Cuis = cuis
	return b
}

func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoRecepcion = codigo
	return b
}

// Build construye la solicitud opaca para la validación de paquetes de compras.
func (b *validacionRecepcionPaqueteComprasBuilder) Build() ValidacionRecepcionPaqueteCompras {
	return ValidacionRecepcionPaqueteCompras{RequestWrapper[facturacion.ValidacionRecepcionPaqueteCompras]{request: b.request}}
}

type anulacionCompraBuilder struct {
	request *facturacion.AnulacionCompra
}

func (b *anulacionCompraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *anulacionCompraBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodigoSucursal = codigoSucursal
	return b
}

func (b *anulacionCompraBuilder) WithCufd(cufd string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.Cufd = cufd
	return b
}

func (b *anulacionCompraBuilder) WithCuis(cuis string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.Cuis = cuis
	return b
}

func (b *anulacionCompraBuilder) WithCodAutorizacion(cod string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodAutorizacion = cod
	return b
}

func (b *anulacionCompraBuilder) WithNitProveedor(nit int64) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.NitProveedor = nit
	return b
}

func (b *anulacionCompraBuilder) WithNroFactura(nro int64) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.NroFactura = nro
	return b
}

func (b *anulacionCompraBuilder) WithNroDuiDim(nro string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.NroDuiDim = nro
	return b
}

func (b *anulacionCompraBuilder) Build() AnulacionCompra {
	return AnulacionCompra{RequestWrapper[facturacion.AnulacionCompra]{request: b.request}}
}

type confirmacionComprasBuilder struct {
	request *facturacion.ConfirmacionCompras
}

func (b *confirmacionComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *confirmacionComprasBuilder) WithCodigoSucursal(codigoSucursal int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CodigoSucursal = codigoSucursal
	return b
}

func (b *confirmacionComprasBuilder) WithCufd(cufd string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Cufd = cufd
	return b
}

func (b *confirmacionComprasBuilder) WithCuis(cuis string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Cuis = cuis
	return b
}

func (b *confirmacionComprasBuilder) WithArchivo(archivo string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Archivo = archivo
	return b
}

func (b *confirmacionComprasBuilder) WithCantidadFacturas(cantidad int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CantidadFacturas = cantidad
	return b
}

func (b *confirmacionComprasBuilder) WithFechaEnvio(fecha time.Time) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

func (b *confirmacionComprasBuilder) WithGestion(gestion int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Gestion = gestion
	return b
}

func (b *confirmacionComprasBuilder) WithHashArchivo(hash string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.HashArchivo = hash
	return b
}

func (b *confirmacionComprasBuilder) WithPeriodo(periodo int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Periodo = periodo
	return b
}

func (b *confirmacionComprasBuilder) Build() ConfirmacionCompras {
	return ConfirmacionCompras{RequestWrapper[facturacion.ConfirmacionCompras]{request: b.request}}
}

type consultaComprasBuilder struct {
	request *facturacion.ConsultaCompras
}

func (b *consultaComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *consultaComprasBuilder) WithCodigoSucursal(codigoSucursal int) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.CodigoSucursal = codigoSucursal
	return b
}

func (b *consultaComprasBuilder) WithCufd(cufd string) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.Cufd = cufd
	return b
}

func (b *consultaComprasBuilder) WithCuis(cuis string) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.Cuis = cuis
	return b
}

func (b *consultaComprasBuilder) WithFecha(fecha time.Time) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.Fecha = datatype.NewTimeSiat(fecha)
	return b
}

func (b *consultaComprasBuilder) Build() ConsultaCompras {
	return ConsultaCompras{RequestWrapper[facturacion.ConsultaCompras]{request: b.request}}
}

type verificarComunicacionRecepcionComprasBuilder struct {
	request *facturacion.VerificarComunicacionRecepcionCompras
}

func (b *verificarComunicacionRecepcionComprasBuilder) Build() VerificarComunicacionRecepcionCompras {
	return VerificarComunicacionRecepcionCompras{RequestWrapper[facturacion.VerificarComunicacionRecepcionCompras]{request: b.request}}
}
