// Package sincronizacion ofrece las operaciones del servicio SOAP FacturacionSincronizacion del protocolo SIAT v3.
package sincronizacion

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ron86i/go-siat/v3/internal/soapclient"
)

// Config agrupa la identidad y los datos de conexión del contribuyente.
type Config struct {
	BaseURL           string
	Token             string
	Nit               int64
	CodigoSistema     string
	CodigoAmbiente    int
	HTTPClient        *http.Client
	MaxRespuestaBytes int64
}

// Fault representa un fallo SOAP devuelto por el SIN.
type Fault = soapclient.Fault

// ErrorComunicacion indica que no se obtuvo una respuesta SOAP utilizable.
type ErrorComunicacion = soapclient.CommunicationError

// Client consulta los catálogos y paramétricas del SIN mediante FacturacionSincronizacion.
type Client struct{ soap *soapclient.Client }

// NewClient crea un cliente para el servicio FacturacionSincronizacion.
func NewClient(c Config) (*Client, error) {
	s, err := soapclient.New(soapclient.Config(c), "FacturacionSincronizacion")
	if err != nil {
		return nil, err
	}
	return &Client{soap: s}, nil
}

type respuestaXML[T any] struct {
	XMLName   xml.Name
	Resultado *T `xml:",any"`
}

func query[T any](ctx context.Context, c *Client, operacion string, r SolicitudSincronizacion) (*T, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	if r.CodigoSucursal < 0 || (r.CodigoPuntoVenta != nil && *r.CodigoPuntoVenta < 0) || r.Cuis == "" {
		return nil, fmt.Errorf("solicitud de sincronización incompleta o inválida: sucursal >= 0 y CUIS son obligatorios")
	}
	cfg := c.soap.Config()
	payload := struct {
		XMLName   xml.Name
		Solicitud solicitudXML
	}{
		XMLName: xml.Name{Local: "ns:" + operacion},
		Solicitud: solicitudXML{
			CodigoAmbiente:   cfg.CodigoAmbiente,
			CodigoPuntoVenta: r.CodigoPuntoVenta,
			CodigoSistema:    cfg.CodigoSistema,
			CodigoSucursal:   r.CodigoSucursal,
			Cuis:             r.Cuis,
			Nit:              cfg.Nit,
		},
	}
	var response respuestaXML[T]
	if err := c.soap.Call(ctx, payload, &response); err != nil {
		return nil, err
	}
	if response.XMLName.Local != operacion+"Response" || response.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de %s", operacion)
	}
	return response.Resultado, nil
}

// SincronizarActividades consulta el catálogo de actividades económicas según CAEB.
func (c *Client) SincronizarActividades(ctx context.Context, r SolicitudSincronizacion) (*ResultadoActividades, error) {
	return query[ResultadoActividades](ctx, c, "sincronizarActividades", r)
}

// SincronizarFechaHora consulta la fecha y hora oficial del servidor del SIN.
func (c *Client) SincronizarFechaHora(ctx context.Context, r SolicitudSincronizacion) (*ResultadoFechaHora, error) {
	return query[ResultadoFechaHora](ctx, c, "sincronizarFechaHora", r)
}

// SincronizarListaActividadesDocumentoSector consulta la relación entre actividades y tipos de documento sector.
func (c *Client) SincronizarListaActividadesDocumentoSector(ctx context.Context, r SolicitudSincronizacion) (*ResultadoActividadesSector, error) {
	return query[ResultadoActividadesSector](ctx, c, "sincronizarListaActividadesDocumentoSector", r)
}

// SincronizarListaLeyendasFactura consulta el catálogo de leyendas legales para facturas.
func (c *Client) SincronizarListaLeyendasFactura(ctx context.Context, r SolicitudSincronizacion) (*ResultadoLeyendas, error) {
	return query[ResultadoLeyendas](ctx, c, "sincronizarListaLeyendasFactura", r)
}

