package services_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/ron86i/go-siat/v2/pkg/utils"
	"github.com/stretchr/testify/assert"
)

type xmlTraceTransport struct {
	t         *testing.T
	base      http.RoundTripper
	outputDir string
	sequence  int
}

func (t *xmlTraceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(bytes.NewReader(body))
	t.sequence++
	sequence := t.sequence
	t.writeXML(sequence, "solicitud", body)
	t.t.Logf("XML enviado al SIAT:\n%s", body)

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = io.NopCloser(bytes.NewReader(responseBody))
	t.writeXML(sequence, "respuesta", responseBody)
	t.t.Logf("XML recibido del SIAT:\n%s", responseBody)

	return resp, nil
}

func (t *xmlTraceTransport) writeXML(sequence int, kind string, body []byte) {
	filename := fmt.Sprintf("%02d-%s.xml", sequence, kind)
	path := filepath.Join(t.outputDir, filename)
	if err := os.WriteFile(path, body, 0o600); err != nil {
		t.t.Logf("no se pudo guardar XML de %s: %v", kind, err)
		return
	}
	t.t.Logf("XML de %s guardado en %s", kind, path)
}

func newXMLTraceHTTPClient(t *testing.T) *http.Client {
	t.Helper()
	outputDir := ".siat-traces"
	if err := os.MkdirAll(outputDir, 0o700); err != nil {
		t.Fatalf("no se pudo crear directorio de trazas XML: %v", err)
	}
	return &http.Client{
		Transport: &xmlTraceTransport{
			t:         t,
			base:      http.DefaultTransport,
			outputDir: outputDir,
		},
	}
}

func TestConsultaPuntoVenta(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{},
	}

	siatClient, err := siat.New(cfg)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	service := siatClient.Operaciones()

	req := models.NewConsultaPuntoVentaBuilder().
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		Build()

	resp, err := service.ConsultaPuntoVenta(context.Background(), req)
	if err == nil && assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Body.Content)
	}
}

func TestSolicitudCuisYConsultaPuntosVenta(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("error cargando .env: %v", err)
	}

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("SIAT_NIT inválido: %v", err)
	}

	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("SIAT_CODIGO_MODALIDAD inválido: %v", err)
	}

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{},
	}

	siatClient, err := siat.New(cfg)
	if err != nil {
		t.Fatalf("error creando cliente: %v", err)
	}

	ctx := context.Background()
	codigoSucursal := 0
	codigoPuntoVenta := 0

	// 1. Obtener CUIS
	cuisReq := models.NewCuisBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCodigoPuntoVenta(codigoPuntoVenta).
		WithCodigoModalidad(codModalidad).
		Build()

	cuisResp, err := siatClient.Codigos().SolicitudCuis(ctx, cuisReq)
	if err != nil {
		t.Fatalf("error solicitando CUIS: %v", err)
	}

	resultadoCuis := cuisResp.Body.Content.RespuestaCuis
	if resultadoCuis.Codigo == "" {
		t.Fatalf("SIAT no devolvió CUIS: %+v", resultadoCuis.MensajesList)
	}

	t.Logf("CUIS obtenido: %s", resultadoCuis.Codigo)

	// 2. Consultar puntos de venta usando el CUIS
	puntosReq := models.NewConsultaPuntoVentaBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCuis(resultadoCuis.Codigo).
		Build()

	puntosResp, err := siatClient.Operaciones().ConsultaPuntoVenta(ctx, puntosReq)
	if err != nil {
		t.Fatalf("error consultando puntos de venta: %v", err)
	}

	resultadoPuntos := puntosResp.Body.Content.Respuesta
	if !resultadoPuntos.Transaccion {
		t.Fatalf("consulta rechazada: %+v", resultadoPuntos.MensajesList)
	}

	for _, punto := range resultadoPuntos.ListaPuntosVentas {
		t.Logf(
			"Punto de venta: código=%d, nombre=%s, tipo=%s",
			punto.CodigoPuntoVenta,
			punto.NombrePuntoVenta,
			punto.TipoPuntoVenta,
		)
	}
}

