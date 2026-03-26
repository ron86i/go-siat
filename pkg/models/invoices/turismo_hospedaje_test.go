package invoices_test

import (
	"context"
	"encoding/xml"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestTurismoHospedaje_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "TURISTA EXTRANJERO"

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 2, 6, 1, 0, cufdControl)

	cabecera := invoices.NewTurismoHospedajeCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("HOTEL TURISTICO").
		WithMunicipio("UYUNI").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("SALAR DE UYUNI").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("99887766").
		WithCodigoCliente("TUR-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(500.00).
		WithCodigoMoneda(2). // Dólares
		WithTipoCambio(6.96).
		WithMontoTotalMoneda(71.84).
		WithLeyenda("Leyenda Turismo").
		WithUsuario("operador01").
		Build()

	detalle := invoices.NewTurismoHospedajeDetalleBuilder().
		WithActividadEconomica("551010").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("HOTEL-001").
		WithDescripcion("HOSPEDAJE NOCHE").
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(500.0).
		WithSubTotal(500.00).
		Build()

	factura := invoices.NewTurismoHospedajeBuilder().
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
		WithCodigoDocumentoSector(6).
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
