package serviciosbasicos_test

import (
	"math"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/serviciosbasicos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate_ErrorsJoin(t *testing.T) {
	// Factura completamente vacía para acumular múltiples errores
	cab := serviciosbasicos.Cabecera{CodigoDocumentoSector: 13}
	f, err := serviciosbasicos.New(cab, nil)
	require.Error(t, err)
	assert.Nil(t, f)

	msg := err.Error()
	assert.Contains(t, msg, "NIT del emisor inválido")
	assert.Contains(t, msg, "número de factura debe ser mayor a cero")
	assert.Contains(t, msg, "fecha de emisión requerida")
	assert.Contains(t, msg, "CUF requerido")
	assert.Contains(t, msg, "CUFD requerido")
	assert.Contains(t, msg, "dirección requerida")
	assert.Contains(t, msg, "número de medidor requerido")
	assert.Contains(t, msg, "código de método de pago inválido")
	assert.Contains(t, msg, "código de moneda inválido")
	assert.Contains(t, msg, "tipo de cambio debe ser mayor a cero")
	assert.Contains(t, msg, "leyenda requerida")
	assert.Contains(t, msg, "usuario requerido")
	assert.Contains(t, msg, "la factura requiere al menos un detalle")
}

func TestValidate_StringLengths(t *testing.T) {
	cab, det := validFixture()
	cab.RazonSocialEmisor = strings.Repeat("A", 201)
	cab.Municipio = strings.Repeat("M", 101)
	tel := strings.Repeat("1", 21)
	cab.Telefono = &tel
	cab.Direccion = strings.Repeat("D", 201)
	cab.NumeroMedidor = strings.Repeat("M", 51)
	rz := strings.Repeat("R", 251)
	cab.NombreRazonSocial = &rz
	dom := strings.Repeat("D", 201)
	cab.DomicilioCliente = &dom
	cab.NumeroDocumento = strings.Repeat("N", 31)
	cab.Complemento = "TOOLONG"
	cab.CodigoCliente = strings.Repeat("C", 101)
	mes := strings.Repeat("M", 21)
	cab.Mes = &mes
	ciudad := strings.Repeat("C", 101)
	cab.Ciudad = &ciudad
	zona := strings.Repeat("Z", 101)
	cab.Zona = &zona
	cab.ComplementoBeneficiario = "TOOLONG"
	detAj := strings.Repeat("A", 201)
	cab.DetalleAjusteNoSujetoIva = &detAj
	detOt := strings.Repeat("O", 201)
	cab.DetalleOtrosPagosNoSujetoIva = &detOt
	cab.Leyenda = strings.Repeat("L", 501)
	cab.Usuario = strings.Repeat("U", 101)
	cab.CodigoInterno = strings.Repeat("I", 51)

	_, err := serviciosbasicos.New(cab, det)
	require.Error(t, err)
	msg := err.Error()

	assert.Contains(t, msg, "razonSocialEmisor excede la longitud máxima")
	assert.Contains(t, msg, "municipio excede la longitud máxima")
	assert.Contains(t, msg, "telefono excede la longitud máxima")
	assert.Contains(t, msg, "direccion excede la longitud máxima")
	assert.Contains(t, msg, "numeroMedidor excede la longitud máxima")
	assert.Contains(t, msg, "nombreRazonSocial excede la longitud máxima")
	assert.Contains(t, msg, "domicilioCliente excede la longitud máxima")
	assert.Contains(t, msg, "numeroDocumento excede la longitud máxima")
	assert.Contains(t, msg, "complemento excede la longitud máxima")
	assert.Contains(t, msg, "codigoCliente excede la longitud máxima")
	assert.Contains(t, msg, "mes excede la longitud máxima")
	assert.Contains(t, msg, "ciudad excede la longitud máxima")
	assert.Contains(t, msg, "zona excede la longitud máxima")
	assert.Contains(t, msg, "complementoBeneficiario excede la longitud máxima")
	assert.Contains(t, msg, "detalleAjusteNoSujetoIva excede la longitud máxima")
	assert.Contains(t, msg, "detalleOtrosPagosNoSujetoIva excede la longitud máxima")
	assert.Contains(t, msg, "leyenda excede la longitud máxima")
	assert.Contains(t, msg, "usuario excede la longitud máxima")
	assert.Contains(t, msg, "codigoInterno excede la longitud máxima")
}

func TestValidate_DetalleFields(t *testing.T) {
	cab, _ := validFixture()
	det := []serviciosbasicos.Detalle{
		{
			ActividadEconomica: "",
			CodigoProductoSin:  0,
			CodigoProducto:     "",
			Descripcion:        "",
			Cantidad:           -1,
			UnidadMedida:       0,
			PrecioNetoUnitario: -10,
		},
	}
	_, err := serviciosbasicos.New(cab, det)
	require.Error(t, err)
	msg := err.Error()

	assert.Contains(t, msg, "detalle 1: actividad económica requerida")
	assert.Contains(t, msg, "detalle 1: código de producto SIN inválido")
	assert.Contains(t, msg, "detalle 1: código de producto requerido")
	assert.Contains(t, msg, "detalle 1: descripción requerida")
	assert.Contains(t, msg, "detalle 1: cantidad debe ser mayor a cero")
	assert.Contains(t, msg, "detalle 1: unidad de medida inválida")
	assert.Contains(t, msg, "detalle 1: precio neto unitario no puede ser negativo")
}

func TestValidate_DecimalsFinitos(t *testing.T) {
	cab, det := validFixture()
	cab.SubTotalNeto = math.NaN()
	_, err := serviciosbasicos.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la cabecera contiene un valor decimal no finito")

	cab, det = validFixture()
	det[0].PrecioNetoUnitario = math.Inf(1)
	_, err = serviciosbasicos.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "detalle 1 contiene un valor decimal no finito")
}
