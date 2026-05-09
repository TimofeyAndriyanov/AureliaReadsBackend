package utils

import "slices"

func Find[T interface{ ~[]E }, E any](slice T, f func(E) bool) (*E, bool) {
	index := slices.IndexFunc(slice, f)

	if index == -1 {
		return nil, false
	}

	return &slice[index], true
}

func Filter[T ~[]E, E any](slice T, predicate func(E) bool) T {
	result := make(T, 0)

	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
