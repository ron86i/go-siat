package siat_test

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models"
)

type siatCapture struct {
	path, body, token string
}

func siatTestServer(t *testing.T) (*httptest.Server, <-chan siatCapture) {
	t.Helper()
	captures := make(chan siatCapture, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		captures <- siatCapture{r.URL.Path, string(body), r.Header.Get("apiKey")}
		w.Header().Set("Content-Type", "text/xml")
		if strings.Contains(string(body), "<ns:verificarComunicacion>") {
			io.WriteString(w, `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body><verificarComunicacionResponse xmlns="https://siat.impuestos.gob.bo/"><return><transaccion>true</transaccion></return></verificarComunicacionResponse></Body></Envelope>`)
			return
		}
		if strings.Contains(string(body), "<ns:cufd>") {
			io.WriteString(w, `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body><cufdResponse xmlns="https://siat.impuestos.gob.bo/"><RespuestaCufd><codigo>CUFD</codigo><codigoControl>CONTROL</codigoControl><codigoInterno>INTERNO</codigoInterno><transaccion>true</transaccion></RespuestaCufd></cufdResponse></Body></Envelope>`)
			return
		}
		if strings.Contains(string(body), "<ns:sincronizarListaProductosServicios>") {
			io.WriteString(w, `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body><sincronizarListaProductosServiciosResponse xmlns="https://siat.impuestos.gob.bo/"><RespuestaListaProductos><transaccion>true</transaccion><listaCodigos><codigoActividad>001234</codigoActividad><codigoProducto>123</codigoProducto><descripcionProducto>Producto</descripcionProducto><nandina>01</nandina><nandina>02</nandina></listaCodigos></RespuestaListaProductos></sincronizarListaProductosServiciosResponse></Body></Envelope>`)
			return
		}
		io.WriteString(w, `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body><recepcionPaqueteFacturaResponse xmlns="https://siat.impuestos.gob.bo/"><RespuestaServicioFacturacion><codigoEstado>901</codigoEstado><codigoRecepcion>RECIBO</codigoRecepcion><mensajesList><codigo>1001</codigo><descripcion>Detalle</descripcion><advertencia>true</advertencia><numeroArchivo>2</numeroArchivo><numeroDetalle>3</numeroDetalle></mensajesList><transaccion>true</transaccion></RespuestaServicioFacturacion></recepcionPaqueteFacturaResponse></Body></Envelope>`)
	}))
	t.Cleanup(server.Close)
	return server, captures
}

func siatTestConfig(url string) siat.Config {
	return siat.Config{BaseURL: url, Token: "test", Nit: 123456789, CodigoSistema: "SISTEMA", CodigoAmbiente: siat.AmbientePruebas}
}

// siatRequestFields checks the actual SOAP namespace and ordered request children.
func siatRequestFields(t *testing.T, body, operation, request string) []string {
	t.Helper()
	d := xml.NewDecoder(strings.NewReader(body))
	var fields []string
	depth, requestDepth := 0, -1
	found := false
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		switch e := token.(type) {
		case xml.StartElement:
			depth++
			if e.Name.Local == operation {
				found = true
				if e.Name.Space != "https://siat.impuestos.gob.bo/" {
					t.Errorf("operation namespace = %q", e.Name.Space)
				}
			}
			if e.Name.Local == request {
				requestDepth = depth
			} else if requestDepth >= 0 && depth == requestDepth+1 {
				fields = append(fields, e.Name.Local)
			}
		case xml.EndElement:
			if depth == requestDepth {
				requestDepth = -1
			}
			depth--
		}
	}
	if !found {
		t.Errorf("operation %s not found in SOAP body", operation)
	}
	return fields
}

