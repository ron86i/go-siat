// Package codigos ofrece las operaciones del servicio SOAP FacturacionCodigos del protocolo SIAT v3.
package codigos

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ron86i/go-siat/v3/internal/soapclient"
)

// Config contiene los datos de conexión e identidad del contribuyente.
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

// Client se conecta al servicio SOAP FacturacionCodigos.
type Client struct {
	soap      *soapclient.Client
	modalidad int
}

// NewClient crea un cliente para el servicio FacturacionCodigos del SIAT v3.
func NewClient(c Config) (*Client, error) {
	s, err := soapclient.New(soapclient.Config(c), "FacturacionCodigos")
	if err != nil {
		return nil, err
	}
	return &Client{soap: s}, nil
}

// WithModalidad define una modalidad por defecto (1: Electrónica, 2: Computarizada) para las solicitudes de este cliente.
func (c *Client) WithModalidad(modalidad int) *Client {
	if c == nil {
		return nil
	}
	return &Client{soap: c.soap, modalidad: modalidad}
}

// SolicitarCUIS solicita un Código Único de Inicio de Sistemas (CUIS).
func (c *Client) SolicitarCUIS(ctx context.Context, r SolicitudCUIS) (*RespuestaCUIS, error) {
	modalidad := r.CodigoModalidad
	if modalidad == 0 && c != nil {
		modalidad = c.modalidad
	}
	var response struct {
		XMLName   xml.Name       `xml:"cuisResponse"`
		Resultado *RespuestaCUIS `xml:"RespuestaCuis"`
	}
	if err := c.send(ctx, "cuis", modalidad, r.CodigoPuntoVenta, r.CodigoSucursal, nil, &response); err != nil {
		return nil, err
	}
	if response.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado CUIS")
	}
	return response.Resultado, nil
}

// SolicitarCUFD solicita un Código Único de Facturación Diaria (CUFD).
func (c *Client) SolicitarCUFD(ctx context.Context, r SolicitudCUFD) (*RespuestaCUFD, error) {
	modalidad := r.CodigoModalidad
	if modalidad == 0 && c != nil {
		modalidad = c.modalidad
	}
	var response struct {
		XMLName   xml.Name       `xml:"cufdResponse"`
		Resultado *RespuestaCUFD `xml:"RespuestaCufd"`
	}
	if err := c.send(ctx, "cufd", modalidad, r.CodigoPuntoVenta, r.CodigoSucursal, &r.Cuis, &response); err != nil {
		return nil, err
	}
	if response.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado CUFD")
	}
	return response.Resultado, nil
}

func (c *Client) send(ctx context.Context, operacion string, modalidad int, punto *int, sucursal int, cuis *string, result any) error {
	if c == nil || c.soap == nil {
		return fmt.Errorf("cliente SIAT no inicializado")
	}
	if modalidad != 1 && modalidad != 2 {
		return fmt.Errorf("solicitud de códigos incompleta: CodigoModalidad debe ser 1 (Electrónica) o 2 (Computarizada)")
	}
	if sucursal < 0 || (punto != nil && *punto < 0) {
		return fmt.Errorf("solicitud de códigos inválida: sucursal o punto de venta no pueden ser negativos")
	}
	if cuis != nil && *cuis == "" {
		return fmt.Errorf("solicitud de códigos incompleta: CUIS requerido para CUFD")
	}

	cfg := c.soap.Config()
	dato := solicitudSoap{
		CodigoAmbiente:   cfg.CodigoAmbiente,
		CodigoModalidad:  modalidad,
		CodigoPuntoVenta: punto,
		CodigoSistema:    cfg.CodigoSistema,
		CodigoSucursal:   sucursal,
		Cuis:             cuis,
		Nit:              cfg.Nit,
	}
	if operacion == "cuis" {
		dato.XMLName.Local = "SolicitudCuis"
	} else {
		dato.XMLName.Local = "SolicitudCufd"
	}
	req := struct {
		XMLName   xml.Name
		Solicitud solicitudSoap
	}{
		XMLName:   xml.Name{Local: "ns:" + operacion},
		Solicitud: dato,
	}
	return c.soap.Call(ctx, req, result)
}
