// Package compras implementa el servicio ServicioRecepcionCompras del protocolo SIAT v3.
package compras

import (
	"context"
	"encoding/base64"
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

// Fault representa un error SOAP devuelto por el SIN.
type Fault = soapclient.Fault

// ErrorComunicacion indica que no se obtuvo una respuesta SOAP utilizable.
type ErrorComunicacion = soapclient.CommunicationError

// Client gestiona las operaciones de recepción, confirmación, consulta y anulación de compras.
type Client struct{ soap *soapclient.Client }

// NewClient crea un cliente para el servicio ServicioRecepcionCompras.
func NewClient(c Config) (*Client, error) {
	s, err := soapclient.New(soapclient.Config(c), "ServicioRecepcionCompras")
	if err != nil {
		return nil, err
	}
	return &Client{soap: s}, nil
}

func (c *Client) branchData(d DatosSucursal) (baseXML, error) {
	if c == nil || c.soap == nil {
		return baseXML{}, fmt.Errorf("cliente SIAT no inicializado")
	}
	if d.CodigoSucursal < 0 || d.CodigoPuntoVenta < 0 || d.Cufd == "" || d.Cuis == "" {
		return baseXML{}, fmt.Errorf("datos de sucursal incompletos o inválidos")
	}
	cfg := c.soap.Config()
	return baseXML{
		CodigoAmbiente:   cfg.CodigoAmbiente,
		CodigoPuntoVenta: d.CodigoPuntoVenta,
		CodigoSistema:    cfg.CodigoSistema,
		CodigoSucursal:   d.CodigoSucursal,
		Cufd:             d.Cufd,
		Cuis:             d.Cuis,
		Nit:              cfg.Nit,
	}, nil
}

type respuestaXML[T any] struct {
	XMLName   xml.Name
	Resultado *T `xml:"RespuestaServicioFacturacion"`
}

func send[T any](ctx context.Context, c *Client, operacion string, solicitud any) (*T, error) {
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

func (c *Client) sendPackage(ctx context.Context, operacion, elemento string, r SolicitudPaquete) (*RespuestaRecepcion, error) {
	base, err := c.branchData(r.DatosSucursal)
	if err != nil {
		return nil, err
	}
	if len(r.Archivo) == 0 || r.CantidadFacturas <= 0 || r.FechaEnvio.IsZero() || r.Gestion <= 0 || r.Periodo < 1 || r.Periodo > 12 || r.HashArchivo == "" {
		return nil, fmt.Errorf("solicitud de paquete de compras incompleta o inválida")
	}
	dato := datosPaqueteXML{
		XMLName:          xml.Name{Local: elemento},
		baseXML:          base,
		Archivo:          base64.StdEncoding.EncodeToString(r.Archivo),
		CantidadFacturas: r.CantidadFacturas,
		FechaEnvio:       r.FechaEnvio.Format("2006-01-02T15:04:05.000"),
		Gestion:          r.Gestion,
		HashArchivo:      r.HashArchivo,
		Periodo:          r.Periodo,
	}
	return send[RespuestaRecepcion](ctx, c, operacion, dato)
}

// RecepcionPaqueteCompras envía un paquete comprimido de compras preparado.
func (c *Client) RecepcionPaqueteCompras(ctx context.Context, r SolicitudPaquete) (*RespuestaRecepcion, error) {
	return c.sendPackage(ctx, "recepcionPaqueteCompras", "SolicitudRecepcionCompras", r)
}

// ConfirmacionCompras envía un paquete de confirmaciones previamente preparado.
func (c *Client) ConfirmacionCompras(ctx context.Context, r SolicitudPaquete) (*RespuestaRecepcion, error) {
	return c.sendPackage(ctx, "confirmacionCompras", "SolicitudConfirmacionCompras", r)
}

// ValidacionRecepcionPaqueteCompras consulta el resultado de validación de un paquete de compras.
func (c *Client) ValidacionRecepcionPaqueteCompras(ctx context.Context, r SolicitudValidacion) (*RespuestaRecepcion, error) {
	base, err := c.branchData(r.DatosSucursal)
	if err != nil {
		return nil, err
	}
	if r.CodigoRecepcion == "" {
		return nil, fmt.Errorf("código de recepción obligatorio")
	}
	dato := datosValidacionXML{
		XMLName:         xml.Name{Local: "SolicitudValidacionRecepcionCompras"},
		baseXML:         base,
		CodigoRecepcion: r.CodigoRecepcion,
	}
	return send[RespuestaRecepcion](ctx, c, "validacionRecepcionPaqueteCompras", dato)
}

// AnulacionCompra solicita la anulación de un registro de compra.
func (c *Client) AnulacionCompra(ctx context.Context, r SolicitudAnulacion) (*RespuestaRecepcion, error) {
	base, err := c.branchData(r.DatosSucursal)
	if err != nil {
		return nil, err
	}
	if r.NitProveedor <= 0 || r.NroFactura < 0 {
		return nil, fmt.Errorf("identificación de compra inválida")
	}
	dato := datosAnulacionXML{
		XMLName:         xml.Name{Local: "SolicitudAnulacionCompra"},
		baseXML:         base,
		CodAutorizacion: r.CodAutorizacion,
		NitProveedor:    r.NitProveedor,
		NroDuiDim:       r.NroDuiDim,
		NroFactura:      r.NroFactura,
	}
	return send[RespuestaRecepcion](ctx, c, "anulacionCompra", dato)
}

// ConsultaCompras descarga y decodifica el archivo de compras registradas para la fecha solicitada.
func (c *Client) ConsultaCompras(ctx context.Context, r SolicitudConsulta) (*RespuestaConsulta, error) {
	base, err := c.branchData(r.DatosSucursal)
	if err != nil {
		return nil, err
	}
	if r.Fecha.IsZero() {
		return nil, fmt.Errorf("fecha de consulta obligatoria")
	}
	dato := datosConsultaXML{
		XMLName: xml.Name{Local: "SolicitudConsultaCompras"},
		baseXML: base,
		Fecha:   r.Fecha.Format("2006-01-02T15:04:05.000"),
	}
	type respuesta struct {
		Archivo     string    `xml:"archivo"`
		Mensajes    []Mensaje `xml:"mensajesList"`
		Transaccion *bool     `xml:"transaccion"`
	}
	result, err := send[respuesta](ctx, c, "consultaCompras", dato)
	if err != nil {
		return nil, err
	}
	archivo, err := base64.StdEncoding.DecodeString(result.Archivo)
	if err != nil {
		return nil, fmt.Errorf("archivo de compras devuelto no es base64 válido: %w", err)
	}
	return &RespuestaConsulta{
		Archivo:     archivo,
		Mensajes:    result.Mensajes,
		Transaccion: result.Transaccion,
	}, nil
}

// VerificarComunicacion consulta la conectividad con el servicio de compras.
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
