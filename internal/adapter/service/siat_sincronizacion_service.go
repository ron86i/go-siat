package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3/client"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/sincronizacion"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
)

type SiatSincronizacionService struct {
	Url        string
	HttpClient *client.Client
}

func (s *SiatSincronizacionService) SincronizarActividades(ctx context.Context, config config.Config, req *sincronizacion.SincronizarActividades) (*sincronizacion.SincronizarActividadesResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarActividadesResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaActividadesDocumentoSector(ctx context.Context, config config.Config, req *sincronizacion.SincronizarListaActividadesDocumentoSector) (*sincronizacion.SincronizarListaActividadesDocumentoSectorResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaActividadesDocumentoSectorResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaLeyendasFactura(ctx context.Context, config config.Config, req *sincronizacion.SincronizarListaLeyendasFactura) (*sincronizacion.SincronizarListaLeyendasFacturaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaLeyendasFacturaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaMensajesServicios(ctx context.Context, config config.Config, req *sincronizacion.SincronizarListaMensajesServicios) (*sincronizacion.SincronizarListaMensajesServiciosResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaMensajesServiciosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaProductosServicios(ctx context.Context, config config.Config, req *sincronizacion.SincronizarListaProductosServicios) (*sincronizacion.SincronizarListaProductosServiciosResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaProductosServiciosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaEventosSignificativos(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaEventosSignificativos) (*sincronizacion.SincronizarParametricaEventosSignificativosResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaEventosSignificativosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaMotivoAnulacion(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaMotivoAnulacion) (*sincronizacion.SincronizarParametricaMotivoAnulacionResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaMotivoAnulacionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaPaisOrigen(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaPaisOrigen) (*sincronizacion.SincronizarParametricaPaisOrigenResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaPaisOrigenResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTipoDocumentoIdentidad) (*sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoSector(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTipoDocumentoSector) (*sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoEmision(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTipoEmision) (*sincronizacion.SincronizarParametricaTipoEmisionResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoEmisionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoHabitacion(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTipoHabitacion) (*sincronizacion.SincronizarParametricaTipoHabitacionResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoHabitacionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMetodoPago(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTipoMetodoPago) (*sincronizacion.SincronizarParametricaTipoMetodoPagoResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMetodoPagoResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMoneda(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTipoMoneda) (*sincronizacion.SincronizarParametricaTipoMonedaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMonedaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoPuntoVenta(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTipoPuntoVenta) (*sincronizacion.SincronizarParametricaTipoPuntoVentaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoPuntoVentaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTiposFactura(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaTiposFactura) (*sincronizacion.SincronizarParametricaTiposFacturaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTiposFacturaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaUnidadMedida(ctx context.Context, config config.Config, req *sincronizacion.SincronizarParametricaUnidadMedida) (*sincronizacion.SincronizarParametricaUnidadMedidaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaUnidadMedidaResponse](s, ctx, config, req)
}

// executeSincronizacion es un helper genérico privado para ejecutar consultas de sincronización SOAP.
func executeSincronizacion[V any, K any](s *SiatSincronizacionService, ctx context.Context, config config.Config, req K) (*V, error) {
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.HttpClient.Post(fullURLSincronizacion(s.Url), client.Config{
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

	envelope, err := parseSoapResponse[V](resp)
	if err != nil {
		return nil, err
	}

	return &envelope.Body.Content, nil
}

// NewSiatSincronizacionService crea una nueva instancia del servicio SiatSincronizacionService.
func NewSiatSincronizacionService(url string, httpClient *client.Client) (*SiatSincronizacionService, error) {
	cleanUrl := strings.TrimSpace(url)
	if cleanUrl == "" {
		return nil, fmt.Errorf("la URL base del SIAT no puede estar vacía")
	}

	if httpClient == nil {
		httpClient = client.New()
		httpClient.SetTimeout(15 * time.Second)
	}

	return &SiatSincronizacionService{
		Url:        cleanUrl,
		HttpClient: httpClient,
	}, nil
}

var _ port.SiatSincronizacionCatalogoService = (*SiatSincronizacionService)(nil)
