package invoices_test

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/models/invoices"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestCompraVenta_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 1, 1, 0, cufdControl)

	nombreRazonSocial := "JUAN PEREZ"
	codigoPuntoVenta := 0

	cabecera := invoices.NewCompraVentaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("Ronaldo Rua").
		WithMunicipio("Tarija").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AVENIDA LA PAZ").
		WithCodigoPuntoVenta(&codigoPuntoVenta).
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("5115889").
		WithCodigoCliente("1").
		WithCodigoMetodoPago(1).
		WithMontoTotal(100).
		WithMontoTotalSujetoIva(100).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(100).
		WithLeyenda("Ley N° 453: Tienes derecho a recibir información...").
		WithUsuario("usuario").
		WithCodigoDocumentoSector(1).
		Build()

	detalle := invoices.NewCompraVentaDetalleBuilder().
		WithActividadEconomica("477300").
		WithCodigoProductoSin(622539).
		WithCodigoProducto("abc123").
		WithDescripcion("GASA").
		WithCantidad(1).
		WithUnidadMedida(1).
		WithPrecioUnitario(100).
		WithSubTotal(100).
		Build()

	factura := invoices.NewCompraVentaBuilder().
		WithModalidad(tc.Modalidad).
		WithCabecera(cabecera).
		AddDetalle(detalle).
		Build()

	xmlData, _ := xml.Marshal(factura)
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error firmando XML: %v", err)
	}
	log.Println(string(xmlData))
	hashString, encodedArchivo, err := utils.CompressAndHash(signedXML)
	if err != nil {
		t.Fatalf("error preparando archivo: %v", err)
	}

	req := models.CompraVenta().NewRecepcionFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoSistema(tc.Sistema).
		WithNit(tc.Nit).
		WithCodigoSucursal(0).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(1).
		WithCodigoPuntoVenta(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		Build()

	resp, err := tc.Client.CompraVenta().RecepcionFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}

func TestCompraVenta_Masiva(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)
	fechaEmision := time.Now()
	codigoPuntoVenta := 0
	for i := 1; i <= 5; i++ {
		cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, siat.EmisionMasiva, 1, 1, i, 0, cufdControl)
		nombreRazonSocial := "JUAN PEREZ"
		cabecera := invoices.NewCompraVentaCabeceraBuilder().
			WithNitEmisor(tc.Nit).
			WithRazonSocialEmisor("Ronaldo Rua").
			WithMunicipio("Tarija").
			WithNumeroFactura(int64(i)).
			WithCuf(cuf).
			WithCufd(cufd).
			WithCodigoSucursal(0).
			WithDireccion("AVENIDA LA PAZ").
			WithCodigoPuntoVenta(&codigoPuntoVenta).
			WithFechaEmision(fechaEmision).
			WithNombreRazonSocial(&nombreRazonSocial).
			WithCodigoTipoDocumentoIdentidad(1).
			WithNumeroDocumento("5115889").
			WithCodigoCliente(strconv.Itoa(i)).
			WithCodigoMetodoPago(1).
			WithMontoTotal(100).
			WithMontoTotalSujetoIva(100).
			WithCodigoMoneda(1).
			WithTipoCambio(1).
			WithMontoTotalMoneda(100).
			WithLeyenda("Ley N° 453").
			WithUsuario("usuario").
			WithCodigoDocumentoSector(1).
			Build()

		detalle := invoices.NewCompraVentaDetalleBuilder().
			WithActividadEconomica("477300").
			WithCodigoProductoSin(622539).
			WithCodigoProducto("abc123").
			WithDescripcion("GASA").
			WithCantidad(1).
			WithUnidadMedida(1).
			WithPrecioUnitario(100).
			WithSubTotal(100).
			Build()

		factura := invoices.NewCompraVentaBuilder().
			WithModalidad(tc.Modalidad).
			WithCabecera(cabecera).
			AddDetalle(detalle).
			Build()

		xmlData, _ := xml.Marshal(factura)
		signedXML, _ := utils.SignXML(xmlData, "key.pem", "cert.crt")

		hdr := &tar.Header{
			Name: fmt.Sprintf("factura_%d.xml", i),
			Mode: 0600,
			Size: int64(len(signedXML)),
		}
		tw.WriteHeader(hdr)
		tw.Write(signedXML)
	}
	tw.Close()

	hashString, encodedArchivo, err := utils.CompressAndHash(tarBuf.Bytes())
	if err != nil {
		t.Fatalf("error preparando paquete masivo: %v", err)
	}

	req := models.CompraVenta().NewRecepcionMasivaFacturaBuilder().
		WithCodigoAmbiente(tc.Ambiente).
		WithCodigoDocumentoSector(1).
		WithCodigoEmision(siat.EmisionMasiva).
		WithCodigoModalidad(tc.Modalidad).
		WithCodigoPuntoVenta(0).
		WithCodigoSistema(tc.Sistema).
		WithCodigoSucursal(0).
		WithCufd(cufd).
		WithCuis(cuis).
		WithNit(tc.Nit).
		WithTipoFacturaDocumento(1).
		WithArchivo(encodedArchivo).
		WithFechaEnvio(fechaEmision).
		WithHashArchivo(hashString).
		WithCantidadFacturas(5).
		Build()

	resp, err := tc.Client.CompraVenta().RecepcionMasivaFactura(context.Background(), tc.Config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}
	assert.Nil(t, resp.Body.Fault)
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}
