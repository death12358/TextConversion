package float64conversion

import "fmt"

func Decimal(f []float64) {

	for _, k := range f {
		fmt.Printf("decimal.NewFromFloat(%v)\n", k)
	}
}
