package common

import (
	"golang.org/x/exp/constraints"
)

func Contain[T comparable](data []T, d T) bool {
	for _, d2 := range data {
		if d == d2 {
			return true
		}
	}
	return false
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Negative[T Number](data T) bool {
	return data < 0
}
