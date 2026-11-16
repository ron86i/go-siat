package facturacion_test

import (
	"testing"
	"time"

	compraventaV3 "github.com/ron86i/go-siat/v3/documents/compraventa"
	serviciosbasicosV3 "github.com/ron86i/go-siat/v3/documents/serviciosbasicos"
	tasaceroV3 "github.com/ron86i/go-siat/v3/documents/tasacero"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type documentoProtocoloAjeno struct{}

func (documentoProtocoloAjeno) Validate() error                         { return nil }
func (documentoProtocoloAjeno) CodigoDocumentoSector() int              { return 1 }
func (documentoProtocoloAjeno) CodigoSucursal() int                     { return 0 }
func (documentoProtocoloAjeno) CodigoPuntoVenta() *int                  { return nil }
func (documentoProtocoloAjeno) Cufd() string                            { return "CUFD" }
func (documentoProtocoloAjeno) Protocolo() int                          { return 2 }
func (documentoProtocoloAjeno) MarshalXMLModalidad(int) ([]byte, error) { return nil, nil }

func TestIncompatibilidadProtocolo_RechazaModeloAjeno(t *testing.T) {
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: facturacion.ModalidadComputarizada}
	_, err := req.PrepararFactura(documentoProtocoloAjeno{}, nil)
	require.ErrorContains(t, err, "documento incompatible con SIAT v3: el documento pertenece al protocolo v2")

	paquete := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: req}
	_, err = paquete.PrepararPaquete(t.Context(), []facturacion.Documento{documentoProtocoloAjeno{}}, nil)
	require.ErrorContains(t, err, "documento 1: documento incompatible con SIAT v3: el documento pertenece al protocolo v2")
}

func TestPrepararFacturaGenerica_AceptaCompraVentaYAlquileres(t *testing.T) {
	req := facturacion.SolicitudRecepcionFactura{
		CodigoModalidad: facturacion.ModalidadComputarizada,
		Cufd:            "CUFD-TEST",
		Cuis:            "CUIS-TEST",
		FechaEnvio:      time.Now(),
	}

	// 1. Compra-Venta via API genérica PrepararFactura
	cvDoc := newTestFactura(t, func(c *compraventaV3.Cabecera, d *[]compraventaV3.Detalle) {
		c.Cufd = "CUFD-TEST"
	})
	resCV, err := req.PrepararFactura(cvDoc, nil)
	require.NoError(t, err)
	assert.Equal(t, 1, resCV.CodigoDocumentoSector)
	assert.NotEmpty(t, resCV.Archivo)
	assert.NotEmpty(t, resCV.HashArchivo)

	// 2. Alquileres via misma API genérica PrepararFactura
	alqDoc := newTestAlquilerFactura(t, 0, "CUFD-TEST")
	resAlq, err := req.PrepararFactura(alqDoc, nil)
	require.NoError(t, err)
	assert.Equal(t, 2, resAlq.CodigoDocumentoSector)
	assert.NotEmpty(t, resAlq.Archivo)
	assert.NotEmpty(t, resAlq.HashArchivo)

	// 3. PrepararPaquete genérico con mezcla de sectores debe ser rechazado
	paq := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: req}
	_, err = paq.PrepararPaquete(t.Context(), []facturacion.Documento{cvDoc, alqDoc}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "sector incompatible")
}

func TestPrepararFactura_ToleranciaPunteroNil(t *testing.T) {
	req := facturacion.SolicitudRecepcionFactura{CodigoModalidad: 2}

	// 1. Interface nil directa
	_, err := req.PrepararFactura(nil, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "el documento no puede ser nil")

	// 2. Typed nil pointer dentro de Documento
	var typedNil *compraventaV3.Factura = nil
	var doc facturacion.Documento = typedNil
	_, err = req.PrepararFactura(doc, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "el documento no puede ser nil")

	// 3. Paquete con typed nil pointer
	paq := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: req}
	_, err = paq.PrepararPaquete(t.Context(), []facturacion.Documento{doc}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "documento 1: el documento no puede ser nil")
}

func TestAsDocumentos_Ergonomia(t *testing.T) {
	cvDoc := newTestFactura(t)
	facturas := []*compraventaV3.Factura{cvDoc}

	// Usar helper genérico AsDocumentos para transformar []*compraventa.Factura a []Documento
	documentos := facturacion.AsDocumentos(facturas)
	require.Len(t, documentos, 1)
	assert.Equal(t, 1, documentos[0].CodigoDocumentoSector())
}

