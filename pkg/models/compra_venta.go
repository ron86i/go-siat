package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/compra_venta"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// --- Interfaces opacas para restringir el acceso a los atributos ---

type VentaAnexo struct {
	requestWrapper[compra_venta.VentaAnexo]
}

// RecepcionAnexos representa una solicitud para la recepción de anexos de una factura.
type RecepcionAnexos struct {
	requestWrapper[compra_venta.RecepcionAnexos]
}

// VerificacionEstadoFactura representa una solicitud para la verificación del estado de una factura.
type VerificacionEstadoFactura struct {
	requestWrapper[facturacion.VerificacionEstadoFactura]
}

// ValidacionRecepcionMasivaFactura representa una solicitud para la validación de la recepción masiva de facturas.
type ValidacionRecepcionMasivaFactura struct {
	requestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

// RecepcionMasivaFactura representa una solicitud para la recepción masiva de facturas.
type RecepcionMasivaFactura struct {
	requestWrapper[facturacion.RecepcionMasivaFactura]
}

// VerificarComunicacionCompraVenta representa una solicitud para verificar la comunicación con el SIAT.
type VerificarComunicacionCompraVenta struct {
	requestWrapper[facturacion.VerificarComunicacion]
}

// ValidacionRecepcionPaqueteFactura representa una solicitud para validar la recepción de paquetes de facturas.
type ValidacionRecepcionPaqueteFactura struct {
	requestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

// RecepcionPaqueteFactura representa una solicitud para la recepción de paquetes de facturas.
type RecepcionPaqueteFactura struct {
	requestWrapper[facturacion.RecepcionPaqueteFactura]
}

// ReversionAnulacionFactura representa una solicitud para la reversión de anulación de factura.
type ReversionAnulacionFactura struct {
	requestWrapper[facturacion.ReversionAnulacionFactura]
}

// AnulacionFactura representa una solicitud para anular una factura emitida.
type AnulacionFactura struct {
	requestWrapper[facturacion.AnulacionFactura]
}

// RecepcionFactura representa una solicitud para el envío de una factura al SIAT.
type RecepcionFactura struct {
	requestWrapper[facturacion.RecepcionFactura]
}



type compraVentaNamespace struct{}

// CompraVenta expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func CompraVenta() compraVentaNamespace {
	return compraVentaNamespace{}
}

// --- Builders para la creación de solicitudes ---

func (compraVentaNamespace) NewRecepcionAnexosBuilder() *RecepcionAnexosBuilder {
	return &RecepcionAnexosBuilder{
		request: &compra_venta.RecepcionAnexos{},
	}
}

type RecepcionAnexosBuilder struct {
	request *compra_venta.RecepcionAnexos
}

func (b *RecepcionAnexosBuilder) WithCodigoAmbiente(codigoAmbiente int) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RecepcionAnexosBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RecepcionAnexosBuilder) WithCodigoSistema(codigoSistema string) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoSistema = codigoSistema
	return b
}

func (b *RecepcionAnexosBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoSucursal = codigoSucursal
	return b
}

func (b *RecepcionAnexosBuilder) WithCuis(cuis string) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cuis = cuis
	return b
}

func (b *RecepcionAnexosBuilder) WithNit(nit int64) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Nit = nit
	return b
}

func (b *RecepcionAnexosBuilder) WithCuf(cuf string) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cuf = cuf
	return b
}

func (b *RecepcionAnexosBuilder) AddAnexo(anexo VentaAnexo) *RecepcionAnexosBuilder {
	if anexo.request != nil {
		b.request.SolicitudRecepcionAnexos.AnexosList = append(b.request.SolicitudRecepcionAnexos.AnexosList, *anexo.request)
	}
	return b
}

// NewVentaAnexoBuilder crea el constructor para un ítem de anexo de venta.
func (compraVentaNamespace) NewVentaAnexoBuilder() *VentaAnexoBuilder {
	return &VentaAnexoBuilder{
		anexo: &compra_venta.VentaAnexo{},
	}
}

