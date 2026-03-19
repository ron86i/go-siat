package siat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/adapter/service"
	"github.com/ron86i/go-siat/internal/core/port"
)

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

// SiatServices agrupa todas las implementaciones de los servicios del SIAT
// accesibles a través de un único punto de entrada orientado a métodos.
type SiatServices struct {
	operaciones    port.SiatOperacionesPort
	sincronizacion port.SiatSincronizacionService
	codigos        port.SiatCodigosService
	compraVenta    port.SiatCompraVentaService
	computarizada  port.SiatComputarizadaService
}

// Operaciones retorna el servicio para la gestión de puntos de venta y eventos significativos.
func (s *SiatServices) Operaciones() port.SiatOperacionesPort {
	return s.operaciones
}

// Sincronizacion retorna el servicio para la obtención de catálogos y parametrizaciones del SIAT.
func (s *SiatServices) Sincronizacion() port.SiatSincronizacionService {
	return s.sincronizacion
}

// Codigos retorna el servicio para la solicitud de códigos CUIS y CUFD, y validación de NIT.
func (s *SiatServices) Codigos() port.SiatCodigosService {
	return s.codigos
}

// CompraVenta retorna el servicio para el envío y anulación de facturas comerciales.
func (s *SiatServices) CompraVenta() port.SiatCompraVentaService {
	return s.compraVenta
}

// Computarizada retorna el servicio para el envío y anulación de facturas comerciales.
func (s *SiatServices) Computarizada() port.SiatComputarizadaService {
	return s.computarizada
}

// New crea e inicializa una nueva instancia unificada de los servicios del SIAT.
// Requiere la URL base del servicio (Pruebas o Producción) y un cliente HTTP opcional.
// Si httpClient es nil, se utilizará uno por defecto con un tiempo de espera (timeout) de 15 segundos.
func New(baseUrl string, httpClient *http.Client) (*SiatServices, error) {
	if httpClient != nil {
		clonedClient := *httpClient
		httpClient = &clonedClient
	} else {
		httpClient = &http.Client{Timeout: 15 * time.Second}
	}

	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	operaciones, err := service.NewSiatOperacionesService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	sincronizacion, err := service.NewSiatSincronizacionService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	codigos, err := service.NewSiatCodigosService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	compraVenta, err := service.NewSiatCompraVentaService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}

	computarizada, err := service.NewSiatComputarizadaService(baseUrl, httpClient)
	if err != nil {
		return nil, err
	}
	return &SiatServices{
		operaciones:    operaciones,
		sincronizacion: sincronizacion,
		codigos:        codigos,
		compraVenta:    compraVenta,
		computarizada:  computarizada,
	}, nil
}
