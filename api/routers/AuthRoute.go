package routers

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases"
)

type AuthHandler struct {
	signInUseCase usecases.SignInUseCase
	signUpUseCase usecases.SignUpUseCase
}

func NewAuthHandler(
	signInUseCase usecases.SignInUseCase,
	signUpUseCase usecases.SignUpUseCase,
) *AuthHandler {
	return &AuthHandler{
		signInUseCase: signInUseCase,
		signUpUseCase: signUpUseCase,
	}
}

func (h AuthHandler) AuthRouter() {
}

func (h AuthHandler) signInRoute() {

}

func (h AuthHandler) signUpRoute() {

}
