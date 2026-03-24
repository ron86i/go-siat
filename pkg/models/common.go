package models

import (
	"encoding/xml"
)

// RequestWrapper es una envoltura genérica utilizada para ocultar la implementación concreta
// de una solicitud y satisfacer las interfaces opacas del SDK.
//
// Aunque el tipo es público, su contenido es privado para evitar el acceso directo desde fuera del SDK.
type RequestWrapper[T any] struct {
	request *T
}

// internal retorna el objeto interno de la solicitud (privado para el SDK).
// NewRequestWrapper crea una nueva instancia de RequestWrapper con el objeto interno configurado.
// Esta función es para uso interno del SDK al construir las solicitudes.
func NewRequestWrapper[T any](request *T) RequestWrapper[T] {
	return RequestWrapper[T]{request: request}
}

func (r RequestWrapper[T]) internal() *T {
	return r.request
}

// MarshalXML implementa la interfaz xml.Marshaler para delegar la serialización
// al objeto interno, evitando que la etiqueta raíz sea "RequestWrapper".
func (r RequestWrapper[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(r.request)
}

// UnwrapInternalRequest es un helper para extraer el request interno de un modelo opaco.
// Se desaconseja su uso directo fuera del SDK, ya que requiere tipos internos.
func UnwrapInternalRequest[T any](req any) *T {
	if r, ok := req.(interface{ internal() *T }); ok {
		return r.internal()
	}
	return nil
}
