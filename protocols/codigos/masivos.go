package codigos

import (
	"context"
	"encoding/xml"
	"fmt"
)

// PuntoCuis identifica una sucursal y un punto de venta para solicitar CUIS.
type PuntoCuis struct {
	CodigoPuntoVenta *int `xml:"codigoPuntoVenta,omitempty"`
	CodigoSucursal   int  `xml:"codigoSucursal"`
}

// PuntoCufd identifica el CUIS y ubicación para solicitar CUFD.
type PuntoCufd struct {
	CodigoPuntoVenta *int   `xml:"codigoPuntoVenta,omitempty"`
	CodigoSucursal   int    `xml:"codigoSucursal"`
	Cuis             string `xml:"cuis"`
}

// SolicitudCuisMasivo agrupa los puntos de una modalidad.
type SolicitudCuisMasivo struct {
	CodigoModalidad int
	Puntos          []PuntoCuis
}

// SolicitudCufdMasivo agrupa los puntos de una modalidad.
type SolicitudCufdMasivo struct {
	CodigoModalidad int
	Puntos          []PuntoCufd
}

// ResultadoCuisMasivo conserva un resultado por punto de venta.
type ResultadoCuisMasivo struct {
	Codigo           string    `xml:"codigo"`
	CodigoPuntoVenta *int      `xml:"codigoPuntoVenta"`
	CodigoSucursal   *int      `xml:"codigoSucursal"`
	FechaVigencia    string    `xml:"fechaVigencia"`
	Mensajes         []Mensaje `xml:"mensajeServicioList"`
	Transaccion      *bool     `xml:"transaccion"`
}

// ResultadoCufdMasivo conserva un resultado CUFD por punto de venta.
type ResultadoCufdMasivo struct {
	Codigo           string    `xml:"codigo"`
	CodigoControl    string    `xml:"codigoControl"`
	CodigoInterno    string    `xml:"codigoInterno"`
	CodigoPuntoVenta *int      `xml:"codigoPuntoVenta"`
	CodigoSucursal   *int      `xml:"codigoSucursal"`
	Cuis             string    `xml:"cuis"`
	Direccion        string    `xml:"direccion"`
	FechaVigencia    string    `xml:"fechaVigencia"`
	Mensajes         []Mensaje `xml:"mensajeServicioList"`
	Transaccion      *bool     `xml:"transaccion"`
}

// RespuestaCuisMasivo contiene resultados individuales y estado general.
type RespuestaCuisMasivo struct {
	Resultados  []ResultadoCuisMasivo `xml:"listaRespuestasCuis"`
	Mensajes    []Mensaje             `xml:"mensajesList"`
	Transaccion *bool                 `xml:"transaccion"`
}

// RespuestaCufdMasivo contiene resultados individuales y estado general.
type RespuestaCufdMasivo struct {
	Resultados  []ResultadoCufdMasivo `xml:"listaRespuestasCufd"`
	Mensajes    []Mensaje             `xml:"mensajesList"`
	Transaccion *bool                 `xml:"transaccion"`
}

// CuisMasivo solicita CUIS para varios puntos de venta.
func (c *Client) CuisMasivo(ctx context.Context, r SolicitudCuisMasivo) (RespuestaCuisMasivo, error) {
	if c == nil || c.soap == nil {
		return RespuestaCuisMasivo{}, fmt.Errorf("cliente SIAT no inicializado")
	}
	if (r.CodigoModalidad != 1 && r.CodigoModalidad != 2) || len(r.Puntos) == 0 {
		return RespuestaCuisMasivo{}, fmt.Errorf("solicitud CUIS masivo incompleta o inválida")
	}
	for i, p := range r.Puntos {
		if p.CodigoSucursal < 0 || (p.CodigoPuntoVenta != nil && *p.CodigoPuntoVenta < 0) {
			return RespuestaCuisMasivo{}, fmt.Errorf("punto %d inválido", i+1)
		}
	}
	cfg := c.soap.Config()
	dato := struct {
		CodigoAmbiente  int         `xml:"codigoAmbiente"`
		CodigoModalidad int         `xml:"codigoModalidad"`
		CodigoSistema   string      `xml:"codigoSistema"`
		Puntos          []PuntoCuis `xml:"datosSolicitud"`
		Nit             int64       `xml:"nit"`
	}{cfg.CodigoAmbiente, r.CodigoModalidad, cfg.CodigoSistema, r.Puntos, cfg.Nit}
	req := struct {
		XMLName   xml.Name `xml:"ns:cuisMasivo"`
		Solicitud any      `xml:"SolicitudCuisMasivoSistemas"`
	}{Solicitud: dato}
	var resp struct {
		XMLName   xml.Name             `xml:"cuisMasivoResponse"`
		Resultado *RespuestaCuisMasivo `xml:"RespuestaCuisMasivo"`
	}
	if err := c.soap.Call(ctx, req, &resp); err != nil {
		return RespuestaCuisMasivo{}, err
	}
	if resp.Resultado == nil {
		return RespuestaCuisMasivo{}, fmt.Errorf("respuesta SIAT sin resultado CUIS masivo")
	}
	return *resp.Resultado, nil
}

// CufdMasivo solicita CUFD para varios puntos de venta.
func (c *Client) CufdMasivo(ctx context.Context, r SolicitudCufdMasivo) (RespuestaCufdMasivo, error) {
	if c == nil || c.soap == nil {
		return RespuestaCufdMasivo{}, fmt.Errorf("cliente SIAT no inicializado")
	}
	if (r.CodigoModalidad != 1 && r.CodigoModalidad != 2) || len(r.Puntos) == 0 {
		return RespuestaCufdMasivo{}, fmt.Errorf("solicitud CUFD masivo incompleta o inválida")
	}
	for i, p := range r.Puntos {
		if p.CodigoSucursal < 0 || (p.CodigoPuntoVenta != nil && *p.CodigoPuntoVenta < 0) || p.Cuis == "" {
			return RespuestaCufdMasivo{}, fmt.Errorf("punto %d inválido", i+1)
		}
	}
	cfg := c.soap.Config()
	dato := struct {
		CodigoAmbiente  int         `xml:"codigoAmbiente"`
		CodigoModalidad int         `xml:"codigoModalidad"`
		CodigoSistema   string      `xml:"codigoSistema"`
		Puntos          []PuntoCufd `xml:"datosSolicitud"`
		Nit             int64       `xml:"nit"`
	}{cfg.CodigoAmbiente, r.CodigoModalidad, cfg.CodigoSistema, r.Puntos, cfg.Nit}
	req := struct {
		XMLName   xml.Name `xml:"ns:cufdMasivo"`
		Solicitud any      `xml:"SolicitudCufdMasivo"`
	}{Solicitud: dato}
	var resp struct {
		XMLName   xml.Name             `xml:"cufdMasivoResponse"`
		Resultado *RespuestaCufdMasivo `xml:"RespuestaCufdMasivo"`
	}
	if err := c.soap.Call(ctx, req, &resp); err != nil {
		return RespuestaCufdMasivo{}, err
	}
	if resp.Resultado == nil {
		return RespuestaCufdMasivo{}, fmt.Errorf("respuesta SIAT sin resultado CUFD masivo")
	}
	return *resp.Resultado, nil
}
