package entity

import (
	"errors"
	"regexp"
	"strconv"
)

type CPF struct {
	Number string `json:"number"`
}

func NewCPF(number string) (*CPF, error) {
	clean := cleanCPF(number)

	if !IsValidCPF(clean) {
		return nil, errors.New("invalid CPF")
	}

	return &CPF{Number: clean}, nil
}

func cleanCPF(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(s, "")
}

func IsValidCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	invalidSeq := []string{
		"00000000000", "11111111111", "22222222222", "33333333333",
		"44444444444", "55555555555", "66666666666", "77777777777",
		"88888888888", "99999999999",
	}

	for _, seq := range invalidSeq {
		if cpf == seq {
			return false
		}
	}

	if !validateCPFCheckDigit(cpf, 9) {
		return false
	}

	if !validateCPFCheckDigit(cpf, 10) {
		return false
	}

	return true
}

func validateCPFCheckDigit(cpf string, length int) bool {
	sum := 0
	weight := length + 1

	for i := 0; i < length; i++ {
		n, _ := strconv.Atoi(string(cpf[i]))
		sum += n * weight
		weight--
	}

	rest := sum % 11
	if rest < 2 {
		rest = 0
	} else {
		rest = 11 - rest
	}

	digit, _ := strconv.Atoi(string(cpf[length]))
	return rest == digit
}

var CleanCPFForTest = cleanCPF
var ValidateCPFCheckDigitForTest = validateCPFCheckDigit
