package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/siat/codigos"
)

// --- Interfaces opacas para restringir el acceso a los atributos ---

// VerificarComunicacionCodigos representa una solicitud para verificar la comunicación con el SIAT.
type VerificarComunicacionCodigos struct {
	requestWrapper[codigos.VerificarComunicacion]
}

// VerificarNit representa una solicitud para validar un NIT.
type VerificarNit struct {
	requestWrapper[codigos.VerificarNit]
}

// Cuis representa una solicitud para el Código Único de Inicio de Sistemas.
type Cuis struct {
	requestWrapper[codigos.Cuis]
}

// Cufd representa una solicitud para el Código Único de Facturación Diaria.
type Cufd struct {
	requestWrapper[codigos.Cufd]
}

// CuisMasivo representa una solicitud masiva de CUIS.
type CuisMasivo struct {
	requestWrapper[codigos.CuisMasivo]
}

// CufdMasivo representa una solicitud masiva de CUFD.
type CufdMasivo struct {
	requestWrapper[codigos.CufdMasivo]
}

// NotificaCertificadoRevocado representa una notificación de certificado revocado.
type NotificaCertificadoRevocado struct {
	requestWrapper[codigos.NotificaCertificadoRevocado]
}

type codigosNamespace struct{}

// Codigos expone constructores de solicitudes para el módulo de Gestión de Códigos del SIAT.
func Codigos() codigosNamespace {
	return codigosNamespace{}
}

// NewVerificarNitBuilder inicia la construcción de una solicitud para validar un NIT.
func (codigosNamespace) NewVerificarNitBuilder() *VerificarNitBuilder {
	return &VerificarNitBuilder{
		request: &codigos.VerificarNit{},
	}
}

// VerificarNitBuilder facilita la configuración de la validación de un NIT.
type VerificarNitBuilder struct {
	request *codigos.VerificarNit
}

func (b *VerificarNitBuilder) WithCodigoAmbiente(codigoAmbiente int) *VerificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *VerificarNitBuilder) WithCodigoModalidad(codigoModalidad int) *VerificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoModalidad = codigoModalidad
	return b
}

func (b *VerificarNitBuilder) WithCodigoSistema(codigoSistema string) *VerificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoSistema = codigoSistema
	return b
}

func (b *VerificarNitBuilder) WithCodigoSucursal(codigoSucursal int) *VerificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoSucursal = codigoSucursal
	return b
}

func (b *VerificarNitBuilder) WithCuis(cuis string) *VerificarNitBuilder {
	b.request.SolicitudVerificarNit.Cuis = cuis
	return b
}

func (b *VerificarNitBuilder) WithNit(nit int64) *VerificarNitBuilder {
	b.request.SolicitudVerificarNit.Nit = nit
	return b
}

func (b *VerificarNitBuilder) WithNitParaVerificacion(nitParaVerificacion int64) *VerificarNitBuilder {
	b.request.SolicitudVerificarNit.NitParaVerificacion = nitParaVerificacion
	return b
}

// Build retorna la solicitud de verificación de NIT lista para ser enviada.
func (b *VerificarNitBuilder) Build() VerificarNit {
	return VerificarNit{requestWrapper[codigos.VerificarNit]{request: b.request}}
}

// NewCuisBuilder inicia la construcción de una solicitud para el Código Único de Inicio de Sistemas.
func (codigosNamespace) NewCuisBuilder() *CuisBuilder {
	return &CuisBuilder{
		request: &codigos.Cuis{},
	}
}

// CuisBuilder ayuda a configurar los parámetros para solicitar un CUIS.
type CuisBuilder struct {
	request *codigos.Cuis
}

func (b *CuisBuilder) WithCodigoAmbiente(codigoAmbiente int) *CuisBuilder {
	b.request.SolicitudCuis.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *CuisBuilder) WithCodigoModalidad(codigoModalidad int) *CuisBuilder {
	b.request.SolicitudCuis.CodigoModalidad = codigoModalidad
	return b
}

