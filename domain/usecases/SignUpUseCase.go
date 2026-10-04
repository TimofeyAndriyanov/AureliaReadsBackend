package usecases

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	userRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/user"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_up"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/services"
)

type SignUpUseCase struct {
	userRepository userRepo.Repository
	hashService    services.HashService
}

func NewSignUpUseCase(
	userRepository userRepo.Repository,
	hashService services.HashService,
) *SignUpUseCase {
	return &SignUpUseCase{
		userRepository: userRepository,
		hashService:    hashService,
	}
}

func (uc *SignUpUseCase) Execute(value entities.SignUpForm) sign_up.SignUpResult {
	if value.IsEmpty() {
		return sign_up.EmptyFields{}
	}

	//_, ok := uc.userRepository.FindUserCredentialsByUsername(value.Username)

	//if ok {
	//	return sign_up.UserAlreadyExists{}
	//}

	//hashPassword := uc.hashService.Hashing(value.Password)

	//newUser := uc.userRepository.AddUser(value.CopyPass(hashPassword))

	//tokens, err := uc.jwtService.NewJwt(newUser)

	//if err != nil {
	//	return nil
	//}

	return sign_up.Success{Data: nil}
}
