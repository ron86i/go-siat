package facturacion

import (
	"encoding/xml"
	"time"
)

// SolicitudRecepcionFactura representa los datos para la recepción individual de una factura.
// Archivo contiene los bytes del XML (firmado si es electrónica) comprimido en GZIP.
type SolicitudRecepcionFactura struct {
	CodigoDocumentoSector int
	CodigoEmision         int
	CodigoModalidad       int
	CodigoPuntoVenta      *int
	CodigoSucursal        int
	Cufd                  string
	Cuis                  string
	TipoFacturaDocumento  int
	Archivo               []byte
	FechaEnvio            time.Time
	HashArchivo           string
}

// SolicitudEstadoFactura identifica la factura para consultar su estado o revertir anulación.
type SolicitudEstadoFactura struct {
	CodigoDocumentoSector int
	CodigoEmision         int
	CodigoModalidad       int
	CodigoPuntoVenta      *int
	CodigoSucursal        int
	Cufd                  string
	Cuis                  string
	TipoFacturaDocumento  int
	Cuf                   string
}

// SolicitudAnulacionFactura incluye el motivo y los datos de la factura a anular.
type SolicitudAnulacionFactura struct {
	SolicitudEstadoFactura
	CodigoMotivo int
}

// SolicitudRecepcionPaquete contiene un TAR.GZ de facturas de contingencia.
type SolicitudRecepcionPaquete struct {
	SolicitudRecepcionFactura
	CantidadFacturas int
	CodigoEvento     int64
}

// SolicitudValidacionPaquete identifica un paquete recibido por el SIN para validar su resultado.
type SolicitudValidacionPaquete struct {
	CodigoDocumentoSector int
	CodigoEmision         int
	CodigoModalidad       int
	CodigoPuntoVenta      *int
	CodigoSucursal        int
	Cufd                  string
	Cuis                  string
	TipoFacturaDocumento  int
	CodigoRecepcion       string
}

// SolicitudRecepcionMasiva contiene un TAR.GZ de emisión masiva.
type SolicitudRecepcionMasiva struct {
	SolicitudRecepcionFactura
	CantidadFacturas int
}

// SolicitudValidacionMasiva identifica una recepción masiva para validar su procesamiento.
type SolicitudValidacionMasiva SolicitudValidacionPaquete

// Serialización XML interna según el contrato WSDL del SIN:

type solicitudRecepcionFacturaXML struct {
	XMLName               xml.Name `xml:"SolicitudServicioRecepcionFactura"`
	CodigoAmbiente        int      `xml:"codigoAmbiente"`
	CodigoDocumentoSector int      `xml:"codigoDocumentoSector"`
	CodigoEmision         int      `xml:"codigoEmision"`
	CodigoModalidad       int      `xml:"codigoModalidad"`
	CodigoPuntoVenta      *int     `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema         string   `xml:"codigoSistema"`
	CodigoSucursal        int      `xml:"codigoSucursal"`
	Cufd                  string   `xml:"cufd"`
	Cuis                  string   `xml:"cuis"`
	Nit                   int64    `xml:"nit"`
	TipoFacturaDocumento  int      `xml:"tipoFacturaDocumento"`
	Archivo               string   `xml:"archivo"`
	FechaEnvio            string   `xml:"fechaEnvio"`
	HashArchivo           string   `xml:"hashArchivo"`
}

type datosOperacionFacturaXML struct {
	XMLName               xml.Name
	CodigoAmbiente        int    `xml:"codigoAmbiente"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector"`
	CodigoEmision         int    `xml:"codigoEmision"`
	CodigoModalidad       int    `xml:"codigoModalidad"`
	CodigoPuntoVenta      *int   `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema         string `xml:"codigoSistema"`
	CodigoSucursal        int    `xml:"codigoSucursal"`
	Cufd                  string `xml:"cufd"`
	Cuis                  string `xml:"cuis"`
	Nit                   int64  `xml:"nit"`
	TipoFacturaDocumento  int    `xml:"tipoFacturaDocumento"`
	CodigoMotivo          *int   `xml:"codigoMotivo,omitempty"`
	Cuf                   string `xml:"cuf"`
}

type datosPaqueteXML struct {
	XMLName               xml.Name
	CodigoAmbiente        int    `xml:"codigoAmbiente"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector"`
	CodigoEmision         int    `xml:"codigoEmision"`
	CodigoModalidad       int    `xml:"codigoModalidad"`
	CodigoPuntoVenta      *int   `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema         string `xml:"codigoSistema"`
	CodigoSucursal        int    `xml:"codigoSucursal"`
	Cufd                  string `xml:"cufd"`
	Cuis                  string `xml:"cuis"`
	Nit                   int64  `xml:"nit"`
	TipoFacturaDocumento  int    `xml:"tipoFacturaDocumento"`
	Archivo               string `xml:"archivo"`
	FechaEnvio            string `xml:"fechaEnvio"`
	HashArchivo           string `xml:"hashArchivo"`
	CantidadFacturas      int    `xml:"cantidadFacturas"`
	CodigoEvento          int64  `xml:"codigoEvento"`
}

type datosValidacionPaqueteXML struct {
	XMLName               xml.Name
	CodigoAmbiente        int    `xml:"codigoAmbiente"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector"`
	CodigoEmision         int    `xml:"codigoEmision"`
	CodigoModalidad       int    `xml:"codigoModalidad"`
	CodigoPuntoVenta      *int   `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema         string `xml:"codigoSistema"`
	CodigoSucursal        int    `xml:"codigoSucursal"`
	Cufd                  string `xml:"cufd"`
	Cuis                  string `xml:"cuis"`
	Nit                   int64  `xml:"nit"`
	TipoFacturaDocumento  int    `xml:"tipoFacturaDocumento"`
	CodigoRecepcion       string `xml:"codigoRecepcion"`
}

type datosMasivaXML struct {
	XMLName               xml.Name
	CodigoAmbiente        int    `xml:"codigoAmbiente"`
	CodigoDocumentoSector int    `xml:"codigoDocumentoSector"`
	CodigoEmision         int    `xml:"codigoEmision"`
	CodigoModalidad       int    `xml:"codigoModalidad"`
	CodigoPuntoVenta      *int   `xml:"codigoPuntoVenta,omitempty"`
	CodigoSistema         string `xml:"codigoSistema"`
	CodigoSucursal        int    `xml:"codigoSucursal"`
	Cufd                  string `xml:"cufd"`
	Cuis                  string `xml:"cuis"`
	Nit                   int64  `xml:"nit"`
	TipoFacturaDocumento  int    `xml:"tipoFacturaDocumento"`
	Archivo               string `xml:"archivo"`
	FechaEnvio            string `xml:"fechaEnvio"`
	HashArchivo           string `xml:"hashArchivo"`
	CantidadFacturas      int    `xml:"cantidadFacturas"`
}
