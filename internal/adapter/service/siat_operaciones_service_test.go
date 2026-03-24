package service_test

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"

	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// TestRegistroPuntoVenta valida la creación de un nuevo punto de venta ante el SIAT.
// Los puntos de venta son necesarios para organizar la facturación por sucursales o terminales.
func TestRegistroPuntoVenta(t *testing.T) {
	godotenv.Load()

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()

	req := models.Operaciones().NewRegistroPuntoVentaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCodigoTipoPuntoVenta(2). // Tipo de punto de venta (2,3,4,5,6)
		WithCuis("197C8240").
		WithDescripcion("Punto de Venta Prueba 1").
		WithNit(nit).
		WithNombrePuntoVenta("PV1").
		Build()

	resp, err := service.RegistroPuntoVenta(context.Background(), config, req)

	if !assert.NoError(t, err) {
		t.Fatalf("Fallo en la comunicación con el SIAT: %v", err)
	}
	xmlBody, err := xml.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Printf("Error al convertir a XML: %v", err)
	}
	log.Printf("Respuesta: %s", string(xmlBody))
	if assert.NotNil(t, resp) {
		resultado := resp.Body.Content.Respuesta
		log.Printf("Resultado Registro Punto Venta - Transacción: %v", resultado.Transaccion)
		if resultado.Transaccion {
			log.Printf("Código Punto Venta: %d", resultado.CodigoPuntoVenta)
		}
		for _, msg := range resultado.MensajesList {
			log.Printf("Mensaje SIAT [%d]: %s", msg.Codigo, msg.Descripcion)
		}
	}
}

func TestRegistroPuntoVentaComisionista(t *testing.T) {
	godotenv.Load()

	nit, err := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	codModalidad, err := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()

	now := time.Now()
	req := models.Operaciones().NewRegistroPuntoVentaComisionistaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		WithDescripcion("Comisionista Prueba SIAT").
		WithFechaInicio(now).
		WithFechaFin(now.Add(24 * time.Hour)).
		WithNit(nit).
		WithNitComisionista(425951025). // Reemplazar por un NIT comisionista válido si es necesario
		WithNombrePuntoVenta("COM-SIAT").
		WithNumeroContrato("CONT-001").
		Build()

	resp, err := service.RegistroPuntoVentaComisionista(context.Background(), config, req)

	if err != nil {
		log.Printf("Error/Soap Fault en RegistroPuntoVentaComisionista: %v", err)
		t.Errorf("Error en RegistroPuntoVentaComisionista: %v", err)
	}

	if resp != nil {
		resultado := resp.Body.Content.Respuesta
		log.Printf("Resultado Registro Comisionista - Transacción: %v", resultado.Transaccion)
		for _, msg := range resultado.MensajesList {
			log.Printf("Mensaje SIAT [%d]: %s", msg.Codigo, msg.Descripcion)
		}
	}
}

// TestOperacionesVerificarComunicacion valida la conectividad SOAP con el servicio de Operaciones.
func TestOperacionesVerificarComunicacion(t *testing.T) {
	godotenv.Load()
	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()
	req := models.Operaciones().NewVerificarComunicacionBuilder().Build()
	resp, err := service.VerificarComunicacion(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta: %v", resp)
		log.Printf("Resultado VerificarComunicacion: %v", resp.Body.Content.Return.Transaccion)
		assert.True(t, resp.Body.Content.Return.Transaccion)
	}
}

func TestConsultaPuntoVenta(t *testing.T) {
	godotenv.Load()
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()

	req := models.Operaciones().NewConsultaPuntoVentaBuilder().
		WithCodigoAmbiente(2).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		WithNit(nit).
		Build()

	resp, err := service.ConsultaPuntoVenta(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		res := resp.Body.Content.Respuesta
		log.Printf("Puntos de Venta encontrados: %d", len(res.ListaPuntosVentas))
		for _, pv := range res.ListaPuntosVentas {
			log.Printf("PV: %s (Código: %d)", pv.NombrePuntoVenta, pv.CodigoPuntoVenta)
		}
	}
}

func TestCierrePuntoVenta(t *testing.T) {
	godotenv.Load()
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()

	req := models.Operaciones().NewCierrePuntoVentaBuilder().
		WithCodigoAmbiente(2).
		WithCodigoPuntoVenta(13). // Un código que probablemente no exista o sea de prueba
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		WithNit(nit).
		Build()

	resp, err := service.CierrePuntoVenta(context.Background(), config, req)

	if err == nil && resp != nil {
		log.Printf("Resultado Cierre PV: %v", resp.Body.Content.Respuesta.Transaccion)
	} else {
		log.Printf("Cierre PV finalizado con (posible error de negocio): %v", err)
	}
}

func TestCierreOperacionesSistema(t *testing.T) {
	godotenv.Load()
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()

	req := models.Operaciones().NewCierreOperacionesSistemaBuilder().
		WithCodigoAmbiente(2).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis("E97E7A16").
		WithNit(nit).
		WithCodigoModalidad(1).
		WithCodigoPuntoVenta(1).
		Build()

	resp, err := service.CierreOperacionesSistema(context.Background(), config, req)

	if err == nil && resp != nil {
		log.Printf("Resultado Cierre Sistema: %v", resp.Body.Content.Respuesta)
	} else {
		log.Printf("Cierre Sistema finalizado con: %v", err)
	}
}

// TestRegistroEventosSignificativos reporta sucesos que impiden la facturación en línea (contingencias).
// Es obligatorio reportar el inicio y fin de estos eventos para justificar la facturación offline.
func TestRegistroEventosSignificativos(t *testing.T) {
	godotenv.Load()
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	serviceCodigos := siatClient.Codigos()
	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	now := time.Now()
	req := models.Operaciones().NewRegistroEventoSignificativoBuilder().
		WithCodigoAmbiente(2).
		WithCodigoMotivoEvento(4).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithCufdEvento("FBQT5CwqE4TERBI5RjlGOEM3MDc=QkFQMVZVTERhV...").
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithDescripcion("Prueba de evento significativo").
		WithFechaHoraInicioEvento(now.Add(-1 * time.Minute)).
		WithFechaHoraFinEvento(now).
		WithNit(nit).
		WithCodigoPuntoVenta(0).
		Build()

	resp, err := service.RegistroEventosSignificativos(context.Background(), config, req)

	if err == nil && resp != nil {
		log.Printf("Resultado Registro Evento: %v", resp.Body.Content)
	} else {
		log.Printf("Registro Evento finalizado con: %v", err)
	}
}

func TestConsultaEventosSignificativos(t *testing.T) {
	godotenv.Load()
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Operaciones()

	req := models.Operaciones().NewConsultaEventoSignificativoBuilder().
		WithCodigoAmbiente(2).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		WithNit(nit).
		WithFechaEvento(time.Now()).
		WithCodigoPuntoVenta(0).
		Build()

	resp, err := service.ConsultaEventosSignificativos(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		res := resp.Body.Content.Respuesta
		log.Printf("Eventos encontrados: %d", len(res.ListaCodigos))
	}
}
