package routers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/dto"

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
	mux.HandleFunc("articles", h.newArticle)
	mux.HandleFunc("articles", h.allArticles)

	mux.HandleFunc("articles/{id}", h.getArticleById)
	mux.HandleFunc("articles/{id}", h.deleteArticleById)
}

func (h ArticleHandler) newArticle(w http.ResponseWriter, r *http.Request) {
	var dest dto.NewArticleDTO

	if jsonErr := json.NewDecoder(r.Body).Decode(&dest); jsonErr != nil {
		h.logger.Error("Ошибка сериализации тела запроса.", jsonErr)
	}

}

func (h ArticleHandler) getArticleById(_ http.ResponseWriter, _ *http.Request) {

}

func (h ArticleHandler) deleteArticleById(_ http.ResponseWriter, _ *http.Request) {

}

func (h ArticleHandler) allArticles(_ http.ResponseWriter, _ *http.Request) {

}
