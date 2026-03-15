package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"bytes"
	"net/http"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatSincronizacionService struct {
	url        string
	HttpClient *http.Client
}

func (s *SiatSincronizacionService) SincronizarActividades(ctx context.Context, config config.Config, req models.SincronizarActividades) (*soap.EnvelopeResponse[sincronizacion.SincronizarActividadesResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarActividadesResponse, sincronizacion.SincronizarActividades](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaActividadesDocumentoSector(ctx context.Context, config config.Config, req models.SincronizarListaActividadesDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaActividadesDocumentoSectorResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaActividadesDocumentoSectorResponse, sincronizacion.SincronizarListaActividadesDocumentoSector](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaLeyendasFactura(ctx context.Context, config config.Config, req models.SincronizarListaLeyendasFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaLeyendasFacturaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaLeyendasFacturaResponse, sincronizacion.SincronizarListaLeyendasFactura](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaMensajesServicios(ctx context.Context, config config.Config, req models.SincronizarListaMensajesServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaMensajesServiciosResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaMensajesServiciosResponse, sincronizacion.SincronizarListaMensajesServicios](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaProductosServicios(ctx context.Context, config config.Config, req models.SincronizarListaProductosServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaProductosServiciosResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaProductosServiciosResponse, sincronizacion.SincronizarListaProductosServicios](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaEventosSignificativos(ctx context.Context, config config.Config, req models.SincronizarParametricaEventosSignificativos) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaEventosSignificativosResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaEventosSignificativosResponse, sincronizacion.SincronizarParametricaEventosSignificativos](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaMotivoAnulacion(ctx context.Context, config config.Config, req models.SincronizarParametricaMotivoAnulacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaMotivoAnulacionResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaMotivoAnulacionResponse, sincronizacion.SincronizarParametricaMotivoAnulacion](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaPaisOrigen(ctx context.Context, config config.Config, req models.SincronizarParametricaPaisOrigen) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaPaisOrigenResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaPaisOrigenResponse, sincronizacion.SincronizarParametricaPaisOrigen](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoDocumentoIdentidad) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse, sincronizacion.SincronizarParametricaTipoDocumentoIdentidad](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoSector(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse, sincronizacion.SincronizarParametricaTipoDocumentoSector](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoEmision(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoEmision) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoEmisionResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoEmisionResponse, sincronizacion.SincronizarParametricaTipoEmision](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoHabitacion(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoHabitacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoHabitacionResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoHabitacionResponse, sincronizacion.SincronizarParametricaTipoHabitacion](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMetodoPago(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoMetodoPago) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMetodoPagoResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMetodoPagoResponse, sincronizacion.SincronizarParametricaTipoMetodoPago](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMoneda(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoMoneda) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMonedaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMonedaResponse, sincronizacion.SincronizarParametricaTipoMoneda](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoPuntoVenta(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoPuntoVenta) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoPuntoVentaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoPuntoVentaResponse, sincronizacion.SincronizarParametricaTipoPuntoVenta](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTiposFactura(ctx context.Context, config config.Config, req models.SincronizarParametricaTiposFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTiposFacturaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTiposFacturaResponse, sincronizacion.SincronizarParametricaTiposFactura](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaUnidadMedida(ctx context.Context, config config.Config, req models.SincronizarParametricaUnidadMedida) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaUnidadMedidaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaUnidadMedidaResponse, sincronizacion.SincronizarParametricaUnidadMedida](s, ctx, config, req)
}

func (s *SiatSincronizacionService) VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionSincronizacion) (*soap.EnvelopeResponse[sincronizacion.VerificarComunicacionResponse], error) {
	req := models.GetInternalRequest[sincronizacion.VerificarComunicacion](opaqueReq)
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
	return parseSoapResponse[sincronizacion.VerificarComunicacionResponse](resp)
}

// executeSincronizacion es un helper genérico privado para ejecutar consultas de sincronización SOAP.
func executeSincronizacion[V any, K any](s *SiatSincronizacionService, ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[V], error) {
	concreteReq := models.GetInternalRequest[K](req)
	if concreteReq == nil {
		return nil, fmt.Errorf("error to convert request")
	}

	xmlBody, err := buildRequest(concreteReq)
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

	return parseSoapResponse[V](resp)
}

// NewSiatSincronizacionService crea una nueva instancia del servicio SiatSincronizacionService.
func NewSiatSincronizacionService(baseUrl string, httpClient *http.Client) (*SiatSincronizacionService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	return &SiatSincronizacionService{
		url:        fullURL(baseUrl, SiatSincronizacion),
		HttpClient: httpClient,
	}, nil
}

var _ port.SiatSincronizacionCatalogoService = (*SiatSincronizacionService)(nil)
