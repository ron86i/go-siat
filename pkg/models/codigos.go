package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/siat/codigos"
)

// --- Interfaces opacas para la gestión de Códigos ---

type VerificarComunicacionCodigos struct {
	RequestWrapper[codigos.VerificarComunicacion]
}

type VerificarNit struct {
	RequestWrapper[codigos.VerificarNit]
}

type Cuis struct {
	RequestWrapper[codigos.Cuis]
}

type Cufd struct {
	RequestWrapper[codigos.Cufd]
}

type CuisMasivo struct {
	RequestWrapper[codigos.CuisMasivo]
}

type CufdMasivo struct {
	RequestWrapper[codigos.CufdMasivo]
}

type NotificaCertificadoRevocado struct {
	RequestWrapper[codigos.NotificaCertificadoRevocado]
}

// --- Constructores de Builders a nivel de paquete ---

func NewVerificarNitBuilder() *verificarNitBuilder {
	return &verificarNitBuilder{
		request: &codigos.VerificarNit{},
	}
}

func NewCuisBuilder() *cuisBuilder {
	return &cuisBuilder{
		request: &codigos.Cuis{},
	}
}

func NewCufdBuilder() *cufdBuilder {
	return &cufdBuilder{
		request: &codigos.Cufd{},
	}
}

func NewCuisMasivoBuilder() *cuisMasivoBuilder {
	return &cuisMasivoBuilder{
		request: &codigos.CuisMasivo{},
	}
}

func NewCufdMasivoBuilder() *cufdMasivoBuilder {
	return &cufdMasivoBuilder{
		request: &codigos.CufdMasivo{},
	}
}

func NewSolicitudListaCuisDtoBuilder() *SolicitudListaCuisDtoBuilder {
	return &SolicitudListaCuisDtoBuilder{
		request: &codigos.SolicitudListaCuisDto{},
	}
}

func NewSolicitudListaCufdDtoBuilder() *SolicitudListaCufdDtoBuilder {
	return &SolicitudListaCufdDtoBuilder{
		request: &codigos.SolicitudListaCufdDto{},
	}
}

func NewNotificaCertificadoRevocadoBuilder() *notificaCertificadoRevocadoBuilder {
	return &notificaCertificadoRevocadoBuilder{
		request: &codigos.NotificaCertificadoRevocado{},
	}
}

func NewVerificarComunicacionCodigosBuilder() *verificarComunicacionCodigosBuilder {
	return &verificarComunicacionCodigosBuilder{
		request: &codigos.VerificarComunicacion{},
	}
}

// --- Implementaciones de Builders ---

type verificarNitBuilder struct {
	request *codigos.VerificarNit
}

func (b *verificarNitBuilder) WithCodigoSucursal(codigoSucursal int) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoSucursal = codigoSucursal
	return b
}

func (b *verificarNitBuilder) WithCuis(cuis string) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.Cuis = cuis
	return b
}

func (b *verificarNitBuilder) WithNitParaVerificacion(nitParaVerificacion int64) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.NitParaVerificacion = nitParaVerificacion
	return b
}

func (b *verificarNitBuilder) WithCodigoModalidad(codigoModalidad int) *verificarNitBuilder {
	b.request.SolicitudVerificarNit.CodigoModalidad = codigoModalidad
	return b
}

func (b *verificarNitBuilder) Build() VerificarNit {
	return VerificarNit{RequestWrapper: NewRequestWrapper(b.request)}
}

type cuisBuilder struct {
	request *codigos.Cuis
}

