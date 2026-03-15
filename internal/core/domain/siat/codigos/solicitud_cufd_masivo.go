package codigos

import (
	"encoding/xml"
	"time"
)

// CufdMasivo define el sobre SOAP para la solicitud masiva de múltiples códigos CUFD.
type CufdMasivo struct {
	XMLName             xml.Name            `xml:"ns:cufdMasivo" json:"-"`
	SolicitudCufdMasivo SolicitudCufdMasivo `xml:"SolicitudCufdMasivo" json:"solicitudCufdMasivo"`
}

// SolicitudCufdMasivo agrupa las credenciales del sistema y una lista de solicitudes individuales
// para tramitar varios códigos CUFD en un solo llamado al servicio.
type SolicitudCufdMasivo struct {
	CodigoAmbiente  int                     `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad int                     `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoSistema   string                  `xml:"codigoSistema" json:"codigoSistema"`
	DatosSolicitud  []SolicitudListaCufdDto `xml:"datosSolicitud" json:"datosSolicitud"`
	Nit             int64                   `xml:"nit" json:"nit"`
}

// SolicitudListaCufdDto define los datos específicos para cada CUFD solicitado masivamente,
// asociándonlo a un punto de venta y sucursal determinados.
type SolicitudListaCufdDto struct {
	CodigoPuntoVenta *int   `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSucursal   int    `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis             string `xml:"cuis" json:"cuis"`
}

// CufdMasivoResponse define el sobre de respuesta tras una operación de solicitud masiva de CUFD.
type CufdMasivoResponse struct {
	XMLName             xml.Name            `xml:"cufdMasivoResponse" json:"-"`
	RespuestaCufdMasivo RespuestaCufdMasivo `xml:"RespuestaCufdMasivo" json:"respuestaCufdMasivo"`
}

// RespuestaCufdMasivo contiene la lista de respuestas individuales para cada CUFD solicitado,
// permitiendo verificar el éxito o falla de cada trámite dentro de la operación masiva.
type RespuestaCufdMasivo struct {
	ListaRespuestasCufd []RespuestaListaRegistroCufdSoapDto `xml:"listaRespuestasCufd,omitempty" json:"listaRespuestasCufd,omitempty"`
	MensajesList        []MensajeServicio                   `xml:"mensajesList,omitempty" json:"mensajesList,omitempty"`
	Transaccion         bool                                `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}

// RespuestaListaRegistroCufdSoapDto encapsula los datos individuales de un CUFD generado
// en una solicitud masiva, incluyendo su código, control, dirección y vigencia.
type RespuestaListaRegistroCufdSoapDto struct {
	Codigo              string            `xml:"codigo" json:"codigo"`
	CodigoControl       string            `xml:"codigoControl" json:"codigoControl"`
	CodigoPuntoVenta    *int              `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSucursal      *int              `xml:"codigoSucursal" json:"codigoSucursal"`
	Cuis                string            `xml:"cuis" json:"cuis"`
	Direccion           string            `xml:"direccion" json:"direccion"`
	FechaVigencia       time.Time         `xml:"fechaVigencia" json:"fechaVigencia"`
	MensajeServicioList []MensajeServicio `xml:"mensajeServicioList" json:"mensajeServicioList"`
	Transaccion         bool              `xml:"transaccion" json:"transaccion"`
}
