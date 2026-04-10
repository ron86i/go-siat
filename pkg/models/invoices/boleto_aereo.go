package invoices

import (
	"encoding/xml"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documents"
	"github.com/ron86i/go-siat/pkg/models"
	"github.com/ron86i/go-siat/pkg/utils"
)

// BoletoAereo representa la estructura completa de una factura de boleto aéreo lista para ser procesada.
// Esta estructura encapsula la factura del sector 30 para su uso en los servicios del SDK.
type BoletoAereo struct {
	models.RequestWrapper[documents.FacturaBoletoAereo]
}

// BoletoAereoCabecera representa la sección de cabecera de una factura de boleto aéreo.
type BoletoAereoCabecera struct {
	models.RequestWrapper[documents.CabeceraBoletoAereo]
}

// boletoAereoBuilder permite la construcción fluida de una factura de Boleto Aéreo.
type boletoAereoBuilder struct {
	factura *documents.FacturaBoletoAereo
}

// NewBoletoAereoBuilder inicializa un nuevo constructor para una factura de Boleto Aéreo.
func NewBoletoAereoBuilder() *boletoAereoBuilder {
	return &boletoAereoBuilder{
		factura: &documents.FacturaBoletoAereo{
			XMLName:           xml.Name{Local: "facturaElectronicaBoletoAereo"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaBoletoAereo.xsd",
		},
	}
}

// WithModalidad establece el nombre de las etiquetas XML y el esquema según la modalidad (Electrónica o Computarizada).
func (b *boletoAereoBuilder) WithModalidad(tipo int) *boletoAereoBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaBoletoAereo"}
		b.factura.XsiSchemaLocation = "facturaElectronicaBoletoAereo.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaBoletoAereo"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaBoletoAereo.xsd"
	}
	return b
}

// WithCabecera asocia la cabecera construida a la factura.
func (b *boletoAereoBuilder) WithCabecera(req BoletoAereoCabecera) *boletoAereoBuilder {
	if internal := models.UnwrapInternalRequest[documents.CabeceraBoletoAereo](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

// Build finaliza la construcción y retorna el objeto BoletoAereo para su envío.
func (b *boletoAereoBuilder) Build() BoletoAereo {
	return BoletoAereo{models.NewRequestWrapper(b.factura)}
}

// --- Cabecera Builder ---

// boletoAereoCabeceraBuilder permite la construcción fluida de la cabecera de la factura.
type boletoAereoCabeceraBuilder struct {
	cabecera *documents.CabeceraBoletoAereo
}

// NewBoletoAereoCabeceraBuilder inicializa un nuevo constructor para la cabecera del Boleto Aéreo.
func NewBoletoAereoCabeceraBuilder() *boletoAereoCabeceraBuilder {
	return &boletoAereoCabeceraBuilder{
		cabecera: &documents.CabeceraBoletoAereo{
			CodigoDocumentoSector: 30,
		},
	}
}

func (b *boletoAereoCabeceraBuilder) WithNitEmisor(v int64) *boletoAereoCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithRazonSocialEmisor(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithNumeroFactura(v int64) *boletoAereoCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCuf(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCufd(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoSucursal(v int) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithDireccion(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoPuntoVenta(v *int) *boletoAereoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &val}
	}
	return b
}

func (b *boletoAereoCabeceraBuilder) WithFechaEmision(v time.Time) *boletoAereoCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *boletoAereoCabeceraBuilder) WithNombreRazonSocial(v *string) *boletoAereoCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithNumeroDocumento(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithNombrePasajero(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.NombrePasajero = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithNumeroDocumentoPasajero(v *string) *boletoAereoCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroDocumentoPasajero = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.NumeroDocumentoPasajero = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoIataLineaAerea(v int) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoIataLineaAerea = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoIataAgenteViajes(v *string) *boletoAereoCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoIataAgenteViajes = datatype.Nilable[string]{Value: nil}
	} else {
		val := *v
		b.cabecera.CodigoIataAgenteViajes = datatype.Nilable[string]{Value: &val}
	}
	return b
}

func (b *boletoAereoCabeceraBuilder) WithNitAgenteViajes(v *int64) *boletoAereoCabeceraBuilder {
	if v == nil {
		b.cabecera.NitAgenteViajes = datatype.Nilable[int64]{Value: nil}
	} else {
		val := *v
		b.cabecera.NitAgenteViajes = datatype.Nilable[int64]{Value: &val}
	}
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoOrigenServicio(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoOrigenServicio = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoMetodoPago(v int) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithMontoTarifa(v float64) *boletoAereoCabeceraBuilder {
	b.cabecera.MontoTarifa = utils.Round(v, 2)
	return b
}

func (b *boletoAereoCabeceraBuilder) WithMontoTotal(v float64) *boletoAereoCabeceraBuilder {
	b.cabecera.MontoTotal = utils.Round(v, 2)
	return b
}

func (b *boletoAereoCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *boletoAereoCabeceraBuilder {
	b.cabecera.MontoTotalSujetoIva = utils.Round(v, 2)
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoMoneda(v int) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithTipoCambio(v float64) *boletoAereoCabeceraBuilder {
	b.cabecera.TipoCambio = utils.Round(v, 2)
	return b
}

func (b *boletoAereoCabeceraBuilder) WithMontoTotalMoneda(v float64) *boletoAereoCabeceraBuilder {
	b.cabecera.MontoTotalMoneda = utils.Round(v, 2)
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoTipoTransaccion(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoTipoTransaccion = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithLeyenda(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithUsuario(v string) *boletoAereoCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *boletoAereoCabeceraBuilder) WithCodigoDocumentoSector(v int) *boletoAereoCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *boletoAereoCabeceraBuilder) Build() BoletoAereoCabecera {
	return BoletoAereoCabecera{models.NewRequestWrapper(b.cabecera)}
}
