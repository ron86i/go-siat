package service_test

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/internal/adapter/service"
	"github.com/ron86i/go-siat/pkg/config"

	"github.com/ron86i/go-siat/internal/core/domain/facturacion/operaciones"
	"github.com/ron86i/go-siat/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestRegistroPuntoVenta(t *testing.T) {
	godotenv.Load()

	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	codModalidad, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	service, err := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)
	if err != nil {
		t.Fatalf("No se pudo inicializar el servicio de operaciones: %v", err)
	}

	req := operaciones.RegistroPuntoVenta{
		SolicitudRegistroPuntoVenta: operaciones.SolicitudRegistroPuntoVenta{
			CodigoAmbiente:       codAmbiente,
			CodigoModalidad:      codModalidad,
			CodigoSistema:        os.Getenv("SIAT_CODIGO_SISTEMA"),
			CodigoSucursal:       0,
			CodigoTipoPuntoVenta: 2, // Tipo de punto de venta (2,3,4,5,6)
			Cuis:                 "197C8240",
			Descripcion:          "Punto de Venta Prueba 1",
			Nit:                  nit,
			NombrePuntoVenta:     "PV1",
		},
	}

	resp, err := service.RegistroPuntoVenta(context.Background(), config, &req)

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

	nit, err := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	if err != nil {
		t.Fatalf("la variable SIAT_NIT debe ser un número válido: %v", err)
	}
	codAmbiente, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_AMBIENTE debe ser un número válido: %v", err)
	}
	codModalidad, err := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	if err != nil {
		t.Fatalf("la variable SIAT_CODIGO_MODALIDAD debe ser un número válido: %v", err)
	}

	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	service, err := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)
	if err != nil {
		t.Fatalf("No se pudo inicializar el servicio de operaciones: %v", err)
	}

	now := time.Now()
	req := operaciones.RegistroPuntoVentaComisionista{
		SolicitudPuntoVentaComisionista: operaciones.SolicitudPuntoVentaComisionista{
			CodigoAmbiente:   codAmbiente,
			CodigoModalidad:  codModalidad,
			CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
			CodigoSucursal:   0,
			Cuis:             "197C8240",
			Descripcion:      "Comisionista Prueba SIAT",
			FechaInicio:      now,
			FechaFin:         now.Add(24 * time.Hour),
			Nit:              nit,
			NitComisionista:  425951025, // Reemplazar por un NIT comisionista válido si es necesario
			NombrePuntoVenta: "COM-SIAT",
			NumeroContrato:   "CONT-001",
		},
	}

	resp, err := service.RegistroPuntoVentaComisionista(context.Background(), config, &req)

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

func TestOperacionesVerificarComunicacion(t *testing.T) {
	godotenv.Load()
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, _ := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)

	resp, err := service.VerificarComunicacion(context.Background(), config)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta: %v", resp)
		log.Printf("Resultado VerificarComunicacion: %v", resp.Body.Content.Return.Transaccion)
		assert.True(t, resp.Body.Content.Return.Transaccion)
	}
}

func TestConsultaPuntoVenta(t *testing.T) {
	godotenv.Load()
	nit, _ := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, _ := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)

	req := operaciones.ConsultaPuntoVenta{
		SolicitudConsultaPuntoVenta: operaciones.SolicitudConsultaPuntoVenta{
			CodigoAmbiente: 2,
			CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
			CodigoSucursal: 0,
			Cuis:           "197C8240",
			Nit:            nit,
		},
	}

	resp, err := service.ConsultaPuntoVenta(context.Background(), config, &req)

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
	nit, _ := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, _ := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)

	req := operaciones.CierrePuntoVenta{
		SolicitudCierrePuntoVenta: operaciones.SolicitudCierrePuntoVenta{
			CodigoAmbiente:   2,
			CodigoPuntoVenta: 13, // Un código que probablemente no exista o sea de prueba
			CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
			CodigoSucursal:   0,
			Cuis:             "197C8240",
			Nit:              nit,
		},
	}

	resp, err := service.CierrePuntoVenta(context.Background(), config, &req)

	if err == nil && resp != nil {
		log.Printf("Resultado Cierre PV: %v", resp.Body.Content.Respuesta.Transaccion)
	} else {
		log.Printf("Cierre PV finalizado con (posible error de negocio): %v", err)
	}
}

func TestCierreOperacionesSistema(t *testing.T) {
	godotenv.Load()
	nit, _ := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, _ := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)

	req := operaciones.CierreOperacionesSistema{
		SolicitudOperaciones: operaciones.SolicitudOperaciones{
			CodigoAmbiente:   2,
			CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
			CodigoSucursal:   0,
			Cuis:             "E97E7A16",
			Nit:              nit,
			CodigoModalidad:  1,
			CodigoPuntoVenta: 1,
		},
	}

	resp, err := service.CierreOperacionesSistema(context.Background(), config, &req)

	if err == nil && resp != nil {
		log.Printf("Resultado Cierre Sistema: %v", resp.Body.Content.Respuesta)
	} else {
		log.Printf("Cierre Sistema finalizado con: %v", err)
	}
}

func TestRegistroEventosSignificativos(t *testing.T) {
	godotenv.Load()
	nit, _ := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, _ := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)

	now := time.Now()
	req := operaciones.RegistroEventoSignificativo{
		SolicitudEventoSignificativo: operaciones.SolicitudEventoSignificativo{
			CodigoAmbiente:        2,
			CodigoMotivoEvento:    4,
			CodigoSistema:         os.Getenv("SIAT_CODIGO_SISTEMA"),
			CodigoSucursal:        0,
			Cuis:                  "197C8240",
			CufdEvento:            "FBQT5...",
			Cufd:                  "FBQT5...",
			Descripcion:           "Prueba de evento significativo",
			FechaHoraInicioEvento: now.Add(-1 * time.Hour),
			FechaHoraFinEvento:    now,
			Nit:                   nit,
			CodigoPuntoVenta:      0,
		},
	}

	resp, err := service.RegistroEventosSignificativos(context.Background(), config, &req)

	if err == nil && resp != nil {
		log.Printf("Resultado Registro Evento: %v", resp.Body.Content)
	} else {
		log.Printf("Registro Evento finalizado con: %v", err)
	}
}

func TestConsultaEventosSignificativos(t *testing.T) {
	godotenv.Load()
	nit, _ := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := config.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}
	service, _ := service.NewSiatOperacionesService(os.Getenv("SIAT_URL"), nil)

	req := operaciones.ConsultaEventoSignificativo{
		SolicitudConsultaEvento: operaciones.SolicitudConsultaEvento{
			CodigoAmbiente:   2,
			CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
			CodigoSucursal:   0,
			Cuis:             "197C8240",
			Nit:              nit,
			FechaEvento:      time.Now(),
			CodigoPuntoVenta: 0,
		},
	}

	resp, err := service.ConsultaEventosSignificativos(context.Background(), config, &req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		res := resp.Body.Content.Respuesta
		log.Printf("Eventos encontrados: %d", len(res.ListaCodigos))
	}
}
