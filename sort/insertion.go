package sort

func insertion(arr []int, pi, pf int) []int {
	for i := pi + 1; i <= pf; i++ {
		j := i
		v := arr[i]
		for ; j > pi && arr[j-1] > v; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = v
	}
	return arr
}
