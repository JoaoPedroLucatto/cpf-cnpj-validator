package postgres

import (
	"cpf-cnpj-api/internal/entity"
	"errors"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (db *Postgres) CreateDocument(document *entity.Document) (*entity.Document, bool, error) {
	result := db.DB.
		Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "number"}},
				DoNothing: true,
			},
		).
		Create(&document)

	if result.Error != nil {
		return nil, false, result.Error
	}

	existed := result.RowsAffected == 0
	if existed {

		var existing entity.Document

		err := db.DB.Unscoped().Where("number = ?", document.Number).First(&existing).Error
		if err != nil {
			return nil, true, err
		}

		if existing.DeletedAt.Valid {
			existing.DeletedAt = gorm.DeletedAt{}
			if err := db.DB.Save(&existing).Error; err != nil {
				return nil, true, err
			}
		}

		return &existing, true, nil
	}

	return document, false, nil
}

func (db *Postgres) GetDocument(input *entity.Document) (*entity.Document, error) {
	var document entity.Document

	if result := db.DB.First(&document, "number = ?", input.Number); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	return &document, nil
}

func (db *Postgres) UpdateDocument(input *entity.Document, id string) (*entity.Document, error) {
	if err := db.DB.
		Model(&entity.Document{}).
		Where("id = ?", id).
		Update("number", input.Number).Error; err != nil {

		return nil, err
	}

	var document entity.Document
	if err := db.DB.
		Where("id = ?", id).
		First(&document).Error; err != nil {

		return nil, err
	}

	return &document, nil
}

func (db *Postgres) DeleteDocument(id string) (*entity.Document, error) {
	var doc entity.Document

	if err := db.DB.
		Where("id = ?", id).
		First(&doc).Error; err != nil {
		return nil, err
	}

	if err := db.DB.Delete(&doc).Error; err != nil {
		return nil, err
	}

	return &doc, nil
}

func (db *Postgres) ListDocuments(document, docType, sortBy, order string) (*[]entity.Document, error) {
	var documents []entity.Document
	query := db.DB

	if document != "" {
		query = query.Where("number = ?", document)
	}

	if docType != "" {
		query = query.Where("type = ?", strings.ToUpper(docType))
	}

	validSortColumns := map[string]bool{
		"number":     true,
		"type":       true,
		"created_at": true,
		"updated_at": true,
	}

	if !validSortColumns[sortBy] {
		sortBy = "created_at"
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	query = query.Order(sortBy + " " + order)

	if result := query.Find(&documents); result.Error != nil {
		return nil, result.Error
	}

	return &documents, nil
}

func (db *Postgres) UpdateDocumentBlocklist(id string, blocked bool) error {
	result := db.DB.Model(&entity.Document{}).Where("id = ?", id).Update("blocked", blocked)

	return result.Error
}
