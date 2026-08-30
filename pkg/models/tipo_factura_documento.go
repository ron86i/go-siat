package models

// Tipos de documento fiscal definidos por el SIAT.
// Se declaran también en el paquete raíz para mantener su API pública; esta
// copia evita que models tenga que importar al paquete raíz y crear un ciclo.
const (
	TipoFacturaConDerechoCreditoFiscal = iota + 1
	TipoFacturaSinDerechoCreditoFiscal
	TipoNotaCreditoDebito
	TipoBoletoAereo
)

// DefinicionTipoFactura describe el tipo fiscal predeterminado de un sector
// implementado por el SDK.
type DefinicionTipoFactura struct {
	CodigoDocumentoSector int
	TipoFacturaDocumento  int
}

var definicionesTipoFactura = []DefinicionTipoFactura{
	{1, TipoFacturaConDerechoCreditoFiscal},
	{2, TipoFacturaConDerechoCreditoFiscal},
	{3, TipoFacturaSinDerechoCreditoFiscal},
	{4, TipoFacturaSinDerechoCreditoFiscal},
	{5, TipoFacturaSinDerechoCreditoFiscal},
	{6, TipoFacturaSinDerechoCreditoFiscal},
	{7, TipoFacturaSinDerechoCreditoFiscal},
	{8, TipoFacturaSinDerechoCreditoFiscal},
	{9, TipoFacturaSinDerechoCreditoFiscal},
	{10, TipoFacturaSinDerechoCreditoFiscal},
	{11, TipoFacturaConDerechoCreditoFiscal},
	{12, TipoFacturaConDerechoCreditoFiscal},
	{13, TipoFacturaConDerechoCreditoFiscal},
	{14, TipoFacturaConDerechoCreditoFiscal},
	{15, TipoFacturaConDerechoCreditoFiscal},
	{16, TipoFacturaConDerechoCreditoFiscal},
	{17, TipoFacturaConDerechoCreditoFiscal},
	{18, TipoFacturaConDerechoCreditoFiscal},
	{19, TipoFacturaConDerechoCreditoFiscal},
	{20, TipoFacturaSinDerechoCreditoFiscal},
	{21, TipoFacturaConDerechoCreditoFiscal},
	{22, TipoFacturaConDerechoCreditoFiscal},
	{23, TipoFacturaConDerechoCreditoFiscal},
	{24, TipoNotaCreditoDebito},
	{28, TipoFacturaSinDerechoCreditoFiscal},
	{29, TipoNotaCreditoDebito},
	{30, TipoBoletoAereo},
	{31, TipoFacturaConDerechoCreditoFiscal},
	{34, TipoFacturaConDerechoCreditoFiscal},
	{35, TipoFacturaConDerechoCreditoFiscal},
	{36, TipoFacturaSinDerechoCreditoFiscal},
	{37, TipoFacturaConDerechoCreditoFiscal},
	{38, TipoFacturaConDerechoCreditoFiscal},
	{39, TipoFacturaConDerechoCreditoFiscal},
	{40, TipoFacturaSinDerechoCreditoFiscal},
	{41, TipoFacturaConDerechoCreditoFiscal},
	{42, TipoFacturaSinDerechoCreditoFiscal},
	{43, TipoFacturaSinDerechoCreditoFiscal},
	{44, TipoFacturaConDerechoCreditoFiscal},
	{45, TipoFacturaSinDerechoCreditoFiscal},
	{46, TipoFacturaSinDerechoCreditoFiscal},
	{47, TipoNotaCreditoDebito},
	{48, TipoNotaCreditoDebito},
	{49, TipoFacturaSinDerechoCreditoFiscal},
	{50, TipoFacturaSinDerechoCreditoFiscal},
	{51, TipoFacturaConDerechoCreditoFiscal},
	{52, TipoFacturaSinDerechoCreditoFiscal},
	{53, TipoFacturaConDerechoCreditoFiscal},
	{54, TipoFacturaSinDerechoCreditoFiscal},
	{55, TipoFacturaConDerechoCreditoFiscal},
}

var tipoFacturaPorSector = func() map[int]int {
	tipos := make(map[int]int, len(definicionesTipoFactura))
	for _, definicion := range definicionesTipoFactura {
		tipos[definicion.CodigoDocumentoSector] = definicion.TipoFacturaDocumento
	}
	return tipos
}()

// DefinicionesTipoFactura devuelve una copia del catálogo para que los
// consumidores no puedan alterar las definiciones internas.
func DefinicionesTipoFactura() []DefinicionTipoFactura {
	return append([]DefinicionTipoFactura(nil), definicionesTipoFactura...)
}

// MetadatosFactura contiene la información fiscal que acompaña a un documento
// durante su preparación y recepción. No se serializa dentro del XML fiscal.
//
// El tipo se inicializa desde el catálogo SIAT y puede sobrescribirse en el
// builder de la factura cuando una integración lo requiera.
type MetadatosFactura struct {
	codigoDocumentoSector int
	tipoFacturaDocumento  int
}

// NuevosMetadatosFactura crea metadatos con el tipo normativo del sector.
func NuevosMetadatosFactura(codigoDocumentoSector int) MetadatosFactura {
	return MetadatosFactura{
		codigoDocumentoSector: codigoDocumentoSector,
		tipoFacturaDocumento:  TipoFacturaDocumentoPredeterminado(codigoDocumentoSector),
	}
}

// CodigoDocumentoSector identifica el diseño fiscal del documento.
func (m MetadatosFactura) CodigoDocumentoSector() int { return m.codigoDocumentoSector }

// TipoFacturaDocumento devuelve el tipo fiscal que debe viajar en el CUF y
// en la solicitud SOAP.
func (m MetadatosFactura) TipoFacturaDocumento() int { return m.tipoFacturaDocumento }

// WithTipoFacturaDocumento sobrescribe el tipo fiscal predeterminado.
func (m *MetadatosFactura) WithTipoFacturaDocumento(tipo int) {
	m.tipoFacturaDocumento = tipo
}

// TipoFacturaDocumentoPredeterminado devuelve el tipo fiscal normativo para
// un documento sector. Los sectores desconocidos conservan el tipo 1, que es
// el valor general de las facturas ordinarias.
//
// El valor puede sobrescribirse con WithTipoFacturaDocumento cuando una
// integración requiera una excepción explícita.
func TipoFacturaDocumentoPredeterminado(codigoDocumentoSector int) int {
	if tipo, existe := tipoFacturaPorSector[codigoDocumentoSector]; existe {
		return tipo
	}
	return TipoFacturaConDerechoCreditoFiscal
}
