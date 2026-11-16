// Package siatpolicy validates SIAT protocol endpoints and facturation routes.
package siatpolicy

import (
	"fmt"
	"net"
	"net/url"
	"reflect"
	"strings"
)

const (
	ProtocolV3 = 3

	ServicioCodigos            = "FacturacionCodigos"
	ServicioOperaciones        = "FacturacionOperaciones"
	ServicioSincronizacion     = "FacturacionSincronizacion"
	ServicioRecepcionCompras   = "ServicioRecepcionCompras"
	ServicioCompraVenta        = "ServicioFacturacionCompraVenta"
	ServicioElectronica        = "ServicioFacturacionElectronica"
	ServicioComputarizada      = "ServicioFacturacionComputarizada"
	ServicioDocumentoAjuste    = "ServicioFacturacionDocumentoAjuste"
	ServicioTelecomunicaciones = "ServicioFacturacionTelecomunicaciones"
	ServicioBasico             = "ServicioFacturacionServicioBasico"
	ServicioEntidadFinanciera  = "ServicioFacturacionEntidadFinanciera"
	ServicioBoletoAereo        = "ServicioFacturacionBoletoAereo"
	MaxFacturasPaquete         = 500
	MaxFacturasMasiva          = 1000
)

var servicios = map[string]struct{}{
	ServicioCodigos: {}, ServicioOperaciones: {}, ServicioSincronizacion: {}, ServicioRecepcionCompras: {},
	ServicioCompraVenta: {}, ServicioElectronica: {}, ServicioComputarizada: {}, ServicioDocumentoAjuste: {},
	ServicioTelecomunicaciones: {}, ServicioBasico: {}, ServicioEntidadFinanciera: {}, ServicioBoletoAereo: {},
}

// Documento es la interfaz común que implementan los modelos fiscales del SDK.
type Documento interface {
	Validate() error
	CodigoDocumentoSector() int
	CodigoSucursal() int
	CodigoPuntoVenta() *int
	Cufd() string
	Protocolo() int
	MarshalXMLModalidad(modalidad int) ([]byte, error)
}

// ResolveBaseURL normalizes a base URL for SIAT 3.0.
// Root URLs are allowed only for loopback hosts used by local mock servers.
func ResolveBaseURL(raw string, ambiente int) (string, error) {
	const path = "/sfd"
	if ambiente != 1 && ambiente != 2 {
		return "", fmt.Errorf("ambiente SIAT inválido: %d", ambiente)
	}
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil || u == nil || (u.Scheme != "https" && u.Scheme != "http") || u.Host == "" || u.User != nil || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" || u.Opaque != "" {
		return "", fmt.Errorf("BaseURL SIAT inválida")
	}
	if u.Scheme == "http" && !isLoopback(u.Hostname()) {
		return "", fmt.Errorf("BaseURL SIAT debe usar HTTPS fuera de loopback")
	}
	if u.EscapedPath() != u.Path {
		return "", fmt.Errorf("BaseURL SIAT no admite segmentos escapados")
	}
	if u.Path == "" || u.Path == "/" {
		if !isLoopback(u.Hostname()) {
			return "", fmt.Errorf("BaseURL debe terminar exactamente en %s", path)
		}
		u.Path = path
	} else if u.Path != path && u.Path != path+"/" {
		return "", fmt.Errorf("BaseURL incompatible con SIAT v3: se requiere la ruta %s", path)
	} else {
		u.Path = path
	}
	u.RawPath = ""
	return strings.TrimSuffix(u.String(), "/"), nil
}

// ValidService rejects service names outside the SDK's SIAT service set.
func ValidService(service string) bool {
	if strings.Contains(service, "/") || strings.Contains(service, "\\") {
		return false
	}
	_, ok := servicios[service]
	return ok
}

