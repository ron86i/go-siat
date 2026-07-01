package services

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	siatErrors "github.com/ron86i/go-siat/internal/core/errors"
	"github.com/ron86i/go-siat/internal/core/ports"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatService define los diferentes servicios disponibles en el SIAT.
type SiatService string

const (
	SiatCodigos            SiatService = "FacturacionCodigos"
	SiatOperaciones        SiatService = "FacturacionOperaciones"
	SiatSincronizacion     SiatService = "FacturacionSincronizacion"
	SiatCompraVenta        SiatService = "ServicioFacturacionCompraVenta"
	SiatComputarizada      SiatService = "ServicioFacturacionComputarizada"
	SiatElectronica        SiatService = "ServicioFacturacionElectronica"
	SiatDocumentoAjuste    SiatService = "ServicioFacturacionDocumentoAjuste"
	SiatTelecomunicaciones SiatService = "ServicioFacturacionTelecomunicaciones"
	SiatServicioBasico     SiatService = "ServicioFacturacionServicioBasico"
	SiatEntidadFinanciera  SiatService = "ServicioFacturacionEntidadFinanciera"
	SiatBoletoAereo        SiatService = "ServicioFacturacionBoletoAereo"
	SiatRecepcionCompras   SiatService = "ServicioRecepcionCompras"
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

// getInternalRequest desempaqueta la estructura de solicitud concreta desde una interfaz opaca
// delegando la operación al paquete models. Esto mantiene la opacidad hacia el usuario final
// mientras permite que las capas internas accedan a los datos necesarios para la comunicación.
func getInternalRequest[T any](req any) *T {
	return models.UnwrapInternalRequest[T](req)
}

// injectFields recorre recursivamente el struct para rellenar los campos comunes del contribuyente.
func injectFields(v reflect.Value, config ports.Config) {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := v.Type().Field(i)

		// Si es un struct anidado, procesarlo de forma recursiva
		if fv.Kind() == reflect.Struct {
			if fv.CanAddr() {
				injectFields(fv.Addr(), config)
			} else {
				injectFields(fv, config)
			}
			continue
		}

		if fv.CanSet() {
			switch ft.Name {
			case "Nit", "NIT":
				// Solo inyecta si el campo no fue seteado explícitamente en el request.
				if fv.Kind() == reflect.Int64 && fv.Int() == 0 {
					fv.SetInt(config.Nit)
				}
			case "CodigoSistema":
				// Solo inyecta si el campo no fue seteado explícitamente en el request.
				if fv.Kind() == reflect.String && fv.String() == "" {
					fv.SetString(config.CodigoSistema)
				}
			case "CodigoAmbiente":
				// Solo inyecta si el campo no fue seteado explícitamente en el request.
				switch fv.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if fv.Int() == 0 {
						fv.SetInt(int64(config.CodigoAmbiente))
					}
				}
			}
		}
	}
}

// injectCredenciales inyecta las credenciales del contribuyente en el request.
func injectCredenciales(req any, config ports.Config) {
	v := reflect.ValueOf(req)
	injectFields(v, config)
}

/*
performSoapRequest es una función genérica que encapsula el flujo completo de una solicitud SOAP al SIAT:

1. Obtiene la solicitud interna desde la interfaz opaca.

2. Inyecta automáticamente los parámetros globales de la sesión (NIT, Sistema, Ambiente, Modalidad).

3. Construye el cuerpo XML (Envelope SOAP).

4. Crea la solicitud HTTP POST con el contexto y headers necesarios (incluyendo el token de API).

5. Ejecuta la solicitud a través del cliente HTTP.

6. Procesa y decodifica la respuesta SOAP.
*/
func performSoapRequest[TReq any, TResp any](ctx context.Context, httpClient *http.Client, url string, config ports.Config, opaqueReq any) (*soap.EnvelopeResponse[TResp], error) {

	req := getInternalRequest[TReq](opaqueReq)
	injectCredenciales(req, config)

	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	ua := config.UserAgent
	if strings.TrimSpace(ua) == "" {
		ua = "go-siat"
	}
	httpReq.Header.Set("User-Agent", ua)
	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	// Inyectar X-Trace-ID si está disponible
	if strings.TrimSpace(config.TraceId) != "" {
		httpReq.Header.Set("X-Trace-ID", config.TraceId)
	}

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, siatErrors.NewNetworkError("fallo en solicitud HTTP al SIAT", err)
	}
	return parseSoapResponse[TResp](resp)
}
