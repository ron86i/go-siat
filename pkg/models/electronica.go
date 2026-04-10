package models

import (
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
)

// AnulacionFacturaElectronica representa la solicitud opaca para anular una factura en la modalidad electrónica.
type AnulacionFacturaElectronica struct {
	RequestWrapper[facturacion.AnulacionFactura]
}

// RecepcionFacturaElectronica representa la solicitud opaca para la recepción de una factura electrónica.
type RecepcionFacturaElectronica struct {
	RequestWrapper[facturacion.RecepcionFactura]
}

// VerificarComunicacionElectronica representa la solicitud opaca para verificar la conexión con el SIAT.
type VerificarComunicacionElectronica struct {
	RequestWrapper[facturacion.VerificarComunicacion]
}

// ReversionAnulacionFacturaElectronica representa la solicitud opaca para la reversión de anulación de una factura electrónica.
type ReversionAnulacionFacturaElectronica struct {
	RequestWrapper[facturacion.ReversionAnulacionFactura]
}

type RecepcionPaqueteFacturaElectronica struct {
	RequestWrapper[facturacion.RecepcionPaqueteFactura]
}

type ValidacionRecepcionPaqueteFacturaElectronica struct {
	RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

type RecepcionMasivaFacturaElectronica struct {
	RequestWrapper[facturacion.RecepcionMasivaFactura]
}

type RecepcionAnexosSuministroEnergiaElectronica struct {
	RequestWrapper[facturacion.RecepcionAnexosSuministroEnergia]
}

type SuministroEnergiaAnexoElectronica struct {
	RequestWrapper[facturacion.SuministroEnergiaAnexo]
}

type ValidacionRecepcionMasivaFacturaElectronica struct {
	RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

type VerificacionEstadoFacturaElectronica struct {
	RequestWrapper[facturacion.VerificacionEstadoFactura]
}

// --- Namespace ---

type electronicaNamespace struct{}

// Electronica expone utilidades y constructores de solicitudes para el módulo de Facturación del SIAT.
func Electronica() electronicaNamespace {
	return electronicaNamespace{}
}

// --- Constructores de Builders ---

// NewRecepcionFacturaBuilder crea un constructor para la solicitud de recepción de factura.
func (electronicaNamespace) NewRecepcionFacturaBuilder() *recepcionFacturaElectronicaBuilder {
	return &recepcionFacturaElectronicaBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

// NewAnulacionFacturaBuilder crea un constructor para la solicitud de anulación de factura.
func (electronicaNamespace) NewAnulacionFacturaBuilder() *anulacionFacturaElectronicaBuilder {
	return &anulacionFacturaElectronicaBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

// NewVerificarComunicacionBuilder crea un constructor para la solicitud de verificación de comunicación.
func (electronicaNamespace) NewVerificarComunicacionBuilder() *verificarComunicacionElectronicaBuilder {
	return &verificarComunicacionElectronicaBuilder{
		request: &facturacion.VerificarComunicacion{},
	}
}

func (electronicaNamespace) NewReversionAnulacionFacturaBuilder() *reversionAnulacionFacturaElectronicaBuilder {
	return &reversionAnulacionFacturaElectronicaBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

func (electronicaNamespace) NewRecepcionPaqueteFacturaBuilder() *recepcionPaqueteFacturaElectronicaBuilder {
	return &recepcionPaqueteFacturaElectronicaBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionPaqueteFacturaBuilder crea un constructor para la solicitud de validación de paquete de factura.
func (electronicaNamespace) NewValidacionRecepcionPaqueteFacturaBuilder() *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	return &validacionRecepcionPaqueteFacturaElectronicaBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

// NewValidacionRecepcionMasivaFacturaBuilder crea un constructor para la solicitud de validación de recepción masiva de factura.
func (electronicaNamespace) NewValidacionRecepcionMasivaFacturaBuilder() *validacionRecepcionMasivaFacturaElectronicaBuilder {
	return &validacionRecepcionMasivaFacturaElectronicaBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
}

// NewVerificacionEstadoFacturaBuilder crea un constructor para la solicitud de verificación de estado de factura.
func (electronicaNamespace) NewVerificacionEstadoFacturaBuilder() *verificacionEstadoFacturaElectronicaBuilder {
	return &verificacionEstadoFacturaElectronicaBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

// NewRecepcionMasivaFacturaBuilder crea un constructor para la solicitud de recepción masiva de factura.
func (electronicaNamespace) NewRecepcionMasivaFacturaBuilder() *recepcionMasivaFacturaElectronicaBuilder {
	return &recepcionMasivaFacturaElectronicaBuilder{
		request: &facturacion.RecepcionMasivaFactura{},
	}
}

func (electronicaNamespace) NewRecepcionAnexosSuministroEnergiaBuilder() *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	return &recepcionAnexosSuministroEnergiaElectronicaBuilder{
		request: &facturacion.RecepcionAnexosSuministroEnergia{},
	}
}

func (electronicaNamespace) NewSuministroEnergiaAnexoBuilder() *suministroEnergiaAnexoElectronicaBuilder {
	return &suministroEnergiaAnexoElectronicaBuilder{
		request: &facturacion.SuministroEnergiaAnexo{},
	}
}

// --- Implementaciones de Builders ---

type recepcionAnexosSuministroEnergiaElectronicaBuilder struct {
	request *facturacion.RecepcionAnexosSuministroEnergia
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCufd(cufd string) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithCuis(cuis string) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithNit(nit int64) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithTipoFacturaDocumento(tipo int) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.TipoFacturaDocumento = tipo
	return b
}

// AddAnexos añade uno o más anexos de suministro de energía a la solicitud.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) AddAnexos(anexos ...SuministroEnergiaAnexoElectronica) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	for _, anexo := range anexos {
		if internal := UnwrapInternalRequest[facturacion.SuministroEnergiaAnexo](anexo); internal != nil {
			b.request.SolicitudRecepcionSuministroAnexos.AnexosList = append(b.request.SolicitudRecepcionSuministroAnexos.AnexosList, *internal)
		}
	}
	return b
}

// WithGiftCard establece el monto de gift card si aplica.
func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) WithGiftCard(giftCard int64) *recepcionAnexosSuministroEnergiaElectronicaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.GiftCard = giftCard
	return b
}

func (b *recepcionAnexosSuministroEnergiaElectronicaBuilder) Build() RecepcionAnexosSuministroEnergiaElectronica {
	return RecepcionAnexosSuministroEnergiaElectronica{RequestWrapper[facturacion.RecepcionAnexosSuministroEnergia]{request: b.request}}
}

type suministroEnergiaAnexoElectronicaBuilder struct {
	request *facturacion.SuministroEnergiaAnexo
}

// WithCufFactSuministro establece el CUF de la factura de suministro asociada.
func (b *suministroEnergiaAnexoElectronicaBuilder) WithCufFactSuministro(cuf string) *suministroEnergiaAnexoElectronicaBuilder {
	b.request.CufFactSuministro = cuf
	return b
}

// WithFechaRecarga establece la fecha de recarga.
func (b *suministroEnergiaAnexoElectronicaBuilder) WithFechaRecarga(fecha time.Time) *suministroEnergiaAnexoElectronicaBuilder {
	b.request.FechaRecarga = datatype.NewTimeSiat(fecha)
	return b
}

// WithMontoRecarga establece el monto de la recarga.
func (b *suministroEnergiaAnexoElectronicaBuilder) WithMontoRecarga(monto float64) *suministroEnergiaAnexoElectronicaBuilder {
	b.request.MontoRecarga = monto
	return b
}

func (b *suministroEnergiaAnexoElectronicaBuilder) Build() SuministroEnergiaAnexoElectronica {
	return SuministroEnergiaAnexoElectronica{RequestWrapper[facturacion.SuministroEnergiaAnexo]{request: b.request}}
}

type recepcionMasivaFacturaElectronicaBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCufd(cufd string) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCuis(cuis string) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithNit(nit int64) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithArchivo establece el archivo XML (tar.gz) en Base64.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithArchivo(archivo string) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de emisión.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

// WithHashArchivo establece el hash del archivo.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithHashArchivo(hashArchivo string) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.HashArchivo = hashArchivo
	return b
}

// WithCantidadFacturas establece la cantidad de facturas.
func (b *recepcionMasivaFacturaElectronicaBuilder) WithCantidadFacturas(cantidadFacturas int) *recepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidadFacturas
	return b
}

