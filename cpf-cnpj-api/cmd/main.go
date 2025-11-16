package main

import (
	"context"
	"cpf-cnpj-api/internal/db/postgres"
	"cpf-cnpj-api/internal/entity"
	"cpf-cnpj-api/internal/handler"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	log := newLog()
	log.Info().Msgf("Running")

	dbUrl := entity.PostgresConnection{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	database, err := postgres.NewPostgres(dbUrl)
	if err != nil {
		log.Fatal().Msgf("error on run server %v", err)

		return
	}

	s := handler.NewServer(context.Background(), log, database)
	if err := s.Server().Run(os.Getenv("HOST")); err != nil {
		log.Fatal().Msgf("error on run server %v", err)
	}
}

func newLog() *zerolog.Logger {
	service := "cpf-cnpj-api"
	log := zerolog.New(os.Stdout).With().
		Timestamp().Str("service", service).Logger()
	logLevel := "debug"

	loggerLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		loggerLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(loggerLevel)

	return &log
}
