package models

import (
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/sincronizacion"
)

// --- Interfaces opacas para las solicitudes de Sincronización ---

type SincronizarActividades struct {
	RequestWrapper[sincronizacion.SincronizarActividades]
}
type SincronizarListaActividadesDocumentoSector struct {
	RequestWrapper[sincronizacion.SincronizarListaActividadesDocumentoSector]
}
type SincronizarListaLeyendasFactura struct {
	RequestWrapper[sincronizacion.SincronizarListaLeyendasFactura]
}
type SincronizarListaMensajesServicios struct {
	RequestWrapper[sincronizacion.SincronizarListaMensajesServicios]
}
type SincronizarListaProductosServicios struct {
	RequestWrapper[sincronizacion.SincronizarListaProductosServicios]
}
type SincronizarParametricaEventosSignificativos struct {
	RequestWrapper[sincronizacion.SincronizarParametricaEventosSignificativos]
}
type SincronizarParametricaMotivoAnulacion struct {
	RequestWrapper[sincronizacion.SincronizarParametricaMotivoAnulacion]
}
type SincronizarParametricaPaisOrigen struct {
	RequestWrapper[sincronizacion.SincronizarParametricaPaisOrigen]
}
type SincronizarParametricaTipoDocumentoIdentidad struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad]
}
type SincronizarParametricaTipoDocumentoSector struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoSector]
}
type SincronizarParametricaTipoEmision struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTipoEmision]
}
type SincronizarParametricaTipoHabitacion struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTipoHabitacion]
}
type SincronizarParametricaTipoMetodoPago struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTipoMetodoPago]
}
type SincronizarParametricaTipoMoneda struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTipoMoneda]
}
type SincronizarParametricaTipoPuntoVenta struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTipoPuntoVenta]
}
type SincronizarParametricaTiposFactura struct {
	RequestWrapper[sincronizacion.SincronizarParametricaTiposFactura]
}
type SincronizarParametricaUnidadMedida struct {
	RequestWrapper[sincronizacion.SincronizarParametricaUnidadMedida]
}
type SincronizarFechaHora struct {
	RequestWrapper[sincronizacion.SincronizarFechaHora]
}
type VerificarComunicacionSincronizacion struct {
	RequestWrapper[sincronizacion.VerificarComunicacion]
}

// --- Constructors a nivel de paquete ---

func NewVerificarComunicacionSincronizacionBuilder() *verificarComunicacionSincronizacionBuilder {
	return &verificarComunicacionSincronizacionBuilder{
		request: &sincronizacion.VerificarComunicacion{},
	}
}

func NewSincronizarActividadesBuilder() SincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades] {
	req := &sincronizacion.SincronizarActividades{}
	return &sincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarActividades]) SincronizarActividades {
			return SincronizarActividades{rw}
		},
	}
}

func NewSincronizarListaActividadesDocumentoSectorBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector] {
	req := &sincronizacion.SincronizarListaActividadesDocumentoSector{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaActividadesDocumentoSector]) SincronizarListaActividadesDocumentoSector {
			return SincronizarListaActividadesDocumentoSector{rw}
		},
	}
}

func NewSincronizarListaLeyendasFacturaBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura] {
	req := &sincronizacion.SincronizarListaLeyendasFactura{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaLeyendasFactura]) SincronizarListaLeyendasFactura {
			return SincronizarListaLeyendasFactura{rw}
		},
	}
}

func NewSincronizarListaMensajesServiciosBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios] {
	req := &sincronizacion.SincronizarListaMensajesServicios{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaMensajesServicios]) SincronizarListaMensajesServicios {
			return SincronizarListaMensajesServicios{rw}
		},
	}
}

func NewSincronizarFechaHoraBuilder() SincronizacionBuilder[sincronizacion.SincronizarFechaHora, SincronizarFechaHora] {
	req := &sincronizacion.SincronizarFechaHora{}
	return &sincronizacionBuilder[sincronizacion.SincronizarFechaHora, SincronizarFechaHora]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarFechaHora]) SincronizarFechaHora {
			return SincronizarFechaHora{rw}
		},
	}
}

func NewSincronizarListaProductosServiciosBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios] {
	req := &sincronizacion.SincronizarListaProductosServicios{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaProductosServicios]) SincronizarListaProductosServicios {
			return SincronizarListaProductosServicios{rw}
		},
	}
}

func NewSincronizarParametricaEventosSignificativosBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos] {
	req := &sincronizacion.SincronizarParametricaEventosSignificativos{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaEventosSignificativos]) SincronizarParametricaEventosSignificativos {
			return SincronizarParametricaEventosSignificativos{rw}
		},
	}
}

func NewSincronizarParametricaMotivoAnulacionBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion] {
	req := &sincronizacion.SincronizarParametricaMotivoAnulacion{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaMotivoAnulacion]) SincronizarParametricaMotivoAnulacion {
			return SincronizarParametricaMotivoAnulacion{rw}
		},
	}
}

func NewSincronizarParametricaPaisOrigenBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen] {
	req := &sincronizacion.SincronizarParametricaPaisOrigen{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaPaisOrigen]) SincronizarParametricaPaisOrigen {
			return SincronizarParametricaPaisOrigen{rw}
		},
	}
}

func NewSincronizarParametricaTipoDocumentoIdentidadBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoIdentidad{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad]) SincronizarParametricaTipoDocumentoIdentidad {
			return SincronizarParametricaTipoDocumentoIdentidad{rw}
		},
	}
}

func NewSincronizarParametricaTipoDocumentoSectorBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoSector{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoSector]) SincronizarParametricaTipoDocumentoSector {
			return SincronizarParametricaTipoDocumentoSector{rw}
		},
	}
}

func NewSincronizarParametricaTipoEmisionBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision] {
	req := &sincronizacion.SincronizarParametricaTipoEmision{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoEmision]) SincronizarParametricaTipoEmision {
			return SincronizarParametricaTipoEmision{rw}
		},
	}
}

func NewSincronizarParametricaTipoHabitacionBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion] {
	req := &sincronizacion.SincronizarParametricaTipoHabitacion{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoHabitacion]) SincronizarParametricaTipoHabitacion {
			return SincronizarParametricaTipoHabitacion{rw}
		},
	}
}

func NewSincronizarParametricaTipoMetodoPagoBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago] {
	req := &sincronizacion.SincronizarParametricaTipoMetodoPago{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoMetodoPago]) SincronizarParametricaTipoMetodoPago {
			return SincronizarParametricaTipoMetodoPago{rw}
		},
	}
}

func NewSincronizarParametricaTipoMonedaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda] {
	req := &sincronizacion.SincronizarParametricaTipoMoneda{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoMoneda]) SincronizarParametricaTipoMoneda {
			return SincronizarParametricaTipoMoneda{rw}
		},
	}
}

func NewSincronizarParametricaTipoPuntoVentaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta] {
	req := &sincronizacion.SincronizarParametricaTipoPuntoVenta{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoPuntoVenta]) SincronizarParametricaTipoPuntoVenta {
			return SincronizarParametricaTipoPuntoVenta{rw}
		},
	}
}

func NewSincronizarParametricaTiposFacturaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura] {
	req := &sincronizacion.SincronizarParametricaTiposFactura{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTiposFactura]) SincronizarParametricaTiposFactura {
			return SincronizarParametricaTiposFactura{rw}
		},
	}
}

func NewSincronizarParametricaUnidadMedidaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida, SincronizarParametricaUnidadMedida] {
	req := &sincronizacion.SincronizarParametricaUnidadMedida{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida, SincronizarParametricaUnidadMedida]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaUnidadMedida]) SincronizarParametricaUnidadMedida {
			return SincronizarParametricaUnidadMedida{rw}
		},
	}
}

// --- Implementations ---

type verificarComunicacionSincronizacionBuilder struct {
	request *sincronizacion.VerificarComunicacion
}

func (b *verificarComunicacionSincronizacionBuilder) Build() VerificarComunicacionSincronizacion {
	return VerificarComunicacionSincronizacion{RequestWrapper[sincronizacion.VerificarComunicacion]{request: b.request}}
}

// SincronizacionBuilder define la interfaz para configurar las diversas solicitudes de sincronización del SIAT.
type SincronizacionBuilder[T any, R any] interface {
	// WithCodigoPuntoVenta establece el código del punto de venta.
	WithCodigoPuntoVenta(int) SincronizacionBuilder[T, R]
	// WithCodigoSucursal establece el código de la sucursal.
	WithCodigoSucursal(int) SincronizacionBuilder[T, R]
	// WithCuis establece el Código Único de Inicio de Sistemas.
	WithCuis(string) SincronizacionBuilder[T, R]
	// Build construye y retorna la solicitud configurada.
	Build() R
}

// sincronizacionBuilder es un generador genérico para configurar solicitudes de sincronización.
type sincronizacionBuilder[T any, R any] struct {
	request *T
	sol     *sincronizacion.SolicitudSincronizacion
	wrap    func(RequestWrapper[T]) R
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *sincronizacionBuilder[T, R]) WithCodigoPuntoVenta(codigoPuntoVenta int) SincronizacionBuilder[T, R] {
	b.sol.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *sincronizacionBuilder[T, R]) WithCodigoSucursal(codigoSucursal int) SincronizacionBuilder[T, R] {
	b.sol.CodigoSucursal = codigoSucursal
	return b
}

// WithCuis establece el CUIS.
func (b *sincronizacionBuilder[T, R]) WithCuis(cuis string) SincronizacionBuilder[T, R] {
	b.sol.Cuis = cuis
	return b
}

// Build entrega el objeto de solicitud configurado.
func (b *sincronizacionBuilder[T, R]) Build() R {
	return b.wrap(RequestWrapper[T]{request: b.request})
}
