// Package boletoaereo proporciona documentos para el sector Boleto Aéreo en SIAT v3.
package boletoaereo

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/shopspring/decimal"
)

// Factura representa el documento fiscal para SIAT v3.
type Factura struct {
	Modalidad int
	Cabecera  Cabecera
}

// Cabecera contiene los datos de cabecera del documento fiscal v3.
type Cabecera struct {
	NitEmisor                    int64
	NumeroFactura                int64
	Cufd                         string
	Cuf                          string
	CodigoSucursal               int
	CodigoPuntoVenta             *int
	Direccion                    string
	FechaEmision                 time.Time
	RazonSocialEmisor            string
	NombreRazonSocial            *string
	NumeroDocumento              string
	CodigoTipoDocumentoIdentidad int
	NombrePasajero               string
	NumeroDocumentoPasajero      *string
	CodigoIataLineaAerea         int
	CodigoIataAgenteViajes       *string
	NitAgenteViajes              *int64
	CodigoOrigenServicio         string
	CodigoMoneda                 int
	TipoCambio                   float64
	MontoTarifa                  float64
	TotalPrecioNeto              float64
	Iva                          float64
	MontoAPagar                  float64
	MontoTotalMoneda             float64
	CodigoMetodoPago             int
	CodigoTipoTransaccion        string
	Usuario                      string
	Leyenda                      string
	CodigoInterno                string
	CodigoDocumentoSector        int
}

func (d Factura) Validate() error {
	if d.Modalidad != 0 && d.Modalidad != 1 && d.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", d.Modalidad)
	}
	if d.Cabecera.CodigoDocumentoSector != 0 && d.Cabecera.CodigoDocumentoSector != 30 {
		return fmt.Errorf("código de sector %d no corresponde a Boleto Aéreo (esperado 30)", d.Cabecera.CodigoDocumentoSector)
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

func (d Factura) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := d.Validate(); err != nil {
		return err
	}
	nombre := "facturaElectronicaBoletoAereo"
	if d.Modalidad == 2 {
		nombre = "facturaComputarizadaBoletoAereo"
	}
	cab := documents.CabeceraBoletoAereo{
		NitEmisor:                    d.Cabecera.NitEmisor,
		NumeroFactura:                d.Cabecera.NumeroFactura,
		Cufd:                         d.Cabecera.Cufd,
		Cuf:                          d.Cabecera.Cuf,
		CodigoSucursal:               d.Cabecera.CodigoSucursal,
		CodigoPuntoVenta:             datatype.Nilable[int]{Value: d.Cabecera.CodigoPuntoVenta},
		Direccion:                    d.Cabecera.Direccion,
		FechaEmision:                 datatype.NewTimeSiat(d.Cabecera.FechaEmision),
		RazonSocialEmisor:            d.Cabecera.RazonSocialEmisor,
		NombreRazonSocial:            datatype.Nilable[string]{Value: d.Cabecera.NombreRazonSocial},
		NumeroDocumento:              d.Cabecera.NumeroDocumento,
		CodigoTipoDocumentoIdentidad: d.Cabecera.CodigoTipoDocumentoIdentidad,
		NombrePasajero:               d.Cabecera.NombrePasajero,
		NumeroDocumentoPasajero:      datatype.Nilable[string]{Value: d.Cabecera.NumeroDocumentoPasajero},
		CodigoIataLineaAerea:         d.Cabecera.CodigoIataLineaAerea,
		CodigoIataAgenteViajes:       datatype.Nilable[string]{Value: d.Cabecera.CodigoIataAgenteViajes},
		NitAgenteViajes:              datatype.Nilable[int64]{Value: d.Cabecera.NitAgenteViajes},
		CodigoOrigenServicio:         d.Cabecera.CodigoOrigenServicio,
		CodigoMoneda:                 d.Cabecera.CodigoMoneda,
		TipoCambio:                   round(d.Cabecera.TipoCambio),
		MontoTarifa:                  round(d.Cabecera.MontoTarifa),
		TotalPrecioNeto:              round(d.Cabecera.TotalPrecioNeto),
		Iva:                          round(d.Cabecera.Iva),
		MontoAPagar:                  round(d.Cabecera.MontoAPagar),
		MontoTotalMoneda:             round(d.Cabecera.MontoTotalMoneda),
		CodigoMetodoPago:             d.Cabecera.CodigoMetodoPago,
		CodigoTipoTransaccion:        d.Cabecera.CodigoTipoTransaccion,
		Usuario:                      d.Cabecera.Usuario,
		Leyenda:                      d.Cabecera.Leyenda,
		CodigoInterno:                d.Cabecera.CodigoInterno,
		CodigoDocumentoSector:        d.Cabecera.CodigoDocumentoSector,
	}
	if cab.CodigoDocumentoSector == 0 {
		cab.CodigoDocumentoSector = 30
	}
	doc := documents.FacturaBoletoAereo{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cab,
	}
	return e.Encode(doc)
}

// CodigoDocumentoSector devuelve el sector fijo del documento si no fue indicado en la cabecera.
func (d *Factura) CodigoDocumentoSector() int {
	if d == nil {
		return 0
	}
	if d.Cabecera.CodigoDocumentoSector != 0 {
		return d.Cabecera.CodigoDocumentoSector
	}
	return 30
}

// CodigoSucursal devuelve la sucursal consignada en la factura.
func (d *Factura) CodigoSucursal() int {
	if d == nil {
		return 0
	}
	return d.Cabecera.CodigoSucursal
}

// CodigoPuntoVenta devuelve una copia del punto de venta consignado en la factura.
func (d *Factura) CodigoPuntoVenta() *int {
	if d == nil || d.Cabecera.CodigoPuntoVenta == nil {
		return nil
	}
	return new(*d.Cabecera.CodigoPuntoVenta)
}

// Cufd devuelve el CUFD consignado en la factura.
func (d *Factura) Cufd() string {
	if d == nil {
		return ""
	}
	return d.Cabecera.Cufd
}

// Protocolo devuelve la versión SIAT del modelo.
func (d *Factura) Protocolo() int {
	if d == nil {
		return 0
	}
	return 3
}

// MarshalXMLModalidad serializa la factura con la modalidad indicada.
func (d *Factura) MarshalXMLModalidad(modalidad int) ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("la factura no puede ser nil")
	}
	factura := *d
	factura.Modalidad = modalidad
	return xml.Marshal(&factura)
}
