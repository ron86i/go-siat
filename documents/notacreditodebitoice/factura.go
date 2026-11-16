// Package notacreditodebitoice proporciona documentos para el sector Nota Crédito Débito ICE en SIAT v3.
package notacreditodebitoice

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
	TotalIcePorcentualOriginal   float64
	TotalIceEspecificoOriginal   float64
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

// Detalle contiene los ítems del documento fiscal v3.
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
	MarcaIce                 int
	PrecioNetoVentaIce       *float64
	AlicuotaEspecifica       *float64
	AlicuotaPorcentual       *float64
	MontoIceEspecifico       *float64
	MontoIcePorcentual       *float64
	CantidadIce              *float64
	CodigoDetalleTransaccion int
}

func (d NotaCreditoDebito) Validate() error {
	if d.Modalidad != 0 && d.Modalidad != 1 && d.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", d.Modalidad)
	}
	if d.Cabecera.CodigoDocumentoSector != 0 && d.Cabecera.CodigoDocumentoSector != 48 {
		return fmt.Errorf("código de sector %d no corresponde a Nota Crédito Débito ICE (esperado 48)", d.Cabecera.CodigoDocumentoSector)
	}
	return d.validateImportes()
}

func (d NotaCreditoDebito) validateImportes() error {
	if len(d.Detalles) == 0 {
		return nil
	}

	originales := make(map[int]Detalle)
	var subTotalOriginal, subTotalDevuelto, totalIceEspecifico, totalIcePorcentual decimal.Decimal
	hayDevolucion := false
	for i, item := range d.Detalles {
		if item.NroItem < 1 || item.NroItem > 9999 {
			return fmt.Errorf("detalles[%d].nroItem debe estar entre 1 y 9999", i)
		}
		if item.CodigoDetalleTransaccion != 1 && item.CodigoDetalleTransaccion != 2 {
			return fmt.Errorf("detalles[%d].codigoDetalleTransaccion debe ser 1 o 2", i)
		}
		if item.MarcaIce != 1 && item.MarcaIce != 2 {
			return fmt.Errorf("detalles[%d].marcaIce debe ser 1 o 2", i)
		}

		cantidad, precio := amount5(item.Cantidad), amount5(item.PrecioNetoUnitario)
		descuento := optionalAmount5(item.MontoDescuento)
		if cantidad.Sign() <= 0 || precio.Sign() <= 0 || descuento.Sign() < 0 {
			return fmt.Errorf("detalles[%d]: cantidad y precio deben ser positivos; el descuento no puede ser negativo", i)
		}
		subTotal := cantidad.Mul(precio).Sub(descuento).Round(5)
		if subTotal.Sign() <= 0 {
			return fmt.Errorf("detalles[%d].subTotal debe ser positivo", i)
		}
		if err := validateAmount(fmt.Sprintf("detalles[%d].subTotal", i), amount5(item.SubTotal), subTotal, 5); err != nil {
			return err
		}
		if item.CodigoDetalleTransaccion == 1 {
			if _, exists := originales[item.NroItem]; exists {
				return fmt.Errorf("detalles[%d]: nroItem %d repetido en la transacción original", i, item.NroItem)
			}
			originales[item.NroItem] = item
			subTotalOriginal = subTotalOriginal.Add(amount5(item.SubTotal))
			totalIceEspecifico = totalIceEspecifico.Add(optionalAmount5(item.MontoIceEspecifico))
			totalIcePorcentual = totalIcePorcentual.Add(optionalAmount5(item.MontoIcePorcentual))
		} else {
			hayDevolucion = true
			subTotalDevuelto = subTotalDevuelto.Add(amount5(item.SubTotal))
		}
	}

	c := d.Cabecera
	subTotalOriginal = subTotalOriginal.Round(2)
	if err := validateAmount("cabecera.subTotalNetoOriginal", amount2(c.SubTotalNetoOriginal), subTotalOriginal, 2); err != nil {
		return err
	}
	totalIceEspecifico = totalIceEspecifico.Round(2)
	totalIcePorcentual = totalIcePorcentual.Round(2)
	if err := validateAmount("cabecera.totalIceEspecificoOriginal", amount2(c.TotalIceEspecificoOriginal), totalIceEspecifico, 2); err != nil {
		return err
	}
	if err := validateAmount("cabecera.totalIcePorcentualOriginal", amount2(c.TotalIcePorcentualOriginal), totalIcePorcentual, 2); err != nil {
		return err
	}
	descuentoAdicional := amount2(c.DescuentoAdicionalOriginal)
	totalPrecioNetoOriginal := subTotalOriginal.Sub(descuentoAdicional).Round(2)
	if err := validateAmount("cabecera.totalPrecioNetoOriginal", amount2(c.TotalPrecioNetoOriginal), totalPrecioNetoOriginal, 2); err != nil {
		return err
	}
	ivaOriginal := totalPrecioNetoOriginal.Mul(decimal.NewFromFloat(0.13)).Round(2)
	if err := validateAmount("cabecera.ivaOriginal", amount2(c.IvaOriginal), ivaOriginal, 2); err != nil {
		return err
	}
	montoAPagarOriginal := totalPrecioNetoOriginal.Add(ivaOriginal).Add(totalIceEspecifico).Add(totalIcePorcentual).Round(2)
	if err := validateAmount("cabecera.montoaPagarOriginal", amount2(c.MontoAPagarOriginal), montoAPagarOriginal, 2); err != nil {
		return err
	}

	subTotalDevuelto = subTotalDevuelto.Round(2)
	if err := validateAmount("cabecera.subTotalNetoDevuelto", amount2(c.SubTotalNetoDevuelto), subTotalDevuelto, 2); err != nil {
		return err
	}
	descuentoDevuelto, err := d.descuentoDevuelto(originales, subTotalOriginal, totalIceEspecifico, totalIcePorcentual, hayDevolucion)
	if err != nil {
		return err
	}
	if err := validateAmount("cabecera.montoDescuentoDevuelto", amount2(c.MontoDescuentoDevuelto), descuentoDevuelto, 2); err != nil {
		return err
	}
	totalPrecioNetoDevuelto := subTotalDevuelto.Sub(descuentoDevuelto.Round(2)).Round(2)
	if err := validateAmount("cabecera.totalPrecioNetoDevuelto", amount2(c.TotalPrecioNetoDevuelto), totalPrecioNetoDevuelto, 2); err != nil {
		return err
	}
	ivaDevuelto := totalPrecioNetoDevuelto.Mul(decimal.NewFromFloat(0.13)).Round(2)
	if err := validateAmount("cabecera.ivaDevuelto", amount2(c.IvaDevuelto), ivaDevuelto, 2); err != nil {
		return err
	}
	montoAPagarDevuelto := totalPrecioNetoDevuelto.Add(ivaDevuelto).Round(2)
	return validateAmount("cabecera.montoaPagarDevuelto", amount2(c.MontoAPagarDevuelto), montoAPagarDevuelto, 2)
}

