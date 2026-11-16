package facturacion_test

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/protocols/facturacion"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_RecepcionFactura(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.URL.Path != "/sfd/ServicioFacturacionCompraVenta" || r.Header.Get("apiKey") != "TokenApi token" {
			t.Errorf("solicitud inesperada: %s", r.URL.Path)
		}
		b, _ := io.ReadAll(r.Body)
		for _, expected := range []string{"<codigoPuntoVenta>0</codigoPuntoVenta>", "<archivo>" + base64.StdEncoding.EncodeToString([]byte("archivo")) + "</archivo>", "<fechaEnvio>2026-10-31T12:30:00.000</fechaEnvio>"} {
			if !strings.Contains(string(b), expected) {
				t.Errorf("falta %s en %s", expected, b)
			}
		}
		w.Header().Set("Content-Type", "application/xml")
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><recepcionFacturaResponse><RespuestaServicioFacturacion><codigoEstado>901</codigoEstado><mensajesList><codigo>123</codigo><descripcion>rechazo</descripcion></mensajesList><transaccion>false</transaccion></RespuestaServicioFacturacion></recepcionFacturaResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL + "/sfd", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	result, err := client.RecepcionFactura(context.Background(), facturacion.SolicitudRecepcionFactura{
		CodigoDocumentoSector: 1,
		CodigoEmision:         1,
		CodigoModalidad:       2,
		CodigoPuntoVenta:      new(0),
		Cufd:                  "cufd",
		Cuis:                  "cuis",
		TipoFacturaDocumento:  1,
		Archivo:               []byte("archivo"),
		FechaEnvio:            time.Date(2026, 10, 31, 12, 30, 0, 0, time.UTC),
		HashArchivo:           "hash",
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, result.Transaccion)
	assert.False(t, *result.Transaccion)
	assert.Equal(t, 901, result.CodigoEstado)
	require.Len(t, result.Mensajes, 1)
	assert.Equal(t, 123, result.Mensajes[0].Codigo)
	assert.Equal(t, 1, calls)
}

func TestClient_SectoresConEndpointPropio(t *testing.T) {
	var paths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><recepcionFacturaResponse><RespuestaServicioFacturacion><transaccion>true</transaccion></RespuestaServicioFacturacion></recepcionFacturaResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	config := facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2}
	compraVenta, err := facturacion.NewClient(config)
	require.NoError(t, err)
	servicioBasico, err := facturacion.NewServicioBasicoClient(config)
	require.NoError(t, err)
	requestFor := func(sector int) facturacion.SolicitudRecepcionFactura {
		return facturacion.SolicitudRecepcionFactura{
			CodigoDocumentoSector: sector,
			CodigoEmision:         1,
			CodigoModalidad:       1,
			CodigoPuntoVenta:      new(0),
			Cufd:                  "cufd",
			Cuis:                  "cuis",
			TipoFacturaDocumento:  1,
			Archivo:               []byte("archivo"),
			FechaEnvio:            time.Date(2026, 10, 31, 12, 30, 0, 0, time.UTC),
			HashArchivo:           "hash",
		}
	}

	for _, tc := range []struct {
		client  *facturacion.Client
		sectors []int
		path    string
	}{
		{compraVenta, []int{1, 35, 41}, "/sfd/ServicioFacturacionCompraVenta"},
		{servicioBasico, []int{13, 40}, "/sfd/ServicioFacturacionServicioBasico"},
	} {
		for _, sector := range tc.sectors {
			_, err := tc.client.RecepcionFactura(t.Context(), requestFor(sector))
			require.NoErrorf(t, err, "sector %d", sector)
			require.Equal(t, tc.path, paths[len(paths)-1])
		}
	}

	requestCount := len(paths)
	_, err = compraVenta.RecepcionFactura(t.Context(), requestFor(13))
	require.Error(t, err)
	_, err = servicioBasico.RecepcionFactura(t.Context(), requestFor(35))
	require.Error(t, err)
	require.Len(t, paths, requestCount, "los sectores deben rechazarse antes del envío HTTP")
}

func TestClient_FaultYRespuestaIncompleta(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><soap:Fault><faultcode>server</faultcode><faultstring>error</faultstring></soap:Fault></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	_, err = client.RecepcionFactura(context.Background(), facturacion.SolicitudRecepcionFactura{
		CodigoDocumentoSector: 1,
		CodigoEmision:         1,
		CodigoModalidad:       2,
		Cufd:                  "cufd",
		Cuis:                  "cuis",
		TipoFacturaDocumento:  1,
		Archivo:               []byte("archivo"),
		FechaEnvio:            time.Now(),
		HashArchivo:           "hash",
	})
	var fault *facturacion.Fault
	require.True(t, errors.As(err, &fault))
	assert.Equal(t, "server", fault.Code)
}

