package articles

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository"
)

type repositoryImpl struct{}

func NewArticlesRepository() repository.ArticlesRepository {
	return &repositoryImpl{}
}

func (i repositoryImpl) AddArticle(uid entities.UserID, article entities.NewArticle) {
	panic("")
}

func (i repositoryImpl) GetArticleById(id entities.ArticleID) (*entities.Article, bool) {
	panic("")
}

func (i repositoryImpl) DeleteArticleById(id entities.ArticleID, uid entities.UserID) {
	panic("")
}

func (i repositoryImpl) GetArticlesByUid(uid entities.UserID) []entities.Article {
	panic("")
}

func (i repositoryImpl) AllArticles() []entities.Article {
	panic("")
}
