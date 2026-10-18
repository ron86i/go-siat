package services_test

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/stretchr/testify/require"
)

// Las pruebas reales de códigos requieren SIAT_RUN=1 y ambiente 2.
// SIAT_PV (por defecto 1) y SIAT_SUCURSAL (por defecto 0) se comparten
// entre las solicitudes de CUIS/CUFD. Los XML se guardan por test.

func TestSolicitudCuis(t *testing.T) {
	client, options := newCodigosIntegration(t)
	requestCodigosCuis(t, client.Codigos(), options)
}

func TestSolicitudCufd(t *testing.T) {
	client, options := newCodigosIntegration(t)
	cuis := codigosCuisForDependentTest(t, client.Codigos(), options)
	requestCodigosCufd(t, client.Codigos(), options, cuis)
}

// TestSolicitudCufdAll solicita SIAT_VECES CUFD (10 por defecto), con un
// solo CUIS para todas las repeticiones. SIAT_CUIS permite reutilizar uno vigente.
func TestSolicitudCufdAll(t *testing.T) {
	client, options := newCodigosIntegration(t)
	runCodigosCufdIterations(t, client.Codigos(), options)
}

func TestSolicitudCufdMasivo(t *testing.T) {
	client, options := newCodigosIntegration(t)
	cuis := codigosCuisForDependentTest(t, client.Codigos(), options)
	request := models.NewCufdMasivoBuilder().WithCodigoModalidad(options.modalidad).
		WithDatosSolicitud(models.NewSolicitudListaCufdDtoBuilder().
			WithCodigoSucursal(options.sucursal).WithCodigoPuntoVenta(options.puntoVenta).WithCuis(cuis)).Build()
	response, err := client.Codigos().SolicitudCufdMasivo(t.Context(), request)
	require.NoError(t, err)
	require.NotNil(t, response)
	content, err := response.GetContent()
	require.NoError(t, err)
	result := content.RespuestaCufdMasivo
	require.True(t, result.Transaccion, "CUFD masivo rechazado: %+v", result.MensajesList)
	require.Len(t, result.ListaRespuestasCufd, 1, "debe responder a la entrada enviada")
	entry := result.ListaRespuestasCufd[0]
	require.True(t, entry.Transaccion, "entrada CUFD rechazada: %+v", entry.MensajeServicioList)
	require.NotEmpty(t, entry.Codigo)
	require.NotEmpty(t, entry.CodigoControl)
	require.Equal(t, cuis, entry.Cuis)
}

func TestSolicitudCuisMasivo(t *testing.T) {
	client, options := newCodigosIntegration(t)
	request := models.NewCuisMasivoBuilder().WithCodigoModalidad(options.modalidad).
		WithDatosSolicitud(models.NewSolicitudListaCuisDtoBuilder().
			WithCodigoSucursal(options.sucursal).WithCodigoPuntoVenta(options.puntoVenta)).Build()
	response, err := client.Codigos().SolicitudCuisMasivo(t.Context(), request)
	require.NoError(t, err)
	require.NotNil(t, response)
	content, err := response.GetContent()
	require.NoError(t, err)
	result := content.RespuestaCuisMasivo
	require.True(t, result.Transaccion, "CUIS masivo rechazado: %+v", result.MensajesList)
	require.Len(t, result.ListaRespuestasCuis, 1, "debe responder a la entrada enviada")
	entry := result.ListaRespuestasCuis[0]
	require.True(t, entry.Transaccion, "entrada CUIS rechazada: %+v", entry.MensajeServicioList)
	require.NotEmpty(t, entry.Codigo)
}

func TestCodigosVerificarComunicacion(t *testing.T) {
	client, _ := newCodigosIntegration(t)
	response, err := client.Codigos().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionCodigosBuilder().Build())
	require.NoError(t, err)
	require.NotNil(t, response)
	content, err := response.GetContent()
	require.NoError(t, err)
	result := content.RespuestaComunicacion
	require.True(t, result.Transaccion, "comunicación rechazada: %+v", result.MensajesList)
}

func TestVerificarNit(t *testing.T) {
	client, options := newCodigosIntegration(t)
	nit := options.config.Nit
	if value := strings.TrimSpace(os.Getenv("SIAT_NIT_VERIFICAR")); value != "" {
		var err error
		nit, err = strconv.ParseInt(value, 10, 64)
		require.NoError(t, err, "SIAT_NIT_VERIFICAR debe ser un entero")
		require.Positive(t, nit)
	}
	cuis := codigosCuisForDependentTest(t, client.Codigos(), options)
	request := models.NewVerificarNitBuilder().WithCodigoModalidad(options.modalidad).
		WithCodigoSucursal(options.sucursal).WithCuis(cuis).WithNitParaVerificacion(nit).Build()
	response, err := client.Codigos().VerificarNit(t.Context(), request)
	require.NoError(t, err)
	require.NotNil(t, response)
	content, err := response.GetContent()
	require.NoError(t, err)
	result := content.RespuestaVerificarNit
	require.True(t, result.Transaccion, "verificación de NIT rechazada: %+v", result.MensajesList)
}

// TestNotificaCertificadoRevocado necesita un certificado real de pruebas en
// SIAT_CERT_REVOCADO. No usa el certificado de firma del cliente por defecto.
func TestNotificaCertificadoRevocado(t *testing.T) {
	client, options := newCodigosIntegration(t)
	path := strings.TrimSpace(os.Getenv("SIAT_CERT_REVOCADO"))
	if path == "" {
		t.Skip("revocación no configurada: SIAT_CERT_REVOCADO debe apuntar a un certificado de pruebas")
	}
	certificate, err := os.ReadFile(path)
	require.NoError(t, err)
	require.NotEmpty(t, strings.TrimSpace(string(certificate)))
	cuis := codigosCuisForDependentTest(t, client.Codigos(), options)
	date := time.Now()
	request := models.NewNotificaCertificadoRevocadoBuilder().WithCodigoSucursal(options.sucursal).
		WithCuis(cuis).WithFechaRevocacion(&date).
		WithRazonRevocacion("Prueba de revocación por sistema").WithCertificado(string(certificate)).Build()
	response, err := client.Codigos().NotificaCertificadoRevocado(t.Context(), request)
	require.NoError(t, err)
	require.NotNil(t, response)
	content, err := response.GetContent()
	require.NoError(t, err)
	result := content.RespuestaNotificaRevocado
	require.True(t, result.Transaccion, "notificación de revocación rechazada: %+v", result.MensajesList)
}
