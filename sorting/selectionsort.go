package sorting

import "golang.org/x/exp/constraints"

func SelectionSort[T constraints.Ordered](list []T) {
	n := len(list)
	for i := 0; i < n; i++ {
		pos := i
		for j := i + 1; j < n; j++ {
			if list[j] < list[pos] {
				pos = j
			}
		}
		list[i], list[pos] = list[pos], list[i]
	}
}