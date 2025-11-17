package handler

import (
	"context"
	"cpf-cnpj-api/internal/db"
	"cpf-cnpj-api/internal/usecase"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type API struct {
	Router *gin.Engine
}

type Server struct {
	Usecase   *usecase.Usecase
	Log       *zerolog.Logger
	StartTime time.Time
}

func NewServer(ctx context.Context, log *zerolog.Logger, repository db.Repository) *Server {
	return &Server{
		Usecase:   usecase.NewUsecaseService(context.Background(), log, repository),
		Log:       log,
		StartTime: time.Now(),
	}
}

func (server *Server) Server() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	gin.DefaultWriter = io.Discard

	router := gin.Default()

	router.Use(server.AddCorsMiddleware())
	router.Use(server.RequestCounterMiddleware())

	router.GET("/", Home)
	router.Static("/docs", "./docs")

	statusPaths := router.Group("status")
	{
		statusPaths.GET("/health", Health(server))
		statusPaths.GET("/metrics", Metrics(server, server.StartTime))
		statusPaths.GET("/ready", Ready(server))
	}

	router.Use(AuthMiddleware())

	documentsPath := router.Group("/documents")
	{
		documentsPath.POST("", PostDocument(server))
		documentsPath.GET("", GetDocuments(server))
		documentsPath.GET("/:document", GetDocument(server))
		documentsPath.PATCH("/:id", PatchDocument(server))
		documentsPath.DELETE("/:id", DeleteDocument(server))
		documentsPath.PATCH("/:id/blocklist", PatchDocumentBlocklist(server))
	}

	return router
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"api": "cpf and cnpf"})
}
