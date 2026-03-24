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

// HospitalClinica representa la estructura completa de una factura de Hospital o Clínica lista para ser procesada.
type HospitalClinica struct {
	models.RequestWrapper[documentos.FacturaHospitalClinica]
}

// HospitalClinicaCabecera representa la sección de cabecera de la factura.
type HospitalClinicaCabecera struct {
	models.RequestWrapper[documentos.CabeceraHospitalClinica]
}

// HospitalClinicaDetalle representa un ítem individual dentro del detalle de la factura.
type HospitalClinicaDetalle struct {
	models.RequestWrapper[documentos.DetalleHospitalClinica]
}

// NewHospitalClinicaBuilder inicia el proceso de construcción de la factura.
func NewHospitalClinicaBuilder() *hospitalClinicaBuilder {
	return &hospitalClinicaBuilder{
		factura: &documentos.FacturaHospitalClinica{
			XMLName:           xml.Name{Local: "facturaElectronicaHospitalClinica"},
			XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
			XsiSchemaLocation: "facturaElectronicaHospitalClinica.xsd",
		},
	}
}

// NewHospitalClinicaCabeceraBuilder crea una instancia del constructor para la cabecera.
func NewHospitalClinicaCabeceraBuilder() *hospitalClinicaCabeceraBuilder {
	return &hospitalClinicaCabeceraBuilder{
		cabecera: &documentos.CabeceraHospitalClinica{
			CodigoDocumentoSector: 17, // Sector 17
		},
	}
}

// NewHospitalClinicaDetalleBuilder crea una instancia del constructor para los ítems de detalle.
func NewHospitalClinicaDetalleBuilder() *hospitalClinicaDetalleBuilder {
	return &hospitalClinicaDetalleBuilder{
		detalle: &documentos.DetalleHospitalClinica{},
	}
}

type hospitalClinicaBuilder struct {
	factura *documentos.FacturaHospitalClinica
}

func (b *hospitalClinicaBuilder) WithCabecera(req HospitalClinicaCabecera) *hospitalClinicaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.CabeceraHospitalClinica](req); internal != nil {
		b.factura.Cabecera = *internal
	}
	return b
}

func (b *hospitalClinicaBuilder) AddDetalle(req HospitalClinicaDetalle) *hospitalClinicaBuilder {
	if internal := models.UnwrapInternalRequest[documentos.DetalleHospitalClinica](req); internal != nil {
		b.factura.Detalle = append(b.factura.Detalle, *internal)
	}
	return b
}

func (b *hospitalClinicaBuilder) WithModalidad(tipo int) *hospitalClinicaBuilder {
	switch tipo {
	case siat.ModalidadElectronica:
		b.factura.XMLName = xml.Name{Local: "facturaElectronicaHospitalClinica"}
		b.factura.XsiSchemaLocation = "facturaElectronicaHospitalClinica.xsd"
	case siat.ModalidadComputarizada:
		b.factura.XMLName = xml.Name{Local: "facturaComputarizadaHospitalClinica"}
		b.factura.XsiSchemaLocation = "facturaComputarizadaHospitalClinica.xsd"
	}
	return b
}

func (b *hospitalClinicaBuilder) Build() HospitalClinica {
	return HospitalClinica{models.NewRequestWrapper(b.factura)}
}

type hospitalClinicaCabeceraBuilder struct {
	cabecera *documentos.CabeceraHospitalClinica
}

