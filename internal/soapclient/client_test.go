package soapclient_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/internal/soapclient"
)

func TestNoEnviarTokenEnRedireccion(t *testing.T) {
	visits := 0
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { visits++ }))
	defer target.Close()
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusTemporaryRedirect)
	}))
	defer origin.Close()
	client, err := soapclient.New(soapclient.Config{BaseURL: origin.URL, Token: "secreto", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2}, "FacturacionCodigos")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Call(t.Context(), struct{}{}, &struct{}{})
	if err == nil || visits != 0 {
		t.Fatalf("error=%v visitas=%d", err, visits)
	}
}

func TestErroresComunicacionConservanCausa(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusBadGateway) }))
	defer server.Close()
	client, err := soapclient.New(soapclient.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2}, "FacturacionCodigos")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Call(t.Context(), struct{}{}, &struct{}{})
	httpErr, ok := errors.AsType[*soapclient.CommunicationError](err)
	if !ok || httpErr.Etapa != "http" || httpErr.StatusCode != http.StatusBadGateway {
		t.Fatalf("error HTTP: %v", err)
	}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	err = client.Call(ctx, struct{}{}, &struct{}{})
	networkErr, ok := errors.AsType[*soapclient.CommunicationError](err)
	if !ok || networkErr.Etapa != "solicitud" || !errors.Is(err, context.Canceled) {
		t.Fatalf("error de solicitud: %v", err)
	}
}

func TestLimiteRespuestaConfigurable(t *testing.T) {
	contenido := strings.Repeat("x", 4<<20)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><resultado><valor>`)
		io.WriteString(w, contenido)
		io.WriteString(w, `</valor></resultado></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()
	cfg := soapclient.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2}
	client, err := soapclient.New(cfg, "FacturacionCodigos")
	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		Valor string `xml:"valor"`
	}
	if err := client.Call(t.Context(), struct{}{}, &result); err != nil {
		t.Fatal(err)
	}
	if result.Valor != contenido {
		t.Fatalf("longitud=%d", len(result.Valor))
	}
	cfg.MaxRespuestaBytes = 128
	client, err = soapclient.New(cfg, "FacturacionCodigos")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Call(t.Context(), struct{}{}, &result)
	comm, ok := errors.AsType[*soapclient.CommunicationError](err)
	if !ok || comm.Etapa != "lectura" {
		t.Fatalf("límite: %v", err)
	}
	cfg.MaxRespuestaBytes = -1
	if _, err = soapclient.New(cfg, "FacturacionCodigos"); err == nil {
		t.Fatal("aceptó límite negativo")
	}
}

func TestFaultConservaDetalle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><soap:Fault><faultcode>SOAP-1</faultcode><faultstring>error del SIN</faultstring><detail>detalle del SIN</detail></soap:Fault></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()
	client, err := soapclient.New(soapclient.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2}, "FacturacionCodigos")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Call(t.Context(), struct{}{}, &struct{}{})
	fault, ok := errors.AsType[*soapclient.Fault](err)
	if !ok || fault.Code != "SOAP-1" || fault.Message != "error del SIN" || fault.Detail != "detalle del SIN" {
		t.Fatalf("fault: %+v", err)
	}
}

func TestNewFijaLaRutaSIAT3(t *testing.T) {
	for _, tc := range []struct {
		name     string
		baseURL  string
		wantPath string
		wantErr  bool
	}{
		{name: "fija /sfd", baseURL: "https://example.test/sfd", wantPath: "/sfd"},
		{name: "rechaza ruta v2", baseURL: "https://example.test/v2", wantErr: true},
		{name: "rechaza ruta similar", baseURL: "https://example.test/sfd-evil", wantErr: true},
		{name: "rechaza ruta anidada", baseURL: "https://example.test/proxy/sfd", wantErr: true},
		{name: "mock loopback añade ruta", baseURL: "http://127.0.0.1:1234", wantPath: "/sfd"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := soapclient.ResolveBaseURL(tc.baseURL, 2)
			if tc.wantErr {
				if err == nil {
					t.Fatal("se esperaba rechazar la URL")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if want := tc.wantPath; !strings.HasSuffix(got, want) {
				t.Fatalf("BaseURL=%q, se esperaba sufijo %q", got, want)
			}
		})
	}
}

func TestNewRechazaServicioDesconocido(t *testing.T) {
	_, err := soapclient.New(soapclient.Config{BaseURL: "https://example.test/sfd", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2}, "../../otro")
	if err == nil {
		t.Fatal("se aceptó un servicio arbitrario")
	}
}
