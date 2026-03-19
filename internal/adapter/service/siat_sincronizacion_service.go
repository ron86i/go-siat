package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"net/http"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatSincronizacionService struct {
	url        string
	httpClient *http.Client
}

func (s *SiatSincronizacionService) SincronizarActividades(ctx context.Context, config config.Config, req models.SincronizarActividades) (*soap.EnvelopeResponse[sincronizacion.SincronizarActividadesResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarActividades, sincronizacion.SincronizarActividadesResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarListaActividadesDocumentoSector(ctx context.Context, config config.Config, req models.SincronizarListaActividadesDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaActividadesDocumentoSectorResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarListaActividadesDocumentoSector, sincronizacion.SincronizarListaActividadesDocumentoSectorResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarListaLeyendasFactura(ctx context.Context, config config.Config, req models.SincronizarListaLeyendasFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaLeyendasFacturaResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarListaLeyendasFactura, sincronizacion.SincronizarListaLeyendasFacturaResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarListaMensajesServicios(ctx context.Context, config config.Config, req models.SincronizarListaMensajesServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaMensajesServiciosResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarListaMensajesServicios, sincronizacion.SincronizarListaMensajesServiciosResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarListaProductosServicios(ctx context.Context, config config.Config, req models.SincronizarListaProductosServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaProductosServiciosResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarListaProductosServicios, sincronizacion.SincronizarListaProductosServiciosResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaEventosSignificativos(ctx context.Context, config config.Config, req models.SincronizarParametricaEventosSignificativos) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaEventosSignificativosResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaEventosSignificativos, sincronizacion.SincronizarParametricaEventosSignificativosResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaMotivoAnulacion(ctx context.Context, config config.Config, req models.SincronizarParametricaMotivoAnulacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaMotivoAnulacionResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaMotivoAnulacion, sincronizacion.SincronizarParametricaMotivoAnulacionResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaPaisOrigen(ctx context.Context, config config.Config, req models.SincronizarParametricaPaisOrigen) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaPaisOrigenResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaPaisOrigen, sincronizacion.SincronizarParametricaPaisOrigenResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoDocumentoIdentidad) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoDocumentoSector(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTipoDocumentoSector, sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoEmision(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoEmision) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoEmisionResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTipoEmision, sincronizacion.SincronizarParametricaTipoEmisionResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoHabitacion(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoHabitacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoHabitacionResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTipoHabitacion, sincronizacion.SincronizarParametricaTipoHabitacionResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMetodoPago(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoMetodoPago) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMetodoPagoResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTipoMetodoPago, sincronizacion.SincronizarParametricaTipoMetodoPagoResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoMoneda(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoMoneda) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMonedaResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTipoMoneda, sincronizacion.SincronizarParametricaTipoMonedaResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTipoPuntoVenta(ctx context.Context, config config.Config, req models.SincronizarParametricaTipoPuntoVenta) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoPuntoVentaResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTipoPuntoVenta, sincronizacion.SincronizarParametricaTipoPuntoVentaResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaTiposFactura(ctx context.Context, config config.Config, req models.SincronizarParametricaTiposFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTiposFacturaResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaTiposFactura, sincronizacion.SincronizarParametricaTiposFacturaResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) SincronizarParametricaUnidadMedida(ctx context.Context, config config.Config, req models.SincronizarParametricaUnidadMedida) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaUnidadMedidaResponse], error) {
	return performSoapRequest[sincronizacion.SincronizarParametricaUnidadMedida, sincronizacion.SincronizarParametricaUnidadMedidaResponse](ctx, s.httpClient, s.url, config.Token, req)
}

func (s *SiatSincronizacionService) VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionSincronizacion) (*soap.EnvelopeResponse[sincronizacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[sincronizacion.VerificarComunicacion, sincronizacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
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
		httpClient: httpClient,
	}, nil
}

var _ port.SiatSincronizacionService = (*SiatSincronizacionService)(nil)
