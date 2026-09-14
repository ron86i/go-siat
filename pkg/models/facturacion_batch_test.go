package models

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"io"
	"math/big"
	"runtime"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v2/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type batchTestInvoice struct {
	XMLName struct{} `xml:"factura"`
	Nro     int      `xml:"nro"`
}

type preparedBatchInvoice struct {
	document []byte
	calls    int32
}

func (f *preparedBatchInvoice) MarshalXMLBytes() ([]byte, error) {
	atomic.AddInt32(&f.calls, 1)
	return f.document, nil
}

type concurrentBatchSigner struct {
	active int32
	max    int32
}

func (s *concurrentBatchSigner) ConcurrentXMLSigning() bool { return true }

func (s *concurrentBatchSigner) SignXML(document []byte) ([]byte, error) {
	active := atomic.AddInt32(&s.active, 1)
	defer atomic.AddInt32(&s.active, -1)
	for {
		current := atomic.LoadInt32(&s.max)
		if active <= current || atomic.CompareAndSwapInt32(&s.max, current, active) {
			break
		}
	}
	time.Sleep(5 * time.Millisecond)
	return append([]byte("<firmada>"), append(document, []byte("</firmada>")...)...), nil
}

func TestRecepcionMasivaFacturaBuilder_WithFacturasEnLote(t *testing.T) {
	facturas := make([]any, 8)
	for i := range facturas {
		facturas[i] = batchTestInvoice{Nro: i + 1}
	}

	signer := &concurrentBatchSigner{}
	builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadElectronica)
	require.NoError(t, builder.WithFacturasEnLote(facturas, signer, &FacturasEnLoteOptions{Workers: 3}))

	request := builder.request.SolicitudServicioRecepcionMasiva
	assert.Equal(t, len(facturas), request.CantidadFacturas)
	assert.GreaterOrEqual(t, atomic.LoadInt32(&signer.max), int32(2))
	assert.LessOrEqual(t, atomic.LoadInt32(&signer.max), int32(3))

	archive, err := base64.StdEncoding.DecodeString(request.SolicitudRecepcionFactura.Archivo)
	require.NoError(t, err)
	gr, err := gzip.NewReader(bytes.NewReader(archive))
	require.NoError(t, err)
	tr := tar.NewReader(gr)
	for i := range facturas {
		header, err := tr.Next()
		require.NoError(t, err)
		assert.Equal(t, "factura_"+strconv.Itoa(i+1)+".xml", header.Name)
		document, err := io.ReadAll(tr)
		require.NoError(t, err)
		assert.Contains(t, string(document), "<nro>"+strconv.Itoa(i+1)+"</nro>")
	}
	_, err = tr.Next()
	assert.Equal(t, io.EOF, err)
}

func TestRecepcionMasivaFacturaBuilder_WithFacturasEnLoteXMLPreconstruido(t *testing.T) {
	factura := &preparedBatchInvoice{document: []byte("<factura><nro>1</nro></factura>")}
	builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadComputarizada)
	require.NoError(t, builder.WithFacturasEnLote([]any{factura}, nil, &FacturasEnLoteOptions{Workers: 1}))
	assert.Equal(t, int32(1), atomic.LoadInt32(&factura.calls))

	archive, err := base64.StdEncoding.DecodeString(builder.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo)
	require.NoError(t, err)
	gr, err := gzip.NewReader(bytes.NewReader(archive))
	require.NoError(t, err)
	tr := tar.NewReader(gr)
	_, err = tr.Next()
	require.NoError(t, err)
	document, err := io.ReadAll(tr)
	require.NoError(t, err)
	assert.Equal(t, factura.document, document)
}

func TestFacturasEnLoteOptions_ProtegeFirmadoresNoConcurrentes(t *testing.T) {
	facturas := []any{batchTestInvoice{Nro: 1}, batchTestInvoice{Nro: 2}}
	signer := &serialBatchSigner{}
	builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadElectronica)
	require.NoError(t, builder.WithFacturasEnLote(facturas, signer, &FacturasEnLoteOptions{Workers: 4}))
	assert.Equal(t, int32(1), atomic.LoadInt32(&signer.max))
}

func TestWorkersFacturas(t *testing.T) {
	options := &FacturasEnLoteOptions{Workers: 4}
	assert.Equal(t, 4, workersFacturas(10, &concurrentBatchSigner{}, true, options))
	assert.Equal(t, 1, workersFacturas(10, &serialBatchSigner{}, true, options))
	assert.Equal(t, 4, workersFacturas(10, nil, false, options))
	assert.Equal(t, min(10, max(1, runtime.GOMAXPROCS(0)-1)), workersFacturas(10, &concurrentBatchSigner{}, true, nil))
}

