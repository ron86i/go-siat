// Package turismohospedaje proporciona documentos del sector 6 para SIAT v3.
package turismohospedaje

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/v3/internal/core/domain/documents"
	"github.com/shopspring/decimal"
)

// Cabecera contiene los datos fiscales de Turismo y Hospedaje.
type Cabecera struct {
	NitEmisor         int64
	RazonSocialEmisor string
	Municipio         string
	Telefono          *string
	NumeroFactura     int64
	Cuf               string

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

	RazonSocialOperadorTurismo *string
	CantidadHuespedes          *int
	CantidadHabitaciones       *int
	CantidadMayores            *int
	CantidadMenores            *int
	FechaIngresoHospedaje      *time.Time
	CodigoMetodoPago           int
	NumeroTarjeta              *int64
	SubTotalNeto               float64
	TotalPrecioNeto            float64
	Iva                        float64
	MontoAPagar                float64

	CodigoMoneda          int
	TipoCambio            float64
	MontoTotalMoneda      float64
	MontoGiftCard         *float64
	DescuentoAdicional    *float64
	CodigoExcepcion       *int
	Leyenda               string
	Usuario               string
	CodigoInterno         string
	CodigoDocumentoSector int
}

// Detalle contiene un servicio turístico o de hospedaje.
type Detalle struct {
	ActividadEconomica   string
	CodigoProductoSin    int64
	CodigoProducto       string
	Descripcion          string
	CodigoTipoHabitacion *int
	Cantidad             float64
	UnidadMedida         int
	PrecioNetoUnitario   float64
	MontoDescuento       *float64
	SubTotal             float64
	DetalleHuespedes     *string
}

// Modalidades de facturación admitidas por el SIAT.
const (
	ModalidadElectronica   = 1
	ModalidadComputarizada = 2
)

// Factura conserva una copia de los datos públicos de entrada validados por New.
type Factura struct {
	cabecera Cabecera
	detalles []Detalle
}

// New valida los datos y asigna el sector sin calcular importes tributarios.
func New(cabecera Cabecera, detalles []Detalle) (*Factura, error) {
	if cabecera.CodigoDocumentoSector == 0 {
		cabecera.CodigoDocumentoSector = 6
	}
	f := &Factura{cabecera: cloneCabecera(cabecera), detalles: cloneDetalles(detalles)}
	if err := f.Validate(); err != nil {
		return nil, err
	}
	return f, nil
}

// Cabecera devuelve una copia defensiva profunda de los datos de cabecera de la factura.
func (f *Factura) Cabecera() Cabecera {
	if f == nil {
		return Cabecera{}
	}
	return cloneCabecera(f.cabecera)
}

// Detalles devuelve una copia defensiva profunda de los ítems de detalle.
func (f *Factura) Detalles() []Detalle {
	if f == nil {
		return nil
	}
	return cloneDetalles(f.detalles)
}

func (f *Factura) CodigoDocumentoSector() int {
	if f == nil {
		return 0
	}
	return f.cabecera.CodigoDocumentoSector
}

// CodigoSucursal devuelve el código de sucursal de la cabecera.
func (f *Factura) CodigoSucursal() int {
	if f == nil {
		return 0
	}
	return f.cabecera.CodigoSucursal
}

// CodigoPuntoVenta devuelve el código de punto de venta si está definido.
func (f *Factura) CodigoPuntoVenta() *int {
	if f == nil || f.cabecera.CodigoPuntoVenta == nil {
		return nil
	}
	return new(*f.cabecera.CodigoPuntoVenta)
}

// Cufd devuelve el CUFD consignado en la factura.
func (f *Factura) Cufd() string {
	if f == nil {
		return ""
	}
	return f.cabecera.Cufd
}

// Protocolo devuelve la versión del protocolo SIAT (3 para SIAT v3).
func (f *Factura) Protocolo() int {
	if f == nil {
		return 0
	}
	return 3
}

