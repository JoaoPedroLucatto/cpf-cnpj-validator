package entity_test

import (
	"cpf-cnpj-api/internal/entity"
	"testing"
)

func TestPostgresConnectionInitialization(t *testing.T) {
	conn := entity.PostgresConnection{
		Host:     "localhost",
		User:     "postgres",
		Password: "secret",
		DBName:   "mydb",
		Port:     "5432",
	}

	if conn.Host != "localhost" {
		t.Errorf("expected Host 'localhost', got '%s'", conn.Host)
	}
	if conn.User != "postgres" {
		t.Errorf("expected User 'postgres', got '%s'", conn.User)
	}
	if conn.Password != "secret" {
		t.Errorf("expected Password 'secret', got '%s'", conn.Password)
	}
	if conn.DBName != "mydb" {
		t.Errorf("expected DBName 'mydb', got '%s'", conn.DBName)
	}
	if conn.Port != "5432" {
		t.Errorf("expected Port '5432', got '%s'", conn.Port)
	}
}
