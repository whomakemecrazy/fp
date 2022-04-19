package common

import (
	"reflect"
)

type Functor[T any] func(T) (T, error)
type Filter[T any] func(T) bool
type TransformerFunctor[T any, R any] func(T) R

// func Functors[T any](fn Functor[T]) func([]T) ([]T, error) {
// 	return func(data []T) ([]T, error) {
// 		ch := make(chan Result[T], len(data))

// 		var wg sync.WaitGroup
// 		wg.Add(1)

// 		result := make([]T, 0, len(data))
// 		go func(ch chan Result[T], result []T, wg *sync.WaitGroup) {
// 			defer wg.Done()
// 			for v := range ch {
// 				result = append(result, v.V)
// 			}
// 		}(ch, result, &wg)

// 		for _, v := range data {
// 			go func(v T) {
// 				ch <- Just(fn(v))
// 			}(v)
// 		}

// 		close(ch)
// 		wg.Wait()
// 		return result, nil
// 	}
// }

func Merge[T comparable, R any](base map[T]R) func(map[T]R) map[T]R {
	return func(map2 map[T]R) map[T]R {
		newMap := make(map[T]R, len(base)+len(map2))
		if base != nil {
			for k, v := range base {
				newMap[k] = v
			}
		}
		if map2 != nil {
			for k, v := range map2 {
				newMap[k] = v
			}
		}
		return newMap
	}
}

func Concat[T any](mine []T, slices ...[]T) []T {
	var mineLen = len(mine)
	var totalLen = mineLen

	for _, slice := range slices {
		if slice == nil {
			continue
		}

		var targetLen = len(slice)
		totalLen += targetLen
	}
	var newOne = make([]T, totalLen)

	for i, item := range mine {
		newOne[i] = item
	}
	totalIndex := mineLen

	for _, slice := range slices {
		if slice == nil {
			continue
		}

		var target = slice
		var targetLen = len(target)
		for j, item := range target {
			newOne[totalIndex+j] = item
		}
		totalIndex += targetLen
	}

	return newOne
}

func IsNil(obj interface{}) bool {
	val := reflect.ValueOf(obj)

	if Kind(obj) == reflect.Ptr {
		return val.IsNil()
	}
	return !val.IsValid()
}

// IsPtr Check is it a Ptr
func IsPtr(obj interface{}) bool {
	return Kind(obj) == reflect.Ptr
}

// Kind Get Kind by reflection
func Kind(obj interface{}) reflect.Kind {
	return reflect.ValueOf(obj).Kind()
}

// type MaxType[T any] interface {
// 	Compare(T) bool
// }

// func Max[T any](m ...MaxType[T]) T {

// }
