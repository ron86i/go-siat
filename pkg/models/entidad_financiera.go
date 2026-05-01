package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// AnulacionFacturaEntidadFinanciera representa la solicitud opaca para anular una factura en la modalidad entidad financiera.
type AnulacionFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.AnulacionFactura]
}

// RecepcionFacturaEntidadFinanciera representa la solicitud opaca para la recepción de una factura entidad financiera.
type RecepcionFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.RecepcionFactura]
}

// VerificarComunicacionEntidadFinanciera representa la solicitud opaca para verificar la conexión con el SIAT.
type VerificarComunicacionEntidadFinanciera struct {
	RequestWrapper[facturacion.VerificarComunicacion]
}

// ReversionAnulacionFacturaEntidadFinanciera representa la solicitud opaca para la reversión de anulación de una factura entidad financiera.
type ReversionAnulacionFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.ReversionAnulacionFactura]
}

// RecepcionPaqueteFacturaEntidadFinanciera representa la solicitud opaca para el envío de paquetes de facturas en la modalidad entidad financiera.
type RecepcionPaqueteFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.RecepcionPaqueteFactura]
}

// ValidacionRecepcionPaqueteFacturaEntidadFinanciera representa la solicitud opaca para validar la recepción de un paquete de facturas entidad financiera.
type ValidacionRecepcionPaqueteFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

// RecepcionMasivaFacturaEntidadFinanciera representa la solicitud opaca para el envío masivo de facturas en la modalidad entidad financiera.
type RecepcionMasivaFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.RecepcionMasivaFactura]
}

// ValidacionRecepcionMasivaFacturaEntidadFinanciera representa la solicitud opaca para validar la recepción masiva de facturas entidad financiera.
type ValidacionRecepcionMasivaFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

// VerificacionEstadoFacturaEntidadFinanciera representa la solicitud opaca para consultar el estado de una factura entidad financiera.
type VerificacionEstadoFacturaEntidadFinanciera struct {
	RequestWrapper[facturacion.VerificacionEstadoFactura]
}

// --- Namespace ---

type entidadFinancieraNamespace struct{}

// EntidadFinanciera expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func EntidadFinanciera() entidadFinancieraNamespace {
	return entidadFinancieraNamespace{}
}

// --- Constructores de Builders ---

