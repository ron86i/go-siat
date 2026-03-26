package siat

import (
	"net/http"

	"github.com/ron86i/go-siat/internal/core/middleware"
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
// Es útil para agregar logging, métricas, retry, circuit breaker, etc.
//
// Parámetros:
//   - baseUrl: URL base del SIAT
//   - httpClient: Cliente HTTP (opcional, usa defaults si es nil)
//   - middlewares: Middlewares a aplicar en orden (primero = más externo)
//
// Los middlewares se aplican al Transport del cliente, permitiendo interceptar todas las solicitudes.
//
// Ejemplo completo con logging:
//
//	type LoggingMiddleware struct{}
//	func (LoggingMiddleware) WrapTransport(base http.RoundTripper) http.RoundTripper {
//	    return &loggingRoundTripper{base: base}
//	}
//	type loggingRoundTripper struct { base http.RoundTripper }
//	func (rt *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
//	    log.Printf("-> %s %s", req.Method, req.URL.Path)
//	    resp, err := rt.base.RoundTrip(req)
//	    if err == nil {
//	        log.Printf("<- %d", resp.StatusCode)
//	    }
//	    return resp, err
//	}
//
//	// Uso
//	s, err := siat.NewWithMiddleware(baseUrl, nil, &LoggingMiddleware{})
func NewWithMiddleware(baseUrl string, httpClient *http.Client, middlewares ...HTTPMiddleware) (*SiatServices, error) {
	// Crear cliente con defaults si no se proporciona
	if httpClient == nil {
		httpClient = NewHTTPClient(DefaultHTTPConfig())
	} else {
		// Clonar para no modificar el cliente del usuario
		clonedClient := *httpClient
		httpClient = &clonedClient
	}

	// Aplicar middlewares si los hay
	if len(middlewares) > 0 && httpClient.Transport != nil {
		chainedMW := middleware.ChainMiddlewares(middlewares...)
		if chainedMW != nil {
			httpClient.Transport = chainedMW.WrapTransport(httpClient.Transport)
		}
	}

	// Usar el New() existente
	return New(baseUrl, httpClient)
}
