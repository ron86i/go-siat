package facturacion

import (
	"context"
	"fmt"
	"time"

	"github.com/ron86i/go-siat/v3/internal/pipeline"
	"github.com/ron86i/go-siat/v3/internal/siatpolicy"
)

// Modalidades de facturación oficiales.
const (
	ModalidadElectronica   = 1
	ModalidadComputarizada = 2
)

// XMLSigner firma el XML de una factura electrónica.
type XMLSigner = pipeline.XMLSigner

func validateLocation(sucursalSolicitud int, puntoSolicitud *int, sucursalFactura int, puntoFactura *int) error {
	if sucursalSolicitud != sucursalFactura {
		return fmt.Errorf("la sucursal de factura y solicitud no coincide")
	}
	if pointValue(puntoSolicitud) != pointValue(puntoFactura) {
		return fmt.Errorf("el punto de venta de factura y solicitud no coincide")
	}
	return nil
}

func pointValue(point *int) int {
	if point == nil {
		return 0
	}
	return *point
}

// PrepararFactura devuelve una copia de la solicitud con el documento fiscal preparado.
// Acepta cualquier documento fiscal que implemente Documento (Compra-Venta, Alquileres, etc.).
// La modalidad de emisión debe estar configurada en r.CodigoModalidad (ModalidadElectronica o ModalidadComputarizada).
func (r SolicitudRecepcionFactura) PrepararFactura(doc Documento, signer XMLSigner) (SolicitudRecepcionFactura, error) {
	validated, err := siatpolicy.ValidateDocument(doc)
	if err != nil {
		return r, err
	}
	doc = validated
	if r.CodigoModalidad != ModalidadElectronica && r.CodigoModalidad != ModalidadComputarizada {
		return r, fmt.Errorf("debe especificar CodigoModalidad (ModalidadElectronica o ModalidadComputarizada) en la solicitud antes de preparar la factura")
	}
	if err := siatpolicy.ValidateDocumentMode(doc, r.CodigoModalidad); err != nil {
		return r, err
	}
	if r.CodigoModalidad == ModalidadElectronica && signer == nil {
		return r, fmt.Errorf("la modalidad electrónica requiere un firmador digital (XMLSigner)")
	}
	if r.CodigoModalidad == ModalidadComputarizada && signer != nil {
		return r, fmt.Errorf("la modalidad computarizada no admite firma digital")
	}

	sector := doc.CodigoDocumentoSector()
	if sector <= 0 {
		return r, fmt.Errorf("código de documento sector inválido: %d", sector)
	}
	if r.CodigoDocumentoSector == 0 {
		r.CodigoDocumentoSector = sector
	} else if r.CodigoDocumentoSector != sector {
		return r, fmt.Errorf("el sector de la solicitud (%d) no coincide con el sector del documento (%d)", r.CodigoDocumentoSector, sector)
	}
	if err := siatpolicy.ValidateDocumentDestination(sector, r.CodigoModalidad); err != nil {
		return r, err
	}
	if r.CodigoEmision != 0 && r.CodigoEmision != 1 {
		return r, fmt.Errorf("codigoEmision %d no corresponde a recepción individual (esperado 1)", r.CodigoEmision)
	}

	if r.CodigoSucursal == 0 && doc.CodigoSucursal() != 0 {
		r.CodigoSucursal = doc.CodigoSucursal()
	}
	if r.CodigoPuntoVenta == nil && doc.CodigoPuntoVenta() != nil {
		r.CodigoPuntoVenta = doc.CodigoPuntoVenta()
	}
	if err := validateLocation(r.CodigoSucursal, r.CodigoPuntoVenta, doc.CodigoSucursal(), doc.CodigoPuntoVenta()); err != nil {
		return r, err
	}
	if r.Cufd == "" && doc.Cufd() != "" {
		r.Cufd = doc.Cufd()
	} else if r.Cufd != "" && doc.Cufd() != "" && r.Cufd != doc.Cufd() {
		return r, fmt.Errorf("el CUFD del documento y la solicitud no coincide")
	}
	archivo, hash, err := pipeline.Factura(doc, r.CodigoModalidad, signer)
	if err != nil {
		return r, err
	}
	r.Archivo = archivo
	r.HashArchivo = hash
	if r.TipoFacturaDocumento == 0 {
		r.TipoFacturaDocumento = 1
	}
	r.CodigoEmision = 1
	if r.FechaEnvio.IsZero() {
		r.FechaEnvio = time.Now()
	}
	return r, nil
}

