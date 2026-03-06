package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/facturacion/operaciones"
)

type operacionesNamespace struct{}

// Operaciones expone constructores de solicitudes para el módulo de Operaciones del SIAT.
var Operaciones = operacionesNamespace{}

// --- Interfaces opacas para las solicitudes de Operaciones ---

// RegistroPuntoVentaRequest representa una solicitud para registrar un nuevo punto de venta.
type RegistroPuntoVentaRequest interface{ commonRequest() }

// ConsultaPuntoVentaRequest representa una solicitud para listar puntos de venta registrados.
type ConsultaPuntoVentaRequest interface{ commonRequest() }

// CierrePuntoVentaRequest representa una solicitud para cerrar un punto de venta.
type CierrePuntoVentaRequest interface{ commonRequest() }

// RegistroPuntoVentaComisionistaRequest representa una solicitud para registrar un comisionista.
type RegistroPuntoVentaComisionistaRequest interface{ commonRequest() }

// CierreOperacionesSistemaRequest representa una solicitud para el cierre del sistema.
type CierreOperacionesSistemaRequest interface{ commonRequest() }

// ConsultaEventoSignificativoRequest representa una solicitud para consultar eventos registrados.
type ConsultaEventoSignificativoRequest interface{ commonRequest() }

// RegistroEventoSignificativoRequest representa una solicitud para informar un evento significativo.
type RegistroEventoSignificativoRequest interface{ commonRequest() }

// NewRegistroPuntoVentaRequest inicia la construcción de una solicitud para registrar un nuevo punto de venta.
func (operacionesNamespace) NewRegistroPuntoVentaRequest() *RegistroPuntoVentaBuilder {
	return &RegistroPuntoVentaBuilder{
		request: &operaciones.RegistroPuntoVenta{},
	}
}

// RegistroPuntoVentaBuilder ayuda a configurar la solicitud de registro de punto de venta.
type RegistroPuntoVentaBuilder struct {
	request *operaciones.RegistroPuntoVenta
}

func (b *RegistroPuntoVentaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RegistroPuntoVentaBuilder) WithCodigoModalidad(codigoModalidad int) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoModalidad = codigoModalidad
	return b
}

func (b *RegistroPuntoVentaBuilder) WithCodigoSistema(codigoSistema string) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoSistema = codigoSistema
	return b
}

func (b *RegistroPuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *RegistroPuntoVentaBuilder) WithCodigoTipoPuntoVenta(codigoTipoPuntoVenta int) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoTipoPuntoVenta = codigoTipoPuntoVenta
	return b
}

func (b *RegistroPuntoVentaBuilder) WithCuis(cuis string) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.Cuis = cuis
	return b
}

func (b *RegistroPuntoVentaBuilder) WithDescripcion(descripcion string) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.Descripcion = descripcion
	return b
}

func (b *RegistroPuntoVentaBuilder) WithNit(nit int64) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.Nit = nit
	return b
}

func (b *RegistroPuntoVentaBuilder) WithNombrePuntoVenta(nombrePuntoVenta string) *RegistroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.NombrePuntoVenta = nombrePuntoVenta
	return b
}

// Build retorna el objeto RegistroPuntoVenta configurado.
func (b *RegistroPuntoVentaBuilder) Build() RegistroPuntoVentaRequest {
	return requestWrapper[operaciones.RegistroPuntoVenta]{request: b.request}
}

// NewConsultaPuntoVentaRequest inicia la construcción de una solicitud para consultar puntos de venta.
func (operacionesNamespace) NewConsultaPuntoVentaRequest() *ConsultaPuntoVentaBuilder {
	return &ConsultaPuntoVentaBuilder{
		request: &operaciones.ConsultaPuntoVenta{},
	}
}

// ConsultaPuntoVentaBuilder facilita la configuración de la consulta de puntos de venta.
type ConsultaPuntoVentaBuilder struct {
	request *operaciones.ConsultaPuntoVenta
}

func (b *ConsultaPuntoVentaBuilder) WithCodigoAmbiente(codigoAmbiente int) *ConsultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *ConsultaPuntoVentaBuilder) WithCodigoSistema(codigoSistema string) *ConsultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.CodigoSistema = codigoSistema
	return b
}

func (b *ConsultaPuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *ConsultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *ConsultaPuntoVentaBuilder) WithCuis(cuis string) *ConsultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.Cuis = cuis
	return b
}

func (b *ConsultaPuntoVentaBuilder) WithNit(nit int64) *ConsultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.Nit = nit
	return b
}

// Build retorna el objeto ConsultaPuntoVenta configurado.
func (b *ConsultaPuntoVentaBuilder) Build() ConsultaPuntoVentaRequest {
	return requestWrapper[operaciones.ConsultaPuntoVenta]{request: b.request}
}

// NewCierrePuntoVentaRequest inicia la construcción de una solicitud para cerrar un punto de venta.
func (operacionesNamespace) NewCierrePuntoVentaRequest() *CierrePuntoVentaBuilder {
	return &CierrePuntoVentaBuilder{
		request: &operaciones.CierrePuntoVenta{},
	}
}

