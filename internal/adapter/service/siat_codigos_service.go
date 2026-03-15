package service

import (
	"context"
	"fmt"

	"strings"
	"time"

	"bytes"
	"net/http"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/codigos"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatCodigosService implementa el puerto port.SiatCodigosService para interactuar con el SIAT.
// Esta estructura utiliza un cliente HTTP personalizado para realizar peticiones SOAP a los endpoints de impuestos.
type SiatCodigosService struct {
	url        string
	HttpClient *http.Client
}

// VerificarComunicacion realiza una prueba de conectividad con el servicio de códigos del SIAT.
// Es útil para validar que las credenciales base (Token, URL) y la conexión de red
// estén funcionando correctamente antes de realizar operaciones de negocio.
func (s *SiatCodigosService) VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionCodigos) (*soap.EnvelopeResponse[codigos.VerificarComunicacionResponse], error) {
	req := models.GetInternalRequest[codigos.VerificarComunicacion](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return parseSoapResponse[codigos.VerificarComunicacionResponse](resp)
}

// NotificaCertificadoRevocado informa al SIAT la revocación de un certificado digital.
// Este procedimiento es crítico cuando un certificado ha sido comprometido o ya no es válido,
// asegurando que las futuras firmas electrónicas asociadas no sean procesadas.
func (s *SiatCodigosService) NotificaCertificadoRevocado(ctx context.Context, config config.Config, opaqueReq models.NotificaCertificadoRevocado) (*soap.EnvelopeResponse[codigos.NotificaCertificadoRevocadoResponse], error) {
	req := models.GetInternalRequest[codigos.NotificaCertificadoRevocado](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return parseSoapResponse[codigos.NotificaCertificadoRevocadoResponse](resp)
}

// SolicitudCufd solicita el Código Único de Facturación Diaria (CUFD) al SIAT.
// Este código es indispensable para la emisión de facturas y tiene una vigencia de 24 horas.
// Configura automáticamente los parámetros base (Ambiente, Modalidad, Sistema, NIT).
func (s *SiatCodigosService) SolicitudCufd(ctx context.Context, config config.Config, opaqueReq models.Cufd) (*soap.EnvelopeResponse[codigos.CufdResponse], error) {
	req := models.GetInternalRequest[codigos.Cufd](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return parseSoapResponse[codigos.CufdResponse](resp)
}

// SolicitudCufdMasivo permite la generación masiva de Códigos Únicos de Facturación Diaria (CUFD).
// Es especialmente útil para sistemas que gestionan múltiples puntos de venta o sucursales de forma centralizada,
// optimizando el proceso de obtención de credenciales de facturación.
func (s *SiatCodigosService) SolicitudCufdMasivo(ctx context.Context, config config.Config, opaqueReq models.CufdMasivo) (*soap.EnvelopeResponse[codigos.CufdMasivoResponse], error) {
	req := models.GetInternalRequest[codigos.CufdMasivo](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[codigos.CufdMasivoResponse](resp)
}

// SolicitudCuis solicita el Código Único de Inicio de Sistemas (CUIS) al SIAT.
// Este código es necesario para iniciar operaciones y tiene una vigencia determinada.
func (s *SiatCodigosService) SolicitudCuis(ctx context.Context, config config.Config, opaqueReq models.Cuis) (*soap.EnvelopeResponse[codigos.CuisResponse], error) {
	req := models.GetInternalRequest[codigos.Cuis](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return parseSoapResponse[codigos.CuisResponse](resp)
}

// SolicitudCuisMasivo permite la generación masiva de Códigos Únicos de Inicio de Sistemas (CUIS).
// Esta función facilita la configuración inicial de múltiples puntos de venta o sucursales de forma simultánea,
// reduciendo la latencia de red y simplificando la gestión de credenciales.
func (s *SiatCodigosService) SolicitudCuisMasivo(ctx context.Context, config config.Config, opaqueReq models.CuisMasivo) (*soap.EnvelopeResponse[codigos.CuisMasivoResponse], error) {
	req := models.GetInternalRequest[codigos.CuisMasivo](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[codigos.CuisMasivoResponse](resp)
}

// VerificarNit verifica la validez de un Número de Identificación Tributaria (NIT) directamente con el servicio SIAT.
// El proceso incluye la construcción de un sobre SOAP con las credenciales y parámetros de configuración (Ambiente, Modalidad, Sistema),
// la ejecución de una petición HTTP POST y la posterior decodificación de la respuesta XML para determinar si el NIT se encuentra activo.
func (s *SiatCodigosService) VerificarNit(ctx context.Context, config config.Config, opaqueReq models.VerificarNit) (*soap.EnvelopeResponse[codigos.VerificarNitResponse], error) {
	req := models.GetInternalRequest[codigos.VerificarNit](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	// Intentar extraer el resultado de la respuesta SOAP
	return parseSoapResponse[codigos.VerificarNitResponse](resp)
}

// NewSiatCodigosService crea una nueva instancia del servicio SiatCodigosService.
func NewSiatCodigosService(baseUrl string, httpClient *http.Client) (*SiatCodigosService, error) {
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

	return &SiatCodigosService{
		url:        fullURL(baseUrl, SiatCodigos),
		HttpClient: httpClient,
	}, nil
}

var _ port.SiatCodigosService = (*SiatCodigosService)(nil)