// WithFactura prepara un documento fiscal seleccionado para recepción individual.
func (r SolicitudRecepcionFactura) WithFactura(f Documento, signer XMLSigner) (SolicitudRecepcionFactura, error) {
	return r.PrepararFactura(f, signer)
}

// PrepararFactura prepara la solicitud de recepción para cualquier documento fiscal detectando conflictos de modalidad.
func (c *Client) PrepararFactura(req SolicitudRecepcionFactura, doc Documento, signer XMLSigner) (SolicitudRecepcionFactura, error) {
	if c != nil && c.modalidad != 0 {
		if req.CodigoModalidad != 0 && req.CodigoModalidad != c.modalidad {
			return req, fmt.Errorf("conflicto de modalidad: el cliente está configurado para modalidad %d pero la solicitud indica %d", c.modalidad, req.CodigoModalidad)
		}
		req.CodigoModalidad = c.modalidad
	}
	validated, err := siatpolicy.ValidateDocument(doc)
	if err != nil {
		return req, err
	}
	if err := c.validateDestination(validated.CodigoDocumentoSector(), req.CodigoModalidad); err != nil {
		return req, err
	}
	return req.PrepararFactura(validated, signer)
}

// WithDocumento prepara un documento fiscal seleccionado para su recepción individual.
func (r SolicitudRecepcionFactura) WithDocumento(doc Documento, sector, modalidad int, signer XMLSigner) (SolicitudRecepcionFactura, error) {
	if sector <= 0 {
		return r, fmt.Errorf("el código de documento sector debe ser mayor a cero")
	}
	if modalidad != 1 && modalidad != 2 {
		return r, fmt.Errorf("modalidad inválida: %d", modalidad)
	}
	validated, err := siatpolicy.ValidateDocument(doc)
	if err != nil {
		return r, err
	}
	if validated.CodigoDocumentoSector() != sector {
		return r, fmt.Errorf("sector declarado (%d) no coincide con el documento (%d)", sector, validated.CodigoDocumentoSector())
	}
	if r.CodigoModalidad == 0 {
		r.CodigoModalidad = modalidad
	} else if r.CodigoModalidad != modalidad {
		return r, fmt.Errorf("modalidad de documento y solicitud no coinciden")
	}
	if r.CodigoDocumentoSector != 0 && r.CodigoDocumentoSector != sector {
		return r, fmt.Errorf("sector de documento y solicitud no coinciden")
	}
	r.CodigoDocumentoSector = sector
	return r.PrepararFactura(validated, signer)
}

// ConDocumento es un alias para WithDocumento.
func (r SolicitudRecepcionFactura) ConDocumento(doc Documento, sector, modalidad int, signer XMLSigner) (SolicitudRecepcionFactura, error) {
	return r.WithDocumento(doc, sector, modalidad, signer)
}

// PrepararPaquete devuelve una copia de la solicitud con el paquete de documentos preparado.
// Todos los documentos deben pertenecer al mismo sector y protocolo SIAT v3.
func (r SolicitudRecepcionPaquete) PrepararPaquete(ctx context.Context, facturas []Documento, signer XMLSigner) (SolicitudRecepcionPaquete, error) {
	prepared, count, err := prepararDocumentos(ctx, r.SolicitudRecepcionFactura, facturas, signer, 2, siatpolicy.MaxFacturasPaquete)
	if err != nil {
		return r, err
	}
	r.SolicitudRecepcionFactura = prepared
	r.CantidadFacturas = count
	return r, nil
}

// WithFacturas devuelve una copia de la solicitud con un paquete de documentos v3 seleccionados preparado.
func (r SolicitudRecepcionPaquete) WithFacturas(ctx context.Context, facturas []Documento, signer XMLSigner) (SolicitudRecepcionPaquete, error) {
	return r.PrepararPaquete(ctx, facturas, signer)
}

// ConPaqueteDocumentos prepara un paquete de documentos fiscales seleccionados para cualquier sector.
func (r SolicitudRecepcionPaquete) ConPaqueteDocumentos(ctx context.Context, documentos []Documento, sector, modalidad int, signer XMLSigner) (SolicitudRecepcionPaquete, error) {
	if sector <= 0 {
		return r, fmt.Errorf("el código de documento sector debe ser mayor a cero")
	}
	if modalidad != 1 && modalidad != 2 {
		return r, fmt.Errorf("modalidad inválida: %d", modalidad)
	}
	docs := make([]Documento, len(documentos))
	for i, doc := range documentos {
		validated, err := siatpolicy.ValidateDocument(doc)
		if err != nil {
			return r, fmt.Errorf("documento %d: %w", i+1, err)
		}
		if validated.CodigoDocumentoSector() != sector {
			return r, fmt.Errorf("documento %d: sector declarado (%d) no coincide con el documento (%d)", i+1, sector, validated.CodigoDocumentoSector())
		}
		docs[i] = validated
	}
	if r.CodigoModalidad != 0 && r.CodigoModalidad != modalidad {
		return r, fmt.Errorf("modalidad de documentos y solicitud no coinciden")
	}
	if r.CodigoDocumentoSector != 0 && r.CodigoDocumentoSector != sector {
		return r, fmt.Errorf("sector de documentos y solicitud no coinciden")
	}
	r.CodigoModalidad = modalidad
	r.CodigoDocumentoSector = sector
	return r.PrepararPaquete(ctx, docs, signer)
}

