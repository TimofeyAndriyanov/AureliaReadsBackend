package api

import (
	articleRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases"
)

func MainRoute(
	signUpUseCase *usecases.SignUpUseCase,
	signInUseCase *usecases.SignInUseCase,
	repository articleRepo.Repository,
) {

}
