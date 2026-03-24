package siat

const (
	// AmbienteProduccion (1): Operaciones reales con validez tributaria.
	AmbienteProduccion = iota + 1
	// AmbientePruebas (2): Entorno de desarrollo, pruebas y certificación.
	AmbientePruebas
)

const (
	// ModalidadElectronica (1): Requiere firma digital de los documentos XML.
	ModalidadElectronica = iota + 1
	// ModalidadComputarizada (2): No requiere firma digital, usa código de control.
	ModalidadComputarizada
)

const (
	// EmisionOnline (1): La emisión se realizó con conexión al SIAT.
	EmisionOnline = iota + 1
	// EmisionOffline (2): La emisión se realizó fuera de línea (Contingencia).
	EmisionOffline
	// EmisionMasiva (3): Para procesos de alta demanda de facturación.
	EmisionMasiva
)
