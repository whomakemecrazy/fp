package stream

type Stream[T any] struct {
	data []T
	err  error
}

func (s Stream[T]) Contains() Stream[T] {
	return s
}
func (s Stream[T]) Filter() Stream[T] {
	return s
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
