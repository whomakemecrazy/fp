package common

func ComposeE[A0 any, T1 any, T2 any](
	fn0 func(A0) (T1, error),
	fn1 func(T1, error) (T2, error),
) func(A0) (T2, error) {
	return func(data A0) (T2, error) {
		return fn1(fn0(data))
	}
}

// func Compose2[A0 any, T1 any, T2 any, T3 any, T4 any](fn0 func(A0) (T1, T2), fn1 func(T1, T2) (T3, T4)) func(A0) (T3, T4) {
// 	return func(data A0) (T3, T4) {
// 		return fn1(fn0(data))
// 	}
// }

func Compose[A0 any, T1 any, T2 any](fn0 func(A0) T1, fn1 func(T1) T2) func(A0) T2 {
	return func(data A0) T2 {
		return fn1(fn0(data))
	}
}
func Map[TValue any, TResult any](iterator func(TValue) TResult) func([]TValue) []TResult {
	return func(data []TValue) []TResult {
		out := make([]TResult, len(data))
		for i, v := range data {
			out[i] = iterator(v)
		}
		return out
	}
}

func Reduce[TValue any, TAcc any](iterator func(TAcc, TValue, int) TAcc, initialAcc TAcc) func([]TValue) TAcc {
	return func(data []TValue) TAcc {
		out := initialAcc
		for i, v := range data {
			out = iterator(out, v, i)
		}
		return out
	}
}
