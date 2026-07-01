package invoices_test

import (
	"context"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/ron86i/go-siat/v2/pkg/models/invoices"
	"github.com/ron86i/go-siat/v2/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestBiodiesel_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "MINISTERIO DE HIDROCARBUROS"

	// Sector 54 = Biodiesel, Tipo Factura 2 = Sin Derecho a Crédito Fiscal
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 54, 1, 0, cufdControl)

	cabecera := invoices.NewBiodieselCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("PLANTA BIODIESEL S.A.").
		WithMunicipio("SANTA CRUZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("PARQUE INDUSTRIAL").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(5). // NIT
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-BIO-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(10000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(10000.00).
		WithLeyenda("Leyenda Biodiesel").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewBiodieselDetalleBuilder().
		WithActividadEconomica("104000").
		WithCodigoProductoSin(54321).
		WithCodigoProducto("BIO-001").
		WithDescripcion("SUMINISTRO DE BIODIESEL").
		WithCantidad(1000.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(10.0).
		WithSubTotal(10000.0).
		Build()

	factura := invoices.NewBiodieselBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(54).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(2).
		WithFechaEnvio(fechaEmision)

	err := builderReq.WithFactura(factura, tc.Client.Config())
	if err != nil {
		t.Fatalf("error al preparar factura: %v", err)
	}

	req := builderReq.Build()

	resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}

func TestBiodiesel_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "MINISTERIO DE HIDROCARBUROS"

	// Sector 54 = Biodiesel, Tipo Factura 2 = Sin Derecho a Crédito Fiscal
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 54, 2, 0, cufdControl)

	cabecera := invoices.NewBiodieselCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("PLANTA BIODIESEL S.A.").
		WithMunicipio("SANTA CRUZ").
		WithNumeroFactura(2).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("PARQUE INDUSTRIAL").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(5). // NIT
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-BIO-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(5000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(5000.00).
		WithLeyenda("Leyenda Biodiesel").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewBiodieselDetalleBuilder().
		WithActividadEconomica("104000").
		WithCodigoProductoSin(54321).
		WithCodigoProducto("BIO-002").
		WithDescripcion("SUMINISTRO DE BIODIESEL B100").
		WithCantidad(500.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(10.0).
		WithSubTotal(5000.0).
		Build()

	factura := invoices.NewBiodieselBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(54).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(2).
		WithFechaEnvio(fechaEmision)

	err := builderReq.WithFactura(factura, tc.Client.Config())
	if err != nil {
		t.Fatalf("error al preparar factura: %v", err)
	}

	req := builderReq.Build()

	resp, err := tc.Client.Computarizada().RecepcionFactura(context.Background(), req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
