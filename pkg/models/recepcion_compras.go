package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
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

// --- Namespace ---

type recepcionComprasNamespace struct{}

// RecepcionCompras expone constructores de solicitudes para el módulo de Recepción de Compras del SIAT.
func RecepcionCompras() recepcionComprasNamespace {
	return recepcionComprasNamespace{}
}

// --- Builders ---

// NewRecepcionPaqueteComprasBuilder crea un nuevo constructor para recepción de paquetes de compras.
func (recepcionComprasNamespace) NewRecepcionPaqueteComprasBuilder() *recepcionPaqueteComprasBuilder {
	return &recepcionPaqueteComprasBuilder{
		request: &facturacion.RecepcionPaqueteCompras{},
	}
}

// NewValidacionRecepcionPaqueteComprasBuilder crea un nuevo constructor para validación de paquetes de compras.
func (recepcionComprasNamespace) NewValidacionRecepcionPaqueteComprasBuilder() *validacionRecepcionPaqueteComprasBuilder {
	return &validacionRecepcionPaqueteComprasBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteCompras{},
	}
}

// NewAnulacionCompraBuilder crea un nuevo constructor para anulación de compra.
func (recepcionComprasNamespace) NewAnulacionCompraBuilder() *anulacionCompraBuilder {
	return &anulacionCompraBuilder{
		request: &facturacion.AnulacionCompra{},
	}
}

// NewConfirmacionComprasBuilder crea un nuevo constructor para confirmación de compras.
func (recepcionComprasNamespace) NewConfirmacionComprasBuilder() *confirmacionComprasBuilder {
	return &confirmacionComprasBuilder{
		request: &facturacion.ConfirmacionCompras{},
	}
}

// NewConsultaComprasBuilder crea un nuevo constructor para consulta de compras.
func (recepcionComprasNamespace) NewConsultaComprasBuilder() *consultaComprasBuilder {
	return &consultaComprasBuilder{
		request: &facturacion.ConsultaCompras{},
	}
}

// NewVerificarComunicacionBuilder crea un nuevo constructor para verificar comunicación.
func (recepcionComprasNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionRecepcionComprasBuilder {
	return &verificarComunicacionRecepcionComprasBuilder{
		request: &facturacion.VerificarComunicacionRecepcionCompras{},
	}
}

// --- Implementaciones de Builders ---

// recepcionPaqueteComprasBuilder
type recepcionPaqueteComprasBuilder struct {
	request *facturacion.RecepcionPaqueteCompras
}

// WithCodigoAmbiente establece el código de ambiente (Piloto o Producción).
func (b *recepcionPaqueteComprasBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionPaqueteComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema informático autorizado.
func (b *recepcionPaqueteComprasBuilder) WithCodigoSistema(codigoSistema string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionPaqueteComprasBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el Código Único de Facturación Diaria.
func (b *recepcionPaqueteComprasBuilder) WithCufd(cufd string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Cufd = cufd
	return b
}

// WithCuis establece el Código Único de Identificación del Sistema.
func (b *recepcionPaqueteComprasBuilder) WithCuis(cuis string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Cuis = cuis
	return b
}

// WithNit establece el Número de Identificación Tributaria del emisor.
func (b *recepcionPaqueteComprasBuilder) WithNit(nit int64) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Nit = nit
	return b
}

// WithArchivo establece el contenido del paquete de facturas (comprimido en GZIP).
func (b *recepcionPaqueteComprasBuilder) WithArchivo(archivo string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Archivo = archivo
	return b
}

// WithCantidadFacturas establece la cantidad de facturas contenidas en el paquete.
func (b *recepcionPaqueteComprasBuilder) WithCantidadFacturas(cantidad int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.CantidadFacturas = cantidad
	return b
}

// WithFechaEnvio establece la fecha y hora de envío del paquete.
func (b *recepcionPaqueteComprasBuilder) WithFechaEnvio(fecha time.Time) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithGestion establece la gestión (año) de las compras enviadas.
func (b *recepcionPaqueteComprasBuilder) WithGestion(gestion int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Gestion = gestion
	return b
}

// WithHashArchivo establece el hash SHA-256 del archivo enviado.
func (b *recepcionPaqueteComprasBuilder) WithHashArchivo(hash string) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.HashArchivo = hash
	return b
}

// WithPeriodo establece el periodo (mes) de las compras enviadas.
func (b *recepcionPaqueteComprasBuilder) WithPeriodo(periodo int) *recepcionPaqueteComprasBuilder {
	b.request.SolicitudRecepcionCompras.Periodo = periodo
	return b
}

// Build construye la solicitud opaca para la recepción de paquetes de compras.
func (b *recepcionPaqueteComprasBuilder) Build() RecepcionPaqueteCompras {
	return RecepcionPaqueteCompras{RequestWrapper[facturacion.RecepcionPaqueteCompras]{request: b.request}}
}

// validacionRecepcionPaqueteComprasBuilder permite construir una solicitud de validación de paquetes de compras.
type validacionRecepcionPaqueteComprasBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteCompras
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionPaqueteComprasBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionPaqueteComprasBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionPaqueteComprasBuilder) WithNit(nit int64) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.Nit = nit
	return b
}

