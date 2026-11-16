package alquileres_test

import (
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/alquileres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type xmlNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Text     string     `xml:",chardata"`
	Children []xmlNode  `xml:",any"`
}

func expectedAlquilerCabeceraSequence() []string {
	return []string{
		"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura",
		"cuf", "cufd", "codigoSucursal", "direccion", "codigoPuntoVenta", "fechaEmision",
		"nombreRazonSocial", "codigoTipoDocumentoIdentidad", "numeroDocumento", "complemento",
		"codigoCliente", "periodoFacturado", "codigoMetodoPago", "numeroTarjeta", "subTotalNeto",
		"totalPrecioNeto", "iva", "montoaPagar", "codigoMoneda", "tipoCambio", "montoTotalMoneda",
		"descuentoAdicional", "codigoExcepcion", "leyenda", "usuario", "codigoInterno",
		"codigoDocumentoSector",
	}
}

func expectedAlquilerDetalleSequence() []string {
	return []string{
		"actividadEconomica", "codigoProductoSin", "codigoProducto", "descripcion",
		"cantidad", "unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal",
	}
}

func TestConformidadAlquileres_SecuenciaXSD(t *testing.T) {
	for _, modalidad := range []int{alquileres.ModalidadElectronica, alquileres.ModalidadComputarizada} {
		cab, det := validFixture()
		f, err := alquileres.New(cab, det)
		require.NoError(t, err)

		data, err := f.XML(modalidad)
		require.NoError(t, err)

		var root xmlNode
		err = xml.Unmarshal(data, &root)
		require.NoError(t, err)

		rootExpected := "facturaElectronicaAlquilerBienInmueble"
		if modalidad == alquileres.ModalidadComputarizada {
			rootExpected = "facturaComputarizadaAlquilerBienInmueble"
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
		expectedCab := expectedAlquilerCabeceraSequence()
		require.Equal(t, len(expectedCab), len(cabeceraNode.Children), "cantidad de campos en cabecera difiere del XSD")
		for i, exp := range expectedCab {
			assert.Equal(t, exp, cabeceraNode.Children[i].XMLName.Local, "campo %d de cabecera fuera de secuencia", i+1)
		}

		// Verificar secuencia de etiquetas de Detalle
		expectedDet := expectedAlquilerDetalleSequence()
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
		if schemaPath := os.Getenv("SIAT_ALQUILER_XSD"); schemaPath != "" && modalidad == alquileres.ModalidadComputarizada {
			tmpFile := filepath.Join(t.TempDir(), "factura_alq.xml")
			err = os.WriteFile(tmpFile, data, 0o600)
			require.NoError(t, err)

			cmd := exec.CommandContext(t.Context(), "xmllint", "--nonet", "--noout", "--schema", schemaPath, tmpFile)
			out, err := cmd.CombinedOutput()
			assert.NoError(t, err, "xmllint falló con XSD oficial: %s", string(out))
		}
	}
}

func TestConformidadAlquileres_CamposNulosXsiNil(t *testing.T) {
	cab, det := validFixture()
	// Asignar nulos explícitos
	cab.Telefono = nil
	cab.CodigoPuntoVenta = nil
	cab.NombreRazonSocial = nil
	cab.NumeroTarjeta = nil
	cab.DescuentoAdicional = nil
	cab.CodigoExcepcion = nil
	det[0].MontoDescuento = nil

	f, err := alquileres.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(alquileres.ModalidadComputarizada)
	require.NoError(t, err)

	var root xmlNode
	err = xml.Unmarshal(data, &root)
	require.NoError(t, err)

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

func TestConformidadAlquileres_EscapadoCaracteresEspeciales(t *testing.T) {
	cab, det := validFixture()
	cab.RazonSocialEmisor = "INMOBILIARIA <EL SOL> & CIA \"LTDA\" 'BOLIVIA'"
	det[0].Descripcion = "DEPARTAMENTO 4B & PARQUEO <EDIFICIO CENTRO> \"PREMIUM\""

	f, err := alquileres.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(alquileres.ModalidadComputarizada)
	require.NoError(t, err)
	xmlStr := string(data)

	assert.NotContains(t, xmlStr, "<EL SOL>")
	assert.NotContains(t, xmlStr, "<EDIFICIO CENTRO>")
	assert.Contains(t, xmlStr, "&lt;EL SOL&gt;")
	assert.Contains(t, xmlStr, "&lt;EDIFICIO CENTRO&gt;")
}
