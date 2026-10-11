package utils

import "net/http"

func Map[T any, R any](items []T, mapper func(T) R) []R {
	result := make([]R, len(items))

	for i, item := range items {
		result[i] = mapper(item)
	}

	return result
}

func Chain(
	handler http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
