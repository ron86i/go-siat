package invoices_test

import (
	"encoding/xml"
	"testing"

	"github.com/ron86i/go-siat/v3"
	"github.com/ron86i/go-siat/v3/pkg/models/invoices"
)

func TestHotelRespetaSecuenciaXSD(t *testing.T) {
	for _, modalidad := range []int{siat.ModalidadElectronica, siat.ModalidadComputarizada} {
		factura := invoices.NewHotelBuilder().
			WithCabecera(invoices.NewHotelCabeceraBuilder().Build()).
			AddDetalle(invoices.NewHotelDetalleBuilder().Build()).
			WithModalidad(modalidad).
			Build()

		xmlData, err := xml.Marshal(factura)
		if err != nil {
			t.Fatalf("modalidad %d: no se pudo serializar: %v", modalidad, err)
		}
		assertXMLElementSequence(t, xmlData, "cabecera", []string{
			"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura", "cuf", "cufd",
			"codigoSucursal", "direccion", "codigoPuntoVenta", "fechaEmision", "nombreRazonSocial",
			"codigoTipoDocumentoIdentidad", "numeroDocumento", "complemento", "codigoCliente",
			"cantidadHuespedes", "cantidadHabitaciones", "cantidadMayores", "cantidadMenores",
			"fechaIngresoHospedaje", "codigoMetodoPago", "numeroTarjeta", "subTotalNeto", "totalPrecioNeto",
			"iva", "montoaPagar", "codigoMoneda", "tipoCambio", "montoTotalMoneda", "montoGiftCard",
			"descuentoAdicional", "codigoExcepcion", "leyenda", "usuario", "codigoInterno",
			"codigoDocumentoSector",
		})
		assertXMLElementSequence(t, xmlData, "detalle", []string{
			"actividadEconomica", "codigoProductoSin", "codigoProducto", "codigoTipoHabitacion", "descripcion",
			"cantidad", "unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal", "detalleHuespedes",
		})
	}
}
