// Package notaconciliacion proporciona documentos para el sector Nota de Conciliación en SIAT v3.
package notaconciliacion

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/shopspring/decimal"
)

// NotaCreditoDebito representa el documento fiscal para SIAT v3.
type NotaCreditoDebito struct {
	Modalidad int
	Cabecera  Cabecera
	Detalles  []Detalle
}

// Cabecera contiene los datos de cabecera del documento fiscal v3.
type Cabecera struct {
	NitEmisor                    int64
	RazonSocialEmisor            string
	Municipio                    string
	Telefono                     *string
	NumeroNotaConciliacion       int64
	Cuf                          string
	Cufd                         string
	CodigoSucursal               int
	Direccion                    string
	CodigoPuntoVenta             *int
	FechaEmision                 time.Time
	NombreRazonSocial            *string
	CodigoTipoDocumentoIdentidad int
	NumeroDocumento              string
	Complemento                  string
	CodigoCliente                string
	NumeroFactura                int64
	NumeroAutorizacionCuf        string
	CodigoControl                *string
	FechaEmisionFactura          time.Time
	TotalPrecioNetoOriginal      float64
	MontoTotalConciliado         float64
	IvaConciliado                float64
	MontoActualizacion           float64
	CreditoFiscalIva             float64
	DebitoFiscalIva              float64
	UfvFactura                   float64
	UfvConciliacion              float64
	CodigoExcepcion              *int
	Leyenda                      string
	Usuario                      string
	CodigoInterno                string
	CodigoDocumentoSector        int
}

// Detalle contiene los ítems del documento fiscal v3.
type Detalle struct {
	ActividadEconomica   string
	CodigoProductoSin    int64
	CodigoProducto       string
	Descripcion          string
	TotalPrecioNeto      float64
	TotalPrecioNetoFinal float64
	MontoConciliado      float64
}

func (d NotaCreditoDebito) Validate() error {
	if d.Modalidad != 0 && d.Modalidad != 1 && d.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", d.Modalidad)
	}
	if d.Cabecera.CodigoDocumentoSector != 0 && d.Cabecera.CodigoDocumentoSector != 29 {
		return fmt.Errorf("código de sector %d no corresponde a Nota de Conciliación (esperado 29)", d.Cabecera.CodigoDocumentoSector)
	}
	return nil
}

func round(v float64) float64 { return decimal.NewFromFloat(v).Round(2).InexactFloat64() }
func roundOptional(v *float64) *float64 {
	if v == nil {
		return nil
	}
	return new(round(*v))
}

func (d NotaCreditoDebito) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := d.Validate(); err != nil {
		return err
	}
	nombre := "notaElectronicaConciliacion"
	if d.Modalidad == 2 {
		nombre = "notaComputarizadaConciliacion"
	}
	cab := documents.CabeceraConciliacion{
		NitEmisor:                    d.Cabecera.NitEmisor,
		RazonSocialEmisor:            d.Cabecera.RazonSocialEmisor,
		Municipio:                    d.Cabecera.Municipio,
		Telefono:                     datatype.Nilable[string]{Value: d.Cabecera.Telefono},
		NumeroNotaConciliacion:       d.Cabecera.NumeroNotaConciliacion,
		Cuf:                          d.Cabecera.Cuf,
		Cufd:                         d.Cabecera.Cufd,
		CodigoSucursal:               d.Cabecera.CodigoSucursal,
		Direccion:                    d.Cabecera.Direccion,
		CodigoPuntoVenta:             datatype.Nilable[int]{Value: d.Cabecera.CodigoPuntoVenta},
		FechaEmision:                 datatype.NewTimeSiat(d.Cabecera.FechaEmision),
		NombreRazonSocial:            datatype.Nilable[string]{Value: d.Cabecera.NombreRazonSocial},
		CodigoTipoDocumentoIdentidad: d.Cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:              d.Cabecera.NumeroDocumento,
		Complemento:                  d.Cabecera.Complemento,
		CodigoCliente:                d.Cabecera.CodigoCliente,
		NumeroFactura:                d.Cabecera.NumeroFactura,
		NumeroAutorizacionCuf:        d.Cabecera.NumeroAutorizacionCuf,
		CodigoControl:                datatype.Nilable[string]{Value: d.Cabecera.CodigoControl},
		FechaEmisionFactura:          datatype.NewTimeSiat(d.Cabecera.FechaEmisionFactura),
		TotalPrecioNetoOriginal:      round(d.Cabecera.TotalPrecioNetoOriginal),
		MontoTotalConciliado:         round(d.Cabecera.MontoTotalConciliado),
		IvaConciliado:                round(d.Cabecera.IvaConciliado),
		MontoActualizacion:           round(d.Cabecera.MontoActualizacion),
		CreditoFiscalIva:             round(d.Cabecera.CreditoFiscalIva),
		DebitoFiscalIva:              round(d.Cabecera.DebitoFiscalIva),
		UfvFactura:                   round(d.Cabecera.UfvFactura),
		UfvConciliacion:              round(d.Cabecera.UfvConciliacion),
		CodigoExcepcion:              datatype.Nilable[int]{Value: d.Cabecera.CodigoExcepcion},
		Leyenda:                      d.Cabecera.Leyenda,
		Usuario:                      d.Cabecera.Usuario,
		CodigoInterno:                d.Cabecera.CodigoInterno,
		CodigoDocumentoSector:        d.Cabecera.CodigoDocumentoSector,
	}
	if cab.CodigoDocumentoSector == 0 {
		cab.CodigoDocumentoSector = 29
	}
	dets := make([]documents.DetalleConciliacion, 0, len(d.Detalles))
	for _, item := range d.Detalles {
		dets = append(dets, documents.DetalleConciliacion{
			ActividadEconomica:   item.ActividadEconomica,
			CodigoProductoSin:    item.CodigoProductoSin,
			CodigoProducto:       item.CodigoProducto,
			Descripcion:          item.Descripcion,
			TotalPrecioNeto:      round(item.TotalPrecioNeto),
			TotalPrecioNetoFinal: round(item.TotalPrecioNetoFinal),
			MontoConciliado:      round(item.MontoConciliado),
		})
	}
	doc := documents.NotaConciliacion{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cab,
		// Detalle omitido en favor de DetalleOriginal y DetalleConciliacion
	}
	return e.Encode(doc)
}
