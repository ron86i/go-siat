package codigos

import (
	"encoding/xml"
	"time"
)

// Cufd representa el sobre SOAP para la solicitud de un Código Único de Facturación Diaria.
type Cufd struct {
	XMLName       xml.Name      `xml:"ns:cufd"`
	SolicitudCufd SolicitudCufd `xml:"SolicitudCufd"`
}

// SolicitudCufd contiene los parámetros requeridos para obtener un CUFD, incluyendo la sucursal,
// el punto de venta y las credenciales del sistema para un emisor específico.
type SolicitudCufd struct {
	CodigoAmbiente   int    `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad  int    `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoPuntoVenta *int   `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"` // Opcional, puede ser nulo
	CodigoSistema    string `xml:"codigoSistema" json:"codigoSistema"`
	CodigoSucursal   int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis             string `xml:"cuis" json:"cuis"`
	NIT              int64  `xml:"nit" json:"nit"`
}

// CufdResponse define la estructura del sobre de respuesta SOAP tras solicitar un CUFD.
type CufdResponse struct {
	XMLName       xml.Name      `xml:"cufdResponse" json:"-"`
	RespuestaCufd RespuestaCufd `xml:"RespuestaCufd" json:"respuestaCufd"`
}

// RespuestaCufd encapsula los datos del CUFD generado por el SIAT, como el código propiamente dicho,
// el código de control, la dirección física asociada y la fecha de vigencia.
type RespuestaCufd struct {
	Codigo        string            `xml:"codigo,omitempty" json:"codigo,omitempty"`
	CodigoControl string            `xml:"codigoControl,omitempty" json:"codigoControl,omitempty"`
	Direccion     string            `xml:"direccion,omitempty" json:"direccion,omitempty"`
	FechaVigencia time.Time         `xml:"fechaVigencia,omitempty" json:"fechaVigencia,omitempty"`
	MensajesList  []MensajeServicio `xml:"mensajesList,omitempty" json:"mensajesList,omitempty"`
	Transaccion   bool              `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}
