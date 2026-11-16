package services_test

import (
	"context"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func siatRecepcionComprasIntegrationConfig(t *testing.T) siat.Config {
	if os.Getenv("SIAT_RUN") != "1" {
		t.Skip("prueba de integración real: requiere SIAT_RUN=1")
	}

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

// TestSiatRecepcionCompras_AnulacionCompra anula una compra existente de
// certificación mediante el constructor público. Requiere SIAT_RUN=1 y datos
// reales del ambiente 2; nunca usa un proveedor o autorización ilustrativos.
// CUIS y CUFD deben corresponder a SIAT_PV y SIAT_SUCURSAL (ambos 0 por defecto).
func TestSiatRecepcionCompras_AnulacionCompra(t *testing.T) {
	_ = godotenv.Load(".env")
	if os.Getenv("SIAT_RUN") != "1" {
		t.Skip("defina SIAT_RUN=1 para anular una compra real de certificación")
	}
	config := siatRecepcionComprasIntegrationConfig(t)
	require.Equal(t, siat.AmbientePruebas, config.CodigoAmbiente, "esta prueba sólo admite ambiente 2")
	valores := make(map[string]string)
	for _, nombre := range []string{"SIAT_CUIS", "SIAT_CUFD", "SIAT_COMPRA_NIT", "SIAT_COMPRA_AUT", "SIAT_COMPRA_NUM"} {
		valores[nombre] = strings.TrimSpace(os.Getenv(nombre))
		require.NotEmpty(t, valores[nombre], "falta %s; configure una compra existente del ambiente 2", nombre)
	}
	nit, err := strconv.ParseInt(valores["SIAT_COMPRA_NIT"], 10, 64)
	require.NoError(t, err)
	require.Positive(t, nit)
	numero, err := strconv.ParseInt(valores["SIAT_COMPRA_NUM"], 10, 64)
	require.NoError(t, err)
	require.Positive(t, numero)
	parametro := func(nombre string) int {
		t.Helper()
		valor := os.Getenv(nombre)
		if valor == "" {
			return 0
		}
		n, err := strconv.Atoi(valor)
		require.NoError(t, err, "valor inválido para %s", nombre)
		require.GreaterOrEqual(t, n, 0)
		return n
	}
	client, err := siat.New(config)
	require.NoError(t, err)
	req := models.NewAnulacionCompraBuilder().
		WithCodigoSucursal(parametro("SIAT_SUCURSAL")).
		WithCodigoPuntoVenta(parametro("SIAT_PV")).
		WithCuis(valores["SIAT_CUIS"]).
		WithCufd(valores["SIAT_CUFD"]).
		WithCodAutorizacion(valores["SIAT_COMPRA_AUT"]).
		WithNitProveedor(nit).
		WithNroDuiDim(os.Getenv("SIAT_COMPRA_DIM")).
		WithNroFactura(numero).
		Build()
	ctx, cancel := context.WithTimeout(t.Context(), 60*time.Second)
	defer cancel()
	resp, err := client.RecepcionCompras().AnulacionCompra(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	contenido, err := resp.GetContent()
	require.NoError(t, err, "SIAT devolvió un SOAP Fault")
	resultado := contenido.RespuestaRecepcion
	t.Logf("anulación de compra: estado=%d descripción=%s transacción=%t mensajes=%+v",
		resultado.CodigoEstado, resultado.CodigoDescripcion, resultado.Transaccion, resultado.MensajesList)
	// Una respuesta transportada correctamente o un rechazo 5154 no prueban
	// la anulación. Se exige el resultado específico de compras reportado por
	// el ejecutor de certificación: 8415 / ELIMINADO.
	require.Equal(t, 8415, resultado.CodigoEstado, "la compra no quedó eliminada: %+v", resultado)
	require.Equal(t, "ELIMINADO", strings.ToUpper(strings.TrimSpace(resultado.CodigoDescripcion)))
}
