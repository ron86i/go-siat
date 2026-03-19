package siat

const (
	// ModalidadElectronica requiere firma digital de los documentos XML.
	ModalidadElectronica = 1
	// ModalidadComputarizada no requiere firma digital, usa un código de control.
	ModalidadComputarizada = 2
	// AmbienteProduccion para operaciones reales con validez tributaria.
	AmbienteProduccion = 1
	// AmbientePruebas para entornos de desarrollo y certificación.
	AmbientePruebas = 2
	// EmisionOnline emisión se realizó en línea
	EmisionOnline = 1
	// EmisionOffline emisión se realizó fuera de línea
	EmisionOffline = 2
	// EmisionMasiva para emisión masiva de factura
	EmisionMasiva = 3
)
