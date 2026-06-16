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

func TestSeguros_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 1
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	emitirSegurosIndividual(t, tc, cuis, cufd, cufdControl, 1)
}

func TestSeguros_ElectronicaAll(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	tc.PuntoVenta = 1
	tc.Sucursal = 0
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	for i := 1; i <= 125; i++ {
		t.Run(fmt.Sprintf("FacturaSeguros_%d", i), func(t *testing.T) {
			emitirSegurosIndividual(t, tc, cuis, cufd, cufdControl, i)
			time.Sleep(50 * time.Millisecond)
		})
	}
}

func emitirSegurosIndividual(t *testing.T, tc *TestContext, cuis, cufd, cufdControl string, nroFactura int) {
	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	// Sector 34: Seguros
	cuf, err := utils.GenerarCUF(tc.Nit, fechaEmision, tc.Sucursal, tc.Modalidad, 1, 1, 34, int64(nroFactura), tc.PuntoVenta, cufdControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	cabecera := invoices.NewSegurosCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("COMPAÑIA DE SEGUROS S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(int64(nroFactura)).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(tc.Sucursal).
		WithCodigoPuntoVenta(&tc.PuntoVenta).
		WithDireccion("AV. ARCE 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-034").
		WithCodigoMetodoPago(1).
		WithMontoTotal(1000.00).
		WithAjusteAfectacionIva(0.00).
		WithMontoTotalSujetoIva(1000.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(1000.00).
		WithLeyenda("Leyenda Seguros").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewSegurosDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(35270).
		WithCodigoProducto("SEG-001").
		WithDescripcion("SEGURO DE VIDA INDIVIDUAL").
		WithCantidad(1.0).
		WithUnidadMedida(58).
		WithPrecioUnitario(1000.00).
		WithSubTotal(1000.00).
		Build()

	factura := invoices.NewSegurosBuilder().
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
		WithCodigoSucursal(tc.Sucursal).
		WithCodigoDocumentoSector(34).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()
	// 1. EMITIR FACTURA SEGUROS ELECTRONICA
	resp, err := tc.Client.Electronica().RecepcionFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("Factura %d - error en solicitud: %v", nroFactura, err)
	}

	if resp.Body.Fault != nil {
		t.Errorf("Factura %d - Fault SIAT: %s", nroFactura, resp.Body.Fault.String())
	} else {
		if resp.Body.Content.RespuestaServicioFacturacion.Transaccion {
			log.Printf("Factura %d - EXITOSA: %s", nroFactura, resp.Body.Content.RespuestaServicioFacturacion.CodigoRecepcion)
		} else {
			mensajes := ""
			for _, m := range resp.Body.Content.RespuestaServicioFacturacion.MensajesList {
				mensajes += fmt.Sprintf("[%d: %s] ", m.Codigo, m.Descripcion)
			}
			t.Errorf("Factura %d - RECHAZADA: %s", nroFactura, mensajes)
		}
	}
	// Pequeño delay para que el SIAT procese el estado
	time.Sleep(50 * time.Millisecond)

	// 2. ANULAR FACTURA
	reqAnulacion := models.Electronica().NewAnulacionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(34).
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

	respAnulacion, err := tc.Client.Electronica().AnulacionFactura(context.Background(), tc.Config, reqAnulacion)
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
	reqReversion := models.Electronica().NewReversionAnulacionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoPuntoVenta(tc.PuntoVenta).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(tc.Sucursal).
		WithNit(tc.Nit).
		WithCodigoDocumentoSector(34).
		WithTipoFacturaDocumento(1).
		WithCodigoEmision(1).
		WithCodigoModalidad(tc.Modalidad).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCuis(cuis).
		Build()

	respReversion, err := tc.Client.Electronica().ReversionAnulacionFactura(context.Background(), tc.Config, reqReversion)
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
