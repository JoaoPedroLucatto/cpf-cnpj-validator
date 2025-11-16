package handler_test

import (
	"context"
	"cpf-cnpj-api/internal/handler"
	"cpf-cnpj-api/internal/usecase"
	"cpf-cnpj-api/tests/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestStatusHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := new(mocks.RepositoryMock)
	repo.On("CountRequests").Return(10, nil)

	logger := zerolog.New(io.Discard)

	uc := &usecase.Usecase{
		Repository: repo,
		Logger:     &logger,
		Context:    context.Background(),
	}

	server := &handler.Server{
		Usecase: uc,
		Log:     &logger,
	}

	start := time.Now().Add(-2 * time.Second)
	handlerFn := handler.Status(server, start)

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	handlerFn(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"requestsCount":10`)
	assert.Contains(t, w.Body.String(), `"uptime"`)

	repo.AssertExpectations(t)
}
