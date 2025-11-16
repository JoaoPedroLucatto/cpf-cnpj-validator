package postgres

import (
	"cpf-cnpj-api/internal/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB           *gorm.DB
	requestCount int64
}

func NewPostgres(dbUrl entity.PostgresConnection) (*Postgres, error) {
	db, err := connect(dbUrl)
	if err != nil {
		return nil, err
	}

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		DB: db,
	}, nil
}

func connect(dbUrl entity.PostgresConnection) (*gorm.DB, error) {
	UrlConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbUrl.Host, dbUrl.User, dbUrl.Password, dbUrl.DBName, dbUrl.Port)

	dbConneted, err := gorm.Open(postgres.Open(UrlConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConneted, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Document{},
	)
}
