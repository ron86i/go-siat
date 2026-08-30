package invoices_test

import (
	"encoding/xml"
	"strings"
	"testing"

	siat "github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
	"github.com/ron86i/go-siat/v2/pkg/models/invoices"
)

// sectorCase describe lo que el catálogo normativo del SIAT espera de un builder de sector:
// el código de documento sector y el elemento raíz del XML en cada modalidad.
//
// Los valores esperados provienen del catálogo del SIAT, no del código, así que este test
// detecta el error más común al agregar un sector: clonar un archivo existente y olvidar
// cambiar el código o el nombre del elemento raíz. Ese error no rompe la compilación y el
// SIAT lo rechaza recién en producción con el código 931, 932 o 939.
type sectorCase struct {
	codigo        int
	nombre        string
	electronica   string
	computarizada string
	build         func(modalidad int) any
}

func sectorCases() []sectorCase {
	return []sectorCase{
		{1, "Compra y Venta", "facturaElectronicaCompraVenta", "facturaComputarizadaCompraVenta",
			func(m int) any { return invoices.NewCompraVentaBuilder().WithModalidad(m).Build() }},
		{2, "Alquiler de Bienes Inmuebles", "facturaElectronicaAlquilerBienInmueble", "facturaComputarizadaAlquilerBienInmueble",
			func(m int) any { return invoices.NewAlquilerBienInmuebleBuilder().WithModalidad(m).Build() }},
		{3, "Comercial de Exportación", "facturaElectronicaComercialExportacion", "facturaComputarizadaComercialExportacion",
			func(m int) any { return invoices.NewComercialExportacionBuilder().WithModalidad(m).Build() }},
		{4, "Libre Consignación", "facturaElectronicaLibreConsignacion", "facturaComputarizadaLibreConsignacion",
			func(m int) any { return invoices.NewLibreConsignacionBuilder().WithModalidad(m).Build() }},
		{5, "Venta en Zona Franca", "facturaElectronicaZonaFranca", "facturaComputarizadaZonaFranca",
			func(m int) any { return invoices.NewZonaFrancaBuilder().WithModalidad(m).Build() }},
		{6, "Servicio Turístico y Hospedaje", "facturaElectronicaServicioTuristicoHospedaje", "facturaComputarizadaServicioTuristicoHospedaje",
			func(m int) any { return invoices.NewTurismoHospedajeBuilder().WithModalidad(m).Build() }},
		{7, "Seguridad Alimentaria", "facturaElectronicaSeguridadAlimentaria", "facturaComputarizadaSeguridadAlimentaria",
			func(m int) any { return invoices.NewSeguridadAlimentariaBuilder().WithModalidad(m).Build() }},
		{8, "Tasa Cero", "facturaElectronicaTasaCero", "facturaComputarizadaTasaCero",
			func(m int) any { return invoices.NewTasaCeroBuilder().WithModalidad(m).Build() }},
		{9, "Moneda Extranjera", "facturaElectronicaMonedaExtranjera", "facturaComputarizadaMonedaExtranjera",
			func(m int) any { return invoices.NewMonedaExtranjeraBuilder().WithModalidad(m).Build() }},
		{10, "Dutty Free", "facturaElectronicaDuttyFree", "facturaComputarizadaDuttyFree",
			func(m int) any { return invoices.NewDuttyFreeBuilder().WithModalidad(m).Build() }},
		{11, "Sector Educativo", "facturaElectronicaSectorEducativo", "facturaComputarizadaSectorEducativo",
			func(m int) any { return invoices.NewSectorEducativoBuilder().WithModalidad(m).Build() }},
		{12, "Comercialización de Hidrocarburos", "facturaElectronicaComercializacionHidrocarburo", "facturaComputarizadaComercializacionHidrocarburo",
			func(m int) any { return invoices.NewComercializacionHidroBuilder().WithModalidad(m).Build() }},
		{13, "Servicios Básicos", "facturaElectronicaServicioBasico", "facturaComputarizadaServicioBasico",
			func(m int) any { return invoices.NewServicioBasicoBuilder().WithModalidad(m).Build() }},
		{14, "Alcanzada por el ICE", "facturaElectronicaAlcanzadaIce", "facturaComputarizadaAlcanzadaIce",
			func(m int) any { return invoices.NewAlcanzadaIceBuilder().WithModalidad(m).Build() }},
		{15, "Entidades Financieras", "facturaElectronicaEntidadFinanciera", "facturaComputarizadaEntidadFinanciera",
			func(m int) any { return invoices.NewEntidadFinancieraBuilder().WithModalidad(m).Build() }},
		{16, "Hoteles", "facturaElectronicaHotel", "facturaComputarizadaHotel",
			func(m int) any { return invoices.NewHotelBuilder().WithModalidad(m).Build() }},
		{17, "Hospitales / Clínicas", "facturaElectronicaHospitalClinica", "facturaComputarizadaHospitalClinica",
			func(m int) any { return invoices.NewHospitalClinicaBuilder().WithModalidad(m).Build() }},
		{18, "Juegos de Azar", "facturaElectronicaJuegoAzar", "facturaComputarizadaJuegoAzar",
			func(m int) any { return invoices.NewJuegoAzarBuilder().WithModalidad(m).Build() }},
		{19, "Hidrocarburos Alcanzada IEHD", "facturaElectronicaHidrocarburoAlcanzadaIehd", "facturaComputarizadaHidrocarburoAlcanzadaIehd",
			func(m int) any { return invoices.NewHidrocarburoAlcanzadaIehdBuilder().WithModalidad(m).Build() }},
		{20, "Exportación de Minerales", "facturaElectronicaComercialExportacionMinera", "facturaComputarizadaComercialExportacionMinera",
			func(m int) any { return invoices.NewComercialExportacionMineraBuilder().WithModalidad(m).Build() }},
		{21, "Venta de Minerales", "facturaElectronicaVentaMineral", "facturaComputarizadaVentaMineral",
			func(m int) any { return invoices.NewVentaMineralBuilder().WithModalidad(m).Build() }},
		{22, "Telecomunicaciones", "facturaElectronicaTelecomunicacion", "facturaComputarizadaTelecomunicacion",
			func(m int) any { return invoices.NewTelecomunicacionesBuilder().WithModalidad(m).Build() }},
		{23, "Prevalorada", "facturaElectronicaPrevalorada", "facturaComputarizadaPrevalorada",
			func(m int) any { return invoices.NewPrevaloradaBuilder().WithModalidad(m).Build() }},
		{24, "Nota de Crédito-Débito", "notaFiscalElectronicaCreditoDebito", "notaFiscalComputarizadaCreditoDebito",
			func(m int) any { return invoices.NewNotaCreditoDebitoBuilder().WithModalidad(m).Build() }},
		{24, "Nota Fiscal de Crédito-Débito", "notaFiscalElectronicaCreditoDebito", "notaFiscalComputarizadaCreditoDebito",
			func(m int) any { return invoices.NewNotaFiscalCreditoDebitoBuilder().WithModalidad(m).Build() }},
		{28, "Exportación de Servicios", "facturaElectronicaComercialExportacionServicio", "facturaComputarizadaComercialExportacionServicio",
			func(m int) any { return invoices.NewComercialExportacionServicioBuilder().WithModalidad(m).Build() }},
		{29, "Nota de Conciliación", "notaElectronicaConciliacion", "notaComputarizadaConciliacion",
			func(m int) any { return invoices.NewNotaConciliacionBuilder().WithModalidad(m).Build() }},
		{30, "Boleto Aéreo", "facturaElectronicaBoletoAereo", "facturaComputarizadaBoletoAereo",
			func(m int) any { return invoices.NewBoletoAereoBuilder().WithModalidad(m).Build() }},
		{31, "Suministro de Energía", "facturaElectronicaSuministroEnergia", "facturaComputarizadaSuministroEnergia",
			func(m int) any { return invoices.NewSuministroEnergiaBuilder().WithModalidad(m).Build() }},
		{34, "Seguros", "facturaElectronicaSeguros", "facturaComputarizadaSeguros",
			func(m int) any { return invoices.NewSegurosBuilder().WithModalidad(m).Build() }},
		{35, "Compra Venta Bonificaciones", "facturaElectronicaCompraVentaBon", "facturaComputarizadaCompraVentaBon",
			func(m int) any { return invoices.NewCompraVentaBonificacionesBuilder().WithModalidad(m).Build() }},
		{36, "Prevalorada Sin Derecho a Crédito Fiscal", "facturaElectronicaPrevaloradaSD", "facturaComputarizadaPrevaloradaSD",
			func(m int) any {
				return invoices.NewPrevaloradaSinDerechoCreditoFiscalBuilder().WithModalidad(m).Build()
			}},
		{37, "Comercialización de GNV", "facturaElectronicaComercializacionGnv", "facturaComputarizadaComercializacionGnv",
			func(m int) any { return invoices.NewComercializacionGnvBuilder().WithModalidad(m).Build() }},
		{38, "Hidrocarburos No Alcanzada IEHD", "facturaElectronicaHidrocarburoNoAlcanzadaIehd", "facturaComputarizadaHidrocarburoNoAlcanzadaIehd",
			func(m int) any { return invoices.NewHidrocarburoNoAlcanzadaIehdBuilder().WithModalidad(m).Build() }},
		{39, "Comercialización de GN y GLP", "facturaElectronicaComercializacionGnGlp", "facturaComputarizadaComercializacionGnGlp",
			func(m int) any { return invoices.NewComercializacionGnGlpBuilder().WithModalidad(m).Build() }},
		{40, "Servicios Básicos Zona Franca", "facturaElectronicaServicioBasicoZf", "facturaComputarizadaServicioBasicoZf",
			func(m int) any { return invoices.NewServicioBasicoZFBuilder().WithModalidad(m).Build() }},
		{41, "Compra Venta Tasas", "facturaElectronicaCompraVentaTasas", "facturaComputarizadaCompraVentaTasas",
			func(m int) any { return invoices.NewCompraVentaTasasBuilder().WithModalidad(m).Build() }},
		{42, "Alquiler Zona Franca", "facturaElectronicaAlquilerZF", "facturaComputarizadaAlquilerZF",
			func(m int) any { return invoices.NewAlquilerZFBuilder().WithModalidad(m).Build() }},
		{43, "Exportación de Hidrocarburos", "facturaElectronicaComercialExportacionHidro", "facturaComputarizadaComercialExportacionHidro",
			func(m int) any { return invoices.NewComercialExportacionHidroBuilder().WithModalidad(m).Build() }},
		{44, "Importación y Comercialización de Lubricantes", "facturaElectronicaImportacionComercializacionLubricantes", "facturaComputarizadaImportacionComercializacionLubricantes",
			func(m int) any {
				return invoices.NewImportacionComercializacionLubricantesBuilder().WithModalidad(m).Build()
			}},
		{45, "Exportación Precio Venta", "facturaElectronicaComercialExportacionPVenta", "facturaComputarizadaComercialExportacionPVenta",
			func(m int) any { return invoices.NewComercialExportacionPVentaBuilder().WithModalidad(m).Build() }},
		{46, "Sector Educativo Zona Franca", "facturaElectronicaSectorEducativoZF", "facturaComputarizadaSectorEducativoZF",
			func(m int) any { return invoices.NewSectorEducativoZFBuilder().WithModalidad(m).Build() }},
		{47, "Nota Crédito Débito Descuentos", "notaElectronicaCreditoDebitoDescuento", "notaComputarizadaCreditoDebitoDescuento",
			func(m int) any { return invoices.NewNotaCreditoDebitoDescuentoBuilder().WithModalidad(m).Build() }},
		{48, "Nota Crédito Débito ICE", "notaElectronicaCreditoDebitoIce", "notaComputarizadaCreditoDebitoIce",
			func(m int) any { return invoices.NewNotaCreditoDebitoIceBuilder().WithModalidad(m).Build() }},
		{49, "Telecomunicaciones Zona Franca", "facturaElectronicaTelecomunicacionZF", "facturaComputarizadaTelecomunicacionZF",
			func(m int) any { return invoices.NewTelecomunicacionesZFBuilder().WithModalidad(m).Build() }},
		{50, "Hospitales / Clínicas Zona Franca", "facturaElectronicaHospitalClinicaZF", "facturaComputarizadaHospitalClinicaZF",
			func(m int) any { return invoices.NewHospitalClinicaZFBuilder().WithModalidad(m).Build() }},
		{51, "Engarrafadoras", "facturaElectronicaEngarrafadoras", "facturaComputarizadaEngarrafadoras",
			func(m int) any { return invoices.NewEngarrafadorasBuilder().WithModalidad(m).Build() }},
		{52, "Venta de Minerales al Banco Central", "facturaElectronicaVentaMineralBCB", "facturaComputarizadaVentaMineralBCB",
			func(m int) any { return invoices.NewVentaMineralBCBBuilder().WithModalidad(m).Build() }},
		{53, "Lubricantes IEHD", "facturaElectronicaImportacionComercializacionLubricantesIEHD", "facturaComputarizadaImportacionComercializacionLubricantesIEHD",
			func(m int) any { return invoices.NewLubricantesIehdBuilder().WithModalidad(m).Build() }},
		{54, "Insumos para Biodiésel", "facturaElectronicaBiodiesel", "facturaComputarizadaBiodiesel",
			func(m int) any { return invoices.NewBiodieselBuilder().WithModalidad(m).Build() }},
		{55, "Comercialización de Combustible", "facturaElectronicaVentaCombustibleSinSubvencion", "facturaComputarizadaVentaCombustibleSinSubvencion",
			func(m int) any { return invoices.NewVentaCombustibleSinSubvencionBuilder().WithModalidad(m).Build() }},
	}
}

