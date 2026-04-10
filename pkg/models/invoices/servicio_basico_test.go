package invoices

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestServicioBasicoBuilder(t *testing.T) {
	fechaEmision := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	telefono := "2222222"
	nombre := "JUAN PEREZ"
	mes := "ENERO"
	gestion := 2024
	ciudad := "LA PAZ"
	zona := "CENTRO"

	cabecera := NewServicioBasicoCabeceraBuilder().
		WithNitEmisor(1234567).
		WithRazonSocialEmisor("EMPRESA ELECTRICA").
		WithMunicipio("LA PAZ").
		WithTelefono(&telefono).
		WithNumeroFactura(1).
		WithCuf("ABC123CUF").
		WithCufd("XYZ789CUFD").
		WithCodigoSucursal(0).
		WithDireccion("AV. MARISCAL SANTA CRUZ 123").
		WithMes(&mes).
		WithGestion(&gestion).
		WithCiudad(&ciudad).
		WithZona(&zona).
		WithNumeroMedidor("MED-001").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(150.50).
		WithMontoTotalSujetoIva(150.50).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(150.50).
		WithLeyenda("Leyenda SIAT").
		WithUsuario("operador1").
		Build()

	detalle := NewServicioBasicoDetalleBuilder().
		WithActividadEconomica("351000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("CONSUMO DE ENERGIA").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(150.50).
		WithSubTotal(150.50).
		Build()

	t.Run("Modalidad Electronica", func(t *testing.T) {
		factura := NewServicioBasicoBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		internal := models.UnwrapInternalRequest[documents.FacturaServicioBasico](factura)
		output, _ := xml.Marshal(internal)
		xmlStr := string(output)

		assert.Contains(t, xmlStr, "facturaElectronicaServicioBasico")
		assert.Contains(t, xmlStr, "<mes>ENERO</mes>")
		assert.Contains(t, xmlStr, "<gestion>2024</gestion>")
		assert.Contains(t, xmlStr, "<codigoDocumentoSector>13</codigoDocumentoSector>")
	})
}

func TestServicioBasicoIntegration(t *testing.T) {
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
	serviceBasico := siatClient.ServicioBasico()

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
	// 3. Generar CUF
	cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 13, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)

	// 4. Construir Factura
	nombre := "JUAN PEREZ"
	mes := "ABRIL"
	gestion := 2024
	cabecera := NewServicioBasicoCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("EMPRESA ELECTRICA S.A.").
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
		WithCodigoCliente("CLI-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley 453...").
		WithUsuario("operador").
		WithMes(&mes).
		WithGestion(&gestion).
		WithNumeroMedidor("123456").
		Build()

	detalle := NewServicioBasicoDetalleBuilder().
		WithActividadEconomica("351000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Servicio Electrico").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := NewServicioBasicoBuilder().
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
	req := models.ServicioBasico().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(1).
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
		Build()

	resp, err := serviceBasico.RecepcionFactura(context.Background(), config, req)

	if err == nil && resp != nil {
		log.Printf("Respuesta Recepcion Servicio Basico: %+v", resp.Body.Content)
	}
}
