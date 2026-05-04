package invoices_test

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
)

func TestNotaCreditoDebitoDescuentoBuilder(t *testing.T) {
	fechaEmision := time.Date(2023, 10, 24, 15, 30, 0, 0, time.UTC)
	fechaEmisionFactura := time.Date(2023, 10, 20, 10, 0, 0, 0, time.UTC)

	telefono := "1234567"
	puntoVenta := 1
	nombreRazonSocial := "Juan Perez"
	complemento := "1A"
	descuentoAdicional := 10.0
	montoDescuentoCreditoDebito := 5.0
	codigoExcepcion := 1
	montoDescuento := 0.0

	cabecera := invoices.NewNotaCreditoDebitoDescuentoCabeceraBuilder().
		WithNitEmisor(123456789).
		WithRazonSocialEmisor("Mi Empresa S.A.").
		WithMunicipio("La Paz").
		WithTelefono(&telefono).
		WithNumeroNotaCreditoDebito(1001).
		WithCuf("12345ABCDE").
		WithCufd("ABCDE12345").
		WithCodigoSucursal(0).
		WithDireccion("Av. 16 de Julio").
		WithCodigoPuntoVenta(&puntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("9876543").
		WithComplemento(&complemento).
		WithCodigoCliente("CLI-001").
		WithNumeroFactura(500).
		WithNumeroAutorizacionCuf("AUTH-999").
		WithFechaEmisionFactura(fechaEmisionFactura).
		WithMontoTotalOriginal(1000.50).
		WithDescuentoAdicional(&descuentoAdicional).
		WithMontoTotalDevuelto(500.00).
		WithMontoDescuentoCreditoDebito(&montoDescuentoCreditoDebito).
		WithMontoEfectivoCreditoDebito(495.00).
		WithCodigoExcepcion(&codigoExcepcion).
		WithLeyenda("Leyenda de prueba").
		WithUsuario("admin").
		Build()

	detalle1 := invoices.NewNotaDetalleCreditoDebitoDescuentoBuilder().
		WithNroItem(1).
		WithActividadEconomica("123456").
		WithCodigoProductoSin(98765).
		WithCodigoProducto("PROD-1").
		WithDescripcion("Producto de prueba 1").
		WithCantidad(2.0).
		WithUnidadMedida(58).
		WithPrecioUnitario(100.0).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(200.0).
		WithCodigoDetalleTransaccion(1).
		Build()

	nota := invoices.NewNotaCreditoDebitoDescuentoBuilder().
		WithModalidad(siat.ModalidadElectronica).
		WithCabecera(cabecera).
		AddDetalle(detalle1).
		Build()

	xmlBytes, err := xml.MarshalIndent(nota, "", "  ")
	if err != nil {
		t.Fatalf("Error marshaling XML: %v", err)
	}

	xmlStr := string(xmlBytes)

	expectedTags := []string{
		"<notaElectronicaCreditoDebitoDescuento",
		"<cabecera>",
		"<nitEmisor>123456789</nitEmisor>",
		"<numeroNotaCreditoDebito>1001</numeroNotaCreditoDebito>",
		"<descuentoAdicional>10</descuentoAdicional>",
		"<detalle>",
		"<nroItem>1</nroItem>",
		"<actividadEconomica>123456</actividadEconomica>",
		"</detalle>",
	}

	for _, tag := range expectedTags {
		if !strings.Contains(xmlStr, tag) {
			t.Errorf("Expected XML to contain %s", tag)
		}
	}
}

func TestNotaCreditoDebitoDescuentoIntegration_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)

	service := tc.Client.DocumentoAjuste()

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

	descuentoAdicional := 10.0
	montoDescuentoCreditoDebito := 0.0

	// 4. Construir XML
	nota := invoices.NewNotaCreditoDebitoDescuentoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoDescuentoCabeceraBuilder().
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
			WithDescuentoAdicional(&descuentoAdicional).
			WithMontoTotalDevuelto(500.00).
			WithMontoDescuentoCreditoDebito(&montoDescuentoCreditoDebito).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoDescuentoBuilder().
			WithNroItem(1).
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
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoDescuentoBuilder().
			WithNroItem(2).
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
	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(47).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(tc.Nit).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()

	// 7. Intentar envío
	resp, err := service.RecepcionDocumentoAjuste(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("Error en la comunicación con el SIAT: %v", err)
	}

	t.Logf("Respuesta SIAT: %+v", resp.Body.Content.RespuestaRecepcionFactura)
}

func TestNotaCreditoDebitoDescuento_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	emitirNotaDescuentoIndividual(t, tc, cuis, cufd, cufdControl, 1)
}

func TestNotaCreditoDebitoDescuento_ElectronicaAll(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 1
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	for i := 1; i <= 125; i++ {
		t.Run(fmt.Sprintf("NotaDescuento_%d", i), func(t *testing.T) {
			emitirNotaDescuentoIndividual(t, tc, cuis, cufd, cufdControl, i)
			time.Sleep(50 * time.Millisecond)
		})
	}
}

