package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Number    string         `json:"number" gorm:"uniqueIndex"`
	Type      string         `json:"type"`
	Blocked   bool           `json:"blocked" gorm:"default:false"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func NewDocument(document string) (*Document, error) {
	if len(strings.TrimSpace(document)) == 0 {
		return nil, ErrCannotDocument
	}

	documentClean := clenDocument(document)

	documentType, err := isCPFOrCNPJ(documentClean)
	if err != nil {
		return nil, err
	}

	documentValid, err := validDocument(documentClean, documentType)
	if err != nil {
		return nil, err
	}

	return &Document{
		Number: documentValid,
		Type:   documentType,
	}, nil

}

func clenDocument(document string) string {
	document = strings.ReplaceAll(document, ".", "")
	document = strings.ReplaceAll(document, "-", "")
	document = strings.ReplaceAll(document, "/", "")

	return document
}

func isCPFOrCNPJ(document string) (string, error) {
	switch len(document) {
	case 11:
		return "CPF", nil
	case 14:
		return "CNPJ", nil
	}

	fmt.Println(len(document))

	return "", ErrUnrecognizableDocument
}

func validDocument(document, documentType string) (string, error) {
	if documentType == "CPF" {
		isValid := IsValidCPF(document)
		if !isValid {
			return document, ErrInvalidCPF
		}

		return document, nil
	}

	isValid := IsValidCNPJ(document)
	if !isValid {
		return document, ErrInvalidCNPJ
	}

	return document, nil
}

var ClenDocumentForTest = clenDocument
var IsCPFOrCNPJForTest = isCPFOrCNPJ
var ValidDocumentForTest = validDocument
