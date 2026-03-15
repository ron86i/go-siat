package port

import (
	"context"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/operaciones"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatOperacionesPort define el contrato para el Servicio Web de Operaciones del SIAT.
// Este puerto facilita la gestión de puntos de venta, el registro de eventos significativos
// y el control de cierres de operaciones y sistemas.
type SiatOperacionesPort interface {
	// VerificarComunicacion realiza una prueba de conectividad con el servicio de operaciones del SIAT.
	VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionOperaciones) (*soap.EnvelopeResponse[operaciones.VerificarComunicacionResponse], error)

	// RegistroPuntoVenta gestiona la creación y habilitación de un nuevo punto de venta para el contribuyente.
	RegistroPuntoVenta(ctx context.Context, config config.Config, req models.RegistroPuntoVenta) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaResponse], error)

	// ConsultaPuntoVenta recupera la información detallada de los puntos de venta asociados a un contribuyente.
	ConsultaPuntoVenta(ctx context.Context, config config.Config, req models.ConsultaPuntoVenta) (*soap.EnvelopeResponse[operaciones.ConsultaPuntoVentaResponse], error)

	// CierrePuntoVenta tramita la deshabilitación o cierre definitivo de un punto de venta registrado.
	CierrePuntoVenta(ctx context.Context, config config.Config, req models.CierrePuntoVenta) (*soap.EnvelopeResponse[operaciones.CierrePuntoVentaResponse], error)

	// RegistroPuntoVentaComisionista registra un punto de venta gestionado por un tercero comisionista.
	RegistroPuntoVentaComisionista(ctx context.Context, config config.Config, req models.RegistroPuntoVentaComisionista) (*soap.EnvelopeResponse[operaciones.RegistroPuntoVentaComisionistaResponse], error)

	// RegistroEventosSignificativos comunica al SIAT la ocurrencia de sucesos que impiden la facturación en línea.
	RegistroEventosSignificativos(ctx context.Context, config config.Config, req models.RegistroEventoSignificativo) (*soap.EnvelopeResponse[operaciones.RegistroEventoSignificativoResponse], error)

	// ConsultaEventosSignificativos permite obtener el listado de eventos de contingencia reportados al SIAT.
	ConsultaEventosSignificativos(ctx context.Context, config config.Config, req models.ConsultaEventoSignificativo) (*soap.EnvelopeResponse[operaciones.ConsultaEventoSignificativoResponse], error)

	// CierreOperacionesSistema finaliza las actividades de facturación del sistema para un periodo o evento dado.
	CierreOperacionesSistema(ctx context.Context, config config.Config, req models.CierreOperacionesSistema) (*soap.EnvelopeResponse[operaciones.CierreOperacionesSistemaResponse], error)
}
