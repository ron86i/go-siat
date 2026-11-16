// Package ventaminerales proporciona documentos para el sector Venta de Minerales en SIAT v3.
package ventaminerales

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
	DireccionComprador           string
	CodigoTipoDocumentoIdentidad int
	NumeroDocumento              string
	Complemento                  *string
	Nim                          *string
	DescripcionProducto          string
	Localidad                    string
	DepartamentoOrigen           *string
	MunicipioOrigen              *string
	CodigoMoneda                 int
	TipoCambio                   float64
	NumeroLote                   string
	NumeroContrato               string
	KilosNetosHumedos            float64
	HumedadPorcentaje            *float64
	HumedadValor                 *float64
	MermaPorcentaje              *float64
	MermaValor                   *float64
	KilosNetosSecos              float64
	CodigoMetodoPago             int
	NumeroTarjeta                *int64
	SubTotalNeto                 float64
	TotalPrecioNeto              float64
	Iva                          float64
	MontoAPagar                  float64
	MontoTotalMoneda             float64
	GastosRealizacion            float64
	LiquidacionPreliminar        *float64
	OtrosDatos                   *string
	MontoGiftCard                *float64
	DescuentoAdicional           *float64
	CodigoExcepcion              *int
	Leyenda                      string
	Usuario                      string
	CodigoInterno                string
	CodigoDocumentoSector        int
}

// Detalle contiene los ítems del documento fiscal v3.
type Detalle struct {
	ActividadEconomica     string
	CodigoProductoSin      int64
	CodigoProducto         string
	CodigoNandina          string
	Descripcion            string
	DescripcionLeyes       string
	CantidadExtraccion     float64
	Cantidad               float64
	UnidadMedidaExtraccion int
	UnidadMedida           int
	PrecioNetoUnitario     float64
	MontoDescuento         float64
	SubTotal               float64
}

func (d Factura) Validate() error {
	if d.Modalidad != 0 && d.Modalidad != 1 && d.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", d.Modalidad)
	}
	if d.Cabecera.CodigoDocumentoSector != 0 && d.Cabecera.CodigoDocumentoSector != 21 {
		return fmt.Errorf("código de sector %d no corresponde a Venta de Minerales (esperado 21)", d.Cabecera.CodigoDocumentoSector)
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
	nombre := "facturaElectronicaVentaMineral"
	if d.Modalidad == 2 {
		nombre = "facturaComputarizadaVentaMineral"
	}
	cab := documents.CabeceraVentaMineral{
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
		FechaEmision:                 datatype.NewTimeSiat(d.Cabecera.FechaEmision),
		NombreRazonSocial:            datatype.Nilable[string]{Value: d.Cabecera.NombreRazonSocial},
		DireccionComprador:           d.Cabecera.DireccionComprador,
		CodigoTipoDocumentoIdentidad: d.Cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:              d.Cabecera.NumeroDocumento,
		Complemento:                  datatype.Nilable[string]{Value: d.Cabecera.Complemento},
		Nim:                          datatype.Nilable[string]{Value: d.Cabecera.Nim},
		DescripcionProducto:          d.Cabecera.DescripcionProducto,
		Localidad:                    d.Cabecera.Localidad,
		DepartamentoOrigen:           datatype.Nilable[string]{Value: d.Cabecera.DepartamentoOrigen},
		MunicipioOrigen:              datatype.Nilable[string]{Value: d.Cabecera.MunicipioOrigen},
		CodigoMoneda:                 d.Cabecera.CodigoMoneda,
		TipoCambio:                   round(d.Cabecera.TipoCambio),
		NumeroLote:                   d.Cabecera.NumeroLote,
		NumeroContrato:               d.Cabecera.NumeroContrato,
		KilosNetosHumedos:            round(d.Cabecera.KilosNetosHumedos),
		HumedadPorcentaje:            datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.HumedadPorcentaje)},
		HumedadValor:                 datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.HumedadValor)},
		MermaPorcentaje:              datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.MermaPorcentaje)},
		MermaValor:                   datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.MermaValor)},
		KilosNetosSecos:              round(d.Cabecera.KilosNetosSecos),
		CodigoMetodoPago:             d.Cabecera.CodigoMetodoPago,
		NumeroTarjeta:                datatype.Nilable[int64]{Value: d.Cabecera.NumeroTarjeta},
		SubTotalNeto:                 round(d.Cabecera.SubTotalNeto),
		TotalPrecioNeto:              round(d.Cabecera.TotalPrecioNeto),
		Iva:                          round(d.Cabecera.Iva),
		MontoAPagar:                  round(d.Cabecera.MontoAPagar),
		MontoTotalMoneda:             round(d.Cabecera.MontoTotalMoneda),
		GastosRealizacion:            round(d.Cabecera.GastosRealizacion),
		LiquidacionPreliminar:        datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.LiquidacionPreliminar)},
		OtrosDatos:                   datatype.Nilable[string]{Value: d.Cabecera.OtrosDatos},
		MontoGiftCard:                datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.MontoGiftCard)},
		DescuentoAdicional:           datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.DescuentoAdicional)},
		CodigoExcepcion:              datatype.Nilable[int]{Value: d.Cabecera.CodigoExcepcion},
		Leyenda:                      d.Cabecera.Leyenda,
		Usuario:                      d.Cabecera.Usuario,
		CodigoInterno:                d.Cabecera.CodigoInterno,
		CodigoDocumentoSector:        d.Cabecera.CodigoDocumentoSector,
	}
	if cab.CodigoDocumentoSector == 0 {
		cab.CodigoDocumentoSector = 21
	}
	dets := make([]documents.DetalleVentaMineral, 0, len(d.Detalles))
	for _, item := range d.Detalles {
		dets = append(dets, documents.DetalleVentaMineral{
			ActividadEconomica:     item.ActividadEconomica,
			CodigoProductoSin:      item.CodigoProductoSin,
			CodigoProducto:         item.CodigoProducto,
			CodigoNandina:          item.CodigoNandina,
			Descripcion:            item.Descripcion,
			DescripcionLeyes:       item.DescripcionLeyes,
			CantidadExtraccion:     round(item.CantidadExtraccion),
			Cantidad:               round(item.Cantidad),
			UnidadMedidaExtraccion: item.UnidadMedidaExtraccion,
			UnidadMedida:           item.UnidadMedida,
			PrecioNetoUnitario:     round(item.PrecioNetoUnitario),
			MontoDescuento:         round(item.MontoDescuento),
			SubTotal:               round(item.SubTotal),
		})
	}
	doc := documents.FacturaVentaMineral{
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
	return 21
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
