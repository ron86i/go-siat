package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"net/http"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/documento_ajuste"
	"github.com/ron86i/go-siat/v2/internal/core/ports"

	"github.com/ron86i/go-siat/v2/pkg/models"
)

type SiatDocumentoAjusteService struct {
	url        string
	httpClient *http.Client
	config     ports.Config
}

func (s *SiatDocumentoAjusteService) RecepcionDocumentoAjuste(ctx context.Context, req models.RecepcionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.RecepcionDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.RecepcionDocumentoAjuste, documento_ajuste.RecepcionDocumentoAjusteResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatDocumentoAjusteService) AnulacionDocumentoAjuste(ctx context.Context, req models.AnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.AnulacionDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.AnulacionDocumentoAjuste, documento_ajuste.AnulacionDocumentoAjusteResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatDocumentoAjusteService) ReversionAnulacionDocumentoAjuste(ctx context.Context, req models.ReversionAnulacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.ReversionAnulacionDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.ReversionAnulacionDocumentoAjuste, documento_ajuste.ReversionAnulacionDocumentoAjusteResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatDocumentoAjusteService) VerificacionEstadoDocumentoAjuste(ctx context.Context, req models.VerificacionEstadoDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificacionEstadoDocumentoAjusteResponse], error) {
	return performSoapRequest[documento_ajuste.VerificacionEstadoDocumentoAjuste, documento_ajuste.VerificacionEstadoDocumentoAjusteResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatDocumentoAjusteService) VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionDocumentoAjuste) (*soap.EnvelopeResponse[documento_ajuste.VerificarComunicacionResponse], error) {
	return performSoapRequest[documento_ajuste.VerificarComunicacion, documento_ajuste.VerificarComunicacionResponse](ctx, s.httpClient, s.url, s.config, req)
}

// NewSiatDocumentoAjusteService crea una nueva instancia de SiatDocumentoAjusteService.
func NewSiatDocumentoAjusteService(baseUrl string, httpClient *http.Client, config ports.Config) (*SiatDocumentoAjusteService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatDocumentoAjusteService{
		url:        fullURL(baseUrl, SiatDocumentoAjuste),
		httpClient: httpClient,
		config:     config,
	}, nil
}

var _ ports.SiatDocumentoAjusteService = (*SiatDocumentoAjusteService)(nil)