func prepararDocumentos(ctx context.Context, r SolicitudRecepcionFactura, facturas []Documento, signer XMLSigner, emision, limite int) (SolicitudRecepcionFactura, int, error) {
	if len(facturas) == 0 || len(facturas) > limite {
		return r, 0, fmt.Errorf("cantidad de documentos fuera del límite de la operación: debe ser entre 1 y %d", limite)
	}
	if r.CodigoModalidad != ModalidadElectronica && r.CodigoModalidad != ModalidadComputarizada {
		return r, 0, fmt.Errorf("debe especificar CodigoModalidad antes de preparar los documentos")
	}
	if r.CodigoModalidad == ModalidadElectronica && signer == nil {
		return r, 0, fmt.Errorf("la modalidad electrónica requiere un firmador digital (XMLSigner)")
	}
	if r.CodigoModalidad == ModalidadComputarizada && signer != nil {
		return r, 0, fmt.Errorf("la modalidad computarizada no admite firma digital")
	}
	if r.CodigoEmision != 0 && r.CodigoEmision != emision {
		return r, 0, fmt.Errorf("codigoEmision %d no corresponde a esta operación (esperado %d)", r.CodigoEmision, emision)
	}

	validated := make([]Documento, len(facturas))
	for i, factura := range facturas {
		doc, err := siatpolicy.ValidateDocument(factura)
		if err != nil {
			return r, 0, fmt.Errorf("documento %d: %w", i+1, err)
		}
		if err := siatpolicy.ValidateDocumentMode(doc, r.CodigoModalidad); err != nil {
			return r, 0, fmt.Errorf("documento %d: %w", i+1, err)
		}
		if i == 0 {
			if r.CodigoDocumentoSector != 0 && r.CodigoDocumentoSector != doc.CodigoDocumentoSector() {
				return r, 0, fmt.Errorf("el sector de la solicitud (%d) no coincide con el documento (%d)", r.CodigoDocumentoSector, doc.CodigoDocumentoSector())
			}
			if err := siatpolicy.ValidateDocumentDestination(doc.CodigoDocumentoSector(), r.CodigoModalidad); err != nil {
				return r, 0, err
			}
			r.CodigoDocumentoSector = doc.CodigoDocumentoSector()
			if r.CodigoSucursal == 0 && doc.CodigoSucursal() != 0 {
				r.CodigoSucursal = doc.CodigoSucursal()
			}
			if r.CodigoPuntoVenta == nil && doc.CodigoPuntoVenta() != nil {
				point := *doc.CodigoPuntoVenta()
				r.CodigoPuntoVenta = &point
			}
			if r.Cufd == "" {
				r.Cufd = doc.Cufd()
			}
		} else {
			if doc.CodigoDocumentoSector() != r.CodigoDocumentoSector {
				return r, 0, fmt.Errorf("documento %d: sector incompatible (%d != %d)", i+1, doc.CodigoDocumentoSector(), r.CodigoDocumentoSector)
			}
			if doc.Cufd() != validated[0].Cufd() {
				return r, 0, fmt.Errorf("documento %d: CUFD distinto al primer documento", i+1)
			}
		}
		if err := validateLocation(r.CodigoSucursal, r.CodigoPuntoVenta, doc.CodigoSucursal(), doc.CodigoPuntoVenta()); err != nil {
			return r, 0, fmt.Errorf("documento %d: %w", i+1, err)
		}
		validated[i] = doc
	}
	archivo, hash, err := pipeline.Paquete(ctx, validated, r.CodigoModalidad, signer)
	if err != nil {
		return r, 0, err
	}
	r.CodigoEmision = emision
	r.Archivo = archivo
	r.HashArchivo = hash
	if r.TipoFacturaDocumento == 0 {
		r.TipoFacturaDocumento = 1
	}
	if r.FechaEnvio.IsZero() {
		r.FechaEnvio = time.Now()
	}
	return r, len(validated), nil
}
