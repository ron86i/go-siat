package services_test

import (
	"encoding/xml"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/codigos"
	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/common"
)

func TestCodigosOptionsFromEnvironment(t *testing.T) {
	base := map[string]string{
		"SIAT_TOKEN": "test", "SIAT_NIT": "123456789",
		"SIAT_CODIGO_SISTEMA": "SISTEMA", "SIAT_CODIGO_AMBIENTE": "2",
		"SIAT_CODIGO_MODALIDAD": "1", "SIAT_URL": "https://pilot.invalid/sfd",
	}
	for _, tc := range []struct {
		name, key, value, wantError string
	}{
		{name: "defaults"},
		{name: "point", key: "SIAT_PV", value: "0"},
		{name: "invalid point", key: "SIAT_PV", value: "...", wantError: "SIAT_PV"},
		{name: "negative point", key: "SIAT_PV", value: "-1", wantError: "SIAT_PV"},
		{name: "invalid modality", key: "SIAT_CODIGO_MODALIDAD", value: "invalid", wantError: "SIAT_CODIGO_MODALIDAD"},
		{name: "unsupported modality", key: "SIAT_CODIGO_MODALIDAD", value: "0", wantError: "SIAT_CODIGO_MODALIDAD"},
		{name: "old endpoint", key: "SIAT_URL", value: "https://pilot.invalid/v2", wantError: "SIAT_URL"},
		{name: "invalid URL", key: "SIAT_URL", value: "/sfd", wantError: "SIAT_URL"},
		{name: "missing token", key: "SIAT_TOKEN", value: "", wantError: "SIAT_TOKEN"},
		{name: "invalid NIT", key: "SIAT_NIT", value: "...", wantError: "SIAT_NIT"},
		{name: "production", key: "SIAT_CODIGO_AMBIENTE", value: "1", wantError: "SIAT_CODIGO_AMBIENTE"},
		{name: "zero repetitions", key: "SIAT_VECES", value: "0", wantError: "SIAT_VECES"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			env := maps.Clone(base)
			if tc.key != "" {
				env[tc.key] = tc.value
			}
			options, err := readCodigosTestOptions(func(key string) string { return env[key] })
			if tc.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantError) {
					t.Fatalf("error = %v; want %s", err, tc.wantError)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			wantPoint := 1
			if tc.key == "SIAT_PV" {
				wantPoint = 0
			}
			if options.puntoVenta != wantPoint || options.sucursal != 0 || options.veces != 10 {
				t.Fatalf("unexpected options: point=%d branch=%d repetitions=%d", options.puntoVenta, options.sucursal, options.veces)
			}
		})
	}

	t.Run("long environment keys", func(t *testing.T) {
		env := maps.Clone(base)
		env["SIAT_CODIGO_PUNTO_VENTA"] = "3"
		env["SIAT_CODIGO_SUCURSAL"] = "2"
		options, err := readCodigosTestOptions(func(key string) string { return env[key] })
		if err != nil || options.puntoVenta != 3 || options.sucursal != 2 {
			t.Fatalf("aliases not applied: error=%v", err)
		}
	})
}

func TestCodigosCuisResponseValidation(t *testing.T) {
	for _, tc := range []struct {
		name      string
		response  codigos.RespuestaCuis
		wantError bool
	}{
		{"accepted", codigos.RespuestaCuis{Codigo: "CUIS", Transaccion: true}, false},
		{"existing", codigos.RespuestaCuis{Codigo: "CUIS", MensajesList: []common.MensajeServicio{{Codigo: 980}}}, false},
		{"missing code", codigos.RespuestaCuis{Transaccion: true}, true},
		{"existing without code", codigos.RespuestaCuis{MensajesList: []common.MensajeServicio{{Codigo: 980}}}, true},
		{"966", codigos.RespuestaCuis{MensajesList: []common.MensajeServicio{{Codigo: 966, Descripcion: "NO SE PUEDE RECUPERAR LOS DATOS DEL CONTRIBUYENTE"}}}, true},
		{"mixed error", codigos.RespuestaCuis{Codigo: "CUIS", MensajesList: []common.MensajeServicio{{Codigo: 980}, {Codigo: 966}}}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := validateCodigosCuisResponse(tc.response)
			if (err != nil) != tc.wantError {
				t.Fatalf("error = %v; want error=%t", err, tc.wantError)
			}
		})
	}
}

