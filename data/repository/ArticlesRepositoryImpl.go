package repository

import (
	utilsData "AureliaReadsBackend/data/utils"
	"AureliaReadsBackend/domain/entities"
	"AureliaReadsBackend/domain/repository"
	"AureliaReadsBackend/domain/utils"
	"slices"
	"time"
)

type articlesRepositoryImpl struct{}

func NewArticlesRepository() repository.ArticlesRepository {
	return &articlesRepositoryImpl{}
}

var articles []entities.Article

func (i articlesRepositoryImpl) AddArticle(uid entities.UserID, article entities.NewArticle) {
	id := utils.NewArticleID()
	currentTime := time.Now()

	newArticle := entities.Article{
		Id:        id,
		UserID:    uid,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: currentTime.String(),
	}
	articles = append(articles, newArticle)
}

func (i articlesRepositoryImpl) GetArticleById(id entities.ArticleID) (*entities.Article, bool) {
	return utilsData.Find(articles, func(a entities.Article) bool {
		return a.Id == id
	})
}

func (i articlesRepositoryImpl) DeleteArticleById(id entities.ArticleID, uid entities.UserID) {
	slices.DeleteFunc(articles, func(a entities.Article) bool {
		return a.Id == id && a.UserID == uid
	})
}

func (i articlesRepositoryImpl) GetArticlesByUid(uid entities.UserID) []entities.Article {
	return utilsData.Filter(articles, func(a entities.Article) bool {
		return a.UserID == uid
	})
}

func (i articlesRepositoryImpl) AllArticles() []entities.Article {
	return articles
}
