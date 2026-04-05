package services

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/documento_ajuste"
	"github.com/ron86i/go-siat/internal/core/ports"
	"github.com/ron86i/go-siat/pkg/models"
)

type siatDocumentoAjusteService struct {
	url        string
	httpClient *http.Client
}

// RecepcionDocumentoAjuste permite el envío de documentos de ajuste para su validación y recepción por parte del SIAT.
func (s *siatDocumentoAjusteService) RecepcionDocumentoAjuste(ctx context.Context, config ports.Config, opaqueReq models.RecepcionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.RecepcionDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.RecepcionDocumentoAjuste, documento_ajuste.RecepcionDocumentoAjusteResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// AnulacionDocumentoAjuste permite la anulación de documentos de ajuste previamente emitidos y aceptados.
func (s *siatDocumentoAjusteService) AnulacionDocumentoAjuste(ctx context.Context, config ports.Config, opaqueReq models.AnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.AnulacionDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.AnulacionDocumentoAjuste, documento_ajuste.AnulacionDocumentoAjusteResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ReversionAnulacionDocumentoAjuste permite la reversión de la anulación de documentos de ajuste.
func (s *siatDocumentoAjusteService) ReversionAnulacionDocumentoAjuste(ctx context.Context, config ports.Config, opaqueReq models.ReversionAnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.ReversionAnulacionDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.ReversionAnulacionDocumentoAjuste, documento_ajuste.ReversionAnulacionDocumentoAjusteResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificacionEstadoDocumentoAjuste permite verificar el estado de los documentos de ajuste.
func (s *siatDocumentoAjusteService) VerificacionEstadoDocumentoAjuste(ctx context.Context, config ports.Config, opaqueReq models.VerificacionEstadoDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificacionEstadoDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.VerificacionEstadoDocumentoAjuste, documento_ajuste.VerificacionEstadoDocumentoAjusteResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificarComunicacion permite verificar la comunicación con el servicio de documentos de ajuste.
func (s *siatDocumentoAjusteService) VerificarComunicacion(ctx context.Context, config ports.Config, opaqueReq models.VerificarComunicacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificarComunicacionResponse], error) {
	return performSoapRequest[documento_ajuste.VerificarComunicacion, documento_ajuste.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// NewSiatDocumentoAjusteService crea una nueva instancia del servicio de documentos de ajuste.
func NewSiatDocumentoAjusteService(baseUrl string, httpClient *http.Client) (*siatDocumentoAjusteService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}
	// Si no se inyecta un cliente, creamos uno con configuraciones seguras por defecto
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &siatDocumentoAjusteService{
		url:        fullURL(baseUrl, SiatDocumentoAjuste),
		httpClient: httpClient,
	}, nil
}

var _ ports.SiatDocumentoAjusteService = (*siatDocumentoAjusteService)(nil)
