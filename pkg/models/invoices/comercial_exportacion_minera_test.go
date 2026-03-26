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

func TestComercialExportacionMinera_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "MINING CORP INTERNATIONAL"
	ruex := "12345-RUEX"
	nim := "67890-NIM"

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 20, 1, 0, cufdControl)

	cabecera := invoices.NewComercialExportacionMineraCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("MINERA BOLIVIA S.A.").
		WithMunicipio("POTOSI").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("CALLE MINERO 456").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithDireccionComprador("123 Global Street, New York").
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("99887766").
		WithRuex(&ruex).
		WithNim(&nim).
		WithConcentradoGranel("CONCENTRADO DE ZINC").
		WithOrigen("POTOSI - BOLIVIA").
		WithPuertoDestino("ANTOFAGASTA").
		WithPaisDestino(41).
		WithIncoterm("CIF").
		WithCodigoCliente("EXP-MIN-001").
		WithCodigoMoneda(2).
		WithTipoCambio(6.96).
		WithTipoCambioANB(6.96).
		WithNumeroLote("LOT-2024-001").
		WithKilosNetosHumedos(10000.00).
		WithKilosNetosSecos(9500.00).
		WithCodigoMetodoPago(1).
		WithMontoTotal(50000.00).
		WithMontoTotalMoneda(7183.91).
		WithGastosRealizacion(500.00).
		WithLeyenda("Leyenda Minera").
		WithUsuario("operador_min").
		Build()

	detalle := invoices.NewComercialExportacionMineraDetalleBuilder().
		WithActividadEconomica("072900").
		WithCodigoProductoSin(54321).
		WithCodigoProducto("MIN-ZN-001").
		WithCodigoNandina("2608.00.00.00").
		WithDescripcion("CONCENTRADO DE ZINC").
		WithDescripcionLeyes("ZN: 50%, AG: 150 G/T").
		WithCantidadExtraccion(10.0).
		WithCantidad(10.0).
		WithUnidadMedidaExtraccion(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(5000.00).
		WithSubTotal(50000.00).
		Build()

	factura := invoices.NewComercialExportacionMineraBuilder().
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
		WithCodigoDocumentoSector(20).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(2).
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
