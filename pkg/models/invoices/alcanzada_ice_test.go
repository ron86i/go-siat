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

func TestAlcanzadaIce_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 15, 1, 0, cufdControl)

	nombreRazonSocial := "COMPRADOR BEBIDAS"
	codigoPuntoVenta := 0
	cantidad := 10.0
	precioUnitario := 20.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	iceEspecifico := 15.0
	icePorcentual := 10.0
	cabecera := invoices.NewAlcanzadaIceCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("CERVECERIA S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. MONTES").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-ICE-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(montoTotal).
		WithMontoIceEspecifico(&iceEspecifico).
		WithMontoIcePorcentual(&icePorcentual).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Leyenda ICE").
		WithUsuario("usuario").
		Build()

	detalle := invoices.NewAlcanzadaIceDetalleBuilder().
		WithActividadEconomica("110300").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("BEB-01").
		WithDescripcion("CERVEZA PRUEBA").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		WithMarcaIce(1).
		WithAlicuotaIva(&montoDescuento).
		WithAlicuotaEspecifica(&iceEspecifico).
		WithAlicuotaPorcentual(&icePorcentual).
		WithCantidadIce(&cantidad).
		Build()

	factura := invoices.NewAlcanzadaIceBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(15).
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
