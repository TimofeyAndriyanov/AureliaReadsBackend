package article

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"
)

type DeleteArticleUseCase struct {
	repo article.Repository
}

func NewDeleteArticleUseCase(repo article.Repository) *DeleteArticleUseCase {
	return &DeleteArticleUseCase{repo: repo}
}

func (uc DeleteArticleUseCase) Execute(id entities.ArticleID) {

}