// TestSectores_ElementoRaiz verifica que cada builder emita el elemento raíz correcto en
// ambas modalidades. Un elemento raíz equivocado es rechazado con el código 939.
func TestSectores_ElementoRaiz(t *testing.T) {
	for _, c := range sectorCases() {
		t.Run(c.nombre, func(t *testing.T) {
			for _, m := range []struct {
				modalidad int
				esperado  string
			}{
				{siat.ModalidadElectronica, c.electronica},
				{siat.ModalidadComputarizada, c.computarizada},
			} {
				raw, err := xml.Marshal(c.build(m.modalidad))
				if err != nil {
					t.Fatalf("sector %d: error al serializar: %v", c.codigo, err)
				}
				if !strings.HasPrefix(string(raw), "<"+m.esperado) {
					t.Errorf("sector %d modalidad %d: elemento raíz es %.70s, se esperaba <%s",
						c.codigo, m.modalidad, raw, m.esperado)
				}
			}
		})
	}
}

// TestSectores_ElementoRaizUnico verifica que dos sectores distintos no compartan elemento
// raíz. Clonar un archivo de sector y olvidar renombrar el elemento produce documentos que
// compilan pero llegan al SIAT identificados como otro sector.
//
// El código 24 es la excepción conocida: NotaCreditoDebito y NotaFiscalCreditoDebito son
// dos diseños del mismo documento y comparten raíz a propósito.
func TestSectores_ElementoRaizUnico(t *testing.T) {
	vistos := map[string]int{}

	for _, c := range sectorCases() {
		if anterior, existe := vistos[c.electronica]; existe && anterior != c.codigo {
			t.Errorf("elemento raíz %q compartido por los sectores %d y %d",
				c.electronica, anterior, c.codigo)
		}
		vistos[c.electronica] = c.codigo
	}
}

