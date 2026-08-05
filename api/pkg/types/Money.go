package types

import "math"

type Money int64

func FloatToMoney(value float64) Money {
	return Money(math.Round(value * 100))
}

func (m Money) ToFloat() float64 {
	return float64(m) / 100
}
