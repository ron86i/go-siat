package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// CompraVenta representa la estructura completa de una factura lista para ser procesada.
type CompraVenta struct {
	requestWrapper[documentos.FacturaCompraVenta]
}

// CompraVentaCabecera representa la sección de cabecera de una factura de compra y venta.
type CompraVentaCabecera struct {
	requestWrapper[documentos.CabeceraCompraVenta]
}

// CompraVentaDetalle representa un ítem individual dentro del detalle de una factura.
type CompraVentaDetalle struct {
	requestWrapper[documentos.DetalleCompraVenta]
}

// NewCompraVentaBuilder inicia el proceso de construcción de una CompraVenta,
// configurando los nombres de nodo XML necesarios para el SIAT.
//
// Por defecto, la factura se configura para ser emitida electrónica.
func NewCompraVentaBuilder() *compraVentaBuilder {
	return &compraVentaBuilder{
		factura: &documentos.FacturaCompraVenta{
			XMLName:           xml.Name{Local: "facturaElectronicaCompraVenta"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaCompraVenta.xsd",
		},
	}
}

// NewCompraVentaCabeceraBuilder crea una nueva instancia del constructor para la cabecera
// de facturas de compra y venta.
func NewCompraVentaCabeceraBuilder() *compraVentaCabeceraBuilder {
	return &compraVentaCabeceraBuilder{
		cabecera: &documentos.CabeceraCompraVenta{},
	}
}

// NewCompraVentaDetalleBuilder crea una nueva instancia del constructor para los ítems
// de detalle de la factura.
func NewCompraVentaDetalleBuilder() *detalleBuilder {
	return &detalleBuilder{
		detalle: &documentos.DetalleCompraVenta{},
	}
}

type compraVentaBuilder struct {
	factura *documentos.FacturaCompraVenta
}

// WithCabecera asocia la cabecera construida previamente a la factura.
func (b *compraVentaBuilder) WithCabecera(req CompraVentaCabecera) *compraVentaBuilder {
	if req.request != nil {
		b.factura.Cabecera = *req.request
	}
	return b
}

// AddDetalle añade un ítem de detalle a la lista de detalles de la factura.
func (b *compraVentaBuilder) AddDetalle(req CompraVentaDetalle) *compraVentaBuilder {
	if req.request != nil {
		b.factura.Detalle = append(b.factura.Detalle, *req.request)
	}
	return b
}

// WithModalidad configura los metadatos XML de la factura según la modalidad (Electrónica o Computarizada).
func (b *compraVentaBuilder) WithModalidad(tipo int) *compraVentaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaElectronicaCompraVenta.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaCompraVenta"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaCompraVenta.xsd"
	}
	return b
}

// Build finaliza la construcción y retorna la estructura opaca lista para ser firmada y enviada.
func (b *compraVentaBuilder) Build() CompraVenta {
	return CompraVenta{requestWrapper[documentos.FacturaCompraVenta]{request: b.factura}}
}

// compraVentaCabeceraBuilder ayuda a configurar la cabecera de la factura de compra y venta.
type compraVentaCabeceraBuilder struct {
	cabecera *documentos.CabeceraCompraVenta
}

// WithNitEmisor configura el NIT del emisor de la factura.
func (b *compraVentaCabeceraBuilder) WithNitEmisor(v int64) *compraVentaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

// WithRazonSocialEmisor configura la razón social o nombre del emisor.
func (b *compraVentaCabeceraBuilder) WithRazonSocialEmisor(v string) *compraVentaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

