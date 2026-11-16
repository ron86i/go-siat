package compras_test

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/protocols/compras"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlujosCompras(t *testing.T) {
	llamadas := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas++
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		body := string(data)
		if r.URL.Path != "/sfd/ServicioRecepcionCompras" || r.Header.Get("apiKey") != "TokenApi token" {
			t.Errorf("ruta %s", r.URL.Path)
		}
		for _, field := range []string{"<codigoAmbiente>2</codigoAmbiente>", "<codigoPuntoVenta>0</codigoPuntoVenta>", "<codigoSucursal>0</codigoSucursal>", "<cufd>CUFD</cufd>", "<cuis>CUIS</cuis>", "<nit>123</nit>"} {
			if !strings.Contains(body, field) {
				t.Errorf("falta %s en %s", field, body)
			}
		}
		var op, result string
		switch llamadas {
		case 1:
			op = "recepcionPaqueteCompras"
			for _, field := range []string{"<SolicitudRecepcionCompras>", "<archivo>" + base64.StdEncoding.EncodeToString([]byte("paquete")) + "</archivo>", "<cantidadFacturas>1</cantidadFacturas>", "<fechaEnvio>2026-10-31T12:00:00.000</fechaEnvio>", "<gestion>2026</gestion>", "<periodo>10</periodo>"} {
				if !strings.Contains(body, field) {
					t.Errorf("falta %s en %s", field, body)
				}
			}
			result = `<codigoRecepcion>REC</codigoRecepcion><transaccion>true</transaccion>`
		case 2:
			op = "validacionRecepcionPaqueteCompras"
			if !strings.Contains(body, "<SolicitudValidacionRecepcionCompras>") || !strings.Contains(body, "<codigoRecepcion>REC</codigoRecepcion>") {
				t.Errorf("validación: %s", body)
			}
			result = `<codigoEstado>901</codigoEstado><mensajesList><codigo>4333</codigo><descripcion>rechazo</descripcion></mensajesList><transaccion>false</transaccion>`
		case 3:
			op = "consultaCompras"
			if !strings.Contains(body, "<SolicitudConsultaCompras>") || !strings.Contains(body, "<fecha>2026-10-31T12:00:00.000</fecha>") {
				t.Errorf("consulta: %s", body)
			}
			result = `<archivo>Y29tcHJhcw==</archivo><transaccion>true</transaccion>`
		case 4:
			op = "confirmacionCompras"
			if !strings.Contains(body, "<SolicitudConfirmacionCompras>") {
				t.Errorf("confirmación: %s", body)
			}
			result = `<codigoRecepcion>CONF</codigoRecepcion><transaccion>true</transaccion>`
		default:
			op = "anulacionCompra"
			if !strings.Contains(body, "<SolicitudAnulacionCompra>") || !strings.Contains(body, "<nitProveedor>456</nitProveedor>") || !strings.Contains(body, "<nroFactura>9</nroFactura>") {
				t.Errorf("anulación: %s", body)
			}
			result = `<transaccion>true</transaccion>`
		}
		fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><%sResponse><RespuestaServicioFacturacion>%s</RespuestaServicioFacturacion></%sResponse></soap:Body></soap:Envelope>`, op, result, op)
	}))
	defer server.Close()

	c, err := compras.NewClient(compras.Config{BaseURL: server.URL + "/sfd", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	base := compras.DatosSucursal{Cufd: "CUFD", Cuis: "CUIS"}
	fecha := time.Date(2026, 10, 31, 12, 0, 0, 0, time.UTC)
	req := compras.SolicitudPaquete{DatosSucursal: base, Archivo: []byte("paquete"), CantidadFacturas: 1, FechaEnvio: fecha, Gestion: 2026, HashArchivo: "hash", Periodo: 10}

	recepcion, err := c.RecepcionPaqueteCompras(t.Context(), req)
	require.NoError(t, err)
	require.NotNil(t, recepcion)
	require.NotNil(t, recepcion.Transaccion)
	assert.True(t, *recepcion.Transaccion)
	assert.Equal(t, "REC", recepcion.CodigoRecepcion)

	validacion, err := c.ValidacionRecepcionPaqueteCompras(t.Context(), compras.SolicitudValidacion{DatosSucursal: base, CodigoRecepcion: recepcion.CodigoRecepcion})
	require.NoError(t, err)
	require.NotNil(t, validacion)
	require.NotNil(t, validacion.Transaccion)
	assert.False(t, *validacion.Transaccion)
	require.Len(t, validacion.Mensajes, 1)
	assert.Equal(t, 4333, validacion.Mensajes[0].Codigo)

	consulta, err := c.ConsultaCompras(t.Context(), compras.SolicitudConsulta{DatosSucursal: base, Fecha: fecha})
	require.NoError(t, err)
	require.NotNil(t, consulta)
	assert.Equal(t, "compras", string(consulta.Archivo))

	confirmacion, err := c.ConfirmacionCompras(t.Context(), req)
	require.NoError(t, err)
	require.NotNil(t, confirmacion)
	assert.Equal(t, "CONF", confirmacion.CodigoRecepcion)

	anulacion, err := c.AnulacionCompra(t.Context(), compras.SolicitudAnulacion{DatosSucursal: base, NitProveedor: 456, NroFactura: 9})
	require.NoError(t, err)
	require.NotNil(t, anulacion)
	require.NotNil(t, anulacion.Transaccion)
	assert.True(t, *anulacion.Transaccion)
	assert.Equal(t, 5, llamadas)
}

func TestCompras_VerificarComunicacion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarComunicacionResponse><return><transaccion>true</transaccion><mensajesList><codigo>926</codigo><descripcion>COMUNICACION EXITOSA</descripcion></mensajesList></return></verificarComunicacionResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()

	c, err := compras.NewClient(compras.Config{BaseURL: server.URL + "/sfd", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	resp, err := c.VerificarComunicacion(t.Context())
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Transaccion)
	assert.True(t, *resp.Transaccion)
	require.Len(t, resp.Mensajes, 1)
	assert.Equal(t, 926, resp.Mensajes[0].Codigo)
}
