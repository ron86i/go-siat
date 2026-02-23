package service

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3/client"
	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
)

// fullURLCodigos construye la URL completa para acceder al servicio de facturación de códigos del SIAT,
// concatenando la URL base del ambiente con el endpoint específico del servicio.
func fullURLCodigos(url string) string {
	return url + "/FacturacionCodigos"
}

func fullURLOperaciones(url string) string {
	return url + "/FacturacionOperaciones"
}

func fullURLSincronizacion(url string) string {
	return url + "/FacturacionSincronizacion"
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
		return nil, fmt.Errorf("error al serializar body SOAP: %w", err)
	}
	return []byte(xml.Header + string(xmlBody)), nil
}

// parseSoapResponse procesa y valida una respuesta HTTP proveniente del servicio para extraer el contenido SOAP esperado.
// Esta función centraliza la validación de códigos de estado HTTP, la detección de errores de negocio (SOAP Fault)
// y el desempaquetado del XML en una estructura genérica, simplificando el flujo de los servicios.
func parseSoapResponse[T any](resp *client.Response) (*soap.EnvelopeResponse[T], error) {
	var result soap.EnvelopeResponse[T]

	// Intentar parsear la respuesta XML en la estructura de respuesta SOAP (puede ser éxito o Fault)
	errUnmarshal := xml.Unmarshal(resp.Body(), &result)

	// Si el servicio devolvió un SOAP Fault, priorizar este error descriptivo de negocio
	if errUnmarshal == nil && result.Body.Fault != nil {
		return nil, fmt.Errorf("SOAP Fault [%s]: %s", result.Body.Fault.FaultCode, result.Body.Fault.FaultString)
	}

	// Si el código de estado HTTP no es 200, informar el error de estado.
	// Esto maneja correctamente casos donde el cuerpo está vacío (EOF) en errores 500.
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("HTTP status inesperado: %d", resp.StatusCode())
	}

	// Si el status es 200 pero hubo un error de parseo, informar el error de XML
	if errUnmarshal != nil {
		log.Printf("[ERROR] Error parseando XML respuesta: %v", errUnmarshal)
		return nil, fmt.Errorf("error al parsear respuesta SOAP: %w", errUnmarshal)
	}

	return &result, nil
}
