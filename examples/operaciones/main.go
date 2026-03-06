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

	operService := s.Operaciones

	ctx := context.Background()
	cfg := config.Config{Token: "TU_TOKEN_API"}

	// 2. Consultar Puntos de Venta
	// Nota: El resultado es una interfaz opaca ConsultaPuntoVentaRequest
	consultaReq := models.Operaciones.NewConsultaPuntoVentaRequest().
		WithCodigoAmbiente(1).
		WithNit(123456789).
		WithCodigoSistema("ABC123DEF").
		WithCuis("C2FC682B").
		WithCodigoSucursal(0).
		Build()

	resp, err := operService.ConsultaPuntoVenta(ctx, cfg, consultaReq)
	if err != nil {
		log.Fatalf("Error en consulta: %v", err)
	}

	if resp != nil && resp.Body.Content.Respuesta.Transaccion {
		fmt.Println("Consulta de puntos de venta exitosa:")
		for _, pv := range resp.Body.Content.Respuesta.ListaPuntosVentas {
			fmt.Printf("- [%d] %s (Tipo: %d)\n", pv.CodigoPuntoVenta, pv.NombrePuntoVenta, pv.CodigoPuntoVenta)
		}
	}

	// 3. Registro de un nuevo Punto de Venta
	registroReq := models.Operaciones.NewRegistroPuntoVentaRequest().
		WithCodigoAmbiente(1).
		WithNit(123456789).
		WithCodigoSistema("ABC123DEF").
		WithCuis("C2FC682B").
		WithCodigoSucursal(0).
		WithCodigoTipoPuntoVenta(2). // p.ej., Cajero
		WithNombrePuntoVenta("Caja 01").
		WithDescripcion("Punto de venta para pruebas").
		Build()

	regResp, err := operService.RegistroPuntoVenta(ctx, cfg, registroReq)
	if err == nil && regResp.Body.Content.Respuesta.Transaccion {
		fmt.Printf("Nuevo punto de venta registrado: %d\n", regResp.Body.Content.Respuesta.CodigoPuntoVenta)
	}
}
