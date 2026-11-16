package serviciosbasicos_test

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/serviciosbasicos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validFixture() (serviciosbasicos.Cabecera, []serviciosbasicos.Detalle) {
	fecha := time.Date(2026, 10, 31, 12, 30, 15, 123000000, time.UTC)
	tel := "2222333"
	pv := 1
	rz := "CLIENTE EPSAS"
	dom := "Av. Principal 456"
	tarjeta := int64(1234567812345678)
	desc := 5.0
	excepcion := 0
	mes := "OCTUBRE"
	gestion := 2026
	ciudad := "La Paz"
	zona := "Miraflores"
	consumo := 25.5
	benLey := int64(123456)
	descLey := 10.0
	descSoc := 5.0
	tasaAseo := 15.0
	tasaAlumb := 10.5
	ajusteNoIva := 2.0
	detAjuste := "Ajuste tarifario"
	otrosPagos := 1.5
	detOtros := "Mantenimiento medidor"
	otrasTasas := 0.5

	cab := serviciosbasicos.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "EPSAS V3",
		Municipio:                    "La Paz",
		Telefono:                     &tel,
		NumeroFactura:                1001,
		Cuf:                          "123456789ABCDEF",
		Cufd:                         "CUFD-TEST-123",
		CodigoSucursal:               0,
		Direccion:                    "Av. Arce 123",
		CodigoPuntoVenta:             &pv,
		Mes:                          &mes,
		Gestion:                      &gestion,
		Ciudad:                       &ciudad,
		Zona:                         &zona,
		NumeroMedidor:                "MED-9988",
		FechaEmision:                 fecha,
		NombreRazonSocial:            &rz,
		DomicilioCliente:             &dom,
		CodigoTipoDocumentoIdentidad: 1,
		NumeroDocumento:              "9876543",
		Complemento:                  "",
		CodigoCliente:                "CLI-01",
		CodigoMetodoPago:             1,
		NumeroTarjeta:                &tarjeta,
		SubTotalNeto:                 100.00,
		TotalPrecioNeto:              100.00,
		Iva:                          13.00,
		MontoAPagar:                  120.50,
		ConsumoPeriodo:               &consumo,
		BeneficiarioLey1886:          &benLey,
		ComplementoBeneficiario:      "",
		MontoDescuentoLey1886:        &descLey,
		MontoDescuentoSocial:         &descSoc,
		TasaAseo:                     &tasaAseo,
		TasaAlumbrado:                &tasaAlumb,
		AjusteNoSujetoIva:            &ajusteNoIva,
		DetalleAjusteNoSujetoIva:     &detAjuste,
		OtrosPagosNoSujetoIva:        &otrosPagos,
		DetalleOtrosPagosNoSujetoIva: &detOtros,
		OtrasTasas:                   &otrasTasas,
		CodigoMoneda:                 1,
		TipoCambio:                   1.0,
		MontoTotalMoneda:             120.50,
		DescuentoAdicional:           &desc,
		CodigoExcepcion:              &excepcion,
		Leyenda:                      "Ley N 453: Tienes derecho a un trato digno.",
		Usuario:                      "operador",
		CodigoInterno:                "SERV-01",
		CodigoDocumentoSector:        13,
	}

	detDesc := 2.50
	detalles := []serviciosbasicos.Detalle{
		{
			ActividadEconomica: "360000",
			CodigoProductoSin:  11223,
			CodigoProducto:     "AGUA-01",
			Descripcion:        "Consumo Agua Potable",
			Cantidad:           1.0,
			UnidadMedida:       1,
			PrecioNetoUnitario: 100.00,
			MontoDescuento:     &detDesc,
			SubTotal:           97.50,
		},
	}

	return cab, detalles
}

func TestNew_Success(t *testing.T) {
	cab, det := validFixture()
	f, err := serviciosbasicos.New(cab, det)
	require.NoError(t, err)
	require.NotNil(t, f)

	assert.Equal(t, int64(123456789), f.Cabecera().NitEmisor)
	assert.Equal(t, 13, f.Cabecera().CodigoDocumentoSector)
	assert.Equal(t, "MED-9988", f.Cabecera().NumeroMedidor)
	assert.Equal(t, 3, f.Protocolo())
	assert.Len(t, f.Detalles(), 1)
}

