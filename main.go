package main

import (
	"github.com/death12358/TextConversion.git/decimalConversion"
	float64conversion "github.com/death12358/TextConversion.git/float64Conversion"
	"github.com/shopspring/decimal"
)

func main() {
	// 初始化数据
	// 初始化数据
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
	decimalConversion.Float64(initValues)
	float64conversion.Decimal([]float64{123, 456})
}
