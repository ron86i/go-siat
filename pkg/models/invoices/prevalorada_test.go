package invoices_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
)

func TestPrevalorada_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	emitirPrevaloradaIndividual(t, tc, cuis, cufd, cufdControl, 1)
}

func TestPrevalorada_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	emitirPrevaloradaIndividual(t, tc, cuis, cufd, cufdControl, 1)
}

func TestPrevalorada_ElectronicaAll(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 0
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)
	for i := 1; i <= 125; i++ {
		t.Run(fmt.Sprintf("Factura_%d", i), func(t *testing.T) {
			emitirPrevaloradaIndividual(t, tc, cuis, cufd, cufdControl, i)
		})
	}
}

func emitirPrevaloradaIndividual(t *testing.T, tc *TestContext, cuis, cufd, cufdControl string, nroFactura int) {
	fechaEmision := time.Now()
	// Sector 23: Prevalorada
	cuf, err := utils.GenerarCUF(tc.Nit, fechaEmision, tc.Sucursal, tc.Modalidad, 1, 1, 23, int64(nroFactura), tc.PuntoVenta, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	cantidad := 1.0
	precioUnitario := 10.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	cabecera := invoices.NewPrevaloradaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("TELEFONICA S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(int64(nroFactura)).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. BALLIVIAN").
		WithCodigoPuntoVenta(&tc.PuntoVenta).
		WithFechaEmision(fechaEmision).
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Leyenda Prevalorada").
		WithUsuario("usuario").
		Build()

	detalle := invoices.NewPrevaloradaDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(99100).
		WithCodigoProducto("abc").
		WithDescripcion("PRODUCTO").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := invoices.NewPrevaloradaBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		WithDetalle(detalle).
		Build()

	if tc.Modalidad == siat.ModalidadElectronica {
		builderReq := models.NewRecepcionFacturaBuilder().
			WithCodigoModalidad(tc.Modalidad).
			WithCodigoSucursal(tc.Sucursal).
			WithCodigoDocumentoSector(23).
			WithCodigoEmision(1).
			WithCodigoPuntoVenta(tc.PuntoVenta).
			WithCufd(cufd).
			WithCuis(cuis).
			WithTipoFacturaDocumento(1).
			WithFechaEnvio(fechaEmision)

		err := builderReq.WithFactura(factura, tc.Client.Config())
		if err != nil {
			t.Fatalf("error al preparar factura: %v", err)
		}

		req := builderReq.Build()
		// 1. EMITIR PREVALORADA ELECTRONICA
		resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), req)
		procesarRespuestaSIAT(t, nroFactura, resp, err)

		// Pequeño delay para que el SIAT procese el estado
		time.Sleep(50 * time.Millisecond)

		// 2. ANULAR FACTURA
		reqAnulacion := models.NewAnulacionFacturaBuilder().
			WithCodigoDocumentoSector(23).
			WithCodigoEmision(siat.EmisionOnline).
			WithTipoFacturaDocumento(1).
			WithCodigoModalidad(tc.Modalidad).
			WithCodigoPuntoVenta(tc.PuntoVenta).
			WithCodigoSucursal(tc.Sucursal).
			WithCufd(cufd).
			WithCuis(cuis).
			WithCuf(cuf).
			WithCodigoMotivo(1). // 1: Factura mal emitida
			Build()

		respAnulacion, err := tc.Client.Electronica().AnulacionFactura(context.Background(), reqAnulacion)
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
		reqReversion := models.NewReversionAnulacionFacturaBuilder().
			WithCodigoPuntoVenta(tc.PuntoVenta).
			WithCodigoSucursal(tc.Sucursal).
			WithCodigoDocumentoSector(23).
			WithTipoFacturaDocumento(1).
			WithCodigoEmision(1).
			WithCodigoModalidad(tc.Modalidad).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCuis(cuis).
			Build()

		respReversion, err := tc.Client.Electronica().ReversionAnulacionFactura(context.Background(), reqReversion)
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
	} else {
		builderReq := models.NewRecepcionFacturaBuilder().
			WithCodigoModalidad(tc.Modalidad).
			WithCodigoSucursal(tc.Sucursal).
			WithCodigoDocumentoSector(23).
			WithCodigoEmision(1).
			WithCodigoPuntoVenta(tc.PuntoVenta).
			WithCufd(cufd).
			WithCuis(cuis).
			WithTipoFacturaDocumento(1).
			WithFechaEnvio(fechaEmision)

		err := builderReq.WithFactura(factura, tc.Client.Config())
		if err != nil {
			t.Fatalf("error al preparar factura: %v", err)
		}

		req := builderReq.Build()

		resp, err := tc.Client.Computarizada().RecepcionFactura(context.Background(), req)
		procesarRespuestaSIAT(t, nroFactura, resp, err)
	}
}

func procesarRespuestaSIAT[T any](t *testing.T, nro int, resp *soap.EnvelopeResponse[T], err error) {
	if err != nil {
		t.Fatalf("Factura %d - error de red/SOAP: %v", nro, err)
	}

	if resp.Body.Fault != nil {
		t.Errorf("Factura %d - Fault SIAT: %s", nro, resp.Body.Fault.String())
		return
	}

	log.Printf("Factura %d - Respuesta: %+v", nro, resp.Body.Content)
}
