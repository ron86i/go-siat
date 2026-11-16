// Package operaciones ofrece las operaciones del servicio SOAP FacturacionOperaciones del protocolo SIAT v3.
package operaciones

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ron86i/go-siat/v3/internal/soapclient"
)

// Config agrupa la identidad y los datos de conexión del contribuyente.
type Config struct {
	BaseURL           string
	Token             string
	Nit               int64
	CodigoSistema     string
	CodigoAmbiente    int
	HTTPClient        *http.Client
	MaxRespuestaBytes int64
}

// Fault representa un fallo SOAP devuelto por el SIN.
type Fault = soapclient.Fault

// ErrorComunicacion indica que no se obtuvo una respuesta SOAP utilizable.
type ErrorComunicacion = soapclient.CommunicationError

// Client gestiona operaciones administrativas de contingencias, puntos de venta y cierres mediante FacturacionOperaciones.
type Client struct{ soap *soapclient.Client }

// NewClient crea un cliente para el servicio FacturacionOperaciones.
func NewClient(c Config) (*Client, error) {
	s, err := soapclient.New(soapclient.Config(c), "FacturacionOperaciones")
	if err != nil {
		return nil, err
	}
	return &Client{soap: s}, nil
}

type respuestaXML[T any] struct {
	XMLName   xml.Name
	Resultado *T `xml:",any"`
}

func callOp[T any](ctx context.Context, c *Client, operacion string, solicitud any) (*T, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	req := struct {
		XMLName   xml.Name
		Solicitud any `xml:",any"`
	}{
		XMLName:   xml.Name{Local: "ns:" + operacion},
		Solicitud: solicitud,
	}
	var resp respuestaXML[T]
	if err := c.soap.Call(ctx, req, &resp); err != nil {
		return nil, err
	}
	if resp.XMLName.Local != operacion+"Response" || resp.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de %s", operacion)
	}
	return resp.Resultado, nil
}

// RegistroEventoSignificativo registra una contingencia ya finalizada.
func (c *Client) RegistroEventoSignificativo(ctx context.Context, r SolicitudRegistroEvento) (*RespuestaEventos, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoMotivoEvento <= 0 || r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cufd == "" || r.CufdEvento == "" || r.Cuis == "" || r.Descripcion == "" || r.FechaHoraInicioEvento.IsZero() || r.FechaHoraFinEvento.IsZero() {
		return nil, fmt.Errorf("registro de evento incompleto o inválido")
	}
	inicio := r.FechaHoraInicioEvento.Format("2006-01-02T15:04:05.000")
	fin := r.FechaHoraFinEvento.Format("2006-01-02T15:04:05.000")
	if fin <= inicio {
		return nil, fmt.Errorf("la fecha de fin del evento debe ser posterior al inicio")
	}
	cfg := c.soap.Config()
	payload := solicitudRegistroEventoXML{
		CodigoAmbiente:        cfg.CodigoAmbiente,
		CodigoMotivoEvento:    r.CodigoMotivoEvento,
		CodigoPuntoVenta:      r.CodigoPuntoVenta,
		CodigoSistema:         cfg.CodigoSistema,
		CodigoSucursal:        r.CodigoSucursal,
		Cufd:                  r.Cufd,
		CufdEvento:            r.CufdEvento,
		Cuis:                  r.Cuis,
		Descripcion:           r.Descripcion,
		FechaHoraFinEvento:    fin,
		FechaHoraInicioEvento: inicio,
		Nit:                   cfg.Nit,
	}
	return callOp[RespuestaEventos](ctx, c, "registroEventoSignificativo", payload)
}

// ConsultaEventoSignificativo consulta eventos por fecha y CUFD.
func (c *Client) ConsultaEventoSignificativo(ctx context.Context, r SolicitudConsultaEvento) (*RespuestaEventos, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cufd == "" || r.Cuis == "" || r.FechaEvento.IsZero() {
		return nil, fmt.Errorf("consulta de evento incompleta o inválida")
	}
	cfg := c.soap.Config()
	payload := solicitudConsultaEventoXML{
		CodigoAmbiente:   cfg.CodigoAmbiente,
		CodigoPuntoVenta: r.CodigoPuntoVenta,
		CodigoSistema:    cfg.CodigoSistema,
		CodigoSucursal:   r.CodigoSucursal,
		Cufd:             r.Cufd,
		Cuis:             r.Cuis,
		FechaEvento:      r.FechaEvento.Format("2006-01-02T15:04:05.000"),
		Nit:              cfg.Nit,
	}
	return callOp[RespuestaEventos](ctx, c, "consultaEventoSignificativo", payload)
}

// RegistroPuntoVenta solicita el registro de un nuevo punto de venta ante el SIN.
func (c *Client) RegistroPuntoVenta(ctx context.Context, r SolicitudRegistroPuntoVenta) (*RespuestaRegistroPuntoVenta, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if (r.CodigoModalidad != 1 && r.CodigoModalidad != 2) || r.CodigoSucursal < 0 || r.CodigoTipoPuntoVenta <= 0 || r.Cuis == "" || r.Descripcion == "" || r.NombrePuntoVenta == "" {
		return nil, fmt.Errorf("solicitud de registro de punto de venta incompleta o inválida")
	}
	cfg := c.soap.Config()
	payload := solicitudRegistroPuntoVentaXML{
		CodigoAmbiente:       cfg.CodigoAmbiente,
		CodigoModalidad:      r.CodigoModalidad,
		CodigoSistema:        cfg.CodigoSistema,
		CodigoSucursal:       r.CodigoSucursal,
		CodigoTipoPuntoVenta: r.CodigoTipoPuntoVenta,
		Cuis:                 r.Cuis,
		Descripcion:          r.Descripcion,
		Nit:                  cfg.Nit,
		NombrePuntoVenta:     r.NombrePuntoVenta,
	}
	return callOp[RespuestaRegistroPuntoVenta](ctx, c, "registroPuntoVenta", payload)
}

