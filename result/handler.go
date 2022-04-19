package result

type Functor[T any] func(T) (T, error)
type Filter[T any] func(T) bool
