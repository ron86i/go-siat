// Package facturacion implementa los servicios de facturación para el protocolo SIAT v3.
package facturacion

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ron86i/go-siat/v3/internal/siatpolicy"
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

// Client gestiona las operaciones de emisión y administración de facturas en SIAT v3.
type Client struct {
	soap      *soapclient.Client
	servicio  string
	modalidad int
}

const (
	servicioCompraVenta = "ServicioFacturacionCompraVenta"
	servicioBasico      = siatpolicy.ServicioBasico
)

// NewClient crea el cliente de Compra-Venta (sectores 1, 35 y 41).
func NewClient(c Config) (*Client, error) {
	return newClient(c, servicioCompraVenta, 0)
}

// NewElectronicaClient crea el cliente de facturación electrónica general (modalidad 1).
func NewElectronicaClient(c Config) (*Client, error) {
	return newClient(c, "ServicioFacturacionElectronica", 1)
}

// NewComputarizadaClient crea el cliente de facturación computarizada general (modalidad 2).
func NewComputarizadaClient(c Config) (*Client, error) {
	return newClient(c, "ServicioFacturacionComputarizada", 2)
}

// NewServicioBasicoClient crea el cliente dedicado a los sectores 13 y 40.
func NewServicioBasicoClient(c Config) (*Client, error) {
	return newClient(c, servicioBasico, 0)
}

func NewDocumentoAjusteClient(c Config) (*Client, error) {
	return newClient(c, siatpolicy.ServicioDocumentoAjuste, 0)
}

func NewTelecomunicacionesClient(c Config) (*Client, error) {
	return newClient(c, siatpolicy.ServicioTelecomunicaciones, 0)
}

func NewEntidadFinancieraClient(c Config) (*Client, error) {
	return newClient(c, siatpolicy.ServicioEntidadFinanciera, 0)
}

func NewBoletoAereoClient(c Config) (*Client, error) {
	return newClient(c, siatpolicy.ServicioBoletoAereo, 0)
}

func newClient(c Config, servicio string, modalidad int) (*Client, error) {
	client, err := soapclient.New(soapclient.Config(c), servicio)
	if err != nil {
		return nil, err
	}
	return &Client{soap: client, servicio: servicio, modalidad: modalidad}, nil
}

func (c *Client) validateDestination(sector, modalidad int) error {
	if c == nil {
		return fmt.Errorf("cliente SIAT no inicializado")
	}
	if err := siatpolicy.ValidateDestination(c.servicio, sector, modalidad); err != nil {
		return err
	}
	if c.modalidad != 0 && modalidad != c.modalidad {
		return fmt.Errorf("el sector o modalidad no corresponde al servicio de facturación")
	}
	return nil
}

// VerificarComunicacion comprueba la conectividad con el servicio de facturación.
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
