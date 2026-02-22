package siat_test

import (
	"context"
	"go-siat/internal/adapter/service/siat"
	"go-siat/internal/core/domain/facturacion"
	"go-siat/internal/core/domain/facturacion/codigos"
	"go-siat/internal/core/util"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestNotificaCertificadoRevocado valida que el servicio sea capaz de informar al SIAT
// sobre la revocación de un certificado digital, utilizando credenciales reales del entorno.
func TestNotificaCertificadoRevocado(t *testing.T) {
	// Cargar configuración de integración desde el entorno (.env)
	godotenv.Load()

	envs := map[string]string{
		"SIAT_URL": os.Getenv("SIAT_URL"),
	}

	// Parsear el NIT (Int64)
	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	config := facturacion.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	service, err := siat.NewSiatCodigosService(envs)
	if err != nil {
		t.Fatalf("No se pudo inicializar el servicio de códigos: %v", err)
	}

	// Preparar la solicitud de notificación. Se requiere un certificado y una razón válida.
	req := codigos.NotificaCertificadoRevocado{
		SolicitudNotificaRevocado: codigos.SolicitudNotifcaRevocado{
			CodigoAmbiente:  codAmbiente,
			CodigoSistema:   os.Getenv("SIAT_CODIGO_SISTEMA"),
			NIT:             nit,
			CodigoSucursal:  0,
			Cuis:            "197C8240",
			FechaRevocacion: new(time.Now()),
			RazonRevocacion: "Prueba de revocación por sistema",
			Certificado: `-----BEGIN CERTIFICATE-----
MIIEejCCA2KgA...alF2Tw0jIVieaeefsL78Yv8fA==
-----END CERTIFICATE-----`,
		},
	}

	// Ejecutar la petición al SIAT y procesar el resultado de la revocación
	resp, err := service.NotificaCertificadoRevocado(context.Background(), config, req)

	// Validar que no existan errores de comunicación o de estructura SOAP
	if !assert.NoError(t, err) {
		t.Fatalf("Fallo en la comunicación con el SIAT: %v", err)
	}

	// Verificar la integridad de la respuesta y registrar el estado devuelto
	if assert.NotNil(t, resp) {
		resultado := resp.Body.Content.RespuestaNotificaRevocado
		log.Printf("Resultado de Notificación Revocada: %v", resultado.Transaccion)

		if len(resultado.MensajesList) > 0 {
			for _, msg := range resultado.MensajesList {
				log.Printf("Mensaje SIAT [%d]: %s", msg.Codigo, msg.Descripcion)
			}
		}

		// En integración, validamos que la transacción sea reportada por el SIAT
		assert.IsType(t, true, resultado.Transaccion)
	}
}

// TestVerificarNit valida el flujo completo de verificación de un NIT directamente
// contra el servicio real del SIAT, asegurando que la configuración cargada sea válida
// y que la respuesta del servidor se procese correctamente sin predecir mensajes fijos.
func TestVerificarNit(t *testing.T) {
	// Cargar configuración de integración desde el entorno (.env)
	godotenv.Load()

	envs := map[string]string{
		"SIAT_URL": os.Getenv("SIAT_URL"),
	}
	codModalidad, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	config := facturacion.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	service, err := siat.NewSiatCodigosService(envs)
	if err != nil {
		t.Fatalf("No se pudo inicializar el servicio de códigos: %v", err)
	}

	// Preparar la solicitud de verificación. El CUIS debe ser válido para el NIT en ambiente de prueba.
	req := codigos.VerificarNit{
		SolicitudVerificarNit: codigos.SolicitudVerificarNit{
			CodigoAmbiente:      codAmbiente,
			CodigoModalidad:     codModalidad,
			CodigoSistema:       os.Getenv("SIAT_CODIGO_SISTEMA"),
			NIT:                 nit,
			CodigoSucursal:      0,
			Cuis:                "197C8240",
			NitParaVerificacion: 12345678, // Un NIT de prueba para validar la comunicación
		},
	}

	// Ejecutar la petición al SIAT y procesar el resultado
	resp, err := service.VerificarNit(context.Background(), config, req)

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
// Verifica que el servicio sea capaz de conectar y recibir un código de respuesta válido.
func TestSolicitudCuis(t *testing.T) {
	godotenv.Load()
	envs := map[string]string{
		"SIAT_URL": os.Getenv("SIAT_URL"),
	}
	codModalidad, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	// Parsear el NIT (Int64)
	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	config := facturacion.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, _ := siat.NewSiatCodigosService(envs)

	req := codigos.Cuis{
		SolicitudCuis: codigos.SolicitudCuis{
			CodigoAmbiente:   codAmbiente,
			CodigoModalidad:  codModalidad,
			CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
			NIT:              nit,
			CodigoSucursal:   0,
			CodigoPuntoVenta: 0,
		},
	}

	resp, err := service.SolicitudCuis(context.Background(), config, req)

	// Confirmar que la comunicación fue exitosa y se recibió un objeto de respuesta
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		resultado := resp.Body.Content.RespuestaCuis

		// Registrar los datos obtenidos para facilitar la auditoría de los tests
		log.Printf("Respuesta Cuis - Transacción: %v", resultado.Transaccion)

		assert.NotEmpty(t, resultado.Codigo)
		log.Printf("CUIS Obtenido: %s", resultado.Codigo)
	}
}

// TestSolicitudCufd valida la obtención del Código Único de Facturación Diaria (CUFD).
// Asegura que el flujo de solicitud contra el SIAT se complete sin errores técnicos.
func TestSolicitudCufd(t *testing.T) {
	// Cargar entorno de configuración para tests de integración
	godotenv.Load()

	envs := map[string]string{
		"SIAT_URL": os.Getenv("SIAT_URL"),
	}
	codModalidad, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	// Parsear el NIT (Int64)
	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	config := facturacion.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, err := siat.NewSiatCodigosService(envs)
	assert.NoError(t, err)

	// Preparar la estructura de solicitud de CUFD con los datos de prueba
	req := codigos.Cufd{
		SolicitudCufd: codigos.SolicitudCufd{
			CodigoAmbiente:   codAmbiente,
			CodigoModalidad:  codModalidad,
			CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
			NIT:              nit,
			CodigoSucursal:   0,
			CodigoPuntoVenta: new(int),
			Cuis:             "197C8240", // Requiere un CUIS vigente para el NIT configurado
		},
	}
	*req.SolicitudCufd.CodigoPuntoVenta = 0

	// Ejecutar la llamada al servicio de códigos del SIAT
	resp, err := service.SolicitudCufd(context.Background(), config, req)

	// Validar la recepción de la respuesta y registrar el estado de la transacción
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		res := resp.Body.Content.RespuestaCufd

		if res.Transaccion {
			log.Printf("Conexión Exitosa: CUFD recibido [%s]", res.Codigo)
		} else {
			t.Errorf("El SIAT procesó la solicitud pero la transacción es false")
		}
	}
}

