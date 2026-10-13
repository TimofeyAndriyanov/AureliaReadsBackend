package usecases

import (
	"encoding/base64"
	"log"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/verify_session_token"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type VerifySessionTokenUseCase struct {
	hash     services.HashService
	sessions user.SessionRepository
}

func NewVerifySessionTokenUseCase(
	hash services.HashService,
	sessions user.SessionRepository,
) *VerifySessionTokenUseCase {
	return &VerifySessionTokenUseCase{
		hash:     hash,
		sessions: sessions,
	}
}

func (uc VerifySessionTokenUseCase) Execute(token entities.SessionToken) verify_session_token.VerifySessionToken {
	hash, hashErr := uc.hash.Hashing(string(token))

	if hashErr != nil {
		log.Printf("VerifySessionTokenUseCase hashErr: %v", hashErr)
		return verify_session_token.Failed{}
	}

	tokenBytes, tokenBytesErr := base64.RawStdEncoding.DecodeString(hash)

	if tokenBytesErr != nil {
		log.Printf("VerifySessionTokenUseCase tokenBytesErr: %v", tokenBytesErr)
		return verify_session_token.Failed{}
	}

	_, err := uc.sessions.DoesSessionExist(entities.TokenHash(tokenBytes))

	if err != nil {
		log.Printf("VerifySessionTokenUseCase err: %v", err)
		return verify_session_token.Failed{}
	}

	return verify_session_token.Success{}
}
