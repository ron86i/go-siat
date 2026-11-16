package facturacion

import (
	"github.com/ron86i/go-siat/v3/documents/alcanzadaice"
	"github.com/ron86i/go-siat/v3/documents/alquileres"
	"github.com/ron86i/go-siat/v3/documents/alquilerzf"
	"github.com/ron86i/go-siat/v3/documents/biodiesel"
	"github.com/ron86i/go-siat/v3/documents/boletoaereo"
	"github.com/ron86i/go-siat/v3/documents/combustiblesinsubvencion"
	"github.com/ron86i/go-siat/v3/documents/comercialexportacion"
	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/ron86i/go-siat/v3/documents/compraventabonificaciones"
	"github.com/ron86i/go-siat/v3/documents/duttyfree"
	"github.com/ron86i/go-siat/v3/documents/entidadesfinancieras"
	"github.com/ron86i/go-siat/v3/documents/exportacionhidrocarburos"
	"github.com/ron86i/go-siat/v3/documents/exportacionminerales"
	"github.com/ron86i/go-siat/v3/documents/exportacionprecioventa"
	"github.com/ron86i/go-siat/v3/documents/exportacionservicios"
	"github.com/ron86i/go-siat/v3/documents/gnglsp"
	"github.com/ron86i/go-siat/v3/documents/gnv"
	"github.com/ron86i/go-siat/v3/documents/hidrocarburosiehd"
	"github.com/ron86i/go-siat/v3/documents/hidrocarburosnoiehd"
	"github.com/ron86i/go-siat/v3/documents/hospitales"
	"github.com/ron86i/go-siat/v3/documents/hospitaleszf"
	"github.com/ron86i/go-siat/v3/documents/hoteles"
	"github.com/ron86i/go-siat/v3/documents/juegosazar"
	"github.com/ron86i/go-siat/v3/documents/libreconsignacion"
	"github.com/ron86i/go-siat/v3/documents/lubricantes"
	"github.com/ron86i/go-siat/v3/documents/lubricantesiehd"
	"github.com/ron86i/go-siat/v3/documents/monedaextranjera"
	"github.com/ron86i/go-siat/v3/documents/prevalorada"
	"github.com/ron86i/go-siat/v3/documents/prevaloradasd"
	"github.com/ron86i/go-siat/v3/documents/sectoreducativo"
	"github.com/ron86i/go-siat/v3/documents/sectoreducativozf"
	"github.com/ron86i/go-siat/v3/documents/seguridadalimentaria"
	"github.com/ron86i/go-siat/v3/documents/seguros"
	"github.com/ron86i/go-siat/v3/documents/serviciosbasicos"
	"github.com/ron86i/go-siat/v3/documents/serviciosbasicoszf"
	"github.com/ron86i/go-siat/v3/documents/suministroenergia"
	"github.com/ron86i/go-siat/v3/documents/tasacero"
	"github.com/ron86i/go-siat/v3/documents/telecomunicaciones"
	"github.com/ron86i/go-siat/v3/documents/telecomunicacioneszf"
	"github.com/ron86i/go-siat/v3/documents/turismohospedaje"
	"github.com/ron86i/go-siat/v3/documents/ventamineralbcb"
	"github.com/ron86i/go-siat/v3/documents/ventaminerales"
	"github.com/ron86i/go-siat/v3/documents/vuelocharter"
	"github.com/ron86i/go-siat/v3/documents/zonafranca"
)

// Documento es el contrato común de los modelos de factura admitidos para emisión en SIAT v3.
type Documento interface {
	Validate() error
	CodigoDocumentoSector() int
	CodigoSucursal() int
	CodigoPuntoVenta() *int
	Cufd() string
	Protocolo() int
	MarshalXMLModalidad(modalidad int) ([]byte, error)
}

// DocumentosPermitidos limita T a los modelos de factura sectoriales disponibles en v3.
type DocumentosPermitidos interface {
	Documento
	*alcanzadaice.Factura |
		*alquileres.Factura |
		*alquilerzf.Factura |
		*biodiesel.Factura |
		*boletoaereo.Factura |
		*combustiblesinsubvencion.Factura |
		*comercialexportacion.Factura |
		*compraventa.Factura |
		*compraventabonificaciones.Factura |
		*duttyfree.Factura |
		*entidadesfinancieras.Factura |
		*exportacionhidrocarburos.Factura |
		*exportacionminerales.Factura |
		*exportacionprecioventa.Factura |
		*exportacionservicios.Factura |
		*gnglsp.Factura |
		*gnv.Factura |
		*hidrocarburosiehd.Factura |
		*hidrocarburosnoiehd.Factura |
		*hospitales.Factura |
		*hospitaleszf.Factura |
		*hoteles.Factura |
		*juegosazar.Factura |
		*libreconsignacion.Factura |
		*lubricantes.Factura |
		*lubricantesiehd.Factura |
		*monedaextranjera.Factura |
		*prevalorada.Factura |
		*prevaloradasd.Factura |
		*sectoreducativo.Factura |
		*sectoreducativozf.Factura |
		*seguridadalimentaria.Factura |
		*seguros.Factura |
		*serviciosbasicos.Factura |
		*serviciosbasicoszf.Factura |
		*suministroenergia.Factura |
		*tasacero.Factura |
		*telecomunicaciones.Factura |
		*telecomunicacioneszf.Factura |
		*turismohospedaje.Factura |
		*ventamineralbcb.Factura |
		*ventaminerales.Factura |
		*vuelocharter.Factura |
		*zonafranca.Factura
}

// AsDocumentos convierte un slice tipado de facturas v3 admitidas a []Documento.
func AsDocumentos[T DocumentosPermitidos](facturas []T) []Documento {
	docs := make([]Documento, len(facturas))
	for i, factura := range facturas {
		docs[i] = factura
	}
	return docs
}
