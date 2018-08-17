package simplemath

import "math"
func Sqrt(nData int) int {
	v := math.Sqrt(float64(nData))
	return int(v)
}

