package handler_test

import (
	"cpf-cnpj-api/internal/handler"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_MissingHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	defer os.Unsetenv("API_TOKEN")

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	middleware := handler.AuthMiddleware()
	middleware(ctx)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization header is required")
}

func TestAuthMiddleware_InvalidHeaderFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	defer os.Unsetenv("API_TOKEN")

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	middleware := handler.AuthMiddleware()
	middleware(ctx)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid Authorization header")
}

func TestAuthMiddleware_WrongToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("API_TOKEN", "expected-token")
	defer os.Unsetenv("API_TOKEN")

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer wrong-token")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	middleware := handler.AuthMiddleware()
	middleware(ctx)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid token")
}

func TestAuthMiddleware_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("API_TOKEN", "valid-token")
	defer os.Unsetenv("API_TOKEN")

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	calledNext := false
	nextHandler := func(ctx *gin.Context) {
		calledNext = true
	}

	handler.AuthMiddleware()(ctx)

	if !ctx.IsAborted() {
		nextHandler(ctx)
	}

	assert.True(t, calledNext)
	assert.Equal(t, 200, w.Code)
}