func (b *recepcionMasivaFacturaElectronicaBuilder) Build() RecepcionMasivaFacturaElectronica {
	return RecepcionMasivaFacturaElectronica{RequestWrapper[facturacion.RecepcionMasivaFactura]{request: b.request}}
}

type verificacionEstadoFacturaElectronicaBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCufd(cufd string) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCuis(cuis string) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithNit(nit int64) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithCuf establece el CUF de la factura.
func (b *verificacionEstadoFacturaElectronicaBuilder) WithCuf(cuf string) *verificacionEstadoFacturaElectronicaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

func (b *verificacionEstadoFacturaElectronicaBuilder) Build() VerificacionEstadoFacturaElectronica {
	return VerificacionEstadoFacturaElectronica{RequestWrapper[facturacion.VerificacionEstadoFactura]{request: b.request}}
}

type validacionRecepcionMasivaFacturaElectronicaBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCufd(cufd string) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCuis(cuis string) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithNit(nit int64) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithCodigoRecepcion establece el código de recepción para validar.
func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) WithCodigoRecepcion(codigoRecepcion string) *validacionRecepcionMasivaFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *validacionRecepcionMasivaFacturaElectronicaBuilder) Build() ValidacionRecepcionMasivaFacturaElectronica {
	return ValidacionRecepcionMasivaFacturaElectronica{RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]{request: b.request}}
}

type recepcionPaqueteFacturaElectronicaBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCufd(cufd string) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCuis(cuis string) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithNit(nit int64) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithArchivo establece el archivo XML (tar.gz) en Base64.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithArchivo(archivo string) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Archivo = archivo
	return b
}

// WithFechaEnvio establece la fecha y hora de emisión.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

// WithHashArchivo establece el hash del archivo.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithHashArchivo(hashArchivo string) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.HashArchivo = hashArchivo
	return b
}

// WithCafc establece el CAFC si aplica.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCafc(cafc string) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = cafc
	return b
}

// WithCantidadFacturas establece la cantidad de facturas.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCantidadFacturas(cantidadFacturas int) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidadFacturas
	return b
}

// WithCodigoEvento establece el código de evento de contingencia.
func (b *recepcionPaqueteFacturaElectronicaBuilder) WithCodigoEvento(codigoEvento int64) *recepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = codigoEvento
	return b
}

func (b *recepcionPaqueteFacturaElectronicaBuilder) Build() RecepcionPaqueteFacturaElectronica {
	return RecepcionPaqueteFacturaElectronica{RequestWrapper[facturacion.RecepcionPaqueteFactura]{request: b.request}}
}

type reversionAnulacionFacturaElectronicaBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCufd(cufd string) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCuis(cuis string) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithNit(nit int64) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithCuf establece el CUF de la factura a revertir.
func (b *reversionAnulacionFacturaElectronicaBuilder) WithCuf(cuf string) *reversionAnulacionFacturaElectronicaBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

func (b *reversionAnulacionFacturaElectronicaBuilder) Build() ReversionAnulacionFacturaElectronica {
	return ReversionAnulacionFacturaElectronica{RequestWrapper[facturacion.ReversionAnulacionFactura]{request: b.request}}
}

type verificarComunicacionElectronicaBuilder struct {
	request *facturacion.VerificarComunicacion
}

// Build finaliza la construcción de la solicitud de verificación de comunicación.
func (b *verificarComunicacionElectronicaBuilder) Build() VerificarComunicacionElectronica {
	return VerificarComunicacionElectronica{RequestWrapper[facturacion.VerificarComunicacion]{request: b.request}}
}

type anulacionFacturaElectronicaBuilder struct {
	request *facturacion.AnulacionFactura
}

