package models

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/documento_ajuste"
	"github.com/ron86i/go-siat/pkg/utils"
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

// -- Builders a nivel de paquete --

func NewRecepcionDocumentoAjusteBuilder() *recepcionDocumentoAjusteBuilder {
	return &recepcionDocumentoAjusteBuilder{
		request: &documento_ajuste.RecepcionDocumentoAjuste{},
	}
}

func NewAnulacionDocumentoAjusteBuilder() *anulacionDocumentoAjusteBuilder {
	return &anulacionDocumentoAjusteBuilder{
		request: &documento_ajuste.AnulacionDocumentoAjuste{},
	}
}

func NewReversionAnulacionDocumentoAjusteBuilder() *reversionAnulacionDocumentoAjusteBuilder {
	return &reversionAnulacionDocumentoAjusteBuilder{
		request: &documento_ajuste.ReversionAnulacionDocumentoAjuste{},
	}
}

func NewVerificacionEstadoDocumentoAjusteBuilder() *verificacionEstadoDocumentoAjusteBuilder {
	return &verificacionEstadoDocumentoAjusteBuilder{
		request: &documento_ajuste.VerificacionEstadoDocumentoAjuste{},
	}
}

func NewVerificarComunicacionDocumentoAjusteBuilder() *verificarComunicacionDocumentoAjusteBuilder {
	return &verificarComunicacionDocumentoAjusteBuilder{
		request: &documento_ajuste.VerificarComunicacion{},
	}
}

// -- Implementaciones de Builders --

type recepcionDocumentoAjusteBuilder struct {
	request *documento_ajuste.RecepcionDocumentoAjuste
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoModalidad(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoModalidad = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoDocumentoSector(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoDocumentoSector = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoEmision(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoEmision = v
	return b
}

func (b *recepcionDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *recepcionDocumentoAjusteBuilder {
	b.request.SolicitudRecepcionFactura.CodigoPuntoVenta = v
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

// WithDocumento serializa, firma (si se provee signer), comprime y calcula el hash del documento de ajuste automáticamente,
// mapeando los valores obtenidos en los campos Archivo y HashArchivo de la solicitud.
func (b *recepcionDocumentoAjusteBuilder) WithDocumento(documento any, signer XMLSigner) error {
	xmlData, err := xml.Marshal(documento)
	if err != nil {
		return err
	}

	var xmlToSend = xmlData

	if signer != nil {
		var err error
		xmlToSend, err = signer.SignXML(xmlData)
		if err != nil {
			return err
		}
	}

	hashString, encodedArchivo, err := utils.CompressAndHash(xmlToSend)
	if err != nil {
		return err
	}
	b.request.SolicitudRecepcionFactura.Archivo = encodedArchivo
	b.request.SolicitudRecepcionFactura.HashArchivo = hashString
	return nil
}

func (b *recepcionDocumentoAjusteBuilder) Build() RecepcionDocumentoAjuste {
	return RecepcionDocumentoAjuste{RequestWrapper[documento_ajuste.RecepcionDocumentoAjuste]{request: b.request}}
}

type anulacionDocumentoAjusteBuilder struct {
	request *documento_ajuste.AnulacionDocumentoAjuste
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoDocumentoSector(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoDocumentoSector = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoEmision(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoEmision = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoPuntoVenta = v
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

func (b *anulacionDocumentoAjusteBuilder) WithTipoFacturaDocumento(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.TipoFacturaDocumento = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCuf(v string) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.Cuf = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) WithCodigoMotivo(v int) *anulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioAnulacionDocumentoAjuste.CodigoMotivo = v
	return b
}

func (b *anulacionDocumentoAjusteBuilder) Build() AnulacionDocumentoAjuste {
	return AnulacionDocumentoAjuste{RequestWrapper[documento_ajuste.AnulacionDocumentoAjuste]{request: b.request}}
}

type reversionAnulacionDocumentoAjusteBuilder struct {
	request *documento_ajuste.ReversionAnulacionDocumentoAjuste
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoDocumentoSector(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoDocumentoSector = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoEmision(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoEmision = v
	return b
}

func (b *reversionAnulacionDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *reversionAnulacionDocumentoAjusteBuilder {
	b.request.SolicitudServicioReversionAnulacionDocumentoAjuste.CodigoPuntoVenta = v
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

type verificacionEstadoDocumentoAjusteBuilder struct {
	request *documento_ajuste.VerificacionEstadoDocumentoAjuste
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoDocumentoSector(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoDocumentoSector = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoEmision(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoEmision = v
	return b
}

func (b *verificacionEstadoDocumentoAjusteBuilder) WithCodigoPuntoVenta(v int) *verificacionEstadoDocumentoAjusteBuilder {
	b.request.SolicitudServicioVerificacionEstado.CodigoPuntoVenta = v
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

type verificarComunicacionDocumentoAjusteBuilder struct {
	request *documento_ajuste.VerificarComunicacion
}

func (b *verificarComunicacionDocumentoAjusteBuilder) Build() VerificarComunicacionDocumentoAjuste {
	return VerificarComunicacionDocumentoAjuste{RequestWrapper[documento_ajuste.VerificarComunicacion]{request: b.request}}
}
