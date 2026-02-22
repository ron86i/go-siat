package siat

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3/client"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/sincronizacion"
	"github.com/ron86i/go-siat/internal/core/port"
)

type SiatSincronizacionService struct {
	Url        string
	HttpClient *client.Client
}

func (s *SiatSincronizacionService) SincronizarActividades(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarActividades) (*sincronizacion.SincronizarActividadesResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarActividades, sincronizacion.SincronizarActividadesResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaActividadesDocumentoSector(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaActividadesDocumentoSector) (*sincronizacion.SincronizarListaActividadesDocumentoSectorResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaActividadesDocumentoSector, sincronizacion.SincronizarListaActividadesDocumentoSectorResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaLeyendasFactura(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaLeyendasFactura) (*sincronizacion.SincronizarListaLeyendasFacturaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaLeyendasFactura, sincronizacion.SincronizarListaLeyendasFacturaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaMensajesServicios(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaMensajesServicios) (*sincronizacion.SincronizarListaMensajesServiciosResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaMensajesServicios, sincronizacion.SincronizarListaMensajesServiciosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaProductosServicios(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaProductosServicios) (*sincronizacion.SincronizarListaProductosServiciosResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarListaProductosServicios, sincronizacion.SincronizarListaProductosServiciosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaEventosSignificativos(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaEventosSignificativos) (*sincronizacion.SincronizarParametricaEventosSignificativosResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaEventosSignificativos, sincronizacion.SincronizarParametricaEventosSignificativosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaMotivoAnulacion(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaMotivoAnulacion) (*sincronizacion.SincronizarParametricaMotivoAnulacionResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaMotivoAnulacion, sincronizacion.SincronizarParametricaMotivoAnulacionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaPaisOrigen(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaPaisOrigen) (*sincronizacion.SincronizarParametricaPaisOrigenResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaPaisOrigen, sincronizacion.SincronizarParametricaPaisOrigenResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoDocumentoIdentidad) (*sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoSector(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoDocumentoSector) (*sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoSector, sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoEmision(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoEmision) (*sincronizacion.SincronizarParametricaTipoEmisionResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoEmision, sincronizacion.SincronizarParametricaTipoEmisionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoHabitacion(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoHabitacion) (*sincronizacion.SincronizarParametricaTipoHabitacionResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoHabitacion, sincronizacion.SincronizarParametricaTipoHabitacionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMetodoPago(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoMetodoPago) (*sincronizacion.SincronizarParametricaTipoMetodoPagoResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMetodoPago, sincronizacion.SincronizarParametricaTipoMetodoPagoResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMoneda(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoMoneda) (*sincronizacion.SincronizarParametricaTipoMonedaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMoneda, sincronizacion.SincronizarParametricaTipoMonedaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoPuntoVenta(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoPuntoVenta) (*sincronizacion.SincronizarParametricaTipoPuntoVentaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoPuntoVenta, sincronizacion.SincronizarParametricaTipoPuntoVentaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTiposFactura(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTiposFactura) (*sincronizacion.SincronizarParametricaTiposFacturaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTiposFactura, sincronizacion.SincronizarParametricaTiposFacturaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaUnidadMedida(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaUnidadMedida) (*sincronizacion.SincronizarParametricaUnidadMedidaResponse, error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaUnidadMedida, sincronizacion.SincronizarParametricaUnidadMedidaResponse](s, ctx, config, req)
}

// executeSincronizacion es un helper genérico privado para ejecutar consultas de sincronización SOAP.
func executeSincronizacion[K any, V any](s *SiatSincronizacionService, ctx context.Context, config facturacion.Config, req K) (*V, error) {
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
// Valida que todas las variables de entorno requeridas estén presentes y configura el cliente HTTP.
// Retorna un error si falta alguna configuración o si los valores numéricos son inválidos.
func NewSiatSincronizacionService(envs map[string]string) (*SiatSincronizacionService, error) {
	url := strings.TrimSpace(envs["SIAT_URL"])
	if url == "" {
		return nil, fmt.Errorf("la variable de entorno SIAT_URL es obligatoria")
	}
	return &SiatSincronizacionService{
		Url:        url,
		HttpClient: client.New(),
	}, nil
}

var _ port.SiatSincronizacionCatalogoService = (*SiatSincronizacionService)(nil)