// WithCodigoAmbiente configura el código de ambiente (1 para Producción, 2 para Pruebas).
func (b *anulacionFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector configura el código del documento sector (ej. 1 para Compra Venta, 10 para Dutty Free).
func (b *anulacionFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision configura el tipo de emisión (1 para Online, 2 para Offline).
func (b *anulacionFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad configura la modalidad de facturación (1 para Electrónica).
func (b *anulacionFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta configura el código del punto de venta emisor.
func (b *anulacionFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema configura el código del sistema registrado en el SIAT.
func (b *anulacionFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal configura el código de la sucursal emisora.
func (b *anulacionFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd configura el Código Único de Facturación Diaria vigente al momento de la emisión.
func (b *anulacionFacturaElectronicaBuilder) WithCufd(cufd string) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

// WithCuis configura el Código Único de Inicio de Sistemas vigente.
func (b *anulacionFacturaElectronicaBuilder) WithCuis(cuis string) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

// WithNit configura el NIT del emisor.
func (b *anulacionFacturaElectronicaBuilder) WithNit(nit int64) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

// WithTipoFacturaDocumento configura el tipo de factura (1 para Con crédito fiscal, 2 para Sin crédito fiscal).
func (b *anulacionFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithCuf configura el Código Único de Factura de la factura que se desea anular.
func (b *anulacionFacturaElectronicaBuilder) WithCuf(cuf string) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

// WithCodigoMotivo configura el código del motivo de anulación según el catálogo del SIAT (ej. 1 para Factura Mal Emitida).
func (b *anulacionFacturaElectronicaBuilder) WithCodigoMotivo(codigoMotivo int) *anulacionFacturaElectronicaBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = codigoMotivo
	return b
}

// Build finaliza la construcción de la solicitud de anulación de factura.
func (b *anulacionFacturaElectronicaBuilder) Build() AnulacionFacturaElectronica {
	return AnulacionFacturaElectronica{RequestWrapper[facturacion.AnulacionFactura]{request: b.request}}
}

type recepcionFacturaElectronicaBuilder struct {
	request *facturacion.RecepcionFactura
}

// WithCodigoAmbiente configura el código de ambiente (1 para Producción, 2 para Pruebas).
func (b *recepcionFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector configura el código del documento sector (ej. 1 para Compra Venta, 10 para Dutty Free).
func (b *recepcionFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision configura el tipo de emisión.
func (b *recepcionFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad configura la modalidad de facturación.
func (b *recepcionFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta configura el punto de venta emisor.
func (b *recepcionFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema configura el código del sistema registrado.
func (b *recepcionFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal configura la sucursal emisora.
func (b *recepcionFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd configura el CUFD vigente.
func (b *recepcionFacturaElectronicaBuilder) WithCufd(cufd string) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

// WithCuis configura el CUIS vigente.
func (b *recepcionFacturaElectronicaBuilder) WithCuis(cuis string) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

// WithNit configura el NIT del emisor.
func (b *recepcionFacturaElectronicaBuilder) WithNit(nit int64) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}

// WithTipoFacturaDocumento configura el tipo de factura.
func (b *recepcionFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithArchivo configura el archivo XML de la factura comprimido en GZIP y codificado en Base64.
func (b *recepcionFacturaElectronicaBuilder) WithArchivo(archivo string) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

// WithFechaEnvio configura la fecha y hora de envío al SIAT.
func (b *recepcionFacturaElectronicaBuilder) WithFechaEnvio(fechaEnvio time.Time) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

// WithHashArchivo configura el hash SHA256 del archivo XML original.
func (b *recepcionFacturaElectronicaBuilder) WithHashArchivo(hashArchivo string) *recepcionFacturaElectronicaBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hashArchivo
	return b
}

// Build finaliza la construcción de la solicitud de recepción de factura.
func (b *recepcionFacturaElectronicaBuilder) Build() RecepcionFacturaElectronica {
	return RecepcionFacturaElectronica{RequestWrapper[facturacion.RecepcionFactura]{request: b.request}}
}

type validacionRecepcionPaqueteFacturaElectronicaBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

// WithCodigoAmbiente establece el código de ambiente.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoAmbiente(codigoAmbiente int) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoEmision(codigoEmision int) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

// WithCodigoModalidad establece el código de la modalidad.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoModalidad(codigoModalidad int) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSistema establece el código del sistema.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoSistema(codigoSistema string) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoSucursal(codigoSucursal int) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

// WithCufd establece el CUFD.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCufd(cufd string) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}

// WithCuis establece el CUIS.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCuis(cuis string) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}

// WithNit establece el NIT del emisor.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithNit(nit int64) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// WithCodigoRecepcion establece el código de recepción para validar.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) WithCodigoRecepcion(codigoRecepcion string) *validacionRecepcionPaqueteFacturaElectronicaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigoRecepcion
	return b
}

// Build finaliza la construcción de la solicitud de validación de recepción de paquete.
func (b *validacionRecepcionPaqueteFacturaElectronicaBuilder) Build() ValidacionRecepcionPaqueteFacturaElectronica {
	return ValidacionRecepcionPaqueteFacturaElectronica{RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]{request: b.request}}
}
