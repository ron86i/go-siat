package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/internal/core/domain/siat/compra_venta"
	"github.com/ron86i/go-siat/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/internal/core/port"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/models"
)

// SiatCompraVentaService implementa el puerto port.SiatCompraVentaService para interactuar
// con los servicios de recepción y anulación de facturas del SIAT.
type SiatCompraVentaService struct {
	url        string
	httpClient *http.Client
}

// RecepcionAnexos
func (s *SiatCompraVentaService) RecepcionAnexos(ctx context.Context, config config.Config, opaqueReq models.RecepcionAnexosCompraVenta) (*soap.EnvelopeResponse[compra_venta.RecepcionAnexosResponse], error) {
	return performSoapRequest[compra_venta.RecepcionAnexos, compra_venta.RecepcionAnexosResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// ValidacionRecepcionMasivaFactura
func (s *SiatCompraVentaService) ValidacionRecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionMasivaFacturaCompraVenta) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// VerificacionEstadoFactura
func (s *SiatCompraVentaService) VerificacionEstadoFactura(ctx context.Context, config config.Config, opaqueReq models.VerificacionEstadoFacturaCompraVenta) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RecepcionMasivaFactura permite recibir paquetes de facturas emitidas bajo la modalidad
// de Facturación Electrónica en Línea de forma masiva (hasta 1000 facturas por paquete).
// La periodicidad del envío (diario, semanal o mensual) se configura en el portal de la Administración Tributaria.
// Retorna un código de recepción si es aceptado, o códigos de error/advertencia.
// Recibe una solicitud opaca de tipo RecepcionMasivaFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) RecepcionMasivaFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionMasivaFacturaCompraVenta) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// VerificarComunicacion permite verificar la comunicación con el SIAT.
// Recibe una solicitud opaca de tipo VerificarComunicacionRequest construida vía Builder.
func (s *SiatCompraVentaService) VerificarComunicacion(ctx context.Context, config config.Config, opaqueReq models.VerificarComunicacionCompraVenta) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// ValidacionRecepcionPaqueteFactura permite validar la recepción de paquetes de facturas.
// Recibe una solicitud opaca de tipo ValidacionRecepcionPaqueteFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) ValidacionRecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.ValidacionRecepcionPaqueteFacturaCompraVenta) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RecepcionPaqueteFactura permite recibir paquetes de hasta 500 facturas emitidas bajo la modalidad
// de Facturación Electrónica en Línea. El servicio verifica la validez de los parámetros
// y la integridad del paquete, retornando un código de recepción si es aceptado,
// o códigos de error/advertencia en caso contrario.
// Recibe una solicitud opaca de tipo RecepcionPaqueteFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) RecepcionPaqueteFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionPaqueteFacturaCompraVenta) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// ReversionAnulacionFactura permite revertir la anulación de una factura previamente enviada al SIAT.
// Permite revertir el estado de las facturas digitales que fueron anuladas por error una sola vez.
// Recibe una solicitud opaca de tipo ReversionAnulacionFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) ReversionAnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.ReversionAnulacionFacturaCompraVenta) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// AnulacionFactura permite anular una factura previamente aceptada por el SIAT.
// Recibe una solicitud opaca de tipo AnulacionFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) AnulacionFactura(ctx context.Context, config config.Config, opaqueReq models.AnulacionFacturaCompraVenta) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
}

// RecepcionFactura envía una factura firmada, comprimida y codificada al SIAT para su procesamiento.
// Recibe una solicitud opaca de tipo RecepcionFacturaRequest construida vía Builder.
func (s *SiatCompraVentaService) RecepcionFactura(ctx context.Context, config config.Config, opaqueReq models.RecepcionFactura) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, config.Token, opaqueReq)
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
		httpClient: httpClient,
	}, nil
}

var _ port.SiatCompraVentaService = (*SiatCompraVentaService)(nil)
