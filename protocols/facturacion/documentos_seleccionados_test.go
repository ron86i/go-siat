package facturacion_test

import (
	"encoding/xml"
	"strings"
	"testing"

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
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
)

func TestAsDocumentosAceptaTodosLosModelosDeFactura(t *testing.T) {
	documentos := [][]facturacion.Documento{
		facturacion.AsDocumentos([]*alcanzadaice.Factura{nil}),
		facturacion.AsDocumentos([]*alquileres.Factura{nil}),
		facturacion.AsDocumentos([]*alquilerzf.Factura{nil}),
		facturacion.AsDocumentos([]*biodiesel.Factura{nil}),
		facturacion.AsDocumentos([]*boletoaereo.Factura{nil}),
		facturacion.AsDocumentos([]*combustiblesinsubvencion.Factura{nil}),
		facturacion.AsDocumentos([]*comercialexportacion.Factura{nil}),
		facturacion.AsDocumentos([]*compraventa.Factura{nil}),
		facturacion.AsDocumentos([]*compraventabonificaciones.Factura{nil}),
		facturacion.AsDocumentos([]*duttyfree.Factura{nil}),
		facturacion.AsDocumentos([]*entidadesfinancieras.Factura{nil}),
		facturacion.AsDocumentos([]*exportacionhidrocarburos.Factura{nil}),
		facturacion.AsDocumentos([]*exportacionminerales.Factura{nil}),
		facturacion.AsDocumentos([]*exportacionprecioventa.Factura{nil}),
		facturacion.AsDocumentos([]*exportacionservicios.Factura{nil}),
		facturacion.AsDocumentos([]*gnglsp.Factura{nil}),
		facturacion.AsDocumentos([]*gnv.Factura{nil}),
		facturacion.AsDocumentos([]*hidrocarburosiehd.Factura{nil}),
		facturacion.AsDocumentos([]*hidrocarburosnoiehd.Factura{nil}),
		facturacion.AsDocumentos([]*hospitales.Factura{nil}),
		facturacion.AsDocumentos([]*hospitaleszf.Factura{nil}),
		facturacion.AsDocumentos([]*hoteles.Factura{nil}),
		facturacion.AsDocumentos([]*juegosazar.Factura{nil}),
		facturacion.AsDocumentos([]*libreconsignacion.Factura{nil}),
		facturacion.AsDocumentos([]*lubricantes.Factura{nil}),
		facturacion.AsDocumentos([]*lubricantesiehd.Factura{nil}),
		facturacion.AsDocumentos([]*monedaextranjera.Factura{nil}),
		facturacion.AsDocumentos([]*prevalorada.Factura{nil}),
		facturacion.AsDocumentos([]*prevaloradasd.Factura{nil}),
		facturacion.AsDocumentos([]*sectoreducativo.Factura{nil}),
		facturacion.AsDocumentos([]*sectoreducativozf.Factura{nil}),
		facturacion.AsDocumentos([]*seguridadalimentaria.Factura{nil}),
		facturacion.AsDocumentos([]*seguros.Factura{nil}),
		facturacion.AsDocumentos([]*serviciosbasicos.Factura{nil}),
		facturacion.AsDocumentos([]*serviciosbasicoszf.Factura{nil}),
		facturacion.AsDocumentos([]*suministroenergia.Factura{nil}),
		facturacion.AsDocumentos([]*tasacero.Factura{nil}),
		facturacion.AsDocumentos([]*telecomunicaciones.Factura{nil}),
		facturacion.AsDocumentos([]*telecomunicacioneszf.Factura{nil}),
		facturacion.AsDocumentos([]*turismohospedaje.Factura{nil}),
		facturacion.AsDocumentos([]*ventamineralbcb.Factura{nil}),
		facturacion.AsDocumentos([]*ventaminerales.Factura{nil}),
		facturacion.AsDocumentos([]*vuelocharter.Factura{nil}),
		facturacion.AsDocumentos([]*zonafranca.Factura{nil}),
	}

	for i, docs := range documentos {
		if len(docs) != 1 {
			t.Fatalf("conversión %d: cantidad=%d", i, len(docs))
		}
	}
}

func TestWithFacturaAceptaDocumentosSeleccionados(t *testing.T) {
	casos := []struct {
		nombre string
		doc    facturacion.Documento
		sector int
	}{
		{"exportacion", &comercialexportacion.Factura{Cabecera: comercialexportacion.Cabecera{Incoterm: 1}}, 3},
		{"educativo", &sectoreducativo.Factura{}, 11},
		{"hoteles", &hoteles.Factura{}, 16},
		{"hospitales", &hospitales.Factura{}, 17},
		{"bonificaciones", &compraventabonificaciones.Factura{}, 35},
		{"vuelo charter", &vuelocharter.Factura{}, 41},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			solicitud, err := (facturacion.SolicitudRecepcionFactura{CodigoModalidad: facturacion.ModalidadComputarizada}).WithFactura(caso.doc, nil)
			if err != nil {
				t.Fatal(err)
			}
			if solicitud.CodigoDocumentoSector != caso.sector || len(solicitud.Archivo) == 0 {
				t.Fatalf("solicitud: sector=%d archivo=%d", solicitud.CodigoDocumentoSector, len(solicitud.Archivo))
			}
			xmlFactura, err := caso.doc.MarshalXMLModalidad(facturacion.ModalidadComputarizada)
			if err != nil {
				t.Fatal(err)
			}
			var root struct{ XMLName xml.Name }
			if err := xml.Unmarshal(xmlFactura, &root); err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(root.XMLName.Local, "Computarizada") {
				t.Fatalf("raíz XML para modalidad computarizada: %s", root.XMLName.Local)
			}
		})
	}
}
