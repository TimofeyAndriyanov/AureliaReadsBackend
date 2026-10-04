package usecases

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_in"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type SignInUseCase struct {
	userRepository userRepo.Repository
	hashService    services.HashService
}

func NewSignInUseCase(
	userRepository userRepo.Repository,
	hashService services.HashService,
) *SignInUseCase {
	return &SignInUseCase{
		userRepository: userRepository,
		hashService:    hashService,
	}
}

func (uc *SignInUseCase) Execute(value entities.SignInForm) sign_in.SignInResult {
	if value.IsEmpty() {
		return sign_in.EmptyFields{}
	}

	//credential, ok := uc.userRepository.FindUserCredentialsByUsername(value.Username)

	//if !ok && credential == nil {
	//	return sign_in.UserNotFound{}
	//}

	//check := uc.hashService.HashChecking(credential.HashPassword, value.Password)

	//if !check {
	//	return sign_in.WrongPassword{}
	//}

	return sign_in.Success{Data: nil}
}
