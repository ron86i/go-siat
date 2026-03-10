package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

func main() {
	// 1. Inicializar el servicio siat
	s, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	codigosService := s.Codigos()

	// 2. Construir solicitud de CUIS usando el Builder
	// El resultado de Build() es una interfaz opaca CuisRequest
	cuisReq := models.Codigos.NewCuisRequest().
		WithCodigoAmbiente(1).
		WithCodigoModalidad(1).
		WithCodigoPuntoVenta(0).
		WithCodigoSucursal(0).
		WithCodigoSistema("ABC123DEF").
		WithNit(123456789).
		Build()

	ctx := context.Background()
	cfg := config.Config{Token: "TU_TOKEN_API"}

	// 3. Solicitar CUIS
	resp, err := codigosService.SolicitudCuis(ctx, cfg, cuisReq)
	if err != nil {
		log.Fatalf("Error al solicitar CUIS: %v", err)
	}

	if resp != nil && resp.Body.Content.RespuestaCuis.Transaccion {
		fmt.Printf("CUIS obtenido: %s\n", resp.Body.Content.RespuestaCuis.Codigo)

		// 4. Con el CUIS, podemos solicitar un CUFD
		cufdReq := models.Codigos.NewCufdRequest().
			WithCodigoAmbiente(1).
			WithCodigoModalidad(1).
			WithCodigoPuntoVenta(0).
			WithCodigoSucursal(0).
			WithCodigoSistema("ABC123DEF").
			WithNit(123456789).
			WithCuis(resp.Body.Content.RespuestaCuis.Codigo).
			Build()

		cufdResp, err := codigosService.SolicitudCufd(ctx, cfg, cufdReq)
		if err == nil && cufdResp.Body.Content.RespuestaCufd.Transaccion {
			fmt.Printf("CUFD obtenido: %s (Control: %s)\n",
				cufdResp.Body.Content.RespuestaCufd.Codigo,
				cufdResp.Body.Content.RespuestaCufd.CodigoControl)
		}
	}
}
