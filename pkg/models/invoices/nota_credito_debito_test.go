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
)

func TestNotaCreditoDebitoBuilder(t *testing.T) {
	fecha := time.Now()

	cabecera := invoices.NewNotaCreditoDebitoCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("Empresa Test").
		WithMunicipio("La Paz").
		WithNumeroNotaCreditoDebito(1).
		WithCuf("ABC123DEF").
		WithCufd("XYZ789").
		WithCodigoSucursal(0).
		WithDireccion("Av. Principal 123").
		WithFechaEmision(fecha).
		WithNombreRazonSocial(ptr("Cliente Test")).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5544332").
		WithCodigoCliente("CLI-001").
		WithNumeroFactura(100).
		WithNumeroAutorizacionCuf("FACT-ABC").
		WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
		WithMontoTotalOriginal(1000.00).
		WithMontoTotalDevuelto(900.00).
		WithMontoEfectivoCreditoDebito(100.00).
		WithLeyenda("Ley Nro 453").
		WithUsuario("admin").
		Build()

	detalle1 := invoices.NewNotaDetalleCreditoDebitoBuilder().
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P001").
		WithDescripcion("Producto A").
		WithCantidad(10).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(1000).
		WithCodigoDetalleTransaccion(1).
		Build()

	detalle2 := invoices.NewNotaDetalleCreditoDebitoBuilder().
		WithActividadEconomica("Venta").
		WithCodigoProductoSin(123).
		WithCodigoProducto("P002").
		WithDescripcion("Producto B devuelto").
		WithCantidad(5).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(500).
		WithCodigoDetalleTransaccion(2).
		Build()

	nota := invoices.NewNotaCreditoDebitoBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle1).
		AddDetalle(detalle2).
		Build()

	output, err := xml.MarshalIndent(nota, "", "  ")
	if err != nil {
		t.Fatalf("Error al serializar: %v", err)
	}

	t.Logf("Nota XML:\n%s", string(output))
}

func TestNotaCreditoDebitoIntegration_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)

	service := tc.Client.DocumentoAjuste()
	// codModalidad := siat.ModalidadComputarizada

	// 1. Obtener CUIS
	cuis := tc.GetCuis(t)

	// 2. Obtener CUFD
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fecha := time.Now()

	// 3. Generar CUF real para Nota (Sector 47) Tipo Documento (3)
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, 0, tc.Modalidad, siat.EmisionOnline, 3, 47, 1, 0, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroNotaCreditoDebito(1).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("Av. Principal 123").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-001").
			WithNumeroFactura(100).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
			WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
			WithMontoTotalOriginal(1000.00).
			WithMontoTotalDevuelto(500.00).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Producto A").
			WithCantidad(10).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1000).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P002").
			WithDescripcion("Producto B devuelto").
			WithCantidad(5).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(500).
			WithCodigoDetalleTransaccion(2).
			Build()).
		Build()

	xmlBytes, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	// 5. Preparar el envío
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	// 6. Crear solicitud de recepción
	req := models.NewRecepcionDocumentoAjusteBuilder().
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoDocumentoSector(47).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionDocumentoAjuste(context.Background(), req)
	if err != nil {
		t.Fatalf("Error en la comunicación con el SIAT: %v", err)
	}

	t.Logf("Respuesta SIAT: %+v", resp.Body.Content.RespuestaRecepcionFactura)
}

func TestNotaCreditoDebito_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	emitirNotaIndividual(t, tc, cuis, cufd, cufdControl, 1)
}

func TestNotaCreditoDebito_ElectronicaAll(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 1
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	for i := 1; i <= 125; i++ {
		t.Run(fmt.Sprintf("Nota_%d", i), func(t *testing.T) {
			emitirNotaIndividual(t, tc, cuis, cufd, cufdControl, i)
			time.Sleep(50 * time.Millisecond)
		})
	}
}

