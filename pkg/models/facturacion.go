package models

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v2/pkg/utils"
)

// XMLSigner define la interfaz para realizar la firma de documentos XML
type XMLSigner interface {
	SignXML(xmlBytes []byte) ([]byte, error)
}

// XMLBytesMarshaler permite entregar el XML compacto ya construido de una
// factura. WithFacturasEnLote lo usa en lugar de encoding/xml.Marshal, lo que
// evita reflexión y asignaciones para integraciones de alto volumen.
//
// El XML debe representar el documento sin firma: en modalidad electrónica el
// SDK lo firma antes de incorporarlo al archivo. El llamador no debe modificar
// los bytes mientras se procesa el lote.
type XMLBytesMarshaler interface {
	MarshalXMLBytes() ([]byte, error)
}

// FacturasEnLoteOptions configura el procesamiento de facturas para emisión masiva.
//
// Workers indica la cantidad máxima de documentos que se serializan y firman
// simultáneamente. Con cero, WithFacturasEnLote usa los CPU lógicos disponibles
// menos uno, manteniendo capacidad para el resto de la aplicación. Un valor de
// uno o menor conserva el procesamiento serial. El valor se limita
// automáticamente a la cantidad de facturas del lote.
//
// WithFacturasEnLote acepta nil para usar este comportamiento automático.
//
// Para proteger integraciones existentes, un firmador personalizado se procesa
// serialmente, salvo que implemente ConcurrentXMLSigning y retorne true.
type FacturasEnLoteOptions struct {
	// Workers limita los documentos preparados simultáneamente.
	Workers int
	// MaxFacturas rechaza lotes que superen este valor. Cero no aplica límite.
	MaxFacturas int
	// DestinoArchivo recibe una copia del TAR.GZ ya firmado para reintentos.
	// El SDK no cierra este writer.
	DestinoArchivo io.Writer
	// OnComplete recibe métricas al terminar correctamente el procesamiento.
	OnComplete func(FacturasEnLoteMetrics)
}

// FacturasEnLoteMetrics describe el trabajo realizado al preparar una emisión masiva.
// Las duraciones de serialización y firma son acumuladas por factura, por lo
// que pueden ser mayores que DuracionTotal cuando se usan varios workers.
type FacturasEnLoteMetrics struct {
	CantidadFacturas      int
	Workers               int
	DuracionSerializacion time.Duration
	DuracionFirma         time.Duration
	DuracionEmpaquetado   time.Duration
	DuracionTotal         time.Duration
	TamanoArchivo         int64
}

type concurrentXMLSigner interface {
	ConcurrentXMLSigning() bool
}

// ErrFirmaElectronicaRequerida indica que se intentó preparar una factura
// electrónica sin el firmador XML configurado en modo estricto.
var ErrFirmaElectronicaRequerida = errors.New("se requiere un firmador XML para facturación electrónica")

// FacturaConTipoDocumento expone el tipo fiscal predeterminado definido por
// el builder de una factura. No forma parte del XML del documento: se usa al
// construir la solicitud SOAP de recepción.
type FacturaConTipoDocumento interface {
	TipoFacturaDocumento() int
}

// FacturaConMetadatos expone la información fiscal completa del documento.
// Los servicios sólo consumen estos valores y no aplican reglas por sector.
type FacturaConMetadatos interface {
	FacturaConTipoDocumento
	CodigoDocumentoSector() int
}

// --- Interfaces opacas de Facturación ---

type AnulacionFactura struct {
	RequestWrapper[facturacion.AnulacionFactura]
}

type RecepcionFactura struct {
	RequestWrapper[facturacion.RecepcionFactura]
}

type ReversionAnulacionFactura struct {
	RequestWrapper[facturacion.ReversionAnulacionFactura]
}

type RecepcionPaqueteFactura struct {
	RequestWrapper[facturacion.RecepcionPaqueteFactura]
}

type ValidacionRecepcionPaqueteFactura struct {
	RequestWrapper[facturacion.ValidacionRecepcionPaqueteFactura]
}

type RecepcionMasivaFactura struct {
	RequestWrapper[facturacion.RecepcionMasivaFactura]
}