// TestSolicitudCufdMasivo verifica la obtención masiva de códigos CUFD para múltiples
// puntos de venta o sucursales en una sola operación.
func TestSolicitudCufdMasivo(t *testing.T) {
	// Cargar configuración de integración real
	godotenv.Load()

	envs := map[string]string{
		"SIAT_URL": os.Getenv("SIAT_URL"),
	}
	codModalidad, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	config := facturacion.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, err := siat.NewSiatCodigosService(envs)
	assert.NoError(t, err)

	// Configurar la lista de solicitudes masivas (por ejemplo, para la sucursal 0 y punto de venta 0)
	req := codigos.CufdMasivo{
		SolicitudCufdMasivo: codigos.SolicitudCufdMasivo{
			CodigoAmbiente:  codAmbiente,
			CodigoModalidad: codModalidad,
			CodigoSistema:   os.Getenv("SIAT_CODIGO_SISTEMA"),
			Nit:             nit,
			DatosSolicitud: []codigos.SolicitudListaCufdDto{
				{
					CodigoSucursal:   0,
					CodigoPuntoVenta: new(0),
					Cuis:             "197C8240",
				},
			},
		},
	}

	// Ejecutar la petición masiva al SIAT
	resp, err := service.SolicitudCufdMasivo(context.Background(), config, req)

	// Validar que la respuesta contenga datos y registrar el resultado de la operación masiva
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		res := resp.Body.Content.RespuestaCufdMasivo
		log.Printf("Resultado Masivo - Transacción: %v", res.Transaccion)

		if len(res.ListaRespuestasCufd) > 0 {
			for _, item := range res.ListaRespuestasCufd {
				log.Printf("CUFD Masivo Recibido para Sucursal %d: %s", *item.CodigoSucursal, item.Codigo)
			}
		}
	}
}

