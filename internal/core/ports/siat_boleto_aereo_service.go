package ports

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatBoletoAereoService define el puerto para el servicio de Facturación de Boletos Aéreos del SIAT.
// Proporciona los métodos necesarios para la emisión masiva, anulación y validación de facturas para este sector.
// Nota: A diferencia de otros sectores, este servicio se centra en la recepción masiva.
type SiatBoletoAereoService interface {
	// AnulacionFactura envía una solicitud al SIAT para anular un boleto aéreo previamente emitido y aceptado.
	AnulacionFactura(ctx context.Context, config Config, opaqueReq models.AnulacionFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)

	// RecepcionMasivaFactura permite el envío de un paquete de boletos aéreos para procesamiento masivo.
	RecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.RecepcionMasivaFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)

	// ReversionAnulacionFactura revierte la anulación de un boleto aéreo previamente enviada al SIAT.
	ReversionAnulacionFactura(ctx context.Context, config Config, opaqueReq models.ReversionAnulacionFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error)

	// ValidacionRecepcionMasivaFactura verifica el estado del procesamiento de un paquete enviado masivamente.
	ValidacionRecepcionMasivaFactura(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionMasivaFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)

	// VerificacionEstadoFactura consulta el estado actual de un boleto aéreo específico.
	VerificacionEstadoFactura(ctx context.Context, config Config, opaqueReq models.VerificacionEstadoFacturaBoletoAereo) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con el servicio de comunicaciones del SIAT.
	VerificarComunicacion(ctx context.Context, config Config, opaqueReq models.VerificarComunicacionBoletoAereo) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)
}
