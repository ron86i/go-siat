package errors

import "fmt"

// SiatError es el tipo de error que retorna el SDK.
// Permite al usuario distinguir entre diferentes tipos de errores y tomar decisiones.
type SiatError struct {
	// Code es un identificador único del error (ej: "INVALID_CUIS", "NETWORK_ERROR", "AUTH_FAILED")
	Code string

	// Message es la descripción del error en lenguaje natural
	Message string

	// SiatCode es el código de error retornado por el SIAT (si aplica)
	SiatCode int

	// StatusCode es el código HTTP (si aplica)
	StatusCode int

	// IsNetworkError indica si el error fue de red (timeout, connection refused, etc)
	IsNetworkError bool

	// IsRetryable indica si la operación puede/debe ser reintentada
	IsRetryable bool

	// Details contiene información adicional para debugging
	Details map[string]interface{}

	// WrappedErr es el error subyacente (si existe)
	WrappedErr error
}

// Error implementa la interfaz error
func (e *SiatError) Error() string {
	if e == nil {
		return ""
	}
	msg := e.Message
	if e.Code != "" {
		msg = fmt.Sprintf("[%s] %s", e.Code, msg)
	}
	if e.SiatCode != 0 {
		msg += fmt.Sprintf(" (SIAT code: %d)", e.SiatCode)
	}
	return msg
}

// Unwrap retorna el error subyacente (para errors.Is y errors.As)
func (e *SiatError) Unwrap() error {
	return e.WrappedErr
}

// NewNetworkError crea un error de red
func NewNetworkError(msg string, err error) *SiatError {
	return &SiatError{
		Code:           "NETWORK_ERROR",
		Message:        msg,
		IsNetworkError: true,
		IsRetryable:    true,
		WrappedErr:     err,
	}
}

// NewSiatError crea un error del servidor SIAT
func NewSiatError(code int, msg string) *SiatError {
	return &SiatError{
		Code:        "SIAT_SERVER_ERROR",
		Message:     msg,
		SiatCode:    code,
		IsRetryable: false, // El servidor rechazó por algo específico
	}
}

// NewAuthError crea un error de autenticación
func NewAuthError(msg string) *SiatError {
	return &SiatError{
		Code:        "AUTH_FAILED",
		Message:     msg,
		IsRetryable: false, // No tiene sentido reintentar sin credentials válidas
	}
}

// NewTimeoutError crea un error de timeout
func NewTimeoutError(msg string) *SiatError {
	return &SiatError{
		Code:           "TIMEOUT",
		Message:        msg,
		IsNetworkError: true,
		IsRetryable:    true,
	}
}

// IsRetryable es un helper que retorna true si el error es reintentable
func IsRetryable(err error) bool {
	if se, ok := err.(*SiatError); ok {
		return se.IsRetryable
	}
	return false
}

// IsNetworkError es un helper que retorna true si fue error de red
func IsNetworkError(err error) bool {
	if se, ok := err.(*SiatError); ok {
		return se.IsNetworkError
	}
	return false
}
