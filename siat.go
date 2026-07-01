package siat

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ron86i/go-siat/internal/adapter/services"
	"github.com/ron86i/go-siat/internal/core/domain/siat/common"
	"github.com/ron86i/go-siat/internal/core/ports"
)

// SiatServices es el punto de entrada principal del SDK.
// Agrupa todas las implementaciones de los servicios del SIAT
// (Códigos, Sincronización, Operaciones, Compra-Venta, Computarizada, Electrónica, etc.).
// Los usuarios deben crear una instancia usando New().
type SiatServices struct {
	config             Config
	operaciones        ports.SiatOperacionesService
	sincronizacion     ports.SiatSincronizacionService
	codigos            ports.SiatCodigosService
	compraVenta        ports.SiatCompraVentaService
	computarizada      ports.SiatSuministroEnergiaService
	electronica        ports.SiatSuministroEnergiaService
	documentoAjuste    ports.SiatDocumentoAjusteService
	telecomunicaciones ports.FacturacionService
	servicioBasico     ports.FacturacionService
	entidadFinanciera  ports.FacturacionService
	boletoAereo        ports.SiatBoletoAereoService
	recepcionCompras   ports.SiatRecepcionComprasService
}

// Config retorna la configuración actual del cliente.
func (s *SiatServices) Config() Config {
	return s.config
}

// Operaciones retorna el servicio para la gestión de puntos de venta (PV),
// cierre de períodos de facturación y eventos significativos (cambios de modalidad, etc.).
func (s *SiatServices) Operaciones() ports.SiatOperacionesService {
	return s.operaciones
}

// Sincronizacion retorna el servicio que proporciona acceso a catálogos maestros.
func (s *SiatServices) Sincronizacion() ports.SiatSincronizacionService {
	return s.sincronizacion
}

// Codigos retorna el servicio para solicitud de códigos (CUIS, CUFD) y validación de NIT.
func (s *SiatServices) Codigos() ports.SiatCodigosService {
	return s.codigos
}

// CompraVenta retorna el servicio para el sector de compra-venta (Sector 1).
func (s *SiatServices) CompraVenta() ports.SiatCompraVentaService {
	return s.compraVenta
}

// Computarizada retorna el servicio para facturación computarizada.
func (s *SiatServices) Computarizada() ports.SiatSuministroEnergiaService {
	return s.computarizada
}

// Electronica retorna el servicio para facturación electrónica (con firma digital).
func (s *SiatServices) Electronica() ports.SiatSuministroEnergiaService {
	return s.electronica
}

// DocumentoAjuste retorna el servicio para el sector de documento de ajuste.
func (s *SiatServices) DocumentoAjuste() ports.SiatDocumentoAjusteService {
	return s.documentoAjuste
}

// Telecomunicaciones retorna el servicio para el sector de telecomunicaciones.
func (s *SiatServices) Telecomunicaciones() ports.FacturacionService {
	return s.telecomunicaciones
}

// ServicioBasico retorna el servicio para el sector de servicios básicos.
func (s *SiatServices) ServicioBasico() ports.FacturacionService {
	return s.servicioBasico
}

// EntidadFinanciera retorna el servicio para el sector de entidades financieras.
func (s *SiatServices) EntidadFinanciera() ports.FacturacionService {
	return s.entidadFinanciera
}

// BoletoAereo retorna el servicio para el sector de boletos aéreos.
func (s *SiatServices) BoletoAereo() ports.SiatBoletoAereoService {
	return s.boletoAereo
}

// RecepcionCompras retorna el servicio para la recepción de compras.
func (s *SiatServices) RecepcionCompras() ports.SiatRecepcionComprasService {
	return s.recepcionCompras
}

// New crea e inicializa una nueva instancia de SiatServices usando la configuración global.
func New(config Config) (*SiatServices, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	httpClient := config.HTTPClient
	if httpClient != nil {
		clonedClient := *httpClient
		httpClient = &clonedClient
	} else {
		httpClient = services.NewHTTPClient(services.DefaultHTTPConfig())
	}

	baseUrl := strings.TrimSpace(config.BaseURL)

	operaciones, err := services.NewSiatOperacionesService(baseUrl, httpClient, config)
	if err != nil {
		return nil, err
	}
	sincronizacion, err := services.NewSiatSincronizacionService(baseUrl, httpClient, config)
	if err != nil {
		return nil, err
	}
	codigos, err := services.NewSiatCodigosService(baseUrl, httpClient, config)
	if err != nil {
		return nil, err
	}
	compraVenta, err := services.NewCompraVentaService(baseUrl, httpClient, config)
	if err != nil {
		return nil, err
	}
	computarizada, err := services.NewSuministroEnergiaService(baseUrl, httpClient, config, services.SiatComputarizada)
	if err != nil {
		return nil, err
	}
	electronica, err := services.NewSuministroEnergiaService(baseUrl, httpClient, config, services.SiatElectronica)
	if err != nil {
		return nil, err
	}
	documentoAjuste, err := services.NewSiatDocumentoAjusteService(baseUrl, httpClient, config)
	if err != nil {
		return nil, err
	}
	telecomunicaciones, err := services.NewFacturacionService(baseUrl, httpClient, config, services.SiatTelecomunicaciones)
	if err != nil {
		return nil, err
	}
	servicioBasico, err := services.NewFacturacionService(baseUrl, httpClient, config, services.SiatServicioBasico)
	if err != nil {
		return nil, err
	}
	entidadFinanciera, err := services.NewFacturacionService(baseUrl, httpClient, config, services.SiatEntidadFinanciera)
	if err != nil {
		return nil, err
	}
	boletoAereo, err := services.NewSiatBoletoAereoService(baseUrl, httpClient, config)
	if err != nil {
		return nil, err
	}
	recepcionCompras, err := services.NewSiatRecepcionComprasService(baseUrl, httpClient, config)
	if err != nil {
		return nil, err
	}

	return &SiatServices{
		config:             config,
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
	}, nil
}

// Verify analiza una respuesta del SIAT y determina si la operación fue exitosa.
// Si la respuesta contiene errores del SIAT (Transaccion=false o mensajes de error),
// construye y retorna un *SiatError detallado.
func Verify(resp interface{}) error {
	if resp == nil {
		return nil
	}

	// Usar la interfaz común
	if res, ok := resp.(common.Result); ok {
		return checkResult(res.IsSuccess(), res.GetMessages())
	}

	return nil
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
		err.Mensajes = mensajes
		// Enriquecer con metadatos de categoría
		err.IsRetryable = IsRetryableCode(firstErrorCode)
		return err
	}

	return nil
}

// Map is a shortcut for map[string]interface{}, useful for JSON returns
type Map map[string]interface{}

func (m Map) ToJSON() (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

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

func (m Map) ToStruct(v interface{}) error {
	bytes, err := m.ToJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(bytes), v)
}