// ConsultaPuntoVenta consulta los puntos de venta vigentes de una sucursal.
func (c *Client) ConsultaPuntoVenta(ctx context.Context, r SolicitudConsultaPuntoVenta) (*RespuestaConsultaPuntoVenta, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoSucursal < 0 || r.Cuis == "" {
		return nil, fmt.Errorf("solicitud de consulta de punto de venta incompleta o inválida")
	}
	cfg := c.soap.Config()
	payload := solicitudConsultaPuntoVentaXML{
		CodigoAmbiente: cfg.CodigoAmbiente,
		CodigoSistema:  cfg.CodigoSistema,
		CodigoSucursal: r.CodigoSucursal,
		Cuis:           r.Cuis,
		Nit:            cfg.Nit,
	}
	return callOp[RespuestaConsultaPuntoVenta](ctx, c, "consultaPuntoVenta", payload)
}

// CierrePuntoVenta solicita el cierre de un punto de venta identificado.
func (c *Client) CierrePuntoVenta(ctx context.Context, r SolicitudCierrePuntoVenta) (*RespuestaCierrePuntoVenta, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoSucursal < 0 || r.CodigoPuntoVenta < 0 || r.Cuis == "" {
		return nil, fmt.Errorf("solicitud de cierre de punto de venta incompleta o inválida")
	}
	cfg := c.soap.Config()
	payload := solicitudCierrePuntoVentaXML{
		CodigoAmbiente:   cfg.CodigoAmbiente,
		CodigoPuntoVenta: r.CodigoPuntoVenta,
		CodigoSistema:    cfg.CodigoSistema,
		CodigoSucursal:   r.CodigoSucursal,
		Cuis:             r.Cuis,
		Nit:              cfg.Nit,
	}
	return callOp[RespuestaCierrePuntoVenta](ctx, c, "cierrePuntoVenta", payload)
}

// RegistroPuntoVentaComisionista solicita el registro de un punto de venta comisionista.
func (c *Client) RegistroPuntoVentaComisionista(ctx context.Context, r SolicitudRegistroComisionista) (*RespuestaRegistroComisionista, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if (r.CodigoModalidad != 1 && r.CodigoModalidad != 2) || r.CodigoSucursal < 0 || r.Cuis == "" || r.Descripcion == "" || r.FechaInicio.IsZero() || r.FechaFin.IsZero() || r.NitComisionista <= 0 || r.NombrePuntoVenta == "" || r.NumeroContrato == "" {
		return nil, fmt.Errorf("solicitud de comisionista incompleta o inválida")
	}
	inicio := r.FechaInicio.Format("2006-01-02T15:04:05.000")
	fin := r.FechaFin.Format("2006-01-02T15:04:05.000")
	if fin <= inicio {
		return nil, fmt.Errorf("la fecha de fin debe ser posterior al inicio")
	}
	cfg := c.soap.Config()
	payload := solicitudRegistroComisionistaXML{
		CodigoAmbiente:   cfg.CodigoAmbiente,
		CodigoModalidad:  r.CodigoModalidad,
		CodigoSistema:    cfg.CodigoSistema,
		CodigoSucursal:   r.CodigoSucursal,
		Cuis:             r.Cuis,
		Descripcion:      r.Descripcion,
		FechaFin:         fin,
		FechaInicio:      inicio,
		Nit:              cfg.Nit,
		NitComisionista:  r.NitComisionista,
		NombrePuntoVenta: r.NombrePuntoVenta,
		NumeroContrato:   r.NumeroContrato,
	}
	return callOp[RespuestaRegistroComisionista](ctx, c, "registroPuntoVentaComisionista", payload)
}

// CierreOperacionesSistema solicita el cierre de operaciones del sistema.
func (c *Client) CierreOperacionesSistema(ctx context.Context, r SolicitudCierreSistema) (*RespuestaCierreSistema, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if (r.CodigoModalidad != 1 && r.CodigoModalidad != 2) || r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cuis == "" {
		return nil, fmt.Errorf("solicitud de cierre de sistema incompleta o inválida")
	}
	cfg := c.soap.Config()
	payload := solicitudCierreSistemaXML{
		CodigoAmbiente:   cfg.CodigoAmbiente,
		CodigoModalidad:  r.CodigoModalidad,
		CodigoPuntoVenta: r.CodigoPuntoVenta,
		CodigoSistema:    cfg.CodigoSistema,
		CodigoSucursal:   r.CodigoSucursal,
		Cuis:             r.Cuis,
		Nit:              cfg.Nit,
	}
	return callOp[RespuestaCierreSistema](ctx, c, "cierreOperacionesSistema", payload)
}

// VerificarComunicacion comprueba la conectividad con el servicio de operaciones.
func (c *Client) VerificarComunicacion(ctx context.Context) (*RespuestaComunicacion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	req := struct {
		XMLName xml.Name `xml:"ns:verificarComunicacion"`
	}{}
	var resp struct {
		XMLName   xml.Name               `xml:"verificarComunicacionResponse"`
		Resultado *RespuestaComunicacion `xml:"return"`
	}
	if err := c.soap.Call(ctx, req, &resp); err != nil {
		return nil, err
	}
	if resp.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de comunicación")
	}
	return resp.Resultado, nil
}
