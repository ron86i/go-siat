package models

import "github.com/ron86i/go-siat/v2/internal/core/domain/documents"

// ConfirmacionCompra representa un XML individual para un archivo de
// confirmaciones de compras. Puede serializarse directamente o incluirse en un
// archivo TAR.GZ antes de enviarlo al servicio de recepción de compras.
type ConfirmacionCompra struct {
	RequestWrapper[documents.ConfirmacionCompra]
}

// NewConfirmacionCompraBuilder crea un builder para el XML confirmacionCompra.
func NewConfirmacionCompraBuilder() *confirmacionCompraBuilder {
	return &confirmacionCompraBuilder{
		confirmacion: &documents.ConfirmacionCompra{},
	}
}

type confirmacionCompraBuilder struct {
	confirmacion *documents.ConfirmacionCompra
}

// WithNro establece el correlativo de la confirmación (1 a 9999 según XSD).
func (b *confirmacionCompraBuilder) WithNro(nro int) *confirmacionCompraBuilder {
	b.confirmacion.Nro = nro
	return b
}

// WithNitEmisor establece el NIT del emisor de la factura de compra.
func (b *confirmacionCompraBuilder) WithNitEmisor(nitEmisor int64) *confirmacionCompraBuilder {
	b.confirmacion.NitEmisor = nitEmisor
	return b
}

// WithCodigoAutorizacion establece el código de autorización de la factura.
func (b *confirmacionCompraBuilder) WithCodigoAutorizacion(codigo string) *confirmacionCompraBuilder {
	b.confirmacion.CodigoAutorizacion = codigo
	return b
}

// WithNumeroFactura establece el número de factura. Se recibe como string para
// preservar el máximo de 20 dígitos permitido por el XSD.
func (b *confirmacionCompraBuilder) WithNumeroFactura(numero string) *confirmacionCompraBuilder {
	b.confirmacion.NumeroFactura = numero
	return b
}

// WithTipoCompra establece el tipo de compra (0 a 9 según XSD).
func (b *confirmacionCompraBuilder) WithTipoCompra(tipo int) *confirmacionCompraBuilder {
	b.confirmacion.TipoCompra = tipo
	return b
}

// Build construye el XML individual de confirmación de compra.
func (b *confirmacionCompraBuilder) Build() ConfirmacionCompra {
	return ConfirmacionCompra{RequestWrapper: NewRequestWrapper(b.confirmacion)}
}
