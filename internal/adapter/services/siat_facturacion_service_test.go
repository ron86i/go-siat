package services_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestFacturacionService_VerificarComunicacion(t *testing.T) {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		t.Skip("Saltando prueba de integración: .env no encontrado")
	}
	godotenv.Load(".env")

	cfg := siat.Config{
		Token:          os.Getenv("SIAT_TOKEN"),
		Nit:            123456789,
		CodigoSistema:  os.Getenv("SIAT_CODIGO_SISTEMA"),
		CodigoAmbiente: siat.AmbientePruebas,
		BaseURL:        os.Getenv("SIAT_URL"),
		HTTPClient:     &http.Client{},
	}

	siatClient, err := siat.New(cfg)
	if err != nil {
		t.Fatalf("error creating client: %v", err)
	}

	service := siatClient.Computarizada()

	req := models.NewVerificarComunicacionFacturacion()
	resp, err := service.VerificarComunicacion(context.Background(), req)
	if assert.NoError(t, err) && assert.NotNil(t, resp) {
		t.Logf("Respuesta SIAT: %+v", resp.Body.Content)
		assert.True(t, resp.Body.Content.Return.Transaccion)
	}
}
