package float64conversion

import "fmt"

func Decimal() {
	a := []float64{2, 3, 5, 7, 10, 20, 30, 50, 60, 70, 80, 90, 100, 120, 140, 160, 200, 300, 500, 1000}
	for _, k := range a {
		fmt.Println("decimal.NewFromFloat(", k, "),")
	}
}
