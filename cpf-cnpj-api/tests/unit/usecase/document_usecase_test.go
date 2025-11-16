package usecase_test

import (
	"cpf-cnpj-api/internal/entity"
	"cpf-cnpj-api/internal/usecase"
	"cpf-cnpj-api/tests/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupConnectionUsecase() (*usecase.Usecase, *mocks.RepositoryMock) {
	mockRepo := new(mocks.RepositoryMock)
	uc := &usecase.Usecase{Repository: mockRepo}
	return uc, mockRepo
}

func TestUsecase_CreateDocument(t *testing.T) {
	usecase, repo := setupConnectionUsecase()
	doc := &entity.Document{Number: "263.495.000-40"}

	t.Run("success new document", func(t *testing.T) {
		repo.On("CreateDocument", doc).Return(doc, false, nil).Once()

		created, existed, err := usecase.CreateDocument(doc)
		assert.NoError(t, err)
		assert.Equal(t, doc, created)
		assert.False(t, existed)

		repo.AssertExpectations(t)
	})

	t.Run("document already exists", func(t *testing.T) {
		repo.On("CreateDocument", doc).Return(doc, true, nil).Once()

		created, existed, err := usecase.CreateDocument(doc)
		assert.NoError(t, err)
		assert.Equal(t, doc, created)
		assert.True(t, existed)

		repo.AssertExpectations(t)
	})

	t.Run("error creating document", func(t *testing.T) {
		repo.On("CreateDocument", doc).Return((*entity.Document)(nil), false, errors.New("fail")).Once()

		created, existed, err := usecase.CreateDocument(doc)

		assert.Error(t, err)
		assert.Nil(t, created)
		assert.False(t, existed)

		repo.AssertExpectations(t)
	})
}

func TestUsecase_GetDocument(t *testing.T) {
	usecase, repo := setupConnectionUsecase()
	doc := &entity.Document{Number: "1"}

	t.Run("success get document", func(t *testing.T) {
		mockDoc := &entity.Document{Number: "123"}
		repo.On("GetDocument", doc).Return(mockDoc, nil).Once()

		result, err := usecase.GetDocument(doc)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "123", result.Number)

		repo.AssertExpectations(t)
	})

	t.Run("error get document", func(t *testing.T) {
		repo.On("GetDocument", doc).Return((*entity.Document)(nil), errors.New("fail")).Once()

		result, err := usecase.GetDocument(doc)

		assert.Error(t, err)
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})
}
func TestUsecase_UpdateDocument(t *testing.T) {
	usecase, repo := setupConnectionUsecase()
	doc := &entity.Document{Number: "1"}

	t.Run("success update document", func(t *testing.T) {
		repo.On("UpdateDocument", doc, "1").Return(doc, nil).Once()

		result, err := usecase.UpdateDocument(doc, "1")
		assert.NoError(t, err)
		assert.Equal(t, doc, result)

		repo.AssertExpectations(t)
	})

	t.Run("error update document", func(t *testing.T) {
		// Retorno explícito de ponteiro nil para evitar panic
		repo.On("UpdateDocument", doc, "1").Return((*entity.Document)(nil), errors.New("fail")).Once()

		result, err := usecase.UpdateDocument(doc, "1")
		assert.Error(t, err)
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})
}

func TestUsecase_DeleteDocument(t *testing.T) {
	usecase, repo := setupConnectionUsecase()

	t.Run("success delete", func(t *testing.T) {
		doc := &entity.Document{Number: "1"}
		repo.On("DeleteDocument", "1").Return(doc, nil).Once()

		result, err := usecase.DeleteDocument("1")
		assert.NoError(t, err)
		assert.Equal(t, doc, result)

		repo.AssertExpectations(t)
	})

	t.Run("error delete", func(t *testing.T) {
		repo.On("DeleteDocument", "1").Return((*entity.Document)(nil), errors.New("fail")).Once()

		result, err := usecase.DeleteDocument("1")
		assert.Error(t, err)
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})
}

func TestUsecase_ListDocuments(t *testing.T) {
	usecase, repo := setupConnectionUsecase()

	docs := &[]entity.Document{
		{Number: "47072371001", Type: "CPF"},
		{Number: "30922172000183", Type: "CNPJ"},
	}

	numberFilter := "47072371001"
	typeFilter := ""
	sortBy := "created_at"
	order := "asc"

	repo.On("ListDocuments", numberFilter, typeFilter, sortBy, order).Return(docs, nil).Once()

	result, err := usecase.ListDocuments(numberFilter, typeFilter, sortBy, order)

	assert.NoError(t, err)
	assert.Equal(t, docs, result)

	repo.AssertExpectations(t)
}

func TestUsecase_MarkDocumentBlocklist(t *testing.T) {
	usecase, repo := setupConnectionUsecase()

	id := "123"
	blocked := true

	repo.On("UpdateDocumentBlocklist", id, blocked).Return(nil).Once()

	err := usecase.MarkDocumentBlocklist(id, blocked)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}
