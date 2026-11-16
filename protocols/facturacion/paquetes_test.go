package facturacion_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/protocols/facturacion"
)

func TestPaqueteYValidacionSeparados(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		body := string(b)
		var op string
		if calls == 1 {
			op = "recepcionPaqueteFactura"
			for _, v := range []string{"<SolicitudServicioRecepcionPaquete>", "<archivo>cGFxdWV0ZQ==</archivo>", "<cantidadFacturas>2</cantidadFacturas>", "<codigoEvento>99</codigoEvento>"} {
				if !strings.Contains(body, v) {
					t.Errorf("falta %s en %s", v, body)
				}
			}
			if strings.Contains(body, "<cafc>") {
				t.Errorf("CAFC inesperado: %s", body)
			}
		} else {
			op = "validacionRecepcionPaqueteFactura"
			if !strings.Contains(body, "<SolicitudServicioValidacionRecepcionPaquete>") || !strings.Contains(body, "<codigoRecepcion>REC</codigoRecepcion>") {
				t.Errorf("validación: %s", body)
			}
		}
		fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><%sResponse><RespuestaServicioFacturacion><codigoEstado>908</codigoEstado><codigoRecepcion>REC</codigoRecepcion><transaccion>true</transaccion></RespuestaServicioFacturacion></%sResponse></soap:Body></soap:Envelope>`, op, op)
	}))
	defer server.Close()
	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	req := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoDocumentoSector: 1, CodigoEmision: 2, CodigoModalidad: 2, Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1, Archivo: []byte("paquete"), FechaEnvio: time.Now(), HashArchivo: "hash"}, CantidadFacturas: 2, CodigoEvento: 99}
	received, err := client.RecepcionPaqueteFactura(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	if calls != 1 || received.Transaccion == nil || !*received.Transaccion || received.CodigoRecepcion != "REC" {
		t.Fatalf("recepción=%+v llamadas=%d", received, calls)
	}
	validated, err := client.ValidacionRecepcionPaqueteFactura(t.Context(), facturacion.SolicitudValidacionPaquete{CodigoDocumentoSector: 1, CodigoEmision: 2, CodigoModalidad: 2, Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1, CodigoRecepcion: received.CodigoRecepcion})
	if err != nil {
		t.Fatal(err)
	}
	if calls != 2 || validated.Transaccion == nil || !*validated.Transaccion {
		t.Fatalf("validación=%+v llamadas=%d", validated, calls)
	}
}

func TestValidacionPaqueteRequiereCodigoRecepcion(t *testing.T) {
	client, err := facturacion.NewClient(facturacion.Config{BaseURL: "http://localhost", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.ValidacionRecepcionPaqueteFactura(t.Context(), facturacion.SolicitudValidacionPaquete{CodigoDocumentoSector: 1, CodigoEmision: 2, CodigoModalidad: 2, Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1})
	if err == nil {
		t.Fatal("aceptó código de recepción vacío")
	}
}
