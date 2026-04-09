package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/siat/codigos"
)

// --- Interfaces opacas para restringir el acceso a los atributos ---

// VerificarComunicacionCodigos representa una solicitud para verificar la comunicación con el SIAT.
type VerificarComunicacionCodigos struct {
	RequestWrapper[codigos.VerificarComunicacion]
}

// VerificarNit representa una solicitud para validar un NIT.
type VerificarNit struct {
	RequestWrapper[codigos.VerificarNit]
}

// Cuis representa una solicitud para el Código Único de Inicio de Sistemas.
type Cuis struct {
	RequestWrapper[codigos.Cuis]
}

// Cufd representa una solicitud para el Código Único de Facturación Diaria.
type Cufd struct {
	RequestWrapper[codigos.Cufd]
}

// CuisMasivo representa una solicitud masiva de CUIS.
type CuisMasivo struct {
	RequestWrapper[codigos.CuisMasivo]
}

// CufdMasivo representa una solicitud masiva de CUFD.
type CufdMasivo struct {
	RequestWrapper[codigos.CufdMasivo]
}

// NotificaCertificadoRevocado representa una notificación de certificado revocado.
type NotificaCertificadoRevocado struct {
	RequestWrapper[codigos.NotificaCertificadoRevocado]
}

// --- Namespace ---

type codigosNamespace struct{}

// Codigos expone constructores de solicitudes para el módulo de Gestión de Códigos del SIAT.
func Codigos() codigosNamespace {
	return codigosNamespace{}
}

// --- Constructores de Builders ---

// NewVerificarNitBuilder inicia la construcción de una solicitud para validar un NIT.
func (codigosNamespace) NewVerificarNitBuilder() *verificarNitBuilder {
	return &verificarNitBuilder{
		request: &codigos.VerificarNit{},
	}
}

// NewCuisBuilder inicia la construcción de una solicitud para el Código Único de Inicio de Sistemas.
func (codigosNamespace) NewCuisBuilder() *cuisBuilder {
	return &cuisBuilder{
		request: &codigos.Cuis{},
	}
}

// NewCufdBuilder inicia la construcción de una solicitud para el Código Único de Facturación Diaria.
func (codigosNamespace) NewCufdBuilder() *cufdBuilder {
	return &cufdBuilder{
		request: &codigos.Cufd{},
	}
}

// NewCuisMasivoBuilder inicia la construcción de una solicitud masiva de CUIS.
func (codigosNamespace) NewCuisMasivoBuilder() *cuisMasivoBuilder {
	return &cuisMasivoBuilder{
		request: &codigos.CuisMasivo{},
	}
}

// NewCufdMasivoBuilder inicia la construcción de una solicitud masiva de CUFD.
func (codigosNamespace) NewCufdMasivoBuilder() *cufdMasivoBuilder {
	return &cufdMasivoBuilder{
		request: &codigos.CufdMasivo{},
	}
}

// NewSolicitudListaCuisDtoBuilder inicia la construcción de una solicitud individual para CUIS masivo.
func (codigosNamespace) NewSolicitudListaCuisDtoBuilder() *SolicitudListaCuisDtoBuilder {
	return &SolicitudListaCuisDtoBuilder{
		request: &codigos.SolicitudListaCuisDto{},
	}
}

// NewSolicitudListaCufdDtoBuilder inicia la construcción de una solicitud individual para CUFD masivo.
func (codigosNamespace) NewSolicitudListaCufdDtoBuilder() *SolicitudListaCufdDtoBuilder {
	return &SolicitudListaCufdDtoBuilder{
		request: &codigos.SolicitudListaCufdDto{},
	}
}

// NewNotificaCertificadoRevocadoBuilder inicia la construcción de una solicitud para notificar un certificado revocado.
func (codigosNamespace) NewNotificaCertificadoRevocadoBuilder() *notificaCertificadoRevocadoBuilder {
	return &notificaCertificadoRevocadoBuilder{
		request: &codigos.NotificaCertificadoRevocado{},
	}
}

// NewVerificarComunicacionCodigosBuilder inicia la construcción de una prueba de conexión.
func (codigosNamespace) NewVerificarComunicacionCodigosBuilder() *verificarComunicacionCodigosBuilder {
	return &verificarComunicacionCodigosBuilder{
		request: &codigos.VerificarComunicacion{},
	}
}

// --- Implementaciones de Builders ---

// verificarNitBuilder facilita la configuración de la validación de un NIT.
type verificarNitBuilder struct {
	request *codigos.VerificarNit
}

// WithCodigoAmbiente establece el código de ambiente (Piloto o Producción).
func (b *verificarNitBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoModalidad establece el código de la modalidad de facturación.
func (b *verificarNitBuilder) WithCodigoModalidad(codigoModalidad int) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoSistema establece el código del sistema autorizado por el SIN.
func (b *verificarNitBuilder) WithCodigoSistema(codigoSistema string) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *verificarNitBuilder) WithCodigoSucursal(codigoSucursal int) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoSucursal = codigoSucursal
	return b
}

