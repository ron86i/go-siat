package models

import (
	"encoding/xml"
)

// RequestWrapper es una envoltura genérica que encapsula el objeto de solicitud concreto
// hacia el SIAT, ocultando su implementación interna para garantizar el principio de encapsulamiento.
//
// Aunque el tipo es público (para ser utilizado como campo embebido en los tipos de solicitud
// del SDK), su campo interno es privado, lo que impide el acceso o modificación directa
// desde fuera del paquete.
//
// Los constructores (builders) del SDK son la única vía válida para crear y configurar
// estas solicitudes correctamente.
type RequestWrapper[T any] struct {
	request *T
}

// NewRequestWrapper crea una nueva instancia de RequestWrapper con el objeto interno configurado.
// Es de uso exclusivo dentro del SDK para inicializar los tipos de respuesta de los builders.
func NewRequestWrapper[T any](request *T) RequestWrapper[T] {
	return RequestWrapper[T]{request: request}
}

// internal retorna el puntero al objeto de solicitud interno.
// Es un método no exportado, accesible únicamente desde dentro del paquete.
func (r RequestWrapper[T]) internal() *T {
	return r.request
}

// MarshalXML implementa la interfaz [xml.Marshaler] para delegar la serialización XML
// al objeto interno, evitando que la etiqueta raíz generada sea "RequestWrapper"
// en lugar del elemento raíz correcto del mensaje SOAP.
func (r RequestWrapper[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(r.request)
}

// UnwrapInternalRequest extrae el puntero al objeto de solicitud interno de cualquier
// tipo que implemente la interfaz internal(). Se utiliza en los builders de tipo compuesto
// (p.ej. AddAnexos) donde se necesita acceder a la estructura interna de un sub-modelo.
//
// Devuelve nil si el tipo provisto no implementa la interfaz interna esperada.
//
// Nota: su uso directo fuera del SDK no está previsto y puede indicar un diseño incorrecto.
func UnwrapInternalRequest[T any](req any) *T {
	if r, ok := req.(interface{ internal() *T }); ok {
		return r.internal()
	}
	return nil
}

// ModalidadComputarizada es la constante que identifica la modalidad de facturación
// Electronica (valor = 1) y Computarizada (valor = 2) según la nomenclatura del SIAT.
// Se usa internamente para determinar si el XML de la factura debe ser firmado digitalmente.
const (
	ModalidadElectronica   = 1
	ModalidadComputarizada = 2
)
