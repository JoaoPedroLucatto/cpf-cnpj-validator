package entity_test

import (
	"cpf-cnpj-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCPF(t *testing.T) {
	t.Run("valid CPF", func(t *testing.T) {
		cpf, err := entity.NewCPF("512.344.908-62")
		assert.NoError(t, err)
		assert.Equal(t, "51234490862", cpf.Number)
	})

	t.Run("invalid CPF number", func(t *testing.T) {
		cpf, err := entity.NewCPF("512.344.908-63") // último dígito inválido
		assert.Error(t, err)
		assert.Nil(t, cpf)
	})

	t.Run("CPF with all same digits", func(t *testing.T) {
		cpf, err := entity.NewCPF("111.111.111-11")
		assert.Error(t, err)
		assert.Nil(t, cpf)
	})

	t.Run("CPF too short", func(t *testing.T) {
		cpf, err := entity.NewCPF("123.456.789-0")
		assert.Error(t, err)
		assert.Nil(t, cpf)
	})

	t.Run("CPF with special chars cleaned", func(t *testing.T) {
		cpf, err := entity.NewCPF("08951653099")
		assert.NoError(t, err)
		assert.Equal(t, "08951653099", cpf.Number)
	})
}

// /func TestValidateCPFCheckDigit(t *testing.T) {
// /	assert.True(t, entity.ValidateCPFCheckDigitForTest("08951653099", 9))
// /	assert.True(t, entity.ValidateCPFCheckDigitForTest("08951653099", 10))
// /
// /	assert.False(t, entity.ValidateCPFCheckDigitForTest("52998224726", 9))
// /	assert.False(t, entity.ValidateCPFCheckDigitForTest("52998224726", 10))
// /}

func TestCleanCPF(t *testing.T) {
	assert.Equal(t, "08951653099", entity.CleanCPFForTest("089.516.530-99"))
	assert.Equal(t, "12345678900", entity.CleanCPFForTest("123-456.789/00"))
}
