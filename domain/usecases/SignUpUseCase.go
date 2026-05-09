package usecases

import (
	"AureliaReadsBackend/domain/entities"
	"AureliaReadsBackend/domain/repository"
	"AureliaReadsBackend/domain/results/sign_up"
	"AureliaReadsBackend/domain/services"
)

type SignUpUseCase struct {
	userRepository repository.UserRepository
	hashService    services.HashService
	jwtService     services.JwtService
}

func NewSignUpUseCase(
	userRepository repository.UserRepository,
	hashService services.HashService,
	jwtService services.JwtService,
) SignUpUseCase {
	return SignUpUseCase{
		userRepository: userRepository,
		hashService:    hashService,
		jwtService:     jwtService,
	}
}

func (uc *SignUpUseCase) Execute(value entities.SignUpForm) sign_up.SignUpResult {
	if value.IsEmpty() {
		return sign_up.EmptyFields{}
	}

	_, ok := uc.userRepository.FindUserCredentialsByUsername(value.Username)

	if ok {
		return sign_up.UserAlreadyExists{}
	}

	hashPassword := uc.hashService.Hashing(value.Password)

	newUser := uc.userRepository.AddUser(value.CopyPass(hashPassword))

	tokens, err := uc.jwtService.NewJwt(newUser)

	if err != nil {
		return nil
	}

	return sign_up.Success{Data: *tokens}
}
