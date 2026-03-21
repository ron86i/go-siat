package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// --- Interfaces opacas para restringir el acceso a los atributos ---

// AnulacionFacturaComputarizada representa la solicitud opaca para anular una factura en la modalidad computarizada.
type AnulacionFacturaComputarizada struct {
	requestWrapper[facturacion.AnulacionFactura]
}

// RecepcionFacturaComputarizada representa la solicitud opaca para la recepción de una factura computarizada.
type RecepcionFacturaComputarizada struct {
	requestWrapper[facturacion.RecepcionFactura]
}

// VerificarComunicacionComputarizada representa la solicitud opaca para verificar la conexión con el SIAT.
type VerificarComunicacionComputarizada struct {
	requestWrapper[facturacion.VerificarComunicacion]
}

// ReversionAnulacionFacturaComputarizada representa la solicitud opaca para la reversión de anulación de una factura computarizada.
type ReversionAnulacionFacturaComputarizada struct {
	requestWrapper[facturacion.ReversionAnulacionFactura]
}

type RecepcionPaqueteFacturaComputarizada struct {
	requestWrapper[facturacion.RecepcionPaqueteFactura]
}

type ValidacionRecepcionPaqueteFacturaComputarizada struct {
	requestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

type RecepcionMasivaFacturaComputarizada struct {
	requestWrapper[facturacion.RecepcionMasivaFactura]
}

type RecepcionAnexosSuministroEnergiaComputarizada struct {
	requestWrapper[facturacion.RecepcionAnexosSuministroEnergia]
}

type SuministroEnergiaAnexoComputarizada struct {
	requestWrapper[facturacion.SuministroEnergiaAnexo]
}

type ValidacionRecepcionMasivaFacturaComputarizada struct {
	requestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

type VerificacionEstadoFacturaComputarizada struct {
	requestWrapper[facturacion.VerificacionEstadoFactura]
}

// --- Namespace ---

type computarizadaNamespace struct{}

// Computarizada expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func Computarizada() computarizadaNamespace {
	return computarizadaNamespace{}
}

// --- Constructores de Builders ---

// NewRecepcionFacturaBuilder crea un constructor para la solicitud de recepción de factura.
func (computarizadaNamespace) NewRecepcionFacturaBuilder() *recepcionFacturaComputarizadaBuilder {
	return &recepcionFacturaComputarizadaBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

// NewAnulacionFacturaBuilder crea un constructor para la solicitud de anulación de factura.
func (computarizadaNamespace) NewAnulacionFacturaBuilder() *anulacionFacturaComputarizadaBuilder {
	return &anulacionFacturaComputarizadaBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

// NewVerificarComunicacionBuilder crea un constructor para la solicitud de verificación de comunicación.
func (computarizadaNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionComputarizadaBuilder {
	return &verificarComunicacionComputarizadaBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

func (computarizadaNamespace) NewReversionAnulacionFacturaBuilder() *reversionAnulacionFacturaComputarizadaBuilder {
	return &reversionAnulacionFacturaComputarizadaBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

func (computarizadaNamespace) NewRecepcionPaqueteFacturaBuilder() *recepcionPaqueteFacturaComputarizadaBuilder {
	return &recepcionPaqueteFacturaComputarizadaBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionPaqueteFacturaBuilder crea un constructor para la solicitud de validación de paquete de factura.
func (computarizadaNamespace) NewValidacionRecepcionPaqueteFacturaBuilder() *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	return &validacionRecepcionPaqueteFacturaComputarizadaBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionMasivaFacturaBuilder crea un constructor para la solicitud de validación de recepción masiva de factura.
func (computarizadaNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	return &validacionRecepcionMasivaFacturaComputarizadaBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

// NewVerificacionEstadoFacturaBuilder crea un constructor para la solicitud de verificación de estado de factura.
func (computarizadaNamespace) NewVerificacionEstadoFacturaBuilder() *verificacionEstadoFacturaComputarizadaBuilder {
	return &verificacionEstadoFacturaComputarizadaBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

// NewRecepcionMasivaFacturaBuilder crea un constructor para la solicitud de recepción masiva de factura.
func (computarizadaNamespace) NewRecepcionMasivaFacturaBuilder() *recepcionMasivaFacturaComputarizadaBuilder {
	return &recepcionMasivaFacturaComputarizadaBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

func (computarizadaNamespace) NewRecepcionAnexosSuministroEnergiaBuilder() *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	return &recepcionAnexosSuministroEnergiaComputarizadaBuilder{
		request: &facturacion.RecepcionAnexosSuministroEnergia{},
	}
}

func (computarizadaNamespace) NewSuministroEnergiaAnexoBuilder() *suministroEnergiaAnexoComputarizadaBuilder {
	return &suministroEnergiaAnexoComputarizadaBuilder{
		request: &facturacion.SuministroEnergiaAnexo{},
	}
}

// --- Implementaciones de Builders ---

type recepcionAnexosSuministroEnergiaComputarizadaBuilder struct {
	request *facturacion.RecepcionAnexosSuministroEnergia
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoEmision = codigoEmision
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoModalidad = codigoModalidad
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoSistema = codigoSistema
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoSucursal = codigoSucursal
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCufd(cufd string) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Cufd = cufd
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithCuis(cuis string) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Cuis = cuis
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithNit(nit int64) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Nit = nit
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithTipoFacturaDocumento(tipo int) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.TipoFacturaDocumento = tipo
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) AddAnexos(anexos ...SuministroEnergiaAnexoComputarizada) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	for _, anexo := range anexos {
		if anexo.request != nil {
			b.request.SolicitudRecepcionSuministroAnexos.AnexosList = append(b.request.SolicitudRecepcionSuministroAnexos.AnexosList, *anexo.request)
		}
	}
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) WithGiftCard(giftCard int64) *recepcionAnexosSuministroEnergiaComputarizadaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.GiftCard = giftCard
	return b
}

func (b *recepcionAnexosSuministroEnergiaComputarizadaBuilder) Build() RecepcionAnexosSuministroEnergiaComputarizada {
	return RecepcionAnexosSuministroEnergiaComputarizada{requestWrapper[facturacion.RecepcionAnexosSuministroEnergia]{request: b.request}}
}

type suministroEnergiaAnexoComputarizadaBuilder struct {
	request *facturacion.SuministroEnergiaAnexo
}

func (b *suministroEnergiaAnexoComputarizadaBuilder) WithCufFactSuministro(cuf string) *suministroEnergiaAnexoComputarizadaBuilder {
	b.request.CufFactSuministro = cuf
	return b
}

func (b *suministroEnergiaAnexoComputarizadaBuilder) WithFechaRecarga(fecha time.Time) *suministroEnergiaAnexoComputarizadaBuilder {
	b.request.FechaRecarga = datatype.NewTimeSiat(fecha)
	return b
}

func (b *suministroEnergiaAnexoComputarizadaBuilder) WithMontoRecarga(monto float64) *suministroEnergiaAnexoComputarizadaBuilder {
	b.request.MontoRecarga = monto
	return b
}

func (b *suministroEnergiaAnexoComputarizadaBuilder) Build() SuministroEnergiaAnexoComputarizada {
	return SuministroEnergiaAnexoComputarizada{requestWrapper[facturacion.SuministroEnergiaAnexo]{request: b.request}}
}

type recepcionMasivaFacturaComputarizadaBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCufd(cufd string) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCuis(cuis string) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithNit(nit int64) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithArchivo(archivo string) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo = archivo
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithHashArchivo(hashArchivo string) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.HashArchivo = hashArchivo
	return b
}
func (b *recepcionMasivaFacturaComputarizadaBuilder) WithCantidadFacturas(cantidadFacturas int) *recepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidadFacturas
	return b
}

func (b *recepcionMasivaFacturaComputarizadaBuilder) Build() RecepcionMasivaFacturaComputarizada {
	return RecepcionMasivaFacturaComputarizada{requestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}

type verificacionEstadoFacturaComputarizadaBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCufd(cufd string) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCuis(cuis string) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithNit(nit int64) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) WithCuf(cuf string) *verificacionEstadoFacturaComputarizadaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

func (b *verificacionEstadoFacturaComputarizadaBuilder) Build() VerificacionEstadoFacturaComputarizada {
	return VerificacionEstadoFacturaComputarizada{requestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

type validacionRecepcionMasivaFacturaComputarizadaBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCufd(cufd string) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCuis(cuis string) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithNit(nit int64) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) WithCodigoRecepcion(codigoRecepcion string) *validacionRecepcionMasivaFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *validacionRecepcionMasivaFacturaComputarizadaBuilder) Build() ValidacionRecepcionMasivaFacturaComputarizada {
	return ValidacionRecepcionMasivaFacturaComputarizada{requestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}

type recepcionPaqueteFacturaComputarizadaBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCufd(cufd string) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCuis(cuis string) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithNit(nit int64) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithArchivo(archivo string) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Archivo = archivo
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithHashArchivo(hashArchivo string) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.HashArchivo = hashArchivo
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCafc(cafc string) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = cafc
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCantidadFacturas(cantidadFacturas int) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidadFacturas
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) WithCodigoEvento(codigoEvento int64) *recepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = codigoEvento
	return b
}

func (b *recepcionPaqueteFacturaComputarizadaBuilder) Build() RecepcionPaqueteFacturaComputarizada {
	return RecepcionPaqueteFacturaComputarizada{requestWrapper[facturacion.RecepcionPaqueteFactura]{request: b.request}}
}

type reversionAnulacionFacturaComputarizadaBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCufd(cufd string) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCuis(cuis string) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithNit(nit int64) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) WithCuf(cuf string) *reversionAnulacionFacturaComputarizadaBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

func (b *reversionAnulacionFacturaComputarizadaBuilder) Build() ReversionAnulacionFacturaComputarizada {
	return ReversionAnulacionFacturaComputarizada{requestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

type verificarComunicacionComputarizadaBuilder struct {
	request *facturacion.VerificarComunicacion
}

// Build finaliza la construcción de la solicitud de verificación de comunicación.
func (b *verificarComunicacionComputarizadaBuilder) Build() VerificarComunicacionComputarizada {
	return VerificarComunicacionComputarizada{requestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

type anulacionFacturaComputarizadaBuilder struct {
	request *facturacion.AnulacionFactura
}

// WithCodigoAmbiente configura el código de ambiente (1 para Producción, 2 para Pruebas).
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector configura el código del documento sector (ej. 1 para Compra Venta, 10 para Dutty Free).
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision configura el tipo de emisión (1 para Online, 2 para Offline).
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad configura la modalidad de facturación (1 para Electrónica, 2 para Computarizada).
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta configura el código del punto de venta emisor.
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema configura el código del sistema registrado en el SIAT.
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal configura el código de la sucursal emisora.
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd configura el Código Único de Facturación Diaria vigente al momento de la emisión.
func (b *anulacionFacturaComputarizadaBuilder) WithCufd(cufd string) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

// WithCuis configura el Código Único de Inicio de Sistemas vigente.
func (b *anulacionFacturaComputarizadaBuilder) WithCuis(cuis string) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

// WithNit configura el NIT del emisor.
func (b *anulacionFacturaComputarizadaBuilder) WithNit(nit int64) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento configura el tipo de factura (1 para Con crédito fiscal, 2 para Sin crédito fiscal).
func (b *anulacionFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithCuf configura el Código Único de Factura de la factura que se desea anular.
func (b *anulacionFacturaComputarizadaBuilder) WithCuf(cuf string) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

// WithCodigoMotivo configura el código del motivo de anulación según el catálogo del SIAT (ej. 1 para Factura Mal Emitida).
func (b *anulacionFacturaComputarizadaBuilder) WithCodigoMotivo(codigoMotivo int) *anulacionFacturaComputarizadaBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = codigoMotivo
	return b
}

// Build finaliza la construcción de la solicitud de anulación de factura.
func (b *anulacionFacturaComputarizadaBuilder) Build() AnulacionFacturaComputarizada {
	return AnulacionFacturaComputarizada{requestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

type recepcionFacturaComputarizadaBuilder struct {
	request *facturacion.RecepcionFactura
}

// WithCodigoAmbiente configura el código de ambiente (1 para Producción, 2 para Pruebas).
func (b *recepcionFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector configura el código del documento sector (ej. 1 para Compra Venta, 10 para Dutty Free).
func (b *recepcionFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision configura el tipo de emisión.
func (b *recepcionFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad configura la modalidad de facturación.
func (b *recepcionFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta configura el punto de venta emisor.
func (b *recepcionFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema configura el código del sistema registrado.
func (b *recepcionFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal configura la sucursal emisora.
func (b *recepcionFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd configura el CUFD vigente.
func (b *recepcionFacturaComputarizadaBuilder) WithCufd(cufd string) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

// WithCuis configura el CUIS vigente.
func (b *recepcionFacturaComputarizadaBuilder) WithCuis(cuis string) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

// WithNit configura el NIT del emisor.
func (b *recepcionFacturaComputarizadaBuilder) WithNit(nit int64) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}

// WithTipoFacturaDocumento configura el tipo de factura.
func (b *recepcionFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithArchivo configura el archivo XML de la factura comprimido en GZIP y codificado en Base64.
func (b *recepcionFacturaComputarizadaBuilder) WithArchivo(archivo string) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

// WithFechaEnvio configura la fecha y hora de envío al SIAT.
func (b *recepcionFacturaComputarizadaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

// WithHashArchivo configura el hash SHA256 del archivo XML original.
func (b *recepcionFacturaComputarizadaBuilder) WithHashArchivo(hashArchivo string) *recepcionFacturaComputarizadaBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hashArchivo
	return b
}

// Build finaliza la construcción de la solicitud de recepción de factura.
func (b *recepcionFacturaComputarizadaBuilder) Build() RecepcionFacturaComputarizada {
	return RecepcionFacturaComputarizada{requestWrapper[facturacion.RecepcionFactura]{request: b.request}}
}

type validacionRecepcionPaqueteFacturaComputarizadaBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithNit(nit int64) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}
func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) WithCodigoRecepcion(codigoRecepcion string) *validacionRecepcionPaqueteFacturaComputarizadaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *validacionRecepcionPaqueteFacturaComputarizadaBuilder) Build() ValidacionRecepcionPaqueteFacturaComputarizada {
	return ValidacionRecepcionPaqueteFacturaComputarizada{requestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]{request: b.request}}
}