func (b *CuisBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *CuisBuilder {
	b.request.SolicitudCuis.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *CuisBuilder) WithCodigoSucursal(codigoSucursal int) *CuisBuilder {
	b.request.SolicitudCuis.CodigoSucursal = codigoSucursal
	return b
}

func (b *CuisBuilder) WithCodigoSistema(codigoSistema string) *CuisBuilder {
	b.request.SolicitudCuis.CodigoSistema = codigoSistema
	return b
}

func (b *CuisBuilder) WithNit(nit int64) *CuisBuilder {
	b.request.SolicitudCuis.Nit = nit
	return b
}

// Build entrega el objeto Cuis configurado.
func (b *CuisBuilder) Build() Cuis {
	return Cuis{requestWrapper[codigos.Cuis]{request: b.request}}
}

// NewCufdBuilder inicia la construcción de una solicitud para el Código Único de Facturación Diaria.
func (codigosNamespace) NewCufdBuilder() *CufdBuilder {
	return &CufdBuilder{
		request: &codigos.Cufd{},
	}
}

// CufdBuilder ayuda a configurar los parámetros para solicitar un CUFD.
type CufdBuilder struct {
	request *codigos.Cufd
}

func (b *CufdBuilder) WithCodigoAmbiente(codigoAmbiente int) *CufdBuilder {
	b.request.SolicitudCufd.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *CufdBuilder) WithCodigoModalidad(codigoModalidad int) *CufdBuilder {
	b.request.SolicitudCufd.CodigoModalidad = codigoModalidad
	return b
}

func (b *CufdBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *CufdBuilder {
	b.request.SolicitudCufd.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *CufdBuilder) WithCodigoSistema(codigoSistema string) *CufdBuilder {
	b.request.SolicitudCufd.CodigoSistema = codigoSistema
	return b
}

func (b *CufdBuilder) WithCodigoSucursal(codigoSucursal int) *CufdBuilder {
	b.request.SolicitudCufd.CodigoSucursal = codigoSucursal
	return b
}

func (b *CufdBuilder) WithCuis(cuis string) *CufdBuilder {
	b.request.SolicitudCufd.Cuis = cuis
	return b
}

func (b *CufdBuilder) WithNit(nit int64) *CufdBuilder {
	b.request.SolicitudCufd.Nit = nit
	return b
}

// Build retorna el objeto Cufd configurado.
func (b *CufdBuilder) Build() Cufd {
	return Cufd{requestWrapper[codigos.Cufd]{request: b.request}}
}

// NewCuisMasivoBuilder inicia la construcción de una solicitud masiva de CUIS.
func (codigosNamespace) NewCuisMasivoBuilder() *CuisMasivoBuilder {
	return &CuisMasivoBuilder{
		request: &codigos.CuisMasivo{},
	}
}

// CuisMasivoBuilder facilita la configuración de solicitudes masivas de CUIS.
type CuisMasivoBuilder struct {
	request *codigos.CuisMasivo
}

func (b *CuisMasivoBuilder) WithCodigoAmbiente(codigoAmbiente int) *CuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *CuisMasivoBuilder) WithCodigoModalidad(codigoModalidad int) *CuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.CodigoModalidad = codigoModalidad
	return b
}

func (b *CuisMasivoBuilder) WithCodigoSistema(codigoSistema string) *CuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.CodigoSistema = codigoSistema
	return b
}

func (b *CuisMasivoBuilder) WithNit(nit int64) *CuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.Nit = nit
	return b
}

func (b *CuisMasivoBuilder) WithDatosSolicitud(datosSolicitud []codigos.SolicitudListaCuisDto) *CuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.DatosSolicitud = datosSolicitud
	return b
}

// Build retorna el objeto CuisMasivo configurado.
func (b *CuisMasivoBuilder) Build() CuisMasivo {
	return CuisMasivo{requestWrapper[codigos.CuisMasivo]{request: b.request}}
}

// NewCufdMasivoBuilder inicia la construcción de una solicitud masiva de CUFD.
func (codigosNamespace) NewCufdMasivoBuilder() *CufdMasivoBuilder {
	return &CufdMasivoBuilder{
		request: &codigos.CufdMasivo{},
	}
}

