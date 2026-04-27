package services_test

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"

	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// TestSiatCompraVentaService_RecepcionAnexos verifica el envío de anexos técnicos para facturas específicas.
// Valida que el builder construya correctamente la solicitud con todos los campos obligatorios
// y que el servicio procese la respuesta del SIAT usando el mapeo XML estandarizado.
func TestSiatCompraVentaService_RecepcionAnexos(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	req := models.CompraVenta().NewRecepcionAnexosBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithCuf("D5340CCDF031F0CFDB...").
		AddAnexos(models.CompraVenta().NewVentaAnexoBuilder().
			WithCodigo("1").
			WithCodigoProducto("86111").
			Build()).
		Build()

	resp, err := serviceCompraVenta.RecepcionAnexos(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion Anexos: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_VerificacionEstadoFactura valida la consulta del estado actual de una factura en el SIAT.
// Verifica que se retorne la información de recepción y estado (Válida, Anulada, etc.) correctamente.
func TestSiatCompraVentaService_VerificacionEstadoFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)

	req := models.CompraVenta().NewVerificacionEstadoFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithCuf("D5340CCDF031F0CFDBF...").
		Build()

	resp, err := serviceCompraVenta.VerificacionEstadoFactura(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Verificacion Estado: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_VerificarComunicacion valida la conectividad básica con el
// Servicio de Facturación Compra Venta del SIAT.
func TestSiatCompraVentaService_VerificarComunicacion(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.CompraVenta()

	req := models.CompraVenta().NewVerificarComunicacionBuilder().Build()
	resp, err := service.VerificarComunicacion(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		assert.True(t, resp.Body.Content.Return.Transaccion)
		log.Printf("Respuesta Comunicacion: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_RecepcionMasivaFactura prueba el flujo de envío de múltiples facturas
// bajo la modalidad de emisión masiva. Asegura que el CUF generado sea consistente con el tipo de emisión.
func TestSiatCompraVentaService_RecepcionMasivaFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente := siat.AmbientePruebas
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error calculando CUIS: %v", err)
	}

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error calculando CUFD: %v", err)
	}

	// Construir paquete de 5 facturas para recepción masiva
	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)
	fechaEmision := time.Now()
	codigoPuntoVenta := 0
	for i := 1; i <= 5; i++ {
		nombreRazonSocial := "JUAN PEREZ"
		// Para masiva debe ser emisión Masiva (3)
		cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, siat.EmisionMasiva, 1, 1, i, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
		log.Printf("CUF #%d: %s", i, cuf)
		cabecera := invoices.NewCompraVentaCabeceraBuilder().
			WithNitEmisor(nit).
			WithRazonSocialEmisor("Ronaldo Rua").
			WithMunicipio("Tarija").
			WithNumeroFactura(int64(i)).
			WithCuf(cuf).
			WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
			WithCodigoSucursal(0).
			WithDireccion("ESQUINA AVENIDA LA PAZ").
			WithCodigoPuntoVenta(&codigoPuntoVenta).
			WithFechaEmision(fechaEmision).
			WithNombreRazonSocial(&nombreRazonSocial).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5115889").
			WithCodigoCliente(strconv.Itoa(i)).
			WithCodigoMetodoPago(1).
			WithMontoTotal(100).
			WithMontoTotalSujetoIva(100).
			WithCodigoMoneda(1).
			WithTipoCambio(1).
			WithMontoTotalMoneda(100).
			WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
			WithUsuario("usuario").
			WithCodigoDocumentoSector(1).
			Build()

		detalle := invoices.NewCompraVentaDetalleBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("abc123").
			WithDescripcion("GASA").
			WithCantidad(1).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(100).
			Build()
		factura := invoices.NewCompraVentaBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		xmlData, _ := xml.Marshal(factura)
		signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")

		hdr := &tar.Header{
			Name: fmt.Sprintf("factura_%d.xml", i),
			Mode: 0600,
			Size: int64(len(signedXML)),
		}
		tw.WriteHeader(hdr)
		tw.Write(signedXML)
	}
	tw.Close()

	// Comprimir el TAR con Gzip y preparar para SIAT
	hashString, encodedArchivo, err := utils.CompressAndHash(tarBuf.Bytes())
	if err != nil {
		t.Fatalf("error preparando paquete masivo: %v", err)
	}

	req := models.CompraVenta().NewRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		WithCantidadFacturas(5).
		Build()

	resp, err := serviceCompraVenta.RecepcionMasivaFactura(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion Masiva: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_ValidacionRecepcionMasivaFactura verifica el estado de procesamiento
// de un paquete de facturas enviado previamente mediante RecepcionMasivaFactura.
func TestSiatCompraVentaService_ValidacionRecepcionMasivaFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)

	req := models.CompraVenta().NewValidacionRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithCodigoRecepcion("755d4aab-1ce6-11f1-8c52-99bc8e8492c6").
		Build()

	resp, err := serviceCompraVenta.ValidacionRecepcionMasivaFactura(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Validacion Masiva: %+v", resp.Body.Content)
	}
}

// TestSiatCompraVentaService_RecepcionPaqueteFactura implementa una prueba compleja de empaquetado.
// Genera 500 facturas en memoria, las firma, crea un archivo TAR.GZ en un buffer (sin archivos físicos)
// y lo envía al SIAT codificado en Base64 para su validación.
func TestSiatCompraVentaService_RecepcionPaqueteFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad := siat.ModalidadElectronica
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	// 1. Obtener CUIS y CUFD para Punto Venta 1
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoPuntoVenta(1).
		WithCodigoSucursal(0).
		Build()
	cuisResp, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo
	// cufdReq := models.Codigos().NewCufdBuilder().
	// 	WithCodigoAmbiente(codAmbiente).
	// 	WithCodigoModalidad(codModalidad).
	// 	WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
	// 	WithCodigoSucursal(0).
	// 	WithCodigoPuntoVenta(1).
	// 	WithNit(nit).
	// 	WithCuis(cuis).
	// 	Build()

	// cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	codigoEvento := int64(9670864)
	cufdEvento := "FBQT5CwqE4TERBI5RjlGOEM3MDc=QjlsMmVLY0VhVUMzcxQUFCRDA1Q0"
	cufdControlEvento := "0046A97840CAF74"

	// 3. Construir Paquete de 500 Facturas
	// La fecha de emisión debe estar DENTRO del rango de contingencia del evento
	// Fecha inicio 2026-04-27T11:26:03.171, Fecha fin 2026-04-27T11:28:03.171
	fechaEmision, _ := time.Parse("2006-01-02T15:04:05.000", "2026-04-27T11:27:00.000")
	archivosMap := make(map[string][]byte)
	now := time.Now()
	for i := 1; i <= 500; i++ {
		// Generar CUF con el CodigoControl del CUFD del evento
		cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, siat.EmisionOffline, 1, 1, i, 1, cufdControlEvento)

		factura := invoices.NewCompraVentaBuilder().
			WithModalidad(codModalidad).
			WithCabecera(invoices.NewCompraVentaCabeceraBuilder().
				WithNitEmisor(nit).
				WithRazonSocialEmisor("Ronaldo Rua").
				WithMunicipio("La Paz").
				WithNumeroFactura(int64(i)).
				WithCuf(cuf).
				WithCufd(cufdEvento). // Las facturas contingentes van con el CUFD del evento
				WithCodigoSucursal(0).
				WithCodigoPuntoVenta(utils.IntPtr(1)).
				WithDireccion("Calle 1").
				WithFechaEmision(fechaEmision).
				WithCodigoTipoDocumentoIdentidad(1).
				WithNumeroDocumento("1234567").
				WithCodigoCliente("CLI01").
				WithCodigoMetodoPago(1).
				WithMontoTotal(100.0).
				WithMontoTotalSujetoIva(100.0). // Requerido para evitar error 1015
				WithCodigoMoneda(1).
				WithTipoCambio(1.0).
				WithMontoTotalMoneda(100.0).
				WithLeyenda("Leyenda").
				WithUsuario("user").
				Build()).
			AddDetalle(invoices.NewCompraVentaDetalleBuilder().
				WithActividadEconomica("477300").
				WithCodigoProductoSin(622539).
				WithCodigoProducto("abc123").
				WithDescripcion("GASA").
				WithCantidad(1.0).
				WithUnidadMedida(1).
				WithPrecioUnitario(100.0).
				WithSubTotal(100.0).
				Build()).
			Build()

		xmlData, _ := xml.Marshal(factura)
		signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")
		archivosMap[fmt.Sprintf("factura_%d.xml", i)] = signedXML
	}

	tarGz, _ := utils.CreateTarGz(archivosMap)
	hashArchivo := utils.SHA256Hex(tarGz)
	encodedArchivo := base64.StdEncoding.EncodeToString(tarGz)

	// 4. Preparar solicitud RecepcionPaqueteFactura
	// Usamos fechaEnvio = la fecha actual en formato UTC extendido sin zona horaria
	reqPaquete := models.CompraVenta().NewRecepcionPaqueteFacturaBuilder().
		WithCodigoAmbiente(codAmbiente). // 2
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis).
		WithCufd(cufdEvento).                   // Enviar el CUFD del evento al paquete
		WithCodigoDocumentoSector(1).           // Sector 1: Compra Venta
		WithCodigoEmision(siat.EmisionOffline). // 2
		WithCodigoEvento(codigoEvento).
		WithCodigoPuntoVenta(1).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(1). // 1: Con Crédito Fiscal
		WithArchivo(encodedArchivo).
		WithFechaEnvio(now).
		WithHashArchivo(hashArchivo).
		WithCantidadFacturas(500).
		Build()

	// 5. Ejecutar
	respPaquete, err := serviceCompraVenta.RecepcionPaqueteFactura(context.Background(), config, reqPaquete)
	if err != nil {
		t.Fatalf("error en RecepcionPaqueteFactura: %v", err)
	}

	assert.NotNil(t, respPaquete)
	log.Printf("Respuesta RecepcionPaqueteFactura: %+v", respPaquete.Body.Content)
}

