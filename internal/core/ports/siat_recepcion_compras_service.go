package ports

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatRecepcionComprasService define el puerto para el servicio de Recepción de Compras del SIAT.
// Permite la gestión de facturas de compras, incluyendo anulación, confirmación y consulta.
type SiatRecepcionComprasService interface {
	// AnulacionCompra solicita la anulación de una compra previamente registrada.
	AnulacionCompra(ctx context.Context, config Config, opaqueReq models.AnulacionCompra) (*soap.EnvelopeResponse[facturacion.AnulacionCompraResponse], error)

	// ConfirmacionCompras confirma las compras enviadas.
	ConfirmacionCompras(ctx context.Context, config Config, opaqueReq models.ConfirmacionCompras) (*soap.EnvelopeResponse[facturacion.ConfirmacionComprasResponse], error)

	// ConsultaCompras permite consultar las compras enviadas en un periodo determinado.
	ConsultaCompras(ctx context.Context, config Config, opaqueReq models.ConsultaCompras) (*soap.EnvelopeResponse[facturacion.ConsultaComprasResponse], error)

	// RecepcionPaqueteCompras envía un paquete de facturas de compras para su procesamiento.
	RecepcionPaqueteCompras(ctx context.Context, config Config, opaqueReq models.RecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteComprasResponse], error)

	// ValidacionRecepcionPaqueteCompras valida el estado de un envío de paquete de compras.
	ValidacionRecepcionPaqueteCompras(ctx context.Context, config Config, opaqueReq models.ValidacionRecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteComprasResponse], error)

	// VerificarComunicacion realiza una prueba de conectividad con este servicio específico.
	VerificarComunicacion(ctx context.Context, config Config, opaqueReq models.VerificarComunicacionRecepcionCompras) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)
}
