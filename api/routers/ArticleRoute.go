package routers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/dto"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/middleware"

	usecases "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/usecases/article"
)

type ArticleHandler struct {
	createArticle *usecases.CreateArticleUseCase
	allArticle    *usecases.GetAllArticlesUseCase
	findByArticle *usecases.FindArticleByIdUseCase
	deleteArticle *usecases.DeleteArticleUseCase
	logger        *slog.Logger
}

func NewArticleHandler(
	createArticle *usecases.CreateArticleUseCase,
	allArticle *usecases.GetAllArticlesUseCase,
	findByArticle *usecases.FindArticleByIdUseCase,
	deleteArticle *usecases.DeleteArticleUseCase,
	logger *slog.Logger,
) *ArticleHandler {
	return &ArticleHandler{
		createArticle: createArticle,
		allArticle:    allArticle,
		findByArticle: findByArticle,
		deleteArticle: deleteArticle,
		logger:        logger,
	}
}

func (h ArticleHandler) ArticleRouter(mux *http.ServeMux) {
	mux.Handle(
		"POST /api/articles",
		middleware.AuthMiddleware(http.HandlerFunc(h.newArticle)),
	)

	mux.Handle(
		"PATCH /api/articles/{id}",
		middleware.AuthMiddleware(http.HandlerFunc(h.patch)),
	)

	mux.Handle(
		"DELETE /api/articles/{id}",
		middleware.AuthMiddleware(http.HandlerFunc(h.deleteArticleById)),
	)

	mux.HandleFunc("GET api/articles", h.allArticles)
	mux.HandleFunc("GET api/articles/{id}", h.getArticleById)
}

func (h ArticleHandler) newArticle(w http.ResponseWriter, r *http.Request) {
	var dest dto.NewArticleDTO

	if jsonErr := json.NewDecoder(r.Body).Decode(&dest); jsonErr != nil {
		h.logger.Error("Ошибка сериализации тела запроса.", jsonErr)
	}

}

func (h ArticleHandler) allArticles(_ http.ResponseWriter, _ *http.Request) {

}

func (h ArticleHandler) getArticleById(_ http.ResponseWriter, _ *http.Request) {

}

func (h ArticleHandler) patch(resp http.ResponseWriter, r *http.Request) {
	var dest dto.StatusFormDTO

	if jsonErr := json.NewDecoder(r.Body).Decode(&dest); jsonErr != nil {
		h.logger.Error("Ошибка сериализации тела запроса.", jsonErr)
		resp.Write(nil)
	}
}

func (h ArticleHandler) deleteArticleById(_ http.ResponseWriter, _ *http.Request) {

}