func TestSIAT_ServicesAndPackageContract(t *testing.T) {
	server, captures := siatTestServer(t)
	client, err := siat.New(siatTestConfig(server.URL + "/sfd"))
	if err != nil {
		t.Fatal(err)
	}
	request := models.NewRecepcionPaqueteFacturaBuilder().
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoDocumentoSector(1).WithTipoFacturaDocumento(1).
		WithCuis("CUIS").WithCufd("CUFD").WithArchivo("QQ==").
		WithHashArchivo("HASH").
		WithCantidadFacturas(10).WithCodigoEvento(123).Build()
	for _, tc := range []struct {
		name string
		send func() error
	}{
		{"ServicioFacturacionCompraVenta", func() error { _, err := client.CompraVenta().RecepcionPaqueteFactura(t.Context(), request); return err }},
		{"ServicioFacturacionElectronica", func() error { _, err := client.Electronica().RecepcionPaqueteFactura(t.Context(), request); return err }},
		{"ServicioFacturacionComputarizada", func() error {
			_, err := client.Computarizada().RecepcionPaqueteFactura(t.Context(), request)
			return err
		}},
		{"ServicioFacturacionServicioBasico", func() error {
			_, err := client.ServicioBasico().RecepcionPaqueteFactura(t.Context(), request)
			return err
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.send(); err != nil {
				t.Fatal(err)
			}
			got := <-captures
			if got.path != "/sfd/"+tc.name || got.token != "TokenApi test" {
				t.Errorf("unexpected endpoint/auth: %s %s", got.path, got.token)
			}
			fields := siatRequestFields(t, got.body, "recepcionPaqueteFactura", "SolicitudServicioRecepcionPaquete")
			want := []string{"codigoAmbiente", "codigoDocumentoSector", "codigoEmision", "codigoModalidad", "codigoPuntoVenta", "codigoSistema", "codigoSucursal", "cufd", "cuis", "nit", "tipoFacturaDocumento", "archivo", "fechaEnvio", "hashArchivo", "cantidadFacturas", "codigoEvento"}
			if !slices.Equal(fields, want) {
				t.Errorf("SOAP fields = %v; want %v", fields, want)
			}
			if !strings.Contains(got.body, "<nit>123456789</nit>") || !strings.Contains(got.body, "<codigoPuntoVenta>0</codigoPuntoVenta>") {
				t.Error("global identity or zero-valued point of sale missing")
			}
		})
	}

}

func TestSIAT_CufdAndResponseDetails(t *testing.T) {
	server, captures := siatTestServer(t)
	client, err := siat.New(siatTestConfig(server.URL + "/sfd"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Codigos().SolicitudCufd(t.Context(), models.NewCufdBuilder().WithCuis("CUIS").Build())
	if err != nil {
		t.Fatal(err)
	}
	if got := <-captures; got.path != "/sfd/FacturacionCodigos" {
		t.Errorf("endpoint = %s", got.path)
	}
	if resp.Body.Content.RespuestaCufd.CodigoInterno != "INTERNO" {
		t.Error("codigoInterno not decoded")
	}
	result, err := client.CompraVenta().RecepcionPaqueteFactura(t.Context(), models.NewRecepcionPaqueteFacturaBuilder().Build())
	if err != nil {
		t.Fatal(err)
	}
	<-captures
	messages := result.Body.Content.RespuestaServicioFacturacion.MensajesList
	if len(messages) != 1 || !messages[0].Advertencia || messages[0].NumeroArchivo != 2 || messages[0].NumeroDetalle != 3 {
		t.Fatalf("response details lost: %+v", messages)
	}
}

func TestSIAT_InvalidConfig(t *testing.T) {
	if client, err := siat.New(siat.Config{}); err == nil || client != nil {
		t.Fatal("invalid config accepted")
	}
}

func TestSIAT_AllServiceEndpoints(t *testing.T) {
	server, captures := siatTestServer(t)
	client, err := siat.New(siatTestConfig(" " + server.URL + "/sfd/ "))
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name string
		send func() error
	}{
		{"FacturacionCodigos", func() error {
			_, err := client.Codigos().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionCodigosBuilder().Build())
			return err
		}},
		{"FacturacionSincronizacion", func() error {
			_, err := client.Sincronizacion().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionSincronizacionBuilder().Build())
			return err
		}},
		{"FacturacionOperaciones", func() error {
			_, err := client.Operaciones().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionOperacionesBuilder().Build())
			return err
		}},
		{"ServicioRecepcionCompras", func() error {
			_, err := client.RecepcionCompras().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionRecepcionComprasBuilder().Build())
			return err
		}},
		{"ServicioFacturacionDocumentoAjuste", func() error {
			_, err := client.DocumentoAjuste().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionDocumentoAjusteBuilder().Build())
			return err
		}},
		{"ServicioFacturacionCompraVenta", func() error {
			_, err := client.CompraVenta().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionFacturacion())
			return err
		}},
		{"ServicioFacturacionElectronica", func() error {
			_, err := client.Electronica().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionFacturacion())
			return err
		}},
		{"ServicioFacturacionComputarizada", func() error {
			_, err := client.Computarizada().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionFacturacion())
			return err
		}},
		{"ServicioFacturacionServicioBasico", func() error {
			_, err := client.ServicioBasico().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionFacturacion())
			return err
		}},
		{"ServicioFacturacionTelecomunicaciones", func() error {
			_, err := client.Telecomunicaciones().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionFacturacion())
			return err
		}},
		{"ServicioFacturacionEntidadFinanciera", func() error {
			_, err := client.EntidadFinanciera().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionFacturacion())
			return err
		}},
		{"ServicioFacturacionBoletoAereo", func() error {
			_, err := client.BoletoAereo().VerificarComunicacion(t.Context(), models.NewVerificarComunicacionFacturacion())
			return err
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.send(); err != nil {
				t.Fatal(err)
			}
			got := <-captures
			if got.path != "/sfd/"+tc.name {
				t.Errorf("endpoint = %s", got.path)
			}
			siatRequestFields(t, got.body, "verificarComunicacion", "")
		})
	}
}

