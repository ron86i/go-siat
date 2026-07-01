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

func TestSectorEducativoZF_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "PADRE DE FAMILIA"

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 46, 1, 0, cufdControl)

	cabecera := invoices.NewSectorEducativoZFCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("COLEGIO INTERNACIONAL ZF").
		WithMunicipio("IQUIQUE").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. LIBERTAD 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-001").
		WithNombreEstudiante("JUANITO PEREZ").
		WithPeriodoFacturado("MARZO 2024").
		WithCodigoMetodoPago(1).
		WithMontoTotal(1500.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(1500.00).
		WithLeyenda("Leyenda Educativa").
		WithUsuario("operador_edu").
		Build()

	detalle := invoices.NewSectorEducativoZFDetalleBuilder().
		WithActividadEconomica("851000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("EDU-001").
		WithDescripcion("MENSULIDAD MARZO").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(1500.00).
		WithSubTotal(1500.00).
		Build()

	factura := invoices.NewSectorEducativoZFBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	builderReq := models.NewRecepcionFacturaBuilder().
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(46).
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
