package invoices_test

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
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

func TestBoletoAereoBuilder(t *testing.T) {
	fechaEmision := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	nombre := "JUAN PEREZ"

	cabecera := invoices.NewBoletoAereoCabeceraBuilder().
		WithNitEmisor(1234567).
		WithRazonSocialEmisor("AEROLINEAS PRUEBA").
		WithNumeroFactura(1).
		WithCuf("ABC123CUF").
		WithCufd("XYZ789CUFD").
		WithCodigoSucursal(0).
		WithDireccion("AV. PRINCIPAL 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombre).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithNombrePasajero("JUAN PEREZ").
		WithCodigoIataLineaAerea(900).
		WithCodigoOrigenServicio("VVI").
		WithCodigoMetodoPago(1).
		WithMontoTarifa(500.00).
		WithMontoTotal(550.00).
		WithMontoTotalSujetoIva(500.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(550.00).
		WithCodigoTipoTransaccion("I").
		WithLeyenda("Leyenda de prueba").
		WithUsuario("tester").
		Build()

	t.Run("Modalidad Electronica", func(t *testing.T) {
		factura := invoices.NewBoletoAereoBuilder().
			WithModalidad(siat.ModalidadElectronica).
			WithCabecera(cabecera).
			Build()

		internal := models.UnwrapInternalRequest[documents.FacturaBoletoAereo](factura)
		output, _ := xml.Marshal(internal)
		xmlStr := string(output)

		assert.Contains(t, xmlStr, "facturaElectronicaBoletoAereo")
		assert.Contains(t, xmlStr, "<nombrePasajero>JUAN PEREZ</nombrePasajero>")
		assert.Contains(t, xmlStr, "<codigoDocumentoSector>30</codigoDocumentoSector>")
	})
}

func TestBoletoAereoIntegration(t *testing.T) {
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
			service := tc.Client.BoletoAereo()
			cuis := tc.GetCuis(t)
			cufd, cufdControl := tc.GetCufd(t, cuis)

			nombre := "JUAN PEREZ"

			factura := invoices.NewBoletoAereoBuilder().
				WithModalidad(tc.Modalidad).
				WithCabecera(invoices.NewBoletoAereoCabeceraBuilder().
					WithNitEmisor(tc.Nit).
					WithRazonSocialEmisor("AEROLINEAS PRUEBA S.A.").
					WithNumeroFactura(1).
					WithCuf(tc.GetCuf(t, 30, 3, 4, 1, 0, cufdControl)).
					WithCufd(cufd).
					WithCodigoSucursal(tc.Sucursal).
					WithDireccion("Av. Principal 123").
					WithFechaEmision(time.Now()).
					WithNombreRazonSocial(&nombre).
					WithCodigoTipoDocumentoIdentidad(1).
					WithNumeroDocumento("1234567").
					WithNombrePasajero("JUAN PEREZ").
					WithCodigoIataLineaAerea(900).
					WithCodigoOrigenServicio("VVI").
					WithCodigoMetodoPago(1).
					WithMontoTarifa(500.00).
					WithMontoTotal(500.00).
					WithMontoTotalSujetoIva(500.00).
					WithCodigoMoneda(1).
					WithTipoCambio(1).
					WithMontoTotalMoneda(500.00).
					WithCodigoTipoTransaccion("I").
					WithLeyenda("Ley 453...").
					WithUsuario("operador").
					Build()).
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
				processedXML = xmlData
			}

			// Boleto Aereo usa envio masivo (se requiere comprimir en tar.gz)
			tarGz, err := utils.CreateTarGz(map[string][]byte{"factura.xml": processedXML})
			if err != nil {
				t.Fatalf("Error creando tar.gz: %v", err)
			}

			hashSum := sha256.Sum256(tarGz)
			hashString := fmt.Sprintf("%x", hashSum)
			encodedArchivo := base64.StdEncoding.EncodeToString(tarGz)

			req := models.BoletoAereo().NewRecepcionMasivaFacturaBuilder().
				WithCodigoAmbiente(tc.Ambiente).
				WithCodigoDocumentoSector(30).
				WithCodigoEmision(3).
				WithCodigoModalidad(tc.Modalidad).
				WithCodigoPuntoVenta(tc.PuntoVenta).
				WithCodigoSistema(tc.Sistema).
				WithCodigoSucursal(tc.Sucursal).
				WithCufd(cufd).
				WithCuis(cuis).
				WithNit(tc.Nit).
				WithTipoFacturaDocumento(4).
				WithArchivo(encodedArchivo).
				WithFechaEnvio(time.Now()).
				WithHashArchivo(hashString).
				WithCantidadFacturas(1).
				Build()

			resp, err := service.RecepcionMasivaFactura(context.Background(), tc.Config, req)

			if err == nil && resp != nil {
				log.Printf("Respuesta Recepcion Masiva Boleto Aereo (%s): %+v", name, resp.Body.Content.RespuestaServicioFacturacion)
			} else if err != nil {
				t.Fatalf("Error en comunicación con SIAT (%s): %v", name, err)
			}
		})
	}
}
