package service_test

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/ron86i/go-siat/internal/adapter/service"
	"github.com/ron86i/go-siat/internal/core/domain/datatype"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/codigos"
	"github.com/ron86i/go-siat/internal/core/domain/facturacion/compra_venta"
	"github.com/ron86i/go-siat/pkg/config"
	"github.com/ron86i/go-siat/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestSiatCompraVentaService_RecepcionFactura(t *testing.T) {
	godotenv.Load(".env")

	// ... (Mantener parseo de variables de entorno igual hasta la creación de servicios)
	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	serviceCodigos, _ := service.NewSiatCodigosService(os.Getenv("SIAT_URL"), client)

	// Solicitar CUIS y CUFD (Misma lógica que tenías)
	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), config, &codigos.Cuis{
		SolicitudCuis: codigos.SolicitudCuis{
			CodigoAmbiente: codAmbiente, CodigoModalidad: codModalidad,
			CodigoSistema: os.Getenv("SIAT_CODIGO_SISTEMA"), Nit: nit,
		},
	})
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), config, &codigos.Cufd{
		SolicitudCufd: codigos.SolicitudCufd{
			CodigoAmbiente: codAmbiente, CodigoModalidad: codModalidad,
			CodigoSistema: os.Getenv("SIAT_CODIGO_SISTEMA"), Nit: nit,
			Cuis: cuis.Body.Content.RespuestaCuis.Codigo,
		},
	})
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}

	serviceCompraVenta, _ := service.NewSiatCompraVentaService(os.Getenv("SIAT_URL"), client)

	fechaEmision := time.Now()
	// 1. Generar CUF
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 1, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// Crear objeto de factura
	factura := compra_venta.FacturaCompraVenta{
		XMLName:           xml.Name{Local: "facturaElectronicaCompraVenta"},
		XmlnsXsi:          "http://www.w3.org/2001/XMLSchema-instance",
		XsiSchemaLocation: "facturaElectronicaCompraVenta.xsd",
		Cabecera: compra_venta.Cabecera{
			NitEmisor:                    nit,
			RazonSocialEmisor:            "Ronaldo Rua",
			Municipio:                    "Tarija",
			NumeroFactura:                1,
			Cuf:                          cuf,
			Cufd:                         cufd.Body.Content.RespuestaCufd.Codigo,
			CodigoSucursal:               0,
			Direccion:                    "ESQUINA AVENIDA LA PAZ",
			CodigoPuntoVenta:             0,
			FechaEmision:                 fechaEmision.Format("2006-01-02T15:04:05.000"), // Formato SIAT preciso
			NombreRazonSocial:            compra_venta.Nilable[string]{Value: new("JUAN PEREZ")},
			CodigoTipoDocumentoIdentidad: 1,
			NumeroDocumento:              "5115889",
			CodigoCliente:                "1",
			CodigoMetodoPago:             1,
			MontoTotal:                   100,
			MontoTotalSujetoIva:          100,
			CodigoMoneda:                 1,
			TipoCambio:                   1,
			MontoTotalMoneda:             100,
			Leyenda:                      "Ley N° 453: Tienes derecho a recibir información...",
			Usuario:                      "usuario",
			CodigoDocumentoSector:        1,
		},
		Detalle: []compra_venta.Detalle{{
			ActividadEconomica: "477300", CodigoProductoSin: "622539", CodigoProducto: "abc123",
			Descripcion: "GASA", Cantidad: 1, UnidadMedida: 1, PrecioUnitario: 100, SubTotal: 100,
		}},
	}

	// 2. Serializar y Firmar
	xmlData, _ := xml.Marshal(factura)
	signedXML, err := utils.SignXML(xmlData, "key.pem", "cert.crt")
	if err != nil {
		t.Fatalf("error firmando XML: %v", err)
	}

	// 3. Comprimir con Gzip
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(signedXML); err != nil {
		t.Fatalf("error comprimiendo: %v", err)
	}
	zw.Close()

	compressedBytes := buf.Bytes()

	// 4. Calcular Hash SHA256 sobre los bytes COMPRIMIDOS (antes del base64)
	hash := sha256.Sum256(compressedBytes)
	hashString := hex.EncodeToString(hash[:])

	// 5. CODIFICAR A BASE64
	encodedArchivo := base64.StdEncoding.EncodeToString(compressedBytes)

	req := &compra_venta.RecepcionFactura{
		SolicitudServicioRecepcionFactura: compra_venta.SolicitudRecepcionFactura{
			SolicitudRecepcion: compra_venta.SolicitudRecepcion{
				CodigoAmbiente: codAmbiente, CodigoModalidad: codModalidad,
				CodigoSistema: os.Getenv("SIAT_CODIGO_SISTEMA"), Nit: nit,
				CodigoSucursal: 0, CodigoDocumentoSector: 1, CodigoEmision: 1,
				CodigoPuntoVenta: 0, Cufd: cufd.Body.Content.RespuestaCufd.Codigo,
				Cuis: cuis.Body.Content.RespuestaCuis.Codigo, TipoFacturaDocumento: 1,
			},
			Archivo:     encodedArchivo, // Se envía la cadena base64 como slice de bytes
			FechaEnvio:  datatype.NewTimeSiat(fechaEmision),
			HashArchivo: hashString,
		},
	}

	resp, err := serviceCompraVenta.RecepcionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta SIAT: %+v", resp.Body.Content)
}

