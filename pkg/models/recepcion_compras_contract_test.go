package models_test

import (
	"encoding/xml"
	"io"
	"slices"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v3/pkg/models"
)

func TestAnulacionCompraBuilderRespetaSecuenciaXML(t *testing.T) {
	req := models.NewAnulacionCompraBuilder().
		WithCodAutorizacion("AUTORIZACION-PRUEBA").
		WithNitProveedor(123456789).
		WithNroFactura(123).
		WithNroDuiDim("DIM-PRUEBA").
		Build()
	contenido, err := xml.Marshal(models.UnwrapInternalRequest[facturacion.AnulacionCompra](req))
	if err != nil {
		t.Fatal(err)
	}
	secuencia := "<codAutorizacion>AUTORIZACION-PRUEBA</codAutorizacion>" +
		"<nitProveedor>123456789</nitProveedor>" +
		"<nroDuiDim>DIM-PRUEBA</nroDuiDim>" +
		"<nroFactura>123</nroFactura>"
	if !strings.Contains(string(contenido), secuencia) {
		t.Fatalf("anulacionCompra no respeta la secuencia XML requerida: %s", contenido)
	}
}

// Las secuencias se contrastaron con el WSDL histórico del servicio piloto v2:
// https://pilotosiatservicios.impuestos.gob.bo/v2/ServicioRecepcionCompras?wsdl=ServicioFacturacion.wsdl
// La comparación con /sfd/ queda pendiente mientras su WSDL responde HTTP 503.
func TestRecepcionComprasBuildersSecuenciasXML(t *testing.T) {
	common := []string{"codigoAmbiente", "codigoPuntoVenta", "codigoSistema", "codigoSucursal", "cufd", "cuis", "nit"}
	packageFields := []string{"archivo", "cantidadFacturas", "fechaEnvio", "gestion", "hashArchivo", "periodo"}
	cases := []struct {
		name    string
		request any
		extra   []string
	}{
		{"recepcionPaqueteCompras", models.NewRecepcionPaqueteComprasBuilder().Build(), packageFields},
		{"validacionRecepcionPaqueteCompras", models.NewValidacionRecepcionPaqueteComprasBuilder().Build(), []string{"codigoRecepcion"}},
		{"anulacionCompra", models.NewAnulacionCompraBuilder().Build(), []string{"codAutorizacion", "nitProveedor", "nroDuiDim", "nroFactura"}},
		{"confirmacionCompras", models.NewConfirmacionComprasBuilder().Build(), packageFields},
		{"consultaCompras", models.NewConsultaComprasBuilder().Build(), []string{"fecha"}},
		{"verificarComunicacion", models.NewVerificarComunicacionRecepcionComprasBuilder().Build(), nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := xml.Marshal(tc.request)
			if err != nil {
				t.Fatal(err)
			}
			decoder := xml.NewDecoder(strings.NewReader(string(data)))
			depth := 0
			var fields []string
			for {
				token, err := decoder.Token()
				if err == io.EOF {
					break
				}
				if err != nil {
					t.Fatal(err)
				}
				switch element := token.(type) {
				case xml.StartElement:
					depth++
					if depth == 1 && element.Name.Local != tc.name {
						t.Fatalf("operación XML inesperada: %s", element.Name.Local)
					}
					if depth == 3 {
						fields = append(fields, element.Name.Local)
					}
				case xml.EndElement:
					depth--
				}
			}
			var expected []string
			if tc.name != "verificarComunicacion" {
				expected = append(slices.Clone(common), tc.extra...)
			}
			if !slices.Equal(fields, expected) {
				t.Fatalf("secuencia XML incorrecta: recibido %v; esperado %v", fields, expected)
			}
		})
	}
}
