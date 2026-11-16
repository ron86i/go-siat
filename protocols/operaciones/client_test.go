package operaciones_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/protocols/operaciones"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistroYConsultaDeEvento(t *testing.T) {
	llamadas := 0
	servidor := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas++
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		texto := string(body)
		if llamadas == 1 {
			for _, want := range []string{"<ns:registroEventoSignificativo>", "<SolicitudEventoSignificativo>", "<cufd>VIGENTE</cufd>", "<cufdEvento>EVENTO</cufdEvento>", "<fechaHoraFinEvento>2026-10-31T12:05:00.123</fechaHoraFinEvento>", "<fechaHoraInicioEvento>2026-10-31T12:00:00.123</fechaHoraInicioEvento>"} {
				if !strings.Contains(texto, want) {
					t.Errorf("falta %s en %s", want, texto)
				}
			}
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><registroEventoSignificativoResponse><RespuestaListaEventos><codigoRecepcionEventoSignificativo>99</codigoRecepcionEventoSignificativo><transaccion>true</transaccion></RespuestaListaEventos></registroEventoSignificativoResponse></soap:Body></soap:Envelope>`)
		} else {
			if !strings.Contains(texto, "<ns:consultaEventoSignificativo>") || !strings.Contains(texto, "<cufd>VIGENTE</cufd>") || !strings.Contains(texto, "<fechaEvento>2026-10-31T12:00:00.123</fechaEvento>") {
				t.Errorf("consulta: %s", texto)
			}
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><consultaEventoSignificativoResponse><RespuestaListaEventos><listaCodigos><codigoEvento>5</codigoEvento><codigoRecepcionEventoSignificativo>99</codigoRecepcionEventoSignificativo><fechaInicio>2026-10-31T12:00:00</fechaInicio></listaCodigos><mensajesList><codigo>17</codigo><descripcion>aviso</descripcion></mensajesList><transaccion>true</transaccion></RespuestaListaEventos></consultaEventoSignificativoResponse></soap:Body></soap:Envelope>`)
		}
	}))
	defer servidor.Close()

	client, err := operaciones.NewClient(operaciones.Config{BaseURL: servidor.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	inicio := time.Date(2026, 10, 31, 12, 0, 0, 123000000, time.FixedZone("Bolivia", -4*60*60))
	solicitud := operaciones.SolicitudRegistroEvento{
		CodigoMotivoEvento:    5,
		Cufd:                  "VIGENTE",
		CufdEvento:            "EVENTO",
		Cuis:                  "CUIS",
		Descripcion:           "Corte",
		FechaHoraInicioEvento: inicio,
		FechaHoraFinEvento:    inicio.Add(5 * time.Minute),
	}
	respuesta, err := client.RegistroEventoSignificativo(t.Context(), solicitud)
	require.NoError(t, err)
	require.NotNil(t, respuesta)
	require.NotNil(t, respuesta.Transaccion)
	assert.True(t, *respuesta.Transaccion)
	assert.Equal(t, int64(99), respuesta.CodigoRecepcionEventoSignificativo)

	resultado, err := client.ConsultaEventoSignificativo(t.Context(), operaciones.SolicitudConsultaEvento{Cufd: "VIGENTE", Cuis: "CUIS", FechaEvento: inicio})
	require.NoError(t, err)
	require.NotNil(t, resultado)
	require.NotNil(t, resultado.Transaccion)
	assert.True(t, *resultado.Transaccion)
	require.Len(t, resultado.Eventos, 1)
	assert.Equal(t, int64(99), resultado.Eventos[0].CodigoRecepcionEventoSignificativo)
	require.Len(t, resultado.Mensajes, 1)

	// Intervalo inválido (fin <= inicio)
	solicitud.FechaHoraFinEvento = inicio
	_, err = client.RegistroEventoSignificativo(t.Context(), solicitud)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "posterior al inicio")

	assert.Equal(t, 2, llamadas)
}

