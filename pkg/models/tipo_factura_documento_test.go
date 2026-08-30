package models_test

import (
	"encoding/xml"
	"errors"
	"strings"
	"testing"

	"github.com/ron86i/go-siat/v2"
	"github.com/ron86i/go-siat/v2/pkg/models"
)

type documentoFiscalPrueba struct{}

func (documentoFiscalPrueba) TipoFacturaDocumento() int { return siat.TipoNotaCreditoDebito }

func (documentoFiscalPrueba) CodigoDocumentoSector() int { return 24 }

func (documentoFiscalPrueba) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	return e.EncodeElement(struct{}{}, xml.StartElement{Name: xml.Name{Local: "notaPrueba"}})
}

func TestTipoFacturaDocumentoPredeterminado(t *testing.T) {
	casos := map[int]int{
		1:  siat.TipoFacturaConDerechoCreditoFiscal,
		3:  siat.TipoFacturaSinDerechoCreditoFiscal,
		6:  siat.TipoFacturaSinDerechoCreditoFiscal,
		8:  siat.TipoFacturaSinDerechoCreditoFiscal,
		13: siat.TipoFacturaConDerechoCreditoFiscal,
		24: siat.TipoNotaCreditoDebito,
		30: siat.TipoBoletoAereo,
		47: siat.TipoNotaCreditoDebito,
		52: siat.TipoFacturaSinDerechoCreditoFiscal,
	}
	for sector, esperado := range casos {
		if obtenido := models.TipoFacturaDocumentoPredeterminado(sector); obtenido != esperado {
			t.Errorf("sector %d: tipo = %d; se esperaba %d", sector, obtenido, esperado)
		}
	}
}

func TestRecepcionesAdoptanMetadatosDelDocumento(t *testing.T) {
	casos := []struct {
		nombre string
		build  func() (any, error)
	}{
		{
			nombre: "factura",
			build: func() (any, error) {
				solicitud := models.NewRecepcionFacturaBuilder()
				if err := solicitud.WithDocumentoFiscal(documentoFiscalPrueba{}, nil); err != nil {
					return nil, err
				}
				return solicitud.Build(), nil
			},
		},
		{
			nombre: "documento de ajuste",
			build: func() (any, error) {
				solicitud := models.NewRecepcionDocumentoAjusteBuilder()
				if err := solicitud.WithDocumentoFiscal(documentoFiscalPrueba{}, nil); err != nil {
					return nil, err
				}
				return solicitud.Build(), nil
			},
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			solicitud, err := caso.build()
			if err != nil {
				t.Fatalf("serializar documento: %v", err)
			}
			raw, err := xml.Marshal(solicitud)
			if err != nil {
				t.Fatalf("serializar solicitud: %v", err)
			}
			if !strings.Contains(string(raw), "<tipoFacturaDocumento>3</tipoFacturaDocumento>") {
				t.Fatalf("la solicitud no adoptó el tipo del documento: %s", raw)
			}
			if !strings.Contains(string(raw), "<codigoDocumentoSector>24</codigoDocumentoSector>") {
				t.Fatalf("la solicitud no adoptó el sector del documento: %s", raw)
			}
		})
	}
}

func TestRecepcionFacturaRespetaMetadatosIndicadosManualmente(t *testing.T) {
	solicitud := models.NewRecepcionFacturaBuilder().
		WithCodigoDocumentoSector(99).
		WithTipoFacturaDocumento(siat.TipoFacturaConDerechoCreditoFiscal)
	if err := solicitud.WithFactura(documentoFiscalPrueba{}, nil); err != nil {
		t.Fatalf("serializar documento: %v", err)
	}
	raw, err := xml.Marshal(solicitud.Build())
	if err != nil {
		t.Fatalf("serializar solicitud: %v", err)
	}
	if !strings.Contains(string(raw), "<codigoDocumentoSector>99</codigoDocumentoSector>") {
		t.Fatalf("la solicitud sobrescribió el sector manual: %s", raw)
	}
	if !strings.Contains(string(raw), "<tipoFacturaDocumento>1</tipoFacturaDocumento>") {
		t.Fatalf("la solicitud sobrescribió el tipo manual: %s", raw)
	}
}

func TestRecepcionesExigenFirmaElectronicaCuandoSeSolicita(t *testing.T) {
	casos := []struct {
		nombre   string
		preparar func() error
	}{
		{
			nombre: "factura",
			preparar: func() error {
				return models.NewRecepcionFacturaBuilder().
					WithCodigoModalidad(siat.ModalidadElectronica).
					WithFirmaElectronicaRequerida().
					WithDocumentoFiscal(documentoFiscalPrueba{}, nil)
			},
		},
		{
			nombre: "documento de ajuste",
			preparar: func() error {
				return models.NewRecepcionDocumentoAjusteBuilder().
					WithCodigoModalidad(siat.ModalidadElectronica).
					WithFirmaElectronicaRequerida().
					WithDocumentoFiscal(documentoFiscalPrueba{}, nil)
			},
		},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			if err := caso.preparar(); !errors.Is(err, models.ErrFirmaElectronicaRequerida) {
				t.Fatalf("error = %v; se esperaba ErrFirmaElectronicaRequerida", err)
			}
		})
	}
}
