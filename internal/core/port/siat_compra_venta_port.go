package port

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/compra_venta"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatCompraVentaService interface {
	// AnulacionFactura anula una factura previamente enviada al SIAT.
	AnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.AnulacionFactura) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)

	// RecepcionFactura envía una factura al SIAT para su procesamiento y validación.
	RecepcionFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionFactura) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error)

	// ReversionAnulacionFactura revierte la anulación de una factura previamente enviada al SIAT.
	// Permite revertir el estado de las facturas digitales que fueron anuladas por error una sola vez.
	ReversionAnulacionFactura(ctx context.Context, config config.Config, req models.ReversionAnulacionFactura) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error)

	// RecepcionPaqueteFactura recibe paquetes de hasta 500 facturas emitidas bajo la modalidad
	// de Facturación Electrónica en Línea. El servicio verifica la validez de los parámetros
	// y la integridad del paquete, retornando un código de recepción si es aceptado,
	// o códigos de error/advertencia en caso contrario.
	RecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error)

	// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas.
	ValidacionRecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error)

	// VerificarComunicacion permite verificar la comunicación con el SIAT.
	VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionCompraVenta) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)

	// RecepcionMasivaFactura recibe paquetes de facturas emitidas bajo la modalidad Electrónica
	// en Línea de forma masiva (hasta 1000 facturas por paquete). La periodicidad del envío
	// (diario, semanal o mensual) se configura en el portal de la Administración Tributaria.
	// Retorna un código de recepción si es aceptado, o códigos de error/advertencia.
	RecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)

	// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
	ValidacionRecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)

	// VerificacionEstadoFactura consulta el estado actual de una factura específica en los registros del SIAT.
	VerificacionEstadoFactura(ctx context.Context, config config.Config, opaqueReq models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)

	RecepcionAnexos(ctx context.Context, config config.Config, opaqueReq models.RecepcionAnexos) (*soap.EnvelopeResponse[compra_venta.RecepcionAnexosResponse], error)
}
