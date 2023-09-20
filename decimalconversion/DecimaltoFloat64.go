package decimalConversion

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func Float64(d []decimal.Decimal) {

	b := make([]float64, len(d))

	for i, val := range d {

		b[i], _ = val.Float64()
	}
	fmt.Println(b)
}
