package entity_test

import (
	"cpf-cnpj-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCNPJ(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
	}{
		{"válido", "11222333000181", false},
		{"tamanho inválido", "123", true},
		{"dígitos inválidos", "11222333000182", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cnpj, err := entity.NewCNPJ(tt.input)

			if tt.wantError {
				require.Error(t, err)
				assert.Nil(t, cnpj)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, cnpj)
			assert.Equal(t, tt.input, cnpj.Number)
		})
	}
}

func TestIsValidCNPJ(t *testing.T) {
	valid := "11222333000181"

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"válido", valid, true},
		{"tamanho errado", "123", false},
		{"inválido", "11222333000182", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := entity.IsValidCNPJ(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculateDigitsViaIsValid(t *testing.T) {
	assert.True(t, entity.IsValidCNPJ("11222333000181"))
	assert.False(t, entity.IsValidCNPJ("11222333000180"))
}
