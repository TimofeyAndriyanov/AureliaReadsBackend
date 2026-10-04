package article

import (
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"
	"github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"
)

type CreateArticleUseCase struct {
	repository article.Repository
}

func NewCreateArticleUseCase() *CreateArticleUseCase {
	return &CreateArticleUseCase{}
}

func (uc CreateArticleUseCase) Execute(uid entities.UserID, value entities.NewArticle) {

}
