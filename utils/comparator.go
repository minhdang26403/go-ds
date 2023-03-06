package utils

import "golang.org/x/exp/constraints"

type Comparator[T constraints.Ordered] func(a, b T) bool

func Less[T constraints.Ordered](a, b T) bool {
	return a < b
}

func Greater[T constraints.Ordered](a, b T) bool {
	return a > b
}
