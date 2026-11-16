package tasacero_test

import (
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/tasacero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type xmlNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Text     string     `xml:",chardata"`
	Children []xmlNode  `xml:",any"`
}

func expectedTasaCeroCabeceraSequence() []string {
	return []string{
		"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura",
		"cuf", "cufd", "codigoSucursal", "direccion", "codigoPuntoVenta",
		"fechaEmision", "nombreRazonSocial", "codigoTipoDocumentoIdentidad", "numeroDocumento",
		"complemento", "codigoCliente", "codigoMetodoPago", "numeroTarjeta",
		"subTotalNeto", "totalPrecioNeto", "iva", "montoaPagar",
		"codigoMoneda", "tipoCambio", "montoTotalMoneda", "montoGiftCard",
		"descuentoAdicional", "codigoExcepcion", "leyenda", "usuario", "codigoInterno",
		"codigoDocumentoSector",
	}
}

func expectedTasaCeroDetalleSequence() []string {
	return []string{
		"actividadEconomica", "codigoProductoSin", "codigoProducto", "descripcion",
		"cantidad", "unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal",
	}
}

func TestConformidadTasaCero_SecuenciaXSD(t *testing.T) {
	for _, modalidad := range []int{tasacero.ModalidadElectronica, tasacero.ModalidadComputarizada} {
		cab, det := validFixture()
		f, err := tasacero.New(cab, det)
		require.NoError(t, err)

		data, err := f.XML(modalidad)
		require.NoError(t, err)

		var root xmlNode
		err = xml.Unmarshal(data, &root)
		require.NoError(t, err)

		rootExpected := "facturaElectronicaTasaCero"
		if modalidad == tasacero.ModalidadComputarizada {
			rootExpected = "facturaComputarizadaTasaCero"
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
			if child.XMLName.Local == "cabecera" {
				cabeceraNode = child
			} else if child.XMLName.Local == "detalle" {
				detalleNode = child
			}
		}
		require.NotNil(t, cabeceraNode, "sección <cabecera> no encontrada")
		require.NotNil(t, detalleNode, "sección <detalle> no encontrada")

		// Verificar secuencia de etiquetas de Cabecera
		expectedCab := expectedTasaCeroCabeceraSequence()
		require.Equal(t, len(expectedCab), len(cabeceraNode.Children), "cantidad de campos en cabecera difiere del XSD")
		for i, exp := range expectedCab {
			assert.Equal(t, exp, cabeceraNode.Children[i].XMLName.Local, "campo %d de cabecera fuera de secuencia", i+1)
		}

		// Verificar secuencia de etiquetas de Detalle
		expectedDet := expectedTasaCeroDetalleSequence()
		require.Equal(t, len(expectedDet), len(detalleNode.Children), "cantidad de campos en detalle difiere del XSD")
		for i, exp := range expectedDet {
			assert.Equal(t, exp, detalleNode.Children[i].XMLName.Local, "campo %d de detalle fuera de secuencia", i+1)
		}

		// Validar ausencia total de campos obsoletos de v2
		xmlStr := string(data)
		for _, obsoleto := range []string{"<montoTotal>", "<montoTotalSujetoIva>", "<cafc>", "<precioUnitario>"} {
			assert.NotContains(t, xmlStr, obsoleto, "el XML v3 contiene campo obsoleto: "+obsoleto)
		}
	}
}

func TestConformidadTasaCero_NilablesConAtributoXsiNil(t *testing.T) {
	cab, det := validFixture()
	cab.Telefono = nil
	cab.CodigoPuntoVenta = nil
	cab.NombreRazonSocial = nil
	cab.NumeroTarjeta = nil
	cab.MontoGiftCard = nil
	cab.DescuentoAdicional = nil
	cab.CodigoExcepcion = nil
	det[0].MontoDescuento = nil

	f, err := tasacero.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(tasacero.ModalidadElectronica)
	require.NoError(t, err)

	var root xmlNode
	err = xml.Unmarshal(data, &root)
	require.NoError(t, err)

	var cabNode *xmlNode
	for i := range root.Children {
		if root.Children[i].XMLName.Local == "cabecera" {
			cabNode = &root.Children[i]
			break
		}
	}
	require.NotNil(t, cabNode)

	nilFields := map[string]bool{
		"telefono":           false,
		"codigoPuntoVenta":   false,
		"nombreRazonSocial":  false,
		"numeroTarjeta":      false,
		"montoGiftCard":      false,
		"descuentoAdicional": false,
		"codigoExcepcion":    false,
	}

	for _, child := range cabNode.Children {
		if _, ok := nilFields[child.XMLName.Local]; ok {
			for _, attr := range child.Attrs {
				if attr.Name.Space == "http://www.w3.org/2001/XMLSchema-instance" && attr.Name.Local == "nil" && attr.Value == "true" {
					nilFields[child.XMLName.Local] = true
				}
			}
		}
	}

	for field, foundNil := range nilFields {
		assert.True(t, foundNil, "el campo %s no incluyó xsi:nil=\"true\"", field)
	}
}

func TestConformidadTasaCero_CaracteresEspecialesEscapados(t *testing.T) {
	cab, det := validFixture()
	rz := "EDITORIAL & DISTRIBUIDORA <LA PAZ>"
	cab.NombreRazonSocial = &rz
	det[0].Descripcion = "Libro de \"Arte & Historia\" <Volumen 1>"

	f, err := tasacero.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(tasacero.ModalidadElectronica)
	require.NoError(t, err)

	xmlStr := string(data)
	assert.Contains(t, xmlStr, "EDITORIAL &amp; DISTRIBUIDORA &lt;LA PAZ&gt;")
	assert.Contains(t, xmlStr, "Libro de &#34;Arte &amp; Historia&#34; &lt;Volumen 1&gt;")

	var parsed xmlNode
	err = xml.Unmarshal(data, &parsed)
	require.NoError(t, err, "el XML con caracteres escapados debe ser sintácticamente válido")
}

func TestConformidadTasaCero_ValidacionExternaXSD(t *testing.T) {
	xsdPath := os.Getenv("SIAT_XSD_PATH")
	if xsdPath == "" {
		t.Skip("SIAT_XSD_PATH no configurado, omitiendo validación externa XSD con xmllint")
	}

	schemaFile := filepath.Join(xsdPath, "facturaElectronicaTasaCero.xsd")
	if _, err := os.Stat(schemaFile); os.IsNotExist(err) {
		t.Skipf("archivo XSD no encontrado en %s, omitiendo", schemaFile)
	}

	xmllint, err := exec.LookPath("xmllint")
	if err != nil {
		t.Skip("herramienta xmllint no disponible en PATH, omitiendo validación")
	}

	cab, det := validFixture()
	f, err := tasacero.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(tasacero.ModalidadElectronica)
	require.NoError(t, err)

	tmpXml, err := os.CreateTemp("", "factura_tasacero_*.xml")
	require.NoError(t, err)
	defer os.Remove(tmpXml.Name())

	_, err = tmpXml.Write(data)
	require.NoError(t, err)
	require.NoError(t, tmpXml.Close())

	cmd := exec.Command(xmllint, "--noout", "--schema", schemaFile, tmpXml.Name())
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "xmllint validación XSD falló: %s", string(out))
}
