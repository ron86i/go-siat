package models_test

import (
	"testing"

	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models/invoices"
)

func TestBuildersFacturaDefinenTipoFiscalPredeterminado(t *testing.T) {
	casos := []struct {
		nombre   string
		factura  interface{ TipoFacturaDocumento() int }
		esperado int
	}{
		{"compra y venta", invoices.NewCompraVentaBuilder().Build(), siat.TipoFacturaConDerechoCreditoFiscal},
		{"exportación", invoices.NewComercialExportacionBuilder().Build(), siat.TipoFacturaSinDerechoCreditoFiscal},
		{"turismo", invoices.NewTurismoHospedajeBuilder().Build(), siat.TipoFacturaSinDerechoCreditoFiscal},
		{"tasa cero", invoices.NewTasaCeroBuilder().Build(), siat.TipoFacturaSinDerechoCreditoFiscal},
		{"hotel", invoices.NewHotelBuilder().Build(), siat.TipoFacturaConDerechoCreditoFiscal},
	}
	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			if obtenido := caso.factura.TipoFacturaDocumento(); obtenido != caso.esperado {
				t.Fatalf("tipo = %d; se esperaba %d", obtenido, caso.esperado)
			}
		})
	}
}

func TestBuilderFacturaPermiteSobrescribirTipoFiscal(t *testing.T) {
	casos := []struct {
		nombre   string
		factura  interface{ TipoFacturaDocumento() int }
		esperado int
	}{
		{
			nombre: "builder con configuración propia",
			factura: invoices.NewTasaCeroBuilder().
				WithTipoFacturaDocumento(siat.TipoFacturaConDerechoCreditoFiscal).
				Build(),
			esperado: siat.TipoFacturaConDerechoCreditoFiscal,
		},
		{
			nombre: "builder centralizado",
			factura: invoices.NewLibreConsignacionBuilder().
				WithTipoFacturaDocumento(siat.TipoFacturaConDerechoCreditoFiscal).
				Build(),
			esperado: siat.TipoFacturaConDerechoCreditoFiscal,
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			if obtenido := caso.factura.TipoFacturaDocumento(); obtenido != caso.esperado {
				t.Fatalf("tipo = %d; se esperaba %d", obtenido, caso.esperado)
			}
		})
	}
}
