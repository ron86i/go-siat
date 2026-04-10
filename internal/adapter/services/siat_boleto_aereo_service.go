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

// SiatBoletoAereoService es la implementación concreta del puerto SiatBoletoAereoService
// para la comunicación con los servicios de Impuestos Nacionales.
type SiatBoletoAereoService struct {
	url        string
	httpClient *http.Client
}

// AnulacionFactura envía una solicitud al SIAT para anular un boleto aéreo previamente emitido y aceptada.
func (s *SiatBoletoAereoService) AnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.AnulacionFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionMasivaFactura permite el envío de un paquete de boletos aéreos para procesamiento masivo.
func (s *SiatBoletoAereoService) RecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionMasivaFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ReversionAnulacionFactura revierte la anulación de un boleto aéreo previamente enviada al SIAT.
func (s *SiatBoletoAereoService) ReversionAnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.ReversionAnulacionFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
func (s *SiatBoletoAereoService) ValidacionRecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionMasivaFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificacionEstadoFactura consulta el estado actual de un boleto aéreo específico.
func (s *SiatBoletoAereoService) VerificacionEstadoFactura(ctx context.Context, config ports.Config, opaqueReq models.VerificacionEstadoFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
func (s *SiatBoletoAereoService) VerificarComunicacion(ctx context.Context, config ports.Config, opaqueReq models.VerificarComunicacionBoletoAereo) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

func NewSiatBoletoAereoService(baseUrl string, httpClient *http.Client) (*SiatBoletoAereoService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatBoletoAereoService{
		url:        fullURL(baseUrl, SiatBoletoAereo),
		httpClient: httpClient,
	}, nil
}

var _ ports.SiatBoletoAereoService = (*SiatBoletoAereoService)(nil)
