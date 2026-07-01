package ports

import (
	"context"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/documento_ajuste"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

type SiatDocumentoAjusteService interface {
	RecepcionDocumentoAjuste(ctx context.Context, opaqueReq models.RecepcionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.RecepcionDocumentoAjusteResponse], error)
	AnulacionDocumentoAjuste(ctx context.Context, opaqueReq models.AnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.AnulacionDocumentoAjusteResponse], error)
	ReversionAnulacionDocumentoAjuste(ctx context.Context, opaqueReq models.ReversionAnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.ReversionAnulacionDocumentoAjusteResponse], error)
	VerificacionEstadoDocumentoAjuste(ctx context.Context, opaqueReq models.VerificacionEstadoDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificacionEstadoDocumentoAjusteResponse], error)
	VerificarComunicacion(ctx context.Context, opaqueReq models.VerificarComunicacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificarComunicacionResponse], error)
}
