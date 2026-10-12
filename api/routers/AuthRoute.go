package routers

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/dto"
	signInResult "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_in"
	signUpResult "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/sign_up"

	usecases "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases/auth"
)

type AuthHandler struct {
	signInUseCase *usecases.SignInUseCase
	signUpUseCase *usecases.SignUpUseCase
	logger        *slog.Logger
}

func NewAuthHandler(
	signInUseCase *usecases.SignInUseCase,
	signUpUseCase *usecases.SignUpUseCase,
	logger *slog.Logger,
) *AuthHandler {
	return &AuthHandler{
		signInUseCase: signInUseCase,
		signUpUseCase: signUpUseCase,
		logger:        logger,
	}
}

func (h AuthHandler) AuthRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/sign_in", h.signInRoute)
	mux.HandleFunc("POST /api/sign_up", h.signUpRoute)
}

func (h AuthHandler) signInRoute(w http.ResponseWriter, r *http.Request) {
	var dest dto.SignInFormDTO

	if jsonErr := json.NewDecoder(r.Body).Decode(&dest); jsonErr != nil {
		h.logger.Error("Ошибка сериализации тела запроса.", jsonErr)
		w.Write(nil)
	}

	switch result := h.signInUseCase.Execute(dest.ToDomain()).(type) {
	case signInResult.Success:
		token := dto.ToDTO(result.Data)
		log.Printf("signInRoute Token %v", token)
		if err := json.NewEncoder(w).Encode(token); err != nil {
			return
		}

	case signInResult.WrongPassword:
		msg := dto.MessageDTO{
			Message: "wrong_password",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}

	case signInResult.EmptyFields:
		msg := dto.MessageDTO{
			Message: "emply_fields",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}

	case signInResult.UserNotFound:
		msg := dto.MessageDTO{
			Message: "user_not_found",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}

	case signInResult.Unknown:
		msg := dto.MessageDTO{
			Message: "unknown",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}

	}
}

func (h AuthHandler) signUpRoute(w http.ResponseWriter, r *http.Request) {
	var dest dto.SignUpFormDTO

	if jsonErr := json.NewDecoder(r.Body).Decode(&dest); jsonErr != nil {
		h.logger.Error("Ошибка сериализации тела запроса: %s", jsonErr)
		return
	}

	switch result := h.signUpUseCase.Execute(dest.ToDomain()).(type) {
	case signUpResult.Success:
		token := dto.ToDTO(result.Data)
		log.Printf("signUpRoute Token %v", token)
		if err := json.NewEncoder(w).Encode(token); err != nil {
			return
		}

	case signUpResult.UserAlreadyExists:
		msg := dto.MessageDTO{
			Message: "user_already_exists",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}

	case signUpResult.EmptyFields:
		msg := dto.MessageDTO{
			Message: "empty_fields",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}

	case signUpResult.Unknown:
		msg := dto.MessageDTO{
			Message: "unknown",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}
	}
}
