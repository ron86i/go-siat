package siat_test

import (
	"testing"

	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestMap_ToJSON(t *testing.T) {
	m := siat.Map{"key": "value", "num": 123}
	jsonStr, err := m.ToJSON()
	assert.NoError(t, err)
	assert.Contains(t, jsonStr, `"key":"value"`)
	assert.Contains(t, jsonStr, `"num":123`)
}

func TestMap_Sum(t *testing.T) {
	m := siat.Map{
		"a": float64(10.5),
		"b": 20,
		"c": int64(30),
		"d": "not a number",
	}
	sum := m.Sum()
	assert.Equal(t, 60.5, sum)
}

func TestMap_ToStruct(t *testing.T) {
	m := siat.Map{"name": "Juan", "age": 30}
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var p Person
	err := m.ToStruct(&p)
	assert.NoError(t, err)
	assert.Equal(t, "Juan", p.Name)
	assert.Equal(t, 30, p.Age)
}

func TestNew(t *testing.T) {
	t.Run("Valid Config", func(t *testing.T) {
		cfg := siat.Config{
			Token:          "test",
			Nit:            123456789,
			CodigoSistema:  "test",
			CodigoAmbiente: siat.AmbientePruebas,
			BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/v2",
		}
		services, err := siat.New(cfg)
		assert.NoError(t, err)
		assert.NotNil(t, services)
	})

	models.NewCuisBuilder().Build()

	t.Run("Invalid Config - Empty BaseUrl", func(t *testing.T) {
		cfg := siat.Config{
			Token:          "test",
			Nit:            123456789,
			CodigoSistema:  "test",
			CodigoAmbiente: siat.AmbientePruebas,
			BaseURL:        "",
		}
		services, err := siat.New(cfg)
		assert.Error(t, err)
		assert.Nil(t, services)
		assert.Equal(t, "BaseURL es obligatorio", err.Error())
	})
}

func TestConfig_SignXML(t *testing.T) {
	t.Run("No credentials configured", func(t *testing.T) {
		cfg := siat.Config{}
		xmlBytes := []byte("<root></root>")
		_, err := cfg.SignXML(xmlBytes)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no se configuraron credenciales válidas")
	})

	t.Run("P12 file not found", func(t *testing.T) {
		cfg := siat.Config{
			CredentialSign: siat.NewP12Credential("nonexistent.p12", "password"),
		}
		xmlBytes := []byte("<root></root>")
		_, err := cfg.SignXML(xmlBytes)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error al leer archivo P12")
	})
}

func TestCredencialFirma_Types(t *testing.T) {
	t.Run("Generic constructors setting types", func(t *testing.T) {
		cfPemBytes := siat.NewPEMCredential([]byte("cert"), []byte("key"))
		assert.Equal(t, "PEM", cfPemBytes.GetType())

		cfP12Bytes := siat.NewP12Credential([]byte("p12"), "password")
		assert.Equal(t, "P12", cfP12Bytes.GetType())
	})
}

func TestSiatError_Mensajes(t *testing.T) {
	err := siat.NewSiatError(1000, "Error de prueba")
	err.Mensajes = []siat.MensajeServicio{
		{Codigo: 1000, Descripcion: "Descripción 1"},
		{Codigo: 2005, Descripcion: "Advertencia 1"}, // Un warning
		{Codigo: 3008, Descripcion: "Advertencia 2"}, // CodeWarnCuisExpira (warning)
	}

	assert.Equal(t, 1000, err.SiatCode)
	assert.Len(t, err.Mensajes, 3)
	assert.True(t, err.HasCode(1000))
	assert.True(t, err.HasCode(2005))
	assert.False(t, err.HasCode(9999))

	warnings := err.GetWarnings()
	assert.Len(t, warnings, 2)
	assert.Equal(t, 2005, warnings[0].Codigo)
	assert.Equal(t, 3008, warnings[1].Codigo)
}
