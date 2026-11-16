// Package notacreditodebito proporciona documentos de nota de crédito-débito para SIAT v3.
package notacreditodebito

import (
	"encoding/xml"
	"fmt"
	"math"
	"time"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/shopspring/decimal"
)

// NotaCreditoDebito representa un documento de ajuste Nota Crédito-Débito para SIAT v3.
// Modalidad cero selecciona electrónica.
type NotaCreditoDebito struct {
	Modalidad int
	Cabecera  Cabecera
	Detalles  []Detalle
}

// Cabecera contiene los datos de cabecera de la nota crédito-débito v3.
type Cabecera struct {
	NitEmisor                    int64
	RazonSocialEmisor            string
	Municipio                    string
	Telefono                     *string
	NumeroNotaCreditoDebito      int64
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
	FechaEmisionFactura          time.Time
	SubTotalNetoOriginal         float64
	TotalPrecioNetoOriginal      float64
	IvaOriginal                  float64
	MontoAPagarOriginal          float64
	SubTotalNetoDevuelto         float64
	TotalPrecioNetoDevuelto      float64
	IvaDevuelto                  float64
	MontoAPagarDevuelto          float64
	DescuentoAdicionalOriginal   float64
	MontoDescuentoDevuelto       float64
	CodigoExcepcion              *int
	Leyenda                      string
	Usuario                      string
	CodigoInterno                string
	CodigoDocumentoSector        int
}

// Detalle contiene un ítem de la nota crédito-débito v3.
type Detalle struct {
	NroItem                  int
	ActividadEconomica       string
	CodigoProductoSin        int64
	CodigoProducto           string
	Descripcion              string
	Cantidad                 float64
	UnidadMedida             int
	PrecioNetoUnitario       float64
	MontoDescuento           *float64
	SubTotal                 float64
	CodigoDetalleTransaccion int
}

