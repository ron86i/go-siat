package services_test

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/codigos"
	"github.com/ron86i/go-siat/v3/internal/core/ports"
	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/stretchr/testify/require"
)

type codigosTestOptions struct {
	config                          siat.Config
	modalidad, sucursal, puntoVenta int
	veces                           int
	cuis                            string
}

// Configuración del ejecutor de pruebas; no añade validaciones al SDK.
func readCodigosTestOptions(getenv func(string) string) (codigosTestOptions, error) {
	var options codigosTestOptions
	get := func(name string) string { return strings.TrimSpace(getenv(name)) }
	for _, name := range []string{"SIAT_TOKEN", "SIAT_NIT", "SIAT_CODIGO_SISTEMA", "SIAT_CODIGO_AMBIENTE", "SIAT_CODIGO_MODALIDAD", "SIAT_URL"} {
		if get(name) == "" {
			return options, fmt.Errorf("falta %s", name)
		}
	}
	nit, err := strconv.ParseInt(get("SIAT_NIT"), 10, 64)
	if err != nil || nit <= 0 {
		return options, fmt.Errorf("SIAT_NIT debe ser un entero positivo")
	}
	if get("SIAT_CODIGO_AMBIENTE") != "2" {
		return options, fmt.Errorf("SIAT_CODIGO_AMBIENTE debe ser 2 para estas pruebas")
	}
	options.modalidad, err = strconv.Atoi(get("SIAT_CODIGO_MODALIDAD"))
	if err != nil || (options.modalidad != siat.ModalidadElectronica && options.modalidad != siat.ModalidadComputarizada) {
		return options, fmt.Errorf("SIAT_CODIGO_MODALIDAD debe ser 1 o 2")
	}
	endpoint, err := url.Parse(get("SIAT_URL"))
	if err != nil || endpoint.Host == "" || (endpoint.Scheme != "https" && endpoint.Scheme != "http") ||
		!strings.HasSuffix(strings.TrimRight(endpoint.Path, "/"), "/sfd") || endpoint.RawQuery != "" || endpoint.Fragment != "" {
		return options, fmt.Errorf("SIAT_URL debe ser una dirección base terminada en /sfd")
	}
	number := func(short, alias string, fallback int) (int, error) {
		name, value := short, get(short)
		if value == "" && alias != "" {
			name, value = alias, get(alias)
		}
		if value == "" {
			return fallback, nil
		}
		n, err := strconv.Atoi(value)
		if err != nil || n < 0 {
			return 0, fmt.Errorf("%s debe ser un entero no negativo", name)
		}
		return n, nil
	}
	if options.sucursal, err = number("SIAT_SUCURSAL", "SIAT_CODIGO_SUCURSAL", 0); err != nil {
		return options, err
	}
	if options.puntoVenta, err = number("SIAT_PV", "SIAT_CODIGO_PUNTO_VENTA", 1); err != nil {
		return options, err
	}
	if options.veces, err = number("SIAT_VECES", "", 10); err != nil {
		return options, err
	}
	if options.veces == 0 {
		return options, fmt.Errorf("SIAT_VECES debe ser mayor a cero")
	}
	options.cuis = get("SIAT_CUIS")
	options.config = siat.Config{
		Token: get("SIAT_TOKEN"), Nit: nit, CodigoSistema: get("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas, BaseURL: get("SIAT_URL"),
	}
	return options, nil
}

func newCodigosIntegration(t *testing.T) (*siat.SiatServices, codigosTestOptions) {
	t.Helper()
	if os.Getenv("SIAT_RUN") != "1" {
		t.Skip("prueba real de códigos: requiere SIAT_RUN=1")
	}
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		t.Fatalf("no se pudo cargar .env: %v", err)
	}
	options, err := readCodigosTestOptions(os.Getenv)
	require.NoError(t, err)
	options.config.HTTPClient = &http.Client{
		Timeout: 25 * time.Second,
		Transport: &xmlTraceTransport{
			t: t, base: http.DefaultTransport,
			outputDir: filepath.Join(".siat-traces", t.Name()),
		},
	}
	client, err := siat.New(options.config)
	require.NoError(t, err)
	t.Logf("Códigos SIAT: sucursal=%d puntoVenta=%d modalidad=%d; XML en .siat-traces/%s", options.sucursal, options.puntoVenta, options.modalidad, t.Name())
	return client, options
}

func validateCodigosCuisResponse(result codigos.RespuestaCuis) error {
	if !result.Transaccion {
		if len(result.MensajesList) == 0 {
			return fmt.Errorf("CUIS rechazado sin mensajes")
		}
		for _, message := range result.MensajesList {
			if message.Codigo != 980 {
				return fmt.Errorf("CUIS rechazado: mensajes=%+v", result.MensajesList)
			}
		}
	}
	if strings.TrimSpace(result.Codigo) == "" {
		return fmt.Errorf("SIAT no devolvió un código CUIS utilizable: transaccion=%t mensajes=%+v", result.Transaccion, result.MensajesList)
	}
	return nil
}

func requestCodigosCuis(t *testing.T, service ports.SiatCodigosService, options codigosTestOptions) string {
	t.Helper()
	request := models.NewCuisBuilder().WithCodigoModalidad(options.modalidad).
		WithCodigoSucursal(options.sucursal).WithCodigoPuntoVenta(options.puntoVenta).Build()
	response, err := service.SolicitudCuis(t.Context(), request)
	require.NoError(t, err)
	require.NotNil(t, response)
	content, err := response.GetContent()
	require.NoError(t, err)
	require.NoError(t, validateCodigosCuisResponse(content.RespuestaCuis))
	return content.RespuestaCuis.Codigo
}

func codigosCuisForDependentTest(t *testing.T, service ports.SiatCodigosService, options codigosTestOptions) string {
	t.Helper()
	if options.cuis != "" {
		return options.cuis
	}
	return requestCodigosCuis(t, service, options)
}

func requestCodigosCufd(t *testing.T, service ports.SiatCodigosService, options codigosTestOptions, cuis string) {
	t.Helper()
	request := models.NewCufdBuilder().WithCodigoModalidad(options.modalidad).
		WithCodigoSucursal(options.sucursal).WithCodigoPuntoVenta(options.puntoVenta).WithCuis(cuis).Build()
	response, err := service.SolicitudCufd(t.Context(), request)
	require.NoError(t, err)
	require.NotNil(t, response)
	content, err := response.GetContent()
	require.NoError(t, err)
	result := content.RespuestaCufd
	require.True(t, result.Transaccion, "CUFD rechazado: mensajes=%+v", result.MensajesList)
	require.NotEmpty(t, result.Codigo, "SIAT no devolvió CUFD")
	require.NotEmpty(t, result.CodigoControl, "SIAT no devolvió código de control")
	require.False(t, result.FechaVigencia.IsZero(), "SIAT no devolvió vigencia de CUFD")
}

func runCodigosCufdIterations(t *testing.T, service ports.SiatCodigosService, options codigosTestOptions) {
	t.Helper()
	cuis := codigosCuisForDependentTest(t, service, options)
	for i := range options.veces {
		if !t.Run(fmt.Sprintf("Iteracion_%d", i+1), func(t *testing.T) {
			requestCodigosCufd(t, service, options, cuis)
		}) {
			break // No repetir peticiones cuando SIAT ya rechazó una iteración.
		}
	}
}
