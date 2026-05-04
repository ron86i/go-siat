package services_test

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/adapter/services"
	"github.com/ron86i/go-siat/internal/core/ports"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load()
	}
	os.Exit(m.Run())
}

func runSincronizacionTest[V any, ReqType any](
	t *testing.T,
	name string,
	req ReqType,
	fn func(context.Context, ports.Config, ReqType) (*V, error),
) {
	config := ports.Config{
		Token: os.Getenv("SIAT_TOKEN"),
	}

	resp, err := fn(context.Background(), config, req)
	if err != nil {
		t.Fatalf("Error en %s: %v", name, err)
	}

	assert.NotNil(t, resp)
	xmlData, _ := xml.MarshalIndent(resp, "", "  ")
	log.Printf("Resultado de %s: %s", name, string(xmlData))
}

func getSiatClient(t *testing.T) *siat.SiatServices {
	cfg := services.DefaultHTTPConfig()
	httpClient := services.NewHTTPClient(cfg)
	siatClient, err := siat.New(os.Getenv("SIAT_URL"), httpClient)
	if err != nil {
		t.Fatalf("No se pudo inicializar el cliente SIAT: %v", err)
	}
	return siatClient
}

func buildSincronizacion[T any, R any](builder models.SincronizacionBuilder[T, R]) R {
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))

	return builder.
		WithCodigoAmbiente(codAmbiente).
		WithNit(nit).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(1).
		WithCuis("B3F82775").
		Build()
}

// --- Tests Individuales ---

