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

func TestVentaMineralBCB_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "BANCO CENTRAL DE BOLIVIA"

	// Sector 52 = Venta de Mineral al BCB
	// Tipo Factura 1 = Compra Venta
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 52, 1, 0, cufdControl)

	cabecera := facturas.NewVentaMineralBCBCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("MINERA AURIFERA S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("CALLE ORO 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithDireccionComprador("CALLE AYACUCHO ESQ. MERCADO").
		WithCodigoTipoDocumentoIdentidad(5). // NIT
		WithNumeroDocumento("1020304050").
		WithConcentradoGranel("ORO BULION").
		WithOrigen("MINA LA RIVA").
		WithCodigoCliente("BCB-001").
		WithCodigoMoneda(1). // BOB
		WithNumeroLote("LOTE-BCB-2024").
		WithKilosNetosHumedos(10.00).
		WithKilosNetosSecos(10.00).
		WithCodigoMetodoPago(1).
		WithMontoTotal(50000.00).
		WithMontoTotalMoneda(50000.00).
		WithSubTotal(50000.00).
		WithGastosRealizacion(0.00).
		WithIva(0.00).
		WithLeyenda("Leyenda Mineria BCB").
		WithUsuario("operador").
		Build()

	detalle := facturas.NewVentaMineralBCBDetalleBuilder().
		WithActividadEconomica("072200").
		WithCodigoProductoSin(123456).
		WithCodigoProducto("ORO-001").
		WithCodigoNandina("7108120000").
		WithDescripcion("ORO EN BARRAS").
		WithDescripcionLeyes("PUREZA 0.999").
		WithCantidadExtraccion(10.0).
		WithCantidad(10.0).
		WithUnidadMedidaExtraccion(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(5000.0).
		WithSubTotal(50000.0).
		Build()

	factura := facturas.NewVentaMineralBCBBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
		Build()

	xmlData, _ := xml.Marshal(factura)
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error firmando XML: %v", err)
	}

	hashString, encodedArchivo, err := utils.CompressAndHash(signedXML)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.Electronica().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(52).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
