// Package soapclient implements the shared SIAT SOAP transport.
package soapclient

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/v3/internal/siatpolicy"
)

// Config contains the connection details shared by SIAT operations.
type Config struct {
	BaseURL           string
	Token             string
	Nit               int64
	CodigoSistema     string
	CodigoAmbiente    int
	HTTPClient        *http.Client
	MaxRespuestaBytes int64
}

// Client sends requests to one SIAT service endpoint.
type Client struct {
	config   Config
	endpoint string
	http     *http.Client
}

// New validates connection details without sending a request.
func New(c Config, service string) (*Client, error) {
	baseURL, err := ResolveBaseURL(c.BaseURL, c.CodigoAmbiente)
	if err != nil {
		return nil, err
	}
	if !siatpolicy.ValidService(service) {
		return nil, fmt.Errorf("servicio SIAT desconocido: %q", service)
	}
	if strings.TrimSpace(c.Token) == "" || c.Nit <= 0 || strings.TrimSpace(c.CodigoSistema) == "" || (c.CodigoAmbiente != 1 && c.CodigoAmbiente != 2) {
		return nil, fmt.Errorf("configuración SIAT incompleta o inválida")
	}
	if c.MaxRespuestaBytes < 0 || c.MaxRespuestaBytes == math.MaxInt64 {
		return nil, fmt.Errorf("límite de respuesta inválido")
	}
	if c.MaxRespuestaBytes == 0 {
		c.MaxRespuestaBytes = 64 << 20
	}
	hc := c.HTTPClient
	if hc == nil {
		hc = &http.Client{Timeout: 45 * time.Second}
	}
	client := *hc
	client.CheckRedirect = func(_ *http.Request, _ []*http.Request) error { return http.ErrUseLastResponse }
	return &Client{config: c, endpoint: baseURL + "/" + service, http: &client}, nil
}

// ResolveBaseURL validates and normalizes the configured SIAT 3.0 environment root.
func ResolveBaseURL(baseURL string, ambiente int) (string, error) {
	return siatpolicy.ResolveBaseURL(baseURL, ambiente)
}

// Config returns a copy of the connection configuration.
func (c *Client) Config() Config { return c.config }

// Fault is a SOAP 1.1 fault returned by SIAT.
type Fault struct {
	Code    string `xml:"faultcode"`
	Message string `xml:"faultstring"`
	Detail  string `xml:"detail"`
}

func (f *Fault) Error() string { return "SOAP fault: " + f.Code + ": " + f.Message }

// CommunicationError indica que no se obtuvo una respuesta SOAP utilizable.
// Una operación enviada puede haber sido procesada por el SIN.
type CommunicationError struct {
	Etapa      string
	StatusCode int
	Err        error
}

func (e *CommunicationError) Error() string {
	if e.StatusCode != 0 {
		return fmt.Sprintf("comunicación SIAT: %s (HTTP %d)", e.Etapa, e.StatusCode)
	}
	return fmt.Sprintf("comunicación SIAT: %s: %v", e.Etapa, e.Err)
}

// Unwrap conserva la causa para errors.Is y errors.As.
func (e *CommunicationError) Unwrap() error { return e.Err }

// Call sends one SOAP request. It does not retry failed requests.
func (c *Client) Call(ctx context.Context, payload any, result any) error {
	envelope := struct {
		XMLName xml.Name `xml:"soapenv:Envelope"`
		Soap    string   `xml:"xmlns:soapenv,attr"`
		NS      string   `xml:"xmlns:ns,attr"`
		Body    struct {
			Content any `xml:",any"`
		} `xml:"soapenv:Body"`
	}{Soap: "http://schemas.xmlsoap.org/soap/envelope/", NS: "https://siat.impuestos.gob.bo/"}
	envelope.Body.Content = payload
	data, err := xml.Marshal(envelope)
	if err != nil {
		return fmt.Errorf("serializar solicitud SOAP: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(append([]byte(xml.Header), data...)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("User-Agent", "go-siat")
	req.Header.Set("apiKey", "TokenApi "+c.config.Token)
	resp, err := c.http.Do(req)
	if err != nil {
		return &CommunicationError{Etapa: "solicitud", Err: err}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, c.config.MaxRespuestaBytes+1))
	if err != nil {
		return &CommunicationError{Etapa: "lectura", Err: err}
	}
	if int64(len(body)) > c.config.MaxRespuestaBytes {
		return &CommunicationError{Etapa: "lectura", Err: fmt.Errorf("respuesta excede %d bytes", c.config.MaxRespuestaBytes)}
	}
	var env struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			Fault   *Fault `xml:"Fault"`
			Content []byte `xml:",innerxml"`
		} `xml:"Body"`
	}
	err = xml.Unmarshal(body, &env)
	if err == nil && env.Body.Fault != nil {
		return env.Body.Fault
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &CommunicationError{Etapa: "http", StatusCode: resp.StatusCode}
	}
	if err != nil {
		return &CommunicationError{Etapa: "respuesta SOAP", Err: err}
	}
	if err = xml.Unmarshal(env.Body.Content, result); err != nil {
		return &CommunicationError{Etapa: "contenido SOAP", Err: err}
	}
	return nil
}
