package models

import (
	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
)

// --- Wrappers ---

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

// --- Namespace ---

type sincronizacionNamespace struct{}

func Sincronizacion() sincronizacionNamespace {
	return sincronizacionNamespace{}
}

// --- Constructors ---

func (sincronizacionNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionSincronizacionBuilder {
	return &verificarComunicacionSincronizacionBuilder{
		request: &sincronizacion.VerificarComunicacion{},
	}
}

// NewSincronizarActividadesBuilder inicia la construcción de una solicitud para sincronizar actividades económicas.
func (sincronizacionNamespace) NewSincronizarActividadesBuilder() SincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades] {
	req := &sincronizacion.SincronizarActividades{}
	return &sincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarActividades]) SincronizarActividades {
			return SincronizarActividades{rw}
		},
	}
}

// NewSincronizarListaActividadesDocumentoSectorBuilder inicia la construcción para la relación actividad-sector.
func (sincronizacionNamespace) NewSincronizarListaActividadesDocumentoSectorBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector] {
	req := &sincronizacion.SincronizarListaActividadesDocumentoSector{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaActividadesDocumentoSector]) SincronizarListaActividadesDocumentoSector {
			return SincronizarListaActividadesDocumentoSector{rw}
		},
	}
}

// NewSincronizarListaLeyendasFacturaBuilder inicia la construcción para obtener leyendas de invoices.
func (sincronizacionNamespace) NewSincronizarListaLeyendasFacturaBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura] {
	req := &sincronizacion.SincronizarListaLeyendasFactura{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaLeyendasFactura]) SincronizarListaLeyendasFactura {
			return SincronizarListaLeyendasFactura{rw}
		},
	}
}

// NewSincronizarListaMensajesServiciosBuilder inicia la construcción para obtener mensajes del servicio.
func (sincronizacionNamespace) NewSincronizarListaMensajesServiciosBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios] {
	req := &sincronizacion.SincronizarListaMensajesServicios{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaMensajesServicios]) SincronizarListaMensajesServicios {
			return SincronizarListaMensajesServicios{rw}
		},
	}
}

// NewSincronizarFechaHoraBuilder inicia la construcción para obtener la fecha y hora oficial del servidor del SIAT.
func (sincronizacionNamespace) NewSincronizarFechaHoraBuilder() SincronizacionBuilder[sincronizacion.SincronizarFechaHora, SincronizarFechaHora] {
	req := &sincronizacion.SincronizarFechaHora{}
	return &sincronizacionBuilder[sincronizacion.SincronizarFechaHora, SincronizarFechaHora]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarFechaHora]) SincronizarFechaHora {
			return SincronizarFechaHora{rw}
		},
	}
}

// NewSincronizarListaProductosServiciosBuilder inicia la construcción para sincronizar productos y servicios.
func (sincronizacionNamespace) NewSincronizarListaProductosServiciosBuilder() SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios] {
	req := &sincronizacion.SincronizarListaProductosServicios{}
	return &sincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarListaProductosServicios]) SincronizarListaProductosServicios {
			return SincronizarListaProductosServicios{rw}
		},
	}
}

// NewSincronizarParametricaEventosSignificativosBuilder inicia la construcción para el catálogo de eventos significativos.
func (sincronizacionNamespace) NewSincronizarParametricaEventosSignificativosBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos] {
	req := &sincronizacion.SincronizarParametricaEventosSignificativos{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaEventosSignificativos]) SincronizarParametricaEventosSignificativos {
			return SincronizarParametricaEventosSignificativos{rw}
		},
	}
}

// NewSincronizarParametricaMotivoAnulacionBuilder inicia la construcción para motivos de anulación.
func (sincronizacionNamespace) NewSincronizarParametricaMotivoAnulacionBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion] {
	req := &sincronizacion.SincronizarParametricaMotivoAnulacion{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaMotivoAnulacion]) SincronizarParametricaMotivoAnulacion {
			return SincronizarParametricaMotivoAnulacion{rw}
		},
	}
}

// NewSincronizarParametricaPaisOrigenBuilder inicia la construcción para el catálogo de países.
func (sincronizacionNamespace) NewSincronizarParametricaPaisOrigenBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen] {
	req := &sincronizacion.SincronizarParametricaPaisOrigen{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaPaisOrigen]) SincronizarParametricaPaisOrigen {
			return SincronizarParametricaPaisOrigen{rw}
		},
	}
}

// NewSincronizarParametricaTipoDocumentoIdentidadBuilder inicia la construcción para tipos de documento de identidad.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoIdentidadBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoIdentidad{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad]) SincronizarParametricaTipoDocumentoIdentidad {
			return SincronizarParametricaTipoDocumentoIdentidad{rw}
		},
	}
}

// NewSincronizarParametricaTipoDocumentoSectorBuilder inicia la construcción para tipos de documento sector.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoSectorBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoSector{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoSector]) SincronizarParametricaTipoDocumentoSector {
			return SincronizarParametricaTipoDocumentoSector{rw}
		},
	}
}

