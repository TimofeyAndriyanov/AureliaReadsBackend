package article

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"

type DeleteArticleUseCase struct {
	repository article.Repository
}

func NewDeleteArticleUseCase() *DeleteArticleUseCase {
	return &DeleteArticleUseCase{}
}

func (uc DeleteArticleUseCase) Execute() {

}
