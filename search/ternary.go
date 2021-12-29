package search

import "math"

func ternary(a, b, c float64, f func(n float64) float64) float64 {
	if math.Abs(a-b) > 1e-10 {
		p1 := a - (a-b)/2.
		if f(p1) < f(b) {
			return ternary(p1, p1-(p1-c)/2, c, f)
		}
	}
	if math.Abs(b-c) > 1e-10 {
		p1 := b - (b-c)/2.
		if f(p1) < f(b) {
			return ternary(a, a-(a-c)/2., p1, f)
		}
	}
	return b
}
