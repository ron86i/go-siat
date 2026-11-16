package codigos_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/protocols/codigos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_SolicitarCUIS_Y_CUFD(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.URL.Path != "/sfd/FacturacionCodigos" || r.Header.Get("apiKey") != "TokenApi token123" {
			t.Errorf("ruta o token incorrectos")
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		body := string(data)
		if !strings.Contains(body, "<codigoAmbiente>2</codigoAmbiente>") ||
			!strings.Contains(body, "<codigoSistema>SISTEMA</codigoSistema>") ||
			!strings.Contains(body, "<nit>123456</nit>") {
			t.Errorf("identidad incorrecta: %s", body)
		}
		if calls == 1 {
			// Solicitud CUIS
			if !strings.Contains(body, "<ns:cuis>") ||
				!strings.Contains(body, "<SolicitudCuis>") ||
				!strings.Contains(body, "<codigoModalidad>2</codigoModalidad>") ||
				!strings.Contains(body, "<codigoSucursal>0</codigoSucursal>") {
				t.Errorf("cuerpo CUIS incorrecto: %s", body)
			}
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><cuisResponse><RespuestaCuis><codigo>CUIS-PRUEBA-123</codigo><fechaVigencia>2026-10-31T12:00:00.000</fechaVigencia><transaccion>true</transaccion></RespuestaCuis></cuisResponse></soap:Body></soap:Envelope>`)
		} else {
			// Solicitud CUFD
			if !strings.Contains(body, "<ns:cufd>") ||
				!strings.Contains(body, "<SolicitudCufd>") ||
				!strings.Contains(body, "<codigoPuntoVenta>0</codigoPuntoVenta>") ||
				!strings.Contains(body, "<cuis>CUIS-PRUEBA-123</cuis>") {
				t.Errorf("cuerpo CUFD incorrecto: %s", body)
			}
			fmt.Fprint(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><cufdResponse><RespuestaCufd><codigo>CUFD-PRUEBA-XYZ</codigo><codigoControl>CTRL-99</codigoControl><codigoInterno>INTERNO-V3</codigoInterno><direccion>Av. Principal</direccion><fechaVigencia>2026-11-01T12:00:00-04:00</fechaVigencia><mensajesList><codigo>17</codigo><descripcion>aviso</descripcion></mensajesList><transaccion>true</transaccion></RespuestaCufd></cufdResponse></soap:Body></soap:Envelope>`)
		}
	}))
	defer srv.Close()

	client, err := codigos.NewClient(codigos.Config{
		BaseURL:        srv.URL + "/sfd",
		Token:          "token123",
		Nit:            123456,
		CodigoSistema:  "SISTEMA",
		CodigoAmbiente: 2,
	})
	require.NoError(t, err)

	// 1. Solicitar CUIS
	puntoVenta := 0
	cuis, err := client.SolicitarCUIS(t.Context(), codigos.SolicitudCUIS{
		CodigoModalidad:  2,
		CodigoSucursal:   0,
		CodigoPuntoVenta: &puntoVenta,
	})
	require.NoError(t, err)
	require.NotNil(t, cuis)
	assert.Equal(t, "CUIS-PRUEBA-123", cuis.Codigo)
	assert.Equal(t, "2026-10-31T12:00:00.000", cuis.FechaVigencia)
	require.NotNil(t, cuis.Transaccion)
	assert.True(t, *cuis.Transaccion)

	// 2. Solicitar CUFD con código interno de SIAT v3
	cufd, err := client.SolicitarCUFD(t.Context(), codigos.SolicitudCUFD{
		CodigoModalidad:  2,
		CodigoSucursal:   0,
		CodigoPuntoVenta: &puntoVenta,
		Cuis:             cuis.Codigo,
	})
	require.NoError(t, err)
	require.NotNil(t, cufd)
	assert.Equal(t, "CUFD-PRUEBA-XYZ", cufd.Codigo)
	assert.Equal(t, "CTRL-99", cufd.CodigoControl)
	assert.Equal(t, "INTERNO-V3", cufd.CodigoInterno, "SIAT v3 debe procesar codigoInterno")
	assert.Equal(t, "Av. Principal", cufd.Direccion)
	require.NotNil(t, cufd.Transaccion)
	assert.True(t, *cufd.Transaccion)
	require.Len(t, cufd.Mensajes, 1)
	assert.Equal(t, 17, cufd.Mensajes[0].Codigo)

	assert.Equal(t, 2, calls)
}