type VentaAnexoBuilder struct {
	anexo *compra_venta.VentaAnexo
}

// WithCodigo configura el código del anexo.
func (b *VentaAnexoBuilder) WithCodigo(codigo string) *VentaAnexoBuilder {
	b.anexo.Codigo = codigo
	return b
}

// WithCodigoProducto configura el código de producto del anexo.
func (b *VentaAnexoBuilder) WithCodigoProducto(codigoProducto string) *VentaAnexoBuilder {
	b.anexo.CodigoProducto = codigoProducto
	return b
}

// WithCodigoProductoSin configura el código SIN del producto del anexo.
func (b *VentaAnexoBuilder) WithCodigoProductoSin(codigoProductoSin int64) *VentaAnexoBuilder {
	b.anexo.CodigoProductoSin = codigoProductoSin
	return b
}

// WithTipoCodigo configura el tipo de código del anexo.
func (b *VentaAnexoBuilder) WithTipoCodigo(tipoCodigo string) *VentaAnexoBuilder {
	b.anexo.TipoCodigo = tipoCodigo
	return b
}

// Build finaliza la construcción del anexo retornando la estructura opaca.
func (b *VentaAnexoBuilder) Build() VentaAnexo {
	return VentaAnexo{requestWrapper[compra_venta.VentaAnexo]{request: b.anexo}}
}
func (b *RecepcionAnexosBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *RecepcionAnexosBuilder) WithCodigoEmision(codigoEmision int) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoEmision = codigoEmision
	return b
}

func (b *RecepcionAnexosBuilder) WithCodigoModalidad(codigoModalidad int) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoModalidad = codigoModalidad
	return b
}

func (b *RecepcionAnexosBuilder) WithCufd(cufd string) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cufd = cufd
	return b
}

func (b *RecepcionAnexosBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *RecepcionAnexosBuilder) Build() RecepcionAnexos {
	return RecepcionAnexos{requestWrapper[compra_venta.RecepcionAnexos]{request: b.request}}
}

