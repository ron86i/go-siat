// Package tasacero proporciona el modelo de factura Tasa Cero (sector 8) para SIAT v3.
package tasacero

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ron86i/go-siat/v3/internal/core/domain/datatype"
	"github.com/shopspring/decimal"
)

// Cabecera representa los datos generales de cabecera de la factura Tasa Cero (sector 8) para SIAT v3.
type Cabecera struct {
	NitEmisor int64

	RazonSocialEmisor string
	Municipio         string
	Telefono          *string
	NumeroFactura     int64
	Cuf               string
	Cufd              string
	CodigoSucursal    int
	Direccion         string

	CodigoPuntoVenta             *int
	FechaEmision                 time.Time
	NombreRazonSocial            *string
	CodigoTipoDocumentoIdentidad int
	NumeroDocumento              string

	Complemento      string
	CodigoCliente    string
	CodigoMetodoPago int
	NumeroTarjeta    *int64
	SubTotalNeto     float64

	TotalPrecioNeto  float64
	Iva              float64
	MontoAPagar      float64
	CodigoMoneda     int
	TipoCambio       float64
	MontoTotalMoneda float64

	MontoGiftCard         *float64
	DescuentoAdicional    *float64
	CodigoExcepcion       *int
	Leyenda               string
	Usuario               string
	CodigoInterno         string
	CodigoDocumentoSector int
}

// Detalle representa un ítem o producto dentro de la factura Tasa Cero para SIAT v3.
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

// Modalidades de facturación admitidas por el SIAT.
const (
	ModalidadElectronica   = 1
	ModalidadComputarizada = 2
)

