package invoices

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
)

func TestTelecomunicacionesZFBuilder(t *testing.T) {
	fechaEmision := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	telefono := "2222222"
	nombre := "JUAN PEREZ (ZONA FRANCA)"
	nitConjunto := int64(123456789)
	nSerie := "SN-ZF-123"

	cabecera := NewTelecomunicacionesZFCabeceraBuilder().
		WithNitEmisor(1234567).
		WithRazonSocialEmisor("EMPRESA TELECOM ZF").
		WithMunicipio("IQUIQUE (EXT)").
		WithTelefono(&telefono).
		WithNitConjunto(&nitConjunto).
		WithNumeroFactura(5).
		WithCuf("ABC123ZF").
		WithCufd("XYZ789ZF").
		WithCodigoSucursal(0).
		WithDireccion("PUERTO ZF 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-ZF-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(500.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(500.00).
		WithLeyenda("Venta en Zona Franca").
		WithUsuario("operador_zf").
		Build()

	detalle := NewTelecomunicacionesZFDetalleBuilder().
		WithActividadEconomica("611000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("CONEXION SATELITAL ZF").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(500.00).
		WithSubTotal(500.00).
		WithNumeroSerie(&nSerie).
		Build()

	t.Run("Modalidad Electronica ZF", func(t *testing.T) {
		factura := NewTelecomunicacionesZFBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		internal := models.UnwrapInternalRequest[documents.FacturaTelecomunicacionesZF](factura)
		if internal == nil {
			t.Fatalf("No se pudo extraer el request interno")
		}
		output, err := xml.MarshalIndent(internal, "", "  ")
		if err != nil {
			t.Fatalf("Error al serializar: %v", err)
		}

		xmlStr := string(output)
		if !strings.Contains(xmlStr, "<facturaElectronicaTelecomunicacionZF") {
			t.Errorf("Nodo raíz incorrecto para Electronica ZF")
		}
		if !strings.Contains(xmlStr, "<montoTotalSujetoIva>0</montoTotalSujetoIva>") {
			t.Errorf("montoTotalSujetoIva debe ser 0 para ZF")
		}
	})
}

func TestTelecomunicacionesZFIntegration(t *testing.T) {
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
	serviceTelecom := siatClient.Telecomunicaciones()

	// 1. Obtener CUIS
	cuisReq := models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build()
	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, cuisReq)

	// 2. Obtener CUFD
	cufdReq := models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build()
	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, cufdReq)

	fechaEmision := time.Now()
	// 3. Generar CUF (Sector 49, TipoFactura 2)
	cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 2, 49, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)

	// 4. Construir Factura ZF
	nombre := "JUAN PEREZ (ZF)"
	cabecera := NewTelecomunicacionesZFCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("EMPRESA TELECOM ZF").
		WithMunicipio("La Paz").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("Av. Principal 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("ZF-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(200).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(200).
		WithLeyenda("Venta en Zona Franca").
		WithUsuario("operador_zf").
		Build()

	detalle := NewTelecomunicacionesZFDetalleBuilder().
		WithActividadEconomica("611000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Servicio ZF").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(200).
		WithSubTotal(200).
		Build()

	factura := NewTelecomunicacionesZFBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 5. Serializar, Firmar, Comprimir
	xmlData, _ := xml.Marshal(factura)
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("Error firmando: %v", err)
	}
	hashString, encodedArchivo, _ := utils.CompressAndHash(signedXML)

	// 6. Recepción
	req := models.Telecomunicaciones().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(49).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(2). // 2 para ZF
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceTelecom.RecepcionFactura(context.Background(), config, req)

	if err == nil && resp != nil {
		log.Printf("Respuesta Recepcion ZF: %+v", resp.Body.Content)
	}
}
