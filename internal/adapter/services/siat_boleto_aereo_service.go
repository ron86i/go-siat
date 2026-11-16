package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"net/http"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v3/internal/core/ports"
	"github.com/ron86i/go-siat/v3/internal/siatpolicy"

	"github.com/ron86i/go-siat/v3/pkg/models"
)

// SiatBoletoAereoService es la implementación concreta del puerto SiatBoletoAereoService
// para la comunicación con los servicios de Impuestos Nacionales.
type SiatBoletoAereoService struct {
	url        string
	httpClient *http.Client
	config     ports.Config
}

// AnulacionFactura anula una factura electrónicado.
func (s *SiatBoletoAereoService) AnulacionFactura(ctx context.Context, req models.AnulacionFactura) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(SiatBoletoAereo, req, 0, func(r *facturacion.AnulacionFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudAnulacion.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

// RecepcionMasivaFactura recepciona un conjunto de facturas electrónicas de forma masiva.
func (s *SiatBoletoAereoService) RecepcionMasivaFactura(ctx context.Context, req models.RecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(SiatBoletoAereo, req, 3, func(r *facturacion.RecepcionMasivaFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	if count := getInternalRequest[facturacion.RecepcionMasivaFactura](req).SolicitudServicioRecepcionMasiva.CantidadFacturas; count <= 0 || count > siatpolicy.MaxFacturasMasiva {
		return nil, fmt.Errorf("cantidad de facturas masiva fuera de límites")
	}
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatBoletoAereoService) ReversionAnulacionFactura(ctx context.Context, req models.ReversionAnulacionFactura) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(SiatBoletoAereo, req, 0, func(r *facturacion.ReversionAnulacionFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudReversionAnulacion.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatBoletoAereoService) ValidacionRecepcionMasivaFactura(ctx context.Context, req models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(SiatBoletoAereo, req, 3, func(r *facturacion.ValidacionRecepcionMasivaFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioValidacionRecepcionMasivaFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	if getInternalRequest[facturacion.ValidacionRecepcionMasivaFactura](req).SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion == "" {
		return nil, fmt.Errorf("codigo de recepción masiva obligatorio")
	}
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatBoletoAereoService) VerificacionEstadoFactura(ctx context.Context, req models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(SiatBoletoAereo, req, 0, func(r *facturacion.VerificacionEstadoFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioVerificacionEstadoFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *SiatBoletoAereoService) VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionFacturacion) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, s.config, req)
}

// NewSiatBoletoAereoService crea una nueva instancia de SiatBoletoAereoService.
func NewSiatBoletoAereoService(baseUrl string, httpClient *http.Client, config ports.Config) (*SiatBoletoAereoService, error) {
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
		config:     config,
	}, nil
}

var _ ports.SiatBoletoAereoService = (*SiatBoletoAereoService)(nil)