func TestRecepcionMasivaFacturaBuilder_WithFacturasEnLoteContextCancelada(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadComputarizada)
	err := builder.WithFacturasEnLoteContext(ctx, []any{batchTestInvoice{Nro: 1}}, nil, nil)
	require.ErrorIs(t, err, context.Canceled)
	request := builder.request.SolicitudServicioRecepcionMasiva
	assert.Empty(t, request.SolicitudRecepcionFactura.Archivo)
	assert.Empty(t, request.SolicitudRecepcionFactura.HashArchivo)
	assert.Zero(t, request.CantidadFacturas)
}

func TestRecepcionMasivaFacturaBuilder_WithFacturasEnLoteOptions(t *testing.T) {
	facturas := []any{batchTestInvoice{Nro: 1}, batchTestInvoice{Nro: 2}}
	builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadComputarizada)
	err := builder.WithFacturasEnLote(facturas, nil, &FacturasEnLoteOptions{MaxFacturas: 1})
	require.ErrorContains(t, err, "supera el máximo configurado")

	var persisted bytes.Buffer
	var metrics FacturasEnLoteMetrics
	builder = NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadComputarizada)
	require.NoError(t, builder.WithFacturasEnLote(facturas, nil, &FacturasEnLoteOptions{
		Workers:        1,
		DestinoArchivo: &persisted,
		OnComplete: func(value FacturasEnLoteMetrics) {
			metrics = value
		},
	}))
	request := builder.request.SolicitudServicioRecepcionMasiva
	archive, err := base64.StdEncoding.DecodeString(request.SolicitudRecepcionFactura.Archivo)
	require.NoError(t, err)
	assert.Equal(t, archive, persisted.Bytes())
	assert.Equal(t, int64(len(archive)), metrics.TamanoArchivo)
	assert.Equal(t, len(facturas), metrics.CantidadFacturas)
	assert.Equal(t, 1, metrics.Workers)
	assert.Positive(t, metrics.DuracionTotal)
}

func BenchmarkRecepcionMasivaFacturaBuilder_WithFacturasEnLote(b *testing.B) {
	for _, count := range []int{100, 500, 1000} {
		b.Run(strconv.Itoa(count)+"_facturas", func(b *testing.B) {
			facturas := make([]any, count)
			for index := range facturas {
				facturas[index] = batchTestInvoice{Nro: index + 1}
			}
			options := &FacturasEnLoteOptions{Workers: 1}
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadComputarizada)
				if err := builder.WithFacturasEnLote(facturas, nil, options); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
	b.Run("1000_xml_preconstruido", func(b *testing.B) {
		facturas := make([]any, 1000)
		for index := range facturas {
			facturas[index] = &preparedBatchInvoice{document: []byte("<factura><nro>1</nro></factura>")}
		}
		options := &FacturasEnLoteOptions{Workers: 1}
		b.ReportAllocs()
		b.ResetTimer()
		for iteration := 0; iteration < b.N; iteration++ {
			builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadComputarizada)
			if err := builder.WithFacturasEnLote(facturas, nil, options); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkRecepcionMasivaFacturaBuilder_WithFacturasEnLoteFirmaElectronica(b *testing.B) {
	signer := newBenchmarkXMLSigner(b)
	for _, count := range []int{100, 500, 1000} {
		b.Run(strconv.Itoa(count)+"_facturas_firmadas", func(b *testing.B) {
			facturas := make([]any, count)
			for index := range facturas {
				facturas[index] = &preparedBatchInvoice{document: []byte("<factura><nro>1</nro></factura>")}
			}
			options := &FacturasEnLoteOptions{Workers: max(1, runtime.GOMAXPROCS(0)-1)}
			b.ReportAllocs()
			b.ResetTimer()
			for iteration := 0; iteration < b.N; iteration++ {
				builder := NewRecepcionMasivaFacturaBuilder().WithCodigoModalidad(ModalidadElectronica)
				if err := builder.WithFacturasEnLote(facturas, signer, options); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func newBenchmarkXMLSigner(b *testing.B) *utils.XMLDocumentSigner {
	b.Helper()
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		b.Fatal(err)
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "go-siat benchmark"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certificateDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		b.Fatal(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	certificatePEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificateDER})
	signer, err := utils.NewXMLDocumentSigner(keyPEM, certificatePEM)
	if err != nil {
		b.Fatal(err)
	}
	return signer
}

type serialBatchSigner struct {
	active int32
	max    int32
}

func (s *serialBatchSigner) SignXML(document []byte) ([]byte, error) {
	active := atomic.AddInt32(&s.active, 1)
	defer atomic.AddInt32(&s.active, -1)
	atomic.StoreInt32(&s.max, active)
	return document, nil
}
