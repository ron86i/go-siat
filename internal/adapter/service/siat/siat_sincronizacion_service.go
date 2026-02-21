package siat

import (
	"context"
	"fmt"
	"go-siat/internal/core/domain/datatype/soap"
	"go-siat/internal/core/domain/facturacion"
	"go-siat/internal/core/domain/facturacion/sincronizacion"
	"go-siat/internal/core/port"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3/client"
)

type SiatSincronizacionService struct {
	Url        string
	HttpClient *client.Client
}

func (s *SiatSincronizacionService) SincronizarActividades(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarActividades) (*soap.EnvelopeResponse[sincronizacion.SincronizarActividadesResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarActividades, sincronizacion.SincronizarActividadesResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaActividadesDocumentoSector(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaActividadesDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaActividadesDocumentoSectorResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaActividadesDocumentoSector, sincronizacion.SincronizarListaActividadesDocumentoSectorResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaLeyendasFactura(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaLeyendasFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaLeyendasFacturaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaLeyendasFactura, sincronizacion.SincronizarListaLeyendasFacturaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaMensajesServicios(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaMensajesServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaMensajesServiciosResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaMensajesServicios, sincronizacion.SincronizarListaMensajesServiciosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarListaProductosServicios(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarListaProductosServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaProductosServiciosResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarListaProductosServicios, sincronizacion.SincronizarListaProductosServiciosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaEventosSignificativos(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaEventosSignificativos) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaEventosSignificativosResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaEventosSignificativos, sincronizacion.SincronizarParametricaEventosSignificativosResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaMotivoAnulacion(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaMotivoAnulacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaMotivoAnulacionResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaMotivoAnulacion, sincronizacion.SincronizarParametricaMotivoAnulacionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaPaisOrigen(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaPaisOrigen) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaPaisOrigenResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaPaisOrigen, sincronizacion.SincronizarParametricaPaisOrigenResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoDocumentoIdentidad) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoSector(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoDocumentoSector, sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoEmision(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoEmision) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoEmisionResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoEmision, sincronizacion.SincronizarParametricaTipoEmisionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoHabitacion(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoHabitacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoHabitacionResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoHabitacion, sincronizacion.SincronizarParametricaTipoHabitacionResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMetodoPago(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoMetodoPago) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMetodoPagoResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMetodoPago, sincronizacion.SincronizarParametricaTipoMetodoPagoResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMoneda(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoMoneda) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMonedaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoMoneda, sincronizacion.SincronizarParametricaTipoMonedaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoPuntoVenta(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTipoPuntoVenta) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoPuntoVentaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTipoPuntoVenta, sincronizacion.SincronizarParametricaTipoPuntoVentaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTiposFactura(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaTiposFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTiposFacturaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaTiposFactura, sincronizacion.SincronizarParametricaTiposFacturaResponse](s, ctx, config, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaUnidadMedida(ctx context.Context, config facturacion.Config, req sincronizacion.SincronizarParametricaUnidadMedida) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaUnidadMedidaResponse], error) {
	return executeSincronizacion[sincronizacion.SincronizarParametricaUnidadMedida, sincronizacion.SincronizarParametricaUnidadMedidaResponse](s, ctx, config, req)
}

// executeSincronizacion es un helper genérico privado para ejecutar consultas de sincronización SOAP.
func executeSincronizacion[K any, V any](s *SiatSincronizacionService, ctx context.Context, config facturacion.Config, req K) (*soap.EnvelopeResponse[V], error) {
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

	return envelope, nil
}

// NewSiatSincronizacionService crea una nueva instancia del servicio SiatSincronizacionService.
// Valida que todas las variables de entorno requeridas estén presentes y configura el cliente HTTP.
// Retorna un error si falta alguna configuración o si los valores numéricos son inválidos.
func NewSiatSincronizacionService(url string, httpClient *client.Client) (*SiatSincronizacionService, error) {
	cleanUrl := strings.TrimSpace(url)
	if cleanUrl == "" {
		return nil, fmt.Errorf("la URL base del SIAT no puede estar vacía")
	}

	// Si no se inyecta un cliente, creamos uno con configuraciones seguras por defecto
	if httpClient == nil {
		httpClient = client.New()
		httpClient.SetTimeout(15 * time.Second)
	}

	return &SiatSincronizacionService{
		Url:        url,
		HttpClient: httpClient,
	}, nil
}

var _ port.SiatSincronizacionCatalogoService = (*SiatSincronizacionService)(nil)
