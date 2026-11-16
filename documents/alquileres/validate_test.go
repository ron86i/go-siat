package alquileres_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/alquileres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate_ErrorsJoin(t *testing.T) {
	// Factura completamente vacía para acumular múltiples errores
	cab := alquileres.Cabecera{CodigoDocumentoSector: 2}
	f, err := alquileres.New(cab, nil)
	require.Error(t, err)
	assert.Nil(t, f)

	msg := err.Error()
	assert.Contains(t, msg, "NIT del emisor inválido")
	assert.Contains(t, msg, "número de factura debe ser mayor a cero")
	assert.Contains(t, msg, "fecha de emisión requerida")
	assert.Contains(t, msg, "CUF requerido")
	assert.Contains(t, msg, "CUFD requerido")
	assert.Contains(t, msg, "dirección requerida")
	assert.Contains(t, msg, "periodo facturado requerido")
	assert.Contains(t, msg, "código de método de pago inválido")
	assert.Contains(t, msg, "código de moneda inválido")
	assert.Contains(t, msg, "tipo de cambio debe ser mayor a cero")
	assert.Contains(t, msg, "leyenda requerida")
	assert.Contains(t, msg, "usuario requerido")
	assert.Contains(t, msg, "la factura requiere al menos un detalle")
}

func TestValidate_StringLengths(t *testing.T) {
	cab, det := validFixture()
	longStr := strings.Repeat("A", 201)
	cab.RazonSocialEmisor = longStr
	cab.Municipio = strings.Repeat("M", 101)
	tel := strings.Repeat("1", 21)
	cab.Telefono = &tel
	cab.Direccion = strings.Repeat("D", 201)
	rz := strings.Repeat("R", 251)
	cab.NombreRazonSocial = &rz
	cab.NumeroDocumento = strings.Repeat("N", 31)
	cab.Complemento = "TOOLONG"
	cab.CodigoCliente = strings.Repeat("C", 101)
	cab.PeriodoFacturado = strings.Repeat("P", 201)
	cab.Leyenda = strings.Repeat("L", 501)
	cab.Usuario = strings.Repeat("U", 101)
	cab.CodigoInterno = strings.Repeat("I", 51)

	_, err := alquileres.New(cab, det)
	require.Error(t, err)
	msg := err.Error()

	assert.Contains(t, msg, "razonSocialEmisor excede la longitud máxima")
	assert.Contains(t, msg, "municipio excede la longitud máxima")
	assert.Contains(t, msg, "telefono excede la longitud máxima")
	assert.Contains(t, msg, "direccion excede la longitud máxima")
	assert.Contains(t, msg, "nombreRazonSocial excede la longitud máxima")
	assert.Contains(t, msg, "numeroDocumento excede la longitud máxima")
	assert.Contains(t, msg, "complemento excede la longitud máxima")
	assert.Contains(t, msg, "codigoCliente excede la longitud máxima")
	assert.Contains(t, msg, "periodoFacturado excede la longitud máxima")
	assert.Contains(t, msg, "leyenda excede la longitud máxima")
	assert.Contains(t, msg, "usuario excede la longitud máxima")
	assert.Contains(t, msg, "codigoInterno excede la longitud máxima")
}

func TestValidate_DetalleFields(t *testing.T) {
	cab, _ := validFixture()
	det := []alquileres.Detalle{
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
	_, err := alquileres.New(cab, det)
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
	_, err := alquileres.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura contiene un valor decimal no finito")

	cab, det = validFixture()
	cab.TipoCambio = math.Inf(1)
	_, err = alquileres.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura contiene un valor decimal no finito")

	cab, det = validFixture()
	det[0].PrecioNetoUnitario = math.NaN()
	_, err = alquileres.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura contiene un valor decimal no finito")
}

func TestValidate_SucursalNegativa(t *testing.T) {
	cab, det := validFixture()
	cab.CodigoSucursal = -1
	_, err := alquileres.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "código de sucursal inválido")
}

func TestValidate_FechaCero(t *testing.T) {
	cab, det := validFixture()
	cab.FechaEmision = time.Time{}
	_, err := alquileres.New(cab, det)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "fecha de emisión requerida")
}