// CufdMasivoBuilder ayuda a configurar la solicitud masiva de códigos CUFD.
type CufdMasivoBuilder struct {
	request *codigos.CufdMasivo
}

func (b *CufdMasivoBuilder) WithCodigoAmbiente(codigoAmbiente int) *CufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *CufdMasivoBuilder) WithCodigoModalidad(codigoModalidad int) *CufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.CodigoModalidad = codigoModalidad
	return b
}

func (b *CufdMasivoBuilder) WithCodigoSistema(codigoSistema string) *CufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.CodigoSistema = codigoSistema
	return b
}

func (b *CufdMasivoBuilder) WithNit(nit int64) *CufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.Nit = nit
	return b
}

func (b *CufdMasivoBuilder) WithDatosSolicitud(datosSolicitud []codigos.SolicitudListaCufdDto) *CufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.DatosSolicitud = datosSolicitud
	return b
}

// Build retorna el objeto CufdMasivo configurado.
func (b *CufdMasivoBuilder) Build() CufdMasivo {
	return CufdMasivo{requestWrapper[codigos.CufdMasivo]{request: b.request}}
}

// NewNotificaCertificadoRevocadoBuilder inicia la construcción de una solicitud para notificar un certificado revocado.
func (codigosNamespace) NewNotificaCertificadoRevocadoBuilder() *NotificaCertificadoRevocadoBuilder {
	return &NotificaCertificadoRevocadoBuilder{
		request: &codigos.NotificaCertificadoRevocado{},
	}
}

// NotificaCertificadoRevocadoBuilder facilita la configuración de la notificación de certificados revocados.
type NotificaCertificadoRevocadoBuilder struct {
	request *codigos.NotificaCertificadoRevocado
}

func (b *NotificaCertificadoRevocadoBuilder) WithCertificado(certificado string) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Certificado = certificado
	return b
}

func (b *NotificaCertificadoRevocadoBuilder) WithCodigoAmbiente(codigoAmbiente int) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *NotificaCertificadoRevocadoBuilder) WithCodigoSistema(codigoSistema string) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.CodigoSistema = codigoSistema
	return b
}

func (b *NotificaCertificadoRevocadoBuilder) WithCodigoSucursal(codigoSucursal int) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.CodigoSucursal = codigoSucursal
	return b
}

func (b *NotificaCertificadoRevocadoBuilder) WithCuis(cuis string) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Cuis = cuis
	return b
}

func (b *NotificaCertificadoRevocadoBuilder) WithNit(nit int64) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Nit = nit
	return b
}

func (b *NotificaCertificadoRevocadoBuilder) WithRazonRevocacion(razonRevocacion string) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.RazonRevocacion = razonRevocacion
	return b
}

func (b *NotificaCertificadoRevocadoBuilder) WithFechaRevocacion(fechaRevocacion *time.Time) *NotificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.FechaRevocacion = fechaRevocacion
	return b
}

// Build retorna el objeto NotificaCertificadoRevocado configurado.
func (b *NotificaCertificadoRevocadoBuilder) Build() NotificaCertificadoRevocado {
	return NotificaCertificadoRevocado{requestWrapper[codigos.NotificaCertificadoRevocado]{request: b.request}}
}

// NewVerificarComunicacionCodigosBuilder inicia la construcción de una prueba de conexión.
func (codigosNamespace) NewVerificarComunicacionCodigosBuilder() *VerificarComunicacionCodigosBuilder {
	return &VerificarComunicacionCodigosBuilder{
		request: &codigos.VerificarComunicacion{},
	}
}

// VerificarComunicacionCodigosBuilder facilita la verificación de comunicación con el SIAT.
type VerificarComunicacionCodigosBuilder struct {
	request *codigos.VerificarComunicacion
}

// Build retorna el objeto de verificación configurado.
func (b *VerificarComunicacionCodigosBuilder) Build() VerificarComunicacionCodigos {
	return VerificarComunicacionCodigos{requestWrapper[codigos.VerificarComunicacion]{request: b.request}}
}
