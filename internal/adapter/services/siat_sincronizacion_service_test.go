package services_test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/internal/core/ports"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/stretchr/testify/require"
)

// sincronizacionIntegrationContext obtiene un CUIS una vez para todos los
// catálogos. SIAT_CUIS permite reutilizar uno vigente; si no existe, se genera
// con SIAT_CODIGO_MODALIDAD.
func sincronizacionIntegrationContext(t *testing.T) (context.Context, ports.SiatSincronizacionService, int, int, string) {
	t.Helper()
	cfg := siatRecepcionComprasIntegrationConfig(t)
	// Los catálogos son consultas de lectura y se ejecutan en bloque: no se
	// guardan sus XML SOAP para evitar llenar .siat-traces. Las pruebas de
	// recepción conservan el cliente trazable para diagnosticar fallos SIAT.
	cfg.HTTPClient = nil
	client, err := siat.New(cfg)
	require.NoError(t, err, "no se pudo crear el cliente SIAT")

	sucursal := codigoSIATDesdeEntorno(t, "SIAT_CODIGO_SUCURSAL", 0)
	puntoVenta := codigoSIATDesdeEntorno(t, "SIAT_CODIGO_PUNTO_VENTA", 0)
	cuis := strings.TrimSpace(os.Getenv("SIAT_CUIS"))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	t.Cleanup(cancel)
	if cuis == "" {
		modalidadRaw := strings.TrimSpace(os.Getenv("SIAT_CODIGO_MODALIDAD"))
		if modalidadRaw == "" {
			t.Skip("prueba de integración omitida; defina SIAT_CUIS o SIAT_CODIGO_MODALIDAD")
		}
		modalidad, err := strconv.Atoi(modalidadRaw)
		require.NoError(t, err, "SIAT_CODIGO_MODALIDAD debe ser un entero válido")
		response, err := client.Codigos().SolicitudCuis(ctx, models.NewCuisBuilder().
			WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCodigoModalidad(modalidad).Build())
		require.NoError(t, err, "no se pudo solicitar CUIS para sincronización")
		content, err := response.GetContent()
		require.NoError(t, err, "SIAT devolvió un SOAP Fault al solicitar CUIS")
		cuis = strings.TrimSpace(content.RespuestaCuis.Codigo)
		// El código 980 no es un error operativo: SIAT devuelve el CUIS vigente
		// con Transaccion=false cuando ya existe uno para ese punto de venta.
		if !content.RespuestaCuis.Transaccion {
			require.True(t, respuestaCUISYaVigente(content.RespuestaCuis.MensajesList),
				"SIAT rechazó la solicitud CUIS: %+v", content.RespuestaCuis.MensajesList)
		}
		require.NotEmpty(t, cuis, "SIAT no devolvió un CUIS")
	}
	return ctx, client.Sincronizacion(), sucursal, puntoVenta, cuis
}

func respuestaCUISYaVigente(mensajes any) bool {
	value := reflect.ValueOf(mensajes)
	if value.Kind() != reflect.Slice {
		return false
	}
	for index := 0; index < value.Len(); index++ {
		mensaje := value.Index(index)
		if mensaje.Kind() == reflect.Pointer {
			mensaje = mensaje.Elem()
		}
		if mensaje.Kind() != reflect.Struct {
			continue
		}
		codigo := mensaje.FieldByName("Codigo")
		if codigo.IsValid() && codigo.Kind() == reflect.Int && codigo.Int() == 980 {
			return true
		}
	}
	return false
}

func codigoSIATDesdeEntorno(t *testing.T, variable string, fallback int) int {
	t.Helper()
	value := strings.TrimSpace(os.Getenv(variable))
	if value == "" {
		return fallback
	}
	code, err := strconv.Atoi(value)
	require.NoError(t, err, "%s debe ser un entero válido", variable)
	require.GreaterOrEqual(t, code, 0, "%s no puede ser negativo", variable)
	return code
}

// Las respuestas de catálogo difieren en su tipo concreto, pero todas poseen
// GetContent y un indicador Transaccion dentro de su contenido.
func validarRespuestaSincronizacion(response any) error {
	if response == nil {
		return fmt.Errorf("SIAT no devolvió respuesta")
	}
	method := reflect.ValueOf(response).MethodByName("GetContent")
	if !method.IsValid() {
		return fmt.Errorf("respuesta SIAT sin GetContent")
	}
	values := method.Call(nil)
	if len(values) != 2 {
		return fmt.Errorf("respuesta SIAT con forma inválida")
	}
	if !values[1].IsNil() {
		return values[1].Interface().(error)
	}
	if transaccion, found := encontrarTransaccion(values[0]); !found {
		return fmt.Errorf("respuesta SIAT sin campo Transaccion")
	} else if !transaccion {
		return fmt.Errorf("SIAT rechazó la sincronización")
	}
	return nil
}

