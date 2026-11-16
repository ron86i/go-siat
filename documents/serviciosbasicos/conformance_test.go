package serviciosbasicos_test

import (
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/ron86i/go-siat/v3/documents/serviciosbasicos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type xmlNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	Text     string     `xml:",chardata"`
	Children []xmlNode  `xml:",any"`
}

func expectedServicioBasicoCabeceraSequence() []string {
	return []string{
		"nitEmisor", "razonSocialEmisor", "municipio", "telefono", "numeroFactura",
		"cuf", "cufd", "codigoSucursal", "direccion", "codigoPuntoVenta",
		"mes", "gestion", "ciudad", "zona", "numeroMedidor", "fechaEmision",
		"nombreRazonSocial", "domicilioCliente", "codigoTipoDocumentoIdentidad", "numeroDocumento",
		"complemento", "codigoCliente", "codigoMetodoPago", "numeroTarjeta",
		"subTotalNeto", "totalPrecioNeto", "iva", "montoaPagar", "consumoPeriodo",
		"beneficiarioLey1886", "complementoBeneficiario", "montoDescuentoLey1886", "montoDescuentoSocial",
		"tasaAseo", "tasaAlumbrado", "ajusteNoSujetoIva", "detalleAjusteNoSujetoIva",
		"otrosPagosNoSujetoIva", "detalleOtrosPagosNoSujetoIva", "otrasTasas",
		"codigoMoneda", "tipoCambio", "montoTotalMoneda", "descuentoAdicional",
		"codigoExcepcion", "leyenda", "usuario", "codigoInterno", "codigoDocumentoSector",
	}
}

func expectedServicioBasicoDetalleSequence() []string {
	return []string{
		"actividadEconomica", "codigoProductoSin", "codigoProducto", "descripcion",
		"cantidad", "unidadMedida", "precioNetoUnitario", "montoDescuento", "subTotal",
	}
}

