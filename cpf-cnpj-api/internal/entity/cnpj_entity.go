package entity

import (
	"errors"
	"strconv"
)

type CNPJ struct {
	Number string `json:"number"`
}

func NewCNPJ(number string) (*CNPJ, error) {
	if !IsValidCNPJ(number) {
		return nil, errors.New("invalid CNPJ")
	}

	return &CNPJ{Number: number}, nil
}

func IsValidCNPJ(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}

	d1 := calculateDigitCNPJ(cnpj[:12], []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	d2 := calculateDigitCNPJ(cnpj[:12]+strconv.Itoa(d1), []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	return cnpj == cnpj[:12]+strconv.Itoa(d1)+strconv.Itoa(d2)
}

func calculateDigitCNPJ(base string, weights []int) int {
	sum := 0

	for i, weight := range weights {
		n, _ := strconv.Atoi(string(base[i]))
		sum += n * weight
	}

	rest := sum % 11
	if rest < 2 {
		return 0
	}

	return 11 - rest
}