func emitirNotaDescuentoIndividual(t *testing.T, tc *TestContext, cuis, cufd, cufdControl string, nroNota int) {
	fecha := time.Now()
	// Sector 47: Nota Crédito Débito Descuento, Tipo Doc 3
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, tc.Sucursal, tc.Modalidad, siat.EmisionOnline, 3, 47, nroNota, tc.PuntoVenta, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	descuentoAdicional := 10.0
	montoDescuentoCreditoDebito := 0.0

	nota := invoices.NewNotaCreditoDebitoDescuentoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoDescuentoCabeceraBuilder().
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
			WithDescuentoAdicional(&descuentoAdicional).
			WithMontoTotalDevuelto(500.00).
			WithMontoDescuentoCreditoDebito(&montoDescuentoCreditoDebito).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoDescuentoBuilder().
			WithNroItem(1).
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
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoDescuentoBuilder().
			WithNroItem(2).
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

	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(47).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(tc.Nit).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()
	// 1. EMITIR NOTA DE CREDITO DEBITO DESCUENTO
	resp, err := tc.Client.DocumentoAjuste().RecepcionDocumentoAjuste(context.Background(), tc.Config, req)
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
	reqAnulacion := models.DocumentoAjuste().NewAnulacionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(47).
		WithCodigoEmision(siat.EmisionOnline).
		WithTipoFacturaDocumento(3).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithNit(tc.Nit). // NIT es requerido en anulación
		WithCufd(cufd).
		WithCuis(cuis).
		WithCuf(cuf).
		WithCodigoMotivo(2). // 2: NOTA DE CREDITO-DEBITO MAL EMITIDA
		Build()

	respAnulacion, err := tc.Client.DocumentoAjuste().AnulacionDocumentoAjuste(context.Background(), tc.Config, reqAnulacion)
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
	reqReversion := models.DocumentoAjuste().NewReversionAnulacionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithNit(tc.Nit).
		WithCodigoDocumentoSector(47).
		WithTipoFacturaDocumento(3).
		WithCodigoEmision(1).
		WithCodigoModalidad(tc.Modalidad).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCuis(cuis).
		Build()

	respReversion, err := tc.Client.DocumentoAjuste().ReversionAnulacionDocumentoAjuste(context.Background(), tc.Config, reqReversion)
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

func TestNotaCreditoDebitoDescuento_ComputarizadaAll(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	for i := 1; i <= 125; i++ {
		t.Run(fmt.Sprintf("NotaDescuentoComp_%d", i), func(t *testing.T) {
			emitirNotaDescuentoComputarizadaIndividual(t, tc, cuis, cufd, cufdControl, i)
		})
	}
}

func emitirNotaDescuentoComputarizadaIndividual(t *testing.T, tc *TestContext, cuis, cufd, cufdControl string, nroNota int) {
	fecha := time.Now()
	// Sector 47: Nota Crédito Débito Descuento Computarizada
	cuf, err := utils.GenerarCUF(tc.Nit, fecha, tc.Sucursal, tc.Modalidad, siat.EmisionOnline, 3, 47, nroNota, tc.PuntoVenta, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	descuentoAdicional := 10.0
	montoDescuentoCreditoDebito := 0.0

	nota := invoices.NewNotaCreditoDebitoDescuentoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(invoices.NewNotaCreditoDebitoDescuentoCabeceraBuilder().
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
			WithDescuentoAdicional(&descuentoAdicional).
			WithMontoTotalDevuelto(500.00).
			WithMontoDescuentoCreditoDebito(&montoDescuentoCreditoDebito).
			WithMontoEfectivoCreditoDebito(65.00).
			WithLeyenda("Ley Nro 453").
			WithUsuario("admin").
			Build()).
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoDescuentoBuilder().
			WithNroItem(1).
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
		AddDetalle(invoices.NewNotaDetalleCreditoDebitoDescuentoBuilder().
			WithNroItem(2).
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

	req := models.DocumentoAjuste().NewRecepcionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(47).
		WithCodigoEmision(siat.EmisionOnline).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(tc.Nit).
		WithTipoFacturaDocumento(3).
		WithArchivo(encoded).
		WithFechaEnvio(fecha).
		WithHashArchivo(hash).
		Build()

	resp, err := tc.Client.DocumentoAjuste().RecepcionDocumentoAjuste(context.Background(), tc.Config, req)
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
	time.Sleep(500 * time.Millisecond)

	// 2. ANULAR FACTURA
	reqAnulacion := models.DocumentoAjuste().NewAnulacionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(47).
		WithCodigoEmision(siat.EmisionOnline).
		WithTipoFacturaDocumento(3).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithNit(tc.Nit). // NIT es requerido en anulación
		WithCufd(cufd).
		WithCuis(cuis).
		WithCuf(cuf).
		WithCodigoMotivo(2). // 2: NOTA DE CREDITO-DEBITO MAL EMITIDA
		Build()

	respAnulacion, err := tc.Client.DocumentoAjuste().AnulacionDocumentoAjuste(context.Background(), tc.Config, reqAnulacion)
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
	time.Sleep(500 * time.Millisecond)

	// 3. REVERTIR ANULACIÓN
	reqReversion := models.DocumentoAjuste().NewReversionAnulacionBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithNit(tc.Nit).
		WithCodigoDocumentoSector(47).
		WithTipoFacturaDocumento(3).
		WithCodigoEmision(1).
		WithCodigoModalidad(tc.Modalidad).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCuis(cuis).
		Build()

	respReversion, err := tc.Client.DocumentoAjuste().ReversionAnulacionDocumentoAjuste(context.Background(), tc.Config, reqReversion)
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
