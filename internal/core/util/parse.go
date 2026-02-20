package util

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseIntSafe convierte un string a int eliminando espacios en blanco.
// Retorna un error descriptivo si la conversión falla.
func ParseIntSafe(valStr string) (int, error) {
	cleanVal := strings.TrimSpace(valStr)
	val, err := strconv.Atoi(cleanVal)
	if err != nil {
		return 0, fmt.Errorf("la variable debe ser un número entero válido: %w", err)
	}
	return val, nil
}

// ParseInt64Safe convierte un string a int64 (ideal para el NIT).
// Retorna un error descriptivo si la conversión falla.
func ParseInt64Safe(valStr string) (int64, error) {
	cleanVal := strings.TrimSpace(valStr)
	val, err := strconv.ParseInt(cleanVal, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("la variable debe ser un número (int64) válido: %w", err)
	}
	return val, nil
}
