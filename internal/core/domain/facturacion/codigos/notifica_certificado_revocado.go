package codigos

import (
	"encoding/xml"
	"time"
)

// NotificaCertificadoRevocado encapsula la solicitud necesaria para informar al SIAT sobre la revocación de un certificado digital.
type NotificaCertificadoRevocado struct {
	XMLName                   xml.Name                 `xml:"ns:notificaCertificadoRevocado" json:"-"`
	SolicitudNotificaRevocado SolicitudNotifcaRevocado `xml:"SolicitudNotificaRevocado" json:"solicitudNotificaRevocado"`
}

// SolicitudNotifcaRevocado contiene los detalles técnicos y de negocio requeridos para revocar un certificado,
// incluyendo credenciales del sistema, datos del emisor y la justificación del evento.
type SolicitudNotifcaRevocado struct {
	Certificado     string     `xml:"certificado" json:"certificado"`
	CodigoAmbiente  int        `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoSistema   string     `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal  int        `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis            string     `xml:"cuis" json:"cuis"`
	FechaRevocacion *time.Time `xml:"fechaRevocacion" json:"fechaRevocacion"`
	NIT             int64      `xml:"nit" json:"nit"`
	RazonRevocacion string     `xml:"razonRevocacion" json:"razonRevocacion"`
}

// NotificaCertificadoRevocadoResponse define la estructura del sobre de respuesta tras una notificación de revocación.
type NotificaCertificadoRevocadoResponse struct {
	XMLName                   xml.Name                  `xml:"notificaCertificadoRevocadoResponse" json:"-"`
	RespuestaNotificaRevocado RespuestaNotificaRevocado `xml:"RespuestaNotificaRevocado" json:"respuestaNotificaRevocado"`
}

// RespuestaNotificaRevocado contiene el resultado procesado por el SIAT, indicando si la transacción
// fue exitosa e incluyendo cualquier mensaje informativo o de error generado.
type RespuestaNotificaRevocado struct {
	MensajesList []MensajeServicio `xml:"mensajesList,omitempty" json:"mensajesList,omitempty"`
	Transaccion  bool              `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}