func TestSIAT_DynamicIdentityDoesNotMutatePackage(t *testing.T) {
	server, captures := siatTestServer(t)
	client, err := siat.New(siatTestConfig(server.URL + "/sfd"))
	if err != nil {
		t.Fatal(err)
	}
	request := models.NewRecepcionPaqueteFacturaBuilder().Build()
	ctx := siat.WithDynamicConfig(t.Context(), siat.Config{Nit: 999, Token: "tenant", CodigoSistema: "TENANT"})
	if _, err := client.CompraVenta().RecepcionPaqueteFactura(ctx, request); err != nil {
		t.Fatal(err)
	}
	got := <-captures
	if got.token != "TokenApi tenant" || !strings.Contains(got.body, "<nit>999</nit>") || !strings.Contains(got.body, "<codigoSistema>TENANT</codigoSistema>") {
		t.Error("dynamic identity not applied")
	}
	if _, err := client.CompraVenta().RecepcionPaqueteFactura(t.Context(), request); err != nil {
		t.Fatal(err)
	}
	got = <-captures
	if got.token != "TokenApi test" || !strings.Contains(got.body, "<nit>123456789</nit>") {
		t.Error("dynamic identity leaked into next request")
	}
}

func TestSIAT_ProductResponse(t *testing.T) {
	server, captures := siatTestServer(t)
	client, err := siat.New(siatTestConfig(server.URL + "/sfd"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Sincronizacion().SincronizarListaProductosServicios(t.Context(), models.NewSincronizarListaProductosServiciosBuilder().Build())
	if err != nil {
		t.Fatal(err)
	}
	<-captures
	products := resp.Body.Content.RespuestaListaProductos.ListaCodigos
	if len(products) != 1 {
		t.Fatalf("products = %+v", products)
	}
	if products[0].CodigoActividad != "001234" || products[0].CodigoProducto != 123 || !slices.Equal(products[0].Nandina, []string{"01", "02"}) {
		t.Errorf("product fields lost: %+v", products[0])
	}
}

func TestSIAT_ZeroRequestReturnsError(t *testing.T) {
	client, err := siat.New(siatTestConfig("http://unused.invalid/sfd"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.CompraVenta().RecepcionFactura(t.Context(), models.RecepcionFactura{}); err == nil {
		t.Fatal("zero request accepted")
	}
}
