package protocols_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	siat "github.com/ron86i/go-siat/v3/protocols"
)

func TestVerificacionComunicacionPorServicio(t *testing.T) {
	paths := []string{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if !strings.Contains(string(body), "<ns:verificarComunicacion>") {
			t.Errorf("solicitud: %s", body)
		}
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><verificarComunicacionResponse><return><mensajesList><codigo>926</codigo><descripcion>comunicación</descripcion></mensajesList><transaccion>true</transaccion></return></verificarComunicacionResponse></soap:Body></soap:Envelope>`)
	}))
	defer srv.Close()
	c, err := siat.NewClient(siat.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	a, err := c.Facturacion.VerificarComunicacion(t.Context())
	if err != nil || a.Transaccion == nil || !*a.Transaccion || len(a.Mensajes) != 1 || a.Mensajes[0].Codigo != 926 {
		t.Fatalf("compra-venta=%+v error=%v", a, err)
	}
	b, err := c.Electronica.VerificarComunicacion(t.Context())
	if err != nil || b.Transaccion == nil || !*b.Transaccion {
		t.Fatalf("electrónica=%+v error=%v", b, err)
	}
	basico, err := c.ServicioBasico.VerificarComunicacion(t.Context())
	if err != nil || basico.Transaccion == nil || !*basico.Transaccion {
		t.Fatalf("servicios básicos=%+v error=%v", basico, err)
	}
	d, err := c.Computarizada.VerificarComunicacion(t.Context())
	if err != nil || d.Transaccion == nil || !*d.Transaccion {
		t.Fatalf("computarizada=%+v error=%v", d, err)
	}
	e, err := c.Operaciones.VerificarComunicacion(t.Context())
	if err != nil || e.Transaccion == nil || !*e.Transaccion {
		t.Fatalf("operaciones=%+v error=%v", e, err)
	}
	f, err := c.Sincronizacion.VerificarComunicacion(t.Context())
	if err != nil || f.Transaccion == nil || !*f.Transaccion {
		t.Fatalf("sincronización=%+v error=%v", f, err)
	}
	g, err := c.Compras.VerificarComunicacion(t.Context())
	if err != nil || g.Transaccion == nil || !*g.Transaccion {
		t.Fatalf("compras=%+v error=%v", g, err)
	}
	want := []string{"/sfd/ServicioFacturacionCompraVenta", "/sfd/ServicioFacturacionElectronica", "/sfd/ServicioFacturacionServicioBasico", "/sfd/ServicioFacturacionComputarizada", "/sfd/FacturacionOperaciones", "/sfd/FacturacionSincronizacion", "/sfd/ServicioRecepcionCompras"}
	if len(paths) != len(want) {
		t.Fatalf("rutas=%v", paths)
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Fatalf("ruta %d: %s", i, paths[i])
		}
	}
}
