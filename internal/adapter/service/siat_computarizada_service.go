package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/computarizada"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type siatComputarizadaService struct {
	url        string
	httpClient *http.Client
}

// RecepcionAnexosSuministroEnergia implements [port.SiatComputarizadaService].
func (s *siatComputarizadaService) RecepcionAnexosSuministroEnergia(ctx context.Context, config config.Config, opaqueReq models.RecepcionAnexosSuministroEnergia) (*soap.EnvelopeResponse[computarizada.RecepcionAnexosSuministroEnergiaResponse], error) {
	return performSoapRequest[computarizada.RecepcionAnexosSuministroEnergia, computarizada.RecepcionAnexosSuministroEnergiaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RecepcionMasivaFactura implements [port.SiatComputarizadaService].
func (s *siatComputarizadaService) RecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionMasivaFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// ValidacionRecepcionMasivaFactura envia una solicitud al SIAT para validar la recepción de un paquete de facturas.
func (s *siatComputarizadaService) ValidacionRecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// ValidacionRecepcionPaqueteFactura envia una solicitud al SIAT para validar la recepción de un paquete de facturas.
func (s *siatComputarizadaService) ValidacionRecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// VerificacionEstadoFactura envia una solicitud al SIAT para verificar el estado de una factura.
func (s *siatComputarizadaService) VerificacionEstadoFactura(ctx context.Context, config config.Config, opaqueReq models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// ReversionAnulacionFactura envia una solicitud al SIAT para revertir una anulación de factura previamente emitida y aceptada.
func (s *siatComputarizadaService) ReversionAnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.ReversionAnulacionFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// AnulacionFactura envía una solicitud al SIAT para anular una factura previamente emitida y aceptada.
func (s *siatComputarizadaService) AnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.AnulacionFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RecepcionFactura envía una factura computarizada al SIAT para su validación y recepción.
func (s *siatComputarizadaService) RecepcionFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// VerificarComunicacion comprueba la disponibilidad de los servicios del SIAT.
func (s *siatComputarizadaService) VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionComputarizada) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RecepcionPaqueteFactura envía un paquete de facturas computarizadas al SIAT para su validación y recepción.
func (s *siatComputarizadaService) RecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionPaqueteFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// NewSiatComputarizadaService crea una nueva instancia del servicio de facturación computarizada.
func NewSiatComputarizadaService(baseUrl string, httpClient *http.Client) (*siatComputarizadaService, error) {
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

	return &siatComputarizadaService{
		url:        fullURL(baseUrl, SiatComputarizada),
		httpClient: httpClient,
	}, nil
}

var _ port.SiatComputarizadaService = (*siatComputarizadaService)(nil)
