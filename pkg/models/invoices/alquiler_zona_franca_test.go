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

func TestAlquilerZonaFranca_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 22, 1, 0, cufdControl)

	nombreRazonSocial := "LOCATARIO ZF"
	codigoPuntoVenta := 0

	cabecera := invoices.NewAlquilerZFCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("ALQUILERES ZF").
		WithMunicipio("IQUIQUE").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("PUERTO SECO").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("99887766").
		WithCodigoCliente("ALQ-ZF-01").
		WithPeriodoFacturado("MARZO 2024").
		WithCodigoMetodoPago(1).
		WithMontoTotal(2000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(2000.00).
		WithLeyenda("Leyenda Alquiler ZF").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewAlquilerZFDetalleBuilder().
		WithActividadEconomica("681000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("ALQ-001").
		WithDescripcion("ALQUILER DE GALPON").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(2000.0).
		WithSubTotal(2000.00).
		Build()

	factura := invoices.NewAlquilerZFBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(22).
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
}
