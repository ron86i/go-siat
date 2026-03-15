package models

import (
	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
)

type sincronizacionNamespace struct{}

func Sincronizacion() sincronizacionNamespace {
	return sincronizacionNamespace{}
}

// --- Interfaces opacas para las solicitudes de Sincronización ---

type SincronizarActividades interface{}
type SincronizarListaActividadesDocumentoSector interface{}
type SincronizarListaLeyendasFactura interface{}
type SincronizarListaMensajesServicios interface{}
type SincronizarListaProductosServicios interface{}
type SincronizarParametricaEventosSignificativos interface{}
type SincronizarParametricaMotivoAnulacion interface{}
type SincronizarParametricaPaisOrigen interface{}
type SincronizarParametricaTipoDocumentoIdentidad interface{}
type SincronizarParametricaTipoDocumentoSector interface{}
type SincronizarParametricaTipoEmision interface{}
type SincronizarParametricaTipoHabitacion interface{}
type SincronizarParametricaTipoMetodoPago interface{}
type SincronizarParametricaTipoMoneda interface{}
type SincronizarParametricaTipoPuntoVenta interface{}
type SincronizarParametricaTiposFactura interface{}
type SincronizarParametricaUnidadMedida interface{}
type VerificarComunicacionSincronizacion interface{}

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
	return requestWrapper[sincronizacion.VerificarComunicacion]{request: b.request}
}

// SincronizacionBuilder es un generador genérico para configurar solicitudes de sincronización.
type SincronizacionBuilder[T any, R any] struct {
	request *T
	sol     *sincronizacion.SolicitudSincronizacion
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
	var ret any = requestWrapper[T]{request: b.request}
	return ret.(R)
}

// NewSincronizarActividadesRequest inicia la construcción de una solicitud para sincronizar actividades económicas.
func (sincronizacionNamespace) NewSincronizarActividadesBuilder() *SincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades] {
	req := &sincronizacion.SincronizarActividades{}
	return &SincronizacionBuilder[sincronizacion.SincronizarActividades, SincronizarActividades]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaActividadesDocumentoSectorRequest inicia la construcción para la relación actividad-sector.
func (sincronizacionNamespace) NewSincronizarListaActividadesDocumentoSectorBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector] {
	req := &sincronizacion.SincronizarListaActividadesDocumentoSector{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector, SincronizarListaActividadesDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaLeyendasFacturaRequest inicia la construcción para obtener leyendas de facturas.
func (sincronizacionNamespace) NewSincronizarListaLeyendasFacturaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura] {
	req := &sincronizacion.SincronizarListaLeyendasFactura{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura, SincronizarListaLeyendasFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaMensajesServiciosRequest inicia la construcción para obtener mensajes del servicio.
func (sincronizacionNamespace) NewSincronizarListaMensajesServiciosBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios] {
	req := &sincronizacion.SincronizarListaMensajesServicios{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios, SincronizarListaMensajesServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaProductosServiciosRequest inicia la construcción para sincronizar productos y servicios.
func (sincronizacionNamespace) NewSincronizarListaProductosServiciosBuilder() *SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios] {
	req := &sincronizacion.SincronizarListaProductosServicios{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios, SincronizarListaProductosServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// Paramétricas

// NewSincronizarParametricaEventosSignificativosRequest inicia la construcción para el catálogo de eventos significativos.
func (sincronizacionNamespace) NewSincronizarParametricaEventosSignificativosBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos] {
	req := &sincronizacion.SincronizarParametricaEventosSignificativos{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos, SincronizarParametricaEventosSignificativos]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaMotivoAnulacionRequest inicia la construcción para motivos de anulación.
func (sincronizacionNamespace) NewSincronizarParametricaMotivoAnulacionBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion] {
	req := &sincronizacion.SincronizarParametricaMotivoAnulacion{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion, SincronizarParametricaMotivoAnulacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaPaisOrigenRequest inicia la construcción para el catálogo de países.
func (sincronizacionNamespace) NewSincronizarParametricaPaisOrigenBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen] {
	req := &sincronizacion.SincronizarParametricaPaisOrigen{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen, SincronizarParametricaPaisOrigen]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoDocumentoIdentidadRequest inicia la construcción para tipos de documento de identidad.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoIdentidadBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoIdentidad{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad, SincronizarParametricaTipoDocumentoIdentidad]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoDocumentoSectorRequest inicia la construcción para tipos de documento sector.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoSectorBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoSector{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector, SincronizarParametricaTipoDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoEmisionRequest inicia la construcción para tipos de emisión.
func (sincronizacionNamespace) NewSincronizarParametricaTipoEmisionBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision] {
	req := &sincronizacion.SincronizarParametricaTipoEmision{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision, SincronizarParametricaTipoEmision]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoHabitacionRequest inicia la construcción para tipos de habitación.
func (sincronizacionNamespace) NewSincronizarParametricaTipoHabitacionBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion] {
	req := &sincronizacion.SincronizarParametricaTipoHabitacion{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion, SincronizarParametricaTipoHabitacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoMetodoPagoRequest inicia la construcción para métodos de pago.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMetodoPagoBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago] {
	req := &sincronizacion.SincronizarParametricaTipoMetodoPago{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago, SincronizarParametricaTipoMetodoPago]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoMonedaRequest inicia la construcción para tipos de moneda.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMonedaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda] {
	req := &sincronizacion.SincronizarParametricaTipoMoneda{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda, SincronizarParametricaTipoMoneda]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoPuntoVentaRequest inicia la construcción para tipos de punto de venta.
func (sincronizacionNamespace) NewSincronizarParametricaTipoPuntoVentaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta] {
	req := &sincronizacion.SincronizarParametricaTipoPuntoVenta{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta, SincronizarParametricaTipoPuntoVenta]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTiposFacturaRequest inicia la construcción para tipos de factura.
func (sincronizacionNamespace) NewSincronizarParametricaTiposFacturaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura] {
	req := &sincronizacion.SincronizarParametricaTiposFactura{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura, SincronizarParametricaTiposFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaUnidadMedidaRequest inicia la construcción para unidades de medida.
func (sincronizacionNamespace) NewSincronizarParametricaUnidadMedidaBuilder() *SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida, SincronizarParametricaUnidadMedida] {
	req := &sincronizacion.SincronizarParametricaUnidadMedida{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida, SincronizarParametricaUnidadMedida]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}
