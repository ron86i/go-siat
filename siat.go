package siat

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/adapter/service"
	"github.com/ron86i/go-siat/internal/core/port"
)

// SiatServices agrupa todas las implementaciones de los servicios del SIAT
// accesibles a través de un único punto de entrada orientado a métodos.
type SiatServices struct {
	operaciones    port.SiatOperacionesPort
	sincronizacion port.SiatSincronizacionCatalogoService
	codigos        port.SiatCodigosService
	compraVenta    port.SiatCompraVentaService
}

// Operaciones retorna el servicio para la gestión de puntos de venta y eventos significativos.
func (s *SiatServices) Operaciones() port.SiatOperacionesPort {
	return s.operaciones
}

// Sincronizacion retorna el servicio para la obtención de catálogos y parametrizaciones del SIAT.
func (s *SiatServices) Sincronizacion() port.SiatSincronizacionCatalogoService {
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

// New crea e inicializa una nueva instancia unificada de los servicios del SIAT.
// Requiere la URL base del servicio (Pruebas o Producción) y un cliente HTTP opcional.
// Si httpClient es nil, se utilizará uno por defecto con un tiempo de espera (timeout) de 15 segundos.
func New(baseUrl string, httpClient *http.Client) (*SiatServices, error) {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("la URL base del SIAT no puede estar vacía")
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
	return &SiatServices{
		operaciones:    operaciones,
		sincronizacion: sincronizacion,
		codigos:        codigos,
		compraVenta:    compraVenta,
	}, nil
}
