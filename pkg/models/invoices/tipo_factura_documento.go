package invoices

import "github.com/ron86i/go-siat/v2/pkg/models"

// DefinicionTipoFactura se mantiene como alias para compatibilidad con los
// consumidores del paquete invoices. El catálogo vive en models, junto a los
// metadatos que lo consumen.
type DefinicionTipoFactura = models.DefinicionTipoFactura

// DefinicionesTipoFactura devuelve el catálogo fiscal compartido.
func DefinicionesTipoFactura() []DefinicionTipoFactura {
	return models.DefinicionesTipoFactura()
}

// Los builders mantienen los metadatos en la factura construida. Estas
// funciones permiten sobrescribir el tipo fiscal cuando una operación lo
// requiere, sin ocultar los métodos promovidos de MetadatosFactura.
func (b *libreConsignacionBuilder) WithTipoFacturaDocumento(tipo int) *libreConsignacionBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *zonaFrancaBuilder) WithTipoFacturaDocumento(tipo int) *zonaFrancaBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *seguridadAlimentariaBuilder) WithTipoFacturaDocumento(tipo int) *seguridadAlimentariaBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *monedaExtranjeraBuilder) WithTipoFacturaDocumento(tipo int) *monedaExtranjeraBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *duttyFreeBuilder) WithTipoFacturaDocumento(tipo int) *duttyFreeBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercializacionHidroBuilder) WithTipoFacturaDocumento(tipo int) *comercializacionHidroBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *servicioBasicoBuilder) WithTipoFacturaDocumento(tipo int) *servicioBasicoBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *alcanzadaIceBuilder) WithTipoFacturaDocumento(tipo int) *alcanzadaIceBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *entidadFinancieraBuilder) WithTipoFacturaDocumento(tipo int) *entidadFinancieraBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *juegoAzarBuilder) WithTipoFacturaDocumento(tipo int) *juegoAzarBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *hidrocarburoAlcanzadaIehdBuilder) WithTipoFacturaDocumento(tipo int) *hidrocarburoAlcanzadaIehdBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercialExportacionMineraBuilder) WithTipoFacturaDocumento(tipo int) *comercialExportacionMineraBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *ventaMineralBuilder) WithTipoFacturaDocumento(tipo int) *ventaMineralBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *telecomunicacionesBuilder) WithTipoFacturaDocumento(tipo int) *telecomunicacionesBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *prevaloradaBuilder) WithTipoFacturaDocumento(tipo int) *prevaloradaBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *notaCreditoDebitoBuilder) WithTipoFacturaDocumento(tipo int) *notaCreditoDebitoBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *notaFiscalCreditoDebitoBuilder) WithTipoFacturaDocumento(tipo int) *notaFiscalCreditoDebitoBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercialExportacionServicioBuilder) WithTipoFacturaDocumento(tipo int) *comercialExportacionServicioBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *notaConciliacionBuilder) WithTipoFacturaDocumento(tipo int) *notaConciliacionBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *boletoAereoBuilder) WithTipoFacturaDocumento(tipo int) *boletoAereoBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *suministroEnergiaBuilder) WithTipoFacturaDocumento(tipo int) *suministroEnergiaBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *segurosBuilder) WithTipoFacturaDocumento(tipo int) *segurosBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *prevaloradaSinDerechoCreditoFiscalBuilder) WithTipoFacturaDocumento(tipo int) *prevaloradaSinDerechoCreditoFiscalBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercializacionGnvBuilder) WithTipoFacturaDocumento(tipo int) *comercializacionGnvBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *hidrocarburoNoAlcanzadaIehdBuilder) WithTipoFacturaDocumento(tipo int) *hidrocarburoNoAlcanzadaIehdBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercializacionGnGlpBuilder) WithTipoFacturaDocumento(tipo int) *comercializacionGnGlpBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *servicioBasicoZFBuilder) WithTipoFacturaDocumento(tipo int) *servicioBasicoZFBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *alquilerZFBuilder) WithTipoFacturaDocumento(tipo int) *alquilerZFBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercialExportacionHidroBuilder) WithTipoFacturaDocumento(tipo int) *comercialExportacionHidroBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *importacionComercializacionLubricantesBuilder) WithTipoFacturaDocumento(tipo int) *importacionComercializacionLubricantesBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *comercialExportacionPVentaBuilder) WithTipoFacturaDocumento(tipo int) *comercialExportacionPVentaBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *sectorEducativoZFBuilder) WithTipoFacturaDocumento(tipo int) *sectorEducativoZFBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *notaCreditoDebitoDescuentoBuilder) WithTipoFacturaDocumento(tipo int) *notaCreditoDebitoDescuentoBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *notaCreditoDebitoIceBuilder) WithTipoFacturaDocumento(tipo int) *notaCreditoDebitoIceBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *telecomunicacionesZFBuilder) WithTipoFacturaDocumento(tipo int) *telecomunicacionesZFBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *hospitalClinicaZFBuilder) WithTipoFacturaDocumento(tipo int) *hospitalClinicaZFBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *engarrafadorasBuilder) WithTipoFacturaDocumento(tipo int) *engarrafadorasBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *ventaMineralBCBBuilder) WithTipoFacturaDocumento(tipo int) *ventaMineralBCBBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *lubricantesIehdBuilder) WithTipoFacturaDocumento(tipo int) *lubricantesIehdBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *biodieselBuilder) WithTipoFacturaDocumento(tipo int) *biodieselBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}

func (b *ventaCombustibleSinSubvencionBuilder) WithTipoFacturaDocumento(tipo int) *ventaCombustibleSinSubvencionBuilder {
	b.metadatos.WithTipoFacturaDocumento(tipo)
	return b
}
