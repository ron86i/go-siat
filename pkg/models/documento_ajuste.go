package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/documento_ajuste"
)

// -- Interfaces opacas para restringir el acceso a los atributos --

type RecepcionDocumentoAjuste struct {
	RequestWrapper[documento_ajuste.RecepcionDocumentoAjuste]
}

type AnulacionDocumentoAjuste struct {
	RequestWrapper[documento_ajuste.AnulacionDocumentoAjuste]
}

type ReversionAnulacionDocumentoAjuste struct {
	RequestWrapper[documento_ajuste.ReversionAnulacionDocumentoAjuste]
}

type VerificacionEstadoDocumentoAjuste struct {
	RequestWrapper[documento_ajuste.VerificacionEstadoDocumentoAjuste]
}

type VerificarComunicacionDocumentoAjuste struct {
	RequestWrapper[documento_ajuste.VerificarComunicacion]
}

// -- Namespace --

type documentoAjusteNamespace struct{}

// DocumentoAjuste expone utilidades y constructores de solicitudes para el módulo de Documento de Ajuste del SIAT.
func DocumentoAjuste() documentoAjusteNamespace {
	return documentoAjusteNamespace{}
}

// -- Builders --
type recepcionDocumentoAjusteBuilder struct {
	request *documento_ajuste.RecepcionDocumentoAjuste
}

type anulacionDocumentoAjusteBuilder struct {
	request *documento_ajuste.AnulacionDocumentoAjuste
}

type reversionAnulacionDocumentoAjusteBuilder struct {
	request *documento_ajuste.ReversionAnulacionDocumentoAjuste
}

type verificacionEstadoDocumentoAjusteBuilder struct {
	request *documento_ajuste.VerificacionEstadoDocumentoAjuste
}

type verificarComunicacionDocumentoAjusteBuilder struct {
	request *documento_ajuste.VerificarComunicacion
}

func (documentoAjusteNamespace) NewRecepcionDocumentoAjuste() *recepcionDocumentoAjusteBuilder {
	return &recepcionDocumentoAjusteBuilder{
		request: &documento_ajuste.RecepcionDocumentoAjuste{},
	}
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoAmbiente(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoAmbiente = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithDocumentoSector(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoDocumentoSector = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoEmision(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoEmision = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoModalidad(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoModalidad = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoPuntoVenta = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoSistema(v string) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoSistema = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoSucursal(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoSucursal = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCufd(v string) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.Cufd = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCuis(v string) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.Cuis = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithNit(v int64) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.Nit = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithTipoFacturaDocumento(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.TipoFacturaDocumento = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithArchivo(v string) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.Archivo = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithFechaEnvio(v time.Time) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.FechaEnvio = datatype.TimeSiat(v)
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithHashArchivo(v string) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.HashArchivo = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) Build() RecepcionDocumentoAjuste {
	return RecepcionDocumentoAjuste{RequestWrapper[documento_ajuste.RecepcionDocumentoAjuste]{request: b.request}}
}

func (documentoAjusteNamespace) NewAnulacionDocumentoAjuste() *anulacionDocumentoAjusteBuilder {
	return &anulacionDocumentoAjusteBuilder{
		request: &documento_ajuste.AnulacionDocumentoAjuste{},
	}
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoAmbiente(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoAmbiente = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithDocumentoSector(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoDocumentoSector = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoEmision(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoEmision = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoModalidad(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoModalidad = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoPuntoVenta = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoSistema(v string) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoSistema = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoSucursal(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoSucursal = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCufd(v string) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.Cufd = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCuis(v string) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.Cuis = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithNit(v int64) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.Nit = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithTipoFacturaDocumento(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.TipoFacturaDocumento = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) Build() AnulacionDocumentoAjuste {
	return AnulacionDocumentoAjuste{RequestWrapper[documento_ajuste.AnulacionDocumentoAjuste]{request: b.request}}
}

func (documentoAjusteNamespace) NewReversionAnulacionDocumentoAjuste() *reversionAnulacionDocumentoAjusteBuilder {
	return &reversionAnulacionDocumentoAjusteBuilder{
		request: &documento_ajuste.ReversionAnulacionDocumentoAjuste{},
	}
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoAmbiente(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoAmbiente = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithDocumentoSector(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoDocumentoSector = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoEmision(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoEmision = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoModalidad(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoModalidad = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoPuntoVenta = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoSistema(v string) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoSistema = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoSucursal(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoSucursal = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCufd(v string) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.Cufd = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCuis(v string) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.Cuis = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithNit(v int64) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.Nit = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithTipoFacturaDocumento(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.TipoFacturaDocumento = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCuf(v string) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.Cuf = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) Build() ReversionAnulacionDocumentoAjuste {
	return ReversionAnulacionDocumentoAjuste{RequestWrapper[documento_ajuste.ReversionAnulacionDocumentoAjuste]{request: b.request}}
}

func (documentoAjusteNamespace) NewVerificacionEstadoDocumentoAjuste() *verificacionEstadoDocumentoAjusteBuilder {
	return &verificacionEstadoDocumentoAjusteBuilder{
		request: &documento_ajuste.VerificacionEstadoDocumentoAjuste{},
	}
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoAmbiente(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoAmbiente = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithDocumentoSector(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoDocumentoSector = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoEmision(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoEmision = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoModalidad(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoModalidad = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoPuntoVenta = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoSistema(v string) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoSistema = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoSucursal(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoSucursal = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCufd(v string) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.Cufd = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCuis(v string) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.Cuis = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithNit(v int64) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.Nit = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithTipoFacturaDocumento(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.TipoFacturaDocumento = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCuf(v string) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.Cuf = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) Build() VerificacionEstadoDocumentoAjuste {
	return VerificacionEstadoDocumentoAjuste{RequestWrapper[documento_ajuste.VerificacionEstadoDocumentoAjuste]{request: b.request}}
}

func (documentoAjusteNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionDocumentoAjusteBuilder {
	return &verificarComunicacionDocumentoAjusteBuilder{
		request: &documento_ajuste.VerificarComunicacion{},
	}
}

func (b *verificarComunicacionDocumentoAjusteBuilder) Build() VerificarComunicacionDocumentoAjuste {
	return VerificarComunicacionDocumentoAjuste{RequestWrapper[documento_ajuste.VerificarComunicacion]{request: b.request}}
}
