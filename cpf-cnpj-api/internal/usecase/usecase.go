package usecase

import (
	"context"
	"cpf-cnpj-api/internal/db"

	"github.com/rs/zerolog"
)

type Usecase struct {
	Context    context.Context
	Logger     *zerolog.Logger
	Repository db.Repository
}

func NewUsecaseService(ctx context.Context, log *zerolog.Logger,
	repository db.Repository) *Usecase {
	return &Usecase{
		Context:    ctx,
		Logger:     log,
		Repository: repository,
	}
}