func TestConformidadServiciosBasicos_SecuenciaXSD(t *testing.T) {
	for _, modalidad := range []int{serviciosbasicos.ModalidadElectronica, serviciosbasicos.ModalidadComputarizada} {
		cab, det := validFixture()
		f, err := serviciosbasicos.New(cab, det)
		require.NoError(t, err)

		data, err := f.XML(modalidad)
		require.NoError(t, err)

		var root xmlNode
		err = xml.Unmarshal(data, &root)
		require.NoError(t, err)

		rootExpected := "facturaElectronicaServicioBasico"
		if modalidad == serviciosbasicos.ModalidadComputarizada {
			rootExpected = "facturaComputarizadaServicioBasico"
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
		expectedCab := expectedServicioBasicoCabeceraSequence()
		require.Equal(t, len(expectedCab), len(cabeceraNode.Children), "cantidad de campos en cabecera difiere del XSD")
		for i, exp := range expectedCab {
			assert.Equal(t, exp, cabeceraNode.Children[i].XMLName.Local, "campo %d de cabecera fuera de secuencia", i+1)
		}

		// Verificar secuencia de etiquetas de Detalle
		expectedDet := expectedServicioBasicoDetalleSequence()
		require.Equal(t, len(expectedDet), len(detalleNode.Children), "cantidad de campos en detalle difiere del XSD")
		for i, exp := range expectedDet {
			assert.Equal(t, exp, detalleNode.Children[i].XMLName.Local, "campo %d de detalle fuera de secuencia", i+1)
		}

		// Validar ausencia total de campos obsoletos de v2
		xmlStr := string(data)
		for _, obsoleto := range []string{"<montoTotal>", "<montoTotalSujetoIva>", "<cafc>", "<precioUnitario>", "<montoDescuentoTarifaDignidad>", "<ajusteSujetoIva>"} {
			assert.NotContains(t, xmlStr, obsoleto, "el XML v3 contiene campo obsoleto: "+obsoleto)
		}
	}
}

func TestConformidadServiciosBasicos_NilablesConAtributoXsiNil(t *testing.T) {
	cab, det := validFixture()
	// Campos nilables fijados explícitamente en nil
	cab.Telefono = nil
	cab.CodigoPuntoVenta = nil
	cab.Mes = nil
	cab.Gestion = nil
	cab.Ciudad = nil
	cab.Zona = nil
	cab.NombreRazonSocial = nil
	cab.DomicilioCliente = nil
	cab.NumeroTarjeta = nil
	cab.ConsumoPeriodo = nil
	cab.BeneficiarioLey1886 = nil
	cab.MontoDescuentoLey1886 = nil
	cab.MontoDescuentoSocial = nil
	cab.TasaAseo = nil
	cab.TasaAlumbrado = nil
	cab.AjusteNoSujetoIva = nil
	cab.DetalleAjusteNoSujetoIva = nil
	cab.OtrosPagosNoSujetoIva = nil
	cab.DetalleOtrosPagosNoSujetoIva = nil
	cab.OtrasTasas = nil
	cab.DescuentoAdicional = nil
	cab.CodigoExcepcion = nil
	det[0].MontoDescuento = nil

	f, err := serviciosbasicos.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(serviciosbasicos.ModalidadElectronica)
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
		"telefono":                     false,
		"codigoPuntoVenta":             false,
		"mes":                          false,
		"gestion":                      false,
		"ciudad":                       false,
		"zona":                         false,
		"nombreRazonSocial":            false,
		"domicilioCliente":             false,
		"numeroTarjeta":                false,
		"consumoPeriodo":               false,
		"beneficiarioLey1886":          false,
		"montoDescuentoLey1886":        false,
		"montoDescuentoSocial":         false,
		"tasaAseo":                     false,
		"tasaAlumbrado":                false,
		"ajusteNoSujetoIva":            false,
		"detalleAjusteNoSujetoIva":     false,
		"otrosPagosNoSujetoIva":        false,
		"detalleOtrosPagosNoSujetoIva": false,
		"otrasTasas":                   false,
		"descuentoAdicional":           false,
		"codigoExcepcion":              false,
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

func TestConformidadServiciosBasicos_CaracteresEspecialesEscapados(t *testing.T) {
	cab, det := validFixture()
	rz := "EMPRESA & HIJOS <S.A.>"
	cab.NombreRazonSocial = &rz
	det[0].Descripcion = "Servicio de Agua & Alcantarillado \"Sector A\" <Tarifa 1>"

	f, err := serviciosbasicos.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(serviciosbasicos.ModalidadElectronica)
	require.NoError(t, err)

	xmlStr := string(data)
	assert.Contains(t, xmlStr, "EMPRESA &amp; HIJOS &lt;S.A.&gt;")
	assert.Contains(t, xmlStr, "Servicio de Agua &amp; Alcantarillado &#34;Sector A&#34; &lt;Tarifa 1&gt;")

	var parsed xmlNode
	err = xml.Unmarshal(data, &parsed)
	require.NoError(t, err, "el XML con caracteres escapados debe ser sintácticamente válido")
}

func TestConformidadServiciosBasicos_ValidacionExternaXSD(t *testing.T) {
	xsdPath := os.Getenv("SIAT_XSD_PATH")
	if xsdPath == "" {
		t.Skip("SIAT_XSD_PATH no configurado, omitiendo validación externa XSD con xmllint")
	}

	schemaFile := filepath.Join(xsdPath, "facturaElectronicaServicioBasico.xsd")
	if _, err := os.Stat(schemaFile); os.IsNotExist(err) {
		t.Skipf("archivo XSD no encontrado en %s, omitiendo", schemaFile)
	}

	xmllint, err := exec.LookPath("xmllint")
	if err != nil {
		t.Skip("herramienta xmllint no disponible en PATH, omitiendo validación")
	}

	cab, det := validFixture()
	f, err := serviciosbasicos.New(cab, det)
	require.NoError(t, err)

	data, err := f.XML(serviciosbasicos.ModalidadElectronica)
	require.NoError(t, err)

	tmpXml, err := os.CreateTemp("", "factura_serviciosbasicos_*.xml")
	require.NoError(t, err)
	defer os.Remove(tmpXml.Name())

	_, err = tmpXml.Write(data)
	require.NoError(t, err)
	require.NoError(t, tmpXml.Close())

	cmd := exec.Command(xmllint, "--noout", "--schema", schemaFile, tmpXml.Name())
	out, err := cmd.CombinedOutput()
	require.NoError(t, err, "xmllint validación XSD falló: %s", string(out))
}
