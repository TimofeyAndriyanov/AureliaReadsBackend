package api

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases"
)

func MainRoute(
	signUpUseCase usecases.SignUpUseCase,
	signInUseCase usecases.SignInUseCase,
	repository repository.ArticlesRepository,
) {

}
