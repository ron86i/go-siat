package facturas_test

import (
	"context"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
)

// TestContext agrupa los elementos comunes necesarios para ejecutar pruebas de integración con el SIAT.
type TestContext struct {
	Client     *siat.SiatServices
	Config     siat.Config
	Nit        int64
	Ambiente   int
	Modalidad  int
	PuntoVenta int
	Sucursal   int
	Sistema    string
}

// setupTestContext inicializa el entorno de pruebas, cargando variables de entorno y configurando el cliente SIAT.
func setupTestContext(t *testing.T, modalidad int) *TestContext {
	// Intentar cargar .env desde la raíz del proyecto (3 niveles arriba de pkg/models/facturas)
	_ = godotenv.Load(".env")

	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	conf := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
	sistema := os.Getenv("SIAT_CODIGO_SISTEMA")

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxConnsPerHost:     20,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
			Proxy:               http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		Timeout: 30 * time.Second,
	}
	siatClient, err := siat.New(os.Getenv("SIAT_URL"), client)
	if err != nil {
		t.Fatalf("error al crear el cliente SIAT: %v", err)
	}

	return &TestContext{
		Client:     siatClient,
		Config:     conf,
		Nit:        nit,
		Ambiente:   codAmbiente,
		Modalidad:  modalidad,
		PuntoVenta: 0,
		Sucursal:   0,
		Sistema:    sistema,
	}
}

// GetCuis solicita un código CUIS al SIAT simplificando el flujo de los tests.
func (tc *TestContext) GetCuis(t *testing.T) string {
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		Build()

	resp, err := tc.Client.Codigos().SolicitudCuis(context.Background(), tc.Config, cuisReq)
	if err != nil {
		t.Fatalf("error al obtener CUIS: %v", err)
	}
	if resp.Body.Content.RespuestaCuis.Codigo == "" {
		t.Fatalf("CUIS obtenido está vacío: %+v", resp.Body.Content.RespuestaCuis)
	}
	return resp.Body.Content.RespuestaCuis.Codigo
}

// GetCufd solicita un código CUFD y su código de control al SIAT.
func (tc *TestContext) GetCufd(t *testing.T, cuis string) (string, string) {
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCuis(cuis).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		Build()

	resp, err := tc.Client.Codigos().SolicitudCufd(context.Background(), tc.Config, cufdReq)
	if err != nil {
		t.Fatalf("error al obtener CUFD: %v", err)
	}
	return resp.Body.Content.RespuestaCufd.Codigo, resp.Body.Content.RespuestaCufd.CodigoControl
}
