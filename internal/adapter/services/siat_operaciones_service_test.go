package services_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestConsultaPuntoVenta(t *testing.T) {
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

	service := siatClient.Operaciones()

	req := models.NewConsultaPuntoVentaBuilder().
		WithCodigoSucursal(0).
		WithCuis("197C8240").
		Build()

	resp, err := service.ConsultaPuntoVenta(context.Background(), req)
	if err == nil && assert.NotNil(t, resp) {
		assert.NotNil(t, resp.Body.Content)
	}
}
