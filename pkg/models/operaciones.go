package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/siat/operaciones"
)

type operacionesNamespace struct{}

func Operaciones() operacionesNamespace {
	return operacionesNamespace{}
}

// --- Interfaces opacas para las solicitudes de Operaciones ---

type VerificarComunicacionOperaciones struct {
	requestWrapper[operaciones.VerificarComunicacion]
}

// RegistroPuntoVenta representa una solicitud para registrar un nuevo punto de venta.
type RegistroPuntoVenta struct {
	requestWrapper[operaciones.RegistroPuntoVenta]
}

// ConsultaPuntoVenta representa una solicitud para listar puntos de venta registrados.
type ConsultaPuntoVenta struct {
	requestWrapper[operaciones.ConsultaPuntoVenta]
}

// CierrePuntoVenta representa una solicitud para cerrar un punto de venta.
type CierrePuntoVenta struct {
	requestWrapper[operaciones.CierrePuntoVenta]
}

// RegistroPuntoVentaComisionista representa una solicitud para registrar un comisionista.
type RegistroPuntoVentaComisionista struct {
	requestWrapper[operaciones.RegistroPuntoVentaComisionista]
}

// CierreOperacionesSistema representa una solicitud para el cierre del sistema.
type CierreOperacionesSistema struct {
	requestWrapper[operaciones.CierreOperacionesSistema]
}

// ConsultaEventoSignificativo representa una solicitud para consultar eventos registrados.
type ConsultaEventoSignificativo struct {
	requestWrapper[operaciones.ConsultaEventoSignificativo]
}

// RegistroEventoSignificativo representa una solicitud para informar un evento significativo.
type RegistroEventoSignificativo struct {
	requestWrapper[operaciones.RegistroEventoSignificativo]
}

