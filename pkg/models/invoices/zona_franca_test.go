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

func TestZonaFranca_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 5, 1, 0, cufdControl)

	cabecera := invoices.NewZonaFrancaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("ZONA FRANCA IQUIQUE").
		WithMunicipio("ORURO").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("PUERTO SECO").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-ZF-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(2000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(2000.00).
		WithLeyenda("Leyenda Zona Franca").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewZonaFrancaDetalleBuilder().
		WithActividadEconomica("469000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("ZF-001").
		WithDescripcion("MERCADERIA ZONA FRANCA").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(2000.0).
		WithSubTotal(2000.00).
		Build()

	factura := invoices.NewZonaFrancaBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(5).
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
