package article

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type Repository interface {
	AddArticle(uid entities.UserID, article entities.NewArticle) error
	GetArticleById(id entities.ArticleID) (*entities.Article, error)
	GetArticlesByUid(uid entities.UserID) ([]entities.Article, error)
	AllArticles() ([]entities.Article, error)
	DeleteArticleById(id entities.ArticleID, uid entities.UserID) error
}
