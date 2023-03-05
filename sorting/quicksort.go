package sorting

import "golang.org/x/exp/constraints"

func partition[T constraints.Ordered](list []T, left, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

func quickSort[T constraints.Ordered](list []T, left, right int) {
	if left >= right {
		return
	}
	pivotIdx := partition(list, left, right)
	quickSort(list, left, pivotIdx-1)
	quickSort(list, pivotIdx+1, right)
}

func QuickSort[T constraints.Ordered](list []T) {
	quickSort(list, 0, len(list)-1)
}