// WithMunicipio configura el municipio donde se emite la factura.
func (b *compraVentaCabeceraBuilder) WithMunicipio(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

// WithTelefono configura el número de teléfono (opcional).
func (b *compraVentaCabeceraBuilder) WithTelefono(telefono *string) *compraVentaCabeceraBuilder {
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
func (b *compraVentaCabeceraBuilder) WithNumeroFactura(v int64) *compraVentaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

// WithCuf configura el Código Único de Factura (CUF).
func (b *compraVentaCabeceraBuilder) WithCuf(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

// WithCufd configura el Código Único de Facturación Diaria (CUFD).
func (b *compraVentaCabeceraBuilder) WithCufd(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

// WithCodigoSucursal configura el código de la sucursal emisora (0 para casa matriz).
func (b *compraVentaCabeceraBuilder) WithCodigoSucursal(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

// WithDireccion configura la dirección física del establecimiento emisor.
func (b *compraVentaCabeceraBuilder) WithDireccion(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

// WithCodigoPuntoVenta configura el código del punto de venta registrado ante el SIAT (nillable, pasar nil para omitir).
func (b *compraVentaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

// WithFechaEmision configura la fecha y hora de emisión de la factura en formato string (ISO8601).
func (b *compraVentaCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *compraVentaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

// WithNombreRazonSocial configura el nombre o razón social del cliente (opcional).
func (b *compraVentaCabeceraBuilder) WithNombreRazonSocial(v *string) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

// WithCodigoTipoDocumentoIdentidad configura el código del tipo de documento de identidad del cliente.
func (b *compraVentaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

// WithNumeroDocumento configura el número de documento de identidad del cliente.
func (b *compraVentaCabeceraBuilder) WithNumeroDocumento(v string) *compraVentaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

// WithComplemento configura el complemento del documento de identidad (opcional).
func (b *compraVentaCabeceraBuilder) WithComplemento(v *string) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

// WithCodigoCliente configura un código interno único para identificar al cliente.
func (b *compraVentaCabeceraBuilder) WithCodigoCliente(v string) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

// WithCodigoMetodoPago configura el código de la paramétrica del método de pago utilizado.
func (b *compraVentaCabeceraBuilder) WithCodigoMetodoPago(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

// WithNumeroTarjeta configura los primeros y últimos dígitos de la tarjeta (opcional, enmascarado).
func (b *compraVentaCabeceraBuilder) WithNumeroTarjeta(v *int64) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

// WithMontoTotal configura el monto total de la factura, redondeado automáticamente a 2 decimales.
func (b *compraVentaCabeceraBuilder) WithMontoTotal(v float64) *compraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

// WithMontoTotalSujetoIva configura el monto base para el crédito fiscal IVA, redondeado automáticamente.
func (b *compraVentaCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *compraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

// WithCodigoMoneda configura el código de la moneda utilizada (ej. 1 para Bolivianos).
func (b *compraVentaCabeceraBuilder) WithCodigoMoneda(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

// WithTipoCambio configura el tipo de cambio respecto al Boliviano, redondeado automáticamente.
func (b *compraVentaCabeceraBuilder) WithTipoCambio(v float64) *compraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

// WithMontoTotalMoneda configura el monto total expresado en la moneda original, redondeado automáticamente.
func (b *compraVentaCabeceraBuilder) WithMontoTotalMoneda(v float64) *compraVentaCabeceraBuilder {
	// Asegurar que el valor sea redondeado a 2 decimales
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

// WithMontoGiftCard configura el monto pagado con tarjeta de regalo o prepago (opcional, redondeado).
func (b *compraVentaCabeceraBuilder) WithMontoGiftCard(v *float64) *compraVentaCabeceraBuilder {
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
func (b *compraVentaCabeceraBuilder) WithDescuentoAdicional(v *float64) *compraVentaCabeceraBuilder {
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
func (b *compraVentaCabeceraBuilder) WithCodigoExcepcion(v *int64) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int64]{Value: &value}
	return b
}

// WithCafc configura el Código de Autorización de Facturas por Contingencia (opcional).
func (b *compraVentaCabeceraBuilder) WithCafc(v *string) *compraVentaCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

// WithLeyenda configura la leyenda obligatoria según la normativa de Impuestos Nacionales.
func (b *compraVentaCabeceraBuilder) WithLeyenda(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

// WithUsuario configura el identificador del usuario que emite la factura.
func (b *compraVentaCabeceraBuilder) WithUsuario(v string) *compraVentaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

// WithCodigoDocumentoSector configura el código que identifica el diseño o sector de la factura.
func (b *compraVentaCabeceraBuilder) WithCodigoDocumentoSector(v int) *compraVentaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

// Build finaliza la construcción de la cabecera retornando la estructura opaca.
func (b *compraVentaCabeceraBuilder) Build() CompraVentaCabecera {
	return CompraVentaCabecera{requestWrapper[documentos.CabeceraCompraVenta]{request: b.cabecera}}
}

// detalleBuilder ayuda a configurar un ítem individual de detalle de la factura.
type detalleBuilder struct {
	detalle *documentos.DetalleCompraVenta
}

// WithActividadEconomica configura el código de actividad económica asociado al producto/servicio.
func (b *detalleBuilder) WithActividadEconomica(v string) *detalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

// WithCodigoProductoSin configura el código de producto según el catálogo del SIAT (integer).
func (b *detalleBuilder) WithCodigoProductoSin(v int64) *detalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

// WithCodigoProducto configura un código interno propio de la empresa para el producto.
func (b *detalleBuilder) WithCodigoProducto(v string) *detalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

// WithDescripcion configura la descripción detallada del artículo o servicio.
func (b *detalleBuilder) WithDescripcion(v string) *detalleBuilder {
	b.detalle.Descripcion = v
	return b
}

// WithCantidad configura la cantidad vendida, según XSD acepta hasta 5 decimales.
func (b *detalleBuilder) WithCantidad(v float64) *detalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.Cantidad = v
	return b
}

// WithUnidadMedida configura el código de la paramétrica de unidad de medida (ej. 1 para unidades).
func (b *detalleBuilder) WithUnidadMedida(v int) *detalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

// WithPrecioUnitario configura el precio unitario, según XSD acepta hasta 5 decimales.
func (b *detalleBuilder) WithPrecioUnitario(v float64) *detalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

// WithMontoDescuento configura un descuento aplicado al ítem (opcional). Acepta hasta 5 decimales.
func (b *detalleBuilder) WithMontoDescuento(v *float64) *detalleBuilder {
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
func (b *detalleBuilder) WithSubTotal(v float64) *detalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 5, 64), 64)
	b.detalle.SubTotal = v
	return b
}

// WithNumeroSerie configura el número de serie del producto (opcional).
func (b *detalleBuilder) WithNumeroSerie(v *string) *detalleBuilder {
	if v == nil {
		b.detalle.NumeroSerie = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.detalle.NumeroSerie = datatype.Nilable[string]{Value: &value}
	return b
}

// WithNumeroImei configura el número IMEI si se trata de equipos telefónicos (opcional).
func (b *detalleBuilder) WithNumeroImei(v *string) *detalleBuilder {
	if v == nil {
		b.detalle.NumeroImei = datatype.Nilable[string]{Value: nil}
		return b
	}

	value := *v
	b.detalle.NumeroImei = datatype.Nilable[string]{Value: &value}
	return b
}

// Build finaliza la construcción del detalle retornando la estructura opaca.
func (b *detalleBuilder) Build() CompraVentaDetalle {
	return CompraVentaDetalle{requestWrapper[documentos.DetalleCompraVenta]{request: b.detalle}}
}
