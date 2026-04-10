package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// AnulacionFacturaBoletoAereo representa la solicitud opaca para anular una factura en la modalidad boleto aéreo.
type AnulacionFacturaBoletoAereo struct {
	RequestWrapper[facturacion.AnulacionFactura]
}

// VerificarComunicacionBoletoAereo representa la solicitud opaca para verificar la conexión con el SIAT.
type VerificarComunicacionBoletoAereo struct {
	RequestWrapper[facturacion.VerificarComunicacion]
}

// ReversionAnulacionFacturaBoletoAereo representa la solicitud opaca para la reversión de anulación de una factura boleto aéreo.
type ReversionAnulacionFacturaBoletoAereo struct {
	RequestWrapper[facturacion.ReversionAnulacionFactura]
}

// RecepcionMasivaFacturaBoletoAereo representa la solicitud opaca para el envío masivo de facturas en la modalidad boleto aéreo.
type RecepcionMasivaFacturaBoletoAereo struct {
	RequestWrapper[facturacion.RecepcionMasivaFactura]
}

// ValidacionRecepcionMasivaFacturaBoletoAereo representa la solicitud opaca para validar la recepción masiva de facturas boleto aéreo.
type ValidacionRecepcionMasivaFacturaBoletoAereo struct {
	RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

// VerificacionEstadoFacturaBoletoAereo representa la solicitud opaca para consultar el estado de una factura boleto aéreo.
type VerificacionEstadoFacturaBoletoAereo struct {
	RequestWrapper[facturacion.VerificacionEstadoFactura]
}

// --- Namespace ---

type boletoAereoNamespace struct{}

// BoletoAereo expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func BoletoAereo() boletoAereoNamespace {
	return boletoAereoNamespace{}
}

// --- Constructores de Builders ---

