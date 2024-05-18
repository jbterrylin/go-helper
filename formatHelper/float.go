package formathelper

import "math"

// RoundFloat 函数将浮点数四舍五入到指定小数位数
func RoundFloat[T ConvertableToFloat](value T, decimalPlaces int) T {
	pow := math.Pow(10, float64(decimalPlaces))
	return T(math.Round(float64(value)*pow) / pow)
}

// CeilFloat 函数将浮点数向上舍入到指定小数位数
func CeilFloat[T ConvertableToFloat](value T, decimalPlaces int) T {
	pow := math.Pow(10, float64(decimalPlaces))
	return T(math.Ceil(float64(value)*pow) / pow)
}

// FloorFloat 函数将浮点数向下舍入到指定小数位数
func FloorFloat[T ConvertableToFloat](value T, decimalPlaces int) T {
	pow := math.Pow(10, float64(decimalPlaces))
	return T(math.Floor(float64(value)*pow) / pow)
}

// TruncateFloat 函数将浮点数截断到指定小数位数
func TruncateFloat[T ConvertableToFloat](value T, decimalPlaces int) T {
	pow := math.Pow(10, float64(decimalPlaces))
	return T(math.Trunc(float64(value)*pow) / pow)
}
