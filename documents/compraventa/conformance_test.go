package compraventa_test

import (
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type xmlNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Text     string     `xml:",chardata"`
	Children []xmlNode  `xml:",any"`
}

func expectedCompraVentaCabeceraSequence() []string {
	return []string{
		"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura",
		"cuf", "cufd", "codigoSucursal", "direccion", "codigoPuntoVenta", "fechaEmision",
		"nombreRazonSocial", "codigoTipoDocumentoIdentidad", "numeroDocumento", "complemento",
		"codigoCliente", "codigoMetodoPago", "numeroTarjeta", "subTotalNeto", "totalPrecioNeto",
		"iva", "montoaPagar", "codigoMoneda", "tipoCambio", "montoTotalMoneda", "montoGiftCard",
		"descuentoAdicional", "codigoExcepcion", "leyenda", "usuario", "codigoInterno",
		"codigoDocumentoSector",
	}
}

func expectedCompraVentaDetalleSequence() []string {
	return []string{
		"actividadEconomica", "codigoProductoSin", "codigoProducto", "descripcion",
		"cantidad", "unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal",
		"numeroSerie", "numeroImei",
	}
}

func TestConformidadCompraVenta_SecuenciaXSD(t *testing.T) {
	for _, modalidad := range []int{compraventa.ModalidadElectronica, compraventa.ModalidadComputarizada} {
		cab, det := fixtureFactura()
		f, err := compraventa.New(cab, det)
		require.NoError(t, err)

		data, err := f.XML(modalidad)
		require.NoError(t, err)

		var root xmlNode
		err = xml.Unmarshal(data, &root)
		require.NoError(t, err)

		rootExpected := "facturaElectronicaCompraVenta"
		if modalidad == compraventa.ModalidadComputarizada {
			rootExpected = "facturaComputarizadaCompraVenta"
		}
		assert.Equal(t, rootExpected, root.XMLName.Local)

		// Verificar atributo xsi:noNamespaceSchemaLocation
		foundSchema := false
		for _, attr := range root.Attrs {
			if attr.Name.Space == "http://www.w3.org/2001/XMLSchema-instance" && attr.Name.Local == "noNamespaceSchemaLocation" {
				foundSchema = (attr.Value == rootExpected+".xsd")
			}
		}
		assert.True(t, foundSchema, "atributo xsi:noNamespaceSchemaLocation inválido o ausente")

		// Extraer cabecera y detalle
		var cabeceraNode, detalleNode *xmlNode
		for i := range root.Children {
			child := &root.Children[i]
			switch child.XMLName.Local {
			case "cabecera":
				cabeceraNode = child
			case "detalle":
				detalleNode = child
			}
		}
		require.NotNil(t, cabeceraNode, "sección <cabecera> no encontrada")
		require.NotNil(t, detalleNode, "sección <detalle> no encontrada")

		// Verificar secuencia de etiquetas de Cabecera
		expectedCab := expectedCompraVentaCabeceraSequence()
		require.Equal(t, len(expectedCab), len(cabeceraNode.Children), "cantidad de campos en cabecera difiere del XSD")
		for i, exp := range expectedCab {
			assert.Equal(t, exp, cabeceraNode.Children[i].XMLName.Local, "campo %d de cabecera fuera de secuencia", i+1)
		}

		// Verificar secuencia de etiquetas de Detalle
		expectedDet := expectedCompraVentaDetalleSequence()
		require.Equal(t, len(expectedDet), len(detalleNode.Children), "cantidad de campos en detalle difiere del XSD")
		for i, exp := range expectedDet {
			assert.Equal(t, exp, detalleNode.Children[i].XMLName.Local, "campo %d de detalle fuera de secuencia", i+1)
		}

		// Validar ausencia total de campos obsoletos de v2
		xmlStr := string(data)
		for _, obsoleto := range []string{"<montoTotal>", "<montoTotalSujetoIva>", "<cafc>", "<precioUnitario>"} {
			assert.NotContains(t, xmlStr, obsoleto, "el XML v3 contiene campo obsoleto: "+obsoleto)
		}

		// Si se especifica el XSD oficial mediante variable de entorno, ejecutar validación xmllint externa
		if schemaPath := os.Getenv("SIAT_COMPRA_VENTA_XSD"); schemaPath != "" && modalidad == compraventa.ModalidadComputarizada {
			tmpFile := filepath.Join(t.TempDir(), "factura_cv.xml")
			err = os.WriteFile(tmpFile, data, 0o600)
			require.NoError(t, err)

			cmd := exec.CommandContext(t.Context(), "xmllint", "--nonet", "--noout", "--schema", schemaPath, tmpFile)
			out, err := cmd.CombinedOutput()
			assert.NoError(t, err, "xmllint falló con XSD oficial: %s", string(out))
		}
	}
}

