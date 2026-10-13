package middleware

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/dto"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/results/verify_session_token"
	usecases "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases/auth"
)

const (
	HeaderAuthorization = "Authorization"
)

type AuthMiddlewareHandler struct {
	usecase *usecases.VerifySessionTokenUseCase
}

func NewAuthMiddlewareHandler(usecase *usecases.VerifySessionTokenUseCase) *AuthMiddlewareHandler {
	return &AuthMiddlewareHandler{
		usecase: usecase,
	}
}

func (h AuthMiddlewareHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(HeaderAuthorization)

		if token == "" {
			respondUnauthorized(w)
			log.Printf("AuthMiddleware Token emply")
			return
		}

		tokenBytes, base64Err := base64.StdEncoding.DecodeString(token)

		if base64Err != nil {
			respondUnauthorized(w)
			log.Printf("AuthMiddleware base64Err: %v", base64Err)
			return
		}

		sessionToken := entities.SessionToken(tokenBytes)

		switch h.usecase.Execute(sessionToken).(type) {
		case verify_session_token.Success:
			next.ServeHTTP(w, r)

		case verify_session_token.Failed:
			respondUnauthorized(w)
			return
		}
	})
}

func respondUnauthorized(w http.ResponseWriter) {
	msg := dto.MessageDTO{
		Message: "unauthorized",
	}
	w.WriteHeader(http.StatusUnauthorized)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		log.Printf("AuthMiddleware respondUnauthorized: %v", err)
	}
}
