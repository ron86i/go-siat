package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/compra_venta"
	"github.com/ron86i/go-siat/v3/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v3/internal/core/ports"
	"github.com/ron86i/go-siat/v3/internal/siatpolicy"
	"github.com/ron86i/go-siat/v3/pkg/models"
)

type FacturacionServiceImpl struct {
	url        string
	service    SiatService
	httpClient *http.Client
	config     ports.Config
}

func NewFacturacionService(baseUrl string, httpClient *http.Client, config ports.Config, serviceName SiatService) (*FacturacionServiceImpl, error) {
	if !siatpolicy.ValidService(string(serviceName)) {
		return nil, fmt.Errorf("servicio SIAT desconocido: %q", serviceName)
	}
	return &FacturacionServiceImpl{
		url:        fullURL(baseUrl, serviceName),
		service:    serviceName,
		httpClient: httpClient,
		config:     config,
	}, nil
}

func (s *FacturacionServiceImpl) AnulacionFactura(ctx context.Context, req models.AnulacionFactura) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 0, func(r *facturacion.AnulacionFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudAnulacion.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) RecepcionFactura(ctx context.Context, req models.RecepcionFactura) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 1, func(r *facturacion.RecepcionFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioRecepcionFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) ReversionAnulacionFactura(ctx context.Context, req models.ReversionAnulacionFactura) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 0, func(r *facturacion.ReversionAnulacionFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudReversionAnulacion.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) RecepcionPaqueteFactura(ctx context.Context, req models.RecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 2, func(r *facturacion.RecepcionPaqueteFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioRecepcionPaquete.SolicitudRecepcionFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	paquete := getInternalRequest[facturacion.RecepcionPaqueteFactura](req).SolicitudServicioRecepcionPaquete
	if paquete.CantidadFacturas <= 0 || paquete.CantidadFacturas > siatpolicy.MaxFacturasPaquete || paquete.CodigoEvento <= 0 {
		return nil, fmt.Errorf("solicitud de paquete incompleta o fuera de límites")
	}
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) ValidacionRecepcionPaqueteFactura(ctx context.Context, req models.ValidacionRecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 2, func(r *facturacion.ValidacionRecepcionPaqueteFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioValidacionRecepcionPaquete.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	if getInternalRequest[facturacion.ValidacionRecepcionPaqueteFactura](req).SolicitudServicioValidacionRecepcionPaquete.CodigoRecepcion == "" {
		return nil, fmt.Errorf("codigo de recepción de paquete obligatorio")
	}
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionFacturacion) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) RecepcionMasivaFactura(ctx context.Context, req models.RecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 3, func(r *facturacion.RecepcionMasivaFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioRecepcionMasiva.SolicitudRecepcionFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	if count := getInternalRequest[facturacion.RecepcionMasivaFactura](req).SolicitudServicioRecepcionMasiva.CantidadFacturas; count <= 0 || count > siatpolicy.MaxFacturasMasiva {
		return nil, fmt.Errorf("cantidad de facturas masiva fuera de límites")
	}
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) ValidacionRecepcionMasivaFactura(ctx context.Context, req models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 3, func(r *facturacion.ValidacionRecepcionMasivaFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioValidacionRecepcionMasivaFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	if getInternalRequest[facturacion.ValidacionRecepcionMasivaFactura](req).SolicitudServicioValidacionRecepcionMasivaFactura.CodigoRecepcion == "" {
		return nil, fmt.Errorf("codigo de recepción masiva obligatorio")
	}
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) VerificacionEstadoFactura(ctx context.Context, req models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	if err := validateLegacyFacturacionRequest(s.service, req, 0, func(r *facturacion.VerificacionEstadoFactura) facturacion.SolicitudRecepcion {
		return r.SolicitudServicioVerificacionEstadoFactura.SolicitudRecepcion
	}); err != nil {
		return nil, err
	}
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func validateLegacyFacturacionRequest[T any](service SiatService, req any, emission int, reception func(*T) facturacion.SolicitudRecepcion) error {
	internal := getInternalRequest[T](req)
	if internal == nil {
		return fmt.Errorf("solicitud de facturación inválida")
	}
	payload := reception(internal)
	if err := siatpolicy.ValidateDestination(string(service), payload.CodigoDocumentoSector, payload.CodigoModalidad); err != nil {
		return err
	}
	if emission != 0 && payload.CodigoEmision != emission {
		return fmt.Errorf("codigoEmision %d incompatible con la operación (esperado %d)", payload.CodigoEmision, emission)
	}
	return nil
}

var _ ports.FacturacionService = (*FacturacionServiceImpl)(nil)

// CompraVenta
type CompraVentaServiceImpl struct {
	FacturacionServiceImpl
}

func NewCompraVentaService(baseUrl string, httpClient *http.Client, config ports.Config) (*CompraVentaServiceImpl, error) {
	fs, err := NewFacturacionService(baseUrl, httpClient, config, SiatCompraVenta)
	if err != nil {
		return nil, err
	}
	return &CompraVentaServiceImpl{
		FacturacionServiceImpl: *fs,
	}, nil
}

func (s *CompraVentaServiceImpl) RecepcionAnexos(ctx context.Context, req models.RecepcionAnexosCompraVenta) (*soap.EnvelopeResponse[compra_venta.RecepcionAnexosResponse], error) {
	return performSoapRequest[compra_venta.RecepcionAnexos, compra_venta.RecepcionAnexosResponse](ctx, s.httpClient, s.url, s.config, req)
}

var _ ports.SiatCompraVentaService = (*CompraVentaServiceImpl)(nil)

// SuministroEnergia (Computarizada y Electronica)
type SuministroEnergiaServiceImpl struct {
	FacturacionServiceImpl
}

func NewSuministroEnergiaService(baseUrl string, httpClient *http.Client, config ports.Config, serviceName SiatService) (*SuministroEnergiaServiceImpl, error) {
	fs, err := NewFacturacionService(baseUrl, httpClient, config, serviceName)
	if err != nil {
		return nil, err
	}
	return &SuministroEnergiaServiceImpl{
		FacturacionServiceImpl: *fs,
	}, nil
}

func (s *SuministroEnergiaServiceImpl) RecepcionAnexosSuministroEnergia(ctx context.Context, req models.RecepcionAnexosSuministroEnergia) (*soap.EnvelopeResponse[facturacion.RecepcionAnexosSuministroEnergiaResponse], error) {
	return performSoapRequest[facturacion.RecepcionAnexosSuministroEnergia, facturacion.RecepcionAnexosSuministroEnergiaResponse](ctx, s.httpClient, s.url, s.config, req)
}

var _ ports.SiatSuministroEnergiaService = (*SuministroEnergiaServiceImpl)(nil)
