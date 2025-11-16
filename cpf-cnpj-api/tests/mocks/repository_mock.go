package mocks

import (
	"cpf-cnpj-api/internal/entity"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) PingDatabase() error {
	args := r.Called()

	return args.Error(0)
}

func (r *RepositoryMock) CountRequests() (int, error) {
	args := r.Called()
	return args.Int(0), args.Error(1)
}

func (r *RepositoryMock) IncrementRequest() {}

func (r *RepositoryMock) CreateDocument(document *entity.Document) (*entity.Document, bool, error) {
	args := r.Called(document)

	return args.Get(0).(*entity.Document), args.Bool(1), args.Error(2)
}

func (r *RepositoryMock) GetDocument(document *entity.Document) (*entity.Document, error) {
	args := r.Called(document)
	return args.Get(0).(*entity.Document), args.Error(1)
}

func (r *RepositoryMock) UpdateDocument(document *entity.Document, id string) (*entity.Document, error) {
	args := r.Called(document, id)

	return args.Get(0).(*entity.Document), args.Error(1)
}

func (r *RepositoryMock) DeleteDocument(id string) (*entity.Document, error) {
	args := r.Called(id)

	return args.Get(0).(*entity.Document), args.Error(1)
}

func (r *RepositoryMock) ListDocuments(number, docType, sortBy, order string) (*[]entity.Document, error) {
	args := r.Called(number, docType, sortBy, order)

	return args.Get(0).(*[]entity.Document), args.Error(1)
}

func (r *RepositoryMock) UpdateDocumentBlocklist(id string, blocked bool) error {
	args := r.Called(id, blocked)

	return args.Error(0)
}
