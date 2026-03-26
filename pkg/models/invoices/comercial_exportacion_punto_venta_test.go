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

func TestComercialExportacionPVenta_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "CLIENTE EXP PVENTA"
	codigoPuntoVenta := 0
	cantidad := 2.0
	precioUnitario := 500.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 45, 1, codigoPuntoVenta, cufdControl)

	cabecera := invoices.NewComercialExportacionPVentaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EMPRESA PVENTA EXPORT").
		WithMunicipio("SCZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. EXPORTADORES P").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-PVE-01").
		WithDireccionComprador("AV. COMPRADOR INT").
		WithIncoterm("FOB").
		WithIncotermDetalle("FOB DETALLE").
		WithPuertoDestino("PUERTO ABC").
		WithLugarDestino("DESTINO FINAL").
		WithCodigoPais(1).
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(montoTotal). // Wait, builder sets 0, but tests can inject this. Let's just use 0 here maybe if SIAT requires 0. Let's provide montoTotal.
		WithMontoDetalle(montoTotal).
		WithPrecioValorBruto(&montoTotal).
		WithCodigoMoneda(2).
		WithTipoCambio(6.96).
		WithMontoTotalMoneda(montoTotal / 6.96).
		WithLeyenda("Leyenda Exp PVenta").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewComercialExportacionPVentaDetalleBuilder().
		WithActividadEconomica("466100").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("PVE-01").
		WithCodigoNandina("11223344").
		WithDescripcion("PRODUCTO PARA EXPORTACION").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := invoices.NewComercialExportacionPVentaBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
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
		WithCodigoDocumentoSector(45).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(codigoPuntoVenta).
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
