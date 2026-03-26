package middleware

import "net/http"

// HTTPMiddleware es una interfaz que permite al usuario extender el comportamiento HTTP del SDK.
// Es el patrón go-idiomatic de "middleware" para RoundTripper.
//
// Ejemplo: Agregar logging a todas las solicitudes
//
//	type LoggingMiddleware struct {
//	    logger Logger
//	}
//
//	func (m *LoggingMiddleware) WrapTransport(base http.RoundTripper) http.RoundTripper {
//	    return &loggingRoundTripper{
//	        base: base,
//	        logger: m.logger,
//	    }
//	}
//
//	type loggingRoundTripper struct {
//	    base http.RoundTripper
//	    logger Logger
//	}
//
//	func (rt *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
//	    log.Printf("Request: %s %s", req.Method, req.URL)
//	    resp, err := rt.base.RoundTrip(req)
//	    log.Printf("Response: %d", resp.StatusCode)
//	    return resp, err
//	}
//
//	// Uso en el SDK
//	middleware := &LoggingMiddleware{logger: myLogger}
//	client := &http.Client{Transport: &http.Transport{...}}
//	wrappedClient := siat.NewWithMiddleware("https://...", client, middleware)
type HTTPMiddleware interface {
	// WrapTransport recibe el transport base y retorna un transport envuelto.
	// Permite interceptar, modificar y registrar todas las solicitudes y respuestas HTTP.
	WrapTransport(base http.RoundTripper) http.RoundTripper
}

// ChainMiddlewares encadena múltiples middlewares en orden.
// El primer middleware es el más externo (wrapper inicial).
func ChainMiddlewares(middlewares ...HTTPMiddleware) HTTPMiddleware {
	if len(middlewares) == 0 {
		return nil
	}
	if len(middlewares) == 1 {
		return middlewares[0]
	}
	return &chainedMiddleware{middlewares: middlewares}
}

type chainedMiddleware struct {
	middlewares []HTTPMiddleware
}

func (cm *chainedMiddleware) WrapTransport(base http.RoundTripper) http.RoundTripper {
	// Aplicar middlewares en orden inverso para que el primero sea el más externo
	for i := len(cm.middlewares) - 1; i >= 0; i-- {
		base = cm.middlewares[i].WrapTransport(base)
	}
	return base
}