// WithCuis establece el Código Único de Inicio de Sistemas.
func (b *verificarNitBuilder) WithCuis(cuis string) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *verificarNitBuilder) WithNit(nit int64) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.Nit = nit
	return b
}

// WithNitParaVerificacion establece el NIT que se desea validar ante el SIN.
func (b *verificarNitBuilder) WithNitParaVerificacion(nitParaVerificacion int64) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.NitParaVerificacion = nitParaVerificacion
	return b
}

// Build retorna la solicitud de verificación de NIT lista para ser enviada.
func (b *verificarNitBuilder) Build() VerificarNit {
	return VerificarNit{RequestWrapper[codigos.VerificarNit]{request: b.request}}
}

// cuisBuilder ayuda a configurar los parámetros para solicitar un CUIS.
type cuisBuilder struct {
	request *codigos.Cuis
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *cuisBuilder) WithCodigoAmbiente(codigoAmbiente int) *cuisBuilder {
	b.request.SolicitudCuis.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *cuisBuilder) WithCodigoModalidad(codigoModalidad int) *cuisBuilder {
	b.request.SolicitudCuis.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *cuisBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cuisBuilder {
	b.request.SolicitudCuis.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *cuisBuilder) WithCodigoSucursal(codigoSucursal int) *cuisBuilder {
	b.request.SolicitudCuis.CodigoSucursal = codigoSucursal
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *cuisBuilder) WithCodigoSistema(codigoSistema string) *cuisBuilder {
	b.request.SolicitudCuis.CodigoSistema = codigoSistema
	return b
}

// WithNit establece el NIT del emisor.
func (b *cuisBuilder) WithNit(nit int64) *cuisBuilder {
	b.request.SolicitudCuis.Nit = nit
	return b
}

// Build entrega el objeto Cuis configurado.
func (b *cuisBuilder) Build() Cuis {
	return Cuis{RequestWrapper[codigos.Cuis]{request: b.request}}
}

// cufdBuilder ayuda a configurar los parámetros para solicitar un CUFD.
type cufdBuilder struct {
	request *codigos.Cufd
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *cufdBuilder) WithCodigoAmbiente(codigoAmbiente int) *cufdBuilder {
	b.request.SolicitudCufd.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *cufdBuilder) WithCodigoModalidad(codigoModalidad int) *cufdBuilder {
	b.request.SolicitudCufd.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *cufdBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cufdBuilder {
	b.request.SolicitudCufd.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *cufdBuilder) WithCodigoSistema(codigoSistema string) *cufdBuilder {
	b.request.SolicitudCufd.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *cufdBuilder) WithCodigoSucursal(codigoSucursal int) *cufdBuilder {
	b.request.SolicitudCufd.CodigoSucursal = codigoSucursal
	return b
}

// WithCuis establece el CUIS.
func (b *cufdBuilder) WithCuis(cuis string) *cufdBuilder {
	b.request.SolicitudCufd.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *cufdBuilder) WithNit(nit int64) *cufdBuilder {
	b.request.SolicitudCufd.Nit = nit
	return b
}

// Build retorna el objeto Cufd configurado.
func (b *cufdBuilder) Build() Cufd {
	return Cufd{RequestWrapper[codigos.Cufd]{request: b.request}}
}

// cuisMasivoBuilder facilita la configuración de solicitudes masivas de CUIS.
type cuisMasivoBuilder struct {
	request *codigos.CuisMasivo
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *cuisMasivoBuilder) WithCodigoAmbiente(codigoAmbiente int) *cuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *cuisMasivoBuilder) WithCodigoModalidad(codigoModalidad int) *cuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *cuisMasivoBuilder) WithCodigoSistema(codigoSistema string) *cuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.CodigoSistema = codigoSistema
	return b
}

// WithNit establece el NIT del emisor.
func (b *cuisMasivoBuilder) WithNit(nit int64) *cuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.Nit = nit
	return b
}

// WithDatosSolicitud establece la lista de sucursales y puntos de venta para los cuales se solicita CUIS.
func (b *cuisMasivoBuilder) WithDatosSolicitud(builders ...*SolicitudListaCuisDtoBuilder) *cuisMasivoBuilder {
	datos := make([]codigos.SolicitudListaCuisDto, len(builders))
	for i, builder := range builders {
		datos[i] = builder.Build()
	}
	b.request.SolicitudCuisMasivoSistemas.DatosSolicitud = datos
	return b
}

// Build retorna el objeto CuisMasivo configurado.
func (b *cuisMasivoBuilder) Build() CuisMasivo {
	return CuisMasivo{RequestWrapper[codigos.CuisMasivo]{request: b.request}}
}

// cufdMasivoBuilder ayuda a configurar la solicitud masiva de códigos CUFD.
type cufdMasivoBuilder struct {
	request *codigos.CufdMasivo
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *cufdMasivoBuilder) WithCodigoAmbiente(codigoAmbiente int) *cufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *cufdMasivoBuilder) WithCodigoModalidad(codigoModalidad int) *cufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *cufdMasivoBuilder) WithCodigoSistema(codigoSistema string) *cufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.CodigoSistema = codigoSistema
	return b
}