func (b *hospitalClinicaCabeceraBuilder) WithNitEmisor(v int64) *hospitalClinicaCabeceraBuilder {
	b.cabecera.NitEmisor = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithRazonSocialEmisor(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.RazonSocialEmisor = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithMunicipio(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.Municipio = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithTelefono(telefono *string) *hospitalClinicaCabeceraBuilder {
	if telefono == nil {
		b.cabecera.Telefono = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *telefono
	b.cabecera.Telefono = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithNumeroFactura(v int64) *hospitalClinicaCabeceraBuilder {
	b.cabecera.NumeroFactura = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCuf(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.Cuf = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCufd(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.Cufd = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoSucursal(v int) *hospitalClinicaCabeceraBuilder {
	b.cabecera.CodigoSucursal = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithDireccion(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.Direccion = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoPuntoVenta(v *int) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoPuntoVenta = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithFechaEmision(v time.Time) *hospitalClinicaCabeceraBuilder {
	b.cabecera.FechaEmision = datatype.NewTimeSiat(v)
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithNombreRazonSocial(v *string) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NombreRazonSocial = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoTipoDocumentoIdentidad(v int) *hospitalClinicaCabeceraBuilder {
	b.cabecera.CodigoTipoDocumentoIdentidad = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithNumeroDocumento(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.NumeroDocumento = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithComplemento(v *string) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.Complemento = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Complemento = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoCliente(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.CodigoCliente = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithModalidadServicio(v *string) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.ModalidadServicio = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.ModalidadServicio = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoMetodoPago(v int) *hospitalClinicaCabeceraBuilder {
	b.cabecera.CodigoMetodoPago = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithNumeroTarjeta(v *int64) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.NumeroTarjeta = datatype.Nilable[int64]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithMontoTotal(v float64) *hospitalClinicaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotal = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithMontoTotalSujetoIva(v float64) *hospitalClinicaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalSujetoIva = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoMoneda(v int) *hospitalClinicaCabeceraBuilder {
	b.cabecera.CodigoMoneda = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithTipoCambio(v float64) *hospitalClinicaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.TipoCambio = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithMontoTotalMoneda(v float64) *hospitalClinicaCabeceraBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.cabecera.MontoTotalMoneda = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithMontoGiftCard(v *float64) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.MontoGiftCard = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithDescuentoAdicional(v *float64) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.cabecera.DescuentoAdicional = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoExcepcion(v *int) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.CodigoExcepcion = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCafc(v *string) *hospitalClinicaCabeceraBuilder {
	if v == nil {
		b.cabecera.Cafc = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.cabecera.Cafc = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithLeyenda(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.Leyenda = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithUsuario(v string) *hospitalClinicaCabeceraBuilder {
	b.cabecera.Usuario = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) WithCodigoDocumentoSector(v int) *hospitalClinicaCabeceraBuilder {
	b.cabecera.CodigoDocumentoSector = v
	return b
}

func (b *hospitalClinicaCabeceraBuilder) Build() HospitalClinicaCabecera {
	return HospitalClinicaCabecera{models.NewRequestWrapper(b.cabecera)}
}

type hospitalClinicaDetalleBuilder struct {
	detalle *documentos.DetalleHospitalClinica
}

func (b *hospitalClinicaDetalleBuilder) WithActividadEconomica(v string) *hospitalClinicaDetalleBuilder {
	b.detalle.ActividadEconomica = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithCodigoProductoSin(v int64) *hospitalClinicaDetalleBuilder {
	b.detalle.CodigoProductoSin = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithCodigoProducto(v string) *hospitalClinicaDetalleBuilder {
	b.detalle.CodigoProducto = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithDescripcion(v string) *hospitalClinicaDetalleBuilder {
	b.detalle.Descripcion = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithEspecialidad(v *string) *hospitalClinicaDetalleBuilder {
	if v == nil {
		b.detalle.Especialidad = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.Especialidad = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithEspecialidadDetalle(v *string) *hospitalClinicaDetalleBuilder {
	if v == nil {
		b.detalle.EspecialidadDetalle = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.EspecialidadDetalle = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithNroQuirofanoSalaOperaciones(v int) *hospitalClinicaDetalleBuilder {
	b.detalle.NroQuirofanoSalaOperaciones = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithEspecialidadMedico(v *string) *hospitalClinicaDetalleBuilder {
	if v == nil {
		b.detalle.EspecialidadMedico = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.EspecialidadMedico = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithNombreApellidoMedico(v string) *hospitalClinicaDetalleBuilder {
	b.detalle.NombreApellidoMedico = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithNitDocumentoMedico(v int64) *hospitalClinicaDetalleBuilder {
	b.detalle.NitDocumentoMedico = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithNroMatriculaMedico(v *string) *hospitalClinicaDetalleBuilder {
	if v == nil {
		b.detalle.NroMatriculaMedico = datatype.Nilable[string]{Value: nil}
		return b
	}
	value := *v
	b.detalle.NroMatriculaMedico = datatype.Nilable[string]{Value: &value}
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithNroFacturaMedico(v *int) *hospitalClinicaDetalleBuilder {
	if v == nil {
		b.detalle.NroFacturaMedico = datatype.Nilable[int]{Value: nil}
		return b
	}
	value := *v
	b.detalle.NroFacturaMedico = datatype.Nilable[int]{Value: &value}
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithCantidad(v float64) *hospitalClinicaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.Cantidad = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithUnidadMedida(v int) *hospitalClinicaDetalleBuilder {
	b.detalle.UnidadMedida = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithPrecioUnitario(v float64) *hospitalClinicaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.PrecioUnitario = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithMontoDescuento(v *float64) *hospitalClinicaDetalleBuilder {
	if v == nil {
		b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: nil}
		return b
	}
	value := *v
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', 2, 64), 64)
	b.detalle.MontoDescuento = datatype.Nilable[float64]{Value: &value}
	return b
}

func (b *hospitalClinicaDetalleBuilder) WithSubTotal(v float64) *hospitalClinicaDetalleBuilder {
	v, _ = strconv.ParseFloat(strconv.FormatFloat(v, 'f', 2, 64), 64)
	b.detalle.SubTotal = v
	return b
}

func (b *hospitalClinicaDetalleBuilder) Build() HospitalClinicaDetalle {
	return HospitalClinicaDetalle{models.NewRequestWrapper(b.detalle)}
}
