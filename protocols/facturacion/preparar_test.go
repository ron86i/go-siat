package facturacion_test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
)

type signerFunc func([]byte) ([]byte, error)

func (f signerFunc) SignXML(b []byte) ([]byte, error) { return f(b) }

func TestPrepararFactura(t *testing.T) {
	factura := newTestFactura(t)
	original := facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2} // Computarizada
	req, err := original.WithFactura(factura, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(original.Archivo) != 0 || req.CodigoDocumentoSector != 1 || req.TipoFacturaDocumento != 1 {
		t.Fatalf("solicitud: %+v", req)
	}
	if want := fmt.Sprintf("%x", sha256.Sum256(req.Archivo)); req.HashArchivo != want {
		t.Fatalf("hash %s != %s", req.HashArchivo, want)
	}
	reader, err := gzip.NewReader(bytes.NewReader(req.Archivo))
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	want, err := factura.XML(2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, want) {
		t.Fatalf("XML descomprimido no coincide")
	}
}

func TestPrepararFacturaExigeModalidadEnSolicitud(t *testing.T) {
	factura := newTestFactura(t)
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: 0}
	if _, err := req.WithFactura(factura, nil); err == nil {
		t.Fatal("debe rechazar solicitud sin CodigoModalidad")
	}
}

func TestPrepararFacturaRechazaPunteroNilTipado(t *testing.T) {
	var factura *compraventa.Factura
	var doc facturacion.Documento = factura
	_, err := (facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2}).PrepararFactura(doc, nil)
	if err == nil {
		t.Fatal("se esperaba error para Documento que contiene un puntero nil")
	}
}

func TestPrepararFacturaConflictoModalidad(t *testing.T) {
	client, err := facturacion.NewComputarizadaClient(facturacion.Config{
		BaseURL:        "http://localhost",
		Token:          "token",
		Nit:            123,
		CodigoSistema:  "sistema",
		CodigoAmbiente: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	factura := newTestFactura(t)
	// Solicitud especifica Electronica (1) pero el cliente es Computarizada (2)
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: facturacion.ModalidadElectronica}
	_, err = client.PrepararFactura(req, factura, nil)
	if err == nil || !strings.Contains(err.Error(), "conflicto de modalidad") {
		t.Fatalf("se esperaba error de conflicto de modalidad, obtuvo: %v", err)
	}
}

func TestPrepararFacturaComputarizadaRechazaFirma(t *testing.T) {
	factura := newTestFactura(t)
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: facturacion.ModalidadComputarizada}
	signer := signerFunc(func(d []byte) ([]byte, error) { return d, nil })
	_, err := req.WithFactura(factura, signer)
	if err == nil || !strings.Contains(err.Error(), "no admite firma digital") {
		t.Fatalf("se esperaba rechazo de firma en computarizada, obtuvo: %v", err)
	}
}

func TestPrepararFacturaElectronica(t *testing.T) {
	factura := newTestFactura(t)
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: 1} // Electrónica
	if _, err := req.WithFactura(factura, nil); err == nil {
		t.Fatal("firma requerida")
	}
	signed := false
	result, err := req.WithFactura(factura, signerFunc(func(data []byte) ([]byte, error) { signed = true; return append(data, []byte("firma")...), nil }))
	if err != nil || !signed {
		t.Fatalf("firma: %v", err)
	}
	reader, err := gzip.NewReader(bytes.NewReader(result.Archivo))
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil || !bytes.HasSuffix(data, []byte("firma")) {
		t.Fatalf("XML firmado: %v", err)
	}
	if _, err := (facturacion.SolicitudRecepcionFactura{CodigoModalidad: 1, CodigoDocumentoSector: 2}).WithFactura(factura, nil); err == nil {
		t.Fatal("sector incompatible")
	}
}

