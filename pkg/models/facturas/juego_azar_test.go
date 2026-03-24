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

func TestJuegoAzar_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	// Sector 18 = Juego de Azar
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 18, 1, 0, cufdControl)

	cabecera := facturas.NewJuegoAzarCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("CASINO ROYALE").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. 16 DE JULIO 456").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-18-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(1130.50).
		WithMontoTotalIj(130.50).
		WithMontoTotalSujetoIpj(870.00).
		WithMontoTotalSujetoIva(1000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(1130.50).
		WithLeyenda("Leyenda Juego de Azar").
		WithUsuario("admin").
		Build()

	detalle := facturas.NewJuegoAzarDetalleBuilder().
		WithActividadEconomica("920000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("P-001").
		WithDescripcion("FICHAS DE JUEGO").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(1000.0).
		WithSubTotal(1000.0).
		Build()

	factura := facturas.NewJuegoAzarBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
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
		WithCodigoDocumentoSector(18).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
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

func TestJuegoAzar_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	// Sector 18 = Juego de Azar
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 18, 2, 0, cufdControl)

	cabecera := facturas.NewJuegoAzarCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("CASINO ROYALE").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(2).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. 16 DE JULIO 456").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-18-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(1130.50).
		WithMontoTotalIj(130.50).
		WithMontoTotalSujetoIpj(870.00).
		WithMontoTotalSujetoIva(1000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(1130.50).
		WithLeyenda("Leyenda Juego de Azar").
		WithUsuario("admin").
		Build()

	detalle := facturas.NewJuegoAzarDetalleBuilder().
		WithActividadEconomica("920000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("P-002").
		WithDescripcion("DERECHO DE JUEGO").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(1000.0).
		WithSubTotal(1000.0).
		Build()

	factura := facturas.NewJuegoAzarBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	xmlData, _ := xml.Marshal(factura)
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlData)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Computarizada().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(18).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := tc.Client.Computarizada().RecepcionFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
