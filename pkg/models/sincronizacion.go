package models

import (
	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
)

type sincronizacionNamespace struct{}

func Sincronizacion() sincronizacionNamespace {
	return sincronizacionNamespace{}
}

// --- Interfaces opacas para las solicitudes de Sincronización ---

type SincronizarActividades struct {
	requestWrapper[sincronizacion.SincronizarActividades]
}
type SincronizarListaActividadesDocumentoSector struct {
	requestWrapper[sincronizacion.SincronizarListaActividadesDocumentoSector]
}
type SincronizarListaLeyendasFactura struct {
	requestWrapper[sincronizacion.SincronizarListaLeyendasFactura]
}
type SincronizarListaMensajesServicios struct {
	requestWrapper[sincronizacion.SincronizarListaMensajesServicios]
}
type SincronizarListaProductosServicios struct {
	requestWrapper[sincronizacion.SincronizarListaProductosServicios]
}
type SincronizarParametricaEventosSignificativos struct {
	requestWrapper[sincronizacion.SincronizarParametricaEventosSignificativos]
}
type SincronizarParametricaMotivoAnulacion struct {
	requestWrapper[sincronizacion.SincronizarParametricaMotivoAnulacion]
}
type SincronizarParametricaPaisOrigen struct {
	requestWrapper[sincronizacion.SincronizarParametricaPaisOrigen]
}
type SincronizarParametricaTipoDocumentoIdentidad struct {
	requestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad]
}
type SincronizarParametricaTipoDocumentoSector struct {
	requestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoSector]
}
type SincronizarParametricaTipoEmision struct {
	requestWrapper[sincronizacion.SincronizarParametricaTipoEmision]
}
type SincronizarParametricaTipoHabitacion struct {
	requestWrapper[sincronizacion.SincronizarParametricaTipoHabitacion]
}
type SincronizarParametricaTipoMetodoPago struct {
	requestWrapper[sincronizacion.SincronizarParametricaTipoMetodoPago]
}
type SincronizarParametricaTipoMoneda struct {
	requestWrapper[sincronizacion.SincronizarParametricaTipoMoneda]
}
type SincronizarParametricaTipoPuntoVenta struct {
	requestWrapper[sincronizacion.SincronizarParametricaTipoPuntoVenta]
}
type SincronizarParametricaTiposFactura struct {
	requestWrapper[sincronizacion.SincronizarParametricaTiposFactura]
}
type SincronizarParametricaUnidadMedida struct {
	requestWrapper[sincronizacion.SincronizarParametricaUnidadMedida]
}
type VerificarComunicacionSincronizacion struct {
	requestWrapper[sincronizacion.VerificarComunicacion]
}

func (sincronizacionNamespace) NewVerificarComunicacionBuilder() *VerificarComunicacionSincronizacionBuilder {
	return &VerificarComunicacionSincronizacionBuilder{
		request: &sincronizacion.VerificarComunicacion{},
	}
}

// RegistroEventoSignificativoBuilder ayuda a configurar el registro de un evento significativo.
type VerificarComunicacionSincronizacionBuilder struct {
	request *sincronizacion.VerificarComunicacion
}

func (b *VerificarComunicacionSincronizacionBuilder) Build() VerificarComunicacionSincronizacion {
	return VerificarComunicacionSincronizacion{requestWrapper[sincronizacion.VerificarComunicacion]{request: b.request}}
}

// SincronizacionBuilder es un generador genérico para configurar solicitudes de sincronización.
type SincronizacionBuilder[T any, R any] struct {
	request *T
	sol     *sincronizacion.SolicitudSincronizacion
	wrap    func(requestWrapper[T]) R
}

func (b *SincronizacionBuilder[T, R]) WithCodigoAmbiente(codigoAmbiente int) *SincronizacionBuilder[T, R] {
	b.sol.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *SincronizacionBuilder[T, R]) WithCodigoPuntoVenta(codigoPuntoVenta int) *SincronizacionBuilder[T, R] {
	b.sol.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *SincronizacionBuilder[T, R]) WithCodigoSistema(codigoSistema string) *SincronizacionBuilder[T, R] {
	b.sol.CodigoSistema = codigoSistema
	return b
}

