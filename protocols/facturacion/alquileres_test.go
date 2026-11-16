package facturacion_test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/alquileres"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
)

func newTestAlquilerFactura(t *testing.T, sucursal int, cufd string) *alquileres.Factura {
	t.Helper()
	cab := alquileres.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "ARRENDADORA S.A.",
		Municipio:                    "La Paz",
		NumeroFactura:                1001,
		Cuf:                          "CUF-ALQ-123",
		Cufd:                         cufd,
		CodigoSucursal:               sucursal,
		Direccion:                    "Av. Arce 123",
		FechaEmision:                 time.Date(2026, 10, 31, 10, 0, 0, 0, time.UTC),
		CodigoTipoDocumentoIdentidad: 1,
		NumeroDocumento:              "9876543",
		CodigoCliente:                "CLI-01",
		PeriodoFacturado:             "OCTUBRE 2026",
		CodigoMetodoPago:             1,
		SubTotalNeto:                 100.0,
		TotalPrecioNeto:              100.0,
		Iva:                          13.0,
		MontoAPagar:                  100.0,
		CodigoMoneda:                 1,
		TipoCambio:                   1.0,
		MontoTotalMoneda:             100.0,
		Leyenda:                      "Ley N 453",
		Usuario:                      "operador",
		CodigoDocumentoSector:        2,
	}
	detalles := []alquileres.Detalle{
		{
			ActividadEconomica: "681000",
			CodigoProductoSin:  99100,
			CodigoProducto:     "ALQ-01",
			Descripcion:        "Alquiler Octubre",
			Cantidad:           1.0,
			UnidadMedida:       58,
			PrecioNetoUnitario: 100.0,
			SubTotal:           100.0,
		},
	}
	f, err := alquileres.New(cab, detalles)
	if err != nil {
		t.Fatalf("crear factura de prueba: %v", err)
	}
	return f
}

func TestFacturaAlquilerSeEnviaAlServicioComputarizado(t *testing.T) {
	llamadas := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		llamadas++
		if r.URL.Path != "/sfd/ServicioFacturacionComputarizada" {
			t.Errorf("ruta: %s", r.URL.Path)
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		if !strings.Contains(string(data), "<codigoDocumentoSector>2</codigoDocumentoSector>") {
			t.Errorf("solicitud: %s", data)
		}
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><recepcionFacturaResponse><RespuestaServicioFacturacion><codigoRecepcion>REC</codigoRecepcion><transaccion>true</transaccion></RespuestaServicioFacturacion></recepcionFacturaResponse></soap:Body></soap:Envelope>`)
	}))
	defer srv.Close()
	cfg := facturacion.Config{BaseURL: srv.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2}
	client, err := facturacion.NewComputarizadaClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	factura := newTestAlquilerFactura(t, 0, "cufd")
	req, err := (facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2, CodigoEmision: 1, Cufd: "cufd", Cuis: "cuis", FechaEnvio: time.Now()}).WithFactura(factura, nil)
	if err != nil {
		t.Fatal(err)
	}
	if req.CodigoDocumentoSector != 2 || req.TipoFacturaDocumento != 1 {
		t.Fatalf("solicitud: %+v", req)
	}
	gz, err := gzip.NewReader(bytes.NewReader(req.Archivo))
	if err != nil {
		t.Fatal(err)
	}
	defer gz.Close()
	var doc struct{ XMLName xml.Name }
	if err := xml.NewDecoder(gz).Decode(&doc); err != nil || doc.XMLName.Local != "facturaComputarizadaAlquilerBienInmueble" {
		t.Fatalf("documento: %+v %v", doc, err)
	}
	result, err := client.RecepcionFactura(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	if result.Transaccion == nil || !*result.Transaccion || llamadas != 1 {
		t.Fatalf("respuesta: %+v llamadas=%d", result, llamadas)
	}
	compra, err := facturacion.NewClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = compra.RecepcionFactura(t.Context(), req); err == nil {
		t.Fatal("envió sector 2 a compra-venta")
	}
	electronica, err := facturacion.NewElectronicaClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = electronica.RecepcionFactura(t.Context(), req); err == nil {
		t.Fatal("envió modalidad computarizada a electrónica")
	}
	if llamadas != 1 {
		t.Fatalf("llamadas adicionales: %d", llamadas)
	}
}

func TestPaqueteAlquiler(t *testing.T) {
	req := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2}}
	f1 := newTestAlquilerFactura(t, 0, "cufd")
	f2 := newTestAlquilerFactura(t, 0, "cufd")
	paquete, err := req.WithFacturas(t.Context(), facturacion.AsDocumentos([]*alquileres.Factura{f1, f2}), nil)
	if err != nil {
		t.Fatal(err)
	}
	if paquete.CodigoDocumentoSector != 2 || paquete.CantidadFacturas != 2 {
		t.Fatalf("paquete: %+v", paquete)
	}
	gz, err := gzip.NewReader(bytes.NewReader(paquete.Archivo))
	if err != nil {
		t.Fatal(err)
	}
	defer gz.Close()
	reader := tar.NewReader(gz)
	for _, name := range []string{"factura_1.xml", "factura_2.xml"} {
		header, err := reader.Next()
		if err != nil || header.Name != name {
			t.Fatalf("archivo: %+v %v", header, err)
		}
	}

	// Modalidad electrónica sin signer debe fallar
	reqElec := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 1}}
	if _, err = reqElec.WithFacturas(t.Context(), facturacion.AsDocumentos([]*alquileres.Factura{f1}), nil); err == nil {
		t.Fatal("aceptó modalidad electrónica sin signer")
	}
}

func TestAlquilerRechazaSucursalDistinta(t *testing.T) {
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2, CodigoSucursal: 1}
	factura := newTestAlquilerFactura(t, 2, "cufd")
	if _, err := req.WithFactura(factura, nil); err == nil {
		t.Fatal("aceptó sucursal distinta")
	}
	paquete := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: req}
	if _, err := paquete.WithFacturas(t.Context(), facturacion.AsDocumentos([]*alquileres.Factura{factura}), nil); err == nil {
		t.Fatal("aceptó sucursal distinta en paquete")
	}
}
