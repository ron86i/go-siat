// Package alcanzadaice proporciona documentos para el sector Alcanzada por el ICE en SIAT v3.
package alcanzadaice

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
	TotalIceEspecifico           *float64
	TotalIcePorcentual           *float64
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

// Detalle contiene los ítems del documento fiscal v3.
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
	MarcaIce           int
	PrecioNetoVentaIce *float64
	AlicuotaEspecifica *float64
	AlicuotaPorcentual *float64
	MontoIceEspecifico *float64
	MontoIcePorcentual *float64
	CantidadIce        *float64
}

func (d Factura) Validate() error {
	if d.Modalidad != 0 && d.Modalidad != 1 && d.Modalidad != 2 {
		return fmt.Errorf("modalidad %d no soportada", d.Modalidad)
	}
	if d.Cabecera.CodigoDocumentoSector != 0 && d.Cabecera.CodigoDocumentoSector != 14 {
		return fmt.Errorf("código de sector %d no corresponde a Alcanzada por el ICE (esperado 14)", d.Cabecera.CodigoDocumentoSector)
	}
	return d.validateImportes()
}

func (d Factura) validateImportes() error {
	if len(d.Detalles) == 0 {
		return nil
	}

	var subTotalNeto, totalIceEspecifico, totalIcePorcentual decimal.Decimal
	for i, item := range d.Detalles {
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
		subTotalNeto = subTotalNeto.Add(amount5(item.SubTotal))

		especifico, porcentual, err := validateICE(item, subTotal, i)
		if err != nil {
			return err
		}
		totalIceEspecifico = totalIceEspecifico.Add(especifico)
		totalIcePorcentual = totalIcePorcentual.Add(porcentual)
	}

	c := d.Cabecera
	subTotalNeto = subTotalNeto.Round(2)
	if err := validateAmount("cabecera.subTotalNeto", amount2(c.SubTotalNeto), subTotalNeto, 2); err != nil {
		return err
	}
	totalPrecioNeto := subTotalNeto.Sub(optionalAmount2(c.DescuentoAdicional)).Round(2)
	if err := validateAmount("cabecera.totalPrecioNeto", amount2(c.TotalPrecioNeto), totalPrecioNeto, 2); err != nil {
		return err
	}
	iva := totalPrecioNeto.Mul(decimal.NewFromFloat(0.13)).Round(2)
	if err := validateAmount("cabecera.iva", amount2(c.Iva), iva, 2); err != nil {
		return err
	}
	if err := validateOptionalAmount("cabecera.totalIceEspecifico", c.TotalIceEspecifico, totalIceEspecifico, 2); err != nil {
		return err
	}
	if err := validateOptionalAmount("cabecera.totalIcePorcentual", c.TotalIcePorcentual, totalIcePorcentual, 2); err != nil {
		return err
	}
	montoAPagar := totalPrecioNeto.Add(iva).Add(totalIceEspecifico.Round(2)).Add(totalIcePorcentual.Round(2)).Round(2)
	if err := validateAmount("cabecera.montoaPagar", amount2(c.MontoAPagar), montoAPagar, 2); err != nil {
		return err
	}
	tipoCambio := amount2(c.TipoCambio)
	if tipoCambio.Sign() <= 0 {
		return fmt.Errorf("cabecera.tipoCambio debe ser positivo al redondear a dos decimales")
	}
	montoTotalMoneda := montoAPagar.Div(tipoCambio).Round(2)
	if err := validateAmount("cabecera.montoTotalMoneda", amount2(c.MontoTotalMoneda), montoTotalMoneda, 2); err != nil {
		return err
	}
	return nil
}

