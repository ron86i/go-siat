package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// FacturaCompraVenta representa la estructura completa de una factura lista para ser procesada.
type FacturaCompraVenta struct {
	requestWrapper[documentos.FacturaCompraVenta]
}

// FacturaCompraVentaCabecera representa la sección de cabecera de una factura de compra y venta.
type FacturaCompraVentaCabecera struct {
	requestWrapper[documentos.CabeceraCompraVenta]
}

// FacturaCompraVentaDetalle representa un ítem individual dentro del detalle de una factura.
type FacturaCompraVentaDetalle struct {
	requestWrapper[documentos.DetalleCompraVenta]
}

// NewFacturaCompraVentaBuilder inicia el proceso de construcción de una FacturaCompraVenta,
// configurando los nombres de nodo XML necesarios para el SIAT.
//
// Por defecto, la factura se configura para ser emitida electrónica.
func NewFacturaCompraVentaBuilder() *FacturaCompraVentaBuilder {
	return &FacturaCompraVentaBuilder{
		factura: &documentos.FacturaCompraVenta{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVenta"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVenta.xsd",
		},
	}
}

// NewFacturaCompraVentaCabeceraBuilder crea una nueva instancia del constructor para la cabecera
// de facturas de compra y venta.
func NewFacturaCompraVentaCabeceraBuilder() *FacturaCompraVentaCabeceraBuilder {
	return &FacturaCompraVentaCabeceraBuilder{
		cabecera: &documentos.CabeceraCompraVenta{},
	}
}

// NewFacturaCompraVentaDetalleBuilder crea una nueva instancia del constructor para los ítems
// de detalle de la factura.
func NewFacturaCompraVentaDetalleBuilder() *DetalleBuilder {
	return &DetalleBuilder{
		detalle: &documentos.DetalleCompraVenta{},
	}
}

type FacturaCompraVentaBuilder struct {
	factura *documentos.FacturaCompraVenta
}

// WithCabecera asocia la cabecera construida previamente a la factura.
func (b *FacturaCompraVentaBuilder) WithCabecera(req FacturaCompraVentaCabecera) *FacturaCompraVentaBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

// AddDetalle añade un ítem de detalle a la lista de detalles de la factura.
func (b *FacturaCompraVentaBuilder) AddDetalle(req FacturaCompraVentaDetalle) *FacturaCompraVentaBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}



// WithModalidad configura los metadatos XML de la factura según la modalidad (Electrónica o Computarizada).
func (b *FacturaCompraVentaBuilder) WithModalidad(tipo int) *FacturaCompraVentaBuilder {
	switch tipo {
	case ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVenta.xsd"
	case ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVenta.xsd"
	}
	return b
}

// Build finaliza la construcción y retorna la estructura opaca lista para ser firmada y enviada.
func (b *FacturaCompraVentaBuilder) Build() FacturaCompraVenta {
	return FacturaCompraVenta{requestWrapper[documentos.FacturaCompraVenta]{request: b.factura}}
}

// FacturaCompraVentaCabeceraBuilder ayuda a configurar la cabecera de la factura de compra y venta.
type FacturaCompraVentaCabeceraBuilder struct {
	cabecera *documentos.CabeceraCompraVenta
}

// WithNitEmisor configura el NIT del emisor de la factura.
func (b *FacturaCompraVentaCabeceraBuilder) WithNitEmisor(v int64) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

// WithRazonSocialEmisor configura la razón social o nombre del emisor.
func (b *FacturaCompraVentaCabeceraBuilder) WithRazonSocialEmisor(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

// WithMunicipio configura el municipio donde se emite la factura.
func (b *FacturaCompraVentaCabeceraBuilder) WithMunicipio(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

// WithTelefono configura el número de teléfono (opcional).
func (b *FacturaCompraVentaCabeceraBuilder) WithTelefono(telefono *string) *FacturaCompraVentaCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}

	// Creamos una copia física del valor en una nueva dirección de memoria
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}

	return b
}

// WithNumeroFactura configura el número correlativo de la factura.
func (b *FacturaCompraVentaCabeceraBuilder) WithNumeroFactura(v int64) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

