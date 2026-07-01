package invoices_test

import (
	"context"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestAlquilerBienInmueble_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 2, 1, 0, cufdControl)

	cabecera := invoices.NewAlquilerBienInmuebleCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("MI EMPRESA S.R.L.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. 6 DE AGOSTO 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-001").
		WithPeriodoFacturado("ENERO 2024").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100.00).
		WithMontoTotalSujetoIva(100.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(100.00).
		WithLeyenda("Leyenda Factura").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewAlquilerBienInmuebleDetalleBuilder().
		WithActividadEconomica("681011").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("ALQ-001").
		WithDescripcion("ALQUILER DE OFICINA CENTRAL").
		WithCantidad(1.0).
		WithUnidadMedida(58).
		WithPrecioUnitario(100.00).
		WithSubTotal(100.00).
		Build()

	factura := invoices.NewAlquilerBienInmuebleBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(2).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
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
