package services_test

import (
	"context"
	"encoding/xml"
	"log"
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

func TestSiatServicioBasicoService_VerificarComunicacion(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}

	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.ServicioBasico()

	req := models.ServicioBasico().NewVerificarComunicacionBuilder().Build()
	resp, err := service.VerificarComunicacion(context.Background(), config, req)

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		assert.True(t, resp.Body.Content.Return.Transaccion)
	}
}

func TestSiatServicioBasicoService_RecepcionFactura(t *testing.T) {
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

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 13, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)

	nombre := "JUAN PEREZ"
	mes, gestion := "ABRIL", 2024
	cabecera := invoices.NewServicioBasicoCabeceraBuilder().
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

	detalle := invoices.NewServicioBasicoDetalleBuilder().
		WithActividadEconomica("351000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Servicio Electrico").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := invoices.NewServicioBasicoBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, _ := xml.Marshal(factura)
	signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")
	hashString, encodedArchivo, _ := utils.CompressAndHash(signedXML)

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

	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepcion Servicio Basico: %+v", resp.Body.Content)
	}
}

func TestSiatServicioBasicoService_VerificacionEstadoFactura(t *testing.T) {
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

	req := models.ServicioBasico().NewVerificacionEstadoFacturaBuilder().
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
		WithCuf("ABC123FAKE").
		Build()

	resp, err := serviceBasico.VerificacionEstadoFactura(context.Background(), config, req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Estado Servicio Basico: %+v", resp.Body.Content)
	}
}

func TestSiatServicioBasicoService_AnulacionFactura(t *testing.T) {
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

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build())

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(codAmbiente).
		WithCodigoModalidad(codModalidad).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build())

	req := models.ServicioBasico().NewAnulacionFacturaBuilder().
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
		WithCuf("ABC123FAKE").
		WithCodigoMotivo(1).
		Build()

	resp, err := serviceBasico.AnulacionFactura(context.Background(), config, req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Anulación: %+v", resp.Body.Content)
	}
}

func TestSiatServicioBasicoService_ReversionAnulacionFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	serviceCodigos := siatClient.Codigos()
	serviceBasico := siatClient.ServicioBasico()

	cuis, _ := serviceCodigos.SolicitudCuis(context.Background(), config, models.Codigos().NewCuisBuilder().
		WithCodigoAmbiente(siat.AmbientePruebas).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		Build())

	cufd, _ := serviceCodigos.SolicitudCufd(context.Background(), config, models.Codigos().NewCufdBuilder().
		WithCodigoAmbiente(siat.AmbientePruebas).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithNit(nit).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		Build())

	req := models.ServicioBasico().NewReversionAnulacionFacturaBuilder().
		WithCodigoAmbiente(siat.AmbientePruebas).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(1).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithCufd(cufd.Body.Content.RespuestaCufd.Codigo).
		WithCuis(cuis.Body.Content.RespuestaCuis.Codigo).
		WithNit(nit).
		WithTipoFacturaDocumento(1).
		WithCuf("ABC123FAKE").
		Build()

	resp, err := serviceBasico.ReversionAnulacionFactura(context.Background(), config, req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Reversión Anulación: %+v", resp.Body.Content)
	}
}

func TestSiatServicioBasicoService_RecepcionMasivaFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.ServicioBasico()

	req := models.ServicioBasico().NewRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(siat.AmbientePruebas).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(1).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithArchivo("ZHVtbXk="). // "dummy" in base64
		WithHashArchivo("dummyhash").
		Build()

	resp, err := service.RecepcionMasivaFactura(context.Background(), config, req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepción Masiva: %+v", resp.Body.Content)
	}
}

func TestSiatServicioBasicoService_ValidacionRecepcionMasivaFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.ServicioBasico()

	req := models.ServicioBasico().NewValidacionRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(siat.AmbientePruebas).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(1).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithCodigoRecepcion("123456").
		Build()

	resp, err := service.ValidacionRecepcionMasivaFactura(context.Background(), config, req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Validación Masiva: %+v", resp.Body.Content)
	}
}

func TestSiatServicioBasicoService_RecepcionPaqueteFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.ServicioBasico()

	req := models.ServicioBasico().NewRecepcionPaqueteFacturaBuilder().
		WithCodigoAmbiente(siat.AmbientePruebas).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(1).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithArchivo("ZHVtbXk=").
		WithHashArchivo("dummyhash").
		WithCantidadFacturas(10).
		WithCodigoEvento(12345).
		Build()

	resp, err := service.RecepcionPaqueteFactura(context.Background(), config, req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Recepción Paquete: %+v", resp.Body.Content)
	}
}

func TestSiatServicioBasicoService_ValidacionRecepcionPaqueteFactura(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	config := siat.Config{Token: os.Getenv("SIAT_TOKEN")}
	siatClient, _ := siat.New(os.Getenv("SIAT_URL"), nil)
	service := siatClient.ServicioBasico()

	req := models.ServicioBasico().NewValidacionRecepcionPaqueteFacturaBuilder().
		WithCodigoAmbiente(siat.AmbientePruebas).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(1).
		WithCodigoModalidad(siat.ModalidadElectronica).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(os.Getenv("SIAT_CODIGO_SISTEMA")).
		WithCodigoSucursal(0).
		WithNit(nit).
		WithCodigoRecepcion("123456").
		Build()

	resp, err := service.ValidacionRecepcionPaqueteFactura(context.Background(), config, req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		log.Printf("Respuesta Validación Paquete: %+v", resp.Body.Content)
	}
}
