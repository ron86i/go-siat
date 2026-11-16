package compraventa_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate_FacturaNil(t *testing.T) {
	var f *compraventa.Factura
	err := f.Validate()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura no puede ser nil")
}

func TestValidate_ErrorsJoinConsolidado(t *testing.T) {
	// Factura vacía: debe reportar todos los campos obligatorios ausentes en una sola pasada
	_, err := compraventa.New(compraventa.Cabecera{}, nil)
	require.Error(t, err)

	msg := err.Error()
	assert.Contains(t, msg, "NIT del emisor inválido")
	assert.Contains(t, msg, "número de factura debe ser mayor a cero")
	assert.Contains(t, msg, "fecha de emisión requerida")
	assert.Contains(t, msg, "CUF requerido")
	assert.Contains(t, msg, "CUFD requerido")
	assert.Contains(t, msg, "dirección requerida")
	assert.Contains(t, msg, "código de método de pago inválido")
	assert.Contains(t, msg, "código de moneda inválido")
	assert.Contains(t, msg, "tipo de cambio debe ser mayor a cero")
	assert.Contains(t, msg, "leyenda requerida")
	assert.Contains(t, msg, "usuario requerido")
	assert.Contains(t, msg, "la factura requiere al menos un detalle")
}

func TestValidate_Sector(t *testing.T) {
	cab := validCabecera()
	cab.CodigoDocumentoSector = 2 // Sector incompatible
	_, err := compraventa.New(cab, validDetalles())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sector incompatible con compra-venta")
}

func TestValidate_CabeceraCamposObligatorios(t *testing.T) {
	tests := []struct {
		name        string
		modify      func(c *compraventa.Cabecera)
		errContains string
	}{
		{
			name:        "NitEmisor negativo o cero",
			modify:      func(c *compraventa.Cabecera) { c.NitEmisor = 0 },
			errContains: "NIT del emisor inválido",
		},
		{
			name:        "NumeroFactura negativo o cero",
			modify:      func(c *compraventa.Cabecera) { c.NumeroFactura = 0 },
			errContains: "número de factura debe ser mayor a cero",
		},
		{
			name:        "FechaEmision zero",
			modify:      func(c *compraventa.Cabecera) { c.FechaEmision = time.Time{} },
			errContains: "fecha de emisión requerida",
		},
		{
			name:        "CUF vacio",
			modify:      func(c *compraventa.Cabecera) { c.Cuf = "" },
			errContains: "CUF requerido",
		},
		{
			name:        "CUFD vacio",
			modify:      func(c *compraventa.Cabecera) { c.Cufd = "" },
			errContains: "CUFD requerido",
		},
		{
			name:        "CodigoSucursal negativo",
			modify:      func(c *compraventa.Cabecera) { c.CodigoSucursal = -1 },
			errContains: "código de sucursal inválido",
		},
		{
			name:        "Direccion vacia",
			modify:      func(c *compraventa.Cabecera) { c.Direccion = "" },
			errContains: "dirección requerida",
		},
		{
			name:        "CodigoMetodoPago cero",
			modify:      func(c *compraventa.Cabecera) { c.CodigoMetodoPago = 0 },
			errContains: "código de método de pago inválido",
		},
		{
			name:        "CodigoMoneda cero",
			modify:      func(c *compraventa.Cabecera) { c.CodigoMoneda = 0 },
			errContains: "código de moneda inválido",
		},
		{
			name:        "TipoCambio negativo o cero",
			modify:      func(c *compraventa.Cabecera) { c.TipoCambio = 0 },
			errContains: "tipo de cambio debe ser mayor a cero",
		},
		{
			name:        "Leyenda vacia",
			modify:      func(c *compraventa.Cabecera) { c.Leyenda = "" },
			errContains: "leyenda requerida",
		},
		{
			name:        "Usuario vacio",
			modify:      func(c *compraventa.Cabecera) { c.Usuario = "" },
			errContains: "usuario requerido",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cab := validCabecera()
			tt.modify(&cab)
			_, err := compraventa.New(cab, validDetalles())
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errContains)
		})
	}
}

