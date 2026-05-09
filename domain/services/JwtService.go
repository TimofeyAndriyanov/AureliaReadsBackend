package services

import (
	"AureliaReadsBackend/domain/entities"
	"AureliaReadsBackend/domain/utils"
)

type JwtService interface {
	NewJwt(uid entities.UserID) (*entities.JwtTokens, error)
	VerifyToken(token string, typeToken utils.TokenType) (*entities.JwtPayload, error)
}
