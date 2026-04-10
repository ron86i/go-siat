package documents

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

// FacturaBoletoAereo representa la estructura completa de una factura de Boleto Aéreo para el SIAT.
// Este sector (30) es especial porque toda su información se encuentra en la cabecera y no utiliza una lista de detalles.
type FacturaBoletoAereo struct {
	XMLName           xml.Name            `json:"-"`
	XmlnsXsi          string              `xml:"xmlns:xsi,attr" json:"-"`
	XsiSchemaLocation string              `xml:"xsi:noNamespaceSchemaLocation,attr" json:"-"`
	Cabecera          CabeceraBoletoAereo `xml:"cabecera" json:"cabecera"`
}

// CabeceraBoletoAereo contiene la información técnica y comercial para la emisión de boletos aéreos.
// El orden de los campos es EXTREMADAMENTE CRÍTICO y sigue el XSD ver.23/08/2021.
type CabeceraBoletoAereo struct {
	NitEmisor                    int64                     `xml:"nitEmisor" json:"nitEmisor"`
	NumeroFactura                int64                     `xml:"numeroFactura" json:"numeroFactura"`
	Cufd                         string                    `xml:"cufd" json:"cufd"`
	Cuf                          string                    `xml:"cuf" json:"cuf"`
	CodigoSucursal               int                       `xml:"codigoSucursal" json:"codigoSucursal"`
	CodigoPuntoVenta             datatype.Nilable[int]     `xml:"codigoPuntoVenta" json:"codigoPuntoVenta"`
	Direccion                    string                    `xml:"direccion" json:"direccion"`
	FechaEmision                 datatype.TimeSiat         `xml:"fechaEmision" json:"fechaEmision"`
	RazonSocialEmisor            string                    `xml:"razonSocialEmisor" json:"razonSocialEmisor"`
	NombreRazonSocial            datatype.Nilable[string]  `xml:"nombreRazonSocial" json:"nombreRazonSocial"`
	NumeroDocumento              string                    `xml:"numeroDocumento" json:"numeroDocumento"`
	CodigoTipoDocumentoIdentidad int                       `xml:"codigoTipoDocumentoIdentidad" json:"codigoTipoDocumentoIdentidad"`
	NombrePasajero               string                    `xml:"nombrePasajero" json:"nombrePasajero"`
	NumeroDocumentoPasajero      datatype.Nilable[string]  `xml:"numeroDocumentoPasajero" json:"numeroDocumentoPasajero"`
	CodigoIataLineaAerea         int                       `xml:"codigoIataLineaAerea" json:"codigoIataLineaAerea"`
	CodigoIataAgenteViajes       datatype.Nilable[string]  `xml:"codigoIataAgenteViajes" json:"codigoIataAgenteViajes"`
	NitAgenteViajes              datatype.Nilable[int64]   `xml:"nitAgenteViajes" json:"nitAgenteViajes"`
	CodigoOrigenServicio         string                    `xml:"codigoOrigenServicio" json:"codigoOrigenServicio"`
	CodigoMoneda                 int                       `xml:"codigoMoneda" json:"codigoMoneda"`
	TipoCambio                   float64                   `xml:"tipoCambio" json:"tipoCambio"`
	MontoTarifa                  float64                   `xml:"montoTarifa" json:"montoTarifa"`
	MontoTotal                   float64                   `xml:"montoTotal" json:"montoTotal"`
	MontoTotalMoneda             float64                   `xml:"montoTotalMoneda" json:"montoTotalMoneda"`
	MontoTotalSujetoIva          float64                   `xml:"montoTotalSujetoIva" json:"montoTotalSujetoIva"`
	CodigoMetodoPago             int                       `xml:"codigoMetodoPago" json:"codigoMetodoPago"`
	CodigoTipoTransaccion        string                    `xml:"codigoTipoTransaccion" json:"codigoTipoTransaccion"`
	Usuario                      string                    `xml:"usuario" json:"usuario"`
	Leyenda                      string                    `xml:"leyenda" json:"leyenda"`
	CodigoDocumentoSector        int                       `xml:"codigoDocumentoSector" json:"codigoDocumentoSector"`
}
