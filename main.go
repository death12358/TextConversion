package Cfunc

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func NumTypeTrans() {
	a := []float64{2, 3, 5, 7, 10, 20, 30, 50, 60, 70, 80, 90, 100, 120, 140, 160, 200, 300, 500, 1000}
	for _, k := range a {
		fmt.Println("decimal.NewFromFloat(", k, "),")
	}

	initValues := []decimal.Decimal{
		decimal.NewFromFloat(1),
		decimal.NewFromFloat(10),
		decimal.NewFromFloat(20),
		decimal.NewFromFloat(30),
		decimal.NewFromFloat(40),
		decimal.NewFromFloat(50),
		decimal.NewFromFloat(60),
		decimal.NewFromFloat(70),
		decimal.NewFromFloat(80),
		decimal.NewFromFloat(90),
		decimal.NewFromFloat(100),
		decimal.NewFromFloat(110),
		decimal.NewFromFloat(120),
		decimal.NewFromFloat(130),
		decimal.NewFromFloat(140),
		decimal.NewFromFloat(150),
		decimal.NewFromFloat(160),
		decimal.NewFromFloat(170),
		decimal.NewFromFloat(180),
		decimal.NewFromFloat(190),
		decimal.NewFromFloat(200),
		decimal.NewFromFloat(220),
		decimal.NewFromFloat(240),
		decimal.NewFromFloat(260),
		decimal.NewFromFloat(280),
		decimal.NewFromFloat(300),
		decimal.NewFromFloat(350),
		decimal.NewFromFloat(400),
		decimal.NewFromFloat(450),
		decimal.NewFromFloat(500),
		decimal.NewFromFloat(1000),
	}

	b := make([]float64, len(initValues))

	for i, val := range initValues {

		b[i], _ = val.Float64()
	}

	fmt.Println(b)

}

func dot() {
	vector1 := []float64{1, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 220, 240, 260, 280, 300, 350, 400, 450, 500, 1000}
	vector2 := []float64{700, 570, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 20, 20, 10, 10, 10}

	dotProd, err := dotProduct(vector1, vector2)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}

	fmt.Printf("内积结果：%.2f\n", dotProd)
}

// []float64{700, 570, 570, 600, 550, 500, 400, 300, 100, 60, 60, 50, 50, 50, 20, 20, 20, 10, 10, 10},
// 	[]float64{0, 0, 0, 0, 0, 0, 500, 350, 400, 60, 60, 50, 50, 20, 20, 10, 0, 0, 0, 0}, //	21	FISHC06 財神

func dotProduct(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("切片长度不相等")
	}

	var result float64
	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}
	return result, nil
}
