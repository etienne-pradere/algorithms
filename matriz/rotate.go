package matriz

func rotateLeft(mat [][]int) [][]int {
	n := len(mat)
	var m int
	if n != 0 {
		m = len(mat[0])
	}
	res := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[i]))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[j][i] = mat[n-1-i][j]
		}
	}
	return res
}

func rotateRight(mat [][]int) [][]int {
	n := len(mat)
	var m int
	if n != 0 {
		m = len(mat[0])
	}
	res := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[i]))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[j][i] = mat[i][m-1-j]
		}
	}
	return res
}
