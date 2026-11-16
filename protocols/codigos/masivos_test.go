package codigos_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/protocols/codigos"
)

func TestCodigosMasivosConResultadosPorPunto(t *testing.T) {
	llamadas := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas++
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		body := string(data)
		if strings.Count(body, "<datosSolicitud>") != 2 || strings.Count(body, "<codigoPuntoVenta>0</codigoPuntoVenta>") != 1 {
			t.Errorf("puntos: %s", body)
		}
		var op, result string
		if llamadas == 1 {
			op = "cuisMasivo"
			if !strings.Contains(body, "<SolicitudCuisMasivoSistemas>") {
				t.Errorf("CUIS: %s", body)
			}
			result = `<RespuestaCuisMasivo><listaRespuestasCuis><codigo>CUIS</codigo><codigoPuntoVenta>0</codigoPuntoVenta><codigoSucursal>0</codigoSucursal><transaccion>true</transaccion></listaRespuestasCuis><listaRespuestasCuis><codigoSucursal>1</codigoSucursal><mensajeServicioList><codigo>991</codigo><descripcion>rechazo</descripcion></mensajeServicioList><transaccion>false</transaccion></listaRespuestasCuis><transaccion>false</transaccion></RespuestaCuisMasivo>`
		} else {
			op = "cufdMasivo"
			if !strings.Contains(body, "<SolicitudCufdMasivo>") || strings.Count(body, "<cuis>CUIS</cuis>") != 2 {
				t.Errorf("CUFD: %s", body)
			}
			result = `<RespuestaCufdMasivo><listaRespuestasCufd><codigo>CUFD</codigo><codigoControl>CTRL</codigoControl><codigoInterno>INTERNO</codigoInterno><codigoPuntoVenta>0</codigoPuntoVenta><codigoSucursal>0</codigoSucursal><cuis>CUIS</cuis><transaccion>true</transaccion></listaRespuestasCufd><transaccion>true</transaccion></RespuestaCufdMasivo>`
		}
		fmt.Fprintf(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><%sResponse>%s</%sResponse></soap:Body></soap:Envelope>`, op, result, op)
	}))
	defer srv.Close()
	c, err := codigos.NewClient(codigos.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	cuis, err := c.CuisMasivo(t.Context(), codigos.SolicitudCuisMasivo{CodigoModalidad: 2, Puntos: []codigos.PuntoCuis{{CodigoPuntoVenta: new(0)}, {CodigoSucursal: 1}}})
	if err != nil {
		t.Fatal(err)
	}
	if cuis.Transaccion == nil || *cuis.Transaccion || len(cuis.Resultados) != 2 || cuis.Resultados[0].CodigoPuntoVenta == nil || *cuis.Resultados[0].CodigoPuntoVenta != 0 || cuis.Resultados[1].CodigoPuntoVenta != nil || cuis.Resultados[1].Transaccion == nil || *cuis.Resultados[1].Transaccion || len(cuis.Resultados[1].Mensajes) != 1 || cuis.Resultados[1].Mensajes[0].Codigo != 991 {
		t.Fatalf("CUIS: %+v", cuis)
	}
	cufd, err := c.CufdMasivo(t.Context(), codigos.SolicitudCufdMasivo{CodigoModalidad: 2, Puntos: []codigos.PuntoCufd{{CodigoPuntoVenta: new(0), Cuis: "CUIS"}, {CodigoSucursal: 1, Cuis: "CUIS"}}})
	if err != nil {
		t.Fatal(err)
	}
	if cufd.Transaccion == nil || !*cufd.Transaccion || len(cufd.Resultados) != 1 || cufd.Resultados[0].CodigoControl != "CTRL" {
		t.Fatalf("CUFD: %+v", cufd)
	}
	if cufd.Resultados[0].CodigoInterno != "INTERNO" {
		t.Fatalf("interno: %+v", cufd)
	}
	if llamadas != 2 {
		t.Fatalf("llamadas=%d", llamadas)
	}
	if _, err := c.CufdMasivo(t.Context(), codigos.SolicitudCufdMasivo{CodigoModalidad: 2, Puntos: []codigos.PuntoCufd{{}}}); err == nil || llamadas != 2 {
		t.Fatal("aceptó CUIS ausente")
	}
}
