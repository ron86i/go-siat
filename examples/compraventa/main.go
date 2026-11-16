// Package main ilustra la construcción y emisión de una factura en SIAT v3 (3.0).
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ron86i/go-siat/v3/documents/compraventa"
	"github.com/ron86i/go-siat/v3/pkg/utils"
	"github.com/ron86i/go-siat/v3/protocols"
	"github.com/ron86i/go-siat/v3/protocols/codigos"
	"github.com/ron86i/go-siat/v3/protocols/facturacion"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 1. Configuración del cliente para protocolo SIAT v3
	config := protocols.Config{
		BaseURL:        "https://pilotosiatservicios.impuestos.gob.bo/sfd",
		Token:          "TuTokenDelegadoSiat",
		Nit:            123456789,
		CodigoSistema:  "SISTEMA-VENTAS-TEST",
		CodigoAmbiente: 2, // 1: Producción, 2: Pruebas / Piloto
	}

	client, err := protocols.NewClient(config)
	if err != nil {
		log.Fatalf("Error al inicializar cliente SIAT v3: %v", err)
	}

	sucursal := 0
	puntoVenta := 0
	modalidad := 2   // 1: Electrónica en Línea, 2: Computarizada en Línea
	tipoEmision := 1 // 1: En Línea, 2: Fuera de línea (contingencia)
	tipoFactura := 1 // 1: Con derecho a crédito fiscal
	numeroFactura := int64(1001)
	fechaEmision := time.Now()

	fmt.Println("--- Paso 1: Obtener CUIS ---")
	cuisResp, err := client.Codigos.SolicitarCUIS(ctx, codigos.SolicitudCUIS{
		CodigoModalidad:  modalidad,
		CodigoSucursal:   sucursal,
		CodigoPuntoVenta: &puntoVenta,
	})
	if err != nil {
		fmt.Printf("Nota: llamada de prueba a CUIS simulada o sin conexión activa: %v\n", err)
		// En un entorno offline o demostración usamos un código de prueba
		cuisResp = &codigos.RespuestaCUIS{Codigo: "CUIS-TEST-12345"}
	}
	fmt.Printf("CUIS obtenido: %s\n", cuisResp.Codigo)

	fmt.Println("\n--- Paso 2: Obtener CUFD ---")
	cufdResp, err := client.Codigos.SolicitarCUFD(ctx, codigos.SolicitudCUFD{
		CodigoModalidad:  modalidad,
		CodigoSucursal:   sucursal,
		CodigoPuntoVenta: &puntoVenta,
		Cuis:             cuisResp.Codigo,
	})
	if err != nil {
		fmt.Printf("Nota: llamada de prueba a CUFD simulada: %v\n", err)
		cufdResp = &codigos.RespuestaCUFD{Codigo: "CUFD-TEST-ABCDE", CodigoControl: "A1B2C3D4"}
	}
	fmt.Printf("CUFD obtenido: %s (Control: %s)\n", cufdResp.Codigo, cufdResp.CodigoControl)

	fmt.Println("\n--- Paso 3: Generar CUF ---")
	cuf, err := utils.CUFParams{
		Nit:                 config.Nit,
		FechaHora:           fechaEmision,
		Sucursal:            sucursal,
		Modalidad:           modalidad,
		TipoEmision:         tipoEmision,
		TipoFactura:         tipoFactura,
		TipoDocumentoSector: 1, // Compra-Venta
		NumeroFactura:       numeroFactura,
		PuntoVenta:          puntoVenta,
		CodigoControl:       cufdResp.CodigoControl,
	}.Generate()
	if err != nil {
		log.Fatalf("Error al generar CUF: %v", err)
	}
	fmt.Printf("CUF generado: %s\n", cuf)

	fmt.Println("\n--- Paso 4: Construir Factura con compraventa.New() ---")
	nombreCliente := "JUAN PEREZ"
	telefono := "22114455"
	descuentoZero := float64(0.0)

	cabecera := compraventa.Cabecera{
		NitEmisor:                    config.Nit,
		RazonSocialEmisor:            "MI EMPRESA S.R.L.",
		Municipio:                    "La Paz",
		Telefono:                     &telefono,
		NumeroFactura:                numeroFactura,
		Cuf:                          cuf,
		Cufd:                         cufdResp.Codigo,
		CodigoSucursal:               sucursal,
		Direccion:                    "Av. 6 de Agosto Nro 1234",
		CodigoPuntoVenta:             &puntoVenta,
		FechaEmision:                 fechaEmision,
		NombreRazonSocial:            &nombreCliente,
		CodigoTipoDocumentoIdentidad: 1, // CI
		NumeroDocumento:              "1234567",
		CodigoCliente:                "CLI-001",
		CodigoMetodoPago:             1, // Efectivo
		// En SIAT v3 los campos son netos (SubTotalNeto, TotalPrecioNeto, Iva, MontoAPagar):
		SubTotalNeto:          100.00,
		TotalPrecioNeto:       100.00,
		Iva:                   13.00,
		MontoAPagar:           100.00,
		CodigoMoneda:          1, // Bolivianos
		TipoCambio:            1.0,
		MontoTotalMoneda:      100.00,
		Leyenda:               "Ley N° 453: Tienes derecho a un trato equitativo y digno.",
		Usuario:               "operador1",
		CodigoInterno:         "FAC-2026-001",
		CodigoDocumentoSector: 1,
	}

	detalles := []compraventa.Detalle{
		{
			ActividadEconomica: "620100",
			CodigoProductoSin:  83141,
			CodigoProducto:     "SERV-001",
			Descripcion:        "Servicios de desarrollo de software",
			Cantidad:           1.0,
			UnidadMedida:       58, // Unidad (Servicios)
			PrecioNetoUnitario: 100.00,
			MontoDescuento:     &descuentoZero,
			SubTotal:           100.00,
		},
	}

	// Construcción asistida y validación obligatoria con New()
	factura, err := compraventa.New(cabecera, detalles)
	if err != nil {
		log.Fatalf("Error al construir factura: %v", err)
	}

	// Serialización directa a XML SIAT v3 (utilizando la modalidad deseada)
	xmlData, err := factura.XML(modalidad)
	if err != nil {
		log.Fatalf("Error al serializar XML: %v", err)
	}
	fmt.Printf("XML generado correctamente (%d bytes)\n", len(xmlData))

	fmt.Println("\n--- Paso 5: Preparar y emitir al SIAT ---")
	solicitud := facturacion.SolicitudRecepcionFactura{
		CodigoModalidad: modalidad, // Modalidad configurada en la solicitud/servicio
		Cuis:            cuisResp.Codigo,
	}

	// PrepararFactura valida, comprime a GZIP y calcula el hash SHA256 automáticamente
	solicitudPreparada, err := solicitud.PrepararFactura(factura, nil)
	if err != nil {
		log.Fatalf("Error al preparar solicitud: %v", err)
	}
	fmt.Printf("Solicitud lista con archivo comprimido (%d bytes, SHA256: %s)\n",
		len(solicitudPreparada.Archivo), solicitudPreparada.HashArchivo)

	// Recepción individual en el SIAT
	resp, err := client.Facturacion.RecepcionFactura(ctx, solicitudPreparada)
	if err != nil {
		fmt.Printf("Envío simulado (sin conexión real al SIN): %v\n", err)
		return
	}

	fmt.Printf("Estado SIN: %s (Código: %d, Recepción: %s)\n",
		resp.CodigoDescripcion, resp.CodigoEstado, resp.CodigoRecepcion)
}
