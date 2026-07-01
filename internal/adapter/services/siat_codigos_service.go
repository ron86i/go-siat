package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"net/http"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/codigos"
	"github.com/ron86i/go-siat/internal/core/ports"

	"github.com/ron86i/go-siat/pkg/models"
)

type SiatCodigosService struct {
	url        string
	httpClient *http.Client
	config     ports.Config
}

func (s *SiatCodigosService) NotificaCertificadoRevocado(ctx context.Context, req models.NotificaCertificadoRevocado) (*soap.EnvelopeResponse[codigos.NotificaCertificadoRevocadoResponse], error) {
	return performSoapRequest[codigos.NotificaCertificadoRevocado, codigos.NotificaCertificadoRevocadoResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatCodigosService) SolicitudCufd(ctx context.Context, req models.Cufd) (*soap.EnvelopeResponse[codigos.CufdResponse], error) {
	return performSoapRequest[codigos.Cufd, codigos.CufdResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatCodigosService) SolicitudCufdMasivo(ctx context.Context, req models.CufdMasivo) (*soap.EnvelopeResponse[codigos.CufdMasivoResponse], error) {
	return performSoapRequest[codigos.CufdMasivo, codigos.CufdMasivoResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatCodigosService) SolicitudCuis(ctx context.Context, req models.Cuis) (*soap.EnvelopeResponse[codigos.CuisResponse], error) {
	return performSoapRequest[codigos.Cuis, codigos.CuisResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatCodigosService) SolicitudCuisMasivo(ctx context.Context, req models.CuisMasivo) (*soap.EnvelopeResponse[codigos.CuisMasivoResponse], error) {
	return performSoapRequest[codigos.CuisMasivo, codigos.CuisMasivoResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatCodigosService) VerificarNit(ctx context.Context, req models.VerificarNit) (*soap.EnvelopeResponse[codigos.VerificarNitResponse], error) {
	return performSoapRequest[codigos.VerificarNit, codigos.VerificarNitResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatCodigosService) VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionCodigos) (*soap.EnvelopeResponse[codigos.VerificarComunicacionResponse], error) {
	return performSoapRequest[codigos.VerificarComunicacion, codigos.VerificarComunicacionResponse](ctx, s.httpClient, s.url, s.config, req)
}

// NewSiatCodigosService crea una nueva instancia del servicio de códigos del SIAT.
func NewSiatCodigosService(baseUrl string, httpClient *http.Client, config ports.Config) (*SiatCodigosService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatCodigosService{
		url:        fullURL(baseUrl, SiatCodigos),
		httpClient: httpClient,
		config:     config,
	}, nil
}

var _ ports.SiatCodigosService = (*SiatCodigosService)(nil)
