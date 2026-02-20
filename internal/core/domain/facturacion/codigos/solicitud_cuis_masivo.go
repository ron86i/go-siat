package codigos

import (
	"encoding/xml"
	"time"
)

// CuisMasivo define el sobre SOAP para realizar la solicitud masiva de códigos CUIS.
type CuisMasivo struct {
	XMLName                     xml.Name                    `xml:"ns:cuisMasivo" json:"-"`
	SolicitudCuisMasivoSistemas SolicitudCuisMasivoSistemas `xml:"SolicitudCuisMasivoSistemas" json:"solicitudCuisMasivoSistemas"`
}

// SolicitudCuisMasivoSistemas agrupa los datos de configuración del sistema y una lista de solicitudes
// específicas para generar múltiples códigos CUIS en un único proceso.
type SolicitudCuisMasivoSistemas struct {
	CodigoAmbiente  int                     `xml:"codigoAmbiente" json:"codigoAmbiente"`
	CodigoModalidad int                     `xml:"codigoModalidad" json:"codigoModalidad"`
	CodigoSistema   string                  `xml:"codigoSistema" json:"codigoSistema"`
	DatosSolicitud  []SolicitudListaCuisDto `xml:"datosSolicitud" json:"datosSolicitud"`
	NIT             int64                   `xml:"nit" json:"nit"`
}

// SolicitudListaCuisDto detalla la sucursal y el punto de venta específicos para los cuales
// se solicita un CUIS dentro de una operación masiva.
type SolicitudListaCuisDto struct {
	CodigoPuntoVenta *int `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	CodigoSucursal   int  `xml:"codigoSucursal" json:"codigoSucursal"`
}

// CuisMasivoResponse define la estructura del sobre de respuesta tras una tramitación masiva de CUIS.
type CuisMasivoResponse struct {
	XMLName             xml.Name            `xml:"cuisMasivoResponse" json:"-"`
	RespuestaCuisMasivo RespuestaCuisMasivo `xml:"RespuestaCuisMasivo" json:"respuestaCuisMasivo"`
}

// RespuestaCuisMasivo contiene el conjunto de resultados para cada CUIS solicitado masivamente,
// permitiendo identificar de manera agregada los éxitos y fallos del proceso.
type RespuestaCuisMasivo struct {
	ListaRespuestasCuis []RespuestaListaRegistroCuisSoapDto `xml:"listaRespuestasCuis,omitempty" json:"listaRespuestasCuis,omitempty"`
	MensajesList        []Mensaje                           `xml:"mensajesList,omitempty" json:"mensajesList,omitempty"`
	Transaccion         bool                                `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}

// RespuestaListaRegistroCuisSoapDto representa el resultado individual de un CUIS tramitado masivamente,
// incluyendo su código, vigencia y mensajes de servicio específicos.
type RespuestaListaRegistroCuisSoapDto struct {
	Codigo              string    `xml:"codigo,omitempty" json:"codigo,omitempty"`
	CodigoPuntoVenta    *int32    `xml:"codigoPuntoVenta,omitempty" json:"codigoPuntoVenta,omitempty"`
	CodigoSucursal      *int32    `xml:"codigoSucursal,omitempty" json:"codigoSucursal,omitempty"`
	FechaVigencia       time.Time `xml:"fechaVigencia,omitempty" json:"fechaVigencia,omitempty"`
	MensajeServicioList []Mensaje `xml:"mensajeServicioList,omitempty" json:"mensajeServicioList,omitempty"`
	Transaccion         bool      `xml:"transaccion,omitempty" json:"transaccion,omitempty"`
}
