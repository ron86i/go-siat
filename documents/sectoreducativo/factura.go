// Package sectoreducativo proporciona documentos del sector educativo para SIAT v3.
package sectoreducativo

import (
	"encoding/xml"
	"fmt"
	"math"
	"time"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/shopspring/decimal"
)

// Factura representa una factura del Sector Educativo para SIAT v3.
// Modalidad cero selecciona electrónica.
type Factura struct {
	Modalidad int
	Cabecera  Cabecera
	Detalles  []Detalle
}

// Cabecera contiene los datos de cabecera de la factura del sector educativo v3.
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
	FechaEmision                 time.Time
	NombreRazonSocial            *string
	CodigoTipoDocumentoIdentidad int
	NumeroDocumento              string
	Complemento                  string
	CodigoCliente                string
	NombreEstudiante             string
	PeriodoFacturado             string
	CodigoMetodoPago             int
	NumeroTarjeta                *int64
	SubTotalNeto                 float64
	TotalPrecioNeto              float64
	Iva                          float64
	MontoAPagar                  float64
	CodigoMoneda                 int
	TipoCambio                   float64
	MontoTotalMoneda             float64
	MontoGiftCard                *float64
	DescuentoAdicional           *float64
	CodigoExcepcion              *int
	Leyenda                      string
	Usuario                      string
	CodigoInterno                string
	CodigoDocumentoSector        int
}

// Detalle contiene un ítem de la factura del sector educativo v3.
type Detalle struct {
	ActividadEconomica string
	CodigoProductoSin  int64
	CodigoProducto     string
	Descripcion        string
	Cantidad           float64
	UnidadMedida       int
	PrecioNetoUnitario float64
	MontoDescuento     *float64
	SubTotal           float64
}

// Validate verifica las reglas de dominio para la factura v3.

// CodigoDocumentoSector devuelve el sector de la cabecera, usando 11 si está omitido.
func (f *Factura) CodigoDocumentoSector() int {
	if f == nil {
		return 0
	}
	if f.Cabecera.CodigoDocumentoSector == 0 {
		return 11
	}
	return f.Cabecera.CodigoDocumentoSector
}

// CodigoSucursal devuelve la sucursal de la cabecera.
func (f *Factura) CodigoSucursal() int {
	if f == nil {
		return 0
	}
	return f.Cabecera.CodigoSucursal
}

// CodigoPuntoVenta devuelve una copia del punto de venta de la cabecera.
func (f *Factura) CodigoPuntoVenta() *int {
	if f == nil || f.Cabecera.CodigoPuntoVenta == nil {
		return nil
	}
	return new(*f.Cabecera.CodigoPuntoVenta)
}

// Cufd devuelve el CUFD de la cabecera.
func (f *Factura) Cufd() string {
	if f == nil {
		return ""
	}
	return f.Cabecera.Cufd
}

// Protocolo devuelve 3 para SIAT v3.
func (f *Factura) Protocolo() int {
	if f == nil {
		return 0
	}
	return 3
}

// MarshalXMLModalidad serializa la factura con la modalidad solicitada.
func (f *Factura) MarshalXMLModalidad(modalidad int) ([]byte, error) {
	if f == nil {
		return nil, fmt.Errorf("la factura no puede ser nil")
	}
	if modalidad != 1 && modalidad != 2 {
		return nil, fmt.Errorf("modalidad inválida: %d", modalidad)
	}
	documento := *f
	documento.Modalidad = modalidad
	return xml.Marshal(documento)
}

