package routers

import (
	"AureliaReadsBackend/api/dto"
	"AureliaReadsBackend/domain/results/sign_in"
	"AureliaReadsBackend/domain/results/sign_up"
	"AureliaReadsBackend/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h AuthHandler) AuthRouter(r *gin.RouterGroup) {
	r.POST("/sign_in", h.signInRoute)
	r.POST("/sign_up", h.signUpRoute)
}

func (h AuthHandler) signInRoute(c *gin.Context) {
	var signInForm dto.SignInFormDTO

	if err := c.ShouldBindJSON(&signInForm); err != nil {
		return
	}

	switch result := h.signInUseCase.Execute(signInForm.ToDomain()).(type) {
	case sign_in.Success:
		c.JSON(
			http.StatusOK,
			gin.H{
				"access":  result.Data.Access,
				"refresh": result.Data.Refresh,
			},
		)
	case sign_in.WrongPassword:
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Неправильный пароль.",
			},
		)
	case sign_in.EmptyFields:
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Поля ввода пустые.",
			},
		)
	case sign_in.UserNotFound:
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Такого пользователя не существует.",
			},
		)

	}
}

func (h AuthHandler) signUpRoute(c *gin.Context) {
	var signUpForm dto.SignUpFormDTO

	if err := c.ShouldBindJSON(&signUpForm); err != nil {
		return
	}

	switch result := h.signUpUseCase.Execute(signUpForm.ToDomain()).(type) {
	case sign_up.Success:
		c.JSON(
			http.StatusOK,
			gin.H{
				"access":  result.Data.Access,
				"refresh": result.Data.Refresh,
			},
		)
	case sign_up.UserAlreadyExists:
		c.JSON(
			http.StatusOK,
			gin.H{
				"access": "Такой пользователь уже существует.",
			},
		)
	case sign_up.EmptyFields:
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Поля ввода пустые.",
			},
		)
	}
}
