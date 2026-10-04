package routers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/TimofeyAndriyanov/AureliaReadsBackend/api/dto"
	articleRepo "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"
)

type ArticleHandler struct {
	repository articleRepo.Repository
	logger     *slog.Logger
}

func NewArticleHandler(
	repository articleRepo.Repository,
) *ArticleHandler {
	return &ArticleHandler{
		repository: repository,
	}
}

func (h ArticleHandler) ArticleRouter() {
	http.HandleFunc("articles", h.newArticle)
	http.HandleFunc("articles", h.allArticles)

	http.HandleFunc("articles/{id}", h.getArticleById)
	http.HandleFunc("articles/{id}", h.deleteArticleById)
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
