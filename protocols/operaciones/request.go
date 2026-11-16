package operaciones

import (
	"encoding/xml"
	"time"
)

// SolicitudRegistroEvento contiene los datos para registrar un evento significativo (contingencia).
type SolicitudRegistroEvento struct {
	CodigoMotivoEvento    int
	CodigoPuntoVenta      *int
	CodigoSucursal        int
	Cufd                  string
	CufdEvento            string
	Cuis                  string
	Descripcion           string
	FechaHoraInicioEvento time.Time
	FechaHoraFinEvento    time.Time
}

// SolicitudConsultaEvento identifica la fecha y los datos para consultar eventos significativos.
type SolicitudConsultaEvento struct {
	CodigoPuntoVenta *int
	CodigoSucursal   int
	Cufd             string
	Cuis             string
	FechaEvento      time.Time
}

// SolicitudRegistroPuntoVenta contiene los datos para registrar un punto de venta.
type SolicitudRegistroPuntoVenta struct {
	CodigoModalidad      int
	CodigoSucursal       int
	CodigoTipoPuntoVenta int
	Cuis                 string
	Descripcion          string
	NombrePuntoVenta     string
}

// SolicitudConsultaPuntoVenta identifica la sucursal y el CUIS para consultar puntos de venta.
type SolicitudConsultaPuntoVenta struct {
	CodigoSucursal int
	Cuis           string
}

// SolicitudCierrePuntoVenta identifica el punto de venta a cerrar.
type SolicitudCierrePuntoVenta struct {
	CodigoPuntoVenta int
	CodigoSucursal   int
	Cuis             string
}

// SolicitudRegistroComisionista contiene los datos y vigencia del contrato de un punto comisionista.
type SolicitudRegistroComisionista struct {
	CodigoModalidad  int
	CodigoSucursal   int
	Cuis             string
	Descripcion      string
	FechaFin         time.Time
	FechaInicio      time.Time
	NitComisionista  int64
	NombrePuntoVenta string
	NumeroContrato   string
}

// SolicitudCierreSistema contiene las credenciales y ubicación para cerrar operaciones del sistema.
type SolicitudCierreSistema struct {
	CodigoModalidad  int
	CodigoPuntoVenta *int
	CodigoSucursal   int
	Cuis             string
}

// Serialización XML interna según el esquema WSDL del SIN.

type solicitudRegistroEventoXML struct {
	XMLName               xml.Name `xml:"SolicitudEventoSignificativo"`
	CodigoAmbiente        int      `xml:"codigoAmbiente"`
	CodigoMotivoEvento    int      `xml:"codigoMotivoEvento"`
	CodigoPuntoVenta      *int     `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema         string   `xml:"codigoSistema"`
	CodigoSucursal        int      `xml:"codigoSucursal"`
	Cufd                  string   `xml:"cufd"`
	CufdEvento            string   `xml:"cufdEvento"`
	Cuis                  string   `xml:"cuis"`
	Descripcion           string   `xml:"descripcion"`
	FechaHoraFinEvento    string   `xml:"fechaHoraFinEvento"`
	FechaHoraInicioEvento string   `xml:"fechaHoraInicioEvento"`
	Nit                   int64    `xml:"nit"`
}

type solicitudConsultaEventoXML struct {
	XMLName          xml.Name `xml:"SolicitudConsultaEvento"`
	CodigoAmbiente   int      `xml:"codigoAmbiente"`
	CodigoPuntoVenta *int     `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema    string   `xml:"codigoSistema"`
	CodigoSucursal   int      `xml:"codigoSucursal"`
	Cufd             string   `xml:"cufd"`
	Cuis             string   `xml:"cuis"`
	FechaEvento      string   `xml:"fechaEvento"`
	Nit              int64    `xml:"nit"`
}

type solicitudRegistroPuntoVentaXML struct {
	XMLName              xml.Name `xml:"SolicitudRegistroPuntoVenta"`
	CodigoAmbiente       int      `xml:"codigoAmbiente"`
	CodigoModalidad      int      `xml:"codigoModalidad"`
	CodigoSistema        string   `xml:"codigoSistema"`
	CodigoSucursal       int      `xml:"codigoSucursal"`
	CodigoTipoPuntoVenta int      `xml:"codigoTipoPuntoVenta"`
	Cuis                 string   `xml:"cuis"`
	Descripcion          string   `xml:"descripcion"`
	Nit                  int64    `xml:"nit"`
	NombrePuntoVenta     string   `xml:"nombrePuntoVenta"`
}

type solicitudConsultaPuntoVentaXML struct {
	XMLName        xml.Name `xml:"SolicitudConsultaPuntoVenta"`
	CodigoAmbiente int      `xml:"codigoAmbiente"`
	CodigoSistema  string   `xml:"codigoSistema"`
	CodigoSucursal int      `xml:"codigoSucursal"`
	Cuis           string   `xml:"cuis"`
	Nit            int64    `xml:"nit"`
}

type solicitudCierrePuntoVentaXML struct {
	XMLName          xml.Name `xml:"SolicitudCierrePuntoVenta"`
	CodigoAmbiente   int      `xml:"codigoAmbiente"`
	CodigoPuntoVenta int      `xml:"codigoPuntoVenta"`
	CodigoSistema    string   `xml:"codigoSistema"`
	CodigoSucursal   int      `xml:"codigoSucursal"`
	Cuis             string   `xml:"cuis"`
	Nit              int64    `xml:"nit"`
}

type solicitudRegistroComisionistaXML struct {
	XMLName          xml.Name `xml:"SolicitudPuntoVentaComisionista"`
	CodigoAmbiente   int      `xml:"codigoAmbiente"`
	CodigoModalidad  int      `xml:"codigoModalidad"`
	CodigoSistema    string   `xml:"codigoSistema"`
	CodigoSucursal   int      `xml:"codigoSucursal"`
	Cuis             string   `xml:"cuis"`
	Descripcion      string   `xml:"descripcion"`
	FechaFin         string   `xml:"fechaFin"`
	FechaInicio      string   `xml:"fechaInicio"`
	Nit              int64    `xml:"nit"`
	NitComisionista  int64    `xml:"nitComisionista"`
	NombrePuntoVenta string   `xml:"nombrePuntoVenta"`
	NumeroContrato   string   `xml:"numeroContrato"`
}

type solicitudCierreSistemaXML struct {
	XMLName          xml.Name `xml:"SolicitudOperaciones"`
	CodigoAmbiente   int      `xml:"codigoAmbiente"`
	CodigoModalidad  int      `xml:"codigoModalidad"`
	CodigoPuntoVenta *int     `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema    string   `xml:"codigoSistema"`
	CodigoSucursal   int      `xml:"codigoSucursal"`
	Cuis             string   `xml:"cuis"`
	Nit              int64    `xml:"nit"`
}