type ValidacionRecepcionMasivaFactura struct {
	RequestWrapper[facturacion.ValidacionRecepcionMasivaFactura]
}

type VerificacionEstadoFactura struct {
	RequestWrapper[facturacion.VerificacionEstadoFactura]
}

type VerificarComunicacionFacturacion struct {
	RequestWrapper[facturacion.VerificarComunicacion]
}

type RecepcionAnexosSuministroEnergia struct {
	RequestWrapper[facturacion.RecepcionAnexosSuministroEnergia]
}

type SuministroEnergiaAnexo struct {
	RequestWrapper[facturacion.SuministroEnergiaAnexo]
}

// --- Constructores y Builders ---

func NewVerificarComunicacionFacturacion() VerificarComunicacionFacturacion {
	return VerificarComunicacionFacturacion{
		RequestWrapper: NewRequestWrapper(&facturacion.VerificarComunicacion{}),
	}
}

// AnulacionFacturaBuilder
type AnulacionFacturaBuilder struct {
	request *facturacion.AnulacionFactura
}

func NewAnulacionFacturaBuilder() *AnulacionFacturaBuilder {
	return &AnulacionFacturaBuilder{
		request: &facturacion.AnulacionFactura{},
	}
}

func (b *AnulacionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoSucursal = codigoSucursal
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *AnulacionFacturaBuilder) WithCuis(cuis string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cuis = cuis
	return b
}