// SincronizarListaProductosServicios consulta el catálogo de productos y servicios homologados.
func (c *Client) SincronizarListaProductosServicios(ctx context.Context, r SolicitudSincronizacion) (*ResultadoProductos, error) {
	return query[ResultadoProductos](ctx, c, "sincronizarListaProductosServicios", r)
}

// SincronizarListaMensajesServicios consulta los mensajes de error e informativos del SIN.
func (c *Client) SincronizarListaMensajesServicios(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarListaMensajesServicios", r)
}

// SincronizarParametricaEventosSignificativos consulta el catálogo de motivos de contingencia o eventos significativos.
func (c *Client) SincronizarParametricaEventosSignificativos(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaEventosSignificativos", r)
}

// SincronizarParametricaMotivoAnulacion consulta los motivos válidos de anulación de facturas.
func (c *Client) SincronizarParametricaMotivoAnulacion(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaMotivoAnulacion", r)
}

// SincronizarParametricaPaisOrigen consulta el catálogo de países de origen.
func (c *Client) SincronizarParametricaPaisOrigen(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaPaisOrigen", r)
}

// SincronizarParametricaTipoDocumentoIdentidad consulta los tipos de documentos de identidad admisibles.
func (c *Client) SincronizarParametricaTipoDocumentoIdentidad(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTipoDocumentoIdentidad", r)
}

// SincronizarParametricaTipoDocumentoSector consulta los códigos y tipos de documento sector.
func (c *Client) SincronizarParametricaTipoDocumentoSector(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTipoDocumentoSector", r)
}

// SincronizarParametricaTipoEmision consulta los tipos de emisión (en línea, fuera de línea / contingencia).
func (c *Client) SincronizarParametricaTipoEmision(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTipoEmision", r)
}

// SincronizarParametricaTipoHabitacion consulta los tipos de habitación para sector hotelero.
func (c *Client) SincronizarParametricaTipoHabitacion(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTipoHabitacion", r)
}

// SincronizarParametricaTipoMetodoPago consulta los métodos de pago (efectivo, tarjeta, transferencia, etc.).
func (c *Client) SincronizarParametricaTipoMetodoPago(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTipoMetodoPago", r)
}

// SincronizarParametricaTipoMoneda consulta las monedas aceptadas por el SIN.
func (c *Client) SincronizarParametricaTipoMoneda(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTipoMoneda", r)
}

// SincronizarParametricaTipoPuntoVenta consulta los tipos de punto de venta (fijo, móvil, etc.).
func (c *Client) SincronizarParametricaTipoPuntoVenta(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTipoPuntoVenta", r)
}

// SincronizarParametricaTiposFactura consulta los tipos de factura (con derecho a crédito fiscal, sin crédito fiscal, etc.).
func (c *Client) SincronizarParametricaTiposFactura(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaTiposFactura", r)
}

// SincronizarParametricaUnidadMedida consulta las unidades de medida oficiales.
func (c *Client) SincronizarParametricaUnidadMedida(ctx context.Context, r SolicitudSincronizacion) (*ResultadoParametricas, error) {
	return query[ResultadoParametricas](ctx, c, "sincronizarParametricaUnidadMedida", r)
}

// VerificarComunicacion comprueba la conectividad con el servicio de sincronización.
func (c *Client) VerificarComunicacion(ctx context.Context) (*RespuestaComunicacion, error) {
	if c == nil || c.soap == nil {
		return nil, fmt.Errorf("cliente SIAT no inicializado")
	}
	req := struct {
		XMLName xml.Name `xml:"ns:verificarComunicacion"`
	}{}
	var resp struct {
		XMLName   xml.Name               `xml:"verificarComunicacionResponse"`
		Resultado *RespuestaComunicacion `xml:"return"`
	}
	if err := c.soap.Call(ctx, req, &resp); err != nil {
		return nil, err
	}
	if resp.Resultado == nil {
		return nil, fmt.Errorf("respuesta SIAT sin resultado de comunicación")
	}
	return resp.Resultado, nil
}
