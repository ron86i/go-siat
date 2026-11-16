package protocols_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	siat "github.com/ron86i/go-siat/v3/protocols"
	"github.com/ron86i/go-siat/v3/protocols/sincronizacion"
)

func TestClienteSincronizaProductos(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if r.URL.Path != "/sfd/FacturacionSincronizacion" || !strings.Contains(string(body), "<ns:sincronizarListaProductosServicios>") || !strings.Contains(string(body), "<cuis>CUIS</cuis>") {
			t.Errorf("solicitud %s %s", r.URL.Path, body)
		}
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><sincronizarListaProductosServiciosResponse><RespuestaListaProductos><mensajesList><codigo>17</codigo><descripcion>aviso</descripcion></mensajesList><transaccion>true</transaccion><listaCodigos><codigoActividad>001234</codigoActividad><codigoProducto>987</codigoProducto><descripcionProducto>Servicio</descripcionProducto><nandina>00</nandina></listaCodigos></RespuestaListaProductos></sincronizarListaProductosServiciosResponse></soap:Body></soap:Envelope>`)
	}))
	defer srv.Close()
	c, err := siat.NewClient(siat.Config{BaseURL: srv.URL + "/sfd", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	if c.Codigos == nil || c.Facturacion == nil || c.Sincronizacion == nil || c.Compras == nil || c.Operaciones == nil {
		t.Fatal("servicios no inicializados")
	}
	got, err := c.Sincronizacion.SincronizarListaProductosServicios(t.Context(), sincronizacion.Solicitud{Cuis: "CUIS"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Transaccion == nil || !*got.Transaccion || len(got.Mensajes) != 1 || got.Mensajes[0].Codigo != 17 || len(got.Productos) != 1 || got.Productos[0].CodigoActividad != "001234" || got.Productos[0].CodigoProducto != 987 {
		t.Fatalf("resultado: %+v", got)
	}
}

func TestClienteRechazaConfiguracionInvalida(t *testing.T) {
	if _, err := siat.NewClient(siat.Config{}); err == nil {
		t.Fatal("aceptó configuración vacía")
	}
}
