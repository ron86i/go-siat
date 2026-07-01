package models

import (
	"time"

	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/operaciones"
)

// --- Interfaces opacas para las solicitudes de Operaciones ---

type VerificarComunicacionOperaciones struct {
	RequestWrapper[operaciones.VerificarComunicacion]
}

type RegistroPuntoVenta struct {
	RequestWrapper[operaciones.RegistroPuntoVenta]
}

type ConsultaPuntoVenta struct {
	RequestWrapper[operaciones.ConsultaPuntoVenta]
}

type CierrePuntoVenta struct {
	RequestWrapper[operaciones.CierrePuntoVenta]
}

type RegistroPuntoVentaComisionista struct {
	RequestWrapper[operaciones.RegistroPuntoVentaComisionista]
}

type CierreOperacionesSistema struct {
	RequestWrapper[operaciones.CierreOperacionesSistema]
}

type ConsultaEventoSignificativo struct {
	RequestWrapper[operaciones.ConsultaEventoSignificativo]
}

type RegistroEventoSignificativo struct {
	RequestWrapper[operaciones.RegistroEventoSignificativo]
}

// --- Constructores de Builders a nivel de paquete ---

func NewRegistroPuntoVentaBuilder() *registroPuntoVentaBuilder {
	return &registroPuntoVentaBuilder{
		request: &operaciones.RegistroPuntoVenta{},
	}
}

func NewConsultaPuntoVentaBuilder() *consultaPuntoVentaBuilder {
	return &consultaPuntoVentaBuilder{
		request: &operaciones.ConsultaPuntoVenta{},
	}
}

func NewCierrePuntoVentaBuilder() *cierrePuntoVentaBuilder {
	return &cierrePuntoVentaBuilder{
		request: &operaciones.CierrePuntoVenta{},
	}
}

func NewRegistroPuntoVentaComisionistaBuilder() *registroPuntoVentaComisionistaBuilder {
	return &registroPuntoVentaComisionistaBuilder{
		request: &operaciones.RegistroPuntoVentaComisionista{},
	}
}

func NewCierreOperacionesSistemaBuilder() *cierreOperacionesSistemaBuilder {
	return &cierreOperacionesSistemaBuilder{
		request: &operaciones.CierreOperacionesSistema{},
	}
}

func NewVerificarComunicacionOperacionesBuilder() *verificarComunicacionOperacionesBuilder {
	return &verificarComunicacionOperacionesBuilder{
		request: &operaciones.VerificarComunicacion{},
	}
}

func NewRegistroEventoSignificativoBuilder() *registroEventoSignificativoBuilder {
	return &registroEventoSignificativoBuilder{
		request: &operaciones.RegistroEventoSignificativo{},
	}
}

func NewConsultaEventoSignificativoBuilder() *consultaEventoSignificativoBuilder {
	return &consultaEventoSignificativoBuilder{
		request: &operaciones.ConsultaEventoSignificativo{},
	}
}

// --- Implementaciones de Builders ---

type registroPuntoVentaBuilder struct {
	request *operaciones.RegistroPuntoVenta
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

func (b *registroPuntoVentaBuilder) WithNombrePuntoVenta(nombre string) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.NombrePuntoVenta = nombre
	return b
}

func (b *registroPuntoVentaBuilder) WithDescripcion(descripcion string) *registroPuntoVentaBuilder {
	b.request.SolicitudRegistroPuntoVenta.Descripcion = descripcion
	return b
}

func (b *registroPuntoVentaBuilder) Build() RegistroPuntoVenta {
	return RegistroPuntoVenta{RequestWrapper: NewRequestWrapper(b.request)}
}

type consultaPuntoVentaBuilder struct {
	request *operaciones.ConsultaPuntoVenta
}

func (b *consultaPuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *consultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *consultaPuntoVentaBuilder) WithCuis(cuis string) *consultaPuntoVentaBuilder {
	b.request.SolicitudConsultaPuntoVenta.Cuis = cuis
	return b
}

func (b *consultaPuntoVentaBuilder) Build() ConsultaPuntoVenta {
	return ConsultaPuntoVenta{RequestWrapper: NewRequestWrapper(b.request)}
}

type cierrePuntoVentaBuilder struct {
	request *operaciones.CierrePuntoVenta
}

func (b *cierrePuntoVentaBuilder) WithCodigoSucursal(codigoSucursal int) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoSucursal = codigoSucursal
	return b
}

