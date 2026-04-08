package ports

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"

	"github.com/ron86i/go-siat/pkg/models"
)

// SiatTelecomunicacionesService define las operaciones para el servicio de Facturación de Telecomunicaciones.
// Estas operaciones permiten la emisión, anulación y validación de facturas bajo la modalidad de telecomunicaciones.
type SiatTelecomunicacionesService interface {
	// AnulacionFactura envía una solicitud al SIAT para anular una factura de telecomunicaciones previamente emitida y aceptada.
	AnulacionFactura(ctx context.Context, config Config, opaqueReq models.AnulacionFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)

	// RecepcionFactura envía una factura de telecomunicaciones al SIAT para su validación y recepción.
	RecepcionFactura(ctx context.Context, config Config, opaqueReq models.RecepcionFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error)

	// ReversionAnulacionFactura revierte la anulación de una factura de telecomunicaciones previamente enviada al SIAT.
	ReversionAnulacionFactura(ctx context.Context, config Config, opaqueReq models.ReversionAnulacionFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error)

	// RecepcionPaqueteFactura recibe paquetes de hasta 500 facturas de telecomunicaciones.
	RecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.RecepcionPaqueteFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error)

	// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas de telecomunicaciones.
	ValidacionRecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
	VerificarComunicacion(ctx context.Context, config Config, opaqueReq models.VerificarComunicacionTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)

	// RecepcionMasivaFactura permite el envío de un paquete de facturas (mínimo 501, máximo 2000) de telecomunicaciones.
	RecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.RecepcionMasivaFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)

	// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
	ValidacionRecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionMasivaFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)

	// VerificacionEstadoFactura consulta el estado actual de una factura de telecomunicaciones específica.
	VerificacionEstadoFactura(ctx context.Context, config Config, opaqueReq models.VerificacionEstadoFacturaTelecomunicaciones) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)
}