func encontrarTransaccion(value reflect.Value) (bool, bool) {
	if value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return false, false
		}
		return encontrarTransaccion(value.Elem())
	}
	if value.Kind() != reflect.Struct {
		return false, false
	}
	field := value.FieldByName("Transaccion")
	if field.IsValid() && field.Kind() == reflect.Bool {
		return field.Bool(), true
	}
	for index := 0; index < value.NumField(); index++ {
		if transaccion, found := encontrarTransaccion(value.Field(index)); found {
			return transaccion, true
		}
	}
	return false, false
}

func TestSiatSincronizacion_Catalogos(t *testing.T) {
	ctx, service, sucursal, puntoVenta, cuis := sincronizacionIntegrationContext(t)

	tests := []struct {
		name string
		call func() (any, error)
	}{
		{"actividades", func() (any, error) {
			return service.SincronizarActividades(ctx, models.NewSincronizarActividadesBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"actividades_documento_sector", func() (any, error) {
			return service.SincronizarListaActividadesDocumentoSector(ctx, models.NewSincronizarListaActividadesDocumentoSectorBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"leyendas_factura", func() (any, error) {
			return service.SincronizarListaLeyendasFactura(ctx, models.NewSincronizarListaLeyendasFacturaBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"mensajes_servicio", func() (any, error) {
			return service.SincronizarListaMensajesServicios(ctx, models.NewSincronizarListaMensajesServiciosBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"productos_servicios", func() (any, error) {
			return service.SincronizarListaProductosServicios(ctx, models.NewSincronizarListaProductosServiciosBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"eventos_significativos", func() (any, error) {
			return service.SincronizarParametricaEventosSignificativos(ctx, models.NewSincronizarParametricaEventosSignificativosBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"motivos_anulacion", func() (any, error) {
			return service.SincronizarParametricaMotivoAnulacion(ctx, models.NewSincronizarParametricaMotivoAnulacionBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"paises_origen", func() (any, error) {
			return service.SincronizarParametricaPaisOrigen(ctx, models.NewSincronizarParametricaPaisOrigenBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_documento_identidad", func() (any, error) {
			return service.SincronizarParametricaTipoDocumentoIdentidad(ctx, models.NewSincronizarParametricaTipoDocumentoIdentidadBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_documento_sector", func() (any, error) {
			return service.SincronizarParametricaTipoDocumentoSector(ctx, models.NewSincronizarParametricaTipoDocumentoSectorBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_emision", func() (any, error) {
			return service.SincronizarParametricaTipoEmision(ctx, models.NewSincronizarParametricaTipoEmisionBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_habitacion", func() (any, error) {
			return service.SincronizarParametricaTipoHabitacion(ctx, models.NewSincronizarParametricaTipoHabitacionBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_metodo_pago", func() (any, error) {
			return service.SincronizarParametricaTipoMetodoPago(ctx, models.NewSincronizarParametricaTipoMetodoPagoBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_moneda", func() (any, error) {
			return service.SincronizarParametricaTipoMoneda(ctx, models.NewSincronizarParametricaTipoMonedaBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_punto_venta", func() (any, error) {
			return service.SincronizarParametricaTipoPuntoVenta(ctx, models.NewSincronizarParametricaTipoPuntoVentaBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"tipos_factura", func() (any, error) {
			return service.SincronizarParametricaTiposFactura(ctx, models.NewSincronizarParametricaTiposFacturaBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
		{"unidades_medida", func() (any, error) {
			return service.SincronizarParametricaUnidadMedida(ctx, models.NewSincronizarParametricaUnidadMedidaBuilder().WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, err := test.call()
			require.NoError(t, err, "falló la solicitud de sincronización")
			if method := reflect.ValueOf(response).MethodByName("GetContent"); method.IsValid() {
				res := method.Call(nil)
				if len(res) > 0 && res[0].IsValid() {
					t.Logf("[%s] Content: %+v", test.name, res[0].Interface())
				}
			}
			require.NoError(t, validarRespuestaSincronizacion(response))
		})
	}
}

func TestSiatSincronizacion_FechaHora(t *testing.T) {
	ctx, service, sucursal, puntoVenta, cuis := sincronizacionIntegrationContext(t)
	response, err := service.SincronizarFechaHora(ctx, models.NewSincronizarFechaHoraBuilder().
		WithCodigoSucursal(sucursal).WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).Build())
	require.NoError(t, err, "falló la sincronización de fecha y hora")
	require.NoError(t, validarRespuestaSincronizacion(response))
}

func TestSiatSincronizacion_CuisInvalido(t *testing.T) {
	ctx, service, sucursal, puntoVenta, _ := sincronizacionIntegrationContext(t)
	cuisInvalido := "CUIS_INVALIDO_12345"

	response, err := service.SincronizarParametricaEventosSignificativos(ctx,
		models.NewSincronizarParametricaEventosSignificativosBuilder().
			WithCodigoSucursal(sucursal).
			WithCodigoPuntoVenta(puntoVenta).
			WithCuis(cuisInvalido).
			Build(),
	)
	require.NoError(t, err, "la llamada SOAP debe retornar la estructura de respuesta sin fallar en transporte")
	content, err := response.GetContent()
	require.NoError(t, err)

	t.Logf("[CUIS Invalido Real SIAT] Content: %+v", content)
}
