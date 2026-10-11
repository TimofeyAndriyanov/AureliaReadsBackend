package usecases

import (
	"database/sql"
	"errors"
	"log"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_in"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type SignInUseCase struct {
	userRepository    userRepo.Repository
	sessionRepository userRepo.SessionRepository
	hashService       services.HashService
	tokenHashService  services.HashService
}

func NewSignInUseCase(
	userRepository userRepo.Repository,
	sessionRepository userRepo.SessionRepository,
	hashService services.HashService,
	tokenHashService services.HashService,
) *SignInUseCase {
	return &SignInUseCase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		hashService:       hashService,
		tokenHashService:  tokenHashService,
	}
}

func (uc *SignInUseCase) Execute(value entities.SignInForm) sign_in.SignInResult {
	if value.IsEmpty() {
		return sign_in.EmptyFields{}
	}

	credential, credentialErr := uc.userRepository.FindUserCredentialsByUsername(value.Username)

	if errors.Is(credentialErr, sql.ErrNoRows) {
		return sign_in.UserNotFound{}
	}

	ok := uc.hashService.HashChecking(credential.HashPassword, value.Password)

	if !ok {
		return sign_in.WrongPassword{}
	}
	sessionToken, _ := entities.NewSessionToken()

	tokenHash, _ := uc.tokenHashService.Hashing(string(sessionToken))

	sessionErr := uc.sessionRepository.Add(credential.Id, value.DeviceName, entities.TokenHash([]byte(tokenHash)))

	if sessionErr != nil {
		log.Fatal(sessionErr)
		return sign_in.Unknown{}
	}

	return sign_in.Success{Data: nil}
}
