package handler

import (
	"cpf-cnpj-api/internal/entity"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentRequest struct {
	Document string `json:"document"`
}

func PostDocument(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.Log.Info().Msg("handler create document")

		var input DocumentRequest

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		documentValid, ok := validateDocument(ctx, server, input.Document)
		if !ok {
			return
		}

		documentCreated, existed, err := server.Usecase.CreateDocument(documentValid)
		if err != nil {
			server.Log.Error().Err(err).Msg("failed to create document")
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to create document",
				"details": err.Error(),
			})

			return
		}

		if existed {
			server.Log.Error().Err(err).Msg("documento existed")
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "documento existed",
			})

			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"document": documentCreated,
		})
	}
}

func GetDocument(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.Log.Info().Msg("handler get document")

		input := ctx.Param("document")

		documentValid, ok := validateDocument(ctx, server, input)
		if !ok {
			return
		}

		document, err := server.Usecase.GetDocument(documentValid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				server.Log.Warn().Msg("document not found")

				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "document not found",
				})
				return
			}

			server.Log.Error().Err(err).Msg("failed to get document")
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to get document",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"document": document,
		})
	}
}

func PatchDocument(server *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		server.Log.Info().Msg("handler update document")

		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing document id"})

			return
		}

		var input DocumentRequest
		if err := ctx.ShouldBindJSON(&input); err != nil {
			server.Log.Error().Err(err).Msg("invalid request body")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})

			return
		}

		documentValid, ok := validateDocument(ctx, server, input.Document)
		if !ok {
			return
		}

		documentUpdated, err := server.Usecase.UpdateDocument(documentValid, id)
		if err != nil {
			server.Log.Error().Err(err).Msg("failed to update document")

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to get document",
				"details": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"document": documentUpdated})
	})
}

func DeleteDocument(server *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		server.Log.Info().Msg("handler update document")

		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing document id"})

			return
		}

		documentDeleted, err := server.Usecase.DeleteDocument(id)
		if err != nil {
			server.Log.Error().Err(err).Msg("failed to delete document")

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to delete document",
				"details": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"document": documentDeleted})
	})
}

func GetDocuments(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.Log.Info().Msg("handler get document")

		document := ctx.Query("document")
		docType := ctx.Query("type")
		sortBy := ctx.DefaultQuery("sortBy", "created_at")
		order := ctx.DefaultQuery("order", "asc")

		var documentNumber string

		if document != "" {
			documentClean, _ := validateDocument(ctx, server, document)
			documentNumber = documentClean.Number
		}

		documentList, err := server.Usecase.ListDocuments(documentNumber, docType, sortBy, order)
		if err != nil {
			server.Log.Error().Err(err).Msg("failed to get document")
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to list document",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"documents": documentList})
	}
}

func PatchDocumentBlocklist(server *Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var body struct {
			Blocked bool `json:"blocked"`
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})

			return
		}

		err := server.Usecase.MarkDocumentBlocklist(id, body.Blocked)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{"id": id, "blocked": body.Blocked})
	}
}

func validateDocument(ctx *gin.Context, server *Server, input string) (*entity.Document, bool) {
	document, err := entity.NewDocument(input)
	if err != nil {
		server.Log.Error().Err(err).Msg("invalid document")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid document",
			"details": err.Error(),
		})

		return nil, false
	}

	return document, true
}
