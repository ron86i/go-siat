// Package comercialexportacion proporciona documentos de exportación comercial para SIAT v3.
package comercialexportacion

import (
	"encoding/xml"
	"fmt"
	"math"
	"time"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/shopspring/decimal"
)

// Factura representa una factura Comercial de Exportación para SIAT v3.
// Modalidad cero selecciona electrónica.
type Factura struct {
	Modalidad int
	Cabecera  Cabecera
	Detalles  []Detalle
}

// Cabecera contiene los datos de cabecera de la factura de exportación v3.
type Cabecera struct {
	NitEmisor                       int64
	RazonSocialEmisor               string
	Municipio                       string
	Telefono                        *string
	NumeroFactura                   int64
	Cuf                             string
	Cufd                            string
	CodigoSucursal                  int
	Direccion                       string
	CodigoPuntoVenta                *int
	FechaEmision                    time.Time
	NombreRazonSocial               *string
	CodigoTipoDocumentoIdentidad    int
	NumeroDocumento                 string
	Complemento                     string
	DireccionComprador              string
	CodigoCliente                   string
	Incoterm                        int
	IncotermAlcance                 string
	IncotermDetalle                 string
	PuertoDestino                   string
	LugarDestino                    string
	CodigoPais                      int
	CodigoMetodoPago                int
	NumeroTarjeta                   *int64
	TotalPrecioNeto                 float64
	MontoAPagar                     float64
	CostosGastosNacionales          *string
	TotalGastosNacionalesFob        float64
	CostosGastosInternacionales     *string
	TotalGastosInternacionales      *float64
	SubTotalNeto                    float64
	Iva                             float64
	CodigoMoneda                    int
	TipoCambio                      float64
	MontoTotalMoneda                float64
	NumeroDescripcionPaquetesBultos *string
	InformacionAdicional            *string
	DescuentoAdicional              *float64
	CodigoExcepcion                 *int
	Leyenda                         string
	Usuario                         string
	CodigoInterno                   string
	CodigoDocumentoSector           int
}

// Detalle contiene un ítem de la factura de exportación v3.
type Detalle struct {
	ActividadEconomica string
	CodigoProductoSin  int64
	CodigoProducto     string
	CodigoNandina      string
	Descripcion        string
	Cantidad           float64
	UnidadMedida       int
	PrecioNetoUnitario float64
	MontoDescuento     *float64
	SubTotal           float64
}

// Validate verifica las restricciones de dominio de la factura v3.

