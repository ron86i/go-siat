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

func TestSectorEducativo_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 11, 1, 0, cufdControl)

	nombreRazonSocial := "PADRE DE FAMILIA"
	codigoPuntoVenta := 0
	cantidad := 1.0
	precioUnitario := 1000.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	cabecera := invoices.NewSectorEducativoCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("COLEGIO PRUEBA").
		WithMunicipio("TARIJA").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("CALLE ESTUDIANTES").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-EDU-01").
		WithNombreEstudiante("JUANITO PEREZ").
		WithPeriodoFacturado("MARZO 2024").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Leyenda Educación").
		WithUsuario("usuario").
		Build()

	detalle := invoices.NewSectorEducativoDetalleBuilder().
		WithActividadEconomica("851000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("EDU-01").
		WithDescripcion("PENSION MARZO").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := invoices.NewSectorEducativoBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, _ := xml.Marshal(factura)
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
		WithCodigoDocumentoSector(11).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
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
