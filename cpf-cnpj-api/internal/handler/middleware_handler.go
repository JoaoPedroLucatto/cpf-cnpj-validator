package handler

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) AddCorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)

			return
		}

		ctx.Next()
	}
}

func (server *Server) RequestCounterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		server.Usecase.Repository.IncrementRequest()

		c.Next()
	}
}