func (d NotaCreditoDebito) descuentoDevuelto(originales map[int]Detalle, subTotalNetoOriginal, totalIceEspecifico, totalIcePorcentual decimal.Decimal, hayDevolucion bool) (decimal.Decimal, error) {
	if !hayDevolucion {
		return decimal.Zero, nil
	}
	baseOriginal := subTotalNetoOriginal.Sub(totalIceEspecifico).Sub(totalIcePorcentual)
	if baseOriginal.Cmp(decimal.Zero) <= 0 {
		return decimal.Zero, fmt.Errorf("no se puede prorratear el descuento: la base original debe ser positiva")
	}

	descuentoAdicional := amount2(d.Cabecera.DescuentoAdicionalOriginal)
	var total decimal.Decimal
	for i, devuelto := range d.Detalles {
		if devuelto.CodigoDetalleTransaccion != 2 {
			continue
		}
		original, exists := originales[devuelto.NroItem]
		if !exists {
			return decimal.Zero, fmt.Errorf("detalles[%d]: nroItem %d no existe en la transacción original", i, devuelto.NroItem)
		}
		cantidadOriginal := amount5(original.Cantidad)
		if cantidadOriginal.Sign() <= 0 {
			return decimal.Zero, fmt.Errorf("detalles[%d]: la cantidad original de nroItem %d debe ser positiva", i, devuelto.NroItem)
		}
		baseItem := amount5(original.SubTotal).Sub(optionalAmount5(original.MontoIcePorcentual)).Sub(optionalAmount5(original.MontoIceEspecifico))
		descuentoItem := baseItem.Div(baseOriginal).Div(cantidadOriginal).Mul(descuentoAdicional).Mul(amount5(devuelto.Cantidad))
		total = total.Add(descuentoItem)
	}
	return total.Round(2), nil
}