// Validate verifica las restricciones de dominio de la nota v3.
func (n NotaCreditoDebito) Validate() error {
	if n.Modalidad != 0 && n.Modalidad != 1 && n.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", n.Modalidad)
	}
	if n.Cabecera.CodigoDocumentoSector != 0 && n.Cabecera.CodigoDocumentoSector != 24 {
		return fmt.Errorf("código de sector %d no corresponde a Nota de Crédito-Débito (esperado 24)", n.Cabecera.CodigoDocumentoSector)
	}
	valores := []float64{
		n.Cabecera.SubTotalNetoOriginal,
		n.Cabecera.TotalPrecioNetoOriginal,
		n.Cabecera.IvaOriginal,
		n.Cabecera.MontoAPagarOriginal,
		n.Cabecera.SubTotalNetoDevuelto,
		n.Cabecera.TotalPrecioNetoDevuelto,
		n.Cabecera.IvaDevuelto,
		n.Cabecera.MontoAPagarDevuelto,
		n.Cabecera.DescuentoAdicionalOriginal,
		n.Cabecera.MontoDescuentoDevuelto,
	}
	for _, d := range n.Detalles {
		valores = append(valores, d.Cantidad, d.PrecioNetoUnitario, d.SubTotal)
		if d.MontoDescuento != nil {
			valores = append(valores, *d.MontoDescuento)
		}
	}
	for _, valor := range valores {
		if math.IsNaN(valor) || math.IsInf(valor, 0) {
			return fmt.Errorf("el documento contiene un valor decimal no finito")
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

// MarshalXML serializa la nota según la especificación de nota crédito-débito v3.
func (n NotaCreditoDebito) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := n.Validate(); err != nil {
		return err
	}
	nombre := "notaFiscalElectronicaCreditoDebito"
	if n.Modalidad == 2 {
		nombre = "notaFiscalComputarizadaCreditoDebito"
	}
	cabecera := documents.CabeceraNotaCreditoDebito{
		NitEmisor:                    n.Cabecera.NitEmisor,
		RazonSocialEmisor:            n.Cabecera.RazonSocialEmisor,
		Municipio:                    n.Cabecera.Municipio,
		Telefono:                     datatype.Nilable[string]{Value: n.Cabecera.Telefono},
		NumeroNotaCreditoDebito:      n.Cabecera.NumeroNotaCreditoDebito,
		Cuf:                          n.Cabecera.Cuf,
		Cufd:                         n.Cabecera.Cufd,
		CodigoSucursal:               n.Cabecera.CodigoSucursal,
		Direccion:                    n.Cabecera.Direccion,
		CodigoPuntoVenta:             datatype.Nilable[int]{Value: n.Cabecera.CodigoPuntoVenta},
		FechaEmision:                 datatype.NewTimeSiat(n.Cabecera.FechaEmision),
		NombreRazonSocial:            datatype.Nilable[string]{Value: n.Cabecera.NombreRazonSocial},
		CodigoTipoDocumentoIdentidad: n.Cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:              n.Cabecera.NumeroDocumento,
		Complemento:                  n.Cabecera.Complemento,
		CodigoCliente:                n.Cabecera.CodigoCliente,
		NumeroFactura:                n.Cabecera.NumeroFactura,
		NumeroAutorizacionCuf:        n.Cabecera.NumeroAutorizacionCuf,
		FechaEmisionFactura:          datatype.NewTimeSiat(n.Cabecera.FechaEmisionFactura),
		SubTotalNetoOriginal:         round(n.Cabecera.SubTotalNetoOriginal),
		TotalPrecioNetoOriginal:      round(n.Cabecera.TotalPrecioNetoOriginal),
		IvaOriginal:                  round(n.Cabecera.IvaOriginal),
		MontoAPagarOriginal:          round(n.Cabecera.MontoAPagarOriginal),
		SubTotalNetoDevuelto:         round(n.Cabecera.SubTotalNetoDevuelto),
		TotalPrecioNetoDevuelto:      round(n.Cabecera.TotalPrecioNetoDevuelto),
		IvaDevuelto:                  round(n.Cabecera.IvaDevuelto),
		MontoAPagarDevuelto:          round(n.Cabecera.MontoAPagarDevuelto),
		DescuentoAdicionalOriginal:   round(n.Cabecera.DescuentoAdicionalOriginal),
		MontoDescuentoDevuelto:       round(n.Cabecera.MontoDescuentoDevuelto),
		CodigoExcepcion:              datatype.Nilable[int]{Value: n.Cabecera.CodigoExcepcion},
		Leyenda:                      n.Cabecera.Leyenda,
		Usuario:                      n.Cabecera.Usuario,
		CodigoInterno:                n.Cabecera.CodigoInterno,
		CodigoDocumentoSector:        n.Cabecera.CodigoDocumentoSector,
	}
	if cabecera.CodigoDocumentoSector == 0 {
		cabecera.CodigoDocumentoSector = 24
	}
	detalles := make([]documents.DetalleNotaCreditoDebito, 0, len(n.Detalles))
	for _, d := range n.Detalles {
		detalles = append(detalles, documents.DetalleNotaCreditoDebito{
			NroItem:                  d.NroItem,
			ActividadEconomica:       d.ActividadEconomica,
			CodigoProductoSin:        d.CodigoProductoSin,
			CodigoProducto:           d.CodigoProducto,
			Descripcion:              d.Descripcion,
			Cantidad:                 round(d.Cantidad),
			UnidadMedida:             d.UnidadMedida,
			PrecioNetoUnitario:       round(d.PrecioNetoUnitario),
			MontoDescuento:           datatype.Nilable[float64]{Value: roundOptional(d.MontoDescuento)},
			SubTotal:                 round(d.SubTotal),
			CodigoDetalleTransaccion: d.CodigoDetalleTransaccion,
		})
	}
	return e.Encode(documents.NotaCreditoDebito{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cabecera,
		Detalle:           detalles,
	})
}
