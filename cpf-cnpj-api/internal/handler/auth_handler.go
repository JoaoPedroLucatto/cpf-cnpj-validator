package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[1] == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header"})
			ctx.Abort()
			return
		}

		token := parts[1]
		expectedToken := os.Getenv("API_TOKEN")

		if token != expectedToken {
			ctx.JSON(http.StatusForbidden, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
