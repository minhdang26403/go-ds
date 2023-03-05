package sorting

import "golang.org/x/exp/constraints"

func merge[T constraints.Ordered](list []T, left, mid, right int) {
	leftSize := mid + 1 - left
	rightSize := right - mid
	left_list := make([]T, leftSize)
	for i := 0; i < leftSize; i++ {
		left_list[i] = list[left + i]
	}
	right_list := make([]T, rightSize)
	for j := 0; j < rightSize; j++ {
		right_list[j] = list[mid + 1 + j]
	}
	i, j, k := 0, 0, left
	for ; i < leftSize && j < rightSize; k++ {
		if left_list[i] < right_list[j] {
			list[k] = left_list[i]
			i++
		} else {
			list[k] = right_list[j]
			j++
		}
	}

	for ; i < leftSize; k, i = k + 1, i + 1 {
		list[k] = left_list[i]
	}

	for ; j < rightSize; k, j = k + 1, j + 1 {
		list[k] = right_list[j]
	}
}

func mergeSort[T constraints.Ordered](list []T, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right - left) / 2
	mergeSort(list, left, mid)
	mergeSort(list, mid + 1, right)
	merge(list, left, mid, right)
}

func MergeSort[T constraints.Ordered](list []T) {
	mergeSort(list, 0, len(list) - 1)
}