func TestSolicitudCuisYCierreOperaciones(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("error cargando .env: %v", err)
	}

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("SIAT_NIT inválido: %v", err)
	}

	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("SIAT_CODIGO_MODALIDAD inválido: %v", err)
	}

	siatClient, err := siat.New(siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{},
	})
	if err != nil {
		t.Fatalf("error creando cliente: %v", err)
	}

	ctx := context.Background()
	codigoSucursal := 0
	codigoPuntoVenta := 0

	// 1. Obtener un CUIS para la sucursal y punto de venta.
	cuisReq := models.NewCuisBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCodigoPuntoVenta(codigoPuntoVenta).
		WithCodigoModalidad(codModalidad).
		Build()

	cuisResp, err := siatClient.Codigos().SolicitudCuis(ctx, cuisReq)
	if err != nil {
		t.Fatalf("error solicitando CUIS: %v", err)
	}

	resultadoCuis := cuisResp.Body.Content.RespuestaCuis
	if resultadoCuis.Codigo == "" {
		t.Fatalf("SIAT no devolvió CUIS: %+v", resultadoCuis.MensajesList)
	}
	t.Logf("CUIS obtenido: %s (transacción: %v, vigencia: %s)",
		resultadoCuis.Codigo,
		resultadoCuis.Transaccion,
		resultadoCuis.FechaVigencia.Format(time.RFC3339),
	)

	// 2. Cerrar las operaciones del sistema usando el CUIS obtenido.
	cierreReq := models.NewCierreOperacionesSistemaBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCodigoPuntoVenta(codigoPuntoVenta).
		WithCodigoModalidad(codModalidad).
		WithCuis(resultadoCuis.Codigo).
		Build()

	cierreResp, err := siatClient.Operaciones().CierreOperacionesSistema(ctx, cierreReq)
	if err != nil {
		t.Fatalf("error cerrando operaciones: %v", err)
	}

	resultadoCierre := cierreResp.Body.Content.Respuesta
	for _, mensaje := range resultadoCierre.MensajesList {
		t.Logf("Mensaje SIAT [%d]: %s", mensaje.Codigo, mensaje.Descripcion)
	}

	assert.True(t, resultadoCierre.Transaccion,
		"el cierre de operaciones fue rechazado: %+v", resultadoCierre.MensajesList)
}

func TestSolicitudCuisYRegistroPuntoVenta(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("error cargando .env: %v", err)
	}

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("SIAT_NIT inválido: %v", err)
	}

	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("SIAT_CODIGO_MODALIDAD inválido: %v", err)
	}

	codTipoPuntoVenta := 5

	siatClient, err := siat.New(siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{},
	})
	if err != nil {
		t.Fatalf("error creando cliente: %v", err)
	}

	ctx := context.Background()
	codigoSucursal := 0
	// El CUIS de registro debe corresponder al sistema de la sucursal,
	// normalmente el punto de venta principal (0).
	codigoPuntoVentaPrincipal := 0

	cuisReq := models.NewCuisBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCodigoPuntoVenta(codigoPuntoVentaPrincipal).
		WithCodigoModalidad(codModalidad).
		Build()

	cuisResp, err := siatClient.Codigos().SolicitudCuis(ctx, cuisReq)
	if err != nil {
		t.Fatalf("error solicitando CUIS: %v", err)
	}

	resultadoCuis := cuisResp.Body.Content.RespuestaCuis
	if resultadoCuis.Codigo == "" {
		t.Fatalf("SIAT no devolvió CUIS: %+v", resultadoCuis.MensajesList)
	}
	t.Logf("CUIS obtenido: %s", resultadoCuis.Codigo)

	nombre := "PV-PRUEBA-" + time.Now().Format("20060102-150405")
	registroReq := models.NewRegistroPuntoVentaBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCodigoModalidad(codModalidad).
		WithCodigoTipoPuntoVenta(codTipoPuntoVenta).
		WithCuis(resultadoCuis.Codigo).
		WithNombrePuntoVenta(nombre).
		WithDescripcion("Punto de venta creado por prueba de integración").
		Build()

	registroResp, err := siatClient.Operaciones().RegistroPuntoVenta(ctx, registroReq)
	if err != nil {
		t.Fatalf("error registrando punto de venta: %v", err)
	}

	resultadoRegistro := registroResp.Body.Content.Respuesta
	for _, mensaje := range resultadoRegistro.MensajesList {
		t.Logf("Mensaje SIAT [%d]: %s", mensaje.Codigo, mensaje.Descripcion)
	}

	assert.True(t, resultadoRegistro.Transaccion,
		"registro rechazado: %+v", resultadoRegistro.MensajesList)
	if resultadoRegistro.Transaccion {
		t.Logf("Punto de venta registrado: código=%d, nombre=%s",
			resultadoRegistro.CodigoPuntoVenta, nombre)
	}
}

