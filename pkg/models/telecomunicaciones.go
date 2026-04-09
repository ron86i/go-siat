package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// AnulacionFacturaTelecomunicaciones representa la solicitud opaca para anular una factura en la modalidad telecomunicaciones.
type AnulacionFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.AnulacionFactura]
}

// RecepcionFacturaTelecomunicaciones representa la solicitud opaca para la recepción de una factura telecomunicaciones.
type RecepcionFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.RecepcionFactura]
}

// VerificarComunicacionTelecomunicaciones representa la solicitud opaca para verificar la conexión con el SIAT.
type VerificarComunicacionTelecomunicaciones struct {
	RequestWrapper[facturacion.VerificarComunicacion]
}

// ReversionAnulacionFacturaTelecomunicaciones representa la solicitud opaca para la reversión de anulación de una factura telecomunicaciones.
type ReversionAnulacionFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.ReversionAnulacionFactura]
}

type RecepcionPaqueteFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.RecepcionPaqueteFactura]
}

type ValidacionRecepcionPaqueteFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

type RecepcionMasivaFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.RecepcionMasivaFactura]
}

type ValidacionRecepcionMasivaFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

type VerificacionEstadoFacturaTelecomunicaciones struct {
	RequestWrapper[facturacion.VerificacionEstadoFactura]
}

// --- Namespace ---

type telecomunicacionesNamespace struct{}

// Telecomunicaciones expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func Telecomunicaciones() telecomunicacionesNamespace {
	return telecomunicacionesNamespace{}
}

// --- Constructores de Builders ---