func validateICE(item Detalle, subTotal decimal.Decimal, index int) (decimal.Decimal, decimal.Decimal, error) {
	var especifico, porcentual decimal.Decimal
	if item.AlicuotaEspecifica != nil || item.MontoIceEspecifico != nil {
		if item.AlicuotaEspecifica == nil || item.MontoIceEspecifico == nil || item.CantidadIce == nil {
			return especifico, porcentual, fmt.Errorf("detalles[%d]: alícuota específica, cantidad ICE e importe ICE deben informarse juntos", index)
		}
		expected := amount5(*item.CantidadIce).Mul(amount5(*item.AlicuotaEspecifica))
		if err := validateAmount(fmt.Sprintf("detalles[%d].montoIceEspecifico", index), amount5(*item.MontoIceEspecifico), expected, 2); err != nil {
			return especifico, porcentual, err
		}
		especifico = amount5(*item.MontoIceEspecifico)
	}
	if item.AlicuotaPorcentual != nil || item.MontoIcePorcentual != nil {
		if item.AlicuotaPorcentual == nil || item.MontoIcePorcentual == nil {
			return especifico, porcentual, fmt.Errorf("detalles[%d]: alícuota porcentual e importe ICE deben informarse juntos", index)
		}
		expected := subTotal.Mul(amount5(*item.AlicuotaPorcentual))
		if err := validateAmount(fmt.Sprintf("detalles[%d].montoIcePorcentual", index), amount5(*item.MontoIcePorcentual), expected, 2); err != nil {
			return especifico, porcentual, err
		}
		porcentual = amount5(*item.MontoIcePorcentual)
	}
	return especifico, porcentual, nil
}

func amount2(v float64) decimal.Decimal { return decimal.NewFromFloat(v).Round(2) }
func amount5(v float64) decimal.Decimal { return decimal.NewFromFloat(v).Round(5) }
func optionalAmount2(v *float64) decimal.Decimal {
	if v == nil {
		return decimal.Zero
	}
	return amount2(*v)
}
func optionalAmount5(v *float64) decimal.Decimal {
	if v == nil {
		return decimal.Zero
	}
	return amount5(*v)
}
func validateOptionalAmount(field string, actual *float64, expected decimal.Decimal, places int32) error {
	if actual == nil {
		if expected.Round(places).IsZero() {
			return nil
		}
		return fmt.Errorf("%s debe reflejar el total calculado %s", field, expected.StringFixed(places))
	}
	return validateAmount(field, decimal.NewFromFloat(*actual), expected, places)
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

func (d Factura) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := d.Validate(); err != nil {
		return err
	}
	nombre := "facturaElectronicaAlcanzadaIce"
	if d.Modalidad == 2 {
		nombre = "facturaComputarizadaAlcanzadaIce"
	}
	cab := documents.CabeceraAlcanzadaIce{
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
		TotalIceEspecifico:           datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.TotalIceEspecifico)},
		TotalIcePorcentual:           datatype.Nilable[float64]{Value: roundOptional(d.Cabecera.TotalIcePorcentual)},
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
		cab.CodigoDocumentoSector = 14
	}
	dets := make([]documents.DetalleAlcanzadaIce, 0, len(d.Detalles))
	for _, item := range d.Detalles {
		dets = append(dets, documents.DetalleAlcanzadaIce{
			ActividadEconomica: item.ActividadEconomica,
			CodigoProductoSin:  item.CodigoProductoSin,
			CodigoProducto:     item.CodigoProducto,
			Descripcion:        item.Descripcion,
			Cantidad:           round5(item.Cantidad),
			UnidadMedida:       item.UnidadMedida,
			PrecioNetoUnitario: round5(item.PrecioNetoUnitario),
			MontoDescuento:     datatype.Nilable[float64]{Value: roundOptional5(item.MontoDescuento)},
			SubTotal:           round5(item.SubTotal),
			MarcaIce:           item.MarcaIce,
			PrecioNetoVentaIce: datatype.Nilable[float64]{Value: roundOptional5(item.PrecioNetoVentaIce)},
			AlicuotaEspecifica: datatype.Nilable[float64]{Value: roundOptional5(item.AlicuotaEspecifica)},
			AlicuotaPorcentual: datatype.Nilable[float64]{Value: roundOptional5(item.AlicuotaPorcentual)},
			MontoIceEspecifico: datatype.Nilable[float64]{Value: roundOptional5(item.MontoIceEspecifico)},
			MontoIcePorcentual: datatype.Nilable[float64]{Value: roundOptional5(item.MontoIcePorcentual)},
			CantidadIce:        datatype.Nilable[float64]{Value: roundOptional5(item.CantidadIce)},
		})
	}
	doc := documents.FacturaAlcanzadaIce{
		XMLName:           xml.Name{Local: nombre},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: "/alcanzadaIce/" + nombre + ".xsd",
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
	return 14
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
