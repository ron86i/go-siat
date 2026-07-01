package ports

import (
	"context"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/codigos"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

// SiatCodigosService define el contrato para la gestión de códigos de facturación ante el SIAT.
type SiatCodigosService interface {
	NotificaCertificadoRevocado(ctx context.Context, req models.NotificaCertificadoRevocado) (*soap.EnvelopeResponse[codigos.NotificaCertificadoRevocadoResponse], error)
	SolicitudCufd(ctx context.Context, req models.Cufd) (*soap.EnvelopeResponse[codigos.CufdResponse], error)
	SolicitudCufdMasivo(ctx context.Context, req models.CufdMasivo) (*soap.EnvelopeResponse[codigos.CufdMasivoResponse], error)
	SolicitudCuis(ctx context.Context, req models.Cuis) (*soap.EnvelopeResponse[codigos.CuisResponse], error)
	SolicitudCuisMasivo(ctx context.Context, req models.CuisMasivo) (*soap.EnvelopeResponse[codigos.CuisMasivoResponse], error)
	VerificarNit(ctx context.Context, req models.VerificarNit) (*soap.EnvelopeResponse[codigos.VerificarNitResponse], error)
	VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionCodigos) (*soap.EnvelopeResponse[codigos.VerificarComunicacionResponse], error)
}
