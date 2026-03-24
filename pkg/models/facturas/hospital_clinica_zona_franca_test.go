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

func TestHospitalClinicaZonaFranca_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "Juan Perez"
	especialidad := "Cardiologia"
	especialidadDetalle := "Pediatrica"
	especialidadMedico := "Cardiologo Pediatra"
	nroMatriculaMedico := "12345"
	nroFacturaMedico := 123
	montoDescuento := 0.0
	cantidad := 1.0
	precioUnitario := 100.0
	subTotalItem := (cantidad * precioUnitario) - montoDescuento
	montoTotal := subTotalItem

	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 50, 1, 0, cufdControl)

	cabecera := facturas.NewHospitalClinicaZFCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("HOSPITAL ZF TEST").
		WithMunicipio("SCZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. MEDICINA 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("CLI-ZF-01").
		WithCodigoMetodoPago(1).
		WithMontoTotal(montoTotal).
		WithMontoTotalSujetoIva(0).
		WithCodigoMoneda(1).
		WithTipoCambio(1).
		WithMontoTotalMoneda(montoTotal).
		WithLeyenda("Leyenda Factura").
		WithUsuario("operador01").
		Build()

	detalle1 := facturas.NewHospitalClinicaZonaFrancaDetalleBuilder().
		WithActividadEconomica("861000").
		WithCodigoProductoSin(12345).
		WithCodigoProducto("PROD-01").
		WithDescripcion("CONSULTA CARDIOLOGIA").
		WithEspecialidad(&especialidad).
		WithEspecialidadDetalle(&especialidadDetalle).
		WithNroQuirofanoSalaOperaciones(1).
		WithEspecialidadMedico(&especialidadMedico).
		WithNombreApellidoMedico("DR. SMITH").
		WithNitDocumentoMedico(987654321).
		WithNroMatriculaMedico(&nroMatriculaMedico).
		WithNroFacturaMedico(&nroFacturaMedico).
		WithCantidad(cantidad).
		WithUnidadMedida(58).
		WithPrecioUnitario(precioUnitario).
		WithMontoDescuento(&montoDescuento).
		WithSubTotal(subTotalItem).
		Build()

	factura := facturas.NewHospitalClinicaZFBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle1).
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
		WithCodigoDocumentoSector(50).
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
