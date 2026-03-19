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

func (r requestWrapper[T]) internal() *T {
	return r.request
}

// GetInternalRequest es un helper para extraer el request interno de un modelo opaco.
// Aunque es público para permitir su uso desde el paquete de servicios, no se recomienda
// su uso directo por parte de los usuarios del SDK.
func GetInternalRequest[T any](req any) *T {
	if r, ok := req.(interface{ internal() *T }); ok {
		return r.internal()
	}
	return nil
}
