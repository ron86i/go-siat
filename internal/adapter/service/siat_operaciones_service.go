package service

import (
	"context"
	"fmt"

	"strings"
	"time"

	"net/http"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/operaciones"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatOperacionesService struct {
	url        string
	httpClient *http.Client
}

// ConsultaPuntoVenta envia una solicitud al SIAT para consultar los puntos de venta registrados.
func (s *SiatOperacionesService) ConsultaPuntoVenta(ctx context.Context, config config.Config, opaqueReq models.ConsultaPuntoVenta) (*soap.EnvelopeResponse[operaciones.ConsultaPuntoVentaResponse], error) {
	return performSoapRequest[operaciones.ConsultaPuntoVenta, operaciones.ConsultaPuntoVentaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// CierreOperacionesSistema envia una solicitud al SIAT para cerrar las operaciones del sistema.
func (s *SiatOperacionesService) CierreOperacionesSistema(ctx context.Context, config config.Config, opaqueReq models.CierreOperacionesSistema) (*soap.EnvelopeResponse[operaciones.CierreOperacionesSistemaResponse], error) {
	return performSoapRequest[operaciones.CierreOperacionesSistema, operaciones.CierreOperacionesSistemaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// CierrePuntoVenta envia una solicitud al SIAT para cerrar un punto de venta.
func (s *SiatOperacionesService) CierrePuntoVenta(ctx context.Context, config config.Config, opaqueReq models.CierrePuntoVenta) (*soap.EnvelopeResponse[operaciones.CierrePuntoVentaResponse], error) {
	return performSoapRequest[operaciones.CierrePuntoVenta, operaciones.CierrePuntoVentaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// ConsultaEventosSignificativos envia una solicitud al SIAT para consultar los eventos significativos registrados.
func (s *SiatOperacionesService) ConsultaEventosSignificativos(ctx context.Context, config config.Config, opaqueReq models.ConsultaEventoSignificativo) (*soap.EnvelopeResponse[operaciones.ConsultaEventoSignificativoResponse], error) {
	return performSoapRequest[operaciones.ConsultaEventoSignificativo, operaciones.ConsultaEventoSignificativoResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RegistroEventosSignificativos envia una solicitud al SIAT para registrar un evento significativo.
func (s *SiatOperacionesService) RegistroEventosSignificativos(ctx context.Context, config config.Config, opaqueReq models.RegistroEventoSignificativo) (*soap.EnvelopeResponse[operaciones.RegistroEventoSignificativoResponse], error) {
	return performSoapRequest[operaciones.RegistroEventoSignificativo, operaciones.RegistroEventoSignificativoResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// VerificarComunicacion envia una solicitud al SIAT para verificar la comunicación.
func (s *SiatOperacionesService) VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionOperaciones) (*soap.EnvelopeResponse[operaciones.VerificarComunicacionResponse], error) {
	return performSoapRequest[operaciones.VerificarComunicacion, operaciones.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RegistroPuntoVenta envia una solicitud al SIAT para registrar un punto de venta.
func (s *SiatOperacionesService) RegistroPuntoVenta(ctx context.Context, config config.Config, opaqueReq models.RegistroPuntoVenta) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaResponse], error) {
	return performSoapRequest[operaciones.RegistroPuntoVenta, operaciones.RegistroPuntoVentaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RegistroPuntoVentaComisionista envia una solicitud al SIAT para registrar un punto de venta comisionista.
func (s *SiatOperacionesService) RegistroPuntoVentaComisionista(ctx context.Context, config config.Config, opaqueReq models.RegistroPuntoVentaComisionista) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaComisionistaResponse], error) {
	return performSoapRequest[operaciones.RegistroPuntoVentaComisionista, operaciones.RegistroPuntoVentaComisionistaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// NewSiatOperacionesService crea una nueva instancia de SiatOperacionesService.
func NewSiatOperacionesService(baseUrl string, httpClient *http.Client) (*SiatOperacionesService, error) {
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

	return &SiatOperacionesService{
		url:        fullURL(baseUrl, SiatOperaciones),
		httpClient: httpClient,
	}, nil
}

var _ port.SiatOperacionesPort = (*SiatOperacionesService)(nil)
