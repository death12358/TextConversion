package decimalConversion

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func Float64(initValues []decimal.Decimal) {

	b := make([]float64, len(initValues))

	for i, val := range initValues {

		b[i], _ = val.Float64()
	}
	fmt.Println(b)
}
