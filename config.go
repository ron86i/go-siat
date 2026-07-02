package siat

import (
	"fmt"
	"strings"

	"github.com/ron86i/go-siat/v2/internal/core/ports"
)

// Config agrupa la identidad global del contribuyente y la configuración de conexión.
type Config = ports.Config

// CredentialSign agrupa los distintos tipos de credenciales de firma digital (PEM o P12).
type CredentialSign = ports.CredentialSign

// Constructores de credenciales re-exportados para uso público del cliente:
var (
	NewPEMCredential  = ports.NewPEMCredential
	NewP12Credential  = ports.NewP12Credential
	WithDynamicConfig = ports.WithDynamicConfig
)

// validateConfig comprueba que todos los campos requeridos estén presentes y sean válidos.
func validateConfig(c Config) error {
	if strings.TrimSpace(c.Token) == "" {
		return fmt.Errorf("Token es obligatorio")
	}
	if c.Nit <= 0 {
		return fmt.Errorf("Nit es obligatorio y debe ser mayor a cero")
	}
	if strings.TrimSpace(c.CodigoSistema) == "" {
		return fmt.Errorf("CodigoSistema es obligatorio")
	}
	if c.CodigoAmbiente != AmbienteProduccion && c.CodigoAmbiente != AmbientePruebas {
		return fmt.Errorf("CodigoAmbiente inválido (%d)", c.CodigoAmbiente)
	}
	if strings.TrimSpace(c.BaseURL) == "" {
		return fmt.Errorf("BaseURL es obligatorio")
	}
	return nil
}
