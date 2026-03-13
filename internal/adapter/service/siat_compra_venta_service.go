package service

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/compra_venta"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatCompraVentaService implementa el puerto port.SiatCompraVentaService para interactuar
// con los servicios de recepción y anulación de facturas del SIAT.
type SiatCompraVentaService struct {
	url        string
	HttpClient *http.Client
}

// RecepcionAnexos
func (s *SiatCompraVentaService) RecepcionAnexos(ctx context.Context, config config.Config, opaqueReq models.RecepcionAnexos) (*soap.EnvelopeResponse[compra_venta.RecepcionAnexosResponse], error) {
	req := models.GetInternalRequest[compra_venta.RecepcionAnexos](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.RecepcionAnexosResponse](resp)
}

// ValidacionRecepcionMasivaFactura
func (s *SiatCompraVentaService) ValidacionRecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[compra_venta.ValidacionRecepcionMasivaFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.ValidacionRecepcionMasivaFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.ValidacionRecepcionMasivaFacturaResponse](resp)
}

// VerificacionEstadoFactura
func (s *SiatCompraVentaService) VerificacionEstadoFactura(ctx context.Context, config config.Config, opaqueReq models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[compra_venta.VerificacionEstadoFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.VerificacionEstadoFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.VerificacionEstadoFacturaResponse](resp)
}

// RecepcionMasivaFactura permite recibir paquetes de facturas emitidas bajo la modalidad
// de Facturación Electrónica en Línea de forma masiva (hasta 1000 facturas por paquete).
// La periodicidad del envío (diario, semanal o mensual) se configura en el portal de la Administración Tributaria.
// Retorna un código de recepción si es aceptado, o códigos de error/advertencia.
// Recibe una solicitud opaca de tipo RecepcionMasivaFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) RecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionMasivaFactura) (*soap.EnvelopeResponse[compra_venta.RecepcionMasivaFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.RecepcionMasivaFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.RecepcionMasivaFacturaResponse](resp)
}

// VerificarComunicacion permite verificar la comunicación con el SIAT.
// Recibe una solicitud opaca de tipo VerificarComunicacionRequest construida vía Builder.
func (s *SiatCompraVentaService) VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionCompraVenta) (*soap.EnvelopeResponse[compra_venta.VerificarComunicacionResponse], error) {
	req := models.GetInternalRequest[compra_venta.VerificarComunicacion](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.VerificarComunicacionResponse](resp)
}

// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas.
// Recibe una solicitud opaca de tipo ValidacionRecepcionPaqueteFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) ValidacionRecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionPaqueteFactura) (*soap.EnvelopeResponse[compra_venta.ValidacionRecepcionPaqueteFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.ValidacionRecepcionPaqueteFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.ValidacionRecepcionPaqueteFacturaResponse](resp)
}

// RecepcionPaqueteFactura permite recibir paquetes de hasta 500 facturas emitidas bajo la modalidad
// de Facturación Electrónica en Línea. El servicio verifica la validez de los parámetros
// y la integridad del paquete, retornando un código de recepción si es aceptado,
// o códigos de error/advertencia en caso contrario.
// Recibe una solicitud opaca de tipo RecepcionPaqueteFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) RecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionPaqueteFactura) (*soap.EnvelopeResponse[compra_venta.RecepcionPaqueteFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.RecepcionPaqueteFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.RecepcionPaqueteFacturaResponse](resp)
}

// ReversionAnulacionFactura permite revertir la anulación de una factura previamente enviada al SIAT.
// Permite revertir el estado de las facturas digitales que fueron anuladas por error una sola vez.
// Recibe una solicitud opaca de tipo ReversionAnulacionFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) ReversionAnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.ReversionAnulacionFactura) (*soap.EnvelopeResponse[compra_venta.ReversionAnulacionFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.ReversionAnulacionFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.ReversionAnulacionFacturaResponse](resp)
}

// AnulacionFactura permite anular una factura previamente aceptada por el SIAT.
// Recibe una solicitud opaca de tipo AnulacionFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) AnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.AnulacionFactura) (*soap.EnvelopeResponse[compra_venta.AnulacionFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.AnulacionFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.AnulacionFacturaResponse](resp)
}

// RecepcionFactura envía una factura firmada, comprimida y codificada al SIAT para su procesamiento.
// Recibe una solicitud opaca de tipo RecepcionFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) RecepcionFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionFactura) (*soap.EnvelopeResponse[compra_venta.RecepcionFacturaResponse], error) {
	req := models.GetInternalRequest[compra_venta.RecepcionFactura](opaqueReq)
	xmlBody, err := buildRequest(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Request XML: %s", string(xmlBody))

	httpReq, err := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(xmlBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/xml")
	httpReq.Header.Set("apiKey", fmt.Sprintf("TokenApi %s", config.Token))

	resp, err := s.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return parseSoapResponse[compra_venta.RecepcionFacturaResponse](resp)
}

// NewSiatCompraVentaService crea una nueva instancia del servicio de Compra y Venta.
// Inicializa la URL específica del servicio basándose en la baseUrl proporcionada.
func NewSiatCompraVentaService(baseUrl string, httpClient *http.Client) (*SiatCompraVentaService, error) {
	baseUrl = strings.TrimSpace(baseUrl)
	if baseUrl == "" {
		return nil, fmt.Errorf("baseUrl is empty")
	}

	// Si no se inyecta un cliente, creamos uno con configuraciones seguras por defecto
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}
	return &SiatCompraVentaService{
		url:        fullURL(baseUrl, SiatCompraVenta),
		HttpClient: httpClient,
	}, nil
}

var _ port.SiatCompraVentaService = (*SiatCompraVentaService)(nil)
