package utils

import (
	"fmt"
	"strconv"
)

// Round rounds a float64 value to the specified number of decimals.
func Round(v float64, decimals int) float64 {
	rounded, _ := strconv.ParseFloat(fmt.Sprintf("%.*f", decimals, v), 64)
	return rounded
}
