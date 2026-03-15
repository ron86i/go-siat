package port

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatSincronizacionCatalogoService define las operaciones para la sincronización de catálogos y parámetros del SIAT.
// Este puerto permite mantener actualizada la información base necesaria para la emisión de facturas,
// como actividades económicas, leyendas, productos, servicios y parámetros técnicos.
type SiatSincronizacionCatalogoService interface {
	// SincronizarActividades obtiene el listado oficial de actividades económicas autorizadas.
	SincronizarActividades(ctx context.Context, config config.Config, opaqueReq models.SincronizarActividades) (*soap.EnvelopeResponse[sincronizacion.SincronizarActividadesResponse], error)

	// SincronizarListaActividadesDocumentoSector sincroniza la relación entre actividades y tipos de documentos sector.
	SincronizarListaActividadesDocumentoSector(ctx context.Context, config config.Config, opaqueReq models.SincronizarListaActividadesDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaActividadesDocumentoSectorResponse], error)

	// SincronizarListaLeyendasFactura recupera las leyendas oficiales que deben imprimirse en las facturas.
	SincronizarListaLeyendasFactura(ctx context.Context, config config.Config, opaqueReq models.SincronizarListaLeyendasFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaLeyendasFacturaResponse], error)

	// SincronizarListaMensajesServicios sincroniza los mensajes informativos emitidos por los servicios del SIAT.
	SincronizarListaMensajesServicios(ctx context.Context, config config.Config, opaqueReq models.SincronizarListaMensajesServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaMensajesServiciosResponse], error)

	// SincronizarListaProductosServicios obtiene el catálogo de productos y servicios homologados.
	SincronizarListaProductosServicios(ctx context.Context, config config.Config, opaqueReq models.SincronizarListaProductosServicios) (*soap.EnvelopeResponse[sincronizacion.SincronizarListaProductosServiciosResponse], error)

	// SincronizarParametricaEventosSignificativos sincroniza los tipos de eventos significativos (contingencias).
	SincronizarParametricaEventosSignificativos(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaEventosSignificativos) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaEventosSignificativosResponse], error)

	// SincronizarParametricaMotivoAnulacion obtiene los motivos válidos para la anulación de facturas.
	SincronizarParametricaMotivoAnulacion(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaMotivoAnulacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaMotivoAnulacionResponse], error)

	// SincronizarParametricaPaisOrigen sincroniza el catálogo de países de origen para exportaciones.
	SincronizarParametricaPaisOrigen(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaPaisOrigen) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaPaisOrigenResponse], error)

	// SincronizarParametricaTipoDocumentoIdentidad recupera los tipos de documentos de identidad aceptados.
	SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTipoDocumentoIdentidad) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoIdentidadResponse], error)

	// SincronizarParametricaTipoDocumentoSector sincroniza los diferentes tipos de documentos fiscales sectoriales.
	SincronizarParametricaTipoDocumentoSector(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTipoDocumentoSector) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoDocumentoSectorResponse], error)

	// SincronizarParametricaTipoEmision obtiene las modalidades de emisión (en línea, fuera de línea, etc.).
	SincronizarParametricaTipoEmision(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTipoEmision) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoEmisionResponse], error)

	// SincronizarParametricaTipoHabitacion sincroniza los tipos de habitación para el sector hotelero.
	SincronizarParametricaTipoHabitacion(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTipoHabitacion) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoHabitacionResponse], error)

	// SincronizarParametricaTipoMetodoPago obtiene los métodos de pago autorizados (efectivo, tarjeta, etc.).
	SincronizarParametricaTipoMetodoPago(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTipoMetodoPago) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMetodoPagoResponse], error)

	// SincronizarParametricaTipoMoneda sincroniza los tipos de moneda y sus códigos respectivos.
	SincronizarParametricaTipoMoneda(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTipoMoneda) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoMonedaResponse], error)

	// SincronizarParametricaTipoPuntoVenta recupera los tipos de puntos de venta (fijo, móvil, etc.).
	SincronizarParametricaTipoPuntoVenta(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTipoPuntoVenta) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTipoPuntoVentaResponse], error)

	// SincronizarParametricaTiposFactura sincroniza la clasificación de facturas (con derecho a crédito fiscal, sin derecho, etc.).
	SincronizarParametricaTiposFactura(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaTiposFactura) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaTiposFacturaResponse], error)

	// SincronizarParametricaUnidadMedida obtiene las unidades de medida estandarizadas por el SIAT.
	SincronizarParametricaUnidadMedida(ctx context.Context, config config.Config, opaqueReq models.SincronizarParametricaUnidadMedida) (*soap.EnvelopeResponse[sincronizacion.SincronizarParametricaUnidadMedidaResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con el servicio de sincronización del SIAT.
	VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionSincronizacion) (*soap.EnvelopeResponse[sincronizacion.VerificarComunicacionResponse], error)
}