func TestClient_ModalidadPorDefectoEnConfig(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		if !strings.Contains(string(data), "<codigoModalidad>1</codigoModalidad>") {
			t.Errorf("debe usar modalidad 1 de la configuración del cliente: %s", string(data))
		}
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><cuisResponse><RespuestaCuis><codigo>CUIS-DEFAULT</codigo><transaccion>true</transaccion></RespuestaCuis></cuisResponse></soap:Body></soap:Envelope>`)
	}))
	defer srv.Close()

	// Configuración con modalidad por defecto = 1 (Electrónica)
	client, err := codigos.NewClient(codigos.Config{
		BaseURL:        srv.URL,
		Token:          "token",
		Nit:            123,
		CodigoSistema:  "SISTEMA",
		CodigoAmbiente: 2,
	})
	require.NoError(t, err)
	client = client.WithModalidad(1)

	// Solicitud sin especificar CodigoModalidad: debe tomar la del cliente
	res, err := client.SolicitarCUIS(t.Context(), codigos.SolicitudCUIS{
		CodigoSucursal: 0,
	})
	require.NoError(t, err)
	assert.Equal(t, "CUIS-DEFAULT", res.Codigo)
}

func TestClient_ValidacionesParametros(t *testing.T) {
	client, err := codigos.NewClient(codigos.Config{
		BaseURL:        "http://localhost",
		Token:          "token",
		Nit:            123,
		CodigoSistema:  "SISTEMA",
		CodigoAmbiente: 2,
	})
	require.NoError(t, err)

	ctx := t.Context()

	// 1. Modalidad ausente o inválida
	_, err = client.SolicitarCUIS(ctx, codigos.SolicitudCUIS{CodigoModalidad: 0})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "CodigoModalidad")

	// 2. Sucursal negativa
	_, err = client.SolicitarCUIS(ctx, codigos.SolicitudCUIS{CodigoModalidad: 1, CodigoSucursal: -1})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sucursal o punto de venta")

	// 3. Punto de venta negativo
	pvNegativo := -5
	_, err = client.SolicitarCUIS(ctx, codigos.SolicitudCUIS{CodigoModalidad: 1, CodigoPuntoVenta: &pvNegativo})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sucursal o punto de venta")

	// 4. CUFD sin CUIS
	_, err = client.SolicitarCUFD(ctx, codigos.SolicitudCUFD{CodigoModalidad: 1, Cuis: ""})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "CUIS requerido")
}

func TestClient_RespuestaRechazada(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><cuisResponse><RespuestaCuis><mensajesList><codigo>991</codigo><descripcion>rechazo</descripcion></mensajesList><transaccion>false</transaccion></RespuestaCuis></cuisResponse></soap:Body></soap:Envelope>`)
	}))
	defer srv.Close()

	client, err := codigos.NewClient(codigos.Config{
		BaseURL:        srv.URL,
		Token:          "token",
		Nit:            123,
		CodigoSistema:  "SISTEMA",
		CodigoAmbiente: 2,
	})
	require.NoError(t, err)

	res, err := client.SolicitarCUIS(t.Context(), codigos.SolicitudCUIS{CodigoModalidad: 1})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.Transaccion)
	assert.False(t, *res.Transaccion)
	require.Len(t, res.Mensajes, 1)
	assert.Equal(t, 991, res.Mensajes[0].Codigo)
}
