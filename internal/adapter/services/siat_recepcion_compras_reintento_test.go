package services_test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/beevik/etree"
	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

// TestSiatRecepcionCompras_ReintentarPaquete reutiliza la solicitud SOAP guardada
// en SIAT_COMPRA_XML (por defecto .siat-traces/05-solicitud.xml). Renueva sólo
// fechaEnvio y reutiliza factura, CUIS y CUFD. Requiere SIAT_RUN=1 y ambiente 2.
// SIAT_COMPRA_NOMBRE permite cambiar únicamente el nombre de la entrada TAR
// para diagnosticar rechazos; por sí solo conserva el XML byte por byte.
// SIAT_COMPRA_TRAZA=0 permite comparar el transporte HTTP sin el capturador.
// SIAT_COMPRA_CONTROL cambia codigoControl sin volver a emitir la factura.
// SIAT_COMPRA_CONTROL_NULO=1 compara codigoControl vacío con xsi:nil="true".
// SIAT_COMPRA_DECIMALES=2 serializa los importes monetarios con dos decimales.
// SIAT_COMPRA_FECHA_SIN_ZONA=1 conserva la hora local y quita zona y fracciones de fechaEmision.
func TestSiatRecepcionCompras_ReintentarPaquete(t *testing.T) {
	_ = godotenv.Load(".env")
	if os.Getenv("SIAT_RUN") != "1" {
		t.Skip("defina SIAT_RUN=1 para reintentar un paquete de compras de pruebas")
	}
	source := os.Getenv("SIAT_COMPRA_XML")
	if source == "" {
		source = ".siat-traces/05-solicitud.xml"
	}
	data, err := os.ReadFile(source)
	require.NoError(t, err)
	var envelope struct {
		Solicitud facturacion.SolicitudRecepcionCompras `xml:"Body>recepcionPaqueteCompras>SolicitudRecepcionCompras"`
	}
	require.NoError(t, xml.Unmarshal(data, &envelope))
	original := envelope.Solicitud
	require.NotEmpty(t, original.Archivo, "el XML debe contener recepcionPaqueteCompras")
	cfg := siatRecepcionComprasIntegrationConfig(t)
	require.Equal(t, siat.AmbientePruebas, cfg.CodigoAmbiente)
	require.Equal(t, cfg.CodigoAmbiente, original.CodigoAmbiente)
	require.Equal(t, cfg.Nit, original.Nit)
	require.Equal(t, cfg.CodigoSistema, original.CodigoSistema)
	output, err := os.MkdirTemp(".siat-traces", "reintento-")
	require.NoError(t, err)
	cfg.HTTPClient = &http.Client{Timeout: 45 * time.Second, Transport: &xmlTraceTransport{
		t: t, base: http.DefaultTransport, outputDir: output,
	}}
	if os.Getenv("SIAT_COMPRA_TRAZA") == "0" {
		cfg.HTTPClient.Transport = http.DefaultTransport
	}
	t.Logf("Trazas de reintento: %s; se reutilizan los códigos y la factura existentes", output)
	name := os.Getenv("SIAT_COMPRA_NOMBRE")
	controlValue := os.Getenv("SIAT_COMPRA_CONTROL")
	controlNulo := os.Getenv("SIAT_COMPRA_CONTROL_NULO") == "1"
	decimales := os.Getenv("SIAT_COMPRA_DECIMALES") == "2"
	fechaSinZona := os.Getenv("SIAT_COMPRA_FECHA_SIN_ZONA") == "1"
	if name != "" || controlValue != "" || controlNulo || decimales || fechaSinZona {
		require.Equal(t, 1, original.CantidadFacturas, "el diagnóstico por nombre admite una sola compra")
		compressed, err := base64.StdEncoding.DecodeString(original.Archivo)
		require.NoError(t, err)
		gz, err := gzip.NewReader(bytes.NewReader(compressed))
		require.NoError(t, err)
		defer gz.Close()
		tr := tar.NewReader(gz)
		header, err := tr.Next()
		require.NoError(t, err)
		require.Equal(t, byte(tar.TypeReg), header.Typeflag)
		require.LessOrEqual(t, header.Size, int64(8<<20))
		xmlData, err := io.ReadAll(io.LimitReader(tr, 8<<20))
		require.NoError(t, err)
		_, err = tr.Next()
		require.ErrorIs(t, err, io.EOF, "el TAR debe contener un solo XML")
		if name == "" {
			name = header.Name
		}
		if controlValue != "" {
			doc := etree.NewDocument()
			require.NoError(t, doc.ReadFromBytes(xmlData))
			control := doc.FindElement("registroCompra/codigoControl")
			require.NotNil(t, control)
			if control.Text() != controlValue {
				require.Empty(t, control.Text(), "no se reemplaza un código de control existente")
				control.SetText(controlValue)
			}
			xmlData, err = doc.WriteToBytes()
			require.NoError(t, err)
			t.Logf("Variante de diagnóstico: codigoControl=%q", controlValue)
		} else if controlNulo {
			doc := etree.NewDocument()
			require.NoError(t, doc.ReadFromBytes(xmlData))
			control := doc.FindElement("registroCompra/codigoControl")
			require.NotNil(t, control)
			require.Empty(t, control.Text(), "no se reemplaza un código de control existente")
			doc.Root().CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
			control.CreateAttr("xsi:nil", "true")
			xmlData, err = doc.WriteToBytes()
			require.NoError(t, err)
			t.Log("Variante de diagnóstico: codigoControl con xsi:nil=true")
		}
		if decimales {
			doc := etree.NewDocument()
			require.NoError(t, doc.ReadFromBytes(xmlData))
			for _, field := range []string{
				"montoTotalCompra", "importeIce", "importeIehd", "importeIpj", "tasas",
				"otroNoSujetoCredito", "importesExentos", "importeTasaCero", "subTotal",
				"descuento", "montoGiftCard", "montoTotalSujetoIva", "creditoFiscal",
			} {
				element := doc.FindElement("registroCompra/" + field)
				require.NotNil(t, element, "falta el campo monetario %s", field)
				amount, err := decimal.NewFromString(element.Text())
				require.NoError(t, err, "campo monetario inválido: %s", field)
				element.SetText(amount.StringFixed(2))
			}
			xmlData, err = doc.WriteToBytes()
			require.NoError(t, err)
			t.Log("Variante de diagnóstico: importes con dos decimales")
		}
		if fechaSinZona {
			doc := etree.NewDocument()
			require.NoError(t, doc.ReadFromBytes(xmlData))
			fechaEmision := doc.FindElement("registroCompra/fechaEmision")
			require.NotNil(t, fechaEmision)
			fecha, err := time.Parse(time.RFC3339Nano, fechaEmision.Text())
			require.NoError(t, err, "fechaEmision debe incluir una fecha RFC3339 válida")
			fechaEmision.SetText(fecha.Format("2006-01-02T15:04:05"))
			xmlData, err = doc.WriteToBytes()
			require.NoError(t, err)
			t.Log("Variante de diagnóstico: fechaEmision sin zona ni fracciones de segundo")
		}
		archive := utils.NewTarGzBase64Writer()
		require.NoError(t, archive.Add(name, xmlData))
		require.NoError(t, archive.Close())
		original.Archivo, original.HashArchivo = archive.Encoded(), archive.Hash()
		t.Logf("Nombre de entrada TAR: %s -> %s", header.Name, name)
	}
	client, err := siat.New(cfg)
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(t.Context(), 3*time.Minute)
	defer cancel()
	resp, err := client.RecepcionCompras().RecepcionPaqueteCompras(ctx,
		models.NewRecepcionPaqueteComprasBuilder().WithCodigoSucursal(original.CodigoSucursal).
			WithCodigoPuntoVenta(original.CodigoPuntoVenta).WithCuis(original.Cuis).WithCufd(original.Cufd).
			WithArchivo(original.Archivo).WithHashArchivo(original.HashArchivo).
			WithCantidadFacturas(original.CantidadFacturas).WithGestion(original.Gestion).WithPeriodo(original.Periodo).
			WithFechaEnvio(time.Now().UTC()).Build())
	require.NoError(t, err)
	content, err := resp.GetContent()
	require.NoError(t, err)
	result := content.RespuestaRecepcion
	t.Logf("Recepción de reintento: %+v", result)
	require.True(t, result.Transaccion, "%+v", result)
	require.NotEmpty(t, result.CodigoRecepcion)
	validarCompraRegistrada(t, ctx, client, original.Cuis, original.Cufd,
		original.CodigoSucursal, original.CodigoPuntoVenta, result.CodigoRecepcion)
}
