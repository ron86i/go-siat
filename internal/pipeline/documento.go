// Package pipeline serializa, firma y empaqueta documentos para su recepción en el SIAT.
package pipeline

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"reflect"

	"github.com/ron86i/go-siat/v3/pkg/utils"
)

// Modalidades de facturación oficiales.
const (
	ModalidadElectronica   = 1
	ModalidadComputarizada = 2
)

// XMLSigner firma el XML electrónico antes de comprimirlo.
type XMLSigner interface{ SignXML([]byte) ([]byte, error) }

// ModalidadMarshaler permite a un documento serializarse con el elemento raíz
// correspondiente a la modalidad de emisión indicada (electrónica o computarizada).
type ModalidadMarshaler interface {
	MarshalXMLModalidad(modalidad int) ([]byte, error)
}

// Factura serializa, firma cuando corresponde y comprime el documento.
func Factura(documento any, modalidad int, signer XMLSigner) ([]byte, string, error) {
	contenido, err := xmlInvoice(documento, modalidad, signer)
	if err != nil {
		return nil, "", err
	}
	archivo, err := utils.Gzip(contenido)
	if err != nil {
		return nil, "", fmt.Errorf("comprimir factura: %w", err)
	}
	return archivo, fmt.Sprintf("%x", sha256.Sum256(archivo)), nil
}

func xmlInvoice(documento any, modalidad int, signer XMLSigner) ([]byte, error) {
	if modalidad != ModalidadElectronica && modalidad != ModalidadComputarizada {
		return nil, fmt.Errorf("modalidad inválida: %d", modalidad)
	}
	if modalidad == ModalidadElectronica && signer == nil {
		return nil, fmt.Errorf("la modalidad electrónica requiere un firmador digital (XMLSigner)")
	}
	if modalidad == ModalidadComputarizada && signer != nil {
		return nil, fmt.Errorf("la modalidad computarizada no admite firma digital")
	}
	var contenido []byte
	var err error
	if mm, ok := documento.(ModalidadMarshaler); ok {
		contenido, err = mm.MarshalXMLModalidad(modalidad)
	} else {
		contenido, err = xml.Marshal(documento)
	}
	if err != nil {
		return nil, fmt.Errorf("serializar factura: %w", err)
	}
	if modalidad == ModalidadElectronica {
		contenido, err = signer.SignXML(contenido)
		if err != nil {
			return nil, fmt.Errorf("firmar factura: %w", err)
		}
		if len(contenido) == 0 {
			return nil, fmt.Errorf("el firmador devolvió XML vacío")
		}
	}
	return contenido, nil
}

// Paquete prepara los XML en orden y devuelve el TAR.GZ y su hash.
// Acepta []any, []xml.Marshaler, []ModalidadMarshaler o cualquier slice de documentos.
func Paquete(ctx context.Context, documentos any, modalidad int, signer XMLSigner) ([]byte, string, error) {
	if documentos == nil {
		return nil, "", fmt.Errorf("el paquete no contiene facturas")
	}
	var docs []any
	switch d := documentos.(type) {
	case []any:
		docs = d
	case []xml.Marshaler:
		docs = make([]any, len(d))
		for i, v := range d {
			docs[i] = v
		}
	default:
		rv := reflect.ValueOf(documentos)
		if rv.Kind() == reflect.Slice {
			docs = make([]any, rv.Len())
			for i := 0; i < rv.Len(); i++ {
				docs[i] = rv.Index(i).Interface()
			}
		} else {
			return nil, "", fmt.Errorf("el paquete requiere un slice de documentos")
		}
	}
	if len(docs) == 0 {
		return nil, "", fmt.Errorf("el paquete no contiene facturas")
	}
	var comprimido bytes.Buffer
	archivo := utils.NewTarGzBase64WriterWithOutput(&comprimido)
	for i, documento := range docs {
		if err := ctx.Err(); err != nil {
			return nil, "", err
		}
		contenido, err := xmlInvoice(documento, modalidad, signer)
		if err != nil {
			return nil, "", fmt.Errorf("factura %d: %w", i+1, err)
		}
		if err := archivo.Add(fmt.Sprintf("factura_%d.xml", i+1), contenido); err != nil {
			return nil, "", fmt.Errorf("agregar factura %d: %w", i+1, err)
		}
	}
	if err := ctx.Err(); err != nil {
		return nil, "", err
	}
	if err := archivo.Close(); err != nil {
		return nil, "", fmt.Errorf("cerrar paquete: %w", err)
	}
	return comprimido.Bytes(), archivo.Hash(), nil
}
