package codigos

import (
	"context"
	"encoding/xml"
	"fmt"
)

// SolicitudVerificarNit identifica el NIT consultado y las credenciales de sucursal.
type SolicitudVerificarNit struct {
	CodigoModalidad     int
	CodigoSucursal      int
	Cuis                string
	NitParaVerificacion int64
}

// RespuestaVerificarNit conserva el resultado y los mensajes del SIN.
type RespuestaVerificarNit struct {
	Mensajes    []Mensaje `xml:"mensajesList"`
	Transaccion *bool     `xml:"transaccion"`
}

// RespuestaComunicacion conserva la respuesta del servicio.
type RespuestaComunicacion struct {
	Mensajes    []Mensaje `xml:"mensajesList"`
	Transaccion *bool     `xml:"transaccion"`
}

// VerificarNit consulta el estado de un NIT en el SIN.
func (c *Client) VerificarNit(ctx context.Context, r SolicitudVerificarNit) (RespuestaVerificarNit, error) {
	if c == nil || c.soap == nil {
		return RespuestaVerificarNit{}, fmt.Errorf("cliente SIAT no inicializado")
	}
	if (r.CodigoModalidad != 1 && r.CodigoModalidad != 2) || r.CodigoSucursal < 0 || r.Cuis == "" || r.NitParaVerificacion <= 0 {
		return RespuestaVerificarNit{}, fmt.Errorf("solicitud de verificación de NIT incompleta o inválida")
	}
	cfg := c.soap.Config()
	data := struct {
		CodigoAmbiente      int    `xml:"codigoAmbiente"`
		CodigoModalidad     int    `xml:"codigoModalidad"`
		CodigoSistema       string `xml:"codigoSistema"`
		CodigoSucursal      int    `xml:"codigoSucursal"`
		Cuis                string `xml:"cuis"`
		Nit                 int64  `xml:"nit"`
		NitParaVerificacion int64  `xml:"nitParaVerificacion"`
	}{cfg.CodigoAmbiente, r.CodigoModalidad, cfg.CodigoSistema, r.CodigoSucursal, r.Cuis, cfg.Nit, r.NitParaVerificacion}
	req := struct {
		XMLName   xml.Name `xml:"ns:verificarNit"`
		Solicitud any      `xml:"SolicitudVerificarNit"`
	}{Solicitud: data}
	var response struct {
		XMLName   xml.Name               `xml:"verificarNitResponse"`
		Resultado *RespuestaVerificarNit `xml:"RespuestaVerificarNit"`
	}
	if err := c.soap.Call(ctx, req, &response); err != nil {
		return RespuestaVerificarNit{}, err
	}
	if response.Resultado == nil {
		return RespuestaVerificarNit{}, fmt.Errorf("respuesta SIAT sin resultado de verificación de NIT")
	}
	return *response.Resultado, nil
}

// VerificarComunicacion consulta la disponibilidad del servicio de códigos.
func (c *Client) VerificarComunicacion(ctx context.Context) (RespuestaComunicacion, error) {
	if c == nil || c.soap == nil {
		return RespuestaComunicacion{}, fmt.Errorf("cliente SIAT no inicializado")
	}
	req := struct {
		XMLName xml.Name `xml:"ns:verificarComunicacion"`
	}{}
	var response struct {
		XMLName   xml.Name               `xml:"verificarComunicacionResponse"`
		Resultado *RespuestaComunicacion `xml:"RespuestaComunicacion"`
	}
	if err := c.soap.Call(ctx, req, &response); err != nil {
		return RespuestaComunicacion{}, err
	}
	if response.Resultado == nil {
		return RespuestaComunicacion{}, fmt.Errorf("respuesta SIAT sin resultado de comunicación")
	}
	return *response.Resultado, nil
}
