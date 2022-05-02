package stream

import (
	"github.com/whomakemecrazy/fp/common"
)

type Stream[T any] struct {
	data []T
	err  error
}

func Just[T any](args []T, e error) Stream[T] {
	return Stream[T]{
		data: args,
		err:  e,
	}
}

func (s Stream[T]) Contains(value T) bool {
	for _, v := range s.data {
		if common.Equal(v, value) {
			return true
		}
	}
	return false
}
func (s Stream[T]) Filter(fn func(T) bool) Stream[T] {
	result := make([]T, 0, len(s.data))
	for _, v := range s.data {
		if fn(v) {
			result = append(result, v)
		}
	}
	return Just(result, nil)
}

func (s Stream[T]) Distinct() Stream[T] {
	return s
}
func (s Stream[T]) IsSubset() Stream[T] {
	return s
}
func (s Stream[T]) Clone() Stream[T] {
	return s
}
func (s Stream[T]) Intersection() Stream[T] {
	return s
}
func (s Stream[T]) Remove() Stream[T] {
	return s
}
func (s Stream[T]) Append() Stream[T] {
	return s
}
func (s Stream[T]) Len() int {
	return len(s.data)
}
func (s Stream[T]) Concat() Stream[T] {
	return s
}
func (s Stream[T]) Extend() Stream[T] {
	return s
}

func (s Stream[T]) Reverse() Stream[T] {
	return s
}

func (s Stream[T]) Sort() Stream[T] {
	return s
}

type Set struct{}

type SetDef[T any] interface {
	GetKey() int64
	GetValue() T
	Equal(SetDef[T]) bool
}
