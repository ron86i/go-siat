package siat

import (
	"github.com/ron86i/go-siat/internal/core/errors"
)

// SiatError es el tipo de error que retorna el SDK.
// Permite al usuario distinguir entre diferentes tipos de errores:
// - Errores de red (timeout, conexión rechazada)
// - Errores del servidor SIAT (validaciones rechazadas)
// - Errores de autenticación (token inválido)
// - Y saber cuáles pueden/deben ser reintentados
//
// Ejemplo:
//
//	resp, err := s.Codigos().SolicitudCuis(ctx, cfg, req)
//	if err != nil {
//	    var siatErr *siat.SiatError
//	    if errors.As(err, &siatErr) {
//	        if siat.IsRetryable(err) {
//	            // Reintentar después
//	        } else if siat.IsNetworkError(err) {
//	            // Error temporal de conexión
//	        } else {
//	            // Error del servidor o autenticación
//	        }
//	    }
//	}
type SiatError = errors.SiatError

// NewNetworkError crea un error de red (timeout, connection refused, etc).
// Estos errores son reintentables.
func NewNetworkError(msg string, err error) *SiatError {
	return errors.NewNetworkError(msg, err)
}

// NewSiatError crea un error del servidor SIAT.
// Indica que el servidor rechazó la solicitud por alguna razón específica.
func NewSiatError(code int, msg string) *SiatError {
	return errors.NewSiatError(code, msg)
}

// NewAuthError crea un error de autenticación.
// No tiene sentido reintentar sin credentials válidas.
func NewAuthError(msg string) *SiatError {
	return errors.NewAuthError(msg)
}

// NewTimeoutError crea un error de timeout.
// Es reintentable porque es un problema temporal.
func NewTimeoutError(msg string) *SiatError {
	return errors.NewTimeoutError(msg)
}

// IsRetryable indica si un error puede/debe ser reintentado.
// Útil para implementar lógica de retry en la aplicación.
func IsRetryable(err error) bool {
	return errors.IsRetryable(err)
}

// IsNetworkError indica si un error fue de red.
// Útil para distinguir entre problemas de conectividad y errores del servidor.
func IsNetworkError(err error) bool {
	return errors.IsNetworkError(err)
}
