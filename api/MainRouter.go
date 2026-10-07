package api

import (
	articleRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"
	authUseCases "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases/auth"
)

func MainRoute(
	signUpUseCase *authUseCases.SignUpUseCase,
	signInUseCase *authUseCases.SignInUseCase,
	repository articleRepo.Repository,
) {

}
