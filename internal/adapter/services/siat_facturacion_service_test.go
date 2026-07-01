package services_test

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

type FacturacionTestContext struct {
	Config      siat.Config
	Client      *siat.SiatServices
	Nit         int64
	Ambiente    int
	Modalidad   int
	Sistema     string
	Cuis        string
	Cufd        string
	CufdControl string
}

type testXMLSigner struct {
	privateKeyPath string
	certPath       string
}

func (s *testXMLSigner) SignXML(xmlBytes []byte) ([]byte, error) {
	return utils.SignXML(xmlBytes, s.privateKeyPath, s.certPath)
}

func setupFacturacionTestContext(t *testing.T) *FacturacionTestContext {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	codModalidad := siat.ModalidadElectronica // Para la mayoría de las pruebas

	sistema := os.Getenv("SIAT_CODIGO_SISTEMA")

	config := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            nit,
		CodigoSistema:  sistema,
		CodigoAmbiente: codAmbiente,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}},
	}

	siatClient, err := siat.New(config)
	if err != nil {
		t.Fatalf("Error creando SIAT client: %v", err)
	}

	serviceCodigos := siatClient.Codigos()

	// 1. Obtener CUIS
	cuisReq := models.NewCuisBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		Build()

	cuisResp, err := serviceCodigos.SolicitudCuis(context.Background(), cuisReq)
	if err != nil || cuisResp == nil || cuisResp.Body.Content.RespuestaCuis.Codigo == "" {
		t.Fatalf("Error obteniendo CUIS: %v", err)
	}
	cuis := cuisResp.Body.Content.RespuestaCuis.Codigo

	// 2. Obtener CUFD
	cufdReq := models.NewCufdBuilder().
		WithCodigoModalidad(codModalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCuis(cuis).
		Build()

	cufdResp, err := serviceCodigos.SolicitudCufd(context.Background(), cufdReq)
	if err != nil || cufdResp == nil || cufdResp.Body.Content.RespuestaCufd.Codigo == "" {
		t.Fatalf("Error obteniendo CUFD: %v", err)
	}
	cufd := cufdResp.Body.Content.RespuestaCufd.Codigo
	cufdControl := cufdResp.Body.Content.RespuestaCufd.CodigoControl

	return &FacturacionTestContext{
		Config: config,
		Client: siatClient,

		Nit:         nit,
		Ambiente:    codAmbiente,
		Modalidad:   codModalidad,
		Sistema:     sistema,
		Cuis:        cuis,
		Cufd:        cufd,
		CufdControl: cufdControl,
	}
}

func buildDuttyFree(t *testing.T, tc *FacturacionTestContext, numeroFactura int64, fechaEmision time.Time, tipoEmision int) (string, any) {
	cuf, err := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, tipoEmision, 2, 10, numeroFactura, 0, tc.CufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0
	cantidad := 1.0
	precioUnitario := 100.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	cabecera := invoices.NewDuttyFreeCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Empresa Test").
		WithMunicipio("La Paz").
		WithNumeroFactura(numeroFactura).
		WithCuf(cuf).
		WithCufd(tc.Cufd).
		WithCodigoSucursal(0).
		WithDireccion("Calle 1").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Leyenda").
		WithUsuario("user").
		Build()

	detalle := invoices.NewDuttyFreeDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("P01").
		WithDescripcion("Producto").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithSubTotal(subTotalItem).
		Build()

	factura := invoices.NewDuttyFreeBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	return cuf, factura
}

func TestFacturacionService_VerificarComunicacion(t *testing.T) {
	tc := setupFacturacionTestContext(t)

	req := models.NewVerificarComunicacionFacturacion()
	resp, err := tc.Client.Electronica().VerificarComunicacion(context.Background(), req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta VerificarComunicacion: %+v", resp.Body.Content)
		assert.True(t, resp.Body.Content.Return.Transaccion)
	}
}

func TestFacturacionService_RecepcionFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)
	fechaEmision := time.Now()

	_, factura := buildDuttyFree(t, tc, 1, fechaEmision, siat.EmisionOnline)

	reqBuilder := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(10). // Sector 10: Dutty Free
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCufd(tc.Cufd).
		WithCuis(tc.Cuis).
		WithTipoFacturaDocumento(2). // Sin Crédito Fiscal
		WithFechaEnvio(fechaEmision)

	signer := &testXMLSigner{privateKeyPath: "key.pem", certPath: "cert.crt"}
	if err := reqBuilder.WithFactura(factura, signer); err != nil {
		t.Fatalf("Error procesando factura: %v", err)
	}

	req := reqBuilder.Build()

	resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta RecepcionFactura: %+v", resp.Body.Content)
	}
}

