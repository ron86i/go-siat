package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/siat/operaciones"
)

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

// --- Namespace ---

type operacionesNamespace struct{}

func Operaciones() operacionesNamespace {
	return operacionesNamespace{}
}

// --- Constructores de Builders ---

// NewRegistroPuntoVentaBuilder inicializa un builder para registrar un nuevo punto de venta.
func (operacionesNamespace) NewRegistroPuntoVentaBuilder() *registroPuntoVentaBuilder {
	return &registroPuntoVentaBuilder{
		request: &operaciones.RegistroPuntoVenta{},
	}
}

// NewConsultaPuntoVentaBuilder inicializa un builder para la consulta de puntos de venta.
func (operacionesNamespace) NewConsultaPuntoVentaBuilder() *consultaPuntoVentaBuilder {
	return &consultaPuntoVentaBuilder{
		request: &operaciones.ConsultaPuntoVenta{},
	}
}

// NewCierrePuntoVentaBuilder inicializa un builder para cerrar un punto de venta.
func (operacionesNamespace) NewCierrePuntoVentaBuilder() *cierrePuntoVentaBuilder {
	return &cierrePuntoVentaBuilder{
		request: &operaciones.CierrePuntoVenta{},
	}
}

// NewRegistroPuntoVentaComisionistaBuilder inicializa un builder para registrar un comisionista en un punto de venta.
func (operacionesNamespace) NewRegistroPuntoVentaComisionistaBuilder() *registroPuntoVentaComisionistaBuilder {
	return &registroPuntoVentaComisionistaBuilder{
		request: &operaciones.RegistroPuntoVentaComisionista{},
	}
}

// NewCierreOperacionesSistemaBuilder inicializa la configuración base para el cierre de operaciones.
func (operacionesNamespace) NewCierreOperacionesSistemaBuilder() *cierreOperacionesSistemaBuilder {
	return &cierreOperacionesSistemaBuilder{
		request: &operaciones.CierreOperacionesSistema{},
	}
}

func (operacionesNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionOperacionesBuilder {
	return &verificarComunicacionOperacionesBuilder{
		request: &operaciones.VerificarComunicacion{},
	}
}

// NewRegistroEventoSignificativoBuilder inicializa un builder para registrar un evento significativo.
func (operacionesNamespace) NewRegistroEventoSignificativoBuilder() *registroEventoSignificativoBuilder {
	return &registroEventoSignificativoBuilder{
		request: &operaciones.RegistroEventoSignificativo{},
	}
}

// NewConsultaEventoSignificativoBuilder inicializa un builder para consultar eventos significativos.
func (operacionesNamespace) NewConsultaEventoSignificativoBuilder() *consultaEventoSignificativoBuilder {
	return &consultaEventoSignificativoBuilder{
		request: &operaciones.ConsultaEventoSignificativo{},
	}
}

// --- Implementaciones de Builders ---

// registroPuntoVentaBuilder ayuda a configurar la solicitud de registro de punto de venta.
type registroPuntoVentaBuilder struct {
	request *operaciones.RegistroPuntoVenta
}

func (b *registroPuntoVentaBuilder) WithCodigoAmbiente(codigoAmbiente int) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *registroPuntoVentaBuilder) WithCodigoModalidad(codigoModalidad int) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoModalidad = codigoModalidad
	return b
}

func (b *registroPuntoVentaBuilder) WithCodigoSistema(codigoSistema string) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoSistema = codigoSistema
	return b
}

func (b *registroPuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *registroPuntoVentaBuilder) WithCodigoTipoPuntoVenta(codigoTipoPuntoVenta int) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.CodigoTipoPuntoVenta = codigoTipoPuntoVenta
	return b
}

func (b *registroPuntoVentaBuilder) WithCuis(cuis string) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.Cuis = cuis
	return b
}

func (b *registroPuntoVentaBuilder) WithDescripcion(descripcion string) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.Descripcion = descripcion
	return b
}

func (b *registroPuntoVentaBuilder) WithNit(nit int64) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.Nit = nit
	return b
}

func (b *registroPuntoVentaBuilder) WithNombrePuntoVenta(nombrePuntoVenta string) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.NombrePuntoVenta = nombrePuntoVenta
	return b
}

