package handler_test

import (
	"bytes"
	"context"
	"cpf-cnpj-api/internal/entity"
	"cpf-cnpj-api/internal/handler"
	"cpf-cnpj-api/internal/usecase"
	"cpf-cnpj-api/tests/mocks"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type NewDocumentFuncType func(string) (*entity.Document, error)

func mockNewDocumentSuccess(input string) (*entity.Document, error) {
	return &entity.Document{Number: input}, nil
}

func mockNewDocumentError(input string) (*entity.Document, error) {
	return nil, errors.New("invalid document")
}

func newTestServer(repo *mocks.RepositoryMock, newDocumentFunc NewDocumentFuncType) *handler.Server {
	logger := zerolog.New(io.Discard)

	uc := &usecase.Usecase{
		Context:    context.TODO(),
		Logger:     &logger,
		Repository: repo,
	}

	return &handler.Server{
		Usecase: uc,
		Log:     &logger,
	}
}

func setup(t *testing.T) {
	gin.SetMode(gin.TestMode)
}

func TestPostDocument_Success(t *testing.T) {
	setup(t)

	repo := new(mocks.RepositoryMock)
	server := newTestServer(repo, mockNewDocumentSuccess)

	body := `{"document":"18627977062"}`
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/documents", bytes.NewBufferString(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	saved := &entity.Document{Number: "18627977062"}
	repo.On("CreateDocument", mock.Anything).Return(saved, false, nil)

	handler.PostDocument(server)(ctx)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), `"document"`)
	repo.AssertExpectations(t)
}

func TestGetDocument(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("document invalid", func(t *testing.T) {
		repo := new(mocks.RepositoryMock)

		server := newTestServer(repo, mockNewDocumentError)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = gin.Params{{Key: "document", Value: "123"}}
		ctx.Request = httptest.NewRequest("GET", "/documents/123", nil)

		handler.GetDocument(server)(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"invalid document"`)

		repo.AssertNotCalled(t, "GetDocument", mock.Anything)
	})

	t.Run("document not found", func(t *testing.T) {
		repo := new(mocks.RepositoryMock)
		server := newTestServer(repo, mockNewDocumentSuccess)

		repo.On("GetDocument", mock.Anything).
			Return((*entity.Document)(nil), gorm.ErrRecordNotFound).
			Once()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = gin.Params{{Key: "document", Value: "18627977062"}}
		ctx.Request = httptest.NewRequest("GET", "/documents/18627977062", nil)

		handler.GetDocument(server)(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), `"document not found"`)

		repo.AssertExpectations(t)
	})
}

func TestPatchDocument(t *testing.T) {
	setup(t)

	t.Run("success", func(t *testing.T) {
		repo := new(mocks.RepositoryMock)
		server := newTestServer(repo, mockNewDocumentSuccess)

		updated := &entity.Document{Number: "18627977062"}

		repo.On("UpdateDocument", mock.Anything, "10").
			Return(updated, nil).Once()

		body := `{"document":"18627977062"}`
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest("PATCH", "/documents/10", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}

		handler.PatchDocument(server)(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.Contains(t, w.Body.String(), `"18627977062"`)

		repo.AssertExpectations(t)
	})

	t.Run("missing id", func(t *testing.T) {
		repo := new(mocks.RepositoryMock)
		server := newTestServer(repo, mockNewDocumentSuccess)

		body := `{"document":"999"}`
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("PATCH", "/documents/", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		handler.PatchDocument(server)(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"missing document id"`)

		repo.AssertExpectations(t)
	})

	t.Run("invalid body", func(t *testing.T) {
		repo := new(mocks.RepositoryMock)
		server := newTestServer(repo, mockNewDocumentSuccess)

		body := `invalid-json`
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest("PATCH", "/documents/10", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}

		handler.PatchDocument(server)(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"invalid request body"`)

		repo.AssertExpectations(t)
	})

	t.Run("invalid document", func(t *testing.T) {
		repo := new(mocks.RepositoryMock)

		server := newTestServer(repo, mockNewDocumentError)

		body := `{"document":"123"}`
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest("PATCH", "/documents/10", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}

		handler.PatchDocument(server)(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"invalid document"`)

		repo.AssertExpectations(t)
	})

	t.Run("update_error", func(t *testing.T) {
		repo := new(mocks.RepositoryMock)
		server := newTestServer(repo, mockNewDocumentSuccess)

		body := `{"document":"26349500040"}`

		repo.On("UpdateDocument", mock.Anything, "10").
			Return((*entity.Document)(nil), errors.New("fail")).Once()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("PATCH", "/documents/10", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "id", Value: "10"}}

		handler.PatchDocument(server)(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"failed to get document"`)

		repo.AssertExpectations(t)
	})
}

func TestDeleteDocument(t *testing.T) {
	setup(t)

	repo := new(mocks.RepositoryMock)
	server := newTestServer(repo, mockNewDocumentSuccess)

	t.Run("missing id", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request = httptest.NewRequest("DELETE", "/documents/", nil)

		handler.DeleteDocument(server)(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"missing document id"`)
		repo.AssertExpectations(t)
	})

	t.Run("delete error", func(t *testing.T) {
		repo.On("DeleteDocument", "15").
			Return((*entity.Document)(nil), errors.New("fail")).Once()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Params = gin.Params{{Key: "id", Value: "15"}}
		ctx.Request = httptest.NewRequest("DELETE", "/documents/15", nil)

		handler.DeleteDocument(server)(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"failed to delete document"`)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		doc := &entity.Document{Number: "111"}

		repo.On("DeleteDocument", "15").
			Return(doc, nil).Once()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{{Key: "id", Value: "15"}}
		ctx.Request = httptest.NewRequest("DELETE", "/documents/15", nil)

		handler.DeleteDocument(server)(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		assert.Contains(t, w.Body.String(), `"111"`)

		repo.AssertExpectations(t)
	})
}

func TestGetDocuments(t *testing.T) {
	setup(t)

	repo := new(mocks.RepositoryMock)
	server := newTestServer(repo, mockNewDocumentSuccess)

	t.Run("success", func(t *testing.T) {
		list := &[]entity.Document{
			{Number: "1"},
			{Number: "2"},
		}

		repo.On("ListDocuments", "", "", "created_at", "asc").
			Return(list, nil).Once()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest("GET", "/documents", nil)

		handler.GetDocuments(server)(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"documents"`)
		repo.AssertExpectations(t)
	})

	t.Run("list error", func(t *testing.T) {
		repo.On("ListDocuments", "", "", "created_at", "asc").
			Return((*[]entity.Document)(nil), errors.New("fail")).Once()

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest("GET", "/documents", nil)

		handler.GetDocuments(server)(ctx)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"failed to list document"`)
		repo.AssertExpectations(t)
	})
}

func TestPatchDocumentBlocklist_Success(t *testing.T) {
	setup(t)

	repo := new(mocks.RepositoryMock)
	server := newTestServer(repo, mockNewDocumentSuccess)

	repo.On("UpdateDocumentBlocklist", "21", true).Return(nil)

	body := `{"blocked":true}`
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{{Key: "id", Value: "21"}}
	ctx.Request = httptest.NewRequest("PATCH", "/documents/21/blocklist", bytes.NewBufferString(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	handler.PatchDocumentBlocklist(server)(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"blocked":true`)
	repo.AssertExpectations(t)
}
