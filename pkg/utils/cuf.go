package utils

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

// CUFParams agrupa todos los parámetros necesarios para generar un CUF (Código Único de Factura).
//
// Es la alternativa recomendada a [GenerarCUF] cuando se construye el CUF dentro de un contexto
// con muchas variables, ya que los campos con nombre previenen errores silenciosos por orden
// incorrecto de parámetros.
//
// Ejemplo de uso:
//
//	cuf, err := utils.CUFParams{
//		Nit:                 tc.Nit,
//		FechaHora:           fechaEmision,
//		Sucursal:            tc.Sucursal,
//		Modalidad:           tc.Modalidad,
//		TipoEmision:         siat.EmisionOnline,
//		TipoFactura:         1,
//		TipoDocumentoSector: 1,
//		NumeroFactura:       1,
//		PuntoVenta:          tc.PuntoVenta,
//		CodigoControl:       cufdControl,
//	}.Generar()
type CUFParams struct {
	// Nit es el NIT del contribuyente emisor (13 dígitos).
	Nit int64
	// FechaHora es la fecha y hora exacta de emisión de la factura.
	FechaHora time.Time
	// Sucursal es el código de la sucursal emisora (4 dígitos).
	Sucursal int
	// Modalidad es el código de modalidad de facturación (1 = Electrónica, 2 = Computarizada).
	Modalidad int
	// TipoEmision es el código de tipo de emisión (1 = Online, 2 = Offline).
	TipoEmision int
	// TipoFactura es el tipo de documento fiscal (1 = Con crédito fiscal, 2 = Sin crédito fiscal, etc.).
	TipoFactura int
	// TipoDocumentoSector es el código del sector del documento (2 dígitos), ej. 1 = Compra y Venta.
	TipoDocumentoSector int
	// NumeroFactura es el número correlativo de la factura (10 dígitos).
	NumeroFactura int64
	// PuntoVenta es el código del punto de venta (4 dígitos).
	PuntoVenta int
	// CodigoControl es la parte de control extraída del CUFD vigente.
	CodigoControl string
}

// Generate genera el CUF completo a partir de los parámetros del receptor.
// Es la variante recomendada de [GenerarCUF] cuando se trabaja con structs o se
// prefiere claridad en los parámetros.
//
// El CUF se calcula siguiendo el algoritmo oficial del SIAT:
//  1. Formateo de cada campo a longitud fija.
//  2. Concatenación de la cadena base (53 dígitos).
//  3. Cálculo del dígito verificador Módulo 11.
//  4. Conversión de la cadena numérica a hexadecimal en mayúsculas.
//  5. Concatenación del código de control (extraído del CUFD).
//
// Ejemplo:
//
//	cuf, err := utils.CUFParams{
//		Nit:       tc.Nit,
//		FechaHora: fechaEmision,
//		// ...
//	}.Generate()
func (p CUFParams) Generate() (string, error) {
	return GenerarCUF(p.Nit, p.FechaHora, p.Sucursal, p.Modalidad, p.TipoEmision,
		p.TipoFactura, p.TipoDocumentoSector, p.NumeroFactura, p.PuntoVenta, p.CodigoControl)
}

// CUFBuilder construye un CUF de forma fluida encadenando llamadas With*.
// El método terminal es [CUFBuilder.Generate], que ejecuta el cálculo y retorna el CUF.
//
// Ejemplo:
//
//	cuf, err := utils.NewCUF().
//		WithNit(tc.Nit).
//		WithFechaHora(fechaEmision).
//		WithSucursal(tc.Sucursal).
//		WithModalidad(tc.Modalidad).
//		WithTipoEmision(siat.EmisionOnline).
//		WithTipoFactura(1).
//		WithTipoDocumentoSector(34).
//		WithNumeroFactura(int64(nroFactura)).
//		WithPuntoVenta(tc.PuntoVenta).
//		WithCodigoControl(cufdControl).
//		Generate()
type CUFBuilder struct {
	params CUFParams
}

// NewCUF crea un nuevo [CUFBuilder] listo para encadenar parámetros.
func NewCUF() *CUFBuilder {
	return &CUFBuilder{}
}

// WithNit establece el NIT del contribuyente emisor (13 dígitos).
func (b *CUFBuilder) WithNit(nit int64) *CUFBuilder {
	b.params.Nit = nit
	return b
}

// WithFechaHora establece la fecha y hora exacta de emisión de la factura.
func (b *CUFBuilder) WithFechaHora(t time.Time) *CUFBuilder {
	b.params.FechaHora = t
	return b
}

// WithSucursal establece el código de la sucursal emisora (4 dígitos).
func (b *CUFBuilder) WithSucursal(sucursal int) *CUFBuilder {
	b.params.Sucursal = sucursal
	return b
}

// WithModalidad establece el código de modalidad de facturación.
// (1 = Electrónica, 2 = Computarizada).
func (b *CUFBuilder) WithModalidad(modalidad int) *CUFBuilder {
	b.params.Modalidad = modalidad
	return b
}

// WithTipoEmision establece el tipo de emisión.
// (1 = Online/En línea, 2 = Offline/Contingencia).
func (b *CUFBuilder) WithTipoEmision(tipoEmision int) *CUFBuilder {
	b.params.TipoEmision = tipoEmision
	return b
}

// WithTipoFactura establece el tipo de documento fiscal.
// (1 = Con crédito fiscal, 2 = Sin crédito fiscal, etc.)
func (b *CUFBuilder) WithTipoFactura(tipoFactura int) *CUFBuilder {
	b.params.TipoFactura = tipoFactura
	return b
}

