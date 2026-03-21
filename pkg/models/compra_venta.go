package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/compra_venta"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// --- Interfaces opacas para restringir el acceso a los atributos ---

type VentaAnexoCompraVenta struct {
	requestWrapper[compra_venta.VentaAnexo]
}

// RecepcionAnexosCompraVenta representa una solicitud para la recepción de anexos de una factura.
type RecepcionAnexosCompraVenta struct {
	requestWrapper[compra_venta.RecepcionAnexos]
}

// VerificacionEstadoFacturaCompraVenta representa una solicitud para la verificación del estado de una factura.
type VerificacionEstadoFacturaCompraVenta struct {
	requestWrapper[facturacion.VerificacionEstadoFactura]
}

// ValidacionRecepcionMasivaFacturaCompraVenta representa una solicitud para la validación de la recepción masiva de facturas.
type ValidacionRecepcionMasivaFacturaCompraVenta struct {
	requestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

// RecepcionMasivaFacturaCompraVenta representa una solicitud para la recepción masiva de facturas.
type RecepcionMasivaFacturaCompraVenta struct {
	requestWrapper[facturacion.RecepcionMasivaFactura]
}

// VerificarComunicacionCompraVenta representa una solicitud para verificar la comunicación con el SIAT.
type VerificarComunicacionCompraVenta struct {
	requestWrapper[facturacion.VerificarComunicacion]
}

// ValidacionRecepcionPaqueteFacturaCompraVenta representa una solicitud para validar la recepción de paquetes de facturas.
type ValidacionRecepcionPaqueteFacturaCompraVenta struct {
	requestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

// RecepcionPaqueteFacturaCompraVenta representa una solicitud para la recepción de paquetes de facturas.
type RecepcionPaqueteFacturaCompraVenta struct {
	requestWrapper[facturacion.RecepcionPaqueteFactura]
}

// ReversionAnulacionFacturaCompraVenta representa una solicitud para la reversión de anulación de factura.
type ReversionAnulacionFacturaCompraVenta struct {
	requestWrapper[facturacion.ReversionAnulacionFactura]
}

// AnulacionFacturaCompraVenta representa una solicitud para anular una factura emitida.
type AnulacionFacturaCompraVenta struct {
	requestWrapper[facturacion.AnulacionFactura]
}

// RecepcionFactura representa una solicitud para el envío de una factura al SIAT.
type RecepcionFactura struct {
	requestWrapper[facturacion.RecepcionFactura]
}

// --- Namespace ---

type compraVentaNamespace struct{}

// CompraVenta expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func CompraVenta() compraVentaNamespace {
	return compraVentaNamespace{}
}

// --- Constructores de Builders ---

func (compraVentaNamespace) NewRecepcionAnexosBuilder() *recepcionAnexosBuilder {
	return &recepcionAnexosBuilder{
		request: &compra_venta.RecepcionAnexos{},
	}
}

// NewVentaAnexoBuilder crea el constructor para un ítem de anexo de venta.
func (compraVentaNamespace) NewVentaAnexoBuilder() *ventaAnexoCompraVentaBuilder {
	return &ventaAnexoCompraVentaBuilder{
		anexo: &compra_venta.VentaAnexo{},
	}
}

func (compraVentaNamespace) NewRecepcionPaqueteFacturaBuilder() *recepcionPaqueteFacturaBuilder {
	return &recepcionPaqueteFacturaBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

// ReversionAnulacionFacturaRequest inicia la construcción de una solicitud de reversión de anulación de factura.
func (compraVentaNamespace) NewReversionAnulacionFacturaBuilder() *reversionAnulacionFacturaBuilder {
	return &reversionAnulacionFacturaBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

// NewAnulacionFacturaRequest inicia la construcción de una solicitud de anulación.
func (compraVentaNamespace) NewAnulacionFacturaBuilder() *anulacionFacturaBuilder {
	return &anulacionFacturaBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

// NewRecepcionFacturaRequest inicia la construcción de una solicitud de recepción de factura.
func (compraVentaNamespace) NewRecepcionFacturaBuilder() *recepcionFacturaBuilder {
	return &recepcionFacturaBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

// NewRecepcionMasivaFactura inicia la construcción de una solicitud de recepción masiva de factura.
func (compraVentaNamespace) NewRecepcionMasivaFacturaBuilder() *recepcionMasivaFacturaBuilder {
	return &recepcionMasivaFacturaBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

// NewVerificarComunicacion inicia la construcción de una solicitud de verificación de comunicación.
func (compraVentaNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionBuilder {
	return &verificarComunicacionBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

// NewValidacionRecepcionPaqueteFactura inicia la construcción de una solicitud de validación de paquete de factura.
func (compraVentaNamespace) NewValidacionRecepcionPaqueteFacturaBuilder() *validacionRecepcionPaqueteFacturaBuilder {
	return &validacionRecepcionPaqueteFacturaBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

func (compraVentaNamespace) NewVerificacionEstadoFacturaBuilder() *verificacionEstadoFacturaBuilder {
	return &verificacionEstadoFacturaBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

func (compraVentaNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *validacionRecepcionMasivaFacturaBuilder {
	return &validacionRecepcionMasivaFacturaBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

// --- Implementaciones de Builders ---

type recepcionAnexosBuilder struct {
	request *compra_venta.RecepcionAnexos
}

func (b *recepcionAnexosBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *recepcionAnexosBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *recepcionAnexosBuilder) WithCodigoSistema(codigoSistema string) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoSistema = codigoSistema
	return b
}

func (b *recepcionAnexosBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoSucursal = codigoSucursal
	return b
}

func (b *recepcionAnexosBuilder) WithCuis(cuis string) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cuis = cuis
	return b
}

func (b *recepcionAnexosBuilder) WithNit(nit int64) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Nit = nit
	return b
}

func (b *recepcionAnexosBuilder) WithCuf(cuf string) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cuf = cuf
	return b
}

func (b *recepcionAnexosBuilder) AddAnexos(anexo ...VentaAnexoCompraVenta) *recepcionAnexosBuilder {
	for _, a := range anexo {
		if a.request != nil {
			b.request.SolicitudRecepcionAnexos.AnexosList = append(b.request.SolicitudRecepcionAnexos.AnexosList, *a.request)
		}
	}
	return b
}

type ventaAnexoCompraVentaBuilder struct {
	anexo *compra_venta.VentaAnexo
}

// WithCodigo configura el código del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithCodigo(codigo string) *ventaAnexoCompraVentaBuilder {
	b.anexo.Codigo = codigo
	return b
}

// WithCodigoProducto configura el código de producto del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithCodigoProducto(codigoProducto string) *ventaAnexoCompraVentaBuilder {
	b.anexo.CodigoProducto = codigoProducto
	return b
}

// WithCodigoProductoSin configura el código SIN del producto del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithCodigoProductoSin(codigoProductoSin int64) *ventaAnexoCompraVentaBuilder {
	b.anexo.CodigoProductoSin = codigoProductoSin
	return b
}

// WithTipoCodigo configura el tipo de código del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithTipoCodigo(tipoCodigo string) *ventaAnexoCompraVentaBuilder {
	b.anexo.TipoCodigo = tipoCodigo
	return b
}

// Build finaliza la construcción del anexo retornando la estructura opaca.
func (b *ventaAnexoCompraVentaBuilder) Build() VentaAnexoCompraVenta {
	return VentaAnexoCompraVenta{
		requestWrapper[compra_venta.VentaAnexo]{request: b.anexo},
	}
}
func (b *recepcionAnexosBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *recepcionAnexosBuilder) WithCodigoEmision(codigoEmision int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoEmision = codigoEmision
	return b
}

func (b *recepcionAnexosBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoModalidad = codigoModalidad
	return b
}

func (b *recepcionAnexosBuilder) WithCufd(cufd string) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cufd = cufd
	return b
}

func (b *recepcionAnexosBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *recepcionAnexosBuilder) Build() RecepcionAnexosCompraVenta {
	return RecepcionAnexosCompraVenta{requestWrapper[compra_venta.RecepcionAnexos]{request: b.request}}
}

type recepcionPaqueteFacturaBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoEmision(codigoEmision int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoSistema(codigoSistema string) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCufd(cufd string) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCuis(cuis string) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithNit(nit int64) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithArchivo(archivo string) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Archivo = archivo
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithHashArchivo(hashArchivo string) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.HashArchivo = hashArchivo
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCafc(cafc string) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = cafc
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCantidadFacturas(cantidadFacturas int) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidadFacturas
	return b
}

func (b *recepcionPaqueteFacturaBuilder) WithCodigoEvento(codigoEvento int64) *recepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = codigoEvento
	return b
}

func (b *recepcionPaqueteFacturaBuilder) Build() RecepcionPaqueteFacturaCompraVenta {
	return RecepcionPaqueteFacturaCompraVenta{requestWrapper[facturacion.RecepcionPaqueteFactura]{request: b.request}}
}

type reversionAnulacionFacturaBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

func (b *reversionAnulacionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCodigoEmision(codigoEmision int) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCodigoSistema(codigoSistema string) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCufd(cufd string) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCuis(cuis string) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithNit(nit int64) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *reversionAnulacionFacturaBuilder) WithCuf(cuf string) *reversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

func (b *reversionAnulacionFacturaBuilder) Build() ReversionAnulacionFacturaCompraVenta {
	return ReversionAnulacionFacturaCompraVenta{requestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

type anulacionFacturaBuilder struct {
	request *facturacion.AnulacionFactura
}

func (b *anulacionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *anulacionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *anulacionFacturaBuilder) WithCodigoEmision(codigoEmision int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

func (b *anulacionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

func (b *anulacionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *anulacionFacturaBuilder) WithCodigoSistema(codigoSistema string) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

func (b *anulacionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

func (b *anulacionFacturaBuilder) WithCufd(cufd string) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

func (b *anulacionFacturaBuilder) WithCuf(cuf string) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

func (b *anulacionFacturaBuilder) WithCuis(cuis string) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

func (b *anulacionFacturaBuilder) WithNit(nit int64) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

func (b *anulacionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *anulacionFacturaBuilder) WithCodigoMotivo(codigoMotivo int) *anulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = codigoMotivo
	return b
}

func (b *anulacionFacturaBuilder) Build() AnulacionFacturaCompraVenta {
	return AnulacionFacturaCompraVenta{requestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

type recepcionFacturaBuilder struct {
	request *facturacion.RecepcionFactura
}

func (b *recepcionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *recepcionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *recepcionFacturaBuilder) WithCodigoEmision(codigoEmision int) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *recepcionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *recepcionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *recepcionFacturaBuilder) WithCodigoSistema(codigoSistema string) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

func (b *recepcionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *recepcionFacturaBuilder) WithCufd(cufd string) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *recepcionFacturaBuilder) WithCuis(cuis string) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *recepcionFacturaBuilder) WithNit(nit int64) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}

func (b *recepcionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *recepcionFacturaBuilder) WithArchivo(archivo string) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

func (b *recepcionFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *recepcionFacturaBuilder) WithHashArchivo(hashArchivo string) *recepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hashArchivo
	return b
}

func (b *recepcionFacturaBuilder) Build() RecepcionFactura {
	return RecepcionFactura{requestWrapper[facturacion.RecepcionFactura]{request: b.request}}
}

type recepcionMasivaFacturaBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

func (b *recepcionMasivaFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCodigoEmision(codigoEmision int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCodigoSistema(codigoSistema string) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCufd(cufd string) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCuis(cuis string) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithNit(nit int64) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithArchivo(archivo string) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo = archivo
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithHashArchivo(hashArchivo string) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.HashArchivo = hashArchivo
	return b
}
func (b *recepcionMasivaFacturaBuilder) WithCantidadFacturas(cantidadFacturas int) *recepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidadFacturas
	return b
}

func (b *recepcionMasivaFacturaBuilder) Build() RecepcionMasivaFacturaCompraVenta {
	return RecepcionMasivaFacturaCompraVenta{requestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}

type verificarComunicacionBuilder struct {
	request *facturacion.VerificarComunicacion
}

func (b *verificarComunicacionBuilder) Build() VerificarComunicacionCompraVenta {
	return VerificarComunicacionCompraVenta{requestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

type validacionRecepcionPaqueteFacturaBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithNit(nit int64) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}
func (b *validacionRecepcionPaqueteFacturaBuilder) WithCodigoRecepcion(codigoRecepcion string) *validacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *validacionRecepcionPaqueteFacturaBuilder) Build() ValidacionRecepcionPaqueteFacturaCompraVenta {
	return ValidacionRecepcionPaqueteFacturaCompraVenta{requestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]{request: b.request}}
}

type verificacionEstadoFacturaBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

func (b *verificacionEstadoFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCodigoEmision(codigoEmision int) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCodigoSistema(codigoSistema string) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCufd(cufd string) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCuis(cuis string) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithNit(nit int64) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *verificacionEstadoFacturaBuilder) WithCuf(cuf string) *verificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

func (b *verificacionEstadoFacturaBuilder) Build() VerificacionEstadoFacturaCompraVenta {
	return VerificacionEstadoFacturaCompraVenta{requestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

type validacionRecepcionMasivaFacturaBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCufd(cufd string) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCuis(cuis string) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithNit(nit int64) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) WithCodigoRecepcion(codigoRecepcion string) *validacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *validacionRecepcionMasivaFacturaBuilder) Build() ValidacionRecepcionMasivaFacturaCompraVenta {
	return ValidacionRecepcionMasivaFacturaCompraVenta{requestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}
