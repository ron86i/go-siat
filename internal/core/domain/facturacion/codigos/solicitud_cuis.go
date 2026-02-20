package codigos

import (
	"encoding/xml"
	"time"
)

// Cuis representa el sobre SOAP para la solicitud de un Código Único de Inicio de Sistemas.
type Cuis struct {
	XMLName       xml.Name      `xml:"ns:cuis"`
	SolicitudCuis SolicitudCuis `xml:"SolicitudCuis"`
}

// SolicitudCuis contiene los parámetros necesarios para obtener un CUIS, identificando de manera única
// al sistema, la sucursal y el punto de venta que iniciarán operaciones.
type SolicitudCuis struct {
	CodigoAmbiente   int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad  int    `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSistema    string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal   int    `xml:"codigoSucursal" json:"codigoSucursal"`
	NIT              int64  `xml:"nit" json:"nit"`
}

// CuisResponse define la estructura del sobre de respuesta SOAP tras solicitar un CUIS.
type CuisResponse struct {
	XMLName       xml.Name      `xml:"cuisResponse" json:"-"`
	RespuestaCuis RespuestaCuis `xml:"RespuestaCuis" json:"respuestaCuis"`
}

// RespuestaCuis encapsula el resultado de la solicitud de CUIS, incluyendo el código generado,
// su fecha de vigencia y cualquier notificación o error emitido por el SIAT.
type RespuestaCuis struct {
	Codigo        string    `xml:"codigo,omitempty" json:"codigo,omitempty"`
	FechaVigencia time.Time `xml:"fechaVigencia,omitempty" json:"fechaVigencia,omitempty"`
	MensajesList  []Mensaje `xml:"mensajesList,omitempty" json:"mensajesList,omitempty"`
	Transaccion   bool      `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}
