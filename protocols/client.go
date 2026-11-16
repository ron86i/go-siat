// Package protocols agrupa los clientes de los servicios SOAP de SIAT.
package protocols

import (
	"github.com/ron86i/go-siat/v3/protocols/codigos"
	"github.com/ron86i/go-siat/v3/protocols/compras"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
	"github.com/ron86i/go-siat/v3/protocols/operaciones"
	"github.com/ron86i/go-siat/v3/protocols/sincronizacion"
)

// Config contiene la conexión e identidad utilizada por los servicios.
type Config = facturacion.Config

// ErrorComunicacion indica que no se obtuvo una respuesta SOAP utilizable.
type ErrorComunicacion = facturacion.ErrorComunicacion

// Client agrupa los servicios disponibles del protocolo.
type Client struct {
	Codigos            *codigos.Client
	Compras            *compras.Client
	Facturacion        *facturacion.Client
	Electronica        *facturacion.Client
	Computarizada      *facturacion.Client
	ServicioBasico     *facturacion.Client
	DocumentoAjuste    *facturacion.Client
	Telecomunicaciones *facturacion.Client
	EntidadFinanciera  *facturacion.Client
	BoletoAereo        *facturacion.Client
	Sincronizacion     *sincronizacion.Client
	Operaciones        *operaciones.Client
}

// NewClient crea los servicios con una configuración común.
func NewClient(config Config) (*Client, error) {
	cod, err := codigos.NewClient(codigos.Config(config))
	if err != nil {
		return nil, err
	}
	fac, err := facturacion.NewClient(config)
	if err != nil {
		return nil, err
	}
	sinc, err := sincronizacion.NewClient(sincronizacion.Config(config))
	if err != nil {
		return nil, err
	}
	ops, err := operaciones.NewClient(operaciones.Config(config))
	if err != nil {
		return nil, err
	}
	electronica, err := facturacion.NewElectronicaClient(config)
	if err != nil {
		return nil, err
	}
	computarizada, err := facturacion.NewComputarizadaClient(config)
	if err != nil {
		return nil, err
	}
	servicioBasico, err := facturacion.NewServicioBasicoClient(config)
	if err != nil {
		return nil, err
	}
	documentoAjuste, err := facturacion.NewDocumentoAjusteClient(config)
	if err != nil {
		return nil, err
	}
	telecomunicaciones, err := facturacion.NewTelecomunicacionesClient(config)
	if err != nil {
		return nil, err
	}
	entidadFinanciera, err := facturacion.NewEntidadFinancieraClient(config)
	if err != nil {
		return nil, err
	}
	boletoAereo, err := facturacion.NewBoletoAereoClient(config)
	if err != nil {
		return nil, err
	}
	comprasClient, err := compras.NewClient(compras.Config(config))
	if err != nil {
		return nil, err
	}
	return &Client{Codigos: cod, Compras: comprasClient, Facturacion: fac, Sincronizacion: sinc, Operaciones: ops, Electronica: electronica, Computarizada: computarizada, ServicioBasico: servicioBasico, DocumentoAjuste: documentoAjuste, Telecomunicaciones: telecomunicaciones, EntidadFinanciera: entidadFinanciera, BoletoAereo: boletoAereo}, nil
}
