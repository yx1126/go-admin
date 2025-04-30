package util

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, cb func(item T, index int) bool) (result []T) {
	for i, v := range slice {
		if cb(v, i) {
			result = append(result, v)
		}
	}
	return
}

func Map[T any, R any](slice []T, cb func(item T, index int) R) (result []R) {
	for i, v := range slice {
		result = append(result, cb(v, i))
	}
	return
}

func Reduce[T any, R any](slice []T, cb func(acc R, item T, index int) R, initial R) (result R) {
	result = initial
	for i, v := range slice {
		result = cb(result, v, i)
	}
	return
}
