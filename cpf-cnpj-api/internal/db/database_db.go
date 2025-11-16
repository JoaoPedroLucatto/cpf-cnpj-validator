package db

import "cpf-cnpj-api/internal/entity"

type Repository interface {
	PingDatabase() error
	CountRequests() (int, error)
	IncrementRequest()

	CreateDocument(*entity.Document) (*entity.Document, bool, error)
	GetDocument(*entity.Document) (*entity.Document, error)
	UpdateDocument(*entity.Document, string) (*entity.Document, error)
	DeleteDocument(string) (*entity.Document, error)
	ListDocuments(string, string, string, string) (*[]entity.Document, error)
	UpdateDocumentBlocklist(string, bool) error
}
