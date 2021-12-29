package sort

func merge(arr []int, tmpArr []int, pi int, pf int) {
	if pf <= pi {
		return
	}
	m := (pi+pf)/2 + 1
	var i, j, k int
	merge(arr, tmpArr, pi, m-1)
	merge(arr, tmpArr, m, pf)
	for i, j, k = pi, m, pi; i <= m-1 && j <= pf; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= m-1; k++ {
		tmpArr[k] = arr[i]
		i++
	}

	for k = pi; k < j; k++ {
		arr[k] = tmpArr[k]
	}
}
