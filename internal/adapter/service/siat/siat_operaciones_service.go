package siat

import (
	"context"
	"fmt"
	"go-siat/internal/core/domain/datatype/soap"
	"go-siat/internal/core/domain/facturacion"
	"go-siat/internal/core/domain/facturacion/operaciones"
	"go-siat/internal/core/port"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3/client"
)

type SiatOperacionesService struct {
	Url        string
	HttpClient *client.Client
}

// ConsultaPuntoVenta implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) ConsultaPuntoVenta(ctx context.Context, config facturacion.Config, req *operaciones.ConsultaPuntoVenta) (*soap.EnvelopeResponse[operaciones.ConsultaPuntoVentaResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}
	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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

	return parseSoapResponse[operaciones.ConsultaPuntoVentaResponse](resp)
}

// CierreOperacionesSistema implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) CierreOperacionesSistema(ctx context.Context, config facturacion.Config, req *operaciones.CierreOperacionesSistema) (*soap.EnvelopeResponse[operaciones.CierreOperacionesSistemaResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}
	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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

	return parseSoapResponse[operaciones.CierreOperacionesSistemaResponse](resp)
}

// CierrePuntoVenta implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) CierrePuntoVenta(ctx context.Context, config facturacion.Config, req *operaciones.CierrePuntoVenta) (*soap.EnvelopeResponse[operaciones.CierrePuntoVentaResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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

	return parseSoapResponse[operaciones.CierrePuntoVentaResponse](resp)
}

// ConsultaEventosSignificativos implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) ConsultaEventosSignificativos(ctx context.Context, config facturacion.Config, req *operaciones.ConsultaEventoSignificativo) (*soap.EnvelopeResponse[operaciones.ConsultaEventoSignificativoResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}
	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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

	return parseSoapResponse[operaciones.ConsultaEventoSignificativoResponse](resp)
}

// RegistroEventosSignificativos implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) RegistroEventosSignificativos(ctx context.Context, config facturacion.Config, req *operaciones.RegistroEventoSignificativo) (*soap.EnvelopeResponse[operaciones.RegistroEventoSignificativoResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}
	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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

	return parseSoapResponse[operaciones.RegistroEventoSignificativoResponse](resp)
}

// VerificarComunicacion implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) VerificarComunicacion(ctx context.Context, config facturacion.Config) (*soap.EnvelopeResponse[operaciones.VerificarComunicacionResponse], error) {
	req := operaciones.VerificarComunicacion{}
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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

	return parseSoapResponse[operaciones.VerificarComunicacionResponse](resp)
}

// RegistroPuntoVenta implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) RegistroPuntoVenta(ctx context.Context, config facturacion.Config, req *operaciones.RegistroPuntoVenta) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}
	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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

	return parseSoapResponse[operaciones.RegistroPuntoVentaResponse](resp)
}

// RegistroPuntoVentaComisionista implements [port.SiatOperacionesPort].
func (s *SiatOperacionesService) RegistroPuntoVentaComisionista(ctx context.Context, config facturacion.Config, req *operaciones.RegistroPuntoVentaComisionista) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaComisionistaResponse], error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	// Ejecutar la petición HTTP utilizando el cliente configurado
	resp, err := s.HttpClient.Post(fullURLOperaciones(s.Url), client.Config{
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
	return parseSoapResponse[operaciones.RegistroPuntoVentaComisionistaResponse](resp)
}

func NewSiatOperacionesService(url string, httpClient *client.Client) (*SiatOperacionesService, error) {
	cleanUrl := strings.TrimSpace(url)
	if cleanUrl == "" {
		return nil, fmt.Errorf("la URL base del SIAT no puede estar vacía")
	}

	// Si no se inyecta un cliente, creamos uno con configuraciones seguras por defecto
	if httpClient == nil {
		httpClient = client.New()
		httpClient.SetTimeout(15 * time.Second)
	}

	return &SiatOperacionesService{
		Url:        cleanUrl,
		HttpClient: httpClient,
	}, nil
}

var _ port.SiatOperacionesPort = (*SiatOperacionesService)(nil)
