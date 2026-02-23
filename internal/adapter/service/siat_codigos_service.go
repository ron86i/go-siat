package service

import (
	"context"
	"fmt"
	"log"

	"strings"
	"time"

	"github.com/gofiber/fiber/v3/client"
	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/codigos"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
)

// SiatCodigosService implementa el puerto port.SiatCodigosService para interactuar con el SIAT.
// Esta estructura utiliza un cliente HTTP personalizado para realizar peticiones SOAP a los endpoints de impuestos.
type SiatCodigosService struct {
	// Url es la dirección base del servicio web del SIAT (ej. ambiente de prueba o producción).
	Url string
	// HttpClient es el cliente encargado de gestionar las peticiones HTTP, timeouts y configuraciones de red.
	HttpClient *client.Client
}

// VerificarComunicacion realiza una prueba de conectividad con el servicio de códigos del SIAT.
// Es útil para validar que las credenciales base (Token, URL) y la conexión de red
// estén funcionando correctamente antes de realizar operaciones de negocio.
func (s *SiatCodigosService) VerificarComunicacion(ctx context.Context, config config.Config, req *codigos.VerificarComunicacion) (*soap.EnvelopeResponse[codigos.VerificarComunicacionResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLCodigos(s.Url), client.Config{
		Ctx:  ctx,
		Body: xmlBody,
		Header: map[string]string{
			"Content-Type": "application/xml",
			"apiKey":       fmt.Sprintf("TokenApi %s", config.Token),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP: %w", err)
	}

	return parseSoapResponse[codigos.VerificarComunicacionResponse](resp)
}

// NotificaCertificadoRevocado informa al SIAT la revocación de un certificado digital.
// Este procedimiento es crítico cuando un certificado ha sido comprometido o ya no es válido,
// asegurando que las futuras firmas electrónicas asociadas no sean procesadas.
func (s *SiatCodigosService) NotificaCertificadoRevocado(ctx context.Context, config config.Config, req *codigos.NotificaCertificadoRevocado) (*soap.EnvelopeResponse[codigos.NotificaCertificadoRevocadoResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLCodigos(s.Url), client.Config{
		Ctx:  ctx,
		Body: xmlBody,
		Header: map[string]string{
			"Content-Type": "application/xml",
			"apiKey":       fmt.Sprintf("TokenApi %s", config.Token),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP: %w", err)
	}

	return parseSoapResponse[codigos.NotificaCertificadoRevocadoResponse](resp)
}

// SolicitudCufd solicita el Código Único de Facturación Diaria (CUFD) al SIAT.
// Este código es indispensable para la emisión de facturas y tiene una vigencia de 24 horas.
// Configura automáticamente los parámetros base (Ambiente, Modalidad, Sistema, NIT).
func (s *SiatCodigosService) SolicitudCufd(ctx context.Context, config config.Config, req *codigos.Cufd) (*soap.EnvelopeResponse[codigos.CufdResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLCodigos(s.Url), client.Config{
		Ctx: ctx,
		Header: map[string]string{
			"Content-Type": "application/xml",
			"apiKey":       fmt.Sprintf("TokenApi %s", config.Token),
		},
		Body: xmlBody,
	})
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP: %w", err)
	}

	return parseSoapResponse[codigos.CufdResponse](resp)
}

// SolicitudCufdMasivo permite la generación masiva de Códigos Únicos de Facturación Diaria (CUFD).
// Es especialmente útil para sistemas que gestionan múltiples puntos de venta o sucursales de forma centralizada,
// optimizando el proceso de obtención de credenciales de facturación.
func (s *SiatCodigosService) SolicitudCufdMasivo(ctx context.Context, config config.Config, req *codigos.CufdMasivo) (*soap.EnvelopeResponse[codigos.CufdMasivoResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP POST hacia el servicio de facturación masiva
	resp, err := s.HttpClient.Post(fullURLCodigos(s.Url), client.Config{
		Ctx: ctx,
		Header: map[string]string{
			"Content-Type": "application/xml",
			"apiKey":       fmt.Sprintf("TokenApi %s", config.Token),
		},
		Body: xmlBody,
	})
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP masivo: %w", err)
	}
	return parseSoapResponse[codigos.CufdMasivoResponse](resp)
}

// SolicitudCuis solicita el Código Único de Inicio de Sistemas (CUIS) al SIAT.
// Este código es necesario para iniciar operaciones y tiene una vigencia determinada.
func (s *SiatCodigosService) SolicitudCuis(ctx context.Context, config config.Config, req *codigos.Cuis) (*soap.EnvelopeResponse[codigos.CuisResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLCodigos(s.Url), client.Config{
		Ctx: ctx,
		Header: map[string]string{
			"Content-Type": "application/xml",
			"apiKey":       fmt.Sprintf("TokenApi %s", config.Token),
		},
		Body: xmlBody,
	})
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP: %w", err)
	}

	return parseSoapResponse[codigos.CuisResponse](resp)
}

// SolicitudCuisMasivo permite la generación masiva de Códigos Únicos de Inicio de Sistemas (CUIS).
// Esta función facilita la configuración inicial de múltiples puntos de venta o sucursales de forma simultánea,
// reduciendo la latencia de red y simplificando la gestión de credenciales.
func (s *SiatCodigosService) SolicitudCuisMasivo(ctx context.Context, config config.Config, req *codigos.CuisMasivo) (*soap.EnvelopeResponse[codigos.CuisMasivoResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP POST hacia el servicio de códigos masivos del SIAT
	resp, err := s.HttpClient.Post(fullURLCodigos(s.Url), client.Config{
		Ctx:  ctx,
		Body: xmlBody,
		Header: map[string]string{
			"Content-Type": "application/xml",
			"apiKey":       fmt.Sprintf("TokenApi %s", config.Token),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP cuis masivo: %w", err)
	}
	return parseSoapResponse[codigos.CuisMasivoResponse](resp)
}

// VerificarNit verifica la validez de un Número de Identificación Tributaria (NIT) directamente con el servicio SIAT.
// El proceso incluye la construcción de un sobre SOAP con las credenciales y parámetros de configuración (Ambiente, Modalidad, Sistema),
// la ejecución de una petición HTTP POST y la posterior decodificación de la respuesta XML para determinar si el NIT se encuentra activo.
func (s *SiatCodigosService) VerificarNit(ctx context.Context, config config.Config, req *codigos.VerificarNit) (*soap.EnvelopeResponse[codigos.VerificarNitResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLCodigos(s.Url), client.Config{
		Ctx:  ctx,
		Body: xmlBody,
		Header: map[string]string{
			"Content-Type": "application/xml",
			"apiKey":       fmt.Sprintf("TokenApi %s", config.Token),
		},
	})
	log.Println("Response:", config.Token)
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP: %w", err)
	}

	// Intentar extraer el resultado de la respuesta SOAP
	return parseSoapResponse[codigos.VerificarNitResponse](resp)
}

// NewSiatCodigosService crea una nueva instancia del servicio SiatCodigosService.
func NewSiatCodigosService(url string, httpClient *client.Client) (*SiatCodigosService, error) {
	cleanUrl := strings.TrimSpace(url)
	if cleanUrl == "" {
		return nil, fmt.Errorf("la URL base del SIAT no puede estar vacía")
	}

	// Si no se inyecta un cliente, creamos uno con configuraciones seguras por defecto
	if httpClient == nil {
		httpClient = client.New()
		httpClient.SetTimeout(15 * time.Second)
	}

	return &SiatCodigosService{
		Url:        cleanUrl,
		HttpClient: httpClient,
	}, nil
}

var _ port.SiatCodigosService = (*SiatCodigosService)(nil)
