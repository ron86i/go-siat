package facturas_test

import (
	"context"
	"encoding/xml"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/facturas"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestMonedaExtranjera_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "CLIENTE CAMBIO MONEDA"
	codigoPuntoVenta := 0

	// Datos de la transacción (Venta de 100 USD @ 6.96)
	montoTotal := 696.00       // Equivalente en BS
	montoTotalMoneda := 100.00 // En USD
	tipoCambio := 6.96
	tipoCambioOficial := 6.96
	ingresoDiferenciaCambio := 0.00 // Ajuste si aplica

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 9, 1, codigoPuntoVenta, cufdControl)

	cabecera := facturas.NewMonedaExtranjeraCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("CASA DE CAMBIO XYZ").
		WithMunicipio("LPZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. CAMACHO").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-001").
		WithCodigoTipoOperacion(1). // Venta
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(0). // Usualmente 0 en Sector 9
		WithIngresoDiferenciaCambio(ingresoDiferenciaCambio).
		WithCodigoMoneda(2). // DOLARES
		WithTipoCambio(tipoCambio).
		WithMontoTotalMoneda(montoTotalMoneda).
		WithLeyenda("Leyenda Moneda Extranjera").
		WithUsuario("operador01").
		WithTipoCambioOficial(tipoCambioOficial).
		Build()

	detalle := facturas.NewMonedaExtranjeraDetalleBuilder().
		WithActividadEconomica("661100").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("USD-SELL").
		WithDescripcion("VENTA DE DOLARES").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(montoTotal).
		WithSubTotal(montoTotal).
		Build()

	factura := facturas.NewMonedaExtranjeraBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, _ := xml.Marshal(factura)
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error firmando XML: %v", err)
	}

	hashString, encodedArchivo, err := utils.CompressAndHash(signedXML)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Electronica().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(9). // Sector 9
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(codigoPuntoVenta).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(2).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
