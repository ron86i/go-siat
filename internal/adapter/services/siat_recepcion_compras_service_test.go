package services_test

import (
	"context"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/stretchr/testify/require"
)

func siatRecepcionComprasIntegrationConfig(t *testing.T) siat.Config {
	t.Helper()

	// Las variables exportadas tienen prioridad; .env solo completa las ausentes.
	_ = godotenv.Load(".env")

	required := map[string]string{
		"SIAT_TOKEN":           os.Getenv("SIAT_TOKEN"),
		"SIAT_NIT":             os.Getenv("SIAT_NIT"),
		"SIAT_CODIGO_SISTEMA":  os.Getenv("SIAT_CODIGO_SISTEMA"),
		"SIAT_CODIGO_AMBIENTE": os.Getenv("SIAT_CODIGO_AMBIENTE"),
		"SIAT_URL":             os.Getenv("SIAT_URL"),
	}
	missing := make([]string, 0, len(required))
	for name, value := range required {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		t.Skipf("prueba de integración omitida; faltan variables de entorno: %s", strings.Join(missing, ", "))
	}

	nit, err := strconv.ParseInt(os.Getenv("SIAT_NIT"), 10, 64)
	require.NoError(t, err, "SIAT_NIT debe ser un entero válido")
	ambiente, err := strconv.Atoi(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	require.NoError(t, err, "SIAT_CODIGO_AMBIENTE debe ser un entero válido")

	return siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: ambiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     newXMLTraceHTTPClient(t),
	}
}

func TestSiatRecepcionCompras_VerificarComunicacion(t *testing.T) {
	config := siatRecepcionComprasIntegrationConfig(t)
	client, err := siat.New(config)
	require.NoError(t, err, "no se pudo crear el cliente SIAT")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := client.RecepcionCompras().VerificarComunicacion(
		ctx,
		models.NewVerificarComunicacionRecepcionComprasBuilder().Build(),
	)
	require.NoError(t, err, "falló la comunicación con el SIAT")
	require.NotNil(t, response)

	content, err := response.GetContent()
	require.NoError(t, err, "el SIAT devolvió un SOAP Fault")

	t.Logf("respuesta de recepción de compras: %+v", content)
	require.True(t, content.Return.Transaccion,
		"el SIAT rechazó la verificación: %+v", content.Return.MensajesList)
}
