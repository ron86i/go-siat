package alquileres_test

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/alquileres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validFixture() (alquileres.Cabecera, []alquileres.Detalle) {
	fecha := time.Date(2026, 10, 31, 12, 30, 15, 123000000, time.UTC)
	tel := "2222333"
	pv := 1
	rz := "CLIENTE S.A."
	tarjeta := int64(1234567812345678)
	desc := 5.0
	excepcion := 0

	cab := alquileres.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "EMPRESA ARRENDADORA S.A.",
		Municipio:                    "La Paz",
		Telefono:                     &tel,
		NumeroFactura:                1001,
		Cuf:                          "123456789ABCDEF",
		Cufd:                         "CUFD-TEST-123",
		CodigoSucursal:               0,
		Direccion:                    "Av. Arce 123",
		CodigoPuntoVenta:             &pv,
		FechaEmision:                 fecha,
		NombreRazonSocial:            &rz,
		CodigoTipoDocumentoIdentidad: 1,
		NumeroDocumento:              "9876543",
		Complemento:                  "",
		CodigoCliente:                "CLI-01",
		PeriodoFacturado:             "OCTUBRE 2026",
		CodigoMetodoPago:             1,
		NumeroTarjeta:                &tarjeta,
		SubTotalNeto:                 1000.00,
		TotalPrecioNeto:              1000.00,
		Iva:                          130.00,
		MontoAPagar:                  995.00,
		CodigoMoneda:                 1,
		TipoCambio:                   1.0,
		MontoTotalMoneda:             995.00,
		DescuentoAdicional:           &desc,
		CodigoExcepcion:              &excepcion,
		Leyenda:                      "Ley N 453",
		Usuario:                      "operador",
		CodigoInterno:                "ALQ-001",
		CodigoDocumentoSector:        2,
	}

	detDesc := 2.50
	detalles := []alquileres.Detalle{
		{
			ActividadEconomica: "681000",
			CodigoProductoSin:  99100,
			CodigoProducto:     "ALQ-OFICINA",
			Descripcion:        "Alquiler de Oficina 301 - Octubre",
			Cantidad:           1.0,
			UnidadMedida:       58,
			PrecioNetoUnitario: 1000.00,
			MontoDescuento:     &detDesc,
			SubTotal:           997.50,
		},
	}

	return cab, detalles
}

func TestNew_Success(t *testing.T) {
	cab, det := validFixture()
	f, err := alquileres.New(cab, det)
	require.NoError(t, err)
	require.NotNil(t, f)

	assert.Equal(t, int64(123456789), f.Cabecera().NitEmisor)
	assert.Equal(t, 2, f.Cabecera().CodigoDocumentoSector)
	assert.Equal(t, "OCTUBRE 2026", f.Cabecera().PeriodoFacturado)
	assert.Len(t, f.Detalles(), 1)
}

func TestNew_AutoAsignaSector(t *testing.T) {
	cab, det := validFixture()
	cab.CodigoDocumentoSector = 0
	f, err := alquileres.New(cab, det)
	require.NoError(t, err)
	assert.Equal(t, 2, f.Cabecera().CodigoDocumentoSector)
}

func TestNew_RechazaSectorIncorrecto(t *testing.T) {
	cab, det := validFixture()
	cab.CodigoDocumentoSector = 1 // Compra-Venta
	_, err := alquileres.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sector incompatible")
}

func TestFactura_InmutabilidadProfunda(t *testing.T) {
	cab, det := validFixture()
	f, err := alquileres.New(cab, det)
	require.NoError(t, err)

	// Modificar variables originales
	*cab.Telefono = "9999999"
	*cab.CodigoPuntoVenta = 99
	*cab.NombreRazonSocial = "OTRO NOMBRE"
	*cab.NumeroTarjeta = 999
	*cab.DescuentoAdicional = 999.99
	*cab.CodigoExcepcion = 99
	det[0].Descripcion = "MODIFICADO"
	*det[0].MontoDescuento = 99.99

	// Verificar que el estado interno de Factura permanece inalterado
	cabAfter := f.Cabecera()
	assert.Equal(t, "2222333", *cabAfter.Telefono)
	assert.Equal(t, 1, *cabAfter.CodigoPuntoVenta)
	assert.Equal(t, "CLIENTE S.A.", *cabAfter.NombreRazonSocial)
	assert.Equal(t, int64(1234567812345678), *cabAfter.NumeroTarjeta)
	assert.Equal(t, 5.0, *cabAfter.DescuentoAdicional)
	assert.Equal(t, 0, *cabAfter.CodigoExcepcion)

	detAfter := f.Detalles()
	assert.Equal(t, "Alquiler de Oficina 301 - Octubre", detAfter[0].Descripcion)
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
	f, err := alquileres.New(cab, det)
	require.NoError(t, err)

	// 1. Modalidad Electrónica
	xmlElec, err := f.XML(alquileres.ModalidadElectronica)
	require.NoError(t, err)
	strElec := string(xmlElec)
	assert.True(t, strings.HasPrefix(strElec, "<facturaElectronicaAlquilerBienInmueble"))
	assert.Contains(t, strElec, `xsi:noNamespaceSchemaLocation="facturaElectronicaAlquilerBienInmueble.xsd"`)
	assert.Contains(t, strElec, "<nitEmisor>123456789</nitEmisor>")
	assert.Contains(t, strElec, "<periodoFacturado>OCTUBRE 2026</periodoFacturado>")
	assert.Contains(t, strElec, "<montoaPagar>995</montoaPagar>")
	assert.Contains(t, strElec, "<codigoDocumentoSector>2</codigoDocumentoSector>")

	// 2. Modalidad Computarizada
	xmlComp, err := f.XML(alquileres.ModalidadComputarizada)
	require.NoError(t, err)
	strComp := string(xmlComp)
	assert.True(t, strings.HasPrefix(strComp, "<facturaComputarizadaAlquilerBienInmueble"))
	assert.Contains(t, strComp, `xsi:noNamespaceSchemaLocation="facturaComputarizadaAlquilerBienInmueble.xsd"`)
	assert.Contains(t, strComp, "<periodoFacturado>OCTUBRE 2026</periodoFacturado>")

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
	var f *alquileres.Factura
	assert.Equal(t, alquileres.Cabecera{}, f.Cabecera())
	assert.Nil(t, f.Detalles())
	assert.Error(t, f.Validate())

	_, err := f.XML(alquileres.ModalidadElectronica)
	assert.Error(t, err)

	_, err = alquileres.Marshal(f, alquileres.ModalidadElectronica)
	assert.Error(t, err)
}
