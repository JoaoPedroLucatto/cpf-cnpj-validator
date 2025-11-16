package entity_test

import (
	"cpf-cnpj-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDocument(t *testing.T) {
	t.Run("deve criar documento CPF válido", func(t *testing.T) {
		doc, err := entity.NewDocument("529.982.247-25")

		require.NoError(t, err)
		require.NotNil(t, doc)
		assert.Equal(t, "52998224725", doc.Number)
		assert.Equal(t, "CPF", doc.Type)
	})

	t.Run("deve criar documento CNPJ válido", func(t *testing.T) {
		doc, err := entity.NewDocument("11.222.333/0001-81")

		require.NoError(t, err)
		require.NotNil(t, doc)
		assert.Equal(t, "11222333000181", doc.Number)
		assert.Equal(t, "CNPJ", doc.Type)
	})

	t.Run("erro: documento vazio", func(t *testing.T) {
		doc, err := entity.NewDocument("   ")

		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Equal(t, entity.ErrCannotDocument, err)
	})

	t.Run("erro: documento não reconhecido", func(t *testing.T) {
		doc, err := entity.NewDocument("123456")

		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Equal(t, entity.ErrUnrecognizableDocument, err)
	})

	t.Run("erro: CPF inválido", func(t *testing.T) {
		doc, err := entity.NewDocument("11144477710")

		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Equal(t, entity.ErrInvalidCPF, err)
	})

	t.Run("erro: CNPJ inválido", func(t *testing.T) {
		doc, err := entity.NewDocument("11222333000180")

		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Equal(t, entity.ErrInvalidCNPJ, err)
	})
}

func TestClenDocument(t *testing.T) {
	got := entity.ClenDocumentForTest("11.222.333/0001-81")
	assert.Equal(t, "11222333000181", got)
}

func TestIsCPFOrCNPJ(t *testing.T) {
	t.Run("CPF", func(t *testing.T) {
		docType, err := entity.IsCPFOrCNPJForTest("52998224725")
		require.NoError(t, err)
		assert.Equal(t, "CPF", docType)
	})

	t.Run("CNPJ", func(t *testing.T) {
		docType, err := entity.IsCPFOrCNPJForTest("11222333000181")
		require.NoError(t, err)
		assert.Equal(t, "CNPJ", docType)
	})

	t.Run("erro: tamanho inválido", func(t *testing.T) {
		docType, err := entity.IsCPFOrCNPJForTest("1234")

		require.Error(t, err)
		assert.Empty(t, docType)
		assert.Equal(t, entity.ErrUnrecognizableDocument, err)
	})
}

func TestValidDocument(t *testing.T) {
	t.Run("CPF válido", func(t *testing.T) {
		doc, err := entity.ValidDocumentForTest("52998224725", "CPF")
		require.NoError(t, err)
		assert.Equal(t, "52998224725", doc)
	})

	t.Run("CPF inválido", func(t *testing.T) {
		_, err := entity.ValidDocumentForTest("00011122233", "CPF")
		require.Error(t, err)
		assert.Equal(t, entity.ErrInvalidCPF, err)
	})

	t.Run("CNPJ válido", func(t *testing.T) {
		doc, err := entity.ValidDocumentForTest("11222333000181", "CNPJ")
		require.NoError(t, err)
		assert.Equal(t, "11222333000181", doc)
	})

	t.Run("CNPJ inválido", func(t *testing.T) {
		_, err := entity.ValidDocumentForTest("11222333000180", "CNPJ")
		require.Error(t, err)
		assert.Equal(t, entity.ErrInvalidCNPJ, err)
	})
}