// TestSectores_CoberturaCatalogo verifica que la tabla de casos cubra exactamente los
// sectores que el SDK declara soportar. Si alguien agrega un builder sin sumarlo acá,
// este test lo detecta antes de que el sector quede sin ninguna prueba.
func TestSectores_CoberturaCatalogo(t *testing.T) {
	// Códigos del catálogo normativo del SIAT que el SDK implementa.
	// El 33 (Tasa Cero IVA Ley N° 1613) está en el catálogo pero no tiene builder.
	esperados := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 28, 29, 30, 31, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
		44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
	}

	cubiertos := map[int]bool{}
	for _, c := range sectorCases() {
		cubiertos[c.codigo] = true
	}

	for _, codigo := range esperados {
		if !cubiertos[codigo] {
			t.Errorf("el sector %d no tiene ningún caso en la tabla", codigo)
		}
	}

	if len(cubiertos) != len(esperados) {
		t.Errorf("la tabla cubre %d sectores, el SDK implementa %d",
			len(cubiertos), len(esperados))
	}

	if cubiertos[33] {
		t.Error("el sector 33 no tiene builder en el SDK; si se agregó, actualizá la documentación")
	}
}

func TestSectores_TipoFacturaDocumentoPredeterminado(t *testing.T) {
	catalogo := invoices.DefinicionesTipoFactura()
	definiciones := make(map[int]int, len(catalogo))
	for _, definicion := range catalogo {
		if _, duplicado := definiciones[definicion.CodigoDocumentoSector]; duplicado {
			t.Fatalf("el sector %d está duplicado en DefinicionesTipoFactura", definicion.CodigoDocumentoSector)
		}
		definiciones[definicion.CodigoDocumentoSector] = definicion.TipoFacturaDocumento
	}
	for _, c := range sectorCases() {
		t.Run(c.nombre, func(t *testing.T) {
			documento, ok := c.build(siat.ModalidadElectronica).(models.FacturaConMetadatos)
			if !ok {
				t.Fatal("la factura no expone sus metadatos fiscales")
			}
			if obtenido := documento.CodigoDocumentoSector(); obtenido != c.codigo {
				t.Fatalf("sector del metadato = %d; se esperaba %d", obtenido, c.codigo)
			}
			esperado, existe := definiciones[c.codigo]
			if !existe {
				t.Fatalf("el sector %d no está definido en DefinicionesTipoFactura", c.codigo)
			}
			if obtenido := documento.TipoFacturaDocumento(); obtenido != esperado {
				t.Fatalf("tipo de factura = %d; se esperaba %d", obtenido, esperado)
			}
		})
	}
	if len(definiciones) != len(sectorCases())-1 {
		t.Fatalf("definiciones = %d; se esperaban %d sectores únicos", len(definiciones), len(sectorCases())-1)
	}
}

func TestSectores_CatalogoTiposFacturaEsInmutable(t *testing.T) {
	catalogo := invoices.DefinicionesTipoFactura()
	original := catalogo[0].TipoFacturaDocumento
	catalogo[0].TipoFacturaDocumento = 999
	if obtenido := invoices.DefinicionesTipoFactura()[0].TipoFacturaDocumento; obtenido != original {
		t.Fatalf("el catálogo público fue modificado: %d", obtenido)
	}
}
