package siat

import (
	"net/http"

	"github.com/ron86i/go-siat/v3/internal/adapter/services"
)

// HTTPConfig es la configuración pública para personalizar el cliente HTTP del SDK.
// Permite que los usuarios customicen timeouts, connection pooling y TLS sin pasar su propio http.Client.
//
// Campos principales:
//   - Timeout: Timeout total para solicitudes (default: 45s)
//   - MaxIdleConns: Pool global de conexiones ociosas (default: 100)
//   - MaxConnsPerHost: Conexiones simultáneas por host (default: 0, sin límite)
//   - TLSMinVersion: Versión mínima de TLS (default: TLS 1.2)
//
// Ejemplo:
//
//	cfg := siat.DefaultHTTPConfig()
//	cfg.Timeout = 60 * time.Second
//	cfg.MaxConnsPerHost = 5
//
//	s, _ := siat.New(siat.Config{
//		Token:          "...",
//		Nit:            123456789,
//		CodigoSistema:  "...",
//		CodigoAmbiente: siat.AmbientePruebas,
//		BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
//		HTTPClient:     siat.NewHTTPClient(cfg),
//	})
type HTTPConfig = services.HTTPConfig

// DefaultHTTPConfig retorna la configuración HTTP recomendada para producción.
//
// Valores:
//   - Timeout: 45 segundos
//   - MaxIdleConns: 100
//   - MaxConnsPerHost: 0 (Sin límite)
//   - MaxIdleConnsPerHost: 100
//   - TLSMinVersion: TLS 1.2
//
// Esta es la configuración usada por defecto si no pasas un http.Client a New().
func DefaultHTTPConfig() HTTPConfig {
	return services.DefaultHTTPConfig()
}

// NewHTTPClient crea un cliente HTTP optimizado basado en HTTPConfig.
// Es útil para customizar la configuración sin reescribir todo el Transport.
//
// Ejemplo:
//
//	cfg := siat.DefaultHTTPConfig()
//	cfg.Timeout = 60 * time.Second
//
//	config.HTTPClient = siat.NewHTTPClient(cfg)
//	s, err := siat.New(config)
func NewHTTPClient(cfg HTTPConfig) *http.Client {
	return services.NewHTTPClient(cfg)
}
