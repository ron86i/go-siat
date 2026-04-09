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

// SiatServicioBasicoService es la implementación concreta del puerto SiatServicioBasicoService
// para la comunicación con los servicios de Impuestos Nacionales.
type SiatServicioBasicoService struct {
	url        string
	httpClient *http.Client
}

// AnulacionFactura envía una solicitud al SIAT para anular una factura de servicios básicos previamente emitida y aceptada.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) AnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.AnulacionFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionFactura envía una factura de servicios básicos al SIAT para su validación y recepción.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) RecepcionFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionMasivaFactura permite el envío de un paquete de facturas (mínimo 501, máximo 2000) de servicios básicos.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) RecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionMasivaFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionPaqueteFactura recibe paquetes de hasta 500 facturas de servicios básicos.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) RecepcionPaqueteFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionPaqueteFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ReversionAnulacionFactura revierte la anulación de una factura de servicios básicos previamente enviada al SIAT.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) ReversionAnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.ReversionAnulacionFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) ValidacionRecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionMasivaFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas de servicios básicos.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) ValidacionRecepcionPaqueteFactura(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificacionEstadoFactura consulta el estado actual de una factura de servicios básicos específica.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) VerificacionEstadoFactura(ctx context.Context, config ports.Config, opaqueReq models.VerificacionEstadoFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
// Implementa [ports.SiatServicioBasicoService].
func (s *SiatServicioBasicoService) VerificarComunicacion(ctx context.Context, config ports.Config, opaqueReq models.VerificarComunicacionServicioBasico) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

func NewSiatServicioBasicoService(baseUrl string, httpClient *http.Client) (*SiatServicioBasicoService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatServicioBasicoService{
		url:        fullURL(baseUrl, SiatServicioBasico),
		httpClient: httpClient,
	}, nil
}

var _ ports.SiatServicioBasicoService = (*SiatServicioBasicoService)(nil)
