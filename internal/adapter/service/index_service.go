package service

import (
	"encoding/xml"

	"io"
	"net/http"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
)

// SiatService define los diferentes servicios disponibles en el SIAT.
type SiatService string

const (
	SiatCodigos        SiatService = "FacturacionCodigos"
	SiatOperaciones    SiatService = "FacturacionOperaciones"
	SiatSincronizacion SiatService = "FacturacionSincronizacion"
	SiatCompraVenta    SiatService = "ServicioFacturacionCompraVenta"
	SiatComputarizada  SiatService = "ServicioFacturacionComputarizada"
)

// fullURL construye la URL completa para acceder a un servicio específico del SIAT,
// concatenando la URL base del ambiente con el endpoint del servicio solicitado.
func fullURL(baseURL string, service SiatService) string {
	return baseURL + "/" + string(service)
}

// buildRequest encapsula un objeto de solicitud genérico dentro de un sobre SOAP estándar (Envelope),
// añadiendo los namespaces requeridos por el SIAT y serializando el resultado a formato XML.
func buildRequest(req any) ([]byte, error) {
	requestBody := soap.Envelope[any]{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsNs:      "https://siat.impuestos.gob.bo/",
		Body: soap.EnvelopeBody[any]{
			Content: req,
		},
	}

	xmlBody, err := xml.MarshalIndent(requestBody, "", "  ")
	if err != nil {
		return nil, err
	}
	return []byte(xml.Header + string(xmlBody)), nil
}

// parseSoapResponse procesa y valida una respuesta HTTP proveniente del servicio para extraer el contenido SOAP esperado.
func parseSoapResponse[T any](resp *http.Response) (*soap.EnvelopeResponse[T], error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result soap.EnvelopeResponse[T]

	// Intentar parsear la respuesta XML en la estructura de respuesta SOAP
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
