package invoices_test

import (
	"context"
	"encoding/xml"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
)

func TestTelecomunicacionesBuilder(t *testing.T) {
	fechaEmision := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	telefono := "2222222"
	nombre := "JUAN PEREZ"
	nitConjunto := int64(123456789)
	nSerie := "SN123456"
	nImei := "IMEI987654"

	cabecera := invoices.NewTelecomunicacionesCabeceraBuilder().
		WithNitEmisor(1234567).
		WithRazonSocialEmisor("EMPRESA TELECOM").
		WithMunicipio("LA PAZ").
		WithTelefono(&telefono).
		WithNitConjunto(&nitConjunto).
		WithNumeroFactura(1).
		WithCuf("ABC123CUF").
		WithCufd("XYZ789CUFD").
		WithCodigoSucursal(0).
		WithDireccion("AV. PRINCIPAL 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100.50).
		WithMontoTotalSujetoIva(100.50).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100.50).
		WithLeyenda("Leyenda SIAT").
		WithUsuario("operador1").
		Build()

	detalle := invoices.NewTelecomunicacionesDetalleBuilder().
		WithActividadEconomica("611000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("SERVICIO DE INTERNET").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(100.50).
		WithSubTotal(100.50).
		WithNumeroSerie(&nSerie).
		WithNumeroImei(&nImei).
		Build()

	t.Run("Modalidad Electronica", func(t *testing.T) {
		factura := invoices.NewTelecomunicacionesBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		internal := models.UnwrapInternalRequest[documents.FacturaTelecomunicaciones](factura)
		if internal == nil {
			t.Fatalf("No se pudo extraer el request interno")
		}
		output, err := xml.MarshalIndent(internal, "", "  ")
		if err != nil {
			t.Fatalf("Error al serializar: %v", err)
		}

		xmlStr := string(output)
		if !strings.Contains(xmlStr, "<facturaElectronicaTelecomunicacion") {
			t.Errorf("Nodo raíz incorrecto para Electronica")
		}
		if !strings.Contains(xmlStr, "<nitConjunto>123456789</nitConjunto>") {
			t.Errorf("nitConjunto no encontrado")
		}
		if !strings.Contains(xmlStr, "<numeroSerie>SN123456</numeroSerie>") {
			t.Errorf("numeroSerie no encontrado")
		}
		if !strings.Contains(xmlStr, "<numeroImei>IMEI987654</numeroImei>") {
			t.Errorf("numeroImei no encontrado")
		}
		if !strings.Contains(xmlStr, "<codigoDocumentoSector>22</codigoDocumentoSector>") {
			t.Errorf("codigoDocumentoSector incorrecto")
		}
	})

	t.Run("Modalidad Computarizada", func(t *testing.T) {
		factura := invoices.NewTelecomunicacionesBuilder().
			WithModalidad(siat.ModalidadComputarizada).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		internal := models.UnwrapInternalRequest[documents.FacturaTelecomunicaciones](factura)
		if internal == nil {
			t.Fatalf("No se pudo extraer el request interno")
		}
		output, err := xml.MarshalIndent(internal, "", "  ")
		if err != nil {
			t.Fatalf("Error al serializar: %v", err)
		}

		xmlStr := string(output)
		if !strings.Contains(xmlStr, "<facturaComputarizadaTelecomunicacion") {
			t.Errorf("Nodo raíz incorrecto para Computarizada")
		}
	})
}

func TestTelecomunicacionesIntegration(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)

	service := tc.Client.Telecomunicaciones()

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	// 3. Generar CUF (Sector 22)
	cuf := tc.GetCuf(t, 22, siat.EmisionOnline, 1, 1, 0, cufdControl)

	// 4. Construir Factura
	nombre := "JUAN PEREZ"
	cabecera := invoices.NewTelecomunicacionesCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EMPRESA TELECOM").
		WithMunicipio("La Paz").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("Av. Principal 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley 453...").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewTelecomunicacionesDetalleBuilder().
		WithActividadEconomica("611000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Internet").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := invoices.NewTelecomunicacionesBuilder().
		WithModalidad(tc.Modalidad).
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

	// 6. Solicitud de recepción
	req := models.Telecomunicaciones().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(22).
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
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionFactura(context.Background(), tc.Config, req)

	if err == nil && resp != nil {
		log.Printf("Respuesta Recepcion: %+v", resp.Body.Content)
	}
}
