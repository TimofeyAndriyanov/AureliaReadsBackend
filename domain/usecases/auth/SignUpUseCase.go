package usecases

import (
	"database/sql"
	"errors"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_up"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type SignUpUseCase struct {
	userRepository    userRepo.Repository
	sessionRepository userRepo.SessionRepository
	hashService       services.HashService
}

func NewSignUpUseCase(
	userRepository userRepo.Repository,
	sessionRepository userRepo.SessionRepository,
	hashService services.HashService,
) *SignUpUseCase {
	return &SignUpUseCase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		hashService:       hashService,
	}
}

func (uc *SignUpUseCase) Execute(value entities.SignUpForm) sign_up.SignUpResult {
	if value.IsEmpty() {
		return sign_up.EmptyFields{}
	}

	credential, credentialErr := uc.userRepository.FindUserCredentialsByUsername(value.Username)

	if !errors.Is(credentialErr, sql.ErrNoRows) && credential != nil {
		return sign_up.UserAlreadyExists{}
	}

	hashPassword, hashPasswordErr := uc.hashService.Hashing(value.Password)

	if hashPasswordErr != nil {
		return sign_up.Unknown{}
	}

	newUser, newUserErr := uc.userRepository.AddUser(value.CopyPass(hashPassword))

	if newUserErr != nil {
		return sign_up.Unknown{}
	}

	token, tokenErr := uc.sessionRepository.Add(*newUser, value.DeviceName)

	if tokenErr != nil {
		return sign_up.Unknown{}
	}

	return sign_up.Success{Data: token}
}
