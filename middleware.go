package siat

import (
	"net/http"

	"github.com/ron86i/go-siat/v2/internal/core/middleware"
)

// HTTPMiddleware es la interfaz pública que permite agregar comportamiento HTTP personalizado.
// Los usuarios pueden implementarla para agregar logging, métricas, retry, etc.
// sin necesidad de modificar el SDK.
//
// Ejemplo: Implementar retry con exponential backoff:
//
//	type RetryMiddleware struct {
//	    maxAttempts int
//	}
//
//	func (m *RetryMiddleware) WrapTransport(base http.RoundTripper) http.RoundTripper {
//	    return &retryRoundTripper{
//	        base: base,
//	        maxAttempts: m.maxAttempts,
//	    }
//	}
//
//	// Uso
//	s, err := siat.NewWithMiddleware(
//	    baseUrl,
//	    nil, // usa defaults
//	    &RetryMiddleware{maxAttempts: 3},
//	)
type HTTPMiddleware = middleware.HTTPMiddleware

// NewWithMiddleware crea una instancia de SiatServices con middlewares HTTP personalizados.
//
// Parámetros:
//   - config: Configuración global del cliente.
//   - middlewares: Middlewares a aplicar en orden (primero = más externo).
//
// Los middlewares se aplican al Transport del cliente, permitiendo interceptar todas las solicitudes.
func NewWithMiddleware(config Config, middlewares ...HTTPMiddleware) (*SiatServices, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	// Crear cliente con defaults si no se proporciona
	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = NewHTTPClient(DefaultHTTPConfig())
	} else {
		// Clonar para no modificar el cliente del usuario
		clonedClient := *httpClient
		httpClient = &clonedClient
	}

	// Aplicar middlewares si los hay
	if len(middlewares) > 0 {
		base := httpClient.Transport
		if base == nil {
			base = http.DefaultTransport
		}
		chainedMW := middleware.ChainMiddlewares(middlewares...)
		if chainedMW != nil {
			httpClient.Transport = chainedMW.WrapTransport(base)
		}
	}

	// Inyectar el cliente HTTP configurado en la estructura de configuración
	config.HTTPClient = httpClient

	// Usar el New() existente
	return New(config)
}
