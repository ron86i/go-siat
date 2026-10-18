package invoices_test

import (
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
)

func TestCompraVentaContratoNeto(t *testing.T) {
	for _, modalidad := range []int{1, 2} {
		cabecera := invoices.NewCompraVentaCabeceraBuilder().
			WithNitEmisor(123456789).WithRazonSocialEmisor("EMISOR DE PRUEBA").
			WithMunicipio("LA PAZ").WithNumeroFactura(1).WithCuf("ABC123").
			WithCufd("ABC123").WithDireccion("DIRECCION DE PRUEBA").
			WithFechaEmision(time.Date(2027, 6, 1, 9, 19, 48, 675000000, time.UTC)).
			WithCodigoTipoDocumentoIdentidad(1).WithNumeroDocumento("123456").
			WithCodigoCliente("CLIENTE1").WithCodigoMetodoPago(1).
			WithSubTotalNeto(100.126).WithTotalPrecioNeto(100.126).
			WithIva(13.126).WithMontoAPagar(113.256).
			WithCodigoMoneda(1).WithTipoCambio(1).WithMontoTotalMoneda(113.256).
			WithLeyenda("FACTURA DE PRUEBA").WithUsuario("prueba").
			WithCodigoInterno("ABC123").Build()
		detalle := invoices.NewCompraVentaDetalleBuilder().
			WithActividadEconomica("620100").WithCodigoProductoSin(83141).
			WithCodigoProducto("P1").WithDescripcion("PRODUCTO DE PRUEBA").
			WithCantidad(1.004).WithUnidadMedida(58).
			WithPrecioNetoUnitario(100.126).WithMontoDescuento(new(0.004)).
			WithSubTotal(100.126).Build()
		factura := invoices.NewCompraVentaBuilder().WithModalidad(modalidad).
			WithCabecera(cabecera).AddDetalle(detalle).Build()
		data, err := xml.Marshal(factura)
		if err != nil {
			t.Fatal(err)
		}
		s := string(data)
		for _, expected := range []string{
			"<subTotalNeto>100.13</subTotalNeto><totalPrecioNeto>100.13</totalPrecioNeto><iva>13.13</iva><montoaPagar>113.26</montoaPagar><codigoMoneda>",
			"</usuario><codigoInterno>ABC123</codigoInterno><codigoDocumentoSector>",
			"</unidadMedida><precioNetoUnitario>100.13</precioNetoUnitario><montoDescuento",
		} {
			if !strings.Contains(s, expected) {
				t.Fatalf("modalidad %d: falta secuencia %s en %s", modalidad, expected, s)
			}
		}
		for _, obsolete := range []string{"<montoTotal>", "<montoTotalSujetoIva>", "<cafc", "<precioUnitario>"} {
			if strings.Contains(s, obsolete) {
				t.Fatalf("campo obsoleto %s", obsolete)
			}
		}
		// La ruta externa permite validar contra el XSD del usuario sin copiarlo
		// al repositorio. La electrónica requiere además una firma XMLDSig.
		if schema := os.Getenv("SIAT_COMPRA_VENTA_XSD"); schema != "" && modalidad == 2 {
			file := filepath.Join(t.TempDir(), "factura.xml")
			if err := os.WriteFile(file, data, 0o600); err != nil {
				t.Fatal(err)
			}
			if out, err := exec.CommandContext(t.Context(), "xmllint", "--nonet", "--noout", "--schema", schema, file).CombinedOutput(); err != nil {
				t.Fatalf("XSD: %v\n%s", err, out)
			}
		}
	}
}
