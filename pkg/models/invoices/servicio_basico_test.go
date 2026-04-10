package invoices_test

import (
	"context"
	"encoding/xml"
	"log"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
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

	cabecera := invoices.NewServicioBasicoCabeceraBuilder().
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

	detalle := invoices.NewServicioBasicoDetalleBuilder().
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
		factura := invoices.NewServicioBasicoBuilder().
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
	tc := setupTestContext(t, siat.ModalidadComputarizada)

	service := tc.Client.ServicioBasico()

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	// 4. Construir Factura
	nombre := "JUAN PEREZ"
	mes := "ABRIL"
	gestion := 2024

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

	// 4. Construir XML
	factura := invoices.NewServicioBasicoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewServicioBasicoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("EMPRESA ELECTRICA S.A.").
			WithMunicipio("La Paz").
			WithNumeroFactura(1).
			WithCuf(tc.GetCuf(t, 13, 1, 1, 1, 0, cufdControl)).
			WithCufd(cufd).
			WithCodigoSucursal(tc.Sucursal).
			WithDireccion("Av. Principal 123").
			WithFechaEmision(time.Now()).
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
			Build()).
		AddDetalle(detalle).
		Build()

	// 5. Serializar, Firmar, Comprimir
	xmlData, _ := xml.Marshal(factura)
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("Error firmando: %v", err)
	}
	hashString, encodedArchivo, _ := utils.CompressAndHash(signedXML)

	// 6. Solicitud de recepción
	req := models.ServicioBasico().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(13).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(tc.Nit).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(time.Now()).
		WithHashArchivo(hashString).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionFactura(context.Background(), tc.Config, req)

	if err == nil && resp != nil {
		log.Printf("Respuesta Recepcion Servicio Basico: %+v", resp.Body.Content)
	}
}