func TestSincronizarActividades(t *testing.T) {
	sincronizarActividades(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarListaActividadesDocumentoSector(t *testing.T) {
	sincronizarListaActividadesDocumentoSector(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarListaLeyendasFactura(t *testing.T) {
	sincronizarListaLeyendasFactura(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarListaMensajesServicios(t *testing.T) {
	sincronizarListaMensajesServicios(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarListaProductosServicios(t *testing.T) {
	sincronizarListaProductosServicios(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaEventosSignificativos(t *testing.T) {
	sincronizarParametricaEventosSignificativos(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaMotivoAnulacion(t *testing.T) {
	sincronizarParametricaMotivoAnulacion(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaPaisOrigen(t *testing.T) {
	sincronizarParametricaPaisOrigen(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTipoDocumentoIdentidad(t *testing.T) {
	sincronizarParametricaTipoDocumentoIdentidad(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTipoDocumentoSector(t *testing.T) {
	sincronizarParametricaTipoDocumentoSector(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTipoEmision(t *testing.T) {
	sincronizarParametricaTipoEmision(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTipoHabitacion(t *testing.T) {
	sincronizarParametricaTipoHabitacion(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTipoMetodoPago(t *testing.T) {
	sincronizarParametricaTipoMetodoPago(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTipoMoneda(t *testing.T) {
	sincronizarParametricaTipoMoneda(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTipoPuntoVenta(t *testing.T) {
	sincronizarParametricaTipoPuntoVenta(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaTiposFactura(t *testing.T) {
	sincronizarParametricaTiposFactura(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarParametricaUnidadMedida(t *testing.T) {
	sincronizarParametricaUnidadMedida(t, getSiatClient(t).Sincronizacion())
}

func TestSincronizarFechaHora(t *testing.T) {
	sincronizarFechaHora(t, getSiatClient(t).Sincronizacion())
}

// --- Lógica de Negocio de los Tests ---

func sincronizarActividades(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarActividadesBuilder())
	runSincronizacionTest(t, "SincronizarActividades", req, service.SincronizarActividades)
}

func sincronizarListaActividadesDocumentoSector(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaActividadesDocumentoSectorBuilder())
	runSincronizacionTest(t, "SincronizarListaActividadesDocumentoSector", req, service.SincronizarListaActividadesDocumentoSector)
}

func sincronizarListaLeyendasFactura(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaLeyendasFacturaBuilder())
	runSincronizacionTest(t, "SincronizarListaLeyendasFactura", req, service.SincronizarListaLeyendasFactura)
}

func sincronizarListaMensajesServicios(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaMensajesServiciosBuilder())
	runSincronizacionTest(t, "SincronizarListaMensajesServicios", req, service.SincronizarListaMensajesServicios)
}

func sincronizarListaProductosServicios(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarListaProductosServiciosBuilder())
	runSincronizacionTest(t, "SincronizarListaProductosServicios", req, service.SincronizarListaProductosServicios)
}

func sincronizarParametricaEventosSignificativos(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaEventosSignificativosBuilder())
	runSincronizacionTest(t, "SincronizarParametricaEventosSignificativos", req, service.SincronizarParametricaEventosSignificativos)
}

func sincronizarParametricaMotivoAnulacion(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaMotivoAnulacionBuilder())
	runSincronizacionTest(t, "SincronizarParametricaMotivoAnulacion", req, service.SincronizarParametricaMotivoAnulacion)
}

func sincronizarParametricaPaisOrigen(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaPaisOrigenBuilder())
	runSincronizacionTest(t, "SincronizarParametricaPaisOrigen", req, service.SincronizarParametricaPaisOrigen)
}

func sincronizarParametricaTipoDocumentoIdentidad(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoDocumentoIdentidadBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoIdentidad", req, service.SincronizarParametricaTipoDocumentoIdentidad)
}

func sincronizarParametricaTipoDocumentoSector(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoDocumentoSectorBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoSector", req, service.SincronizarParametricaTipoDocumentoSector)
}

func sincronizarParametricaTipoEmision(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoEmisionBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTipoEmision", req, service.SincronizarParametricaTipoEmision)
}

func sincronizarParametricaTipoHabitacion(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoHabitacionBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTipoHabitacion", req, service.SincronizarParametricaTipoHabitacion)
}

func sincronizarParametricaTipoMetodoPago(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoMetodoPagoBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTipoMetodoPago", req, service.SincronizarParametricaTipoMetodoPago)
}

func sincronizarParametricaTipoMoneda(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoMonedaBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTipoMoneda", req, service.SincronizarParametricaTipoMoneda)
}

func sincronizarParametricaTipoPuntoVenta(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTipoPuntoVentaBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTipoPuntoVenta", req, service.SincronizarParametricaTipoPuntoVenta)
}

func sincronizarParametricaTiposFactura(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaTiposFacturaBuilder())
	runSincronizacionTest(t, "SincronizarParametricaTiposFactura", req, service.SincronizarParametricaTiposFactura)
}

func sincronizarParametricaUnidadMedida(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarParametricaUnidadMedidaBuilder())
	runSincronizacionTest(t, "SincronizarParametricaUnidadMedida", req, service.SincronizarParametricaUnidadMedida)
}

func sincronizarFechaHora(t *testing.T, service ports.SiatSincronizacionService) {
	req := buildSincronizacion(models.Sincronizacion().NewSincronizarFechaHoraBuilder())
	runSincronizacionTest(t, "SincronizarFechaHora", req, service.SincronizarFechaHora)
}

// TestSincronizacionAll ejecuta secuencialmente todos los tests del servicio de sincronización.
func TestSincronizacionAll(t *testing.T) {
	// Inicializar el cliente una sola vez para habilitar el pooling de conexiones
	siatClient := getSiatClient(t)
	service := siatClient.Sincronizacion()

	// Ejecución secuencial de las iteraciones
	for i := 0; i < 50; i++ {
		// Usamos un subtest para agrupar los resultados de esta iteración
		t.Run(fmt.Sprintf("Iteracion_%d", i), func(t *testing.T) {
			t.Run("SincronizarActividades", func(t *testing.T) { sincronizarActividades(t, service) })
			t.Run("SincronizarListaActividadesDocumentoSector", func(t *testing.T) { sincronizarListaActividadesDocumentoSector(t, service) })
			t.Run("SincronizarListaLeyendasFactura", func(t *testing.T) { sincronizarListaLeyendasFactura(t, service) })
			t.Run("SincronizarListaMensajesServicios", func(t *testing.T) { sincronizarListaMensajesServicios(t, service) })
			t.Run("SincronizarListaProductosServicios", func(t *testing.T) { sincronizarListaProductosServicios(t, service) })
			t.Run("SincronizarParametricaEventosSignificativos", func(t *testing.T) { sincronizarParametricaEventosSignificativos(t, service) })
			t.Run("SincronizarParametricaMotivoAnulacion", func(t *testing.T) { sincronizarParametricaMotivoAnulacion(t, service) })
			t.Run("SincronizarParametricaPaisOrigen", func(t *testing.T) { sincronizarParametricaPaisOrigen(t, service) })
			t.Run("SincronizarParametricaTipoDocumentoIdentidad", func(t *testing.T) { sincronizarParametricaTipoDocumentoIdentidad(t, service) })
			t.Run("SincronizarParametricaTipoDocumentoSector", func(t *testing.T) { sincronizarParametricaTipoDocumentoSector(t, service) })
			t.Run("SincronizarParametricaTipoEmision", func(t *testing.T) { sincronizarParametricaTipoEmision(t, service) })
			t.Run("SincronizarParametricaTipoHabitacion", func(t *testing.T) { sincronizarParametricaTipoHabitacion(t, service) })
			t.Run("SincronizarParametricaTipoMetodoPago", func(t *testing.T) { sincronizarParametricaTipoMetodoPago(t, service) })
			t.Run("SincronizarParametricaTipoMoneda", func(t *testing.T) { sincronizarParametricaTipoMoneda(t, service) })
			t.Run("SincronizarParametricaTipoPuntoVenta", func(t *testing.T) { sincronizarParametricaTipoPuntoVenta(t, service) })
			t.Run("SincronizarParametricaTiposFactura", func(t *testing.T) { sincronizarParametricaTiposFactura(t, service) })
			t.Run("SincronizarParametricaUnidadMedida", func(t *testing.T) { sincronizarParametricaUnidadMedida(t, service) })
			t.Run("SincronizarFechaHora", func(t *testing.T) { sincronizarFechaHora(t, service) })
		})
	}
}
