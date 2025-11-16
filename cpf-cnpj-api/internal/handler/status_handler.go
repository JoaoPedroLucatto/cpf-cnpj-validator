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

func Status(server *Server, startTime time.Time) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		server.Log.Info().Msg("handler get status")

		uptime := time.Since(startTime).String()
		requests, _ := server.Usecase.Repository.CountRequests()

		ctx.JSON(http.StatusOK, gin.H{"status": StatusResponse{
			Uptime:        uptime,
			RequestsCount: requests,
		}})

	})
}
