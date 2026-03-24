package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/internal/core/port"

	"github.com/ron86i/go-siat/pkg/models"
)

type SiatElectronicaService struct {
	url        string
	httpClient *http.Client
}

// AnulacionFactura envía una solicitud al SIAT para anular una factura previamente emitida y aceptada.
func (s *SiatElectronicaService) AnulacionFactura(ctx context.Context, config Config, opaqueReq models.AnulacionFacturaElectronica) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionFactura envía una factura electrónica al SIAT para su validación y recepción.
func (s *SiatElectronicaService) RecepcionFactura(ctx context.Context, config Config, opaqueReq models.RecepcionFacturaElectronica) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ReversionAnulacionFactura envía una solicitud al SIAT para revertir una anulación de factura previamente emitida y aceptada.
func (s *SiatElectronicaService) ReversionAnulacionFactura(ctx context.Context, config Config, opaqueReq models.ReversionAnulacionFacturaElectronica) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionPaqueteFactura envía un paquete de facturas electrónicas al SIAT para su validación y recepción.
func (s *SiatElectronicaService) RecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.RecepcionPaqueteFacturaElectronica) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionPaqueteFactura envía una solicitud al SIAT para validar la recepción de un paquete de facturas.
func (s *SiatElectronicaService) ValidacionRecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaElectronica) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificarComunicacion comprueba la disponibilidad de los servicios del SIAT.
func (s *SiatElectronicaService) VerificarComunicacion(ctx context.Context, config Config, opaqueReq models.VerificarComunicacionElectronica) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionMasivaFactura envía de forma masiva un paquete de facturas al SIAT.
func (s *SiatElectronicaService) RecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.RecepcionMasivaFacturaElectronica) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// ValidacionRecepcionMasivaFactura envía una solicitud al SIAT para validar la recepción masiva de facturas.
func (s *SiatElectronicaService) ValidacionRecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionMasivaFacturaElectronica) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// VerificacionEstadoFactura envía una solicitud al SIAT para verificar el estado de una factura.
func (s *SiatElectronicaService) VerificacionEstadoFactura(ctx context.Context, config Config, opaqueReq models.VerificacionEstadoFacturaElectronica) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// RecepcionAnexosSuministroEnergia envía al SIAT la información detallada de los anexos correspondientes a las recargas de suministro de energía eléctrica.
func (s *SiatElectronicaService) RecepcionAnexosSuministroEnergia(ctx context.Context, config Config, opaqueReq models.RecepcionAnexosSuministroEnergiaElectronica) (*soap.EnvelopeResponse[facturacion.RecepcionAnexosSuministroEnergiaResponse], error) {
	return performSoapRequest[facturacion.RecepcionAnexosSuministroEnergia, facturacion.RecepcionAnexosSuministroEnergiaResponse](ctx, s.httpClient, s.url, config, opaqueReq)
}

// NewSiatElectronicaService crea una nueva instancia del servicio de facturación electrónica.
func NewSiatElectronicaService(baseUrl string, httpClient *http.Client) (*SiatElectronicaService, error) {
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

	return &SiatElectronicaService{
		url:        fullURL(baseUrl, SiatElectronica),
		httpClient: httpClient,
	}, nil
}

var _ port.SiatElectronicaService = (*SiatElectronicaService)(nil)
