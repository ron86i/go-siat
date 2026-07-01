package services

import (
	"context"
	"net/http"

	"github.com/ron86i/go-siat/v2/internal/core/domain/datatype/soap"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/compra_venta"
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/facturacion"
	"github.com/ron86i/go-siat/v2/internal/core/ports"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

type FacturacionServiceImpl struct {
	url        string
	httpClient *http.Client
	config     ports.Config
}

func NewFacturacionService(baseUrl string, httpClient *http.Client, config ports.Config, serviceName SiatService) (*FacturacionServiceImpl, error) {
	return &FacturacionServiceImpl{
		url:        fullURL(baseUrl, serviceName),
		httpClient: httpClient,
		config:     config,
	}, nil
}

func (s *FacturacionServiceImpl) AnulacionFactura(ctx context.Context, req models.AnulacionFactura) (*soap.EnvelopeResponse[facturacion.AnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.AnulacionFactura, facturacion.AnulacionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) RecepcionFactura(ctx context.Context, req models.RecepcionFactura) (*soap.EnvelopeResponse[facturacion.RecepcionFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionFactura, facturacion.RecepcionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) ReversionAnulacionFactura(ctx context.Context, req models.ReversionAnulacionFactura) (*soap.EnvelopeResponse[facturacion.ReversionAnulacionFacturaResponse], error) {
	return performSoapRequest[facturacion.ReversionAnulacionFactura, facturacion.ReversionAnulacionFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) RecepcionPaqueteFactura(ctx context.Context, req models.RecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.RecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionPaqueteFactura, facturacion.RecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) ValidacionRecepcionPaqueteFactura(ctx context.Context, req models.ValidacionRecepcionPaqueteFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionPaqueteFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionPaqueteFactura, facturacion.ValidacionRecepcionPaqueteFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) VerificarComunicacion(ctx context.Context, req models.VerificarComunicacionFacturacion) (*soap.EnvelopeResponse[facturacion.VerificarComunicacionResponse], error) {
	return performSoapRequest[facturacion.VerificarComunicacion, facturacion.VerificarComunicacionResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) RecepcionMasivaFactura(ctx context.Context, req models.RecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.RecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.RecepcionMasivaFactura, facturacion.RecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) ValidacionRecepcionMasivaFactura(ctx context.Context, req models.ValidacionRecepcionMasivaFactura) (*soap.EnvelopeResponse[facturacion.ValidacionRecepcionMasivaFacturaResponse], error) {
	return performSoapRequest[facturacion.ValidacionRecepcionMasivaFactura, facturacion.ValidacionRecepcionMasivaFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
}

func (s *FacturacionServiceImpl) VerificacionEstadoFactura(ctx context.Context, req models.VerificacionEstadoFactura) (*soap.EnvelopeResponse[facturacion.VerificacionEstadoFacturaResponse], error) {
	return performSoapRequest[facturacion.VerificacionEstadoFactura, facturacion.VerificacionEstadoFacturaResponse](ctx, s.httpClient, s.url, s.config, req)
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
