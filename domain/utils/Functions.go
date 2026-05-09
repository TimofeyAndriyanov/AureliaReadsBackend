package utils

import (
	"AureliaReadsBackend/domain/entities"

	"github.com/google/uuid"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func NewUserID() entities.UserID {
	return entities.UserID(uuid.New())
}

func NewArticleID() entities.ArticleID {
	return entities.ArticleID(uuid.New())
}
