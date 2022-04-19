package maybe

import (
	"github.com/pkg/errors"

	"github.com/whomakemecrazy/fp/common"
)

type Maybe[T any] struct {
	V T
	E error
}

func Just[T any](element T) Maybe[T] {
	if common.IsNil(element) {
		return Maybe[T]{E: errors.New("empty element")}
	}
	return Maybe[T]{V: element}
}

func (m Maybe[T]) IsError() bool {
	return m.E != nil
}

func (m Maybe[T]) Or(element T) Maybe[T] {
	if m.IsError() {
		return Maybe[T]{V: element}
	}
	return m
}

func (m Maybe[T]) Then(m2 Maybe[T]) Maybe[T] {
	if m.IsError() {
		return m2
	}
	return m

}
func (m Maybe[T]) Error(e string) Maybe[T] {
	m.E = errors.Wrap(m.E, e)
	return m
}
func (m Maybe[T]) Map(fn func(T) T) Maybe[T] {
	if m.IsError() {
		return m
	}
	return Just(fn(m.V))
}

func Transform[T, R any](fn common.TransformerFunctor[T, R]) func(Maybe[T]) Maybe[R] {
	return func(m Maybe[T]) Maybe[R] {
		return Just(fn(m.V))
	}
}
