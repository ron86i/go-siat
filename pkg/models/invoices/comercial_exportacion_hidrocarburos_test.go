package invoices_test

import (
	"context"
	"encoding/xml"
	"log"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestComercialExportacionHidro_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "CLIENTE HIDROCARBUROS"
	codigoPuntoVenta := 0

	// Datos de la transacción (Valores grandes para probar precisión)
	montoDetalle := 99500.03
	costosNacionales := map[string]float64{"GASTOS_ADMIN": 500.00}
	totalCostosNacionales := 500.00
	totalGastosNacionalesFob := montoDetalle + totalCostosNacionales // 100000.03

	totalGastosInternacionales := 0.00
	montoTotal := totalGastosNacionalesFob + totalGastosInternacionales // 100000.03

	tipoCambio := 6.96
	montoTotalBs := montoTotal * tipoCambio

	// El CUF debe generarse con tipoFactura = 2 (Exportación) para el Sector 43
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 43, 1, 0, cufdControl)

	cabecera := invoices.NewComercialExportacionHidroCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EMPRESA PETROLERA").
		WithMunicipio("SCZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. PIRAI").
		WithCodigoPuntoVenta(utils.IntPtr(0)).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-HIDRO").
		WithDireccionComprador("AV. BRASIL").
		WithIncoterm("FOB").
		WithIncotermDetalle("CIUDAD BOLIVIA").
		WithPuertoDestino("SANTOS").
		WithLugarDestino("BRASIL").
		WithCodigoPais(1).
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotalBs).
		WithCostosGastosNacionales(costosNacionales).
		WithTotalGastosNacionalesFob(totalGastosNacionalesFob).
		WithCostosGastosInternacionales(nil).
		WithTotalGastosInternacionales(utils.Float64Ptr(totalGastosInternacionales)).
		WithMontoDetalle(montoDetalle).
		WithMontoTotalSujetoIva(0).
		WithCodigoMoneda(2). // USD
		WithTipoCambio(tipoCambio).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Leyenda Hidrocarburos").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewComercialExportacionHidroDetalleBuilder().
		WithActividadEconomica("466100").
		WithCodigoProductoSin(9988).
		WithCodigoProducto("GAS-001").
		WithCodigoNandina("2701110000").
		WithDescripcion("GAS NATURAL").
		WithCantidad(100.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(995.0003).
		WithSubTotal(montoDetalle).
		Build()

	factura := invoices.NewComercialExportacionHidroBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, _ := xml.Marshal(factura)
	// Log del XML para verificar los totales enviados
	log.Printf("XML Generado (Sector 43): %s", string(xmlData))

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
		WithCodigoDocumentoSector(43).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(codigoPuntoVenta).
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
