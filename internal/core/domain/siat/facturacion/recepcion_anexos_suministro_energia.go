package facturacion

import (
	"encoding/xml"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
)

type RecepcionAnexosSuministroEnergia struct {
	XMLName                            xml.Name                                  `xml:"ns:recepcionAnexosSuministroEnergia" json:"-"`
	SolicitudRecepcionSuministroAnexos SolicitudRecepcionAnexosSuministroEnergia `xml:"SolicitudRecepcionSuministroAnexos" json:"solicitudRecepcionSuministroAnexos"`
}

type SolicitudRecepcionAnexosSuministroEnergia struct {
	SolicitudRecepcion
	AnexosList []SuministroEnergiaAnexo `xml:"anexosList" json:"anexosList"`
	GiftCard   int64                    `xml:"giftCard" json:"giftCard"`
}

type RecepcionAnexosSuministroEnergiaResponse struct {
	RespuestaRecepcionAnexos RespuestaRecepcion `xml:"RespuestaRecepcionAnexos" json:"respuestaRecepcionAnexos"`
}

type SuministroEnergiaAnexo struct {
	CufFactSuministro string            `xml:"cufFactSuministro" json:"cufFactSuministro"`
	FechaRecarga      datatype.TimeSiat `xml:"fechaRecarga" json:"fechaRecarga"`
	MontoRecarga      float64           `xml:"montoRecarga" json:"montoRecarga"`
}