// NewSincronizarParametricaTipoEmisionBuilder inicia la construcción para tipos de emisión.
func (sincronizacionNamespace) NewSincronizarParametricaTipoEmisionBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision] {
	req := &sincronizacion.SincronizarParametricaTipoEmision{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoEmision]) SincronizarParametricaTipoEmision {
			return SincronizarParametricaTipoEmision{rw}
		},
	}
}

// NewSincronizarParametricaTipoHabitacionBuilder inicia la construcción para tipos de habitación.
func (sincronizacionNamespace) NewSincronizarParametricaTipoHabitacionBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion] {
	req := &sincronizacion.SincronizarParametricaTipoHabitacion{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoHabitacion]) SincronizarParametricaTipoHabitacion {
			return SincronizarParametricaTipoHabitacion{rw}
		},
	}
}

// NewSincronizarParametricaTipoMetodoPagoBuilder inicia la construcción para métodos de pago.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMetodoPagoBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago] {
	req := &sincronizacion.SincronizarParametricaTipoMetodoPago{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoMetodoPago]) SincronizarParametricaTipoMetodoPago {
			return SincronizarParametricaTipoMetodoPago{rw}
		},
	}
}

// NewSincronizarParametricaTipoMonedaBuilder inicia la construcción para tipos de moneda.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMonedaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda] {
	req := &sincronizacion.SincronizarParametricaTipoMoneda{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoMoneda]) SincronizarParametricaTipoMoneda {
			return SincronizarParametricaTipoMoneda{rw}
		},
	}
}

// NewSincronizarParametricaTipoPuntoVentaBuilder inicia la construcción para tipos de punto de venta.
func (sincronizacionNamespace) NewSincronizarParametricaTipoPuntoVentaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta] {
	req := &sincronizacion.SincronizarParametricaTipoPuntoVenta{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTipoPuntoVenta]) SincronizarParametricaTipoPuntoVenta {
			return SincronizarParametricaTipoPuntoVenta{rw}
		},
	}
}

// NewSincronizarParametricaTiposFacturaBuilder inicia la construcción para tipos de factura.
func (sincronizacionNamespace) NewSincronizarParametricaTiposFacturaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura] {
	req := &sincronizacion.SincronizarParametricaTiposFactura{}
	return &sincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw RequestWrapper[sincronizacion.SincronizarParametricaTiposFactura]) SincronizarParametricaTiposFactura {
			return SincronizarParametricaTiposFactura{rw}
		},
	}
}

// NewSincronizarParametricaUnidadMedidaBuilder inicia la construcción para unidades de medida.
func (sincronizacionNamespace) NewSincronizarParametricaUnidadMedidaBuilder() SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida, SincronizarParametricaUnidadMedida] {
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

// RegistroEventoSignificativoBuilder ayuda a configurar el registro de un evento significativo.
type verificarComunicacionSincronizacionBuilder struct {
	request *sincronizacion.VerificarComunicacion
}

func (b *verificarComunicacionSincronizacionBuilder) Build() VerificarComunicacionSincronizacion {
	return VerificarComunicacionSincronizacion{RequestWrapper[sincronizacion.VerificarComunicacion]{request: b.request}}
}

// SincronizacionBuilder define la interfaz para configurar solicitudes de sincronización.
type SincronizacionBuilder[T any, R any] interface {
	WithCodigoAmbiente(int) SincronizacionBuilder[T, R]
	WithCodigoPuntoVenta(int) SincronizacionBuilder[T, R]
	WithCodigoSistema(string) SincronizacionBuilder[T, R]
	WithCodigoSucursal(int) SincronizacionBuilder[T, R]
	WithCuis(string) SincronizacionBuilder[T, R]
	WithNit(int64) SincronizacionBuilder[T, R]
	Build() R
}

// sincronizacionBuilder es un generador genérico para configurar solicitudes de sincronización.
type sincronizacionBuilder[T any, R any] struct {
	request *T
	sol     *sincronizacion.SolicitudSincronizacion
	wrap    func(RequestWrapper[T]) R
}

func (b *sincronizacionBuilder[T, R]) WithCodigoAmbiente(codigoAmbiente int) SincronizacionBuilder[T, R] {
	b.sol.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *sincronizacionBuilder[T, R]) WithCodigoPuntoVenta(codigoPuntoVenta int) SincronizacionBuilder[T, R] {
	b.sol.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *sincronizacionBuilder[T, R]) WithCodigoSistema(codigoSistema string) SincronizacionBuilder[T, R] {
	b.sol.CodigoSistema = codigoSistema
	return b
}

func (b *sincronizacionBuilder[T, R]) WithCodigoSucursal(codigoSucursal int) SincronizacionBuilder[T, R] {
	b.sol.CodigoSucursal = codigoSucursal
	return b
}

func (b *sincronizacionBuilder[T, R]) WithCuis(cuis string) SincronizacionBuilder[T, R] {
	b.sol.Cuis = cuis
	return b
}

func (b *sincronizacionBuilder[T, R]) WithNit(nit int64) SincronizacionBuilder[T, R] {
	b.sol.NIT = nit
	return b
}

// Build entrega el objeto de solicitud configurado.
func (b *sincronizacionBuilder[T, R]) Build() R {
	return b.wrap(RequestWrapper[T]{request: b.request})
}
