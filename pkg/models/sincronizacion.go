package models

import (
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/sincronizacion"
)

type sincronizacionNamespace struct{}

// Sincronizacion expone constructores de solicitudes para el módulo de Sincronización de Catálogos del SIAT.
var Sincronizacion = sincronizacionNamespace{}

// --- Interfaces opacas para las solicitudes de Sincronización ---

// SincronizacionRequest representa una solicitud genérica de sincronización o consulta de catálogos.
type SincronizacionRequest interface{ commonRequest() }

// SincronizacionBuilder es un generador genérico para configurar solicitudes de sincronización.
type SincronizacionBuilder[T any] struct {
	request *T
	sol     *sincronizacion.SolicitudSincronizacion
}

func (b *SincronizacionBuilder[T]) WithCodigoAmbiente(codigoAmbiente int) *SincronizacionBuilder[T] {
	b.sol.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *SincronizacionBuilder[T]) WithCodigoPuntoVenta(codigoPuntoVenta int) *SincronizacionBuilder[T] {
	b.sol.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *SincronizacionBuilder[T]) WithCodigoSistema(codigoSistema string) *SincronizacionBuilder[T] {
	b.sol.CodigoSistema = codigoSistema
	return b
}

func (b *SincronizacionBuilder[T]) WithCodigoSucursal(codigoSucursal int) *SincronizacionBuilder[T] {
	b.sol.CodigoSucursal = codigoSucursal
	return b
}

func (b *SincronizacionBuilder[T]) WithCuis(cuis string) *SincronizacionBuilder[T] {
	b.sol.Cuis = cuis
	return b
}

func (b *SincronizacionBuilder[T]) WithNit(nit int64) *SincronizacionBuilder[T] {
	b.sol.NIT = nit
	return b
}

// Build entrega el objeto de solicitud configurado.
func (b *SincronizacionBuilder[T]) Build() SincronizacionRequest {
	return requestWrapper[T]{request: b.request}
}

// NewSincronizarActividadesRequest inicia la construcción de una solicitud para sincronizar actividades económicas.
func (sincronizacionNamespace) NewSincronizarActividadesRequest() *SincronizacionBuilder[sincronizacion.SincronizarActividades] {
	req := &sincronizacion.SincronizarActividades{}
	return &SincronizacionBuilder[sincronizacion.SincronizarActividades]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaActividadesDocumentoSectorRequest inicia la construcción para la relación actividad-sector.
func (sincronizacionNamespace) NewSincronizarListaActividadesDocumentoSectorRequest() *SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector] {
	req := &sincronizacion.SincronizarListaActividadesDocumentoSector{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaActividadesDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaLeyendasFacturaRequest inicia la construcción para obtener leyendas de facturas.
func (sincronizacionNamespace) NewSincronizarListaLeyendasFacturaRequest() *SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura] {
	req := &sincronizacion.SincronizarListaLeyendasFactura{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaLeyendasFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaMensajesServiciosRequest inicia la construcción para obtener mensajes del servicio.
func (sincronizacionNamespace) NewSincronizarListaMensajesServiciosRequest() *SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios] {
	req := &sincronizacion.SincronizarListaMensajesServicios{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaMensajesServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarListaProductosServiciosRequest inicia la construcción para sincronizar productos y servicios.
func (sincronizacionNamespace) NewSincronizarListaProductosServiciosRequest() *SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios] {
	req := &sincronizacion.SincronizarListaProductosServicios{}
	return &SincronizacionBuilder[sincronizacion.SincronizarListaProductosServicios]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// Paramétricas

// NewSincronizarParametricaEventosSignificativosRequest inicia la construcción para el catálogo de eventos significativos.
func (sincronizacionNamespace) NewSincronizarParametricaEventosSignificativosRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos] {
	req := &sincronizacion.SincronizarParametricaEventosSignificativos{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaEventosSignificativos]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaMotivoAnulacionRequest inicia la construcción para motivos de anulación.
func (sincronizacionNamespace) NewSincronizarParametricaMotivoAnulacionRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion] {
	req := &sincronizacion.SincronizarParametricaMotivoAnulacion{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaMotivoAnulacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaPaisOrigenRequest inicia la construcción para el catálogo de países.
func (sincronizacionNamespace) NewSincronizarParametricaPaisOrigenRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen] {
	req := &sincronizacion.SincronizarParametricaPaisOrigen{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaPaisOrigen]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoDocumentoIdentidadRequest inicia la construcción para tipos de documento de identidad.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoIdentidadRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoIdentidad{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoIdentidad]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoDocumentoSectorRequest inicia la construcción para tipos de documento sector.
func (sincronizacionNamespace) NewSincronizarParametricaTipoDocumentoSectorRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector] {
	req := &sincronizacion.SincronizarParametricaTipoDocumentoSector{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoDocumentoSector]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoEmisionRequest inicia la construcción para tipos de emisión.
func (sincronizacionNamespace) NewSincronizarParametricaTipoEmisionRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision] {
	req := &sincronizacion.SincronizarParametricaTipoEmision{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoEmision]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoHabitacionRequest inicia la construcción para tipos de habitación.
func (sincronizacionNamespace) NewSincronizarParametricaTipoHabitacionRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion] {
	req := &sincronizacion.SincronizarParametricaTipoHabitacion{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoHabitacion]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoMetodoPagoRequest inicia la construcción para métodos de pago.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMetodoPagoRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago] {
	req := &sincronizacion.SincronizarParametricaTipoMetodoPago{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMetodoPago]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoMonedaRequest inicia la construcción para tipos de moneda.
func (sincronizacionNamespace) NewSincronizarParametricaTipoMonedaRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda] {
	req := &sincronizacion.SincronizarParametricaTipoMoneda{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoMoneda]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTipoPuntoVentaRequest inicia la construcción para tipos de punto de venta.
func (sincronizacionNamespace) NewSincronizarParametricaTipoPuntoVentaRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta] {
	req := &sincronizacion.SincronizarParametricaTipoPuntoVenta{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTipoPuntoVenta]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaTiposFacturaRequest inicia la construcción para tipos de factura.
func (sincronizacionNamespace) NewSincronizarParametricaTiposFacturaRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura] {
	req := &sincronizacion.SincronizarParametricaTiposFactura{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaTiposFactura]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}

// NewSincronizarParametricaUnidadMedidaRequest inicia la construcción para unidades de medida.
func (sincronizacionNamespace) NewSincronizarParametricaUnidadMedidaRequest() *SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida] {
	req := &sincronizacion.SincronizarParametricaUnidadMedida{}
	return &SincronizacionBuilder[sincronizacion.SincronizarParametricaUnidadMedida]{
		request: req,
		sol:     &req.SolicitudSincronizacion,
	}
}