// Build retorna el objeto RegistroPuntoVenta configurado.
func (b *registroPuntoVentaBuilder) Build() RegistroPuntoVenta {
	return RegistroPuntoVenta{requestWrapper[operaciones.RegistroPuntoVenta]{request: b.request}}
}

// consultaPuntoVentaBuilder facilita la configuración de la consulta de puntos de venta.
type consultaPuntoVentaBuilder struct {
	request *operaciones.ConsultaPuntoVenta
}

func (b *consultaPuntoVentaBuilder) WithCodigoAmbiente(codigoAmbiente int) *consultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *consultaPuntoVentaBuilder) WithCodigoSistema(codigoSistema string) *consultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.CodigoSistema = codigoSistema
	return b
}

func (b *consultaPuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *consultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *consultaPuntoVentaBuilder) WithCuis(cuis string) *consultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.Cuis = cuis
	return b
}

func (b *consultaPuntoVentaBuilder) WithNit(nit int64) *consultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.Nit = nit
	return b
}

// Build retorna el objeto ConsultaPuntoVenta configurado.
func (b *consultaPuntoVentaBuilder) Build() ConsultaPuntoVenta {
	return ConsultaPuntoVenta{requestWrapper[operaciones.ConsultaPuntoVenta]{request: b.request}}
}

// cierrePuntoVentaBuilder ayuda a configurar el cierre de un punto de venta.
type cierrePuntoVentaBuilder struct {
	request *operaciones.CierrePuntoVenta
}

func (b *cierrePuntoVentaBuilder) WithCodigoAmbiente(codigoAmbiente int) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *cierrePuntoVentaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *cierrePuntoVentaBuilder) WithCodigoSistema(codigoSistema string) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoSistema = codigoSistema
	return b
}

func (b *cierrePuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *cierrePuntoVentaBuilder) WithCuis(cuis string) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.Cuis = cuis
	return b
}

func (b *cierrePuntoVentaBuilder) WithNit(nit int64) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.Nit = nit
	return b
}

// Build retorna el objeto CierrePuntoVenta configurado.
func (b *cierrePuntoVentaBuilder) Build() CierrePuntoVenta {
	return CierrePuntoVenta{requestWrapper[operaciones.CierrePuntoVenta]{request: b.request}}
}

// registroPuntoVentaComisionistaBuilder facilita la configuración de puntos de venta comisionistas.
type registroPuntoVentaComisionistaBuilder struct {
	request *operaciones.RegistroPuntoVentaComisionista
}