func TestNew_AutoAsignaSector(t *testing.T) {
	cab, det := validFixture()
	cab.CodigoDocumentoSector = 0
	f, err := serviciosbasicos.New(cab, det)
	require.NoError(t, err)
	assert.Equal(t, 13, f.Cabecera().CodigoDocumentoSector)
}

func TestNew_RechazaSectorIncorrecto(t *testing.T) {
	cab, det := validFixture()
	cab.CodigoDocumentoSector = 1 // Compra-Venta
	_, err := serviciosbasicos.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sector incompatible")
}

func TestFactura_InmutabilidadProfunda(t *testing.T) {
	cab, det := validFixture()
	f, err := serviciosbasicos.New(cab, det)
	require.NoError(t, err)

	// Modificar variables originales
	*cab.Telefono = "9999999"
	*cab.CodigoPuntoVenta = 99
	*cab.Mes = "DICIEMBRE"
	*cab.ConsumoPeriodo = 999.99
	*cab.TasaAseo = 0.0
	det[0].Descripcion = "MODIFICADO"
	*det[0].MontoDescuento = 99.99

	// Verificar que el estado interno de Factura permanece inalterado
	cabAfter := f.Cabecera()
	assert.Equal(t, "2222333", *cabAfter.Telefono)
	assert.Equal(t, 1, *cabAfter.CodigoPuntoVenta)
	assert.Equal(t, "OCTUBRE", *cabAfter.Mes)
	assert.Equal(t, 25.5, *cabAfter.ConsumoPeriodo)
	assert.Equal(t, 15.0, *cabAfter.TasaAseo)

	detAfter := f.Detalles()
	assert.Equal(t, "Consumo Agua Potable", detAfter[0].Descripcion)
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
	f, err := serviciosbasicos.New(cab, det)
	require.NoError(t, err)

	// 1. Modalidad Electrónica
	xmlElec, err := f.XML(serviciosbasicos.ModalidadElectronica)
	require.NoError(t, err)
	strElec := string(xmlElec)
	assert.True(t, strings.HasPrefix(strElec, "<facturaElectronicaServicioBasico"))
	assert.Contains(t, strElec, `xsi:noNamespaceSchemaLocation="facturaElectronicaServicioBasico.xsd"`)
	assert.Contains(t, strElec, "<nitEmisor>123456789</nitEmisor>")
	assert.Contains(t, strElec, "<numeroMedidor>MED-9988</numeroMedidor>")
	assert.Contains(t, strElec, "<mes>OCTUBRE</mes>")
	assert.Contains(t, strElec, "<consumoPeriodo>25.5</consumoPeriodo>")
	assert.Contains(t, strElec, "<montoaPagar>120.5</montoaPagar>")
	assert.Contains(t, strElec, "<codigoDocumentoSector>13</codigoDocumentoSector>")

	// 2. Modalidad Computarizada
	xmlComp, err := f.XML(serviciosbasicos.ModalidadComputarizada)
	require.NoError(t, err)
	strComp := string(xmlComp)
	assert.True(t, strings.HasPrefix(strComp, "<facturaComputarizadaServicioBasico"))
	assert.Contains(t, strComp, `xsi:noNamespaceSchemaLocation="facturaComputarizadaServicioBasico.xsd"`)
	assert.Contains(t, strComp, "<numeroMedidor>MED-9988</numeroMedidor>")

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
	var f *serviciosbasicos.Factura
	assert.Equal(t, serviciosbasicos.Cabecera{}, f.Cabecera())
	assert.Nil(t, f.Detalles())
	assert.Equal(t, 0, f.CodigoDocumentoSector())
	assert.Equal(t, 0, f.CodigoSector())
	assert.Equal(t, 0, f.CodigoSucursal())
	assert.Nil(t, f.CodigoPuntoVenta())
	assert.Equal(t, "", f.Cufd())
	assert.Equal(t, 0, f.Protocolo())
	assert.Error(t, f.Validate())

	_, err := f.XML(serviciosbasicos.ModalidadElectronica)
	assert.Error(t, err)

	_, err = serviciosbasicos.Marshal(f, serviciosbasicos.ModalidadElectronica)
	assert.Error(t, err)
}