// Factura contiene los datos fiscales de una factura Tasa Cero (sector 8).
// Sus campos internos son privados para garantizar que toda instancia
// se construya mediante New(), asegurando su integridad y validación estricta.
type Factura struct {
	cabecera Cabecera
	detalles []Detalle
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

// CodigoDocumentoSector devuelve el código de documento sector (8 para Tasa Cero).
func (f *Factura) CodigoDocumentoSector() int {
	if f == nil {
		return 0
	}
	return f.cabecera.CodigoDocumentoSector
}

// CodigoSector es un alias para CodigoDocumentoSector.
func (f *Factura) CodigoSector() int {
	return f.CodigoDocumentoSector()
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
	v := *f.cabecera.CodigoPuntoVenta
	return &v
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

// New construye, asigna metadatos técnicos y valida una nueva Factura Tasa Cero v3.
// Si CodigoDocumentoSector en cabecera es 0, se asigna automáticamente 8.
// Si se proporciona un código de sector distinto de 8, se rechaza explícitamente.
// Realiza copia defensiva profunda de la cabecera y detalles para garantizar inmutabilidad.
func New(cabecera Cabecera, detalles []Detalle) (*Factura, error) {
	if cabecera.CodigoDocumentoSector == 0 {
		cabecera.CodigoDocumentoSector = 8
	} else if cabecera.CodigoDocumentoSector != 8 {
		return nil, errors.New("sector incompatible con tasa cero")
	}

	factura := &Factura{
		cabecera: cloneCabecera(cabecera),
		detalles: cloneDetalles(detalles),
	}

	if err := factura.Validate(); err != nil {
		return nil, err
	}

	return factura, nil
}

func cloneCabecera(c Cabecera) Cabecera {
	res := c
	if c.Telefono != nil {
		v := *c.Telefono
		res.Telefono = &v
	}
	if c.CodigoPuntoVenta != nil {
		v := *c.CodigoPuntoVenta
		res.CodigoPuntoVenta = &v
	}
	if c.NombreRazonSocial != nil {
		v := *c.NombreRazonSocial
		res.NombreRazonSocial = &v
	}
	if c.NumeroTarjeta != nil {
		v := *c.NumeroTarjeta
		res.NumeroTarjeta = &v
	}
	if c.MontoGiftCard != nil {
		v := *c.MontoGiftCard
		res.MontoGiftCard = &v
	}
	if c.DescuentoAdicional != nil {
		v := *c.DescuentoAdicional
		res.DescuentoAdicional = &v
	}
	if c.CodigoExcepcion != nil {
		v := *c.CodigoExcepcion
		res.CodigoExcepcion = &v
	}
	return res
}

func cloneDetalle(d Detalle) Detalle {
	res := d
	if d.MontoDescuento != nil {
		v := *d.MontoDescuento
		res.MontoDescuento = &v
	}
	return res
}

func cloneDetalles(detalles []Detalle) []Detalle {
	if detalles == nil {
		return nil
	}
	cloned := make([]Detalle, len(detalles))
	for i, d := range detalles {
		cloned[i] = cloneDetalle(d)
	}
	return cloned
}

// Validate comprueba las reglas estructurales normativas del SIN usando errors.Join para recopilar todas las violaciones encontradas en una sola pasada.
func (f *Factura) Validate() error {
	if f == nil {
		return errors.New("la factura no puede ser nil")
	}

	var errs []error

	f.validateCabecera(&errs)
	f.validateDetalles(&errs)
	f.validateDecimals(&errs)

	return errors.Join(errs...)
}

// Validaciones de longitud máxima de cadenas según especificación técnica del SIN.
func (f *Factura) validateCabecera(errs *[]error) {
	c := &f.cabecera

	if c.CodigoDocumentoSector != 8 {
		*errs = append(*errs, errors.New("el sector de tasa cero debe ser 8"))
	}
	if c.NitEmisor <= 0 {
		*errs = append(*errs, errors.New("NIT del emisor inválido"))
	}
	if c.NumeroFactura <= 0 {
		*errs = append(*errs, errors.New("número de factura debe ser mayor a cero"))
	}
	if c.FechaEmision.IsZero() {
		*errs = append(*errs, errors.New("fecha de emisión requerida"))
	}
	if c.Cuf == "" {
		*errs = append(*errs, errors.New("CUF requerido"))
	}
	if c.Cufd == "" {
		*errs = append(*errs, errors.New("CUFD requerido"))
	}
	if c.CodigoSucursal < 0 {
		*errs = append(*errs, errors.New("código de sucursal inválido"))
	}
	if c.Direccion == "" {
		*errs = append(*errs, errors.New("dirección requerida"))
	}
	if c.CodigoMetodoPago <= 0 {
		*errs = append(*errs, errors.New("código de método de pago inválido"))
	}
	if c.CodigoMoneda <= 0 {
		*errs = append(*errs, errors.New("código de moneda inválido"))
	}
	if c.TipoCambio <= 0 {
		*errs = append(*errs, errors.New("tipo de cambio debe ser mayor a cero"))
	}
	if c.Leyenda == "" {
		*errs = append(*errs, errors.New("leyenda requerida"))
	}
	if c.Usuario == "" {
		*errs = append(*errs, errors.New("usuario requerido"))
	}

	if len(c.RazonSocialEmisor) > 200 {
		*errs = append(*errs, errors.New("razonSocialEmisor excede la longitud máxima (200)"))
	}
	if len(c.Municipio) > 100 {
		*errs = append(*errs, errors.New("municipio excede la longitud máxima (100)"))
	}
	if c.Telefono != nil && len(*c.Telefono) > 20 {
		*errs = append(*errs, errors.New("telefono excede la longitud máxima (20)"))
	}
	if len(c.Direccion) > 200 {
		*errs = append(*errs, errors.New("direccion excede la longitud máxima (200)"))
	}
	if c.NombreRazonSocial != nil && len(*c.NombreRazonSocial) > 250 {
		*errs = append(*errs, errors.New("nombreRazonSocial excede la longitud máxima (250)"))
	}
	if len(c.NumeroDocumento) > 30 {
		*errs = append(*errs, errors.New("numeroDocumento excede la longitud máxima (30)"))
	}
	if len(c.Complemento) > 5 {
		*errs = append(*errs, errors.New("complemento excede la longitud máxima (5)"))
	}
	if len(c.CodigoCliente) > 100 {
		*errs = append(*errs, errors.New("codigoCliente excede la longitud máxima (100)"))
	}
	if len(c.Leyenda) > 500 {
		*errs = append(*errs, errors.New("leyenda excede la longitud máxima (500)"))
	}
	if len(c.Usuario) > 100 {
		*errs = append(*errs, errors.New("usuario excede la longitud máxima (100)"))
	}
	if len(c.CodigoInterno) > 50 {
		*errs = append(*errs, errors.New("codigoInterno excede la longitud máxima (50)"))
	}
}

func (f *Factura) validateDetalles(errs *[]error) {
	if len(f.detalles) == 0 {
		*errs = append(*errs, errors.New("la factura requiere al menos un detalle"))
		return
	}

	for i, d := range f.detalles {
		idx := i + 1
		if d.ActividadEconomica == "" {
			*errs = append(*errs, fmt.Errorf("detalle %d: actividad económica requerida", idx))
		}
		if d.CodigoProductoSin <= 0 {
			*errs = append(*errs, fmt.Errorf("detalle %d: código de producto SIN inválido", idx))
		}
		if d.CodigoProducto == "" {
			*errs = append(*errs, fmt.Errorf("detalle %d: código de producto requerido", idx))
		}
		if d.Descripcion == "" {
			*errs = append(*errs, fmt.Errorf("detalle %d: descripción requerida", idx))
		}
		if d.Cantidad <= 0 {
			*errs = append(*errs, fmt.Errorf("detalle %d: cantidad debe ser mayor a cero", idx))
		}
		if d.UnidadMedida <= 0 {
			*errs = append(*errs, fmt.Errorf("detalle %d: unidad de medida inválida", idx))
		}
		if d.PrecioNetoUnitario < 0 {
			*errs = append(*errs, fmt.Errorf("detalle %d: precio neto unitario no puede ser negativo", idx))
		}
		if d.SubTotal < 0 {
			*errs = append(*errs, fmt.Errorf("detalle %d: subtotal no puede ser negativo", idx))
		}
		if d.MontoDescuento != nil && *d.MontoDescuento < 0 {
			*errs = append(*errs, fmt.Errorf("detalle %d: monto de descuento no puede ser negativo", idx))
		}

		if len(d.ActividadEconomica) > 20 {
			*errs = append(*errs, fmt.Errorf("detalle %d: actividadEconomica excede longitud máxima (20)", idx))
		}
		if len(d.CodigoProducto) > 50 {
			*errs = append(*errs, fmt.Errorf("detalle %d: codigoProducto excede longitud máxima (50)", idx))
		}
		if len(d.Descripcion) > 500 {
			*errs = append(*errs, fmt.Errorf("detalle %d: descripcion excede longitud máxima (500)", idx))
		}
	}
}

func (f *Factura) validateDecimals(errs *[]error) {
	c := &f.cabecera
	valoresCabecera := []float64{
		c.SubTotalNeto,
		c.TotalPrecioNeto,
		c.Iva,
		c.MontoAPagar,
		c.TipoCambio,
		c.MontoTotalMoneda,
	}
	if c.MontoGiftCard != nil {
		valoresCabecera = append(valoresCabecera, *c.MontoGiftCard)
	}
	if c.DescuentoAdicional != nil {
		valoresCabecera = append(valoresCabecera, *c.DescuentoAdicional)
	}

	for _, v := range valoresCabecera {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			*errs = append(*errs, errors.New("la cabecera contiene un valor decimal no finito"))
			break
		}
	}

	for i, d := range f.detalles {
		valoresDetalle := []float64{d.Cantidad, d.PrecioNetoUnitario, d.SubTotal}
		if d.MontoDescuento != nil {
			valoresDetalle = append(valoresDetalle, *d.MontoDescuento)
		}
		for _, v := range valoresDetalle {
			if math.IsNaN(v) || math.IsInf(v, 0) {
				*errs = append(*errs, fmt.Errorf("detalle %d contiene un valor decimal no finito", i+1))
				break
			}
		}
	}
}

// Marshal serializa la factura a XML normalizado según la modalidad indicada.
func Marshal(f *Factura, modalidad int) ([]byte, error) {
	if f == nil {
		return nil, errors.New("la factura no puede ser nil")
	}
	return f.XML(modalidad)
}

// XML serializa la factura a XML normalizado según la modalidad exigida (ModalidadElectronica o ModalidadComputarizada). La modalidad es obligatoria para evitar generar accidentalmente un documento con esquema incorrecto.
func (f *Factura) XML(modalidad int) ([]byte, error) {
	return f.MarshalXMLModalidad(modalidad)
}

// MarshalXMLModalidad implementa pipeline.ModalidadMarshaler, serializando el XML con el elemento raíz correspondiente a la modalidad indicada por el servicio.
func (f *Factura) MarshalXMLModalidad(modalidad int) ([]byte, error) {
	if f == nil {
		return nil, errors.New("la factura no puede ser nil")
	}
	if modalidad != ModalidadElectronica && modalidad != ModalidadComputarizada {
		return nil, fmt.Errorf("modalidad inválida: %d (use ModalidadElectronica o ModalidadComputarizada)", modalidad)
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
	return errors.New("tasacero.Factura requiere especificar la modalidad de emisión (use factura.XML(modalidad) o el pipeline de facturacion)")
}

type xmlCabecera struct {
	NitEmisor                    int64                     `xml:"nitEmisor"`
	RazonSocialEmisor            string                    `xml:"razonSocialEmisor"`
	Municipio                    string                    `xml:"municipio"`
	Telefono                     datatype.Nilable[string]  `xml:"telefono"`
	NumeroFactura                int64                     `xml:"numeroFactura"`
	Cuf                          string                    `xml:"cuf"`
	Cufd                         string                    `xml:"cufd"`
	CodigoSucursal               int                       `xml:"codigoSucursal"`
	Direccion                    string                    `xml:"direccion"`
	CodigoPuntoVenta             datatype.Nilable[int]     `xml:"codigoPuntoVenta"`
	FechaEmision                 datatype.TimeSiat         `xml:"fechaEmision"`
	NombreRazonSocial            datatype.Nilable[string]  `xml:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int                       `xml:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string                    `xml:"numeroDocumento"`
	Complemento                  string                    `xml:"complemento"`
	CodigoCliente                string                    `xml:"codigoCliente"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago"`
	NumeroTarjeta                datatype.Nilable[int64]   `xml:"numeroTarjeta"`
	SubTotalNeto                 float64                   `xml:"subTotalNeto"`
	TotalPrecioNeto              float64                   `xml:"totalPrecioNeto"`
	Iva                          float64                   `xml:"iva"`
	MontoAPagar                  float64                   `xml:"montoaPagar"`
	CodigoMoneda                 int                       `xml:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda"`
	MontoGiftCard                datatype.Nilable[float64] `xml:"montoGiftCard"`
	DescuentoAdicional           datatype.Nilable[float64] `xml:"descuentoAdicional"`
	CodigoExcepcion              datatype.Nilable[int]     `xml:"codigoExcepcion"`
	Leyenda                      string                    `xml:"leyenda"`
	Usuario                      string                    `xml:"usuario"`
	CodigoInterno                string                    `xml:"codigoInterno"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector"`
}

type xmlDetalle struct {
	ActividadEconomica string                    `xml:"actividadEconomica"`
	CodigoProductoSin  int64                     `xml:"codigoProductoSin"`
	CodigoProducto     string                    `xml:"codigoProducto"`
	Descripcion        string                    `xml:"descripcion"`
	Cantidad           float64                   `xml:"cantidad"`
	UnidadMedida       int                       `xml:"unidadMedida"`
	PrecioNetoUnitario float64                   `xml:"precioNetoUnitario"`
	MontoDescuento     datatype.Nilable[float64] `xml:"montoDescuento"`
	SubTotal           float64                   `xml:"subTotal"`
}

func (f *Factura) encodeXML(e *xml.Encoder, modalidad int) error {
	if err := f.Validate(); err != nil {
		return err
	}
	nombre := "facturaElectronicaTasaCero"
	if modalidad == ModalidadComputarizada {
		nombre = "facturaComputarizadaTasaCero"
	}
	cabecera := xmlCabecera{
		NitEmisor:                    f.cabecera.NitEmisor,
		RazonSocialEmisor:            f.cabecera.RazonSocialEmisor,
		Municipio:                    f.cabecera.Municipio,
		Telefono:                     datatype.Nilable[string]{Value: f.cabecera.Telefono},
		NumeroFactura:                f.cabecera.NumeroFactura,
		Cuf:                          f.cabecera.Cuf,
		Cufd:                         f.cabecera.Cufd,
		CodigoSucursal:               f.cabecera.CodigoSucursal,
		Direccion:                    f.cabecera.Direccion,
		CodigoPuntoVenta:             datatype.Nilable[int]{Value: f.cabecera.CodigoPuntoVenta},
		FechaEmision:                 datatype.NewTimeSiat(f.cabecera.FechaEmision),
		NombreRazonSocial:            datatype.Nilable[string]{Value: f.cabecera.NombreRazonSocial},
		CodigoTipoDocumentoIdentidad: f.cabecera.CodigoTipoDocumentoIdentidad,
		NumeroDocumento:              f.cabecera.NumeroDocumento,
		Complemento:                  f.cabecera.Complemento,
		CodigoCliente:                f.cabecera.CodigoCliente,
		CodigoMetodoPago:             f.cabecera.CodigoMetodoPago,
		NumeroTarjeta:                datatype.Nilable[int64]{Value: f.cabecera.NumeroTarjeta},
		SubTotalNeto:                 round(f.cabecera.SubTotalNeto),
		TotalPrecioNeto:              round(f.cabecera.TotalPrecioNeto),
		Iva:                          round(f.cabecera.Iva),
		MontoAPagar:                  round(f.cabecera.MontoAPagar),
		CodigoMoneda:                 f.cabecera.CodigoMoneda,
		TipoCambio:                   round(f.cabecera.TipoCambio),
		MontoTotalMoneda:             round(f.cabecera.MontoTotalMoneda),
		MontoGiftCard:                datatype.Nilable[float64]{Value: roundOptional(f.cabecera.MontoGiftCard)},
		DescuentoAdicional:           datatype.Nilable[float64]{Value: roundOptional(f.cabecera.DescuentoAdicional)},
		CodigoExcepcion:              datatype.Nilable[int]{Value: f.cabecera.CodigoExcepcion},
		Leyenda:                      f.cabecera.Leyenda,
		Usuario:                      f.cabecera.Usuario,
		CodigoInterno:                f.cabecera.CodigoInterno,
		CodigoDocumentoSector:        f.cabecera.CodigoDocumentoSector,
	}
	detalles := make([]xmlDetalle, 0, len(f.detalles))
	for _, d := range f.detalles {
		detalles = append(detalles, xmlDetalle{
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
	documento := struct {
		XMLName           xml.Name     `json:"-"`
		XmlnsXsi          string       `xml:"xmlns:xsi,attr" json:"-"`
		XsiSchemaLocation string       `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
		Cabecera          xmlCabecera  `xml:"cabecera" json:"cabecera"`
		Detalle           []xmlDetalle `xml:"detalle" json:"detalle"`
	}{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: nombre + ".xsd",
		Cabecera:          cabecera,
		Detalle:           detalles,
	}
	return e.Encode(documento)
}

func round(v float64) float64 {
	return decimal.NewFromFloat(v).Round(2).InexactFloat64()
}

func roundOptional(v *float64) *float64 {
	if v == nil {
		return nil
	}
	return new(round(*v))
}
