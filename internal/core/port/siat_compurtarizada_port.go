package port

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatComputarizadaService interface {
	// AnulacionFactura envía una solicitud al SIAT para anular una factura previamente emitida y aceptada.
	AnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.AnulacionFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)

	// RecepcionFactura envía una factura computarizada al SIAT para su validación y recepción.
	RecepcionFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error)

	// ReversionAnulacionFactura revierte la anulación de una factura previamente enviada al SIAT.
	// Permite revertir el estado de las facturas digitales que fueron anuladas por error una sola vez.
	ReversionAnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.ReversionAnulacionFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error)

	// RecepcionPaqueteFactura recibe paquetes de hasta 500 facturas emitidas bajo la modalidad
	// de Facturación Electrónica en Línea. El servicio verifica la validez de los parámetros
	// y la integridad del paquete, retornando un código de recepción si es aceptado,
	// o códigos de error/advertencia en caso contrario.
	RecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionPaqueteFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error)

	// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas.
	ValidacionRecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con el servicio de operaciones del SIAT.
	VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionComputarizada) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)

	// RecepcionMasivaFactura permite el envío de un paquete de facturas (mínimo 501, máximo 2000)
	// para su procesamiento masivo por parte del SIAT.
	RecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionMasivaFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)

	// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
	ValidacionRecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionMasivaFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)

	// VerificacionEstadoFactura consulta el estado actual de una factura específica en los registros del SIAT.
	VerificacionEstadoFactura(ctx context.Context, config config.Config, opaqueReq models.VerificacionEstadoFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)

	// RecepcionAnexosSuministroEnergia informa al SIAT sobre los anexos de recargas de suministro
	// de energía eléctrica. Permite registrar el detalle de las recargas asociadas a facturas de
	// suministro, incluyendo montos y fechas de recarga, así como el uso de gift cards.
	RecepcionAnexosSuministroEnergia(ctx context.Context, config config.Config, opaqueReq models.RecepcionAnexosSuministroEnergiaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionAnexosSuministroEnergiaResponse], error)
}
