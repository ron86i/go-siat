package siat_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/internal/adapter/service/siat"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/sincronizacion"
	"github.com/ron86i/go-siat/internal/core/util"
	"github.com/stretchr/testify/assert"
)

func runSincronizacionTest[K any, V any](
	t *testing.T,
	name string,
	req K,
	fn func(context.Context, facturacion.Config, K) (*V, error),
) {
	t.Run(name, func(t *testing.T) {
		godotenv.Load()

		envs := map[string]string{
			"SIAT_URL": os.Getenv("SIAT_URL"),
		}
		config := facturacion.Config{
			Token: os.Getenv("SIAT_TOKEN"),
		}

		_, err := siat.NewSiatSincronizacionService(envs)
		if err != nil {
			t.Fatalf("No se pudo inicializar el servicio: %v", err)
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
	nit, _ := util.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := util.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))

	return sincronizacion.SolicitudSincronizacion{
		CodigoAmbiente:   codAmbiente,
		NIT:              nit,
		CodigoSistema:    os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoSucursal:   0,
		CodigoPuntoVenta: 0,
		Cuis:             "C2FC682C", // CUIS de prueba
	}
}

func TestSincronizacionCompleta(t *testing.T) {
	godotenv.Load()
	envs := map[string]string{"SIAT_URL": os.Getenv("SIAT_URL")}
	service, _ := siat.NewSiatSincronizacionService(envs)
	solicitud := getCommonRequest(t)

	runSincronizacionTest(t, "SincronizarActividades", sincronizacion.SincronizarActividades{SolicitudSincronizacion: solicitud}, service.SincronizarActividades)
	runSincronizacionTest(t, "SincronizarListaActividadesDocumentoSector", sincronizacion.SincronizarListaActividadesDocumentoSector{SolicitudSincronizacion: solicitud}, service.SincronizarListaActividadesDocumentoSector)
	runSincronizacionTest(t, "SincronizarListaLeyendasFactura", sincronizacion.SincronizarListaLeyendasFactura{SolicitudSincronizacion: solicitud}, service.SincronizarListaLeyendasFactura)
	runSincronizacionTest(t, "SincronizarListaMensajesServicios", sincronizacion.SincronizarListaMensajesServicios{SolicitudSincronizacion: solicitud}, service.SincronizarListaMensajesServicios)
	runSincronizacionTest(t, "SincronizarListaProductosServicios", sincronizacion.SincronizarListaProductosServicios{SolicitudSincronizacion: solicitud}, service.SincronizarListaProductosServicios)
	runSincronizacionTest(t, "SincronizarParametricaEventosSignificativos", sincronizacion.SincronizarParametricaEventosSignificativos{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaEventosSignificativos)
	runSincronizacionTest(t, "SincronizarParametricaMotivoAnulacion", sincronizacion.SincronizarParametricaMotivoAnulacion{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaMotivoAnulacion)
	runSincronizacionTest(t, "SincronizarParametricaPaisOrigen", sincronizacion.SincronizarParametricaPaisOrigen{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaPaisOrigen)
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoIdentidad", sincronizacion.SincronizarParametricaTipoDocumentoIdentidad{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTipoDocumentoIdentidad)
	runSincronizacionTest(t, "SincronizarParametricaTipoDocumentoSector", sincronizacion.SincronizarParametricaTipoDocumentoSector{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTipoDocumentoSector)
	runSincronizacionTest(t, "SincronizarParametricaTipoEmision", sincronizacion.SincronizarParametricaTipoEmision{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTipoEmision)
	runSincronizacionTest(t, "SincronizarParametricaTipoHabitacion", sincronizacion.SincronizarParametricaTipoHabitacion{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTipoHabitacion)
	runSincronizacionTest(t, "SincronizarParametricaTipoMetodoPago", sincronizacion.SincronizarParametricaTipoMetodoPago{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTipoMetodoPago)
	runSincronizacionTest(t, "SincronizarParametricaTipoMoneda", sincronizacion.SincronizarParametricaTipoMoneda{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTipoMoneda)
	runSincronizacionTest(t, "SincronizarParametricaTipoPuntoVenta", sincronizacion.SincronizarParametricaTipoPuntoVenta{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTipoPuntoVenta)
	runSincronizacionTest(t, "SincronizarParametricaTiposFactura", sincronizacion.SincronizarParametricaTiposFactura{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaTiposFactura)
	runSincronizacionTest(t, "SincronizarParametricaUnidadMedida", sincronizacion.SincronizarParametricaUnidadMedida{SolicitudSincronizacion: solicitud}, service.SincronizarParametricaUnidadMedida)
}