func (b *cierrePuntoVentaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *cierrePuntoVentaBuilder) WithCuis(cuis string) *cierrePuntoVentaBuilder {
	b.request.SolicitudCierrePuntoVenta.Cuis = cuis
	return b
}

func (b *cierrePuntoVentaBuilder) Build() CierrePuntoVenta {
	return CierrePuntoVenta{RequestWrapper: NewRequestWrapper(b.request)}
}

type registroPuntoVentaComisionistaBuilder struct {
	request *operaciones.RegistroPuntoVentaComisionista
}

func (b *registroPuntoVentaComisionistaBuilder) WithCodigoSucursal(codigoSucursal int) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.CodigoSucursal = codigoSucursal
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithCuis(cuis string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Cuis = cuis
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithNombrePuntoVenta(nombre string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NombrePuntoVenta = nombre
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithDescripcion(descripcion string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.Descripcion = descripcion
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithNitComisionista(nit int64) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NitComisionista = nit
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithNumeroContrato(numero string) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.NumeroContrato = numero
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithFechaInicio(t time.Time) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.FechaInicio = t
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) WithFechaFin(t time.Time) *registroPuntoVentaComisionistaBuilder {
	b.request.SolicitudPuntoVentaComisionista.FechaFin = t
	return b
}

func (b *registroPuntoVentaComisionistaBuilder) Build() RegistroPuntoVentaComisionista {
	return RegistroPuntoVentaComisionista{RequestWrapper: NewRequestWrapper(b.request)}
}

type cierreOperacionesSistemaBuilder struct {
	request *operaciones.CierreOperacionesSistema
}

func (b *cierreOperacionesSistemaBuilder) WithCodigoSucursal(codigoSucursal int) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoSucursal = codigoSucursal
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *cierreOperacionesSistemaBuilder) WithCuis(cuis string) *cierreOperacionesSistemaBuilder {
	b.request.SolicitudOperaciones.Cuis = cuis
	return b
}

func (b *cierreOperacionesSistemaBuilder) Build() CierreOperacionesSistema {
	return CierreOperacionesSistema{RequestWrapper: NewRequestWrapper(b.request)}
}

type consultaEventoSignificativoBuilder struct {
	request *operaciones.ConsultaEventoSignificativo
}

func (b *consultaEventoSignificativoBuilder) WithCodigoSucursal(codigoSucursal int) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoSucursal = codigoSucursal
	return b
}

func (b *consultaEventoSignificativoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *consultaEventoSignificativoBuilder {
	b.request.SolicitudConsultaEvento.CodigoPuntoVenta = codigoPuntoVenta
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

func (b *consultaEventoSignificativoBuilder) Build() ConsultaEventoSignificativo {
	return ConsultaEventoSignificativo{RequestWrapper: NewRequestWrapper(b.request)}
}

type registroEventoSignificativoBuilder struct {
	request *operaciones.RegistroEventoSignificativo
}

func (b *registroEventoSignificativoBuilder) WithCodigoSucursal(codigoSucursal int) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoSucursal = codigoSucursal
	return b
}

func (b *registroEventoSignificativoBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *registroEventoSignificativoBuilder) WithCuis(cuis string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Cuis = cuis
	return b
}

func (b *registroEventoSignificativoBuilder) WithCufd(cufd string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Cufd = cufd
	return b
}

func (b *registroEventoSignificativoBuilder) WithCufdEvento(cufdEvento string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CufdEvento = cufdEvento
	return b
}

func (b *registroEventoSignificativoBuilder) WithCodigoMotivoEvento(motivo int) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.CodigoMotivoEvento = motivo
	return b
}

func (b *registroEventoSignificativoBuilder) WithDescripcion(desc string) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.Descripcion = desc
	return b
}

func (b *registroEventoSignificativoBuilder) WithFechaInicio(t time.Time) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.FechaHoraInicioEvento = t
	return b
}

func (b *registroEventoSignificativoBuilder) WithFechaFin(t time.Time) *registroEventoSignificativoBuilder {
	b.request.SolicitudEventoSignificativo.FechaHoraFinEvento = t
	return b
}

func (b *registroEventoSignificativoBuilder) Build() RegistroEventoSignificativo {
	return RegistroEventoSignificativo{RequestWrapper: NewRequestWrapper(b.request)}
}

type verificarComunicacionOperacionesBuilder struct {
	request *operaciones.VerificarComunicacion
}

func (b *verificarComunicacionOperacionesBuilder) Build() VerificarComunicacionOperaciones {
	return VerificarComunicacionOperaciones{RequestWrapper: NewRequestWrapper(b.request)}
}
