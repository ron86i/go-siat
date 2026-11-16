package services_test

import (
	"context"
	"encoding/xml"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/ron86i/go-siat/v3/pkg/models"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	siatfacturacion "github.com/ron86i/go-siat/v3/protocols/facturacion"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

type registroCompraE2E struct {
	XMLName             xml.Name  `xml:"registroCompra"`
	Nro                 int       `xml:"nro"`
	NitEmisor           int64     `xml:"nitEmisor"`
	RazonSocialEmisor   string    `xml:"razonSocialEmisor"`
	CodigoAutorizacion  string    `xml:"codigoAutorizacion"`
	NumeroFactura       int64     `xml:"numeroFactura"`
	NumeroDuiDim        string    `xml:"numeroDuiDim"`
	FechaEmision        time.Time `xml:"fechaEmision"`
	MontoTotalCompra    string    `xml:"montoTotalCompra"`
	ImporteIce          string    `xml:"importeIce"`
	ImporteIehd         string    `xml:"importeIehd"`
	ImporteIpj          string    `xml:"importeIpj"`
	Tasas               string    `xml:"tasas"`
	OtroNoSujetoCredito string    `xml:"otroNoSujetoCredito"`
	ImportesExentos     string    `xml:"importesExentos"`
	ImporteTasaCero     string    `xml:"importeTasaCero"`
	SubTotal            string    `xml:"subTotal"`
	Descuento           string    `xml:"descuento"`
	MontoGiftCard       string    `xml:"montoGiftCard"`
	MontoTotalSujetoIva string    `xml:"montoTotalSujetoIva"`
	CreditoFiscal       string    `xml:"creditoFiscal"`
	TipoCompra          int       `xml:"tipoCompra"`
	CodigoControl       string    `xml:"codigoControl"`
}

func TestRegistroCompraE2EXMLUsesOfficialSchemaSequence(t *testing.T) {
	registro := registroCompraE2E{
		Nro:                 1,
		NitEmisor:           123456789,
		RazonSocialEmisor:   "Proveedor de prueba",
		CodigoAutorizacion:  "CUF-PRUEBA",
		NumeroFactura:       123,
		NumeroDuiDim:        "0",
		FechaEmision:        time.Date(2026, time.October, 16, 12, 0, 0, 0, time.FixedZone("America/La_Paz", -4*60*60)),
		MontoTotalCompra:    "100.01",
		ImporteIce:          "0",
		ImporteIehd:         "0",
		ImporteIpj:          "0",
		Tasas:               "0",
		OtroNoSujetoCredito: "0",
		ImportesExentos:     "0",
		ImporteTasaCero:     "0",
		SubTotal:            "100.01",
		Descuento:           "0",
		MontoGiftCard:       "0",
		MontoTotalSujetoIva: "88.50",
		CreditoFiscal:       "11.51",
		TipoCompra:          1,
		CodigoControl:       "0",
	}

	data, err := xml.Marshal(registro)
	require.NoError(t, err)

	campos := []string{
		"nro", "nitEmisor", "razonSocialEmisor", "codigoAutorizacion",
		"numeroFactura", "numeroDuiDim", "fechaEmision", "montoTotalCompra",
		"importeIce", "importeIehd", "importeIpj", "tasas",
		"otroNoSujetoCredito", "importesExentos", "importeTasaCero", "subTotal",
		"descuento", "montoGiftCard", "montoTotalSujetoIva", "creditoFiscal",
		"tipoCompra", "codigoControl",
	}
	contenido := string(data)
	posicion := 0
	for _, campo := range campos {
		indice := strings.Index(contenido[posicion:], "<"+campo+">")
		require.NotEqual(t, -1, indice, "falta el campo <%s> o no respeta la secuencia", campo)
		posicion += indice + len(campo) + 2
	}
}

// TestSiatRecepcionCompras_FlujoCompleto emite una factura al propio NIT y la
// declara mediante RecepcionPaqueteCompras. Sólo admite ambiente 2.
// SIAT_RUN=1 habilita las operaciones reales; SIAT_COMPRA_RAZON_SOCIAL debe
// contener la razón social registrada para el NIT usado en la prueba. Los
// códigos y datos de la compra se obtienen en esta ejecución. Exige estado
// final 5154. La factura queda vigente.
func TestSiatRecepcionCompras_FlujoCompleto(t *testing.T) {
	_ = godotenv.Load(".env")
	if os.Getenv("SIAT_RUN") != "1" {
		t.Skip("defina SIAT_RUN=1 para emitir y declarar una compra de pruebas")
	}
	razonSocial := strings.TrimSpace(os.Getenv("SIAT_COMPRA_RAZON_SOCIAL"))
	require.NotEmpty(t, razonSocial, "defina SIAT_COMPRA_RAZON_SOCIAL con la razón social registrada en SIAT")
	cfg := siatRecepcionComprasIntegrationConfig(t)
	require.Equal(t, siat.AmbientePruebas, cfg.CodigoAmbiente)
	cfg.HTTPClient.Timeout = 45 * time.Second
	modalidad := siat.ModalidadElectronica
	if value := os.Getenv("SIAT_CODIGO_MODALIDAD"); value != "" {
		var err error
		modalidad, err = strconv.Atoi(value)
		require.NoError(t, err)
	}
	require.Contains(t, []int{siat.ModalidadElectronica, siat.ModalidadComputarizada}, modalidad)
	if modalidad == siat.ModalidadElectronica {
		path := os.Getenv("SIAT_P12_PATH")
		if path == "" {
			path = "key.p12"
		}
		cfg.CredentialSign = siat.NewP12Credential(path, os.Getenv("SIAT_P12_PASSWORD"))
	}
	client, err := siat.New(cfg)
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(t.Context(), 4*time.Minute)
	defer cancel()
	// Una sola pareja CUIS/CUFD, sucursal y punto de venta 0 en todo el flujo.
	cuisResp, err := client.Codigos().SolicitudCuis(ctx, models.NewCuisBuilder().
		WithCodigoSucursal(0).WithCodigoPuntoVenta(0).WithCodigoModalidad(modalidad).Build())
	require.NoError(t, err)
	cuisContent, err := cuisResp.GetContent()
	require.NoError(t, err)
	cuis := cuisContent.RespuestaCuis.Codigo
	require.NotEmpty(t, cuis, "%+v", cuisContent.RespuestaCuis.MensajesList)
	cufdResp, err := client.Codigos().SolicitudCufd(ctx, models.NewCufdBuilder().
		WithCodigoSucursal(0).WithCodigoPuntoVenta(0).WithCodigoModalidad(modalidad).WithCuis(cuis).Build())
	require.NoError(t, err)
	cufdContent, err := cufdResp.GetContent()
	require.NoError(t, err)
	cufd := cufdContent.RespuestaCufd
	require.NotEmpty(t, cufd.Codigo)
	require.NotEmpty(t, cufd.CodigoControl)
	productsResp, err := client.Sincronizacion().SincronizarListaProductosServicios(ctx,
		models.NewSincronizarListaProductosServiciosBuilder().WithCodigoSucursal(0).
			WithCodigoPuntoVenta(0).WithCuis(cuis).Build())
	require.NoError(t, err)
	products, err := productsResp.GetContent()
	require.NoError(t, err)
	require.True(t, products.RespuestaListaProductos.Transaccion, "%+v", products.RespuestaListaProductos)
	require.NotEmpty(t, products.RespuestaListaProductos.ListaCodigos)
	product := products.RespuestaListaProductos.ListaCodigos[0]
	fecha := time.Now().In(time.FixedZone("America/La_Paz", -4*60*60)).Truncate(time.Millisecond)
	numero := fecha.UnixMilli() % 1000000000
	cuf, err := utils.GenerarCUF(cfg.Nit, fecha, 0, modalidad, siat.EmisionOnline, 1, 1, numero, 0, cufd.CodigoControl)
	require.NoError(t, err)
	nit := strconv.FormatInt(cfg.Nit, 10)
	cabecera := compraventa.Cabecera{
		NitEmisor:                    cfg.Nit,
		RazonSocialEmisor:            razonSocial,
		Municipio:                    "LA PAZ",
		Direccion:                    "DOMICILIO DE PRUEBAS",
		NumeroFactura:                numero,
		Cuf:                          cuf,
		Cufd:                         cufd.Codigo,
		CodigoSucursal:               0,
		CodigoPuntoVenta:             new(0),
		FechaEmision:                 fecha,
		NombreRazonSocial:            &razonSocial,
		CodigoTipoDocumentoIdentidad: 5,
		NumeroDocumento:              nit,
		Complemento:                  "0",
		CodigoCliente:                nit,
		CodigoMetodoPago:             1,
		SubTotalNeto:                 88.50,
		TotalPrecioNeto:              88.50,
		Iva:                          11.51,
		MontoAPagar:                  100.01,
		CodigoInterno:                cufd.CodigoInterno,
		CodigoMoneda:                 1,
		TipoCambio:                   1,
		MontoTotalMoneda:             100.01,
		Leyenda:                      "FACTURA DE PRUEBA DE INTEGRACION",
		Usuario:                      "integracion",
		CodigoDocumentoSector:        1,
	}
	detalles := []compraventa.Detalle{
		{
			ActividadEconomica: product.CodigoActividad,
			CodigoProductoSin:  product.CodigoProducto,
			CodigoProducto:     "PRUEBA-RCV",
			Descripcion:        product.DescripcionProducto,
			Cantidad:           1,
			UnidadMedida:       58,
			PrecioNetoUnitario: 88.50,
			SubTotal:           88.50,
		},
	}
	factura, err := compraventa.New(cabecera, detalles)
	require.NoError(t, err)
	emisionClient, err := siatfacturacion.NewClient(siatfacturacion.Config{
		BaseURL:        cfg.BaseURL,
		Token:          cfg.Token,
		Nit:            cfg.Nit,
		CodigoSistema:  cfg.CodigoSistema,
		CodigoAmbiente: cfg.CodigoAmbiente,
		HTTPClient:     cfg.HTTPClient,
	})
	require.NoError(t, err)
	var signer siatfacturacion.XMLSigner
	if modalidad == siat.ModalidadElectronica {
		signer = cfg.CredentialSign
	}
	emision := siatfacturacion.SolicitudRecepcionFactura{
		CodigoDocumentoSector: 1,
		CodigoEmision:         siat.EmisionOnline,
		CodigoModalidad:       modalidad,
		CodigoPuntoVenta:      new(0),
		CodigoSucursal:        0,
		Cufd:                  cufd.Codigo,
		Cuis:                  cuis,
		TipoFacturaDocumento:  1,
		FechaEnvio:            fecha,
	}
	preparada, err := emisionClient.PrepararFactura(emision, factura, signer)
	require.NoError(t, err)
	emitida, err := emisionClient.RecepcionFactura(ctx, preparada)
	require.NoError(t, err)
	require.NotNil(t, emitida.Transaccion, "%+v", emitida)
	require.True(t, *emitida.Transaccion, "%+v", emitida)
	require.Equal(t, 908, emitida.CodigoEstado, "%+v", emitida)
	t.Logf("Factura emitida: número=%d CUF=%s; emisor y comprador usan el mismo NIT", numero, cuf)
	// Este escenario declara la compra manualmente. ConsultaCompras y
	// ConfirmacionCompras son otro flujo: sólo confirman compras ya reportadas.
	registro := registroCompraE2E{
		Nro:                 1,
		NitEmisor:           cfg.Nit,
		RazonSocialEmisor:   razonSocial,
		CodigoAutorizacion:  cuf,
		NumeroFactura:       numero,
		NumeroDuiDim:        "0",
		FechaEmision:        fecha,
		MontoTotalCompra:    "100.01",
		ImporteIce:          "0",
		ImporteIehd:         "0",
		ImporteIpj:          "0",
		Tasas:               "0",
		OtroNoSujetoCredito: "0",
		ImportesExentos:     "0",
		ImporteTasaCero:     "0",
		SubTotal:            "100.01",
		Descuento:           "0",
		MontoGiftCard:       "0",
		MontoTotalSujetoIva: "88.50",
		CreditoFiscal:       "11.51",
		TipoCompra:          1,
		CodigoControl:       "0",
	}
	xmlData, err := xml.Marshal(registro)
	require.NoError(t, err)
	xmlData = append([]byte(xml.Header), xmlData...)
	t.Logf("XML literal registroCompra enviado:\n%s", xmlData)
	archive := utils.NewTarGzBase64Writer()
	require.NoError(t, archive.Add("registroCompra_1.xml", xmlData))
	require.NoError(t, archive.Close())
	packageResp, err := client.RecepcionCompras().RecepcionPaqueteCompras(ctx, models.NewRecepcionPaqueteComprasBuilder().
		WithCodigoSucursal(0).WithCodigoPuntoVenta(0).WithCuis(cuis).WithCufd(cufd.Codigo).
		WithArchivo(archive.Encoded()).WithHashArchivo(archive.Hash()).WithCantidadFacturas(1).
		WithGestion(fecha.Year()).WithPeriodo(int(fecha.Month())).
		WithFechaEnvio(time.Now().UTC()).Build())
	require.NoError(t, err)
	packageContent, err := packageResp.GetContent()
	require.NoError(t, err)
	result := packageContent.RespuestaRecepcion
	t.Logf("Recepción de paquete de compra: estado=%d recepción=%s descripción=%s transacción=%t mensajes=%+v", result.CodigoEstado, result.CodigoRecepcion, result.CodigoDescripcion, result.Transaccion, result.MensajesList)
	require.True(t, result.Transaccion, "%+v", result)
	require.Contains(t, []int{5152, 5154}, result.CodigoEstado, "%+v", result)
	require.NotEmpty(t, result.CodigoRecepcion, "la recepción no devolvió código de recepción")
	validarCompraRegistrada(t, ctx, client, cuis, cufd.Codigo, 0, 0, result.CodigoRecepcion)
}

func validarCompraRegistrada(t *testing.T, ctx context.Context, client *siat.SiatServices, cuis, cufd string, sucursal, puntoVenta int, recepcion string) {
	t.Helper()
	for intento := range 30 {
		validation, err := client.RecepcionCompras().ValidacionRecepcionPaqueteCompras(ctx,
			models.NewValidacionRecepcionPaqueteComprasBuilder().WithCodigoSucursal(sucursal).
				WithCodigoPuntoVenta(puntoVenta).WithCuis(cuis).WithCufd(cufd).
				WithCodigoRecepcion(recepcion).Build())
		require.NoError(t, err)
		content, err := validation.GetContent()
		require.NoError(t, err)
		final := content.RespuestaRecepcion
		t.Logf("Validación compra (%d): estado=%d descripción=%s transacción=%t mensajes=%+v", intento+1, final.CodigoEstado, final.CodigoDescripcion, final.Transaccion, final.MensajesList)
		if final.CodigoEstado != 5152 && final.CodigoEstado != 901 {
			require.True(t, final.Transaccion, "%+v", final)
			require.Equal(t, 5154, final.CodigoEstado, "se requiere una compra procesada, no sólo la recepción del archivo: %+v", final)
			return
		}
		select {
		case <-ctx.Done():
			t.Fatal(ctx.Err())
		case <-time.After(2 * time.Second):
		}
	}
	t.Fatalf("SIAT no terminó de procesar la compra; código de recepción: %s", recepcion)
}
