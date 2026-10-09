package article

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/repository/article"

type GetAllArticlesUseCase struct {
	repo article.Repository
}

func NewGetAllArticlesUseCase(repo article.Repository) *GetAllArticlesUseCase {
	return &GetAllArticlesUseCase{repo: repo}
}

func (uc GetAllArticlesUseCase) Execute() {

}
