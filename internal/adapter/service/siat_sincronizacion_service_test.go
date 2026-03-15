package service_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"

	"github.com/ron86i/go-siat/internal/core/domain/siat/sincronizacion"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func runSincronizacionTest[V any, ReqType any](
	t *testing.T,
	name string,
	req ReqType,
	fn func(context.Context, config.Config, ReqType) (*V, error),
) {
	t.Run(name, func(t *testing.T) {
		godotenv.Load()

		config := config.Config{
			Token: os.Getenv("SIAT_TOKEN"),
		}

		_, err := siat.New(os.Getenv("SIAT_URL"), nil)
		if err != nil {
			t.Fatalf("No se pudo inicializar el cliente SIAT: %v", err)
		}
		// En sincronización, el servicio se obtiene de siatClient.Sincronizacion()
		// Pero aquí runSincronizacionTest recibe el método.
		// Ajustamos el llamador de este helper.

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
		Cuis:             "C2FC682C", // CUIS de prueba
	}
}

func buildSyncReq[T any, R any](b *models.SincronizacionBuilder[T, R], sol sincronizacion.SolicitudSincronizacion) R {
	return b.WithCodigoAmbiente(sol.CodigoAmbiente).
		WithCodigoPuntoVenta(sol.CodigoPuntoVenta).
		WithCodigoSistema(sol.CodigoSistema).
		WithCodigoSucursal(sol.CodigoSucursal).
		WithCuis(sol.Cuis).
		WithNit(sol.NIT).
		Build()
}

// TestSincronizacionCompleta ejecuta secuencialmente todas las solicitudes de sincronización.
// Asegura que el sistema pueda obtener los catálogos oficiales actualizados del SIAT.
func TestSincronizacionCompleta(t *testing.T) {
	godotenv.Load()
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.Sincronizacion()
	solicitud := getCommonRequest(t)

	runSincronizacionTest(t, "SincronizarActividades", buildSyncReq(models.Sincronizacion().NewSincronizarActividadesBuilder(), solicitud), service.SincronizarActividades)
	runSincronizacionTest(t, "SincronizarListaActividadesDocumentoSector", buildSyncReq(models.Sincronizacion().NewSincronizarListaActividadesDocumentoSectorBuilder(), solicitud), service.SincronizarListaActividadesDocumentoSector)
	runSincronizacionTest(t, "SincronizarListaLeyendasFactura", buildSyncReq(models.Sincronizacion().NewSincronizarListaLeyendasFacturaBuilder(), solicitud), service.SincronizarListaLeyendasFactura)
	runSincronizacionTest(t, "SincronizarListaMensajesServicios", buildSyncReq(models.Sincronizacion().NewSincronizarListaMensajesServiciosBuilder(), solicitud), service.SincronizarListaMensajesServicios)
	runSincronizacionTest(t, "SincronizarListaProductosServicios", buildSyncReq(models.Sincronizacion().NewSincronizarListaProductosServiciosBuilder(), solicitud), service.SincronizarListaProductosServicios)
	runSincronizacionTest(t, "SincronizarParametricaEventosSignificativos", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaEventosSignificativosBuilder(), solicitud), service.SincronizarParametricaEventosSignificativos)
	runSincronizacionTest(t, "SincronizarParametricaMotivoAnulacion", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaMotivoAnulacionBuilder(), solicitud), service.SincronizarParametricaMotivoAnulacion)
	runSincronizacionTest(t, "SincronizarParametricaPaisOrigen", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaPaisOrigenBuilder(), solicitud), service.SincronizarParametricaPaisOrigen)
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoIdentidad", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTipoDocumentoIdentidadBuilder(), solicitud), service.SincronizarParametricaTipoDocumentoIdentidad)
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoSector", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTipoDocumentoSectorBuilder(), solicitud), service.SincronizarParametricaTipoDocumentoSector)
	runSincronizacionTest(t, "SincronizarParametricaTipoEmision", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTipoEmisionBuilder(), solicitud), service.SincronizarParametricaTipoEmision)
	runSincronizacionTest(t, "SincronizarParametricaTipoHabitacion", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTipoHabitacionBuilder(), solicitud), service.SincronizarParametricaTipoHabitacion)
	runSincronizacionTest(t, "SincronizarParametricaTipoMetodoPago", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTipoMetodoPagoBuilder(), solicitud), service.SincronizarParametricaTipoMetodoPago)
	runSincronizacionTest(t, "SincronizarParametricaTipoMoneda", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTipoMonedaBuilder(), solicitud), service.SincronizarParametricaTipoMoneda)
	runSincronizacionTest(t, "SincronizarParametricaTipoPuntoVenta", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTipoPuntoVentaBuilder(), solicitud), service.SincronizarParametricaTipoPuntoVenta)
	runSincronizacionTest(t, "SincronizarParametricaTiposFactura", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaTiposFacturaBuilder(), solicitud), service.SincronizarParametricaTiposFactura)
	runSincronizacionTest(t, "SincronizarParametricaUnidadMedida", buildSyncReq(models.Sincronizacion().NewSincronizarParametricaUnidadMedidaBuilder(), solicitud), service.SincronizarParametricaUnidadMedida)
}