// WithCuf configura el Código Único de Factura (CUF).
func (b *FacturaCompraVentaCabeceraBuilder) WithCuf(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

// WithCufd configura el Código Único de Facturación Diaria (CUFD).
func (b *FacturaCompraVentaCabeceraBuilder) WithCufd(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

// WithCodigoSucursal configura el código de la sucursal emisora (0 para casa matriz).
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoSucursal(v int) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

// WithDireccion configura la dirección física del establecimiento emisor.
func (b *FacturaCompraVentaCabeceraBuilder) WithDireccion(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

// WithCodigoPuntoVenta configura el código del punto de venta registrado ante el SIAT (nillable, pasar nil para omitir).
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

// WithFechaEmision configura la fecha y hora de emisión de la factura en formato string (ISO8601).
func (b *FacturaCompraVentaCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

// WithNombreRazonSocial configura el nombre o razón social del cliente (opcional).
func (b *FacturaCompraVentaCabeceraBuilder) WithNombreRazonSocial(v *string) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

// WithCodigoTipoDocumentoIdentidad configura el código del tipo de documento de identidad del cliente.
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

// WithNumeroDocumento configura el número de documento de identidad del cliente.
func (b *FacturaCompraVentaCabeceraBuilder) WithNumeroDocumento(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

// WithComplemento configura el complemento del documento de identidad (opcional).
func (b *FacturaCompraVentaCabeceraBuilder) WithComplemento(v *string) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

// WithCodigoCliente configura un código interno único para identificar al cliente.
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoCliente(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

// WithCodigoMetodoPago configura el código de la paramétrica del método de pago utilizado.
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoMetodoPago(v int) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

// WithNumeroTarjeta configura los primeros y últimos dígitos de la tarjeta (opcional, enmascarado).
func (b *FacturaCompraVentaCabeceraBuilder) WithNumeroTarjeta(v *int64) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

// WithMontoTotal configura el monto total de la factura, redondeado automáticamente a 2 decimales.
func (b *FacturaCompraVentaCabeceraBuilder) WithMontoTotal(v float64) *FacturaCompraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

// WithMontoTotalSujetoIva configura el monto base para el crédito fiscal IVA, redondeado automáticamente.
func (b *FacturaCompraVentaCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *FacturaCompraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoMoneda configura el código de la moneda utilizada (ej. 1 para Bolivianos).
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoMoneda(v int) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

// WithTipoCambio configura el tipo de cambio respecto al Boliviano, redondeado automáticamente.
func (b *FacturaCompraVentaCabeceraBuilder) WithTipoCambio(v float64) *FacturaCompraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

// WithMontoTotalMoneda configura el monto total expresado en la moneda original, redondeado automáticamente.
func (b *FacturaCompraVentaCabeceraBuilder) WithMontoTotalMoneda(v float64) *FacturaCompraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

// WithMontoGiftCard configura el monto pagado con tarjeta de regalo o prepago (opcional, redondeado).
func (b *FacturaCompraVentaCabeceraBuilder) WithMontoGiftCard(v *float64) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}

	value := *v
	// Asegurar que el valor sea redondeado a 2 decimales
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

// WithDescuentoAdicional configura un descuento global aplicado a toda la factura (opcional, redondeado).
func (b *FacturaCompraVentaCabeceraBuilder) WithDescuentoAdicional(v *float64) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}

	value := *v
	// Asegurar que el valor sea redondeado a 2 decimales
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

// WithCodigoExcepcion configura un código de excepción si los datos del cliente son inválidos pero autorizados (opcional).
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoExcepcion(v *int64) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &value}
	return b
}

// WithCafc configura el Código de Autorización de Facturas por Contingencia (opcional).
func (b *FacturaCompraVentaCabeceraBuilder) WithCafc(v *string) *FacturaCompraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

// WithLeyenda configura la leyenda obligatoria según la normativa de Impuestos Nacionales.
func (b *FacturaCompraVentaCabeceraBuilder) WithLeyenda(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

// WithUsuario configura el identificador del usuario que emite la factura.
func (b *FacturaCompraVentaCabeceraBuilder) WithUsuario(v string) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *FacturaCompraVentaCabeceraBuilder) WithCodigoDocumentoSector(v int) *FacturaCompraVentaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

// Build finaliza la construcción de la cabecera retornando la estructura opaca.
func (b *FacturaCompraVentaCabeceraBuilder) Build() FacturaCompraVentaCabecera {
	return FacturaCompraVentaCabecera{requestWrapper[documentos.CabeceraCompraVenta]{request: b.cabecera}}
}

// DetalleBuilder ayuda a configurar un ítem individual de detalle de la factura.
type DetalleBuilder struct {
	detalle *documentos.DetalleCompraVenta
}

// WithActividadEconomica configura el código de actividad económica asociado al producto/servicio.
func (b *DetalleBuilder) WithActividadEconomica(v string) *DetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

// WithCodigoProductoSin configura el código de producto según el catálogo del SIAT (integer).
func (b *DetalleBuilder) WithCodigoProductoSin(v int64) *DetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

// WithCodigoProducto configura un código interno propio de la empresa para el producto.
func (b *DetalleBuilder) WithCodigoProducto(v string) *DetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

// WithDescripcion configura la descripción detallada del artículo o servicio.
func (b *DetalleBuilder) WithDescripcion(v string) *DetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

// WithCantidad configura la cantidad vendida, según XSD acepta hasta 5 decimales.
func (b *DetalleBuilder) WithCantidad(v float64) *DetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

// WithUnidadMedida configura el código de la paramétrica de unidad de medida (ej. 1 para unidades).
func (b *DetalleBuilder) WithUnidadMedida(v int) *DetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

// WithPrecioUnitario configura el precio unitario, según XSD acepta hasta 5 decimales.
func (b *DetalleBuilder) WithPrecioUnitario(v float64) *DetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

// WithMontoDescuento configura un descuento aplicado al ítem (opcional). Acepta hasta 5 decimales.
func (b *DetalleBuilder) WithMontoDescuento(v *float64) *DetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}

	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 5, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

// WithSubTotal configura el subtotal del ítem. Acepta hasta 5 decimales.
func (b *DetalleBuilder) WithSubTotal(v float64) *DetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

// WithNumeroSerie configura el número de serie del producto (opcional).
func (b *DetalleBuilder) WithNumeroSerie(v *string) *DetalleBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &value}
	return b
}

// WithNumeroImei configura el número IMEI si se trata de equipos telefónicos (opcional).
func (b *DetalleBuilder) WithNumeroImei(v *string) *DetalleBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &value}
	return b
}

// Build finaliza la construcción del detalle retornando la estructura opaca.
func (b *DetalleBuilder) Build() FacturaCompraVentaDetalle {
	return FacturaCompraVentaDetalle{requestWrapper[documentos.DetalleCompraVenta]{request: b.detalle}}
}
