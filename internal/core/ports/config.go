package ports

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ron86i/go-siat/pkg/utils"
)

// CredentialSign representa las credenciales de firma digital de manera tipada y estructurada.
type CredentialSign struct {
	CertBytes       []byte
	PrivateKeyBytes []byte
	P12Bytes        []byte
	P12Password     string
	err             error // Guarda errores de lectura de archivo durante la inicialización
}

// NewPEMCredential crea una credencial de firma a partir de un certificado y llave privada (.crt/.key, pueden ser rutas string o bytes []byte).
func NewPEMCredential(cert, privateKey any) CredentialSign {
	cf := CredentialSign{}
	switch c := cert.(type) {
	case string:
		data, err := os.ReadFile(c)
		if err != nil {
			cf.err = fmt.Errorf("error al leer archivo certificado PEM: %w", err)
		} else {
			cf.CertBytes = data
		}
	case []byte:
		cf.CertBytes = c
	}
	switch k := privateKey.(type) {
	case string:
		data, err := os.ReadFile(k)
		if err != nil {
			cf.err = fmt.Errorf("error al leer archivo llave privada PEM: %w", err)
		} else {
			cf.PrivateKeyBytes = data
		}
	case []byte:
		cf.PrivateKeyBytes = k
	}
	return cf
}

// NewP12Credential crea una credencial de firma a partir de un PKCS#12 (.p12/.pfx, puede ser ruta string o bytes []byte) y su contraseña.
func NewP12Credential(p12 any, password string) CredentialSign {
	cf := CredentialSign{
		P12Password: password,
	}
	switch p := p12.(type) {
	case string:
		data, err := os.ReadFile(p)
		if err != nil {
			cf.err = fmt.Errorf("error al leer archivo P12: %w", err)
		} else {
			cf.P12Bytes = data
		}
	case []byte:
		cf.P12Bytes = p
	}
	return cf
}

// GetType retorna el tipo de la credencial ("PEM", "P12" o "UNKNOWN").
func (cf CredentialSign) GetType() string {
	if len(cf.P12Bytes) > 0 {
		return "P12"
	}
	if len(cf.CertBytes) > 0 && len(cf.PrivateKeyBytes) > 0 {
		return "PEM"
	}
	return "UNKNOWN"
}

// SignXML firma un documento XML utilizando esta credencial.
func (cf CredentialSign) SignXML(xmlBytes []byte) ([]byte, error) {
	if cf.err != nil {
		return nil, cf.err
	}

	// 1. Intentar firmar con P12
	if len(cf.P12Bytes) > 0 {
		return utils.SignWithP12Bytes(xmlBytes, cf.P12Bytes, cf.P12Password)
	}

	// 2. Intentar firmar con PEM (cert/key)
	if len(cf.CertBytes) > 0 && len(cf.PrivateKeyBytes) > 0 {
		return utils.SignXMLBytes(xmlBytes, cf.PrivateKeyBytes, cf.CertBytes)
	}

	return nil, fmt.Errorf("no se configuraron credenciales válidas de firma digital (.crt o .p12)")
}

// Config agrupa la identidad global del contribuyente y la configuración de conexión.
type Config struct {
	// Token de autenticación proporcionado por el SIAT (obligatorio)
	Token string

	// Nit del contribuyente emisor (obligatorio)
	Nit int64

	// CodigoSistema autorizado por el SIN (obligatorio)
	CodigoSistema string

	// CodigoAmbiente: siat.AmbienteProduccion o siat.AmbientePruebas (obligatorio)
	CodigoAmbiente int

	// BaseURL del SIAT (obligatorio)
	BaseURL string

	// UserAgent personalizado (opcional, default: "go-siat")
	UserAgent string

	// TraceId para correlacionar solicitudes (opcional)
	TraceId string

	// HTTPClient personalizado (opcional)
	HTTPClient *http.Client

	// Credenciales de firma digital (opcional, requerida para facturación electrónica)
	CredentialSign CredentialSign
}

// SignXML firma un documento XML utilizando las credenciales configuradas en Config.
func (c Config) SignXML(xmlBytes []byte) ([]byte, error) {
	return c.CredentialSign.SignXML(xmlBytes)
}
