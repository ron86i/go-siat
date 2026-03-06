package port

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/operaciones"
	"github.com/ron86i/go-siat/pkg/config"
)

type SiatOperacionesPort interface {
	VerificarComunicacion(ctx context.Context, config config.Config) (*soap.EnvelopeResponse[operaciones.VerificarComunicacionResponse], error)
	RegistroPuntoVenta(ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaResponse], error)
	ConsultaPuntoVenta(ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[operaciones.ConsultaPuntoVentaResponse], error)
	CierrePuntoVenta(ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[operaciones.CierrePuntoVentaResponse], error)
	RegistroPuntoVentaComisionista(ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaComisionistaResponse], error)
	RegistroEventosSignificativos(ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[operaciones.RegistroEventoSignificativoResponse], error)
	ConsultaEventosSignificativos(ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[operaciones.ConsultaEventoSignificativoResponse], error)
	CierreOperacionesSistema(ctx context.Context, config config.Config, req any) (*soap.EnvelopeResponse[operaciones.CierreOperacionesSistemaResponse], error)
}
