package siat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/ron86i/go-siat/internal/adapter/services"
	"github.com/ron86i/go-siat/internal/core/domain/siat/common"
	"github.com/ron86i/go-siat/internal/core/ports"
)

// Map es un alias para map[string]interface{} que proporciona métodos de utilidad
// para trabajar con datos JSON de forma más cómda.
// Es especialmente útil al trabajar con respuestas heterogéneas del SIAT.
type Map map[string]interface{}

// ToJSON convierte el Map a su representación en string JSON.
// Retorna un error si la codificación falla.
func (m Map) ToJSON() (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Sum retorna la suma de todos los valores numéricos en el Map.
// Soporta tipos float64, float32, int, int64 e int32.
// Los valores no numéricos se ignoran.
func (m Map) Sum() float64 {
	var total float64
	for _, v := range m {
		switch val := v.(type) {
		case float64:
			total += val
		case float32:
			total += float64(val)
		case int:
			total += float64(val)
		case int64:
			total += float64(val)
		case int32:
			total += float64(val)
		}
	}
	return total
}

// ToStruct convierte el Map en la estructura Go especificada.
// Utiliza encoding/json internamente, por lo que se requiere que v sea un puntero
// a una estructura con etiquetas json apropiadas.
func (m Map) ToStruct(v interface{}) error {
	bytes, err := m.ToJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(bytes), v)
}

// SiatServices es el punto de entrada principal del SDK.
// Agrupa todas las implementaciones de los servicios del SIAT
// (Códigos, Sincronización, Operaciones, Compra-Venta, Computarizada, Electrónica)
// y proporciona acceso a ellos a través de métodos orientados a objetivos.
// Los usuarios deben crear una instancia usando New().
type SiatServices struct {
	operaciones        ports.SiatOperacionesPort
	sincronizacion     ports.SiatSincronizacionService
	codigos            ports.SiatCodigosService
	compraVenta        ports.SiatCompraVentaService
	computarizada      ports.SiatComputarizadaService
	electronica        ports.SiatElectronicaService
	documentoAjuste    ports.SiatDocumentoAjusteService
	telecomunicaciones ports.SiatTelecomunicacionesService
	servicioBasico     ports.SiatServicioBasicoService
	entidadFinanciera  ports.SiatEntidadFinancieraService
	boletoAereo        ports.SiatBoletoAereoService
	recepcionCompras   ports.SiatRecepcionComprasService
	traceID            string // Opcional, para correlacionar solicitudes en sistemas distribuidos
}

// Operaciones retorna el servicio para la gestión de puntos de venta (PV),
// cierre de períodos de facturación y eventos significativos (cambios de modalidad, etc.).
func (s *SiatServices) Operaciones() ports.SiatOperacionesPort {
	return s.operaciones
}

// Sincronizacion retorna el servicio que proporciona acceso a catálogos maestros:
// actividades económicas, documentos fiscales, monedas, tipos de cambio, etc.
// Estos catálogos son esenciales para validar datos antes de emitir invoices.
func (s *SiatServices) Sincronizacion() ports.SiatSincronizacionService {
	return s.sincronizacion
}

// Codigos retorna el servicio para:
// - Solicitud de códigos CUIS (Código Único de Identificación de Sistemas)
// - Solicitud de códigos CUFD (Código Único de Facturación por Dirección)
// - Validación de números NIT (Rol Tributario)
// Los códigos CUIS y CUFD son obligatorios para emitir invoices.
func (s *SiatServices) Codigos() ports.SiatCodigosService {
	return s.codigos
}

// CompraVenta retorna el servicio para el sector de compra-venta (Sector 1).
// Permite enviar, recibir y anular facturas comerciales estándar.
// Este es el sector más común para comercios generales.
func (s *SiatServices) CompraVenta() ports.SiatCompraVentaService {
	return s.compraVenta
}

// Computarizada retorna el servicio para facturación computarizada
// (sin firma digital, basada en máquinas registradoras fiscales).
// Permite enviar, recibir y anular facturas de este tipo.
func (s *SiatServices) Computarizada() ports.SiatComputarizadaService {
	return s.computarizada
}

// Electronica retorna el servicio para facturación electrónica (con firma digital).
// Permite enviar, recibir y anular facturas electrónicas de todos los sectores.
// Este es el tipo de facturación más moderno y flexible del SIAT.
func (s *SiatServices) Electronica() ports.SiatElectronicaService {
	return s.electronica
}