func (b *registroPuntoVentaComisionistaBuilder) WithCodigoAmbiente(codigoAmbiente int) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithCodigoModalidad(codigoModalidad int) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoModalidad = codigoModalidad
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithCodigoSistema(codigoSistema string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoSistema = codigoSistema
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithCodigoSucursal(codigoSucursal int) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoSucursal = codigoSucursal
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithCuis(cuis string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Cuis = cuis
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithNit(nit int64) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Nit = nit
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithNitComisionista(nitComisionista int64) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NitComisionista = nitComisionista
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithNombrePuntoVenta(nombrePuntoVenta string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NombrePuntoVenta = nombrePuntoVenta
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithDescripcion(descripcion string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Descripcion = descripcion
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithNumeroContrato(numeroContrato string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NumeroContrato = numeroContrato
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithFechaInicio(fechaInicio time.Time) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.FechaInicio = fechaInicio
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithFechaFin(fechaFin time.Time) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.FechaFin = fechaFin
	return b
}

// Build retorna el objeto RegistroPuntoVentaComisionista configurado.
func (b *registroPuntoVentaComisionistaBuilder) Build() RegistroPuntoVentaComisionista {
	return RegistroPuntoVentaComisionista{requestWrapper[operaciones.RegistroPuntoVentaComisionista]{request: b.request}}
}

// cierreOperacionesSistemaBuilder ayuda a configurar el cierre de operaciones del sistema.
type cierreOperacionesSistemaBuilder struct {
	request *operaciones.CierreOperacionesSistema
}

func (b *cierreOperacionesSistemaBuilder) WithCodigoAmbiente(codigoAmbiente int) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithCodigoModalidad(codigoModalidad int) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoModalidad = codigoModalidad
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithCodigoSistema(codigoSistema string) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoSistema = codigoSistema
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithCodigoSucursal(codigoSucursal int) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoSucursal = codigoSucursal
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithCuis(cuis string) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.Cuis = cuis
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithNit(nit int64) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.Nit = nit
	return b
}

// Build retorna el objeto CierreOperacionesSistema configurado.
func (b *cierreOperacionesSistemaBuilder) Build() CierreOperacionesSistema {
	return CierreOperacionesSistema{requestWrapper[operaciones.CierreOperacionesSistema]{request: b.request}}
}

// RegistroEventoSignificativoBuilder ayuda a configurar el registro de un evento significativo.
type verificarComunicacionOperacionesBuilder struct {
	request *operaciones.VerificarComunicacion
}

func (b *verificarComunicacionOperacionesBuilder) Build() VerificarComunicacionOperaciones {
	return VerificarComunicacionOperaciones{requestWrapper[operaciones.VerificarComunicacion]{request: b.request}}
}

// registroEventoSignificativoBuilder ayuda a configurar el registro de un evento significativo.
type registroEventoSignificativoBuilder struct {
	request *operaciones.RegistroEventoSignificativo
}

func (b *registroEventoSignificativoBuilder) WithCodigoAmbiente(codigoAmbiente int) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *registroEventoSignificativoBuilder) WithCodigoMotivoEvento(codigoMotivoEvento int) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoMotivoEvento = codigoMotivoEvento
	return b
}

func (b *registroEventoSignificativoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *registroEventoSignificativoBuilder) WithCodigoSistema(codigoSistema string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoSistema = codigoSistema
	return b
}

func (b *registroEventoSignificativoBuilder) WithCodigoSucursal(codigoSucursal int) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoSucursal = codigoSucursal
	return b
}

func (b *registroEventoSignificativoBuilder) WithCuis(cuis string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Cuis = cuis
	return b
}

func (b *registroEventoSignificativoBuilder) WithCufdEvento(cufdEvento string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CufdEvento = cufdEvento
	return b
}

func (b *registroEventoSignificativoBuilder) WithCufd(cufd string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Cufd = cufd
	return b
}

func (b *registroEventoSignificativoBuilder) WithDescripcion(descripcion string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Descripcion = descripcion
	return b
}

func (b *registroEventoSignificativoBuilder) WithFechaHoraInicioEvento(fecha time.Time) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.FechaHoraInicioEvento = fecha
	return b
}

func (b *registroEventoSignificativoBuilder) WithFechaHoraFinEvento(fecha time.Time) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.FechaHoraFinEvento = fecha
	return b
}

func (b *registroEventoSignificativoBuilder) WithNit(nit int64) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Nit = nit
	return b
}

func (b *registroEventoSignificativoBuilder) Build() RegistroEventoSignificativo {
	return RegistroEventoSignificativo{requestWrapper[operaciones.RegistroEventoSignificativo]{request: b.request}}
}

// consultaEventoSignificativoBuilder ayuda a configurar la consulta de eventos significativos.
type consultaEventoSignificativoBuilder struct {
	request *operaciones.ConsultaEventoSignificativo
}

func (b *consultaEventoSignificativoBuilder) WithCodigoAmbiente(codigoAmbiente int) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *consultaEventoSignificativoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *consultaEventoSignificativoBuilder) WithCodigoSistema(codigoSistema string) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoSistema = codigoSistema
	return b
}

func (b *consultaEventoSignificativoBuilder) WithCodigoSucursal(codigoSucursal int) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoSucursal = codigoSucursal
	return b
}

func (b *consultaEventoSignificativoBuilder) WithCuis(cuis string) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.Cuis = cuis
	return b
}

func (b *consultaEventoSignificativoBuilder) WithFechaEvento(fecha time.Time) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.FechaEvento = fecha
	return b
}

func (b *consultaEventoSignificativoBuilder) WithNit(nit int64) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.Nit = nit
	return b
}

func (b *consultaEventoSignificativoBuilder) Build() ConsultaEventoSignificativo {
	return ConsultaEventoSignificativo{requestWrapper[operaciones.ConsultaEventoSignificativo]{request: b.request}}
}