func amount2(v float64) decimal.Decimal { return decimal.NewFromFloat(v).Round(2) }
func amount5(v float64) decimal.Decimal { return decimal.NewFromFloat(v).Round(5) }
func optionalAmount5(v *float64) decimal.Decimal {
	if v == nil {
		return decimal.Zero
	}
	return amount5(*v)
}
func validateAmount(field string, actual, expected decimal.Decimal, places int32) error {
	actual, expected = actual.Round(places), expected.Round(places)
	if actual.Cmp(expected) != 0 {
		return fmt.Errorf("%s no coincide: recibido %s, esperado %s", field, actual.StringFixed(places), expected.StringFixed(places))
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
func round5(v float64) float64 { return decimal.NewFromFloat(v).Round(5).InexactFloat64() }
func roundOptional5(v *float64) *float64 {
	if v == nil {
		return nil
	}
	return new(round5(*v))
}

func (d NotaCreditoDebito) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := d.Validate(); err != nil {
		return err
	}
	nombre := "notaElectronicaCreditoDebitoIce"
	if d.Modalidad == 2 {
		nombre = "notaComputarizadaCreditoDebitoIce"
	}
	cab := documents.CabeceraNotaCreditoDebitoIce{
		NitEmisor:                    d.Cabecera.NitEmisor,
		RazonSocialEmisor:            d.Cabecera.RazonSocialEmisor,
		Municipio:                    d.Cabecera.Municipio,
		Telefono:                     datatype.Nilable[string]{Value: d.Cabecera.Telefono},
		NumeroNotaCreditoDebito:      d.Cabecera.NumeroNotaCreditoDebito,
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
		FechaEmisionFactura:          datatype.NewTimeSiat(d.Cabecera.FechaEmisionFactura),
		SubTotalNetoOriginal:         round(d.Cabecera.SubTotalNetoOriginal),
		TotalPrecioNetoOriginal:      round(d.Cabecera.TotalPrecioNetoOriginal),
		TotalIcePorcentualOriginal:   round(d.Cabecera.TotalIcePorcentualOriginal),
		TotalIceEspecificoOriginal:   round(d.Cabecera.TotalIceEspecificoOriginal),
		IvaOriginal:                  round(d.Cabecera.IvaOriginal),
		MontoAPagarOriginal:          round(d.Cabecera.MontoAPagarOriginal),
		SubTotalNetoDevuelto:         round(d.Cabecera.SubTotalNetoDevuelto),
		TotalPrecioNetoDevuelto:      round(d.Cabecera.TotalPrecioNetoDevuelto),
		IvaDevuelto:                  round(d.Cabecera.IvaDevuelto),
		MontoAPagarDevuelto:          round(d.Cabecera.MontoAPagarDevuelto),
		DescuentoAdicionalOriginal:   round(d.Cabecera.DescuentoAdicionalOriginal),
		MontoDescuentoDevuelto:       round(d.Cabecera.MontoDescuentoDevuelto),
		CodigoExcepcion:              datatype.Nilable[int]{Value: d.Cabecera.CodigoExcepcion},
		Leyenda:                      d.Cabecera.Leyenda,
		Usuario:                      d.Cabecera.Usuario,
		CodigoInterno:                d.Cabecera.CodigoInterno,
		CodigoDocumentoSector:        d.Cabecera.CodigoDocumentoSector,
	}
	if cab.CodigoDocumentoSector == 0 {
		cab.CodigoDocumentoSector = 48
	}
	dets := make([]documents.DetalleNotaCreditoDebitoIce, 0, len(d.Detalles))
	for _, item := range d.Detalles {
		dets = append(dets, documents.DetalleNotaCreditoDebitoIce{
			NroItem:                  item.NroItem,
			ActividadEconomica:       item.ActividadEconomica,
			CodigoProductoSin:        item.CodigoProductoSin,
			CodigoProducto:           item.CodigoProducto,
			Descripcion:              item.Descripcion,
			Cantidad:                 round5(item.Cantidad),
			UnidadMedida:             item.UnidadMedida,
			PrecioNetoUnitario:       round5(item.PrecioNetoUnitario),
			MontoDescuento:           datatype.Nilable[float64]{Value: roundOptional5(item.MontoDescuento)},
			SubTotal:                 round5(item.SubTotal),
			MarcaIce:                 item.MarcaIce,
			PrecioNetoVentaIce:       datatype.Nilable[float64]{Value: roundOptional5(item.PrecioNetoVentaIce)},
			AlicuotaEspecifica:       datatype.Nilable[float64]{Value: roundOptional5(item.AlicuotaEspecifica)},
			AlicuotaPorcentual:       datatype.Nilable[float64]{Value: roundOptional5(item.AlicuotaPorcentual)},
			MontoIceEspecifico:       datatype.Nilable[float64]{Value: roundOptional5(item.MontoIceEspecifico)},
			MontoIcePorcentual:       datatype.Nilable[float64]{Value: roundOptional5(item.MontoIcePorcentual)},
			CantidadIce:              datatype.Nilable[float64]{Value: roundOptional5(item.CantidadIce)},
			CodigoDetalleTransaccion: item.CodigoDetalleTransaccion,
		})
	}
	doc := documents.NotaCreditoDebitoIce{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: "/creditoDebito/" + nombre + ".xsd",
		Cabecera:          cab,
		Detalle:           dets,
	}
	return e.Encode(doc)
}