func TestFacturacionService_AnulacionFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)

	req := models.NewAnulacionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCuis(tc.Cuis).
		WithCufd(tc.Cufd).
		WithCodigoMotivo(1).
		WithCuf("D5340CCDF031F2596FC03311F6F76AB5334D0A86A626F497FCE6AAF74").
		Build()

	resp, err := tc.Client.Electronica().AnulacionFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta AnulacionFactura: %+v", resp.Body.Content)
	}
}

func TestFacturacionService_ReversionAnulacionFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)

	req := models.NewReversionAnulacionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCuis(tc.Cuis).
		WithCufd(tc.Cufd).
		WithCodigoDocumentoSector(10).
		WithTipoFacturaDocumento(2).
		WithCodigoEmision(siat.EmisionOnline).
		WithCuf("D5340CCDF031F2596FC03311F6F76AB5334D0A86A626F497FCE6AAF74").
		Build()

	resp, err := tc.Client.Electronica().ReversionAnulacionFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta ReversionAnulacionFactura: %+v", resp.Body.Content)
	}
}

func TestFacturacionService_RecepcionPaqueteFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)
	fechaEmision := time.Now()

	_, factura := buildDuttyFree(t, tc, 2, fechaEmision, siat.EmisionOffline)

	reqBuilder := models.NewRecepcionPaqueteFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCuis(tc.Cuis).
		WithCufd(tc.Cufd).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionOffline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2).
		WithFechaEnvio(fechaEmision).
		WithCodigoEvento(0) // Solo si corresponde

	signer := &testXMLSigner{privateKeyPath: "key.pem", certPath: "cert.crt"}
	if err := reqBuilder.WithFacturas([]any{factura}, signer); err != nil {
		t.Fatalf("Error procesando paquete de facturas: %v", err)
	}

	req := reqBuilder.Build()

	resp, err := tc.Client.Electronica().RecepcionPaqueteFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta RecepcionPaqueteFactura: %+v", resp.Body.Content)
	}
}

func TestFacturacionService_ValidacionRecepcionPaqueteFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)

	req := models.NewValidacionRecepcionPaqueteFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCuis(tc.Cuis).
		WithCufd(tc.Cufd).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionOffline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoRecepcion("123ABCD").
		WithTipoFacturaDocumento(2).
		Build()

	resp, err := tc.Client.Electronica().ValidacionRecepcionPaqueteFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta ValidacionRecepcionPaqueteFactura: %+v", resp.Body.Content)
	}
}

func TestFacturacionService_RecepcionMasivaFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)
	fechaEmision := time.Now()

	_, factura := buildDuttyFree(t, tc, 3, fechaEmision, siat.EmisionMasiva)

	reqBuilder := models.NewRecepcionMasivaFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCuis(tc.Cuis).
		WithCufd(tc.Cufd).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2).
		WithFechaEnvio(fechaEmision)

	signer := &testXMLSigner{privateKeyPath: "key.pem", certPath: "cert.crt"}
	if err := reqBuilder.WithFacturas([]any{factura}, signer); err != nil {
		t.Fatalf("Error procesando facturas masivas: %v", err)
	}

	req := reqBuilder.Build()

	resp, err := tc.Client.Electronica().RecepcionMasivaFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta RecepcionMasivaFactura: %+v", resp.Body.Content)
	}
}

func TestFacturacionService_ValidacionRecepcionMasivaFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)

	req := models.NewValidacionRecepcionMasivaFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCuis(tc.Cuis).
		WithCufd(tc.Cufd).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2).
		WithCodigoRecepcion("123ABCD").
		Build()

	resp, err := tc.Client.Electronica().ValidacionRecepcionMasivaFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta ValidacionRecepcionMasivaFactura: %+v", resp.Body.Content)
	}
}

func TestFacturacionService_VerificacionEstadoFactura(t *testing.T) {
	tc := setupFacturacionTestContext(t)

	req := models.NewVerificacionEstadoFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCuis(tc.Cuis).
		WithCufd(tc.Cufd).
		WithCodigoDocumentoSector(10).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithTipoFacturaDocumento(2).
		WithCuf("D5340CCDF0...").
		Build()

	resp, err := tc.Client.Electronica().VerificacionEstadoFactura(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta VerificacionEstadoFactura: %+v", resp.Body.Content)
	}
}
