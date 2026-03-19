package port

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/computarizada"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

type SiatComputarizadaService interface {
	AnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.AnulacionFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)

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

	RecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionMasivaFacturaComputarizada) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)

	// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
	ValidacionRecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)

	// VerificacionEstadoFactura consulta el estado actual de una factura específica en los registros del SIAT.
	VerificacionEstadoFactura(ctx context.Context, config config.Config, opaqueReq models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)

	RecepcionAnexosSuministroEnergia(ctx context.Context, config config.Config, opaqueReq models.RecepcionAnexosSuministroEnergia) (*soap.EnvelopeResponse[computarizada.RecepcionAnexosSuministroEnergiaResponse], error)
}