// WithNit establece el NIT del emisor.
func (b *cufdMasivoBuilder) WithNit(nit int64) *cufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.Nit = nit
	return b
}

// WithDatosSolicitud establece la lista de sucursales, puntos de venta y CUIS para los cuales se solicita CUFD.
func (b *cufdMasivoBuilder) WithDatosSolicitud(builders ...*SolicitudListaCufdDtoBuilder) *cufdMasivoBuilder {
	datos := make([]codigos.SolicitudListaCufdDto, len(builders))
	for i, builder := range builders {
		datos[i] = builder.Build()
	}
	b.request.SolicitudCufdMasivo.DatosSolicitud = datos
	return b
}

// Build retorna el objeto CufdMasivo configurado.
func (b *cufdMasivoBuilder) Build() CufdMasivo {
	return CufdMasivo{RequestWrapper[codigos.CufdMasivo]{request: b.request}}
}

// notificaCertificadoRevocadoBuilder facilita la configuración de la notificación de certificados revocados.
type notificaCertificadoRevocadoBuilder struct {
	request *codigos.NotificaCertificadoRevocado
}

// WithCertificado establece el certificado revocado en formato Base64.
func (b *notificaCertificadoRevocadoBuilder) WithCertificado(certificado string) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Certificado = certificado
	return b
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *notificaCertificadoRevocadoBuilder) WithCodigoAmbiente(codigoAmbiente int) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *notificaCertificadoRevocadoBuilder) WithCodigoSistema(codigoSistema string) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *notificaCertificadoRevocadoBuilder) WithCodigoSucursal(codigoSucursal int) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.CodigoSucursal = codigoSucursal
	return b
}

// WithCuis establece el CUIS.
func (b *notificaCertificadoRevocadoBuilder) WithCuis(cuis string) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *notificaCertificadoRevocadoBuilder) WithNit(nit int64) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Nit = nit
	return b
}

// WithRazonRevocacion establece el motivo de la revocación del certificado.
func (b *notificaCertificadoRevocadoBuilder) WithRazonRevocacion(razonRevocacion string) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.RazonRevocacion = razonRevocacion
	return b
}

// WithFechaRevocacion establece la fecha y hora de la revocación.
func (b *notificaCertificadoRevocadoBuilder) WithFechaRevocacion(fechaRevocacion *time.Time) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.FechaRevocacion = fechaRevocacion
	return b
}

// Build retorna el objeto NotificaCertificadoRevocado configurado.
func (b *notificaCertificadoRevocadoBuilder) Build() NotificaCertificadoRevocado {
	return NotificaCertificadoRevocado{RequestWrapper[codigos.NotificaCertificadoRevocado]{request: b.request}}
}

// SolicitudListaCuisDtoBuilder facilita la configuración de una solicitud individual de CUIS.
type SolicitudListaCuisDtoBuilder struct {
	request *codigos.SolicitudListaCuisDto
}

// WithCodigoPuntoVenta establece el código del punto de venta para la solicitud masiva.
func (b *SolicitudListaCuisDtoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *SolicitudListaCuisDtoBuilder {
	b.request.CodigoPuntoVenta = &codigoPuntoVenta
	return b
}

// WithCodigoSucursal establece el código de la sucursal para la solicitud masiva.
func (b *SolicitudListaCuisDtoBuilder) WithCodigoSucursal(codigoSucursal int) *SolicitudListaCuisDtoBuilder {
	b.request.CodigoSucursal = codigoSucursal
	return b
}

func (b *SolicitudListaCuisDtoBuilder) Build() codigos.SolicitudListaCuisDto {
	return *b.request
}

// SolicitudListaCufdDtoBuilder facilita la configuración de una solicitud individual de CUFD.
type SolicitudListaCufdDtoBuilder struct {
	request *codigos.SolicitudListaCufdDto
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *SolicitudListaCufdDtoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *SolicitudListaCufdDtoBuilder {
	b.request.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *SolicitudListaCufdDtoBuilder) WithCodigoSucursal(codigoSucursal int) *SolicitudListaCufdDtoBuilder {
	b.request.CodigoSucursal = codigoSucursal
	return b
}

// WithCuis establece el CUIS asociado a la sucursal/punto de venta.
func (b *SolicitudListaCufdDtoBuilder) WithCuis(cuis string) *SolicitudListaCufdDtoBuilder {
	b.request.Cuis = cuis
	return b
}

func (b *SolicitudListaCufdDtoBuilder) Build() codigos.SolicitudListaCufdDto {
	return *b.request
}

// verificarComunicacionCodigosBuilder facilita la verificación de comunicación con el SIAT.
type verificarComunicacionCodigosBuilder struct {
	request *codigos.VerificarComunicacion
}

// Build retorna el objeto de verificación configurado.
func (b *verificarComunicacionCodigosBuilder) Build() VerificarComunicacionCodigos {
	return VerificarComunicacionCodigos{RequestWrapper[codigos.VerificarComunicacion]{request: b.request}}
}
