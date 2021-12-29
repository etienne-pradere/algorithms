package search

import "math"

func binary(k, s, dato int, f func(a int) int) int {
	if k == s {
		return k
	}
	p := int(math.Ceil(float64((k + 1.) / 2.)))
	if dato > f(p) {
		return binary(p, s, dato, f)
	}
	return binary(k, p-1, dato, f)
}
