package siat_test

import (
	"fmt"
	"testing"

	"github.com/ron86i/go-siat"
	"github.com/stretchr/testify/assert"
)

func TestGetMensaje(t *testing.T) {
	tests := []struct {
		code     int
		expected string
	}{
		{903, "Recepción Procesada"},
		{926, "Comunicación Exitosa"},
		{1013, "El Calculo Del Monto Total Es Erróneo"},
		{3010, "La Factura Ya Se Encuentra Utilizada o Consolidada"},
		{9999, "Código de respuesta SIAT desconocido: 9999"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Code %d", tt.code), func(t *testing.T) {
			assert.Equal(t, tt.expected, siat.GetMensaje(tt.code))
		})
	}
}
