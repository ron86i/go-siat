package siat

import "github.com/ron86i/go-siat/v2/pkg/models"

const (
	// AmbienteProduccion (1): Operaciones reales con validez tributaria.
	AmbienteProduccion = models.AmbienteProduccion
	// AmbientePruebas (2): Entorno de desarrollo, pruebas y certificación.
	AmbientePruebas = models.AmbientePruebas
)

const (
	// ModalidadElectronica (1): Requiere firma digital de los documentos XML.
	ModalidadElectronica = models.ModalidadElectronica
	// ModalidadComputarizada (2): No requiere firma digital, usa código de control.
	ModalidadComputarizada = models.ModalidadComputarizada
)

const (
	// EmisionOnline (1): La emisión se realizó con conexión al SIAT.
	EmisionOnline = models.EmisionOnline
	// EmisionOffline (2): La emisión se realizó fuera de línea (Contingencia).
	EmisionOffline = models.EmisionOffline
	// EmisionMasiva (3): Para procesos de alta demanda de facturación.
	EmisionMasiva = models.EmisionMasiva
)

// Tipos de documento fiscal definidos por el SIAT.
const (
	// Estos alias mantienen la API del paquete raíz; la definición canónica
	// está en pkg/models, que es donde la consumen los builders de facturas.
	TipoFacturaConDerechoCreditoFiscal = models.TipoFacturaConDerechoCreditoFiscal
	TipoFacturaSinDerechoCreditoFiscal = models.TipoFacturaSinDerechoCreditoFiscal
	TipoNotaCreditoDebito              = models.TipoNotaCreditoDebito
	TipoBoletoAereo                    = models.TipoBoletoAereo
)
