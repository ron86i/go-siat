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

func TestVentaMineral_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "MINING COMPANY CORP"

	// Sector 21 = Venta de Mineral
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 21, 1, 0, cufdControl)

	cabecera := invoices.NewVentaMineralCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("MINERIA BOLIVIANA S.R.L.").
		WithMunicipio("POTOSI").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("CALLE MINERO 789").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithDireccionComprador("MAIN ST. 456, NEW YORK").
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithConcentradoGranel("CONCENTRADO DE PLATA").
		WithOrigen("MINA POTOSI").
		WithIncoterm("CIF").
		WithCodigoCliente("MIN-001").
		WithCodigoMoneda(2). // USD
		WithTipoCambio(6.96).
		WithTipoCambioANB(6.96).
		WithNumeroLote("LOTE-2024-001").
		WithKilosNetosHumedos(1000.00).
		WithKilosNetosSecos(1000.00).
		WithCodigoMetodoPago(1).
		WithMontoTotal(396000.06).
		WithMontoTotalSujetoIva(56896.56).
		WithMontoTotalMoneda(56896.56).
		WithSubTotal(50000.01).
		WithGastosRealizacion(500.00).
		WithIva(7396.55).
		WithLeyenda("Leyenda Mineria").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewVentaMineralDetalleBuilder().
		WithActividadEconomica("071000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("MIN-AG").
		WithCodigoNandina("2603.00.00.00").
		WithDescripcion("CONCENTRADO DE PLATA").
		WithDescripcionLeyes("LEY 99.9%").
		WithCantidadExtraccion(1000.0).
		WithCantidad(950.0).
		WithUnidadMedidaExtraccion(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(52.63159).
		WithSubTotal(50000.0105).
		Build()

	factura := invoices.NewVentaMineralBuilder().
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
		WithCodigoDocumentoSector(21).
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

func TestVentaMineral_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "MINING COMPANY CORP"

	// Sector 21 = Venta de Mineral
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 21, 2, 0, cufdControl)

	cabecera := invoices.NewVentaMineralCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("MINERIA BOLIVIANA S.R.L.").
		WithMunicipio("POTOSI").
		WithNumeroFactura(2).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("CALLE MINERO 789").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithDireccionComprador("MAIN ST. 456, NEW YORK").
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithConcentradoGranel("CONCENTRADO DE COBRE").
		WithOrigen("MINA ORURO").
		WithIncoterm("FOB").
		WithCodigoCliente("MIN-002").
		WithCodigoMoneda(1). // BOB
		WithTipoCambio(1.0).
		WithTipoCambioANB(1.0).
		WithNumeroLote("LOTE-2024-002").
		WithKilosNetosHumedos(1000.00).
		WithKilosNetosSecos(1000.00).
		WithCodigoMetodoPago(1).
		WithMontoTotal(56896.56).
		WithMontoTotalSujetoIva(56896.56).
		WithMontoTotalMoneda(56896.56).
		WithSubTotal(50000.01). // Correspondiente al detalle
		WithGastosRealizacion(500.00).
		WithIva(7396.55).
		WithLeyenda("Leyenda Mineria").
		WithUsuario("operador").
		Build()

	detalle := invoices.NewVentaMineralDetalleBuilder().
		WithActividadEconomica("071000").
		WithCodigoProductoSin(12346).
		WithCodigoProducto("MIN-CU").
		WithCodigoNandina("2603.00.00.00").
		WithDescripcion("CONCENTRADO DE COBRE").
		WithDescripcionLeyes("LEY 25%").
		WithCantidadExtraccion(1000.0).
		WithCantidad(1000.0).
		WithUnidadMedidaExtraccion(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(50.00001).
		WithSubTotal(50000.01).
		Build()

	factura := invoices.NewVentaMineralBuilder().
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
		WithCodigoDocumentoSector(21).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
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
