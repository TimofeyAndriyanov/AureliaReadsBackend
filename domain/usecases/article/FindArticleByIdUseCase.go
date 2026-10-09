package article

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"

type FindArticleByIdUseCase struct {
	repo article.Repository
}

func NewFindArticleByIdUseCase(repo article.Repository) *FindArticleByIdUseCase {
	return &FindArticleByIdUseCase{repo: repo}
}

func (uc FindArticleByIdUseCase) Execute() {

}