func (b *SincronizacionBuilder[T, R]) WithCodigoSucursal(codigoSucursal int) *SincronizacionBuilder[T, R] {
	b.sol.CodigoSucursal = codigoSucursal
	return b
}

func (b *SincronizacionBuilder[T, R]) WithCuis(cuis string) *SincronizacionBuilder[T, R] {
	b.sol.Cuis = cuis
	return b
}

func (b *SincronizacionBuilder[T, R]) WithNit(nit int64) *SincronizacionBuilder[T, R] {
	b.sol.NIT = nit
	return b
}

// Build entrega el objeto de solicitud configurado.
func (b *SincronizacionBuilder[T, R]) Build() R {
	return b.wrap(requestWrapper[T]{request: b.request})
}

// NewSincronizarActividadesRequest inicia la construcción de una solicitud para sincronizar actividades económicas.
func (sincronizacionNamespace) NewSincronizarActividadesBuilder() *SincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades] {
	req := &sincronizacion.SincronizarActividades{}
	return &SincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap:    func(rw requestWrapper[sincronizacion.SincronizarActividades]) SincronizarActividades {
			return SincronizarActividades{rw}
		},
	}
}

// NewSincronizarListaActividadesDocumentoSectorRequest inicia la construcción para la relación actividad-sector.
func (sincronizacionNamespace) NewSincronizarListaActividadesDocumentoSectorBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector] {
	req := &sincronizacion.SincronizarListaActividadesDocumentoSector{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap:    func(rw requestWrapper[sincronizacion.SincronizarListaActividadesDocumentoSector]) SincronizarListaActividadesDocumentoSector {
			return SincronizarListaActividadesDocumentoSector{rw}
		},
	}
}

// NewSincronizarListaLeyendasFacturaRequest inicia la construcción para obtener leyendas de facturas.
func (sincronizacionNamespace) NewSincronizarListaLeyendasFacturaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura] {
	req := &sincronizacion.SincronizarListaLeyendasFactura{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarListaLeyendasFactura]) SincronizarListaLeyendasFactura {
			return SincronizarListaLeyendasFactura{rw}
		},
	}
}

// NewSincronizarListaMensajesServiciosRequest inicia la construcción para obtener mensajes del servicio.
func (sincronizacionNamespace) NewSincronizarListaMensajesServiciosBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios] {
	req := &sincronizacion.SincronizarListaMensajesServicios{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarListaMensajesServicios]) SincronizarListaMensajesServicios {
			return SincronizarListaMensajesServicios{rw}
		},
	}
}

// NewSincronizarListaProductosServiciosRequest inicia la construcción para sincronizar productos y servicios.
func (sincronizacionNamespace) NewSincronizarListaProductosServiciosBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios] {
	req := &sincronizacion.SincronizarListaProductosServicios{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarListaProductosServicios]) SincronizarListaProductosServicios {
			return SincronizarListaProductosServicios{rw}
		},
	}
}

// Paramétricas

// NewSincronizarParametricaEventosSignificativosRequest inicia la construcción para el catálogo de eventos significativos.
func (sincronizacionNamespace) NewSincronizarParametricaEventosSignificativosBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos] {
	req := &sincronizacion.SincronizarParametricaEventosSignificativos{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaEventosSignificativos]) SincronizarParametricaEventosSignificativos {
			return SincronizarParametricaEventosSignificativos{rw}
		},
	}
}

// NewSincronizarParametricaMotivoAnulacionRequest inicia la construcción para motivos de anulación.
func (sincronizacionNamespace) NewSincronizarParametricaMotivoAnulacionBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion] {
	req := &sincronizacion.SincronizarParametricaMotivoAnulacion{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaMotivoAnulacion]) SincronizarParametricaMotivoAnulacion {
			return SincronizarParametricaMotivoAnulacion{rw}
		},
	}
}