func TestCodigosCufdIterationsReuseCuis(t *testing.T) {
	for _, configured := range []bool{false, true} {
		t.Run(fmt.Sprintf("configuredCUIS_%t", configured), func(t *testing.T) {
			var cuisCalls, cufdCalls atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Error(err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if r.URL.Path != "/sfd/FacturacionCodigos" {
					t.Errorf("unexpected route: %s", r.URL.Path)
				}
				w.Header().Set("Content-Type", "text/xml")
				operation, response := "cufd", "<codigo>CUFD</codigo><codigoControl>CONTROL</codigoControl><fechaVigencia>2026-12-01T23:59:59-04:00</fechaVigencia>"
				if strings.Contains(string(body), "<ns:cuis>") {
					cuisCalls.Add(1)
					if !strings.Contains(string(body), "<codigoPuntoVenta>3</codigoPuntoVenta>") || !strings.Contains(string(body), "<codigoSucursal>2</codigoSucursal>") {
						t.Error("CUIS request uses a different point of sale or branch")
					}
					operation, response = "cuis", "<codigo>CUIS-PRUEBA</codigo>"
				} else {
					cufdCalls.Add(1)
					var payload struct {
						Point  int    `xml:"Body>cufd>SolicitudCufd>codigoPuntoVenta"`
						Branch int    `xml:"Body>cufd>SolicitudCufd>codigoSucursal"`
						Cuis   string `xml:"Body>cufd>SolicitudCufd>cuis"`
					}
					if err := xml.Unmarshal(body, &payload); err != nil {
						t.Error(err)
					}
					if payload.Point != 3 || payload.Branch != 2 || payload.Cuis != "CUIS-PRUEBA" {
						t.Errorf("CUFD request uses inconsistent identity: %+v", payload)
					}
				}
				wrapper := "RespuestaCufd"
				if operation == "cuis" {
					wrapper = "RespuestaCuis"
				}
				fmt.Fprintf(w, `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body><%sResponse xmlns="https://siat.impuestos.gob.bo/"><%s>%s<transaccion>true</transaccion></%s></%sResponse></Body></Envelope>`, operation, wrapper, response, wrapper, operation)
			}))
			t.Cleanup(server.Close)
			client, err := siat.New(siat.Config{
				Token: "test", Nit: 123456789, CodigoSistema: "SISTEMA", CodigoAmbiente: siat.AmbientePruebas,
				BaseURL: server.URL + "/sfd", HTTPClient: server.Client(),
			})
			if err != nil {
				t.Fatal(err)
			}
			options := codigosTestOptions{modalidad: 1, sucursal: 2, puntoVenta: 3, veces: 3}
			wantCuis := int32(1)
			if configured {
				options.cuis, wantCuis = "CUIS-PRUEBA", 0
			}
			runCodigosCufdIterations(t, client.Codigos(), options)
			if cuisCalls.Load() != wantCuis || cufdCalls.Load() != 3 {
				t.Errorf("calls: CUIS=%d CUFD=%d; want %d and 3", cuisCalls.Load(), cufdCalls.Load(), wantCuis)
			}
		})
	}
}

func TestCodigosIntegrationRequiresOptIn(t *testing.T) {
	t.Setenv("SIAT_RUN", "0")
	reached := false
	t.Run("disabled", func(t *testing.T) {
		newCodigosIntegration(t)
		reached = true
	})
	if reached {
		t.Fatal("real integration setup ran without SIAT_RUN=1")
	}
}
