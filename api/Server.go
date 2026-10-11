package api

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/dto"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/middleware"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/routers"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/utils"
)

func NewServer(host, port string, auth *routers.AuthHandler, article *routers.ArticleHandler) {
	addr := net.JoinHostPort(host, port)

	mux := http.NewServeMux()

	middlewares := utils.Chain(
		mux,
		middleware.BaseMiddleware,
		middleware.LoggingMiddleware,
	)

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		msg := dto.MessageDTO{
			Message: "Привет, как дела?",
		}

		if err := json.NewEncoder(w).Encode(msg); err != nil {
			return
		}
	})

	auth.AuthRouter(mux)
	article.ArticleRouter(mux)

	log.Printf("Сервер успешно запущен на %s", addr)

	server := &http.Server{
		Addr:        addr,
		Handler:     middlewares,
		ReadTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