func TestFacturaPreparadaHastaRespuestaSOAP(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var env struct {
			Body struct {
				Recepcion struct {
					Solicitud struct {
						Archivo string `xml:"archivo"`
						Hash    string `xml:"hashArchivo"`
					} `xml:"SolicitudServicioRecepcionFactura"`
				} `xml:"recepcionFactura"`
			} `xml:"Body"`
		}
		if err := xml.NewDecoder(r.Body).Decode(&env); err != nil {
			t.Errorf("SOAP: %v", err)
			return
		}
		archivo, err := base64.StdEncoding.DecodeString(env.Body.Recepcion.Solicitud.Archivo)
		if err != nil {
			t.Errorf("archivo: %v", err)
			return
		}
		if want := fmt.Sprintf("%x", sha256.Sum256(archivo)); env.Body.Recepcion.Solicitud.Hash != want {
			t.Errorf("hash incorrecto")
		}
		gz, err := gzip.NewReader(bytes.NewReader(archivo))
		if err != nil {
			t.Errorf("gzip: %v", err)
			return
		}
		defer gz.Close()
		var doc struct{ XMLName xml.Name }
		if err := xml.NewDecoder(gz).Decode(&doc); err != nil || doc.XMLName.Local != "facturaComputarizadaCompraVenta" {
			t.Errorf("documento=%s error=%v", doc.XMLName.Local, err)
		}
		io.WriteString(w, `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><recepcionFacturaResponse><RespuestaServicioFacturacion><codigoEstado>908</codigoEstado><codigoRecepcion>REC-PRUEBA</codigoRecepcion><transaccion>true</transaccion></RespuestaServicioFacturacion></recepcionFacturaResponse></soap:Body></soap:Envelope>`)
	}))
	defer server.Close()
	client, err := facturacion.NewClient(facturacion.Config{BaseURL: server.URL, Token: "token", Nit: 123, CodigoSistema: "sistema", CodigoAmbiente: 2})
	if err != nil {
		t.Fatal(err)
	}
	factura := newTestFactura(t)
	req, err := (facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2, CodigoEmision: 1, Cufd: "cufd", Cuis: "cuis", FechaEnvio: time.Now()}).WithFactura(factura, nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.RecepcionFactura(t.Context(), req)
	if err != nil {
		t.Fatal(err)
	}
	if response.Transaccion == nil || !*response.Transaccion || response.CodigoRecepcion != "REC-PRUEBA" {
		t.Fatalf("respuesta: %+v", response)
	}
}

func TestPrepararPaquete(t *testing.T) {
	f1 := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		(*d)[0].Descripcion = "primera"
	})
	f2 := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		c.NumeroFactura = 101
		(*d)[0].Descripcion = "segunda"
	})
	facturas := []*compraventa.Factura{f1, f2}
	original := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2}}
	req, err := original.WithFacturas(t.Context(), facturacion.AsDocumentos(facturas), nil)
	if err != nil {
		t.Fatal(err)
	}
	if req.CantidadFacturas != 2 || req.CodigoDocumentoSector != 1 || req.TipoFacturaDocumento != 1 || len(original.Archivo) != 0 {
		t.Fatalf("solicitud: %+v", req)
	}
	if want := fmt.Sprintf("%x", sha256.Sum256(req.Archivo)); req.HashArchivo != want {
		t.Fatal("hash incorrecto")
	}
	gz, err := gzip.NewReader(bytes.NewReader(req.Archivo))
	if err != nil {
		t.Fatal(err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	for i, want := range []string{"primera", "segunda"} {
		h, err := tr.Next()
		if err != nil {
			t.Fatal(err)
		}
		if h.Name != fmt.Sprintf("factura_%d.xml", i+1) {
			t.Fatalf("nombre %s", h.Name)
		}
		data, err := io.ReadAll(tr)
		if err != nil || !strings.Contains(string(data), want) {
			t.Fatalf("XML %s: %v", data, err)
		}
	}
	if _, err := tr.Next(); !errors.Is(err, io.EOF) {
		t.Fatalf("archivo adicional: %v", err)
	}
}

func TestPrepararPaqueteRechazaErroresSinModificarSolicitud(t *testing.T) {
	original := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2, Archivo: []byte("anterior"), HashArchivo: "anterior"}, CantidadFacturas: 7}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	factura := newTestFactura(t)
	result, err := original.WithFacturas(ctx, facturacion.AsDocumentos([]*compraventa.Factura{factura}), nil)
	if !errors.Is(err, context.Canceled) || string(result.Archivo) != "anterior" || result.HashArchivo != "anterior" || result.CantidadFacturas != 7 {
		t.Fatalf("cancelación: %+v %v", result, err)
	}
	if _, err := original.WithFacturas(t.Context(), nil, nil); err == nil {
		t.Fatal("aceptó paquete vacío")
	}
	electronic := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 1}}
	if _, err := electronic.WithFacturas(t.Context(), facturacion.AsDocumentos([]*compraventa.Factura{factura}), nil); err == nil {
		t.Fatal("aceptó factura electrónica sin firma")
	}
	calls := 0
	signed, err := electronic.WithFacturas(t.Context(), facturacion.AsDocumentos([]*compraventa.Factura{factura, factura}), signerFunc(func(b []byte) ([]byte, error) { calls++; return append(b, []byte("firma")...), nil }))
	if err != nil || calls != 2 || signed.CantidadFacturas != 2 {
		t.Fatalf("firmas=%d error=%v", calls, err)
	}
}

