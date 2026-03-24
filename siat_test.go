package siat_test

import (
	"testing"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
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
	t.Run("Valid BaseUrl", func(t *testing.T) {
		services, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
		assert.NoError(t, err)
		assert.NotNil(t, services)

	})
	models.Codigos().NewCuisBuilder().Build()

	t.Run("Empty BaseUrl", func(t *testing.T) {
		services, err := siat.New("", nil)
		assert.Error(t, err)
		assert.Nil(t, services)
		assert.Equal(t, "baseUrl is empty", err.Error())
	})
}
