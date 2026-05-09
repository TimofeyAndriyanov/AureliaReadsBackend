package usecases

import (
	"AureliaReadsBackend/domain/entities"
	"AureliaReadsBackend/domain/repository"
	"AureliaReadsBackend/domain/results/sign_in"
	"AureliaReadsBackend/domain/services"
)

type SignInUseCase struct {
	userRepository repository.UserRepository
	hashService    services.HashService
	jwtService     services.JwtService
}

func NewSignInUseCase(
	userRepository repository.UserRepository,
	hashService services.HashService,
	jwtService services.JwtService,
) SignInUseCase {
	return SignInUseCase{
		userRepository: userRepository,
		hashService:    hashService,
		jwtService:     jwtService,
	}
}

func (uc *SignInUseCase) Execute(value entities.SignInForm) sign_in.SignInResult {
	if value.IsEmpty() {
		return sign_in.EmptyFields{}
	}

	credential, ok := uc.userRepository.FindUserCredentialsByUsername(value.Username)

	if !ok && credential == nil {
		return sign_in.UserNotFound{}
	}

	check := uc.hashService.HashChecking(credential.HashPassword, value.Password)

	if !check {
		return sign_in.WrongPassword{}
	}

	tokens, err := uc.jwtService.NewJwt(credential.Id)

	if err != nil {
		return nil
	}

	return sign_in.Success{Data: *tokens}
}