func TestOperacionesPuntoVenta(t *testing.T) {
	count := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		body := string(data)
		if !strings.Contains(body, "<codigoAmbiente>2</codigoAmbiente>") || !strings.Contains(body, "<nit>123</nit>") || !strings.Contains(body, "<cuis>CUIS</cuis>") {
			t.Errorf("solicitud: %s", body)
		}
		var op, element, content string
		switch count {
		case 1:
			op = "registroPuntoVenta"
			element = "RespuestaRegistroPuntoVenta"
			content = "<codigoPuntoVenta>7</codigoPuntoVenta>"
			for _, s := range []string{"<SolicitudRegistroPuntoVenta>", "<codigoModalidad>2</codigoModalidad>", "<codigoTipoPuntoVenta>1</codigoTipoPuntoVenta>", "<nombrePuntoVenta>CAJA</nombrePuntoVenta>"} {
				if !strings.Contains(body, s) {
					t.Errorf("falta %s en %s", s, body)
				}
			}
		case 2:
			op = "consultaPuntoVenta"
			element = "RespuestaConsultaPuntoVenta"
			content = "<listaPuntosVentas><codigoPuntoVenta>7</codigoPuntoVenta><nombrePuntoVenta>CAJA</nombrePuntoVenta><tipoPuntoVenta>FIJO</tipoPuntoVenta></listaPuntosVentas>"
			if !strings.Contains(body, "<SolicitudConsultaPuntoVenta>") {
				t.Errorf("consulta: %s", body)
			}
		default:
			op = "cierrePuntoVenta"
			element = "RespuestaCierrePuntoVenta"
			content = "<codigoPuntoVenta>7</codigoPuntoVenta>"
			if !strings.Contains(body, "<SolicitudCierrePuntoVenta>") || !strings.Contains(body, "<codigoPuntoVenta>7</codigoPuntoVenta>") {
				t.Errorf("cierre: %s", body)
			}
		}
		fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><%sResponse><%s>%s<mensajesList><codigo>123</codigo><descripcion>mensaje</descripcion></mensajesList><transaccion>true</transaccion></%s></%sResponse></soap:Body></soap:Envelope>`, op, element, content, element, op)
	}))
	defer srv.Close()

	client, err := operaciones.NewClient(operaciones.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	require.NoError(t, err)

	reg, err := client.RegistroPuntoVenta(t.Context(), operaciones.SolicitudRegistroPuntoVenta{CodigoModalidad: 2, CodigoTipoPuntoVenta: 1, Cuis: "CUIS", Descripcion: "Caja", NombrePuntoVenta: "CAJA"})
	require.NoError(t, err)
	require.NotNil(t, reg)
	require.NotNil(t, reg.Transaccion)
	assert.True(t, *reg.Transaccion)
	require.NotNil(t, reg.CodigoPuntoVenta)
	assert.Equal(t, 7, *reg.CodigoPuntoVenta)
	require.Len(t, reg.Mensajes, 1)

	lista, err := client.ConsultaPuntoVenta(t.Context(), operaciones.SolicitudConsultaPuntoVenta{Cuis: "CUIS"})
	require.NoError(t, err)
	require.NotNil(t, lista)
	require.NotNil(t, lista.Transaccion)
	assert.True(t, *lista.Transaccion)
	require.Len(t, lista.PuntosVentas, 1)
	assert.Equal(t, 7, lista.PuntosVentas[0].CodigoPuntoVenta)

	cierre, err := client.CierrePuntoVenta(t.Context(), operaciones.SolicitudCierrePuntoVenta{CodigoPuntoVenta: 7, Cuis: "CUIS"})
	require.NoError(t, err)
	require.NotNil(t, cierre)
	require.NotNil(t, cierre.Transaccion)
	assert.True(t, *cierre.Transaccion)
	require.NotNil(t, cierre.CodigoPuntoVenta)
	assert.Equal(t, 7, *cierre.CodigoPuntoVenta)
	assert.Equal(t, 3, count)

	_, err = client.CierrePuntoVenta(t.Context(), operaciones.SolicitudCierrePuntoVenta{CodigoPuntoVenta: -1, Cuis: "CUIS"})
	require.Error(t, err)
	assert.Equal(t, 3, count)
}

func TestComisionistaYCierreSistema(t *testing.T) {
	llamadas := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas++
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		s := string(data)
		if llamadas == 1 {
			for _, v := range []string{"<SolicitudPuntoVentaComisionista>", "<fechaFin>2026-11-01T12:00:00.000</fechaFin>", "<fechaInicio>2026-10-31T12:00:00.000</fechaInicio>", "<nitComisionista>456</nitComisionista>", "<numeroContrato>ABC</numeroContrato>"} {
				if !strings.Contains(s, v) {
					t.Errorf("falta %s en %s", v, s)
				}
			}
			if strings.Index(s, "<fechaFin>") > strings.Index(s, "<fechaInicio>") {
				t.Errorf("orden esperado en WSDL: fechaFin antes de fechaInicio: %s", s)
			}
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><registroPuntoVentaComisionistaResponse><RespuestaPuntoVentaComisionista><codigoPuntoVenta>4</codigoPuntoVenta><transaccion>true</transaccion></RespuestaPuntoVentaComisionista></registroPuntoVentaComisionistaResponse></soap:Body></soap:Envelope>`)
		} else {
			for _, v := range []string{"<SolicitudOperaciones>", "<codigoPuntoVenta>0</codigoPuntoVenta>", "<codigoSistema>SISTEMA</codigoSistema>"} {
				if !strings.Contains(s, v) {
					t.Errorf("falta %s en %s", v, s)
				}
			}
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><cierreOperacionesSistemaResponse><RespuestaCierreSistemas><codigoSistema>SISTEMA</codigoSistema><mensajesList><codigo>991</codigo><descripcion>rechazo</descripcion></mensajesList><transaccion>false</transaccion></RespuestaCierreSistemas></cierreOperacionesSistemaResponse></soap:Body></soap:Envelope>`)
		}
	}))
	defer srv.Close()

	c, err := operaciones.NewClient(operaciones.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "SISTEMA", CodigoAmbiente: 2})
	require.NoError(t, err)

	inicio := time.Date(2026, 10, 31, 12, 0, 0, 0, time.FixedZone("Bolivia", -4*60*60))
	req := operaciones.SolicitudRegistroComisionista{
		CodigoModalidad:  2,
		Cuis:             "CUIS",
		Descripcion:      "sucursal",
		FechaInicio:      inicio,
		FechaFin:         inicio.Add(24 * time.Hour),
		NitComisionista:  456,
		NombrePuntoVenta: "CAJA",
		NumeroContrato:   "ABC",
	}
	result, err := c.RegistroPuntoVentaComisionista(t.Context(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, result.Transaccion)
	assert.True(t, *result.Transaccion)
	require.NotNil(t, result.CodigoPuntoVenta)
	assert.Equal(t, 4, *result.CodigoPuntoVenta)

	puntoCero := 0
	closeResult, err := c.CierreOperacionesSistema(t.Context(), operaciones.SolicitudCierreSistema{
		CodigoModalidad:  2,
		CodigoPuntoVenta: &puntoCero,
		Cuis:             "CUIS",
	})
	require.NoError(t, err)
	require.NotNil(t, closeResult)
	require.NotNil(t, closeResult.Transaccion)
	assert.False(t, *closeResult.Transaccion)
	require.Len(t, closeResult.Mensajes, 1)
	assert.Equal(t, 991, closeResult.Mensajes[0].Codigo)
	assert.Equal(t, 2, llamadas)

	// Validación fecha fin <= fecha inicio
	req.FechaFin = inicio
	_, err = c.RegistroPuntoVentaComisionista(t.Context(), req)
	require.Error(t, err)
	assert.Equal(t, 2, llamadas)
}

func TestVerificarComunicacion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarComunicacionResponse><return><transaccion>true</transaccion><mensajesList><codigo>926</codigo><descripcion>COMUNICACION EXITOSA</descripcion></mensajesList></return></verificarComunicacionResponse></soap:Body></soap:Envelope>`)
	}))
	defer srv.Close()

	c, err := operaciones.NewClient(operaciones.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "SISTEMA", CodigoAmbiente: 2})
	require.NoError(t, err)

	com, err := c.VerificarComunicacion(t.Context())
	require.NoError(t, err)
	require.NotNil(t, com)
	require.NotNil(t, com.Transaccion)
	assert.True(t, *com.Transaccion)
	require.Len(t, com.Mensajes, 1)
	assert.Equal(t, 926, com.Mensajes[0].Codigo)
}
