package result

type Result[T any] struct {
	V T
	E error
}

func New[T any](element T) Result[T] {
	return Result[T]{V: element}
}

func Just[T any](element T, e error) Result[T] {
	return Result[T]{V: element, E: e}
}

func Transform[T any, R any](fn func(T) (R, error)) func(Result[T]) Result[R] {
	return func(r1 Result[T]) Result[R] {
		switch {
		case r1.IsFailed():
			return Result[R]{
				E: r1.E,
			}
		}
		return Just(fn(r1.V))
	}
}

func (r Result[T]) IsFailed() bool {
	return r.E != nil
}

func (r Result[T]) Map(fn Functor[T]) Result[T] {
	if r.IsFailed() {
		return r
	}
	result, err := fn(r.V)
	return Result[T]{V: result, E: err}
}

func (r Result[T]) SetError(err error) Result[T] {
	r.E = err
	return r
}

func (r Result[T]) Value() T {
	return r.V
}

func (r Result[T]) Error() error {
	return r.E
}
