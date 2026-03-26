package invoices_test

import (
	"context"
	"encoding/xml"
	"log"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSuministroEnergia_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "JUAN PEREZ"

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 31, 1, 0, cufdControl)

	cabecera := invoices.NewSuministroEnergiaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("EMPRESA ELECTRICA").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. BLANCO GALINDO").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-ELEC-01").
		WithPlacaVehiculo("000000").
		WithCodigoMetodoPago(1).
		WithMontoTotal(150.50).
		WithMontoTotalSujetoIva(150.50).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(150.50).
		WithLeyenda("Leyenda Energía").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewSuministroEnergiaDetalleBuilder().
		WithActividadEconomica("351000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("ELEC-001").
		WithDescripcion("CONSUMO DE ENERGIA ELECTRICA").
		WithCantidad(1.0).
		WithUnidadMedida(58).
		WithPrecioUnitario(150.5).
		WithSubTotal(150.50).
		Build()

	factura := invoices.NewSuministroEnergiaBuilder().
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
		WithCodigoDocumentoSector(31).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
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
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}
