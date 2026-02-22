package port

import (
	"context"
	"go-siat/internal/core/domain/datatype/soap"
	"go-siat/internal/core/domain/facturacion"
	"go-siat/internal/core/domain/facturacion/operaciones"
)

type SiatOperacionesPort interface {
	VerificarComunicacion(ctx context.Context, config facturacion.Config) (*soap.EnvelopeResponse[operaciones.VerificarComunicacionResponse], error)
	RegistroPuntoVenta(ctx context.Context, config facturacion.Config, req *operaciones.RegistroPuntoVenta) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaResponse], error)
	ConsultaPuntoVenta(ctx context.Context, config facturacion.Config, req *operaciones.ConsultaPuntoVenta) (*soap.EnvelopeResponse[operaciones.ConsultaPuntoVentaResponse], error)
	CierrePuntoVenta(ctx context.Context, config facturacion.Config, req *operaciones.CierrePuntoVenta) (*soap.EnvelopeResponse[operaciones.CierrePuntoVentaResponse], error)
	RegistroPuntoVentaComisionista(ctx context.Context, config facturacion.Config, req *operaciones.RegistroPuntoVentaComisionista) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaComisionistaResponse], error)
	RegistroEventosSignificativos(ctx context.Context, config facturacion.Config, req *operaciones.RegistroEventoSignificativo) (*soap.EnvelopeResponse[operaciones.RegistroEventoSignificativoResponse], error)
	ConsultaEventosSignificativos(ctx context.Context, config facturacion.Config, req *operaciones.ConsultaEventoSignificativo) (*soap.EnvelopeResponse[operaciones.ConsultaEventoSignificativoResponse], error)
	CierreOperacionesSistema(ctx context.Context, config facturacion.Config, req *operaciones.CierreOperacionesSistema) (*soap.EnvelopeResponse[operaciones.CierreOperacionesSistemaResponse], error)
}
