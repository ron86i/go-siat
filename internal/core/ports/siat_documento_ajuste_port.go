package ports

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/documento_ajuste"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatDocumentoAjusteService interface {
	// RecepcionDocumentoAjuste permite el envío de documentos de ajuste para su validación y recepción por parte del SIAT.
	RecepcionDocumentoAjuste(ctx context.Context, config Config, opaqueReq models.RecepcionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.RecepcionDocumentoAjusteResponse], error)
	// AnulacionDocumentoAjuste permite la anulación de documentos de ajuste previamente emitidos y aceptados.
	AnulacionDocumentoAjuste(ctx context.Context, config Config, opaqueReq models.AnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.AnulacionDocumentoAjusteResponse], error)
	// ReversionAnulacionDocumentoAjuste permite la reversión de la anulación de documentos de ajuste.
	ReversionAnulacionDocumentoAjuste(ctx context.Context, config Config, opaqueReq models.ReversionAnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.ReversionAnulacionDocumentoAjusteResponse], error)
	// VerificacionEstadoDocumentoAjuste permite verificar el estado de los documentos de ajuste.
	VerificacionEstadoDocumentoAjuste(ctx context.Context, config Config, opaqueReq models.VerificacionEstadoDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificacionEstadoDocumentoAjusteResponse], error)
	// VerificarComunicacion permite verificar la comunicación con el servicio de documentos de ajuste.
	VerificarComunicacion(ctx context.Context, config Config, opaqueReq models.VerificarComunicacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificarComunicacionResponse], error)
}