// NewRegistroPuntoVentaBuilder inicializa un builder para registrar un nuevo punto de venta.
func (operacionesNamespace) NewRegistroPuntoVentaBuilder() *RegistroPuntoVentaBuilder {
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
func (b *RegistroPuntoVentaBuilder) Build() RegistroPuntoVenta {
	return RegistroPuntoVenta{requestWrapper[operaciones.RegistroPuntoVenta]{request: b.request}}
}

// NewConsultaPuntoVentaBuilder inicializa un builder para la consulta de puntos de venta.
func (operacionesNamespace) NewConsultaPuntoVentaBuilder() *ConsultaPuntoVentaBuilder {
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
func (b *ConsultaPuntoVentaBuilder) Build() ConsultaPuntoVenta {
	return ConsultaPuntoVenta{requestWrapper[operaciones.ConsultaPuntoVenta]{request: b.request}}
}

// NewCierrePuntoVentaBuilder inicializa un builder para cerrar un punto de venta.
func (operacionesNamespace) NewCierrePuntoVentaBuilder() *CierrePuntoVentaBuilder {
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
func (b *CierrePuntoVentaBuilder) Build() CierrePuntoVenta {
	return CierrePuntoVenta{requestWrapper[operaciones.CierrePuntoVenta]{request: b.request}}
}

// NewRegistroPuntoVentaComisionistaBuilder inicializa un builder para registrar un comisionista en un punto de venta.
func (operacionesNamespace) NewRegistroPuntoVentaComisionistaBuilder() *RegistroPuntoVentaComisionistaBuilder {
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
func (b *RegistroPuntoVentaComisionistaBuilder) Build() RegistroPuntoVentaComisionista {
	return RegistroPuntoVentaComisionista{requestWrapper[operaciones.RegistroPuntoVentaComisionista]{request: b.request}}
}

// NewCierreOperacionesSistemaBuilder inicializa la configuración base para el cierre de operaciones.
func (operacionesNamespace) NewCierreOperacionesSistemaBuilder() *CierreOperacionesSistemaBuilder {
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
func (b *CierreOperacionesSistemaBuilder) Build() CierreOperacionesSistema {
	return CierreOperacionesSistema{requestWrapper[operaciones.CierreOperacionesSistema]{request: b.request}}
}

func (operacionesNamespace) NewVerificarComunicacionBuilder() *VerificarComunicacionOperacionesBuilder {
	return &VerificarComunicacionOperacionesBuilder{
		request: &operaciones.VerificarComunicacion{},
	}
}

// RegistroEventoSignificativoBuilder ayuda a configurar el registro de un evento significativo.
type VerificarComunicacionOperacionesBuilder struct {
	request *operaciones.VerificarComunicacion
}

func (b *VerificarComunicacionOperacionesBuilder) Build() VerificarComunicacionOperaciones {
	return VerificarComunicacionOperaciones{requestWrapper[operaciones.VerificarComunicacion]{request: b.request}}
}

// NewRegistroEventoSignificativoBuilder inicializa un builder para registrar un evento significativo.
func (operacionesNamespace) NewRegistroEventoSignificativoBuilder() *RegistroEventoSignificativoBuilder {
	return &RegistroEventoSignificativoBuilder{
		request: &operaciones.RegistroEventoSignificativo{},
	}
}

// RegistroEventoSignificativoBuilder ayuda a configurar el registro de un evento significativo.
type RegistroEventoSignificativoBuilder struct {
	request *operaciones.RegistroEventoSignificativo
}

func (b *RegistroEventoSignificativoBuilder) WithCodigoAmbiente(codigoAmbiente int) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithCodigoMotivoEvento(codigoMotivoEvento int) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoMotivoEvento = codigoMotivoEvento
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithCodigoSistema(codigoSistema string) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoSistema = codigoSistema
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithCodigoSucursal(codigoSucursal int) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoSucursal = codigoSucursal
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithCuis(cuis string) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Cuis = cuis
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithCufdEvento(cufdEvento string) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CufdEvento = cufdEvento
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithCufd(cufd string) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Cufd = cufd
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithDescripcion(descripcion string) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Descripcion = descripcion
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithFechaHoraInicioEvento(fecha time.Time) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.FechaHoraInicioEvento = fecha
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithFechaHoraFinEvento(fecha time.Time) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.FechaHoraFinEvento = fecha
	return b
}

func (b *RegistroEventoSignificativoBuilder) WithNit(nit int64) *RegistroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Nit = nit
	return b
}

func (b *RegistroEventoSignificativoBuilder) Build() RegistroEventoSignificativo {
	return RegistroEventoSignificativo{requestWrapper[operaciones.RegistroEventoSignificativo]{request: b.request}}
}

// NewConsultaEventoSignificativoBuilder inicializa un builder para consultar eventos significativos.
func (operacionesNamespace) NewConsultaEventoSignificativoBuilder() *ConsultaEventoSignificativoBuilder {
	return &ConsultaEventoSignificativoBuilder{
		request: &operaciones.ConsultaEventoSignificativo{},
	}
}

// ConsultaEventoSignificativoBuilder ayuda a configurar la consulta de eventos significativos.
type ConsultaEventoSignificativoBuilder struct {
	request *operaciones.ConsultaEventoSignificativo
}

func (b *ConsultaEventoSignificativoBuilder) WithCodigoAmbiente(codigoAmbiente int) *ConsultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *ConsultaEventoSignificativoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *ConsultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *ConsultaEventoSignificativoBuilder) WithCodigoSistema(codigoSistema string) *ConsultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoSistema = codigoSistema
	return b
}

func (b *ConsultaEventoSignificativoBuilder) WithCodigoSucursal(codigoSucursal int) *ConsultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoSucursal = codigoSucursal
	return b
}

func (b *ConsultaEventoSignificativoBuilder) WithCuis(cuis string) *ConsultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.Cuis = cuis
	return b
}

func (b *ConsultaEventoSignificativoBuilder) WithFechaEvento(fecha time.Time) *ConsultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.FechaEvento = fecha
	return b
}

func (b *ConsultaEventoSignificativoBuilder) WithNit(nit int64) *ConsultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.Nit = nit
	return b
}

func (b *ConsultaEventoSignificativoBuilder) Build() ConsultaEventoSignificativo {
	return ConsultaEventoSignificativo{requestWrapper[operaciones.ConsultaEventoSignificativo]{request: b.request}}
}
