package api

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/routers"
)

func NewServer(host, port string, auth *routers.AuthHandler, article *routers.ArticleHandler) {
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, lisErr := net.Listen("tcp", addr)

	if lisErr != nil {
		log.Fatal(lisErr)
	}

	mux := http.NewServeMux()

	auth.AuthRouter(mux)
	article.ArticleRouter(mux)

	if serveErr := http.Serve(lis, mux); serveErr != nil {
		log.Fatal(serveErr)
	}

	log.Printf("Сервер успешно запущен на %s", addr)
}
