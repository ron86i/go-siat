package invoices_test

import (
	"context"
	"encoding/xml"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/internal/core/domain/documents"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/ron86i/go-siat/v2/pkg/models/invoices"
)

func TestTelecomunicacionesZFBuilder(t *testing.T) {
	fechaEmision := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	telefono := "2222222"
	nombre := "JUAN PEREZ (ZONA FRANCA)"
	nitConjunto := int64(123456789)
	nSerie := "SN-ZF-123"

	cabecera := invoices.NewTelecomunicacionesZFCabeceraBuilder().
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

	detalle := invoices.NewTelecomunicacionesZFDetalleBuilder().
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
		factura := invoices.NewTelecomunicacionesZFBuilder().
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

func TestTelecomunicacionesZF_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)

	service := tc.Client.Telecomunicaciones()

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	// 3. Generar CUF (Sector 49, TipoFactura 2)
	cuf := tc.GetCuf(t, 49, siat.EmisionOnline, 2, 1, 0, cufdControl)

	// 4. Construir Factura ZF
	nombre := "JUAN PEREZ (ZF)"
	cabecera := invoices.NewTelecomunicacionesZFCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EMPRESA TELECOM ZF").
		WithMunicipio("La Paz").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(tc.Sucursal).
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

	detalle := invoices.NewTelecomunicacionesZFDetalleBuilder().
		WithActividadEconomica("611000").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Servicio ZF").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(200).
		WithSubTotal(200).
		Build()

	factura := invoices.NewTelecomunicacionesZFBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	// 5. Serializar, Firmar, Comprimir
	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoDocumentoSector(49).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(2). // 2 para ZF
		WithFechaEnvio(fechaEmision)

	err := builderReq.WithFactura(factura, tc.Client.Config())
	if err != nil {
		t.Fatalf("error al preparar factura: %v", err)
	}

	req := builderReq.Build()

	resp, err := service.RecepcionFactura(context.Background(), req)

	if err == nil && resp != nil {
		log.Printf("Respuesta Recepcion ZF: %+v", resp.Body.Content)
	}
}
