package ports

import (
	"context"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

// SiatRecepcionComprasService define el puerto para el servicio de Recepción de Compras del SIAT.
type SiatRecepcionComprasService interface {
	AnulacionCompra(ctx context.Context, req models.AnulacionCompra) (*soap.EnvelopeResponse[facturacion.AnulacionCompraResponse], error)
	ConfirmacionCompras(ctx context.Context, req models.ConfirmacionCompras) (*soap.EnvelopeResponse[facturacion.ConfirmacionComprasResponse], error)
	ConsultaCompras(ctx context.Context, req models.ConsultaCompras) (*soap.EnvelopeResponse[facturacion.ConsultaComprasResponse], error)
	RecepcionPaqueteCompras(ctx context.Context, req models.RecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteComprasResponse], error)
	ValidacionRecepcionPaqueteCompras(ctx context.Context, req models.ValidacionRecepcionPaqueteCompras) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteComprasResponse], error)
	VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionRecepcionCompras) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)
}
