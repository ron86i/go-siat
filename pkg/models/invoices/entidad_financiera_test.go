package invoices_test

import (
	"context"
	"encoding/xml"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestEntidadFinancieraBuilder(t *testing.T) {
	fechaEmision := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	telefono := "2222222"
	nombre := "JUAN PEREZ"
	montoArrendamiento := 1000.50

	cabecera := invoices.NewEntidadFinancieraCabeceraBuilder().
		WithNitEmisor(1234567).
		WithRazonSocialEmisor("BANCO PRUEBA").
		WithMunicipio("LA PAZ").
		WithTelefono(&telefono).
		WithNumeroFactura(1).
		WithCuf("ABC123CUF").
		WithCufd("XYZ789CUFD").
		WithCodigoSucursal(0).
		WithDireccion("AV. CAMACHO 456").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-001").
		WithCodigoMetodoPago(1).
		WithMontoTotalArrendamientoFinanciero(&montoArrendamiento).
		WithMontoTotal(5000.50).
		WithMontoTotalSujetoIva(5000.50).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(5000.50).
		WithLeyenda("Leyenda Banco").
		WithUsuario("operador1").
		Build()

	detalle := invoices.NewEntidadFinancieraDetalleBuilder().
		WithActividadEconomica("641911").
		WithCodigoProductoSin(456).
		WithCodigoProducto("P001").
		WithDescripcion("INTERESES PRESTAMO").
		WithCantidad(1).
		WithUnidadMedida(58).
		WithPrecioUnitario(5000.50).
		WithSubTotal(5000.50).
		Build()

	t.Run("Modalidad Electronica", func(t *testing.T) {
		factura := invoices.NewEntidadFinancieraBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		internal := models.UnwrapInternalRequest[documents.FacturaEntidadFinanciera](factura)
		output, _ := xml.Marshal(internal)
		xmlStr := string(output)

		assert.Contains(t, xmlStr, "facturaElectronicaEntidadFinanciera")
		assert.Contains(t, xmlStr, "<montoTotalArrendamientoFinanciero>1000.5</montoTotalArrendamientoFinanciero>")
		assert.Contains(t, xmlStr, "<codigoDocumentoSector>15</codigoDocumentoSector>")
	})
}

func TestEntidadFinancieraIntegration(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}

	modalidades := []int{siat.ModalidadElectronica, siat.ModalidadComputarizada}

	for _, mod := range modalidades {
		name := "Electronica"
		if mod == siat.ModalidadComputarizada {
			name = "Computarizada"
		}

		t.Run(name, func(t *testing.T) {
			tc := setupTestContext(t, mod)
			service := tc.Client.EntidadFinanciera()
			cuis := tc.GetCuis(t)
			cufd, cufdControl := tc.GetCufd(t, cuis)

			nombre := "JUAN PEREZ"
			montoArrendamiento := 500.00

			detalle := invoices.NewEntidadFinancieraDetalleBuilder().
				WithActividadEconomica("641911").
				WithCodigoProductoSin(456).
				WithCodigoProducto("P001").
				WithDescripcion("Comision Bancaria").
				WithCantidad(1).
				WithUnidadMedida(58).
				WithPrecioUnitario(100).
				WithSubTotal(100).
				Build()

			montoTotal := 100.0 + montoArrendamiento

			factura := invoices.NewEntidadFinancieraBuilder().
				WithModalidad(tc.Modalidad).
				WithCabecera(invoices.NewEntidadFinancieraCabeceraBuilder().
					WithNitEmisor(tc.Nit).
					WithRazonSocialEmisor("BANCO PRUEBA S.A.").
					WithMunicipio("La Paz").
					WithNumeroFactura(1).
					WithCuf(tc.GetCuf(t, 15, 1, 1, 1, 0, cufdControl)).
					WithCufd(cufd).
					WithCodigoSucursal(tc.Sucursal).
					WithDireccion("Av. Camacho 123").
					WithFechaEmision(time.Now()).
					WithNombreRazonSocial(&nombre).
					WithCodigoTipoDocumentoIdentidad(1).
					WithNumeroDocumento("1234567").
					WithCodigoCliente("CLI-001").
					WithCodigoMetodoPago(1).
					WithMontoTotalArrendamientoFinanciero(&montoArrendamiento).
					WithMontoTotal(montoTotal).
					WithMontoTotalSujetoIva(100).
					WithCodigoMoneda(1).
					WithTipoCambio(1).
					WithMontoTotalMoneda(montoTotal).
					WithLeyenda("Ley 453...").
					WithUsuario("operador").
					Build()).
				AddDetalle(detalle).
				Build()

			xmlData, _ := xml.Marshal(factura)
			var processedXML []byte

			if mod == siat.ModalidadElectronica {
				signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
				if err != nil {
					t.Skip("Saltando envío Electronica: Error firmando (probablemente faltan certificados):", err)
					return
				}
				processedXML = signedXML
			} else {
				// Computarizada no se firma
				processedXML = xmlData
			}

			hashString, encodedArchivo, _ := utils.CompressAndHash(processedXML)

			req := models.EntidadFinanciera().NewRecepcionFacturaBuilder().
				WithCodigoAmbiente(tc.Ambiente).
				WithCodigoDocumentoSector(15).
				WithCodigoEmision(siat.EmisionOnline).
				WithCodigoModalidad(tc.Modalidad).
				WithCodigoPuntoVenta(tc.PuntoVenta).
				WithCodigoSistema(tc.Sistema).
				WithCodigoSucursal(tc.Sucursal).
				WithCufd(cufd).
				WithCuis(cuis).
				WithNit(tc.Nit).
				WithTipoFacturaDocumento(1).
				WithArchivo(encodedArchivo).
				WithFechaEnvio(time.Now()).
				WithHashArchivo(hashString).
				Build()

			resp, err := service.RecepcionFactura(context.Background(), tc.Config, req)

			if err == nil && resp != nil {
				log.Printf("Respuesta Recepcion Entidad Financiera (%s): %+v", name, resp.Body.Content)
			}
		})
	}
}
