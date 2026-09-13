package models_test

import (
	"testing"

	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/ron86i/go-siat/v2/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestConfirmacionCompraBuilder_SerializaXMLDelXSD(t *testing.T) {
	confirmacion := models.NewConfirmacionCompraBuilder().
		WithNro(1).
		WithNitEmisor(154422029).
		WithCodigoAutorizacion("1").
		WithNumeroFactura("11110").
		WithTipoCompra(1).
		Build()

	xmlData, err := utils.MarshalIndentedXML(confirmacion)
	require.NoError(t, err)
	require.Equal(t,
		"<confirmacionCompra>\n  <nro>1</nro>\n  <nitEmisor>154422029</nitEmisor>\n  <codigoAutorizacion>1</codigoAutorizacion>\n  <numeroFactura>11110</numeroFactura>\n  <tipoCompra>1</tipoCompra>\n</confirmacionCompra>",
		string(xmlData),
	)
}