func (f Factura) Validate() error {
	if f.Modalidad != 0 && f.Modalidad != 1 && f.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", f.Modalidad)
	}
	if f.Cabecera.CodigoDocumentoSector != 0 && f.Cabecera.CodigoDocumentoSector != 11 {
		return fmt.Errorf("código de sector %d no corresponde a Sector Educativo (esperado 11)", f.Cabecera.CodigoDocumentoSector)
	}
	valores := []float64{
		f.Cabecera.SubTotalNeto,
		f.Cabecera.TotalPrecioNeto,
		f.Cabecera.Iva,
		f.Cabecera.MontoAPagar,
		f.Cabecera.TipoCambio,
		f.Cabecera.MontoTotalMoneda,
	}
	if f.Cabecera.MontoGiftCard != nil {
		valores = append(valores, *f.Cabecera.MontoGiftCard)
	}
	if f.Cabecera.DescuentoAdicional != nil {
		valores = append(valores, *f.Cabecera.DescuentoAdicional)
	}
	for _, d := range f.Detalles {
		valores = append(valores, d.Cantidad, d.PrecioNetoUnitario, d.SubTotal)
		if d.MontoDescuento != nil {
			valores = append(valores, *d.MontoDescuento)
		}
	}
	for _, valor := range valores {
		if math.IsNaN(valor) || math.IsInf(valor, 0) {
			return fmt.Errorf("la factura contiene un valor decimal no finito")
		}
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

// MarshalXML serializa la factura según la especificación del sector educativo v3.
func (f Factura) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := f.Validate(); err != nil {
		return err
	}
	nombre := "facturaElectronicaSectorEducativo"
	if f.Modalidad == 2 {
		nombre = "facturaComputarizadaSectorEducativo"
	}
	cabecera := documents.CabeceraSectorEducativo{
		NitEmisor:                    f.Cabecera.NitEmisor,
		RazonSocialEmisor:            f.Cabecera.RazonSocialEmisor,
		Municipio:                    f.Cabecera.Municipio,
		Telefono:                     datatype.Nilable[string]{Value: f.Cabecera.Telefono},
		NumeroFactura:                f.Cabecera.NumeroFactura,
		Cuf:                          f.Cabecera.Cuf,
		Cufd:                         f.Cabecera.Cufd,
		CodigoSucursal:               f.Cabecera.CodigoSucursal,
		Direccion:                    f.Cabecera.Direccion,
		CodigoPuntoVenta:             datatype.Nilable[int]{Value: f.Cabecera.CodigoPuntoVenta},
		FechaEmision:                 datatype.NewTimeSiat(f.Cabecera.FechaEmision),
		NombreRazonSocial:            datatype.Nilable[string]{Value: f.Cabecera.NombreRazonSocial},
		CodigoTipoDocumentoIdentidad: f.Cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:              f.Cabecera.NumeroDocumento,
		Complemento:                  f.Cabecera.Complemento,
		CodigoCliente:                f.Cabecera.CodigoCliente,
		NombreEstudiante:             f.Cabecera.NombreEstudiante,
		PeriodoFacturado:             f.Cabecera.PeriodoFacturado,
		CodigoMetodoPago:             f.Cabecera.CodigoMetodoPago,
		NumeroTarjeta:                datatype.Nilable[int64]{Value: f.Cabecera.NumeroTarjeta},
		SubTotalNeto:                 round(f.Cabecera.SubTotalNeto),
		TotalPrecioNeto:              round(f.Cabecera.TotalPrecioNeto),
		Iva:                          round(f.Cabecera.Iva),
		MontoAPagar:                  round(f.Cabecera.MontoAPagar),
		CodigoMoneda:                 f.Cabecera.CodigoMoneda,
		TipoCambio:                   round(f.Cabecera.TipoCambio),
		MontoTotalMoneda:             round(f.Cabecera.MontoTotalMoneda),
		MontoGiftCard:                datatype.Nilable[float64]{Value: roundOptional(f.Cabecera.MontoGiftCard)},
		DescuentoAdicional:           datatype.Nilable[float64]{Value: roundOptional(f.Cabecera.DescuentoAdicional)},
		CodigoExcepcion:              datatype.Nilable[int]{Value: f.Cabecera.CodigoExcepcion},
		Leyenda:                      f.Cabecera.Leyenda,
		Usuario:                      f.Cabecera.Usuario,
		CodigoInterno:                f.Cabecera.CodigoInterno,
		CodigoDocumentoSector:        f.Cabecera.CodigoDocumentoSector,
	}
	if cabecera.CodigoDocumentoSector == 0 {
		cabecera.CodigoDocumentoSector = 11
	}
	detalles := make([]documents.DetalleSectorEducativo, 0, len(f.Detalles))
	for _, d := range f.Detalles {
		detalles = append(detalles, documents.DetalleSectorEducativo{
			ActividadEconomica: d.ActividadEconomica,
			CodigoProductoSin:  d.CodigoProductoSin,
			CodigoProducto:     d.CodigoProducto,
			Descripcion:        d.Descripcion,
			Cantidad:           round(d.Cantidad),
			UnidadMedida:       d.UnidadMedida,
			PrecioNetoUnitario: round(d.PrecioNetoUnitario),
			MontoDescuento:     datatype.Nilable[float64]{Value: roundOptional(d.MontoDescuento)},
			SubTotal:           round(d.SubTotal),
		})
	}
	return e.Encode(documents.FacturaSectorEducativo{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cabecera,
		Detalle:           detalles,
	})
}
