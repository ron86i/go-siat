package port

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/codigos"
	"github.com/ron86i/go-siat/pkg/config"
)

// SiatCodigosService define el contrato para la gestión de códigos de facturación ante el SIAT.
// Este puerto incluye operaciones para la obtención de CUIS y CUFD, validación de NITs
// y comunicación de eventos relacionados con la vigencia de certificados digitales.
type SiatCodigosService interface {
	// NotificaCertificadoRevocado comunica al SIAT la revocación de un certificado digital específico.
	NotificaCertificadoRevocado(ctx context.Context, config config.Config, req *codigos.NotificaCertificadoRevocado) (*soap.EnvelopeResponse[codigos.NotificaCertificadoRevocadoResponse], error)

	// SolicitudCufd gestiona la obtención del Código Único de Facturación Diaria para un punto de venta.
	SolicitudCufd(ctx context.Context, config config.Config, req *codigos.Cufd) (*soap.EnvelopeResponse[codigos.CufdResponse], error)

	// SolicitudCufdMasivo permite la tramitación simultánea de múltiples códigos CUFD.
	SolicitudCufdMasivo(ctx context.Context, config config.Config, req *codigos.CufdMasivo) (*soap.EnvelopeResponse[codigos.CufdMasivoResponse], error)

	// SolicitudCuis obtiene el Código Único de Inicio de Sistemas necesario para operar ante el SIAT.
	SolicitudCuis(ctx context.Context, config config.Config, req *codigos.Cuis) (*soap.EnvelopeResponse[codigos.CuisResponse], error)

	// SolicitudCuisMasivo facilita la generación de múltiples códigos CUIS en una sola operación.
	SolicitudCuisMasivo(ctx context.Context, config config.Config, req *codigos.CuisMasivo) (*soap.EnvelopeResponse[codigos.CuisMasivoResponse], error)

	// VerificarNit valida si un Número de Identificación Tributaria se encuentra activo y habilitado.
	VerificarNit(ctx context.Context, config config.Config, req *codigos.VerificarNit) (*soap.EnvelopeResponse[codigos.VerificarNitResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con el servicio de códigos del SIAT.
	VerificarComunicacion(ctx context.Context, config config.Config, req *codigos.VerificarComunicacion) (*soap.EnvelopeResponse[codigos.VerificarComunicacionResponse], error)
}
