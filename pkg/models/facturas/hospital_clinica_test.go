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

func TestHospitalClinica_Electronica(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadElectronica)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "PACIENTE EJEMPLO"
	modalidadServicio := "INTERNACION"

	// Sector 17 = Hospital/Clinica
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 17, 1, 0, cufdControl)

	cabecera := facturas.NewHospitalClinicaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("CLINICA BOLIVIA S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(1).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. ARCE 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("PAC-001").
		WithModalidadServicio(&modalidadServicio).
		WithCodigoMetodoPago(1).
		WithMontoTotal(2500.00).
		WithMontoTotalSujetoIva(2500.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(2500.00).
		WithLeyenda("Leyenda Hospital").
		WithUsuario("admin").
		Build()

	especialidad := "CIRUGIA GENERAL"
	detalle := facturas.NewHospitalClinicaDetalleBuilder().
		WithActividadEconomica("851100").
		WithCodigoProductoSin(85111).
		WithCodigoProducto("SERV-001").
		WithDescripcion("APENDICECTOMIA").
		WithEspecialidad(&especialidad).
		WithNroQuirofanoSalaOperaciones(1).
		WithNombreApellidoMedico("DR. PEREZ").
		WithNitDocumentoMedico(1234567).
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(2500.0).
		WithSubTotal(2500.0).
		Build()

	factura := facturas.NewHospitalClinicaBuilder().
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
		WithCodigoDocumentoSector(17).
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
	t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
}

func TestHospitalClinica_Computarizada(t *testing.T) {
	tc := setupTestContext(t, siat.ModalidadComputarizada)
	cuis := tc.GetCuis(t)
	cufd, cufdControl := tc.GetCufd(t, cuis)

	fechaEmision := time.Now()
	nombreRazonSocial := "PACIENTE EJEMPLO"

	// Sector 17 = Hospital/Clinica
	cuf, _ := utils.GenerarCUF(tc.Nit, fechaEmision, 0, tc.Modalidad, 1, 1, 17, 2, 0, cufdControl)

	cabecera := facturas.NewHospitalClinicaCabeceraBuilder().
		WithNitEmisor(tc.Nit).
		WithRazonSocialEmisor("CLINICA BOLIVIA S.A.").
		WithMunicipio("LA PAZ").
		WithNumeroFactura(2).
		WithCuf(cuf).
		WithCufd(cufd).
		WithCodigoSucursal(0).
		WithDireccion("AV. ARCE 123").
		WithFechaEmision(fechaEmision).
		WithNombreRazonSocial(&nombreRazonSocial).
		WithCodigoTipoDocumentoIdentidad(1).
		WithNumeroDocumento("1234567").
		WithCodigoCliente("PAC-001").
		WithCodigoMetodoPago(1).
		WithMontoTotal(150.00).
		WithMontoTotalSujetoIva(150.00).
		WithCodigoMoneda(1).
		WithTipoCambio(1.0).
		WithMontoTotalMoneda(150.00).
		WithLeyenda("Leyenda Hospital").
		WithUsuario("admin").
		Build()

	detalle := facturas.NewHospitalClinicaDetalleBuilder().
		WithActividadEconomica("851100").
		WithCodigoProductoSin(85112).
		WithCodigoProducto("SERV-002").
		WithDescripcion("CONSULTA MEDICA").
		WithNroQuirofanoSalaOperaciones(2).
		WithNombreApellidoMedico("DR. GOMEZ").
		WithNitDocumentoMedico(7654321).
		WithCantidad(1.0).
		WithUnidadMedida(1).
		WithPrecioUnitario(150.0).
		WithSubTotal(150.0).
		Build()

	factura := facturas.NewHospitalClinicaBuilder().
		WithCabecera(cabecera).
		AddDetalle(detalle).
		WithModalidad(tc.Modalidad).
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
		WithCodigoDocumentoSector(17).
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
