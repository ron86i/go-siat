// Package serviciosbasicoszf proporciona documentos para el sector Servicios Básicos Zona Franca en SIAT v3.
package serviciosbasicoszf

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
	RazonSocialEmisor            string
	Municipio                    string
	Telefono                     *string
	NumeroFactura                int64
	Cuf                          string
	Cufd                         string
	CodigoSucursal               int
	Direccion                    string
	CodigoPuntoVenta             *int
	Mes                          *string
	Gestion                      *int
	Ciudad                       *string
	Zona                         *string
	NumeroMedidor                string
	FechaEmision                 time.Time
	NombreRazonSocial            *string
	DomicilioCliente             *string
	CodigoTipoDocumentoIdentidad int
	NumeroDocumento              string
	Complemento                  string
	CodigoCliente                string
	CodigoMetodoPago             int
	NumeroTarjeta                *int64
	SubTotalNeto                 float64
	TotalPrecioNeto              float64
	Iva                          float64
	MontoAPagar                  float64
	ConsumoPeriodo               *float64
	BeneficiarioLey1886          *int64
	ComplementoBeneficiario      string
	MontoDescuentoLey1886        *float64
	MontoDescuentoSocial         *float64
	TasaAseo                     *float64
	TasaAlumbrado                *float64
	AjusteNoSujetoIva            *float64
	DetalleAjusteNoSujetoIva     *string
	OtrosPagosNoSujetoIva        *float64
	DetalleOtrosPagosNoSujetoIva *string
	OtrasTasas                   *float64
	CodigoMoneda                 int
	TipoCambio                   float64
	MontoTotalMoneda             float64
	DescuentoAdicional           *float64
	CodigoExcepcion              *int
	Leyenda                      string
	Usuario                      string
	CodigoInterno                string
	CodigoDocumentoSector        int
}

func (d Factura) Validate() error {
	if d.Modalidad != 0 && d.Modalidad != 1 && d.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", d.Modalidad)
	}
	if d.Cabecera.CodigoDocumentoSector != 0 && d.Cabecera.CodigoDocumentoSector != 40 {
		return fmt.Errorf("código de sector %d no corresponde a Servicios Básicos Zona Franca (esperado 40)", d.Cabecera.CodigoDocumentoSector)
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
	nombre := "facturaElectronicaServicioBasicoZf"
	if d.Modalidad == 2 {
		nombre = "facturaComputarizadaServicioBasicoZf"
	}
	cab := documents.CabeceraServicioBasicoZF{
		NitEmisor:                    d.Cabecera.NitEmisor,
		RazonSocialEmisor:            d.Cabecera.RazonSocialEmisor,
		Municipio:                    d.Cabecera.Municipio,
		Telefono:                     datatype.Nilable[string]{Value: d.Cabecera.Telefono},
		NumeroFactura:                d.Cabecera.NumeroFactura,
		Cuf:                          d.Cabecera.Cuf,
		Cufd:                         d.Cabecera.Cufd,
		CodigoSucursal:               d.Cabecera.CodigoSucursal,
		Direccion:                    d.Cabecera.Direccion,
		CodigoPuntoVenta:             datatype.Nilable[int]{Value: d.Cabecera.CodigoPuntoVenta},
		Mes:                          datatype.Nilable[string]{Value: d.Cabecera.Mes},
		Gestion:                      datatype.Nilable[int]{Value: d.Cabecera.Gestion},
		Ciudad:                       datatype.Nilable[string]{Value: d.Cabecera.Ciudad},
		Zona:                         datatype.Nilable[string]{Value: d.Cabecera.Zona},
		NumeroMedidor:                d.Cabecera.NumeroMedidor,
		FechaEmision:                 datatype.NewTimeSiat(d.Cabecera.FechaEmision),
		NombreRazonSocial:            datatype.Nilable[string]{Value: d.Cabecera.NombreRazonSocial},
		DomicilioCliente:             datatype.Nilable[string]{Value: d.Cabecera.DomicilioCliente},
		CodigoTipoDocumentoIdentidad: d.Cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:              d.Cabecera.NumeroDocumento,
		Complemento:                  d.Cabecera.Complemento,
		CodigoCliente:                d.Cabecera.CodigoCliente,
		CodigoMetodoPago:             d.Cabecera.CodigoMetodoPago,
		NumeroTarjeta:                datatype.Nilable[int64]{Value: d.Cabecera.NumeroTarjeta},
		SubTotalNeto:                 round(d.Cabecera.SubTotalNeto),
		TotalPrecioNeto:              round(d.Cabecera.TotalPrecioNeto),
		Iva:                          round(d.Cabecera.Iva),
		MontoAPagar:                  round(d.Cabecera.MontoAPagar),
		ConsumoPeriodo:               datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.ConsumoPeriodo)},
		BeneficiarioLey1886:          datatype.Nilable[int64]{Value: d.Cabecera.BeneficiarioLey1886},
		ComplementoBeneficiario:      d.Cabecera.ComplementoBeneficiario,
		MontoDescuentoLey1886:        datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.MontoDescuentoLey1886)},
		MontoDescuentoSocial:         datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.MontoDescuentoSocial)},
		TasaAseo:                     datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.TasaAseo)},
		TasaAlumbrado:                datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.TasaAlumbrado)},
		AjusteNoSujetoIva:            datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.AjusteNoSujetoIva)},
		DetalleAjusteNoSujetoIva:     datatype.Nilable[string]{Value: d.Cabecera.DetalleAjusteNoSujetoIva},
		OtrosPagosNoSujetoIva:        datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.OtrosPagosNoSujetoIva)},
		DetalleOtrosPagosNoSujetoIva: datatype.Nilable[string]{Value: d.Cabecera.DetalleOtrosPagosNoSujetoIva},
		OtrasTasas:                   datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.OtrasTasas)},
		CodigoMoneda:                 d.Cabecera.CodigoMoneda,
		TipoCambio:                   round(d.Cabecera.TipoCambio),
		MontoTotalMoneda:             round(d.Cabecera.MontoTotalMoneda),
		DescuentoAdicional:           datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.DescuentoAdicional)},
		CodigoExcepcion:              datatype.Nilable[int]{Value: d.Cabecera.CodigoExcepcion},
		Leyenda:                      d.Cabecera.Leyenda,
		Usuario:                      d.Cabecera.Usuario,
		CodigoInterno:                d.Cabecera.CodigoInterno,
		CodigoDocumentoSector:        d.Cabecera.CodigoDocumentoSector,
	}
	if cab.CodigoDocumentoSector == 0 {
		cab.CodigoDocumentoSector = 40
	}
	doc := documents.FacturaServicioBasicoZF{
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
	return 40
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
