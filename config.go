package siat

import "github.com/ron86i/go-siat/internal/core/ports"

// Config agrupa la configuración necesaria para realizar solicitudes autenticadas al SIAT.
// Es un alias a ports.Config y debe ser instanciada cuando se realicen llamadas a los servicios.
//
// Campos:
//   - Token: Token de autenticación proporcionado por el SIAT (obligatorio)
//   - UserAgent: Identificador del cliente HTTP (opcional, pero recomendado para registro)
//
// Ejemplo:
//
//	cfg := siat.Config{
//		Token: "tu_token_api",
//		UserAgent: "MyApp/1.0 (Bolivia)",
//	}
type Config = ports.Config
