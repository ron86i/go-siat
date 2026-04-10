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

func TestServicioBasicoZFBuilder(t *testing.T) {
	fechaEmision := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	telefono := "2222222"
	nombre := "JUAN PEREZ (ZF)"
	mes := "ENERO"
	gestion := 2024
	ciudad := "COBIJA"
	zona := "ZONA FRANCA"

	cabecera := NewServicioBasicoZFCabeceraBuilder().
		WithNitEmisor(1234567).
		WithRazonSocialEmisor("EMPRESA ZF").
		WithMunicipio("COBIJA").
		WithTelefono(&telefono).
		WithNumeroFactura(1).
		WithCuf("ABC123CUF_ZF").
		WithCufd("XYZ789CUFD_ZF").
		WithCodigoSucursal(0).
		WithDireccion("AV. ZONA FRANCA 1").
		WithMes(&mes).
		WithGestion(&gestion).
		WithCiudad(&ciudad).
		WithZona(&zona).
		WithNumeroMedidor("MED-ZF-001").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-ZF-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100.00).
		WithLeyenda("Leyenda ZF").
		WithUsuario("operador_zf").
		Build()

	detalle := NewServicioBasicoZFDetalleBuilder().
		WithActividadEconomica("351000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P-ZF-001").
		WithDescripcion("CONSUMO ZF").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(100.00).
		WithSubTotal(100.00).
		Build()

	t.Run("Modalidad Electronica ZF", func(t *testing.T) {
		factura := NewServicioBasicoZFBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		internal := models.UnwrapInternalRequest[documents.FacturaServicioBasicoZF](factura)
		output, _ := xml.Marshal(internal)
		xmlStr := string(output)

		assert.Contains(t, xmlStr, "facturaElectronicaServicioBasicoZf")
		assert.Contains(t, xmlStr, "<montoTotalSujetoIva>0</montoTotalSujetoIva>")
		assert.Contains(t, xmlStr, "<codigoDocumentoSector>40</codigoDocumentoSector>")
	})
}

func TestServicioBasicoZFIntegration(t *testing.T) {
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
	// 3. Generar CUF (Sector 40)
	cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 2, 40, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)

	// 4. Construir Factura
	nombre := "PRODUCTOS ZF S.A."
	mes := "ABRIL"
	gestion := 2024
	cabecera := NewServicioBasicoZFCabeceraBuilder().
		WithNitEmisor(nit).
		WithRazonSocialEmisor("EMPRESA ZF S.A.").
		WithMunicipio("Cobija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCodigoSucursal(0).
		WithDireccion("Zona Franca Cobija").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-ZF-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(200).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(200).
		WithLeyenda("ZONA FRANCA - EXENTO DE IVA").
		WithUsuario("operador").
		WithMes(&mes).
		WithGestion(&gestion).
		WithNumeroMedidor("ZF-9988").
		Build()

	detalle := NewServicioBasicoZFDetalleBuilder().
		WithActividadEconomica("351000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P-ZF-002").
		WithDescripcion("Suministro ZF").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(200).
		WithSubTotal(200).
		Build()

	factura := NewServicioBasicoZFBuilder().
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
		WithCodigoDocumentoSector(40).
		WithCodigoEmision(1).
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(2).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := serviceBasico.RecepcionFactura(context.Background(), config, req)

	if err == nil && resp != nil {
		log.Printf("Respuesta Recepcion Servicio Basico ZF: %+v", resp.Body.Content)
	}
}
