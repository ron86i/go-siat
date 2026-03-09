package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/compra_venta"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
)

func main() {
	godotenv.Load(".env")

	// 1. Configuración básica desde el entorno
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	siatURL := os.Getenv("SIAT_URL")
	siatToken := os.Getenv("SIAT_TOKEN")
	siatSistema := os.Getenv("SIAT_CODIGO_SISTEMA")

	if siatURL == "" || siatToken == "" {
		log.Fatalf("Las variables de entorno SIAT_URL y SIAT_TOKEN son obligatorias")
	}

	cfg := config.Config{Token: siatToken}
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}

	// 2. Inicializar servicios
	siatService, err := siat.New(siatURL, client)
	if err != nil {
		log.Fatalf("Error inicializando SIAT: %v", err)
	}
	serviceCodigos := siatService.Codigos
	serviceCompraVenta := siatService.CompraVenta

	// 3. Solicitar CUIS
	fmt.Println("Solicitando CUIS...")
	cuisReq := models.Codigos.NewCuisRequest().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(siatSistema).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(0).
		Build()

	respCuis, err := serviceCodigos.SolicitudCuis(context.Background(), cfg, cuisReq)
	if err != nil {
		log.Fatalf("Error obteniendo CUIS: %v", err)
	}
	cuis := respCuis.Body.Content.RespuestaCuis.Codigo
	fmt.Printf("CUIS: %s", cuis)
	// 4. Solicitar CUFD
	fmt.Println("Solicitando CUFD...")
	cufdReq := models.Codigos.NewCufdRequest().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(siatSistema).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(0).
		WithCuis(cuis).
		Build()

	respCufd, err := serviceCodigos.SolicitudCufd(context.Background(), cfg, cufdReq)
	if err != nil {
		log.Fatalf("Error obteniendo CUFD: %v", err)
	}
	cufd := respCufd.Body.Content.RespuestaCufd.Codigo
	codigoControl := respCufd.Body.Content.RespuestaCufd.CodigoControl
	fmt.Printf("CUFD: %s", cufd)
	// 5. Generar CUF
	fechaEmision := time.Now()
	codEmision := 1

	fmt.Printf("MODO: Emisión %d, Modalidad %d\n", codEmision, codModalidad)

	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, codEmision, 1, 1, 1, 0, codigoControl)
	if err != nil {
		log.Fatalf("Error generando CUF: %v", err)
	}

	// 6. Construir Factura usando Builders
	cabecera := models.CompraVenta.NewCabecera().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("ESQUINA AVENIDA LA PAZ").
		WithCodigoPuntoVenta(0).
		WithFechaEmision(fechaEmision.Format("2006-01-02T15:04:05.000")).
		WithNombreRazonSocial("JUAN PEREZ").
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley 453: Tienes derecho a recibir informacion...").
		WithUsuario("usuario").
		WithCodigoDocumentoSector(1).
		Build()

	detalle := models.CompraVenta.NewDetalle().
		WithActividadEconomica("477300").
		WithCodigoProductoSin("622539").
		WithCodigoProducto("abc123").
		WithDescripcion("GASA").
		WithCantidad(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := models.CompraVenta.NewFacturaCompraVenta().
		WithCabecera(cabecera).
		WithModalidad(codModalidad).
		AddDetalle(detalle).
		Build()

	// 7. Serializar, Firmar, Comprimir
	xmlData, err := xml.Marshal(models.GetInternalRequest[compra_venta.FacturaCompraVenta](factura))
	if err != nil {
		log.Fatalf("Error marshaling factura: %v", err)
	}

	fmt.Printf("Invoice XML:\n%s\n", string(xmlData))

	// Intentar cargar certificados desde varias rutas
	keyPath := "key.pem"
	certPath := "cert.crt"
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		// Probar ruta donde están los del test
		keyPath = "internal/adapter/service/key.pem"
		certPath = "internal/adapter/service/cert.crt"
	}

	var signedXML []byte
	if codModalidad == 1 {
		fmt.Println("MODO: Factura Electrónica (Requiere Firma)")
		fmt.Println("Firmando XML...")
		signedXML, err = utils.SignXML(xmlData, keyPath, certPath)
		if err != nil {
			fmt.Println("--- ERROR CRÍTICO ---")
			fmt.Println("No se encontró firma digital válida o falló el proceso.")
			fmt.Println("Las facturas ELECTRÓNICAS (Modalidad 1) RECHAZAN envios sin firma.")
			fmt.Println("---------------------")
			signedXML = xmlData
		}
	} else {
		fmt.Println("MODO: Factura Computarizada (No requiere firma)")
		fmt.Println("Saltando firma para modalidad Computarizada (Modalidad 2)...")
		signedXML = xmlData
	}

	fmt.Println("Comprimiendo archivo...")
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(signedXML); err != nil {
		log.Fatalf("Error comprimiendo: %v", err)
	}
	zw.Close()
	compressedBytes := buf.Bytes()

	// 8. Hash y Base64
	hash := sha256.Sum256(compressedBytes)
	hashString := hex.EncodeToString(hash[:])

	// SIAT requiere el archivo en Base64.
	encodedArchivo := base64.StdEncoding.EncodeToString(compressedBytes)

	// 9. Construir Solicitud de Recepción usando Builder
	recepcionReq := models.CompraVenta.NewRecepcionFacturaRequest().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(codEmision).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(siatSistema).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	// 10. Enviar al SIAT
	fmt.Println("Enviando a recepción...")
	resp, err := serviceCompraVenta.RecepcionFactura(context.Background(), cfg, recepcionReq)
	if err != nil {
		log.Fatalf("Error en recepción SIAT: %v", err)
	}

	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}