func TestSiatCompraVentaService_AnulacionFactura(t *testing.T) {
	godotenv.Load(".env")

	codModalidad, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_MODALIDAD"))
	nit, _ := utils.ParseInt64Safe(os.Getenv("SIAT_NIT"))
	codAmbiente, _ := utils.ParseIntSafe(os.Getenv("SIAT_CODIGO_AMBIENTE"))
	config := config.Config{Token: os.Getenv("SIAT_TOKEN")}

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyFromEnvironment}}
	serviceCodigos, _ := service.NewSiatCodigosService(os.Getenv("SIAT_URL"), client)

	// Solicitar CUIS y CUFD
	cuis, err := serviceCodigos.SolicitudCuis(context.Background(), config, &codigos.Cuis{
		SolicitudCuis: codigos.SolicitudCuis{
			CodigoAmbiente: codAmbiente, CodigoModalidad: codModalidad,
			CodigoSistema: os.Getenv("SIAT_CODIGO_SISTEMA"), Nit: nit,
		},
	})
	if err != nil {
		t.Fatalf("error CUIS: %v", err)
	}

	cufd, err := serviceCodigos.SolicitudCufd(context.Background(), config, &codigos.Cufd{
		SolicitudCufd: codigos.SolicitudCufd{
			CodigoAmbiente: codAmbiente, CodigoModalidad: codModalidad,
			CodigoSistema: os.Getenv("SIAT_CODIGO_SISTEMA"), Nit: nit,
			Cuis: cuis.Body.Content.RespuestaCuis.Codigo,
		},
	})
	if err != nil {
		t.Fatalf("error CUFD: %v", err)
	}

	serviceCompraVenta, _ := service.NewSiatCompraVentaService(os.Getenv("SIAT_URL"), client)

	fechaEmision := time.Now()
	// Generar CUF de la factura que supuestamente vamos a anular
	cuf, err := utils.GenerarCUF(nit, fechaEmision, 0, codModalidad, 1, 1, 1, 1, 0, cufd.Body.Content.RespuestaCufd.CodigoControl)
	if err != nil {
		t.Fatalf("error al generar CUF: %v", err)
	}

	// Usar la estructura directamente sin Builder
	req := &compra_venta.AnulacionFactura{
		SolicitudAnulacion: compra_venta.SolicitudAnulacion{
			SolicitudRecepcion: compra_venta.SolicitudRecepcion{
				CodigoAmbiente:        codAmbiente,
				CodigoDocumentoSector: 1,
				CodigoEmision:         1,
				CodigoModalidad:       codModalidad,
				CodigoPuntoVenta:      0,
				CodigoSistema:         os.Getenv("SIAT_CODIGO_SISTEMA"),
				CodigoSucursal:        0,
				Cufd:                  cufd.Body.Content.RespuestaCufd.Codigo,
			},
			Cuf:          cuf,
			CodigoMotivo: 1,
		},
	}

	resp, err := serviceCompraVenta.AnulacionFactura(context.Background(), config, req)
	if err != nil {
		t.Fatalf("error en solicitud de anulación: %v", err)
	}

	assert.NotNil(t, resp)
	log.Printf("Respuesta Anulación SIAT: %+v", resp.Body.Content)
}
