package services

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// HTTPConfig agrupa configuraciones HTTP para máxima flexibilidad del usuario.
// Permite customizar timeouts, connection pooling y TLS sin cambiar el SDK.
type HTTPConfig struct {
	// Timeout total para solicitudes HTTP (default: 45 segundos)
	Timeout time.Duration

	// MaxIdleConns es el número máximo de conexiones ociosas en el pool global
	// (default: 100). Este es el pool compartido entre todos los hosts.
	MaxIdleConns int

	// MaxConnsPerHost es el número máximo de conexiones simultáneas por host
	// (default: 10). Limita la concurrencia al SIAT.
	MaxConnsPerHost int

	// MaxIdleConnsPerHost es el número máximo de conexiones ociosas por host
	// que se mantienen en el pool (default: 5).
	MaxIdleConnsPerHost int

	// TLSMinVersion especifica la versión mínima de TLS (default: TLS 1.2)
	TLSMinVersion uint16

	// DisableKeepAlives desactiva HTTP Keep-Alive (default: false para mejor performance)
	DisableKeepAlives bool

	// IdleConnTimeout es el tiempo después del cual se cierra una conexión ociosa
	// (default: 90 segundos)
	IdleConnTimeout time.Duration

	// DialTimeout es el timeout para crear nuevas conexiones
	// (default: 10 segundos)
	DialTimeout time.Duration

	// DialKeepAlive es el intervalo entre keep-alive probes TCP
	// (default: 30 segundos)
	DialKeepAlive time.Duration

	// TLSHandshakeTimeout es el timeout para el handshake TLS
	// (default: 10 segundos)
	TLSHandshakeTimeout time.Duration
}

// DefaultHTTPConfig retorna una configuración HTTP segura y optimizada para producción.
// Es el reemplazo recomendado para hardcoded timeouts de 15 segundos.
func DefaultHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Timeout:             45 * time.Second,
		MaxIdleConns:        100,
		MaxConnsPerHost:     10,
		MaxIdleConnsPerHost: 5,
		TLSMinVersion:       tls.VersionTLS12,
		DisableKeepAlives:   false,
		IdleConnTimeout:     90 * time.Second,
		DialTimeout:         10 * time.Second,
		DialKeepAlive:       30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}
}

// NewHTTPClient crea un cliente HTTP optimizado basado en HTTPConfig.
// Es útil cuando el usuario no proporciona su propio http.Client a New().
//
// Ejemplo:
//
//	cfg := service.DefaultHTTPConfig()
//	cfg.Timeout = 60 * time.Second // Customizar si necesario
//	client := service.NewHTTPClient(cfg)
//	s, err := siat.New(baseUrl, client)
func NewHTTPClient(cfg HTTPConfig) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        cfg.MaxIdleConns,
		MaxConnsPerHost:     cfg.MaxConnsPerHost,
		MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
		IdleConnTimeout:     cfg.IdleConnTimeout,
		DisableKeepAlives:   cfg.DisableKeepAlives,
		Proxy:               http.ProxyFromEnvironment, // Respetar proxy del sistema
		DialContext: (&net.Dialer{
			Timeout:   cfg.DialTimeout,
			KeepAlive: cfg.DialKeepAlive,
		}).DialContext,
		TLSHandshakeTimeout: cfg.TLSHandshakeTimeout,
		TLSClientConfig: &tls.Config{
			MinVersion: cfg.TLSMinVersion,
		},
	}

	return &http.Client{
		Transport: transport,
		Timeout:   cfg.Timeout,
	}
}
