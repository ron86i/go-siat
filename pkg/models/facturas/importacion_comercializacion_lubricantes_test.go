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

func TestImportacionLubricantes_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	// Sector 44: Comercializacion Lubricantes
	cuf, err := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 44, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0
	cantidad := 10.0
	precioUnitario := 50.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem
	montoDeduccion := 0.0
	montoTotalSujetoIva := montoTotal - montoDeduccion

	cabecera := facturas.NewImportacionComercializacionLubricantesCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("ESQUINA AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoDeduccionIehdDS25530(&montoDeduccion).
		WithMontoTotalSujetoIva(montoTotalSujetoIva).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalle := facturas.NewImportacionComercializacionLubricantesDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("LUB-123").
		WithDescripcion("LUBRICANTE SINTETICO").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		WithCantidadLitros(0.0).
		WithPorcentajeDeduccionIehdDS25530(nil).
		Build()

	factura := facturas.NewImportacionComercializacionLubricantesBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error marshaling XML: %v", err)
	}

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
		WithCodigoDocumentoSector(44). // Sector 44
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1). // Derecho a credito fiscal
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

func TestImportacionLubricantes_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	// Sector 44: Comercializacion Lubricantes
	cuf, err := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 44, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0
	cantidad := 10.0
	precioUnitario := 50.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem
	montoDeduccion := 0.0
	montoTotalSujetoIva := montoTotal - montoDeduccion

	cabecera := facturas.NewImportacionComercializacionLubricantesCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("ESQUINA AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoDeduccionIehdDS25530(&montoDeduccion).
		WithMontoTotalSujetoIva(montoTotalSujetoIva).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		Build()

	detalle := facturas.NewImportacionComercializacionLubricantesDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("LUB-123").
		WithDescripcion("LUBRICANTE SINTETICO").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		WithCantidadLitros(0.0).
		WithPorcentajeDeduccionIehdDS25530(nil).
		Build()

	factura := facturas.NewImportacionComercializacionLubricantesBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error marshaling XML: %v", err)
	}

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
		WithCodigoDocumentoSector(44). // Sector 44
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1). // Derecho a credito fiscal
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
