package invoices_test

import (
	"bytes"
	"encoding/xml"
	"io"
	"testing"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
)

func TestAlcanzadaIceRespetaSecuenciaXSD(t *testing.T) {
	for _, modalidad := range []int{siat.ModalidadElectronica, siat.ModalidadComputarizada} {
		factura := invoices.NewAlcanzadaIceBuilder().
			WithCabecera(invoices.NewAlcanzadaIceCabeceraBuilder().Build()).
			AddDetalle(invoices.NewAlcanzadaIceDetalleBuilder().Build()).
			WithModalidad(modalidad).
			Build()

		xmlData, err := xml.Marshal(factura)
		if err != nil {
			t.Fatalf("modalidad %d: no se pudo serializar: %v", modalidad, err)
		}

		assertXMLElementSequence(t, xmlData, "cabecera", []string{
			"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura",
			"cuf", "cufd", "codigoSucursal", "direccion", "codigoPuntoVenta", "fechaEmision",
			"nombreRazonSocial", "codigoTipoDocumentoIdentidad", "numeroDocumento", "complemento",
			"codigoCliente", "codigoMetodoPago", "numeroTarjeta", "subTotalNeto", "totalPrecioNeto",
			"iva", "montoaPagar", "totalIceEspecifico", "totalIcePorcentual", "codigoMoneda",
			"tipoCambio", "montoTotalMoneda", "descuentoAdicional", "codigoExcepcion", "leyenda",
			"usuario", "codigoInterno", "codigoDocumentoSector",
		})
		assertXMLElementSequence(t, xmlData, "detalle", []string{
			"actividadEconomica", "codigoProductoSin", "codigoProducto", "descripcion", "cantidad",
			"unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal", "marcaIce",
			"precioNetoVentaIce", "alicuotaEspecifica", "alicuotaPorcentual", "montoIceEspecifico",
			"montoIcePorcentual", "cantidadIce",
		})
	}
}

func assertXMLElementSequence(t *testing.T, data []byte, parent string, want []string) {
	t.Helper()
	decoder := xml.NewDecoder(bytes.NewReader(data))
	var got []string
	depth := 0
	inside := false
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("no se pudo leer XML: %v", err)
		}
		switch token := token.(type) {
		case xml.StartElement:
			if inside {
				if depth == 1 {
					got = append(got, token.Name.Local)
				}
				depth++
			} else if token.Name.Local == parent {
				inside = true
				depth = 1
			}
		case xml.EndElement:
			if inside {
				depth--
				if depth == 0 {
					inside = false
				}
			}
		}
	}
	if len(got) != len(want) {
		t.Fatalf("%s: elementos XML = %v; esperado = %v", parent, got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("%s elemento %d = %q; esperado %q", parent, i+1, got[i], want[i])
		}
	}
}
