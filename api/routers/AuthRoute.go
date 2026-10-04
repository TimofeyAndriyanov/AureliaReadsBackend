package routers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/dto"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/log"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases"
)

type AuthHandler struct {
	signInUseCase *usecases.SignInUseCase
	signUpUseCase *usecases.SignUpUseCase
	repository    log.Repository
	logger        *slog.Logger
}

func NewAuthHandler(
	signInUseCase *usecases.SignInUseCase,
	signUpUseCase *usecases.SignUpUseCase,
	repository log.Repository,
	logger *slog.Logger,
) *AuthHandler {
	return &AuthHandler{
		signInUseCase: signInUseCase,
		signUpUseCase: signUpUseCase,
		logger:        logger,
	}
}

func (h AuthHandler) AuthRouter() {
	http.HandleFunc("sign_in", h.signInRoute)
	http.HandleFunc("sign_up", h.signUpRoute)
}

func (h AuthHandler) signInRoute(_ http.ResponseWriter, r *http.Request) {
	var dest dto.SignInFormDTO

	if jsonErr := json.NewDecoder(r.Body).Decode(&dest); jsonErr != nil {
		h.repository.Add(entities.ERROR, "")
		h.logger.Error("Ошибка сериализации тела запроса.", jsonErr)
	}
}

func (h AuthHandler) signUpRoute(_ http.ResponseWriter, r *http.Request) {
	var dest dto.SignUpFormDTO

	if jsonErr := json.NewDecoder(r.Body).Decode(&dest); jsonErr != nil {
		h.logger.Error("Ошибка сериализации тела запроса.", jsonErr)
	}
}
