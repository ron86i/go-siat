package invoices_test

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestCompraVentaBonificaciones_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 41, 1, 0, cufdControl) // Sector 41?

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	cabecera := invoices.NewCompraVentaBonificacionesCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(90).
		WithMontoTotalSujetoIva(90).
		WithDescuentoAdicional(utils.Float64Ptr(10.0)).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(90).
		WithLeyenda("Ley N° 453").
		WithUsuario("usuario").
		Build()

	detalle := invoices.NewCompraVentaBonificacionesDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("abc123").
		WithDescripcion("PRODUCTO CON BONIFICACION").
		WithCantidad(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithMontoDescuento(utils.Float64Ptr(0.0)).
		WithSubTotal(100).
		Build()

	factura := invoices.NewCompraVentaBonificacionesBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, _ := xml.Marshal(factura)
	hashString, encodedArchivo, _ := utils.CompressAndHash(xmlData)

	req := models.CompraVenta().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(35).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := tc.Client.CompraVenta().RecepcionFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}

func TestCompraVentaBonificaciones_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	emitirCompraVentaBonificacionesElectronicaIndividual(t, tc, cuis, cufd, cufdControl, 1)
}

func TestCompraVentaBonificaciones_ElectronicaAll(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	for i := 1; i <= 125; i++ {
		t.Run(fmt.Sprintf("FacturaBonificacion_%d", i), func(t *testing.T) {
			emitirCompraVentaBonificacionesElectronicaIndividual(t, tc, cuis, cufd, cufdControl, i)
			time.Sleep(50 * time.Millisecond)
		})
	}
}

func emitirCompraVentaBonificacionesElectronicaIndividual(t *testing.T, tc *TestContext, cuis, cufd, cufdControl string, nroFactura int) {
	fechaEmision := time.Now()
	// Sector 35: Compra Venta Bonificaciones
	cuf, err := utils.GenerarCUF(tc.Nit, fechaEmision, tc.Sucursal, tc.Modalidad, siat.EmisionOnline, 1, 35, nroFactura, tc.PuntoVenta, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nombreRazonSocial := "JUAN PEREZ"

	cabecera := invoices.NewCompraVentaBonificacionesCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(int64(nroFactura)).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(tc.Sucursal).
		WithDireccion("AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&tc.PuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(90).
		WithMontoTotalSujetoIva(90.0).
		WithDescuentoAdicional(utils.Float64Ptr(10.0)).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(90.0).
		WithLeyenda("Ley N° 453").
		WithUsuario("usuario").
		Build()

	detalle := invoices.NewCompraVentaBonificacionesDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("abc123").
		WithDescripcion("PRODUCTO CON BONIFICACION").
		WithCantidad(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithMontoDescuento(utils.Float64Ptr(0.0)).
		WithSubTotal(100).
		Build()

	factura := invoices.NewCompraVentaBonificacionesBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlBytes, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	xmlBytes, err = utils.SignXML(xmlBytes, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error al firmar: %v", err)
	}
	hashString, encodedArchivo, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	req := models.CompraVenta().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCodigoSucursal(tc.Sucursal).
		WithCodigoDocumentoSector(35).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()
	// 1. EMITIR FACTURA COMPRA VENTA BONIFICACIONES
	resp, err := tc.Client.CompraVenta().RecepcionFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("Factura %d - error en solicitud: %v", nroFactura, err)
	}

	if resp.Body.Fault != nil {
		t.Errorf("Factura %d - Fault SIAT: %s", nroFactura, resp.Body.Fault.String())
	} else {
		t.Logf("Factura %d - Respuesta SIAT: %+v", nroFactura, resp.Body.Content)
	}
	// Pequeño delay para que el SIAT procese el estado
	time.Sleep(50 * time.Millisecond)

	// 2. ANULAR FACTURA
	reqAnulacion := models.CompraVenta().NewAnulacionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(35).
		WithCodigoEmision(siat.EmisionOnline).
		WithTipoFacturaDocumento(1).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithNit(tc.Nit). // NIT es requerido en anulación
		WithCufd(cufd).
		WithCuis(cuis).
		WithCuf(cuf).
		WithCodigoMotivo(1). // 1: Factura mal emitida
		Build()

	respAnulacion, err := tc.Client.CompraVenta().AnulacionFactura(context.Background(), tc.Config, reqAnulacion)
	if err != nil {
		t.Fatalf("error en anulación Factura %d: %v", nroFactura, err)
	}

	if !respAnulacion.Body.Content.RespuestaServicioFacturacion.Transaccion {
		mensajes := ""
		for _, m := range respAnulacion.Body.Content.RespuestaServicioFacturacion.MensajesList {
			mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
		}
		t.Errorf("Anulación Factura %d falló: %s", nroFactura, mensajes)
		return // No podemos revertir si no se anuló
	}
	log.Printf("Factura %d ANULADA correctamente", nroFactura)

	// Otro delay para la reversión
	time.Sleep(50 * time.Millisecond)

	// 3. REVERTIR ANULACIÓN
	reqReversion := models.CompraVenta().NewReversionAnulacionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithNit(tc.Nit).
		WithCodigoDocumentoSector(35).
		WithTipoFacturaDocumento(1).
		WithCodigoEmision(1).
		WithCodigoModalidad(tc.Modalidad).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCuis(cuis).
		Build()

	respReversion, err := tc.Client.CompraVenta().ReversionAnulacionFactura(context.Background(), tc.Config, reqReversion)
	if err != nil {
		t.Fatalf("error en reversión Factura %d: %v", nroFactura, err)
	}

	if !respReversion.Body.Content.RespuestaServicioFacturacion.Transaccion {
		mensajes := ""
		for _, m := range respReversion.Body.Content.RespuestaServicioFacturacion.MensajesList {
			mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
		}
		t.Errorf("Reversión Factura %d falló: %s", nroFactura, mensajes)
		return
	}
	log.Printf("Factura %d REVERTIDA (vuelve a ser válida)", nroFactura)
}
