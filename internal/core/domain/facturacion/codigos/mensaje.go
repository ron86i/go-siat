package codigos

// MensajeServicio representa una estructura de notificación o error devuelta por los servicios del SIAT.
// Contiene un código identificador y una descripción detallada que permite entender el resultado de una operación.
type MensajeServicio struct {
	Codigo      int    `xml:"codigo" json:"codigo"`
	Descripcion string `xml:"descripcion" json:"descripcion"`
}
