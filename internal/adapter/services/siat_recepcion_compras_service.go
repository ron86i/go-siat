package services

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/internal/core/ports"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatRecepcionComprasService es la implementación del puerto SiatRecepcionComprasService.
type SiatRecepcionComprasService struct {
	url        string
	httpClient *http.Client
}

// AnulacionCompra implementa [ports.SiatRecepcionComprasService].
func (s *SiatRecepcionComprasService) AnulacionCompra(ctx context.Context, config ports.Config, opaqueReq models.AnulacionCompra) (*soap.EnvelopeResponse[facturacion.AnulacionCompraResponse], error) {
	return performSoapRequest[facturacion.AnulacionCompra, facturacion.AnulacionCompraResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ConfirmacionCompras implementa [ports.SiatRecepcionComprasService].
func (s *SiatRecepcionComprasService) ConfirmacionCompras(ctx context.Context, config ports.Config, opaqueReq models.ConfirmacionCompras) (*soap.EnvelopeResponse[facturacion.ConfirmacionComprasResponse], error) {
	return performSoapRequest[facturacion.ConfirmacionCompras, facturacion.ConfirmacionComprasResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ConsultaCompras implementa [ports.SiatRecepcionComprasService].
func (s *SiatRecepcionComprasService) ConsultaCompras(ctx context.Context, config ports.Config, opaqueReq models.ConsultaCompras) (*soap.EnvelopeResponse[facturacion.ConsultaComprasResponse], error) {
	return performSoapRequest[facturacion.ConsultaCompras, facturacion.ConsultaComprasResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionPaqueteCompras implementa [ports.SiatRecepcionComprasService].
func (s *SiatRecepcionComprasService) RecepcionPaqueteCompras(ctx context.Context, config ports.Config, opaqueReq models.RecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteComprasResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteCompras, facturacion.RecepcionPaqueteComprasResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionPaqueteCompras implementa [ports.SiatRecepcionComprasService].
func (s *SiatRecepcionComprasService) ValidacionRecepcionPaqueteCompras(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteComprasResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteCompras, facturacion.ValidacionRecepcionPaqueteComprasResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificarComunicacion implementa [ports.SiatRecepcionComprasService].
func (s *SiatRecepcionComprasService) VerificarComunicacion(ctx context.Context, config ports.Config, opaqueReq models.VerificarComunicacionRecepcionCompras) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacionRecepcionCompras, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

func NewSiatRecepcionComprasService(baseUrl string, httpClient *http.Client) (*SiatRecepcionComprasService, error) {
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
	}, nil
}

var _ ports.SiatRecepcionComprasService = (*SiatRecepcionComprasService)(nil)
