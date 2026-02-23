package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/facturacion/operaciones"
)

// Definimos alias para que el usuario pueda VER los tipos
type RegistroPuntoVenta = operaciones.RegistroPuntoVenta
type SolicitudRegistroPuntoVenta = operaciones.SolicitudRegistroPuntoVenta
type RegistroPuntoVentaResponse = operaciones.RegistroPuntoVentaResponse

type ConsultaPuntoVenta = operaciones.ConsultaPuntoVenta
type SolicitudConsultaPuntoVenta = operaciones.SolicitudConsultaPuntoVenta
type ConsultaPuntoVentaResponse = operaciones.ConsultaPuntoVentaResponse

type CierrePuntoVenta = operaciones.CierrePuntoVenta
type SolicitudCierrePuntoVenta = operaciones.SolicitudCierrePuntoVenta
type CierrePuntoVentaResponse = operaciones.CierrePuntoVentaResponse

type RegistroPuntoVentaComisionista = operaciones.RegistroPuntoVentaComisionista
type SolicitudPuntoVentaComisionista = operaciones.SolicitudPuntoVentaComisionista
type RegistroPuntoVentaComisionistaResponse = operaciones.RegistroPuntoVentaComisionistaResponse

type CierreOperacionesSistema = operaciones.CierreOperacionesSistema
type SolicitudOperaciones = operaciones.SolicitudOperaciones
type CierreOperacionesSistemaResponse = operaciones.CierreOperacionesSistemaResponse

type VerificiarComunicacionOperaciones = operaciones.VerificarComunicacion
type VerificarComunicacionOperacionesResponse = operaciones.VerificarComunicacionResponse

type operacionesNamespace struct{}

var Operaciones = operacionesNamespace{}

func (operacionesNamespace) NewRegistroPuntoVentaRequest(ambiente, modalidad, sucursal int, cuis, sistema string, nit int64) *RegistroPuntoVenta {
	return &operaciones.RegistroPuntoVenta{
		SolicitudRegistroPuntoVenta: operaciones.SolicitudRegistroPuntoVenta{
			CodigoAmbiente:  ambiente,
			CodigoModalidad: modalidad,
			CodigoSistema:   sistema,
			CodigoSucursal:  sucursal,
			Cuis:            cuis,
			Nit:             nit,
		},
	}
}

func (operacionesNamespace) NewConsultaPuntoVentaRequest(ambiente, sucursal int, cuis, sistema string, nit int64) *ConsultaPuntoVenta {
	return &operaciones.ConsultaPuntoVenta{
		SolicitudConsultaPuntoVenta: operaciones.SolicitudConsultaPuntoVenta{
			CodigoAmbiente: ambiente,
			CodigoSistema:  sistema,
			CodigoSucursal: sucursal,
			Cuis:           cuis,
			Nit:            nit,
		},
	}
}

func (operacionesNamespace) NewCierrePuntoVentaRequest(ambiente, sucursal, puntoVenta int, cuis, sistema string, nit int64) *CierrePuntoVenta {
	return &operaciones.CierrePuntoVenta{
		SolicitudCierrePuntoVenta: operaciones.SolicitudCierrePuntoVenta{
			CodigoAmbiente:   ambiente,
			CodigoPuntoVenta: puntoVenta,
			CodigoSistema:    sistema,
			CodigoSucursal:   sucursal,
			Cuis:             cuis,
			Nit:              nit,
		},
	}
}

func (operacionesNamespace) NewRegistroPuntoVentaComisionistaRequest(ambiente, modalidad, sucursal int, cuis, sistema string, nit, nitComisionista int64, nombre, descripcion, contrato string, inicio, fin time.Time) *RegistroPuntoVentaComisionista {
	return &operaciones.RegistroPuntoVentaComisionista{
		SolicitudPuntoVentaComisionista: operaciones.SolicitudPuntoVentaComisionista{
			CodigoAmbiente:   ambiente,
			CodigoModalidad:  modalidad,
			CodigoSistema:    sistema,
			CodigoSucursal:   sucursal,
			Cuis:             cuis,
			Nit:              nit,
			NitComisionista:  nitComisionista,
			NombrePuntoVenta: nombre,
			Descripcion:      descripcion,
			NumeroContrato:   contrato,
			FechaInicio:      inicio,
			FechaFin:         fin,
		},
	}
}

func (operacionesNamespace) NewCierreOperacionesSistemaRequest(ambiente, modalidad, sucursal, puntoVenta int, cuis, sistema string, nit int64) *CierreOperacionesSistema {
	return &operaciones.CierreOperacionesSistema{
		SolicitudOperaciones: operaciones.SolicitudOperaciones{
			CodigoAmbiente:   ambiente,
			CodigoModalidad:  modalidad,
			CodigoPuntoVenta: puntoVenta,
			CodigoSistema:    sistema,
			CodigoSucursal:   sucursal,
			Cuis:             cuis,
			Nit:              nit,
		},
	}
}

func (operacionesNamespace) NewVerificiarComunicacionOperaciones() *VerificiarComunicacionOperaciones {
	return &operaciones.VerificarComunicacion{}
}
