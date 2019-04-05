package helpers

// Mod11 generates a validation number
// using the Mod11 method
func Mod11(input int) int {
	mod := input % 11

	if mod >= 2 {
		return 11 - mod
	}

	return 0
}