func (telecomunicacionesNamespace) NewRecepcionFacturaBuilder() *recepcionFacturaTelecomunicacionesBuilder {
	return &recepcionFacturaTelecomunicacionesBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

func (telecomunicacionesNamespace) NewAnulacionFacturaBuilder() *anulacionFacturaTelecomunicacionesBuilder {
	return &anulacionFacturaTelecomunicacionesBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

func (telecomunicacionesNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionTelecomunicacionesBuilder {
	return &verificarComunicacionTelecomunicacionesBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

func (telecomunicacionesNamespace) NewReversionAnulacionFacturaBuilder() *reversionAnulacionFacturaTelecomunicacionesBuilder {
	return &reversionAnulacionFacturaTelecomunicacionesBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

func (telecomunicacionesNamespace) NewRecepcionPaqueteFacturaBuilder() *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	return &recepcionPaqueteFacturaTelecomunicacionesBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

func (telecomunicacionesNamespace) NewValidacionRecepcionPaqueteFacturaBuilder() *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	return &validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

func (telecomunicacionesNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	return &validacionRecepcionMasivaFacturaTelecomunicacionesBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

func (telecomunicacionesNamespace) NewVerificacionEstadoFacturaBuilder() *verificacionEstadoFacturaTelecomunicacionesBuilder {
	return &verificacionEstadoFacturaTelecomunicacionesBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

func (telecomunicacionesNamespace) NewRecepcionMasivaFacturaBuilder() *recepcionMasivaFacturaTelecomunicacionesBuilder {
	return &recepcionMasivaFacturaTelecomunicacionesBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

// --- Implementaciones de Builders ---

type recepcionFacturaTelecomunicacionesBuilder struct {
	request *facturacion.RecepcionFactura
}

// WithCodigoAmbiente establece el código de ambiente (Piloto o Producción).
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad de facturación.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema autorizado.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithNit(nit int64) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo XML comprimido (.gz).
func (b *recepcionFacturaTelecomunicacionesBuilder) WithArchivo(archivo string) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de emisión.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithFechaEnvio(fecha time.Time) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash SHA-256 del archivo XML.
func (b *recepcionFacturaTelecomunicacionesBuilder) WithHashArchivo(hash string) *recepcionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hash
	return b
}

func (b *recepcionFacturaTelecomunicacionesBuilder) Build() RecepcionFacturaTelecomunicaciones {
	return RecepcionFacturaTelecomunicaciones{RequestWrapper[facturacion.RecepcionFactura]{request: b.request}}
}

type anulacionFacturaTelecomunicacionesBuilder struct {
	request *facturacion.AnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithNit(nit int64) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF de la factura a anular.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCuf(cuf string) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

// WithCodigoMotivo establece el motivo de anulación.
func (b *anulacionFacturaTelecomunicacionesBuilder) WithCodigoMotivo(motivo int) *anulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = motivo
	return b
}

func (b *anulacionFacturaTelecomunicacionesBuilder) Build() AnulacionFacturaTelecomunicaciones {
	return AnulacionFacturaTelecomunicaciones{RequestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

type verificarComunicacionTelecomunicacionesBuilder struct {
	request *facturacion.VerificarComunicacion
}

func (b *verificarComunicacionTelecomunicacionesBuilder) Build() VerificarComunicacionTelecomunicaciones {
	return VerificarComunicacionTelecomunicaciones{RequestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

type reversionAnulacionFacturaTelecomunicacionesBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithNit(nit int64) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF cuya anulación se desea revertir.
func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) WithCuf(cuf string) *reversionAnulacionFacturaTelecomunicacionesBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

func (b *reversionAnulacionFacturaTelecomunicacionesBuilder) Build() ReversionAnulacionFacturaTelecomunicaciones {
	return ReversionAnulacionFacturaTelecomunicaciones{RequestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

type recepcionPaqueteFacturaTelecomunicacionesBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithNit(nit int64) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo comprimido (.tar.gz) en Base64.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithArchivo(archivo string) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de envío del paquete.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithFechaEnvio(fecha time.Time) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash del archivo comprimido.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithHashArchivo(hash string) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.HashArchivo = hash
	return b
}

// WithCafc establece el CAFC si aplica.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCafc(cafc string) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = cafc
	return b
}

// WithCantidadFacturas establece la cantidad de facturas en el paquete.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCantidadFacturas(cantidad int) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidad
	return b
}

// WithCodigoEvento establece el código de evento significativo.
func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoEvento(evento int64) *recepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = evento
	return b
}

func (b *recepcionPaqueteFacturaTelecomunicacionesBuilder) Build() RecepcionPaqueteFacturaTelecomunicaciones {
	return RecepcionPaqueteFacturaTelecomunicaciones{RequestWrapper[facturacion.RecepcionPaqueteFactura]{request: b.request}}
}

type validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithNit(nit int64) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.TipoFacturaDocumento = tipo
	return b
}

// WithCodigoRecepcion establece el código de recepción del paquete a validar.
func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigo
	return b
}

func (b *validacionRecepcionPaqueteFacturaTelecomunicacionesBuilder) Build() ValidacionRecepcionPaqueteFacturaTelecomunicaciones {
	return ValidacionRecepcionPaqueteFacturaTelecomunicaciones{RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]{request: b.request}}
}

type validacionRecepcionMasivaFacturaTelecomunicacionesBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithNit(nit int64) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipo
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigo
	return b
}

func (b *validacionRecepcionMasivaFacturaTelecomunicacionesBuilder) Build() ValidacionRecepcionMasivaFacturaTelecomunicaciones {
	return ValidacionRecepcionMasivaFacturaTelecomunicaciones{RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}

type verificacionEstadoFacturaTelecomunicacionesBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithNit(nit int64) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF de la factura a verificar.
func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) WithCuf(cuf string) *verificacionEstadoFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

func (b *verificacionEstadoFacturaTelecomunicacionesBuilder) Build() VerificacionEstadoFacturaTelecomunicaciones {
	return VerificacionEstadoFacturaTelecomunicaciones{RequestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

type recepcionMasivaFacturaTelecomunicacionesBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoEmision(codigoEmision int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema authorized.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoSistema(codigoSistema string) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCufd(cufd string) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCuis(cuis string) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithNit(nit int64) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo XML (tar.gz) en Base64.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithArchivo(archivo string) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de envío masivo.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithFechaEnvio(fecha time.Time) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash del archivo masivo.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithHashArchivo(hash string) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.HashArchivo = hash
	return b
}

// WithCantidadFacturas establece la cantidad de facturas.
func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) WithCantidadFacturas(cantidad int) *recepcionMasivaFacturaTelecomunicacionesBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidad
	return b
}

func (b *recepcionMasivaFacturaTelecomunicacionesBuilder) Build() RecepcionMasivaFacturaTelecomunicaciones {
	return RecepcionMasivaFacturaTelecomunicaciones{RequestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}
