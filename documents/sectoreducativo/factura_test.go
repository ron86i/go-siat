package sectoreducativo_test

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/sectoreducativo"
)

func TestFacturaV3SectorEducativoMarshal(t *testing.T) {
	factura := sectoreducativo.Factura{
		Modalidad: 1, // Electrónica
		Cabecera: sectoreducativo.Cabecera{
			NitEmisor:                    123456789,
			RazonSocialEmisor:            "Colegio San Simon V3",
			Municipio:                    "Cochabamba",
			NumeroFactura:                11,
			Cuf:                          "CUF123V3",
			Cufd:                         "CUFD123",
			Direccion:                    "Calle Sucre",
			FechaEmision:                 time.Date(2026, 10, 31, 11, 0, 0, 0, time.UTC),
			CodigoTipoDocumentoIdentidad: 1,
			NumeroDocumento:              "1234567",
			CodigoCliente:                "CLI-01",
			NombreEstudiante:             "Juan Perez",
			PeriodoFacturado:             "Octubre 2026",
			CodigoMetodoPago:             1,
			SubTotalNeto:                 500.0,
			TotalPrecioNeto:              500.0,
			Iva:                          65.0,
			MontoAPagar:                  500.0,
			CodigoMoneda:                 1,
			TipoCambio:                   1.0,
			MontoTotalMoneda:             500.0,
			Leyenda:                      "Ley N 453",
			Usuario:                      "cajero",
			CodigoInterno:                "EDU-123",
		},
		Detalles: []sectoreducativo.Detalle{
			{
				ActividadEconomica: "851000",
				CodigoProductoSin:  99123,
				CodigoProducto:     "MATR-01",
				Descripcion:        "Mensualidad escolar",
				Cantidad:           1,
				UnidadMedida:       1,
				PrecioNetoUnitario: 500.0,
				SubTotal:           500.0,
			},
		},
	}

	data, err := xml.Marshal(factura)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}

	xmlStr := string(data)
	if !strings.Contains(xmlStr, "<facturaElectronicaSectorEducativo") {
		t.Errorf("nombre de elemento raíz incorrecto: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<subTotalNeto>500</subTotalNeto>") {
		t.Errorf("subTotalNeto no encontrado: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<precioNetoUnitario>500</precioNetoUnitario>") {
		t.Errorf("precioNetoUnitario no encontrado: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<codigoInterno>EDU-123</codigoInterno>") {
		t.Errorf("codigoInterno no encontrado: %s", xmlStr)
	}
}
