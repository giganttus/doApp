package middlewares

import (
	"doApp/auth"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey = "Authorization"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get(authHeaderKey)
		if token == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := auth.ValidateJwt(token, os.Getenv("ACCESS_KEY"))
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID := int(claims["userID"].(float64))

		ctx.Set("userID", userID)
		ctx.Next()
	}
}
