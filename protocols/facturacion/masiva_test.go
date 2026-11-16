package facturacion_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/alquileres"
	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
)

func TestEmisionMasivaYValidacion(t *testing.T) {
	llamadas := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas++
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		s := string(b)
		op := "recepcionMasivaFactura"
		if llamadas == 1 {
			if !strings.Contains(s, "<SolicitudServicioRecepcionMasiva>") || !strings.Contains(s, "<cantidadFacturas>2</cantidadFacturas>") || strings.Contains(s, "<codigoEvento>") || strings.Contains(s, "<cafc>") {
				t.Errorf("recepción: %s", s)
			}
		} else {
			op = "validacionRecepcionMasivaFactura"
			if !strings.Contains(s, "<SolicitudServicioValidacionRecepcionMasiva>") || !strings.Contains(s, "<codigoRecepcion>REC</codigoRecepcion>") {
				t.Errorf("validación: %s", s)
			}
		}
		fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><%sResponse><RespuestaServicioFacturacion><codigoRecepcion>REC</codigoRecepcion><transaccion>true</transaccion></RespuestaServicioFacturacion></%sResponse></soap:Body></soap:Envelope>`, op, op)
	}))
	defer srv.Close()
	client, err := facturacion.NewClient(facturacion.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	f1 := newTestFactura(t)
	f2 := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		c.NumeroFactura = 101
	})
	req, err := (facturacion.SolicitudRecepcionMasiva{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2, CodigoEmision: 3, Cufd: "cufd", Cuis: "cuis", FechaEnvio: time.Now()}}).WithFacturas(t.Context(), facturacion.AsDocumentos([]*compraventa.Factura{f1, f2}), nil)
	if err != nil {
		t.Fatal(err)
	}
	res, err := client.RecepcionMasivaFactura(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	if llamadas != 1 || res.Transaccion == nil || !*res.Transaccion || res.CodigoRecepcion != "REC" {
		t.Fatalf("recepción=%+v llamadas=%d", res, llamadas)
	}
	validacion := facturacion.SolicitudValidacionMasiva{CodigoDocumentoSector: 1, CodigoEmision: 3, CodigoModalidad: 2, Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1, CodigoRecepcion: res.CodigoRecepcion}
	_, err = client.ValidacionRecepcionMasivaFactura(t.Context(), validacion)
	if err != nil {
		t.Fatal(err)
	}
	if llamadas != 2 {
		t.Fatalf("llamadas=%d", llamadas)
	}
}

func TestEmisionMasivaConAlquilerUsaLaInterfazDeDocumentos(t *testing.T) {
	documento := newTestAlquilerFactura(t, 0, "cufd")
	req := facturacion.SolicitudRecepcionMasiva{
		SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{
			CodigoModalidad: 2,
			Cufd:            "cufd",
			Cuis:            "cuis",
		},
	}

	preparada, err := req.WithFacturas(t.Context(), facturacion.AsDocumentos([]*alquileres.Factura{documento}), nil)
	if err != nil {
		t.Fatal(err)
	}
	if preparada.CodigoDocumentoSector != 2 || preparada.CantidadFacturas != 1 {
		t.Fatalf("solicitud masiva: sector=%d cantidad=%d", preparada.CodigoDocumentoSector, preparada.CantidadFacturas)
	}
}

func TestValidacionPaqueteRechazaDestinoIncorrecto(t *testing.T) {
	client, err := facturacion.NewClient(facturacion.Config{BaseURL: "http://localhost", Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	req := facturacion.SolicitudValidacionPaquete{CodigoDocumentoSector: 2, CodigoEmision: 2, CodigoModalidad: 2, Cufd: "cufd", Cuis: "cuis", TipoFacturaDocumento: 1, CodigoRecepcion: "REC"}
	if _, err = client.ValidacionRecepcionPaqueteFactura(t.Context(), req); err == nil || !strings.Contains(err.Error(), "no ServicioFacturacionCompraVenta") {
		t.Fatalf("destino incorrecto: %v", err)
	}
}
