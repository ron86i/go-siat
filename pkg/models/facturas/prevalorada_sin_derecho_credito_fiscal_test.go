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

func TestPrevaloradaSinDerechoCreditoFiscal_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	// Sector 36 = Prevalorada sin Derecho Crédito Fiscal
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, siat.EmisionOnline, 2, 36, 1, 0, cufdControl)

	montoTotal := 50.0
	codigoPuntoVenta := 0

	cabecera := facturas.NewPrevaloradaSinDerechoCreditoFiscalCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EVENTOS S.A.").
		WithMunicipio("SANTA CRUZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("CALLE EVENTOS 123").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("admin").
		Build()

	detalle := facturas.NewPrevaloradaSinDerechoCreditoFiscalDetalleBuilder().
		WithActividadEconomica("900000").
		WithCodigoProductoSin(99001).
		WithCodigoProducto("ENTRADA-01").
		WithDescripcion("ENTRADA ESPECTACULO PUBLICO").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(montoTotal).
		WithSubTotal(montoTotal).
		Build()

	factura := facturas.NewPrevaloradaSinDerechoCreditoFiscalBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		WithDetalle(detalle).
		Build()

	// Verify XML structure and fixed fields for Sector 36
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
		WithCodigoDocumentoSector(36).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(2).
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
