// Package exportacionprecioventa proporciona documentos para el sector Exportación Precio Venta en SIAT v3.
package exportacionprecioventa

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
	Detalles  []Detalle
}

// Cabecera contiene los datos de cabecera del documento fiscal v3.
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
	AlcanceIncoterm                 string
	PuertoDestino                   string
	LugarDestino                    string
	CodigoPais                      int
	CodigoMetodoPago                int
	NumeroTarjeta                   *int64
	MontoAPagar                     float64
	CostosGastosNacionales          *string
	TotalGastosNacionalesFob        float64
	CostosGastosInternacionales     *string
	TotalGastosInternacionales      *float64
	PrecioValorBruto                *float64
	SubTotalNeto                    float64
	TotalPrecioNeto                 float64
	Iva                             float64
	CodigoMoneda                    int
	TipoCambio                      float64
	MontoTotalMoneda                float64
	NumeroDescripcionPaquetesBultos *string
	InformacionAdicional            *string
	CodigoExcepcion                 *int
	Leyenda                         string
	Usuario                         string
	CodigoInterno                   string
	CodigoDocumentoSector           int
}

// Detalle contiene los ítems del documento fiscal v3.
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

func (d Factura) Validate() error {
	if d.Modalidad != 0 && d.Modalidad != 1 && d.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", d.Modalidad)
	}
	if d.Cabecera.CodigoDocumentoSector != 0 && d.Cabecera.CodigoDocumentoSector != 45 {
		return fmt.Errorf("código de sector %d no corresponde a Exportación Precio Venta (esperado 45)", d.Cabecera.CodigoDocumentoSector)
	}
	if d.Cabecera.Incoterm < 1 || d.Cabecera.Incoterm > 11 {
		return fmt.Errorf("incoterm %d fuera del rango permitido [1, 11]", d.Cabecera.Incoterm)
	}
	return nil
}

func round(v float64) float64 {
	return decimal.NewFromFloat(v).Round(2).InexactFloat64()
}
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

func (d Factura) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := d.Validate(); err != nil {
		return err
	}
	nombre := "facturaElectronicaComercialExportacionPVenta"
	if d.Modalidad == 2 {
		nombre = "facturaComputarizadaComercialExportacionPVenta"
	}
	cab := documents.CabeceraComercialExportacionPVenta{
		NitEmisor:                       d.Cabecera.NitEmisor,
		RazonSocialEmisor:               d.Cabecera.RazonSocialEmisor,
		Municipio:                       d.Cabecera.Municipio,
		Telefono:                        datatype.Nilable[string]{Value: d.Cabecera.Telefono},
		NumeroFactura:                   d.Cabecera.NumeroFactura,
		Cuf:                             d.Cabecera.Cuf,
		Cufd:                            d.Cabecera.Cufd,
		CodigoSucursal:                  d.Cabecera.CodigoSucursal,
		Direccion:                       d.Cabecera.Direccion,
		CodigoPuntoVenta:                datatype.Nilable[int]{Value: d.Cabecera.CodigoPuntoVenta},
		FechaEmision:                    datatype.NewTimeSiat(d.Cabecera.FechaEmision),
		NombreRazonSocial:               datatype.Nilable[string]{Value: d.Cabecera.NombreRazonSocial},
		CodigoTipoDocumentoIdentidad:    d.Cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:                 d.Cabecera.NumeroDocumento,
		Complemento:                     d.Cabecera.Complemento,
		DireccionComprador:              d.Cabecera.DireccionComprador,
		CodigoCliente:                   d.Cabecera.CodigoCliente,
		Incoterm:                        d.Cabecera.Incoterm,
		AlcanceIncoterm:                 d.Cabecera.AlcanceIncoterm,
		PuertoDestino:                   d.Cabecera.PuertoDestino,
		LugarDestino:                    d.Cabecera.LugarDestino,
		CodigoPais:                      d.Cabecera.CodigoPais,
		CodigoMetodoPago:                d.Cabecera.CodigoMetodoPago,
		NumeroTarjeta:                   datatype.Nilable[int64]{Value: d.Cabecera.NumeroTarjeta},
		MontoAPagar:                     round(d.Cabecera.MontoAPagar),
		CostosGastosNacionales:          datatype.Nilable[string]{Value: d.Cabecera.CostosGastosNacionales},
		TotalGastosNacionalesFob:        round(d.Cabecera.TotalGastosNacionalesFob),
		CostosGastosInternacionales:     datatype.Nilable[string]{Value: d.Cabecera.CostosGastosInternacionales},
		TotalGastosInternacionales:      datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.TotalGastosInternacionales)},
		PrecioValorBruto:                datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.PrecioValorBruto)},
		SubTotalNeto:                    round(d.Cabecera.SubTotalNeto),
		TotalPrecioNeto:                 round(d.Cabecera.TotalPrecioNeto),
		Iva:                             round(d.Cabecera.Iva),
		CodigoMoneda:                    d.Cabecera.CodigoMoneda,
		TipoCambio:                      round5(d.Cabecera.TipoCambio),
		MontoTotalMoneda:                round(d.Cabecera.MontoTotalMoneda),
		NumeroDescripcionPaquetesBultos: datatype.Nilable[string]{Value: d.Cabecera.NumeroDescripcionPaquetesBultos},
		InformacionAdicional:            datatype.Nilable[string]{Value: d.Cabecera.InformacionAdicional},
		CodigoExcepcion:                 datatype.Nilable[int]{Value: d.Cabecera.CodigoExcepcion},
		Leyenda:                         d.Cabecera.Leyenda,
		Usuario:                         d.Cabecera.Usuario,
		CodigoInterno:                   d.Cabecera.CodigoInterno,
		CodigoDocumentoSector:           d.Cabecera.CodigoDocumentoSector,
	}
	if cab.CodigoDocumentoSector == 0 {
		cab.CodigoDocumentoSector = 45
	}
	dets := make([]documents.DetalleComercialExportacionPVenta, 0, len(d.Detalles))
	for _, item := range d.Detalles {
		dets = append(dets, documents.DetalleComercialExportacionPVenta{
			ActividadEconomica: item.ActividadEconomica,
			CodigoProductoSin:  item.CodigoProductoSin,
			CodigoProducto:     item.CodigoProducto,
			CodigoNandina:      item.CodigoNandina,
			Descripcion:        item.Descripcion,
			Cantidad:           round5(item.Cantidad),
			UnidadMedida:       item.UnidadMedida,
			PrecioNetoUnitario: round5(item.PrecioNetoUnitario),
			MontoDescuento:     datatype.Nilable[float64]{Value: roundOptional5(item.MontoDescuento)},
			SubTotal:           round5(item.SubTotal),
		})
	}
	doc := documents.FacturaComercialExportacionPVenta{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cab,
		Detalle:           dets,
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
	return 45
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
