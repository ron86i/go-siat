package ports

import (
	"context"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/compra_venta"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

// FacturacionService define el contrato común para todas las modalidades de facturación del SIAT.
type FacturacionService interface {
	AnulacionFactura(ctx context.Context, req models.AnulacionFactura) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error)
	RecepcionFactura(ctx context.Context, req models.RecepcionFactura) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error)
	ReversionAnulacionFactura(ctx context.Context, req models.ReversionAnulacionFactura) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error)
	RecepcionPaqueteFactura(ctx context.Context, req models.RecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error)
	ValidacionRecepcionPaqueteFactura(ctx context.Context, req models.ValidacionRecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error)
	VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionFacturacion) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error)
	RecepcionMasivaFactura(ctx context.Context, req models.RecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error)
	ValidacionRecepcionMasivaFactura(ctx context.Context, req models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error)
	VerificacionEstadoFactura(ctx context.Context, req models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error)
}

// SiatCompraVentaService extiende FacturacionService con métodos exclusivos para Compra-Venta (anexos).
type SiatCompraVentaService interface {
	FacturacionService
	RecepcionAnexos(ctx context.Context, req models.RecepcionAnexosCompraVenta) (*soap.EnvelopeResponse[compra_venta.RecepcionAnexosResponse], error)
}

// SiatSuministroEnergiaService extiende FacturacionService con métodos exclusivos para Suministro de Energía.
type SiatSuministroEnergiaService interface {
	FacturacionService
	RecepcionAnexosSuministroEnergia(ctx context.Context, req models.RecepcionAnexosSuministroEnergia) (*soap.EnvelopeResponse[facturacion.RecepcionAnexosSuministroEnergiaResponse], error)
}
