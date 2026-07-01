package ports

import (
	"context"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/operaciones"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

// SiatOperacionesService define el contrato parssa el Servicio Web de Operaciones del SIAT.
type SiatOperacionesService interface {
	VerificarComunicacion(ctx context.Context, opaqueReq models.VerificarComunicacionOperaciones) (*soap.EnvelopeResponse[operaciones.VerificarComunicacionResponse], error)
	RegistroPuntoVenta(ctx context.Context, req models.RegistroPuntoVenta) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaResponse], error)
	ConsultaPuntoVenta(ctx context.Context, req models.ConsultaPuntoVenta) (*soap.EnvelopeResponse[operaciones.ConsultaPuntoVentaResponse], error)
	CierrePuntoVenta(ctx context.Context, req models.CierrePuntoVenta) (*soap.EnvelopeResponse[operaciones.CierrePuntoVentaResponse], error)
	RegistroPuntoVentaComisionista(ctx context.Context, req models.RegistroPuntoVentaComisionista) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaComisionistaResponse], error)
	RegistroEventosSignificativos(ctx context.Context, req models.RegistroEventoSignificativo) (*soap.EnvelopeResponse[operaciones.RegistroEventoSignificativoResponse], error)
	ConsultaEventosSignificativos(ctx context.Context, req models.ConsultaEventoSignificativo) (*soap.EnvelopeResponse[operaciones.ConsultaEventoSignificativoResponse], error)
	CierreOperacionesSistema(ctx context.Context, req models.CierreOperacionesSistema) (*soap.EnvelopeResponse[operaciones.CierreOperacionesSistemaResponse], error)
}
