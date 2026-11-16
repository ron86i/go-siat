package compras

import (
	"encoding/xml"
	"time"
)

// DatosSucursal contiene los datos comunes de identificación de sucursal y punto de venta.
type DatosSucursal struct {
	CodigoPuntoVenta int
	CodigoSucursal   int
	Cufd             string
	Cuis             string
}

// SolicitudPaquete contiene el archivo comprimido y los datos del período de compras.
type SolicitudPaquete struct {
	DatosSucursal
	Archivo          []byte
	CantidadFacturas int
	FechaEnvio       time.Time
	Gestion          int
	HashArchivo      string
	Periodo          int
}

// SolicitudValidacion identifica un paquete previamente recibido para consultar su validación.
type SolicitudValidacion struct {
	DatosSucursal
	CodigoRecepcion string
}

// SolicitudAnulacion identifica una compra registrada que se desea anular.
type SolicitudAnulacion struct {
	DatosSucursal
	CodAutorizacion string
	NitProveedor    int64
	NroDuiDim       string
	NroFactura      int64
}

// SolicitudConsulta identifica la fecha para descargar las compras registradas.
type SolicitudConsulta struct {
	DatosSucursal
	Fecha time.Time
}

type baseXML struct {
	CodigoAmbiente   int    `xml:"codigoAmbiente"`
	CodigoPuntoVenta int    `xml:"codigoPuntoVenta"`
	CodigoSistema    string `xml:"codigoSistema"`
	CodigoSucursal   int    `xml:"codigoSucursal"`
	Cufd             string `xml:"cufd"`
	Cuis             string `xml:"cuis"`
	Nit              int64  `xml:"nit"`
}

type datosPaqueteXML struct {
	XMLName xml.Name
	baseXML
	Archivo          string `xml:"archivo"`
	CantidadFacturas int    `xml:"cantidadFacturas"`
	FechaEnvio       string `xml:"fechaEnvio"`
	Gestion          int    `xml:"gestion"`
	HashArchivo      string `xml:"hashArchivo"`
	Periodo          int    `xml:"periodo"`
}

type datosValidacionXML struct {
	XMLName xml.Name
	baseXML
	CodigoRecepcion string `xml:"codigoRecepcion"`
}

type datosAnulacionXML struct {
	XMLName xml.Name
	baseXML
	CodAutorizacion string `xml:"codAutorizacion"`
	NitProveedor    int64  `xml:"nitProveedor"`
	NroDuiDim       string `xml:"nroDuiDim"`
	NroFactura      int64  `xml:"nroFactura"`
}

type datosConsultaXML struct {
	XMLName xml.Name
	baseXML
	Fecha string `xml:"fecha"`
}
