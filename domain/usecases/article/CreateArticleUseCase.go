package article

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"
)

type CreateArticleUseCase struct {
	repo article.Repository
}

func NewCreateArticleUseCase(repo article.Repository) *CreateArticleUseCase {
	return &CreateArticleUseCase{repo: repo}
}

func (uc CreateArticleUseCase) Execute(uid entities.UserID, value entities.NewArticle) {
	uc.repo.AddArticle(uid, value)
}
