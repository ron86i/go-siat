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
	// 1. Inicializar el servicio SIAT
	s, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	sincService := s.Sincronizacion

	ctx := context.Background()
	cfg := config.Config{Token: "TU_TOKEN_API"}

	// 2. Sincronizar Actividades
	actReq := models.Sincronizacion.NewSincronizarActividadesRequest().
		WithCodigoAmbiente(1).
		WithNit(123456789).
		WithCodigoSistema("ABC123DEF").
		WithCuis("C2FC682B").
		Build()

	actResp, err := sincService.SincronizarActividades(ctx, cfg, actReq)
	if err == nil && actResp.RespuestaListaActividades.Transaccion {
		fmt.Println("Actividades sincronizadas correctamente.")
	}

	// 3. Sincronizar Lista de Productos y Servicios
	prodReq := models.Sincronizacion.NewSincronizarListaProductosServiciosRequest().
		WithCodigoAmbiente(1).
		WithNit(123456789).
		WithCodigoSistema("ABC123DEF").
		WithCuis("C2FC682B").
		Build()

	prodResp, err := sincService.SincronizarListaProductosServicios(ctx, cfg, prodReq)
	if err == nil && prodResp.RespuestaListaProductos.Transaccion {
		fmt.Printf("Se obtuvieron %d productos/servicios.\n", len(prodResp.RespuestaListaProductos.ListaCodigos))
	}
}
