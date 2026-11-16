package compraventa_test

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validCabecera() compraventa.Cabecera {
	telefono := "22114455"
	nombre := "JUAN PEREZ"
	return compraventa.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "MI EMPRESA S.R.L.",
		Municipio:                    "La Paz",
		Telefono:                     &telefono,
		NumeroFactura:                100,
		Cuf:                          "1234567890ABCDEF1234567890ABCDEF12345678",
		Cufd:                         "BQUEw-MDQ4MzAwM",
		CodigoSucursal:               0,
		Direccion:                    "Av. 6 de Agosto Nro 1234",
		FechaEmision:                 time.Date(2026, 10, 31, 12, 30, 15, 0, time.UTC),
		NombreRazonSocial:            &nombre,
		CodigoTipoDocumentoIdentidad: 1,
		NumeroDocumento:              "1234567",
		CodigoCliente:                "CLI-001",
		CodigoMetodoPago:             1,
		SubTotalNeto:                 100.00,
		TotalPrecioNeto:              100.00,
		Iva:                          13.00,
		MontoAPagar:                  100.00,
		CodigoMoneda:                 1,
		TipoCambio:                   1.0,
		MontoTotalMoneda:             100.00,
		Leyenda:                      "Ley N° 453: Tienes derecho a un trato equitativo y digno.",
		Usuario:                      "operador1",
		CodigoDocumentoSector:        1,
	}
}

func validDetalles() []compraventa.Detalle {
	return []compraventa.Detalle{
		{
			ActividadEconomica: "620100",
			CodigoProductoSin:  83141,
			CodigoProducto:     "SERV-001",
			Descripcion:        "Servicio de software",
			Cantidad:           1.0,
			UnidadMedida:       58,
			PrecioNetoUnitario: 100.00,
			SubTotal:           100.00,
		},
	}
}

func TestFactura_ConstructorNew(t *testing.T) {
	cabecera := validCabecera()
	cabecera.CodigoDocumentoSector = 0 // Asignación automática a 1
	detalles := validDetalles()

	factura, err := compraventa.New(cabecera, detalles)
	require.NoError(t, err)
	require.NotNil(t, factura)

	// Getters
	assert.Equal(t, 1, factura.Cabecera().CodigoDocumentoSector)
	assert.Equal(t, int64(123456789), factura.Cabecera().NitEmisor)
	assert.Len(t, factura.Detalles(), 1)
	assert.Equal(t, "SERV-001", factura.Detalles()[0].CodigoProducto)

	// Copia defensiva de detalles: mutaciones externas no deben afectar la Factura
	detalles[0].CodigoProducto = "MODIFICADO"
	assert.Equal(t, "SERV-001", factura.Detalles()[0].CodigoProducto)

	// Mutación del slice devuelto por Detalles() no afecta la Factura
	devueltos := factura.Detalles()
	devueltos[0].CodigoProducto = "ALTERADO"
	assert.Equal(t, "SERV-001", factura.Detalles()[0].CodigoProducto)
}

func TestFactura_InmutabilidadProfunda(t *testing.T) {
	desc := 10.0
	serie := "12345"
	detalles := []compraventa.Detalle{
		{
			ActividadEconomica: "620100",
			CodigoProductoSin:  83141,
			CodigoProducto:     "SERV-001",
			Descripcion:        "Servicio",
			Cantidad:           1.0,
			UnidadMedida:       58,
			PrecioNetoUnitario: 100.0,
			SubTotal:           90.0,
			MontoDescuento:     &desc,
			NumeroSerie:        &serie,
		},
	}
	cabecera := validCabecera()
	descuentoAdic := 5.0
	cabecera.DescuentoAdicional = &descuentoAdic

	factura, err := compraventa.New(cabecera, detalles)
	require.NoError(t, err)

	// Mutar las variables externas originales
	desc = 999.0
	serie = "MUTADO"
	descuentoAdic = 888.0

	// Verificar que los datos en la factura no cambiaron (copia defensiva profunda en New)
	assert.Equal(t, 10.0, *factura.Detalles()[0].MontoDescuento)
	assert.Equal(t, "12345", *factura.Detalles()[0].NumeroSerie)
	assert.Equal(t, 5.0, *factura.Cabecera().DescuentoAdicional)

	// Mutar a través de los getters no debe alterar el estado interno
	devueltos := factura.Detalles()
	*devueltos[0].MontoDescuento = 777.0
	assert.Equal(t, 10.0, *factura.Detalles()[0].MontoDescuento)

	cabDevuelta := factura.Cabecera()
	*cabDevuelta.DescuentoAdicional = 666.0
	assert.Equal(t, 5.0, *factura.Cabecera().DescuentoAdicional)
}

