package ports

import (
	"context"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/sincronizacion"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

// SiatSincronizacionService define las operaciones para la sincronización de catálogos y parámetros del SIAT.
type SiatSincronizacionService interface {
	SincronizarActividades(ctx context.Context, opaqueReq models.SincronizarActividades) (*soap.EnvelopeResponse[sincronizacion.SincronizarActividadesResponse], error)
	SincronizarListaActividadesDocumentoSector(ctx context.Context, opaqueReq models.SincronizarListaActividadesDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaActividadesDocumentoSectorResponse], error)
	SincronizarListaLeyendasFactura(ctx context.Context, opaqueReq models.SincronizarListaLeyendasFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaLeyendasFacturaResponse], error)
	SincronizarListaMensajesServicios(ctx context.Context, opaqueReq models.SincronizarListaMensajesServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaMensajesServiciosResponse], error)
	SincronizarListaProductosServicios(ctx context.Context, opaqueReq models.SincronizarListaProductosServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaProductosServiciosResponse], error)
	SincronizarParametricaEventosSignificativos(ctx context.Context, opaqueReq models.SincronizarParametricaEventosSignificativos) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaEventosSignificativosResponse], error)
	SincronizarParametricaMotivoAnulacion(ctx context.Context, opaqueReq models.SincronizarParametricaMotivoAnulacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaMotivoAnulacionResponse], error)
	SincronizarParametricaPaisOrigen(ctx context.Context, opaqueReq models.SincronizarParametricaPaisOrigen) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaPaisOrigenResponse], error)
	SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, opaqueReq models.SincronizarParametricaTipoDocumentoIdentidad) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse], error)
	SincronizarParametricaTipoDocumentoSector(ctx context.Context, opaqueReq models.SincronizarParametricaTipoDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse], error)
	SincronizarParametricaTipoEmision(ctx context.Context, opaqueReq models.SincronizarParametricaTipoEmision) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoEmisionResponse], error)
	SincronizarParametricaTipoHabitacion(ctx context.Context, opaqueReq models.SincronizarParametricaTipoHabitacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoHabitacionResponse], error)
	SincronizarParametricaTipoMetodoPago(ctx context.Context, opaqueReq models.SincronizarParametricaTipoMetodoPago) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMetodoPagoResponse], error)
	SincronizarParametricaTipoMoneda(ctx context.Context, opaqueReq models.SincronizarParametricaTipoMoneda) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMonedaResponse], error)
	SincronizarParametricaTipoPuntoVenta(ctx context.Context, opaqueReq models.SincronizarParametricaTipoPuntoVenta) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoPuntoVentaResponse], error)
	SincronizarParametricaTiposFactura(ctx context.Context, opaqueReq models.SincronizarParametricaTiposFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTiposFacturaResponse], error)
	SincronizarParametricaUnidadMedida(ctx context.Context, opaqueReq models.SincronizarParametricaUnidadMedida) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaUnidadMedidaResponse], error)
	SincronizarFechaHora(ctx context.Context, opaqueReq models.SincronizarFechaHora) (*soap.EnvelopeResponse[sincronizacion.SincronizarFechaHoraResponse], error)
	VerificarComunicacion(ctx context.Context, opaqueReq models.VerificarComunicacionSincronizacion) (*soap.EnvelopeResponse[sincronizacion.VerificarComunicacionResponse], error)
}