// ValidateDestination checks the document sector, SOAP service, and modality.
func ValidateDestination(service string, sector, modalidad int) error {
	if sector <= 0 {
		return fmt.Errorf("código de documento sector inválido: %d", sector)
	}
	if modalidad != 1 && modalidad != 2 {
		return fmt.Errorf("modalidad SIAT inválida: %d", modalidad)
	}
	if !ValidService(service) {
		return fmt.Errorf("servicio SIAT desconocido: %q", service)
	}
	wantService, ok := serviceForSector(sector)
	if !ok {
		return fmt.Errorf("sector SIAT desconocido: %d", sector)
	}
	general := wantService == ServicioElectronica
	if service != wantService && !(general && service == ServicioComputarizada) {
		return fmt.Errorf("el sector %d requiere el servicio %s, no %s", sector, wantService, service)
	}
	if service == ServicioElectronica && modalidad != 1 {
		return fmt.Errorf("el servicio electrónico requiere modalidad 1")
	}
	if service == ServicioComputarizada && modalidad != 2 {
		return fmt.Errorf("el servicio computarizado requiere modalidad 2")
	}
	if sector == 52 && (service != ServicioElectronica || modalidad != 1) {
		return fmt.Errorf("el sector 52 solo admite modalidad electrónica (1)")
	}
	return nil
}

// ValidateDocumentDestination checks that the sector is supported by its service for this modality.
func ValidateDocumentDestination(sector, modalidad int) error {
	service, ok := serviceForSector(sector)
	if !ok {
		return fmt.Errorf("sector SIAT desconocido: %d", sector)
	}
	if service == ServicioElectronica && modalidad == 2 {
		service = ServicioComputarizada
	}
	return ValidateDestination(service, sector, modalidad)
}

// ValidateDocumentMode checks the mode declared by the SDK model when it has one.
func ValidateDocumentMode(raw any, modalidad int) error {
	value := reflect.ValueOf(raw)
	for value.IsValid() && (value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface) {
		if value.IsNil() {
			return fmt.Errorf("el documento no puede ser nil")
		}
		value = value.Elem()
	}
	if !value.IsValid() || value.Kind() != reflect.Struct {
		return nil
	}
	mode := value.FieldByName("Modalidad")
	if mode.IsValid() && mode.Kind() == reflect.Int && mode.Int() != 0 && int(mode.Int()) != modalidad {
		return fmt.Errorf("modalidad del documento (%d) no coincide con la solicitud (%d)", mode.Int(), modalidad)
	}
	return nil
}

// ValidateDocument rejects nil values, foreign protocol models, and external wrappers.
func ValidateDocument(raw any) (Documento, error) {
	if raw == nil {
		return nil, fmt.Errorf("el documento no puede ser nil")
	}
	value := reflect.ValueOf(raw)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		if value.IsNil() {
			return nil, fmt.Errorf("el documento no puede ser nil")
		}
	}
	doc, ok := raw.(Documento)
	if !ok {
		return nil, fmt.Errorf("tipo de documento no admitido: %T", raw)
	}
	if doc.Protocolo() != ProtocolV3 {
		return nil, fmt.Errorf("documento incompatible con SIAT v3: el documento pertenece al protocolo v%d", doc.Protocolo())
	}
	typ := reflect.TypeOf(raw)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	const wantPrefix = "github.com/ron86i/go-siat/v3/documents/"
	if typ.PkgPath() == "" || !strings.HasPrefix(typ.PkgPath(), wantPrefix) {
		return nil, fmt.Errorf("el documento debe ser un modelo SIAT v3 del SDK")
	}
	if err := doc.Validate(); err != nil {
		return nil, fmt.Errorf("documento inválido: %w", err)
	}
	return doc, nil
}

func isLoopback(host string) bool {
	if strings.EqualFold(host, "localhost") {
		return true
	}
	ip := net.ParseIP(strings.Trim(host, "[]"))
	return ip != nil && ip.IsLoopback()
}

func serviceForSector(sector int) (string, bool) {
	switch sector {
	case 1, 35, 41:
		return ServicioCompraVenta, true
	case 13, 40:
		return ServicioBasico, true
	case 15:
		return ServicioEntidadFinanciera, true
	case 22, 49:
		return ServicioTelecomunicaciones, true
	case 24, 29, 48:
		return ServicioDocumentoAjuste, true
	case 30:
		return ServicioBoletoAereo, true
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 14, 16, 17, 18, 19, 20, 21, 23, 28, 31, 34, 36, 37, 38, 39, 42, 43, 44, 45, 46, 50, 52, 53, 54, 55:
		return ServicioElectronica, true
	default:
		return "", false
	}
}