// WithCodigoRecepcion establece el código de recepción obtenido al enviar el paquete.
func (b *validacionRecepcionPaqueteComprasBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionPaqueteComprasBuilder {
	b.request.SolicitudValidacionRecepcionCompras.CodigoRecepcion = codigo
	return b
}

// Build construye la solicitud opaca para la validación de paquetes de compras.
func (b *validacionRecepcionPaqueteComprasBuilder) Build() ValidacionRecepcionPaqueteCompras {
	return ValidacionRecepcionPaqueteCompras{RequestWrapper[facturacion.ValidacionRecepcionPaqueteCompras]{request: b.request}}
}

// anulacionCompraBuilder permite construir una solicitud de anulación de compra.
type anulacionCompraBuilder struct {
	request *facturacion.AnulacionCompra
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *anulacionCompraBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *anulacionCompraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *anulacionCompraBuilder) WithCodigoSistema(codigoSistema string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *anulacionCompraBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *anulacionCompraBuilder) WithCufd(cufd string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *anulacionCompraBuilder) WithCuis(cuis string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *anulacionCompraBuilder) WithNit(nit int64) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.Nit = nit
	return b
}

// WithCodAutorizacion establece el código de autorización de la factura de compra.
func (b *anulacionCompraBuilder) WithCodAutorizacion(cod string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.CodAutorizacion = cod
	return b
}

// WithNitProveedor establece el NIT del proveedor de la compra.
func (b *anulacionCompraBuilder) WithNitProveedor(nit int64) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.NitProveedor = nit
	return b
}

// WithNroFactura establece el número de la factura de compra.
func (b *anulacionCompraBuilder) WithNroFactura(nro int64) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.NroFactura = nro
	return b
}

// WithNroDuiDim establece el número de DUI/DIM si aplica.
func (b *anulacionCompraBuilder) WithNroDuiDim(nro string) *anulacionCompraBuilder {
	b.request.SolicitudAnulacionCompra.NroDuiDim = nro
	return b
}

// Build construye la solicitud opaca para la anulación de compra.
func (b *anulacionCompraBuilder) Build() AnulacionCompra {
	return AnulacionCompra{RequestWrapper[facturacion.AnulacionCompra]{request: b.request}}
}

// confirmacionComprasBuilder permite construir una solicitud de confirmación de compras.
type confirmacionComprasBuilder struct {
	request *facturacion.ConfirmacionCompras
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *confirmacionComprasBuilder) WithCodigoAmbiente(codigoAmbiente int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *confirmacionComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *confirmacionComprasBuilder) WithCodigoSistema(codigoSistema string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *confirmacionComprasBuilder) WithCodigoSucursal(codigoSucursal int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *confirmacionComprasBuilder) WithCufd(cufd string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *confirmacionComprasBuilder) WithCuis(cuis string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *confirmacionComprasBuilder) WithNit(nit int64) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Nit = nit
	return b
}

// WithArchivo establece el contenido del paquete confirmado.
func (b *confirmacionComprasBuilder) WithArchivo(archivo string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Archivo = archivo
	return b
}

// WithCantidadFacturas establece la cantidad de facturas confirmadas.
func (b *confirmacionComprasBuilder) WithCantidadFacturas(cantidad int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.CantidadFacturas = cantidad
	return b
}

// WithFechaEnvio establece la fecha y hora de envío de la confirmación.
func (b *confirmacionComprasBuilder) WithFechaEnvio(fecha time.Time) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithGestion establece la gestión (año) de las compras confirmadas.
func (b *confirmacionComprasBuilder) WithGestion(gestion int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Gestion = gestion
	return b
}

// WithHashArchivo establece el hash SHA-256 del archivo confirmado.
func (b *confirmacionComprasBuilder) WithHashArchivo(hash string) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.HashArchivo = hash
	return b
}

// WithPeriodo establece el periodo (mes) de las compras confirmadas.
func (b *confirmacionComprasBuilder) WithPeriodo(periodo int) *confirmacionComprasBuilder {
	b.request.SolicitudConfirmacionCompras.Periodo = periodo
	return b
}

// Build construye la solicitud opaca para la confirmación de compras.
func (b *confirmacionComprasBuilder) Build() ConfirmacionCompras {
	return ConfirmacionCompras{RequestWrapper[facturacion.ConfirmacionCompras]{request: b.request}}
}

// consultaComprasBuilder permite construir una solicitud de consulta de compras.
type consultaComprasBuilder struct {
	request *facturacion.ConsultaCompras
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *consultaComprasBuilder) WithCodigoAmbiente(codigoAmbiente int) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *consultaComprasBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *consultaComprasBuilder) WithCodigoSistema(codigoSistema string) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *consultaComprasBuilder) WithCodigoSucursal(codigoSucursal int) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *consultaComprasBuilder) WithCufd(cufd string) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *consultaComprasBuilder) WithCuis(cuis string) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *consultaComprasBuilder) WithNit(nit int64) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.Nit = nit
	return b
}

// WithFecha establece la fecha de consulta de las compras.
func (b *consultaComprasBuilder) WithFecha(fecha time.Time) *consultaComprasBuilder {
	b.request.SolicitudConsultaCompras.Fecha = datatype.NewTimeSiat(fecha)
	return b
}

// Build construye la solicitud opaca para la consulta de compras.
func (b *consultaComprasBuilder) Build() ConsultaCompras {
	return ConsultaCompras{RequestWrapper[facturacion.ConsultaCompras]{request: b.request}}
}

// verificarComunicacionRecepcionComprasBuilder permite construir una solicitud de verificación de comunicación.
type verificarComunicacionRecepcionComprasBuilder struct {
	request *facturacion.VerificarComunicacionRecepcionCompras
}

// Build construye la solicitud opaca para la verificación de comunicación con el servicio de compras.
func (b *verificarComunicacionRecepcionComprasBuilder) Build() VerificarComunicacionRecepcionCompras {
	return VerificarComunicacionRecepcionCompras{RequestWrapper[facturacion.VerificarComunicacionRecepcionCompras]{request: b.request}}
}
