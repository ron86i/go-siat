package utils

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

// GenerarCUF genera el CUF completo para una factura
//
// Args:
//
//	nit (int64): Nit del contribuyente.
//	fechaHora (time.Time): Fecha y hora de emisión de la factura.
//	sucursal (int): Número de sucursal.
//	modalidad (int): Modalidad de emisión.
//	tipoEmision (int): Tipo de emisión.
//	tipoFactura (int): Tipo de factura.
//	tipoDocumentoSector (int): Tipo de documento sector.
//	numeroFactura (int): Número de la factura.
//	puntoVenta (int): Punto de venta.
//	codigoControl (string): Código de control.
//
// Returns:
//
//	(string, error): CUF completo y error si ocurre.
func GenerarCUF(nit int64, fechaHora time.Time, sucursal, modalidad, tipoEmision, tipoFactura, tipoDocumentoSector, numeroFactura, puntoVenta int, codigoControl string) (string, error) {

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
