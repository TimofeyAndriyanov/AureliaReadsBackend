package api

import (
	"AureliaReadsBackend/api/routers"
	"AureliaReadsBackend/domain/repository"
	"AureliaReadsBackend/domain/services"
	"AureliaReadsBackend/domain/usecases"

	"github.com/gin-gonic/gin"
)

func MainRoute(
	signUpUseCase usecases.SignUpUseCase,
	signInUseCase usecases.SignInUseCase,
	repository repository.ArticlesRepository,
	service services.JwtService,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("api")

	authHandler := routers.NewAuthHandler(
		signInUseCase,
		signUpUseCase,
	)
	authHandler.AuthRouter(api)

	articlesHandler := routers.NewArticleHandler(
		repository,
		service,
	)
	articlesHandler.ArticleRouter(api)

	return r
}