// CodigoDocumentoSector devuelve el sector de la cabecera, usando 3 si está omitido.
func (f *Factura) CodigoDocumentoSector() int {
	if f == nil {
		return 0
	}
	if f.Cabecera.CodigoDocumentoSector == 0 {
		return 3
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
	if f.Cabecera.CodigoDocumentoSector != 0 && f.Cabecera.CodigoDocumentoSector != 3 {
		return fmt.Errorf("código de sector %d no corresponde a Comercial de Exportación (esperado 3)", f.Cabecera.CodigoDocumentoSector)
	}
	if f.Cabecera.Incoterm < 1 || f.Cabecera.Incoterm > 11 {
		return fmt.Errorf("incoterm %d fuera del rango permitido [1, 11]", f.Cabecera.Incoterm)
	}
	valores := []float64{
		f.Cabecera.TotalPrecioNeto,
		f.Cabecera.MontoAPagar,
		f.Cabecera.TotalGastosNacionalesFob,
		f.Cabecera.SubTotalNeto,
		f.Cabecera.Iva,
		f.Cabecera.TipoCambio,
		f.Cabecera.MontoTotalMoneda,
	}
	if f.Cabecera.TotalGastosInternacionales != nil {
		valores = append(valores, *f.Cabecera.TotalGastosInternacionales)
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

func round(v float64) float64  { return decimal.NewFromFloat(v).Round(2).InexactFloat64() }
func round5(v float64) float64 { return decimal.NewFromFloat(v).Round(5).InexactFloat64() }
func roundOptional5(v *float64) *float64 {
	if v == nil {
		return nil
	}
	return new(round5(*v))
}
func roundOptional(v *float64) *float64 {
	if v == nil {
		return nil
	}
	return new(round(*v))
}

// MarshalXML serializa la factura según la especificación de comercial de exportación v3.
func (f Factura) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := f.Validate(); err != nil {
		return err
	}
	nombre := "facturaElectronicaComercialExportacion"
	if f.Modalidad == 2 {
		nombre = "facturaComputarizadaComercialExportacion"
	}
	cabecera := documents.CabeceraComercialExportacion{
		NitEmisor:                       f.Cabecera.NitEmisor,
		RazonSocialEmisor:               f.Cabecera.RazonSocialEmisor,
		Municipio:                       f.Cabecera.Municipio,
		Telefono:                        datatype.Nilable[string]{Value: f.Cabecera.Telefono},
		NumeroFactura:                   f.Cabecera.NumeroFactura,
		Cuf:                             f.Cabecera.Cuf,
		Cufd:                            f.Cabecera.Cufd,
		CodigoSucursal:                  f.Cabecera.CodigoSucursal,
		Direccion:                       f.Cabecera.Direccion,
		CodigoPuntoVenta:                datatype.Nilable[int]{Value: f.Cabecera.CodigoPuntoVenta},
		FechaEmision:                    datatype.NewTimeSiat(f.Cabecera.FechaEmision),
		NombreRazonSocial:               datatype.Nilable[string]{Value: f.Cabecera.NombreRazonSocial},
		CodigoTipoDocumentoIdentidad:    f.Cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:                 f.Cabecera.NumeroDocumento,
		Complemento:                     f.Cabecera.Complemento,
		DireccionComprador:              f.Cabecera.DireccionComprador,
		CodigoCliente:                   f.Cabecera.CodigoCliente,
		Incoterm:                        f.Cabecera.Incoterm,
		IncotermAlcance:                 f.Cabecera.IncotermAlcance,
		IncotermDetalle:                 f.Cabecera.IncotermDetalle,
		PuertoDestino:                   f.Cabecera.PuertoDestino,
		LugarDestino:                    f.Cabecera.LugarDestino,
		CodigoPais:                      f.Cabecera.CodigoPais,
		CodigoMetodoPago:                f.Cabecera.CodigoMetodoPago,
		NumeroTarjeta:                   datatype.Nilable[int64]{Value: f.Cabecera.NumeroTarjeta},
		TotalPrecioNeto:                 round(f.Cabecera.TotalPrecioNeto),
		MontoAPagar:                     round(f.Cabecera.MontoAPagar),
		CostosGastosNacionales:          datatype.Nilable[string]{Value: f.Cabecera.CostosGastosNacionales},
		TotalGastosNacionalesFob:        round(f.Cabecera.TotalGastosNacionalesFob),
		CostosGastosInternacionales:     datatype.Nilable[string]{Value: f.Cabecera.CostosGastosInternacionales},
		TotalGastosInternacionales:      datatype.Nilable[float64]{Value: roundOptional(f.Cabecera.TotalGastosInternacionales)},
		SubTotalNeto:                    round(f.Cabecera.SubTotalNeto),
		Iva:                             round(f.Cabecera.Iva),
		CodigoMoneda:                    f.Cabecera.CodigoMoneda,
		TipoCambio:                      round5(f.Cabecera.TipoCambio),
		MontoTotalMoneda:                round(f.Cabecera.MontoTotalMoneda),
		NumeroDescripcionPaquetesBultos: datatype.Nilable[string]{Value: f.Cabecera.NumeroDescripcionPaquetesBultos},
		InformacionAdicional:            datatype.Nilable[string]{Value: f.Cabecera.InformacionAdicional},
		DescuentoAdicional:              datatype.Nilable[float64]{Value: roundOptional(f.Cabecera.DescuentoAdicional)},
		CodigoExcepcion:                 datatype.Nilable[int]{Value: f.Cabecera.CodigoExcepcion},
		Leyenda:                         f.Cabecera.Leyenda,
		Usuario:                         f.Cabecera.Usuario,
		CodigoInterno:                   f.Cabecera.CodigoInterno,
		CodigoDocumentoSector:           f.Cabecera.CodigoDocumentoSector,
	}
	if cabecera.CodigoDocumentoSector == 0 {
		cabecera.CodigoDocumentoSector = 3
	}
	detalles := make([]documents.DetalleComercialExportacion, 0, len(f.Detalles))
	for _, d := range f.Detalles {
		detalles = append(detalles, documents.DetalleComercialExportacion{
			ActividadEconomica: d.ActividadEconomica,
			CodigoProductoSin:  d.CodigoProductoSin,
			CodigoProducto:     d.CodigoProducto,
			CodigoNandina:      d.CodigoNandina,
			Descripcion:        d.Descripcion,
			Cantidad:           round5(d.Cantidad),
			UnidadMedida:       d.UnidadMedida,
			PrecioNetoUnitario: round5(d.PrecioNetoUnitario),
			MontoDescuento:     datatype.Nilable[float64]{Value: roundOptional5(d.MontoDescuento)},
			SubTotal:           round5(d.SubTotal),
		})
	}
	return e.Encode(documents.FacturaComercialExportacion{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cabecera,
		Detalle:           detalles,
	})
}