func TestValidate_CabeceraLongitudMaxima(t *testing.T) {
	tests := []struct {
		name        string
		modify      func(c *compraventa.Cabecera)
		errContains string
	}{
		{
			name:        "RazonSocialEmisor > 200",
			modify:      func(c *compraventa.Cabecera) { c.RazonSocialEmisor = strings.Repeat("A", 201) },
			errContains: "razonSocialEmisor excede la longitud máxima (200)",
		},
		{
			name:        "Municipio > 100",
			modify:      func(c *compraventa.Cabecera) { c.Municipio = strings.Repeat("M", 101) },
			errContains: "municipio excede la longitud máxima (100)",
		},
		{
			name:        "Telefono > 20",
			modify:      func(c *compraventa.Cabecera) { t := strings.Repeat("1", 21); c.Telefono = &t },
			errContains: "telefono excede la longitud máxima (20)",
		},
		{
			name:        "Direccion > 200",
			modify:      func(c *compraventa.Cabecera) { c.Direccion = strings.Repeat("D", 201) },
			errContains: "direccion excede la longitud máxima (200)",
		},
		{
			name:        "NombreRazonSocial > 250",
			modify:      func(c *compraventa.Cabecera) { n := strings.Repeat("C", 251); c.NombreRazonSocial = &n },
			errContains: "nombreRazonSocial excede la longitud máxima (250)",
		},
		{
			name:        "NumeroDocumento > 30",
			modify:      func(c *compraventa.Cabecera) { c.NumeroDocumento = strings.Repeat("9", 31) },
			errContains: "numeroDocumento excede la longitud máxima (30)",
		},
		{
			name:        "Complemento > 5",
			modify:      func(c *compraventa.Cabecera) { c.Complemento = "123456" },
			errContains: "complemento excede la longitud máxima (5)",
		},
		{
			name:        "CodigoCliente > 100",
			modify:      func(c *compraventa.Cabecera) { c.CodigoCliente = strings.Repeat("K", 101) },
			errContains: "codigoCliente excede la longitud máxima (100)",
		},
		{
			name:        "Leyenda > 500",
			modify:      func(c *compraventa.Cabecera) { c.Leyenda = strings.Repeat("L", 501) },
			errContains: "leyenda excede la longitud máxima (500)",
		},
		{
			name:        "Usuario > 100",
			modify:      func(c *compraventa.Cabecera) { c.Usuario = strings.Repeat("U", 101) },
			errContains: "usuario excede la longitud máxima (100)",
		},
		{
			name:        "CodigoInterno > 50",
			modify:      func(c *compraventa.Cabecera) { c.CodigoInterno = strings.Repeat("I", 51) },
			errContains: "codigoInterno excede la longitud máxima (50)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cab := validCabecera()
			tt.modify(&cab)
			_, err := compraventa.New(cab, validDetalles())
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errContains)
		})
	}
}

func TestValidate_Detalles(t *testing.T) {
	tests := []struct {
		name        string
		modify      func(d *compraventa.Detalle)
		errContains string
	}{
		{
			name:        "ActividadEconomica vacia",
			modify:      func(d *compraventa.Detalle) { d.ActividadEconomica = "" },
			errContains: "actividad económica requerida",
		},
		{
			name:        "ActividadEconomica > 20",
			modify:      func(d *compraventa.Detalle) { d.ActividadEconomica = strings.Repeat("A", 21) },
			errContains: "actividad económica excede la longitud máxima (20)",
		},
		{
			name:        "CodigoProductoSin cero",
			modify:      func(d *compraventa.Detalle) { d.CodigoProductoSin = 0 },
			errContains: "código de producto SIN inválido",
		},
		{
			name:        "CodigoProducto vacio",
			modify:      func(d *compraventa.Detalle) { d.CodigoProducto = "" },
			errContains: "código de producto requerido",
		},
		{
			name:        "CodigoProducto > 50",
			modify:      func(d *compraventa.Detalle) { d.CodigoProducto = strings.Repeat("P", 51) },
			errContains: "código de producto excede la longitud máxima (50)",
		},
		{
			name:        "Descripcion vacia",
			modify:      func(d *compraventa.Detalle) { d.Descripcion = "" },
			errContains: "descripción requerida",
		},
		{
			name:        "Descripcion > 500",
			modify:      func(d *compraventa.Detalle) { d.Descripcion = strings.Repeat("D", 501) },
			errContains: "descripción excede la longitud máxima (500)",
		},
		{
			name:        "Cantidad negativa o cero",
			modify:      func(d *compraventa.Detalle) { d.Cantidad = 0 },
			errContains: "cantidad debe ser mayor a cero",
		},
		{
			name:        "UnidadMedida cero",
			modify:      func(d *compraventa.Detalle) { d.UnidadMedida = 0 },
			errContains: "unidad de medida inválida",
		},
		{
			name:        "PrecioNetoUnitario negativo",
			modify:      func(d *compraventa.Detalle) { d.PrecioNetoUnitario = -10.0 },
			errContains: "precio neto unitario no puede ser negativo",
		},
		{
			name:        "NumeroSerie > 100",
			modify:      func(d *compraventa.Detalle) { s := strings.Repeat("S", 101); d.NumeroSerie = &s },
			errContains: "número de serie excede la longitud máxima (100)",
		},
		{
			name:        "NumeroImei > 100",
			modify:      func(d *compraventa.Detalle) { i := strings.Repeat("I", 101); d.NumeroImei = &i },
			errContains: "número de IMEI excede la longitud máxima (100)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			det := validDetalles()
			tt.modify(&det[0])
			_, err := compraventa.New(validCabecera(), det)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errContains)
		})
	}
}

func TestValidate_ValoresDecimalesNoFinitos(t *testing.T) {
	// NaN en cabecera
	cabNaN := validCabecera()
	cabNaN.Iva = math.NaN()
	_, err := compraventa.New(cabNaN, validDetalles())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura contiene un valor decimal no finito")

	// Inf en cabecera
	cabInf := validCabecera()
	cabInf.MontoAPagar = math.Inf(1)
	_, err = compraventa.New(cabInf, validDetalles())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura contiene un valor decimal no finito")

	// NaN en detalle
	detNaN := validDetalles()
	detNaN[0].PrecioNetoUnitario = math.NaN()
	_, err = compraventa.New(validCabecera(), detNaN)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura contiene un valor decimal no finito")

	// Inf en descuento de detalle
	detInf := validDetalles()
	detInf[0].MontoDescuento = new(math.Inf(-1))
	_, err = compraventa.New(validCabecera(), detInf)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "la factura contiene un valor decimal no finito")
}