func cloneCabecera(c Cabecera) Cabecera {
	if c.Telefono != nil {
		c.Telefono = new(*c.Telefono)
	}
	if c.CodigoPuntoVenta != nil {
		c.CodigoPuntoVenta = new(*c.CodigoPuntoVenta)
	}
	if c.NombreRazonSocial != nil {
		c.NombreRazonSocial = new(*c.NombreRazonSocial)
	}
	if c.RazonSocialOperadorTurismo != nil {
		c.RazonSocialOperadorTurismo = new(*c.RazonSocialOperadorTurismo)
	}
	if c.CantidadHuespedes != nil {
		c.CantidadHuespedes = new(*c.CantidadHuespedes)
	}
	if c.CantidadHabitaciones != nil {
		c.CantidadHabitaciones = new(*c.CantidadHabitaciones)
	}
	if c.CantidadMayores != nil {
		c.CantidadMayores = new(*c.CantidadMayores)
	}
	if c.CantidadMenores != nil {
		c.CantidadMenores = new(*c.CantidadMenores)
	}
	if c.FechaIngresoHospedaje != nil {
		c.FechaIngresoHospedaje = new(*c.FechaIngresoHospedaje)
	}
	if c.NumeroTarjeta != nil {
		c.NumeroTarjeta = new(*c.NumeroTarjeta)
	}
	if c.MontoGiftCard != nil {
		c.MontoGiftCard = new(*c.MontoGiftCard)
	}
	if c.DescuentoAdicional != nil {
		c.DescuentoAdicional = new(*c.DescuentoAdicional)
	}
	if c.CodigoExcepcion != nil {
		c.CodigoExcepcion = new(*c.CodigoExcepcion)
	}
	return c
}

func cloneDetalles(detalles []Detalle) []Detalle {
	if detalles == nil {
		return nil
	}
	cloned := make([]Detalle, len(detalles))
	for i, d := range detalles {
		if d.CodigoTipoHabitacion != nil {
			d.CodigoTipoHabitacion = new(*d.CodigoTipoHabitacion)
		}
		if d.MontoDescuento != nil {
			d.MontoDescuento = new(*d.MontoDescuento)
		}
		if d.DetalleHuespedes != nil {
			d.DetalleHuespedes = new(*d.DetalleHuespedes)
		}
		cloned[i] = d
	}
	return cloned
}

// Validate comprueba las reglas estructurales normativas del SIN usando errors.Join para recopilar todas las violaciones encontradas en una sola pasada.
func (f *Factura) Validate() error {
	if f == nil {
		return errors.New("la factura no puede ser nil")
	}
	c := &f.cabecera
	var errs []error
	for _, regla := range []struct {
		invalida bool
		mensaje  string
	}{
		{c.CodigoDocumentoSector != 6, "el sector de turismo y hospedaje debe ser 6"},
		{c.NitEmisor <= 0, "NIT del emisor inválido"},
		{c.NumeroFactura <= 0, "número de factura inválido"},
		{c.FechaEmision.IsZero(), "fecha de emisión requerida"},
		{c.Cuf == "", "CUF requerido"},
		{c.Cufd == "", "CUFD requerido"},
		{c.CodigoSucursal < 0, "código de sucursal inválido"},
		{c.CodigoPuntoVenta != nil && *c.CodigoPuntoVenta < 0, "código de punto de venta inválido"},
		{c.RazonSocialEmisor == "", "razón social del emisor requerida"},
		{c.Municipio == "", "municipio requerido"},
		{c.Direccion == "", "dirección requerida"},
		{c.CodigoTipoDocumentoIdentidad <= 0, "tipo de documento de identidad inválido"},
		{c.NumeroDocumento == "", "número de documento requerido"},
		{c.CodigoCliente == "", "código de cliente requerido"},
		{c.CodigoMetodoPago <= 0, "código de método de pago inválido"},
		{c.CodigoMoneda <= 0, "código de moneda inválido"},
		{c.TipoCambio <= 0, "tipo de cambio debe ser mayor a cero"},
		{c.Leyenda == "", "leyenda requerida"},
		{c.Usuario == "", "usuario requerido"},
		{c.FechaIngresoHospedaje != nil && c.FechaIngresoHospedaje.IsZero(), "fecha de ingreso inválida"},
		{len(f.detalles) == 0, "la factura requiere al menos un detalle"},
	} {
		if regla.invalida {
			errs = append(errs, errors.New(regla.mensaje))
		}
	}
	for _, cantidad := range []*int{c.CantidadHuespedes, c.CantidadHabitaciones, c.CantidadMayores, c.CantidadMenores} {
		if cantidad != nil && *cantidad < 0 {
			errs = append(errs, errors.New("las cantidades de hospedaje no pueden ser negativas"))
		}
	}
	valores := []float64{c.SubTotalNeto, c.TotalPrecioNeto, c.Iva, c.MontoAPagar, c.TipoCambio, c.MontoTotalMoneda}
	for _, v := range []*float64{c.MontoGiftCard, c.DescuentoAdicional} {
		if v != nil {
			valores = append(valores, *v)
		}
	}
	for _, v := range valores {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			errs = append(errs, errors.New("la cabecera contiene un decimal no finito"))
			break
		}
	}
	for i, d := range f.detalles {
		if d.ActividadEconomica == "" || d.CodigoProductoSin <= 0 || d.CodigoProducto == "" || d.Descripcion == "" || d.UnidadMedida <= 0 {
			errs = append(errs, fmt.Errorf("detalle %d: datos del servicio incompletos o inválidos", i+1))
		}
		if d.Cantidad <= 0 || d.PrecioNetoUnitario < 0 || d.SubTotal < 0 || (d.MontoDescuento != nil && *d.MontoDescuento < 0) {
			errs = append(errs, fmt.Errorf("detalle %d: cantidad o importes inválidos", i+1))
		}
		if d.CodigoTipoHabitacion != nil && *d.CodigoTipoHabitacion <= 0 {
			errs = append(errs, fmt.Errorf("detalle %d: tipo de habitación inválido", i+1))
		}
		valores := []float64{d.Cantidad, d.PrecioNetoUnitario, d.SubTotal}
		if d.MontoDescuento != nil {
			valores = append(valores, *d.MontoDescuento)
		}
		for _, v := range valores {
			if math.IsNaN(v) || math.IsInf(v, 0) {
				errs = append(errs, fmt.Errorf("detalle %d: decimal no finito", i+1))
				break
			}
		}
	}
	return errors.Join(errs...)
}

