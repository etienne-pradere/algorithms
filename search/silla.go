package search

type Silla struct {
	Mat [][]int
}

// Search f es la Función desde punto dado por (a,b) y punto final dado por (c,d)
func (silla *Silla) Search(a, b, c, d, dato int, f func(a, b, c, d int) int) *[]int {
	if d == len(silla.Mat[0]) {
		return nil
	}
	if c < 0 {
		return nil
	}

	if f(a, b, c, d) == dato {
		return &[]int{c, d}
	}
	return silla.Search(a, b, c, d+1, dato, f)
}