func TestClient_TransaccionAusenteEsDistintaDeRechazo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><recepcionFacturaResponse><RespuestaServicioFacturacion><mensajesList><codigo>991</codigo><descripcion>error</descripcion></mensajesList></RespuestaServicioFacturacion></recepcionFacturaResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	result, err := client.RecepcionFactura(t.Context(), facturacion.SolicitudRecepcionFactura{
		CodigoDocumentoSector: 1,
		CodigoEmision:         1,
		CodigoModalidad:       2,
		Cufd:                  "cufd",
		Cuis:                  "cuis",
		TipoFacturaDocumento:  1,
		Archivo:               []byte("archivo"),
		FechaEnvio:            time.Now(),
		HashArchivo:           "hash",
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Nil(t, result.Transaccion)
	require.Len(t, result.Mensajes, 1)
	assert.Equal(t, 991, result.Mensajes[0].Codigo)
}

func TestClient_OperacionesFactura(t *testing.T) {
	for _, operation := range []struct {
		nombre, solicitud string
		motivo            bool
	}{
		{"verificacionEstadoFactura", "SolicitudServicioVerificacionEstadoFactura", false},
		{"anulacionFactura", "SolicitudServicioAnulacionFactura", true},
		{"reversionAnulacionFactura", "SolicitudServicioReversionAnulacionFactura", false},
	} {
		t.Run(operation.nombre, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Error(err)
					return
				}
				text := string(body)
				for _, part := range []string{"<ns:" + operation.nombre + ">", "<" + operation.solicitud + ">", "<codigoPuntoVenta>0</codigoPuntoVenta>", "<cuf>CUF</cuf>"} {
					if !strings.Contains(text, part) {
						t.Errorf("falta %s en %s", part, text)
					}
				}
				pos := strings.Index(text, "<codigoMotivo>")
				if operation.motivo && (pos < 0 || pos > strings.Index(text, "<cuf>")) {
					t.Errorf("motivo fuera de orden: %s", text)
				}
				if !operation.motivo && pos >= 0 {
					t.Errorf("motivo inesperado: %s", text)
				}
				fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><%sResponse><RespuestaServicioFacturacion><codigoEstado>905</codigoEstado><mensajesList><codigo>123</codigo><descripcion>observado</descripcion></mensajesList><transaccion>false</transaccion></RespuestaServicioFacturacion></%sResponse></soap:Body></soap:Envelope>`, operation.nombre, operation.nombre)
			}))
			defer server.Close()

			client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
			require.NoError(t, err)

			req := facturacion.SolicitudEstadoFactura{CodigoDocumentoSector: 1, CodigoEmision: 1, CodigoModalidad: 2, CodigoPuntoVenta: new(0), Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1, Cuf: "CUF"}
			var got *facturacion.RespuestaRecepcion
			switch operation.nombre {
			case "verificacionEstadoFactura":
				got, err = client.VerificacionEstadoFactura(t.Context(), req)
			case "anulacionFactura":
				got, err = client.AnulacionFactura(t.Context(), facturacion.SolicitudAnulacionFactura{SolicitudEstadoFactura: req, CodigoMotivo: 1})
			default:
				got, err = client.ReversionAnulacionFactura(t.Context(), req)
			}
			require.NoError(t, err)
			require.NotNil(t, got)
			assert.Equal(t, 1, calls)
			require.NotNil(t, got.Transaccion)
			assert.False(t, *got.Transaccion)
			assert.Equal(t, 905, got.CodigoEstado)
			require.Len(t, got.Mensajes, 1)
			assert.Equal(t, 123, got.Mensajes[0].Codigo)
		})
	}
}

func TestClient_OperacionRespuestaIncorrecta(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><anulacionFacturaResponse><RespuestaServicioFacturacion><transaccion>true</transaccion></RespuestaServicioFacturacion></anulacionFacturaResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	_, err = client.VerificacionEstadoFactura(t.Context(), facturacion.SolicitudEstadoFactura{CodigoDocumentoSector: 1, CodigoEmision: 1, CodigoModalidad: 2, Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1, Cuf: "CUF"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sin resultado de verificacionEstadoFactura")
}

func TestClient_VerificarComunicacion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarComunicacionResponse><return><transaccion>true</transaccion><mensajesList><codigo>926</codigo><descripcion>COMUNICACION EXITOSA</descripcion></mensajesList></return></verificarComunicacionResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	resp, err := client.VerificarComunicacion(t.Context())
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Transaccion)
	assert.True(t, *resp.Transaccion)
	require.Len(t, resp.Mensajes, 1)
	assert.Equal(t, 926, resp.Mensajes[0].Codigo)
}

func TestClient_ValidacionesParametros(t *testing.T) {
	client, err := facturacion.NewClient(facturacion.Config{BaseURL: "http://localhost", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	ctx := t.Context()

	// Solicitud incompleta en RecepcionFactura
	_, err = client.RecepcionFactura(ctx, facturacion.SolicitudRecepcionFactura{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "incompleta o inválida")

	// Solicitud incompleta en VerificacionEstadoFactura
	_, err = client.VerificacionEstadoFactura(ctx, facturacion.SolicitudEstadoFactura{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "incompleta o inválida")

	// Solicitud incompleta en AnulacionFactura (motivo <= 0)
	_, err = client.AnulacionFactura(ctx, facturacion.SolicitudAnulacionFactura{
		SolicitudEstadoFactura: facturacion.SolicitudEstadoFactura{
			CodigoDocumentoSector: 1, CodigoEmision: 1, CodigoModalidad: 2, Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1, Cuf: "cuf",
		},
		CodigoMotivo: 0,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "incompleta o inválida")
}

func TestClient_Concurrencia(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarComunicacionResponse><return><transaccion>true</transaccion><mensajesList><codigo>926</codigo><descripcion>OK</descripcion></mensajesList></return></verificarComunicacionResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.VerificarComunicacion(context.Background())
			if err != nil || resp == nil || resp.Transaccion == nil || !*resp.Transaccion {
				t.Errorf("error concurrente: %v", err)
			}
		}()
	}
	wg.Wait()
}