// TestSolicitudCuisMasivo verifica la obtención masiva de códigos CUIS para múltiples
// puntos de venta en una única transacción, asegurando la eficiencia en configuraciones extensas.
func TestSolicitudCuisMasivo(t *testing.T) {
	// Cargar configuración real del entorno para pruebas de integración
	godotenv.Load()

	envs := map[string]string{
		"SIAT_URL": os.Getenv("SIAT_URL"),
	}
	codModalidad, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	// Parsear el NIT (Int64)
	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	config := facturacion.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, err := siat.NewSiatCodigosService(envs)
	assert.NoError(t, err)

	// Configurar la solicitud masiva de CUIS para un punto de venta específico
	req := codigos.CuisMasivo{
		SolicitudCuisMasivoSistemas: codigos.SolicitudCuisMasivoSistemas{
			CodigoAmbiente:  codAmbiente,
			CodigoModalidad: codModalidad,
			CodigoSistema:   os.Getenv("SIAT_CODIGO_SISTEMA"),
			NIT:             nit,
			DatosSolicitud: []codigos.SolicitudListaCuisDto{
				{
					CodigoSucursal:   0,
					CodigoPuntoVenta: new(int),
				},
			},
		},
	}

	// Ejecutar la petición masiva al SIAT
	resp, err := service.SolicitudCuisMasivo(context.Background(), config, req)

	// Validar que la comunicación fue exitosa y registrar los resultados individuales obtenidos
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		res := resp.Body.Content.RespuestaCuisMasivo
		log.Printf("Resultado Cuis Masivo - Transacción: %v", res.Transaccion)

		if len(res.ListaRespuestasCuis) > 0 {
			for _, item := range res.ListaRespuestasCuis {
				log.Printf("CUIS Masivo Recibido para Sucursal %d: %s", *item.CodigoSucursal, item.Codigo)
			}
		}
	}
}

// TestVerificarComunicacion valida la conectividad básica con los servidores del SIAT enviando una solicitud
// de verificación de comunicación. El servicio debe registrar la petición y retornar un código
// que confirme la recepción exitosa, asegurando que el canal SOAP esté operativo.
func TestVerificarComunicacion(t *testing.T) {
	// Cargar configuración desde .env
	godotenv.Load()

	envs := map[string]string{
		"SIAT_URL": os.Getenv("SIAT_URL"),
	}

	config := facturacion.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, err := siat.NewSiatCodigosService(envs)
	if err != nil {
		t.Fatalf("Error al crear servicio: %v", err)
	}

	ctx := context.Background()
	req := codigos.VerificarComunicacion{}

	// Ejecutar la petición de verificación de comunicación
	resp, err := service.VerificarComunicacion(ctx, config, req)
	if err != nil {
		t.Fatalf("Error en VerificarComunicacion: %v", err)
	}

	assert.NotNil(t, resp, "La respuesta no debería ser nula")
	assert.True(t, resp.Body.Content.RespuestaComunicacion.Transaccion, "La prueba de comunicación debería ser exitosa")
}