func TestPreparacionesRechazanCantidadesSobreLimite(t *testing.T) {
	factura := newTestFactura(t)
	paquete := make([]*compraventa.Factura, 501)
	for i := range paquete {
		paquete[i] = factura
	}
	if _, err := (facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2}}).WithFacturas(t.Context(), facturacion.AsDocumentos(paquete), nil); err == nil {
		t.Fatal("se aceptó un paquete de más de 500 facturas")
	}

	masiva := make([]*compraventa.Factura, 1001)
	for i := range masiva {
		masiva[i] = factura
	}
	if _, err := (facturacion.SolicitudRecepcionMasiva{SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2}}).WithFacturas(t.Context(), facturacion.AsDocumentos(masiva), nil); err == nil {
		t.Fatal("se aceptó una masiva de más de 1000 facturas")
	}
}

func TestPreparacionCompruebaUbicacionYCUFD(t *testing.T) {
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2, CodigoSucursal: 1, CodigoPuntoVenta: new(2), Cufd: "VIGENTE"}
	factura := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		c.CodigoSucursal = 1
		c.CodigoPuntoVenta = new(2)
		c.Cufd = "VIGENTE"
	})
	if _, err := req.WithFactura(factura, nil); err != nil {
		t.Fatal(err)
	}

	fDiffSuc := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		c.CodigoSucursal = 0
		c.CodigoPuntoVenta = new(2)
		c.Cufd = "VIGENTE"
	})
	if _, err := req.WithFactura(fDiffSuc, nil); err == nil {
		t.Fatal("aceptó sucursal distinta")
	}

	fDiffPunto := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		c.CodigoSucursal = 1
		c.CodigoPuntoVenta = new(3)
		c.Cufd = "VIGENTE"
	})
	if _, err := req.WithFactura(fDiffPunto, nil); err == nil {
		t.Fatal("aceptó punto de venta distinto")
	}

	fDiffCUFD := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		c.CodigoSucursal = 1
		c.CodigoPuntoVenta = new(2)
		c.Cufd = "EVENTO"
	})
	if _, err := req.WithFactura(fDiffCUFD, nil); err == nil {
		t.Fatal("aceptó CUFD individual distinto")
	}

	paquete := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: req}
	if _, err := paquete.WithFacturas(t.Context(), facturacion.AsDocumentos([]*compraventa.Factura{fDiffCUFD}), nil); err != nil {
		t.Fatalf("el paquete debe admitir CUFD del evento distinto: %v", err)
	}

	fDiffPaqSuc := newTestFactura(t, func(c *compraventa.Cabecera, d *[]compraventa.Detalle) {
		c.CodigoSucursal = 9
		c.CodigoPuntoVenta = new(2)
		c.Cufd = "VIGENTE"
	})
	if _, err := paquete.WithFacturas(t.Context(), facturacion.AsDocumentos([]*compraventa.Factura{fDiffPaqSuc}), nil); err == nil {
		t.Fatal("aceptó sucursal distinta en paquete")
	}
}
