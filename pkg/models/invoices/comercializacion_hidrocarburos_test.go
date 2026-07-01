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

func TestComercializacionHidrocarburos_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 12, 1, 0, cufdControl)

	codigoPais := 1
	cabecera := invoices.NewComercializacionHidroCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("HIDROCARBUROS S.A.").
		WithMunicipio("SCZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. PETROLERA").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-HIDRO-01").
		WithPlacaVehiculo("1234ABC").
		WithCodigoPais(&codigoPais).
		WithCodigoMetodoPago(1).
		WithMontoTotal(100.00).
		WithMontoTotalSujetoIva(100.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(100.00).
		WithLeyenda("Leyenda Hidrocarburos").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewComercializacionHidroDetalleBuilder().
		WithActividadEconomica("466100").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("HIDRO-001").
		WithDescripcion("DIESEL OIL").
		WithCantidad(1.0).
		WithUnidadMedida(58).
		WithPrecioUnitario(100.0).
		WithSubTotal(100.00).
		Build()

	factura := invoices.NewComercializacionHidroBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(12).
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

	resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
