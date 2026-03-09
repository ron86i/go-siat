package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// GenerarCUF genera el CUF completo para una factura
func GenerarCUF(nit int64, fechaHora time.Time, sucursal int, modalidad, tipoEmision, tipoFactura int, tipoDocumentoSector int, numeroFactura int, puntoVenta int, codigoControl string) (string, error) {

	// 1. Formatear campos a longitud fija
	nitStr := fmt.Sprintf("%013d", nit)                                // 13
	fechaStr := fechaHora.Format("20060102150405.000")                 // 17 yyyyMMddHHmmss.SSS
	fechaStr = strings.ReplaceAll(fechaStr, ".", "")                   // 17 dígitos
	sucursalStr := fmt.Sprintf("%04d", sucursal)                       // 4
	modalidadStr := fmt.Sprintf("%1d", modalidad)                      // 1
	tipoEmisionStr := fmt.Sprintf("%1d", tipoEmision)                  // 1
	tipoFacturaStr := fmt.Sprintf("%1d", tipoFactura)                  // 1
	tipoDocumentoSectorStr := fmt.Sprintf("%02d", tipoDocumentoSector) // 2
	numeroFacturaStr := fmt.Sprintf("%010d", numeroFactura)            // 10
	puntoVentaStr := fmt.Sprintf("%04d", puntoVenta)                   // 4

	// 2. Concatenar todo hasta punto de venta
	cadena := nitStr + fechaStr + sucursalStr + modalidadStr + tipoEmisionStr +
		tipoFacturaStr + tipoDocumentoSectorStr + numeroFacturaStr + puntoVentaStr

	// 3. Calcular dígito verificador Módulo 11
	digitoVerificador := calculaDigitoMod11(cadena, 1, 9, false)
	cadena += digitoVerificador

	// 4. Convertir a Base16 (hexadecimal como número, no ASCII)
	bigInt := new(big.Int)
	bigInt.SetString(cadena, 10)
	cadenaHex := strings.ToUpper(bigInt.Text(16))

	// 5. Concatenar código de control (CUFD)
	cuf := cadenaHex + codigoControl

	return cuf, nil
}

// CalculaDigitoMod11 calcula el dígito verificador Mod 11
func calculaDigitoMod11(cadena string, numDig, limMult int, x10 bool) string {
	var mult, suma, dig int
	if !x10 {
		numDig = 1
	}
	for n := 1; n <= numDig; n++ {
		suma = 0
		mult = 2
		for i := len(cadena) - 1; i >= 0; i-- {
			num, _ := strconv.Atoi(string(cadena[i]))
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