func TestPrepararFacturaGenerica_ServiciosBasicosV3(t *testing.T) {
	req := facturacion.SolicitudRecepcionFactura{
		CodigoModalidad: facturacion.ModalidadComputarizada,
		Cufd:            "CUFD-TEST",
		Cuis:            "CUIS-TEST",
		FechaEnvio:      time.Now(),
	}

	cab := serviciosbasicosV3.Cabecera{
		NitEmisor:             123,
		NumeroFactura:         1,
		Cuf:                   "CUF",
		Cufd:                  "CUFD-TEST",
		Direccion:             "Av. Principal 123",
		NumeroMedidor:         "MED-001",
		FechaEmision:          time.Now(),
		CodigoMetodoPago:      1,
		SubTotalNeto:          100,
		TotalPrecioNeto:       100,
		Iva:                   13,
		MontoAPagar:           100,
		CodigoMoneda:          1,
		TipoCambio:            1,
		MontoTotalMoneda:      100,
		Leyenda:               "Ley N 453",
		Usuario:               "user",
		CodigoDocumentoSector: 13,
	}
	det := []serviciosbasicosV3.Detalle{
		{
			ActividadEconomica: "360000",
			CodigoProductoSin:  11223,
			CodigoProducto:     "AGUA-01",
			Descripcion:        "Consumo Agua",
			Cantidad:           1,
			UnidadMedida:       1,
			PrecioNetoUnitario: 100,
			SubTotal:           100,
		},
	}
	sbDoc, err := serviciosbasicosV3.New(cab, det)
	require.NoError(t, err)
	assert.Equal(t, 3, sbDoc.Protocolo())
	assert.Equal(t, 13, sbDoc.CodigoDocumentoSector())

	// 1. PrepararFactura individual
	res, err := req.PrepararFactura(sbDoc, nil)
	require.NoError(t, err)
	assert.Equal(t, 13, res.CodigoDocumentoSector)
	assert.NotEmpty(t, res.Archivo)
	assert.NotEmpty(t, res.HashArchivo)

	// 2. PrepararPaquete con AsDocumentos
	paq := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: req}
	docs := facturacion.AsDocumentos([]*serviciosbasicosV3.Factura{sbDoc})
	resPaq, err := paq.PrepararPaquete(t.Context(), docs, nil)
	require.NoError(t, err)
	assert.Equal(t, 13, resPaq.CodigoDocumentoSector)
	assert.Equal(t, 1, resPaq.CantidadFacturas)
	assert.NotEmpty(t, resPaq.Archivo)

	// 3. PrepararMasiva
	mas := facturacion.SolicitudRecepcionMasiva{SolicitudRecepcionFactura: req}
	resMas, err := mas.PrepararMasiva(t.Context(), docs, nil)
	require.NoError(t, err)
	assert.Equal(t, 13, resMas.CodigoDocumentoSector)
	assert.Equal(t, 1, resMas.CantidadFacturas)
	assert.NotEmpty(t, resMas.Archivo)
}

func TestPrepararFacturaGenerica_TasaCeroV3(t *testing.T) {
	req := facturacion.SolicitudRecepcionFactura{
		CodigoModalidad: facturacion.ModalidadComputarizada,
		Cufd:            "CUFD-TEST",
		Cuis:            "CUIS-TEST",
		FechaEnvio:      time.Now(),
	}

	cab := tasaceroV3.Cabecera{
		NitEmisor:             123,
		NumeroFactura:         1,
		Cuf:                   "CUF",
		Cufd:                  "CUFD-TEST",
		Direccion:             "Av. 6 de Agosto",
		FechaEmision:          time.Now(),
		CodigoMetodoPago:      1,
		SubTotalNeto:          100,
		TotalPrecioNeto:       100,
		Iva:                   0,
		MontoAPagar:           100,
		CodigoMoneda:          1,
		TipoCambio:            1,
		MontoTotalMoneda:      100,
		Leyenda:               "Ley N 453",
		Usuario:               "user",
		CodigoDocumentoSector: 8,
	}
	det := []tasaceroV3.Detalle{
		{
			ActividadEconomica: "581100",
			CodigoProductoSin:  83141,
			CodigoProducto:     "LIBRO-01",
			Descripcion:        "Libro",
			Cantidad:           1,
			UnidadMedida:       58,
			PrecioNetoUnitario: 100,
			SubTotal:           100,
		},
	}
	tcDoc, err := tasaceroV3.New(cab, det)
	require.NoError(t, err)
	assert.Equal(t, 3, tcDoc.Protocolo())
	assert.Equal(t, 8, tcDoc.CodigoDocumentoSector())

	// 1. PrepararFactura individual
	res, err := req.PrepararFactura(tcDoc, nil)
	require.NoError(t, err)
	assert.Equal(t, 8, res.CodigoDocumentoSector)
	assert.NotEmpty(t, res.Archivo)
	assert.NotEmpty(t, res.HashArchivo)

	// 2. PrepararPaquete con AsDocumentos
	paq := facturacion.SolicitudRecepcionPaquete{SolicitudRecepcionFactura: req}
	docs := facturacion.AsDocumentos([]*tasaceroV3.Factura{tcDoc})
	resPaq, err := paq.PrepararPaquete(t.Context(), docs, nil)
	require.NoError(t, err)
	assert.Equal(t, 8, resPaq.CodigoDocumentoSector)
	assert.Equal(t, 1, resPaq.CantidadFacturas)
	assert.NotEmpty(t, resPaq.Archivo)

	// 3. PrepararMasiva
	mas := facturacion.SolicitudRecepcionMasiva{SolicitudRecepcionFactura: req}
	resMas, err := mas.PrepararMasiva(t.Context(), docs, nil)
	require.NoError(t, err)
	assert.Equal(t, 8, resMas.CodigoDocumentoSector)
	assert.Equal(t, 1, resMas.CantidadFacturas)
	assert.NotEmpty(t, resMas.Archivo)
}
