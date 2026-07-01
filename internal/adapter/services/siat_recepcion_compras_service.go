package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"net/http"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v2/internal/core/ports"

	"github.com/ron86i/go-siat/v2/pkg/models"
)

type SiatRecepcionComprasService struct {
	url        string
	httpClient *http.Client
	config     ports.Config
}

func (s *SiatRecepcionComprasService) AnulacionCompra(ctx context.Context, req models.AnulacionCompra) (*soap.EnvelopeResponse[facturacion.AnulacionCompraResponse], error) {
	return performSoapRequest[facturacion.AnulacionCompra, facturacion.AnulacionCompraResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatRecepcionComprasService) ConfirmacionCompras(ctx context.Context, req models.ConfirmacionCompras) (*soap.EnvelopeResponse[facturacion.ConfirmacionComprasResponse], error) {
	return performSoapRequest[facturacion.ConfirmacionCompras, facturacion.ConfirmacionComprasResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatRecepcionComprasService) ConsultaCompras(ctx context.Context, req models.ConsultaCompras) (*soap.EnvelopeResponse[facturacion.ConsultaComprasResponse], error) {
	return performSoapRequest[facturacion.ConsultaCompras, facturacion.ConsultaComprasResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatRecepcionComprasService) RecepcionPaqueteCompras(ctx context.Context, req models.RecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteComprasResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteCompras, facturacion.RecepcionPaqueteComprasResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatRecepcionComprasService) ValidacionRecepcionPaqueteCompras(ctx context.Context, req models.ValidacionRecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteComprasResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteCompras, facturacion.ValidacionRecepcionPaqueteComprasResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatRecepcionComprasService) VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionRecepcionCompras) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacionRecepcionCompras, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, s.config, req)
}

// NewSiatRecepcionComprasService crea una nueva instancia de SiatRecepcionComprasService.
func NewSiatRecepcionComprasService(baseUrl string, httpClient *http.Client, config ports.Config) (*SiatRecepcionComprasService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatRecepcionComprasService{
		url:        fullURL(baseUrl, SiatRecepcionCompras),
		httpClient: httpClient,
		config:     config,
	}, nil
}

var _ ports.SiatRecepcionComprasService = (*SiatRecepcionComprasService)(nil)
