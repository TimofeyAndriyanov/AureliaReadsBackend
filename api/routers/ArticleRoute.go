package routers

import (
	"AureliaReadsBackend/api/dto"
	"AureliaReadsBackend/api/middleware"
	"AureliaReadsBackend/api/utils"
	"AureliaReadsBackend/domain/entities"
	"AureliaReadsBackend/domain/repository"
	"AureliaReadsBackend/domain/services"
	domainUtils "AureliaReadsBackend/domain/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	repository repository.ArticlesRepository
	service    services.JwtService
}

func NewArticleHandler(
	repository repository.ArticlesRepository,
	service services.JwtService,
) ArticleHandler {
	return ArticleHandler{
		repository: repository,
		service:    service,
	}
}

func (h ArticleHandler) ArticleRouter(r *gin.RouterGroup) {
	r.GET("articles", h.allArticles)
	r.GET(":user_id/articles/:article_id", h.getArticleById)
	r.GET(":user_id/articles", h.getArticleByUid)

	secure := r.Use(middleware.AuthMiddleware(h.service, domainUtils.AccessToken))
	secure.POST("me/new_article", h.newArticle)
	secure.DELETE(":user_id/articles/:article_id", h.deleteArticleById)
}

func (h ArticleHandler) newArticle(c *gin.Context) {
	var newArticle dto.NewArticleDTO

	userID := c.MustGet("user_id").(entities.UserID)

	if err := c.ShouldBindJSON(&newArticle); err != nil {
		return
	}

	h.repository.AddArticle(userID, newArticle.ToDomain())

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h ArticleHandler) getArticleById(c *gin.Context) {
	articleID, err := utils.ParamUUID(c, "article_id")

	if err != nil {
		return
	}

	h.repository.GetArticleById(entities.ArticleID(articleID))
}

func (h ArticleHandler) getArticleByUid(c *gin.Context) {
	userID, err := utils.ParamUUID(c, "user_id")

	if err != nil {
		return
	}

	list := h.repository.GetArticlesByUid(entities.UserID(userID))

	listDto := utils.Map(list, func(t entities.Article) dto.ArticleDTO {
		return dto.ArticleDTO{
			Id:        t.Id,
			UserID:    t.UserID,
			Title:     t.Title,
			Content:   t.Content,
			CreatedAt: t.CreatedAt,
		}
	})

	c.JSON(http.StatusOK, listDto)
}

func (h ArticleHandler) deleteArticleById(c *gin.Context) {
	userID := c.MustGet("user_id").(entities.UserID)
	articleID, err := utils.ParamUUID(c, "article_id")

	if err != nil {
		return
	}

	h.repository.DeleteArticleById(entities.ArticleID(articleID), userID)

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h ArticleHandler) allArticles(c *gin.Context) {
	list := h.repository.AllArticles()

	listDto := utils.Map(list, func(t entities.Article) dto.ArticleDTO {
		return dto.ArticleDTO{
			Id:        t.Id,
			UserID:    t.UserID,
			Title:     t.Title,
			Content:   t.Content,
			CreatedAt: t.CreatedAt,
		}
	})

	c.JSON(http.StatusOK, listDto)

}
