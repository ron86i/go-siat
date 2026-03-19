package facturas

import (
	"encoding/xml"
)

// requestWrapper es una envoltura genérica utilizada para ocultar la implementación concreta
// de una solicitud y satisfacer las interfaces opacas del SDK.
type requestWrapper[T any] struct {
	request *T
}

// marshalXML implementa la interfaz xml.Marshaler para delegar la serialización
// al objeto interno, evitando que la etiqueta raíz sea "requestWrapper".
func (r requestWrapper[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(r.request)
}
