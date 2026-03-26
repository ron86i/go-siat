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

func TestLibreConsignacion_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	codigoPuntoVenta := 0

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 4, 1, codigoPuntoVenta, cufdControl)

	cabecera := invoices.NewLibreConsignacionCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("IMPORTADORA S.R.L.").
		WithMunicipio("SCZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. BANZER").
		WithFechaEmision(fechaEmision).
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithCodigoMetodoPago(1).
		WithMontoTotal(1000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(1000.00).
		WithLeyenda("Leyenda Libre Consignación").
		WithUsuario("operador01").
		WithCodigoTipoDocumentoIdentidad(5).
		WithCodigoPais(1).
		WithPuertoDestino("PUERTO 1").
		WithCodigoCliente("123456789").
		Build()

	detalle := invoices.NewLibreConsignacionDetalleBuilder().
		WithActividadEconomica("469000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("CONS-001").
		WithDescripcion("MERCADERIA EN CONSIGNACION").
		WithCodigoNandina("12345678").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(1000.0).
		WithSubTotal(1000.00).
		Build()

	factura := invoices.NewLibreConsignacionBuilder().
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
		WithCodigoDocumentoSector(4).
		WithCodigoEmision(1).
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
