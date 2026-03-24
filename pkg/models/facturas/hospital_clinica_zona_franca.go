package facturas

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/ron86i/go-siat"
	"github.com/ron86i/go-siat/pkg/models"

	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/documentos"
)

// HospitalClinicaZF representa la estructura completa de una factura Hospital Clínica Zona Franca.
type HospitalClinicaZF struct {
	models.RequestWrapper[documentos.FacturaHospitalClinicaZF]
}

// HospitalClinicaZFCabecera representa la sección de cabecera.
type HospitalClinicaZFCabecera struct {
	models.RequestWrapper[documentos.CabeceraHospitalClinicaZF]
}

// HospitalClinicaZFDetalle representa un ítem individual dentro del detalle.
type HospitalClinicaZFDetalle struct {
	models.RequestWrapper[documentos.DetalleHospitalClinicaZF]
}

// NewHospitalClinicaZFBuilder inicia el proceso de construcción.
func NewHospitalClinicaZFBuilder() *hospitalClinicaZFBuilder {
	return &hospitalClinicaZFBuilder{
		factura: &documentos.FacturaHospitalClinicaZF{
			XMLName:           xml.Name{Local: "facturaElectronicaHospitalClinicaZF"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaHospitalClinicaZF.xsd",
		},
	}
}

// NewHospitalClinicaZFCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewHospitalClinicaZFCabeceraBuilder() *hospitalClinicaZFCabeceraBuilder {
	return &hospitalClinicaZFCabeceraBuilder{
		cabecera: &documentos.CabeceraHospitalClinicaZF{
			CodigoDocumentoSector: 50, // Sector 50
			MontoTotalSujetoIva:   0,  // Siempre 0
		},
	}
}

// NewHospitalClinicaZonaFrancaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewHospitalClinicaZonaFrancaDetalleBuilder() *hospitalClinicaZFDetalleBuilder {
	return &hospitalClinicaZFDetalleBuilder{
		detalle: &documentos.DetalleHospitalClinicaZF{},
	}
}

type hospitalClinicaZFBuilder struct {
	factura *documentos.FacturaHospitalClinicaZF
}

func (b *hospitalClinicaZFBuilder) WithCabecera(req HospitalClinicaZFCabecera) *hospitalClinicaZFBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraHospitalClinicaZF](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *hospitalClinicaZFBuilder) AddDetalle(req HospitalClinicaZFDetalle) *hospitalClinicaZFBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleHospitalClinicaZF](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *hospitalClinicaZFBuilder) WithModalidad(tipo int) *hospitalClinicaZFBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaHospitalClinicaZF"}
		b.factura.XsiSchemaLocation = "facturaElectronicaHospitalClinicaZF.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaHospitalClinicaZF"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaHospitalClinicaZF.xsd"
	}
	return b
}

func (b *hospitalClinicaZFBuilder) Build() HospitalClinicaZF {
	return HospitalClinicaZF{models.NewRequestWrapper(b.factura)}
}

type hospitalClinicaZFCabeceraBuilder struct {
	cabecera *documentos.CabeceraHospitalClinicaZF
}

func (b *hospitalClinicaZFCabeceraBuilder) WithNitEmisor(v int64) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithRazonSocialEmisor(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithMunicipio(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithTelefono(telefono *string) *hospitalClinicaZFCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithNumeroFactura(v int64) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCuf(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCufd(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoSucursal(v int) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithDireccion(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoPuntoVenta(v *int) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithFechaEmision(fechaEmision time.Time) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(fechaEmision)
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithNombreRazonSocial(v *string) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithNumeroDocumento(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithComplemento(v *string) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoCliente(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithModalidadServicio(v *string) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.ModalidadServicio = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.ModalidadServicio = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoMetodoPago(v int) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithNumeroTarjeta(v *int64) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithMontoTotal(v float64) *hospitalClinicaZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoMoneda(v int) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithTipoCambio(v float64) *hospitalClinicaZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithMontoTotalMoneda(v float64) *hospitalClinicaZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithMontoGiftCard(v *float64) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithDescuentoAdicional(v *float64) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoExcepcion(v *int) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCafc(v *string) *hospitalClinicaZFCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithLeyenda(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithUsuario(v string) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *hospitalClinicaZFCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) WithCodigoDocumentoSector(v int) *hospitalClinicaZFCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *hospitalClinicaZFCabeceraBuilder) Build() HospitalClinicaZFCabecera {
	return HospitalClinicaZFCabecera{models.NewRequestWrapper(b.cabecera)}
}

type hospitalClinicaZFDetalleBuilder struct {
	detalle *documentos.DetalleHospitalClinicaZF
}

func (b *hospitalClinicaZFDetalleBuilder) WithActividadEconomica(v string) *hospitalClinicaZFDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithCodigoProductoSin(v int64) *hospitalClinicaZFDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithCodigoProducto(v string) *hospitalClinicaZFDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithDescripcion(v string) *hospitalClinicaZFDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithEspecialidad(v *string) *hospitalClinicaZFDetalleBuilder {
	if v == nil {
		b.detalle.Especialidad = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.Especialidad = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithEspecialidadDetalle(v *string) *hospitalClinicaZFDetalleBuilder {
	if v == nil {
		b.detalle.EspecialidadDetalle = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.EspecialidadDetalle = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithNroQuirofanoSalaOperaciones(v int) *hospitalClinicaZFDetalleBuilder {
	b.detalle.NroQuirofanoSalaOperaciones = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithEspecialidadMedico(v *string) *hospitalClinicaZFDetalleBuilder {
	if v == nil {
		b.detalle.EspecialidadMedico = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.EspecialidadMedico = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithNombreApellidoMedico(v string) *hospitalClinicaZFDetalleBuilder {
	b.detalle.NombreApellidoMedico = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithNitDocumentoMedico(v int64) *hospitalClinicaZFDetalleBuilder {
	b.detalle.NitDocumentoMedico = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithNroMatriculaMedico(v *string) *hospitalClinicaZFDetalleBuilder {
	if v == nil {
		b.detalle.NroMatriculaMedico = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.NroMatriculaMedico = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithNroFacturaMedico(v *int) *hospitalClinicaZFDetalleBuilder {
	if v == nil {
		b.detalle.NroFacturaMedico = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.detalle.NroFacturaMedico = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithCantidad(v float64) *hospitalClinicaZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithUnidadMedida(v int) *hospitalClinicaZFDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithPrecioUnitario(v float64) *hospitalClinicaZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithMontoDescuento(v *float64) *hospitalClinicaZFDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) WithSubTotal(v float64) *hospitalClinicaZFDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *hospitalClinicaZFDetalleBuilder) Build() HospitalClinicaZFDetalle {
	return HospitalClinicaZFDetalle{models.NewRequestWrapper(b.detalle)}
}
