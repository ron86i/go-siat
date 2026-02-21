package codigos

import "encoding/xml"

// VerificarComunicacion representa una solicitud básica para testear la conectividad con los servicios del SIAT.
type VerificarComunicacion struct {
	XMLName xml.Name `xml:"ns:verificarComunicacion" json:"-"`
}

// VerificarComunicacionResponse define la estructura del sobre de respuesta tras una prueba de comunicación.
type VerificarComunicacionResponse struct {
	XMLName               xml.Name              `xml:"verificarComunicacionResponse" json:"-"`
	RespuestaComunicacion RespuestaComunicacion `xml:"RespuestaComunicacion" json:"respuestaComunicacion"`
}

// RespuestaComunicacion informa sobre el estado de la conexión, indicando si la transacción
// de prueba fue exitosa y reportando cualquier mensaje técnico del servidor.
type RespuestaComunicacion struct {
	MensajesList []MensajeServicio `xml:"mensajesList,omitempty" json:"mensajesList,omitempty"`
	Transaccion  bool              `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}
