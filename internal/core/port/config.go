package port

// Config agrupa la configuración para realizar solicitudes autenticadas al SIAT.
// Debe ser instanciada para cada operación que requiera autenticación.
//
// Campos:
//   - Token: Token de autenticación obtenido del SIAT (obligatorio)
//   - UserAgent: Identificador opcional del cliente HTTP para registro y debugging
//   - TraceId: ID opcional para correlacionar solicitudes en sistemas distribuidos
type Config struct {
	// Token es el código de autenticación proporcionado por el SIAT
	Token string
	// UserAgent es el identificador opcional del cliente para propósitos de logging
	UserAgent string
	// TraceId es un identificador opcional para correlacionar solicitudes en sistemas distribuidos.
	// Si está establecido, se inyecta en el encabezado HTTP "X-Trace-ID".
	TraceId string
}
