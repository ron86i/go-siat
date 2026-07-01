package models

import (
	"github.com/ron86i/go-siat/v2/internal/core/domain/siat/compra_venta"
)

// --- Interfaces opacas para restringir el acceso a los atributos ---

type VentaAnexoCompraVenta struct {
	RequestWrapper[compra_venta.VentaAnexo]
}

// RecepcionAnexosCompraVenta representa una solicitud para la recepción de anexos de una factura.
type RecepcionAnexosCompraVenta struct {
	RequestWrapper[compra_venta.RecepcionAnexos]
}

// --- Constructores de Builders ---

func NewRecepcionAnexosBuilder() *recepcionAnexosBuilder {
	return &recepcionAnexosBuilder{
		request: &compra_venta.RecepcionAnexos{},
	}
}

// NewVentaAnexoBuilder crea el constructor para un ítem de anexo de venta.
func NewVentaAnexoBuilder() *ventaAnexoCompraVentaBuilder {
	return &ventaAnexoCompraVentaBuilder{
		anexo: &compra_venta.VentaAnexo{},
	}
}

// --- Implementaciones de Builders ---

type recepcionAnexosBuilder struct {
	request *compra_venta.RecepcionAnexos
}

// WithCodigoPuntoVenta establece el código del punto de venta.
func (b *recepcionAnexosBuilder) WithCodigoPuntoVenta(codigoPuntoVenta int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoPuntoVenta = codigoPuntoVenta
	return b
}

// WithCodigoSucursal establece el código de la sucursal.
func (b *recepcionAnexosBuilder) WithCodigoSucursal(codigoSucursal int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoSucursal = codigoSucursal
	return b
}

// WithCuis establece el Código Único de Inicio de Sistemas.
func (b *recepcionAnexosBuilder) WithCuis(cuis string) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cuis = cuis
	return b
}

// WithCufd establece el Código Único de Facturación Diaria.
func (b *recepcionAnexosBuilder) WithCufd(cufd string) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cufd = cufd
	return b
}

// WithCuf establece el Código Único de Facturación de la factura asociada.
func (b *recepcionAnexosBuilder) WithCuf(cuf string) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.Cuf = cuf
	return b
}

// WithCodigoDocumentoSector establece el código del documento sector.
func (b *recepcionAnexosBuilder) WithCodigoDocumentoSector(codigoDocumentoSector int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoDocumentoSector = codigoDocumentoSector
	return b
}

// WithCodigoEmision establece el tipo de emisión.
func (b *recepcionAnexosBuilder) WithCodigoEmision(codigoEmision int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.CodigoEmision = codigoEmision
	return b
}

// WithTipoFacturaDocumento establece el tipo de documento.
func (b *recepcionAnexosBuilder) WithTipoFacturaDocumento(tipoFacturaDocumento int) *recepcionAnexosBuilder {
	b.request.SolicitudRecepcionAnexos.TipoFacturaDocumento = tipoFacturaDocumento
	return b
}

// AddAnexos añade uno o más anexos de venta a la solicitud.
func (b *recepcionAnexosBuilder) AddAnexos(anexo ...VentaAnexoCompraVenta) *recepcionAnexosBuilder {
	for _, a := range anexo {
		if internal := UnwrapInternalRequest[compra_venta.VentaAnexo](a); internal != nil {
			b.request.SolicitudRecepcionAnexos.AnexosList = append(b.request.SolicitudRecepcionAnexos.AnexosList, *internal)
		}
	}
	return b
}

func (b *recepcionAnexosBuilder) Build() RecepcionAnexosCompraVenta {
	return RecepcionAnexosCompraVenta{RequestWrapper: NewRequestWrapper(b.request)}
}

type ventaAnexoCompraVentaBuilder struct {
	anexo *compra_venta.VentaAnexo
}

// WithCodigo configura el código del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithCodigo(codigo string) *ventaAnexoCompraVentaBuilder {
	b.anexo.Codigo = codigo
	return b
}

// WithCodigoProducto configura el código de producto del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithCodigoProducto(codigoProducto string) *ventaAnexoCompraVentaBuilder {
	b.anexo.CodigoProducto = codigoProducto
	return b
}

// WithCodigoProductoSin configura el código SIN del producto del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithCodigoProductoSin(codigoProductoSin int64) *ventaAnexoCompraVentaBuilder {
	b.anexo.CodigoProductoSin = codigoProductoSin
	return b
}

// WithTipoCodigo configura el tipo de código del anexo.
func (b *ventaAnexoCompraVentaBuilder) WithTipoCodigo(tipoCodigo string) *ventaAnexoCompraVentaBuilder {
	b.anexo.TipoCodigo = tipoCodigo
	return b
}

// Build finaliza la construcción del anexo retornando la estructura opaca.
func (b *ventaAnexoCompraVentaBuilder) Build() VentaAnexoCompraVenta {
	return VentaAnexoCompraVenta{RequestWrapper: NewRequestWrapper(b.anexo)}
}
