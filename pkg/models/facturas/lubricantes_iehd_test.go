package facturas_test

import (
	"context"
	"encoding/xml"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/facturas"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestLubricantesIehd_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "CLIENTE LUBRI"
	codigoPuntoVenta := 0
	cantidad := 1.0
	precioUnitario := 100.0
	montoDescuento := 0.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 53, 1, codigoPuntoVenta, cufdControl)

	zero := 0.0
	ciudad := "LA PAZ"
	propietario := "JUAN PEREZ"
	representante := "JUAN PEREZ"
	condicion := "CONTADO"
	periodo := "MARZO 2024"

	cabecera := facturas.NewLubricantesIehdCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("LUBRICANTES S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. LUBRICANTES").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-LUB-01").
		WithCiudad(&ciudad).
		WithNombrePropietario(&propietario).
		WithNombreRepresentanteLegal(&representante).
		WithCondicionPago(&condicion).
		WithPeriodoEntrega(&periodo).
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoDeduccionIehdDS25530(&zero).
		WithMontoTotalSujetoIva(montoTotal).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Leyenda Lubricantes").
		WithUsuario("operador01").
		Build()

	detalle := facturas.NewLubricantesIehdDetalleBuilder().
		WithActividadEconomica("466100").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("LUB-01").
		WithDescripcion("LUBRICANTE IEHD").
		WithCantidad(cantidad).
		WithUnidadMedida(1).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		WithCantidadLitros(1.0).
		WithPorcentajeDeduccionIehdDS25530(&zero).
		Build()

	factura := facturas.NewLubricantesIehdBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
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
		WithCodigoDocumentoSector(53).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(codigoPuntoVenta).
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
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
