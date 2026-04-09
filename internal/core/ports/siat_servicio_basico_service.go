package ports

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"

	"github.com/ron86i/go-siat/pkg/models"
)

// SiatServicioBasicoService define las operaciones para el servicio de Facturación de Servicios Básicos.
// Estas operaciones permiten la emisión, anulación y validación de facturas bajo la modalidad de servicios básicos.
type SiatServicioBasicoService interface {
	// AnulacionFactura envía una solicitud al SIAT para anular una factura de servicios básicos previamente emitida y aceptada.
	AnulacionFactura(ctx context.Context, config Config, opaqueReq models.AnulacionFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)

	// RecepcionFactura envía una factura de servicios básicos al SIAT para su validación y recepción.
	RecepcionFactura(ctx context.Context, config Config, opaqueReq models.RecepcionFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error)

	// ReversionAnulacionFactura revierte la anulación de una factura de servicios básicos previamente enviada al SIAT.
	ReversionAnulacionFactura(ctx context.Context, config Config, opaqueReq models.ReversionAnulacionFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error)

	// RecepcionPaqueteFactura recibe paquetes de hasta 500 facturas de servicios básicos.
	RecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.RecepcionPaqueteFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error)

	// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas de servicios básicos.
	ValidacionRecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
	VerificarComunicacion(ctx context.Context, config Config, opaqueReq models.VerificarComunicacionServicioBasico) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)

	// RecepcionMasivaFactura permite el envío de un paquete de facturas (mínimo 501, máximo 2000) de servicios básicos.
	RecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.RecepcionMasivaFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)

	// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
	ValidacionRecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionMasivaFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)

	// VerificacionEstadoFactura consulta el estado actual de una factura de servicios básicos específica.
	VerificacionEstadoFactura(ctx context.Context, config Config, opaqueReq models.VerificacionEstadoFacturaServicioBasico) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)
}
