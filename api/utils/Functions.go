package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ParamUUID(c *gin.Context, key string) (uuid.UUID, error) {
	return uuid.Parse(c.Param(key))
}

func Map[T any, R any](items []T, mapper func(T) R) []R {
	result := make([]R, len(items))

	for i, item := range items {
		result[i] = mapper(item)
	}

	return result
}