// CierrePuntoVentaBuilder ayuda a configurar el cierre de un punto de venta.
type CierrePuntoVentaBuilder struct {
	request *operaciones.CierrePuntoVenta
}

func (b *CierrePuntoVentaBuilder) WithCodigoAmbiente(codigoAmbiente int) *CierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *CierrePuntoVentaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *CierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *CierrePuntoVentaBuilder) WithCodigoSistema(codigoSistema string) *CierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoSistema = codigoSistema
	return b
}

func (b *CierrePuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *CierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *CierrePuntoVentaBuilder) WithCuis(cuis string) *CierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.Cuis = cuis
	return b
}

func (b *CierrePuntoVentaBuilder) WithNit(nit int64) *CierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.Nit = nit
	return b
}

// Build retorna el objeto CierrePuntoVenta configurado.
func (b *CierrePuntoVentaBuilder) Build() CierrePuntoVentaRequest {
	return requestWrapper[operaciones.CierrePuntoVenta]{request: b.request}
}

// NewRegistroPuntoVentaComisionistaRequest inicia la construcción de una solicitud para un comisionista.
func (operacionesNamespace) NewRegistroPuntoVentaComisionistaRequest() *RegistroPuntoVentaComisionistaBuilder {
	return &RegistroPuntoVentaComisionistaBuilder{
		request: &operaciones.RegistroPuntoVentaComisionista{},
	}
}

// RegistroPuntoVentaComisionistaBuilder facilita la configuración de puntos de venta comisionistas.
type RegistroPuntoVentaComisionistaBuilder struct {
	request *operaciones.RegistroPuntoVentaComisionista
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithCodigoModalidad(codigoModalidad int) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoModalidad = codigoModalidad
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithCodigoSistema(codigoSistema string) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoSistema = codigoSistema
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithCodigoSucursal(codigoSucursal int) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoSucursal = codigoSucursal
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithCuis(cuis string) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Cuis = cuis
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithNit(nit int64) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Nit = nit
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithNitComisionista(nitComisionista int64) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NitComisionista = nitComisionista
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithNombrePuntoVenta(nombrePuntoVenta string) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NombrePuntoVenta = nombrePuntoVenta
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithDescripcion(descripcion string) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Descripcion = descripcion
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithNumeroContrato(numeroContrato string) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NumeroContrato = numeroContrato
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithFechaInicio(fechaInicio time.Time) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.FechaInicio = fechaInicio
	return b
}

func (b *RegistroPuntoVentaComisionistaBuilder) WithFechaFin(fechaFin time.Time) *RegistroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.FechaFin = fechaFin
	return b
}

// Build retorna el objeto RegistroPuntoVentaComisionista configurado.
func (b *RegistroPuntoVentaComisionistaBuilder) Build() RegistroPuntoVentaComisionistaRequest {
	return requestWrapper[operaciones.RegistroPuntoVentaComisionista]{request: b.request}
}

// NewCierreOperacionesSistemaRequest inicia la construcción de una solicitud para cerrar operaciones del sistema.
func (operacionesNamespace) NewCierreOperacionesSistemaRequest() *CierreOperacionesSistemaBuilder {
	return &CierreOperacionesSistemaBuilder{
		request: &operaciones.CierreOperacionesSistema{},
	}
}

// CierreOperacionesSistemaBuilder ayuda a configurar el cierre de operaciones del sistema.
type CierreOperacionesSistemaBuilder struct {
	request *operaciones.CierreOperacionesSistema
}

func (b *CierreOperacionesSistemaBuilder) WithCodigoAmbiente(codigoAmbiente int) *CierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *CierreOperacionesSistemaBuilder) WithCodigoModalidad(codigoModalidad int) *CierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoModalidad = codigoModalidad
	return b
}

func (b *CierreOperacionesSistemaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *CierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *CierreOperacionesSistemaBuilder) WithCodigoSistema(codigoSistema string) *CierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoSistema = codigoSistema
	return b
}

func (b *CierreOperacionesSistemaBuilder) WithCodigoSucursal(codigoSucursal int) *CierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoSucursal = codigoSucursal
	return b
}

func (b *CierreOperacionesSistemaBuilder) WithCuis(cuis string) *CierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.Cuis = cuis
	return b
}

func (b *CierreOperacionesSistemaBuilder) WithNit(nit int64) *CierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.Nit = nit
	return b
}

// Build retorna el objeto CierreOperacionesSistema configurado.
func (b *CierreOperacionesSistemaBuilder) Build() CierreOperacionesSistemaRequest {
	return requestWrapper[operaciones.CierreOperacionesSistema]{request: b.request}
}

func (operacionesNamespace) NewVerificiarComunicacionOperaciones() *operaciones.VerificarComunicacion {
	return &operaciones.VerificarComunicacion{}
}