// TestSiatCompraVentaService_ValidacionRecepcionPaqueteFactura consulta el resultado de la validación
// de un paquete de facturas (TAR.GZ) enviado al SIAT.
func TestSiatCompraVentaService_ValidacionRecepcionPaqueteFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(1).
		WithNit(nit).
		Build()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCodigoPuntoVenta(1).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)

	req := models.CompraVenta().NewValidacionRecepcionPaqueteFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(siat.EmisionOffline).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(1).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithCodigoRecepcion("9c081e42-425f-11f1-b837-337cb4b633c2").
		Build()

	resp, err := serviceCompraVenta.ValidacionRecepcionPaqueteFactura(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Validacion Paquete: %+v", resp.Body.Content)
	}
}

func TestSiatCompraVentaService_ReversionAnulacionFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}

	cuf := "ABCD1234"
	resp, err := serviceCompraVenta.ReversionAnulacionFactura(context.Background(), config, models.CompraVenta().
		NewReversionAnulacionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1).
		Build())

	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatCompraVentaService_RecepcionCompraVenta valida el flujo técnico de emisión de una factura individual.
// Proceso: Construcción -> Firmado XML -> Compresión Gzip -> Codificación Base64 -> Envío SOAP.
func TestSiatCompraVentaService_RecepcionCompraVenta(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()

	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}

	serviceCompraVenta := siatClient.CompraVenta()

	fechaEmision := time.Now()
	// 1. Generar CUF
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 1, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0
	// Crear objeto de factura usando el nuevo paquete facturas
	cabecera := invoices.NewCompraVentaCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("ESQUINA AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		WithCodigoDocumentoSector(1).
		Build()

	detalle := invoices.NewCompraVentaDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("abc123").
		WithDescripcion("GASA").
		WithCantidad(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := invoices.NewCompraVentaBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 2. Serializar y Firmar
	xmlData, _ := xml.Marshal(factura)
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error firmando XML: %v", err)
	}

	// 3, 4, 5. Preparar archivo (Gzip + Hash SHA256 + Base64)
	hashString, encodedArchivo, err := utils.CompressAndHash(signedXML)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.CompraVenta().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceCompraVenta.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

// TestSiatCompraVentaService_AnulacionFactura valida el proceso de anulación de una factura emitida.
// Se debe especificar un código de motivo de anulación válido según la paramétrica del SIAT.
// TestSiatCompraVentaService_AnulacionFactura prueba el flujo de anulación de una factura
// proporcionando el motivo de anulación y el CUF correspondiente.
func TestSiatCompraVentaService_AnulacionFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), client)
	serviceCodigos := siatClient.Codigos()
	serviceCompraVenta := siatClient.CompraVenta()

	// Solicitar CUIS y CUFD
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()

	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}

	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}

	// Generar CUF de la factura que supuestamente vamos a anular
	cuf := "ABCD1234"
	// Usar el Builder en lugar de instanciar directamente
	req := models.CompraVenta().NewAnulacionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuf(cuf).
		WithCodigoMotivo(1).
		Build()

	resp, err := serviceCompraVenta.AnulacionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud de anulación: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta Anulación SIAT: %+v", resp.Body.Content)
}
