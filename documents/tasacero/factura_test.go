package tasacero_test

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/tasacero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validFixture() (tasacero.Cabecera, []tasacero.Detalle) {
	fecha := time.Date(2026, 10, 31, 12, 30, 15, 123000000, time.UTC)
	tel := "2222333"
	pv := 1
	rz := "EDITORIAL LIBROS S.A."
	tarjeta := int64(1234567812345678)
	gift := 20.0
	desc := 5.0
	excepcion := 0

	cab := tasacero.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "EDITORIAL LIBROS S.A.",
		Municipio:                    "La Paz",
		Telefono:                     &tel,
		NumeroFactura:                1001,
		Cuf:                          "123456789ABCDEF",
		Cufd:                         "CUFD-TEST-123",
		CodigoSucursal:               0,
		Direccion:                    "Av. 6 de Agosto 123",
		CodigoPuntoVenta:             &pv,
		FechaEmision:                 fecha,
		NombreRazonSocial:            &rz,
		CodigoTipoDocumentoIdentidad: 1,
		NumeroDocumento:              "9876543",
		Complemento:                  "",
		CodigoCliente:                "CLI-01",
		CodigoMetodoPago:             1,
		NumeroTarjeta:                &tarjeta,
		SubTotalNeto:                 100.00,
		TotalPrecioNeto:              100.00,
		Iva:                          0.00, // Tasa cero alícuota 0
		MontoAPagar:                  75.00,
		CodigoMoneda:                 1,
		TipoCambio:                   1.0,
		MontoTotalMoneda:             75.00,
		MontoGiftCard:                &gift,
		DescuentoAdicional:           &desc,
		CodigoExcepcion:              &excepcion,
		Leyenda:                      "Ley N 453: Tienes derecho a un trato equitativo y digno.",
		Usuario:                      "operador",
		CodigoInterno:                "TC-001",
		CodigoDocumentoSector:        8,
	}

	detDesc := 2.50
	detalles := []tasacero.Detalle{
		{
			ActividadEconomica: "581100",
			CodigoProductoSin:  83141,
			CodigoProducto:     "LIBRO-01",
			Descripcion:        "Libro de Programación Go",
			Cantidad:           1.0,
			UnidadMedida:       58,
			PrecioNetoUnitario: 100.00,
			MontoDescuento:     &detDesc,
			SubTotal:           97.50,
		},
	}

	return cab, detalles
}

func TestNew_Success(t *testing.T) {
	cab, det := validFixture()
	f, err := tasacero.New(cab, det)
	require.NoError(t, err)
	require.NotNil(t, f)

	assert.Equal(t, int64(123456789), f.Cabecera().NitEmisor)
	assert.Equal(t, 8, f.Cabecera().CodigoDocumentoSector)
	assert.Equal(t, 3, f.Protocolo())
	assert.Len(t, f.Detalles(), 1)
}

func TestNew_AutoAsignaSector(t *testing.T) {
	cab, det := validFixture()
	cab.CodigoDocumentoSector = 0
	f, err := tasacero.New(cab, det)
	require.NoError(t, err)
	assert.Equal(t, 8, f.Cabecera().CodigoDocumentoSector)
}

func TestNew_RechazaSectorIncorrecto(t *testing.T) {
	cab, det := validFixture()
	cab.CodigoDocumentoSector = 1 // Compra-Venta
	_, err := tasacero.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sector incompatible")
}

func TestFactura_InmutabilidadProfunda(t *testing.T) {
	cab, det := validFixture()
	f, err := tasacero.New(cab, det)
	require.NoError(t, err)

	// Modificar variables originales
	*cab.Telefono = "9999999"
	*cab.CodigoPuntoVenta = 99
	*cab.NombreRazonSocial = "OTRO NOMBRE"
	*cab.MontoGiftCard = 999.99
	det[0].Descripcion = "MODIFICADO"
	*det[0].MontoDescuento = 99.99

	// Verificar que el estado interno permanece inalterado
	cabAfter := f.Cabecera()
	assert.Equal(t, "2222333", *cabAfter.Telefono)
	assert.Equal(t, 1, *cabAfter.CodigoPuntoVenta)
	assert.Equal(t, "EDITORIAL LIBROS S.A.", *cabAfter.NombreRazonSocial)
	assert.Equal(t, 20.0, *cabAfter.MontoGiftCard)

	detAfter := f.Detalles()
	assert.Equal(t, "Libro de Programación Go", detAfter[0].Descripcion)
	assert.Equal(t, 2.50, *detAfter[0].MontoDescuento)

	// Modificar a través de los getters
	*cabAfter.Telefono = "0000000"
	*detAfter[0].MontoDescuento = 0.0

	// Comprobar nuevamente
	assert.Equal(t, "2222333", *f.Cabecera().Telefono)
	assert.Equal(t, 2.50, *f.Detalles()[0].MontoDescuento)
}

func TestFacturaXML_SerializacionModalidades(t *testing.T) {
	cab, det := validFixture()
	f, err := tasacero.New(cab, det)
	require.NoError(t, err)

	// 1. Modalidad Electrónica
	xmlElec, err := f.XML(tasacero.ModalidadElectronica)
	require.NoError(t, err)
	strElec := string(xmlElec)
	assert.True(t, strings.HasPrefix(strElec, "<facturaElectronicaTasaCero"))
	assert.Contains(t, strElec, `xsi:noNamespaceSchemaLocation="facturaElectronicaTasaCero.xsd"`)
	assert.Contains(t, strElec, "<nitEmisor>123456789</nitEmisor>")
	assert.Contains(t, strElec, "<iva>0</iva>")
	assert.Contains(t, strElec, "<montoaPagar>75</montoaPagar>")
	assert.Contains(t, strElec, "<montoGiftCard>20</montoGiftCard>")
	assert.Contains(t, strElec, "<codigoDocumentoSector>8</codigoDocumentoSector>")

	// 2. Modalidad Computarizada
	xmlComp, err := f.XML(tasacero.ModalidadComputarizada)
	require.NoError(t, err)
	strComp := string(xmlComp)
	assert.True(t, strings.HasPrefix(strComp, "<facturaComputarizadaTasaCero"))
	assert.Contains(t, strComp, `xsi:noNamespaceSchemaLocation="facturaComputarizadaTasaCero.xsd"`)
	assert.Contains(t, strComp, "<iva>0</iva>")

	// 3. Modalidad inválida
	_, err = f.XML(99)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "modalidad inválida")

	// 4. MarshalXML directo sin modalidad debe ser rechazado
	_, err = xml.Marshal(f)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "requiere especificar la modalidad de emisión")
}

func TestNilFactura(t *testing.T) {
	var f *tasacero.Factura
	assert.Equal(t, tasacero.Cabecera{}, f.Cabecera())
	assert.Nil(t, f.Detalles())
	assert.Equal(t, 0, f.CodigoDocumentoSector())
	assert.Equal(t, 0, f.CodigoSector())
	assert.Equal(t, 0, f.CodigoSucursal())
	assert.Nil(t, f.CodigoPuntoVenta())
	assert.Equal(t, "", f.Cufd())
	assert.Equal(t, 0, f.Protocolo())
	assert.Error(t, f.Validate())

	_, err := f.XML(tasacero.ModalidadElectronica)
	assert.Error(t, err)

	_, err = tasacero.Marshal(f, tasacero.ModalidadElectronica)
	assert.Error(t, err)
}
