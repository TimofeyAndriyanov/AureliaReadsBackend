package usecases

import (
	"database/sql"
	"errors"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_in"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type SignInUseCase struct {
	userRepository    userRepo.Repository
	sessionRepository userRepo.SessionRepository
	hashService       services.HashService
}

func NewSignInUseCase(
	userRepository userRepo.Repository,
	sessionRepository userRepo.SessionRepository,
	hashService services.HashService,
) *SignInUseCase {
	return &SignInUseCase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		hashService:       hashService,
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

	token, tokenErr := uc.sessionRepository.Add(credential.Id, value.DeviceName)

	if tokenErr != nil {
		return sign_in.Unknown{}
	}

	return sign_in.Success{Data: token}
}
