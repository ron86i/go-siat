package sincronizacion_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/protocols/sincronizacion"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_CatalogosParametricos(t *testing.T) {
	var operacion string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if !strings.Contains(string(b), "<ns:"+operacion+">") || !strings.Contains(string(b), "<codigoPuntoVenta>0</codigoPuntoVenta>") {
			t.Errorf("solicitud: %s", b)
		}
		fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><%sResponse><RespuestaListaParametricas><mensajesList><codigo>991</codigo><descripcion>mensaje</descripcion></mensajesList><transaccion>false</transaccion><listaCodigos><codigoClasificador>4</codigoClasificador><descripcion>valor</descripcion></listaCodigos></RespuestaListaParametricas></%sResponse></soap:Body></soap:Envelope>`, operacion, operacion)
	}))
	defer server.Close()

	c, err := sincronizacion.NewClient(sincronizacion.Config{
		BaseURL:        server.URL,
		Token:          "token",
		Nit:            123,
		CodigoSistema:  "sistema",
		CodigoAmbiente: 2,
	})
	require.NoError(t, err)

	cases := []struct {
		name string
		call func(context.Context, sincronizacion.Solicitud) (*sincronizacion.ResultadoParametricas, error)
	}{
		{"sincronizarListaMensajesServicios", c.SincronizarListaMensajesServicios},
		{"sincronizarParametricaEventosSignificativos", c.SincronizarParametricaEventosSignificativos},
		{"sincronizarParametricaMotivoAnulacion", c.SincronizarParametricaMotivoAnulacion},
		{"sincronizarParametricaPaisOrigen", c.SincronizarParametricaPaisOrigen},
		{"sincronizarParametricaTipoDocumentoIdentidad", c.SincronizarParametricaTipoDocumentoIdentidad},
		{"sincronizarParametricaTipoDocumentoSector", c.SincronizarParametricaTipoDocumentoSector},
		{"sincronizarParametricaTipoEmision", c.SincronizarParametricaTipoEmision},
		{"sincronizarParametricaTipoHabitacion", c.SincronizarParametricaTipoHabitacion},
		{"sincronizarParametricaTipoMetodoPago", c.SincronizarParametricaTipoMetodoPago},
		{"sincronizarParametricaTipoMoneda", c.SincronizarParametricaTipoMoneda},
		{"sincronizarParametricaTipoPuntoVenta", c.SincronizarParametricaTipoPuntoVenta},
		{"sincronizarParametricaTiposFactura", c.SincronizarParametricaTiposFactura},
		{"sincronizarParametricaUnidadMedida", c.SincronizarParametricaUnidadMedida},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			operacion = tc.name
			punto := 0
			got, err := tc.call(t.Context(), sincronizacion.Solicitud{
				Cuis:             "CUIS-123",
				CodigoSucursal:   0,
				CodigoPuntoVenta: &punto,
			})
			require.NoError(t, err)
			require.NotNil(t, got)
			require.NotNil(t, got.Transaccion)
			assert.False(t, *got.Transaccion)
			require.Len(t, got.Mensajes, 1)
			assert.Equal(t, 991, got.Mensajes[0].Codigo)
			require.Len(t, got.Codigos, 1)
			assert.Equal(t, 4, got.Codigos[0].CodigoClasificador)
			assert.Equal(t, "valor", got.Codigos[0].Descripcion)
		})
	}
}

func TestClient_CatalogosEspeciales(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		s := string(b)
		if strings.Contains(s, "<ns:sincronizarActividades>") {
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><sincronizarActividadesResponse><RespuestaListaActividades><transaccion>true</transaccion><listaActividades><codigoCaeb>620100</codigoCaeb><descripcion>Programacion</descripcion><tipoActividad>P</tipoActividad></listaActividades></RespuestaListaActividades></sincronizarActividadesResponse></soap:Body></soap:Envelope>`)
		} else if strings.Contains(s, "<ns:sincronizarFechaHora>") {
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><sincronizarFechaHoraResponse><RespuestaFechaHora><transaccion>true</transaccion><fechaHora>2026-10-31T16:00:00.000</fechaHora></RespuestaFechaHora></sincronizarFechaHoraResponse></soap:Body></soap:Envelope>`)
		} else if strings.Contains(s, "<ns:sincronizarListaActividadesDocumentoSector>") {
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><sincronizarListaActividadesDocumentoSectorResponse><RespuestaListaActividadesDocumentoSector><transaccion>true</transaccion><listaActividadesDocumentoSector><codigoActividad>620100</codigoActividad><codigoDocumentoSector>1</codigoDocumentoSector><tipoDocumentoSector>Factura Compra Venta</tipoDocumentoSector></listaActividadesDocumentoSector></RespuestaListaActividadesDocumentoSector></sincronizarListaActividadesDocumentoSectorResponse></soap:Body></soap:Envelope>`)
		} else if strings.Contains(s, "<ns:sincronizarListaLeyendasFactura>") {
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><sincronizarListaLeyendasFacturaResponse><RespuestaListaParametricasLeyendas><transaccion>true</transaccion><listaLeyendas><codigoActividad>620100</codigoActividad><descripcionLeyenda>Ley 453</descripcionLeyenda></listaLeyendas></RespuestaListaParametricasLeyendas></sincronizarListaLeyendasFacturaResponse></soap:Body></soap:Envelope>`)
		} else if strings.Contains(s, "<ns:sincronizarListaProductosServicios>") {
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><sincronizarListaProductosServiciosResponse><RespuestaListaProductos><transaccion>true</transaccion><listaCodigos><codigoActividad>620100</codigoActividad><codigoProducto>83141</codigoProducto><descripcionProducto>Software</descripcionProducto><nandina>00</nandina></listaCodigos></RespuestaListaProductos></sincronizarListaProductosServiciosResponse></soap:Body></soap:Envelope>`)
		} else if strings.Contains(s, "<ns:verificarComunicacion>") {
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarComunicacionResponse><return><transaccion>true</transaccion><mensajesList><codigo>926</codigo><descripcion>COMUNICACION EXITOSA</descripcion></mensajesList></return></verificarComunicacionResponse></soap:Body></soap:Envelope>`)
		}
	}))
	defer server.Close()

	c, err := sincronizacion.NewClient(sincronizacion.Config{
		BaseURL:        server.URL,
		Token:          "token",
		Nit:            123,
		CodigoSistema:  "sistema",
		CodigoAmbiente: 2,
	})
	require.NoError(t, err)
	ctx := t.Context()
	req := sincronizacion.Solicitud{Cuis: "CUIS-TEST"}

	// 1. Actividades
	act, err := c.SincronizarActividades(ctx, req)
	require.NoError(t, err)
	require.Len(t, act.Actividades, 1)
	assert.Equal(t, "620100", act.Actividades[0].CodigoCaeb)

	// 2. Fecha y Hora
	fh, err := c.SincronizarFechaHora(ctx, req)
	require.NoError(t, err)
	assert.Equal(t, "2026-10-31T16:00:00.000", fh.FechaHora)

	// 3. Actividad - Documento Sector
	ads, err := c.SincronizarListaActividadesDocumentoSector(ctx, req)
	require.NoError(t, err)
	require.Len(t, ads.Actividades, 1)
	assert.Equal(t, 1, ads.Actividades[0].CodigoDocumentoSector)

	// 4. Leyendas
	ley, err := c.SincronizarListaLeyendasFactura(ctx, req)
	require.NoError(t, err)
	require.Len(t, ley.Leyendas, 1)
	assert.Equal(t, "Ley 453", ley.Leyendas[0].DescripcionLeyenda)

	// 5. Productos
	prod, err := c.SincronizarListaProductosServicios(ctx, req)
	require.NoError(t, err)
	require.Len(t, prod.Productos, 1)
	assert.Equal(t, int64(83141), prod.Productos[0].CodigoProducto)

	// 6. Verificar Comunicación
	com, err := c.VerificarComunicacion(ctx)
	require.NoError(t, err)
	require.NotNil(t, com.Transaccion)
	assert.True(t, *com.Transaccion)
	require.Len(t, com.Mensajes, 1)
	assert.Equal(t, 926, com.Mensajes[0].Codigo)
}

func TestClient_ValidacionesSolicitud(t *testing.T) {
	c, err := sincronizacion.NewClient(sincronizacion.Config{
		BaseURL:        "http://localhost",
		Token:          "token",
		Nit:            123,
		CodigoSistema:  "sistema",
		CodigoAmbiente: 2,
	})
	require.NoError(t, err)
	ctx := t.Context()

	// CUIS vacío
	_, err = c.SincronizarActividades(ctx, sincronizacion.Solicitud{Cuis: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "CUIS son obligatorios")

	// Sucursal negativa
	_, err = c.SincronizarActividades(ctx, sincronizacion.Solicitud{Cuis: "CUIS", CodigoSucursal: -1})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sucursal >= 0")

	// Punto de venta negativo
	pv := -1
	_, err = c.SincronizarActividades(ctx, sincronizacion.Solicitud{Cuis: "CUIS", CodigoPuntoVenta: &pv})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "incompleta o inválida")
}
