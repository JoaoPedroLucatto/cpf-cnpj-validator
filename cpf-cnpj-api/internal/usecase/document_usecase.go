package usecase

import "cpf-cnpj-api/internal/entity"

func (usecase *Usecase) CreateDocument(document *entity.Document) (*entity.Document, bool, error) {
	created, existed, err := usecase.Repository.CreateDocument(document)
	if err != nil {
		return nil, false, err
	}

	return created, existed, nil
}

func (usecase *Usecase) GetDocument(input *entity.Document) (*entity.Document, error) {
	document, err := usecase.Repository.GetDocument(input)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (usecase *Usecase) UpdateDocument(document *entity.Document, id string) (*entity.Document, error) {
	updated, err := usecase.Repository.UpdateDocument(document, id)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (usecase *Usecase) DeleteDocument(id string) (*entity.Document, error) {
	deleted, err := usecase.Repository.DeleteDocument(id)
	if err != nil {
		return nil, err
	}

	return deleted, nil
}

func (usecase *Usecase) ListDocuments(document, doctype, sortBy, order string) (*[]entity.Document, error) {
	list, err := usecase.Repository.ListDocuments(document, doctype, sortBy, order)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (usecase *Usecase) MarkDocumentBlocklist(id string, blocked bool) error {
	return usecase.Repository.UpdateDocumentBlocklist(id, blocked)
}

func (usecase *Usecase) CountRequests() (int, error) {
	count, err := usecase.Repository.CountRequests()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (usecase *Usecase) PingDatabase() error {
	return usecase.Repository.PingDatabase()
}
