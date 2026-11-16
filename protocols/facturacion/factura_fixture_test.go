package facturacion_test

import (
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/compraventa"
)

func newTestFactura(t *testing.T, mods ...func(c *compraventa.Cabecera, d *[]compraventa.Detalle)) *compraventa.Factura {
	t.Helper()
	nombre := "JUAN PEREZ"
	telefono := "22114455"
	cab := compraventa.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "MI EMPRESA S.R.L.",
		Municipio:                    "La Paz",
		Telefono:                     &telefono,
		NumeroFactura:                100,
		Cuf:                          "1234567890ABCDEF1234567890ABCDEF12345678",
		Cufd:                         "cufd",
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
	detalles := []compraventa.Detalle{
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
	for _, m := range mods {
		if m != nil {
			m(&cab, &detalles)
		}
	}
	f, err := compraventa.New(cab, detalles)
	if err != nil {
		t.Fatalf("error construyendo factura de prueba: %v", err)
	}
	return f
}
