package compra_venta

import (
	"encoding/json"
	"encoding/xml"
)

// Nilable es un tipo genérico para manejar valores que pueden ser nulos en el XML del SIAT.
// Es necesario porque el SIAT requiere que ciertos campos opcionales se envíen con el atributo
// xsi:nil="true" cuando no tienen valor, en lugar de omitir la etiqueta por completo.
type Nilable[T any] struct {
	Value *T
}

// MarshalXML implementa la interfaz xml.Marshaler para manejar la nulidad explícita con xsi:nil.
func (n Nilable[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if n.Value == nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xsi:nil"}, Value: "true"})
		return e.EncodeElement("", start)
	}
	return e.EncodeElement(*n.Value, start)
}

// UnmarshalXML implementa la interfaz xml.Unmarshaler para manejar la nulidad explícita con xsi:nil.
func (n *Nilable[T]) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var isNil bool
	for _, attr := range start.Attr {
		if attr.Name.Local == "nil" && attr.Value == "true" {
			isNil = true
			break
		}
	}

	if isNil {
		n.Value = nil
		return d.Skip()
	}

	var val T
	if err := d.DecodeElement(&val, &start); err != nil {
		return err
	}
	n.Value = &val
	return nil
}

// MarshalJSON implementa la interfaz json.Marshaler.
func (n Nilable[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Value)
}

// UnmarshalJSON implementa la interfaz json.Unmarshaler.
func (n *Nilable[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.Value)
}

// ------------------ Estructuras de la Factura ------------------

// FacturaCompraVenta representa la estructura completa de una factura de compra-venta para el SIAT.
type FacturaCompraVenta struct {
	XMLName           xml.Name  `json:"-"`
	XmlnsXsi          string    `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string    `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          Cabecera  `xml:"cabecera" json:"cabecera"`
	Detalle           []Detalle `xml:"detalle" json:"detalle"`
}

// Cabecera contiene la información general y del cliente de la factura.
type Cabecera struct {
	NitEmisor                    int64            `xml:"nitEmisor" json:"nitEmisor"`
	RazonSocialEmisor            string           `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	Municipio                    string           `xml:"municipio" json:"municipio"`
	Telefono                     Nilable[string]  `xml:"telefono" json:"telefono"`
	NumeroFactura                int64            `xml:"numeroFactura" json:"numeroFactura"`
	Cuf                          string           `xml:"cuf" json:"cuf"`
	Cufd                         string           `xml:"cufd" json:"cufd"`
	CodigoSucursal               int              `xml:"codigoSucursal" json:"codigoSucursal"`
	Direccion                    string           `xml:"direccion" json:"direccion"`
	CodigoPuntoVenta             int              `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	FechaEmision                 string           `xml:"fechaEmision" json:"fechaEmision"`
	NombreRazonSocial            Nilable[string]  `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	CodigoTipoDocumentoIdentidad int              `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NumeroDocumento              string           `xml:"numeroDocumento" json:"numeroDocumento"`
	Complemento                  Nilable[string]  `xml:"complemento" json:"complemento"`
	CodigoCliente                string           `xml:"codigoCliente" json:"codigoCliente"`
	CodigoMetodoPago             int              `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	NumeroTarjeta                Nilable[int64]   `xml:"numeroTarjeta" json:"numeroTarjeta"`
	MontoTotal                   float64          `xml:"montoTotal" json:"montoTotal"`
	MontoTotalSujetoIva          float64          `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMoneda                 int              `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64          `xml:"tipoCambio" json:"tipoCambio"`
	MontoTotalMoneda             float64          `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	MontoGiftCard                Nilable[float64] `xml:"montoGiftCard" json:"montoGiftCard"`
	DescuentoAdicional           Nilable[float64] `xml:"descuentoAdicional" json:"descuentoAdicional"`
	CodigoExcepcion              Nilable[int64]   `xml:"codigoExcepcion" json:"codigoExcepcion"`
	Cafc                         Nilable[string]  `xml:"cafc" json:"cafc"`
	Leyenda                      string           `xml:"leyenda" json:"leyenda"`
	Usuario                      string           `xml:"usuario" json:"usuario"`
	CodigoDocumentoSector        int              `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}

// Detalle representa un ítem o servicio dentro de la factura.
type Detalle struct {
	ActividadEconomica string           `xml:"actividadEconomica" json:"actividadEconomica"`
	CodigoProductoSin  string           `xml:"codigoProductoSin" json:"codigoProductoSin"`
	CodigoProducto     string           `xml:"codigoProducto" json:"codigoProducto"`
	Descripcion        string           `xml:"descripcion" json:"descripcion"`
	Cantidad           float64          `xml:"cantidad" json:"cantidad"`
	UnidadMedida       int              `xml:"unidadMedida" json:"unidadMedida"`
	PrecioUnitario     float64          `xml:"precioUnitario" json:"precioUnitario"`
	MontoDescuento     Nilable[float64] `xml:"montoDescuento" json:"montoDescuento"`
	SubTotal           float64          `xml:"subTotal" json:"subTotal"`
	NumeroSerie        Nilable[string]  `xml:"numeroSerie" json:"numeroSerie"`
	NumeroImei         Nilable[string]  `xml:"numeroImei" json:"numeroImei"`
}
