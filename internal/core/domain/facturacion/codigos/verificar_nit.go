package codigos

import "encoding/xml"

// VerificarNit representa el sobre SOAP para la solicitud de validación de un NIT ante el SIAT.
type VerificarNit struct {
	XMLName               xml.Name              `xml:"ns:verificarNit" json:"-"`
	SolicitudVerificarNit SolicitudVerificarNit `xml:"SolicitudVerificarNit" json:"solicitudVerificarNit"`
}

// SolicitudVerificarNit contiene los parámetros técnicos y el número identificador
// del contribuyente que se desea validar en el padrón nacional.
type SolicitudVerificarNit struct {
	CodigoAmbiente      int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad     int    `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoSistema       string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal      int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis                string `xml:"cuis" json:"cuis"`
	NIT                 int64  `xml:"nit" json:"nit"`
	NitParaVerificacion int64  `xml:"nitParaVerificacion" json:"nitParaVerificacion"`
}

// VerificarNitResponse define la estructura del sobre de respuesta tras la validación de un NIT.
type VerificarNitResponse struct {
	XMLName               xml.Name              `xml:"verificarNitResponse" json:"-"`
	RespuestaVerificarNit RespuestaVerificarNit `xml:"RespuestaVerificarNit" json:"respuestaVerificarNit"`
}

// RespuestaVerificarNit comunica si el NIT consultado se encuentra vigente y activo,
// adjuntando mensajes informativos en caso de discrepancias o errores técnicos.
type RespuestaVerificarNit struct {
	MensajesList []MensajeServicio `xml:"mensajesList,omitempty" json:"mensajesList,omitempty"`
	Transaccion  bool              `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}