func (b *AnulacionFacturaBuilder) WithCufd(cufd string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cufd = cufd
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *AnulacionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoEmision(codigoEmision int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoEmision = codigoEmision
	return b
}

func (b *AnulacionFacturaBuilder) WithCuf(cuf string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Cuf = cuf
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoMotivo(codigoMotivo int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoMotivo = codigoMotivo
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoModalidad = codigoModalidad
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *AnulacionFacturaBuilder) WithCodigoSistema(codigoSistema string) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.CodigoSistema = codigoSistema
	return b
}

func (b *AnulacionFacturaBuilder) WithNit(nit int64) *AnulacionFacturaBuilder {
	b.request.SolicitudAnulacion.Nit = nit
	return b
}

func (b *AnulacionFacturaBuilder) Build() AnulacionFactura {
	return AnulacionFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// RecepcionFacturaBuilder
type RecepcionFacturaBuilder struct {
	request                   *facturacion.RecepcionFactura
	xmlPreparado              []byte
	hashArchivo               string
	xmlIndentado              bool
	firmaElectronicaRequerida bool
}

func NewRecepcionFacturaBuilder() *RecepcionFacturaBuilder {
	return &RecepcionFacturaBuilder{
		request: &facturacion.RecepcionFactura{},
	}
}

func (b *RecepcionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RecepcionFacturaBuilder) WithCuis(cuis string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *RecepcionFacturaBuilder) WithCufd(cufd string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *RecepcionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoEmision(codigoEmision int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *RecepcionFacturaBuilder) WithArchivo(archivo string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.Archivo = archivo
	return b
}

func (b *RecepcionFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *RecepcionFacturaBuilder) WithHashArchivo(hashArchivo string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hashArchivo
	return b
}

// WithXMLIndentado conserva una representación legible del documento. Debe
// activarse antes de WithFactura: la indentación forma parte del XML firmado.
func (b *RecepcionFacturaBuilder) WithXMLIndentado() *RecepcionFacturaBuilder {
	b.xmlIndentado = true
	return b
}

// WithFirmaElectronicaRequerida exige un XMLSigner cuando la modalidad sea
// electrónica. No afecta la modalidad computarizada ni el comportamiento
// existente mientras no se invoque explícitamente.
func (b *RecepcionFacturaBuilder) WithFirmaElectronicaRequerida() *RecepcionFacturaBuilder {
	b.firmaElectronicaRequerida = true
	return b
}

// WithDocumentoFiscal es la alternativa tipada a WithFactura. Garantiza en
// compilación que el documento expone sector y tipo fiscal.
func (b *RecepcionFacturaBuilder) WithDocumentoFiscal(documento FacturaConMetadatos, signer XMLSigner) error {
	return b.WithFactura(documento, signer)
}

// WithFactura serializa, firma (si es electrónica), comprime y calcula el hash de la factura automáticamente,
// mapeando los valores obtenidos en los campos Archivo y HashArchivo de la solicitud.
func (b *RecepcionFacturaBuilder) WithFactura(factura any, signer XMLSigner) error {
	solicitud := &b.request.SolicitudServicioRecepcionFactura.SolicitudRecepcion
	if documento, ok := factura.(FacturaConMetadatos); ok {
		if solicitud.CodigoDocumentoSector == 0 {
			solicitud.CodigoDocumentoSector = documento.CodigoDocumentoSector()
		}
		if solicitud.TipoFacturaDocumento == 0 {
			solicitud.TipoFacturaDocumento = documento.TipoFacturaDocumento()
		}
	} else if solicitud.TipoFacturaDocumento == 0 {
		if documento, ok := factura.(FacturaConTipoDocumento); ok {
			solicitud.TipoFacturaDocumento = documento.TipoFacturaDocumento()
		}
	}
	var xmlData []byte
	var err error
	if b.xmlIndentado {
		xmlData, err = xml.MarshalIndent(factura, "", "    ")
		xmlData = append([]byte(xml.Header), xmlData...)
	} else {
		xmlData, err = xml.Marshal(factura)
	}
	if err != nil {
		return err
	}

	var xmlToSend = xmlData
	var signRequired bool
	if b.request.SolicitudServicioRecepcionFactura.CodigoModalidad == ModalidadElectronica {
		signRequired = true
	}
	if signRequired && signer == nil && b.firmaElectronicaRequerida {
		return ErrFirmaElectronicaRequerida
	}

	if signRequired && signer != nil {
		var err error
		xmlToSend, err = signer.SignXML(xmlData)
		if err != nil {
			return err
		}
	}
	if b.xmlIndentado && !bytes.HasPrefix(xmlToSend, []byte("<?xml")) {
		xmlToSend = append([]byte(xml.Header), xmlToSend...)
	}

	hashString, encodedArchivo, err := utils.CompressAndHash(xmlToSend)
	if err != nil {
		return err
	}
	b.request.SolicitudServicioRecepcionFactura.Archivo = encodedArchivo
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hashString
	b.xmlPreparado = append(b.xmlPreparado[:0], xmlToSend...)
	b.hashArchivo = hashString
	return nil
}

// XMLPreparado devuelve una copia del XML firmado (o el XML computarizado) y
// el hash exacto incluidos en la recepción SIAT. Sólo está disponible después
// de invocar WithFactura.
func (b *RecepcionFacturaBuilder) XMLPreparado() ([]byte, string, bool) {
	if len(b.xmlPreparado) == 0 || b.hashArchivo == "" {
		return nil, "", false
	}
	return append([]byte(nil), b.xmlPreparado...), b.hashArchivo, true
}

func (b *RecepcionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RecepcionFacturaBuilder) WithCodigoSistema(codigoSistema string) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.CodigoSistema = codigoSistema
	return b
}

func (b *RecepcionFacturaBuilder) WithNit(nit int64) *RecepcionFacturaBuilder {
	b.request.SolicitudServicioRecepcionFactura.Nit = nit
	return b
}

func (b *RecepcionFacturaBuilder) Build() RecepcionFactura {
	return RecepcionFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// ReversionAnulacionFacturaBuilder
type ReversionAnulacionFacturaBuilder struct {
	request *facturacion.ReversionAnulacionFactura
}

func NewReversionAnulacionFacturaBuilder() *ReversionAnulacionFacturaBuilder {
	return &ReversionAnulacionFacturaBuilder{
		request: &facturacion.ReversionAnulacionFactura{},
	}
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSucursal = codigoSucursal
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCuis(cuis string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cuis = cuis
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCufd(cufd string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cufd = cufd
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoEmision(codigoEmision int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoEmision = codigoEmision
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCuf(cuf string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Cuf = cuf
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoModalidad = codigoModalidad
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithCodigoSistema(codigoSistema string) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.CodigoSistema = codigoSistema
	return b
}

func (b *ReversionAnulacionFacturaBuilder) WithNit(nit int64) *ReversionAnulacionFacturaBuilder {
	b.request.SolicitudReversionAnulacion.Nit = nit
	return b
}

func (b *ReversionAnulacionFacturaBuilder) Build() ReversionAnulacionFactura {
	return ReversionAnulacionFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// RecepcionPaqueteFacturaBuilder
type RecepcionPaqueteFacturaBuilder struct {
	request *facturacion.RecepcionPaqueteFactura
}

func NewRecepcionPaqueteFacturaBuilder() *RecepcionPaqueteFacturaBuilder {
	return &RecepcionPaqueteFacturaBuilder{
		request: &facturacion.RecepcionPaqueteFactura{},
	}
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCuis(cuis string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCufd(cufd string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoEmision(codigoEmision int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithArchivo(archivo string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.Archivo = archivo
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithHashArchivo(hashArchivo string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.HashArchivo = hashArchivo
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCafc(cafc *string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.Cafc = datatype.Nilable[string]{Value: cafc}
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCantidadFacturas(cantidadFacturas int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = cantidadFacturas
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoEvento(codigoEvento int64) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.CodigoEvento = codigoEvento
	return b
}

// WithFacturas serializa, firma (si es electrónica), empaqueta en un archivo TAR.GZ y calcula el hash de las facturas automáticamente,
// mapeando los valores obtenidos en los campos Archivo, HashArchivo y CantidadFacturas de la solicitud.
func (b *RecepcionPaqueteFacturaBuilder) WithFacturas(facturas []any, signer XMLSigner) error {
	return b.WithFacturasEnLote(facturas, signer, &FacturasEnLoteOptions{Workers: 1})
}

// WithFacturasEnLote serializa las facturas, las firma cuando corresponde y
// crea el archivo TAR.GZ codificado en Base64 requerido por el SIAT.
//
// Al finalizar actualiza Archivo, HashArchivo y CantidadFacturas de la
// solicitud. El hash se calcula sobre el archivo comprimido. Los XML se
// incorporan al TAR en el mismo orden recibido, incluso si Workers es mayor a
// uno.
//
// Con options nil, o con options.Workers igual a cero, se usa automáticamente
// los CPU lógicos disponibles menos uno. WithFacturas conserva su
// comportamiento serial para compatibilidad.
//
// En modalidad electrónica, el firmador se usa sólo en paralelo si declara
// ConcurrentXMLSigning() bool y retorna true. Si no lo declara, o si Workers
// es menor a dos, el procesamiento es serial. Esto evita asumir que un
// firmador de una aplicación consumidora es seguro para concurrencia.
func (b *RecepcionPaqueteFacturaBuilder) WithFacturasEnLote(facturas []any, signer XMLSigner, options *FacturasEnLoteOptions) error {
	return b.WithFacturasEnLoteContext(context.Background(), facturas, signer, options)
}

// WithFacturasEnLoteContext prepara una emisión masiva respetando la cancelación de ctx.
// Si ctx se cancela, no se asignan Archivo, HashArchivo ni CantidadFacturas.
// La firma que ya esté en ejecución no puede interrumpirse porque XMLSigner no
// recibe contexto, pero no se inicia el siguiente documento.
func (b *RecepcionPaqueteFacturaBuilder) WithFacturasEnLoteContext(ctx context.Context, facturas []any, signer XMLSigner, options *FacturasEnLoteOptions) error {
	signRequired := b.request.SolicitudServicioRecepcionPaquete.CodigoModalidad != 2
	hashString, encodedArchivo, metrics, err := prepararFacturas(ctx, facturas, signer, signRequired, options)
	if err != nil {
		return err
	}
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.Archivo = encodedArchivo
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.HashArchivo = hashString
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = len(facturas)
	if options != nil && options.OnComplete != nil {
		options.OnComplete(metrics)
	}
	return nil
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithCodigoSistema(codigoSistema string) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) WithNit(nit int64) *RecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}

func (b *RecepcionPaqueteFacturaBuilder) Build() RecepcionPaqueteFactura {
	return RecepcionPaqueteFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// ValidacionRecepcionPaqueteFacturaBuilder
type ValidacionRecepcionPaqueteFacturaBuilder struct {
	request *facturacion.ValidacionRecepcionPaqueteFactura
}

func NewValidacionRecepcionPaqueteFacturaBuilder() *ValidacionRecepcionPaqueteFacturaBuilder {
	return &ValidacionRecepcionPaqueteFacturaBuilder{
		request: &facturacion.ValidacionRecepcionPaqueteFactura{},
	}
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCuis(cuis string) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCufd(cufd string) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoEmision(codigoEmision int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoRecepcion(codigoRecepcion string) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) WithNit(nit int64) *ValidacionRecepcionPaqueteFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion.Nit = nit
	return b
}

func (b *ValidacionRecepcionPaqueteFacturaBuilder) Build() ValidacionRecepcionPaqueteFactura {
	return ValidacionRecepcionPaqueteFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// RecepcionMasivaFacturaBuilder
type RecepcionMasivaFacturaBuilder struct {
	request *facturacion.RecepcionMasivaFactura
}

func NewRecepcionMasivaFacturaBuilder() *RecepcionMasivaFacturaBuilder {
	return &RecepcionMasivaFacturaBuilder{
		request: &facturacion.RecepcionMasivaFactura{
			SolicitudServicioRecepcionMasiva: facturacion.SolicitudRecepcionMasiva{
				SolicitudRecepcionFactura: facturacion.SolicitudRecepcionFactura{
					SolicitudRecepcion: facturacion.SolicitudRecepcion{
						CodigoEmision: EmisionMasiva,
					},
				},
			},
		},
	}
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSucursal = codigoSucursal
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCuis(cuis string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cuis = cuis
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCufd(cufd string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Cufd = cufd
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoEmision(codigoEmision int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoEmision = codigoEmision
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithArchivo(archivo string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo = archivo
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithFechaEnvio(fechaEnvio time.Time) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.FechaEnvio = datatype.NewTimeSiat(fechaEnvio)
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithHashArchivo(hashArchivo string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.HashArchivo = hashArchivo
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoModalidad = codigoModalidad
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCodigoSistema(codigoSistema string) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.CodigoSistema = codigoSistema
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithNit(nit int64) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion.Nit = nit
	return b
}

func (b *RecepcionMasivaFacturaBuilder) WithCantidadFacturas(cantidadFacturas int) *RecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = cantidadFacturas
	return b
}

// WithFacturas serializa, firma (si es electrónica), empaqueta en un archivo TAR.GZ y calcula el hash de las facturas automáticamente,
// mapeando los valores obtenidos en los campos Archivo, HashArchivo y CantidadFacturas de la solicitud.
func (b *RecepcionMasivaFacturaBuilder) WithFacturas(facturas []any, signer XMLSigner) error {
	return b.WithFacturasEnLote(facturas, signer, &FacturasEnLoteOptions{Workers: 1})
}

// WithFacturasEnLote serializa las facturas, las firma cuando corresponde y
// crea el archivo TAR.GZ codificado en Base64 para recepción masiva del SIAT.
//
// Al finalizar actualiza Archivo, HashArchivo y CantidadFacturas de la
// solicitud. El hash corresponde al archivo TAR.GZ, no al XML individual. El
// orden de los XML dentro del archivo coincide siempre con el orden de
// facturas, aunque Workers sea mayor a uno.
//
// Con options nil, o con options.Workers igual a cero, se usa automáticamente
// los CPU lógicos disponibles menos uno. WithFacturas conserva su
// comportamiento serial para compatibilidad.
//
// En modalidad electrónica, sólo se firma en paralelo si el firmador declara
// ConcurrentXMLSigning() bool y retorna true. Esta restricción evita carreras
// con firmadores personalizados que no estén preparados para concurrencia.
func (b *RecepcionMasivaFacturaBuilder) WithFacturasEnLote(facturas []any, signer XMLSigner, options *FacturasEnLoteOptions) error {
	return b.WithFacturasEnLoteContext(context.Background(), facturas, signer, options)
}

// WithFacturasEnLoteContext prepara una emisión masiva respetando la cancelación
// de ctx. Si ctx se cancela, la solicitud se conserva sin modificaciones.
// XMLSigner no recibe contexto, por lo que una firma ya iniciada finaliza antes
// de observar la cancelación.
func (b *RecepcionMasivaFacturaBuilder) WithFacturasEnLoteContext(ctx context.Context, facturas []any, signer XMLSigner, options *FacturasEnLoteOptions) error {
	signRequired := b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.CodigoModalidad != 2
	hashString, encodedArchivo, metrics, err := prepararFacturas(ctx, facturas, signer, signRequired, options)
	if err != nil {
		return err
	}
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo = encodedArchivo
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.HashArchivo = hashString
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = len(facturas)
	if options != nil && options.OnComplete != nil {
		options.OnComplete(metrics)
	}
	return nil
}

type facturaPreparada struct {
	xml           []byte
	err           error
	serializacion time.Duration
	firma         time.Duration
}

func prepararFacturas(ctx context.Context, facturas []any, signer XMLSigner, signRequired bool, options *FacturasEnLoteOptions) (hashArchivo, archivo string, metrics FacturasEnLoteMetrics, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	metrics = FacturasEnLoteMetrics{CantidadFacturas: len(facturas)}
	started := time.Now()
	defer func() { metrics.DuracionTotal = time.Since(started) }()
	if err := ctx.Err(); err != nil {
		return "", "", metrics, err
	}
	if options != nil && options.MaxFacturas > 0 && len(facturas) > options.MaxFacturas {
		return "", "", metrics, fmt.Errorf("el lote contiene %d facturas y supera el máximo configurado de %d", len(facturas), options.MaxFacturas)
	}

	var destination io.Writer
	if options != nil {
		destination = options.DestinoArchivo
	}
	archive := utils.NewTarGzBase64WriterWithOutput(destination)
	workers := workersFacturas(len(facturas), signer, signRequired, options)
	metrics.Workers = workers
	if workers == 1 {
		for index, factura := range facturas {
			result := prepararFactura(ctx, factura, index+1, signer, signRequired)
			metrics.DuracionSerializacion += result.serializacion
			metrics.DuracionFirma += result.firma
			if result.err != nil {
				return "", "", metrics, result.err
			}
			if err := ctx.Err(); err != nil {
				return "", "", metrics, err
			}
			startedPackaging := time.Now()
			if err := agregarFacturaAlArchivo(archive, index+1, result.xml); err != nil {
				return "", "", metrics, err
			}
			metrics.DuracionEmpaquetado += time.Since(startedPackaging)
		}
		return cerrarArchivoFacturas(archive, metrics)
	}

	for first := 0; first < len(facturas); first += workers {
		if err := ctx.Err(); err != nil {
			return "", "", metrics, err
		}
		last := min(first+workers, len(facturas))
		prepared := make([]facturaPreparada, last-first)
		var wg sync.WaitGroup
		for index := first; index < last; index++ {
			localIndex := index - first
			wg.Add(1)
			go func(invoice any, invoiceNumber, position int) {
				defer wg.Done()
				prepared[position] = prepararFactura(ctx, invoice, invoiceNumber, signer, signRequired)
			}(facturas[index], index+1, localIndex)
		}
		wg.Wait()
		if err := ctx.Err(); err != nil {
			return "", "", metrics, err
		}

		for index, result := range prepared {
			metrics.DuracionSerializacion += result.serializacion
			metrics.DuracionFirma += result.firma
			if result.err != nil {
				return "", "", metrics, result.err
			}
			startedPackaging := time.Now()
			if err := agregarFacturaAlArchivo(archive, first+index+1, result.xml); err != nil {
				return "", "", metrics, err
			}
			metrics.DuracionEmpaquetado += time.Since(startedPackaging)
		}
	}
	return cerrarArchivoFacturas(archive, metrics)
}

func prepararFactura(ctx context.Context, factura any, invoiceNumber int, signer XMLSigner, signRequired bool) facturaPreparada {
	if err := ctx.Err(); err != nil {
		return facturaPreparada{err: err}
	}
	startedSerialization := time.Now()
	xmlData, err := marshalFacturaXML(factura)
	serializationDuration := time.Since(startedSerialization)
	if err != nil {
		return facturaPreparada{err: fmt.Errorf("error serializando factura %d: %w", invoiceNumber, err), serializacion: serializationDuration}
	}
	result := facturaPreparada{xml: xmlData, serializacion: serializationDuration}
	if signRequired && signer != nil {
		if err := ctx.Err(); err != nil {
			result.err = err
			return result
		}
		startedSigning := time.Now()
		xmlData, err = signer.SignXML(xmlData)
		result.firma = time.Since(startedSigning)
		if err != nil {
			result.err = fmt.Errorf("error firmando XML de factura %d: %w", invoiceNumber, err)
			return result
		}
		result.xml = xmlData
	}
	return result
}

func marshalFacturaXML(factura any) ([]byte, error) {
	if marshaler, ok := factura.(XMLBytesMarshaler); ok {
		return marshaler.MarshalXMLBytes()
	}
	return xml.Marshal(factura)
}

func agregarFacturaAlArchivo(archive *utils.TarGzBase64Writer, invoiceNumber int, xmlData []byte) error {
	if err := archive.Add(nombreFacturaArchivo(invoiceNumber), xmlData); err != nil {
		return fmt.Errorf("error escribiendo factura %d en el archivo: %w", invoiceNumber, err)
	}
	return nil
}

func nombreFacturaArchivo(invoiceNumber int) string {
	var buffer [32]byte
	name := append(buffer[:0], "factura_"...)
	name = strconv.AppendInt(name, int64(invoiceNumber), 10)
	name = append(name, ".xml"...)
	return string(name)
}

func cerrarArchivoFacturas(archive *utils.TarGzBase64Writer, metrics FacturasEnLoteMetrics) (string, string, FacturasEnLoteMetrics, error) {
	startedPackaging := time.Now()
	if err := archive.Close(); err != nil {
		return "", "", metrics, fmt.Errorf("error cerrando archivo comprimido: %w", err)
	}
	metrics.DuracionEmpaquetado += time.Since(startedPackaging)
	metrics.TamanoArchivo = archive.Size()
	return archive.Hash(), archive.Encoded(), metrics, nil
}

func workersFacturas(total int, signer XMLSigner, signRequired bool, options *FacturasEnLoteOptions) int {
	if total < 1 {
		return 1
	}
	workers := 0
	if options != nil {
		workers = options.Workers
	}
	if workers == 0 {
		workers = max(1, runtime.GOMAXPROCS(0)-1)
	}
	if workers < 2 {
		return 1
	}
	if !signRequired || signer == nil {
		return min(workers, total)
	}
	if concurrent, ok := signer.(concurrentXMLSigner); !ok || !concurrent.ConcurrentXMLSigning() {
		return 1
	}
	return min(workers, total)
}

func (b *RecepcionMasivaFacturaBuilder) Build() RecepcionMasivaFactura {
	return RecepcionMasivaFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// ValidacionRecepcionMasivaFacturaBuilder
type ValidacionRecepcionMasivaFacturaBuilder struct {
	request *facturacion.ValidacionRecepcionMasivaFactura
}

func NewValidacionRecepcionMasivaFacturaBuilder() *ValidacionRecepcionMasivaFacturaBuilder {
	return &ValidacionRecepcionMasivaFacturaBuilder{
		request: &facturacion.ValidacionRecepcionMasivaFactura{
			SolicitudServicioValidacionRecepcionMasivaFactura: facturacion.SolicitudValidacionRecepcionMasiva{
				SolicitudRecepcion: facturacion.SolicitudRecepcion{
					CodigoEmision: EmisionMasiva,
				},
			},
		},
	}
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoAmbiente(codigoAmbiente int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoAmbiente = codigoAmbiente
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCuis(cuis string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cuis = cuis
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCufd(cufd string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Cufd = cufd
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoEmision(codigoEmision int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoEmision = codigoEmision
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoRecepcion(codigoRecepcion string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion = codigoRecepcion
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithCodigoSistema(codigoSistema string) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.CodigoSistema = codigoSistema
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) WithNit(nit int64) *ValidacionRecepcionMasivaFacturaBuilder {
	b.request.SolicitudServicioValidacionRecepcionMasivaFactura.Nit = nit
	return b
}

func (b *ValidacionRecepcionMasivaFacturaBuilder) Build() ValidacionRecepcionMasivaFactura {
	return ValidacionRecepcionMasivaFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// VerificacionEstadoFacturaBuilder
type VerificacionEstadoFacturaBuilder struct {
	request *facturacion.VerificacionEstadoFactura
}

func NewVerificacionEstadoFacturaBuilder() *VerificacionEstadoFacturaBuilder {
	return &VerificacionEstadoFacturaBuilder{
		request: &facturacion.VerificacionEstadoFactura{},
	}
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoSucursal(codigoSucursal int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoSucursal = codigoSucursal
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCuis(cuis string) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuis = cuis
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCufd(cufd string) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cufd = cufd
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoEmision(codigoEmision int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoEmision = codigoEmision
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCuf(cuf string) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Cuf = cuf
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithCodigoModalidad(codigoModalidad int) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.CodigoModalidad = codigoModalidad
	return b
}

func (b *VerificacionEstadoFacturaBuilder) WithNit(nit int64) *VerificacionEstadoFacturaBuilder {
	b.request.SolicitudServicioVerificacionEstadoFactura.Nit = nit
	return b
}

func (b *VerificacionEstadoFacturaBuilder) Build() VerificacionEstadoFactura {
	return VerificacionEstadoFactura{RequestWrapper: NewRequestWrapper(b.request)}
}

// RecepcionAnexosSuministroEnergiaBuilder
type RecepcionAnexosSuministroEnergiaBuilder struct {
	request *facturacion.RecepcionAnexosSuministroEnergia
}

func NewRecepcionAnexosSuministroEnergiaBuilder() *RecepcionAnexosSuministroEnergiaBuilder {
	return &RecepcionAnexosSuministroEnergiaBuilder{
		request: &facturacion.RecepcionAnexosSuministroEnergia{},
	}
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithCodigoSucursal(codigoSucursal int) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoSucursal = codigoSucursal
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithCuis(cuis string) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Cuis = cuis
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithCufd(cufd string) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Cufd = cufd
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithCodigoEmision(codigoEmision int) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.CodigoEmision = codigoEmision
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithGiftCard(giftCard int64) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.GiftCard = giftCard
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) AddAnexos(anexos ...SuministroEnergiaAnexo) *RecepcionAnexosSuministroEnergiaBuilder {
	for _, a := range anexos {
		if internal := UnwrapInternalRequest[facturacion.SuministroEnergiaAnexo](a); internal != nil {
			b.request.SolicitudRecepcionSuministroAnexos.AnexosList = append(b.request.SolicitudRecepcionSuministroAnexos.AnexosList, *internal)
		}
	}
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) WithNit(nit int64) *RecepcionAnexosSuministroEnergiaBuilder {
	b.request.SolicitudRecepcionSuministroAnexos.Nit = nit
	return b
}

func (b *RecepcionAnexosSuministroEnergiaBuilder) Build() RecepcionAnexosSuministroEnergia {
	return RecepcionAnexosSuministroEnergia{RequestWrapper: NewRequestWrapper(b.request)}
}

// SuministroEnergiaAnexoBuilder
type SuministroEnergiaAnexoBuilder struct {
	request *facturacion.SuministroEnergiaAnexo
}

func NewSuministroEnergiaAnexoBuilder() *SuministroEnergiaAnexoBuilder {
	return &SuministroEnergiaAnexoBuilder{
		request: &facturacion.SuministroEnergiaAnexo{},
	}
}

func (b *SuministroEnergiaAnexoBuilder) WithCufFactSuministro(cuf string) *SuministroEnergiaAnexoBuilder {
	b.request.CufFactSuministro = cuf
	return b
}

func (b *SuministroEnergiaAnexoBuilder) WithFechaRecarga(fecha time.Time) *SuministroEnergiaAnexoBuilder {
	b.request.FechaRecarga = datatype.NewTimeSiat(fecha)
	return b
}

func (b *SuministroEnergiaAnexoBuilder) WithMontoRecarga(monto float64) *SuministroEnergiaAnexoBuilder {
	b.request.MontoRecarga = monto
	return b
}

func (b *SuministroEnergiaAnexoBuilder) Build() SuministroEnergiaAnexo {
	return SuministroEnergiaAnexo{RequestWrapper: NewRequestWrapper(b.request)}
}