func TestSolicitudCuisYCierrePuntoVenta(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("error cargando .env: %v", err)
	}

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("SIAT_NIT inválido: %v", err)
	}

	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("SIAT_CODIGO_MODALIDAD inválido: %v", err)
	}

	codigoPuntoVenta := 11
	siatClient, err := siat.New(siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     newXMLTraceHTTPClient(t),
	})
	if err != nil {
		t.Fatalf("error creando cliente: %v", err)
	}

	ctx := context.Background()
	codigoSucursal := 0

	// El CUIS identifica el sistema en la sucursal; el punto de venta a
	// cerrar se indica por separado en la solicitud de cierre.
	cuisReq := models.NewCuisBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCodigoPuntoVenta(0).
		WithCodigoModalidad(codModalidad).
		Build()

	cuisResp, err := siatClient.Codigos().SolicitudCuis(ctx, cuisReq)
	if err != nil {
		t.Fatalf("error solicitando CUIS: %v", err)
	}

	resultadoCuis := cuisResp.Body.Content.RespuestaCuis
	if resultadoCuis.Codigo == "" {
		t.Fatalf("SIAT no devolvió CUIS: %+v", resultadoCuis.MensajesList)
	}
	t.Logf("CUIS obtenido para sucursal %d: %s (transacción: %v)",
		codigoSucursal, resultadoCuis.Codigo, resultadoCuis.Transaccion)

	// Confirmar que el punto existe antes de ejecutar una operación destructiva.
	consultaReq := models.NewConsultaPuntoVentaBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCuis(resultadoCuis.Codigo).
		Build()

	consultaResp, err := siatClient.Operaciones().ConsultaPuntoVenta(ctx, consultaReq)
	if err != nil {
		t.Fatalf("error consultando puntos de venta: %v", err)
	}

	resultadoConsulta := consultaResp.Body.Content.Respuesta
	if !resultadoConsulta.Transaccion {
		t.Fatalf("consulta de puntos de venta rechazada: %+v", resultadoConsulta.MensajesList)
	}

	existePuntoVenta := false
	codigosDisponibles := make([]int, 0, len(resultadoConsulta.ListaPuntosVentas))
	for _, punto := range resultadoConsulta.ListaPuntosVentas {
		codigosDisponibles = append(codigosDisponibles, punto.CodigoPuntoVenta)
		if punto.CodigoPuntoVenta == codigoPuntoVenta {
			existePuntoVenta = true
		}
	}
	if !existePuntoVenta {
		t.Fatalf("el punto de venta %d no existe en la sucursal %d; códigos disponibles: %v",
			codigoPuntoVenta, codigoSucursal, codigosDisponibles)
	}

	cierreReq := models.NewCierrePuntoVentaBuilder().
		WithCodigoSucursal(codigoSucursal).
		WithCodigoPuntoVenta(codigoPuntoVenta).
		WithCuis(resultadoCuis.Codigo).
		Build()

	cierreResp, err := siatClient.Operaciones().CierrePuntoVenta(ctx, cierreReq)
	if err != nil {
		t.Fatalf("error cerrando punto de venta: %v", err)
	}

	resultadoCierre := cierreResp.Body.Content.Respuesta
	for _, mensaje := range resultadoCierre.MensajesList {
		t.Logf("Mensaje SIAT [%d]: %s", mensaje.Codigo, mensaje.Descripcion)
	}

	assert.True(t, resultadoCierre.Transaccion,
		"cierre de punto de venta rechazado: %+v", resultadoCierre.MensajesList)
	if resultadoCierre.Transaccion {
		t.Logf("Punto de venta cerrado: código=%d", resultadoCierre.CodigoPuntoVenta)
	}
}
