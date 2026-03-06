package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"bytes"
	"net/http"

	facturacion_sincronizacion "github.com/ron86i/go-siat/internal/core/domain/facturacion/sincronizacion"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatSincronizacionService struct {
	url        string
	HttpClient *http.Client
}

func (s *SiatSincronizacionService) SincronizarActividades(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarActividadesResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarActividadesResponse, facturacion_sincronizacion.SincronizarActividades](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaActividadesDocumentoSector(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaActividadesDocumentoSectorResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarListaActividadesDocumentoSectorResponse, facturacion_sincronizacion.SincronizarListaActividadesDocumentoSector](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaLeyendasFactura(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaLeyendasFacturaResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarListaLeyendasFacturaResponse, facturacion_sincronizacion.SincronizarListaLeyendasFactura](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaMensajesServicios(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaMensajesServiciosResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarListaMensajesServiciosResponse, facturacion_sincronizacion.SincronizarListaMensajesServicios](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaProductosServicios(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaProductosServiciosResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarListaProductosServiciosResponse, facturacion_sincronizacion.SincronizarListaProductosServicios](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaEventosSignificativos(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaEventosSignificativosResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaEventosSignificativosResponse, facturacion_sincronizacion.SincronizarParametricaEventosSignificativos](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaMotivoAnulacion(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaMotivoAnulacionResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaMotivoAnulacionResponse, facturacion_sincronizacion.SincronizarParametricaMotivoAnulacion](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaPaisOrigen(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaPaisOrigenResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaPaisOrigenResponse, facturacion_sincronizacion.SincronizarParametricaPaisOrigen](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse, facturacion_sincronizacion.SincronizarParametricaTipoDocumentoIdentidad](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoSector(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse, facturacion_sincronizacion.SincronizarParametricaTipoDocumentoSector](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoEmision(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoEmisionResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTipoEmisionResponse, facturacion_sincronizacion.SincronizarParametricaTipoEmision](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoHabitacion(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoHabitacionResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTipoHabitacionResponse, facturacion_sincronizacion.SincronizarParametricaTipoHabitacion](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMetodoPago(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoMetodoPagoResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTipoMetodoPagoResponse, facturacion_sincronizacion.SincronizarParametricaTipoMetodoPago](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMoneda(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoMonedaResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTipoMonedaResponse, facturacion_sincronizacion.SincronizarParametricaTipoMoneda](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoPuntoVenta(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoPuntoVentaResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTipoPuntoVentaResponse, facturacion_sincronizacion.SincronizarParametricaTipoPuntoVenta](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTiposFactura(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTiposFacturaResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaTiposFacturaResponse, facturacion_sincronizacion.SincronizarParametricaTiposFactura](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaUnidadMedida(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaUnidadMedidaResponse, error) {
	return executeSincronizacion[facturacion_sincronizacion.SincronizarParametricaUnidadMedidaResponse, facturacion_sincronizacion.SincronizarParametricaUnidadMedida](s, ctx, config, req)
}

// executeSincronizacion es un helper genérico privado para ejecutar consultas de sincronización SOAP.
func executeSincronizacion[V any, K any](s *SiatSincronizacionService, ctx context.Context, config config.Config, req any) (*V, error) {
	concreteReq := models.GetInternalRequest[K](req)
	if concreteReq == nil {
		return nil, fmt.Errorf("tipo de solicitud inválido")
	}

	xmlBody, err := buildRequest(concreteReq)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, fmt.Errorf("error al crear petición HTTP: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error al hacer request HTTP: %w", err)
	}

	envelope, err := parseSoapResponse[V](resp)
	if err != nil {
		return nil, err
	}

	return &envelope.Body.Content, nil
}

// NewSiatSincronizacionService crea una nueva instancia del servicio SiatSincronizacionService.
func NewSiatSincronizacionService(baseUrl string, httpClient *http.Client) (*SiatSincronizacionService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("la URL base del SIAT no puede estar vacía")
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
