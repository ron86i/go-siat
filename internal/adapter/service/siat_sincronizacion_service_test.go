package service_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"

	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
	"github.com/ron86i/go-siat/internal/core/port"

	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func runSincronizacionTest[V any, ReqType any](
	t *testing.T,
	name string,
	req ReqType,
	fn func(context.Context, port.Config, ReqType) (*V, error),
) {
	t.Run(name, func(t *testing.T) {
		godotenv.Load()

		config := siat.Config{
			Token: os.Getenv("SIAT_TOKEN"),
		}

		_, err := siat.New(os.Getenv("SIAT_URL"), nil)
		if err != nil {
			t.Fatalf("No se pudo inicializar el cliente SIAT: %v", err)
		}

		resp, err := fn(context.Background(), config, req)
		if err != nil {
			t.Fatalf("Error en %s: %v", name, err)
		}

		assert.NotNil(t, resp)
		log.Printf("Resultado de %s: %+v", name, resp)
	})
}

func getCommonRequest(_ *testing.T) sincronizacion.SolicitudSincronizacion {
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))

	return sincronizacion.SolicitudSincronizacion{
		CodigoAmbiente:   codAmbiente,
		NIT:              nit,
		CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoSucursal:   0,
		CodigoPuntoVenta: 0,
		Cuis:             "C2FC682C",
	}
}

func buildSincronizacion[T any, R any](b models.SincronizacionBuilder[T, R], sol sincronizacion.SolicitudSincronizacion) R {
	return b.WithCodigoAmbiente(sol.CodigoAmbiente).
		WithCodigoPuntoVenta(sol.CodigoPuntoVenta).
		WithCodigoSistema(sol.CodigoSistema).
		WithCodigoSucursal(sol.CodigoSucursal).
		WithCuis(sol.Cuis).
		WithNit(sol.NIT).
		Build()
}

func TestSincronizarActividades(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarActividadesBuilder(), sol)
	runSincronizacionTest(t, "SincronizarActividades", req, service.SincronizarActividades)
}

func TestSincronizarListaActividadesDocumentoSector(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaActividadesDocumentoSectorBuilder(), sol)
	runSincronizacionTest(t, "SincronizarListaActividadesDocumentoSector", req, service.SincronizarListaActividadesDocumentoSector)
}

func TestSincronizarListaLeyendasFactura(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaLeyendasFacturaBuilder(), sol)
	runSincronizacionTest(t, "SincronizarListaLeyendasFactura", req, service.SincronizarListaLeyendasFactura)
}

func TestSincronizarListaMensajesServicios(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaMensajesServiciosBuilder(), sol)
	runSincronizacionTest(t, "SincronizarListaMensajesServicios", req, service.SincronizarListaMensajesServicios)
}

func TestSincronizarListaProductosServicios(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaProductosServiciosBuilder(), sol)
	runSincronizacionTest(t, "SincronizarListaProductosServicios", req, service.SincronizarListaProductosServicios)
}

func TestSincronizarParametricaEventosSignificativos(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaEventosSignificativosBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaEventosSignificativos", req, service.SincronizarParametricaEventosSignificativos)
}

func TestSincronizarParametricaMotivoAnulacion(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaMotivoAnulacionBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaMotivoAnulacion", req, service.SincronizarParametricaMotivoAnulacion)
}

func TestSincronizarParametricaPaisOrigen(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaPaisOrigenBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaPaisOrigen", req, service.SincronizarParametricaPaisOrigen)
}

func TestSincronizarParametricaTipoDocumentoIdentidad(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoDocumentoIdentidadBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoIdentidad", req, service.SincronizarParametricaTipoDocumentoIdentidad)
}

func TestSincronizarParametricaTipoDocumentoSector(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoDocumentoSectorBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoSector", req, service.SincronizarParametricaTipoDocumentoSector)
}

func TestSincronizarParametricaTipoEmision(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoEmisionBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTipoEmision", req, service.SincronizarParametricaTipoEmision)
}

func TestSincronizarParametricaTipoHabitacion(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoHabitacionBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTipoHabitacion", req, service.SincronizarParametricaTipoHabitacion)
}

func TestSincronizarParametricaTipoMetodoPago(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoMetodoPagoBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTipoMetodoPago", req, service.SincronizarParametricaTipoMetodoPago)
}

func TestSincronizarParametricaTipoMoneda(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoMonedaBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTipoMoneda", req, service.SincronizarParametricaTipoMoneda)
}

func TestSincronizarParametricaTipoPuntoVenta(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoPuntoVentaBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTipoPuntoVenta", req, service.SincronizarParametricaTipoPuntoVenta)
}

func TestSincronizarParametricaTiposFactura(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTiposFacturaBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaTiposFactura", req, service.SincronizarParametricaTiposFactura)
}

func TestSincronizarParametricaUnidadMedida(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	sol := getCommonRequest(t)
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaUnidadMedidaBuilder(), sol)
	runSincronizacionTest(t, "SincronizarParametricaUnidadMedida", req, service.SincronizarParametricaUnidadMedida)
}