func (compraVentaNamespace) NewRecepcionPaqueteFacturaBuilder() *RecepcionPaqueteFacturaBuilder {
	return &RecepcionPaqueteFacturaBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

type RecepcionPaqueteFacturaBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoEmision(codigoEmision int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoSistema(codigoSistema string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCufd(cufd string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCuis(cuis string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithNit(nit int64) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithArchivo(archivo string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Archivo = archivo
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithHashArchivo(hashArchivo string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.HashArchivo = hashArchivo
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCafc(cafc string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = cafc
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCantidadFacturas(cantidadFacturas int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidadFacturas
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoEvento(codigoEvento int64) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = codigoEvento
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) Build() RecepcionPaqueteFactura {
	return RecepcionPaqueteFactura{requestWrapper[facturacion.RecepcionPaqueteFactura]{request: b.request}}
}

// ReversionAnulacionFacturaRequest inicia la construcción de una solicitud de reversión de anulación de factura.
func (compraVentaNamespace) NewReversionAnulacionFacturaBuilder() *ReversionAnulacionFacturaBuilder {
	return &ReversionAnulacionFacturaBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

type ReversionAnulacionFacturaBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoEmision(codigoEmision int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoSistema(codigoSistema string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCufd(cufd string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCuis(cuis string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithNit(nit int64) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCuf(cuf string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

func (b *ReversionAnulacionFacturaBuilder) Build() ReversionAnulacionFactura {
	return ReversionAnulacionFactura{requestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

// NewAnulacionFacturaRequest inicia la construcción de una solicitud de anulación.
func (compraVentaNamespace) NewAnulacionFacturaBuilder() *AnulacionFacturaBuilder {
	return &AnulacionFacturaBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

type AnulacionFacturaBuilder struct {
	request *facturacion.AnulacionFactura
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

func (b *AnulacionFacturaBuilder) Build() AnulacionFactura {
	return AnulacionFactura{requestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

// NewRecepcionFacturaRequest inicia la construcción de una solicitud de recepción de factura.
func (compraVentaNamespace) NewRecepcionFacturaBuilder() *RecepcionFacturaBuilder {
	return &RecepcionFacturaBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

type RecepcionFacturaBuilder struct {
	request *facturacion.RecepcionFactura
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

func (b *RecepcionFacturaBuilder) Build() RecepcionFactura {
	return RecepcionFactura{requestWrapper[facturacion.RecepcionFactura]{request: b.request}}
}

// NewRecepcionMasivaFacturaBuilder inicia la construcción de una solicitud de recepción masiva de factura.
func (compraVentaNamespace) NewRecepcionMasivaFacturaBuilder() *RecepcionMasivaFacturaBuilder {
	return &RecepcionMasivaFacturaBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

type RecepcionMasivaFacturaBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCodigoEmision(codigoEmision int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCodigoSistema(codigoSistema string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCufd(cufd string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCuis(cuis string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithNit(nit int64) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithArchivo(archivo string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo = archivo
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithHashArchivo(hashArchivo string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.HashArchivo = hashArchivo
	return b
}
func (b *RecepcionMasivaFacturaBuilder) WithCantidadFacturas(cantidadFacturas int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidadFacturas
	return b
}

func (b *RecepcionMasivaFacturaBuilder) Build() RecepcionMasivaFactura {
	return RecepcionMasivaFactura{requestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}

// NewVerificarComunicacionBuilder inicia la construcción de una solicitud de verificación de comunicación.
func (compraVentaNamespace) NewVerificarComunicacionBuilder() *VerificarComunicacionBuilder {
	return &VerificarComunicacionBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

type VerificarComunicacionBuilder struct {
	request *facturacion.VerificarComunicacion
}

func (b *VerificarComunicacionBuilder) Build() VerificarComunicacionCompraVenta {
	return VerificarComunicacionCompraVenta{requestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

// NewValidacionRecepcionPaqueteFacturaBuilder inicia la construcción de una solicitud de validación de paquete de factura.
func (compraVentaNamespace) NewValidacionRecepcionPaqueteFacturaBuilder() *ValidacionRecepcionPaqueteFacturaBuilder {
	return &ValidacionRecepcionPaqueteFacturaBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

type ValidacionRecepcionPaqueteFacturaBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoEmision(codigoEmision int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoSistema(codigoSistema string) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCufd(cufd string) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCuis(cuis string) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithNit(nit int64) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}
func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoRecepcion(codigoRecepcion string) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) Build() ValidacionRecepcionPaqueteFactura {
	return ValidacionRecepcionPaqueteFactura{requestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]{request: b.request}}
}

func (compraVentaNamespace) NewVerificacionEstadoFacturaBuilder() *VerificacionEstadoFacturaBuilder {
	return &VerificacionEstadoFacturaBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

type VerificacionEstadoFacturaBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoEmision(codigoEmision int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoSistema(codigoSistema string) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCufd(cufd string) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCuis(cuis string) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithNit(nit int64) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCuf(cuf string) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

func (b *VerificacionEstadoFacturaBuilder) Build() VerificacionEstadoFactura {
	return VerificacionEstadoFactura{requestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

func (compraVentaNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *ValidacionRecepcionMasivaFacturaBuilder {
	return &ValidacionRecepcionMasivaFacturaBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

type ValidacionRecepcionMasivaFacturaBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoEmision(codigoEmision int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoSistema(codigoSistema string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCufd(cufd string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCuis(cuis string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithNit(nit int64) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoRecepcion(codigoRecepcion string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) Build() ValidacionRecepcionMasivaFactura {
	return ValidacionRecepcionMasivaFactura{requestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}