// NewSincronizarParametricaPaisOrigenRequest inicia la construcción para el catálogo de países.
func (sincronizacionNamespace) NewSincronizarParametricaPaisOrigenBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen] {
	req := &sincronizacion.SincronizarParametricaPaisOrigen{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaPaisOrigen]) SincronizarParametricaPaisOrigen {
			return SincronizarParametricaPaisOrigen{rw}
		},
	}
}

// NewSincronizarParametricaTipoDocumentoIdentidadRequest inicia la construcción para tipos de documento de identidad.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoIdentidadBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoIdentidad{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad]) SincronizarParametricaTipoDocumentoIdentidad {
			return SincronizarParametricaTipoDocumentoIdentidad{rw}
		},
	}
}

// NewSincronizarParametricaTipoDocumentoSectorRequest inicia la construcción para tipos de documento sector.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoSectorBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoSector{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTipoDocumentoSector]) SincronizarParametricaTipoDocumentoSector {
			return SincronizarParametricaTipoDocumentoSector{rw}
		},
	}
}

// NewSincronizarParametricaTipoEmisionRequest inicia la construcción para tipos de emisión.
func (sincronizacionNamespace) NewSincronizarParametricaTipoEmisionBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision] {
	req := &sincronizacion.SincronizarParametricaTipoEmision{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTipoEmision]) SincronizarParametricaTipoEmision {
			return SincronizarParametricaTipoEmision{rw}
		},
	}
}

// NewSincronizarParametricaTipoHabitacionRequest inicia la construcción para tipos de habitación.
func (sincronizacionNamespace) NewSincronizarParametricaTipoHabitacionBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion] {
	req := &sincronizacion.SincronizarParametricaTipoHabitacion{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTipoHabitacion]) SincronizarParametricaTipoHabitacion {
			return SincronizarParametricaTipoHabitacion{rw}
		},
	}
}

// NewSincronizarParametricaTipoMetodoPagoRequest inicia la construcción para métodos de pago.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMetodoPagoBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago] {
	req := &sincronizacion.SincronizarParametricaTipoMetodoPago{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTipoMetodoPago]) SincronizarParametricaTipoMetodoPago {
			return SincronizarParametricaTipoMetodoPago{rw}
		},
	}
}

// NewSincronizarParametricaTipoMonedaRequest inicia la construcción para tipos de moneda.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMonedaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda] {
	req := &sincronizacion.SincronizarParametricaTipoMoneda{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTipoMoneda]) SincronizarParametricaTipoMoneda {
			return SincronizarParametricaTipoMoneda{rw}
		},
	}
}

// NewSincronizarParametricaTipoPuntoVentaRequest inicia la construcción para tipos de punto de venta.
func (sincronizacionNamespace) NewSincronizarParametricaTipoPuntoVentaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta] {
	req := &sincronizacion.SincronizarParametricaTipoPuntoVenta{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTipoPuntoVenta]) SincronizarParametricaTipoPuntoVenta {
			return SincronizarParametricaTipoPuntoVenta{rw}
		},
	}
}

// NewSincronizarParametricaTiposFacturaRequest inicia la construcción para tipos de factura.
func (sincronizacionNamespace) NewSincronizarParametricaTiposFacturaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura] {
	req := &sincronizacion.SincronizarParametricaTiposFactura{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaTiposFactura]) SincronizarParametricaTiposFactura {
			return SincronizarParametricaTiposFactura{rw}
		},
	}
}

// NewSincronizarParametricaUnidadMedidaRequest inicia la construcción para unidades de medida.
func (sincronizacionNamespace) NewSincronizarParametricaUnidadMedidaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida, SincronizarParametricaUnidadMedida] {
	req := &sincronizacion.SincronizarParametricaUnidadMedida{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida, SincronizarParametricaUnidadMedida]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
		wrap: func(rw requestWrapper[sincronizacion.SincronizarParametricaUnidadMedida]) SincronizarParametricaUnidadMedida {
			return SincronizarParametricaUnidadMedida{rw}
		},
	}
}
