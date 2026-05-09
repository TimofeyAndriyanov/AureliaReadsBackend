package middleware

import (
	"AureliaReadsBackend/domain/services"
	"AureliaReadsBackend/domain/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(service services.JwtService, tokenType utils.TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing authorization header",
			})
			return
		}

		const bearerPrefix = "Bearer "

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization header",
			})
			return
		}

		token := strings.TrimPrefix(authHeader, bearerPrefix)

		claims, err := service.VerifyToken(token, tokenType)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
