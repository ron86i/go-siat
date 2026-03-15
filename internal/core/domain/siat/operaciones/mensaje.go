package operaciones

// MensajeServicio representa un mensaje devuelto por el servidor del SIAT
type MensajeServicio struct {
	Codigo      int    `xml:"codigo" json:"codigo"`
	Descripcion string `xml:"descripcion" json:"descripcion"`
}
