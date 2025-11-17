package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Uptime        string `json:"uptime"`
	RequestsCount int    `json:"requestsCount"`
}

func Health(server *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		server.Log.Info().Msg("handler health alive")
		ctx.Status(http.StatusOK)
	})
}

func Metrics(server *Server, startTime time.Time) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		server.Log.Info().Msg("handler get metrics")

		uptime := time.Since(startTime).String()
		requests, _ := server.Usecase.CountRequests()

		ctx.JSON(http.StatusOK, gin.H{"status": StatusResponse{
			Uptime:        uptime,
			RequestsCount: requests,
		}})

	})
}

func Ready(server *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		server.Log.Info().Msg("handler health ready")

		err := server.Usecase.PingDatabase()
		if err != nil {
			ctx.Status(http.StatusInternalServerError)

			return
		}

		ctx.Status(http.StatusOK)
	})
}
