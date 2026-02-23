package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/facturacion/codigos"
)

// Definimos alias para que el usuario pueda VER los tipos
type VerificarNit = codigos.VerificarNit
type SolicitudVerificarNit = codigos.SolicitudVerificarNit
type VerificarNitResponse = codigos.VerificarNitResponse

type Cuis = codigos.Cuis
type SolicitudCuis = codigos.SolicitudCuis
type CuisResponse = codigos.CuisResponse

type Cufd = codigos.Cufd
type SolicitudCufd = codigos.SolicitudCufd
type CufdResponse = codigos.CufdResponse

type CuisMasivo = codigos.CuisMasivo
type SolicitudCuisMasivoSistemas = codigos.SolicitudCuisMasivoSistemas
type CuisMasivoResponse = codigos.CuisMasivoResponse

type CufdMasivo = codigos.CufdMasivo
type SolicitudCufdMasivo = codigos.SolicitudCufdMasivo
type CufdMasivoResponse = codigos.CufdMasivoResponse

type NotificaCertificadoRevocado = codigos.NotificaCertificadoRevocado
type SolicitudNotifcaRevocado = codigos.SolicitudNotifcaRevocado
type NotificaCertificadoRevocadoResponse = codigos.NotificaCertificadoRevocadoResponse

type VerificarComunicacionCodigos = codigos.VerificarComunicacion
type VerificarComunicacionCodigosResponse = codigos.VerificarComunicacionResponse

type codigosNamespace struct{}

var Codigos = codigosNamespace{}

func (codigosNamespace) NewVerificarNitRequest(ambiente, modalidad, sucursal int, cuis, sistema string, nit, nitVerif int64) *VerificarNit {
	return &codigos.VerificarNit{
		SolicitudVerificarNit: codigos.SolicitudVerificarNit{
			CodigoAmbiente:      ambiente,
			CodigoModalidad:     modalidad,
			CodigoSistema:       sistema,
			CodigoSucursal:      sucursal,
			Cuis:                cuis,
			NIT:                 nit,
			NitParaVerificacion: nitVerif,
		},
	}
}

func (codigosNamespace) NewCuisRequest(ambiente, modalidad, puntoVenta, sucursal int, sistema string, nit int64) *Cuis {
	return &codigos.Cuis{
		SolicitudCuis: codigos.SolicitudCuis{
			CodigoAmbiente:   ambiente,
			CodigoModalidad:  modalidad,
			CodigoPuntoVenta: puntoVenta,
			CodigoSistema:    sistema,
			CodigoSucursal:   sucursal,
			NIT:              nit,
		},
	}
}

func (codigosNamespace) NewCufdRequest(ambiente, modalidad, sucursal int, puntoVenta *int, cuis, sistema string, nit int64) *Cufd {
	return &codigos.Cufd{
		SolicitudCufd: codigos.SolicitudCufd{
			CodigoAmbiente:   ambiente,
			CodigoModalidad:  modalidad,
			CodigoPuntoVenta: puntoVenta,
			CodigoSistema:    sistema,
			CodigoSucursal:   sucursal,
			Cuis:             cuis,
			NIT:              nit,
		},
	}
}

func (codigosNamespace) NewCuisMasivoRequest(ambiente, modalidad int, sistema string, nit int64, datos []codigos.SolicitudListaCuisDto) *CuisMasivo {
	return &codigos.CuisMasivo{
		SolicitudCuisMasivoSistemas: codigos.SolicitudCuisMasivoSistemas{
			CodigoAmbiente:  ambiente,
			CodigoModalidad: modalidad,
			CodigoSistema:   sistema,
			NIT:             nit,
			DatosSolicitud:  datos,
		},
	}
}

func (codigosNamespace) NewCufdMasivoRequest(ambiente, modalidad int, sistema string, nit int64, datos []codigos.SolicitudListaCufdDto) *CufdMasivo {
	return &codigos.CufdMasivo{
		SolicitudCufdMasivo: codigos.SolicitudCufdMasivo{
			CodigoAmbiente:  ambiente,
			CodigoModalidad: modalidad,
			CodigoSistema:   sistema,
			Nit:             nit,
			DatosSolicitud:  datos,
		},
	}
}

func (codigosNamespace) NewNotificaCertificadoRevocadoRequest(ambiente, sucursal int, cuis, sistema string, nit int64, certificado, razon string, fecha *time.Time) *NotificaCertificadoRevocado {
	return &codigos.NotificaCertificadoRevocado{
		SolicitudNotificaRevocado: codigos.SolicitudNotifcaRevocado{
			CodigoAmbiente:  ambiente,
			CodigoSistema:   sistema,
			CodigoSucursal:  sucursal,
			Cuis:            cuis,
			NIT:             nit,
			Certificado:     certificado,
			RazonRevocacion: razon,
			FechaRevocacion: fecha,
		},
	}
}

func (codigosNamespace) NewVerificarComunicacionCodigos() *VerificarComunicacionCodigos {
	return &codigos.VerificarComunicacion{}
}
