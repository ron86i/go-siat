package soap

import "encoding/xml"

// Envelope genérico para las PETICIONES (Marshal)
// Ahora acepta un tipo genérico T para que no tengas que usar 'any' en el Body
type Envelope[T any] struct {
	XMLName      xml.Name `xml:"soapenv:Envelope"`
	XmlnsSoapenv string   `xml:"xmlns:soapenv,attr"`
	XmlnsNs      string   `xml:"xmlns:ns,attr"`
	// Usamos un puntero para omitir la etiqueta completamente si el Header está vacío
	Header *Header         `xml:"soapenv:Header,omitempty"`
	Body   EnvelopeBody[T] `xml:"soapenv:Body"`
}

type Header struct {
	// Puedes agregar campos aquí si el SIAT requiere un token en la cabecera en el futuro
}

// Envelope genérico para las RESPUESTAS (Unmarshal)
type EnvelopeResponse[T any] struct {
	XMLName   xml.Name        `xml:"Envelope"`
	XmlnsSoap string          `xml:"xmlns:soapenv,attr,omitempty"`
	XmlnsNs   string          `xml:"xmlns:ns,attr,omitempty"`
	Header    Header          `xml:"Header,omitempty"`
	Body      EnvelopeBody[T] `xml:"Body"`
}

// EnvelopeBody representa el cuerpo del sobre SOAP, conteniendo el éxito (Content) o el fallo (Fault).
type EnvelopeBody[T any] struct {
	// Content contiene el contenido interno en caso de éxito (p.ej. verificarNitResponse).
	// ',any' permite inyectar o extraer cualquier tipo genérico.
	Content T `xml:",any"`
	// Fault captura errores estándar de SOAP cuando la operación falla en el servidor.
	Fault *Fault `xml:"Fault,omitempty"`
}

// Fault representa un error SOAP estándar (SOAP Fault 1.1).
// Contiene detalles técnicos sobre por qué falló la solicitud en el servidor SIAT.
type Fault struct {
	XMLName     xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FaultCode   string   `xml:"faultcode"`   // Código identificador del error
	FaultString string   `xml:"faultstring"` // Descripción legible del error
	Detail      string   `xml:"detail"`      // Detalles adicionales o trazas de error
}
