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

func TestIsRetryableCode(t *testing.T) {
	retryableCodes := []int{123, 967, 991, 995, 999}
	for _, code := range retryableCodes {
		assert.True(t, siat.IsRetryableCode(code), "Code %d should be retryable", code)
	}
	assert.False(t, siat.IsRetryableCode(901), "Code 901 should not be retryable")
}

func TestIsValidationCode(t *testing.T) {
	validationCodes := []int{919, 940, 969, 971, 1000, 1061}
	for _, code := range validationCodes {
		assert.True(t, siat.IsValidationCode(code), "Code %d should be validation", code)
	}
	
	nonValidationCodes := []int{901, 908, 967, 991, 995, 2000, 3000}
	for _, code := range nonValidationCodes {
		assert.False(t, siat.IsValidationCode(code), "Code %d should not be validation", code)
	}
}

func TestIsWarningCode(t *testing.T) {
	warningCodes := []int{2000, 2010, 2019, 3008}
	for _, code := range warningCodes {
		assert.True(t, siat.IsWarningCode(code), "Code %d should be warning", code)
	}
	
	nonWarningCodes := []int{1999, 2020, 3000, 3005, 991}
	for _, code := range nonWarningCodes {
		assert.False(t, siat.IsWarningCode(code), "Code %d should not be warning", code)
	}
}

func TestIsConfigCode(t *testing.T) {
	configCodes := []int{910, 911, 912, 917, 958, 959, 975, 989}
	for _, code := range configCodes {
		assert.True(t, siat.IsConfigCode(code), "Code %d should be config", code)
	}
	assert.False(t, siat.IsConfigCode(901), "Code 901 should not be config")
}
