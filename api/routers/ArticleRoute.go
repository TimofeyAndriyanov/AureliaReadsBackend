package routers

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository"
)

type ArticleHandler struct {
	repository repository.ArticlesRepository
}

func NewArticleHandler(
	repository repository.ArticlesRepository,
) *ArticleHandler {
	return &ArticleHandler{
		repository: repository,
	}
}

func (h ArticleHandler) ArticleRouter() {

}

func (h ArticleHandler) newArticle() {

}

func (h ArticleHandler) getArticleById() {

}

func (h ArticleHandler) deleteArticleById() {

}

func (h ArticleHandler) allArticles() {

}
