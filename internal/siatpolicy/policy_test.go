package siatpolicy_test

import (
	"testing"

	"github.com/ron86i/go-siat/v3/internal/siatpolicy"
)

func TestValidateDestinationUsaServicioSectorYModalidad(t *testing.T) {
	for _, tc := range []struct {
		name    string
		service string
		sector  int
		mode    int
		wantErr bool
	}{
		{name: "compra venta", service: siatpolicy.ServicioCompraVenta, sector: 1, mode: 1},
		{name: "compra venta bonificaciones", service: siatpolicy.ServicioCompraVenta, sector: 35, mode: 2},
		{name: "servicio basico", service: siatpolicy.ServicioBasico, sector: 13, mode: 1},
		{name: "ajuste no va por electronica general", service: siatpolicy.ServicioElectronica, sector: 24, mode: 1, wantErr: true},
		{name: "compra venta no va por computarizada general", service: siatpolicy.ServicioComputarizada, sector: 1, mode: 2, wantErr: true},
		{name: "basico no va por electronica general", service: siatpolicy.ServicioElectronica, sector: 13, mode: 1, wantErr: true},
		{name: "sector 52 solo electronica", service: siatpolicy.ServicioElectronica, sector: 52, mode: 1},
		{name: "sector 52 rechaza computarizada", service: siatpolicy.ServicioComputarizada, sector: 52, mode: 2, wantErr: true},
		{name: "modalidad no coincide", service: siatpolicy.ServicioElectronica, sector: 2, mode: 2, wantErr: true},
		{name: "sector inexistente", service: siatpolicy.ServicioElectronica, sector: 47, mode: 1, wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := siatpolicy.ValidateDestination(tc.service, tc.sector, tc.mode)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ValidateDestination() error = %v, wantErr %t", err, tc.wantErr)
			}
		})
	}
}