// WithTipoDocumentoSector establece el código del sector del documento (2 dígitos).
// Por ejemplo: 1 = Compra y Venta, 34 = Seguros, 22 = Telecomunicaciones.
func (b *CUFBuilder) WithTipoDocumentoSector(sector int) *CUFBuilder {
	b.params.TipoDocumentoSector = sector
	return b
}

// WithNumeroFactura establece el número correlativo de la factura (10 dígitos).
func (b *CUFBuilder) WithNumeroFactura(n int64) *CUFBuilder {
	b.params.NumeroFactura = n
	return b
}

// WithPuntoVenta establece el código del punto de venta (4 dígitos).
func (b *CUFBuilder) WithPuntoVenta(puntoVenta int) *CUFBuilder {
	b.params.PuntoVenta = puntoVenta
	return b
}

// WithCodigoControl establece el código de control extraído del CUFD vigente.
func (b *CUFBuilder) WithCodigoControl(codigoControl string) *CUFBuilder {
	b.params.CodigoControl = codigoControl
	return b
}

// Generate ejecuta el cálculo del CUF con todos los parámetros configurados
// y retorna el CUF completo o un error si el cálculo falla.
func (b *CUFBuilder) Generate() (string, error) {
	return b.params.Generate()
}

// GenerarCUF genera el CUF (Código Único de Factura) completo a partir de parámetros posicionales.
//
// Considera usar [GenerarCUFDesde] con [CUFParams] para mayor legibilidad cuando el sitio
// de llamada tenga muchas variables, ya que los parámetros posicionales del mismo tipo (int)
// son susceptibles a errores de orden silenciosos.
//
// El CUF se calcula siguiendo el algoritmo oficial del SIAT:
//  1. Formateo de cada campo a longitud fija.
//  2. Concatenación de la cadena base (53 dígitos).
//  3. Cálculo del dígito verificador Módulo 11.
//  4. Conversión de la cadena numérica a hexadecimal en mayúsculas.
//  5. Concatenación del código de control (extraído del CUFD).
//
// Parámetros:
//   - nit: NIT del contribuyente (13 dígitos).
//   - fechaHora: fecha y hora exacta de emisión.
//   - sucursal: código de sucursal (4 dígitos).
//   - modalidad: modalidad de facturación (1 = Electrónica, 2 = Computarizada).
//   - tipoEmision: tipo de emisión (1 = Online, 2 = Offline).
//   - tipoFactura: tipo de documento fiscal.
//   - tipoDocumentoSector: código de sector del documento (2 dígitos).
//   - numeroFactura: número correlativo de la factura (10 dígitos).
//   - puntoVenta: código del punto de venta (4 dígitos).
//   - codigoControl: código de control extraído del CUFD vigente.
func GenerarCUF(nit int64, fechaHora time.Time, sucursal, modalidad, tipoEmision, tipoFactura, tipoDocumentoSector int, numeroFactura int64, puntoVenta int, codigoControl string) (string, error) {

	// 1. Format fields to fixed length
	nitStr := fmt.Sprintf("%013d", nit)                                // 13
	fechaStr := fechaHora.Format("20060102150405.000")                 // 17 yyyyMMddHHmmss.SSS
	fechaStr = strings.ReplaceAll(fechaStr, ".", "")                   // 17 digits
	sucursalStr := fmt.Sprintf("%04d", sucursal)                       // 4
	modalidadStr := fmt.Sprintf("%1d", modalidad)                      // 1
	tipoEmisionStr := fmt.Sprintf("%1d", tipoEmision)                  // 1
	tipoFacturaStr := fmt.Sprintf("%1d", tipoFactura)                  // 1
	tipoDocumentoSectorStr := fmt.Sprintf("%02d", tipoDocumentoSector) // 2
	numeroFacturaStr := fmt.Sprintf("%010d", numeroFactura)            // 10
	puntoVentaStr := fmt.Sprintf("%04d", puntoVenta)                   // 4

	// 2. Concatenate everything up to point of sale
	cadena := nitStr + fechaStr + sucursalStr + modalidadStr + tipoEmisionStr +
		tipoFacturaStr + tipoDocumentoSectorStr + numeroFacturaStr + puntoVentaStr

	// 3. Calculate Modulo 11 verify digit
	digitoVerificador := calculaDigitoMod11(cadena, 1, 9, false)
	cadena += digitoVerificador

	// 4. Convert to Base16 (hexadecimal as number, not ASCII)
	bigInt := new(big.Int)
	if _, ok := bigInt.SetString(cadena, 10); !ok {
		return "", fmt.Errorf("error converting to BigInt")
	}
	cadenaHex := strings.ToUpper(bigInt.Text(16))

	// 5. Concatenate control code (CUFD)
	cuf := cadenaHex + codigoControl

	return cuf, nil
}

// calculaDigitoMod11 calculates the Mod 11 verify digit
func calculaDigitoMod11(cadena string, numDig, limMult int, x10 bool) string {
	var mult, suma, dig int
	if !x10 {
		numDig = 1
	}
	for n := 1; n <= numDig; n++ {
		suma = 0
		mult = 2
		for i := len(cadena) - 1; i >= 0; i-- {
			num := int(cadena[i] - '0')
			suma += mult * num
			mult++
			if mult > limMult {
				mult = 2
			}
		}
		if x10 {
			dig = ((suma * 10) % 11) % 10
		} else {
			dig = suma % 11
		}
		switch dig {
		case 10:
			cadena += "1"
		case 11:
			cadena += "0"
		default:
			cadena += fmt.Sprintf("%d", dig)
		}
	}
	return cadena[len(cadena)-numDig:]
}
