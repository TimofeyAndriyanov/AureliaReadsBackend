package utils

import (
	"context"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
)

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

type userIDKey struct{}

func WithUserID(ctx context.Context, id entities.UserID) context.Context {
	return context.WithValue(ctx, userIDKey{}, id)
}

func UserID(ctx context.Context) (entities.UserID, bool) {
	userID, ok := ctx.Value(userIDKey{}).(entities.UserID)
	return userID, ok
}