// DocumentoAjuste retorna el servicio para el sector de documento de ajuste.
// Permite enviar, recibir y anular facturas de este tipo.
func (s *SiatServices) DocumentoAjuste() ports.SiatDocumentoAjusteService {
	return s.documentoAjuste
}

// Telecomunicaciones retorna el servicio para el sector de telecomunicaciones.
// Permite enviar, recibir y anular facturas de este tipo.
func (s *SiatServices) Telecomunicaciones() ports.SiatTelecomunicacionesService {
	return s.telecomunicaciones
}

// ServicioBasico retorna el servicio para el sector para el sector de servicios básicos.
// Permite enviar, recibir y anular facturas de este tipo.
func (s *SiatServices) ServicioBasico() ports.SiatServicioBasicoService {
	return s.servicioBasico
}

// EntidadFinanciera retorna el servicio para el sector de entidades financieras.
// Permite enviar, recibir y anular facturas de este tipo.
func (s *SiatServices) EntidadFinanciera() ports.SiatEntidadFinancieraService {
	return s.entidadFinanciera
}

// BoletoAereo retorna el servicio para el sector de boletos aéreos.
// Permite enviar masivamente, recibir y anular facturas de este tipo.
func (s *SiatServices) BoletoAereo() ports.SiatBoletoAereoService {
	return s.boletoAereo
}

// RecepcionCompras retorna el servicio para la recepción de compras.
// Permite enviar, consultar y anular registros de compras ante el SIAT.
func (s *SiatServices) RecepcionCompras() ports.SiatRecepcionComprasService {
	return s.recepcionCompras
}

// WithConfig retorna una nueva instancia de ports.Config con el traceID actual.
// Esto permite que el usuario establezca el traceID una sola vez con WithTraceID()
// y que automáticamente se inyecte en todas las solicitudes posteriores.
//
// El usuario debe proporcionar el Token de autenticación del SIAT.
// El UserAgent se establece en una cadena vacía y puede ser personalizado.
//
// Parámetros:
//   - token: El token de autenticación del SIAT (obligatorio)
//
// Retorna:
//   - ports.Config con TraceID pre-establecido
//
// Ejemplo:
//
//	s.WithTraceID("trace-12345")
//	config := s.WithConfig("myToken123")
//	// config.TraceID es "trace-12345" automáticamente
func (s *SiatServices) WithConfig(token string) ports.Config {
	return ports.Config{
		Token:   token,
		TraceId: s.traceID,
	}
}

// New crea e inicializa una nueva instancia de SiatServices.
//
// Parámetros:
//   - baseUrl: URL base de los servicios SIAT (ej: https://pilotosiatservicios.impuestos.gob.bo/v2)
//   - httpClient: Cliente HTTP personalizado (opcional). Si es nil, se crea uno con configuración segura.
//
// La función configura automáticamente:
//   - Timeouts apropiados (15s handshake, 45s total)
//   - TLS 1.2+ para seguridad
//   - Pools de conexión para alto rendimiento
//   - Proxy desde variables de entorno si están configuradas
//
// Retorna un error si baseUrl está vacía o si alguno de los servicios falla al inicializarse.
//
// Ejemplo:
//
//	s, err := siat.New("https://pilotosiatservicios.impuestos.gob.bo/v2", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
func New(baseUrl string, httpClient *http.Client) (*SiatServices, error) {
	if httpClient != nil {
		clonedClient := *httpClient
		httpClient = &clonedClient
	} else {
		// Usar HTTPConfig para crear cliente optimizado por defecto
		httpClient = services.NewHTTPClient(services.DefaultHTTPConfig())
	}

	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	operaciones, err := services.NewSiatOperacionesService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	sincronizacion, err := services.NewSiatSincronizacionService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	codigos, err := services.NewSiatCodigosService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	compraVenta, err := services.NewSiatCompraVentaService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	computarizada, err := services.NewSiatComputarizadaService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	electronica, err := services.NewSiatElectronicaService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	documentoAjuste, err := services.NewSiatDocumentoAjusteService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	telecomunicaciones, err := services.NewSiatTelecomunicacionesService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	servicioBasico, err := services.NewSiatServicioBasicoService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	entidadFinanciera, err := services.NewSiatEntidadFinancieraService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	boletoAereo, err := services.NewSiatBoletoAereoService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	recepcionCompras, err := services.NewSiatRecepcionComprasService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	return &SiatServices{
		operaciones:        operaciones,
		sincronizacion:     sincronizacion,
		codigos:            codigos,
		compraVenta:        compraVenta,
		computarizada:      computarizada,
		electronica:        electronica,
		documentoAjuste:    documentoAjuste,
		telecomunicaciones: telecomunicaciones,
		servicioBasico:     servicioBasico,
		entidadFinanciera:  entidadFinanciera,
		boletoAereo:        boletoAereo,
		recepcionCompras:   recepcionCompras,
		traceID:            "", // Inicialmente vacío
	}, nil
}

