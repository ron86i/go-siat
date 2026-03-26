package invoices_test

import (
	"context"
	"encoding/xml"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestComercializacionGnGlp_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	// Sector 39 = Comercialización de GN y GLP
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 39, 1, 0, cufdControl)

	cabecera := invoices.NewComercializacionGnGlpCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("GAS DE ORIENTE S.A.").
		WithMunicipio("SANTA CRUZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. PETROLERA 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-39-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(500.00).
		WithMontoTotalSujetoIva(500.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(500.00).
		WithLeyenda("Leyenda GN/GLP").
		WithUsuario("admin").
		Build()

	detalle := invoices.NewComercializacionGnGlpDetalleBuilder().
		WithActividadEconomica("352000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("GAS-01").
		WithDescripcion("SUMINISTRO DE GAS").
		WithCantidad(100.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(5.0).
		WithSubTotal(500.0).
		Build()

	factura := invoices.NewComercializacionGnGlpBuilder().
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
		WithCodigoDocumentoSector(39).
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

func TestComercializacionGnGlp_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	// Sector 39 = Comercialización de GN y GLP
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 39, 1, 0, cufdControl)

	cabecera := invoices.NewComercializacionGnGlpCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("GAS DE ORIENTE S.A.").
		WithMunicipio("SANTA CRUZ").
		WithNumeroFactura(2).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. PETROLERA 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-39-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(300.00).
		WithMontoTotalSujetoIva(300.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(300.00).
		WithLeyenda("Leyenda GN/GLP").
		WithUsuario("admin").
		Build()

	detalle := invoices.NewComercializacionGnGlpDetalleBuilder().
		WithActividadEconomica("352000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("GAS-02").
		WithDescripcion("CARGA GLP").
		WithCantidad(10.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(30.0).
		WithSubTotal(300.0).
		Build()

	factura := invoices.NewComercializacionGnGlpBuilder().
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
		WithCodigoDocumentoSector(39).
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
