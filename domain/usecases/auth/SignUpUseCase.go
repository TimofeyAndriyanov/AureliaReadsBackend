package usecases

import (
	"log"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_up"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type SignUpUseCase struct {
	userRepository    userRepo.Repository
	sessionRepository userRepo.SessionRepository
	hashService       services.HashService
	tokenHashService  services.HashService
}

func NewSignUpUseCase(
	userRepository userRepo.Repository,
	sessionRepository userRepo.SessionRepository,
	hashService services.HashService,
	tokenHashService services.HashService,
) *SignUpUseCase {
	return &SignUpUseCase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		hashService:       hashService,
		tokenHashService:  tokenHashService,
	}
}

func (uc *SignUpUseCase) Execute(value entities.SignUpForm) sign_up.SignUpResult {
	if value.IsEmpty() {
		return sign_up.EmptyFields{}
	}

	credential, credentialErr := uc.userRepository.FindUserCredentialsByUsername(value.Username)

	if credentialErr != nil {
		return sign_up.Unknown{}
	}

	log.Printf("SignUpUseCase FindUserCredentialsByUsername:", credential)

	if credential != nil {
		log.Printf("SignUpUseCase FindUserCredentialsByUsername:", credentialErr)
		return sign_up.UserAlreadyExists{}
	}

	hashPassword, hashPasswordErr := uc.hashService.Hashing(value.Password)

	if hashPasswordErr != nil {
		log.Fatal(hashPasswordErr)
		return sign_up.Unknown{}
	}

	newUser, newUserErr := uc.userRepository.AddUser(value.CopyPass(hashPassword))

	if newUserErr != nil {
		log.Printf("SignUpUseCase AddUser:", newUserErr)
		log.Fatal(newUserErr)
		return sign_up.Unknown{}
	}

	sessionToken, _ := entities.NewSessionToken()

	tokenHash, _ := uc.tokenHashService.Hashing(string(sessionToken))

	sessionErr := uc.sessionRepository.Add(*newUser, value.DeviceName, entities.TokenHash([]byte(tokenHash)))

	if sessionErr != nil {
		log.Fatal(sessionErr)
		return sign_up.Unknown{}
	}

	return sign_up.Success{Data: sessionToken}
}