func emitirNotaIndividual(t *testing.T, tc *TestContext, cuis, cufd, cufdControl string, nroNota int) {
	fecha := time.Now()
	// Sector 24: Nota Crédito Débito, Tipo Doc 3
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, tc.Sucursal, tc.Modalidad, siat.EmisionOnline, 3, 24, int64(nroNota), tc.PuntoVenta, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nota := invoices.NewNotaCreditoDebitoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroNotaCreditoDebito(int64(nroNota)).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(tc.Sucursal).
			WithCodigoPuntoVenta(&tc.PuntoVenta).
			WithDireccion("Av. Principal 123").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-001").
			WithNumeroFactura(100).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
			WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
			WithMontoTotalOriginal(1000.00).
			WithMontoTotalDevuelto(500.00).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Producto A").
			WithCantidad(10).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1000).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P002").
			WithDescripcion("Producto B devuelto").
			WithCantidad(5).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(500).
			WithCodigoDetalleTransaccion(2).
			Build()).
		Build()

	xmlBytes, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	xmlBytes, err = utils.SignXML(xmlBytes, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error al firmar: %v", err)
	}

	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	req := models.NewRecepcionDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(24).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()
	// 1. EMITIR NOTA DE CREDITO DEBITO
	resp, err := tc.Client.DocumentoAjuste().RecepcionDocumentoAjuste(context.Background(), req)
	if err != nil {
		t.Fatalf("Nota %d - Error de red: %v", nroNota, err)
	}

	if resp.Body.Fault != nil {
		t.Errorf("Nota %d - Fault SIAT: %s", nroNota, resp.Body.Fault.String())
	} else {
		if resp.Body.Content.RespuestaRecepcionFactura.Transaccion {
			log.Printf("Nota %d - EXITOSA: %s", nroNota, resp.Body.Content.RespuestaRecepcionFactura.CodigoRecepcion)
		} else {
			mensajes := ""
			for _, m := range resp.Body.Content.RespuestaRecepcionFactura.MensajesList {
				mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
			}
			t.Errorf("Nota %d - RECHAZADA: %s", nroNota, mensajes)
		}
	}
	// Pequeño delay para que el SIAT procese el estado
	time.Sleep(50 * time.Millisecond)

	// 2. ANULAR FACTURA
	reqAnulacion := models.NewAnulacionDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(24).
		WithCodigoEmision(siat.EmisionOnline).
		WithTipoFacturaDocumento(3).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		// NIT es requerido en anulación
		WithCufd(cufd).
		WithCuis(cuis).
		WithCuf(cuf).
		WithCodigoMotivo(2). // 2: NOTA DE CREDITO-DEBITO MAL EMITIDA
		Build()

	respAnulacion, err := tc.Client.DocumentoAjuste().AnulacionDocumentoAjuste(context.Background(), reqAnulacion)
	if err != nil {
		t.Fatalf("error en anulación nota %d: %v", nroNota, err)
	}

	if !respAnulacion.Body.Content.RespuestaServicioFacturacion.Transaccion {
		mensajes := ""
		for _, m := range respAnulacion.Body.Content.RespuestaServicioFacturacion.MensajesList {
			mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
		}
		t.Errorf("Anulación Nota Credito debito %d falló: %s", nroNota, mensajes)
		return // No podemos revertir si no se anuló
	}
	log.Printf("Nota Credito debito %d ANULADA correctamente", nroNota)

	// Otro delay para la reversión
	time.Sleep(50 * time.Millisecond)

	// 3. REVERTIR ANULACIÓN
	reqReversion := models.NewReversionAnulacionDocumentoAjusteBuilder().
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		WithCodigoDocumentoSector(24).
		WithTipoFacturaDocumento(3).
		WithCodigoEmision(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCuis(cuis).
		Build()

	respReversion, err := tc.Client.DocumentoAjuste().ReversionAnulacionDocumentoAjuste(context.Background(), reqReversion)
	if err != nil {
		t.Fatalf("error en reversión Nota Credito debito %d: %v", nroNota, err)
	}

	if !respReversion.Body.Content.RespuestaServicioFacturacion.Transaccion {
		mensajes := ""
		for _, m := range respReversion.Body.Content.RespuestaServicioFacturacion.MensajesList {
			mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
		}
		t.Errorf("Reversión Nota Credito debito %d falló: %s", nroNota, mensajes)
		return
	}
	log.Printf("Nota Credito debito %d REVERTIDA (vuelve a ser válida)", nroNota)
}

func TestNotaCreditoDebito_ComputarizadaAll(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	for i := 1; i <= 125; i++ {
		t.Run(fmt.Sprintf("NotaComp_%d", i), func(t *testing.T) {
			emitirNotaComputarizadaIndividual(t, tc, cuis, cufd, cufdControl, i)
		})
	}
}

