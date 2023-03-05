package sorting

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](list []T) {
	n := len(list)
	swapped := true
	for i := 0; i < n - 1 && swapped; i++ {
		swapped = false
		for j := 0; j < n - 1 - i; j++ {
			if list[j] > list[j + 1] {
				list[j], list[j + 1] = list[j + 1], list[j]
				swapped = true
			}
		}
	}
}