// NewAnulacionFacturaBuilder crea un nuevo constructor para una solicitud de anulación de factura.
func (boletoAereoNamespace) NewAnulacionFacturaBuilder() *anulacionFacturaBoletoAereoBuilder {
	return &anulacionFacturaBoletoAereoBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

// NewVerificarComunicacionBuilder crea un nuevo constructor para una solicitud de verificación de comunicación.
func (boletoAereoNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionBoletoAereoBuilder {
	return &verificarComunicacionBoletoAereoBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

// NewReversionAnulacionFacturaBuilder crea un nuevo constructor para una solicitud de reversión de anulación.
func (boletoAereoNamespace) NewReversionAnulacionFacturaBuilder() *reversionAnulacionFacturaBoletoAereoBuilder {
	return &reversionAnulacionFacturaBoletoAereoBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

// NewValidacionRecepcionMasivaFacturaBuilder crea un nuevo constructor para una solicitud de validación de recepción masiva.
func (boletoAereoNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	return &validacionRecepcionMasivaFacturaBoletoAereoBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

// NewVerificacionEstadoFacturaBuilder crea un nuevo constructor para una solicitud de verificación de estado de factura.
func (boletoAereoNamespace) NewVerificacionEstadoFacturaBuilder() *verificacionEstadoFacturaBoletoAereoBuilder {
	return &verificacionEstadoFacturaBoletoAereoBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

// NewRecepcionMasivaFacturaBuilder crea un nuevo constructor para una solicitud de recepción masiva de facturas.
func (boletoAereoNamespace) NewRecepcionMasivaFacturaBuilder() *recepcionMasivaFacturaBoletoAereoBuilder {
	return &recepcionMasivaFacturaBoletoAereoBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

// --- Implementaciones de Builders ---

// anulacionFacturaBoletoAereoBuilder permite construir una solicitud de anulación de factura.
type anulacionFacturaBoletoAereoBuilder struct {
	request *facturacion.AnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente (Piloto o Producción).
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión (Online/Offline).
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoEmision(codigoEmision int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad de facturación.
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoModalidad(codigoModalidad int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema autorizado por el SIN.
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoSistema(codigoSistema string) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el Código Único de Facturación Diaria.
func (b *anulacionFacturaBoletoAereoBuilder) WithCufd(cufd string) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el Código Único de Inicio de Sistemas.
func (b *anulacionFacturaBoletoAereoBuilder) WithCuis(cuis string) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *anulacionFacturaBoletoAereoBuilder) WithNit(nit int64) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de factura (ej. con derecho a crédito fiscal).
func (b *anulacionFacturaBoletoAereoBuilder) WithTipoFacturaDocumento(tipo int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF de la factura que se desea anular.
func (b *anulacionFacturaBoletoAereoBuilder) WithCuf(cuf string) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

// WithCodigoMotivo establece el motivo de anulación (según catálogo del SIN).
func (b *anulacionFacturaBoletoAereoBuilder) WithCodigoMotivo(motivo int) *anulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = motivo
	return b
}

// Build construye la solicitud opaca para la anulación de factura de boleto aéreo.
func (b *anulacionFacturaBoletoAereoBuilder) Build() AnulacionFacturaBoletoAereo {
	return AnulacionFacturaBoletoAereo{RequestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

// verificarComunicacionBoletoAereoBuilder permite construir una solicitud para verificar comunicación del sector aeronáutico.
type verificarComunicacionBoletoAereoBuilder struct {
	request *facturacion.VerificarComunicacion
}

// Build construye la solicitud opaca para verificar comunicación con el SIAT.
func (b *verificarComunicacionBoletoAereoBuilder) Build() VerificarComunicacionBoletoAereo {
	return VerificarComunicacionBoletoAereo{RequestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

// reversionAnulacionFacturaBoletoAereoBuilder permite construir una solicitud de reversión de anulación para boletos aéreos.
type reversionAnulacionFacturaBoletoAereoBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCodigoAmbiente(codigoAmbiente int) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCodigoEmision(codigoEmision int) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCodigoModalidad(codigoModalidad int) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCodigoSistema(codigoSistema string) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCodigoSucursal(codigoSucursal int) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCufd(cufd string) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCuis(cuis string) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithNit(nit int64) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithTipoFacturaDocumento(tipo int) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el CUF de la factura cuya anulación se desea revertir.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) WithCuf(cuf string) *reversionAnulacionFacturaBoletoAereoBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

// Build construye la solicitud opaca para la reversión de anulación de boleto aéreo.
func (b *reversionAnulacionFacturaBoletoAereoBuilder) Build() ReversionAnulacionFacturaBoletoAereo {
	return ReversionAnulacionFacturaBoletoAereo{RequestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

// validacionRecepcionMasivaFacturaBoletoAereoBuilder permite construir una solicitud de validación para envíos masivos de boletos.
type validacionRecepcionMasivaFacturaBoletoAereoBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCufd(cufd string) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCuis(cuis string) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithNit(nit int64) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithTipoFacturaDocumento(tipo int) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipo
	return b
}

// WithCodigoRecepcion establece el código de recepción masiva a validar.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) WithCodigoRecepcion(codigo string) *validacionRecepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigo
	return b
}

// Build construye la solicitud opaca para la validación masiva de boletos aéreos.
func (b *validacionRecepcionMasivaFacturaBoletoAereoBuilder) Build() ValidacionRecepcionMasivaFacturaBoletoAereo {
	return ValidacionRecepcionMasivaFacturaBoletoAereo{RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}

// verificacionEstadoFacturaBoletoAereoBuilder permite construir una solicitud de verificación de estado de un boleto aéreo.
type verificacionEstadoFacturaBoletoAereoBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCodigoEmision(codigoEmision int) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCodigoModalidad(codigoModalidad int) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCodigoSistema(codigoSistema string) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCodigoSucursal(codigoSucursal int) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCufd(cufd string) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCuis(cuis string) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithNit(nit int64) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithTipoFacturaDocumento(tipo int) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipo
	return b
}

// WithCuf establece el código único de facturación (CUF) a verificar.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) WithCuf(cuf string) *verificacionEstadoFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

// Build construye la solicitud opaca para la verificación del estado de un boleto aéreo.
func (b *verificacionEstadoFacturaBoletoAereoBuilder) Build() VerificacionEstadoFacturaBoletoAereo {
	return VerificacionEstadoFacturaBoletoAereo{RequestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

// recepcionMasivaFacturaBoletoAereoBuilder permite construir una solicitud de recepción masiva de boletos aéreos.
type recepcionMasivaFacturaBoletoAereoBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente para la recepción masiva.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del sector del documento.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCodigoEmision(codigoEmision int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCodigoSistema(codigoSistema string) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCufd(cufd string) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCuis(cuis string) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithNit(nit int64) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithTipoFacturaDocumento(tipo int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.TipoFacturaDocumento = tipo
	return b
}

// WithArchivo establece el archivo XML (tar.gz) en Base64.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithArchivo(archivo string) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de envío masivo.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithFechaEnvio(fecha time.Time) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.FechaEnvio = datatype.NewTimeSiat(fecha)
	return b
}

// WithHashArchivo establece el hash del archivo masivo.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithHashArchivo(hash string) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.HashArchivo = hash
	return b
}

// WithCantidadFacturas establece la cantidad de facturas.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) WithCantidadFacturas(cantidad int) *recepcionMasivaFacturaBoletoAereoBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidad
	return b
}

// Build construye la solicitud opaca para el envío masivo de facturas.
func (b *recepcionMasivaFacturaBoletoAereoBuilder) Build() RecepcionMasivaFacturaBoletoAereo {
	return RecepcionMasivaFacturaBoletoAereo{RequestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}
