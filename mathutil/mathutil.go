package mathutil

import (
	"github.com/shopspring/decimal"
	"math"
)

// Round 返回浮点数的四舍五入后的整数
func Round(f float64) int {
	return int(math.Round(f))
}

// RoundDown 浮点数指定小数点位数向下截取
func RoundDown(f float64, places int32) float64 {
	d := decimal.NewFromFloat(f)
	return d.RoundDown(places).InexactFloat64()
}

// RoundUp 浮点数指定小数点位数向上截取
func RoundUp(f float64, places int32) float64 {
	d := decimal.NewFromFloat(f)
	return d.RoundUp(places).InexactFloat64()
}

// Sum 浮点数精确求和
func Sum(f1, f2 float64) float64 {
	return decimal.Sum(decimal.NewFromFloat(f1), decimal.NewFromFloat(f2)).InexactFloat64()
}

func Sub(f1, f2 float64) float64 {
	return decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).InexactFloat64()
}

func Mul(f1, f2 float64) float64 {
	return decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).InexactFloat64()
}
