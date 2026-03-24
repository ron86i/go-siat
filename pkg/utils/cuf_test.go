package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerarCUF(t *testing.T) {
	nit := int64(123456789)
	fecha := time.Date(2023, 10, 27, 10, 30, 0, 0, time.UTC)
	sucursal := 0
	modalidad := 1
	tipoEmision := 1
	tipoFactura := 1
	tipoDocumentoSector := 1
	numeroFactura := 1
	puntoVenta := 0
	codigoControl := "ABCDE12345"

	cuf, err := GenerarCUF(nit, fecha, sucursal, modalidad, tipoEmision, tipoFactura, tipoDocumentoSector, numeroFactura, puntoVenta, codigoControl)
	assert.NoError(t, err)
	assert.NotEmpty(t, cuf)
	assert.True(t, len(cuf) > 10)
}

func TestCalculaDigitoMod11(t *testing.T) {
	t.Run("Standard Modulo 11", func(t *testing.T) {
		result := calculaDigitoMod11("1234567", 1, 9, false)
		assert.Equal(t, 1, len(result))
	})
}
