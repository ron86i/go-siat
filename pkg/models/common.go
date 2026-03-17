package models

import (
	"encoding/xml"
)

// requestWrapper es una envoltura genérica utilizada para ocultar la implementación concreta
// de una solicitud y satisfacer las interfaces opacas del SDK.
type requestWrapper[T any] struct {
	request *T
}

// MarshalXML implementa la interfaz xml.Marshaler para delegar la serialización
// al objeto interno, evitando que la etiqueta raíz sea "requestWrapper".
func (r requestWrapper[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(r.request)
}

func (r requestWrapper[T]) Internal() *T {
	return r.request
}

// GetInternalRequest desempaqueta la estructura de solicitud concreta desde una interfaz opaca.
// Este método es utilizado internamente por los servicios para acceder a los campos de la solicitud.
// Soporta tanto envolturas (wrappers) como punteros directos para mayor flexibilidad.
func GetInternalRequest[T any](req any) *T {
	if getter, ok := req.(interface{ Internal() *T }); ok {
		return getter.Internal()
	}
	if wrapper, ok := req.(requestWrapper[T]); ok {
		return wrapper.request
	}
	if res, ok := req.(*T); ok {
		return res
	}
	return nil
}
