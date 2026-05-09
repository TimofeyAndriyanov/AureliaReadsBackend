package repository

import "AureliaReadsBackend/domain/entities"

type ArticlesRepository interface {
	AddArticle(uid entities.UserID, article entities.NewArticle)
	GetArticleById(id entities.ArticleID) (*entities.Article, bool)
	GetArticlesByUid(uid entities.UserID) []entities.Article
	AllArticles() []entities.Article
	DeleteArticleById(id entities.ArticleID, uid entities.UserID)
}
