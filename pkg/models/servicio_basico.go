package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// AnulacionFacturaServicioBasico representa la solicitud opaca para anular una factura en la modalidad servicio básico.
type AnulacionFacturaServicioBasico struct {
	RequestWrapper[facturacion.AnulacionFactura]
}

// RecepcionFacturaServicioBasico representa la solicitud opaca para la recepción de una factura servicio básico.
type RecepcionFacturaServicioBasico struct {
	RequestWrapper[facturacion.RecepcionFactura]
}

// VerificarComunicacionServicioBasico representa la solicitud opaca para verificar la conexión con el SIAT.
type VerificarComunicacionServicioBasico struct {
	RequestWrapper[facturacion.VerificarComunicacion]
}

// ReversionAnulacionFacturaServicioBasico representa la solicitud opaca para la reversión de anulación de una factura servicio básico.
type ReversionAnulacionFacturaServicioBasico struct {
	RequestWrapper[facturacion.ReversionAnulacionFactura]
}

// RecepcionPaqueteFacturaServicioBasico representa la solicitud opaca para el envío de paquetes de facturas en la modalidad servicio básico.
type RecepcionPaqueteFacturaServicioBasico struct {
	RequestWrapper[facturacion.RecepcionPaqueteFactura]
}

// ValidacionRecepcionPaqueteFacturaServicioBasico representa la solicitud opaca para validar la recepción de un paquete de facturas servicio básico.
type ValidacionRecepcionPaqueteFacturaServicioBasico struct {
	RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

// RecepcionMasivaFacturaServicioBasico representa la solicitud opaca para el envío masivo de facturas en la modalidad servicio básico.
type RecepcionMasivaFacturaServicioBasico struct {
	RequestWrapper[facturacion.RecepcionMasivaFactura]
}

// ValidacionRecepcionMasivaFacturaServicioBasico representa la solicitud opaca para validar la recepción masiva de facturas servicio básico.
type ValidacionRecepcionMasivaFacturaServicioBasico struct {
	RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

// VerificacionEstadoFacturaServicioBasico representa la solicitud opaca para consultar el estado de una factura servicio básico.
type VerificacionEstadoFacturaServicioBasico struct {
	RequestWrapper[facturacion.VerificacionEstadoFactura]
}

// --- Namespace ---

type servicioBasicoNamespace struct{}

// ServicioBasico expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func ServicioBasico() servicioBasicoNamespace {
	return servicioBasicoNamespace{}
}

// --- Constructores de Builders ---

// NewRecepcionFacturaBuilder crea un nuevo constructor para una solicitud de recepción de factura individual.
func (servicioBasicoNamespace) NewRecepcionFacturaBuilder() *recepcionFacturaServicioBasicoBuilder {
	return &recepcionFacturaServicioBasicoBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

// NewAnulacionFacturaBuilder crea un nuevo constructor para una solicitud de anulación de factura.
func (servicioBasicoNamespace) NewAnulacionFacturaBuilder() *anulacionFacturaServicioBasicoBuilder {
	return &anulacionFacturaServicioBasicoBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

// NewVerificarComunicacionBuilder crea un nuevo constructor para una solicitud de verificación de comunicación.
func (servicioBasicoNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionServicioBasicoBuilder {
	return &verificarComunicacionServicioBasicoBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

// NewReversionAnulacionFacturaBuilder crea un nuevo constructor para una solicitud de reversión de anulación.
func (servicioBasicoNamespace) NewReversionAnulacionFacturaBuilder() *reversionAnulacionFacturaServicioBasicoBuilder {
	return &reversionAnulacionFacturaServicioBasicoBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

// NewRecepcionPaqueteFacturaBuilder crea un nuevo constructor para una solicitud de recepción de paquete de facturas.
func (servicioBasicoNamespace) NewRecepcionPaqueteFacturaBuilder() *recepcionPaqueteFacturaServicioBasicoBuilder {
	return &recepcionPaqueteFacturaServicioBasicoBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionPaqueteFacturaBuilder crea un nuevo constructor para una solicitud de validación de paquete de facturas.
func (servicioBasicoNamespace) NewValidacionRecepcionPaqueteFacturaBuilder() *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	return &validacionRecepcionPaqueteFacturaServicioBasicoBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionMasivaFacturaBuilder crea un nuevo constructor para una solicitud de validación de recepción masiva.
func (servicioBasicoNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	return &validacionRecepcionMasivaFacturaServicioBasicoBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

// NewVerificacionEstadoFacturaBuilder crea un nuevo constructor para una solicitud de verificación de estado de factura.
func (servicioBasicoNamespace) NewVerificacionEstadoFacturaBuilder() *verificacionEstadoFacturaServicioBasicoBuilder {
	return &verificacionEstadoFacturaServicioBasicoBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

// NewRecepcionMasivaFacturaBuilder crea un nuevo constructor para una solicitud de recepción masiva de facturas.
func (servicioBasicoNamespace) NewRecepcionMasivaFacturaBuilder() *recepcionMasivaFacturaServicioBasicoBuilder {
	return &recepcionMasivaFacturaServicioBasicoBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

// --- Implementaciones de Builders ---

// recepcionFacturaServicioBasicoBuilder permite construir una solicitud de recepción de factura.
type recepcionFacturaServicioBasicoBuilder struct {
	request *facturacion.RecepcionFactura
}

// WithCodigoAmbiente establece el código de ambiente (Piloto o Producción).
func (b *recepcionFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad de facturación.
func (b *recepcionFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema autorizado.
func (b *recepcionFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionFacturaServicioBasicoBuilder) WithCufd(cufd string) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionFacturaServicioBasicoBuilder) WithCuis(cuis string) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionFacturaServicioBasicoBuilder) WithNit(nit int64) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo XML comprimido en GZIP y codificado en Base64.
func (b *recepcionFacturaServicioBasicoBuilder) WithArchivo(archivo string) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de emisión del documento.
func (b *recepcionFacturaServicioBasicoBuilder) WithFechaEnvio(fecha time.Time) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash SHA-256 del archivo XML original.
func (b *recepcionFacturaServicioBasicoBuilder) WithHashArchivo(hash string) *recepcionFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hash
	return b
}

// Build construye la solicitud opaca de recepción de factura.
func (b *recepcionFacturaServicioBasicoBuilder) Build() RecepcionFacturaServicioBasico {
	return RecepcionFacturaServicioBasico{RequestWrapper[facturacion.RecepcionFactura]{request: b.request}}
}

// anulacionFacturaServicioBasicoBuilder permite construir una solicitud de anulación de factura.
type anulacionFacturaServicioBasicoBuilder struct {
	request *facturacion.AnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *anulacionFacturaServicioBasicoBuilder) WithCufd(cufd string) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *anulacionFacturaServicioBasicoBuilder) WithCuis(cuis string) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *anulacionFacturaServicioBasicoBuilder) WithNit(nit int64) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *anulacionFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF de la factura a anular.
func (b *anulacionFacturaServicioBasicoBuilder) WithCuf(cuf string) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

// WithCodigoMotivo establece el motivo de anulación.
func (b *anulacionFacturaServicioBasicoBuilder) WithCodigoMotivo(motivo int) *anulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = motivo
	return b
}

// Build construye la solicitud opaca para la anulación de factura.
func (b *anulacionFacturaServicioBasicoBuilder) Build() AnulacionFacturaServicioBasico {
	return AnulacionFacturaServicioBasico{RequestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

// verificarComunicacionServicioBasicoBuilder permite construir una solicitud para verificar comunicación.
type verificarComunicacionServicioBasicoBuilder struct {
	request *facturacion.VerificarComunicacion
}

// Build construye la solicitud opaca para verificar comunicación.
func (b *verificarComunicacionServicioBasicoBuilder) Build() VerificarComunicacionServicioBasico {
	return VerificarComunicacionServicioBasico{RequestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

// reversionAnulacionFacturaServicioBasicoBuilder permite construir una solicitud de reversión de anulación.
type reversionAnulacionFacturaServicioBasicoBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCufd(cufd string) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCuis(cuis string) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithNit(nit int64) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF cuya anulación se desea revertir.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) WithCuf(cuf string) *reversionAnulacionFacturaServicioBasicoBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

// Build construye la solicitud opaca para la reversión de anulación.
func (b *reversionAnulacionFacturaServicioBasicoBuilder) Build() ReversionAnulacionFacturaServicioBasico {
	return ReversionAnulacionFacturaServicioBasico{RequestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

// recepcionPaqueteFacturaServicioBasicoBuilder permite construir una solicitud de recepción de paquete de facturas.
type recepcionPaqueteFacturaServicioBasicoBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCufd(cufd string) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCuis(cuis string) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithNit(nit int64) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo comprimido (.tar.gz) en Base64.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithArchivo(archivo string) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de envío del paquete.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithFechaEnvio(fecha time.Time) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash del archivo comprimido.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithHashArchivo(hash string) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.HashArchivo = hash
	return b
}

// WithCafc establece el CAFC si aplica.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCafc(cafc string) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = cafc
	return b
}

// WithCantidadFacturas establece la cantidad de facturas en el paquete.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCantidadFacturas(cantidad int) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidad
	return b
}

// WithCodigoEvento establece el código de evento significativo.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoEvento(evento int64) *recepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = evento
	return b
}

// Build construye la solicitud opaca para el envío del paquete de facturas.
func (b *recepcionPaqueteFacturaServicioBasicoBuilder) Build() RecepcionPaqueteFacturaServicioBasico {
	return RecepcionPaqueteFacturaServicioBasico{RequestWrapper[facturacion.RecepcionPaqueteFactura]{request: b.request}}
}

// validacionRecepcionPaqueteFacturaServicioBasicoBuilder permite construir una solicitud de validación de paquete.
type validacionRecepcionPaqueteFacturaServicioBasicoBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithNit(nit int64) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.TipoFacturaDocumento = tipo
	return b
}

// WithCodigoRecepcion establece el código de recepción del paquete a validar.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionPaqueteFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigo
	return b
}

// Build construye la solicitud opaca para la validación del paquete.
func (b *validacionRecepcionPaqueteFacturaServicioBasicoBuilder) Build() ValidacionRecepcionPaqueteFacturaServicioBasico {
	return ValidacionRecepcionPaqueteFacturaServicioBasico{RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]{request: b.request}}
}

// validacionRecepcionMasivaFacturaServicioBasicoBuilder permite construir una solicitud de validación masiva.
type validacionRecepcionMasivaFacturaServicioBasicoBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCufd(cufd string) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCuis(cuis string) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithNit(nit int64) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipo
	return b
}

// WithCodigoRecepcion establece el código de recepción masiva a validar.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigo
	return b
}

// Build construye la solicitud opaca para la validación masiva.
func (b *validacionRecepcionMasivaFacturaServicioBasicoBuilder) Build() ValidacionRecepcionMasivaFacturaServicioBasico {
	return ValidacionRecepcionMasivaFacturaServicioBasico{RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}

// verificacionEstadoFacturaServicioBasicoBuilder permite construir una solicitud de verificación de estado.
type verificacionEstadoFacturaServicioBasicoBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCufd(cufd string) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCuis(cuis string) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithNit(nit int64) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF de la factura a verificar.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) WithCuf(cuf string) *verificacionEstadoFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

// Build construye la solicitud opaca para la verificación del estado de factura.
func (b *verificacionEstadoFacturaServicioBasicoBuilder) Build() VerificacionEstadoFacturaServicioBasico {
	return VerificacionEstadoFacturaServicioBasico{RequestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

// recepcionMasivaFacturaServicioBasicoBuilder permite construir una solicitud de recepción masiva.
type recepcionMasivaFacturaServicioBasicoBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCodigoEmision(codigoEmision int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCodigoSistema(codigoSistema string) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCufd(cufd string) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCuis(cuis string) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithNit(nit int64) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo XML (tar.gz) en Base64.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithArchivo(archivo string) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de envío masivo.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithFechaEnvio(fecha time.Time) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash del archivo masivo.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithHashArchivo(hash string) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.HashArchivo = hash
	return b
}

// WithCantidadFacturas establece la cantidad de facturas en el envío masivo.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) WithCantidadFacturas(cantidad int) *recepcionMasivaFacturaServicioBasicoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidad
	return b
}

// Build construye la solicitud opaca para el envío masivo de facturas.
func (b *recepcionMasivaFacturaServicioBasicoBuilder) Build() RecepcionMasivaFacturaServicioBasico {
	return RecepcionMasivaFacturaServicioBasico{RequestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}
