package invoices_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestCompraVenta_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 1, 1, 0, cufdControl)

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	cabecera := invoices.NewCompraVentaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		WithCodigoDocumentoSector(1).
		Build()

	detalle := invoices.NewCompraVentaDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("abc123").
		WithDescripcion("GASA").
		WithCantidad(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := invoices.NewCompraVentaBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	builder := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithFechaEnvio(fechaEmision)

	err := builder.WithFactura(factura, tc.Client.Config())
	if err != nil {
		t.Fatalf("error preparando factura con el constructor: %v", err)
	}
	req := builder.Build()

	resp, err := tc.Client.CompraVenta().RecepcionFactura(context.Background(), req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}

func TestCompraVenta_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)
	tipoFactura := 1 // Factura de compra venta
	var numeroFactura int64 = 1
	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, tc.Sucursal, tc.Modalidad, 1, tipoFactura, 1, numeroFactura, tc.PuntoVenta, cufdControl)

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	cabecera := invoices.NewCompraVentaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(numeroFactura).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		WithCodigoDocumentoSector(1).
		Build()

	detalle := invoices.NewCompraVentaDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("abc123").
		WithDescripcion("GASA").
		WithCantidad(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := invoices.NewCompraVentaBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	builder := models.NewRecepcionFacturaBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithFechaEnvio(fechaEmision)

	err := builder.WithFactura(factura, tc.Client.Config())
	if err != nil {
		t.Fatalf("error preparando factura computarizada con el constructor: %v", err)
	}
	req := builder.Build()

	resp, err := tc.Client.CompraVenta().RecepcionFactura(context.Background(), req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}

func TestCompraVenta_Masiva(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	var facturas []any
	fechaEmision := time.Now()
	codigoPuntoVenta := 0
	for i := 1; i <= 5; i++ {
		cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, siat.EmisionMasiva, 1, 1, int64(i), 0, cufdControl)
		nombreRazonSocial := "JUAN PEREZ"
		cabecera := invoices.NewCompraVentaCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Ronaldo Rua").
			WithMunicipio("Tarija").
			WithNumeroFactura(int64(i)).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("AVENIDA LA PAZ").
			WithCodigoPuntoVenta(&codigoPuntoVenta).
			WithFechaEmision(fechaEmision).
			WithNombreRazonSocial(&nombreRazonSocial).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5115889").
			WithCodigoCliente(strconv.Itoa(i)).
			WithCodigoMetodoPago(1).
			WithMontoTotal(100).
			WithMontoTotalSujetoIva(100).
			WithCodigoMoneda(1).
			WithTipoCambio(1).
			WithMontoTotalMoneda(100).
			WithLeyenda("Ley N° 453").
			WithUsuario("usuario").
			WithCodigoDocumentoSector(1).
			Build()

		detalle := invoices.NewCompraVentaDetalleBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("abc123").
			WithDescripcion("GASA").
			WithCantidad(1).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(100).
			Build()

		factura := invoices.NewCompraVentaBuilder().
			WithModalidad(tc.Modalidad).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		facturas = append(facturas, factura)
	}

	builder := models.NewRecepcionMasivaFacturaBuilder().
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithFechaEnvio(fechaEmision).
		WithCodigoModalidad(tc.Modalidad)

	err := builder.WithFacturas(facturas, tc.Client.Config())
	if err != nil {
		t.Fatalf("error preparando paquete masivo con el constructor: %v", err)
	}
	req := builder.Build()

	resp, err := tc.Client.CompraVenta().RecepcionMasivaFactura(context.Background(), req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