func TestFactura_NilReceivers(t *testing.T) {
	var f *compraventa.Factura
	assert.Empty(t, f.Cabecera())
	assert.Nil(t, f.Detalles())
	_, err := compraventa.Marshal(f, compraventa.ModalidadElectronica)
	require.Error(t, err)
}

func TestFacturaXML_SerializacionPorModalidad(t *testing.T) {
	cabecera := validCabecera()
	cabecera.SubTotalNeto = 12.345
	detalles := validDetalles()
	detalles[0].Cantidad = 1.234
	detalles[0].MontoDescuento = new(0.004)

	factura, err := compraventa.New(cabecera, detalles)
	require.NoError(t, err)

	for _, modalidad := range []int{compraventa.ModalidadElectronica, compraventa.ModalidadComputarizada} {
		got, err := factura.XML(modalidad)
		require.NoError(t, err)

		rootTag := "facturaElectronicaCompraVenta"
		schemaTag := "facturaElectronicaCompraVenta.xsd"
		if modalidad == compraventa.ModalidadComputarizada {
			rootTag = "facturaComputarizadaCompraVenta"
			schemaTag = "facturaComputarizadaCompraVenta.xsd"
		}

		xmlStr := string(got)
		assert.True(t, strings.HasPrefix(xmlStr, "<"+rootTag), "debe iniciar con el tag raíz correspondiente")
		assert.Contains(t, xmlStr, `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"`)
		assert.Contains(t, xmlStr, `xsi:noNamespaceSchemaLocation="`+schemaTag+`"`)
		assert.Contains(t, xmlStr, "<nitEmisor>123456789</nitEmisor>")
		assert.Contains(t, xmlStr, "<subTotalNeto>12.35</subTotalNeto>") // Redondeo a 2 decimales
		assert.Contains(t, xmlStr, "<telefono>22114455</telefono>")
		assert.Contains(t, xmlStr, "<descripcion>Servicio de software</descripcion>")
		assert.Contains(t, xmlStr, "<cantidad>1.23</cantidad>")
	}
}

func TestFacturaXML_ModalidadInvalida(t *testing.T) {
	factura, err := compraventa.New(validCabecera(), validDetalles())
	require.NoError(t, err)

	_, err = factura.XML(99)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "modalidad inválida")
}

func TestFacturaXML_MarshalDirectoRechazado(t *testing.T) {
	factura, err := compraventa.New(validCabecera(), validDetalles())
	require.NoError(t, err)

	// xml.Marshal directo se rechaza para obligar a usar la modalidad correcta
	_, err = xml.Marshal(factura)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "requiere especificar la modalidad")
}

func TestFacturaXML_ValoresOpcionales(t *testing.T) {
	cabecera := validCabecera()
	cabecera.MontoGiftCard = new(1.235)
	cabecera.DescuentoAdicional = new(2.345)
	cabecera.CodigoExcepcion = new(int64(1))

	detalles := validDetalles()
	detalles[0].Cantidad = 1.235
	detalles[0].PrecioNetoUnitario = 2.345
	detalles[0].MontoDescuento = new(0.015)
	detalles[0].SubTotal = 3.455
	detalles[0].NumeroSerie = new("serie123")
	detalles[0].NumeroImei = new("imei456")

	factura, err := compraventa.New(cabecera, detalles)
	require.NoError(t, err)

	got, err := compraventa.Marshal(factura, compraventa.ModalidadComputarizada)
	require.NoError(t, err)

	xmlStr := string(got)
	assert.Contains(t, xmlStr, "<facturaComputarizadaCompraVenta")
	assert.Contains(t, xmlStr, "<montoGiftCard>1.24</montoGiftCard>")
	assert.Contains(t, xmlStr, "<descuentoAdicional>2.35</descuentoAdicional>")
	assert.Contains(t, xmlStr, "<codigoExcepcion>1</codigoExcepcion>")
	assert.Contains(t, xmlStr, "<numeroSerie>serie123</numeroSerie>")
	assert.Contains(t, xmlStr, "<numeroImei>imei456</numeroImei>")
}
