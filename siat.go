package siat

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/adapter/service"
)

// siatServices agrupa todas las implementaciones de los servicios del SIAT
// accesibles a través de un único punto de entrada orientado a métodos.
type siatServices struct {
	operaciones    *service.SiatOperacionesService
	sincronizacion *service.SiatSincronizacionService
	codigos        *service.SiatCodigosService
	compraVenta    *service.SiatCompraVentaService
}

// Operaciones retorna el servicio para la gestión de puntos de venta y eventos significativos.
func (s *siatServices) Operaciones() *service.SiatOperacionesService {
	return s.operaciones
}

// Sincronizacion retorna el servicio para la obtención de catálogos y parametrizaciones del SIAT.
func (s *siatServices) Sincronizacion() *service.SiatSincronizacionService {
	return s.sincronizacion
}

// Codigos retorna el servicio para la solicitud de códigos CUIS y CUFD, y validación de NIT.
func (s *siatServices) Codigos() *service.SiatCodigosService {
	return s.codigos
}

// CompraVenta retorna el servicio para el envío y anulación de facturas comerciales.
func (s *siatServices) CompraVenta() *service.SiatCompraVentaService {
	return s.compraVenta
}

// New crea e inicializa una nueva instancia unificada de los servicios del SIAT.
// Requiere la URL base del servicio (Pruebas o Producción) y un cliente HTTP opcional.
// Si httpClient es nil, se utilizará uno por defecto con un tiempo de espera (timeout) de 15 segundos.
func New(baseUrl string, httpClient *http.Client) (*siatServices, error) {
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
	return &siatServices{
		operaciones:    operaciones,
		sincronizacion: sincronizacion,
		codigos:        codigos,
		compraVenta:    compraVenta,
	}, nil
}
