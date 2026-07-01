package models

import (
	"archive/tar"
	"bytes"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v2/pkg/utils"
)

// XMLSigner define la interfaz para realizar la firma de documentos XML
type XMLSigner interface {
	SignXML(xmlBytes []byte) ([]byte, error)
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
	request *facturacion.RecepcionFactura
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

// WithFactura serializa, firma (si es electrónica), comprime y calcula el hash de la factura automáticamente,
// mapeando los valores obtenidos en los campos Archivo y HashArchivo de la solicitud.
func (b *RecepcionFacturaBuilder) WithFactura(factura any, signer XMLSigner) error {
	xmlData, err := xml.Marshal(factura)
	if err != nil {
		return err
	}

	var xmlToSend = xmlData
	var signRequired bool
	if b.request.SolicitudServicioRecepcionFactura.CodigoModalidad == ModalidadElectronica {
		signRequired = true
	}

	if signRequired && signer != nil {
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
	b.request.SolicitudServicioRecepcionFactura.Archivo = encodedArchivo
	b.request.SolicitudServicioRecepcionFactura.HashArchivo = hashString
	return nil
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
	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)

	var signRequired = true
	if b.request.SolicitudServicioRecepcionPaquete.CodigoModalidad == 2 { // Modalidad Computarizada
		signRequired = false
	}

	for i, f := range facturas {
		xmlData, err := xml.Marshal(f)
		if err != nil {
			return fmt.Errorf("error serializando factura %d: %w", i+1, err)
		}
		var xmlToSend = xmlData
		if signRequired && signer != nil {
			var err error
			xmlToSend, err = signer.SignXML(xmlData)
			if err != nil {
				return fmt.Errorf("error firmando XML de factura %d: %w", i+1, err)
			}
		}
		hdr := &tar.Header{
			Name: fmt.Sprintf("factura_%d.xml", i+1),
			Mode: 0600,
			Size: int64(len(xmlToSend)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return fmt.Errorf("error escribiendo cabecera tar de factura %d: %w", i+1, err)
		}
		if _, err := tw.Write(xmlToSend); err != nil {
			return fmt.Errorf("error escribiendo datos tar de factura %d: %w", i+1, err)
		}
	}
	if err := tw.Close(); err != nil {
		return fmt.Errorf("error cerrando tar: %w", err)
	}
	hashString, encodedArchivo, err := utils.CompressAndHash(tarBuf.Bytes())
	if err != nil {
		return fmt.Errorf("error comprimiendo archivo: %w", err)
	}
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.Archivo = encodedArchivo
	b.request.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.HashArchivo = hashString
	b.request.SolicitudServicioRecepcionPaquete.CantidadFacturas = len(facturas)
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
		request: &facturacion.RecepcionMasivaFactura{},
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
	var tarBuf bytes.Buffer
	tw := tar.NewWriter(&tarBuf)

	var signRequired = true
	if b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.CodigoModalidad == 2 { // Modalidad Computarizada
		signRequired = false
	}

	for i, f := range facturas {
		xmlData, err := xml.Marshal(f)
		if err != nil {
			return fmt.Errorf("error serializando factura %d: %w", i+1, err)
		}
		var xmlToSend = xmlData
		if signRequired && signer != nil {
			var err error
			xmlToSend, err = signer.SignXML(xmlData)
			if err != nil {
				return fmt.Errorf("error firmando XML de factura %d: %w", i+1, err)
			}
		}
		hdr := &tar.Header{
			Name: fmt.Sprintf("factura_%d.xml", i+1),
			Mode: 0600,
			Size: int64(len(xmlToSend)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return fmt.Errorf("error escribiendo cabecera tar de factura %d: %w", i+1, err)
		}
		if _, err := tw.Write(xmlToSend); err != nil {
			return fmt.Errorf("error escribiendo datos tar de factura %d: %w", i+1, err)
		}
	}
	if err := tw.Close(); err != nil {
		return fmt.Errorf("error cerrando tar: %w", err)
	}
	hashString, encodedArchivo, err := utils.CompressAndHash(tarBuf.Bytes())
	if err != nil {
		return fmt.Errorf("error comprimiendo archivo: %w", err)
	}
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.Archivo = encodedArchivo
	b.request.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.HashArchivo = hashString
	b.request.SolicitudServicioRecepcionMasiva.CantidadFacturas = len(facturas)
	return nil
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
		request: &facturacion.ValidacionRecepcionMasivaFactura{},
	}
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