func (b *cuisBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cuisBuilder {
	b.request.SolicitudCuis.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *cuisBuilder) WithCodigoSucursal(codigoSucursal int) *cuisBuilder {
	b.request.SolicitudCuis.CodigoSucursal = codigoSucursal
	return b
}

func (b *cuisBuilder) WithCodigoModalidad(codigoModalidad int) *cuisBuilder {
	b.request.SolicitudCuis.CodigoModalidad = codigoModalidad
	return b
}

func (b *cuisBuilder) Build() Cuis {
	return Cuis{RequestWrapper: NewRequestWrapper(b.request)}
}

type cufdBuilder struct {
	request *codigos.Cufd
}

func (b *cufdBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cufdBuilder {
	b.request.SolicitudCufd.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *cufdBuilder) WithCodigoSucursal(codigoSucursal int) *cufdBuilder {
	b.request.SolicitudCufd.CodigoSucursal = codigoSucursal
	return b
}

func (b *cufdBuilder) WithCuis(cuis string) *cufdBuilder {
	b.request.SolicitudCufd.Cuis = cuis
	return b
}

func (b *cufdBuilder) WithCodigoModalidad(codigoModalidad int) *cufdBuilder {
	b.request.SolicitudCufd.CodigoModalidad = codigoModalidad
	return b
}

func (b *cufdBuilder) Build() Cufd {
	return Cufd{RequestWrapper: NewRequestWrapper(b.request)}
}

type cuisMasivoBuilder struct {
	request *codigos.CuisMasivo
}

func (b *cuisMasivoBuilder) WithDatosSolicitud(builders ...*SolicitudListaCuisDtoBuilder) *cuisMasivoBuilder {
	datos := make([]codigos.SolicitudListaCuisDto, len(builders))
	for i, builder := range builders {
		datos[i] = builder.Build()
	}
	b.request.SolicitudCuisMasivoSistemas.DatosSolicitud = datos
	return b
}

func (b *cuisMasivoBuilder) WithCodigoModalidad(codigoModalidad int) *cuisMasivoBuilder {
	b.request.SolicitudCuisMasivoSistemas.CodigoModalidad = codigoModalidad
	return b
}

func (b *cuisMasivoBuilder) Build() CuisMasivo {
	return CuisMasivo{RequestWrapper: NewRequestWrapper(b.request)}
}

type cufdMasivoBuilder struct {
	request *codigos.CufdMasivo
}

func (b *cufdMasivoBuilder) WithDatosSolicitud(builders ...*SolicitudListaCufdDtoBuilder) *cufdMasivoBuilder {
	datos := make([]codigos.SolicitudListaCufdDto, len(builders))
	for i, builder := range builders {
		datos[i] = builder.Build()
	}
	b.request.SolicitudCufdMasivo.DatosSolicitud = datos
	return b
}

func (b *cufdMasivoBuilder) WithCodigoModalidad(codigoModalidad int) *cufdMasivoBuilder {
	b.request.SolicitudCufdMasivo.CodigoModalidad = codigoModalidad
	return b
}

func (b *cufdMasivoBuilder) Build() CufdMasivo {
	return CufdMasivo{RequestWrapper: NewRequestWrapper(b.request)}
}

type notificaCertificadoRevocadoBuilder struct {
	request *codigos.NotificaCertificadoRevocado
}

func (b *notificaCertificadoRevocadoBuilder) WithCertificado(certificado string) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Certificado = certificado
	return b
}

func (b *notificaCertificadoRevocadoBuilder) WithCodigoSucursal(codigoSucursal int) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.CodigoSucursal = codigoSucursal
	return b
}

func (b *notificaCertificadoRevocadoBuilder) WithCuis(cuis string) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.Cuis = cuis
	return b
}

func (b *notificaCertificadoRevocadoBuilder) WithRazonRevocacion(razonRevocacion string) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.RazonRevocacion = razonRevocacion
	return b
}

func (b *notificaCertificadoRevocadoBuilder) WithFechaRevocacion(fechaRevocacion *time.Time) *notificaCertificadoRevocadoBuilder {
	b.request.SolicitudNotificaRevocado.FechaRevocacion = fechaRevocacion
	return b
}

func (b *notificaCertificadoRevocadoBuilder) Build() NotificaCertificadoRevocado {
	return NotificaCertificadoRevocado{RequestWrapper: NewRequestWrapper(b.request)}
}

type SolicitudListaCuisDtoBuilder struct {
	request *codigos.SolicitudListaCuisDto
}

func (b *SolicitudListaCuisDtoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *SolicitudListaCuisDtoBuilder {
	b.request.CodigoPuntoVenta = &codigoPuntoVenta
	return b
}

func (b *SolicitudListaCuisDtoBuilder) WithCodigoSucursal(codigoSucursal int) *SolicitudListaCuisDtoBuilder {
	b.request.CodigoSucursal = codigoSucursal
	return b
}

func (b *SolicitudListaCuisDtoBuilder) Build() codigos.SolicitudListaCuisDto {
	return *b.request
}

type SolicitudListaCufdDtoBuilder struct {
	request *codigos.SolicitudListaCufdDto
}

func (b *SolicitudListaCufdDtoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *SolicitudListaCufdDtoBuilder {
	b.request.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *SolicitudListaCufdDtoBuilder) WithCodigoSucursal(codigoSucursal int) *SolicitudListaCufdDtoBuilder {
	b.request.CodigoSucursal = codigoSucursal
	return b
}

func (b *SolicitudListaCufdDtoBuilder) WithCuis(cuis string) *SolicitudListaCufdDtoBuilder {
	b.request.Cuis = cuis
	return b
}

func (b *SolicitudListaCufdDtoBuilder) Build() codigos.SolicitudListaCufdDto {
	return *b.request
}

type verificarComunicacionCodigosBuilder struct {
	request *codigos.VerificarComunicacion
}

func (b *verificarComunicacionCodigosBuilder) Build() VerificarComunicacionCodigos {
	return VerificarComunicacionCodigos{RequestWrapper: NewRequestWrapper(b.request)}
}