// Marshal serializa la factura a XML normalizado según la modalidad indicada.
func Marshal(f *Factura, modalidad int) ([]byte, error) { return f.XML(modalidad) }

// XML serializa la factura con la modalidad determinada por facturación.
func (f *Factura) XML(modalidad int) ([]byte, error) { return f.MarshalXMLModalidad(modalidad) }

// MarshalXMLModalidad implementa pipeline.ModalidadMarshaler, serializando el XML con el elemento raíz correspondiente a la modalidad indicada por el servicio.
func (f *Factura) MarshalXMLModalidad(modalidad int) ([]byte, error) {
	if f == nil {
		return nil, errors.New("la factura no puede ser nil")
	}
	if modalidad != ModalidadElectronica && modalidad != ModalidadComputarizada {
		return nil, fmt.Errorf("modalidad inválida: %d", modalidad)
	}
	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	if err := f.encodeXML(enc, modalidad); err != nil {
		return nil, err
	}
	if err := enc.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalXML rechaza la serialización directa sin modalidad para evitar ambigüedades. Debe utilizarse factura.XML(modalidad) o el pipeline de recepción de facturacion.
func (f *Factura) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return errors.New("turismohospedaje.Factura requiere modalidad explícita: use XML(modalidad) o el pipeline de facturación")
}

func round(v float64) float64 { return decimal.NewFromFloat(v).Round(2).InexactFloat64() }
func roundOptional(v *float64) *float64 {
	if v == nil {
		return nil
	}
	return new(round(*v))
}