func emitirNotaComputarizadaIndividual(t *testing.T, tc *TestContext, cuis, cufd, cufdControl string, nroNota int) {
	fecha := time.Now()
	// Sector 47: Nota Crédito Débito Computarizada
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, tc.Sucursal, tc.Modalidad, siat.EmisionOnline, 3, 47, int64(nroNota), tc.PuntoVenta, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	nota := invoices.NewNotaCreditoDebitoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Empresa Test").
			WithMunicipio("La Paz").
			WithNumeroNotaCreditoDebito(int64(nroNota)).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(tc.Sucursal).
			WithCodigoPuntoVenta(&tc.PuntoVenta).
			WithDireccion("Av. Principal 123").
			WithFechaEmision(fecha).
			WithNombreRazonSocial(ptr("Cliente Test")).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5544332").
			WithCodigoCliente("CLI-001").
			WithNumeroFactura(100).
			WithNumeroAutorizacionCuf("DUMMY_AUT_CUF").
			WithFechaEmisionFactura(fecha.Add(-24 * time.Hour)).
			WithMontoTotalOriginal(1000.00).
			WithMontoTotalDevuelto(500.00).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P001").
			WithDescripcion("Producto A").
			WithCantidad(10).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(1000).
			WithCodigoDetalleTransaccion(1).
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("P002").
			WithDescripcion("Producto B devuelto").
			WithCantidad(5).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(500).
			WithCodigoDetalleTransaccion(2).
			Build()).
		Build()

	xmlBytes, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error marshal: %v", err)
	}

	// NO se firma para computarizada
	hash, encoded, err := utils.CompressAndHash(xmlBytes)
	if err != nil {
		t.Fatalf("error compress: %v", err)
	}

	req := models.NewRecepcionDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(47).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()
	// 1. EMITIR NOTA DE CREDITO DEBITO COMPUTARIZADA
	resp, err := tc.Client.DocumentoAjuste().RecepcionDocumentoAjuste(context.Background(), req)
	if err != nil {
		t.Fatalf("Nota %d - Error de red: %v", nroNota, err)
	}

	if resp.Body.Fault != nil {
		t.Errorf("Nota %d - Fault SIAT: %s", nroNota, resp.Body.Fault.String())
	} else {
		if resp.Body.Content.RespuestaRecepcionFactura.Transaccion {
			log.Printf("Nota Comp %d - EXITOSA: %s", nroNota, resp.Body.Content.RespuestaRecepcionFactura.CodigoRecepcion)
		} else {
			mensajes := ""
			for _, m := range resp.Body.Content.RespuestaRecepcionFactura.MensajesList {
				mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
			}
			t.Errorf("Nota Comp %d - RECHAZADA: %s", nroNota, mensajes)
		}
	}

	// Pequeño delay para que el SIAT procese el estado
	time.Sleep(50 * time.Millisecond)

	// 2. ANULAR FACTURA
	reqAnulacion := models.NewAnulacionDocumentoAjusteBuilder().
		WithCodigoDocumentoSector(23).
		WithCodigoEmision(siat.EmisionOnline).
		WithTipoFacturaDocumento(1).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		// NIT es requerido en anulación
		WithCufd(cufd).
		WithCuis(cuis).
		WithCuf(cuf).
		WithCodigoMotivo(1). // 1: Factura mal emitida
		Build()

	respAnulacion, err := tc.Client.DocumentoAjuste().AnulacionDocumentoAjuste(context.Background(), reqAnulacion)
	if err != nil {
		t.Fatalf("error en anulación nota credito debito %d: %v", nroNota, err)
	}

	if !respAnulacion.Body.Content.RespuestaServicioFacturacion.Transaccion {
		mensajes := ""
		for _, m := range respAnulacion.Body.Content.RespuestaServicioFacturacion.MensajesList {
			mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
		}
		t.Errorf("Anulación Nota Credito debito %d falló: %s", nroNota, mensajes)
		return // No podemos revertir si no se anuló
	}
	log.Printf("Nota Credito debito %d ANULADA correctamente", nroNota)

	// Otro delay para la reversión
	time.Sleep(50 * time.Millisecond)

	// 3. REVERTIR ANULACIÓN
	reqReversion := models.NewReversionAnulacionDocumentoAjusteBuilder().
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSucursal(tc.Sucursal).
		WithCodigoDocumentoSector(23).
		WithTipoFacturaDocumento(1).
		WithCodigoEmision(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCuis(cuis).
		Build()

	respReversion, err := tc.Client.DocumentoAjuste().ReversionAnulacionDocumentoAjuste(context.Background(), reqReversion)
	if err != nil {
		t.Fatalf("error en reversión Nota Credito debito %d: %v", nroNota, err)
	}

	if !respReversion.Body.Content.RespuestaServicioFacturacion.Transaccion {
		mensajes := ""
		for _, m := range respReversion.Body.Content.RespuestaServicioFacturacion.MensajesList {
			mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
		}
		t.Errorf("Reversión Nota Credito debito %d falló: %s", nroNota, mensajes)
		return
	}
	log.Printf("Nota Credito debito %d REVERTIDA (vuelve a ser válida)", nroNota)
}
