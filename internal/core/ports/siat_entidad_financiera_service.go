package ports

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatEntidadFinancieraService define el puerto para el servicio de Facturación de Entidades Financieras del SIAT.
// Proporciona los métodos necesarios para la emisión, anulación y validación de facturas para este sector.
type SiatEntidadFinancieraService interface {
	// AnulacionFactura envía una solicitud al SIAT para anular una factura de entidad financiera previamente emitida y aceptada.
	AnulacionFactura(ctx context.Context, config Config, opaqueReq models.AnulacionFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)

	// RecepcionFactura envía una factura de entidad financiera al SIAT para su validación y recepción.
	RecepcionFactura(ctx context.Context, config Config, opaqueReq models.RecepcionFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error)

	// RecepcionMasivaFactura permite el envío de un paquete de facturas de entidad financiera para procesamiento masivo.
	RecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.RecepcionMasivaFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)

	// RecepcionPaqueteFactura recibe paquetes de facturas de entidad financiera.
	RecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.RecepcionPaqueteFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error)

	// ReversionAnulacionFactura revierte la anulación de una factura de entidad financiera previamente enviada al SIAT.
	ReversionAnulacionFactura(ctx context.Context, config Config, opaqueReq models.ReversionAnulacionFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error)

	// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
	ValidacionRecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionMasivaFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)

	// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas de entidad financiera.
	ValidacionRecepcionPaqueteFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error)

	// VerificacionEstadoFactura consulta el estado actual de una factura de entidad financiera específica.
	VerificacionEstadoFactura(ctx context.Context, config Config, opaqueReq models.VerificacionEstadoFacturaEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
	VerificarComunicacion(ctx context.Context, config Config, opaqueReq models.VerificarComunicacionEntidadFinanciera) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)
}
