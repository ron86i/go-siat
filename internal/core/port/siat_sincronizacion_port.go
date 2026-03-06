package port

import (
	"context"

	facturacion_sincronizacion "github.com/ron86i/go-siat/internal/core/domain/facturacion/sincronizacion"
	"github.com/ron86i/go-siat/pkg/config"
)

// SiatSincronizacionCatalogoService define las operaciones para la sincronización de catálogos y parámetros del SIAT.
type SiatSincronizacionCatalogoService interface {
	SincronizarActividades(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarActividadesResponse, error)
	SincronizarListaActividadesDocumentoSector(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaActividadesDocumentoSectorResponse, error)
	SincronizarListaLeyendasFactura(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaLeyendasFacturaResponse, error)
	SincronizarListaMensajesServicios(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaMensajesServiciosResponse, error)
	SincronizarListaProductosServicios(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarListaProductosServiciosResponse, error)
	SincronizarParametricaEventosSignificativos(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaEventosSignificativosResponse, error)
	SincronizarParametricaMotivoAnulacion(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaMotivoAnulacionResponse, error)
	SincronizarParametricaPaisOrigen(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaPaisOrigenResponse, error)
	SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse, error)
	SincronizarParametricaTipoDocumentoSector(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse, error)
	SincronizarParametricaTipoEmision(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoEmisionResponse, error)
	SincronizarParametricaTipoHabitacion(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoHabitacionResponse, error)
	SincronizarParametricaTipoMetodoPago(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoMetodoPagoResponse, error)
	SincronizarParametricaTipoMoneda(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoMonedaResponse, error)
	SincronizarParametricaTipoPuntoVenta(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTipoPuntoVentaResponse, error)
	SincronizarParametricaTiposFactura(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaTiposFacturaResponse, error)
	SincronizarParametricaUnidadMedida(ctx context.Context, config config.Config, req any) (*facturacion_sincronizacion.SincronizarParametricaUnidadMedidaResponse, error)
}
