package codigos_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/protocols/codigos"
)

func TestVerificarNitYComunicacion(t *testing.T) {
	llamadas := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas++
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		body := string(data)
		if llamadas == 1 {
			for _, text := range []string{"<ns:verificarNit>", "<SolicitudVerificarNit>", "<codigoModalidad>2</codigoModalidad>", "<cuis>CUIS</cuis>", "<nit>123</nit>", "<nitParaVerificacion>456</nitParaVerificacion>"} {
				if !strings.Contains(body, text) {
					t.Errorf("falta %s en %s", text, body)
				}
			}
			if strings.Index(body, "<nit>") > strings.Index(body, "<nitParaVerificacion>") {
				t.Errorf("orden: %s", body)
			}
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarNitResponse><RespuestaVerificarNit><mensajesList><codigo>991</codigo><descripcion>NIT observado</descripcion></mensajesList><transaccion>false</transaccion></RespuestaVerificarNit></verificarNitResponse></soap:Body></soap:Envelope>`)
		} else {
			if !strings.Contains(body, "<ns:verificarComunicacion>") || strings.Contains(body, "SolicitudVerificarNit") {
				t.Errorf("comunicación: %s", body)
			}
			io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarComunicacionResponse><RespuestaComunicacion><transaccion>true</transaccion></RespuestaComunicacion></verificarComunicacionResponse></soap:Body></soap:Envelope>`)
		}
	}))
	defer srv.Close()
	client, err := codigos.NewClient(codigos.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	req := codigos.SolicitudVerificarNit{CodigoModalidad: 2, Cuis: "CUIS", NitParaVerificacion: 456}
	result, err := client.VerificarNit(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.Transaccion == nil || *result.Transaccion || len(result.Mensajes) != 1 || result.Mensajes[0].Codigo != 991 {
		t.Fatalf("NIT: %+v", result)
	}
	comunicacion, err := client.VerificarComunicacion(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if comunicacion.Transaccion == nil || !*comunicacion.Transaccion || llamadas != 2 {
		t.Fatalf("comunicación: %+v llamadas=%d", comunicacion, llamadas)
	}
	req.NitParaVerificacion = 0
	if _, err := client.VerificarNit(t.Context(), req); err == nil || llamadas != 2 {
		t.Fatal("aceptó NIT inválido")
	}
}
