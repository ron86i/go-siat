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

// SiatTelecomunicacionesService es la implementación concreta del puerto SiatTelecomunicacionesService
// para la comunicación con los servicios de Impuestos Nacionales.
type SiatTelecomunicacionesService struct {
	url        string
	httpClient *http.Client
}

// AnulacionFactura envía una solicitud al SIAT para anular una factura de telecomunicaciones previamente emitida y aceptada.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) AnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.AnulacionFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionFactura envía una factura de telecomunicaciones al SIAT para su validación y recepción.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) RecepcionFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionMasivaFactura permite el envío de un paquete de facturas (mínimo 501, máximo 2000) de telecomunicaciones.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) RecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionMasivaFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionPaqueteFactura recibe paquetes de hasta 500 facturas de telecomunicaciones.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) RecepcionPaqueteFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionPaqueteFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ReversionAnulacionFactura revierte la anulación de una factura de telecomunicaciones previamente enviada al SIAT.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) ReversionAnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.ReversionAnulacionFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) ValidacionRecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionMasivaFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas de telecomunicaciones.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) ValidacionRecepcionPaqueteFactura(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificacionEstadoFactura consulta el estado actual de una factura de telecomunicaciones específica.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) VerificacionEstadoFactura(ctx context.Context, config ports.Config, opaqueReq models.VerificacionEstadoFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
// Implementa [ports.SiatTelecomunicacionesService].
func (s *SiatTelecomunicacionesService) VerificarComunicacion(ctx context.Context, config ports.Config, opaqueReq models.VerificarComunicacionTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

func NewSiatTelecomunicacionesService(baseUrl string, httpClient *http.Client) (*SiatTelecomunicacionesService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatTelecomunicacionesService{
		url:        fullURL(baseUrl, SiatTelecomunicaciones),
		httpClient: httpClient,
	}, nil
}

var _ ports.SiatTelecomunicacionesService = (*SiatTelecomunicacionesService)(nil)