func (d *Factura) encodeXML(e *xml.Encoder, modalidad int) error {
	if err := d.Validate(); err != nil {
		return err
	}
	nombre := "facturaElectronicaServicioTuristicoHospedaje"
	if modalidad == 2 {
		nombre = "facturaComputarizadaServicioTuristicoHospedaje"
	}
	cab := documents.CabeceraTurismoHospedaje{
		NitEmisor:                    d.cabecera.NitEmisor,
		RazonSocialEmisor:            d.cabecera.RazonSocialEmisor,
		Municipio:                    d.cabecera.Municipio,
		Telefono:                     datatype.Nilable[string]{Value: d.cabecera.Telefono},
		NumeroFactura:                d.cabecera.NumeroFactura,
		Cuf:                          d.cabecera.Cuf,
		Cufd:                         d.cabecera.Cufd,
		CodigoSucursal:               d.cabecera.CodigoSucursal,
		Direccion:                    d.cabecera.Direccion,
		CodigoPuntoVenta:             datatype.Nilable[int]{Value: d.cabecera.CodigoPuntoVenta},
		FechaEmision:                 datatype.NewTimeSiat(d.cabecera.FechaEmision),
		NombreRazonSocial:            datatype.Nilable[string]{Value: d.cabecera.NombreRazonSocial},
		CodigoTipoDocumentoIdentidad: d.cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:              d.cabecera.NumeroDocumento,
		Complemento:                  d.cabecera.Complemento,
		CodigoCliente:                d.cabecera.CodigoCliente,
		RazonSocialOperadorTurismo:   datatype.Nilable[string]{Value: d.cabecera.RazonSocialOperadorTurismo},
		CantidadHuespedes:            datatype.Nilable[int]{Value: d.cabecera.CantidadHuespedes},
		CantidadHabitaciones:         datatype.Nilable[int]{Value: d.cabecera.CantidadHabitaciones},
		CantidadMayores:              datatype.Nilable[int]{Value: d.cabecera.CantidadMayores},
		CantidadMenores:              datatype.Nilable[int]{Value: d.cabecera.CantidadMenores},
		FechaIngresoHospedaje: func() datatype.Nilable[datatype.TimeSiat] {
			if d.cabecera.FechaIngresoHospedaje == nil {
				return datatype.Nilable[datatype.TimeSiat]{}
			}
			return datatype.Nilable[datatype.TimeSiat]{Value: new(datatype.NewTimeSiat(*d.cabecera.FechaIngresoHospedaje))}
		}(),
		CodigoMetodoPago:      d.cabecera.CodigoMetodoPago,
		NumeroTarjeta:         datatype.Nilable[int64]{Value: d.cabecera.NumeroTarjeta},
		SubTotalNeto:          round(d.cabecera.SubTotalNeto),
		TotalPrecioNeto:       round(d.cabecera.TotalPrecioNeto),
		Iva:                   round(d.cabecera.Iva),
		MontoAPagar:           round(d.cabecera.MontoAPagar),
		CodigoMoneda:          d.cabecera.CodigoMoneda,
		TipoCambio:            round(d.cabecera.TipoCambio),
		MontoTotalMoneda:      round(d.cabecera.MontoTotalMoneda),
		MontoGiftCard:         datatype.Nilable[float64]{Value: roundOptional(d.cabecera.MontoGiftCard)},
		DescuentoAdicional:    datatype.Nilable[float64]{Value: roundOptional(d.cabecera.DescuentoAdicional)},
		CodigoExcepcion:       datatype.Nilable[int]{Value: d.cabecera.CodigoExcepcion},
		Leyenda:               d.cabecera.Leyenda,
		Usuario:               d.cabecera.Usuario,
		CodigoInterno:         d.cabecera.CodigoInterno,
		CodigoDocumentoSector: d.cabecera.CodigoDocumentoSector,
	}
	if cab.CodigoDocumentoSector == 0 {
		cab.CodigoDocumentoSector = 6
	}
	dets := make([]documents.DetalleTurismoHospedaje, 0, len(d.detalles))
	for _, item := range d.detalles {
		dets = append(dets, documents.DetalleTurismoHospedaje{
			ActividadEconomica:   item.ActividadEconomica,
			CodigoProductoSin:    item.CodigoProductoSin,
			CodigoProducto:       item.CodigoProducto,
			Descripcion:          item.Descripcion,
			CodigoTipoHabitacion: datatype.Nilable[int]{Value: item.CodigoTipoHabitacion},
			Cantidad:             round(item.Cantidad),
			UnidadMedida:         item.UnidadMedida,
			PrecioNetoUnitario:   round(item.PrecioNetoUnitario),
			MontoDescuento:       datatype.Nilable[float64]{Value: roundOptional(item.MontoDescuento)},
			SubTotal:             round(item.SubTotal),
			DetalleHuespedes:     datatype.Nilable[string]{Value: item.DetalleHuespedes},
		})
	}
	doc := documents.FacturaTurismoHospedaje{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cab,
		Detalle:           dets,
	}
	return e.Encode(doc)
}
