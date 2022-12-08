package quicksort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T constraints.Ordered](arr []T, lo int, hi int) {
	if lo < hi {
		pi := partition(arr, lo, hi)
		quickSort(arr, lo, pi-1)
		quickSort(arr, pi+1, hi)
	}
}

func partition[T constraints.Ordered](arr []T, lo int, hi int) int {
	pivot := arr[hi]
	i := lo - 1

	for j := lo; j <= hi-1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, hi)
	return i + 1
}

func swap[T constraints.Ordered](arr []T, i int, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}
