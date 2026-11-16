package notacreditodebito_test

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/notacreditodebito"
)

func TestNotaV3CreditoDebitoMarshal(t *testing.T) {
	nota := notacreditodebito.NotaCreditoDebito{
		Modalidad: 1, // Electrónica
		Cabecera: notacreditodebito.Cabecera{
			NitEmisor:                    123456789,
			RazonSocialEmisor:            "Comercializadora Central V3",
			Municipio:                    "La Paz",
			NumeroNotaCreditoDebito:      6,
			Cuf:                          "CUF-NC-123V3",
			Cufd:                         "CUFD-NC-123",
			Direccion:                    "Calle 21 Calacoto",
			FechaEmision:                 time.Date(2026, 10, 31, 14, 0, 0, 0, time.UTC),
			CodigoTipoDocumentoIdentidad: 1,
			NumeroDocumento:              "44556677",
			CodigoCliente:                "CLI-RET",
			NumeroFactura:                100,
			NumeroAutorizacionCuf:        "CUF-ORIG-100",
			FechaEmisionFactura:          time.Date(2026, 10, 15, 10, 0, 0, 0, time.UTC),
			SubTotalNetoOriginal:         200.0,
			TotalPrecioNetoOriginal:      200.0,
			IvaOriginal:                  26.0,
			MontoAPagarOriginal:          200.0,
			SubTotalNetoDevuelto:         50.0,
			TotalPrecioNetoDevuelto:      50.0,
			IvaDevuelto:                  6.5,
			MontoAPagarDevuelto:          50.0,
			DescuentoAdicionalOriginal:   0,
			MontoDescuentoDevuelto:       0,
			Leyenda:                      "Ley N 453",
			Usuario:                      "cajero",
			CodigoInterno:                "NC-001",
		},
		Detalles: []notacreditodebito.Detalle{
			{
				NroItem:                  1,
				ActividadEconomica:       "471100",
				CodigoProductoSin:        88776,
				CodigoProducto:           "PROD-01",
				Descripcion:              "Devolución parcial producto",
				Cantidad:                 1,
				UnidadMedida:             1,
				PrecioNetoUnitario:       50.0,
				SubTotal:                 50.0,
				CodigoDetalleTransaccion: 1,
			},
		},
	}

	data, err := xml.Marshal(nota)
	if err != nil {
		t.Fatalf("error serializando: %v", err)
	}

	xmlStr := string(data)
	if !strings.Contains(xmlStr, "<notaFiscalElectronicaCreditoDebito") {
		t.Errorf("elemento raíz incorrecto: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<subTotalNetoOriginal>200</subTotalNetoOriginal>") {
		t.Errorf("subTotalNetoOriginal no encontrado: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<precioNetoUnitario>50</precioNetoUnitario>") {
		t.Errorf("precioNetoUnitario no encontrado: %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<nroItem>1</nroItem>") {
		t.Errorf("nroItem no encontrado: %s", xmlStr)
	}
}