// WithTraceID establece un ID de seguimiento (trace ID) para correlacionar solicitudes
// en sistemas distribuidos. El trace ID se inyecta en el encabezado HTTP "X-Trace-ID"
// en todas las solicitudes posteriores.
//
// El trace ID es completamente opcional y no afecta la funcionalidad del SDK si no se proporciona.
// Es útil para:
//   - Correlacionar logs entre múltiples servicios
//   - Rastrear solicitudes a través de sistemas distribuidos
//   - Debugging de problemas en producción
//
// Parámetros:
//   - id: El ID de seguimiento (puede estar vacío para limpiar el trace ID anterior)
//
// Retorna el receptor (SiatServices) para permitir encadenamiento de métodos.
//
// Ejemplo:
//
//	s.WithTraceID("trace-12345-dev").Operaciones()...
//	// El trace-12345-dev se incluirá en el encabezado X-Trace-ID de todas las solicitudes
func (s *SiatServices) WithTraceID(id string) *SiatServices {
	s.traceID = id
	return s
}

// Verify analiza una respuesta del SIAT y determina si la operación fue exitosa.
// Si la respuesta contiene errores del SIAT (Transaccion=false o mensajes de error),
// construye y retorna un *SiatError detallado.
//
// Es compatible con cualquier objeto de respuesta del SDK (RespuestaCuis, RespuestaRecepcion, etc.)
// utilizando reflexión para extraer los campos 'Transaccion' y 'MensajesList'.
//
// Ejemplo:
//
//	resp, err := s.Codigos().VerificarNit(ctx, cfg, req)
//	if err != nil { return err } // Error de red
//	if err := siat.Verify(resp.Body.Content.RespuestaCuis); err != nil {
//	    return err // Error del SIAT (ej: NIT inválido)
//	}
func Verify(resp interface{}) error {
	if resp == nil {
		return nil
	}

	// Intentar usar la interfaz común primero (más eficiente y seguro)
	if res, ok := resp.(common.Result); ok {
		return checkResult(res.IsSuccess(), res.GetMessages())
	}

	// Fallback por reflexión para objetos que aún no implementan la interfaz
	val := reflect.ValueOf(resp)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil
	}

	// Extraer campo Transaccion (bool)
	transaccionField := val.FieldByName("Transaccion")
	success := true
	if transaccionField.IsValid() && transaccionField.Kind() == reflect.Bool {
		success = transaccionField.Bool()
	}

	// Extraer campo MensajesList (slice)
	mensajesField := val.FieldByName("MensajesList")
	var messages []common.MensajeServicio

	if mensajesField.IsValid() && mensajesField.Kind() == reflect.Slice {
		for i := 0; i < mensajesField.Len(); i++ {
			msgVal := mensajesField.Index(i)
			codeField := msgVal.FieldByName("Codigo")
			descField := msgVal.FieldByName("Descripcion")

			if codeField.IsValid() && descField.IsValid() {
				messages = append(messages, common.MensajeServicio{
					Codigo:      int(codeField.Int()),
					Descripcion: descField.String(),
				})
			}
		}
	}

	return checkResult(success, messages)
}

// checkResult es un helper interno para validar el éxito y categorizar mensajes.
func checkResult(success bool, mensajes []common.MensajeServicio) error {
	var messagesStr []string
	var firstErrorCode int
	hasErrors := false

	for _, m := range mensajes {
		// Categorizar: solo fallar si no es un warning
		if !IsWarningCode(m.Codigo) {
			hasErrors = true
			if firstErrorCode == 0 {
				firstErrorCode = m.Codigo
			}
		}
		messagesStr = append(messagesStr, fmt.Sprintf("[%d] %s", m.Codigo, m.Descripcion))
	}

	// Si Transaccion es false o hay mensajes que no son warnings, es un error
	if !success || hasErrors {
		fullMsg := strings.Join(messagesStr, "; ")
		if fullMsg == "" {
			fullMsg = "Operación rechazada por el SIAT sin mensajes específicos"
		}

		err := NewSiatError(firstErrorCode, fullMsg)
		// Enriquecer con metadatos de categoría
		err.IsRetryable = IsRetryableCode(firstErrorCode)
		return err
	}

	return nil
}