// NewRecepcionFacturaBuilder crea un nuevo constructor para una solicitud de recepción de factura individual.
func (entidadFinancieraNamespace) NewRecepcionFacturaBuilder() *recepcionFacturaEntidadFinancieraBuilder {
	return &recepcionFacturaEntidadFinancieraBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

// NewAnulacionFacturaBuilder crea un nuevo constructor para una solicitud de anulación de factura.
func (entidadFinancieraNamespace) NewAnulacionFacturaBuilder() *anulacionFacturaEntidadFinancieraBuilder {
	return &anulacionFacturaEntidadFinancieraBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

// NewVerificarComunicacionBuilder crea un nuevo constructor para una solicitud de verificación de comunicación.
func (entidadFinancieraNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionEntidadFinancieraBuilder {
	return &verificarComunicacionEntidadFinancieraBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

// NewReversionAnulacionFacturaBuilder crea un nuevo constructor para una solicitud de reversión de anulación.
func (entidadFinancieraNamespace) NewReversionAnulacionFacturaBuilder() *reversionAnulacionFacturaEntidadFinancieraBuilder {
	return &reversionAnulacionFacturaEntidadFinancieraBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

// NewRecepcionPaqueteFacturaBuilder crea un nuevo constructor para una solicitud de recepción de paquete de facturas.
func (entidadFinancieraNamespace) NewRecepcionPaqueteFacturaBuilder() *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	return &recepcionPaqueteFacturaEntidadFinancieraBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionPaqueteFacturaBuilder crea un nuevo constructor para una solicitud de validación de paquete de facturas.
func (entidadFinancieraNamespace) NewValidacionRecepcionPaqueteFacturaBuilder() *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	return &validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionMasivaFacturaBuilder crea un nuevo constructor para una solicitud de validación de recepción masiva.
func (entidadFinancieraNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	return &validacionRecepcionMasivaFacturaEntidadFinancieraBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

// NewVerificacionEstadoFacturaBuilder crea un nuevo constructor para una solicitud de verificación de estado de factura.
func (entidadFinancieraNamespace) NewVerificacionEstadoFacturaBuilder() *verificacionEstadoFacturaEntidadFinancieraBuilder {
	return &verificacionEstadoFacturaEntidadFinancieraBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

// NewRecepcionMasivaFacturaBuilder crea un nuevo constructor para una solicitud de recepción masiva de facturas.
func (entidadFinancieraNamespace) NewRecepcionMasivaFacturaBuilder() *recepcionMasivaFacturaEntidadFinancieraBuilder {
	return &recepcionMasivaFacturaEntidadFinancieraBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

// --- Implementaciones de Builders ---

// recepcionFacturaEntidadFinancieraBuilder permite construir una solicitud de recepción de factura.
type recepcionFacturaEntidadFinancieraBuilder struct {
	request *facturacion.RecepcionFactura
}

// WithCodigoAmbiente establece el código de ambiente (Piloto o Producción).
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector (Entidad Financiera).
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión (Online/Offline).
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad de facturación.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema autorizado.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el Código Único de Facturación Diaria.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.Cufd = cufd
	return b
}

// WithCuis establece el Código Único de Identificación del Sistema.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithNit(nit int64) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de factura (Crédito-Débito, etc).
func (b *recepcionFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el contenido XML de la factura comprimido en GZIP y codificado en Base64.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithArchivo(archivo string) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de emisión del documento.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithFechaEnvio(fecha time.Time) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el SHA-256 del archivo XML original.
func (b *recepcionFacturaEntidadFinancieraBuilder) WithHashArchivo(hash string) *recepcionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hash
	return b
}

// Build construye la solicitud opaca de recepción de factura.
func (b *recepcionFacturaEntidadFinancieraBuilder) Build() RecepcionFacturaEntidadFinanciera {
	return RecepcionFacturaEntidadFinanciera{RequestWrapper[facturacion.RecepcionFactura]{request: b.request}}
}

// anulacionFacturaEntidadFinancieraBuilder permite construir una solicitud de anulación de factura.
type anulacionFacturaEntidadFinancieraBuilder struct {
	request *facturacion.AnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithNit(nit int64) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el Código Único de Factura a anular.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCuf(cuf string) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

// WithCodigoMotivo establece el código del motivo de anulación.
func (b *anulacionFacturaEntidadFinancieraBuilder) WithCodigoMotivo(motivo int) *anulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = motivo
	return b
}

// Build construye la solicitud opaca para la anulación de factura.
func (b *anulacionFacturaEntidadFinancieraBuilder) Build() AnulacionFacturaEntidadFinanciera {
	return AnulacionFacturaEntidadFinanciera{RequestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

// verificarComunicacionEntidadFinancieraBuilder permite construir una solicitud para verificar comunicación.
type verificarComunicacionEntidadFinancieraBuilder struct {
	request *facturacion.VerificarComunicacion
}

// Build construye la solicitud opaca para verificar comunicación.
func (b *verificarComunicacionEntidadFinancieraBuilder) Build() VerificarComunicacionEntidadFinanciera {
	return VerificarComunicacionEntidadFinanciera{RequestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

// reversionAnulacionFacturaEntidadFinancieraBuilder permite construir una solicitud de reversión de anulación.
type reversionAnulacionFacturaEntidadFinancieraBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithNit(nit int64) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF cuya anulación se desea revertir.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) WithCuf(cuf string) *reversionAnulacionFacturaEntidadFinancieraBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

// Build construye la solicitud opaca para la reversión de anulación.
func (b *reversionAnulacionFacturaEntidadFinancieraBuilder) Build() ReversionAnulacionFacturaEntidadFinanciera {
	return ReversionAnulacionFacturaEntidadFinanciera{RequestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

// recepcionPaqueteFacturaEntidadFinancieraBuilder permite construir una solicitud de recepción de paquete de facturas.
type recepcionPaqueteFacturaEntidadFinancieraBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithNit(nit int64) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo comprimido (.tar.gz) en formato Base64.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithArchivo(archivo string) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de envío del paquete.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithFechaEnvio(fecha time.Time) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash SHA-256 del archivo comprimido.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithHashArchivo(hash string) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.HashArchivo = hash
	return b
}

// WithCafc establece el CAFC si aplica.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCafc(cafc *string) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = datatype.Nilable[string]{Value: cafc}
	return b
}

// WithCantidadFacturas establece la cantidad de facturas en el paquete.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCantidadFacturas(cantidad int) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidad
	return b
}

// WithCodigoEvento establece el código de evento significativo.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoEvento(evento int64) *recepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = evento
	return b
}

// Build construye la solicitud opaca para el envío del paquete de facturas.
func (b *recepcionPaqueteFacturaEntidadFinancieraBuilder) Build() RecepcionPaqueteFacturaEntidadFinanciera {
	return RecepcionPaqueteFacturaEntidadFinanciera{RequestWrapper[facturacion.RecepcionPaqueteFactura]{request: b.request}}
}

// validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder permite construir una solicitud de validación de paquete.
type validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithNit(nit int64) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.TipoFacturaDocumento = tipo
	return b
}

// WithCodigoRecepcion establece el código de recepción a validar.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigo
	return b
}

// Build construye la solicitud opaca para la validación del paquete.
func (b *validacionRecepcionPaqueteFacturaEntidadFinancieraBuilder) Build() ValidacionRecepcionPaqueteFacturaEntidadFinanciera {
	return ValidacionRecepcionPaqueteFacturaEntidadFinanciera{RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]{request: b.request}}
}

// validacionRecepcionMasivaFacturaEntidadFinancieraBuilder permite construir una solicitud de validación masiva.
type validacionRecepcionMasivaFacturaEntidadFinancieraBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithNit(nit int64) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipo
	return b
}

// WithCodigoRecepcion establece el código de recepción masiva a validar.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigo
	return b
}

// Build construye la solicitud opaca para la validación masiva.
func (b *validacionRecepcionMasivaFacturaEntidadFinancieraBuilder) Build() ValidacionRecepcionMasivaFacturaEntidadFinanciera {
	return ValidacionRecepcionMasivaFacturaEntidadFinanciera{RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}

// verificacionEstadoFacturaEntidadFinancieraBuilder permite construir una solicitud de verificación de estado.
type verificacionEstadoFacturaEntidadFinancieraBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithNit(nit int64) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF de la factura a verificar.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) WithCuf(cuf string) *verificacionEstadoFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

// Build construye la solicitud opaca para la verificación del estado de factura.
func (b *verificacionEstadoFacturaEntidadFinancieraBuilder) Build() VerificacionEstadoFacturaEntidadFinanciera {
	return VerificacionEstadoFacturaEntidadFinanciera{RequestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

// recepcionMasivaFacturaEntidadFinancieraBuilder permite construir una solicitud de recepción masiva.
type recepcionMasivaFacturaEntidadFinancieraBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoEmision(codigoEmision int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoSistema(codigoSistema string) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCufd(cufd string) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCuis(cuis string) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithNit(nit int64) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo XML (tar.gz) en Base64.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithArchivo(archivo string) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de envío masivo.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithFechaEnvio(fecha time.Time) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash del archivo masivo.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithHashArchivo(hash string) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.HashArchivo = hash
	return b
}

// WithCantidadFacturas establece la cantidad de facturas en el envío masivo.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) WithCantidadFacturas(cantidad int) *recepcionMasivaFacturaEntidadFinancieraBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidad
	return b
}

// Build construye la solicitud opaca para el envío masivo de facturas.
func (b *recepcionMasivaFacturaEntidadFinancieraBuilder) Build() RecepcionMasivaFacturaEntidadFinanciera {
	return RecepcionMasivaFacturaEntidadFinanciera{RequestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}
