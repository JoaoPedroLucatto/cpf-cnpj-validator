package entity

import "errors"

var (
	ErrCannotDocument         = errors.New("document number cannot be empty")
	ErrUnrecognizableDocument = errors.New("document number not recognized as CPF or CNPJ")
	ErrInvalidCPF             = errors.New("invalid CPF number")
	ErrInvalidCNPJ            = errors.New("invalid CNPJ number")
)
