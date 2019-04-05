package cpf

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/brazilian-utils/go-brazilian-utils/helpers"
)

// IsValid returns true if the input
// is a valid CPF with or without mask
func IsValid(cpf string) bool {
	if !regexp.MustCompile(`^\d{3}\.?\d{3}.?\d{3}-?\d{2}$`).MatchString(cpf) {
		return false
	}

	cpfNumbers := helpers.OnlyNumbers(cpf)

	return hasValidLength(cpfNumbers) && !isBlacklisted(cpf) && isValidChecksum(cpfNumbers)
}

// Perform checksum validation
func isValidChecksum(cpf string) bool {
	validity := true

	for _, verifier := range verificationDigits {
		mod := computeMod(strings.Split(cpf[:verifier], ""))

		valid, _ := strconv.Atoi(string(cpf[verifier]))
		validity = validity && (valid == mod)
	}

	return validity
}

// Compute the mod for the current slice of strings
func computeMod(digits []string) (res int) {
	weight := len(digits) + 1

	var mod int
	for _, digitStr := range digits {
		digit, _ := strconv.Atoi(digitStr)
		mod += digit * weight
		weight--
	}

	if mod = mod % 11; mod >= 2 {
		res = 11 - mod
	}

	return
}

// Validates the string length
func hasValidLength(cpf string) bool {
	return len(cpf) == size
}

func isBlacklisted(cpf string) bool {
	return helpers.Contains(blacklist, cpf)
}
