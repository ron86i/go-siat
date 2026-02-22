package operaciones

import (
	"encoding/xml"
)

// VerificarComunicacion es el wrapper para la petición (Request)
type VerificarComunicacion struct {
	XMLName xml.Name `xml:"ns:verificarComunicacion"`
}

// VerificarComunicacionResponse representa el Body de la respuesta SOAP
type VerificarComunicacionResponse struct {
	XMLName xml.Name `xml:"verificarComunicacionResponse"`
	// El SIAT siempre devuelve los datos dentro de un nodo <return>
	Return RespuestaComunicacion `xml:"return"`
}

// RespuestaComunicacion representa la estructura interna de la respuesta
type RespuestaComunicacion struct {
	MensajesList []MensajeServicio `xml:"mensajesList" json:"mensajesList"`
	Transaccion  bool              `xml:"transaccion" json:"transaccion"`
}
