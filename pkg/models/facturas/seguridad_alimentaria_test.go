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

func TestSeguridadAlimentaria_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "ENTIDAD DE APOYO"

	// Sector 7 = Seguridad Alimentaria
	// Tipo Factura 2 = Sin Derecho a Crédito Fiscal
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 7, 1, 0, cufdControl)

	cabecera := facturas.NewSeguridadAlimentariaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EMAPA").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. CAMACHO 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(5). // NIT
		WithNumeroDocumento("1234567").
		WithCodigoCliente("EMAPA-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(100.00).
		WithLeyenda("Leyenda Seguridad Alimentaria").
		WithUsuario("admin").
		Build()

	detalle := facturas.NewSeguridadAlimentariaDetalleBuilder().
		WithActividadEconomica("011100").
		WithCodigoProductoSin(11111).
		WithCodigoProducto("ARROZ-001").
		WithDescripcion("ARROZ DE PRIMERA").
		WithCantidad(2.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(50.0).
		WithSubTotal(100.0).
		Build()

	factura := facturas.NewSeguridadAlimentariaBuilder().
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
		WithCodigoDocumentoSector(7).
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

func TestSeguridadAlimentaria_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "ENTIDAD DE APOYO"

	// Sector 7 = Seguridad Alimentaria
	// Tipo Factura 2 = Sin Derecho a Crédito Fiscal
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 7, 2, 0, cufdControl)

	cabecera := facturas.NewSeguridadAlimentariaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EMAPA").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(2).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. CAMACHO 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(5). // NIT
		WithNumeroDocumento("1234567").
		WithCodigoCliente("EMAPA-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(200.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(200.00).
		WithLeyenda("Leyenda Seguridad Alimentaria").
		WithUsuario("admin").
		Build()

	detalle := facturas.NewSeguridadAlimentariaDetalleBuilder().
		WithActividadEconomica("011100").
		WithCodigoProductoSin(11111).
		WithCodigoProducto("HARINA-001").
		WithDescripcion("HARINA DE TRIGO").
		WithCantidad(4.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(50.0).
		WithSubTotal(200.0).
		Build()

	factura := facturas.NewSeguridadAlimentariaBuilder().
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
		WithCodigoDocumentoSector(7).
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