func TestConformidadCompraVenta_CamposNulosXsiNil(t *testing.T) {
	cab, det := fixtureFactura()
	// Asignar nulos explícitos
	cab.Telefono = nil
	cab.CodigoPuntoVenta = nil
	cab.NombreRazonSocial = nil
	cab.NumeroTarjeta = nil
	cab.MontoGiftCard = nil
	cab.DescuentoAdicional = nil
	cab.CodigoExcepcion = nil
	det[0].MontoDescuento = nil
	det[0].NumeroSerie = nil
	det[0].NumeroImei = nil

	f, err := compraventa.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(compraventa.ModalidadComputarizada)
	require.NoError(t, err)

	var root xmlNode
	err = xml.Unmarshal(data, &root)
	require.NoError(t, err)

	// Verificar xsi:nil="true" en campos nulos de cabecera
	var cabNode *xmlNode
	for i := range root.Children {
		if root.Children[i].XMLName.Local == "cabecera" {
			cabNode = &root.Children[i]
		}
	}
	require.NotNil(t, cabNode)

	nilFields := map[string]bool{
		"telefono":           true,
		"codigoPuntoVenta":   true,
		"nombreRazonSocial":  true,
		"numeroTarjeta":      true,
		"montoGiftCard":      true,
		"descuentoAdicional": true,
		"codigoExcepcion":    true,
	}

	for _, ch := range cabNode.Children {
		if nilFields[ch.XMLName.Local] {
			hasNil := false
			for _, attr := range ch.Attrs {
				if attr.Name.Space == "http://www.w3.org/2001/XMLSchema-instance" && attr.Name.Local == "nil" && attr.Value == "true" {
					hasNil = true
				}
			}
			assert.True(t, hasNil, "campo %s debe tener xsi:nil=\"true\"", ch.XMLName.Local)
		}
	}
}

func TestConformidadCompraVenta_EscapadoCaracteresEspeciales(t *testing.T) {
	cab, det := fixtureFactura()
	cab.RazonSocialEmisor = "COMPAÑÍA <PRUEBA> & CIA \"COMERCIAL\" 'S.A.'"
	det[0].Descripcion = "ÍTEM & SERVICIO <TIPO A> \"SPECIAL\""

	f, err := compraventa.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(compraventa.ModalidadComputarizada)
	require.NoError(t, err)
	xmlStr := string(data)

	// No debe contener caracteres XML sin escapar en los valores de datos
	assert.NotContains(t, xmlStr, "<PRUEBA>")
	assert.NotContains(t, xmlStr, "<TIPO A>")
	assert.Contains(t, xmlStr, "&lt;PRUEBA&gt;")
	assert.Contains(t, xmlStr, "&lt;TIPO A&gt;")
}

func fixtureFactura() (compraventa.Cabecera, []compraventa.Detalle) {
	fecha := time.Date(2026, 10, 31, 14, 30, 0, 456000000, time.UTC)
	tel := "2222333"
	pv := 1
	rz := "CLIENTE INTERNACIONAL S.A."
	tarjeta := int64(1234567812345678)
	gift := 0.0
	desc := 5.0
	excepcion := int64(0)

	cab := compraventa.Cabecera{
		NitEmisor:                    123456789,
		RazonSocialEmisor:            "MI EMPRESA S.A.",
		Municipio:                    "La Paz",
		Telefono:                     &tel,
		NumeroFactura:                5001,
		Cuf:                          "123456789ABCDEF",
		Cufd:                         "CUFD-TEST-123",
		CodigoSucursal:               0,
		Direccion:                    "Av. 6 de Agosto 123",
		CodigoPuntoVenta:             &pv,
		FechaEmision:                 fecha,
		NombreRazonSocial:            &rz,
		CodigoTipoDocumentoIdentidad: 1,
		NumeroDocumento:              "9876543",
		Complemento:                  "",
		CodigoCliente:                "CLI-001",
		CodigoMetodoPago:             1,
		NumeroTarjeta:                &tarjeta,
		SubTotalNeto:                 100.0,
		TotalPrecioNeto:              100.0,
		Iva:                          13.0,
		MontoAPagar:                  95.0,
		CodigoMoneda:                 1,
		TipoCambio:                   1.0,
		MontoTotalMoneda:             95.0,
		MontoGiftCard:                &gift,
		DescuentoAdicional:           &desc,
		CodigoExcepcion:              &excepcion,
		Leyenda:                      "Ley N 453",
		Usuario:                      "operador",
		CodigoInterno:                "INT-001",
		CodigoDocumentoSector:        1,
	}

	detDesc := 2.50
	serie := "SN123456"
	imei := "IMEI987654"
	det := []compraventa.Detalle{
		{
			ActividadEconomica: "620100",
			CodigoProductoSin:  83141,
			CodigoProducto:     "PROD-1",
			Descripcion:        "Servicio de Desarrollo",
			Cantidad:           1.0,
			UnidadMedida:       58,
			PrecioNetoUnitario: 100.0,
			MontoDescuento:     &detDesc,
			SubTotal:           97.50,
			NumeroSerie:        &serie,
			NumeroImei:         &imei,
		},
	}
	return cab, det
}
