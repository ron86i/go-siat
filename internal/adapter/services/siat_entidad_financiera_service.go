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

// SiatEntidadFinancieraService es la implementación concreta del puerto SiatEntidadFinancieraService
// para la comunicación con los servicios de Impuestos Nacionales.
type SiatEntidadFinancieraService struct {
	url        string
	httpClient *http.Client
}

// AnulacionFactura envía una solicitud al SIAT para anular una factura de entidad financiera previamente emitida y aceptada.
func (s *SiatEntidadFinancieraService) AnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.AnulacionFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionFactura envía una factura de entidad financiera al SIAT para su validación y recepción.
func (s *SiatEntidadFinancieraService) RecepcionFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionMasivaFactura permite el envío de un paquete de facturas de entidad financiera para procesamiento masivo.
func (s *SiatEntidadFinancieraService) RecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionMasivaFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionPaqueteFactura recibe paquetes de facturas de entidad financiera.
func (s *SiatEntidadFinancieraService) RecepcionPaqueteFactura(ctx context.Context, config ports.Config, opaqueReq models.RecepcionPaqueteFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ReversionAnulacionFactura revierte la anulación de una factura de entidad financiera previamente enviada al SIAT.
func (s *SiatEntidadFinancieraService) ReversionAnulacionFactura(ctx context.Context, config ports.Config, opaqueReq models.ReversionAnulacionFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
func (s *SiatEntidadFinancieraService) ValidacionRecepcionMasivaFactura(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionMasivaFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas de entidad financiera.
func (s *SiatEntidadFinancieraService) ValidacionRecepcionPaqueteFactura(ctx context.Context, config ports.Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificacionEstadoFactura consulta el estado actual de una factura de entidad financiera específica.
func (s *SiatEntidadFinancieraService) VerificacionEstadoFactura(ctx context.Context, config ports.Config, opaqueReq models.VerificacionEstadoFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
func (s *SiatEntidadFinancieraService) VerificarComunicacion(ctx context.Context, config ports.Config, opaqueReq models.VerificarComunicacionEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

func NewSiatEntidadFinancieraService(baseUrl string, httpClient *http.Client) (*SiatEntidadFinancieraService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatEntidadFinancieraService{
		url:        fullURL(baseUrl, SiatEntidadFinanciera),
		httpClient: httpClient,
	}, nil
}

var _ ports.SiatEntidadFinancieraService = (*SiatEntidadFinancieraService)(nil)
