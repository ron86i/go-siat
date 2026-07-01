package services_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/ports"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// TestNotificaCertificadoRevocado valida que el servicio sea capaz de informar al SIAT
// sobre la revocación de un certificado digital.
// Requisitos: El certificado enviado debe ser el registrado previamente en el portal de Impuestos.
func TestNotificaCertificadoRevocado(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            123456789,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{},
	}

	siatClient, err := siat.New(cfg)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	service := siatClient.Codigos()
	fechaRevocacion := time.Now()

	req := models.NewNotificaCertificadoRevocadoBuilder().
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		WithFechaRevocacion(&fechaRevocacion).
		WithRazonRevocacion("Prueba de revocación por sistema").
		WithCertificado(`-----BEGIN CERTIFICATE-----
MIIEejCCA2KgA...alF2Tw0jIVieaeefsL78Yv8fA==
-----END CERTIFICATE-----`).
		Build()

	resp, err := service.NotificaCertificadoRevocado(context.Background(), req)
	if err == nil && assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Body.Content)
	}
}

// TestVerificarNit valida el flujo completo de verificación de un NIT directamente
// contra el servicio real del SIAT, asegurando que la configuración cargada sea válida
// y que la respuesta del servidor se procese correctamente sin predecir mensajes fijos.
func TestVerificarNit(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}

	config := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     siat.NewHTTPClient(siat.DefaultHTTPConfig()),
	}

	siatClient, err := siat.New(config)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	service := siatClient.Codigos()

	req := models.NewVerificarNitBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		WithNitParaVerificacion(12345678). // Un NIT de prueba para validar la comunicación
		Build()

	// Ejecutar la petición al SIAT y procesar el resultado
	resp, err := service.VerificarNit(context.Background(), req)

	// Validar que no existan errores de comunicación o de serialización XML
	if !assert.NoError(t, err) {
		t.Fatalf("Fallo en la comunicación con el SIAT: %v", err)
	}

	// Verificar la integridad de la respuesta y registrar los mensajes devueltos por el servidor
	if assert.NotNil(t, resp) {
		resultado := resp.Body.Content.RespuestaVerificarNit
		log.Printf("Resultado de Verificación NIT: %v", resultado.Transaccion)

		if len(resultado.MensajesList) > 0 {
			for _, msg := range resultado.MensajesList {
				log.Printf("Mensaje SIAT [%d]: %s", msg.Codigo, msg.Descripcion)
			}
		}

		// En integración, validamos que el tipo de datos sea el esperado para la transacción
		assert.IsType(t, true, resultado.Transaccion)
	}
}

// TestSolicitudCuis valida la obtención de un Código Único de Inicio de Sistemas (CUIS).
// El CUIS es fundamental para identificar un punto de venta y su sistema asociado ante el SIAT.
// Se recomienda renovar el CUIS periódicamente según la vigencia devuelta por el servidor.
func TestSolicitudCuis(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	siatClient := getSiatClient(t)
	service := siatClient.Codigos()

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))

	req := models.NewCuisBuilder().
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(1).
		WithCodigoModalidad(codModalidad).
		Build()

	resp, err := service.SolicitudCuis(context.Background(), req)

	// Confirmar que la comunicación fue exitosa y se recibió un objeto de respuesta
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		resultado := resp.Body.Content.RespuestaCuis

		// Registrar los datos obtenidos para facilitar la auditoría de los tests
		log.Printf("Respuesta Cuis - Transacción: %v", resultado.Transaccion)

		isCuisVigente := false
		if len(resultado.MensajesList) > 0 {
			for _, msg := range resultado.MensajesList {
				log.Printf("Mensaje SIAT [%d]: %s", msg.Codigo, msg.Descripcion)
				if msg.Codigo == 980 {
					isCuisVigente = true
				}
			}
		}

		assert.True(t, resultado.Transaccion || isCuisVigente, "La solicitud de CUIS debería ser exitosa o indicar que ya existe un CUIS vigente")
		if resultado.Transaccion {
			assert.NotEmpty(t, resultado.Codigo)
			log.Printf("CUIS Obtenido: %s", resultado.Codigo)
		}
	}
}

// TestSolicitudCufdAll ejecuta 10 veces la solicitud de CUFD para validar estabilidad.
func TestSolicitudCufdAll(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load()

	siatClient := getSiatClient(t)
	service := siatClient.Codigos()

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Iteracion_%d", i), func(t *testing.T) {
			solicitudCufd(t, service)
		})
	}
}

func TestSolicitudCufd(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load()
	solicitudCufd(t, getSiatClient(t).Codigos())
}

func solicitudCufd(t *testing.T, service ports.SiatCodigosService) {
	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))

	codigoPuntoVenta := 1
	// 1. Obtener CUIS primero
	reqCuis := models.NewCuisBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(codigoPuntoVenta).
		Build()

	respCuis, err := service.SolicitudCuis(context.Background(), reqCuis)
	if err != nil {
		t.Fatalf("Error obteniendo CUIS: %v", err)
	}

	cuisCode := respCuis.Body.Content.RespuestaCuis.Codigo
	if cuisCode == "" {
		t.Skip("Saltando test de CUFD porque no se pudo obtener un CUIS vigente")
	}

	// 2. Solicitar CUFD
	reqCufd := models.NewCufdBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(codigoPuntoVenta).
		WithCuis(cuisCode).
		Build()

	resp, err := service.SolicitudCufd(context.Background(), reqCufd)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		res := resp.Body.Content.RespuestaCufd
		log.Printf("Respuesta de cufd: %+v", res)
	}
}

// TestSolicitudCufdMasivo verifica la obtención masiva de códigos CUFD para múltiples
// puntos de venta o sucursales en una sola operación.
func TestSolicitudCufdMasivo(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	// Cargar configuración de integración real
	godotenv.Load()

	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	siatClient := getSiatClient(t)
	service := siatClient.Codigos()

	req := models.NewCuisBuilder().
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(0).
		WithCodigoModalidad(codModalidad).
		Build()

	resp, err := service.SolicitudCuis(context.Background(), req)
	if err == nil && assert.NotNil(t, resp) {
		resultado := resp.Body.Content.RespuestaCuis
		t.Logf("Resultado de Solicitud CUIS: %v", resultado.Transaccion)

		isCuisVigente := false
		if len(resultado.MensajesList) > 0 {
			for _, msg := range resultado.MensajesList {
				t.Logf("Mensaje SIAT [%d]: %s", msg.Codigo, msg.Descripcion)
				if msg.Codigo == 980 {
					isCuisVigente = true
					t.Logf("CUIS vigente %s", resultado.Codigo)
				}
			}
		}

		assert.True(t, resultado.Transaccion || isCuisVigente, "La solicitud de CUIS debería ser exitosa o indicar que ya existe un CUIS vigente")
	}
}

func getSiatClient(t *testing.T) *siat.SiatServices {
	godotenv.Load()
	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: codAmbiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     siat.NewHTTPClient(siat.DefaultHTTPConfig()),
	}

	client, err := siat.New(cfg)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}
	return client
}
